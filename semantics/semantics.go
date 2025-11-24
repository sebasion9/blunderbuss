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
// this file is big mess
// ~1. Export consts to other package/file~
// 2. Variables per function scope/per block scope
// 3. Manage stack
// 4. Replace "ifs" with "switch case"
// 5. A wrapper for appending text to sections, mov(), add() etc.
// 6. Create struct for handling certain features
// 7. Operator precedence
// 8. lookup wiki
// 9. Semantic errors for compilers
// 10. Go all through visits, check missing parts

type Visitor struct {
	*parsing.BaseBlunderbussVisitor
	codegen.Codegen
	CompilerContext
	text []string
	data []string
	// global variables for now
	vars map[string]any
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
		CompilerContext: NewCompilerContext(),
		
	}
}

func (v *Visitor) VisitProgram(ctx *parsing.ProgramContext) any {
	v.text = append(v.text, "section .text")
	v.data = append(v.data, "section .data")
	// TODO: change to CompilerContext scopes
	v.vars = make(map[string]any)
	v.vars["return"] = 0
	for _, f := range ctx.AllFunc_() {
		v.Visit(f)
	}
	asm := append(v.data, v.text...)
	return strings.Join(asm, "\n")
}

func (v *Visitor) VisitFunc(ctx *parsing.FuncContext) any {
	funcName := ctx.ID().GetText() 
	switch funcName {
		case "main":
			funcType := ctx.TYPE().GetText()
			if funcType != INT {
				return nil
			}
			v.Codegen.GenMainInit(&v.text)

			v.Visit(ctx.Args())
			v.Visit(ctx.Block())

			v.Codegen.GenMainExit(&v.text, v.vars["return"].(int))

		default:
			v.Codegen.GenFuncInit(&v.text, funcName)

			// cache/type ??
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
	if ctx.RETURN() != nil {
		rcStr := v.Visit(ctx.Expr()).(string)
		rc, err := strconv.Atoi(rcStr)
		//TODO: error
		if err != nil {
			return nil
		}
		v.vars["return"] = rc
		// return v.Visit(ctx.Expr())
	}
	if ctx.ASSIGN() != nil && ctx.TYPE() != nil && ctx.TYPE().GetText() == STR {
		//TODO: change this all
		expr := v.Visit(ctx.Expr()).(string)
		v.data = append(v.data,
		fmt.Sprintf("%s: db %s, 0",ctx.ID().GetText(), expr),
		)
		v.vars[ctx.ID().GetText()] = expr//len(expr)+1

	}
	if ctx.Func_call() != nil {
		v.Visit(ctx.Func_call())
	}

	// exprCtx := ctx.Expr()
	// if exprCtx != nil {
	// 	exprVal := v.Visit(ctx.Expr()).(string)
	// 	v.text = append(v.text, fmt.Sprintf("mov rax, %s", exprVal))
	// }
	return nil
}

func (v *Visitor) VisitExpr(ctx *parsing.ExprContext) any {
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
		//TODO: cant' be global
		// return v.vars[ctx.ID().GetText()]
		return ctx.ID().GetText()
	}
	return "0"
}
func (v *Visitor) VisitFunc_call(ctx *parsing.Func_callContext) any {
	if ctx.ID().GetText() == PRINT {
		arg := v.Visit(ctx.Call_args()).(string)
		v.text = append(v.text, 
			"mov rax, 1",
			"mov rdi, 1",
			fmt.Sprintf("mov rsi, %s", arg),
			fmt.Sprintf("mov rdx, %d", len(v.vars[arg].(string)) + 1), //v.vars[arg]),
			"syscall",
		)
	}
	return nil
}
func (v *Visitor) VisitCall_args(ctx *parsing.Call_argsContext) any {
	return v.Visit(ctx.Expr(0))
}





