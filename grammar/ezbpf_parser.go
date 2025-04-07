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
		"'s32'", "'s64'", "'var'", "'const'", "'char'", "'int'", "'long'", "'short'",
		"'uint'", "'void'", "'struct'", "'fn'", "'return'", "'if'", "'elif'",
		"'else'", "'map'", "'BPF_MAP_TYPE_HASH'", "'BPF_MAP_TYPE_ARRAY'", "'BPF_MAP_TYPE_PERCPU_HASH'",
		"'BPF_MAP_TYPE_PERCPU_ARRAY'", "'BPF_MAP_TYPE_LRU_HASH'", "'BPF_MAP_TYPE_LRU_PERCPU_HASH'",
	}
	staticData.SymbolicNames = []string{
		"", "ADD", "SUB", "MUL", "DIV", "MOD", "BIT_AND", "BIT_OR", "BIT_XOR",
		"LSHIFT", "RSHIFT", "EQ", "NEQ", "LT", "GT", "LTE", "GTE", "AND", "OR",
		"NOT", "ASSIGN", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN",
		"MOD_ASSIGN", "LPAR", "RPAR", "LBRA", "RBRA", "LSQ", "RSQ", "COMMA",
		"SEMI", "COLON", "DOT", "HEX_LITERAL", "OCT_LITERAL", "BIN_LITERAL",
		"DEC_LITERAL", "CHAR_LITERAL", "STRING_LITERAL", "UPTR32", "UPTR64",
		"SPTR32", "SPTR64", "U32", "U64", "S32", "S64", "VAR", "CONST", "CHAR",
		"INT", "LONG", "SHORT", "UINT", "VOID", "STRUCT", "FN", "RETURN", "IF",
		"ELSEIF", "ELSE", "MAP", "BPF_MAP_TYPE_HASH", "BPF_MAP_TYPE_ARRAY",
		"BPF_MAP_TYPE_PERCPU_HASH", "BPF_MAP_TYPE_PERCPU_ARRAY", "BPF_MAP_TYPE_LRU_HASH",
		"BPF_MAP_TYPE_LRU_PERCPU_HASH", "IDENTIFIER", "WS", "COMMENT", "MULTI_COMMENT",
	}
	staticData.RuleNames = []string{
		"type", "assign", "compare", "expression", "structFieldAssign", "arg",
		"args", "param", "params", "varInitStmt", "varDeclStmt", "constDeclStmt",
		"mapType", "mapDeclStmt", "structDataMemStmt", "structDeclStmt", "ifStmt",
		"returnStmt", "funcDeclStmt", "stmts", "prog",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 74, 264, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 64, 8, 3, 10,
		3, 12, 3, 67, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 73, 8, 3, 1, 3, 3, 3,
		76, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 5, 3, 115, 8, 3, 10, 3, 12, 3, 118, 9, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 127, 8, 4, 10, 4, 12, 4, 130, 9, 4, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 137, 8, 6, 10, 6, 12, 6, 140, 9, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 149, 8, 8, 10, 8, 12, 8, 152,
		9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 200, 8, 15, 10, 15, 12, 15, 203, 9,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 3, 17, 217, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 3, 18, 225, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 242, 8,
		19, 10, 19, 12, 19, 245, 9, 19, 1, 20, 5, 20, 248, 8, 20, 10, 20, 12, 20,
		251, 9, 20, 1, 20, 5, 20, 254, 8, 20, 10, 20, 12, 20, 257, 9, 20, 1, 20,
		4, 20, 260, 8, 20, 11, 20, 12, 20, 261, 1, 20, 0, 1, 6, 21, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 0, 4,
		3, 0, 42, 49, 52, 56, 71, 71, 1, 0, 20, 25, 1, 0, 11, 19, 1, 0, 65, 70,
		282, 0, 42, 1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 6, 75, 1,
		0, 0, 0, 8, 119, 1, 0, 0, 0, 10, 131, 1, 0, 0, 0, 12, 133, 1, 0, 0, 0,
		14, 141, 1, 0, 0, 0, 16, 145, 1, 0, 0, 0, 18, 153, 1, 0, 0, 0, 20, 159,
		1, 0, 0, 0, 22, 167, 1, 0, 0, 0, 24, 175, 1, 0, 0, 0, 26, 177, 1, 0, 0,
		0, 28, 190, 1, 0, 0, 0, 30, 195, 1, 0, 0, 0, 32, 206, 1, 0, 0, 0, 34, 214,
		1, 0, 0, 0, 36, 220, 1, 0, 0, 0, 38, 243, 1, 0, 0, 0, 40, 249, 1, 0, 0,
		0, 42, 43, 7, 0, 0, 0, 43, 1, 1, 0, 0, 0, 44, 45, 7, 1, 0, 0, 45, 3, 1,
		0, 0, 0, 46, 47, 7, 2, 0, 0, 47, 5, 1, 0, 0, 0, 48, 49, 6, 3, -1, 0, 49,
		76, 5, 36, 0, 0, 50, 76, 5, 37, 0, 0, 51, 76, 5, 38, 0, 0, 52, 76, 5, 39,
		0, 0, 53, 76, 5, 40, 0, 0, 54, 76, 5, 41, 0, 0, 55, 76, 5, 71, 0, 0, 56,
		57, 5, 26, 0, 0, 57, 58, 3, 6, 3, 0, 58, 59, 5, 27, 0, 0, 59, 76, 1, 0,
		0, 0, 60, 61, 5, 71, 0, 0, 61, 65, 5, 28, 0, 0, 62, 64, 3, 8, 4, 0, 63,
		62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0,
		0, 66, 68, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 76, 5, 29, 0, 0, 69, 70,
		5, 71, 0, 0, 70, 72, 5, 26, 0, 0, 71, 73, 3, 12, 6, 0, 72, 71, 1, 0, 0,
		0, 72, 73, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 76, 5, 27, 0, 0, 75, 48,
		1, 0, 0, 0, 75, 50, 1, 0, 0, 0, 75, 51, 1, 0, 0, 0, 75, 52, 1, 0, 0, 0,
		75, 53, 1, 0, 0, 0, 75, 54, 1, 0, 0, 0, 75, 55, 1, 0, 0, 0, 75, 56, 1,
		0, 0, 0, 75, 60, 1, 0, 0, 0, 75, 69, 1, 0, 0, 0, 76, 116, 1, 0, 0, 0, 77,
		78, 10, 14, 0, 0, 78, 79, 5, 1, 0, 0, 79, 115, 3, 6, 3, 15, 80, 81, 10,
		13, 0, 0, 81, 82, 5, 2, 0, 0, 82, 115, 3, 6, 3, 14, 83, 84, 10, 12, 0,
		0, 84, 85, 5, 3, 0, 0, 85, 115, 3, 6, 3, 13, 86, 87, 10, 11, 0, 0, 87,
		88, 5, 4, 0, 0, 88, 115, 3, 6, 3, 12, 89, 90, 10, 10, 0, 0, 90, 91, 5,
		5, 0, 0, 91, 115, 3, 6, 3, 11, 92, 93, 10, 9, 0, 0, 93, 94, 5, 6, 0, 0,
		94, 115, 3, 6, 3, 10, 95, 96, 10, 8, 0, 0, 96, 97, 5, 7, 0, 0, 97, 115,
		3, 6, 3, 9, 98, 99, 10, 7, 0, 0, 99, 100, 5, 8, 0, 0, 100, 115, 3, 6, 3,
		8, 101, 102, 10, 6, 0, 0, 102, 103, 5, 9, 0, 0, 103, 115, 3, 6, 3, 7, 104,
		105, 10, 5, 0, 0, 105, 106, 5, 10, 0, 0, 106, 115, 3, 6, 3, 6, 107, 108,
		10, 4, 0, 0, 108, 109, 3, 4, 2, 0, 109, 110, 3, 6, 3, 5, 110, 115, 1, 0,
		0, 0, 111, 112, 10, 15, 0, 0, 112, 113, 5, 35, 0, 0, 113, 115, 5, 71, 0,
		0, 114, 77, 1, 0, 0, 0, 114, 80, 1, 0, 0, 0, 114, 83, 1, 0, 0, 0, 114,
		86, 1, 0, 0, 0, 114, 89, 1, 0, 0, 0, 114, 92, 1, 0, 0, 0, 114, 95, 1, 0,
		0, 0, 114, 98, 1, 0, 0, 0, 114, 101, 1, 0, 0, 0, 114, 104, 1, 0, 0, 0,
		114, 107, 1, 0, 0, 0, 114, 111, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116,
		114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 7, 1, 0, 0, 0, 118, 116, 1,
		0, 0, 0, 119, 120, 5, 71, 0, 0, 120, 121, 5, 34, 0, 0, 121, 128, 3, 6,
		3, 0, 122, 123, 5, 32, 0, 0, 123, 124, 5, 71, 0, 0, 124, 125, 5, 34, 0,
		0, 125, 127, 3, 6, 3, 0, 126, 122, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128,
		126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 9, 1, 0, 0, 0, 130, 128, 1,
		0, 0, 0, 131, 132, 3, 6, 3, 0, 132, 11, 1, 0, 0, 0, 133, 138, 3, 10, 5,
		0, 134, 135, 5, 32, 0, 0, 135, 137, 3, 10, 5, 0, 136, 134, 1, 0, 0, 0,
		137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		13, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 5, 71, 0, 0, 142, 143,
		5, 34, 0, 0, 143, 144, 3, 0, 0, 0, 144, 15, 1, 0, 0, 0, 145, 150, 3, 14,
		7, 0, 146, 147, 5, 32, 0, 0, 147, 149, 3, 14, 7, 0, 148, 146, 1, 0, 0,
		0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151,
		17, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 154, 5, 50, 0, 0, 154, 155,
		5, 71, 0, 0, 155, 156, 5, 34, 0, 0, 156, 157, 3, 0, 0, 0, 157, 158, 5,
		33, 0, 0, 158, 19, 1, 0, 0, 0, 159, 160, 5, 50, 0, 0, 160, 161, 5, 71,
		0, 0, 161, 162, 5, 34, 0, 0, 162, 163, 3, 0, 0, 0, 163, 164, 3, 2, 1, 0,
		164, 165, 3, 6, 3, 0, 165, 166, 5, 33, 0, 0, 166, 21, 1, 0, 0, 0, 167,
		168, 5, 51, 0, 0, 168, 169, 5, 71, 0, 0, 169, 170, 5, 34, 0, 0, 170, 171,
		3, 0, 0, 0, 171, 172, 3, 2, 1, 0, 172, 173, 3, 6, 3, 0, 173, 174, 5, 33,
		0, 0, 174, 23, 1, 0, 0, 0, 175, 176, 7, 3, 0, 0, 176, 25, 1, 0, 0, 0, 177,
		178, 5, 64, 0, 0, 178, 179, 5, 71, 0, 0, 179, 180, 5, 34, 0, 0, 180, 181,
		3, 24, 12, 0, 181, 182, 5, 13, 0, 0, 182, 183, 3, 0, 0, 0, 183, 184, 5,
		32, 0, 0, 184, 185, 3, 0, 0, 0, 185, 186, 5, 32, 0, 0, 186, 187, 5, 39,
		0, 0, 187, 188, 5, 14, 0, 0, 188, 189, 5, 33, 0, 0, 189, 27, 1, 0, 0, 0,
		190, 191, 5, 71, 0, 0, 191, 192, 5, 34, 0, 0, 192, 193, 3, 0, 0, 0, 193,
		194, 5, 33, 0, 0, 194, 29, 1, 0, 0, 0, 195, 196, 5, 58, 0, 0, 196, 197,
		5, 71, 0, 0, 197, 201, 5, 28, 0, 0, 198, 200, 3, 28, 14, 0, 199, 198, 1,
		0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0,
		0, 202, 204, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 204, 205, 5, 29, 0, 0, 205,
		31, 1, 0, 0, 0, 206, 207, 5, 61, 0, 0, 207, 208, 5, 26, 0, 0, 208, 209,
		3, 6, 3, 0, 209, 210, 5, 27, 0, 0, 210, 211, 5, 28, 0, 0, 211, 212, 3,
		38, 19, 0, 212, 213, 5, 29, 0, 0, 213, 33, 1, 0, 0, 0, 214, 216, 5, 60,
		0, 0, 215, 217, 3, 6, 3, 0, 216, 215, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0,
		217, 218, 1, 0, 0, 0, 218, 219, 5, 33, 0, 0, 219, 35, 1, 0, 0, 0, 220,
		221, 5, 59, 0, 0, 221, 222, 5, 71, 0, 0, 222, 224, 5, 26, 0, 0, 223, 225,
		3, 16, 8, 0, 224, 223, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 226, 1, 0,
		0, 0, 226, 227, 5, 27, 0, 0, 227, 228, 5, 34, 0, 0, 228, 229, 3, 0, 0,
		0, 229, 230, 5, 28, 0, 0, 230, 231, 3, 38, 19, 0, 231, 232, 5, 29, 0, 0,
		232, 37, 1, 0, 0, 0, 233, 242, 3, 18, 9, 0, 234, 242, 3, 20, 10, 0, 235,
		242, 3, 22, 11, 0, 236, 242, 3, 26, 13, 0, 237, 242, 3, 30, 15, 0, 238,
		242, 3, 32, 16, 0, 239, 242, 3, 34, 17, 0, 240, 242, 3, 36, 18, 0, 241,
		233, 1, 0, 0, 0, 241, 234, 1, 0, 0, 0, 241, 235, 1, 0, 0, 0, 241, 236,
		1, 0, 0, 0, 241, 237, 1, 0, 0, 0, 241, 238, 1, 0, 0, 0, 241, 239, 1, 0,
		0, 0, 241, 240, 1, 0, 0, 0, 242, 245, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0,
		243, 244, 1, 0, 0, 0, 244, 39, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 248,
		3, 30, 15, 0, 247, 246, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249, 247, 1,
		0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 255, 1, 0, 0, 0, 251, 249, 1, 0, 0,
		0, 252, 254, 3, 26, 13, 0, 253, 252, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0,
		255, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257,
		255, 1, 0, 0, 0, 258, 260, 3, 36, 18, 0, 259, 258, 1, 0, 0, 0, 260, 261,
		1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 41, 1, 0,
		0, 0, 16, 65, 72, 75, 114, 116, 128, 138, 150, 201, 216, 224, 241, 243,
		249, 255, 261,
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
	ezbpfParserVAR                          = 50
	ezbpfParserCONST                        = 51
	ezbpfParserCHAR                         = 52
	ezbpfParserINT                          = 53
	ezbpfParserLONG                         = 54
	ezbpfParserSHORT                        = 55
	ezbpfParserUINT                         = 56
	ezbpfParserVOID                         = 57
	ezbpfParserSTRUCT                       = 58
	ezbpfParserFN                           = 59
	ezbpfParserRETURN                       = 60
	ezbpfParserIF                           = 61
	ezbpfParserELSEIF                       = 62
	ezbpfParserELSE                         = 63
	ezbpfParserMAP                          = 64
	ezbpfParserBPF_MAP_TYPE_HASH            = 65
	ezbpfParserBPF_MAP_TYPE_ARRAY           = 66
	ezbpfParserBPF_MAP_TYPE_PERCPU_HASH     = 67
	ezbpfParserBPF_MAP_TYPE_PERCPU_ARRAY    = 68
	ezbpfParserBPF_MAP_TYPE_LRU_HASH        = 69
	ezbpfParserBPF_MAP_TYPE_LRU_PERCPU_HASH = 70
	ezbpfParserIDENTIFIER                   = 71
	ezbpfParserWS                           = 72
	ezbpfParserCOMMENT                      = 73
	ezbpfParserMULTI_COMMENT                = 74
)

