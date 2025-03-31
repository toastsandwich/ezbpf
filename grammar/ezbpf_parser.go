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
		"", "", "", "", "", "", "", "'SEC'", "'('", "')'", "'{'", "'}'", "'['",
		"']'", "':'", "';'", "','", "'='", "'+'", "'-'", "'*'", "'/'", "'%'",
		"'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'", "'||'", "'if'",
		"'else'", "'while'", "'for'", "'var'", "'return'", "'struct'", "'bpf_map'",
		"'type'", "'key'", "'value'", "'max_entries'", "'delete'", "'fn'",
	}
	staticData.SymbolicNames = []string{
		"", "Identifier", "INTEGER_LITERAL", "CHAR_LITERAL", "STRING_LITERAL",
		"WS", "COMMENT", "SEC", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK",
		"RBRACK", "COLON", "SEMI", "COMMA", "EQUAL", "PLUS", "MINUS", "STAR",
		"SLASH", "PERCENT", "EQUAL_EQUAL", "BANG_EQUAL", "LT", "GT", "LTE",
		"GTE", "AND_AND", "OR_OR", "IF", "ELSE", "WHILE", "FOR", "VAR", "RETURN",
		"STRUCT", "BPF_MAP", "TYPE", "KEY", "VALUE", "MAX_ENTRIES", "DELETE",
		"FN",
	}
	staticData.RuleNames = []string{
		"compilationUnit", "structDefinition", "structMember", "bpfMapDefinition",
		"bpfMapOption", "globalVariableDeclaration", "functionDefinition", "parameterList",
		"parameter", "block", "statement", "variableDeclarationStatement", "assignmentStatement",
		"returnStatement", "ifStatement", "whileStatement", "forStatement",
		"simpleStatement", "functionCallStatement", "argumentList", "bpfMapOperationStatement",
		"type", "expression", "primaryExpression", "operator",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 282, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 0,
		1, 0, 5, 0, 55, 8, 0, 10, 0, 12, 0, 58, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 66, 8, 1, 10, 1, 12, 1, 69, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 82, 8, 3, 10, 3, 12, 3,
		85, 9, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 108,
		8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 116, 8, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 128, 8, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 133, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 140, 8,
		7, 10, 7, 12, 7, 143, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 151,
		8, 9, 10, 9, 12, 9, 154, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 166, 8, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 174, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 3, 14, 194, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 3, 17, 215, 8, 17, 1, 18, 1, 18, 1, 18, 3, 18,
		220, 8, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 228, 8, 19,
		10, 19, 12, 19, 231, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 3, 20, 253, 8, 20, 1, 21, 1, 21, 3, 21, 257, 8, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 263, 8, 22, 10, 22, 12, 22, 266, 9,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		3, 23, 278, 8, 23, 1, 24, 1, 24, 1, 24, 0, 0, 25, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		0, 1, 1, 0, 17, 30, 293, 0, 56, 1, 0, 0, 0, 2, 61, 1, 0, 0, 0, 4, 72, 1,
		0, 0, 0, 6, 77, 1, 0, 0, 0, 8, 107, 1, 0, 0, 0, 10, 109, 1, 0, 0, 0, 12,
		119, 1, 0, 0, 0, 14, 136, 1, 0, 0, 0, 16, 144, 1, 0, 0, 0, 18, 148, 1,
		0, 0, 0, 20, 165, 1, 0, 0, 0, 22, 167, 1, 0, 0, 0, 24, 177, 1, 0, 0, 0,
		26, 182, 1, 0, 0, 0, 28, 186, 1, 0, 0, 0, 30, 195, 1, 0, 0, 0, 32, 201,
		1, 0, 0, 0, 34, 214, 1, 0, 0, 0, 36, 216, 1, 0, 0, 0, 38, 224, 1, 0, 0,
		0, 40, 252, 1, 0, 0, 0, 42, 254, 1, 0, 0, 0, 44, 258, 1, 0, 0, 0, 46, 277,
		1, 0, 0, 0, 48, 279, 1, 0, 0, 0, 50, 55, 3, 2, 1, 0, 51, 55, 3, 6, 3, 0,
		52, 55, 3, 10, 5, 0, 53, 55, 3, 12, 6, 0, 54, 50, 1, 0, 0, 0, 54, 51, 1,
		0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56,
		54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0, 58, 56, 1, 0, 0,
		0, 59, 60, 5, 0, 0, 1, 60, 1, 1, 0, 0, 0, 61, 62, 5, 37, 0, 0, 62, 63,
		5, 1, 0, 0, 63, 67, 5, 10, 0, 0, 64, 66, 3, 4, 2, 0, 65, 64, 1, 0, 0, 0,
		66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 70, 1,
		0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 71, 5, 11, 0, 0, 71, 3, 1, 0, 0, 0, 72,
		73, 5, 1, 0, 0, 73, 74, 5, 14, 0, 0, 74, 75, 3, 42, 21, 0, 75, 76, 5, 15,
		0, 0, 76, 5, 1, 0, 0, 0, 77, 78, 5, 38, 0, 0, 78, 79, 5, 1, 0, 0, 79, 83,
		5, 10, 0, 0, 80, 82, 3, 8, 4, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0,
		83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 83, 1,
		0, 0, 0, 86, 87, 5, 11, 0, 0, 87, 88, 5, 15, 0, 0, 88, 7, 1, 0, 0, 0, 89,
		90, 5, 39, 0, 0, 90, 91, 5, 14, 0, 0, 91, 92, 5, 1, 0, 0, 92, 108, 5, 15,
		0, 0, 93, 94, 5, 40, 0, 0, 94, 95, 5, 14, 0, 0, 95, 96, 3, 42, 21, 0, 96,
		97, 5, 15, 0, 0, 97, 108, 1, 0, 0, 0, 98, 99, 5, 41, 0, 0, 99, 100, 5,
		14, 0, 0, 100, 101, 3, 42, 21, 0, 101, 102, 5, 15, 0, 0, 102, 108, 1, 0,
		0, 0, 103, 104, 5, 42, 0, 0, 104, 105, 5, 14, 0, 0, 105, 106, 5, 2, 0,
		0, 106, 108, 5, 15, 0, 0, 107, 89, 1, 0, 0, 0, 107, 93, 1, 0, 0, 0, 107,
		98, 1, 0, 0, 0, 107, 103, 1, 0, 0, 0, 108, 9, 1, 0, 0, 0, 109, 110, 5,
		35, 0, 0, 110, 111, 5, 1, 0, 0, 111, 112, 5, 14, 0, 0, 112, 115, 3, 42,
		21, 0, 113, 114, 5, 17, 0, 0, 114, 116, 3, 44, 22, 0, 115, 113, 1, 0, 0,
		0, 115, 116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 5, 15, 0, 0, 118,
		11, 1, 0, 0, 0, 119, 120, 5, 7, 0, 0, 120, 121, 5, 8, 0, 0, 121, 122, 5,
		4, 0, 0, 122, 123, 5, 9, 0, 0, 123, 124, 5, 44, 0, 0, 124, 125, 5, 1, 0,
		0, 125, 127, 5, 8, 0, 0, 126, 128, 3, 14, 7, 0, 127, 126, 1, 0, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 132, 5, 9, 0, 0, 130, 131,
		5, 14, 0, 0, 131, 133, 3, 42, 21, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1,
		0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 3, 18, 9, 0, 135, 13, 1, 0, 0,
		0, 136, 141, 3, 16, 8, 0, 137, 138, 5, 16, 0, 0, 138, 140, 3, 16, 8, 0,
		139, 137, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 15, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 145, 5,
		1, 0, 0, 145, 146, 5, 14, 0, 0, 146, 147, 3, 42, 21, 0, 147, 17, 1, 0,
		0, 0, 148, 152, 5, 10, 0, 0, 149, 151, 3, 20, 10, 0, 150, 149, 1, 0, 0,
		0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153,
		155, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 156, 5, 11, 0, 0, 156, 19,
		1, 0, 0, 0, 157, 166, 3, 22, 11, 0, 158, 166, 3, 24, 12, 0, 159, 166, 3,
		26, 13, 0, 160, 166, 3, 28, 14, 0, 161, 166, 3, 30, 15, 0, 162, 166, 3,
		32, 16, 0, 163, 166, 3, 36, 18, 0, 164, 166, 3, 40, 20, 0, 165, 157, 1,
		0, 0, 0, 165, 158, 1, 0, 0, 0, 165, 159, 1, 0, 0, 0, 165, 160, 1, 0, 0,
		0, 165, 161, 1, 0, 0, 0, 165, 162, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165,
		164, 1, 0, 0, 0, 166, 21, 1, 0, 0, 0, 167, 168, 5, 35, 0, 0, 168, 169,
		5, 1, 0, 0, 169, 170, 5, 14, 0, 0, 170, 173, 3, 42, 21, 0, 171, 172, 5,
		17, 0, 0, 172, 174, 3, 44, 22, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0,
		0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 5, 15, 0, 0, 176, 23, 1, 0, 0, 0,
		177, 178, 5, 1, 0, 0, 178, 179, 5, 17, 0, 0, 179, 180, 3, 44, 22, 0, 180,
		181, 5, 15, 0, 0, 181, 25, 1, 0, 0, 0, 182, 183, 5, 36, 0, 0, 183, 184,
		3, 44, 22, 0, 184, 185, 5, 15, 0, 0, 185, 27, 1, 0, 0, 0, 186, 187, 5,
		31, 0, 0, 187, 188, 5, 8, 0, 0, 188, 189, 3, 44, 22, 0, 189, 190, 5, 9,
		0, 0, 190, 193, 3, 18, 9, 0, 191, 192, 5, 32, 0, 0, 192, 194, 3, 18, 9,
		0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 29, 1, 0, 0, 0, 195,
		196, 5, 33, 0, 0, 196, 197, 5, 8, 0, 0, 197, 198, 3, 44, 22, 0, 198, 199,
		5, 9, 0, 0, 199, 200, 3, 18, 9, 0, 200, 31, 1, 0, 0, 0, 201, 202, 5, 34,
		0, 0, 202, 203, 5, 8, 0, 0, 203, 204, 3, 34, 17, 0, 204, 205, 5, 15, 0,
		0, 205, 206, 3, 44, 22, 0, 206, 207, 5, 15, 0, 0, 207, 208, 3, 34, 17,
		0, 208, 209, 5, 9, 0, 0, 209, 210, 3, 18, 9, 0, 210, 33, 1, 0, 0, 0, 211,
		215, 3, 24, 12, 0, 212, 215, 3, 22, 11, 0, 213, 215, 3, 36, 18, 0, 214,
		211, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 213, 1, 0, 0, 0, 215, 35, 1,
		0, 0, 0, 216, 217, 5, 1, 0, 0, 217, 219, 5, 8, 0, 0, 218, 220, 3, 38, 19,
		0, 219, 218, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221,
		222, 5, 9, 0, 0, 222, 223, 5, 15, 0, 0, 223, 37, 1, 0, 0, 0, 224, 229,
		3, 44, 22, 0, 225, 226, 5, 16, 0, 0, 226, 228, 3, 44, 22, 0, 227, 225,
		1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0,
		0, 0, 230, 39, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 233, 5, 1, 0, 0,
		233, 234, 5, 12, 0, 0, 234, 235, 3, 44, 22, 0, 235, 236, 5, 13, 0, 0, 236,
		237, 5, 17, 0, 0, 237, 238, 3, 44, 22, 0, 238, 239, 5, 15, 0, 0, 239, 253,
		1, 0, 0, 0, 240, 241, 5, 1, 0, 0, 241, 242, 5, 12, 0, 0, 242, 243, 3, 44,
		22, 0, 243, 244, 5, 13, 0, 0, 244, 253, 1, 0, 0, 0, 245, 246, 5, 43, 0,
		0, 246, 247, 5, 1, 0, 0, 247, 248, 5, 12, 0, 0, 248, 249, 3, 44, 22, 0,
		249, 250, 5, 13, 0, 0, 250, 251, 5, 15, 0, 0, 251, 253, 1, 0, 0, 0, 252,
		232, 1, 0, 0, 0, 252, 240, 1, 0, 0, 0, 252, 245, 1, 0, 0, 0, 253, 41, 1,
		0, 0, 0, 254, 256, 5, 1, 0, 0, 255, 257, 5, 20, 0, 0, 256, 255, 1, 0, 0,
		0, 256, 257, 1, 0, 0, 0, 257, 43, 1, 0, 0, 0, 258, 264, 3, 46, 23, 0, 259,
		260, 3, 48, 24, 0, 260, 261, 3, 46, 23, 0, 261, 263, 1, 0, 0, 0, 262, 259,
		1, 0, 0, 0, 263, 266, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264, 265, 1, 0,
		0, 0, 265, 45, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 267, 278, 5, 1, 0, 0,
		268, 278, 5, 2, 0, 0, 269, 278, 5, 3, 0, 0, 270, 278, 5, 4, 0, 0, 271,
		272, 5, 8, 0, 0, 272, 273, 3, 44, 22, 0, 273, 274, 5, 9, 0, 0, 274, 278,
		1, 0, 0, 0, 275, 278, 3, 36, 18, 0, 276, 278, 3, 40, 20, 0, 277, 267, 1,
		0, 0, 0, 277, 268, 1, 0, 0, 0, 277, 269, 1, 0, 0, 0, 277, 270, 1, 0, 0,
		0, 277, 271, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 276, 1, 0, 0, 0, 278,
		47, 1, 0, 0, 0, 279, 280, 7, 0, 0, 0, 280, 49, 1, 0, 0, 0, 20, 54, 56,
		67, 83, 107, 115, 127, 132, 141, 152, 165, 173, 193, 214, 219, 229, 252,
		256, 264, 277,
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
	ezbpfParserEOF             = antlr.TokenEOF
	ezbpfParserIdentifier      = 1
	ezbpfParserINTEGER_LITERAL = 2
	ezbpfParserCHAR_LITERAL    = 3
	ezbpfParserSTRING_LITERAL  = 4
	ezbpfParserWS              = 5
	ezbpfParserCOMMENT         = 6
	ezbpfParserSEC             = 7
	ezbpfParserLPAREN          = 8
	ezbpfParserRPAREN          = 9
	ezbpfParserLBRACE          = 10
	ezbpfParserRBRACE          = 11
	ezbpfParserLBRACK          = 12
	ezbpfParserRBRACK          = 13
	ezbpfParserCOLON           = 14
	ezbpfParserSEMI            = 15
	ezbpfParserCOMMA           = 16
	ezbpfParserEQUAL           = 17
	ezbpfParserPLUS            = 18
	ezbpfParserMINUS           = 19
	ezbpfParserSTAR            = 20
	ezbpfParserSLASH           = 21
	ezbpfParserPERCENT         = 22
	ezbpfParserEQUAL_EQUAL     = 23
	ezbpfParserBANG_EQUAL      = 24
	ezbpfParserLT              = 25
	ezbpfParserGT              = 26
	ezbpfParserLTE             = 27
	ezbpfParserGTE             = 28
	ezbpfParserAND_AND         = 29
	ezbpfParserOR_OR           = 30
	ezbpfParserIF              = 31
	ezbpfParserELSE            = 32
	ezbpfParserWHILE           = 33
	ezbpfParserFOR             = 34
	ezbpfParserVAR             = 35
	ezbpfParserRETURN          = 36
	ezbpfParserSTRUCT          = 37
	ezbpfParserBPF_MAP         = 38
	ezbpfParserTYPE            = 39
	ezbpfParserKEY             = 40
	ezbpfParserVALUE           = 41
	ezbpfParserMAX_ENTRIES     = 42
	ezbpfParserDELETE          = 43
	ezbpfParserFN              = 44
)

