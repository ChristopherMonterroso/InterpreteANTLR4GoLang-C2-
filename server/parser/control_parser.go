// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Control
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

type ControlParser struct {
	*antlr.BaseParser
}

var ControlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func controlParserInit() {
	staticData := &ControlParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "':'", "'?'", "'+='", "'-='", "'['", "']'", "'.'", "'append'",
		"'('", "')'", "'removeLast'", "'remove'", "'at'", "'func'", "'{'", "'return'",
		"'}'", "'->'", "'&'", "','", "'_'", "'inout'", "'println'", "'print'",
		"'if'", "'else'", "'switch'", "'case'", "'break'", "'default'", "'while'",
		"'for'", "'in'", "'...'", "'guard'", "'!'", "'%'", "'*'", "'/'", "'+'",
		"'-'", "'>='", "'>'", "'<='", "'<'", "'=='", "'!='", "'&&'", "'||'",
		"'true'", "'false'", "'isEmpty'", "'count'", "'Int'", "'String'", "'Float'",
		"'Bool'", "'Character'", "'continue'", "'let'", "'var'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "INT", "ID", "FLOAT",
		"STRING", "CHAR", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"prog", "block", "stmt", "assignstmt", "vectorPpts", "funcstmt", "listaParamsCall",
		"listaParams", "listaExp", "printlnstmt", "printstmt", "ifstmt", "switchstmt",
		"cases", "caseblock", "whilestmt", "forstmt", "guardstmt", "expr", "primitivo",
		"transfer_sentence", "var_type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 70, 409, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 49, 8, 1, 10, 1, 12, 1, 52, 9,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 64,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 121, 8, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 144, 8, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 153, 8, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 199, 8, 5, 1, 6, 1, 6, 3, 6, 203, 8,
		6, 1, 6, 3, 6, 206, 8, 6, 1, 6, 1, 6, 1, 6, 3, 6, 211, 8, 6, 1, 7, 3, 7,
		214, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 219, 8, 7, 1, 7, 1, 7, 3, 7, 223, 8,
		7, 1, 7, 1, 7, 1, 7, 3, 7, 228, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 234,
		8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 241, 8, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 3, 11, 258, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 3, 11, 267, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 274,
		8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 283, 8,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 289, 8, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 5, 13, 298, 8, 13, 10, 13, 12, 13, 301, 9,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 308, 8, 14, 1, 14, 1, 14,
		1, 14, 3, 14, 313, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 339, 8, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 347, 8, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 3, 18, 375, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 398, 8, 18, 10, 18, 12,
		18, 401, 9, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 0, 1,
		36, 22, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 0, 11, 2, 0, 22, 22, 64, 64, 1, 0, 51, 52, 1, 0, 39,
		40, 1, 0, 41, 42, 1, 0, 43, 44, 1, 0, 45, 46, 1, 0, 47, 48, 1, 0, 49, 50,
		1, 0, 55, 59, 3, 0, 17, 17, 30, 30, 60, 60, 1, 0, 61, 62, 449, 0, 44, 1,
		0, 0, 0, 2, 50, 1, 0, 0, 0, 4, 63, 1, 0, 0, 0, 6, 120, 1, 0, 0, 0, 8, 143,
		1, 0, 0, 0, 10, 198, 1, 0, 0, 0, 12, 202, 1, 0, 0, 0, 14, 233, 1, 0, 0,
		0, 16, 240, 1, 0, 0, 0, 18, 242, 1, 0, 0, 0, 20, 247, 1, 0, 0, 0, 22, 288,
		1, 0, 0, 0, 24, 290, 1, 0, 0, 0, 26, 299, 1, 0, 0, 0, 28, 312, 1, 0, 0,
		0, 30, 314, 1, 0, 0, 0, 32, 338, 1, 0, 0, 0, 34, 340, 1, 0, 0, 0, 36, 374,
		1, 0, 0, 0, 38, 402, 1, 0, 0, 0, 40, 404, 1, 0, 0, 0, 42, 406, 1, 0, 0,
		0, 44, 45, 3, 2, 1, 0, 45, 46, 5, 0, 0, 1, 46, 1, 1, 0, 0, 0, 47, 49, 3,
		4, 2, 0, 48, 47, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50,
		51, 1, 0, 0, 0, 51, 3, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 64, 3, 6, 3,
		0, 54, 64, 3, 18, 9, 0, 55, 64, 3, 20, 10, 0, 56, 64, 3, 22, 11, 0, 57,
		64, 3, 30, 15, 0, 58, 64, 3, 24, 12, 0, 59, 64, 3, 32, 16, 0, 60, 64, 3,
		34, 17, 0, 61, 64, 3, 8, 4, 0, 62, 64, 3, 10, 5, 0, 63, 53, 1, 0, 0, 0,
		63, 54, 1, 0, 0, 0, 63, 55, 1, 0, 0, 0, 63, 56, 1, 0, 0, 0, 63, 57, 1,
		0, 0, 0, 63, 58, 1, 0, 0, 0, 63, 59, 1, 0, 0, 0, 63, 60, 1, 0, 0, 0, 63,
		61, 1, 0, 0, 0, 63, 62, 1, 0, 0, 0, 64, 5, 1, 0, 0, 0, 65, 66, 5, 64, 0,
		0, 66, 67, 5, 1, 0, 0, 67, 121, 3, 36, 18, 0, 68, 69, 3, 42, 21, 0, 69,
		70, 5, 64, 0, 0, 70, 71, 5, 2, 0, 0, 71, 72, 3, 38, 19, 0, 72, 73, 5, 1,
		0, 0, 73, 74, 3, 36, 18, 0, 74, 121, 1, 0, 0, 0, 75, 76, 3, 42, 21, 0,
		76, 77, 5, 64, 0, 0, 77, 78, 5, 2, 0, 0, 78, 79, 3, 38, 19, 0, 79, 80,
		5, 3, 0, 0, 80, 121, 1, 0, 0, 0, 81, 82, 3, 42, 21, 0, 82, 83, 5, 64, 0,
		0, 83, 84, 5, 1, 0, 0, 84, 85, 3, 36, 18, 0, 85, 121, 1, 0, 0, 0, 86, 87,
		5, 64, 0, 0, 87, 88, 5, 4, 0, 0, 88, 121, 3, 36, 18, 0, 89, 90, 5, 64,
		0, 0, 90, 91, 5, 5, 0, 0, 91, 121, 3, 36, 18, 0, 92, 93, 3, 42, 21, 0,
		93, 94, 5, 64, 0, 0, 94, 95, 5, 2, 0, 0, 95, 96, 5, 6, 0, 0, 96, 97, 3,
		38, 19, 0, 97, 98, 5, 7, 0, 0, 98, 99, 5, 1, 0, 0, 99, 100, 5, 6, 0, 0,
		100, 101, 5, 7, 0, 0, 101, 121, 1, 0, 0, 0, 102, 103, 3, 42, 21, 0, 103,
		104, 5, 64, 0, 0, 104, 105, 5, 2, 0, 0, 105, 106, 5, 6, 0, 0, 106, 107,
		3, 38, 19, 0, 107, 108, 5, 7, 0, 0, 108, 109, 5, 1, 0, 0, 109, 110, 5,
		6, 0, 0, 110, 111, 3, 16, 8, 0, 111, 112, 5, 7, 0, 0, 112, 121, 1, 0, 0,
		0, 113, 114, 5, 64, 0, 0, 114, 115, 5, 6, 0, 0, 115, 116, 3, 36, 18, 0,
		116, 117, 5, 7, 0, 0, 117, 118, 5, 1, 0, 0, 118, 119, 3, 36, 18, 0, 119,
		121, 1, 0, 0, 0, 120, 65, 1, 0, 0, 0, 120, 68, 1, 0, 0, 0, 120, 75, 1,
		0, 0, 0, 120, 81, 1, 0, 0, 0, 120, 86, 1, 0, 0, 0, 120, 89, 1, 0, 0, 0,
		120, 92, 1, 0, 0, 0, 120, 102, 1, 0, 0, 0, 120, 113, 1, 0, 0, 0, 121, 7,
		1, 0, 0, 0, 122, 123, 5, 64, 0, 0, 123, 124, 5, 8, 0, 0, 124, 125, 5, 9,
		0, 0, 125, 126, 5, 10, 0, 0, 126, 127, 3, 36, 18, 0, 127, 128, 5, 11, 0,
		0, 128, 144, 1, 0, 0, 0, 129, 130, 5, 64, 0, 0, 130, 131, 5, 8, 0, 0, 131,
		132, 5, 12, 0, 0, 132, 133, 5, 10, 0, 0, 133, 144, 5, 11, 0, 0, 134, 135,
		5, 64, 0, 0, 135, 136, 5, 8, 0, 0, 136, 137, 5, 13, 0, 0, 137, 138, 5,
		10, 0, 0, 138, 139, 5, 14, 0, 0, 139, 140, 5, 2, 0, 0, 140, 141, 3, 36,
		18, 0, 141, 142, 5, 11, 0, 0, 142, 144, 1, 0, 0, 0, 143, 122, 1, 0, 0,
		0, 143, 129, 1, 0, 0, 0, 143, 134, 1, 0, 0, 0, 144, 9, 1, 0, 0, 0, 145,
		146, 5, 15, 0, 0, 146, 147, 5, 64, 0, 0, 147, 148, 5, 10, 0, 0, 148, 149,
		5, 11, 0, 0, 149, 150, 5, 16, 0, 0, 150, 152, 3, 2, 1, 0, 151, 153, 5,
		17, 0, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 1, 0, 0,
		0, 154, 155, 5, 18, 0, 0, 155, 199, 1, 0, 0, 0, 156, 157, 5, 15, 0, 0,
		157, 158, 5, 64, 0, 0, 158, 159, 5, 10, 0, 0, 159, 160, 3, 14, 7, 0, 160,
		161, 5, 11, 0, 0, 161, 162, 5, 16, 0, 0, 162, 163, 3, 2, 1, 0, 163, 164,
		5, 18, 0, 0, 164, 199, 1, 0, 0, 0, 165, 166, 5, 15, 0, 0, 166, 167, 5,
		64, 0, 0, 167, 168, 5, 10, 0, 0, 168, 169, 5, 11, 0, 0, 169, 170, 5, 19,
		0, 0, 170, 171, 3, 38, 19, 0, 171, 172, 5, 16, 0, 0, 172, 173, 3, 2, 1,
		0, 173, 174, 5, 17, 0, 0, 174, 175, 3, 36, 18, 0, 175, 176, 5, 18, 0, 0,
		176, 199, 1, 0, 0, 0, 177, 178, 5, 15, 0, 0, 178, 179, 5, 64, 0, 0, 179,
		180, 5, 10, 0, 0, 180, 181, 3, 14, 7, 0, 181, 182, 5, 11, 0, 0, 182, 183,
		5, 19, 0, 0, 183, 184, 3, 38, 19, 0, 184, 185, 5, 16, 0, 0, 185, 186, 3,
		2, 1, 0, 186, 187, 5, 17, 0, 0, 187, 188, 3, 36, 18, 0, 188, 189, 5, 18,
		0, 0, 189, 199, 1, 0, 0, 0, 190, 191, 5, 64, 0, 0, 191, 192, 5, 10, 0,
		0, 192, 199, 5, 11, 0, 0, 193, 194, 5, 64, 0, 0, 194, 195, 5, 10, 0, 0,
		195, 196, 3, 12, 6, 0, 196, 197, 5, 11, 0, 0, 197, 199, 1, 0, 0, 0, 198,
		145, 1, 0, 0, 0, 198, 156, 1, 0, 0, 0, 198, 165, 1, 0, 0, 0, 198, 177,
		1, 0, 0, 0, 198, 190, 1, 0, 0, 0, 198, 193, 1, 0, 0, 0, 199, 11, 1, 0,
		0, 0, 200, 201, 5, 64, 0, 0, 201, 203, 5, 2, 0, 0, 202, 200, 1, 0, 0, 0,
		202, 203, 1, 0, 0, 0, 203, 205, 1, 0, 0, 0, 204, 206, 5, 20, 0, 0, 205,
		204, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 210,
		3, 36, 18, 0, 208, 209, 5, 21, 0, 0, 209, 211, 3, 12, 6, 0, 210, 208, 1,
		0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 13, 1, 0, 0, 0, 212, 214, 7, 0, 0,
		0, 213, 212, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215,
		216, 5, 64, 0, 0, 216, 218, 5, 2, 0, 0, 217, 219, 5, 23, 0, 0, 218, 217,
		1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 234, 3, 38,
		19, 0, 221, 223, 7, 0, 0, 0, 222, 221, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0,
		223, 224, 1, 0, 0, 0, 224, 225, 5, 64, 0, 0, 225, 227, 5, 2, 0, 0, 226,
		228, 5, 23, 0, 0, 227, 226, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229,
		1, 0, 0, 0, 229, 230, 3, 38, 19, 0, 230, 231, 5, 21, 0, 0, 231, 232, 3,
		14, 7, 0, 232, 234, 1, 0, 0, 0, 233, 213, 1, 0, 0, 0, 233, 222, 1, 0, 0,
		0, 234, 15, 1, 0, 0, 0, 235, 241, 3, 36, 18, 0, 236, 237, 3, 36, 18, 0,
		237, 238, 5, 21, 0, 0, 238, 239, 3, 16, 8, 0, 239, 241, 1, 0, 0, 0, 240,
		235, 1, 0, 0, 0, 240, 236, 1, 0, 0, 0, 241, 17, 1, 0, 0, 0, 242, 243, 5,
		24, 0, 0, 243, 244, 5, 10, 0, 0, 244, 245, 3, 36, 18, 0, 245, 246, 5, 11,
		0, 0, 246, 19, 1, 0, 0, 0, 247, 248, 5, 25, 0, 0, 248, 249, 5, 10, 0, 0,
		249, 250, 3, 36, 18, 0, 250, 251, 5, 11, 0, 0, 251, 21, 1, 0, 0, 0, 252,
		253, 5, 26, 0, 0, 253, 254, 3, 36, 18, 0, 254, 255, 5, 16, 0, 0, 255, 257,
		3, 2, 1, 0, 256, 258, 3, 40, 20, 0, 257, 256, 1, 0, 0, 0, 257, 258, 1,
		0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 260, 5, 18, 0, 0, 260, 289, 1, 0, 0,
		0, 261, 262, 5, 26, 0, 0, 262, 263, 3, 36, 18, 0, 263, 264, 5, 16, 0, 0,
		264, 266, 3, 2, 1, 0, 265, 267, 3, 40, 20, 0, 266, 265, 1, 0, 0, 0, 266,
		267, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 269, 5, 18, 0, 0, 269, 270,
		5, 27, 0, 0, 270, 271, 5, 16, 0, 0, 271, 273, 3, 2, 1, 0, 272, 274, 3,
		40, 20, 0, 273, 272, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 275, 1, 0,
		0, 0, 275, 276, 5, 18, 0, 0, 276, 289, 1, 0, 0, 0, 277, 278, 5, 26, 0,
		0, 278, 279, 3, 36, 18, 0, 279, 280, 5, 16, 0, 0, 280, 282, 3, 2, 1, 0,
		281, 283, 3, 40, 20, 0, 282, 281, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283,
		284, 1, 0, 0, 0, 284, 285, 5, 18, 0, 0, 285, 286, 5, 27, 0, 0, 286, 287,
		3, 22, 11, 0, 287, 289, 1, 0, 0, 0, 288, 252, 1, 0, 0, 0, 288, 261, 1,
		0, 0, 0, 288, 277, 1, 0, 0, 0, 289, 23, 1, 0, 0, 0, 290, 291, 5, 28, 0,
		0, 291, 292, 3, 36, 18, 0, 292, 293, 5, 16, 0, 0, 293, 294, 3, 26, 13,
		0, 294, 295, 5, 18, 0, 0, 295, 25, 1, 0, 0, 0, 296, 298, 3, 28, 14, 0,
		297, 296, 1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 299,
		300, 1, 0, 0, 0, 300, 27, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 302, 303, 5,
		29, 0, 0, 303, 304, 3, 36, 18, 0, 304, 305, 5, 2, 0, 0, 305, 307, 3, 2,
		1, 0, 306, 308, 5, 30, 0, 0, 307, 306, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0,
		308, 313, 1, 0, 0, 0, 309, 310, 5, 31, 0, 0, 310, 311, 5, 2, 0, 0, 311,
		313, 3, 2, 1, 0, 312, 302, 1, 0, 0, 0, 312, 309, 1, 0, 0, 0, 313, 29, 1,
		0, 0, 0, 314, 315, 5, 32, 0, 0, 315, 316, 3, 36, 18, 0, 316, 317, 5, 16,
		0, 0, 317, 318, 3, 2, 1, 0, 318, 319, 5, 18, 0, 0, 319, 31, 1, 0, 0, 0,
		320, 321, 5, 33, 0, 0, 321, 322, 5, 64, 0, 0, 322, 323, 5, 34, 0, 0, 323,
		324, 3, 36, 18, 0, 324, 325, 5, 35, 0, 0, 325, 326, 3, 36, 18, 0, 326,
		327, 5, 16, 0, 0, 327, 328, 3, 2, 1, 0, 328, 329, 5, 18, 0, 0, 329, 339,
		1, 0, 0, 0, 330, 331, 5, 33, 0, 0, 331, 332, 5, 64, 0, 0, 332, 333, 5,
		34, 0, 0, 333, 334, 3, 36, 18, 0, 334, 335, 5, 16, 0, 0, 335, 336, 3, 2,
		1, 0, 336, 337, 5, 18, 0, 0, 337, 339, 1, 0, 0, 0, 338, 320, 1, 0, 0, 0,
		338, 330, 1, 0, 0, 0, 339, 33, 1, 0, 0, 0, 340, 341, 5, 36, 0, 0, 341,
		342, 3, 36, 18, 0, 342, 343, 5, 27, 0, 0, 343, 344, 5, 16, 0, 0, 344, 346,
		3, 2, 1, 0, 345, 347, 3, 40, 20, 0, 346, 345, 1, 0, 0, 0, 346, 347, 1,
		0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 349, 5, 18, 0, 0, 349, 35, 1, 0, 0,
		0, 350, 351, 6, 18, -1, 0, 351, 352, 5, 37, 0, 0, 352, 375, 3, 36, 18,
		18, 353, 354, 5, 10, 0, 0, 354, 355, 3, 36, 18, 0, 355, 356, 5, 11, 0,
		0, 356, 375, 1, 0, 0, 0, 357, 375, 5, 63, 0, 0, 358, 375, 5, 66, 0, 0,
		359, 375, 7, 1, 0, 0, 360, 375, 5, 65, 0, 0, 361, 375, 5, 67, 0, 0, 362,
		375, 5, 64, 0, 0, 363, 364, 5, 64, 0, 0, 364, 365, 5, 8, 0, 0, 365, 375,
		5, 53, 0, 0, 366, 367, 5, 64, 0, 0, 367, 368, 5, 8, 0, 0, 368, 375, 5,
		54, 0, 0, 369, 370, 5, 64, 0, 0, 370, 371, 5, 6, 0, 0, 371, 372, 3, 36,
		18, 0, 372, 373, 5, 7, 0, 0, 373, 375, 1, 0, 0, 0, 374, 350, 1, 0, 0, 0,
		374, 353, 1, 0, 0, 0, 374, 357, 1, 0, 0, 0, 374, 358, 1, 0, 0, 0, 374,
		359, 1, 0, 0, 0, 374, 360, 1, 0, 0, 0, 374, 361, 1, 0, 0, 0, 374, 362,
		1, 0, 0, 0, 374, 363, 1, 0, 0, 0, 374, 366, 1, 0, 0, 0, 374, 369, 1, 0,
		0, 0, 375, 399, 1, 0, 0, 0, 376, 377, 10, 17, 0, 0, 377, 378, 5, 38, 0,
		0, 378, 398, 3, 36, 18, 18, 379, 380, 10, 16, 0, 0, 380, 381, 7, 2, 0,
		0, 381, 398, 3, 36, 18, 17, 382, 383, 10, 15, 0, 0, 383, 384, 7, 3, 0,
		0, 384, 398, 3, 36, 18, 16, 385, 386, 10, 14, 0, 0, 386, 387, 7, 4, 0,
		0, 387, 398, 3, 36, 18, 15, 388, 389, 10, 13, 0, 0, 389, 390, 7, 5, 0,
		0, 390, 398, 3, 36, 18, 14, 391, 392, 10, 12, 0, 0, 392, 393, 7, 6, 0,
		0, 393, 398, 3, 36, 18, 13, 394, 395, 10, 11, 0, 0, 395, 396, 7, 7, 0,
		0, 396, 398, 3, 36, 18, 12, 397, 376, 1, 0, 0, 0, 397, 379, 1, 0, 0, 0,
		397, 382, 1, 0, 0, 0, 397, 385, 1, 0, 0, 0, 397, 388, 1, 0, 0, 0, 397,
		391, 1, 0, 0, 0, 397, 394, 1, 0, 0, 0, 398, 401, 1, 0, 0, 0, 399, 397,
		1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 37, 1, 0, 0, 0, 401, 399, 1, 0,
		0, 0, 402, 403, 7, 8, 0, 0, 403, 39, 1, 0, 0, 0, 404, 405, 7, 9, 0, 0,
		405, 41, 1, 0, 0, 0, 406, 407, 7, 10, 0, 0, 407, 43, 1, 0, 0, 0, 28, 50,
		63, 120, 143, 152, 198, 202, 205, 210, 213, 218, 222, 227, 233, 240, 257,
		266, 273, 282, 288, 299, 307, 312, 338, 346, 374, 397, 399,
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

// ControlParserInit initializes any static state used to implement ControlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewControlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ControlParserInit() {
	staticData := &ControlParserStaticData
	staticData.once.Do(controlParserInit)
}

// NewControlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewControlParser(input antlr.TokenStream) *ControlParser {
	ControlParserInit()
	this := new(ControlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ControlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Control.g4"

	return this
}

// ControlParser tokens.
const (
	ControlParserEOF          = antlr.TokenEOF
	ControlParserT__0         = 1
	ControlParserT__1         = 2
	ControlParserT__2         = 3
	ControlParserT__3         = 4
	ControlParserT__4         = 5
	ControlParserT__5         = 6
	ControlParserT__6         = 7
	ControlParserT__7         = 8
	ControlParserT__8         = 9
	ControlParserT__9         = 10
	ControlParserT__10        = 11
	ControlParserT__11        = 12
	ControlParserT__12        = 13
	ControlParserT__13        = 14
	ControlParserT__14        = 15
	ControlParserT__15        = 16
	ControlParserT__16        = 17
	ControlParserT__17        = 18
	ControlParserT__18        = 19
	ControlParserT__19        = 20
	ControlParserT__20        = 21
	ControlParserT__21        = 22
	ControlParserT__22        = 23
	ControlParserT__23        = 24
	ControlParserT__24        = 25
	ControlParserT__25        = 26
	ControlParserT__26        = 27
	ControlParserT__27        = 28
	ControlParserT__28        = 29
	ControlParserT__29        = 30
	ControlParserT__30        = 31
	ControlParserT__31        = 32
	ControlParserT__32        = 33
	ControlParserT__33        = 34
	ControlParserT__34        = 35
	ControlParserT__35        = 36
	ControlParserT__36        = 37
	ControlParserT__37        = 38
	ControlParserT__38        = 39
	ControlParserT__39        = 40
	ControlParserT__40        = 41
	ControlParserT__41        = 42
	ControlParserT__42        = 43
	ControlParserT__43        = 44
	ControlParserT__44        = 45
	ControlParserT__45        = 46
	ControlParserT__46        = 47
	ControlParserT__47        = 48
	ControlParserT__48        = 49
	ControlParserT__49        = 50
	ControlParserT__50        = 51
	ControlParserT__51        = 52
	ControlParserT__52        = 53
	ControlParserT__53        = 54
	ControlParserT__54        = 55
	ControlParserT__55        = 56
	ControlParserT__56        = 57
	ControlParserT__57        = 58
	ControlParserT__58        = 59
	ControlParserT__59        = 60
	ControlParserT__60        = 61
	ControlParserT__61        = 62
	ControlParserINT          = 63
	ControlParserID           = 64
	ControlParserFLOAT        = 65
	ControlParserSTRING       = 66
	ControlParserCHAR         = 67
	ControlParserWS           = 68
	ControlParserCOMMENT      = 69
	ControlParserLINE_COMMENT = 70
)