// ezbpfParser rules.
const (
	ezbpfParserRULE_type              = 0
	ezbpfParserRULE_assign            = 1
	ezbpfParserRULE_compare           = 2
	ezbpfParserRULE_expression        = 3
	ezbpfParserRULE_structFieldAssign = 4
	ezbpfParserRULE_arg               = 5
	ezbpfParserRULE_args              = 6
	ezbpfParserRULE_param             = 7
	ezbpfParserRULE_params            = 8
	ezbpfParserRULE_varInitStmt       = 9
	ezbpfParserRULE_varDeclStmt       = 10
	ezbpfParserRULE_constDeclStmt     = 11
	ezbpfParserRULE_mapType           = 12
	ezbpfParserRULE_mapDeclStmt       = 13
	ezbpfParserRULE_structDataMemStmt = 14
	ezbpfParserRULE_structDeclStmt    = 15
	ezbpfParserRULE_ifStmt            = 16
	ezbpfParserRULE_returnStmt        = 17
	ezbpfParserRULE_funcDeclStmt      = 18
	ezbpfParserRULE_stmts             = 19
	ezbpfParserRULE_prog              = 20
)

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	U32() antlr.TerminalNode
	U64() antlr.TerminalNode
	UPTR32() antlr.TerminalNode
	UPTR64() antlr.TerminalNode
	S32() antlr.TerminalNode
	S64() antlr.TerminalNode
	SPTR32() antlr.TerminalNode
	SPTR64() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	SHORT() antlr.TerminalNode
	LONG() antlr.TerminalNode
	UINT() antlr.TerminalNode
	INT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

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

