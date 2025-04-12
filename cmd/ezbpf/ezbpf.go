package ezbpf

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/fatih/color"
	"github.com/toastsandwich/ezbpf/cmd/context"
	"github.com/toastsandwich/ezbpf/compiler"
	"github.com/toastsandwich/ezbpf/grammar"
)

func Main(ctx context.Context) {
	var file string
	if f, ok := ctx.Get("input").(string); ok {
		file = f
	}

	fs, err := antlr.NewFileStream(file)
	if err != nil {
		color.Red(err.Error())
		return
	}
	lexer := grammar.NewezbpfLexer(fs)
	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := grammar.NewezbpfParser(ts)

	compiler := compiler.NewC()
	tree := parser.Prog()
	code := compiler.Compile(tree)
	ctx.Put("output", code)
}