// ControlParser rules.
const (
	ControlParserRULE_prog              = 0
	ControlParserRULE_block             = 1
	ControlParserRULE_stmt              = 2
	ControlParserRULE_assignstmt        = 3
	ControlParserRULE_vectorPpts        = 4
	ControlParserRULE_funcstmt          = 5
	ControlParserRULE_listaParamsCall   = 6
	ControlParserRULE_listaParams       = 7
	ControlParserRULE_listaExp          = 8
	ControlParserRULE_printlnstmt       = 9
	ControlParserRULE_printstmt         = 10
	ControlParserRULE_ifstmt            = 11
	ControlParserRULE_switchstmt        = 12
	ControlParserRULE_cases             = 13
	ControlParserRULE_caseblock         = 14
	ControlParserRULE_whilestmt         = 15
	ControlParserRULE_forstmt           = 16
	ControlParserRULE_guardstmt         = 17
	ControlParserRULE_expr              = 18
	ControlParserRULE_primitivo         = 19
	ControlParserRULE_transfer_sentence = 20
	ControlParserRULE_var_type          = 21
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

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
	p.RuleIndex = ControlParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Block() IBlockContext {
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

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(ControlParserEOF, 0)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ControlParserRULE_prog)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Block()
	}
	{
		p.SetState(45)
		p.Match(ControlParserEOF)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

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
	p.RuleIndex = ControlParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
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

func (s *BlockContext) Stmt(i int) IStmtContext {
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

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ControlParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&774056188456449) != 0 {
		{
			p.SetState(47)
			p.Stmt()
		}

		p.SetState(52)
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

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignstmt() IAssignstmtContext
	Printlnstmt() IPrintlnstmtContext
	Printstmt() IPrintstmtContext
	Ifstmt() IIfstmtContext
	Whilestmt() IWhilestmtContext
	Switchstmt() ISwitchstmtContext
	Forstmt() IForstmtContext
	Guardstmt() IGuardstmtContext
	VectorPpts() IVectorPptsContext
	Funcstmt() IFuncstmtContext

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
	p.RuleIndex = ControlParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Assignstmt() IAssignstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignstmtContext)
}

