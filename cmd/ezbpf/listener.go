package main

import (
	"fmt"
	"strings"

	parser "github.com/toastsandwich/ezbpf/grammar"

	"github.com/antlr4-go/antlr/v4"
)

// MyLangToCListener is a listener that generates C code for eBPF.
type MyLangToCListener struct {
	*parser.BaseezbpfListener

	output strings.Builder
	scopes []map[string]string // For simple variable tracking (you might need something more robust)
	errors []string            // Collect errors during translation
}

// NewMyLangToCListener creates a new MyLangToCListener.
func NewMyLangToCListener() *MyLangToCListener {
	return &MyLangToCListener{
		scopes: []map[string]string{make(map[string]string)},
		errors: []string{},
	}
}

// GetOutput returns the generated C code.
func (l *MyLangToCListener) GetOutput() string {
	return l.output.String()
}

// GetErrors returns the collected errors.
func (l *MyLangToCListener) GetErrors() []string {
	return l.errors
}

// addError adds an error message to the errors list.
func (l *MyLangToCListener) addError(message string, ctx antlr.ParserRuleContext) {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	l.errors = append(l.errors, fmt.Sprintf("line %d:%d: %s", line, column, message))
}

// enterScope creates a new scope.
func (l *MyLangToCListener) enterScope() {
	l.scopes = append(l.scopes, make(map[string]string))
}

// exitScope removes the current scope.
func (l *MyLangToCListener) exitScope() {
	l.scopes = l.scopes[:len(l.scopes)-1]
}

// declareVariable declares a variable in the current scope.
func (l *MyLangToCListener) declareVariable(name string, cType string, ctx antlr.ParserRuleContext) {
	currentScope := l.scopes[len(l.scopes)-1]
	if _, ok := currentScope[name]; ok {
		l.addError(fmt.Sprintf("variable '%s' already declared in this scope", name), ctx)
		return // Don't redeclare
	}
	currentScope[name] = cType
}

// lookupVariable finds the C type of a variable in the current or enclosing scopes.
func (l *MyLangToCListener) lookupVariable(name string, ctx antlr.ParserRuleContext) (string, bool) {
	for i := len(l.scopes) - 1; i >= 0; i-- {
		if cType, ok := l.scopes[i][name]; ok {
			return cType, true
		}
	}
	l.addError(fmt.Sprintf("variable '%s' not declared", name), ctx)
	return "", false
}

// VisitTerminal is called when the listener visits a terminal node.
func (l *MyLangToCListener) VisitTerminal(node antlr.TerminalNode) {
	// Handle simple identifier translations if needed
}

// EnterStructDefinition is called when entering a structDefinition rule.
func (l *MyLangToCListener) EnterStructDefinition(ctx *parser.StructDefinitionContext) {
	structName := ctx.Identifier().GetText()
	l.output.WriteString(fmt.Sprintf("struct %s {\n", structName))
}

// ExitStructDefinition is called when exiting a structDefinition rule.
func (l *MyLangToCListener) ExitStructDefinition(ctx *parser.StructDefinitionContext) {
	l.output.WriteString("};\n\n")
}

// EnterStructMember is called when entering a structMember rule.
func (l *MyLangToCListener) EnterStructMember(ctx *parser.StructMemberContext) {
	memberName := ctx.Identifier().GetText()
	myLangType := ctx.Type_().GetText()
	cType := l.translateType(myLangType, ctx)
	l.output.WriteString(fmt.Sprintf("\t%s %s;\n", cType, memberName))
}

// EnterBpfMapDefinition is called when entering a bpfMapDefinition rule.
func (l *MyLangToCListener) EnterBpfMapDefinition(ctx *parser.BpfMapDefinitionContext) {
	mapName := ctx.Identifier().GetText()
	l.output.WriteString(fmt.Sprintf("struct bpf_map_def SEC(\"maps/%s\") %s = {\n", mapName, mapName))
}

// ExitBpfMapDefinition is called when exiting a bpfMapDefinition rule.
func (l *MyLangToCListener) ExitBpfMapDefinition(ctx *parser.BpfMapDefinitionContext) {
	l.output.WriteString("};\n\n")
}

