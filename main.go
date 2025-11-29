package main

import (
	"blunderbuss/semantics"
	"blunderbuss/parsing"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream("bbuss/array.bbuss")
	lexer := parsing.NewBlunderbussLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parsing.NewBlunderbussParser(tokens)

	tree := parser.Program()

	fmt.Println(tree.ToStringTree([]string{}, parser))

	visitor := semantics.NewBlunderbussVisitor()
	asm := visitor.Visit(tree).(string)
	fmt.Println(asm)


	os.WriteFile("target/out.asm", []byte(asm), 0644)
	
}

