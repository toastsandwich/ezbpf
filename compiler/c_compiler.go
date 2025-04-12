package compiler

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fatih/color"
	"github.com/toastsandwich/ezbpf/grammar"
)

type C struct {
	*grammar.BaseezbpfVisitor
	indent  int
	Builder strings.Builder
}

func (c *C) Visit(tree antlr.ParseTree) any {
	color.Green("-----------------------\n")
	color.Yellow(fmt.Sprintf("visiting %T\n", tree))
	color.Blue(tree.GetText())
	color.Green("-----------------------\n")

	return tree.Accept(c)
}

func (c *C) PushIndent() {
	c.indent++
}

func (c *C) PopIndent() {
	c.indent--
}

func (c *C) Indent() string { return strings.Repeat("\t", c.indent) }

func (c *C) VisitType(ctx *grammar.TypeContext) any {
	switch {
	case ctx.IDENTIFIER() != nil:
		return "struct " + ctx.IDENTIFIER().GetText()
	default:
		return ctx.GetText()
	}
}

func (c *C) VisitAssign(ctx *grammar.AssignContext) any {
	return ctx.GetText()
}

func (c *C) VisitCompare(ctx *grammar.CompareContext) any {
	return ctx.GetText()
}

func (c *C) VisitBinLiteralExpression(ctx *grammar.BinLiteralExpressionContext) any {
	return ctx.GetText()
}

func (c *C) VisitIdentifierExpression(ctx *grammar.IdentifierExpressionContext) any {
	return ctx.GetText()
}

func (c *C) VisitBitorExpression(ctx *grammar.BitorExpressionContext) any {
	exprs := ctx.AllExpression()
	return exprs[0].GetText() + ctx.BIT_OR().GetText() + exprs[1].GetText()
}

func (c *C) VisitOctLiteralExpression(ctx *grammar.OctLiteralExpressionContext) any {
	return ctx.GetText()
}

func (c *C) VisitBitxorExpression(ctx *grammar.BitxorExpressionContext) any {
	exprs := ctx.AllExpression()

	return fmt.Sprintf("%s %s %s", exprs[0].GetText(), ctx.BIT_XOR().GetText(), exprs[1].GetText())
}

func (c *C) VisitStringLiteralExpression(ctx *grammar.StringLiteralExpressionContext) any {
	// return fmt.Sprintf("\"%s\"", ctx.GetText())
	return ctx.GetText()
}

func (c *C) VisitBitandExpression(ctx *grammar.BitandExpressionContext) any {
	exprs := ctx.AllExpression()
	return fmt.Sprintf("%s %s %s", exprs[0].GetText(), ctx.BIT_AND().GetText(), exprs[1].GetText())
}

func (c *C) VisitHexLiteralExpression(ctx *grammar.HexLiteralExpressionContext) any {
	return ctx.GetText()
}

func (c *C) VisitFunccallExpression(ctx *grammar.FunccallExpressionContext) any {
	ident := ctx.FuncCallExpression().IDENTIFIER().GetText()
	args := c.Visit(ctx.FuncCallExpression().Args()).(string)
	return fmt.Sprintf("%s(%s);", ident, args)
}

func (c *C) VisitStructFieldAssignExpression(ctx *grammar.StructFieldAssignContext) any {
	exprs := []string{}
	for _, e := range ctx.AllExpression() {
		expr := c.Visit(e).(string)
		exprs = append(exprs, expr)
	}
	return strings.Join(exprs, ",")
}

func (c *C) VisitStructinitExpression(ctx *grammar.StructinitExpressionContext) any {
	return "{" + c.Visit(ctx.StructInitExpression().StructFieldAssign()).(string) + "}"
}

func (c *C) VisitArg(ctx *grammar.ArgContext) any {
	//exprs
	expr := c.Visit(ctx.Expression()).(string)
	return expr
}

func (c *C) VisitArgs(ctx *grammar.ArgsContext) any {
	args := []string{}
	for _, a := range ctx.AllArg() {
		arg := c.Visit(a).(string)
		args = append(args, arg)
	}
	return strings.Join(args, ",")
}

func (c *C) VisitParam(ctx *grammar.ParamContext) any {
	type_ := c.Visit(ctx.Type_()).(string)
	name := ctx.IDENTIFIER().GetText()
	return fmt.Sprintf("%s %s", type_, name)
}