func (s *StmtContext) Printlnstmt() IPrintlnstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintlnstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintlnstmtContext)
}

func (s *StmtContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *StmtContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *StmtContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *StmtContext) Switchstmt() ISwitchstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchstmtContext)
}

func (s *StmtContext) Forstmt() IForstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForstmtContext)
}

func (s *StmtContext) Guardstmt() IGuardstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardstmtContext)
}

func (s *StmtContext) VectorPpts() IVectorPptsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorPptsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorPptsContext)
}

func (s *StmtContext) Funcstmt() IFuncstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncstmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ControlParserRULE_stmt)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Assignstmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Printlnstmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Printstmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Ifstmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)
			p.Whilestmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(58)
			p.Switchstmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(59)
			p.Forstmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(60)
			p.Guardstmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(61)
			p.VectorPpts()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(62)
			p.Funcstmt()
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

// IAssignstmtContext is an interface to support dynamic dispatch.
type IAssignstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignstmtContext differentiates from other interfaces.
	IsAssignstmtContext()
}

type AssignstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignstmtContext() *AssignstmtContext {
	var p = new(AssignstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_assignstmt
	return p
}

func InitEmptyAssignstmtContext(p *AssignstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_assignstmt
}

func (*AssignstmtContext) IsAssignstmtContext() {}

func NewAssignstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignstmtContext {
	var p = new(AssignstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_assignstmt

	return p
}

func (s *AssignstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignstmtContext) CopyAll(ctx *AssignstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AsignacionContext struct {
	AssignstmtContext
}

func NewAsignacionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionContext {
	var p = new(AsignacionContext)

	InitEmptyAssignstmtContext(&p.AssignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignstmtContext))

	return p
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *AsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *AsignacionContext) Primitivo() IPrimitivoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *AsignacionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (s *AsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitAsignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionVectorContext struct {
	AssignstmtContext
}

func NewAsignacionVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionVectorContext {
	var p = new(AsignacionVectorContext)

	InitEmptyAssignstmtContext(&p.AssignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignstmtContext))

	return p
}

func (s *AsignacionVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionVectorContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *AsignacionVectorContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *AsignacionVectorContext) Primitivo() IPrimitivoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *AsignacionVectorContext) ListaExp() IListaExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExpContext)
}

func (s *AsignacionVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterAsignacionVector(s)
	}
}

func (s *AsignacionVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitAsignacionVector(s)
	}
}

func (s *AsignacionVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitAsignacionVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReasignacionContext struct {
	AssignstmtContext
}

func NewReasignacionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReasignacionContext {
	var p = new(ReasignacionContext)

	InitEmptyAssignstmtContext(&p.AssignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignstmtContext))

	return p
}

func (s *ReasignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReasignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *ReasignacionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReasignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterReasignacion(s)
	}
}

func (s *ReasignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitReasignacion(s)
	}
}

func (s *ReasignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitReasignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncrementoContext struct {
	AssignstmtContext
}

func NewIncrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrementoContext {
	var p = new(IncrementoContext)

	InitEmptyAssignstmtContext(&p.AssignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignstmtContext))

	return p
}

func (s *IncrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *IncrementoContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IncrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterIncremento(s)
	}
}

func (s *IncrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitIncremento(s)
	}
}

func (s *IncrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitIncremento(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReasignacionVectorContext struct {
	AssignstmtContext
}

func NewReasignacionVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReasignacionVectorContext {
	var p = new(ReasignacionVectorContext)

	InitEmptyAssignstmtContext(&p.AssignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignstmtContext))

	return p
}

func (s *ReasignacionVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReasignacionVectorContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *ReasignacionVectorContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ReasignacionVectorContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ReasignacionVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterReasignacionVector(s)
	}
}

func (s *ReasignacionVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitReasignacionVector(s)
	}
}

func (s *ReasignacionVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitReasignacionVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionVectorVacioContext struct {
	AssignstmtContext
}

func NewAsignacionVectorVacioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionVectorVacioContext {
	var p = new(AsignacionVectorVacioContext)

	InitEmptyAssignstmtContext(&p.AssignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignstmtContext))

	return p
}

func (s *AsignacionVectorVacioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionVectorVacioContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *AsignacionVectorVacioContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *AsignacionVectorVacioContext) Primitivo() IPrimitivoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *AsignacionVectorVacioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterAsignacionVectorVacio(s)
	}
}

func (s *AsignacionVectorVacioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitAsignacionVectorVacio(s)
	}
}

func (s *AsignacionVectorVacioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitAsignacionVectorVacio(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionNoExprContext struct {
	AssignstmtContext
}

func NewAsignacionNoExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionNoExprContext {
	var p = new(AsignacionNoExprContext)

	InitEmptyAssignstmtContext(&p.AssignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignstmtContext))

	return p
}

func (s *AsignacionNoExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionNoExprContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *AsignacionNoExprContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *AsignacionNoExprContext) Primitivo() IPrimitivoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *AsignacionNoExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterAsignacionNoExpr(s)
	}
}

func (s *AsignacionNoExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitAsignacionNoExpr(s)
	}
}

func (s *AsignacionNoExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitAsignacionNoExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionNoPrimitiveContext struct {
	AssignstmtContext
}

func NewAsignacionNoPrimitiveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionNoPrimitiveContext {
	var p = new(AsignacionNoPrimitiveContext)

	InitEmptyAssignstmtContext(&p.AssignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignstmtContext))

	return p
}

func (s *AsignacionNoPrimitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionNoPrimitiveContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *AsignacionNoPrimitiveContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *AsignacionNoPrimitiveContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionNoPrimitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterAsignacionNoPrimitive(s)
	}
}

func (s *AsignacionNoPrimitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitAsignacionNoPrimitive(s)
	}
}

func (s *AsignacionNoPrimitiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitAsignacionNoPrimitive(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecrementoContext struct {
	AssignstmtContext
}

func NewDecrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecrementoContext {
	var p = new(DecrementoContext)

	InitEmptyAssignstmtContext(&p.AssignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignstmtContext))

	return p
}

func (s *DecrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *DecrementoContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DecrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterDecremento(s)
	}
}

func (s *DecrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitDecremento(s)
	}
}

func (s *DecrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitDecremento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Assignstmt() (localctx IAssignstmtContext) {
	localctx = NewAssignstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ControlParserRULE_assignstmt)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewReasignacionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.expr(0)
		}

	case 2:
		localctx = NewAsignacionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.Var_type()
		}
		{
			p.SetState(69)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Primitivo()
		}
		{
			p.SetState(72)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.expr(0)
		}

	case 3:
		localctx = NewAsignacionNoExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.Var_type()
		}
		{
			p.SetState(76)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Primitivo()
		}
		{
			p.SetState(79)
			p.Match(ControlParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewAsignacionNoPrimitiveContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(81)
			p.Var_type()
		}
		{
			p.SetState(82)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.expr(0)
		}

	case 5:
		localctx = NewIncrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(86)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(ControlParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.expr(0)
		}

	case 6:
		localctx = NewDecrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(89)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(ControlParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.expr(0)
		}

	case 7:
		localctx = NewAsignacionVectorVacioContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(92)
			p.Var_type()
		}
		{
			p.SetState(93)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Primitivo()
		}
		{
			p.SetState(97)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewAsignacionVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(102)
			p.Var_type()
		}
		{
			p.SetState(103)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Primitivo()
		}
		{
			p.SetState(107)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.ListaExp()
		}
		{
			p.SetState(111)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewReasignacionVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(113)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.expr(0)
		}
		{
			p.SetState(116)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.expr(0)
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

// IVectorPptsContext is an interface to support dynamic dispatch.
type IVectorPptsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVectorPptsContext differentiates from other interfaces.
	IsVectorPptsContext()
}

type VectorPptsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorPptsContext() *VectorPptsContext {
	var p = new(VectorPptsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_vectorPpts
	return p
}

func InitEmptyVectorPptsContext(p *VectorPptsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_vectorPpts
}

func (*VectorPptsContext) IsVectorPptsContext() {}

func NewVectorPptsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorPptsContext {
	var p = new(VectorPptsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_vectorPpts

	return p
}

func (s *VectorPptsContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorPptsContext) CopyAll(ctx *VectorPptsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VectorPptsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorPptsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorRemoveAtContext struct {
	VectorPptsContext
}

func NewVectorRemoveAtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorRemoveAtContext {
	var p = new(VectorRemoveAtContext)

	InitEmptyVectorPptsContext(&p.VectorPptsContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorPptsContext))

	return p
}

func (s *VectorRemoveAtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorRemoveAtContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *VectorRemoveAtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VectorRemoveAtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterVectorRemoveAt(s)
	}
}

func (s *VectorRemoveAtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitVectorRemoveAt(s)
	}
}

