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

	scope := v.cctx.NewScope("program_registers")
	InitRegisters(&scope)


	scope = v.cctx.NewScope("libbbuss")
	for _, decl := range lib.GetDeclaredFuncs() {
		name := decl.Name
		type_ := decl.Type_
		args := decl.Args

		v.cctx.NewScope(name)
		var mappedArgs []ScopeFuncArg
		for i, arg := range args {
			sfa := *NewScopeFuncArg(arg.Name, TypeEnumFromStr(arg.Type_), i)
			mappedArgs = append(mappedArgs, sfa)
		}

		scope[name] = NewScopeFunc(name, mappedArgs, TypeEnumFromStr(type_))
	}


	scope = v.cctx.GetScopeByName("program")

	for _, f := range ctx.AllFunc_() {
		fnName := v.Visit(f).(*ScopeFunc)
		scope[fnName.id] = fnName
	}


	fmt.Printf("scope len %d\n", len(v.cctx.scopes))
	for k, v := range v.cctx.scopes {
		fmt.Printf("%s:%v\n",k, v)
	}

	asm := v.Codegen.StreamAsm()
	return asm
}

func (v *Visitor) VisitFunc(ctx *parsing.FuncContext) any {
	funcName := ctx.ID().GetText() 
	funcType := ctx.TYPE().GetText()
	type_ := TypeEnumFromStr(funcType)
	var args []ScopeFuncArg
	switch funcName {
		case "main":
			//TODO: handle errors
			if funcType != INT {
				return nil
			}
			v.Codegen.GenFuncInit(funcName)
			text := v.GetText()
			frameId := len(*text) - 1

			scope := v.cctx.NewScope(funcName)
			v.Visit(ctx.Args())
			v.Visit(ctx.Block())

			v.GenAlignStack((len(scope)))

			(*text)[frameId].SetSrc(strconv.Itoa(len(scope)*8))

			//TODO: 
			v.Codegen.GenFuncExit(0)

		default:
			v.Codegen.GenFuncInit(funcName)
			text := v.GetText()
			frameId := len(*text) - 1

			// cache/type, this only supports "void"
			scope := v.cctx.NewScope(funcName)
			args = v.Visit(ctx.Args()).([]ScopeFuncArg)
			v.Visit(ctx.Block())


			v.GenAlignStack(len(scope))
			(*text)[frameId].SetSrc(strconv.Itoa(len(scope)*8))
			//TODO: 
			v.Codegen.GenFuncExit(nil)
	}

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
	// scope := v.cctx.GetScopeByName(parent.ID().GetText())
	scope := v.cctx.GetCurrScope()
	registers := v.cctx.GetScopeByName("program_registers")
	if ctx.RETURN() != nil {
		expr := v.Visit(ctx.Expr()).(*ScopeVar)
		registers["rax"].(*Register).Write(expr.Type())
		//TODO: error
		if parent.TYPE().GetText() != StrFromTypeEnum(expr.Type()) {
			fmt.Println("[ERR] return type mismatch")
		}

		v.GenMovRegRelative("rax", expr.Offset())
		return nil
	}
	//TODO: change this all, add rest of assignment parsing
	//TODO: also handle errors for types != primitives
	if ctx.ASSIGN() != nil {
		if (ctx.TYPE() != nil && ctx.TYPE().GetText() == STR) ||
		(ctx.ID() != nil &&
		scope[ctx.ID().GetText()] != nil &&
		scope[ctx.ID().GetText()].Type() == STR_) {

			//TODO: handle if variable exists in scope
			rhs := *v.Visit(ctx.Expr()).(*ScopeVar)
			lhs := scope[ctx.ID().GetText()]

			// TODO: handle
			if rhs.Type() != STR_ {
				fmt.Println("[ERR] not a string")
				return nil
			}

			if lhs == nil {
				offset := len(scope)*8+8
				lhs = NewScopeVar(0, STR_, offset)
			}

			// func, so rax
			if rhs.Raw() == nil {
				v.Codegen.GenMovAddrRelative(lhs.Offset(), "rax")
				rax, _ := registers["rax"].(*Register)
				lhs.(*ScopeVar).expr = rax.Raw()
				scope[ctx.ID().GetText()] = lhs
			} else {
				v.Codegen.GenMovRegRelative("rax", rhs.Offset())
				v.Codegen.GenMovAddrRelative(lhs.Offset(), "rax")
				lhs.(*ScopeVar).expr = rhs.expr
				scope[ctx.ID().GetText()] = lhs

			}

		}
		if (ctx.TYPE() != nil && ctx.TYPE().GetText() == INT) ||
		(ctx.ID() != nil &&
		scope[ctx.ID().GetText()] != nil &&
		scope[ctx.ID().GetText()].Type() == INT_) {

			rhs := *v.Visit(ctx.Expr()).(*ScopeVar)
			lhs := scope[ctx.ID().GetText()]
			//TODO: errors
			if rhs.Type() != INT_ {
				fmt.Println("[ERR] not a number")
				return nil
			}

			if lhs == nil {
				offset := len(scope)*8+8
				lhs = NewScopeVar(0, INT_, offset)
			}

			// func, so rax
			if rhs.Raw() == nil {
				v.Codegen.GenMovAddrRelative(lhs.Offset(), "rax")
				rax, _ := registers["rax"].(*Register)
				lhs.(*ScopeVar).expr = rax.Raw()
				scope[ctx.ID().GetText()] = lhs
			} else {
				v.Codegen.GenMovRegRelative("rax", rhs.Offset())
				v.Codegen.GenMovAddrRelative(lhs.Offset(), "rax")
				lhs.(*ScopeVar).expr = rhs.expr
				scope[ctx.ID().GetText()] = lhs
			}

		}
		return nil
	}

	if ctx.TYPE() != nil && ctx.ID() != nil {
		switch ctx.TYPE().GetText() {
		case INT:
			offset := len(scope)*8+8
			scope[ctx.ID().GetText()] = NewScopeVar(0, INT_, offset)
		case STR:
			offset := len(scope)*8+8
			scope[ctx.ID().GetText()] = NewScopeVar("", STR_, offset)
		}

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
	scope := v.cctx.GetCurrScope()
	if ctx.NUM() != nil {
		id := v.CreateId()
		val, _ := strconv.Atoi(ctx.NUM().GetText())
		text := fmt.Sprintf("qword %d", val)
		v.Codegen.GenIntPrimitive(id, val)
		offset := len(scope)*8+8
		v.GenMovAddrRelative(offset, text)
		svar := NewScopeVar(val, INT_, offset)
		scope[id] = svar
 
		return svar
	}
	if ctx.STRING() != nil {
		id := v.CreateId()
		val := ctx.STRING().GetText()
		v.Codegen.GenStrPrimitive(id, val)
		offset := len(scope)*8+8
		v.GenMovMemory("rax", id)
		v.GenMovAddrRelative(offset, "rax")
		svar := NewScopeVar(id, STR_, offset)
		scope[id] = svar

		return svar
	}
	if ctx.Func_call() != nil {
		offset := len(scope)*8+8
		svar := v.Visit(ctx.Func_call()).(*ScopeVar) 
		svar.offset = offset
		v.GenMovAddrRelative(offset, "rax")
		scope[v.CreateId()] = svar
		return svar
	}
	if ctx.ID() != nil {
		return scope[ctx.ID().GetText()]
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
		// expr := lhs.expr.(int) + rhs.expr.(int)
		offset := len(scope)*8+8
		prod := NewScopeVar(nil, INT_, offset)
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
		// expr := lhs.expr.(int) - rhs.expr.(int)
		offset := len(scope)*8+8
		prod := NewScopeVar(nil, INT_, offset)
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
		// expr := lhs.expr.(int) * rhs.expr.(int)
		offset := len(scope)*8+8
		prod := NewScopeVar(nil, INT_, offset)
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
		// expr := lhs.expr.(int) / rhs.expr.(int)
		offset := len(scope)*8+8
		prod := NewScopeVar(nil, INT_, offset)

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
		offset := len(scope)*8+8
		prod := NewScopeVar(nil, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenCmp("rax", "rbx")
		v.GenSete("al")
		v.GenMovAddrRelative(offset, "al")
		scope[v.CreateId()] = prod
		return prod
	}

	if ctx.LE() != nil {
		lhs := v.Visit(ctx.Expr(0)).(*ScopeVar)
		rhs := v.Visit(ctx.Expr(1)).(*ScopeVar)
		if lhs.Type() != rhs.Type() && lhs.Type() != INT_ {
			fmt.Println("[ERR] mismatched types")
			return nil
		}
		offset := len(scope)*8+8
		prod := NewScopeVar(nil, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenCmp("rax", "rbx")
		v.GenSetle("al")
		v.GenMovAddrRelative(offset, "al")
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
		offset := len(scope)*8+8
		prod := NewScopeVar(nil, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenCmp("rbx", "rax")
		v.GenSetl("al")
		v.GenMovAddrRelative(offset, "al")
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
		offset := len(scope)*8+8
		prod := NewScopeVar(nil, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenCmp("rbx", "rax")
		v.GenSetge("al")
		v.GenMovAddrRelative(offset, "al")
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
		offset := len(scope)*8+8
		prod := NewScopeVar(nil, INT_, offset)
		v.GenMovRegRelative("rax", lhs.offset)
		v.GenMovRegRelative("rbx", rhs.offset)
		v.GenCmp("rbx", "rax")
		v.GenSetg("al")
		v.GenMovAddrRelative(offset, "al")
		scope[v.CreateId()] = prod
		return prod

	}
	return nil
}

func (v *Visitor) VisitFunc_call(ctx *parsing.Func_callContext) any {
	// this is libbbuss part, but stays here for now
	//TODO: libbuss should have predefined print and other standard functions in lib package
	registers := v.cctx.GetScopeByName("program_registers")

	program := v.cctx.GetScopeByName("program")
	lib := v.cctx.GetScopeByName("libbbuss")
	fnName := ctx.ID().GetText()
	fn, ok := program[fnName].(*ScopeFunc)
	if ok {
		scope := v.cctx.GetScopeByName(fnName)
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
			scope[fn.args[i].expr.(string)] = expr
			v.GenPushArg(expr.Offset(), i)
		}
		v.GenCallFunc(fnName)
		v.GenShrinkStackFrame(len(fn.args) - 6)

		rax := registers["rax"].(*Register)
		rax.Write(fn.type_)

		return NewScopeVar(nil, fn.Type(), 0)
	}

	fn, ok = lib[ctx.ID().GetText()].(*ScopeFunc)
	if ok {
		scope := v.cctx.GetScopeByName(fnName)
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

			scope[fn.args[i].expr.(string)] = expr
			v.GenPushArg(expr.Offset(), i)
		}
		v.GenMovMemory("rax", "0")
		v.GenCallFunc(fnName)
		v.GenShrinkStackFrame(len(fn.args) - 6)
		rax := registers["rax"].(*Register)
		rax.Write(fn.type_)

		return NewScopeVar(nil, fn.Type(), 0)

	}

	return nil
}

func (v *Visitor) VisitIf_stmt(ctx *parsing.If_stmtContext) any {
	scopeName := BlockScopeName("if")
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

	return nil
}

func (v *Visitor) VisitFor_stmt(ctx *parsing.For_stmtContext) any {
	scopeName := BlockScopeName("for")
	startLabel := fmt.Sprintf("START__%s", scopeName)
	endLabel := fmt.Sprintf("END__%s", scopeName)
	// v.cctx.NewScope(scopeName)
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




