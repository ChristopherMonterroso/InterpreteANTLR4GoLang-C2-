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
		"'('", "')'", "'removeLast'", "'remove'", "'at'", "','", "'func'", "'{'",
		"'}'", "'->'", "'return'", "'&'", "'_'", "'inout'", "'println'", "'print'",
		"'if'", "'else'", "'switch'", "'case'", "'break'", "'default'", "'while'",
		"'for'", "'in'", "'...'", "'guard'", "'!'", "'-'", "'%'", "'*'", "'/'",
		"'+'", "'>='", "'>'", "'<='", "'<'", "'=='", "'!='", "'&&'", "'||'",
		"'true'", "'false'", "'isEmpty'", "'count'", "'String'", "'Int'", "'Float'",
		"'Bool'", "'Character'", "'continue'", "'let'", "'var'", "'nil'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "NIL", "INT", "ID",
		"FLOAT", "STRING", "CHAR", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"prog", "block", "stmt", "assignstmt", "vectorPpts", "matrixstmt", "defMatrix",
		"listaValores_Mat", "lista_Expresiones", "listaValores_Mat2", "funcstmt",
		"callFunction", "listaParamsCall", "listaParams", "listaExp", "returnstmt",
		"printlnstmt", "printstmt", "ifstmt", "switchstmt", "cases", "caseblock",
		"whilestmt", "forstmt", "guardstmt", "expr", "primitivo", "transfer_sentence",
		"var_type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 538, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 63, 8,
		1, 10, 1, 12, 1, 66, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 81, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 148, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 171, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 197, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 208, 8, 8, 10, 8, 12, 8, 211, 9, 8, 1, 9,
		1, 9, 1, 9, 3, 9, 216, 8, 9, 1, 9, 1, 9, 1, 9, 5, 9, 221, 8, 9, 10, 9,
		12, 9, 224, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 230, 8, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 241, 8,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 249, 8, 10, 1, 10,
		3, 10, 252, 8, 10, 1, 10, 1, 10, 3, 10, 256, 8, 10, 1, 11, 1, 11, 1, 11,
		3, 11, 261, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 267, 8, 12, 1, 12,
		3, 12, 270, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 276, 8, 12, 1, 12,
		3, 12, 279, 8, 12, 1, 12, 5, 12, 282, 8, 12, 10, 12, 12, 12, 285, 9, 12,
		1, 13, 3, 13, 288, 8, 13, 1, 13, 1, 13, 1, 13, 3, 13, 293, 8, 13, 1, 13,
		3, 13, 296, 8, 13, 1, 13, 1, 13, 3, 13, 300, 8, 13, 1, 13, 3, 13, 303,
		8, 13, 1, 13, 1, 13, 1, 13, 3, 13, 308, 8, 13, 1, 13, 3, 13, 311, 8, 13,
		1, 13, 1, 13, 3, 13, 315, 8, 13, 1, 13, 1, 13, 1, 13, 3, 13, 320, 8, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 327, 8, 14, 1, 15, 1, 15, 3,
		15, 331, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 5, 17, 343, 8, 17, 10, 17, 12, 17, 346, 9, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 355, 8, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 366, 8, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 373, 8, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 382, 8, 18, 1, 18, 1, 18, 3, 18, 386,
		8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 5, 20, 395, 8,
		20, 10, 20, 12, 20, 398, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21,
		405, 8, 21, 1, 21, 1, 21, 1, 21, 3, 21, 410, 8, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		3, 23, 436, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 444,
		8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		3, 25, 466, 8, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 504, 8, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		5, 25, 527, 8, 25, 10, 25, 12, 25, 530, 9, 25, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 28, 0, 2, 18, 50, 29, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 0, 11, 2, 0, 22, 22, 65, 65, 1, 0, 51, 52, 1, 0, 40, 41, 2, 0,
		38, 38, 42, 42, 1, 0, 43, 44, 1, 0, 45, 46, 1, 0, 47, 48, 1, 0, 49, 50,
		1, 0, 55, 59, 2, 0, 30, 30, 60, 60, 1, 0, 61, 62, 595, 0, 58, 1, 0, 0,
		0, 2, 64, 1, 0, 0, 0, 4, 80, 1, 0, 0, 0, 6, 147, 1, 0, 0, 0, 8, 170, 1,
		0, 0, 0, 10, 196, 1, 0, 0, 0, 12, 198, 1, 0, 0, 0, 14, 200, 1, 0, 0, 0,
		16, 204, 1, 0, 0, 0, 18, 215, 1, 0, 0, 0, 20, 255, 1, 0, 0, 0, 22, 257,
		1, 0, 0, 0, 24, 266, 1, 0, 0, 0, 26, 319, 1, 0, 0, 0, 28, 326, 1, 0, 0,
		0, 30, 328, 1, 0, 0, 0, 32, 332, 1, 0, 0, 0, 34, 337, 1, 0, 0, 0, 36, 385,
		1, 0, 0, 0, 38, 387, 1, 0, 0, 0, 40, 396, 1, 0, 0, 0, 42, 409, 1, 0, 0,
		0, 44, 411, 1, 0, 0, 0, 46, 435, 1, 0, 0, 0, 48, 437, 1, 0, 0, 0, 50, 503,
		1, 0, 0, 0, 52, 531, 1, 0, 0, 0, 54, 533, 1, 0, 0, 0, 56, 535, 1, 0, 0,
		0, 58, 59, 3, 2, 1, 0, 59, 60, 5, 0, 0, 1, 60, 1, 1, 0, 0, 0, 61, 63, 3,
		4, 2, 0, 62, 61, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64,
		65, 1, 0, 0, 0, 65, 3, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 81, 3, 6, 3,
		0, 68, 81, 3, 32, 16, 0, 69, 81, 3, 34, 17, 0, 70, 81, 3, 36, 18, 0, 71,
		81, 3, 44, 22, 0, 72, 81, 3, 38, 19, 0, 73, 81, 3, 46, 23, 0, 74, 81, 3,
		48, 24, 0, 75, 81, 3, 8, 4, 0, 76, 81, 3, 10, 5, 0, 77, 81, 3, 20, 10,
		0, 78, 81, 3, 22, 11, 0, 79, 81, 3, 30, 15, 0, 80, 67, 1, 0, 0, 0, 80,
		68, 1, 0, 0, 0, 80, 69, 1, 0, 0, 0, 80, 70, 1, 0, 0, 0, 80, 71, 1, 0, 0,
		0, 80, 72, 1, 0, 0, 0, 80, 73, 1, 0, 0, 0, 80, 74, 1, 0, 0, 0, 80, 75,
		1, 0, 0, 0, 80, 76, 1, 0, 0, 0, 80, 77, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0,
		80, 79, 1, 0, 0, 0, 81, 5, 1, 0, 0, 0, 82, 83, 5, 65, 0, 0, 83, 84, 5,
		1, 0, 0, 84, 148, 3, 50, 25, 0, 85, 86, 3, 56, 28, 0, 86, 87, 5, 65, 0,
		0, 87, 88, 5, 2, 0, 0, 88, 89, 3, 52, 26, 0, 89, 90, 5, 1, 0, 0, 90, 91,
		3, 50, 25, 0, 91, 148, 1, 0, 0, 0, 92, 93, 3, 56, 28, 0, 93, 94, 5, 65,
		0, 0, 94, 95, 5, 2, 0, 0, 95, 96, 3, 52, 26, 0, 96, 97, 5, 3, 0, 0, 97,
		148, 1, 0, 0, 0, 98, 99, 3, 56, 28, 0, 99, 100, 5, 65, 0, 0, 100, 101,
		5, 1, 0, 0, 101, 102, 3, 50, 25, 0, 102, 148, 1, 0, 0, 0, 103, 104, 5,
		65, 0, 0, 104, 105, 5, 4, 0, 0, 105, 148, 3, 50, 25, 0, 106, 107, 5, 65,
		0, 0, 107, 108, 5, 5, 0, 0, 108, 148, 3, 50, 25, 0, 109, 110, 3, 56, 28,
		0, 110, 111, 5, 65, 0, 0, 111, 112, 5, 2, 0, 0, 112, 113, 5, 6, 0, 0, 113,
		114, 3, 52, 26, 0, 114, 115, 5, 7, 0, 0, 115, 116, 5, 1, 0, 0, 116, 117,
		5, 6, 0, 0, 117, 118, 5, 7, 0, 0, 118, 148, 1, 0, 0, 0, 119, 120, 3, 56,
		28, 0, 120, 121, 5, 65, 0, 0, 121, 122, 5, 2, 0, 0, 122, 123, 5, 6, 0,
		0, 123, 124, 3, 52, 26, 0, 124, 125, 5, 7, 0, 0, 125, 126, 5, 1, 0, 0,
		126, 127, 5, 6, 0, 0, 127, 128, 3, 28, 14, 0, 128, 129, 5, 7, 0, 0, 129,
		148, 1, 0, 0, 0, 130, 131, 5, 65, 0, 0, 131, 132, 5, 6, 0, 0, 132, 133,
		3, 50, 25, 0, 133, 134, 5, 7, 0, 0, 134, 135, 5, 1, 0, 0, 135, 136, 3,
		50, 25, 0, 136, 148, 1, 0, 0, 0, 137, 138, 5, 65, 0, 0, 138, 139, 5, 6,
		0, 0, 139, 140, 3, 50, 25, 0, 140, 141, 5, 7, 0, 0, 141, 142, 5, 6, 0,
		0, 142, 143, 3, 50, 25, 0, 143, 144, 5, 7, 0, 0, 144, 145, 5, 1, 0, 0,
		145, 146, 3, 50, 25, 0, 146, 148, 1, 0, 0, 0, 147, 82, 1, 0, 0, 0, 147,
		85, 1, 0, 0, 0, 147, 92, 1, 0, 0, 0, 147, 98, 1, 0, 0, 0, 147, 103, 1,
		0, 0, 0, 147, 106, 1, 0, 0, 0, 147, 109, 1, 0, 0, 0, 147, 119, 1, 0, 0,
		0, 147, 130, 1, 0, 0, 0, 147, 137, 1, 0, 0, 0, 148, 7, 1, 0, 0, 0, 149,
		150, 5, 65, 0, 0, 150, 151, 5, 8, 0, 0, 151, 152, 5, 9, 0, 0, 152, 153,
		5, 10, 0, 0, 153, 154, 3, 50, 25, 0, 154, 155, 5, 11, 0, 0, 155, 171, 1,
		0, 0, 0, 156, 157, 5, 65, 0, 0, 157, 158, 5, 8, 0, 0, 158, 159, 5, 12,
		0, 0, 159, 160, 5, 10, 0, 0, 160, 171, 5, 11, 0, 0, 161, 162, 5, 65, 0,
		0, 162, 163, 5, 8, 0, 0, 163, 164, 5, 13, 0, 0, 164, 165, 5, 10, 0, 0,
		165, 166, 5, 14, 0, 0, 166, 167, 5, 2, 0, 0, 167, 168, 3, 50, 25, 0, 168,
		169, 5, 11, 0, 0, 169, 171, 1, 0, 0, 0, 170, 149, 1, 0, 0, 0, 170, 156,
		1, 0, 0, 0, 170, 161, 1, 0, 0, 0, 171, 9, 1, 0, 0, 0, 172, 173, 3, 56,
		28, 0, 173, 174, 5, 65, 0, 0, 174, 175, 5, 2, 0, 0, 175, 176, 5, 6, 0,
		0, 176, 177, 5, 6, 0, 0, 177, 178, 3, 52, 26, 0, 178, 179, 5, 7, 0, 0,
		179, 180, 5, 7, 0, 0, 180, 181, 5, 1, 0, 0, 181, 182, 3, 12, 6, 0, 182,
		197, 1, 0, 0, 0, 183, 184, 3, 56, 28, 0, 184, 185, 5, 65, 0, 0, 185, 186,
		5, 2, 0, 0, 186, 187, 5, 6, 0, 0, 187, 188, 5, 6, 0, 0, 188, 189, 5, 6,
		0, 0, 189, 190, 3, 52, 26, 0, 190, 191, 5, 7, 0, 0, 191, 192, 5, 7, 0,
		0, 192, 193, 5, 7, 0, 0, 193, 194, 5, 1, 0, 0, 194, 195, 3, 12, 6, 0, 195,
		197, 1, 0, 0, 0, 196, 172, 1, 0, 0, 0, 196, 183, 1, 0, 0, 0, 197, 11, 1,
		0, 0, 0, 198, 199, 3, 14, 7, 0, 199, 13, 1, 0, 0, 0, 200, 201, 5, 6, 0,
		0, 201, 202, 3, 18, 9, 0, 202, 203, 5, 7, 0, 0, 203, 15, 1, 0, 0, 0, 204,
		209, 3, 50, 25, 0, 205, 206, 5, 15, 0, 0, 206, 208, 3, 50, 25, 0, 207,
		205, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210,
		1, 0, 0, 0, 210, 17, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 212, 213, 6, 9,
		-1, 0, 213, 216, 3, 14, 7, 0, 214, 216, 3, 16, 8, 0, 215, 212, 1, 0, 0,
		0, 215, 214, 1, 0, 0, 0, 216, 222, 1, 0, 0, 0, 217, 218, 10, 3, 0, 0, 218,
		219, 5, 15, 0, 0, 219, 221, 3, 14, 7, 0, 220, 217, 1, 0, 0, 0, 221, 224,
		1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 19, 1, 0,
		0, 0, 224, 222, 1, 0, 0, 0, 225, 226, 5, 16, 0, 0, 226, 227, 5, 65, 0,
		0, 227, 229, 5, 10, 0, 0, 228, 230, 3, 26, 13, 0, 229, 228, 1, 0, 0, 0,
		229, 230, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 232, 5, 11, 0, 0, 232,
		233, 5, 17, 0, 0, 233, 234, 3, 2, 1, 0, 234, 235, 5, 18, 0, 0, 235, 256,
		1, 0, 0, 0, 236, 237, 5, 16, 0, 0, 237, 238, 5, 65, 0, 0, 238, 240, 5,
		10, 0, 0, 239, 241, 3, 26, 13, 0, 240, 239, 1, 0, 0, 0, 240, 241, 1, 0,
		0, 0, 241, 242, 1, 0, 0, 0, 242, 243, 5, 11, 0, 0, 243, 244, 5, 19, 0,
		0, 244, 245, 3, 52, 26, 0, 245, 246, 5, 17, 0, 0, 246, 248, 3, 2, 1, 0,
		247, 249, 5, 20, 0, 0, 248, 247, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249,
		251, 1, 0, 0, 0, 250, 252, 3, 50, 25, 0, 251, 250, 1, 0, 0, 0, 251, 252,
		1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 254, 5, 18, 0, 0, 254, 256, 1, 0,
		0, 0, 255, 225, 1, 0, 0, 0, 255, 236, 1, 0, 0, 0, 256, 21, 1, 0, 0, 0,
		257, 258, 5, 65, 0, 0, 258, 260, 5, 10, 0, 0, 259, 261, 3, 24, 12, 0, 260,
		259, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263,
		5, 11, 0, 0, 263, 23, 1, 0, 0, 0, 264, 265, 5, 65, 0, 0, 265, 267, 5, 2,
		0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 269, 1, 0, 0, 0,
		268, 270, 5, 21, 0, 0, 269, 268, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270,
		271, 1, 0, 0, 0, 271, 283, 3, 50, 25, 0, 272, 275, 5, 15, 0, 0, 273, 274,
		5, 65, 0, 0, 274, 276, 5, 2, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0,
		0, 0, 276, 278, 1, 0, 0, 0, 277, 279, 5, 21, 0, 0, 278, 277, 1, 0, 0, 0,
		278, 279, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 282, 3, 50, 25, 0, 281,
		272, 1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284,
		1, 0, 0, 0, 284, 25, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 286, 288, 7, 0,
		0, 0, 287, 286, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0,
		289, 290, 5, 65, 0, 0, 290, 292, 5, 2, 0, 0, 291, 293, 5, 23, 0, 0, 292,
		291, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 295, 1, 0, 0, 0, 294, 296,
		5, 6, 0, 0, 295, 294, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 297, 1, 0,
		0, 0, 297, 299, 3, 52, 26, 0, 298, 300, 5, 7, 0, 0, 299, 298, 1, 0, 0,
		0, 299, 300, 1, 0, 0, 0, 300, 320, 1, 0, 0, 0, 301, 303, 7, 0, 0, 0, 302,
		301, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305,
		5, 65, 0, 0, 305, 307, 5, 2, 0, 0, 306, 308, 5, 23, 0, 0, 307, 306, 1,
		0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 310, 1, 0, 0, 0, 309, 311, 5, 6, 0,
		0, 310, 309, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312,
		314, 3, 52, 26, 0, 313, 315, 5, 7, 0, 0, 314, 313, 1, 0, 0, 0, 314, 315,
		1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 317, 5, 15, 0, 0, 317, 318, 3, 26,
		13, 0, 318, 320, 1, 0, 0, 0, 319, 287, 1, 0, 0, 0, 319, 302, 1, 0, 0, 0,
		320, 27, 1, 0, 0, 0, 321, 327, 3, 50, 25, 0, 322, 323, 3, 50, 25, 0, 323,
		324, 5, 15, 0, 0, 324, 325, 3, 28, 14, 0, 325, 327, 1, 0, 0, 0, 326, 321,
		1, 0, 0, 0, 326, 322, 1, 0, 0, 0, 327, 29, 1, 0, 0, 0, 328, 330, 5, 20,
		0, 0, 329, 331, 3, 50, 25, 0, 330, 329, 1, 0, 0, 0, 330, 331, 1, 0, 0,
		0, 331, 31, 1, 0, 0, 0, 332, 333, 5, 24, 0, 0, 333, 334, 5, 10, 0, 0, 334,
		335, 3, 50, 25, 0, 335, 336, 5, 11, 0, 0, 336, 33, 1, 0, 0, 0, 337, 338,
		5, 25, 0, 0, 338, 339, 5, 10, 0, 0, 339, 344, 3, 50, 25, 0, 340, 341, 5,
		15, 0, 0, 341, 343, 3, 50, 25, 0, 342, 340, 1, 0, 0, 0, 343, 346, 1, 0,
		0, 0, 344, 342, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 347, 1, 0, 0, 0,
		346, 344, 1, 0, 0, 0, 347, 348, 5, 11, 0, 0, 348, 35, 1, 0, 0, 0, 349,
		350, 5, 26, 0, 0, 350, 351, 3, 50, 25, 0, 351, 352, 5, 17, 0, 0, 352, 354,
		3, 2, 1, 0, 353, 355, 3, 54, 27, 0, 354, 353, 1, 0, 0, 0, 354, 355, 1,
		0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 357, 5, 18, 0, 0, 357, 358, 5, 27,
		0, 0, 358, 359, 3, 36, 18, 0, 359, 386, 1, 0, 0, 0, 360, 361, 5, 26, 0,
		0, 361, 362, 3, 50, 25, 0, 362, 363, 5, 17, 0, 0, 363, 365, 3, 2, 1, 0,
		364, 366, 3, 54, 27, 0, 365, 364, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366,
		367, 1, 0, 0, 0, 367, 368, 5, 18, 0, 0, 368, 369, 5, 27, 0, 0, 369, 370,
		5, 17, 0, 0, 370, 372, 3, 2, 1, 0, 371, 373, 3, 54, 27, 0, 372, 371, 1,
		0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 375, 5, 18, 0,
		0, 375, 386, 1, 0, 0, 0, 376, 377, 5, 26, 0, 0, 377, 378, 3, 50, 25, 0,
		378, 379, 5, 17, 0, 0, 379, 381, 3, 2, 1, 0, 380, 382, 3, 54, 27, 0, 381,
		380, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 384,
		5, 18, 0, 0, 384, 386, 1, 0, 0, 0, 385, 349, 1, 0, 0, 0, 385, 360, 1, 0,
		0, 0, 385, 376, 1, 0, 0, 0, 386, 37, 1, 0, 0, 0, 387, 388, 5, 28, 0, 0,
		388, 389, 3, 50, 25, 0, 389, 390, 5, 17, 0, 0, 390, 391, 3, 40, 20, 0,
		391, 392, 5, 18, 0, 0, 392, 39, 1, 0, 0, 0, 393, 395, 3, 42, 21, 0, 394,
		393, 1, 0, 0, 0, 395, 398, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 396, 397,
		1, 0, 0, 0, 397, 41, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 399, 400, 5, 29,
		0, 0, 400, 401, 3, 50, 25, 0, 401, 402, 5, 2, 0, 0, 402, 404, 3, 2, 1,
		0, 403, 405, 5, 30, 0, 0, 404, 403, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405,
		410, 1, 0, 0, 0, 406, 407, 5, 31, 0, 0, 407, 408, 5, 2, 0, 0, 408, 410,
		3, 2, 1, 0, 409, 399, 1, 0, 0, 0, 409, 406, 1, 0, 0, 0, 410, 43, 1, 0,
		0, 0, 411, 412, 5, 32, 0, 0, 412, 413, 3, 50, 25, 0, 413, 414, 5, 17, 0,
		0, 414, 415, 3, 2, 1, 0, 415, 416, 5, 18, 0, 0, 416, 45, 1, 0, 0, 0, 417,
		418, 5, 33, 0, 0, 418, 419, 5, 65, 0, 0, 419, 420, 5, 34, 0, 0, 420, 421,
		3, 50, 25, 0, 421, 422, 5, 35, 0, 0, 422, 423, 3, 50, 25, 0, 423, 424,
		5, 17, 0, 0, 424, 425, 3, 2, 1, 0, 425, 426, 5, 18, 0, 0, 426, 436, 1,
		0, 0, 0, 427, 428, 5, 33, 0, 0, 428, 429, 5, 65, 0, 0, 429, 430, 5, 34,
		0, 0, 430, 431, 3, 50, 25, 0, 431, 432, 5, 17, 0, 0, 432, 433, 3, 2, 1,
		0, 433, 434, 5, 18, 0, 0, 434, 436, 1, 0, 0, 0, 435, 417, 1, 0, 0, 0, 435,
		427, 1, 0, 0, 0, 436, 47, 1, 0, 0, 0, 437, 438, 5, 36, 0, 0, 438, 439,
		3, 50, 25, 0, 439, 440, 5, 27, 0, 0, 440, 441, 5, 17, 0, 0, 441, 443, 3,
		2, 1, 0, 442, 444, 3, 54, 27, 0, 443, 442, 1, 0, 0, 0, 443, 444, 1, 0,
		0, 0, 444, 445, 1, 0, 0, 0, 445, 446, 5, 18, 0, 0, 446, 49, 1, 0, 0, 0,
		447, 448, 6, 25, -1, 0, 448, 449, 5, 37, 0, 0, 449, 504, 3, 50, 25, 25,
		450, 451, 5, 38, 0, 0, 451, 504, 3, 50, 25, 24, 452, 453, 5, 10, 0, 0,
		453, 454, 3, 50, 25, 0, 454, 455, 5, 11, 0, 0, 455, 504, 1, 0, 0, 0, 456,
		504, 5, 63, 0, 0, 457, 504, 5, 64, 0, 0, 458, 504, 5, 67, 0, 0, 459, 504,
		7, 1, 0, 0, 460, 504, 5, 66, 0, 0, 461, 504, 5, 68, 0, 0, 462, 463, 5,
		65, 0, 0, 463, 465, 5, 10, 0, 0, 464, 466, 3, 24, 12, 0, 465, 464, 1, 0,
		0, 0, 465, 466, 1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 504, 5, 11, 0, 0,
		468, 504, 5, 65, 0, 0, 469, 470, 5, 65, 0, 0, 470, 471, 5, 8, 0, 0, 471,
		504, 5, 53, 0, 0, 472, 473, 5, 65, 0, 0, 473, 474, 5, 8, 0, 0, 474, 504,
		5, 54, 0, 0, 475, 476, 5, 65, 0, 0, 476, 477, 5, 6, 0, 0, 477, 478, 3,
		50, 25, 0, 478, 479, 5, 7, 0, 0, 479, 504, 1, 0, 0, 0, 480, 481, 5, 65,
		0, 0, 481, 482, 5, 6, 0, 0, 482, 483, 3, 50, 25, 0, 483, 484, 5, 7, 0,
		0, 484, 485, 5, 6, 0, 0, 485, 486, 3, 50, 25, 0, 486, 487, 5, 7, 0, 0,
		487, 504, 1, 0, 0, 0, 488, 489, 5, 55, 0, 0, 489, 490, 5, 10, 0, 0, 490,
		491, 3, 50, 25, 0, 491, 492, 5, 11, 0, 0, 492, 504, 1, 0, 0, 0, 493, 494,
		5, 56, 0, 0, 494, 495, 5, 10, 0, 0, 495, 496, 3, 50, 25, 0, 496, 497, 5,
		11, 0, 0, 497, 504, 1, 0, 0, 0, 498, 499, 5, 57, 0, 0, 499, 500, 5, 10,
		0, 0, 500, 501, 3, 50, 25, 0, 501, 502, 5, 11, 0, 0, 502, 504, 1, 0, 0,
		0, 503, 447, 1, 0, 0, 0, 503, 450, 1, 0, 0, 0, 503, 452, 1, 0, 0, 0, 503,
		456, 1, 0, 0, 0, 503, 457, 1, 0, 0, 0, 503, 458, 1, 0, 0, 0, 503, 459,
		1, 0, 0, 0, 503, 460, 1, 0, 0, 0, 503, 461, 1, 0, 0, 0, 503, 462, 1, 0,
		0, 0, 503, 468, 1, 0, 0, 0, 503, 469, 1, 0, 0, 0, 503, 472, 1, 0, 0, 0,
		503, 475, 1, 0, 0, 0, 503, 480, 1, 0, 0, 0, 503, 488, 1, 0, 0, 0, 503,
		493, 1, 0, 0, 0, 503, 498, 1, 0, 0, 0, 504, 528, 1, 0, 0, 0, 505, 506,
		10, 23, 0, 0, 506, 507, 5, 39, 0, 0, 507, 527, 3, 50, 25, 24, 508, 509,
		10, 22, 0, 0, 509, 510, 7, 2, 0, 0, 510, 527, 3, 50, 25, 23, 511, 512,
		10, 21, 0, 0, 512, 513, 7, 3, 0, 0, 513, 527, 3, 50, 25, 22, 514, 515,
		10, 20, 0, 0, 515, 516, 7, 4, 0, 0, 516, 527, 3, 50, 25, 21, 517, 518,
		10, 19, 0, 0, 518, 519, 7, 5, 0, 0, 519, 527, 3, 50, 25, 20, 520, 521,
		10, 18, 0, 0, 521, 522, 7, 6, 0, 0, 522, 527, 3, 50, 25, 19, 523, 524,
		10, 17, 0, 0, 524, 525, 7, 7, 0, 0, 525, 527, 3, 50, 25, 18, 526, 505,
		1, 0, 0, 0, 526, 508, 1, 0, 0, 0, 526, 511, 1, 0, 0, 0, 526, 514, 1, 0,
		0, 0, 526, 517, 1, 0, 0, 0, 526, 520, 1, 0, 0, 0, 526, 523, 1, 0, 0, 0,
		527, 530, 1, 0, 0, 0, 528, 526, 1, 0, 0, 0, 528, 529, 1, 0, 0, 0, 529,
		51, 1, 0, 0, 0, 530, 528, 1, 0, 0, 0, 531, 532, 7, 8, 0, 0, 532, 53, 1,
		0, 0, 0, 533, 534, 7, 9, 0, 0, 534, 55, 1, 0, 0, 0, 535, 536, 7, 10, 0,
		0, 536, 57, 1, 0, 0, 0, 45, 64, 80, 147, 170, 196, 209, 215, 222, 229,
		240, 248, 251, 255, 260, 266, 269, 275, 278, 283, 287, 292, 295, 299, 302,
		307, 310, 314, 319, 326, 330, 344, 354, 365, 372, 381, 385, 396, 404, 409,
		435, 443, 465, 503, 526, 528,
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
	ControlParserNIL          = 63
	ControlParserINT          = 64
	ControlParserID           = 65
	ControlParserFLOAT        = 66
	ControlParserSTRING       = 67
	ControlParserCHAR         = 68
	ControlParserWHITESPACE   = 69
	ControlParserCOMMENT      = 70
	ControlParserLINE_COMMENT = 71
)

