// Code generated from ezbpf.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // ezbpf
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ezbpfParser struct {
	*antlr.BaseParser
}

var EzbpfParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ezbpfParserInit() {
	staticData := &EzbpfParserStaticData
	staticData.LiteralNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "'%'", "'&'", "'|'", "'^'", "'<<'",
		"'>>'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'", "'||'",
		"'!'", "'='", "'+='", "'-='", "'*='", "'/='", "'%='", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "','", "';'", "':'", "'.'", "", "", "",
		"", "", "", "'__u32'", "'__u64'", "'__s32'", "'__s64'", "'u32'", "'u64'",
		"'s32'", "'s64'", "'__be32'", "'__be64'", "'var'", "'const'", "'char'",
		"'uchar'", "'int'", "'long'", "'short'", "'uint'", "'void'", "'struct'",
		"'ethhdr'", "'iphdr'", "'tcphdr'", "'udphdr'", "'fn'", "'return'", "'if'",
		"'elif'", "'else'", "'map'", "'BPF_MAP_TYPE_HASH'", "'BPF_MAP_TYPE_ARRAY'",
		"'BPF_MAP_TYPE_PERCPU_HASH'", "'BPF_MAP_TYPE_PERCPU_ARRAY'", "'BPF_MAP_TYPE_LRU_HASH'",
		"'BPF_MAP_TYPE_LRU_PERCPU_HASH'",
	}
	staticData.SymbolicNames = []string{
		"", "ADD", "SUB", "MUL", "DIV", "MOD", "BIT_AND", "BIT_OR", "BIT_XOR",
		"LSHIFT", "RSHIFT", "EQ", "NEQ", "LT", "GT", "LTE", "GTE", "AND", "OR",
		"NOT", "ASSIGN", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN",
		"MOD_ASSIGN", "LPAR", "RPAR", "LBRA", "RBRA", "LSQ", "RSQ", "COMMA",
		"SEMI", "COLON", "DOT", "HEX_LITERAL", "OCT_LITERAL", "BIN_LITERAL",
		"DEC_LITERAL", "CHAR_LITERAL", "STRING_LITERAL", "UPTR32", "UPTR64",
		"SPTR32", "SPTR64", "U32", "U64", "S32", "S64", "BEPTR32", "BEPTR64",
		"VAR", "CONST", "CHAR", "UCHAR", "INT", "LONG", "SHORT", "UINT", "VOID",
		"STRUCT", "ETHHDR", "IPHDR", "TCPHDR", "UDPHDR", "FN", "RETURN", "IF",
		"ELSEIF", "ELSE", "MAP", "BPF_MAP_TYPE_HASH", "BPF_MAP_TYPE_ARRAY",
		"BPF_MAP_TYPE_PERCPU_HASH", "BPF_MAP_TYPE_PERCPU_ARRAY", "BPF_MAP_TYPE_LRU_HASH",
		"BPF_MAP_TYPE_LRU_PERCPU_HASH", "IDENTIFIER", "WS", "COMMENT", "MULTI_COMMENT",
	}
	staticData.RuleNames = []string{
		"baseType", "type", "assign", "compare", "expression", "funcCallExpression",
		"structInitExpression", "structFieldAssign", "arg", "args", "param",
		"params", "varInitStmt", "varDeclStmt", "constDeclStmt", "mapType",
		"mapDeclStmt", "structDataMemStmt", "structDeclStmt", "ifStmt", "returnStmt",
		"funcDeclStmt", "funcCallStmt", "stmt", "stmts", "prog",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 81, 304, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 59, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 83, 8, 4, 10, 4, 12, 4, 86, 9, 4,
		1, 4, 1, 4, 3, 4, 90, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 134,
		8, 4, 10, 4, 12, 4, 137, 9, 4, 1, 5, 1, 5, 1, 5, 3, 5, 142, 8, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 3, 6, 149, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 160, 8, 7, 10, 7, 12, 7, 163, 9, 7, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 170, 8, 9, 10, 9, 12, 9, 173, 9, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 5, 11, 182, 8, 11, 10, 11, 12,
		11, 185, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 233, 8, 18, 10, 18,
		12, 18, 236, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 20, 1, 20, 3, 20, 250, 8, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 1, 21, 3, 21, 258, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 279, 8, 23, 1, 24, 5, 24, 282, 8, 24,
		10, 24, 12, 24, 285, 9, 24, 1, 25, 5, 25, 288, 8, 25, 10, 25, 12, 25, 291,
		9, 25, 1, 25, 5, 25, 294, 8, 25, 10, 25, 12, 25, 297, 9, 25, 1, 25, 4,
		25, 300, 8, 25, 11, 25, 12, 25, 301, 1, 25, 0, 1, 8, 26, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 0, 4, 4, 0, 42, 51, 54, 60, 62, 65, 78, 78, 1, 0, 20, 25, 1,
		0, 11, 19, 1, 0, 72, 77, 322, 0, 52, 1, 0, 0, 0, 2, 54, 1, 0, 0, 0, 4,
		60, 1, 0, 0, 0, 6, 62, 1, 0, 0, 0, 8, 89, 1, 0, 0, 0, 10, 138, 1, 0, 0,
		0, 12, 145, 1, 0, 0, 0, 14, 152, 1, 0, 0, 0, 16, 164, 1, 0, 0, 0, 18, 166,
		1, 0, 0, 0, 20, 174, 1, 0, 0, 0, 22, 178, 1, 0, 0, 0, 24, 186, 1, 0, 0,
		0, 26, 192, 1, 0, 0, 0, 28, 200, 1, 0, 0, 0, 30, 208, 1, 0, 0, 0, 32, 210,
		1, 0, 0, 0, 34, 223, 1, 0, 0, 0, 36, 228, 1, 0, 0, 0, 38, 239, 1, 0, 0,
		0, 40, 247, 1, 0, 0, 0, 42, 253, 1, 0, 0, 0, 44, 266, 1, 0, 0, 0, 46, 278,
		1, 0, 0, 0, 48, 283, 1, 0, 0, 0, 50, 289, 1, 0, 0, 0, 52, 53, 7, 0, 0,
		0, 53, 1, 1, 0, 0, 0, 54, 58, 3, 0, 0, 0, 55, 56, 5, 30, 0, 0, 56, 57,
		5, 39, 0, 0, 57, 59, 5, 31, 0, 0, 58, 55, 1, 0, 0, 0, 58, 59, 1, 0, 0,
		0, 59, 3, 1, 0, 0, 0, 60, 61, 7, 1, 0, 0, 61, 5, 1, 0, 0, 0, 62, 63, 7,
		2, 0, 0, 63, 7, 1, 0, 0, 0, 64, 65, 6, 4, -1, 0, 65, 90, 5, 36, 0, 0, 66,
		90, 5, 37, 0, 0, 67, 90, 5, 38, 0, 0, 68, 90, 5, 39, 0, 0, 69, 90, 5, 40,
		0, 0, 70, 90, 5, 41, 0, 0, 71, 90, 5, 78, 0, 0, 72, 90, 3, 10, 5, 0, 73,
		90, 3, 12, 6, 0, 74, 75, 5, 26, 0, 0, 75, 76, 3, 8, 4, 0, 76, 77, 5, 27,
		0, 0, 77, 90, 1, 0, 0, 0, 78, 79, 5, 30, 0, 0, 79, 84, 3, 8, 4, 0, 80,
		81, 5, 32, 0, 0, 81, 83, 3, 8, 4, 0, 82, 80, 1, 0, 0, 0, 83, 86, 1, 0,
		0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 84,
		1, 0, 0, 0, 87, 88, 5, 31, 0, 0, 88, 90, 1, 0, 0, 0, 89, 64, 1, 0, 0, 0,
		89, 66, 1, 0, 0, 0, 89, 67, 1, 0, 0, 0, 89, 68, 1, 0, 0, 0, 89, 69, 1,
		0, 0, 0, 89, 70, 1, 0, 0, 0, 89, 71, 1, 0, 0, 0, 89, 72, 1, 0, 0, 0, 89,
		73, 1, 0, 0, 0, 89, 74, 1, 0, 0, 0, 89, 78, 1, 0, 0, 0, 90, 135, 1, 0,
		0, 0, 91, 92, 10, 11, 0, 0, 92, 93, 5, 3, 0, 0, 93, 134, 3, 8, 4, 12, 94,
		95, 10, 10, 0, 0, 95, 96, 5, 4, 0, 0, 96, 134, 3, 8, 4, 11, 97, 98, 10,
		9, 0, 0, 98, 99, 5, 5, 0, 0, 99, 134, 3, 8, 4, 10, 100, 101, 10, 8, 0,
		0, 101, 102, 5, 1, 0, 0, 102, 134, 3, 8, 4, 9, 103, 104, 10, 7, 0, 0, 104,
		105, 5, 2, 0, 0, 105, 134, 3, 8, 4, 8, 106, 107, 10, 6, 0, 0, 107, 108,
		5, 9, 0, 0, 108, 134, 3, 8, 4, 7, 109, 110, 10, 5, 0, 0, 110, 111, 5, 10,
		0, 0, 111, 134, 3, 8, 4, 6, 112, 113, 10, 4, 0, 0, 113, 114, 5, 6, 0, 0,
		114, 134, 3, 8, 4, 5, 115, 116, 10, 3, 0, 0, 116, 117, 5, 7, 0, 0, 117,
		134, 3, 8, 4, 4, 118, 119, 10, 2, 0, 0, 119, 120, 5, 8, 0, 0, 120, 134,
		3, 8, 4, 3, 121, 122, 10, 1, 0, 0, 122, 123, 3, 6, 3, 0, 123, 124, 3, 8,
		4, 2, 124, 134, 1, 0, 0, 0, 125, 126, 10, 15, 0, 0, 126, 127, 5, 35, 0,
		0, 127, 134, 5, 78, 0, 0, 128, 129, 10, 13, 0, 0, 129, 130, 5, 30, 0, 0,
		130, 131, 3, 8, 4, 0, 131, 132, 5, 31, 0, 0, 132, 134, 1, 0, 0, 0, 133,
		91, 1, 0, 0, 0, 133, 94, 1, 0, 0, 0, 133, 97, 1, 0, 0, 0, 133, 100, 1,
		0, 0, 0, 133, 103, 1, 0, 0, 0, 133, 106, 1, 0, 0, 0, 133, 109, 1, 0, 0,
		0, 133, 112, 1, 0, 0, 0, 133, 115, 1, 0, 0, 0, 133, 118, 1, 0, 0, 0, 133,
		121, 1, 0, 0, 0, 133, 125, 1, 0, 0, 0, 133, 128, 1, 0, 0, 0, 134, 137,
		1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 9, 1, 0, 0,
		0, 137, 135, 1, 0, 0, 0, 138, 139, 5, 78, 0, 0, 139, 141, 5, 26, 0, 0,
		140, 142, 3, 18, 9, 0, 141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142,
		143, 1, 0, 0, 0, 143, 144, 5, 27, 0, 0, 144, 11, 1, 0, 0, 0, 145, 146,
		5, 78, 0, 0, 146, 148, 5, 28, 0, 0, 147, 149, 3, 14, 7, 0, 148, 147, 1,
		0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 5, 29, 0,
		0, 151, 13, 1, 0, 0, 0, 152, 153, 5, 78, 0, 0, 153, 154, 5, 34, 0, 0, 154,
		161, 3, 8, 4, 0, 155, 156, 5, 32, 0, 0, 156, 157, 5, 78, 0, 0, 157, 158,
		5, 34, 0, 0, 158, 160, 3, 8, 4, 0, 159, 155, 1, 0, 0, 0, 160, 163, 1, 0,
		0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 15, 1, 0, 0, 0,
		163, 161, 1, 0, 0, 0, 164, 165, 3, 8, 4, 0, 165, 17, 1, 0, 0, 0, 166, 171,
		3, 16, 8, 0, 167, 168, 5, 32, 0, 0, 168, 170, 3, 16, 8, 0, 169, 167, 1,
		0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0,
		0, 172, 19, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174, 175, 5, 78, 0, 0, 175,
		176, 5, 34, 0, 0, 176, 177, 3, 2, 1, 0, 177, 21, 1, 0, 0, 0, 178, 183,
		3, 20, 10, 0, 179, 180, 5, 32, 0, 0, 180, 182, 3, 20, 10, 0, 181, 179,
		1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0,
		0, 0, 184, 23, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 187, 5, 52, 0, 0,
		187, 188, 5, 78, 0, 0, 188, 189, 5, 34, 0, 0, 189, 190, 3, 2, 1, 0, 190,
		191, 5, 33, 0, 0, 191, 25, 1, 0, 0, 0, 192, 193, 5, 52, 0, 0, 193, 194,
		5, 78, 0, 0, 194, 195, 5, 34, 0, 0, 195, 196, 3, 2, 1, 0, 196, 197, 3,
		4, 2, 0, 197, 198, 3, 8, 4, 0, 198, 199, 5, 33, 0, 0, 199, 27, 1, 0, 0,
		0, 200, 201, 5, 53, 0, 0, 201, 202, 5, 78, 0, 0, 202, 203, 5, 34, 0, 0,
		203, 204, 3, 2, 1, 0, 204, 205, 3, 4, 2, 0, 205, 206, 3, 8, 4, 0, 206,
		207, 5, 33, 0, 0, 207, 29, 1, 0, 0, 0, 208, 209, 7, 3, 0, 0, 209, 31, 1,
		0, 0, 0, 210, 211, 5, 71, 0, 0, 211, 212, 5, 78, 0, 0, 212, 213, 5, 34,
		0, 0, 213, 214, 3, 30, 15, 0, 214, 215, 5, 13, 0, 0, 215, 216, 3, 2, 1,
		0, 216, 217, 5, 32, 0, 0, 217, 218, 3, 2, 1, 0, 218, 219, 5, 32, 0, 0,
		219, 220, 5, 39, 0, 0, 220, 221, 5, 14, 0, 0, 221, 222, 5, 33, 0, 0, 222,
		33, 1, 0, 0, 0, 223, 224, 5, 78, 0, 0, 224, 225, 5, 34, 0, 0, 225, 226,
		3, 2, 1, 0, 226, 227, 5, 33, 0, 0, 227, 35, 1, 0, 0, 0, 228, 229, 5, 61,
		0, 0, 229, 230, 5, 78, 0, 0, 230, 234, 5, 28, 0, 0, 231, 233, 3, 34, 17,
		0, 232, 231, 1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234,
		235, 1, 0, 0, 0, 235, 237, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 238,
		5, 29, 0, 0, 238, 37, 1, 0, 0, 0, 239, 240, 5, 68, 0, 0, 240, 241, 5, 26,
		0, 0, 241, 242, 3, 8, 4, 0, 242, 243, 5, 27, 0, 0, 243, 244, 5, 28, 0,
		0, 244, 245, 3, 48, 24, 0, 245, 246, 5, 29, 0, 0, 246, 39, 1, 0, 0, 0,
		247, 249, 5, 67, 0, 0, 248, 250, 3, 8, 4, 0, 249, 248, 1, 0, 0, 0, 249,
		250, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 252, 5, 33, 0, 0, 252, 41,
		1, 0, 0, 0, 253, 254, 5, 66, 0, 0, 254, 255, 5, 78, 0, 0, 255, 257, 5,
		26, 0, 0, 256, 258, 3, 22, 11, 0, 257, 256, 1, 0, 0, 0, 257, 258, 1, 0,
		0, 0, 258, 259, 1, 0, 0, 0, 259, 260, 5, 27, 0, 0, 260, 261, 5, 34, 0,
		0, 261, 262, 3, 2, 1, 0, 262, 263, 5, 28, 0, 0, 263, 264, 3, 48, 24, 0,
		264, 265, 5, 29, 0, 0, 265, 43, 1, 0, 0, 0, 266, 267, 3, 10, 5, 0, 267,
		268, 5, 33, 0, 0, 268, 45, 1, 0, 0, 0, 269, 279, 3, 24, 12, 0, 270, 279,
		3, 26, 13, 0, 271, 279, 3, 28, 14, 0, 272, 279, 3, 32, 16, 0, 273, 279,
		3, 36, 18, 0, 274, 279, 3, 38, 19, 0, 275, 279, 3, 40, 20, 0, 276, 279,
		3, 42, 21, 0, 277, 279, 3, 44, 22, 0, 278, 269, 1, 0, 0, 0, 278, 270, 1,
		0, 0, 0, 278, 271, 1, 0, 0, 0, 278, 272, 1, 0, 0, 0, 278, 273, 1, 0, 0,
		0, 278, 274, 1, 0, 0, 0, 278, 275, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278,
		277, 1, 0, 0, 0, 279, 47, 1, 0, 0, 0, 280, 282, 3, 46, 23, 0, 281, 280,
		1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0,
		0, 0, 284, 49, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 286, 288, 3, 36, 18, 0,
		287, 286, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 289,
		290, 1, 0, 0, 0, 290, 295, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 292, 294,
		3, 32, 16, 0, 293, 292, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295, 293, 1,
		0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 299, 1, 0, 0, 0, 297, 295, 1, 0, 0,
		0, 298, 300, 3, 42, 21, 0, 299, 298, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0,
		301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 51, 1, 0, 0, 0, 18, 58,
		84, 89, 133, 135, 141, 148, 161, 171, 183, 234, 249, 257, 278, 283, 289,
		295, 301,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ezbpfParserInit initializes any static state used to implement ezbpfParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewezbpfParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EzbpfParserInit() {
	staticData := &EzbpfParserStaticData
	staticData.once.Do(ezbpfParserInit)
}

// NewezbpfParser produces a new parser instance for the optional input antlr.TokenStream.
func NewezbpfParser(input antlr.TokenStream) *ezbpfParser {
	EzbpfParserInit()
	this := new(ezbpfParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EzbpfParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ezbpf.g4"

	return this
}

// ezbpfParser tokens.
const (
	ezbpfParserEOF                          = antlr.TokenEOF
	ezbpfParserADD                          = 1
	ezbpfParserSUB                          = 2
	ezbpfParserMUL                          = 3
	ezbpfParserDIV                          = 4
	ezbpfParserMOD                          = 5
	ezbpfParserBIT_AND                      = 6
	ezbpfParserBIT_OR                       = 7
	ezbpfParserBIT_XOR                      = 8
	ezbpfParserLSHIFT                       = 9
	ezbpfParserRSHIFT                       = 10
	ezbpfParserEQ                           = 11
	ezbpfParserNEQ                          = 12
	ezbpfParserLT                           = 13
	ezbpfParserGT                           = 14
	ezbpfParserLTE                          = 15
	ezbpfParserGTE                          = 16
	ezbpfParserAND                          = 17
	ezbpfParserOR                           = 18
	ezbpfParserNOT                          = 19
	ezbpfParserASSIGN                       = 20
	ezbpfParserADD_ASSIGN                   = 21
	ezbpfParserSUB_ASSIGN                   = 22
	ezbpfParserMUL_ASSIGN                   = 23
	ezbpfParserDIV_ASSIGN                   = 24
	ezbpfParserMOD_ASSIGN                   = 25
	ezbpfParserLPAR                         = 26
	ezbpfParserRPAR                         = 27
	ezbpfParserLBRA                         = 28
	ezbpfParserRBRA                         = 29
	ezbpfParserLSQ                          = 30
	ezbpfParserRSQ                          = 31
	ezbpfParserCOMMA                        = 32
	ezbpfParserSEMI                         = 33
	ezbpfParserCOLON                        = 34
	ezbpfParserDOT                          = 35
	ezbpfParserHEX_LITERAL                  = 36
	ezbpfParserOCT_LITERAL                  = 37
	ezbpfParserBIN_LITERAL                  = 38
	ezbpfParserDEC_LITERAL                  = 39
	ezbpfParserCHAR_LITERAL                 = 40
	ezbpfParserSTRING_LITERAL               = 41
	ezbpfParserUPTR32                       = 42
	ezbpfParserUPTR64                       = 43
	ezbpfParserSPTR32                       = 44
	ezbpfParserSPTR64                       = 45
	ezbpfParserU32                          = 46
	ezbpfParserU64                          = 47
	ezbpfParserS32                          = 48
	ezbpfParserS64                          = 49
	ezbpfParserBEPTR32                      = 50
	ezbpfParserBEPTR64                      = 51
	ezbpfParserVAR                          = 52
	ezbpfParserCONST                        = 53
	ezbpfParserCHAR                         = 54
	ezbpfParserUCHAR                        = 55
	ezbpfParserINT                          = 56
	ezbpfParserLONG                         = 57
	ezbpfParserSHORT                        = 58
	ezbpfParserUINT                         = 59
	ezbpfParserVOID                         = 60
	ezbpfParserSTRUCT                       = 61
	ezbpfParserETHHDR                       = 62
	ezbpfParserIPHDR                        = 63
	ezbpfParserTCPHDR                       = 64
	ezbpfParserUDPHDR                       = 65
	ezbpfParserFN                           = 66
	ezbpfParserRETURN                       = 67
	ezbpfParserIF                           = 68
	ezbpfParserELSEIF                       = 69
	ezbpfParserELSE                         = 70
	ezbpfParserMAP                          = 71
	ezbpfParserBPF_MAP_TYPE_HASH            = 72
	ezbpfParserBPF_MAP_TYPE_ARRAY           = 73
	ezbpfParserBPF_MAP_TYPE_PERCPU_HASH     = 74
	ezbpfParserBPF_MAP_TYPE_PERCPU_ARRAY    = 75
	ezbpfParserBPF_MAP_TYPE_LRU_HASH        = 76
	ezbpfParserBPF_MAP_TYPE_LRU_PERCPU_HASH = 77
	ezbpfParserIDENTIFIER                   = 78
	ezbpfParserWS                           = 79
	ezbpfParserCOMMENT                      = 80
	ezbpfParserMULTI_COMMENT                = 81
)

// ezbpfParser rules.
const (
	ezbpfParserRULE_baseType             = 0
	ezbpfParserRULE_type                 = 1
	ezbpfParserRULE_assign               = 2
	ezbpfParserRULE_compare              = 3
	ezbpfParserRULE_expression           = 4
	ezbpfParserRULE_funcCallExpression   = 5
	ezbpfParserRULE_structInitExpression = 6
	ezbpfParserRULE_structFieldAssign    = 7
	ezbpfParserRULE_arg                  = 8
	ezbpfParserRULE_args                 = 9
	ezbpfParserRULE_param                = 10
	ezbpfParserRULE_params               = 11
	ezbpfParserRULE_varInitStmt          = 12
	ezbpfParserRULE_varDeclStmt          = 13
	ezbpfParserRULE_constDeclStmt        = 14
	ezbpfParserRULE_mapType              = 15
	ezbpfParserRULE_mapDeclStmt          = 16
	ezbpfParserRULE_structDataMemStmt    = 17
	ezbpfParserRULE_structDeclStmt       = 18
	ezbpfParserRULE_ifStmt               = 19
	ezbpfParserRULE_returnStmt           = 20
	ezbpfParserRULE_funcDeclStmt         = 21
	ezbpfParserRULE_funcCallStmt         = 22
	ezbpfParserRULE_stmt                 = 23
	ezbpfParserRULE_stmts                = 24
	ezbpfParserRULE_prog                 = 25
)

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	U32() antlr.TerminalNode
	U64() antlr.TerminalNode
	UPTR32() antlr.TerminalNode
	UPTR64() antlr.TerminalNode
	BEPTR32() antlr.TerminalNode
	BEPTR64() antlr.TerminalNode
	S32() antlr.TerminalNode
	S64() antlr.TerminalNode
	SPTR32() antlr.TerminalNode
	SPTR64() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	UCHAR() antlr.TerminalNode
	SHORT() antlr.TerminalNode
	LONG() antlr.TerminalNode
	UINT() antlr.TerminalNode
	INT() antlr.TerminalNode
	VOID() antlr.TerminalNode
	ETHHDR() antlr.TerminalNode
	IPHDR() antlr.TerminalNode
	TCPHDR() antlr.TerminalNode
	UDPHDR() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_baseType
	return p
}

func InitEmptyBaseTypeContext(p *BaseTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_baseType
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseTypeContext) U32() antlr.TerminalNode {
	return s.GetToken(ezbpfParserU32, 0)
}

func (s *BaseTypeContext) U64() antlr.TerminalNode {
	return s.GetToken(ezbpfParserU64, 0)
}

func (s *BaseTypeContext) UPTR32() antlr.TerminalNode {
	return s.GetToken(ezbpfParserUPTR32, 0)
}

func (s *BaseTypeContext) UPTR64() antlr.TerminalNode {
	return s.GetToken(ezbpfParserUPTR64, 0)
}

func (s *BaseTypeContext) BEPTR32() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBEPTR32, 0)
}