func (s *TypeContext) U32() antlr.TerminalNode {
	return s.GetToken(ezbpfParserU32, 0)
}

func (s *TypeContext) U64() antlr.TerminalNode {
	return s.GetToken(ezbpfParserU64, 0)
}

func (s *TypeContext) UPTR32() antlr.TerminalNode {
	return s.GetToken(ezbpfParserUPTR32, 0)
}

func (s *TypeContext) UPTR64() antlr.TerminalNode {
	return s.GetToken(ezbpfParserUPTR64, 0)
}

func (s *TypeContext) S32() antlr.TerminalNode {
	return s.GetToken(ezbpfParserS32, 0)
}

func (s *TypeContext) S64() antlr.TerminalNode {
	return s.GetToken(ezbpfParserS64, 0)
}

func (s *TypeContext) SPTR32() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSPTR32, 0)
}

func (s *TypeContext) SPTR64() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSPTR64, 0)
}

func (s *TypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCHAR, 0)
}

func (s *TypeContext) SHORT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSHORT, 0)
}

func (s *TypeContext) LONG() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLONG, 0)
}

func (s *TypeContext) UINT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserUINT, 0)
}

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserINT, 0)
}

func (s *TypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *ezbpfParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ezbpfParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-42)) & ^0x3f) == 0 && ((int64(1)<<(_la-42))&536902911) != 0) {
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

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *ezbpfParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ezbpfParserRULE_assign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
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

func (s *CompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterCompare(s)
	}
}