// EnterBpfMapOption is called when entering a bpfMapOption rule.
func (l *MyLangToCListener) EnterBpfMapOption(ctx *parser.BpfMapOptionContext) {
	option := ctx.GetChild(0).GetPayload().(string)
	value := ctx.GetChild(2).GetPayload().(string)

	switch option {
	case "type":
		l.output.WriteString(fmt.Sprintf("\t.type = BPF_MAP_%s,\n", strings.ToUpper(value)))
	case "key":
		cKeyType := l.translateType(value, ctx)
		l.output.WriteString(fmt.Sprintf("\t.key_size = sizeof(%s),\n", cKeyType))
	case "value":
		cValueType := l.translateType(value, ctx)
		l.output.WriteString(fmt.Sprintf("\t.value_size = sizeof(%s),\n", cValueType))
	case "max_entries":
		l.output.WriteString(fmt.Sprintf("\t.max_entries = %s,\n", value))
	}
}

// EnterGlobalVariableDeclaration is called when entering a globalVariableDeclaration rule.
func (l *MyLangToCListener) EnterGlobalVariableDeclaration(ctx *parser.GlobalVariableDeclarationContext) {
	varName := ctx.Identifier().GetText()
	myLangType := ctx.Type_().GetText()
	cType := l.translateType(myLangType, ctx)
	l.declareVariable(varName, cType, ctx) // Register global variable
	l.output.WriteString(fmt.Sprintf("%s %s", cType, varName))
	if ctx.EQUAL() != nil {
		// Basic expression handling for globals (can be expanded)
		l.output.WriteString(" = ")
		l.processExpression(ctx.Expression()) // Rely on expression handling
	}
	l.output.WriteString(";\n")
}

// EnterFunctionDefinition is called when entering a functionDefinition rule.
func (l *MyLangToCListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	secName := strings.Trim(ctx.STRING_LITERAL().GetText(), "\"")
	funcName := ctx.Identifier().GetText()
	returnTypeCtx := ctx.Type_()
	cReturnType := "int" // Default return type
	if returnTypeCtx != nil {
		cReturnType = l.translateType(returnTypeCtx.GetText(), ctx)
	}

	l.output.WriteString(fmt.Sprintf("SEC(\"%s\")\n%s %s(", secName, cReturnType, funcName))
	l.enterScope() // New scope for function

	params := ctx.ParameterList()
	if params != nil {
		for i, param := range params.AllParameter() {
			paramName := param.Identifier().GetText()
			paramType := param.Type_().GetText()
			cParamType := l.translateType(paramType, param)
			l.declareVariable(paramName, cParamType, param) // Register parameter
			l.output.WriteString(fmt.Sprintf("%s %s", cParamType, paramName))
			if i < len(params.AllParameter())-1 {
				l.output.WriteString(", ")
			}
		}
	}
	l.output.WriteString(") {\n")
}

// ExitFunctionDefinition is called when exiting a functionDefinition rule.
func (l *MyLangToCListener) ExitFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	l.exitScope()
	l.output.WriteString("}\n\n")
}

// EnterVariableDeclarationStatement is called when entering a variableDeclarationStatement rule.
func (l *MyLangToCListener) EnterVariableDeclarationStatement(ctx *parser.VariableDeclarationStatementContext) {
	varName := ctx.Identifier().GetText()
	myLangType := ctx.Type_().GetText()
	cType := l.translateType(myLangType, ctx)
	l.declareVariable(varName, cType, ctx) // Register local variable
	l.output.WriteString(fmt.Sprintf("\t%s %s", cType, varName))
	if ctx.EQUAL() != nil {
		l.output.WriteString(" = ")
	}
}

// ExitVariableDeclarationStatement is called when exiting a variableDeclarationStatement rule.
func (l *MyLangToCListener) ExitVariableDeclarationStatement(ctx *parser.VariableDeclarationStatementContext) {
	l.output.WriteString(";\n")
}

// EnterAssignmentStatement is called when entering an assignmentStatement rule.
func (l *MyLangToCListener) EnterAssignmentStatement(ctx *parser.AssignmentStatementContext) {
	varName := ctx.Identifier().GetText()
	if _, ok := l.lookupVariable(varName, ctx); !ok {
		// Error already added in lookupVariable
		return
	}
	l.output.WriteString(fmt.Sprintf("\t%s = ", varName))
}

// ExitAssignmentStatement is called when exiting an assignmentStatement rule.
func (l *MyLangToCListener) ExitAssignmentStatement(ctx *parser.AssignmentStatementContext) {
	l.output.WriteString(";\n")
}

// EnterReturnStatement is called when entering a returnStatement rule.
func (l *MyLangToCListener) EnterReturnStatement(ctx *parser.ReturnStatementContext) {
	l.output.WriteString("\treturn ")
}

// ExitReturnStatement is called when exiting a returnStatement rule.
func (l *MyLangToCListener) ExitReturnStatement(ctx *parser.ReturnStatementContext) {
	l.output.WriteString(";\n")
}

