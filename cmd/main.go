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
	
	parser.RemoveErrorListeners()
	
	listener := &semantics.CustomErrorListener{}
	parser.AddErrorListener(listener)

	tree := parser.Program()

	if len(listener.Errors) > 0 {
		fmt.Println("[ERR] Syntax errors found")
		for _, e := range listener.Errors {
			fmt.Printf("Syntax error at %d:%d near '%s': %s\n", 
				e.Line, e.Col, e.Text, e.Msg)
		}
		os.Exit(1)
	}


	//fmt.Println(tree.ToStringTree([]string{}, parser))

	visitor := semantics.NewBlunderbussVisitor()
	asm := visitor.Visit(tree).(string)
	if len(visitor.Errors) > 0 {
		fmt.Println("[ERR] Compilation failed with errors:")
		for _, e := range visitor.Errors {
			fmt.Println(e)
		}
		os.Exit(1)
	}

	fmt.Println(asm)


	err := os.WriteFile("target/out.asm", []byte(asm), 0644)
	if err != nil {
		fmt.Println("[ERR] failed to write asm to disk", err)
		os.Exit(1)
	}
	
	os.Exit(0)
}