// ControlParser rules.
const (
	ControlParserRULE_prog              = 0
	ControlParserRULE_block             = 1
	ControlParserRULE_stmt              = 2
	ControlParserRULE_assignstmt        = 3
	ControlParserRULE_vectorPpts        = 4
	ControlParserRULE_matrixstmt        = 5
	ControlParserRULE_defMatrix         = 6
	ControlParserRULE_listaValores_Mat  = 7
	ControlParserRULE_lista_Expresiones = 8
	ControlParserRULE_listaValores_Mat2 = 9
	ControlParserRULE_funcstmt          = 10
	ControlParserRULE_callFunction      = 11
	ControlParserRULE_listaParamsCall   = 12
	ControlParserRULE_listaParams       = 13
	ControlParserRULE_listaExp          = 14
	ControlParserRULE_returnstmt        = 15
	ControlParserRULE_printlnstmt       = 16
	ControlParserRULE_printstmt         = 17
	ControlParserRULE_ifstmt            = 18
	ControlParserRULE_switchstmt        = 19
	ControlParserRULE_cases             = 20
	ControlParserRULE_caseblock         = 21
	ControlParserRULE_whilestmt         = 22
	ControlParserRULE_forstmt           = 23
	ControlParserRULE_guardstmt         = 24
	ControlParserRULE_expr              = 25
	ControlParserRULE_primitivo         = 26
	ControlParserRULE_transfer_sentence = 27
	ControlParserRULE_var_type          = 28
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
		p.SetState(58)
		p.Block()
	}
	{
		p.SetState(59)
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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(61)
				p.Stmt()
			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
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
	Matrixstmt() IMatrixstmtContext
	Funcstmt() IFuncstmtContext
	CallFunction() ICallFunctionContext
	Returnstmt() IReturnstmtContext

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

func (s *StmtContext) Matrixstmt() IMatrixstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixstmtContext)
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