// EnterIfStatement is called when entering an ifStatement rule.
func (l *MyLangToCListener) EnterIfStatement(ctx *parser.IfStatementContext) {
	l.output.WriteString("\tif (")
}

// ExitIfStatement is called when exiting an ifStatement rule.
func (l *MyLangToCListener) ExitIfStatement(ctx *parser.IfStatementContext) {
	l.output.WriteString(")\n")
	if ctx.ElseBlock() != nil {
		l.output.WriteString("\telse\n")
	}
}

// EnterWhileStatement is called when entering a whileStatement rule.
func (l *MyLangToCListener) EnterWhileStatement(ctx *parser.WhileStatementContext) {
	l.output.WriteString("\twhile (")
}

// ExitWhileStatement is called when exiting a whileStatement rule.
func (l *MyLangToCListener) ExitWhileStatement(ctx *parser.WhileStatementContext) {
	l.output.WriteString(")\n")
}

// EnterForStatement is called when entering a forStatement rule.
func (l *MyLangToCListener) EnterForStatement(ctx *parser.ForStatementContext) {
	l.output.WriteString("\tfor (")
	// Initialization will be handled in EnterSimpleStatement
}

// ExitForStatement is called when exiting a forStatement rule.
func (l *MyLangToCListener) ExitForStatement(ctx *parser.ForStatementContext) {
	l.output.WriteString(")\n")
}

// EnterSimpleStatement is called when entering a simpleStatement rule within a for loop.
func (l *MyLangToCListener) EnterSimpleStatement(ctx *parser.SimpleStatementContext) {
	// For now, assuming assignment or variable declaration
	if assignCtx := ctx.AssignmentStatement(); assignCtx != nil {
		l.EnterAssignmentStatement(assignCtx)
	} else if varDeclCtx := ctx.VariableDeclarationStatement(); varDeclCtx != nil {
		l.EnterVariableDeclarationStatement(varDeclCtx)
	} else if funcCallCtx := ctx.FunctionCallStatement(); funcCallCtx != nil {
		l.EnterFunctionCallStatement(funcCallCtx)
	}
}

// ExitSimpleStatement is called when exiting a simpleStatement rule within a for loop.
func (l *MyLangToCListener) ExitSimpleStatement(ctx *parser.SimpleStatementContext) {
	if assignCtx := ctx.AssignmentStatement(); assignCtx != nil {
		l.ExitAssignmentStatement(assignCtx)
	} else if varDeclCtx := ctx.VariableDeclarationStatement(); varDeclCtx != nil {
		l.ExitVariableDeclarationStatement(varDeclCtx)
	} else if funcCallCtx := ctx.FunctionCallStatement(); funcCallCtx != nil {
		l.ExitFunctionCallStatement(funcCallCtx)
	}
	if ctx.Parent().(*parser.ForStatementContext) != nil && ctx != ctx.Parent().(*parser.ForStatementContext).GetStop() {
		l.output.WriteString(";")
	}
}

// EnterFunctionCallStatement is called when entering a functionCallStatement rule.
func (l *MyLangToCListener) EnterFunctionCallStatement(ctx *parser.FunctionCallStatementContext) {
	funcName := ctx.Identifier().GetText()
	l.output.WriteString(fmt.Sprintf("\t%s(", funcName))
}

// ExitFunctionCallStatement is called when exiting a functionCallStatement rule.
func (l *MyLangToCListener) ExitFunctionCallStatement(ctx *parser.FunctionCallStatementContext) {
	l.output.WriteString(");\n")
}

// EnterArgumentList is called when entering an argumentList rule.
func (l *MyLangToCListener) EnterArgumentList(ctx *parser.ArgumentListContext) {
	// Arguments will be handled by processExpression
}

// ExitArgumentList is called when exiting an argumentList rule.
func (l *MyLangToCListener) ExitArgumentList(ctx *parser.ArgumentListContext) {
}

func (l *MyLangToCListener) processExpression(ctx *parser.ExpressionContext) {
	// Handle binary operations
	if len(ctx.AllPrimaryExpression()) > 1 && len(ctx.AllOperator()) > 0 {
		// This is a very basic handling, needs proper operator precedence
		l.processPrimaryExpression(ctx.PrimaryExpression(0))
		for i, opCtx := range ctx.AllOperator() {
			l.output.WriteString(fmt.Sprintf(" %s ", opCtx.GetText()))
			l.processPrimaryExpression(ctx.PrimaryExpression(i + 1))
		}
	} else if len(ctx.AllPrimaryExpression()) == 1 {
		l.processPrimaryExpression(ctx.PrimaryExpression(0))
	}
}

