// Code generated from Blunderbuss.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // Blunderbuss
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BlunderbussParser.
type BlunderbussVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BlunderbussParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by BlunderbussParser#extern.
	VisitExtern(ctx *ExternContext) interface{}

	// Visit a parse tree produced by BlunderbussParser#func.
	VisitFunc(ctx *FuncContext) interface{}

	// Visit a parse tree produced by BlunderbussParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by BlunderbussParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by BlunderbussParser#call_args.
	VisitCall_args(ctx *Call_argsContext) interface{}

	// Visit a parse tree produced by BlunderbussParser#func_call.
	VisitFunc_call(ctx *Func_callContext) interface{}

	// Visit a parse tree produced by BlunderbussParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by BlunderbussParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by BlunderbussParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by BlunderbussParser#if_stmt.
	VisitIf_stmt(ctx *If_stmtContext) interface{}

	// Visit a parse tree produced by BlunderbussParser#for_stmt.
	VisitFor_stmt(ctx *For_stmtContext) interface{}
}
