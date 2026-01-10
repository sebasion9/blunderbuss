package semantics_test

import (
	"blunderbuss/cmd/parsing"
	"blunderbuss/cmd/semantics"
	"testing"
	"github.com/antlr4-go/antlr/v4"
	"fmt"
)

func boilerplateMain(input string) string {
	return fmt.Sprintf(`
		func main() int { 
			%s
			return 0;
		}`, input)
}

func TestAssign(t *testing.T) {
    input := boilerplateMain("int a = 10; int a = a; int b = c; int c = \"aaa\" ")
    is := antlr.NewInputStream(input)
    lexer := parsing.NewBlunderbussLexer(is)
    tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    parser := parsing.NewBlunderbussParser(tokens)

    tree := parser.Program()
    visitor := semantics.NewBlunderbussVisitor()
    visitor.Visit(tree)

	// err := visitor.Errors[0]
 //    if err == nil {
 //        t.Fatalf("expected undefined variable error, got nil")
 //    }
}