func (s *StmtContext) CallFunction() ICallFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFunctionContext)
}

func (s *StmtContext) Returnstmt() IReturnstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnstmtContext)
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
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Assignstmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.Printlnstmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.Printstmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.Ifstmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(71)
			p.Whilestmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(72)
			p.Switchstmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(73)
			p.Forstmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(74)
			p.Guardstmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(75)
			p.VectorPpts()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(76)
			p.Matrixstmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(77)
			p.Funcstmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(78)
			p.CallFunction()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(79)
			p.Returnstmt()
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

type ReasignacionMatrixTwoDContext struct {
	AssignstmtContext
}

func NewReasignacionMatrixTwoDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReasignacionMatrixTwoDContext {
	var p = new(ReasignacionMatrixTwoDContext)

	InitEmptyAssignstmtContext(&p.AssignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignstmtContext))

	return p
}

func (s *ReasignacionMatrixTwoDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReasignacionMatrixTwoDContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *ReasignacionMatrixTwoDContext) AllExpr() []IExprContext {
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

func (s *ReasignacionMatrixTwoDContext) Expr(i int) IExprContext {
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

func (s *ReasignacionMatrixTwoDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterReasignacionMatrixTwoD(s)
	}
}

func (s *ReasignacionMatrixTwoDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitReasignacionMatrixTwoD(s)
	}
}

func (s *ReasignacionMatrixTwoDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitReasignacionMatrixTwoD(s)

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
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewReasignacionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
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

	case 2:
		localctx = NewAsignacionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Var_type()
		}
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
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Primitivo()
		}
		{
			p.SetState(89)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.expr(0)
		}

	case 3:
		localctx = NewAsignacionNoExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
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
			p.Primitivo()
		}
		{
			p.SetState(96)
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
			p.SetState(98)
			p.Var_type()
		}
		{
			p.SetState(99)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.expr(0)
		}

	case 5:
		localctx = NewIncrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
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
			p.Match(ControlParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.expr(0)
		}

	case 6:
		localctx = NewDecrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(106)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(ControlParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.expr(0)
		}

	case 7:
		localctx = NewAsignacionVectorVacioContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(109)
			p.Var_type()
		}
		{
			p.SetState(110)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Primitivo()
		}
		{
			p.SetState(114)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
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
			p.SetState(119)
			p.Var_type()
		}
		{
			p.SetState(120)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Primitivo()
		}
		{
			p.SetState(124)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.ListaExp()
		}
		{
			p.SetState(128)
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
			p.SetState(130)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.expr(0)
		}
		{
			p.SetState(133)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.expr(0)
		}

	case 10:
		localctx = NewReasignacionMatrixTwoDContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(137)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.expr(0)
		}
		{
			p.SetState(140)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.expr(0)
		}
		{
			p.SetState(143)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
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
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVectorAppendContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(ControlParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(ControlParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.expr(0)
		}
		{
			p.SetState(154)
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
			p.SetState(156)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Match(ControlParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(ControlParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
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
			p.SetState(161)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Match(ControlParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Match(ControlParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Match(ControlParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.expr(0)
		}
		{
			p.SetState(168)
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

// IMatrixstmtContext is an interface to support dynamic dispatch.
type IMatrixstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMatrixstmtContext differentiates from other interfaces.
	IsMatrixstmtContext()
}

type MatrixstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixstmtContext() *MatrixstmtContext {
	var p = new(MatrixstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_matrixstmt
	return p
}

func InitEmptyMatrixstmtContext(p *MatrixstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_matrixstmt
}

func (*MatrixstmtContext) IsMatrixstmtContext() {}

func NewMatrixstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixstmtContext {
	var p = new(MatrixstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_matrixstmt

	return p
}

func (s *MatrixstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixstmtContext) CopyAll(ctx *MatrixstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MatrixstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MatrixTwoDContext struct {
	MatrixstmtContext
}

func NewMatrixTwoDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrixTwoDContext {
	var p = new(MatrixTwoDContext)

	InitEmptyMatrixstmtContext(&p.MatrixstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*MatrixstmtContext))

	return p
}

func (s *MatrixTwoDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixTwoDContext) Var_type() IVar_typeContext {
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

func (s *MatrixTwoDContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *MatrixTwoDContext) Primitivo() IPrimitivoContext {
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

func (s *MatrixTwoDContext) DefMatrix() IDefMatrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefMatrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefMatrixContext)
}

func (s *MatrixTwoDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterMatrixTwoD(s)
	}
}

func (s *MatrixTwoDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitMatrixTwoD(s)
	}
}

func (s *MatrixTwoDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitMatrixTwoD(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatrixThreeDContext struct {
	MatrixstmtContext
}

func NewMatrixThreeDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrixThreeDContext {
	var p = new(MatrixThreeDContext)

	InitEmptyMatrixstmtContext(&p.MatrixstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*MatrixstmtContext))

	return p
}

func (s *MatrixThreeDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixThreeDContext) Var_type() IVar_typeContext {
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

func (s *MatrixThreeDContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *MatrixThreeDContext) Primitivo() IPrimitivoContext {
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

func (s *MatrixThreeDContext) DefMatrix() IDefMatrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefMatrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefMatrixContext)
}

func (s *MatrixThreeDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterMatrixThreeD(s)
	}
}

func (s *MatrixThreeDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitMatrixThreeD(s)
	}
}

func (s *MatrixThreeDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitMatrixThreeD(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Matrixstmt() (localctx IMatrixstmtContext) {
	localctx = NewMatrixstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ControlParserRULE_matrixstmt)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMatrixTwoDContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Var_type()
		}
		{
			p.SetState(173)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Primitivo()
		}
		{
			p.SetState(178)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.DefMatrix()
		}

	case 2:
		localctx = NewMatrixThreeDContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Var_type()
		}
		{
			p.SetState(184)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Primitivo()
		}
		{
			p.SetState(190)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.DefMatrix()
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

// IDefMatrixContext is an interface to support dynamic dispatch.
type IDefMatrixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ListaValores_Mat() IListaValores_MatContext

	// IsDefMatrixContext differentiates from other interfaces.
	IsDefMatrixContext()
}

type DefMatrixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefMatrixContext() *DefMatrixContext {
	var p = new(DefMatrixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_defMatrix
	return p
}

func InitEmptyDefMatrixContext(p *DefMatrixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_defMatrix
}

func (*DefMatrixContext) IsDefMatrixContext() {}

func NewDefMatrixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefMatrixContext {
	var p = new(DefMatrixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_defMatrix

	return p
}

func (s *DefMatrixContext) GetParser() antlr.Parser { return s.parser }

func (s *DefMatrixContext) ListaValores_Mat() IListaValores_MatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaValores_MatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaValores_MatContext)
}

func (s *DefMatrixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefMatrixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefMatrixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterDefMatrix(s)
	}
}

func (s *DefMatrixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitDefMatrix(s)
	}
}

func (s *DefMatrixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitDefMatrix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) DefMatrix() (localctx IDefMatrixContext) {
	localctx = NewDefMatrixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ControlParserRULE_defMatrix)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.ListaValores_Mat()
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

// IListaValores_MatContext is an interface to support dynamic dispatch.
type IListaValores_MatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ListaValores_Mat2() IListaValores_Mat2Context

	// IsListaValores_MatContext differentiates from other interfaces.
	IsListaValores_MatContext()
}

type ListaValores_MatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaValores_MatContext() *ListaValores_MatContext {
	var p = new(ListaValores_MatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_listaValores_Mat
	return p
}

func InitEmptyListaValores_MatContext(p *ListaValores_MatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_listaValores_Mat
}

func (*ListaValores_MatContext) IsListaValores_MatContext() {}

func NewListaValores_MatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaValores_MatContext {
	var p = new(ListaValores_MatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_listaValores_Mat

	return p
}

func (s *ListaValores_MatContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaValores_MatContext) ListaValores_Mat2() IListaValores_Mat2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaValores_Mat2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaValores_Mat2Context)
}

func (s *ListaValores_MatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaValores_MatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaValores_MatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterListaValores_Mat(s)
	}
}

func (s *ListaValores_MatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitListaValores_Mat(s)
	}
}

func (s *ListaValores_MatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitListaValores_Mat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) ListaValores_Mat() (localctx IListaValores_MatContext) {
	localctx = NewListaValores_MatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ControlParserRULE_listaValores_Mat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(ControlParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.listaValores_Mat2(0)
	}
	{
		p.SetState(202)
		p.Match(ControlParserT__6)
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

// ILista_ExpresionesContext is an interface to support dynamic dispatch.
type ILista_ExpresionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsLista_ExpresionesContext differentiates from other interfaces.
	IsLista_ExpresionesContext()
}

type Lista_ExpresionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_ExpresionesContext() *Lista_ExpresionesContext {
	var p = new(Lista_ExpresionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_lista_Expresiones
	return p
}

func InitEmptyLista_ExpresionesContext(p *Lista_ExpresionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_lista_Expresiones
}

func (*Lista_ExpresionesContext) IsLista_ExpresionesContext() {}

func NewLista_ExpresionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_ExpresionesContext {
	var p = new(Lista_ExpresionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_lista_Expresiones

	return p
}

func (s *Lista_ExpresionesContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_ExpresionesContext) AllExpr() []IExprContext {
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

func (s *Lista_ExpresionesContext) Expr(i int) IExprContext {
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

func (s *Lista_ExpresionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_ExpresionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_ExpresionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterLista_Expresiones(s)
	}
}

func (s *Lista_ExpresionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitLista_Expresiones(s)
	}
}

func (s *Lista_ExpresionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitLista_Expresiones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Lista_Expresiones() (localctx ILista_ExpresionesContext) {
	localctx = NewLista_ExpresionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ControlParserRULE_lista_Expresiones)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.expr(0)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(205)
				p.Match(ControlParserT__14)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(206)
				p.expr(0)
			}

		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaValores_Mat2Context is an interface to support dynamic dispatch.
type IListaValores_Mat2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ListaValores_Mat() IListaValores_MatContext
	Lista_Expresiones() ILista_ExpresionesContext
	ListaValores_Mat2() IListaValores_Mat2Context

	// IsListaValores_Mat2Context differentiates from other interfaces.
	IsListaValores_Mat2Context()
}

type ListaValores_Mat2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaValores_Mat2Context() *ListaValores_Mat2Context {
	var p = new(ListaValores_Mat2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_listaValores_Mat2
	return p
}

func InitEmptyListaValores_Mat2Context(p *ListaValores_Mat2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_listaValores_Mat2
}

func (*ListaValores_Mat2Context) IsListaValores_Mat2Context() {}

func NewListaValores_Mat2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaValores_Mat2Context {
	var p = new(ListaValores_Mat2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_listaValores_Mat2

	return p
}

func (s *ListaValores_Mat2Context) GetParser() antlr.Parser { return s.parser }

func (s *ListaValores_Mat2Context) ListaValores_Mat() IListaValores_MatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaValores_MatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaValores_MatContext)
}

func (s *ListaValores_Mat2Context) Lista_Expresiones() ILista_ExpresionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_ExpresionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_ExpresionesContext)
}

func (s *ListaValores_Mat2Context) ListaValores_Mat2() IListaValores_Mat2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaValores_Mat2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaValores_Mat2Context)
}

