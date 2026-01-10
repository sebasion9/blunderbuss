package semantics

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type SemanticError struct {
	Msg string
	Ctx antlr.ParserRuleContext
}

func(e *SemanticError) Error() string {
	token := e.Ctx.GetStart()
	return fmt.Sprintf("Semantic error at line %d:%d: %s", 
		token.GetLine(), token.GetColumn(), e.Msg)
}

func NewSemanticError(msg string, ctx antlr.ParserRuleContext) *SemanticError {
	return &SemanticError{Msg: msg, Ctx: ctx}
}


type SyntaxError struct {
    Msg  string
    Text string
    Line int
    Col  int
}

type CustomErrorListener struct {
    *antlr.DefaultErrorListener
    Errors []SyntaxError
}

func (l *CustomErrorListener) SyntaxError(
    recognizer antlr.Recognizer,
    offendingSymbol interface{},
    line, column int,
    msg string,
    e antlr.RecognitionException,
) {
    tokenText := ""
    if t, ok := offendingSymbol.(antlr.Token); ok {
        tokenText = t.GetText()
    }

    l.Errors = append(l.Errors, SyntaxError{
        Msg:  msg,
        Text: tokenText,
        Line: line,
        Col:  column,
    })
}