func (s *CompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitCompare(s)
	}
}

func (p *ezbpfParser) Compare() (localctx ICompareContext) {
	localctx = NewCompareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ezbpfParserRULE_compare)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
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

	// Getter signatures
	HEX_LITERAL() antlr.TerminalNode
	OCT_LITERAL() antlr.TerminalNode
	BIN_LITERAL() antlr.TerminalNode
	DEC_LITERAL() antlr.TerminalNode
	CHAR_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAR() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	RPAR() antlr.TerminalNode
	LBRA() antlr.TerminalNode
	RBRA() antlr.TerminalNode
	AllStructFieldAssign() []IStructFieldAssignContext
	StructFieldAssign(i int) IStructFieldAssignContext
	Args() IArgsContext
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	BIT_AND() antlr.TerminalNode
	BIT_OR() antlr.TerminalNode
	BIT_XOR() antlr.TerminalNode
	LSHIFT() antlr.TerminalNode
	RSHIFT() antlr.TerminalNode
	Compare() ICompareContext
	DOT() antlr.TerminalNode

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

func (s *ExpressionContext) HEX_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserHEX_LITERAL, 0)
}

func (s *ExpressionContext) OCT_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserOCT_LITERAL, 0)
}

func (s *ExpressionContext) BIN_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBIN_LITERAL, 0)
}

