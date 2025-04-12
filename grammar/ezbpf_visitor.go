// Code generated from ezbpf.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // ezbpf
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ezbpfParser.
type ezbpfVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ezbpfParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by ezbpfParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by ezbpfParser#compare.
	VisitCompare(ctx *CompareContext) interface{}

	// Visit a parse tree produced by ezbpfParser#dotExpression.
	VisitDotExpression(ctx *DotExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#binLiteralExpression.
	VisitBinLiteralExpression(ctx *BinLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#lshiftExpression.
	VisitLshiftExpression(ctx *LshiftExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#identifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#modExpression.
	VisitModExpression(ctx *ModExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#bitorExpression.
	VisitBitorExpression(ctx *BitorExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#divExpression.
	VisitDivExpression(ctx *DivExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#octLiteralExpression.
	VisitOctLiteralExpression(ctx *OctLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#bitxorExpression.
	VisitBitxorExpression(ctx *BitxorExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#stringLiteralExpression.
	VisitStringLiteralExpression(ctx *StringLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#bitandExpression.
	VisitBitandExpression(ctx *BitandExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#compareExpression.
	VisitCompareExpression(ctx *CompareExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#addExpression.
	VisitAddExpression(ctx *AddExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#hexLiteralExpression.
	VisitHexLiteralExpression(ctx *HexLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#funccallExpression.
	VisitFunccallExpression(ctx *FunccallExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#structinitExpression.
	VisitStructinitExpression(ctx *StructinitExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#charLiteralExpression.
	VisitCharLiteralExpression(ctx *CharLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#rshiftExpression.
	VisitRshiftExpression(ctx *RshiftExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#parExpression.
	VisitParExpression(ctx *ParExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#subExpression.
	VisitSubExpression(ctx *SubExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#decLiteralExpression.
	VisitDecLiteralExpression(ctx *DecLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#mulExpression.
	VisitMulExpression(ctx *MulExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#funcCallExpression.
	VisitFuncCallExpression(ctx *FuncCallExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#structInitExpression.
	VisitStructInitExpression(ctx *StructInitExpressionContext) interface{}

	// Visit a parse tree produced by ezbpfParser#structFieldAssign.
	VisitStructFieldAssign(ctx *StructFieldAssignContext) interface{}

	// Visit a parse tree produced by ezbpfParser#arg.
	VisitArg(ctx *ArgContext) interface{}

	// Visit a parse tree produced by ezbpfParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by ezbpfParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by ezbpfParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by ezbpfParser#varInitStmt.
	VisitVarInitStmt(ctx *VarInitStmtContext) interface{}

	// Visit a parse tree produced by ezbpfParser#varDeclStmt.
	VisitVarDeclStmt(ctx *VarDeclStmtContext) interface{}

	// Visit a parse tree produced by ezbpfParser#constDeclStmt.
	VisitConstDeclStmt(ctx *ConstDeclStmtContext) interface{}

	// Visit a parse tree produced by ezbpfParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by ezbpfParser#mapDeclStmt.
	VisitMapDeclStmt(ctx *MapDeclStmtContext) interface{}

	// Visit a parse tree produced by ezbpfParser#structDataMemStmt.
	VisitStructDataMemStmt(ctx *StructDataMemStmtContext) interface{}

	// Visit a parse tree produced by ezbpfParser#structDeclStmt.
	VisitStructDeclStmt(ctx *StructDeclStmtContext) interface{}

	// Visit a parse tree produced by ezbpfParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by ezbpfParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by ezbpfParser#funcDeclStmt.
	VisitFuncDeclStmt(ctx *FuncDeclStmtContext) interface{}

	// Visit a parse tree produced by ezbpfParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by ezbpfParser#stmts.
	VisitStmts(ctx *StmtsContext) interface{}

	// Visit a parse tree produced by ezbpfParser#prog.
	VisitProg(ctx *ProgContext) interface{}
}