func (s *BaseTypeContext) BEPTR64() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBEPTR64, 0)
}

func (s *BaseTypeContext) S32() antlr.TerminalNode {
	return s.GetToken(ezbpfParserS32, 0)
}

func (s *BaseTypeContext) S64() antlr.TerminalNode {
	return s.GetToken(ezbpfParserS64, 0)
}

func (s *BaseTypeContext) SPTR32() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSPTR32, 0)
}

func (s *BaseTypeContext) SPTR64() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSPTR64, 0)
}

func (s *BaseTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCHAR, 0)
}

func (s *BaseTypeContext) UCHAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserUCHAR, 0)
}

func (s *BaseTypeContext) SHORT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSHORT, 0)
}

func (s *BaseTypeContext) LONG() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLONG, 0)
}

func (s *BaseTypeContext) UINT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserUINT, 0)
}

func (s *BaseTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserINT, 0)
}

func (s *BaseTypeContext) VOID() antlr.TerminalNode {
	return s.GetToken(ezbpfParserVOID, 0)
}

func (s *BaseTypeContext) ETHHDR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserETHHDR, 0)
}

func (s *BaseTypeContext) IPHDR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIPHDR, 0)
}

func (s *BaseTypeContext) TCPHDR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserTCPHDR, 0)
}

