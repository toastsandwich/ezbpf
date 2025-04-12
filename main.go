package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
	"github.com/toastsandwich/ezbpf/cmd/context"
	"github.com/toastsandwich/ezbpf/cmd/ezbpf"
)

func main() {
	output := flag.String("o", "a", "ezbpf -o sample sample.ezbpf")
	flag.Parse()

	if len(flag.Args()) != 1 {
		color.Red("need an input file")
		return
	}
	input := flag.Arg(0)
	ctx := context.ContextWithValue("input", input)

	ezbpf.Main(ctx)

	code := ctx.Get("output").(string)
	f, err := os.Create(*output)
	if err != nil {
		color.Red(err.Error())
		return
	}
	_, err = f.Write([]byte(code))
	if err != nil {
		color.Red(err.Error())
		return
	}
}
