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

// EnterType is called when production type is entered.
func (s *BaseezbpfListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseezbpfListener) ExitType(ctx *TypeContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseezbpfListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseezbpfListener) ExitAssign(ctx *AssignContext) {}

// EnterCompare is called when production compare is entered.
func (s *BaseezbpfListener) EnterCompare(ctx *CompareContext) {}

// ExitCompare is called when production compare is exited.
func (s *BaseezbpfListener) ExitCompare(ctx *CompareContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseezbpfListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseezbpfListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStructFieldAssign is called when production structFieldAssign is entered.
func (s *BaseezbpfListener) EnterStructFieldAssign(ctx *StructFieldAssignContext) {}

// ExitStructFieldAssign is called when production structFieldAssign is exited.
func (s *BaseezbpfListener) ExitStructFieldAssign(ctx *StructFieldAssignContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseezbpfListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseezbpfListener) ExitArg(ctx *ArgContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseezbpfListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseezbpfListener) ExitArgs(ctx *ArgsContext) {}

// EnterParam is called when production param is entered.
func (s *BaseezbpfListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseezbpfListener) ExitParam(ctx *ParamContext) {}

// EnterParams is called when production params is entered.
func (s *BaseezbpfListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseezbpfListener) ExitParams(ctx *ParamsContext) {}

// EnterVarInitStmt is called when production varInitStmt is entered.
func (s *BaseezbpfListener) EnterVarInitStmt(ctx *VarInitStmtContext) {}

// ExitVarInitStmt is called when production varInitStmt is exited.
func (s *BaseezbpfListener) ExitVarInitStmt(ctx *VarInitStmtContext) {}

// EnterVarDeclStmt is called when production varDeclStmt is entered.
func (s *BaseezbpfListener) EnterVarDeclStmt(ctx *VarDeclStmtContext) {}

// ExitVarDeclStmt is called when production varDeclStmt is exited.
func (s *BaseezbpfListener) ExitVarDeclStmt(ctx *VarDeclStmtContext) {}

// EnterConstDeclStmt is called when production constDeclStmt is entered.
func (s *BaseezbpfListener) EnterConstDeclStmt(ctx *ConstDeclStmtContext) {}

// ExitConstDeclStmt is called when production constDeclStmt is exited.
func (s *BaseezbpfListener) ExitConstDeclStmt(ctx *ConstDeclStmtContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseezbpfListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseezbpfListener) ExitMapType(ctx *MapTypeContext) {}

// EnterMapDeclStmt is called when production mapDeclStmt is entered.
func (s *BaseezbpfListener) EnterMapDeclStmt(ctx *MapDeclStmtContext) {}

// ExitMapDeclStmt is called when production mapDeclStmt is exited.
func (s *BaseezbpfListener) ExitMapDeclStmt(ctx *MapDeclStmtContext) {}

// EnterStructDataMemStmt is called when production structDataMemStmt is entered.
func (s *BaseezbpfListener) EnterStructDataMemStmt(ctx *StructDataMemStmtContext) {}

// ExitStructDataMemStmt is called when production structDataMemStmt is exited.
func (s *BaseezbpfListener) ExitStructDataMemStmt(ctx *StructDataMemStmtContext) {}

// EnterStructDeclStmt is called when production structDeclStmt is entered.
func (s *BaseezbpfListener) EnterStructDeclStmt(ctx *StructDeclStmtContext) {}

// ExitStructDeclStmt is called when production structDeclStmt is exited.
func (s *BaseezbpfListener) ExitStructDeclStmt(ctx *StructDeclStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseezbpfListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseezbpfListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseezbpfListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseezbpfListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterFuncDeclStmt is called when production funcDeclStmt is entered.
func (s *BaseezbpfListener) EnterFuncDeclStmt(ctx *FuncDeclStmtContext) {}

// ExitFuncDeclStmt is called when production funcDeclStmt is exited.
func (s *BaseezbpfListener) ExitFuncDeclStmt(ctx *FuncDeclStmtContext) {}

// EnterStmts is called when production stmts is entered.
func (s *BaseezbpfListener) EnterStmts(ctx *StmtsContext) {}

// ExitStmts is called when production stmts is exited.
func (s *BaseezbpfListener) ExitStmts(ctx *StmtsContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseezbpfListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseezbpfListener) ExitProg(ctx *ProgContext) {}