// ezbpfParser rules.
const (
	ezbpfParserRULE_compilationUnit              = 0
	ezbpfParserRULE_structDefinition             = 1
	ezbpfParserRULE_structMember                 = 2
	ezbpfParserRULE_bpfMapDefinition             = 3
	ezbpfParserRULE_bpfMapOption                 = 4
	ezbpfParserRULE_globalVariableDeclaration    = 5
	ezbpfParserRULE_functionDefinition           = 6
	ezbpfParserRULE_parameterList                = 7
	ezbpfParserRULE_parameter                    = 8
	ezbpfParserRULE_block                        = 9
	ezbpfParserRULE_statement                    = 10
	ezbpfParserRULE_variableDeclarationStatement = 11
	ezbpfParserRULE_assignmentStatement          = 12
	ezbpfParserRULE_returnStatement              = 13
	ezbpfParserRULE_ifStatement                  = 14
	ezbpfParserRULE_whileStatement               = 15
	ezbpfParserRULE_forStatement                 = 16
	ezbpfParserRULE_simpleStatement              = 17
	ezbpfParserRULE_functionCallStatement        = 18
	ezbpfParserRULE_argumentList                 = 19
	ezbpfParserRULE_bpfMapOperationStatement     = 20
	ezbpfParserRULE_type                         = 21
	ezbpfParserRULE_expression                   = 22
	ezbpfParserRULE_primaryExpression            = 23
	ezbpfParserRULE_operator                     = 24
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStructDefinition() []IStructDefinitionContext
	StructDefinition(i int) IStructDefinitionContext
	AllBpfMapDefinition() []IBpfMapDefinitionContext
	BpfMapDefinition(i int) IBpfMapDefinitionContext
	AllGlobalVariableDeclaration() []IGlobalVariableDeclarationContext
	GlobalVariableDeclaration(i int) IGlobalVariableDeclarationContext
	AllFunctionDefinition() []IFunctionDefinitionContext
	FunctionDefinition(i int) IFunctionDefinitionContext

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_compilationUnit
	return p
}