func (s *ListaValores_Mat2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaValores_Mat2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaValores_Mat2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterListaValores_Mat2(s)
	}
}

func (s *ListaValores_Mat2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitListaValores_Mat2(s)
	}
}

func (s *ListaValores_Mat2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitListaValores_Mat2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) ListaValores_Mat2() (localctx IListaValores_Mat2Context) {
	return p.listaValores_Mat2(0)
}

func (p *ControlParser) listaValores_Mat2(_p int) (localctx IListaValores_Mat2Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListaValores_Mat2Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListaValores_Mat2Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, ControlParserRULE_listaValores_Mat2, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ControlParserT__5:
		{
			p.SetState(213)
			p.ListaValores_Mat()
		}

	case ControlParserT__9, ControlParserT__36, ControlParserT__37, ControlParserT__50, ControlParserT__51, ControlParserT__54, ControlParserT__55, ControlParserT__56, ControlParserNIL, ControlParserINT, ControlParserID, ControlParserFLOAT, ControlParserSTRING, ControlParserCHAR:
		{
			p.SetState(214)
			p.Lista_Expresiones()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListaValores_Mat2Context(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_listaValores_Mat2)
			p.SetState(217)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(218)
				p.Match(ControlParserT__14)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(219)
				p.ListaValores_Mat()
			}

		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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

type Func_conRetorno_conTipoContext struct {
	FuncstmtContext
	ret antlr.Token
	exp IExprContext
}

func NewFunc_conRetorno_conTipoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_conRetorno_conTipoContext {
	var p = new(Func_conRetorno_conTipoContext)

	InitEmptyFuncstmtContext(&p.FuncstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncstmtContext))

	return p
}

