// Code generated from ezbpf.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // ezbpf
import "github.com/antlr4-go/antlr/v4"

// ezbpfListener is a complete listener for a parse tree produced by ezbpfParser.
type ezbpfListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterStructDefinition is called when entering the structDefinition production.
	EnterStructDefinition(c *StructDefinitionContext)

	// EnterStructMember is called when entering the structMember production.
	EnterStructMember(c *StructMemberContext)

	// EnterBpfMapDefinition is called when entering the bpfMapDefinition production.
	EnterBpfMapDefinition(c *BpfMapDefinitionContext)

	// EnterBpfMapOption is called when entering the bpfMapOption production.
	EnterBpfMapOption(c *BpfMapOptionContext)

	// EnterGlobalVariableDeclaration is called when entering the globalVariableDeclaration production.
	EnterGlobalVariableDeclaration(c *GlobalVariableDeclarationContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVariableDeclarationStatement is called when entering the variableDeclarationStatement production.
	EnterVariableDeclarationStatement(c *VariableDeclarationStatementContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterSimpleStatement is called when entering the simpleStatement production.
	EnterSimpleStatement(c *SimpleStatementContext)

	// EnterFunctionCallStatement is called when entering the functionCallStatement production.
	EnterFunctionCallStatement(c *FunctionCallStatementContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterBpfMapOperationStatement is called when entering the bpfMapOperationStatement production.
	EnterBpfMapOperationStatement(c *BpfMapOperationStatementContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitStructDefinition is called when exiting the structDefinition production.
	ExitStructDefinition(c *StructDefinitionContext)

	// ExitStructMember is called when exiting the structMember production.
	ExitStructMember(c *StructMemberContext)

	// ExitBpfMapDefinition is called when exiting the bpfMapDefinition production.
	ExitBpfMapDefinition(c *BpfMapDefinitionContext)

	// ExitBpfMapOption is called when exiting the bpfMapOption production.
	ExitBpfMapOption(c *BpfMapOptionContext)

	// ExitGlobalVariableDeclaration is called when exiting the globalVariableDeclaration production.
	ExitGlobalVariableDeclaration(c *GlobalVariableDeclarationContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVariableDeclarationStatement is called when exiting the variableDeclarationStatement production.
	ExitVariableDeclarationStatement(c *VariableDeclarationStatementContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitSimpleStatement is called when exiting the simpleStatement production.
	ExitSimpleStatement(c *SimpleStatementContext)

	// ExitFunctionCallStatement is called when exiting the functionCallStatement production.
	ExitFunctionCallStatement(c *FunctionCallStatementContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitBpfMapOperationStatement is called when exiting the bpfMapOperationStatement production.
	ExitBpfMapOperationStatement(c *BpfMapOperationStatementContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)
}