func InitEmptyCompilationUnitContext(p *CompilationUnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_compilationUnit
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(ezbpfParserEOF, 0)
}

func (s *CompilationUnitContext) AllStructDefinition() []IStructDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IStructDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructDefinitionContext); ok {
			tst[i] = t.(IStructDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *CompilationUnitContext) StructDefinition(i int) IStructDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDefinitionContext); ok {
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

	return t.(IStructDefinitionContext)
}

func (s *CompilationUnitContext) AllBpfMapDefinition() []IBpfMapDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBpfMapDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IBpfMapDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBpfMapDefinitionContext); ok {
			tst[i] = t.(IBpfMapDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *CompilationUnitContext) BpfMapDefinition(i int) IBpfMapDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBpfMapDefinitionContext); ok {
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

	return t.(IBpfMapDefinitionContext)
}

func (s *CompilationUnitContext) AllGlobalVariableDeclaration() []IGlobalVariableDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGlobalVariableDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IGlobalVariableDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGlobalVariableDeclarationContext); ok {
			tst[i] = t.(IGlobalVariableDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *CompilationUnitContext) GlobalVariableDeclaration(i int) IGlobalVariableDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalVariableDeclarationContext); ok {
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

	return t.(IGlobalVariableDeclarationContext)
}

func (s *CompilationUnitContext) AllFunctionDefinition() []IFunctionDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDefinitionContext); ok {
			tst[i] = t.(IFunctionDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *CompilationUnitContext) FunctionDefinition(i int) IFunctionDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
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

	return t.(IFunctionDefinitionContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (p *ezbpfParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ezbpfParserRULE_compilationUnit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&446676598912) != 0 {
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ezbpfParserSTRUCT:
			{
				p.SetState(50)
				p.StructDefinition()
			}

		case ezbpfParserBPF_MAP:
			{
				p.SetState(51)
				p.BpfMapDefinition()
			}

		case ezbpfParserVAR:
			{
				p.SetState(52)
				p.GlobalVariableDeclaration()
			}

		case ezbpfParserSEC:
			{
				p.SetState(53)
				p.FunctionDefinition()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(59)
		p.Match(ezbpfParserEOF)
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

// IStructDefinitionContext is an interface to support dynamic dispatch.
type IStructDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStructMember() []IStructMemberContext
	StructMember(i int) IStructMemberContext

	// IsStructDefinitionContext differentiates from other interfaces.
	IsStructDefinitionContext()
}

type StructDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDefinitionContext() *StructDefinitionContext {
	var p = new(StructDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structDefinition
	return p
}

func InitEmptyStructDefinitionContext(p *StructDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structDefinition
}

func (*StructDefinitionContext) IsStructDefinitionContext() {}

func NewStructDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDefinitionContext {
	var p = new(StructDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_structDefinition

	return p
}

func (s *StructDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDefinitionContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSTRUCT, 0)
}

func (s *StructDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *StructDefinitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLBRACE, 0)
}

func (s *StructDefinitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRBRACE, 0)
}

func (s *StructDefinitionContext) AllStructMember() []IStructMemberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructMemberContext); ok {
			len++
		}
	}

	tst := make([]IStructMemberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructMemberContext); ok {
			tst[i] = t.(IStructMemberContext)
			i++
		}
	}

	return tst
}