func (c *C) VisitParams(ctx *grammar.ParamsContext) any {
	params := []string{}
	for _, p := range ctx.AllParam() {
		params = append(params, c.Visit(p).(string))
	}
	return strings.Join(params, ",")
}

func (c *C) VisitAddExpression(ctx *grammar.AddExpressionContext) any {
	exprs := ctx.AllExpression()
	return fmt.Sprintf("%s + %s", exprs[0].GetText(), exprs[1].GetText())
}

func (c *C) VisitSubExpression(ctx *grammar.SubExpressionContext) any {
	exprs := ctx.AllExpression()
	return fmt.Sprintf("%s - %s", exprs[0].GetText(), exprs[1].GetText())
}

func (c *C) VisitMulExpression(ctx *grammar.MulExpressionContext) any {
	exprs := ctx.AllExpression()
	return fmt.Sprintf("%s * %s", exprs[0].GetText(), exprs[1].GetText())
}

func (c *C) VisitDivExpression(ctx *grammar.DivExpressionContext) any {
	exprs := ctx.AllExpression()
	return fmt.Sprintf("%s / %s", exprs[0].GetText(), exprs[1].GetText())
}

func (c *C) VisitModExpression(ctx *grammar.ModExpressionContext) any {
	exprs := ctx.AllExpression()
	mod := "%"
	return fmt.Sprintf("%s %s %s", exprs[0].GetText(), mod, exprs[1].GetText())
}

func (c *C) VisitDotExpression(ctx *grammar.DotExpressionContext) any {
	expr := ctx.Expression().GetText()
	idnt := ctx.IDENTIFIER().GetText()
	return expr + "." + idnt
}

func (c *C) VisitLShiftExpression(ctx *grammar.LshiftExpressionContext) any {
	return ctx.GetText()
}

func (c *C) VisitRshiftExpression(ctx *grammar.RshiftExpressionContext) any {
	return ctx.GetText()
}

func (c *C) VisitVarDeclStmt(ctx *grammar.VarDeclStmtContext) any {
	// var name: type = expression ;
	ident := ctx.IDENTIFIER().GetText()
	type_ := c.Visit(ctx.Type_()).(string)
	expr := c.Visit(ctx.Expression()).(string)

	return fmt.Sprintf("%s %s = %s;", type_, ident, expr)
}

func (c *C) VisitVarInitStmt(ctx *grammar.VarInitStmtContext) any {
	// var name:type ;
	ident := ctx.IDENTIFIER().GetText()
	type_ := c.Visit(ctx.Type_()).(string)
	return fmt.Sprintf("%s %s;", type_, ident)
}

func (c *C) VisitConstDeclStmt(ctx *grammar.ConstDeclStmtContext) any {
	// const name: type = expression;
	ident := ctx.IDENTIFIER().GetText()
	type_ := c.Visit(ctx.Type_()).(string)
	expr := c.Visit(ctx.Expression()).(string)

	return fmt.Sprintf("const %s %s = %s;", type_, ident, expr)
}

func (c *C) VisitMapType(ctx *grammar.MapTypeContext) any {
	return ctx.GetText()
}

func (c *C) VisitMapDeclStmt(ctx *grammar.MapDeclStmtContext) any {
	// map name: map_type<type, type, max_entries>;
	ident := ctx.IDENTIFIER().GetText()
	mapType := c.Visit(ctx.MapType()).(string)
	kType_ := c.Visit(ctx.Type_(0)).(string)
	vType_ := c.Visit(ctx.Type_(1)).(string)
	maxEntries := ctx.DEC_LITERAL().GetText()

	return fmt.Sprintf(`
struct {
	__uint(type, %s);
	__type(key, %s);
	__type(value, %s);
	__uint(max_entries, %s);
} %s .SEC("maps");
`,
		mapType,
		kType_,
		vType_,
		maxEntries,
		ident,
	)
}

func (c *C) VisitStructDataMemStmt(ctx *grammar.StructDataMemStmtContext) any {
	ident := ctx.IDENTIFIER().GetText()
	type_ := c.Visit(ctx.Type_()).(string)
	return fmt.Sprintf("%s %s;", type_, ident)
}

func (c *C) VisitStructDeclStmt(ctx *grammar.StructDeclStmtContext) any {
	ident := ctx.IDENTIFIER().GetText()
	sdms := []string{}

	for _, s := range ctx.AllStructDataMemStmt() {
		sdm := c.Visit(s).(string)
		sdms = append(sdms, sdm)
	}
	return fmt.Sprintf(`
struct  %s {
	%s	
};			
`,
		ident,
		strings.Join(sdms, "\n"),
	)
}

