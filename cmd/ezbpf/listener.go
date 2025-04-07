package ezbpf

import (
	"fmt"
	"strings"

	"github.com/toastsandwich/ezbpf/grammar"
)

// this will be used for development purposes
type DevListener struct {
	*grammar.BaseezbpfListener

	Errors []string
}

func NewDevListener() *DevListener {
	return &DevListener{&grammar.BaseezbpfListener{}, []string{}}
}

// func (d *DevListener) EnterExpression(ctx *grammar.ExpressionContext) {
// 	// fmt.Println("entering expression")
// 	text := ctx.GetText()
// 	fmt.Println(text)
// }

func (d *DevListener) EnterMapDeclStmt(ctx *grammar.MapDeclStmtContext) {
	name := ctx.IDENTIFIER().GetText()
	mapType := ctx.MapType().GetText()
	keyType, valType := ctx.Type_(0).GetText(), ctx.Type_(1).GetText()
	maxEntries := ctx.DEC_LITERAL().GetText()

	fmt.Printf("map %s <type: %s key: %s val: %s, max_entries: %s>\n",
		name,
		mapType,
		keyType,
		valType,
		maxEntries,
	)
}

func (d *DevListener) EnterVarDeclStmt(ctx *grammar.VarDeclStmtContext) {
	name := ctx.IDENTIFIER().GetText()
	_type := ctx.Type_().GetText()
	val := ctx.Expression().GetText()

	fmt.Printf("%s %s = %s\n", name, _type, val)
}

func (d *DevListener) EnterFuncDeclStmt(ctx *grammar.FuncDeclStmtContext) {
	fname := ctx.IDENTIFIER().GetText()
	p := ctx.Params()
	var params string
	if p != nil {
		params = p.GetText()
	}

	fmt.Println("entering func", fname, "with params:", params)
}

func (d *DevListener) ExitFuncDeclStmt(ctx *grammar.FuncDeclStmtContext) {
	fname := ctx.IDENTIFIER().GetText()
	fmt.Println("exiting func", fname)
}

func (d *DevListener) EnterReturnStmt(ctx *grammar.ReturnStmtContext) {
	fmt.Println(ctx.GetText())
	val := ctx.Expression().GetText()
	fmt.Println("return", val)
}

func (d *DevListener) EnterStructDeclStmt(ctx *grammar.StructDeclStmtContext) {
	fmt.Println(ctx.GetText())
}

func (d *DevListener) EnterStructFieldAssignStmt(ctx *grammar.StructFieldAssignContext) {
	fmt.Println(ctx.GetText())
}

// this will be used to compile
//
// TODO Implement a func call system that will
// call the functions of generator on the context
type CompilerListener struct {
	*grammar.BaseezbpfListener
	Ctx     Context
	Builder strings.Builder
}

func NewCompilerListener(ctx Context) *CompilerListener {
	return &CompilerListener{
		&grammar.BaseezbpfListener{},
		ctx,
		strings.Builder{},
	}
}

func (cl *CompilerListener) EnterType(ctx *grammar.TypeContext) {

}

func (cl CompilerListener) EnterVarInitStmt(ctx *grammar.VarInitStmtContext) {

}