func (s *StructDefinitionContext) StructMember(i int) IStructMemberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructMemberContext); ok {
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

	return t.(IStructMemberContext)
}

func (s *StructDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterStructDefinition(s)
	}
}

func (s *StructDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitStructDefinition(s)
	}
}

func (p *ezbpfParser) StructDefinition() (localctx IStructDefinitionContext) {
	localctx = NewStructDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ezbpfParserRULE_structDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(ezbpfParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Match(ezbpfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(ezbpfParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserIdentifier {
		{
			p.SetState(64)
			p.StructMember()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(ezbpfParserRBRACE)
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

// IStructMemberContext is an interface to support dynamic dispatch.
type IStructMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	SEMI() antlr.TerminalNode

	// IsStructMemberContext differentiates from other interfaces.
	IsStructMemberContext()
}

type StructMemberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructMemberContext() *StructMemberContext {
	var p = new(StructMemberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structMember
	return p
}

func InitEmptyStructMemberContext(p *StructMemberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_structMember
}

func (*StructMemberContext) IsStructMemberContext() {}

func NewStructMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemberContext {
	var p = new(StructMemberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_structMember

	return p
}

func (s *StructMemberContext) GetParser() antlr.Parser { return s.parser }

func (s *StructMemberContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *StructMemberContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *StructMemberContext) Type_() ITypeContext {
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

func (s *StructMemberContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *StructMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructMemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterStructMember(s)
	}
}

func (s *StructMemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitStructMember(s)
	}
}

func (p *ezbpfParser) StructMember() (localctx IStructMemberContext) {
	localctx = NewStructMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ezbpfParserRULE_structMember)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(ezbpfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Type_()
	}
	{
		p.SetState(75)
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

// IBpfMapDefinitionContext is an interface to support dynamic dispatch.
type IBpfMapDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BPF_MAP() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	AllBpfMapOption() []IBpfMapOptionContext
	BpfMapOption(i int) IBpfMapOptionContext

	// IsBpfMapDefinitionContext differentiates from other interfaces.
	IsBpfMapDefinitionContext()
}

type BpfMapDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBpfMapDefinitionContext() *BpfMapDefinitionContext {
	var p = new(BpfMapDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_bpfMapDefinition
	return p
}

func InitEmptyBpfMapDefinitionContext(p *BpfMapDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_bpfMapDefinition
}

func (*BpfMapDefinitionContext) IsBpfMapDefinitionContext() {}

func NewBpfMapDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BpfMapDefinitionContext {
	var p = new(BpfMapDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_bpfMapDefinition

	return p
}

func (s *BpfMapDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *BpfMapDefinitionContext) BPF_MAP() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBPF_MAP, 0)
}

func (s *BpfMapDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *BpfMapDefinitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLBRACE, 0)
}

func (s *BpfMapDefinitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRBRACE, 0)
}

func (s *BpfMapDefinitionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *BpfMapDefinitionContext) AllBpfMapOption() []IBpfMapOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBpfMapOptionContext); ok {
			len++
		}
	}

	tst := make([]IBpfMapOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBpfMapOptionContext); ok {
			tst[i] = t.(IBpfMapOptionContext)
			i++
		}
	}

	return tst
}

