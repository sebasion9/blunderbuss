package semantics

import (
	"blunderbuss/cmd/codegen"
	"blunderbuss/cmd/parsing"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

var QWORD = 8

type Visitor struct {
	*parsing.BaseBlunderbussVisitor
	codegen.Codegen
	cctx CompilerContext
	Errors []error
}
func (v *Visitor) Visit(tree antlr.ParseTree) any {
	switch ctx := tree.(type) {
	case *parsing.ProgramContext:
		return v.VisitProgram(ctx)
	case *parsing.FuncContext:
		return v.VisitFunc(ctx)
	case *parsing.ArgsContext:
		return v.VisitArgs(ctx)
	case *parsing.StmtContext:
		return v.VisitStmt(ctx)
	case *parsing.ExprContext:
		return v.VisitExpr(ctx)
	case *parsing.ParamContext:
		return v.VisitParam(ctx)
	case *parsing.BlockContext:
		return v.VisitBlock(ctx)
	case *parsing.For_stmtContext:
		return v.VisitFor_stmt(ctx)
	case *parsing.If_stmtContext:
		return v.VisitIf_stmt(ctx)
	case *parsing.Func_callContext:
		return v.VisitFunc_call(ctx)
	case *parsing.Call_argsContext:
		return v.VisitCall_args(ctx)
	default:
		return ctx.Accept(v)
	}
}

func NewBlunderbussVisitor() *Visitor {
	return &Visitor {
		BaseBlunderbussVisitor: &parsing.BaseBlunderbussVisitor{},
		Codegen: codegen.Codegen{},
		cctx: NewCompilerContext(),
	}
}

func (v *Visitor) VisitProgram(ctx *parsing.ProgramContext) any {
	v.Codegen.GenInit()
	parent := v.cctx.GetCurrScope()

	scope := NewScopeTree("program_registers", parent, OTHER)
	InitRegisters(scope.GetVars())


	scope = parent

	for _, decl := range ctx.AllExtern() {
		if decl.Func_() != nil {
			name := decl.Func_().ID().GetText()
			type_ := decl.Func_().TYPE().GetText()
			args := v.Visit(decl.Func_().Args()).([]ScopeFuncArg)

			var mappedArgs []ScopeFuncArg
			for i, arg := range args {
				sfa := *NewScopeFuncArg(arg.expr.(string), arg.type_, i)
				mappedArgs = append(mappedArgs, sfa)
			}
			(*scope.GetVars())[name] = NewScopeFunc(name, mappedArgs, TypeEnumFromStr(type_), false)
			v.GenExtern(name)
		}
	}



	for _, f := range ctx.AllFunc_() {
		v.Visit(f)
	}


	asm := v.Codegen.StreamAsm()
	return asm
}

func (v *Visitor) VisitFunc(ctx *parsing.FuncContext) any {
	funcName := ctx.ID().GetText() 
	funcType := ctx.TYPE().GetText()
	type_ := TypeEnumFromStr(funcType)
	isCache := ctx.CACHE() != nil
	endFn := EndFnLabel(funcName)
	endFnCached := EndFnLabel(fmt.Sprintf("%s__CACHED", funcName))
	startFn := fmt.Sprintf("START__FN__%s", funcName)



	var args []ScopeFuncArg
	v.Codegen.GenFuncInit(funcName)
	text := v.GetText()
	frameId := len(*text) - 1

	parent := v.cctx.GetCurrScope()
	scope := NewScopeTree(funcName, parent, FUNC)
	v.cctx.currentScope = scope



	args = v.Visit(ctx.Args()).([]ScopeFuncArg)
	var stackArgs []ScopeVar
	for i, a := range args {
		off := scope.GetOff()
		scope.IncrOff(QWORD)
		sv := NewScopeVar(a.expr.(string), a.type_, off)
		(*scope.GetVars())[a.expr.(string)] = sv
		stackArgs = append(stackArgs, *sv)
		v.GenMovAddrRelative(off, v.GetCallArg(i))
	}


	if isCache && funcName != "main" {
		CheckSafe(funcName,&v.Codegen)
		PrepKey(funcName, type_,stackArgs, &v.Codegen, scope)
		CallGetm(&v.Codegen, endFnCached, startFn)
	}

	fn := NewScopeFunc(funcName, args, type_, isCache)
	(*parent.GetVars())[funcName] = fn

	v.GenLabel(startFn)

	v.Visit(ctx.Block())
	v.GenLabel(endFn)

	if isCache && funcName != "main" {
		PrepKey(funcName, type_, stackArgs, &v.Codegen, scope)
		CallSetm(&v.Codegen)
	}

	v.GenLabel(endFnCached)


	// align stack to 16 bytes
	offset := scope.GetOff()
	scaled := (offset - QWORD)
	diff := scaled % (2*QWORD)
	align := scaled + diff

	v.GenAlignStack((align))

	(*text)[frameId].SetSrc(strconv.Itoa(align))

	v.Codegen.GenFuncExit()


	v.cctx.currentScope = parent

	(*parent.GetVars())[funcName] = fn
	return fn
}

func (v *Visitor) VisitArgs(ctx *parsing.ArgsContext) any {
	var args []ScopeFuncArg
	for i, p := range ctx.AllParam() {
		name := p.ID().GetText()
		typeText := p.TYPE().GetText()
		type_ := ANY_
		switch typeText {
		case "int":
			type_ = INT_
		case "str":
			type_ = STR_
		default:
			type_ = ANY_
		}
		args = append(args, *NewScopeFuncArg(name, type_, i))
	}
	return args
}

func (v *Visitor) VisitBlock(ctx *parsing.BlockContext) any {
	for _, s := range ctx.AllStmt() {
		v.Visit(s)
	}
	return nil
}

func (v *Visitor) VisitStmt(ctx *parsing.StmtContext) any {
	parent := GetParentFunc(ctx)
	if parent == nil {
		return nil 
	}
	currScope := v.cctx.GetCurrScope()
	scopeTree := GetWhatFunc(currScope)
	root := GetRootScope(scopeTree)
	registers := *root.GetScopeByName("program_registers").GetVars()



	// assign to scope, initialize
	if ctx.ASSIGN() != nil {
		if ctx.LBRACKET() != nil && ctx.RBRACKET() != nil {
			lhsName := ctx.ID().GetText()
			scope := SearchIdUp(currScope, lhsName)
			lhs := (*scope.GetVars())[lhsName].(*ScopeVar)
			if lhs.Type() != PTR_ {
				v.AddError(NewIndexError(StrFromTypeEnum(lhs.Type()), lhsName, ctx))
				return nil
			}

			off := v.Visit(ctx.Expr(0)).(*ScopeVar).offset
			rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
			loff := lhs.offset
			roff := rhs.offset

			// rbx = index
			v.GenMovRegRelative("rbx", off)
			// rcx = index * qword
			v.GenMovOffset("rcx", "rbx", QWORD)
			// rdx = array
			v.GenMovRegRelative("rdx", loff)
			// rdx = array[index]
			v.GenAddMemory("rdx", "rcx")
			// rax = rhs expr
			if rhs.Raw() == nil {
				rax, _ := registers["rax"].(*Register)
				lhs.expr = rax.Raw()
			} else {
				v.Codegen.GenMovRegRelative("rax", roff)
				lhs.expr = rhs.expr
			}
			// array[index] = rhs expr
			v.GenMovMemory("[rdx]", "rax")
			(*scope.GetVars())[lhsName] = lhs

			return nil
		}
		// check if lhs in scope
		lhsName := ctx.ID().GetText()

		// look for var in scopes
		varScope := SearchIdUp(currScope, lhsName)
		var lhs *ScopeVar


		if varScope != nil {
			// assign existing variable
			existingVar := (*varScope.GetVars())[lhsName]
			if existingVar != nil && ctx.TYPE() == nil {
				lhs = existingVar.(*ScopeVar)
			// variable reinitialization
			} else if existingVar != nil && ctx.TYPE() != nil {
				v.AddError(NewInitializedVariableError(ctx.TYPE().GetText(), lhsName, ctx))
				return nil
			}
		} else {
			// undefined variable
			if ctx.TYPE() == nil {
				v.AddError(NewUndefinedVariableError(lhsName, ctx))
				return nil
			// initialize new variable
			} else {
				varScope = v.cctx.GetCurrScope()
				offset := scopeTree.GetOff()
				scopeTree.IncrOff(QWORD)
				lhs = NewScopeVar(0, TypeEnumFromStr(ctx.TYPE().GetText()), offset)
			}

		}

		rhs := *v.Visit(ctx.Expr(0)).(*ScopeVar)

		if lhs.Type() != rhs.Type() {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}

		loff := lhs.offset
		roff := rhs.offset

		if rhs.Raw() == nil {
			v.Codegen.GenMovAddrRelative(loff, "rax")
			rax, _ := registers["rax"].(*Register)
			lhs.expr = rax.Raw()
			(*varScope.GetVars())[lhsName] = lhs
		} else {
			v.Codegen.GenMovRegRelative("rax", roff)
			v.Codegen.GenMovAddrRelative(loff, "rax")
			lhs.expr = rhs.expr
			(*varScope.GetVars())[lhsName] = lhs
		}
		return nil

	}

	if ctx.TYPE() != nil && ctx.ID() != nil {
		switch ctx.TYPE().GetText() {
		case INT:
			offset := scopeTree.GetOff()
			scopeTree.IncrOff(QWORD)
			(*currScope.GetVars())[ctx.ID().GetText()] = NewScopeVar(0, INT_, offset)
		case STR:
			offset := scopeTree.GetOff()
			scopeTree.IncrOff(QWORD)
			(*currScope.GetVars())[ctx.ID().GetText()] = NewScopeVar("", STR_, offset)
		case PTR:
			offset := scopeTree.GetOff()
			scopeTree.IncrOff(QWORD)
			(*currScope.GetVars())[ctx.ID().GetText()] = NewScopeVar(0, PTR_, offset)
			// mov [rbp - offset], 0
			off := v.WrapEff(fmt.Sprintf("rbp - %d", offset))
			zero := "0"
			v.AddText(codegen.NewInstr("mov", &off, &zero))
		}
		return nil

	}

	if ctx.Func_call() != nil {
		v.Visit(ctx.Func_call())
		return nil
	}

	if ctx.For_stmt() != nil {
		v.Visit(ctx.For_stmt())
		return nil

	}

	if ctx.If_stmt() != nil {
		v.Visit(ctx.If_stmt())
		return nil
	}

	if ctx.BREAK() != nil {
		forScope := GetWhatFor(currScope)
		name := forScope.name
		label := EndBlockLabel(name)
		v.GenJmp(label)
		return nil
	}

	if ctx.NEXT() != nil {
		forScope := GetWhatFor(currScope)
		name := forScope.name
		label := EndStmtLabel(name)
		v.GenJmp(label)
		return nil
	}

	if ctx.RETURN() != nil {
		expr := v.Visit(ctx.Expr(0)).(*ScopeVar)
		registers["rax"].(*Register).Write(expr.Type())
		if parent.TYPE().GetText() != StrFromTypeEnum(expr.Type()) {
			v.AddError(NewTypeMismatchError(parent.TYPE().GetText(), StrFromTypeEnum(expr.Type()),ctx))
			return nil
		}

		endFn := EndFnLabel(parent.ID().GetText())

		v.GenMovRegRelative("r12", expr.Offset())
		v.GenJmp(endFn)
		return nil
	}

	return nil
}

func (v *Visitor) VisitExpr(ctx *parsing.ExprContext) any {
	scopeTree := v.cctx.GetCurrScope()
	funcTree := GetWhatFunc(scopeTree)
	scope := *scopeTree.GetVars()
	if ctx.NUM() != nil {
		id := v.CreateId()
		val, _ := strconv.Atoi(ctx.NUM().GetText())
		text := fmt.Sprintf("qword %d", val)
		v.Codegen.GenIntPrimitive(id, val)
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		v.GenMovAddrRelative(offset, text)
		svar := NewScopeVar(val, INT_, offset)
		scope[id] = svar
 
		return svar
	}
	if ctx.STRING() != nil {
		id := v.CreateId()
		val := ctx.STRING().GetText()
		v.Codegen.GenStrPrimitive(id, val)
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		v.GenMovMemory("rax", id)
		v.GenMovAddrRelative(offset, "rax")
		svar := NewScopeVar(id, STR_, offset)
		scope[id] = svar

		return svar
	}

	if ctx.Func_call() != nil {
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		svarfc := v.Visit(ctx.Func_call())
		if svarfc == nil {
			v.AddError(NewUndefinedVariableError(ctx.Func_call().GetText(), ctx))
			return NewScopeVar(nil, ANY_, 0)
		}
		svar := svarfc.(*ScopeVar)

		svar.offset = offset
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = svar
		return svar
	}

	if ctx.ID() != nil {
		id := ctx.ID().GetText()

		// look for var in scopes
		varScope := SearchIdUp(scopeTree, id)
		// couldnt find scope with that lhs name
		if varScope == nil {
			v.AddError(NewUndefinedVariableError(id, ctx))
			return NewScopeVar(nil, ANY_, 0)
		}
		expr := (*varScope.GetVars())[id]

		return expr
	}

	if ctx.LBRACKET() != nil && ctx.RBRACKET() != nil {
		ptr := v.Visit(ctx.Expr(0)).(*ScopeVar)
		idx := v.Visit(ctx.Expr(1)).(*ScopeVar)

		ptroff := ptr.offset
		idxoff := idx.offset

		// address
		v.GenMovRegRelative("rax", ptroff)
		// index
		v.GenMovRegRelative("rbx", idxoff)
		// index * QWORD
		v.GenMovOffset("rcx", "rbx", QWORD)
		// [rax + rcx] = array[index]
		// rbx = array[index] (value)
		v.GenMovMemory("rbx", "[rax + rcx]")

		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)

		v.GenMovAddrRelative(offset, "rbx")
		prod := NewScopeVar(0, INT_, offset)
		scope[v.CreateId()] = prod
		return prod

	}

	if ctx.LPAREN() != nil && ctx.RPAREN() != nil {
		return v.Visit(ctx.Expr(0))
	}

	// operators
	if ctx.PLUS() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		// offset := len(scope)*QWORD+QWORD
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rsi", rhs.offset)
		v.GenAddMemory("rax", "rsi")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod
	}

	if ctx.MINUS() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rsi", rhs.offset)
		v.GenSubMemory("rax", "rsi")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod
	}

	

	if ctx.MULT() != nil && len(ctx.AllExpr()) == 2 {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenImul(rhs.offset)
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod
	}

	if ctx.DIV() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		prod := NewScopeVar(0, INT_, offset)

		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenXor("rdx", "rdx")
		v.GenDiv()
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod
	}

	if ctx.MOD() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		prod := NewScopeVar(0, INT_, offset)

		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenXor("rdx", "rdx")
		v.GenDiv()
		v.GenMovAddrRelative(offset, "rdx")
		scope[v.CreateId()] = prod
		return prod
	}

	if ctx.EQUAL() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rbx", lhs.offset)
		v.GenMovRegRelative("rcx", rhs.offset)
		v.GenXor("rax", "rax")
		v.GenCmp("rbx", "rcx")
		v.GenSete("al")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod
	}

	if ctx.NOT_EQUAL() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rbx", lhs.offset)
		v.GenMovRegRelative("rcx", rhs.offset)
		v.GenXor("rax", "rax")
		v.GenCmp("rbx", "rcx")
		v.GenSetne("al")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod
	}

	if ctx.LE() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rcx", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenXor("rax", "rax")
		v.GenCmp("rbx", "rcx")
		v.GenSetle("al")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod

	}

	if ctx.LT() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rcx", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenXor("rax", "rax")
		v.GenCmp("rbx", "rcx")
		v.GenSetl("al")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod

	}

	if ctx.GE() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rcx", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenXor("rax", "rax")
		v.GenCmp("rbx", "rcx")
		v.GenSetge("al")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod
	}

	if ctx.GT() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			v.AddError(NewTypeMismatchError(StrFromTypeEnum(lhs.Type()), StrFromTypeEnum(rhs.Type()),ctx))
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rcx", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenXor("rax", "rax")
		v.GenCmp("rbx", "rcx")
		v.GenSetg("al")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod

	}


	// reference
	if ctx.AMPS() != nil {
		if ctx.Expr(0).ID() != nil {
			fnName := ctx.Expr(0).ID().GetText()
			foundScope := SearchIdUp(scopeTree, fnName)
			if fn, ok := (*foundScope.GetVars())[fnName].(*ScopeFunc); ok {
				offset := funcTree.GetOff()
				funcTree.IncrOff(QWORD)
				prod := NewScopeVar(0, PTR_, offset)
				rax := "rax"
				exprOff := fn.Raw().(string)
				newOff := v.WrapEff(fmt.Sprintf("rbp - %d", offset))
				v.AddText(codegen.NewInstr("lea", &rax, &exprOff))
				v.AddText(codegen.NewInstr("mov", &newOff, &rax))
				scope[v.CreateId()] = prod
				return prod
			}
		}

		sv := v.Visit(ctx.Expr(0)).(*ScopeVar)

		offset := funcTree.GetOff()
		funcTree.IncrOff(QWORD)
		prod := NewScopeVar(0, PTR_, offset)

		rax := "rax"
		exprOff := v.WrapEff(sv.offset)
		newOff := v.WrapEff(offset)
		v.AddText(codegen.NewInstr("lea", &rax, &exprOff))
		v.AddText(codegen.NewInstr("mov", &newOff, &rax))

		scope[v.CreateId()] = prod
		return prod
	}

	return nil
}