func (s *Func_conRetorno_conTipoContext) GetRet() antlr.Token { return s.ret }

func (s *Func_conRetorno_conTipoContext) SetRet(v antlr.Token) { s.ret = v }

func (s *Func_conRetorno_conTipoContext) GetExp() IExprContext { return s.exp }

func (s *Func_conRetorno_conTipoContext) SetExp(v IExprContext) { s.exp = v }

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

func (s *Func_conRetorno_conTipoContext) ListaParams() IListaParamsContext {
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

type Func_sinRetornoContext struct {
	FuncstmtContext
}

func NewFunc_sinRetornoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_sinRetornoContext {
	var p = new(Func_sinRetornoContext)

	InitEmptyFuncstmtContext(&p.FuncstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncstmtContext))

	return p
}

func (s *Func_sinRetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_sinRetornoContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *Func_sinRetornoContext) Block() IBlockContext {
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

func (s *Func_sinRetornoContext) ListaParams() IListaParamsContext {
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

func (s *Func_sinRetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterFunc_sinRetorno(s)
	}
}

func (s *Func_sinRetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitFunc_sinRetorno(s)
	}
}

func (s *Func_sinRetornoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitFunc_sinRetorno(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Funcstmt() (localctx IFuncstmtContext) {
	localctx = NewFuncstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ControlParserRULE_funcstmt)
	var _la int

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunc_sinRetornoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__21 || _la == ControlParserID {
			{
				p.SetState(228)
				p.ListaParams()
			}

		}
		{
			p.SetState(231)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Match(ControlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Block()
		}
		{
			p.SetState(234)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunc_conRetorno_conTipoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.Match(ControlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__21 || _la == ControlParserID {
			{
				p.SetState(239)
				p.ListaParams()
			}

		}
		{
			p.SetState(242)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Match(ControlParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Primitivo()
		}
		{
			p.SetState(245)
			p.Match(ControlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Block()
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__19 {
			{
				p.SetState(247)

				var _m = p.Match(ControlParserT__19)

				localctx.(*Func_conRetorno_conTipoContext).ret = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-10)) & ^0x3f) == 0 && ((int64(1)<<(_la-10))&567706441125724161) != 0 {
			{
				p.SetState(250)

				var _x = p.expr(0)

				localctx.(*Func_conRetorno_conTipoContext).exp = _x
			}

		}
		{
			p.SetState(253)
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

// ICallFunctionContext is an interface to support dynamic dispatch.
type ICallFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ListaParamsCall() IListaParamsCallContext

	// IsCallFunctionContext differentiates from other interfaces.
	IsCallFunctionContext()
}

type CallFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallFunctionContext() *CallFunctionContext {
	var p = new(CallFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_callFunction
	return p
}

func InitEmptyCallFunctionContext(p *CallFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_callFunction
}

func (*CallFunctionContext) IsCallFunctionContext() {}

func NewCallFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFunctionContext {
	var p = new(CallFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_callFunction

	return p
}

func (s *CallFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *CallFunctionContext) ListaParamsCall() IListaParamsCallContext {
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

func (s *CallFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *ControlParser) CallFunction() (localctx ICallFunctionContext) {
	localctx = NewCallFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ControlParserRULE_callFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(ControlParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(ControlParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-10)) & ^0x3f) == 0 && ((int64(1)<<(_la-10))&567706441125726209) != 0 {
		{
			p.SetState(259)
			p.ListaParamsCall()
		}

	}
	{
		p.SetState(262)
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
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

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

func (s *ListaParamsCallContext) AllExpr() []IExprContext {
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

func (s *ListaParamsCallContext) Expr(i int) IExprContext {
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

func (s *ListaParamsCallContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ControlParserID)
}

func (s *ListaParamsCallContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ControlParserID, i)
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
	p.EnterRule(localctx, 24, ControlParserRULE_listaParamsCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(266)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(264)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ControlParserT__20 {
		{
			p.SetState(268)

			var _m = p.Match(ControlParserT__20)

			localctx.(*ListaParamsCallContext).ref = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(271)
		p.expr(0)
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ControlParserT__14 {
		{
			p.SetState(272)
			p.Match(ControlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(273)
				p.Match(ControlParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(274)
				p.Match(ControlParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__20 {
			{
				p.SetState(277)

				var _m = p.Match(ControlParserT__20)

				localctx.(*ListaParamsCallContext).ref = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(280)
			p.expr(0)
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
	ext      antlr.Token
	ino      antlr.Token
	isVector antlr.Token
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

func (s *NumParamsContext) GetIsVector() antlr.Token { return s.isVector }

func (s *NumParamsContext) SetExt(v antlr.Token) { s.ext = v }

func (s *NumParamsContext) SetIno(v antlr.Token) { s.ino = v }

func (s *NumParamsContext) SetIsVector(v antlr.Token) { s.isVector = v }

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
	ext      antlr.Token
	ino      antlr.Token
	isVector antlr.Token
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

func (s *OneParamContext) GetIsVector() antlr.Token { return s.isVector }

func (s *OneParamContext) SetExt(v antlr.Token) { s.ext = v }

func (s *OneParamContext) SetIno(v antlr.Token) { s.ino = v }

func (s *OneParamContext) SetIsVector(v antlr.Token) { s.isVector = v }

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
	p.EnterRule(localctx, 26, ControlParserRULE_listaParams)
	var _la int

	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOneParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(287)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(286)

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
			p.SetState(289)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__22 {
			{
				p.SetState(291)

				var _m = p.Match(ControlParserT__22)

				localctx.(*OneParamContext).ino = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__5 {
			{
				p.SetState(294)
				p.Match(ControlParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(297)
			p.Primitivo()
		}
		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__6 {
			{
				p.SetState(298)

				var _m = p.Match(ControlParserT__6)

				localctx.(*OneParamContext).isVector = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		localctx = NewNumParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(302)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(301)

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
			p.SetState(304)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__22 {
			{
				p.SetState(306)

				var _m = p.Match(ControlParserT__22)

				localctx.(*NumParamsContext).ino = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__5 {
			{
				p.SetState(309)
				p.Match(ControlParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(312)
			p.Primitivo()
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__6 {
			{
				p.SetState(313)

				var _m = p.Match(ControlParserT__6)

				localctx.(*NumParamsContext).isVector = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(316)
			p.Match(ControlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
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
	p.EnterRule(localctx, 28, ControlParserRULE_listaExp)
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOneExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.expr(0)
		}

	case 2:
		localctx = NewNumExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.expr(0)
		}
		{
			p.SetState(323)
			p.Match(ControlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
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

// IReturnstmtContext is an interface to support dynamic dispatch.
type IReturnstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsReturnstmtContext differentiates from other interfaces.
	IsReturnstmtContext()
}

type ReturnstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnstmtContext() *ReturnstmtContext {
	var p = new(ReturnstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_returnstmt
	return p
}

func InitEmptyReturnstmtContext(p *ReturnstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_returnstmt
}

func (*ReturnstmtContext) IsReturnstmtContext() {}

func NewReturnstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnstmtContext {
	var p = new(ReturnstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_returnstmt

	return p
}

func (s *ReturnstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnstmtContext) Expr() IExprContext {
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

func (s *ReturnstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterReturnstmt(s)
	}
}

func (s *ReturnstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitReturnstmt(s)
	}
}

func (s *ReturnstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitReturnstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Returnstmt() (localctx IReturnstmtContext) {
	localctx = NewReturnstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ControlParserRULE_returnstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(ControlParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(329)
			p.expr(0)
		}

	} else if p.HasError() { // JIM
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
	p.EnterRule(localctx, 32, ControlParserRULE_printlnstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(ControlParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(333)
		p.Match(ControlParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(334)
		p.expr(0)
	}
	{
		p.SetState(335)
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
	AllExpr() []IExprContext
	Expr(i int) IExprContext

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

func (s *PrintstmtContext) AllExpr() []IExprContext {
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

func (s *PrintstmtContext) Expr(i int) IExprContext {
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
	p.EnterRule(localctx, 34, ControlParserRULE_printstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(ControlParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(338)
		p.Match(ControlParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(339)
		p.expr(0)
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ControlParserT__14 {
		{
			p.SetState(340)
			p.Match(ControlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.expr(0)
		}

		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(347)
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
	p.EnterRule(localctx, 36, ControlParserRULE_ifstmt)
	var _la int

	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewElse_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)
			p.Match(ControlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.expr(0)
		}
		{
			p.SetState(351)
			p.Match(ControlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Block()
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__29 || _la == ControlParserT__59 {
			{
				p.SetState(353)
				p.Transfer_sentence()
			}

		}
		{
			p.SetState(356)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Match(ControlParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)
			p.Ifstmt()
		}

	case 2:
		localctx = NewElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(360)
			p.Match(ControlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.expr(0)
		}
		{
			p.SetState(362)
			p.Match(ControlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.Block()
		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__29 || _la == ControlParserT__59 {
			{
				p.SetState(364)
				p.Transfer_sentence()
			}

		}
		{
			p.SetState(367)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.Match(ControlParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Match(ControlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.Block()
		}
		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__29 || _la == ControlParserT__59 {
			{
				p.SetState(371)
				p.Transfer_sentence()
			}

		}
		{
			p.SetState(374)
			p.Match(ControlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewIfNormalContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(376)
			p.Match(ControlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.expr(0)
		}
		{
			p.SetState(378)
			p.Match(ControlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.Block()
		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__29 || _la == ControlParserT__59 {
			{
				p.SetState(380)
				p.Transfer_sentence()
			}

		}
		{
			p.SetState(383)
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
	p.EnterRule(localctx, 38, ControlParserRULE_switchstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(ControlParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(388)
		p.expr(0)
	}
	{
		p.SetState(389)
		p.Match(ControlParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(390)
		p.Cases()
	}
	{
		p.SetState(391)
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
	p.EnterRule(localctx, 40, ControlParserRULE_cases)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ControlParserT__28 || _la == ControlParserT__30 {
		{
			p.SetState(393)
			p.Caseblock()
		}

		p.SetState(398)
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
	p.EnterRule(localctx, 42, ControlParserRULE_caseblock)
	var _la int

	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ControlParserT__28:
		localctx = NewUnCaseContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(399)
			p.Match(ControlParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.expr(0)
		}
		{
			p.SetState(401)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.Block()
		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ControlParserT__29 {
			{
				p.SetState(403)
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
			p.SetState(406)
			p.Match(ControlParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
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
	p.EnterRule(localctx, 44, ControlParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.Match(ControlParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
		p.expr(0)
	}
	{
		p.SetState(413)
		p.Match(ControlParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(414)
		p.Block()
	}
	{
		p.SetState(415)
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
	p.EnterRule(localctx, 46, ControlParserRULE_forstmt)
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForNormalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(417)
			p.Match(ControlParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Match(ControlParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)
			p.expr(0)
		}
		{
			p.SetState(421)
			p.Match(ControlParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.expr(0)
		}
		{
			p.SetState(423)
			p.Match(ControlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Block()
		}
		{
			p.SetState(425)
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
			p.SetState(427)
			p.Match(ControlParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.Match(ControlParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.expr(0)
		}
		{
			p.SetState(431)
			p.Match(ControlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)
			p.Block()
		}
		{
			p.SetState(433)
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
	p.EnterRule(localctx, 48, ControlParserRULE_guardstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(437)
		p.Match(ControlParserT__35)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(438)
		p.expr(0)
	}
	{
		p.SetState(439)
		p.Match(ControlParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)
		p.Match(ControlParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)
		p.Block()
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ControlParserT__29 || _la == ControlParserT__59 {
		{
			p.SetState(442)
			p.Transfer_sentence()
		}

	}
	{
		p.SetState(445)
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

type ToIntContext struct {
	ExprContext
}

func NewToIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToIntContext {
	var p = new(ToIntContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ToIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToIntContext) Expr() IExprContext {
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

func (s *ToIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterToInt(s)
	}
}

func (s *ToIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitToInt(s)
	}
}

func (s *ToIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitToInt(s)

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

type ToFloatContext struct {
	ExprContext
}

func NewToFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToFloatContext {
	var p = new(ToFloatContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ToFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToFloatContext) Expr() IExprContext {
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

func (s *ToFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterToFloat(s)
	}
}

func (s *ToFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitToFloat(s)
	}
}

func (s *ToFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitToFloat(s)

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

type AccesoMatrixTwoDContext struct {
	ExprContext
}

func NewAccesoMatrixTwoDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccesoMatrixTwoDContext {
	var p = new(AccesoMatrixTwoDContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AccesoMatrixTwoDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoMatrixTwoDContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *AccesoMatrixTwoDContext) AllExpr() []IExprContext {
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

func (s *AccesoMatrixTwoDContext) Expr(i int) IExprContext {
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

func (s *AccesoMatrixTwoDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterAccesoMatrixTwoD(s)
	}
}

func (s *AccesoMatrixTwoDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitAccesoMatrixTwoD(s)
	}
}

func (s *AccesoMatrixTwoDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitAccesoMatrixTwoD(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilExprContext struct {
	ExprContext
}

func NewNilExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilExprContext {
	var p = new(NilExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NilExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilExprContext) NIL() antlr.TerminalNode {
	return s.GetToken(ControlParserNIL, 0)
}

func (s *NilExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterNilExpr(s)
	}
}

func (s *NilExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitNilExpr(s)
	}
}

func (s *NilExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitNilExpr(s)

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

type NegExprContext struct {
	ExprContext
	right IExprContext
}

func NewNegExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegExprContext {
	var p = new(NegExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NegExprContext) GetRight() IExprContext { return s.right }

func (s *NegExprContext) SetRight(v IExprContext) { s.right = v }

func (s *NegExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegExprContext) Expr() IExprContext {
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

func (s *NegExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterNegExpr(s)
	}
}

func (s *NegExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitNegExpr(s)
	}
}

func (s *NegExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitNegExpr(s)

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

type CallFuncAsExprContext struct {
	ExprContext
}

func NewCallFuncAsExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallFuncAsExprContext {
	var p = new(CallFuncAsExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CallFuncAsExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFuncAsExprContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *CallFuncAsExprContext) ListaParamsCall() IListaParamsCallContext {
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

func (s *CallFuncAsExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterCallFuncAsExpr(s)
	}
}

func (s *CallFuncAsExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitCallFuncAsExpr(s)
	}
}

func (s *CallFuncAsExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitCallFuncAsExpr(s)

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

type ToStringContext struct {
	ExprContext
}

func NewToStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToStringContext {
	var p = new(ToStringContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ToStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToStringContext) Expr() IExprContext {
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

func (s *ToStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterToString(s)
	}
}

func (s *ToStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitToString(s)
	}
}

func (s *ToStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitToString(s)

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

func (p *ControlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ControlParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, ControlParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(448)
			p.Match(ControlParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)

			var _x = p.expr(25)

			localctx.(*NotExprContext).right = _x
		}

	case 2:
		localctx = NewNegExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(450)
			p.Match(ControlParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)

			var _x = p.expr(24)

			localctx.(*NegExprContext).right = _x
		}

	case 3:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(452)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
			p.expr(0)
		}
		{
			p.SetState(454)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(456)
			p.Match(ControlParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(457)
			p.Match(ControlParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(458)
			p.Match(ControlParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(459)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ControlParserT__50 || _la == ControlParserT__51) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 8:
		localctx = NewFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(460)
			p.Match(ControlParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewCharExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(461)
			p.Match(ControlParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewCallFuncAsExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(462)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(465)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-10)) & ^0x3f) == 0 && ((int64(1)<<(_la-10))&567706441125726209) != 0 {
			{
				p.SetState(464)
				p.ListaParamsCall()
			}

		}
		{
			p.SetState(467)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(468)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewVectorIsEmptyContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(469)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.Match(ControlParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
			p.Match(ControlParserT__52)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewVectorCountContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(472)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(473)
			p.Match(ControlParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(474)
			p.Match(ControlParserT__53)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewVectorGetElementContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(475)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)
			p.expr(0)
		}
		{
			p.SetState(478)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewAccesoMatrixTwoDContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(480)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(481)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(482)
			p.expr(0)
		}
		{
			p.SetState(483)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(484)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)
			p.expr(0)
		}
		{
			p.SetState(486)
			p.Match(ControlParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		localctx = NewToStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(488)
			p.Match(ControlParserT__54)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(489)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.expr(0)
		}
		{
			p.SetState(491)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 17:
		localctx = NewToIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(493)
			p.Match(ControlParserT__55)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(495)
			p.expr(0)
		}
		{
			p.SetState(496)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 18:
		localctx = NewToFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(498)
			p.Match(ControlParserT__56)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(499)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(500)
			p.expr(0)
		}
		{
			p.SetState(501)
			p.Match(ControlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(528)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(526)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(505)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(506)

					var _m = p.Match(ControlParserT__38)

					localctx.(*OpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(507)

					var _x = p.expr(24)

					localctx.(*OpExprContext).right = _x
				}

			case 2:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(508)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(509)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__39 || _la == ControlParserT__40) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(510)

					var _x = p.expr(23)

					localctx.(*OpExprContext).right = _x
				}

			case 3:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(511)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(512)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__37 || _la == ControlParserT__41) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(513)

					var _x = p.expr(22)

					localctx.(*OpExprContext).right = _x
				}

			case 4:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(514)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(515)

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
					p.SetState(516)

					var _x = p.expr(21)

					localctx.(*OpExprContext).right = _x
				}

			case 5:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(517)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(518)

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
					p.SetState(519)

					var _x = p.expr(20)

					localctx.(*OpExprContext).right = _x
				}

			case 6:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(520)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(521)

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
					p.SetState(522)

					var _x = p.expr(19)

					localctx.(*OpExprContext).right = _x
				}

			case 7:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(523)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(524)

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
					p.SetState(525)

					var _x = p.expr(18)

					localctx.(*OpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(530)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 52, ControlParserRULE_primitivo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(531)
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
	p.EnterRule(localctx, 54, ControlParserRULE_transfer_sentence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(533)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ControlParserT__29 || _la == ControlParserT__59) {
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
	p.EnterRule(localctx, 56, ControlParserRULE_var_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(535)
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
	case 9:
		var t *ListaValores_Mat2Context = nil
		if localctx != nil {
			t = localctx.(*ListaValores_Mat2Context)
		}
		return p.ListaValores_Mat2_Sempred(t, predIndex)

	case 25:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ControlParser) ListaValores_Mat2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ControlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 17)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