func (s *BpfMapDefinitionContext) BpfMapOption(i int) IBpfMapOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBpfMapOptionContext); ok {
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

	return t.(IBpfMapOptionContext)
}

func (s *BpfMapDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BpfMapDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BpfMapDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterBpfMapDefinition(s)
	}
}

func (s *BpfMapDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitBpfMapDefinition(s)
	}
}

func (p *ezbpfParser) BpfMapDefinition() (localctx IBpfMapDefinitionContext) {
	localctx = NewBpfMapDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ezbpfParserRULE_bpfMapDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(ezbpfParserBPF_MAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Match(ezbpfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(ezbpfParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8246337208320) != 0 {
		{
			p.SetState(80)
			p.BpfMapOption()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
		p.Match(ezbpfParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
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

// IBpfMapOptionContext is an interface to support dynamic dispatch.
type IBpfMapOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	KEY() antlr.TerminalNode
	Type_() ITypeContext
	VALUE() antlr.TerminalNode
	MAX_ENTRIES() antlr.TerminalNode
	INTEGER_LITERAL() antlr.TerminalNode

	// IsBpfMapOptionContext differentiates from other interfaces.
	IsBpfMapOptionContext()
}

type BpfMapOptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBpfMapOptionContext() *BpfMapOptionContext {
	var p = new(BpfMapOptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_bpfMapOption
	return p
}

func InitEmptyBpfMapOptionContext(p *BpfMapOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_bpfMapOption
}

func (*BpfMapOptionContext) IsBpfMapOptionContext() {}

func NewBpfMapOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BpfMapOptionContext {
	var p = new(BpfMapOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_bpfMapOption

	return p
}

func (s *BpfMapOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *BpfMapOptionContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserTYPE, 0)
}

func (s *BpfMapOptionContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *BpfMapOptionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *BpfMapOptionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *BpfMapOptionContext) KEY() antlr.TerminalNode {
	return s.GetToken(ezbpfParserKEY, 0)
}

func (s *BpfMapOptionContext) Type_() ITypeContext {
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

func (s *BpfMapOptionContext) VALUE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserVALUE, 0)
}

func (s *BpfMapOptionContext) MAX_ENTRIES() antlr.TerminalNode {
	return s.GetToken(ezbpfParserMAX_ENTRIES, 0)
}

func (s *BpfMapOptionContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserINTEGER_LITERAL, 0)
}

func (s *BpfMapOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BpfMapOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BpfMapOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterBpfMapOption(s)
	}
}

func (s *BpfMapOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitBpfMapOption(s)
	}
}

func (p *ezbpfParser) BpfMapOption() (localctx IBpfMapOptionContext) {
	localctx = NewBpfMapOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ezbpfParserRULE_bpfMapOption)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ezbpfParserTYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Match(ezbpfParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(ezbpfParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(ezbpfParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Match(ezbpfParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ezbpfParserKEY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(ezbpfParserKEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(ezbpfParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Type_()
		}
		{
			p.SetState(96)
			p.Match(ezbpfParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ezbpfParserVALUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Match(ezbpfParserVALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(ezbpfParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Type_()
		}
		{
			p.SetState(101)
			p.Match(ezbpfParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ezbpfParserMAX_ENTRIES:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.Match(ezbpfParserMAX_ENTRIES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Match(ezbpfParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(ezbpfParserINTEGER_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(ezbpfParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IGlobalVariableDeclarationContext is an interface to support dynamic dispatch.
type IGlobalVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	SEMI() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	Expression() IExpressionContext

	// IsGlobalVariableDeclarationContext differentiates from other interfaces.
	IsGlobalVariableDeclarationContext()
}

type GlobalVariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalVariableDeclarationContext() *GlobalVariableDeclarationContext {
	var p = new(GlobalVariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_globalVariableDeclaration
	return p
}

func InitEmptyGlobalVariableDeclarationContext(p *GlobalVariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_globalVariableDeclaration
}

func (*GlobalVariableDeclarationContext) IsGlobalVariableDeclarationContext() {}

func NewGlobalVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalVariableDeclarationContext {
	var p = new(GlobalVariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_globalVariableDeclaration

	return p
}

func (s *GlobalVariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalVariableDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserVAR, 0)
}

func (s *GlobalVariableDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *GlobalVariableDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *GlobalVariableDeclarationContext) Type_() ITypeContext {
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

func (s *GlobalVariableDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *GlobalVariableDeclarationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserEQUAL, 0)
}

func (s *GlobalVariableDeclarationContext) Expression() IExpressionContext {
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

func (s *GlobalVariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalVariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalVariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterGlobalVariableDeclaration(s)
	}
}

func (s *GlobalVariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitGlobalVariableDeclaration(s)
	}
}

func (p *ezbpfParser) GlobalVariableDeclaration() (localctx IGlobalVariableDeclarationContext) {
	localctx = NewGlobalVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ezbpfParserRULE_globalVariableDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(ezbpfParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(ezbpfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Type_()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ezbpfParserEQUAL {
		{
			p.SetState(113)
			p.Match(ezbpfParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Expression()
		}

	}
	{
		p.SetState(117)
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

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEC() antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	FN() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Block() IBlockContext
	ParameterList() IParameterListContext
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_functionDefinition
	return p
}

func InitEmptyFunctionDefinitionContext(p *FunctionDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_functionDefinition
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) SEC() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEC, 0)
}

func (s *FunctionDefinitionContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserLPAREN)
}

func (s *FunctionDefinitionContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserLPAREN, i)
}

func (s *FunctionDefinitionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSTRING_LITERAL, 0)
}

func (s *FunctionDefinitionContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserRPAREN)
}

func (s *FunctionDefinitionContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserRPAREN, i)
}

func (s *FunctionDefinitionContext) FN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserFN, 0)
}

func (s *FunctionDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *FunctionDefinitionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDefinitionContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDefinitionContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *FunctionDefinitionContext) Type_() ITypeContext {
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

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (p *ezbpfParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ezbpfParserRULE_functionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(ezbpfParserSEC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(ezbpfParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(ezbpfParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(ezbpfParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(ezbpfParserFN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(ezbpfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(ezbpfParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ezbpfParserIdentifier {
		{
			p.SetState(126)
			p.ParameterList()
		}

	}
	{
		p.SetState(129)
		p.Match(ezbpfParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ezbpfParserCOLON {
		{
			p.SetState(130)
			p.Match(ezbpfParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Type_()
		}

	}
	{
		p.SetState(134)
		p.Block()
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

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
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

	return t.(IParameterContext)
}

func (s *ParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserCOMMA)
}

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOMMA, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (p *ezbpfParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ezbpfParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Parameter()
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserCOMMA {
		{
			p.SetState(137)
			p.Match(ezbpfParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Parameter()
		}

		p.SetState(143)
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

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *ParameterContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *ParameterContext) Type_() ITypeContext {
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

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *ezbpfParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ezbpfParserRULE_parameter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(ezbpfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(ezbpfParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRBRACE, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *ezbpfParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ezbpfParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(ezbpfParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8927089524738) != 0 {
		{
			p.SetState(149)
			p.Statement()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(155)
		p.Match(ezbpfParserRBRACE)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDeclarationStatement() IVariableDeclarationStatementContext
	AssignmentStatement() IAssignmentStatementContext
	ReturnStatement() IReturnStatementContext
	IfStatement() IIfStatementContext
	WhileStatement() IWhileStatementContext
	ForStatement() IForStatementContext
	FunctionCallStatement() IFunctionCallStatementContext
	BpfMapOperationStatement() IBpfMapOperationStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VariableDeclarationStatement() IVariableDeclarationStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationStatementContext)
}

func (s *StatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) FunctionCallStatement() IFunctionCallStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallStatementContext)
}

func (s *StatementContext) BpfMapOperationStatement() IBpfMapOperationStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBpfMapOperationStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBpfMapOperationStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *ezbpfParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ezbpfParserRULE_statement)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.VariableDeclarationStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.AssignmentStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(159)
			p.ReturnStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(160)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(161)
			p.WhileStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(162)
			p.ForStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(163)
			p.FunctionCallStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(164)
			p.BpfMapOperationStatement()
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

// IVariableDeclarationStatementContext is an interface to support dynamic dispatch.
type IVariableDeclarationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	SEMI() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVariableDeclarationStatementContext differentiates from other interfaces.
	IsVariableDeclarationStatementContext()
}

type VariableDeclarationStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationStatementContext() *VariableDeclarationStatementContext {
	var p = new(VariableDeclarationStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_variableDeclarationStatement
	return p
}

func InitEmptyVariableDeclarationStatementContext(p *VariableDeclarationStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_variableDeclarationStatement
}

func (*VariableDeclarationStatementContext) IsVariableDeclarationStatementContext() {}

func NewVariableDeclarationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationStatementContext {
	var p = new(VariableDeclarationStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_variableDeclarationStatement

	return p
}

func (s *VariableDeclarationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationStatementContext) VAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserVAR, 0)
}

func (s *VariableDeclarationStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *VariableDeclarationStatementContext) COLON() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOLON, 0)
}

func (s *VariableDeclarationStatementContext) Type_() ITypeContext {
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

func (s *VariableDeclarationStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *VariableDeclarationStatementContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserEQUAL, 0)
}

func (s *VariableDeclarationStatementContext) Expression() IExpressionContext {
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

func (s *VariableDeclarationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterVariableDeclarationStatement(s)
	}
}

func (s *VariableDeclarationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitVariableDeclarationStatement(s)
	}
}

func (p *ezbpfParser) VariableDeclarationStatement() (localctx IVariableDeclarationStatementContext) {
	localctx = NewVariableDeclarationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ezbpfParserRULE_variableDeclarationStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(ezbpfParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(ezbpfParserIdentifier)
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
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ezbpfParserEQUAL {
		{
			p.SetState(171)
			p.Match(ezbpfParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Expression()
		}

	}
	{
		p.SetState(175)
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

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	Expression() IExpressionContext
	SEMI() antlr.TerminalNode

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *AssignmentStatementContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserEQUAL, 0)
}

func (s *AssignmentStatementContext) Expression() IExpressionContext {
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

func (s *AssignmentStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (p *ezbpfParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ezbpfParserRULE_assignmentStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(ezbpfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Match(ezbpfParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Expression()
	}
	{
		p.SetState(180)
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

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext
	SEMI() antlr.TerminalNode

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRETURN, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
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

func (s *ReturnStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *ezbpfParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ezbpfParserRULE_returnStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(ezbpfParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Expression()
	}
	{
		p.SetState(184)
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIF, 0)
}

func (s *IfStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLPAREN, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
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

func (s *IfStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRPAREN, 0)
}

func (s *IfStatementContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserELSE, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *ezbpfParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ezbpfParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(ezbpfParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(ezbpfParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.Expression()
	}
	{
		p.SetState(189)
		p.Match(ezbpfParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.Block()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ezbpfParserELSE {
		{
			p.SetState(191)
			p.Match(ezbpfParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Block()
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

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserWHILE, 0)
}

func (s *WhileStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLPAREN, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
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

func (s *WhileStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRPAREN, 0)
}

func (s *WhileStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (p *ezbpfParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ezbpfParserRULE_whileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(ezbpfParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Match(ezbpfParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Expression()
	}
	{
		p.SetState(198)
		p.Match(ezbpfParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Block()
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

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllSimpleStatement() []ISimpleStatementContext
	SimpleStatement(i int) ISimpleStatementContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserFOR, 0)
}

func (s *ForStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLPAREN, 0)
}

func (s *ForStatementContext) AllSimpleStatement() []ISimpleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStatementContext); ok {
			tst[i] = t.(ISimpleStatementContext)
			i++
		}
	}

	return tst
}

func (s *ForStatementContext) SimpleStatement(i int) ISimpleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
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

	return t.(ISimpleStatementContext)
}

func (s *ForStatementContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserSEMI)
}

func (s *ForStatementContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, i)
}

func (s *ForStatementContext) Expression() IExpressionContext {
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

func (s *ForStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRPAREN, 0)
}

func (s *ForStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (p *ezbpfParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ezbpfParserRULE_forStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(ezbpfParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(ezbpfParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.SimpleStatement()
	}
	{
		p.SetState(204)
		p.Match(ezbpfParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Expression()
	}
	{
		p.SetState(206)
		p.Match(ezbpfParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.SimpleStatement()
	}
	{
		p.SetState(208)
		p.Match(ezbpfParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Block()
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

// ISimpleStatementContext is an interface to support dynamic dispatch.
type ISimpleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignmentStatement() IAssignmentStatementContext
	VariableDeclarationStatement() IVariableDeclarationStatementContext
	FunctionCallStatement() IFunctionCallStatementContext

	// IsSimpleStatementContext differentiates from other interfaces.
	IsSimpleStatementContext()
}

type SimpleStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStatementContext() *SimpleStatementContext {
	var p = new(SimpleStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_simpleStatement
	return p
}

func InitEmptySimpleStatementContext(p *SimpleStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_simpleStatement
}

func (*SimpleStatementContext) IsSimpleStatementContext() {}

func NewSimpleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStatementContext {
	var p = new(SimpleStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_simpleStatement

	return p
}

func (s *SimpleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *SimpleStatementContext) VariableDeclarationStatement() IVariableDeclarationStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationStatementContext)
}

func (s *SimpleStatementContext) FunctionCallStatement() IFunctionCallStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallStatementContext)
}

func (s *SimpleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterSimpleStatement(s)
	}
}

func (s *SimpleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitSimpleStatement(s)
	}
}

func (p *ezbpfParser) SimpleStatement() (localctx ISimpleStatementContext) {
	localctx = NewSimpleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ezbpfParserRULE_simpleStatement)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.AssignmentStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.VariableDeclarationStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(213)
			p.FunctionCallStatement()
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

// IFunctionCallStatementContext is an interface to support dynamic dispatch.
type IFunctionCallStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsFunctionCallStatementContext differentiates from other interfaces.
	IsFunctionCallStatementContext()
}

type FunctionCallStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallStatementContext() *FunctionCallStatementContext {
	var p = new(FunctionCallStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_functionCallStatement
	return p
}

func InitEmptyFunctionCallStatementContext(p *FunctionCallStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_functionCallStatement
}

func (*FunctionCallStatementContext) IsFunctionCallStatementContext() {}

func NewFunctionCallStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallStatementContext {
	var p = new(FunctionCallStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_functionCallStatement

	return p
}

func (s *FunctionCallStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *FunctionCallStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLPAREN, 0)
}

func (s *FunctionCallStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRPAREN, 0)
}

func (s *FunctionCallStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *FunctionCallStatementContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterFunctionCallStatement(s)
	}
}

func (s *FunctionCallStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitFunctionCallStatement(s)
	}
}

func (p *ezbpfParser) FunctionCallStatement() (localctx IFunctionCallStatementContext) {
	localctx = NewFunctionCallStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ezbpfParserRULE_functionCallStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(ezbpfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Match(ezbpfParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8796093022494) != 0 {
		{
			p.SetState(218)
			p.ArgumentList()
		}

	}
	{
		p.SetState(221)
		p.Match(ezbpfParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
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

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
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

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
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

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ezbpfParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ezbpfParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (p *ezbpfParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ezbpfParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Expression()
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ezbpfParserCOMMA {
		{
			p.SetState(225)
			p.Match(ezbpfParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Expression()
		}

		p.SetState(231)
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

// IBpfMapOperationStatementContext is an interface to support dynamic dispatch.
type IBpfMapOperationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	RBRACK() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	DELETE() antlr.TerminalNode

	// IsBpfMapOperationStatementContext differentiates from other interfaces.
	IsBpfMapOperationStatementContext()
}

type BpfMapOperationStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBpfMapOperationStatementContext() *BpfMapOperationStatementContext {
	var p = new(BpfMapOperationStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_bpfMapOperationStatement
	return p
}

func InitEmptyBpfMapOperationStatementContext(p *BpfMapOperationStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_bpfMapOperationStatement
}

func (*BpfMapOperationStatementContext) IsBpfMapOperationStatementContext() {}

func NewBpfMapOperationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BpfMapOperationStatementContext {
	var p = new(BpfMapOperationStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_bpfMapOperationStatement

	return p
}

func (s *BpfMapOperationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BpfMapOperationStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *BpfMapOperationStatementContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLBRACK, 0)
}

func (s *BpfMapOperationStatementContext) AllExpression() []IExpressionContext {
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

func (s *BpfMapOperationStatementContext) Expression(i int) IExpressionContext {
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

func (s *BpfMapOperationStatementContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRBRACK, 0)
}

func (s *BpfMapOperationStatementContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserEQUAL, 0)
}

func (s *BpfMapOperationStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSEMI, 0)
}

func (s *BpfMapOperationStatementContext) DELETE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserDELETE, 0)
}

func (s *BpfMapOperationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BpfMapOperationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BpfMapOperationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterBpfMapOperationStatement(s)
	}
}

func (s *BpfMapOperationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitBpfMapOperationStatement(s)
	}
}

func (p *ezbpfParser) BpfMapOperationStatement() (localctx IBpfMapOperationStatementContext) {
	localctx = NewBpfMapOperationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ezbpfParserRULE_bpfMapOperationStatement)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Match(ezbpfParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(ezbpfParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Expression()
		}
		{
			p.SetState(235)
			p.Match(ezbpfParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Match(ezbpfParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Expression()
		}
		{
			p.SetState(238)
			p.Match(ezbpfParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(240)
			p.Match(ezbpfParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(ezbpfParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Expression()
		}
		{
			p.SetState(243)
			p.Match(ezbpfParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(245)
			p.Match(ezbpfParserDELETE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Match(ezbpfParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Match(ezbpfParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.Expression()
		}
		{
			p.SetState(249)
			p.Match(ezbpfParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Match(ezbpfParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	STAR() antlr.TerminalNode

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

func (s *TypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *TypeContext) STAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSTAR, 0)
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
	p.EnterRule(localctx, 42, ezbpfParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(ezbpfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ezbpfParserSTAR {
		{
			p.SetState(255)
			p.Match(ezbpfParserSTAR)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPrimaryExpression() []IPrimaryExpressionContext
	PrimaryExpression(i int) IPrimaryExpressionContext
	AllOperator() []IOperatorContext
	Operator(i int) IOperatorContext

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

func (s *ExpressionContext) AllPrimaryExpression() []IPrimaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryExpressionContext); ok {
			tst[i] = t.(IPrimaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) PrimaryExpression(i int) IPrimaryExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
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

	return t.(IPrimaryExpressionContext)
}

func (s *ExpressionContext) AllOperator() []IOperatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperatorContext); ok {
			len++
		}
	}

	tst := make([]IOperatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperatorContext); ok {
			tst[i] = t.(IOperatorContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Operator(i int) IOperatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
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

	return t.(IOperatorContext)
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
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ezbpfParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.PrimaryExpression()
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2147352576) != 0 {
		{
			p.SetState(259)
			p.Operator()
		}
		{
			p.SetState(260)
			p.PrimaryExpression()
		}

		p.SetState(266)
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

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	INTEGER_LITERAL() antlr.TerminalNode
	CHAR_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	FunctionCallStatement() IFunctionCallStatementContext
	BpfMapOperationStatement() IBpfMapOperationStatementContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ezbpfParserIdentifier, 0)
}

func (s *PrimaryExpressionContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserINTEGER_LITERAL, 0)
}

func (s *PrimaryExpressionContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserCHAR_LITERAL, 0)
}

func (s *PrimaryExpressionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSTRING_LITERAL, 0)
}

func (s *PrimaryExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLPAREN, 0)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
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

func (s *PrimaryExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ezbpfParserRPAREN, 0)
}

func (s *PrimaryExpressionContext) FunctionCallStatement() IFunctionCallStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallStatementContext)
}

func (s *PrimaryExpressionContext) BpfMapOperationStatement() IBpfMapOperationStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBpfMapOperationStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBpfMapOperationStatementContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (p *ezbpfParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ezbpfParserRULE_primaryExpression)
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(267)
			p.Match(ezbpfParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(268)
			p.Match(ezbpfParserINTEGER_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(269)
			p.Match(ezbpfParserCHAR_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(270)
			p.Match(ezbpfParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(271)
			p.Match(ezbpfParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.Expression()
		}
		{
			p.SetState(273)
			p.Match(ezbpfParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(275)
			p.FunctionCallStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(276)
			p.BpfMapOperationStatement()
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

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL_EQUAL() antlr.TerminalNode
	BANG_EQUAL() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GTE() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	STAR() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	PERCENT() antlr.TerminalNode
	AND_AND() antlr.TerminalNode
	OR_OR() antlr.TerminalNode
	EQUAL() antlr.TerminalNode

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_operator
	return p
}

func InitEmptyOperatorContext(p *OperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ezbpfParserRULE_operator
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ezbpfParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) EQUAL_EQUAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserEQUAL_EQUAL, 0)
}

func (s *OperatorContext) BANG_EQUAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserBANG_EQUAL, 0)
}

func (s *OperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLT, 0)
}

func (s *OperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserGT, 0)
}

func (s *OperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserLTE, 0)
}

func (s *OperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(ezbpfParserGTE, 0)
}

func (s *OperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ezbpfParserPLUS, 0)
}

func (s *OperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ezbpfParserMINUS, 0)
}

func (s *OperatorContext) STAR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSTAR, 0)
}

func (s *OperatorContext) SLASH() antlr.TerminalNode {
	return s.GetToken(ezbpfParserSLASH, 0)
}

func (s *OperatorContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(ezbpfParserPERCENT, 0)
}

func (s *OperatorContext) AND_AND() antlr.TerminalNode {
	return s.GetToken(ezbpfParserAND_AND, 0)
}

func (s *OperatorContext) OR_OR() antlr.TerminalNode {
	return s.GetToken(ezbpfParserOR_OR, 0)
}

func (s *OperatorContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ezbpfParserEQUAL, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ezbpfListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *ezbpfParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ezbpfParserRULE_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2147352576) != 0) {
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