func (c *C) VisitIfStmt(ctx *grammar.IfStmtContext) any {
	expr := c.Visit(ctx.Expression()).(string)
	stmts := c.Visit(ctx.Stmts()).(string)
	return fmt.Sprintf(`
if (%s) {
	%s
}
`,
		expr,
		stmts)
}

func (c *C) VisitReturnStmt(ctx *grammar.ReturnStmtContext) any {
	expr := c.Visit(ctx.Expression()).(string)
	return fmt.Sprintf("return %s;", expr)
}

func (c *C) VisitStmt(ctx *grammar.StmtContext) any {
	switch {
	case ctx.VarDeclStmt() != nil:
		return c.Visit(ctx.VarDeclStmt())
	case ctx.VarInitStmt() != nil:
		return c.Visit(ctx.VarInitStmt())
	case ctx.ConstDeclStmt() != nil:
		return c.Visit(ctx.ConstDeclStmt())
	case ctx.MapDeclStmt() != nil:
		return c.Visit(ctx.MapDeclStmt())
	case ctx.StructDeclStmt() != nil:
		return c.Visit(ctx.StructDeclStmt())
	case ctx.IfStmt() != nil:
		return c.Visit(ctx.IfStmt())
	case ctx.ReturnStmt() != nil:
		return c.Visit(ctx.ReturnStmt())
	case ctx.FuncDeclStmt() != nil:
		return c.Visit(ctx.FuncDeclStmt())
	default:
		color.Red("returning nil")
		return nil
	}
}

func (c *C) VisitStmts(ctx *grammar.StmtsContext) any {
	stmts := []string{}
	for _, s := range ctx.AllStmt() {
		if stmt, ok := c.Visit(s).(string); ok {
			stmts = append(stmts, stmt)
		}
	}
	return strings.Join(stmts, "\n")
}

func (c *C) VisitFuncDeclStmt(ctx *grammar.FuncDeclStmtContext) any {
	var params string
	fname := ctx.IDENTIFIER().GetText() // function name
	if ctx.Params() != nil {
		params = c.Visit(ctx.Params()).(string) // params
	}
	retType := c.Visit(ctx.Type_()).(string) // return type
	var stmts string
	if val, ok := c.Visit(ctx.Stmts()).(string); ok {
		stmts = val
	}
	color.Yellow(fmt.Sprintf("ctx.Stmts => %s\n", ctx.Stmts().GetText()))
	// stmts := c.Visit(ctx.Stmts()).(string)
	return fmt.Sprintf(`
%s %s(%s) {
	%s
}
		`,
		retType,
		fname,
		params,
		stmts,
	)
}

func (c *C) VisitProg(ctx *grammar.ProgContext) any {
	for _, sd := range ctx.AllStructDeclStmt() {
		s := c.Visit(sd).(string)
		c.Builder.WriteString(s)
	}
	for _, md := range ctx.AllMapDeclStmt() {
		m := c.Visit(md).(string)
		c.Builder.WriteString(m)
	}
	for _, fd := range ctx.AllFuncDeclStmt() {
		f := c.Visit(fd).(string)
		c.Builder.WriteString(f)
	}
	return c.Builder.String()
}

func NewC() *C {
	return &C{
		BaseezbpfVisitor: &grammar.BaseezbpfVisitor{},
		indent:           0,
		Builder:          strings.Builder{},
	}
}

func (c *C) Compile(prog grammar.IProgContext) string {
	defer func() {
		if r := recover(); r != nil {
			c := fmt.Sprintln(c.Builder.String())
			color.Green(c)
			stk := debug.Stack()
			fmt.Println(r)
			color.Red(string(stk))
		}
	}()
	c.LoadHeaders()
	if code, ok := c.Visit(prog).(string); ok {
		return code
	}
	return "empty"
}

func (c *C) LoadHeaders() {
	c.Builder.WriteString("#include <linux/bpf.h>\n")
	c.Builder.WriteString("#include <linux/if_ether.h>\n")
	c.Builder.WriteString("#include <linux/ip.h>\n")
	c.Builder.WriteString("#include <linux/tcp.h>\n")
	c.Builder.WriteString("#include <linux/udp.h>\n")
	c.Builder.WriteString("#include <bpf/bpf_helpers.h>\n")
}
