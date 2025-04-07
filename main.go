package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
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

	ctx := ezbpf.NewContext()
	ctx.Put("input", input)

	ezbpf.Main(ctx)

	out := ctx.Get("out")

	file, err := os.Create(*output)
	if err != nil {
		color.Red(err.Error())
		return
	}
	_, err = file.Write([]byte(out))
	if err != nil {
		color.Red(err.Error())
		return
	}
}
