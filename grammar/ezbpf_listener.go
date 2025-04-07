// Code generated from ezbpf.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // ezbpf
import "github.com/antlr4-go/antlr/v4"

// ezbpfListener is a complete listener for a parse tree produced by ezbpfParser.
type ezbpfListener interface {
	antlr.ParseTreeListener

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterCompare is called when entering the compare production.
	EnterCompare(c *CompareContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterStructFieldAssign is called when entering the structFieldAssign production.
	EnterStructFieldAssign(c *StructFieldAssignContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterVarInitStmt is called when entering the varInitStmt production.
	EnterVarInitStmt(c *VarInitStmtContext)

	// EnterVarDeclStmt is called when entering the varDeclStmt production.
	EnterVarDeclStmt(c *VarDeclStmtContext)

	// EnterConstDeclStmt is called when entering the constDeclStmt production.
	EnterConstDeclStmt(c *ConstDeclStmtContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterMapDeclStmt is called when entering the mapDeclStmt production.
	EnterMapDeclStmt(c *MapDeclStmtContext)

	// EnterStructDataMemStmt is called when entering the structDataMemStmt production.
	EnterStructDataMemStmt(c *StructDataMemStmtContext)

	// EnterStructDeclStmt is called when entering the structDeclStmt production.
	EnterStructDeclStmt(c *StructDeclStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterFuncDeclStmt is called when entering the funcDeclStmt production.
	EnterFuncDeclStmt(c *FuncDeclStmtContext)

	// EnterStmts is called when entering the stmts production.
	EnterStmts(c *StmtsContext)

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitCompare is called when exiting the compare production.
	ExitCompare(c *CompareContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitStructFieldAssign is called when exiting the structFieldAssign production.
	ExitStructFieldAssign(c *StructFieldAssignContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitVarInitStmt is called when exiting the varInitStmt production.
	ExitVarInitStmt(c *VarInitStmtContext)

	// ExitVarDeclStmt is called when exiting the varDeclStmt production.
	ExitVarDeclStmt(c *VarDeclStmtContext)

	// ExitConstDeclStmt is called when exiting the constDeclStmt production.
	ExitConstDeclStmt(c *ConstDeclStmtContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitMapDeclStmt is called when exiting the mapDeclStmt production.
	ExitMapDeclStmt(c *MapDeclStmtContext)

	// ExitStructDataMemStmt is called when exiting the structDataMemStmt production.
	ExitStructDataMemStmt(c *StructDataMemStmtContext)

	// ExitStructDeclStmt is called when exiting the structDeclStmt production.
	ExitStructDeclStmt(c *StructDeclStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitFuncDeclStmt is called when exiting the funcDeclStmt production.
	ExitFuncDeclStmt(c *FuncDeclStmtContext)

	// ExitStmts is called when exiting the stmts production.
	ExitStmts(c *StmtsContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)
}
