package semantics

import (
	"blunderbuss/codegen"
	"blunderbuss/parsing"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// TODO:
// ~1. Export consts to other package/file~
// ~2. Variables per function scope/per block scope~
// 3. Manage stack
// 4. Replace "ifs" with "switch case"
// 5. A wrapper for appending text to sections, mov(), add() etc.
// 6. Create struct for handling certain features
// 7. Operator precedence
// 8. lookup wiki
// 9. Semantic errors for compilers
// 10. Go all through visits, check missing parts
// 11. GetParent
// 12. Clean up todos/comments
// 13. func calls are complicated (identificators vs expressions), (varName vs "133")
// 14. func calls arguments, are assigning internally: func a(b, c)			a("123", varname) -> a = "123", b = varname -> primitive

type Visitor struct {
	*parsing.BaseBlunderbussVisitor
	codegen.Codegen
	cctx CompilerContext
	text []string
	data []string
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
		text: []string{},
		BaseBlunderbussVisitor: &parsing.BaseBlunderbussVisitor{},
		Codegen: codegen.Codegen{},
		cctx: NewCompilerContext(),
		
	}
}

func (v *Visitor) VisitProgram(ctx *parsing.ProgramContext) any {
	v.text = append(v.text, "section .text")
	v.data = append(v.data, "section .data")

	for _, f := range ctx.AllFunc_() {
		v.Visit(f)
	}

	asm := append(v.data, v.text...)
	fmt.Printf("scope len %d\n", len(v.cctx.scopes))
	for i := 0; i < len(v.cctx.scopes); i++{
		fmt.Printf("%v\n", v.cctx.scopes[i])
	}

	return strings.Join(asm, "\n")
}

func (v *Visitor) VisitFunc(ctx *parsing.FuncContext) any {
	funcName := ctx.ID().GetText() 
	parentScope := *v.cctx.GetCurrScope()
	parentScope[funcName] = NewScopeFunc(funcName)
	switch funcName {
		case "main":
			funcType := ctx.TYPE().GetText()
			if funcType != INT {
				return nil
			}
			v.Codegen.GenMainInit(&v.text)
			scope := *v.cctx.NewScope()

			v.Visit(ctx.Args())
			v.Visit(ctx.Block())

			v.Codegen.GenMainExit(&v.text, scope["return"].Raw().(int))

		default:
			v.Codegen.GenFuncInit(&v.text, funcName)

			// cache/type, this only supports "void"
			_ = *v.cctx.NewScope()
			v.Visit(ctx.Args())
			v.Visit(ctx.Block())
			v.Codegen.GenFuncExit(&v.text)
	}


	return nil
}

//TODO:
func (v *Visitor) VisitArgs(ctx *parsing.ArgsContext) any {
	return nil
}

func (v *Visitor) VisitBlock(ctx *parsing.BlockContext) any {
	for _, s := range ctx.AllStmt() {
		v.Visit(s)
	}
	return nil
}

func (v *Visitor) VisitStmt(ctx *parsing.StmtContext) any {
	scope := *v.cctx.GetCurrScope()
	if ctx.RETURN() != nil {
		rcStr := v.Visit(ctx.Expr()).(string)
		rc, err := strconv.Atoi(rcStr)
		//TODO: error
		if err != nil {
			return nil
		}
		scope["return"] = NewScopeVar(rc)
	}
	//TODO: change this all, add rest of assignment parsing
	//TODO: also handle errors for types != primitives
	if ctx.ASSIGN() != nil && ctx.TYPE() != nil {
		if ctx.TYPE().GetText() == STR {
			//TODO: handle if variable exists in scope
			expr := v.Visit(ctx.Expr()).(string)
			//TODO: handle consts primitive scope (uniqueness)
			scope[ctx.ID().GetText()] = NewScopeVar(expr)

			// v.data = append(v.data,
			// 	fmt.Sprintf("%s: db %s, 0",ctx.ID().GetText(), expr),
			// )
			v.GenStrPrimitive(&v.data, ctx.ID().GetText(), expr)
		}
		if ctx.TYPE().GetText() == INT {
			expr := v.Visit(ctx.Expr()).(string)
			//TODO: also handle errors for types != primitives
			num, err := strconv.Atoi(expr)
			if err != nil {
				return nil
			}
			scope[ctx.ID().GetText()] = NewScopeVar(expr)
			v.GenIntPrimitive(&v.data, ctx.ID().GetText(), num)
		}
	}

	if ctx.Func_call() != nil {
		v.Visit(ctx.Func_call())
	}

	return nil
}

func (v *Visitor) VisitExpr(ctx *parsing.ExprContext) any {
	scope := *v.cctx.GetCurrScope()
	if ctx.NUM() != nil {
		return ctx.NUM().GetText()
	}
	if ctx.STRING() != nil {
		return ctx.STRING().GetText()
	}
	if ctx.Func_call() != nil {
		v.Visit(ctx.Func_call())
	}
	if ctx.ID() != nil {
		parent := ctx.GetParent()
		_, ok := parent.(*parsing.Call_argsContext)
		if !ok {
			return scope[ctx.ID().GetText()].Raw()
		}
		return ctx.ID().GetText()
	}
	return "0"
}

func (v *Visitor) VisitFunc_call(ctx *parsing.Func_callContext) any {
	// scope := *v.cctx.GetCurrScope()
	// this is libbbuss part, but stays here for now
	// handle nums and other
	if ctx.ID().GetText() == PRINT {
		arg := v.Visit(ctx.Call_args()).(string)
		v.text = append(v.text, 
			fmt.Sprintf("mov rdi, %s", arg),
			"call puts",
		)

	}
	return nil
}
func (v *Visitor) VisitCall_args(ctx *parsing.Call_argsContext) any {
	// scope := *v.cctx.GetCurrScope()
	expr := v.Visit(ctx.Expr(0))
	return expr
	
	// return nil
}





