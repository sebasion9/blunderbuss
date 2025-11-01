package codegen
import (
	"fmt"
	"strings"
	"github.com/antlr4-go/antlr/v4"
	"blunderbuss/parsing"
)
// TODO:
// this file is big mess
// 1. Export consts to other package/file
// 2. Variables per function scope
// 3. Manage stack
// 4. Replace "ifs" with "switch case"
// 5. A wrapper for appending text to sections

// TODO: 
// do something with this types
const (
	INT = "int"
	STR = "str"
)
//TODO:
// and do something with these std calls
const (
	PRINT = "print"
)

type CodegenVisitor struct {
	*parsing.BaseBlunderbussVisitor
	text []string
	data []string
	vars map[string]int
	asm []string
}

func (v *CodegenVisitor) Visit(tree antlr.ParseTree) any {
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

func NewBlunderbussVisitor() *CodegenVisitor {
	return &CodegenVisitor {
		text: []string{},
		BaseBlunderbussVisitor: &parsing.BaseBlunderbussVisitor{},
	}
}

func (v *CodegenVisitor) VisitProgram(ctx *parsing.ProgramContext) any {
	v.text = append(v.text, "section .text")
	v.data = append(v.data, "section .data")
	v.vars = make(map[string]int)
	for _, f := range ctx.AllFunc_() {
		v.Visit(f)
	}
	asm := append(v.data, v.text...)
	return strings.Join(asm, "\n")
}

func (v *CodegenVisitor) VisitFunc(ctx *parsing.FuncContext) any {
	funcName := ctx.ID().GetText() 
	switch funcName {
		case "main":
			funcType := ctx.TYPE().GetText()
			if funcType != INT {
				return nil
			}
			v.text = append(v.text, 
				"global _start",
				"_start:",
			)
			v.Visit(ctx.Args())
			returnCode := v.Visit(ctx.Block())
			v.text = append(v.text,
				"mov rax, 60; exit syscall",
				fmt.Sprintf("mov rdi, %s", returnCode),
				"syscall",
			)

		default:
			v.text = append(v.text, 
				fmt.Sprintf("global %s", funcName),
				fmt.Sprintf("%s:", funcName),
			)
			// cache/lazy/type ??
			v.Visit(ctx.Args())
			v.Visit(ctx.Block())
			v.text = append(v.text, "ret")
	}


	return nil
}

func (v *CodegenVisitor) VisitArgs(ctx *parsing.ArgsContext) any {
	return nil
}

func (v *CodegenVisitor) VisitBlock(ctx *parsing.BlockContext) any {
	for _, s := range ctx.AllStmt() {
		stmt := v.Visit(s)
		if stmt != nil {
			return stmt
		}
	}
	return nil
}

func (v *CodegenVisitor) VisitStmt(ctx *parsing.StmtContext) any {
	if ctx.RETURN() != nil {
		return v.Visit(ctx.Expr())
	}
	if ctx.ASSIGN() != nil && ctx.TYPE() != nil && ctx.TYPE().GetText() == STR {
		expr := v.Visit(ctx.Expr()).(string)
		v.data = append(v.data,
		fmt.Sprintf("%s: db %s, 0",ctx.ID().GetText(), expr),
		)
		v.vars[ctx.ID().GetText()] = len(expr)+1

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

func (v *CodegenVisitor) VisitExpr(ctx *parsing.ExprContext) any {
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
		return ctx.ID().GetText()
	}
	return "0"
}
func (v *CodegenVisitor) VisitFunc_call(ctx *parsing.Func_callContext) any {
	if ctx.ID().GetText() == PRINT {
		arg := v.Visit(ctx.Call_args()).(string)
		v.text = append(v.text, 
			"mov rax, 1",
			"mov rdi, 1",
			fmt.Sprintf("mov rsi, %s", arg),
			fmt.Sprintf("mov rdx, %d", v.vars[arg]),
			"syscall",
		)
	}
	return nil
}
func (v *CodegenVisitor) VisitCall_args(ctx *parsing.Call_argsContext) any {
	return v.Visit(ctx.Expr(0))
}