func (s *ExpressionContext) DEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserDEC_LITERAL, 0)
}

func (s *ExpressionContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCHAR_LITERAL, 0)
}

func (s *ExpressionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSTRING_LITERAL, 0)
}

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIDENTIFIER, 0)
}

func (s *ExpressionContext) LPAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLPAR, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionContext) RPAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRPAR, 0)
}

func (s *ExpressionContext) LBRA() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLBRA, 0)
}

func (s *ExpressionContext) RBRA() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRBRA, 0)
}

func (s *ExpressionContext) AllStructFieldAssign() []IStructFieldAssignContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructFieldAssignContext); ok {
			len++
		}
	}

	tst := make([]IStructFieldAssignContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructFieldAssignContext); ok {
			tst[i] = t.(IStructFieldAssignContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) StructFieldAssign(i int) IStructFieldAssignContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldAssignContext); ok {
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

	return t.(IStructFieldAssignContext)
}

func (s *ExpressionContext) Args() IArgsContext {
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

func (s *ExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(ezbpfParserADD, 0)
}

func (s *ExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSUB, 0)
}

func (s *ExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserMUL, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(ezbpfParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(ezbpfParserMOD, 0)
}

func (s *ExpressionContext) BIT_AND() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBIT_AND, 0)
}

func (s *ExpressionContext) BIT_OR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBIT_OR, 0)
}

