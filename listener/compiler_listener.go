package listener

import "github.com/toastsandwich/ezbpf/grammar"

type Listener interface {
	EnterType(*grammar.TypeContext) string
	EnterAssign(*grammar.AssignContext) string
	EnterCompare(*grammar.CompareContext) string
	EnterExpression(*grammar.ExpressionContext) string
	EnterStructFieldAssign(*grammar.StructFieldAssignContext) string
	EnterArg(*grammar.ArgContext) string
	EnterArgs(*grammar.ArgsContext) string
	EnterParam(*grammar.ParamContext) string
	EnterParams(*grammar.ParamsContext) string
	EnterVarInitStmt(*grammar.VarInitStmtContext) string
	EnterVarDeclStmt(*grammar.VarDeclStmtContext) string
	EnterConstDeclStmt(*grammar.ConstDeclStmtContext) string
	EnterMapType(*grammar.MapTypeContext) string
	EnterMapDeclStmt(*grammar.MapDeclStmtContext) string
	EnterStructDeclStmt(*grammar.StructDeclStmtContext) string
	EnterStructDataMemStmt(*grammar.StructDataMemStmtContext) string
	// TODO implement further more
}