func (s *VectorRemoveAtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitVectorRemoveAt(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorAppendContext struct {
	VectorPptsContext
}

func NewVectorAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorAppendContext {
	var p = new(VectorAppendContext)

	InitEmptyVectorPptsContext(&p.VectorPptsContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorPptsContext))

	return p
}

func (s *VectorAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorAppendContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *VectorAppendContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VectorAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterVectorAppend(s)
	}
}

func (s *VectorAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitVectorAppend(s)
	}
}

func (s *VectorAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitVectorAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorRemoveLastContext struct {
	VectorPptsContext
}

func NewVectorRemoveLastContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorRemoveLastContext {
	var p = new(VectorRemoveLastContext)

	InitEmptyVectorPptsContext(&p.VectorPptsContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorPptsContext))

	return p
}

func (s *VectorRemoveLastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorRemoveLastContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *VectorRemoveLastContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterVectorRemoveLast(s)
	}
}

func (s *VectorRemoveLastContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitVectorRemoveLast(s)
	}
}

func (s *VectorRemoveLastContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitVectorRemoveLast(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) VectorPpts() (localctx IVectorPptsContext) {
	localctx = NewVectorPptsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ControlParserRULE_vectorPpts)
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVectorAppendContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(ControlParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(ControlParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.expr(0)
		}
		{
			p.SetState(127)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewVectorRemoveLastContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(ControlParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Match(ControlParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVectorRemoveAtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(ControlParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(ControlParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(ControlParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.expr(0)
		}
		{
			p.SetState(141)
			p.Match(ControlParserT__10)
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

// IFuncstmtContext is an interface to support dynamic dispatch.
type IFuncstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncstmtContext differentiates from other interfaces.
	IsFuncstmtContext()
}

type FuncstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncstmtContext() *FuncstmtContext {
	var p = new(FuncstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_funcstmt
	return p
}

func InitEmptyFuncstmtContext(p *FuncstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_funcstmt
}

func (*FuncstmtContext) IsFuncstmtContext() {}

func NewFuncstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncstmtContext {
	var p = new(FuncstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_funcstmt

	return p
}

func (s *FuncstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncstmtContext) CopyAll(ctx *FuncstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncParams_sinRetornoContext struct {
	FuncstmtContext
}

func NewFuncParams_sinRetornoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncParams_sinRetornoContext {
	var p = new(FuncParams_sinRetornoContext)

	InitEmptyFuncstmtContext(&p.FuncstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncstmtContext))

	return p
}

func (s *FuncParams_sinRetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParams_sinRetornoContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *FuncParams_sinRetornoContext) ListaParams() IListaParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaParamsContext)
}

func (s *FuncParams_sinRetornoContext) Block() IBlockContext {
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

func (s *FuncParams_sinRetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterFuncParams_sinRetorno(s)
	}
}

func (s *FuncParams_sinRetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitFuncParams_sinRetorno(s)
	}
}

func (s *FuncParams_sinRetornoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitFuncParams_sinRetorno(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncParams_ConRetornoContext struct {
	FuncstmtContext
}

func NewFuncParams_ConRetornoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncParams_ConRetornoContext {
	var p = new(FuncParams_ConRetornoContext)

	InitEmptyFuncstmtContext(&p.FuncstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncstmtContext))

	return p
}

func (s *FuncParams_ConRetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParams_ConRetornoContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *FuncParams_ConRetornoContext) ListaParams() IListaParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaParamsContext)
}

func (s *FuncParams_ConRetornoContext) Primitivo() IPrimitivoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *FuncParams_ConRetornoContext) Block() IBlockContext {
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

func (s *FuncParams_ConRetornoContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncParams_ConRetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterFuncParams_ConRetorno(s)
	}
}

func (s *FuncParams_ConRetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitFuncParams_ConRetorno(s)
	}
}

func (s *FuncParams_ConRetornoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitFuncParams_ConRetorno(s)

	default:
		return t.VisitChildren(s)
	}
}

type Func_conRetorno_conTipoContext struct {
	FuncstmtContext
}

func NewFunc_conRetorno_conTipoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_conRetorno_conTipoContext {
	var p = new(Func_conRetorno_conTipoContext)

	InitEmptyFuncstmtContext(&p.FuncstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncstmtContext))

	return p
}

func (s *Func_conRetorno_conTipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_conRetorno_conTipoContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *Func_conRetorno_conTipoContext) Primitivo() IPrimitivoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Func_conRetorno_conTipoContext) Block() IBlockContext {
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

func (s *Func_conRetorno_conTipoContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Func_conRetorno_conTipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterFunc_conRetorno_conTipo(s)
	}
}

func (s *Func_conRetorno_conTipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitFunc_conRetorno_conTipo(s)
	}
}

func (s *Func_conRetorno_conTipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitFunc_conRetorno_conTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallFunctionContext struct {
	FuncstmtContext
}

func NewCallFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFunctionContext {
	var p = new(CallFunctionContext)

	InitEmptyFuncstmtContext(&p.FuncstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncstmtContext))

	return p
}

func (s *CallFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *CallFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterCallFunction(s)
	}
}

func (s *CallFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitCallFunction(s)
	}
}

func (s *CallFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitCallFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallFunctionParamsContext struct {
	FuncstmtContext
}

func NewCallFunctionParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFunctionParamsContext {
	var p = new(CallFunctionParamsContext)

	InitEmptyFuncstmtContext(&p.FuncstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncstmtContext))

	return p
}

func (s *CallFunctionParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionParamsContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *CallFunctionParamsContext) ListaParamsCall() IListaParamsCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaParamsCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaParamsCallContext)
}

func (s *CallFunctionParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterCallFunctionParams(s)
	}
}

func (s *CallFunctionParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitCallFunctionParams(s)
	}
}

func (s *CallFunctionParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitCallFunctionParams(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncSinTipoRetornoContext struct {
	FuncstmtContext
}

func NewFuncSinTipoRetornoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncSinTipoRetornoContext {
	var p = new(FuncSinTipoRetornoContext)

	InitEmptyFuncstmtContext(&p.FuncstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncstmtContext))

	return p
}

func (s *FuncSinTipoRetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSinTipoRetornoContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *FuncSinTipoRetornoContext) Block() IBlockContext {
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

func (s *FuncSinTipoRetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterFuncSinTipoRetorno(s)
	}
}

func (s *FuncSinTipoRetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitFuncSinTipoRetorno(s)
	}
}

