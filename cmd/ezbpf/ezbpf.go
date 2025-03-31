package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/toastsandwich/ezbpf/grammar"
)

func main() {
	i := "var i: int = 10;"

	input := antlr.NewInputStream(i)
	lexer := grammar.NewezbpfLexer(input)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewezbpfParser(tokens)

	tree := parser.CompilationUnit()
	listener := NewPrintListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}
