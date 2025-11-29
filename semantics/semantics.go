package semantics

import (
	"blunderbuss/codegen"
	"blunderbuss/lib"
	"blunderbuss/parsing"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

// TODO:
// ~1. Export consts to other package/file~
// ~2. Variables per function scope/per block scope~
// ~3. Manage stack and func calling conventions!~
// 4. Replace "ifs" with "switch case"
// 9. Semantic errors for compilers
// 10. Go all through visits, check missing parts of grammar
// 12. Clean up todos/comments
// 14. ~structure with registers, new scope struct register~
// 15. check if used variables are in scope...
// 16. ~global -> stack~
// 17. ~return stmt~
// ~18. for, if~
// ~7. Operator precedence~
// 19. array - heap
// 20. scope tree, id lookup
// 21. offset()
// 22. force else if order..

type Visitor struct {
	*parsing.BaseBlunderbussVisitor
	codegen.Codegen
	cctx CompilerContext
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
	case *parsing.Effect_blockContext:
		return v.VisitEffect_block(ctx)
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


	scope = NewScopeTree("libbbuss", parent, OTHER)
	v.cctx.currentScope = scope
	for _, decl := range lib.GetDeclaredFuncs() {
		name := decl.Name
		type_ := decl.Type_
		args := decl.Args

		var mappedArgs []ScopeFuncArg
		for i, arg := range args {
			sfa := *NewScopeFuncArg(arg.Name, TypeEnumFromStr(arg.Type_), i)
			mappedArgs = append(mappedArgs, sfa)
		}
		// (*parent.GetVars())[name] = NewScopeFunc(name, mappedArgs, TypeEnumFromStr(type_))
		(*scope.GetVars())[name] = NewScopeFunc(name, mappedArgs, TypeEnumFromStr(type_))

	}
	v.cctx.currentScope = parent



	for _, f := range ctx.AllFunc_() {
		fn := v.Visit(f).(*ScopeFunc)
		(*parent.GetVars())[fn.id] = fn
	}

	parent.DebugPrint(0)

	asm := v.Codegen.StreamAsm()
	return asm
}

func (v *Visitor) VisitFunc(ctx *parsing.FuncContext) any {
	funcName := ctx.ID().GetText() 
	funcType := ctx.TYPE().GetText()
	type_ := TypeEnumFromStr(funcType)
	var args []ScopeFuncArg
	//TODO: handle errors
	v.Codegen.GenFuncInit(funcName)
	text := v.GetText()
	frameId := len(*text) - 1

	parent := v.cctx.GetCurrScope()
	scope := NewScopeTree(funcName, parent, FUNC)
	v.cctx.currentScope = scope

	v.Visit(ctx.Args())
	v.Visit(ctx.Block())

	endFn := EndFnLabel(funcName)
	v.GenLabel(endFn)


	offset := scope.GetOff()
	// varsLen := GetAllVarsDownLen(scope)
	// v.GenAlignStack(varsLen)
	v.GenAlignStack((offset - 8)/8)
	// v.GenAlignStack((len(scope.vars)))

	// (*text)[frameId].SetSrc(strconv.Itoa(varsLen*8))
	(*text)[frameId].SetSrc(strconv.Itoa(offset - 8))

	v.Codegen.GenFuncExit()


	v.cctx.currentScope = parent
	return NewScopeFunc(funcName, args, type_)
}

//TODO:
func (v *Visitor) VisitArgs(ctx *parsing.ArgsContext) any {
	var args []ScopeFuncArg
	for i, p := range ctx.AllParam() {
		name := p.ID().GetText()
		typeText := p.TYPE().GetText()
		type_ := VOID_
		switch typeText {
		case "int":
			type_ = INT_
		case "str":
			type_ = STR_
		default:
			type_ = VOID_
		}
		args = append(args, *NewScopeFuncArg(name, type_, i))
	}
	return args
}

func (v *Visitor) VisitBlock(ctx *parsing.BlockContext) any {
	//TODO: consider adding scope on block level
	for _, s := range ctx.AllStmt() {
		v.Visit(s)
	}
	return nil
}

func (v *Visitor) VisitStmt(ctx *parsing.StmtContext) any {
	parent := GetParentFunc(ctx)
	//TODO:
	if parent == nil {
		return nil
	}
	currScope := v.cctx.GetCurrScope()
	scopeTree := GetWhatFunc(currScope)
	// scope := *scopeTree.GetVars()
	root := GetRootScope(scopeTree)
	registers := *root.GetScopeByName("program_registers").GetVars()


	if ctx.RETURN() != nil {
		expr := v.Visit(ctx.Expr()).(*ScopeVar)
		registers["rax"].(*Register).Write(expr.Type())
		//TODO: error
		if parent.TYPE().GetText() != StrFromTypeEnum(expr.Type()) {
			fmt.Println("[ERR] return type mismatch")
		}

		endFn := EndFnLabel(parent.ID().GetText())

		v.GenMovRegRelative("rax", expr.Offset())
		v.GenJmp(endFn)
		return nil
	}
	//TODO: change this all, add rest of assignment parsing
	//TODO: also handle errors for types != primitives


	// assign to scope, initialize
	if ctx.ASSIGN() != nil {
		// check if lhs in scope
		lhsName := ctx.ID().GetText()

		// look for var in scopes
		varScope := SearchIdUp(currScope, lhsName)
		var lhs *ScopeVar


		if varScope != nil {
			// assign existing variable
			existingVar := (*varScope.GetVars())[lhsName]
			if existingVar != nil && ctx.TYPE() == nil {
				//(*varScope.GetVars())[lhsName].(*ScopeVar)
				lhs = existingVar.(*ScopeVar)
			// variable reinitialization
			} else if existingVar != nil && ctx.TYPE() != nil {
				fmt.Println("[ERR] variable reinitialization")
				return nil
			}
		} else {
			// undefined variable
			if ctx.TYPE() == nil {
				fmt.Println("[ERR] undefined variable")
				return nil
			// initialize new variable
			} else {
				varScope = v.cctx.GetCurrScope()
				offset := scopeTree.GetOff()
				scopeTree.IncrOff(8)
				lhs = NewScopeVar(0, TypeEnumFromStr(ctx.TYPE().GetText()), offset)
			}

		}
		// if varScope != nil && ctx.TYPE() == nil {
		// } else if varScope != nil && ctx.TYPE() != nil {
		// } else if varScope == nil && ctx.TYPE() == nil {
		// } else {
		// }

		rhs := *v.Visit(ctx.Expr()).(*ScopeVar)

		if lhs.Type() != rhs.Type() {
			fmt.Println("[ERR] type mismatch")
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
			scopeTree.IncrOff(8)
			(*currScope.GetVars())[ctx.ID().GetText()] = NewScopeVar(0, INT_, offset)
		case STR:
			offset := scopeTree.GetOff()
			scopeTree.IncrOff(8)
			(*currScope.GetVars())[ctx.ID().GetText()] = NewScopeVar("", STR_, offset)
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
		// get parent scope -> jmp end
	}

	if ctx.NEXT() != nil {
		// get parent scope -> jmp start
	}


	return nil
}

func (v *Visitor) VisitExpr(ctx *parsing.ExprContext) any {
	//TODO: the primitives should be saved in .data and referenced as variables
	//TODO: then they can return the variables identifier and referenced as in scope
	// if returns nil -> then register
	scopeTree := v.cctx.GetCurrScope()
	funcTree := GetWhatFunc(scopeTree)
	scope := *scopeTree.GetVars()
	if ctx.NUM() != nil {
		id := v.CreateId()
		val, _ := strconv.Atoi(ctx.NUM().GetText())
		text := fmt.Sprintf("qword %d", val)
		v.Codegen.GenIntPrimitive(id, val)
		offset := funcTree.GetOff()
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
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
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
		v.GenMovMemory("rax", id)
		v.GenMovAddrRelative(offset, "rax")
		svar := NewScopeVar(id, STR_, offset)
		scope[id] = svar

		return svar
	}
	if ctx.Func_call() != nil {
		offset := funcTree.GetOff()
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
		svar := v.Visit(ctx.Func_call()).(*ScopeVar) 
		svar.offset = offset
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = svar
		return svar
	}
	if ctx.ID() != nil {
		// svar := (*SearchIdUp(scopeTree, ctx.ID().GetText()).GetVars())[ctx.ID().GetText()]
		id := ctx.ID().GetText()

		// look for var in scopes
		varScope := SearchIdUp(scopeTree, id)
		// couldnt find scope with that lhs name
		//TODO: err
		if varScope == nil {
			fmt.Println("[ERR] undefined variable")
		}
		return (*varScope.GetVars())[id]
	}

	if ctx.LPAREN() != nil && ctx.RPAREN() != nil {
		return v.Visit(ctx.Expr(0))
	}

	// operators
	if ctx.PLUS() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		//TODO: err mismatched types
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			fmt.Println("[ERR] mismatched types")
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
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
		//TODO: err mismatched types
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			fmt.Println("[ERR] mismatched types")
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rsi", rhs.offset)
		v.GenSubMemory("rax", "rsi")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod
	}

	

	if ctx.MULT() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		//TODO: err mismatched types
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			fmt.Println("[ERR] mismatched types")
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
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
			fmt.Println("[ERR] mismatched types")
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
		prod := NewScopeVar(0, INT_, offset)

		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenXor("rdx", "rdx")
		v.GenDiv()
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod
	}

	if ctx.EQUAL() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			fmt.Println("[ERR] mismatched types")
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenCmp("rax", "rbx")
		v.GenSete("al")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod
	}

	if ctx.NOT_EQUAL() != nil {
		return nil
	}

	if ctx.LE() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			fmt.Println("[ERR] mismatched types")
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenCmp("rax", "rbx")
		v.GenSetle("al")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod

	}

	if ctx.LT() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			fmt.Println("[ERR] mismatched types")
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenCmp("rbx", "rax")
		v.GenSetl("al")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod

	}

	if ctx.GE() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			fmt.Println("[ERR] mismatched types")
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenCmp("rbx", "rax")
		v.GenSetge("al")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod
	}

	if ctx.GT() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			fmt.Println("[ERR] mismatched types")
			return nil
		}
		offset := funcTree.GetOff()
		funcTree.IncrOff(8)
		// offset := len(scope)*8+8
		prod := NewScopeVar(0, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenCmp("rbx", "rax")
		v.GenSetg("al")
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = prod
		return prod

	}

	if ctx.AND() != nil {
		return nil

	}

	if ctx.OR() != nil {
		return nil
	}


	return nil
}

