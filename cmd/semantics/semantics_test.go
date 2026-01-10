package semantics_test

import (
	"blunderbuss/cmd/parsing"
	"blunderbuss/cmd/semantics"
	"errors"
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

func boilerplateMain(input string) string {
	return fmt.Sprintf(`
		func main() int { 
			%s
			return 0;
		}`, input)
}

func TestAssign(t *testing.T) {
    input := boilerplateMain("int a = 10; int b = c; int c = \"aaa\"; a[1] = 1; int a = 1;")
    is := antlr.NewInputStream(input)
    lexer := parsing.NewBlunderbussLexer(is)
    tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    parser := parsing.NewBlunderbussParser(tokens)

    tree := parser.Program()
    visitor := semantics.NewBlunderbussVisitor()
    visitor.Visit(tree)

	if len(visitor.Errors) != 5 {
		t.Logf("error count: %d", len(visitor.Errors))
		t.Logf("%v", visitor.Errors)
		t.Fatalf("expected 4(5) errors")
	}

	var uve *semantics.UndefinedVariableError
	var tme *semantics.TypeMismatchError
	var ie *semantics.IndexError
	var ive *semantics.InitializedVariableError
	e1 := visitor.Errors[0]
	if errors.As(e1, &uve) {
        t.Logf("got expected UndefinedVariableError: %s", uve.Msg)
    } else {
        t.Fatalf("expected UndefinedVariableError, got %T: %v", e1, e1)
    }

	// helper error at [1] (undefined variable has type any)

	e2 := visitor.Errors[2]
	if errors.As(e2, &tme) {
        t.Logf("got expected TypeMismatchError: %s", tme.Msg)
    } else {
        t.Fatalf("expected TypeMismatchError, got %T: %v", e2, e2)
    }

	e3 := visitor.Errors[3]
	if errors.As(e3, &ie) {
        t.Logf("got expected IndexError: %s", ie.Msg)
    } else {
        t.Fatalf("expected IndexError, got %T: %v", e2, e2)
	}

	e4 := visitor.Errors[4]
	if errors.As(e4, &ive) {
        t.Logf("got expected InitializedVariableError: %s", ive.Msg)
    } else {
        t.Fatalf("expected InitializedVariableError, got %T: %v", e4, e4)
	}

}