func (s *ExpressionContext) BIT_XOR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBIT_XOR, 0)
}

func (s *ExpressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLSHIFT, 0)
}

func (s *ExpressionContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRSHIFT, 0)
}

func (s *ExpressionContext) Compare() ICompareContext {
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

func (s *ExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserDOT, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitExpression(s)
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
	_startState := 6
	p.EnterRecursionRule(localctx, 6, ezbpfParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(49)
			p.Match(ezbpfParserHEX_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(50)
			p.Match(ezbpfParserOCT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(51)
			p.Match(ezbpfParserBIN_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(52)
			p.Match(ezbpfParserDEC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(53)
			p.Match(ezbpfParserCHAR_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		{
			p.SetState(54)
			p.Match(ezbpfParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		{
			p.SetState(55)
			p.Match(ezbpfParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		{
			p.SetState(56)
			p.Match(ezbpfParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.expression(0)
		}
		{
			p.SetState(58)
			p.Match(ezbpfParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		{
			p.SetState(60)
			p.Match(ezbpfParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Match(ezbpfParserLBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ezbpfParserIDENTIFIER {
			{
				p.SetState(62)
				p.StructFieldAssign()
			}

			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(68)
			p.Match(ezbpfParserRBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		{
			p.SetState(69)
			p.Match(ezbpfParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(ezbpfParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&35184372153345) != 0 {
			{
				p.SetState(71)
				p.Args()
			}

		}
		{
			p.SetState(74)
			p.Match(ezbpfParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(116)
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
			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(78)
					p.Match(ezbpfParserADD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(79)
					p.expression(15)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(81)
					p.Match(ezbpfParserSUB)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(82)
					p.expression(14)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(84)
					p.Match(ezbpfParserMUL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(85)
					p.expression(13)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(87)
					p.Match(ezbpfParserDIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(88)
					p.expression(12)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(90)
					p.Match(ezbpfParserMOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(91)
					p.expression(11)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(93)
					p.Match(ezbpfParserBIT_AND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(94)
					p.expression(10)
				}

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(96)
					p.Match(ezbpfParserBIT_OR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(97)
					p.expression(9)
				}

			case 8:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(99)
					p.Match(ezbpfParserBIT_XOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(100)
					p.expression(8)
				}

			case 9:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(102)
					p.Match(ezbpfParserLSHIFT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(103)
					p.expression(7)
				}

			case 10:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(105)
					p.Match(ezbpfParserRSHIFT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(106)
					p.expression(6)
				}

			case 11:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(108)
					p.Compare()
				}
				{
					p.SetState(109)
					p.expression(5)
				}

			case 12:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ezbpfParserRULE_expression)
				p.SetState(111)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(112)
					p.Match(ezbpfParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(113)
					p.Match(ezbpfParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(118)
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

func (s *StructFieldAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterStructFieldAssign(s)
	}
}

func (s *StructFieldAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitStructFieldAssign(s)
	}
}

func (p *ezbpfParser) StructFieldAssign() (localctx IStructFieldAssignContext) {
	localctx = NewStructFieldAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ezbpfParserRULE_structFieldAssign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.expression(0)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserCOMMA {
		{
			p.SetState(122)
			p.Match(ezbpfParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(ezbpfParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(ezbpfParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.expression(0)
		}

		p.SetState(130)
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

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitArg(s)
	}
}

func (p *ezbpfParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ezbpfParserRULE_arg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
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

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *ezbpfParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ezbpfParserRULE_args)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Arg()
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserCOMMA {
		{
			p.SetState(134)
			p.Match(ezbpfParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Arg()
		}

		p.SetState(140)
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

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *ezbpfParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ezbpfParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
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

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitParams(s)
	}
}

func (p *ezbpfParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ezbpfParserRULE_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Param()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserCOMMA {
		{
			p.SetState(146)
			p.Match(ezbpfParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Param()
		}

		p.SetState(152)
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

func (s *VarInitStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterVarInitStmt(s)
	}
}

func (s *VarInitStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitVarInitStmt(s)
	}
}

func (p *ezbpfParser) VarInitStmt() (localctx IVarInitStmtContext) {
	localctx = NewVarInitStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ezbpfParserRULE_varInitStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(ezbpfParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Type_()
	}
	{
		p.SetState(157)
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

func (s *VarDeclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterVarDeclStmt(s)
	}
}

func (s *VarDeclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitVarDeclStmt(s)
	}
}

func (p *ezbpfParser) VarDeclStmt() (localctx IVarDeclStmtContext) {
	localctx = NewVarDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ezbpfParserRULE_varDeclStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(ezbpfParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Type_()
	}
	{
		p.SetState(163)
		p.Assign()
	}
	{
		p.SetState(164)
		p.expression(0)
	}
	{
		p.SetState(165)
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

func (s *ConstDeclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterConstDeclStmt(s)
	}
}

func (s *ConstDeclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitConstDeclStmt(s)
	}
}

func (p *ezbpfParser) ConstDeclStmt() (localctx IConstDeclStmtContext) {
	localctx = NewConstDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ezbpfParserRULE_constDeclStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(ezbpfParserCONST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Type_()
	}
	{
		p.SetState(171)
		p.Assign()
	}
	{
		p.SetState(172)
		p.expression(0)
	}
	{
		p.SetState(173)
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

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (p *ezbpfParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ezbpfParserRULE_mapType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
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

func (s *MapDeclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterMapDeclStmt(s)
	}
}

func (s *MapDeclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitMapDeclStmt(s)
	}
}

func (p *ezbpfParser) MapDeclStmt() (localctx IMapDeclStmtContext) {
	localctx = NewMapDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ezbpfParserRULE_mapDeclStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(ezbpfParserMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.MapType()
	}
	{
		p.SetState(181)
		p.Match(ezbpfParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Type_()
	}
	{
		p.SetState(183)
		p.Match(ezbpfParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Type_()
	}
	{
		p.SetState(185)
		p.Match(ezbpfParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(ezbpfParserDEC_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(ezbpfParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
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

func (s *StructDataMemStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterStructDataMemStmt(s)
	}
}

func (s *StructDataMemStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitStructDataMemStmt(s)
	}
}

func (p *ezbpfParser) StructDataMemStmt() (localctx IStructDataMemStmtContext) {
	localctx = NewStructDataMemStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ezbpfParserRULE_structDataMemStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
		p.Type_()
	}
	{
		p.SetState(193)
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

func (s *StructDeclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterStructDeclStmt(s)
	}
}

func (s *StructDeclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitStructDeclStmt(s)
	}
}

func (p *ezbpfParser) StructDeclStmt() (localctx IStructDeclStmtContext) {
	localctx = NewStructDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ezbpfParserRULE_structDeclStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(ezbpfParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Match(ezbpfParserLBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserIDENTIFIER {
		{
			p.SetState(198)
			p.StructDataMemStmt()
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(204)
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

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (p *ezbpfParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ezbpfParserRULE_ifStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(ezbpfParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Match(ezbpfParserLPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.expression(0)
	}
	{
		p.SetState(209)
		p.Match(ezbpfParserRPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Match(ezbpfParserLBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Stmts()
	}
	{
		p.SetState(212)
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

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (p *ezbpfParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ezbpfParserRULE_returnStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(ezbpfParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&35184372153345) != 0 {
		{
			p.SetState(215)
			p.expression(0)
		}

	}
	{
		p.SetState(218)
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

func (s *FuncDeclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterFuncDeclStmt(s)
	}
}

func (s *FuncDeclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitFuncDeclStmt(s)
	}
}

func (p *ezbpfParser) FuncDeclStmt() (localctx IFuncDeclStmtContext) {
	localctx = NewFuncDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ezbpfParserRULE_funcDeclStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(ezbpfParserFN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Match(ezbpfParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(ezbpfParserLPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ezbpfParserIDENTIFIER {
		{
			p.SetState(223)
			p.Params()
		}

	}
	{
		p.SetState(226)
		p.Match(ezbpfParserRPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Type_()
	}
	{
		p.SetState(229)
		p.Match(ezbpfParserLBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Stmts()
	}
	{
		p.SetState(231)
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

// IStmtsContext is an interface to support dynamic dispatch.
type IStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVarInitStmt() []IVarInitStmtContext
	VarInitStmt(i int) IVarInitStmtContext
	AllVarDeclStmt() []IVarDeclStmtContext
	VarDeclStmt(i int) IVarDeclStmtContext
	AllConstDeclStmt() []IConstDeclStmtContext
	ConstDeclStmt(i int) IConstDeclStmtContext
	AllMapDeclStmt() []IMapDeclStmtContext
	MapDeclStmt(i int) IMapDeclStmtContext
	AllStructDeclStmt() []IStructDeclStmtContext
	StructDeclStmt(i int) IStructDeclStmtContext
	AllIfStmt() []IIfStmtContext
	IfStmt(i int) IIfStmtContext
	AllReturnStmt() []IReturnStmtContext
	ReturnStmt(i int) IReturnStmtContext
	AllFuncDeclStmt() []IFuncDeclStmtContext
	FuncDeclStmt(i int) IFuncDeclStmtContext

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

func (s *StmtsContext) AllVarInitStmt() []IVarInitStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarInitStmtContext); ok {
			len++
		}
	}

	tst := make([]IVarInitStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarInitStmtContext); ok {
			tst[i] = t.(IVarInitStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtsContext) VarInitStmt(i int) IVarInitStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarInitStmtContext); ok {
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

	return t.(IVarInitStmtContext)
}

func (s *StmtsContext) AllVarDeclStmt() []IVarDeclStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclStmtContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclStmtContext); ok {
			tst[i] = t.(IVarDeclStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtsContext) VarDeclStmt(i int) IVarDeclStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclStmtContext); ok {
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

	return t.(IVarDeclStmtContext)
}

func (s *StmtsContext) AllConstDeclStmt() []IConstDeclStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstDeclStmtContext); ok {
			len++
		}
	}

	tst := make([]IConstDeclStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstDeclStmtContext); ok {
			tst[i] = t.(IConstDeclStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtsContext) ConstDeclStmt(i int) IConstDeclStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclStmtContext); ok {
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

	return t.(IConstDeclStmtContext)
}

func (s *StmtsContext) AllMapDeclStmt() []IMapDeclStmtContext {
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

func (s *StmtsContext) MapDeclStmt(i int) IMapDeclStmtContext {
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

func (s *StmtsContext) AllStructDeclStmt() []IStructDeclStmtContext {
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

func (s *StmtsContext) StructDeclStmt(i int) IStructDeclStmtContext {
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

func (s *StmtsContext) AllIfStmt() []IIfStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfStmtContext); ok {
			len++
		}
	}

	tst := make([]IIfStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfStmtContext); ok {
			tst[i] = t.(IIfStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtsContext) IfStmt(i int) IIfStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
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

	return t.(IIfStmtContext)
}

func (s *StmtsContext) AllReturnStmt() []IReturnStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReturnStmtContext); ok {
			len++
		}
	}

	tst := make([]IReturnStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReturnStmtContext); ok {
			tst[i] = t.(IReturnStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtsContext) ReturnStmt(i int) IReturnStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
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

	return t.(IReturnStmtContext)
}

func (s *StmtsContext) AllFuncDeclStmt() []IFuncDeclStmtContext {
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

func (s *StmtsContext) FuncDeclStmt(i int) IFuncDeclStmtContext {
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

func (s *StmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterStmts(s)
	}
}

func (s *StmtsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitStmts(s)
	}
}

func (p *ezbpfParser) Stmts() (localctx IStmtsContext) {
	localctx = NewStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ezbpfParserRULE_stmts)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-50)) & ^0x3f) == 0 && ((int64(1)<<(_la-50))&20227) != 0 {
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(233)
				p.VarInitStmt()
			}

		case 2:
			{
				p.SetState(234)
				p.VarDeclStmt()
			}

		case 3:
			{
				p.SetState(235)
				p.ConstDeclStmt()
			}

		case 4:
			{
				p.SetState(236)
				p.MapDeclStmt()
			}

		case 5:
			{
				p.SetState(237)
				p.StructDeclStmt()
			}

		case 6:
			{
				p.SetState(238)
				p.IfStmt()
			}

		case 7:
			{
				p.SetState(239)
				p.ReturnStmt()
			}

		case 8:
			{
				p.SetState(240)
				p.FuncDeclStmt()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(245)
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

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *ezbpfParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ezbpfParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserSTRUCT {
		{
			p.SetState(246)
			p.StructDeclStmt()
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserMAP {
		{
			p.SetState(252)
			p.MapDeclStmt()
		}

		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ezbpfParserFN {
		{
			p.SetState(258)
			p.FuncDeclStmt()
		}

		p.SetState(261)
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
	case 3:
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
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
