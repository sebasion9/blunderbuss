package main

import (
	"blunderbuss/parsing"
	"blunderbuss/codegen"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream("bbuss/return.bbuss")
	lexer := parsing.NewBlunderbussLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parsing.NewBlunderbussParser(tokens)

	tree := parser.Program()

	fmt.Println(tree.ToStringTree([]string{}, parser))

	visitor := codegen.NewBlunderbussVisitor()
	asm := visitor.Visit(tree)//.(string)
	fmt.Println(asm)
}

