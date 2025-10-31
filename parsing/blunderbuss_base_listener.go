// Code generated from Blunderbuss.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // Blunderbuss
import "github.com/antlr4-go/antlr/v4"

// BaseBlunderbussListener is a complete listener for a parse tree produced by BlunderbussParser.
type BaseBlunderbussListener struct{}

var _ BlunderbussListener = &BaseBlunderbussListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBlunderbussListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBlunderbussListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBlunderbussListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBlunderbussListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseBlunderbussListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseBlunderbussListener) ExitProgram(ctx *ProgramContext) {}

// EnterFunc is called when production func is entered.
func (s *BaseBlunderbussListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BaseBlunderbussListener) ExitFunc(ctx *FuncContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseBlunderbussListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseBlunderbussListener) ExitArgs(ctx *ArgsContext) {}

// EnterParam is called when production param is entered.
func (s *BaseBlunderbussListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseBlunderbussListener) ExitParam(ctx *ParamContext) {}

// EnterCall_args is called when production call_args is entered.
func (s *BaseBlunderbussListener) EnterCall_args(ctx *Call_argsContext) {}

// ExitCall_args is called when production call_args is exited.
func (s *BaseBlunderbussListener) ExitCall_args(ctx *Call_argsContext) {}

// EnterFunc_call is called when production func_call is entered.
func (s *BaseBlunderbussListener) EnterFunc_call(ctx *Func_callContext) {}

// ExitFunc_call is called when production func_call is exited.
func (s *BaseBlunderbussListener) ExitFunc_call(ctx *Func_callContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBlunderbussListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBlunderbussListener) ExitBlock(ctx *BlockContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseBlunderbussListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseBlunderbussListener) ExitExpr(ctx *ExprContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseBlunderbussListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseBlunderbussListener) ExitStmt(ctx *StmtContext) {}

// EnterEffect_block is called when production effect_block is entered.
func (s *BaseBlunderbussListener) EnterEffect_block(ctx *Effect_blockContext) {}

// ExitEffect_block is called when production effect_block is exited.
func (s *BaseBlunderbussListener) ExitEffect_block(ctx *Effect_blockContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BaseBlunderbussListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BaseBlunderbussListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterFor_stmt is called when production for_stmt is entered.
func (s *BaseBlunderbussListener) EnterFor_stmt(ctx *For_stmtContext) {}

// ExitFor_stmt is called when production for_stmt is exited.
func (s *BaseBlunderbussListener) ExitFor_stmt(ctx *For_stmtContext) {}
