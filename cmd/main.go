package main

import (
	"blunderbuss/cmd/semantics"
	"blunderbuss/cmd/parsing"
	"fmt"
	"os"
	"io"

	"github.com/antlr4-go/antlr/v4"
)

func main() {

	stdin, _ := io.ReadAll(os.Stdin)
	input := antlr.NewInputStream(string(stdin))

	lexer := parsing.NewBlunderbussLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parsing.NewBlunderbussParser(tokens)

	tree := parser.Program()

	fmt.Println(tree.ToStringTree([]string{}, parser))

	visitor := semantics.NewBlunderbussVisitor()
	asm := visitor.Visit(tree).(string)
	fmt.Println(asm)


	err := os.WriteFile("target/out.asm", []byte(asm), 0644)
	if err != nil {
		fmt.Println("[ERR] failed to write asm to disk", err)
	}
	
}