func (s *FuncSinTipoRetornoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitFuncSinTipoRetorno(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Funcstmt() (localctx IFuncstmtContext) {
	localctx = NewFuncstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ControlParserRULE_funcstmt)
	var _la int

	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFuncSinTipoRetornoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Match(ControlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Block()
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__16 {
			{
				p.SetState(151)
				p.Match(ControlParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(154)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFuncParams_sinRetornoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Match(ControlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.ListaParams()
		}
		{
			p.SetState(160)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Block()
		}
		{
			p.SetState(163)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFunc_conRetorno_conTipoContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(165)
			p.Match(ControlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(ControlParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Primitivo()
		}
		{
			p.SetState(171)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Block()
		}
		{
			p.SetState(173)
			p.Match(ControlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.expr(0)
		}
		{
			p.SetState(175)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewFuncParams_ConRetornoContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(177)
			p.Match(ControlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.ListaParams()
		}
		{
			p.SetState(181)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(ControlParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Primitivo()
		}
		{
			p.SetState(184)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Block()
		}
		{
			p.SetState(186)
			p.Match(ControlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.expr(0)
		}
		{
			p.SetState(188)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewCallFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(190)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewCallFunctionParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(193)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.ListaParamsCall()
		}
		{
			p.SetState(196)
			p.Match(ControlParserT__10)
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

// IListaParamsCallContext is an interface to support dynamic dispatch.
type IListaParamsCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRef returns the ref token.
	GetRef() antlr.Token

	// SetRef sets the ref token.
	SetRef(antlr.Token)

	// Getter signatures
	Expr() IExprContext
	ID() antlr.TerminalNode
	ListaParamsCall() IListaParamsCallContext

	// IsListaParamsCallContext differentiates from other interfaces.
	IsListaParamsCallContext()
}

type ListaParamsCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ref    antlr.Token
}

func NewEmptyListaParamsCallContext() *ListaParamsCallContext {
	var p = new(ListaParamsCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_listaParamsCall
	return p
}

func InitEmptyListaParamsCallContext(p *ListaParamsCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_listaParamsCall
}

func (*ListaParamsCallContext) IsListaParamsCallContext() {}

func NewListaParamsCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaParamsCallContext {
	var p = new(ListaParamsCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_listaParamsCall

	return p
}

func (s *ListaParamsCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaParamsCallContext) GetRef() antlr.Token { return s.ref }

func (s *ListaParamsCallContext) SetRef(v antlr.Token) { s.ref = v }

func (s *ListaParamsCallContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListaParamsCallContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *ListaParamsCallContext) ListaParamsCall() IListaParamsCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaParamsCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaParamsCallContext)
}

func (s *ListaParamsCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaParamsCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaParamsCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterListaParamsCall(s)
	}
}

func (s *ListaParamsCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitListaParamsCall(s)
	}
}

func (s *ListaParamsCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitListaParamsCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) ListaParamsCall() (localctx IListaParamsCallContext) {
	localctx = NewListaParamsCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ControlParserRULE_listaParamsCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(200)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ControlParserT__19 {
		{
			p.SetState(204)

			var _m = p.Match(ControlParserT__19)

			localctx.(*ListaParamsCallContext).ref = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(207)
		p.expr(0)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ControlParserT__20 {
		{
			p.SetState(208)
			p.Match(ControlParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.ListaParamsCall()
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

// IListaParamsContext is an interface to support dynamic dispatch.
type IListaParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsListaParamsContext differentiates from other interfaces.
	IsListaParamsContext()
}

type ListaParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaParamsContext() *ListaParamsContext {
	var p = new(ListaParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_listaParams
	return p
}

func InitEmptyListaParamsContext(p *ListaParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_listaParams
}

func (*ListaParamsContext) IsListaParamsContext() {}

func NewListaParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaParamsContext {
	var p = new(ListaParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_listaParams

	return p
}

func (s *ListaParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaParamsContext) CopyAll(ctx *ListaParamsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ListaParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumParamsContext struct {
	ListaParamsContext
	ext antlr.Token
	ino antlr.Token
}

func NewNumParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumParamsContext {
	var p = new(NumParamsContext)

	InitEmptyListaParamsContext(&p.ListaParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListaParamsContext))

	return p
}

func (s *NumParamsContext) GetExt() antlr.Token { return s.ext }

func (s *NumParamsContext) GetIno() antlr.Token { return s.ino }

func (s *NumParamsContext) SetExt(v antlr.Token) { s.ext = v }

func (s *NumParamsContext) SetIno(v antlr.Token) { s.ino = v }

func (s *NumParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumParamsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ControlParserID)
}

func (s *NumParamsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ControlParserID, i)
}

func (s *NumParamsContext) Primitivo() IPrimitivoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *NumParamsContext) ListaParams() IListaParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaParamsContext)
}

func (s *NumParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterNumParams(s)
	}
}

func (s *NumParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitNumParams(s)
	}
}

func (s *NumParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitNumParams(s)

	default:
		return t.VisitChildren(s)
	}
}

type OneParamContext struct {
	ListaParamsContext
	ext antlr.Token
	ino antlr.Token
}

func NewOneParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OneParamContext {
	var p = new(OneParamContext)

	InitEmptyListaParamsContext(&p.ListaParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListaParamsContext))

	return p
}

func (s *OneParamContext) GetExt() antlr.Token { return s.ext }

func (s *OneParamContext) GetIno() antlr.Token { return s.ino }

func (s *OneParamContext) SetExt(v antlr.Token) { s.ext = v }

func (s *OneParamContext) SetIno(v antlr.Token) { s.ino = v }

func (s *OneParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneParamContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ControlParserID)
}

func (s *OneParamContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ControlParserID, i)
}

func (s *OneParamContext) Primitivo() IPrimitivoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *OneParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterOneParam(s)
	}
}

func (s *OneParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitOneParam(s)
	}
}

func (s *OneParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitOneParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) ListaParams() (localctx IListaParamsContext) {
	localctx = NewListaParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ControlParserRULE_listaParams)
	var _la int

	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOneParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(213)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(212)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*OneParamContext).ext = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ControlParserT__21 || _la == ControlParserID) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*OneParamContext).ext = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(215)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__22 {
			{
				p.SetState(217)

				var _m = p.Match(ControlParserT__22)

				localctx.(*OneParamContext).ino = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(220)
			p.Primitivo()
		}

	case 2:
		localctx = NewNumParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(222)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(221)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*NumParamsContext).ext = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ControlParserT__21 || _la == ControlParserID) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*NumParamsContext).ext = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(224)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__22 {
			{
				p.SetState(226)

				var _m = p.Match(ControlParserT__22)

				localctx.(*NumParamsContext).ino = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(229)
			p.Primitivo()
		}
		{
			p.SetState(230)
			p.Match(ControlParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.ListaParams()
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

// IListaExpContext is an interface to support dynamic dispatch.
type IListaExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsListaExpContext differentiates from other interfaces.
	IsListaExpContext()
}

type ListaExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaExpContext() *ListaExpContext {
	var p = new(ListaExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_listaExp
	return p
}

func InitEmptyListaExpContext(p *ListaExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_listaExp
}

func (*ListaExpContext) IsListaExpContext() {}

func NewListaExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaExpContext {
	var p = new(ListaExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_listaExp

	return p
}

func (s *ListaExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaExpContext) CopyAll(ctx *ListaExpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ListaExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OneExprContext struct {
	ListaExpContext
}

func NewOneExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OneExprContext {
	var p = new(OneExprContext)

	InitEmptyListaExpContext(&p.ListaExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListaExpContext))

	return p
}

func (s *OneExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OneExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterOneExpr(s)
	}
}

func (s *OneExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitOneExpr(s)
	}
}

func (s *OneExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitOneExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumExprContext struct {
	ListaExpContext
}

func NewNumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumExprContext {
	var p = new(NumExprContext)

	InitEmptyListaExpContext(&p.ListaExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListaExpContext))

	return p
}

func (s *NumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NumExprContext) ListaExp() IListaExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExpContext)
}

func (s *NumExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterNumExpr(s)
	}
}

func (s *NumExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitNumExpr(s)
	}
}

func (s *NumExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitNumExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) ListaExp() (localctx IListaExpContext) {
	localctx = NewListaExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ControlParserRULE_listaExp)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOneExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.expr(0)
		}

	case 2:
		localctx = NewNumExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.expr(0)
		}
		{
			p.SetState(237)
			p.Match(ControlParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.ListaExp()
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

// IPrintlnstmtContext is an interface to support dynamic dispatch.
type IPrintlnstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsPrintlnstmtContext differentiates from other interfaces.
	IsPrintlnstmtContext()
}

type PrintlnstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintlnstmtContext() *PrintlnstmtContext {
	var p = new(PrintlnstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_printlnstmt
	return p
}

func InitEmptyPrintlnstmtContext(p *PrintlnstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_printlnstmt
}

func (*PrintlnstmtContext) IsPrintlnstmtContext() {}

func NewPrintlnstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintlnstmtContext {
	var p = new(PrintlnstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_printlnstmt

	return p
}

func (s *PrintlnstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintlnstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintlnstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintlnstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterPrintlnstmt(s)
	}
}

func (s *PrintlnstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitPrintlnstmt(s)
	}
}

func (s *PrintlnstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitPrintlnstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Printlnstmt() (localctx IPrintlnstmtContext) {
	localctx = NewPrintlnstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ControlParserRULE_printlnstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(ControlParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(ControlParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.expr(0)
	}
	{
		p.SetState(245)
		p.Match(ControlParserT__10)
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

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (s *PrintstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitPrintstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ControlParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(ControlParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Match(ControlParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.expr(0)
	}
	{
		p.SetState(250)
		p.Match(ControlParserT__10)
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

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) CopyAll(ctx *IfstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfNormalContext struct {
	IfstmtContext
}

func NewIfNormalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfNormalContext {
	var p = new(IfNormalContext)

	InitEmptyIfstmtContext(&p.IfstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfstmtContext))

	return p
}

func (s *IfNormalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfNormalContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfNormalContext) Block() IBlockContext {
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

func (s *IfNormalContext) Transfer_sentence() ITransfer_sentenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransfer_sentenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransfer_sentenceContext)
}

func (s *IfNormalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterIfNormal(s)
	}
}

func (s *IfNormalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitIfNormal(s)
	}
}

func (s *IfNormalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitIfNormal(s)

	default:
		return t.VisitChildren(s)
	}
}

type ElseContext struct {
	IfstmtContext
}

func NewElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseContext {
	var p = new(ElseContext)

	InitEmptyIfstmtContext(&p.IfstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfstmtContext))

	return p
}

func (s *ElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ElseContext) AllBlock() []IBlockContext {
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

func (s *ElseContext) Block(i int) IBlockContext {
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

func (s *ElseContext) AllTransfer_sentence() []ITransfer_sentenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITransfer_sentenceContext); ok {
			len++
		}
	}

	tst := make([]ITransfer_sentenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITransfer_sentenceContext); ok {
			tst[i] = t.(ITransfer_sentenceContext)
			i++
		}
	}

	return tst
}