func (s *BaseTypeContext) UDPHDR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserUDPHDR, 0)
}

func (s *BaseTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitBaseType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) BaseType() (localctx IBaseTypeContext) {
	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ezbpfParserRULE_baseType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-42)) & ^0x3f) == 0 && ((int64(1)<<(_la-42))&68735726591) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BaseType() IBaseTypeContext
	LSQ() antlr.TerminalNode
	DEC_LITERAL() antlr.TerminalNode
	RSQ() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) BaseType() IBaseTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *TypeContext) LSQ() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLSQ, 0)
}

func (s *TypeContext) DEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserDEC_LITERAL, 0)
}

func (s *TypeContext) RSQ() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRSQ, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ezbpfParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.BaseType()
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ezbpfParserLSQ {
		{
			p.SetState(55)
			p.Match(ezbpfParserLSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Match(ezbpfParserDEC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Match(ezbpfParserRSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	ADD_ASSIGN() antlr.TerminalNode
	SUB_ASSIGN() antlr.TerminalNode
	MUL_ASSIGN() antlr.TerminalNode
	DIV_ASSIGN() antlr.TerminalNode
	MOD_ASSIGN() antlr.TerminalNode

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserASSIGN, 0)
}

func (s *AssignContext) ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserADD_ASSIGN, 0)
}

func (s *AssignContext) SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSUB_ASSIGN, 0)
}

