// Code generated from Blunderbuss.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // Blunderbuss
import "github.com/antlr4-go/antlr/v4"

type BaseBlunderbussVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBlunderbussVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBlunderbussVisitor) VisitFunc(ctx *FuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBlunderbussVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBlunderbussVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBlunderbussVisitor) VisitCall_args(ctx *Call_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBlunderbussVisitor) VisitFunc_call(ctx *Func_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBlunderbussVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBlunderbussVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBlunderbussVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBlunderbussVisitor) VisitEffect_block(ctx *Effect_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBlunderbussVisitor) VisitIf_stmt(ctx *If_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBlunderbussVisitor) VisitFor_stmt(ctx *For_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}
