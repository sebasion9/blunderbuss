// Code generated from Blunderbuss.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // Blunderbuss
import "github.com/antlr4-go/antlr/v4"

// BlunderbussListener is a complete listener for a parse tree produced by BlunderbussParser.
type BlunderbussListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterFunc is called when entering the func production.
	EnterFunc(c *FuncContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterCall_args is called when entering the call_args production.
	EnterCall_args(c *Call_argsContext)

	// EnterFunc_call is called when entering the func_call production.
	EnterFunc_call(c *Func_callContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterEffect_block is called when entering the effect_block production.
	EnterEffect_block(c *Effect_blockContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterFor_stmt is called when entering the for_stmt production.
	EnterFor_stmt(c *For_stmtContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitFunc is called when exiting the func production.
	ExitFunc(c *FuncContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitCall_args is called when exiting the call_args production.
	ExitCall_args(c *Call_argsContext)

	// ExitFunc_call is called when exiting the func_call production.
	ExitFunc_call(c *Func_callContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitEffect_block is called when exiting the effect_block production.
	ExitEffect_block(c *Effect_blockContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitFor_stmt is called when exiting the for_stmt production.
	ExitFor_stmt(c *For_stmtContext)
}
