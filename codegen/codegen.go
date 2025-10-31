package codegen
import (
	"fmt"
	"strings"
	"github.com/antlr4-go/antlr/v4"
	"blunderbuss/parsing"
)

type CodegenVisitor struct {
	*parsing.BaseBlunderbussVisitor
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
		asm: []string{},
		BaseBlunderbussVisitor: &parsing.BaseBlunderbussVisitor{},
	}
}

func (v *CodegenVisitor) VisitProgram(ctx *parsing.ProgramContext) any {
	for _, f := range ctx.AllFunc_() {
		v.Visit(f)
	}
	return strings.Join(v.asm, "\n")
}

func (v *CodegenVisitor) VisitFunc(ctx *parsing.FuncContext) any {
	funcName := ctx.ID().GetText()
	v.asm = append(v.asm, 
		fmt.Sprintf("global %s", funcName),
		"section .text",
		fmt.Sprintf("%s:", funcName),
	)
	// cache/lazy/type ??
	v.Visit(ctx.Args())
	v.Visit(ctx.Block())

	v.asm = append(v.asm, "ret")
	return nil
}

func (v *CodegenVisitor) VisitArgs(ctx *parsing.ArgsContext) any {
	return nil
}

func (v *CodegenVisitor) VisitBlock(ctx *parsing.BlockContext) any {
	for _, s := range ctx.AllStmt() {
		v.Visit(s)
	}
	return nil
}

func (v *CodegenVisitor) VisitStmt(ctx *parsing.StmtContext) any {
	exprCtx := ctx.Expr()
	if exprCtx != nil {
		exprVal := v.Visit(ctx.Expr()).(string)
		v.asm = append(v.asm, fmt.Sprintf("mov rax, %s", exprVal))
	}
	return nil
}

func (v *CodegenVisitor) VisitExpr(ctx *parsing.ExprContext) any {
	if ctx.NUM() != nil {
		return ctx.NUM().GetText()
	}
	return "0"
}