func (s *AssignContext) MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserMUL_ASSIGN, 0)
}

func (s *AssignContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserDIV_ASSIGN, 0)
}

func (s *AssignContext) MOD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserMOD_ASSIGN, 0)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ezbpfParserRULE_assign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66060288) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompareContext is an interface to support dynamic dispatch.
type ICompareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	GT() antlr.TerminalNode
	LT() antlr.TerminalNode
	GTE() antlr.TerminalNode
	LTE() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsCompareContext differentiates from other interfaces.
	IsCompareContext()
}

type CompareContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareContext() *CompareContext {
	var p = new(CompareContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_compare
	return p
}

func InitEmptyCompareContext(p *CompareContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_compare
}

func (*CompareContext) IsCompareContext() {}

func NewCompareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareContext {
	var p = new(CompareContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_compare

	return p
}

func (s *CompareContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareContext) EQ() antlr.TerminalNode {
	return s.GetToken(ezbpfParserEQ, 0)
}

func (s *CompareContext) NEQ() antlr.TerminalNode {
	return s.GetToken(ezbpfParserNEQ, 0)
}

func (s *CompareContext) GT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserGT, 0)
}

func (s *CompareContext) LT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLT, 0)
}

func (s *CompareContext) GTE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserGTE, 0)
}

func (s *CompareContext) LTE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLTE, 0)
}

func (s *CompareContext) AND() antlr.TerminalNode {
	return s.GetToken(ezbpfParserAND, 0)
}

func (s *CompareContext) OR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserOR, 0)
}

func (s *CompareContext) NOT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserNOT, 0)
}

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitCompare(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) Compare() (localctx ICompareContext) {
	localctx = NewCompareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ezbpfParserRULE_compare)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1046528) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DotExpressionContext struct {
	ExpressionContext
}

func NewDotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotExpressionContext {
	var p = new(DotExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DotExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserDOT, 0)
}

func (s *DotExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *DotExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitDotExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinLiteralExpressionContext struct {
	ExpressionContext
}

func NewBinLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinLiteralExpressionContext {
	var p = new(BinLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinLiteralExpressionContext) BIN_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBIN_LITERAL, 0)
}

