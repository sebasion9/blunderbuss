package main

import (
	"blunderbuss/parsing"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream("test.bbuss")
	lexer := parsing.NewBlunderbussLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parsing.NewBlunderbussParser(tokens)

	tree := parser.Program()

	fmt.Println(tree.ToStringTree([]string{}, parser))


}

