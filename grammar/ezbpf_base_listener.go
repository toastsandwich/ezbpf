// Code generated from ezbpf.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // ezbpf
import "github.com/antlr4-go/antlr/v4"

// BaseezbpfListener is a complete listener for a parse tree produced by ezbpfParser.
type BaseezbpfListener struct{}

var _ ezbpfListener = &BaseezbpfListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseezbpfListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseezbpfListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseezbpfListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseezbpfListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseezbpfListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseezbpfListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterStructDefinition is called when production structDefinition is entered.
func (s *BaseezbpfListener) EnterStructDefinition(ctx *StructDefinitionContext) {}

// ExitStructDefinition is called when production structDefinition is exited.
func (s *BaseezbpfListener) ExitStructDefinition(ctx *StructDefinitionContext) {}

// EnterStructMember is called when production structMember is entered.
func (s *BaseezbpfListener) EnterStructMember(ctx *StructMemberContext) {}

// ExitStructMember is called when production structMember is exited.
func (s *BaseezbpfListener) ExitStructMember(ctx *StructMemberContext) {}

// EnterBpfMapDefinition is called when production bpfMapDefinition is entered.
func (s *BaseezbpfListener) EnterBpfMapDefinition(ctx *BpfMapDefinitionContext) {}

// ExitBpfMapDefinition is called when production bpfMapDefinition is exited.
func (s *BaseezbpfListener) ExitBpfMapDefinition(ctx *BpfMapDefinitionContext) {}

// EnterBpfMapOption is called when production bpfMapOption is entered.
func (s *BaseezbpfListener) EnterBpfMapOption(ctx *BpfMapOptionContext) {}

// ExitBpfMapOption is called when production bpfMapOption is exited.
func (s *BaseezbpfListener) ExitBpfMapOption(ctx *BpfMapOptionContext) {}

// EnterGlobalVariableDeclaration is called when production globalVariableDeclaration is entered.
func (s *BaseezbpfListener) EnterGlobalVariableDeclaration(ctx *GlobalVariableDeclarationContext) {}

// ExitGlobalVariableDeclaration is called when production globalVariableDeclaration is exited.
func (s *BaseezbpfListener) ExitGlobalVariableDeclaration(ctx *GlobalVariableDeclarationContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseezbpfListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseezbpfListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseezbpfListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseezbpfListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseezbpfListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseezbpfListener) ExitParameter(ctx *ParameterContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseezbpfListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseezbpfListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseezbpfListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseezbpfListener) ExitStatement(ctx *StatementContext) {}

// EnterVariableDeclarationStatement is called when production variableDeclarationStatement is entered.
func (s *BaseezbpfListener) EnterVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) {
}

// ExitVariableDeclarationStatement is called when production variableDeclarationStatement is exited.
func (s *BaseezbpfListener) ExitVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) {
}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseezbpfListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseezbpfListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseezbpfListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseezbpfListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseezbpfListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseezbpfListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseezbpfListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseezbpfListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseezbpfListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseezbpfListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterSimpleStatement is called when production simpleStatement is entered.
func (s *BaseezbpfListener) EnterSimpleStatement(ctx *SimpleStatementContext) {}

// ExitSimpleStatement is called when production simpleStatement is exited.
func (s *BaseezbpfListener) ExitSimpleStatement(ctx *SimpleStatementContext) {}

// EnterFunctionCallStatement is called when production functionCallStatement is entered.
func (s *BaseezbpfListener) EnterFunctionCallStatement(ctx *FunctionCallStatementContext) {}

// ExitFunctionCallStatement is called when production functionCallStatement is exited.
func (s *BaseezbpfListener) ExitFunctionCallStatement(ctx *FunctionCallStatementContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseezbpfListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseezbpfListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterBpfMapOperationStatement is called when production bpfMapOperationStatement is entered.
func (s *BaseezbpfListener) EnterBpfMapOperationStatement(ctx *BpfMapOperationStatementContext) {}

// ExitBpfMapOperationStatement is called when production bpfMapOperationStatement is exited.
func (s *BaseezbpfListener) ExitBpfMapOperationStatement(ctx *BpfMapOperationStatementContext) {}

// EnterType is called when production type is entered.
func (s *BaseezbpfListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseezbpfListener) ExitType(ctx *TypeContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseezbpfListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseezbpfListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseezbpfListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseezbpfListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseezbpfListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseezbpfListener) ExitOperator(ctx *OperatorContext) {}