func (s *BinLiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitBinLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LshiftExpressionContext struct {
	ExpressionContext
}

func NewLshiftExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LshiftExpressionContext {
	var p = new(LshiftExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LshiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LshiftExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LshiftExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LshiftExpressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLSHIFT, 0)
}

func (s *LshiftExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitLshiftExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModExpressionContext struct {
	ExpressionContext
}

func NewModExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModExpressionContext {
	var p = new(ModExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ModExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ModExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ModExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(ezbpfParserMOD, 0)
}

func (s *ModExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitModExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitorExpressionContext struct {
	ExpressionContext
}

func NewBitorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitorExpressionContext {
	var p = new(BitorExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitorExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitorExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitorExpressionContext) BIT_OR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBIT_OR, 0)
}

func (s *BitorExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitBitorExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivExpressionContext struct {
	ExpressionContext
}

func NewDivExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivExpressionContext {
	var p = new(DivExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DivExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DivExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DivExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(ezbpfParserDIV, 0)
}

func (s *DivExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitDivExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type OctLiteralExpressionContext struct {
	ExpressionContext
}

func NewOctLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OctLiteralExpressionContext {
	var p = new(OctLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OctLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OctLiteralExpressionContext) OCT_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserOCT_LITERAL, 0)
}

func (s *OctLiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitOctLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitxorExpressionContext struct {
	ExpressionContext
}

func NewBitxorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitxorExpressionContext {
	var p = new(BitxorExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitxorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitxorExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitxorExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitxorExpressionContext) BIT_XOR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBIT_XOR, 0)
}

func (s *BitxorExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitBitxorExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringLiteralExpressionContext struct {
	ExpressionContext
}

func NewStringLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralExpressionContext {
	var p = new(StringLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StringLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralExpressionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSTRING_LITERAL, 0)
}

func (s *StringLiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitStringLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitandExpressionContext struct {
	ExpressionContext
}

func NewBitandExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitandExpressionContext {
	var p = new(BitandExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitandExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitandExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitandExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitandExpressionContext) BIT_AND() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBIT_AND, 0)
}

func (s *BitandExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitBitandExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareExpressionContext struct {
	ExpressionContext
}

func NewCompareExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpressionContext {
	var p = new(CompareExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CompareExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompareExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompareExpressionContext) Compare() ICompareContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareContext)
}

func (s *CompareExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitCompareExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayLiteralExpressionContext struct {
	ExpressionContext
}

func NewArrayLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralExpressionContext {
	var p = new(ArrayLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralExpressionContext) LSQ() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLSQ, 0)
}

func (s *ArrayLiteralExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLiteralExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayLiteralExpressionContext) RSQ() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRSQ, 0)
}

func (s *ArrayLiteralExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserCOMMA)
}

func (s *ArrayLiteralExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOMMA, i)
}

func (s *ArrayLiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitArrayLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExpressionContext struct {
	ExpressionContext
}

func NewAddExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExpressionContext {
	var p = new(AddExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(ezbpfParserADD, 0)
}

func (s *AddExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitAddExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type HexLiteralExpressionContext struct {
	ExpressionContext
}

func NewHexLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HexLiteralExpressionContext {
	var p = new(HexLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HexLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexLiteralExpressionContext) HEX_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserHEX_LITERAL, 0)
}

func (s *HexLiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitHexLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunccallExpressionContext struct {
	ExpressionContext
}

func NewFunccallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunccallExpressionContext {
	var p = new(FunccallExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunccallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunccallExpressionContext) FuncCallExpression() IFuncCallExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallExpressionContext)
}

func (s *FunccallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitFunccallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructinitExpressionContext struct {
	ExpressionContext
}

func NewStructinitExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructinitExpressionContext {
	var p = new(StructinitExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StructinitExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructinitExpressionContext) StructInitExpression() IStructInitExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructInitExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructInitExpressionContext)
}

func (s *StructinitExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitStructinitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CharLiteralExpressionContext struct {
	ExpressionContext
}

func NewCharLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharLiteralExpressionContext {
	var p = new(CharLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CharLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharLiteralExpressionContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCHAR_LITERAL, 0)
}

func (s *CharLiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitCharLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type RshiftExpressionContext struct {
	ExpressionContext
}

func NewRshiftExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RshiftExpressionContext {
	var p = new(RshiftExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RshiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RshiftExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RshiftExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RshiftExpressionContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRSHIFT, 0)
}

func (s *RshiftExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitRshiftExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexExpressionContext struct {
	ExpressionContext
}

func NewIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExpressionContext {
	var p = new(IndexExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IndexExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexExpressionContext) LSQ() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLSQ, 0)
}

func (s *IndexExpressionContext) RSQ() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRSQ, 0)
}

func (s *IndexExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitIndexExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParExpressionContext struct {
	ExpressionContext
}

func NewParExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExpressionContext {
	var p = new(ParExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExpressionContext) LPAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLPAR, 0)
}

func (s *ParExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParExpressionContext) RPAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRPAR, 0)
}

func (s *ParExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitParExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubExpressionContext struct {
	ExpressionContext
}

func NewSubExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubExpressionContext {
	var p = new(SubExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SubExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SubExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SubExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSUB, 0)
}

func (s *SubExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitSubExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecLiteralExpressionContext struct {
	ExpressionContext
}

func NewDecLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecLiteralExpressionContext {
	var p = new(DecLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DecLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecLiteralExpressionContext) DEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserDEC_LITERAL, 0)
}

func (s *DecLiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitDecLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulExpressionContext struct {
	ExpressionContext
}

func NewMulExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExpressionContext {
	var p = new(MulExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserMUL, 0)
}

func (s *MulExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitMulExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ezbpfParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, ezbpfParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewHexLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(65)
			p.Match(ezbpfParserHEX_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewOctLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)
			p.Match(ezbpfParserOCT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewBinLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(67)
			p.Match(ezbpfParserBIN_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewDecLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(ezbpfParserDEC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewCharLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Match(ezbpfParserCHAR_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStringLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(70)
			p.Match(ezbpfParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(71)
			p.Match(ezbpfParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewFunccallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(72)
			p.FuncCallExpression()
		}

	case 9:
		localctx = NewStructinitExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(73)
			p.StructInitExpression()
		}

	case 10:
		localctx = NewParExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(74)
			p.Match(ezbpfParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.expression(0)
		}
		{
			p.SetState(76)
			p.Match(ezbpfParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewArrayLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(78)
			p.Match(ezbpfParserLSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.expression(0)
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ezbpfParserCOMMA {
			{
				p.SetState(80)
				p.Match(ezbpfParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(81)
				p.expression(0)
			}

			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(87)
			p.Match(ezbpfParserRSQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(133)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(91)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(92)
					p.Match(ezbpfParserMUL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(93)
					p.expression(12)
				}

			case 2:
				localctx = NewDivExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(94)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(95)
					p.Match(ezbpfParserDIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(96)
					p.expression(11)
				}

			case 3:
				localctx = NewModExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(97)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(98)
					p.Match(ezbpfParserMOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(99)
					p.expression(10)
				}

			case 4:
				localctx = NewAddExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(100)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(101)
					p.Match(ezbpfParserADD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(102)
					p.expression(9)
				}

			case 5:
				localctx = NewSubExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(104)
					p.Match(ezbpfParserSUB)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(105)
					p.expression(8)
				}

			case 6:
				localctx = NewLshiftExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(106)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(107)
					p.Match(ezbpfParserLSHIFT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(108)
					p.expression(7)
				}

			case 7:
				localctx = NewRshiftExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(110)
					p.Match(ezbpfParserRSHIFT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(111)
					p.expression(6)
				}

			case 8:
				localctx = NewBitandExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(112)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(113)
					p.Match(ezbpfParserBIT_AND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(114)
					p.expression(5)
				}

			case 9:
				localctx = NewBitorExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(115)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(116)
					p.Match(ezbpfParserBIT_OR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(117)
					p.expression(4)
				}

			case 10:
				localctx = NewBitxorExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(118)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(119)
					p.Match(ezbpfParserBIT_XOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(120)
					p.expression(3)
				}

			case 11:
				localctx = NewCompareExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(121)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(122)
					p.Compare()
				}
				{
					p.SetState(123)
					p.expression(2)
				}

			case 12:
				localctx = NewDotExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(125)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(126)
					p.Match(ezbpfParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(127)
					p.Match(ezbpfParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 13:
				localctx = NewIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(129)
					p.Match(ezbpfParserLSQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(130)
					p.expression(0)
				}
				{
					p.SetState(131)
					p.Match(ezbpfParserRSQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncCallExpressionContext is an interface to support dynamic dispatch.
type IFuncCallExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LPAR() antlr.TerminalNode
	RPAR() antlr.TerminalNode
	Args() IArgsContext

	// IsFuncCallExpressionContext differentiates from other interfaces.
	IsFuncCallExpressionContext()
}

type FuncCallExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallExpressionContext() *FuncCallExpressionContext {
	var p = new(FuncCallExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_funcCallExpression
	return p
}

func InitEmptyFuncCallExpressionContext(p *FuncCallExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_funcCallExpression
}

func (*FuncCallExpressionContext) IsFuncCallExpressionContext() {}

func NewFuncCallExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallExpressionContext {
	var p = new(FuncCallExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_funcCallExpression

	return p
}

func (s *FuncCallExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *FuncCallExpressionContext) LPAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLPAR, 0)
}

func (s *FuncCallExpressionContext) RPAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRPAR, 0)
}

func (s *FuncCallExpressionContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *FuncCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitFuncCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) FuncCallExpression() (localctx IFuncCallExpressionContext) {
	localctx = NewFuncCallExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ezbpfParserRULE_funcCallExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(ezbpfParserLPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&4503599627435025) != 0 {
		{
			p.SetState(140)
			p.Args()
		}

	}
	{
		p.SetState(143)
		p.Match(ezbpfParserRPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructInitExpressionContext is an interface to support dynamic dispatch.
type IStructInitExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LBRA() antlr.TerminalNode
	RBRA() antlr.TerminalNode
	StructFieldAssign() IStructFieldAssignContext

	// IsStructInitExpressionContext differentiates from other interfaces.
	IsStructInitExpressionContext()
}

type StructInitExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructInitExpressionContext() *StructInitExpressionContext {
	var p = new(StructInitExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structInitExpression
	return p
}

func InitEmptyStructInitExpressionContext(p *StructInitExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structInitExpression
}

func (*StructInitExpressionContext) IsStructInitExpressionContext() {}

func NewStructInitExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructInitExpressionContext {
	var p = new(StructInitExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_structInitExpression

	return p
}

func (s *StructInitExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *StructInitExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *StructInitExpressionContext) LBRA() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLBRA, 0)
}

func (s *StructInitExpressionContext) RBRA() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRBRA, 0)
}

func (s *StructInitExpressionContext) StructFieldAssign() IStructFieldAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructFieldAssignContext)
}

func (s *StructInitExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInitExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructInitExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitStructInitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) StructInitExpression() (localctx IStructInitExpressionContext) {
	localctx = NewStructInitExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ezbpfParserRULE_structInitExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(ezbpfParserLBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ezbpfParserIDENTIFIER {
		{
			p.SetState(147)
			p.StructFieldAssign()
		}

	}
	{
		p.SetState(150)
		p.Match(ezbpfParserRBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructFieldAssignContext is an interface to support dynamic dispatch.
type IStructFieldAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsStructFieldAssignContext differentiates from other interfaces.
	IsStructFieldAssignContext()
}

type StructFieldAssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructFieldAssignContext() *StructFieldAssignContext {
	var p = new(StructFieldAssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structFieldAssign
	return p
}

func InitEmptyStructFieldAssignContext(p *StructFieldAssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structFieldAssign
}

func (*StructFieldAssignContext) IsStructFieldAssignContext() {}

func NewStructFieldAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldAssignContext {
	var p = new(StructFieldAssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_structFieldAssign

	return p
}

func (s *StructFieldAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldAssignContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserIDENTIFIER)
}

func (s *StructFieldAssignContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, i)
}

func (s *StructFieldAssignContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserCOLON)
}

func (s *StructFieldAssignContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, i)
}

func (s *StructFieldAssignContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *StructFieldAssignContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StructFieldAssignContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserCOMMA)
}

func (s *StructFieldAssignContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOMMA, i)
}

func (s *StructFieldAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFieldAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitStructFieldAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) StructFieldAssign() (localctx IStructFieldAssignContext) {
	localctx = NewStructFieldAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ezbpfParserRULE_structFieldAssign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.expression(0)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserCOMMA {
		{
			p.SetState(155)
			p.Match(ezbpfParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Match(ezbpfParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Match(ezbpfParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.expression(0)
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_arg
	return p
}

func InitEmptyArgContext(p *ArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_arg
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ezbpfParserRULE_arg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArg() []IArgContext
	Arg(i int) IArgContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_args
	return p
}

func InitEmptyArgsContext(p *ArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_args
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllArg() []IArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgContext); ok {
			len++
		}
	}

	tst := make([]IArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgContext); ok {
			tst[i] = t.(IArgContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) Arg(i int) IArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *ArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserCOMMA)
}

func (s *ArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOMMA, i)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ezbpfParserRULE_args)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Arg()
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserCOMMA {
		{
			p.SetState(167)
			p.Match(ezbpfParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Arg()
		}

		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *ParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *ParamContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ezbpfParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserCOMMA)
}

func (s *ParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOMMA, i)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ezbpfParserRULE_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Param()
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserCOMMA {
		{
			p.SetState(179)
			p.Match(ezbpfParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Param()
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarInitStmtContext is an interface to support dynamic dispatch.
type IVarInitStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	SEMI() antlr.TerminalNode

	// IsVarInitStmtContext differentiates from other interfaces.
	IsVarInitStmtContext()
}

type VarInitStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarInitStmtContext() *VarInitStmtContext {
	var p = new(VarInitStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_varInitStmt
	return p
}

func InitEmptyVarInitStmtContext(p *VarInitStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_varInitStmt
}

func (*VarInitStmtContext) IsVarInitStmtContext() {}

func NewVarInitStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarInitStmtContext {
	var p = new(VarInitStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_varInitStmt

	return p
}

func (s *VarInitStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VarInitStmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserVAR, 0)
}

func (s *VarInitStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *VarInitStmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *VarInitStmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarInitStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *VarInitStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarInitStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarInitStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitVarInitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) VarInitStmt() (localctx IVarInitStmtContext) {
	localctx = NewVarInitStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ezbpfParserRULE_varInitStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(ezbpfParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.Type_()
	}
	{
		p.SetState(190)
		p.Match(ezbpfParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclStmtContext is an interface to support dynamic dispatch.
type IVarDeclStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	Assign() IAssignContext
	Expression() IExpressionContext
	SEMI() antlr.TerminalNode

	// IsVarDeclStmtContext differentiates from other interfaces.
	IsVarDeclStmtContext()
}

type VarDeclStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclStmtContext() *VarDeclStmtContext {
	var p = new(VarDeclStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_varDeclStmt
	return p
}

func InitEmptyVarDeclStmtContext(p *VarDeclStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_varDeclStmt
}

func (*VarDeclStmtContext) IsVarDeclStmtContext() {}

func NewVarDeclStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclStmtContext {
	var p = new(VarDeclStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_varDeclStmt

	return p
}

func (s *VarDeclStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclStmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserVAR, 0)
}

func (s *VarDeclStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *VarDeclStmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *VarDeclStmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarDeclStmtContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *VarDeclStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarDeclStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *VarDeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitVarDeclStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) VarDeclStmt() (localctx IVarDeclStmtContext) {
	localctx = NewVarDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ezbpfParserRULE_varDeclStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(ezbpfParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Type_()
	}
	{
		p.SetState(196)
		p.Assign()
	}
	{
		p.SetState(197)
		p.expression(0)
	}
	{
		p.SetState(198)
		p.Match(ezbpfParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstDeclStmtContext is an interface to support dynamic dispatch.
type IConstDeclStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	Assign() IAssignContext
	Expression() IExpressionContext
	SEMI() antlr.TerminalNode

	// IsConstDeclStmtContext differentiates from other interfaces.
	IsConstDeclStmtContext()
}

type ConstDeclStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDeclStmtContext() *ConstDeclStmtContext {
	var p = new(ConstDeclStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_constDeclStmt
	return p
}

func InitEmptyConstDeclStmtContext(p *ConstDeclStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_constDeclStmt
}

func (*ConstDeclStmtContext) IsConstDeclStmtContext() {}

func NewConstDeclStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDeclStmtContext {
	var p = new(ConstDeclStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_constDeclStmt

	return p
}

func (s *ConstDeclStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDeclStmtContext) CONST() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCONST, 0)
}

func (s *ConstDeclStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *ConstDeclStmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *ConstDeclStmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ConstDeclStmtContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ConstDeclStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConstDeclStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *ConstDeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDeclStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstDeclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitConstDeclStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) ConstDeclStmt() (localctx IConstDeclStmtContext) {
	localctx = NewConstDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ezbpfParserRULE_constDeclStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(ezbpfParserCONST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.Type_()
	}
	{
		p.SetState(204)
		p.Assign()
	}
	{
		p.SetState(205)
		p.expression(0)
	}
	{
		p.SetState(206)
		p.Match(ezbpfParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BPF_MAP_TYPE_HASH() antlr.TerminalNode
	BPF_MAP_TYPE_ARRAY() antlr.TerminalNode
	BPF_MAP_TYPE_PERCPU_HASH() antlr.TerminalNode
	BPF_MAP_TYPE_PERCPU_ARRAY() antlr.TerminalNode
	BPF_MAP_TYPE_LRU_HASH() antlr.TerminalNode
	BPF_MAP_TYPE_LRU_PERCPU_HASH() antlr.TerminalNode

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_mapType
	return p
}

func InitEmptyMapTypeContext(p *MapTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_mapType
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) BPF_MAP_TYPE_HASH() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBPF_MAP_TYPE_HASH, 0)
}

func (s *MapTypeContext) BPF_MAP_TYPE_ARRAY() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBPF_MAP_TYPE_ARRAY, 0)
}

func (s *MapTypeContext) BPF_MAP_TYPE_PERCPU_HASH() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBPF_MAP_TYPE_PERCPU_HASH, 0)
}

func (s *MapTypeContext) BPF_MAP_TYPE_PERCPU_ARRAY() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBPF_MAP_TYPE_PERCPU_ARRAY, 0)
}

func (s *MapTypeContext) BPF_MAP_TYPE_LRU_HASH() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBPF_MAP_TYPE_LRU_HASH, 0)
}

func (s *MapTypeContext) BPF_MAP_TYPE_LRU_PERCPU_HASH() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBPF_MAP_TYPE_LRU_PERCPU_HASH, 0)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitMapType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ezbpfParserRULE_mapType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-72)) & ^0x3f) == 0 && ((int64(1)<<(_la-72))&63) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapDeclStmtContext is an interface to support dynamic dispatch.
type IMapDeclStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	MapType() IMapTypeContext
	LT() antlr.TerminalNode
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	DEC_LITERAL() antlr.TerminalNode
	GT() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsMapDeclStmtContext differentiates from other interfaces.
	IsMapDeclStmtContext()
}

type MapDeclStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapDeclStmtContext() *MapDeclStmtContext {
	var p = new(MapDeclStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_mapDeclStmt
	return p
}

func InitEmptyMapDeclStmtContext(p *MapDeclStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_mapDeclStmt
}

func (*MapDeclStmtContext) IsMapDeclStmtContext() {}

func NewMapDeclStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapDeclStmtContext {
	var p = new(MapDeclStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_mapDeclStmt

	return p
}

func (s *MapDeclStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *MapDeclStmtContext) MAP() antlr.TerminalNode {
	return s.GetToken(ezbpfParserMAP, 0)
}

func (s *MapDeclStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *MapDeclStmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *MapDeclStmtContext) MapType() IMapTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *MapDeclStmtContext) LT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLT, 0)
}

func (s *MapDeclStmtContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *MapDeclStmtContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MapDeclStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserCOMMA)
}

func (s *MapDeclStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOMMA, i)
}

func (s *MapDeclStmtContext) DEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserDEC_LITERAL, 0)
}

func (s *MapDeclStmtContext) GT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserGT, 0)
}

func (s *MapDeclStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *MapDeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapDeclStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapDeclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitMapDeclStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) MapDeclStmt() (localctx IMapDeclStmtContext) {
	localctx = NewMapDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ezbpfParserRULE_mapDeclStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(ezbpfParserMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.MapType()
	}
	{
		p.SetState(214)
		p.Match(ezbpfParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Type_()
	}
	{
		p.SetState(216)
		p.Match(ezbpfParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Type_()
	}
	{
		p.SetState(218)
		p.Match(ezbpfParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.Match(ezbpfParserDEC_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(ezbpfParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Match(ezbpfParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDataMemStmtContext is an interface to support dynamic dispatch.
type IStructDataMemStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	SEMI() antlr.TerminalNode

	// IsStructDataMemStmtContext differentiates from other interfaces.
	IsStructDataMemStmtContext()
}

type StructDataMemStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDataMemStmtContext() *StructDataMemStmtContext {
	var p = new(StructDataMemStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structDataMemStmt
	return p
}

func InitEmptyStructDataMemStmtContext(p *StructDataMemStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structDataMemStmt
}

func (*StructDataMemStmtContext) IsStructDataMemStmtContext() {}

func NewStructDataMemStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDataMemStmtContext {
	var p = new(StructDataMemStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_structDataMemStmt

	return p
}

func (s *StructDataMemStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDataMemStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *StructDataMemStmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *StructDataMemStmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructDataMemStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *StructDataMemStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDataMemStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDataMemStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitStructDataMemStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) StructDataMemStmt() (localctx IStructDataMemStmtContext) {
	localctx = NewStructDataMemStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ezbpfParserRULE_structDataMemStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Type_()
	}
	{
		p.SetState(226)
		p.Match(ezbpfParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclStmtContext is an interface to support dynamic dispatch.
type IStructDeclStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRA() antlr.TerminalNode
	RBRA() antlr.TerminalNode
	AllStructDataMemStmt() []IStructDataMemStmtContext
	StructDataMemStmt(i int) IStructDataMemStmtContext

	// IsStructDeclStmtContext differentiates from other interfaces.
	IsStructDeclStmtContext()
}

type StructDeclStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclStmtContext() *StructDeclStmtContext {
	var p = new(StructDeclStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structDeclStmt
	return p
}

func InitEmptyStructDeclStmtContext(p *StructDeclStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structDeclStmt
}

func (*StructDeclStmtContext) IsStructDeclStmtContext() {}

func NewStructDeclStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclStmtContext {
	var p = new(StructDeclStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_structDeclStmt

	return p
}

func (s *StructDeclStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclStmtContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSTRUCT, 0)
}

func (s *StructDeclStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *StructDeclStmtContext) LBRA() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLBRA, 0)
}

func (s *StructDeclStmtContext) RBRA() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRBRA, 0)
}

func (s *StructDeclStmtContext) AllStructDataMemStmt() []IStructDataMemStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructDataMemStmtContext); ok {
			len++
		}
	}

	tst := make([]IStructDataMemStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructDataMemStmtContext); ok {
			tst[i] = t.(IStructDataMemStmtContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclStmtContext) StructDataMemStmt(i int) IStructDataMemStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDataMemStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDataMemStmtContext)
}

func (s *StructDeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitStructDeclStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) StructDeclStmt() (localctx IStructDeclStmtContext) {
	localctx = NewStructDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ezbpfParserRULE_structDeclStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(ezbpfParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Match(ezbpfParserLBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserIDENTIFIER {
		{
			p.SetState(231)
			p.StructDataMemStmt()
		}

		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(237)
		p.Match(ezbpfParserRBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAR() antlr.TerminalNode
	Expression() IExpressionContext
	RPAR() antlr.TerminalNode
	LBRA() antlr.TerminalNode
	Stmts() IStmtsContext
	RBRA() antlr.TerminalNode

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIF, 0)
}

func (s *IfStmtContext) LPAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLPAR, 0)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) RPAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRPAR, 0)
}

func (s *IfStmtContext) LBRA() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLBRA, 0)
}

func (s *IfStmtContext) Stmts() IStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *IfStmtContext) RBRA() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRBRA, 0)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ezbpfParserRULE_ifStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(ezbpfParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Match(ezbpfParserLPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.expression(0)
	}
	{
		p.SetState(242)
		p.Match(ezbpfParserRPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(ezbpfParserLBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Stmts()
	}
	{
		p.SetState(245)
		p.Match(ezbpfParserRBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRETURN, 0)
}

func (s *ReturnStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *ReturnStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ezbpfParserRULE_returnStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(ezbpfParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&4503599627435025) != 0 {
		{
			p.SetState(248)
			p.expression(0)
		}

	}
	{
		p.SetState(251)
		p.Match(ezbpfParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDeclStmtContext is an interface to support dynamic dispatch.
type IFuncDeclStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAR() antlr.TerminalNode
	RPAR() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	LBRA() antlr.TerminalNode
	Stmts() IStmtsContext
	RBRA() antlr.TerminalNode
	Params() IParamsContext

	// IsFuncDeclStmtContext differentiates from other interfaces.
	IsFuncDeclStmtContext()
}

type FuncDeclStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclStmtContext() *FuncDeclStmtContext {
	var p = new(FuncDeclStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_funcDeclStmt
	return p
}

func InitEmptyFuncDeclStmtContext(p *FuncDeclStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_funcDeclStmt
}

func (*FuncDeclStmtContext) IsFuncDeclStmtContext() {}

func NewFuncDeclStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclStmtContext {
	var p = new(FuncDeclStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_funcDeclStmt

	return p
}

func (s *FuncDeclStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclStmtContext) FN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserFN, 0)
}

func (s *FuncDeclStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *FuncDeclStmtContext) LPAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLPAR, 0)
}

func (s *FuncDeclStmtContext) RPAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRPAR, 0)
}

func (s *FuncDeclStmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *FuncDeclStmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncDeclStmtContext) LBRA() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLBRA, 0)
}