func (v *Visitor) VisitFunc_call(ctx *parsing.Func_callContext) any {
	scope := v.cctx.GetCurrScope()
	root := GetRootScope(scope)
	registers := *root.GetScopeByName("program_registers").GetVars()

	program := *root.GetVars()
	fnName := ctx.ID().GetText()
	fn, ok := program[fnName].(*ScopeFunc)
	if ok {
		fnScopeTree := NewScopeTree(fnName, scope, FNCALL)
		fnScope := *fnScopeTree.GetVars()
		v.cctx.currentScope = fnScopeTree
		
		if len(ctx.Call_args().AllExpr()) != len(fn.args) {
			v.AddError(NewFnMismatchError(
				fnName,
				strconv.Itoa(len(ctx.Call_args().AllExpr())),
				strconv.Itoa(len(fn.args)),
				"count",
				ctx))
			return nil
		}


		pushOff := []int{}
		pushIdx := []int{}
		for i, arg := range ctx.Call_args().AllExpr() {
			expr := v.Visit(arg).(*ScopeVar)
			if fn.args[i].type_ != ANY_ && (fn.args[i].type_ != expr.type_) {
				v.AddError(NewFnMismatchError(
				fnName,
				StrFromTypeEnum(fn.args[i].type_),
				StrFromTypeEnum(expr.type_),
				"type",
				ctx))
				return nil
			}
			fnScope[fn.args[i].expr.(string)] = expr
			pushOff = append(pushOff, expr.Offset())
			pushIdx = append(pushIdx, i)
		}
		for i := range pushOff {
			v.GenPushArg(pushOff[i], pushIdx[i])
		}

		arg := "0"
		if ctx.SAFE() != nil {
			arg = "1"
		}
		r10 := "r10"
		v.AddText(codegen.NewInstr("mov", &r10, &arg))

		v.GenXor("rax", "rax")
		v.GenCallFunc(fnName)
		v.GenShrinkStackFrame(len(fn.args) - 6)

		rax := registers["rax"].(*Register)
		rax.Write(fn.type_)


		v.cctx.currentScope = scope
		return NewScopeVar(nil, fn.Type(), 0)
	}

	return nil
}

