package semantics

import (
	"blunderbuss/codegen"
	"blunderbuss/lib"
	"blunderbuss/parsing"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// TODO:
// ~1. Export consts to other package/file~
// ~2. Variables per function scope/per block scope~
// 3. Manage stack and func calling conventions!
// 4. Replace "ifs" with "switch case"
// 5. A wrapper for appending text to sections, mov(), add() etc.
// 7. Operator precedence
// 8. lookup wiki
// 9. Semantic errors for compilers
// 10. Go all through visits, check missing parts of grammar
// 12. Clean up todos/comments
// 13. In-place primitives should generate labels in .data and save them in scope as such

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

	scope := v.cctx.NewScope("libbbuss")
	for _, pair := range lib.GetDeclaredFuncs() {
		name := pair[0]
		type_ := pair[1]
		scope[name] = NewScopeFunc(name, TypeEnumFromStr(type_))
	}


	//TODO: should be in lib

	v.Codegen.GenFuncInit("print")
	v.Codegen.GenFuncExit(nil)

	scope = v.cctx.GetScopeByName("program")

	for _, f := range ctx.AllFunc_() {
		fnName := v.Visit(f).(*ScopeFunc)
		scope[fnName.id] = fnName
	}


	fmt.Printf("scope len %d\n", len(v.cctx.scopes))
	for k, v := range v.cctx.scopes {
		fmt.Printf("%s:%v\n",k, v)
	}

	asm := v.Codegen.ConcatAsm()
	return strings.Join(asm, "\n")
}

func (v *Visitor) VisitFunc(ctx *parsing.FuncContext) any {
	funcName := ctx.ID().GetText() 
	funcType := ctx.TYPE().GetText()
	type_ := TypeEnumFromStr(funcType)
	switch funcName {
		case "main":
			//TODO: handle errors
			if funcType != INT {
				return nil
			}
			// v.Codegen.GenMainInit(&v.text)
			v.Codegen.GenFuncInit(funcName)

			_ = v.cctx.NewScope(funcName)
			v.Visit(ctx.Args())
			v.Visit(ctx.Block())

			// v.Codegen.GenFuncExit(scope["return"].Raw().(int))
			v.Codegen.GenFuncExit(0)

		default:
			//TODO: declare as global in .text
			//TODO: modify to generate function at top? another array
			//TODO: calling convention
			//TODO: handle stack
			v.Codegen.GenFuncInit(funcName)

			// cache/type, this only supports "void"
			v.cctx.NewScope(funcName)
			v.Visit(ctx.Args())
			v.Visit(ctx.Block())
			//TODO: return in rax
			v.Codegen.GenFuncExit(nil)
	}

	return NewScopeFunc(funcName, type_)
}

//TODO:
func (v *Visitor) VisitArgs(ctx *parsing.ArgsContext) any {
	parent := ctx.GetParent().(*parsing.FuncContext).ID().GetText()
	scope := v.cctx.GetScopeByName(parent)
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
		scope[name] = NewScopeFuncArg(name, type_, i)
	}
	return nil
}

func (v *Visitor) VisitBlock(ctx *parsing.BlockContext) any {
	//TODO: consider adding scope on block level
	for _, s := range ctx.AllStmt() {
		v.Visit(s)
	}
	return nil
}

func (v *Visitor) VisitStmt(ctx *parsing.StmtContext) any {
	parent, ok := ctx.GetParent().GetParent().(*parsing.FuncContext)
	//TODO:
	if !ok {
		return nil
	}
	scope := v.cctx.GetScopeByName(parent.ID().GetText())
	if ctx.RETURN() != nil {
		rcStr := v.Visit(ctx.Expr()).(string)
		rc, err := strconv.Atoi(rcStr)
		//TODO: error
		if err != nil {
			return nil
		}
		scope["return"] = NewScopeVar(rc, INT_)
	}
	//TODO: change this all, add rest of assignment parsing
	//TODO: also handle errors for types != primitives
	if ctx.ASSIGN() != nil && ctx.TYPE() != nil {
		if ctx.TYPE().GetText() == STR {
			//TODO: handle if variable exists in scope
			id := v.Visit(ctx.Expr()).(string)
			rhs := scope[id]
			// TODO: handle
			if rhs.Type() != STR_ {
				return nil
			}
			val := rhs.Raw().(string)
			//TODO: handle consts primitive scope (uniqueness), stack
			scope[ctx.ID().GetText()] = NewScopeVar(val, STR_)

			v.GenStrPrimitive(ctx.ID().GetText(), val)
		}
		if ctx.TYPE().GetText() == INT {
			expr := v.Visit(ctx.Expr()).(string)
			//TODO: also handle errors for types != primitives
			num, err := strconv.Atoi(expr)
			if err != nil {
				return nil
			}
			scope[ctx.ID().GetText()] = NewScopeVar(expr, INT_)
			v.GenIntPrimitive(ctx.ID().GetText(), num)
		}
	}

	if ctx.Func_call() != nil {
		v.Visit(ctx.Func_call())
	}

	return nil
}

func (v *Visitor) VisitExpr(ctx *parsing.ExprContext) any {
	//TODO: the primitives should be saved in .data and referenced as variables
	//TODO: then they can return the variables identifier and referenced as in scope
	scope := v.cctx.GetCurrScope()
	if ctx.NUM() != nil {
		id := v.Codegen.CreateId()
		//TODO: double
		val, _ := strconv.Atoi(ctx.NUM().GetText())
		scope[id] = NewScopeVar(val, INT_)
		v.Codegen.GenIntPrimitive(id, val)
		return id
	}
	if ctx.STRING() != nil {
		id := v.Codegen.CreateId()
		//TODO: double
		val := ctx.STRING().GetText()
		scope[id] = NewScopeVar(val, STR_)
		v.Codegen.GenStrPrimitive(id, val)
		return id
	}
	if ctx.Func_call() != nil {
		//TODO: expr -> id
		v.Visit(ctx.Func_call())
	}
	if ctx.ID() != nil {
		return ctx.ID().GetText()
	}
	//TODO: handle OP precedence
	if ctx.Expr(0) != nil && ctx.Expr(1) != nil && ctx.BIN_OP() != nil {
		return nil
	}
	if ctx.SIN_OP() != nil && ctx.Expr(0) != nil {
		return nil
	}
	return nil
}

func (v *Visitor) VisitFunc_call(ctx *parsing.Func_callContext) any {
	// func call should ask SCOPE for declared func with same name, and grab its declared arguments
	// this is libbbuss part, but stays here for now
	// handle nums and other
	//TODO: libbuss should have predefined print and other standard functions in lib package
	//TODO: calling convntion
	//TODO: returning calling convention (rax)
	if ctx.ID().GetText() == PRINT {
		arg := v.Visit(ctx.Call_args()).(string)
		v.Codegen.GenPrint(arg)
	}
	return nil
}
func (v *Visitor) VisitCall_args(ctx *parsing.Call_argsContext) any {
	expr := v.Visit(ctx.Expr(0))
	return expr
}





