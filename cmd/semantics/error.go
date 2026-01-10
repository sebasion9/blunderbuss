package semantics

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

// type SemanticError struct {
// 	Msg string
// 	Ctx antlr.ParserRuleContext
// }
//
// func(e *SemanticError) Error() string {
// 	token := e.Ctx.GetStart()
// 	return fmt.Sprintf("Semantic error at line %d:%d: %s", 
// 		token.GetLine(), token.GetColumn(), e.Msg)
// }
//
// func NewSemanticError(msg string, ctx antlr.ParserRuleContext) *SemanticError {
// 	return &SemanticError{Msg: msg, Ctx: ctx}
// }
//

type CompilerError interface {
	error
	Context() antlr.ParserRuleContext
}

type BaseCompilerError struct {
	Msg string
	Ctx antlr.ParserRuleContext
}

func (e *BaseCompilerError) Error() string {
	token := e.Ctx.GetStart()
	return fmt.Sprintf("Compiler error at line %d:%d: %s", 
		token.GetLine(), token.GetColumn(), e.Msg)
}

func (e *BaseCompilerError) Context() antlr.ParserRuleContext {
	return e.Ctx
}

type UndefinedVariableError struct {
	*BaseCompilerError
	VarName string
}

type InitializedVariableError struct {
	*BaseCompilerError
	Type string
	VarName string
}

type TypeMismatchError struct {
    *BaseCompilerError
    Expected string
    Got string
}

type IndexError struct {
	*BaseCompilerError
	Type string
	VarName string
}

type FnMismatchError struct {
	*BaseCompilerError
	FnName string
	Expected string
	Got string
	Reason string
}


func NewUndefinedVariableError(varName string, ctx antlr.ParserRuleContext) *UndefinedVariableError {
	return &UndefinedVariableError{
		BaseCompilerError: &BaseCompilerError{
			Msg: fmt.Sprintf("undefined variable: \"%s\"", varName),
			Ctx: ctx,
		},
		VarName: varName,
	}
}

func NewInitializedVariableError(type_, varName string, ctx antlr.ParserRuleContext) *InitializedVariableError{
	return &InitializedVariableError{
		BaseCompilerError: &BaseCompilerError{
			Msg: fmt.Sprintf("variable \"%s %s\" already initialized", type_, varName),
			Ctx: ctx,
		},
		VarName: varName,
		Type: type_,
	}
}

func NewTypeMismatchError(expected, got string, ctx antlr.ParserRuleContext) *TypeMismatchError {
    return &TypeMismatchError{
        BaseCompilerError: &BaseCompilerError {
			Msg: fmt.Sprintf("type mismatch: want: \"%s\" got: \"%s\"", expected, got),
            Ctx: ctx,
        },
        Expected: expected,
        Got:      got,
    }
}

func NewIndexError(type_, varName string, ctx antlr.ParserRuleContext) *IndexError{
    return &IndexError{
        BaseCompilerError: &BaseCompilerError {
			Msg: fmt.Sprintf("can't index that, this is not a pointer \"%s %s\"", type_, varName),
            Ctx: ctx,
        },
		Type: type_,
		VarName: varName,
    }
}

func NewFnMismatchError(fnName, expected, got, reason string, ctx antlr.ParserRuleContext) *FnMismatchError {
	switch reason {
		case "count":
		return &FnMismatchError{
			BaseCompilerError: &BaseCompilerError{
				Msg: fmt.Sprintf("mismatched function call argument count for function: \"%s\" want: \"%s\" got: \"%s\"", fnName, expected, got),
				Ctx: ctx,
			},
			Expected: expected,
			Got: got,
			Reason: reason,
		}
		case "type":
		return &FnMismatchError{
			BaseCompilerError: &BaseCompilerError{
				Msg: fmt.Sprintf("invalid argument type for function: \"%s\" want: \"%s\" got: \"%s\"", fnName, expected, got),
				Ctx: ctx,
			},
			Expected: expected,
			Got: got,
			Reason: reason,
		}
		default:
		return &FnMismatchError{
			BaseCompilerError: &BaseCompilerError{
				Msg: "Some type of function mismatch error occured",
				Ctx: ctx,
			},
			Expected: expected,
			Got: got,
			Reason: reason,
		}
	}

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