func (s *FuncDeclStmtContext) Stmts() IStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *FuncDeclStmtContext) RBRA() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRBRA, 0)
}

func (s *FuncDeclStmtContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FuncDeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitFuncDeclStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) FuncDeclStmt() (localctx IFuncDeclStmtContext) {
	localctx = NewFuncDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ezbpfParserRULE_funcDeclStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(ezbpfParserFN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Match(ezbpfParserLPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ezbpfParserIDENTIFIER {
		{
			p.SetState(256)
			p.Params()
		}

	}
	{
		p.SetState(259)
		p.Match(ezbpfParserRPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Type_()
	}
	{
		p.SetState(262)
		p.Match(ezbpfParserLBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Stmts()
	}
	{
		p.SetState(264)
		p.Match(ezbpfParserRBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncCallStmtContext is an interface to support dynamic dispatch.
type IFuncCallStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncCallExpression() IFuncCallExpressionContext
	SEMI() antlr.TerminalNode

	// IsFuncCallStmtContext differentiates from other interfaces.
	IsFuncCallStmtContext()
}

type FuncCallStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallStmtContext() *FuncCallStmtContext {
	var p = new(FuncCallStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_funcCallStmt
	return p
}

func InitEmptyFuncCallStmtContext(p *FuncCallStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_funcCallStmt
}

func (*FuncCallStmtContext) IsFuncCallStmtContext() {}

func NewFuncCallStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallStmtContext {
	var p = new(FuncCallStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_funcCallStmt

	return p
}

func (s *FuncCallStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallStmtContext) FuncCallExpression() IFuncCallExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallExpressionContext)
}

func (s *FuncCallStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *FuncCallStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitFuncCallStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) FuncCallStmt() (localctx IFuncCallStmtContext) {
	localctx = NewFuncCallStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ezbpfParserRULE_funcCallStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.FuncCallExpression()
	}
	{
		p.SetState(267)
		p.Match(ezbpfParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarInitStmt() IVarInitStmtContext
	VarDeclStmt() IVarDeclStmtContext
	ConstDeclStmt() IConstDeclStmtContext
	MapDeclStmt() IMapDeclStmtContext
	StructDeclStmt() IStructDeclStmtContext
	IfStmt() IIfStmtContext
	ReturnStmt() IReturnStmtContext
	FuncDeclStmt() IFuncDeclStmtContext
	FuncCallStmt() IFuncCallStmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) VarInitStmt() IVarInitStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarInitStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarInitStmtContext)
}

func (s *StmtContext) VarDeclStmt() IVarDeclStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclStmtContext)
}

func (s *StmtContext) ConstDeclStmt() IConstDeclStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDeclStmtContext)
}

func (s *StmtContext) MapDeclStmt() IMapDeclStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapDeclStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapDeclStmtContext)
}