func (v *Visitor) VisitIf_stmt(ctx *parsing.If_stmtContext) any {
	scopeName := BlockScopeName("if")

	parent := v.cctx.GetCurrScope()
	scope := NewScopeTree(scopeName, parent, IF)
	v.cctx.currentScope = scope

	ifLabel := fmt.Sprintf("IF__%s", scopeName)
	elseIfLabel := fmt.Sprintf("ELSEIF__%s", scopeName)
	elseLabel := fmt.Sprintf("ELSE__%s", scopeName)
	endIf := fmt.Sprintf("END__%s", scopeName)
	last := len(ctx.AllBlock()) - 1

	ifCondition := v.Visit(ctx.Expr(0)).(*ScopeVar)
	v.GenCmpAddrReg(ifCondition.offset, "0")
	v.GenJz(elseIfLabel)
	v.GenJmp(ifLabel)

	elseIfLen := len(ctx.AllELSEIF())
	v.GenLabel(elseIfLabel)
	for i := range ctx.AllELSEIF() {
		endElseIfLabel := fmt.Sprintf("END__%s__%d", elseIfLabel, i)
		condition := v.Visit(ctx.Expr(i+1)).(*ScopeVar)


		v.GenCmpAddrReg(condition.offset, "0")
		if(i == elseIfLen - 1) {
			v.GenJz(elseLabel)
		} else {
			v.GenJz(endElseIfLabel)
		}

		v.Visit(ctx.Block(i+1))

		v.GenCmpAddrReg(condition.offset, "1")
		v.GenJz(endIf)

		v.GenLabel(endElseIfLabel)
	}




	v.GenLabel(elseLabel)
	if ctx.ELSE() != nil {
		v.Visit(ctx.Block(last))
	}
	v.GenJmp(endIf)

	v.GenLabel(ifLabel)
	v.Visit(ctx.Block(0))
	v.GenLabel(endIf)


	v.cctx.currentScope = parent

	return nil
}