// EnterExpression is called when entering an expression rule.
func (l *MyLangToCListener) EnterExpression(ctx *parser.ExpressionContext) {
	l.processExpression(ctx)
}

// ExitExpression is called when exiting an expression rule.
func (l *MyLangToCListener) ExitExpression(ctx *parser.ExpressionContext) {
	// Nothing specific to do here for now
}

// processPrimaryExpression is called to process primary expressions.
func (l *MyLangToCListener) processPrimaryExpression(ctx *parser.PrimaryExpressionContext) {
	if id := ctx.Identifier(); id != nil {
		l.output.WriteString(id.GetText())
	} else if intLit := ctx.INTEGER_LITERAL(); intLit != nil {
		l.output.WriteString(intLit.GetText())
	} else if charLit := ctx.CHAR_LITERAL(); charLit != nil {
		// Need to handle escaping properly in a real compiler
		l.output.WriteString(charLit.GetText())
	} else if stringLit := ctx.STRING_LITERAL(); stringLit != nil {
		l.output.WriteString(stringLit.GetText())
	} else if funcCall := ctx.FunctionCallStatement(); funcCall != nil {
		l.EnterFunctionCallStatement(funcCall)
		// Process arguments
		if funcCall.ArgumentList() != nil {
			for i, arg := range funcCall.ArgumentList().AllExpression() {
				l.processExpression(arg)
				if i < len(funcCall.ArgumentList().AllExpression())-1 {
					l.output.WriteString(", ")
				}
			}
		}
		l.ExitFunctionCallStatement(funcCall)

	} else if mapOp := ctx.BpfMapOperationStatement(); mapOp != nil {
		l.EnterBpfMapOperationStatement(mapOp)
		l.ExitBpfMapOperationStatement(mapOp)
	} else if ctx.LPAREN() != nil {
		l.output.WriteString("(")
		l.processExpression(ctx.Parent().(*parser.ExpressionContext)) // Process the expression within the parentheses
		l.output.WriteString(")")
	}
}

// EnterPrimaryExpression is called when entering a primaryExpression rule.
func (l *MyLangToCListener) EnterPrimaryExpression(ctx *parser.PrimaryExpressionContext) {
	l.processPrimaryExpression(ctx)
}

// ExitPrimaryExpression is called when exiting a primaryExpression rule.
func (l *MyLangToCListener) ExitPrimaryExpression(ctx *parser.PrimaryExpressionContext) {
	// Nothing here
}

// EnterBpfMapOperationStatement is called when entering a bpfMapOperationStatement rule.
func (l *MyLangToCListener) EnterBpfMapOperationStatement(ctx *parser.BpfMapOperationStatementContext) {
	mapName := ctx.Identifier(0).GetText()
	l.output.WriteString(mapName)
	l.output.WriteString("[")
	l.processExpression(ctx.Expression(0))
	l.output.WriteString("]")
	if len(ctx.AllExpression()) > 1 && ctx.EQUAL() != nil { // Assignment
		l.output.WriteString(" = ")
		l.processExpression(ctx.Expression(1))
	} else if ctx.DELETE() != nil {
		l.output.WriteString(fmt.Sprintf(" = 0; // Suggesting a delete operation for map '%s'", mapName))
		// Real delete would likely involve a BPF helper function:  bpf_map_delete_elem(&%s, &key);
	}
}

// ExitBpfMapOperationStatement is called when exiting a bpfMapOperationStatement rule.
func (l *MyLangToCListener) ExitBpfMapOperationStatement(ctx *parser.BpfMapOperationStatementContext) {
	if ctx.EQUAL() != nil || ctx.DELETE() != nil {
		l.output.WriteString(";\n")
	}
}

// translateType maps your language's types to C types for eBPF.
func (l *MyLangToCListener) translateType(myLangType string, ctx antlr.ParserRuleContext) string {
	switch myLangType {
	case "int":
		return "int"
	case "char":
		return "char"
	case "long":
		return "long"
	case "xdp.Context":
		return "struct xdp_md*" //  Correct C type for xdp.Context in eBPF
	case "*int":
		return "int*"
	case "*char":
		return "char*"
	case "*long":
		return "long*"
	default:
		// Check if it's a previously defined struct
		if _, ok := l.lookupVariable(myLangType, ctx); ok {
			return fmt.Sprintf("struct %s", myLangType)
		}
		l.addError(fmt.Sprintf("unknown type '%s'", myLangType), ctx)
		return "int" // Default to int to avoid cascading errors, the error will be caught and shown.
	}
}