func (s *ElseContext) Transfer_sentence(i int) ITransfer_sentenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransfer_sentenceContext); ok {
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

	return t.(ITransfer_sentenceContext)
}

func (s *ElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterElse(s)
	}
}

func (s *ElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitElse(s)
	}
}

func (s *ElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type Else_ifContext struct {
	IfstmtContext
}

func NewElse_ifContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Else_ifContext {
	var p = new(Else_ifContext)

	InitEmptyIfstmtContext(&p.IfstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfstmtContext))

	return p
}

func (s *Else_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_ifContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Else_ifContext) Block() IBlockContext {
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

func (s *Else_ifContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *Else_ifContext) Transfer_sentence() ITransfer_sentenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransfer_sentenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransfer_sentenceContext)
}

func (s *Else_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterElse_if(s)
	}
}

func (s *Else_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitElse_if(s)
	}
}

func (s *Else_ifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitElse_if(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ControlParserRULE_ifstmt)
	var _la int

	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfNormalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.Match(ControlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.expr(0)
		}
		{
			p.SetState(254)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Block()
		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152921505680719872) != 0 {
			{
				p.SetState(256)
				p.Transfer_sentence()
			}

		}
		{
			p.SetState(259)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.Match(ControlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.expr(0)
		}
		{
			p.SetState(263)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Block()
		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152921505680719872) != 0 {
			{
				p.SetState(265)
				p.Transfer_sentence()
			}

		}
		{
			p.SetState(268)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Match(ControlParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.Block()
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152921505680719872) != 0 {
			{
				p.SetState(272)
				p.Transfer_sentence()
			}

		}
		{
			p.SetState(275)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewElse_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(277)
			p.Match(ControlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)
			p.expr(0)
		}
		{
			p.SetState(279)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.Block()
		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152921505680719872) != 0 {
			{
				p.SetState(281)
				p.Transfer_sentence()
			}

		}
		{
			p.SetState(284)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.Match(ControlParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.Ifstmt()
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

// ISwitchstmtContext is an interface to support dynamic dispatch.
type ISwitchstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Cases() ICasesContext

	// IsSwitchstmtContext differentiates from other interfaces.
	IsSwitchstmtContext()
}

type SwitchstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchstmtContext() *SwitchstmtContext {
	var p = new(SwitchstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_switchstmt
	return p
}

func InitEmptySwitchstmtContext(p *SwitchstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_switchstmt
}

func (*SwitchstmtContext) IsSwitchstmtContext() {}

func NewSwitchstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstmtContext {
	var p = new(SwitchstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_switchstmt

	return p
}

func (s *SwitchstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchstmtContext) Cases() ICasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *SwitchstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitSwitchstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Switchstmt() (localctx ISwitchstmtContext) {
	localctx = NewSwitchstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ControlParserRULE_switchstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(ControlParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.expr(0)
	}
	{
		p.SetState(292)
		p.Match(ControlParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Cases()
	}
	{
		p.SetState(294)
		p.Match(ControlParserT__17)
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

// ICasesContext is an interface to support dynamic dispatch.
type ICasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCaseblock() []ICaseblockContext
	Caseblock(i int) ICaseblockContext

	// IsCasesContext differentiates from other interfaces.
	IsCasesContext()
}

type CasesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCasesContext() *CasesContext {
	var p = new(CasesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_cases
	return p
}

func InitEmptyCasesContext(p *CasesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_cases
}

func (*CasesContext) IsCasesContext() {}

func NewCasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesContext {
	var p = new(CasesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_cases

	return p
}

func (s *CasesContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesContext) AllCaseblock() []ICaseblockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseblockContext); ok {
			len++
		}
	}

	tst := make([]ICaseblockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseblockContext); ok {
			tst[i] = t.(ICaseblockContext)
			i++
		}
	}

	return tst
}

func (s *CasesContext) Caseblock(i int) ICaseblockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseblockContext); ok {
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

	return t.(ICaseblockContext)
}

func (s *CasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterCases(s)
	}
}

func (s *CasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitCases(s)
	}
}

func (s *CasesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitCases(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Cases() (localctx ICasesContext) {
	localctx = NewCasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ControlParserRULE_cases)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ControlParserT__28 || _la == ControlParserT__30 {
		{
			p.SetState(296)
			p.Caseblock()
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

// ICaseblockContext is an interface to support dynamic dispatch.
type ICaseblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCaseblockContext differentiates from other interfaces.
	IsCaseblockContext()
}

type CaseblockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseblockContext() *CaseblockContext {
	var p = new(CaseblockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_caseblock
	return p
}

func InitEmptyCaseblockContext(p *CaseblockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_caseblock
}

func (*CaseblockContext) IsCaseblockContext() {}

func NewCaseblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseblockContext {
	var p = new(CaseblockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_caseblock

	return p
}

func (s *CaseblockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseblockContext) CopyAll(ctx *CaseblockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CaseblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnCaseContext struct {
	CaseblockContext
}

func NewUnCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnCaseContext {
	var p = new(UnCaseContext)

	InitEmptyCaseblockContext(&p.CaseblockContext)
	p.parser = parser
	p.CopyAll(ctx.(*CaseblockContext))

	return p
}

func (s *UnCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnCaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnCaseContext) Block() IBlockContext {
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

func (s *UnCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterUnCase(s)
	}
}

func (s *UnCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitUnCase(s)
	}
}

func (s *UnCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitUnCase(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnDefaultContext struct {
	CaseblockContext
}

func NewUnDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnDefaultContext {
	var p = new(UnDefaultContext)

	InitEmptyCaseblockContext(&p.CaseblockContext)
	p.parser = parser
	p.CopyAll(ctx.(*CaseblockContext))

	return p
}

func (s *UnDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnDefaultContext) Block() IBlockContext {
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

func (s *UnDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterUnDefault(s)
	}
}

func (s *UnDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitUnDefault(s)
	}
}

func (s *UnDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitUnDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Caseblock() (localctx ICaseblockContext) {
	localctx = NewCaseblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ControlParserRULE_caseblock)
	var _la int

	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ControlParserT__28:
		localctx = NewUnCaseContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(302)
			p.Match(ControlParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.expr(0)
		}
		{
			p.SetState(304)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Block()
		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__29 {
			{
				p.SetState(306)
				p.Match(ControlParserT__29)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case ControlParserT__30:
		localctx = NewUnDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(309)
			p.Match(ControlParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)
			p.Block()
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

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhilestmtContext) Block() IBlockContext {
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

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (s *WhilestmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitWhilestmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ControlParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(ControlParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.expr(0)
	}
	{
		p.SetState(316)
		p.Match(ControlParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.Block()
	}
	{
		p.SetState(318)
		p.Match(ControlParserT__17)
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

// IForstmtContext is an interface to support dynamic dispatch.
type IForstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsForstmtContext differentiates from other interfaces.
	IsForstmtContext()
}

type ForstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForstmtContext() *ForstmtContext {
	var p = new(ForstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_forstmt
	return p
}

func InitEmptyForstmtContext(p *ForstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_forstmt
}

func (*ForstmtContext) IsForstmtContext() {}

func NewForstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForstmtContext {
	var p = new(ForstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_forstmt

	return p
}

func (s *ForstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForstmtContext) CopyAll(ctx *ForstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ForstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForEachContext struct {
	ForstmtContext
}

func NewForEachContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForEachContext {
	var p = new(ForEachContext)

	InitEmptyForstmtContext(&p.ForstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForstmtContext))

	return p
}

func (s *ForEachContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForEachContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *ForEachContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForEachContext) Block() IBlockContext {
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

func (s *ForEachContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterForEach(s)
	}
}

func (s *ForEachContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitForEach(s)
	}
}

func (s *ForEachContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitForEach(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForNormalContext struct {
	ForstmtContext
}

func NewForNormalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForNormalContext {
	var p = new(ForNormalContext)

	InitEmptyForstmtContext(&p.ForstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForstmtContext))

	return p
}

func (s *ForNormalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForNormalContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *ForNormalContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForNormalContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ForNormalContext) Block() IBlockContext {
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

func (s *ForNormalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterForNormal(s)
	}
}

func (s *ForNormalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitForNormal(s)
	}
}

func (s *ForNormalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitForNormal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Forstmt() (localctx IForstmtContext) {
	localctx = NewForstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ControlParserRULE_forstmt)
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForNormalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)
			p.Match(ControlParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(ControlParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.expr(0)
		}
		{
			p.SetState(324)
			p.Match(ControlParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			p.expr(0)
		}
		{
			p.SetState(326)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Block()
		}
		{
			p.SetState(328)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewForEachContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(330)
			p.Match(ControlParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(ControlParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.expr(0)
		}
		{
			p.SetState(334)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Block()
		}
		{
			p.SetState(336)
			p.Match(ControlParserT__17)
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

// IGuardstmtContext is an interface to support dynamic dispatch.
type IGuardstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext
	Transfer_sentence() ITransfer_sentenceContext

	// IsGuardstmtContext differentiates from other interfaces.
	IsGuardstmtContext()
}

type GuardstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardstmtContext() *GuardstmtContext {
	var p = new(GuardstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_guardstmt
	return p
}

func InitEmptyGuardstmtContext(p *GuardstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_guardstmt
}

func (*GuardstmtContext) IsGuardstmtContext() {}

func NewGuardstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardstmtContext {
	var p = new(GuardstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_guardstmt

	return p
}

func (s *GuardstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardstmtContext) Block() IBlockContext {
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

func (s *GuardstmtContext) Transfer_sentence() ITransfer_sentenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransfer_sentenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransfer_sentenceContext)
}

func (s *GuardstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterGuardstmt(s)
	}
}

func (s *GuardstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitGuardstmt(s)
	}
}

func (s *GuardstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitGuardstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Guardstmt() (localctx IGuardstmtContext) {
	localctx = NewGuardstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ControlParserRULE_guardstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.Match(ControlParserT__35)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
		p.expr(0)
	}
	{
		p.SetState(342)
		p.Match(ControlParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Match(ControlParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.Block()
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152921505680719872) != 0 {
		{
			p.SetState(345)
			p.Transfer_sentence()
		}

	}
	{
		p.SetState(348)
		p.Match(ControlParserT__17)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolExprContext struct {
	ExprContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatExprContext struct {
	ExprContext
}

func NewFloatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatExprContext {
	var p = new(FloatExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FloatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatExprContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ControlParserFLOAT, 0)
}

func (s *FloatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterFloatExpr(s)
	}
}

func (s *FloatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitFloatExpr(s)
	}
}

func (s *FloatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitFloatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParExprContext struct {
	ExprContext
}

func NewParExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExprContext {
	var p = new(ParExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterParExpr(s)
	}
}

func (s *ParExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitParExpr(s)
	}
}

func (s *ParExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitParExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorGetElementContext struct {
	ExprContext
}

func NewVectorGetElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorGetElementContext {
	var p = new(VectorGetElementContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorGetElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorGetElementContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *VectorGetElementContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VectorGetElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterVectorGetElement(s)
	}
}

func (s *VectorGetElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitVectorGetElement(s)
	}
}

func (s *VectorGetElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitVectorGetElement(s)

	default:
		return t.VisitChildren(s)
	}
}

type StrExprContext struct {
	ExprContext
}

func NewStrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrExprContext {
	var p = new(StrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(ControlParserSTRING, 0)
}

func (s *StrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterStrExpr(s)
	}
}

func (s *StrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitStrExpr(s)
	}
}

func (s *StrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitStrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorIsEmptyContext struct {
	ExprContext
}

func NewVectorIsEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorIsEmptyContext {
	var p = new(VectorIsEmptyContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorIsEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorIsEmptyContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *VectorIsEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterVectorIsEmpty(s)
	}
}

func (s *VectorIsEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitVectorIsEmpty(s)
	}
}

func (s *VectorIsEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitVectorIsEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExprContext
	right IExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRight() IExprContext { return s.right }

func (s *NotExprContext) SetRight(v IExprContext) { s.right = v }

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntExprContext struct {
	ExprContext
}

func NewIntExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntExprContext {
	var p = new(IntExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExprContext) INT() antlr.TerminalNode {
	return s.GetToken(ControlParserINT, 0)
}

func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewOpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpExprContext {
	var p = new(OpExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OpExprContext) GetOp() antlr.Token { return s.op }

func (s *OpExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpExprContext) GetLeft() IExprContext { return s.left }

func (s *OpExprContext) GetRight() IExprContext { return s.right }

func (s *OpExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *OpExprContext) SetRight(v IExprContext) { s.right = v }

func (s *OpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OpExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *OpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterOpExpr(s)
	}
}

func (s *OpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitOpExpr(s)
	}
}

func (s *OpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitOpExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CharExprContext struct {
	ExprContext
}

func NewCharExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharExprContext {
	var p = new(CharExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CharExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharExprContext) CHAR() antlr.TerminalNode {
	return s.GetToken(ControlParserCHAR, 0)
}

func (s *CharExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterCharExpr(s)
	}
}

func (s *CharExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitCharExpr(s)
	}
}

func (s *CharExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitCharExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorCountContext struct {
	ExprContext
}

func NewVectorCountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorCountContext {
	var p = new(VectorCountContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorCountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorCountContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *VectorCountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterVectorCount(s)
	}
}

func (s *VectorCountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitVectorCount(s)
	}
}

func (s *VectorCountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitVectorCount(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ControlParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, ControlParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(351)
			p.Match(ControlParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)

			var _x = p.expr(18)

			localctx.(*NotExprContext).right = _x
		}

	case 2:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(353)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.expr(0)
		}
		{
			p.SetState(355)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(357)
			p.Match(ControlParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(358)
			p.Match(ControlParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(359)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ControlParserT__50 || _la == ControlParserT__51) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 6:
		localctx = NewFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(360)
			p.Match(ControlParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewCharExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(361)
			p.Match(ControlParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(362)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewVectorIsEmptyContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(363)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.Match(ControlParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Match(ControlParserT__52)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewVectorCountContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(366)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.Match(ControlParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.Match(ControlParserT__53)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewVectorGetElementContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(369)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.expr(0)
		}
		{
			p.SetState(372)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(397)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(376)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(377)

					var _m = p.Match(ControlParserT__37)

					localctx.(*OpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(378)

					var _x = p.expr(18)

					localctx.(*OpExprContext).right = _x
				}

			case 2:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(379)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(380)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__38 || _la == ControlParserT__39) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(381)

					var _x = p.expr(17)

					localctx.(*OpExprContext).right = _x
				}

			case 3:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(382)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(383)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__40 || _la == ControlParserT__41) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(384)

					var _x = p.expr(16)

					localctx.(*OpExprContext).right = _x
				}

			case 4:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(385)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(386)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__42 || _la == ControlParserT__43) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(387)

					var _x = p.expr(15)

					localctx.(*OpExprContext).right = _x
				}

			case 5:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(388)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(389)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__44 || _la == ControlParserT__45) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(390)

					var _x = p.expr(14)

					localctx.(*OpExprContext).right = _x
				}

			case 6:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(391)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(392)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__46 || _la == ControlParserT__47) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(393)

					var _x = p.expr(13)

					localctx.(*OpExprContext).right = _x
				}

			case 7:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(394)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(395)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__48 || _la == ControlParserT__49) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(396)

					var _x = p.expr(12)

					localctx.(*OpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
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

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_primitivo
	return p
}

func InitEmptyPrimitivoContext(p *PrimitivoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_primitivo
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }
func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (s *PrimitivoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitPrimitivo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Primitivo() (localctx IPrimitivoContext) {
	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ControlParserRULE_primitivo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1116892707587883008) != 0) {
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

// ITransfer_sentenceContext is an interface to support dynamic dispatch.
type ITransfer_sentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTransfer_sentenceContext differentiates from other interfaces.
	IsTransfer_sentenceContext()
}

type Transfer_sentenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransfer_sentenceContext() *Transfer_sentenceContext {
	var p = new(Transfer_sentenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_transfer_sentence
	return p
}

func InitEmptyTransfer_sentenceContext(p *Transfer_sentenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_transfer_sentence
}

func (*Transfer_sentenceContext) IsTransfer_sentenceContext() {}

func NewTransfer_sentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Transfer_sentenceContext {
	var p = new(Transfer_sentenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_transfer_sentence

	return p
}

func (s *Transfer_sentenceContext) GetParser() antlr.Parser { return s.parser }
func (s *Transfer_sentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Transfer_sentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Transfer_sentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterTransfer_sentence(s)
	}
}

func (s *Transfer_sentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitTransfer_sentence(s)
	}
}

func (s *Transfer_sentenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitTransfer_sentence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Transfer_sentence() (localctx ITransfer_sentenceContext) {
	localctx = NewTransfer_sentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ControlParserRULE_transfer_sentence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152921505680719872) != 0) {
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

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }
func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (s *Var_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitVar_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ControlParserRULE_var_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ControlParserT__60 || _la == ControlParserT__61) {
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

func (p *ControlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ControlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