func (s *StmtContext) StructDeclStmt() IStructDeclStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclStmtContext)
}

func (s *StmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StmtContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StmtContext) FuncDeclStmt() IFuncDeclStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclStmtContext)
}

func (s *StmtContext) FuncCallStmt() IFuncCallStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallStmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ezbpfParserRULE_stmt)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(269)
			p.VarInitStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			p.VarDeclStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(271)
			p.ConstDeclStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(272)
			p.MapDeclStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(273)
			p.StructDeclStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(274)
			p.IfStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(275)
			p.ReturnStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(276)
			p.FuncDeclStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(277)
			p.FuncCallStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtsContext is an interface to support dynamic dispatch.
type IStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsStmtsContext differentiates from other interfaces.
	IsStmtsContext()
}

type StmtsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtsContext() *StmtsContext {
	var p = new(StmtsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_stmts
	return p
}

func InitEmptyStmtsContext(p *StmtsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_stmts
}

func (*StmtsContext) IsStmtsContext() {}

func NewStmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtsContext {
	var p = new(StmtsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_stmts

	return p
}

func (s *StmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtsContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtsContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *StmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitStmts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) Stmts() (localctx IStmtsContext) {
	localctx = NewStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ezbpfParserRULE_stmts)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-52)) & ^0x3f) == 0 && ((int64(1)<<(_la-52))&67748355) != 0 {
		{
			p.SetState(280)
			p.Stmt()
		}

		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStructDeclStmt() []IStructDeclStmtContext
	StructDeclStmt(i int) IStructDeclStmtContext
	AllMapDeclStmt() []IMapDeclStmtContext
	MapDeclStmt(i int) IMapDeclStmtContext
	AllFuncDeclStmt() []IFuncDeclStmtContext
	FuncDeclStmt(i int) IFuncDeclStmtContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStructDeclStmt() []IStructDeclStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructDeclStmtContext); ok {
			len++
		}
	}

	tst := make([]IStructDeclStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructDeclStmtContext); ok {
			tst[i] = t.(IStructDeclStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) StructDeclStmt(i int) IStructDeclStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclStmtContext)
}

func (s *ProgContext) AllMapDeclStmt() []IMapDeclStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapDeclStmtContext); ok {
			len++
		}
	}

	tst := make([]IMapDeclStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapDeclStmtContext); ok {
			tst[i] = t.(IMapDeclStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) MapDeclStmt(i int) IMapDeclStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapDeclStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapDeclStmtContext)
}

func (s *ProgContext) AllFuncDeclStmt() []IFuncDeclStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDeclStmtContext); ok {
			len++
		}
	}

	tst := make([]IFuncDeclStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDeclStmtContext); ok {
			tst[i] = t.(IFuncDeclStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) FuncDeclStmt(i int) IFuncDeclStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclStmtContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ezbpfVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ezbpfParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ezbpfParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserSTRUCT {
		{
			p.SetState(286)
			p.StructDeclStmt()
		}

		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserMAP {
		{
			p.SetState(292)
			p.MapDeclStmt()
		}

		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ezbpfParserFN {
		{
			p.SetState(298)
			p.FuncDeclStmt()
		}

		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ezbpfParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ezbpfParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 13)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