func (v *Visitor) VisitFunc_call(ctx *parsing.Func_callContext) any {
	// this is libbbuss part, but stays here for now
	//TODO: libbuss should have predefined print and other standard functions in lib package
	scope := v.cctx.GetCurrScope()
	root := GetRootScope(scope)
	registers := *root.GetScopeByName("program_registers").GetVars()

	program := *root.GetVars()
	lib := *root.GetScopeByName("libbbuss").GetVars()
	fnName := ctx.ID().GetText()
	fn, ok := program[fnName].(*ScopeFunc)
	if ok {
		fnScopeTree := NewScopeTree(fnName, scope, FNCALL)
		fnScope := *fnScopeTree.GetVars()
		v.cctx.currentScope = fnScopeTree
		
		//TODO: remove dbg
		for _, a := range fn.args {
			fmt.Printf("%s:%v:%d\n", a.expr.(string), StrFromTypeEnum(a.type_), a.idx)
		}

		if len(ctx.Call_args().AllExpr()) != len(fn.args) {
			fmt.Println("[ERR] mismatched args amount")
			return nil
		}

		for i, arg := range ctx.Call_args().AllExpr() {
			expr := v.Visit(arg).(*ScopeVar)
			//TODO: err invalid type and amount of args

			if fn.args[i].type_ != ANY_ && (fn.args[i].type_ != expr.type_) {
				fmt.Println("[ERR] wrong type")
				return nil
			}
			fnScope[fn.args[i].expr.(string)] = expr
			v.GenPushArg(expr.Offset(), i)
		}
		v.GenCallFunc(fnName)
		v.GenShrinkStackFrame(len(fn.args) - 6)

		rax := registers["rax"].(*Register)
		rax.Write(fn.type_)


		v.cctx.currentScope = scope
		return NewScopeVar(nil, fn.Type(), 0)
	}

	fn, ok = lib[ctx.ID().GetText()].(*ScopeFunc)
	if ok {
		// scope := v.cctx.GetScopeByName(fnName)
		fnScopeTree := NewScopeTree(fnName, scope, FNCALL)
		fnScope := *fnScopeTree.GetVars()
		v.cctx.currentScope = fnScopeTree
		//TODO: remove dbg
		for _, a := range fn.args {
			fmt.Printf("%s:%v:%d\n", a.expr.(string), StrFromTypeEnum(a.type_), a.idx)
		}

		//TODO: err handling (what if doesnt exist etc)
		for i, arg := range ctx.Call_args().AllExpr() {
			expr := v.Visit(arg).(*ScopeVar)
			//TODO: err invalid type

			if fn.args[i].type_ != ANY_ && (fn.args[i].type_ != expr.type_) {
				fmt.Println("[ERR] wrong type")
				return nil
			}

			fnScope[fn.args[i].expr.(string)] = expr
			v.GenPushArg(expr.Offset(), i)
		}
		v.GenMovMemory("rax", "0")
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
	scope := NewScopeTree(scopeName, parent, BLOCK)
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
	startLabel := fmt.Sprintf("START__%s", scopeName)
	endLabel := fmt.Sprintf("END__%s", scopeName)
	parent := v.cctx.GetCurrScope()
	scope := NewScopeTree(scopeName, parent, BLOCK)
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