func (v *Visitor) VisitFor_stmt(ctx *parsing.For_stmtContext) any {
	scopeName := BlockScopeName("for")
	startLabel := StartBlockLabel(scopeName)
	endLabel := EndBlockLabel(scopeName)
	stmtLabel := EndStmtLabel(scopeName)

	parent := v.cctx.GetCurrScope()
	scope := NewScopeTree(scopeName, parent, FOR)
	v.cctx.currentScope = scope

	v.Visit(ctx.Stmt(0))
	v.GenLabel(startLabel)
	// init variable
	// eval expr
	condition := v.Visit(ctx.Expr()).(*ScopeVar)
	// check condition, break
	v.GenCmpAddrReg(condition.offset, "0")
	v.GenJz(endLabel)

	// visit stmts
	v.Visit(ctx.Block())

	// increment
	v.GenLabel(stmtLabel)
	v.Visit(ctx.Stmt(1))

	// jump to start
	v.GenJmp(startLabel)
	v.GenLabel(endLabel)

	v.cctx.currentScope = parent
	return nil
}


func GetParentFunc(t antlr.Tree) *parsing.FuncContext {
	parent := t.GetParent()
	for parent != nil {
		if ctx, ok := parent.(*parsing.FuncContext); ok {
			return ctx
		}
		parent = parent.GetParent()
	}
	return nil 
}


func (v *Visitor) AddError(err error) {
	v.Errors = append(v.Errors, err)
}


