// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ControlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ControlLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func controllexerLexerInit() {
	staticData := &ControlLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='", "':'", "'?'", "'+='", "'-='", "'['", "']'", "'.'", "'append'",
		"'('", "')'", "'removeLast'", "'remove'", "'at'", "','", "'println'",
		"'print'", "'if'", "'{'", "'}'", "'else'", "'switch'", "'case'", "'break'",
		"'default'", "'while'", "'for'", "'in'", "'...'", "'guard'", "'!'",
		"'%'", "'*'", "'/'", "'+'", "'-'", "'>='", "'>'", "'<='", "'<'", "'=='",
		"'!='", "'&&'", "'||'", "'true'", "'false'", "'isEmpty'", "'count'",
		"'Int'", "'String'", "'Float'", "'Bool'", "'Character'", "'continue'",
		"'return'", "'let'", "'var'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "INT", "ID", "FLOAT", "STRING", "CHAR",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
		"T__49", "T__50", "T__51", "T__52", "T__53", "T__54", "T__55", "T__56",
		"INT", "ID", "FLOAT", "STRING", "CHAR", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 63, 414, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1,
		43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45,
		1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52,
		1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1,
		53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54,
		1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 57, 4, 57, 371,
		8, 57, 11, 57, 12, 57, 372, 1, 58, 1, 58, 5, 58, 377, 8, 58, 10, 58, 12,
		58, 380, 9, 58, 1, 59, 4, 59, 383, 8, 59, 11, 59, 12, 59, 384, 1, 59, 1,
		59, 4, 59, 389, 8, 59, 11, 59, 12, 59, 390, 1, 60, 1, 60, 1, 60, 1, 60,
		5, 60, 397, 8, 60, 10, 60, 12, 60, 400, 9, 60, 1, 60, 1, 60, 1, 61, 1,
		61, 1, 61, 1, 61, 1, 62, 4, 62, 409, 8, 62, 11, 62, 12, 62, 410, 1, 62,
		1, 62, 0, 0, 63, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71,
		36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89,
		45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53,
		107, 54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61,
		123, 62, 125, 63, 1, 0, 6, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 10, 10, 13, 13, 34, 34, 3,
		0, 10, 10, 13, 13, 39, 39, 2, 0, 9, 10, 32, 32, 420, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1,
		0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87,
		1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0,
		95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0,
		0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109,
		1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0,
		0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1,
		0, 0, 0, 0, 125, 1, 0, 0, 0, 1, 127, 1, 0, 0, 0, 3, 129, 1, 0, 0, 0, 5,
		131, 1, 0, 0, 0, 7, 133, 1, 0, 0, 0, 9, 136, 1, 0, 0, 0, 11, 139, 1, 0,
		0, 0, 13, 141, 1, 0, 0, 0, 15, 143, 1, 0, 0, 0, 17, 145, 1, 0, 0, 0, 19,
		152, 1, 0, 0, 0, 21, 154, 1, 0, 0, 0, 23, 156, 1, 0, 0, 0, 25, 167, 1,
		0, 0, 0, 27, 174, 1, 0, 0, 0, 29, 177, 1, 0, 0, 0, 31, 179, 1, 0, 0, 0,
		33, 187, 1, 0, 0, 0, 35, 193, 1, 0, 0, 0, 37, 196, 1, 0, 0, 0, 39, 198,
		1, 0, 0, 0, 41, 200, 1, 0, 0, 0, 43, 205, 1, 0, 0, 0, 45, 212, 1, 0, 0,
		0, 47, 217, 1, 0, 0, 0, 49, 223, 1, 0, 0, 0, 51, 231, 1, 0, 0, 0, 53, 237,
		1, 0, 0, 0, 55, 241, 1, 0, 0, 0, 57, 244, 1, 0, 0, 0, 59, 248, 1, 0, 0,
		0, 61, 254, 1, 0, 0, 0, 63, 256, 1, 0, 0, 0, 65, 258, 1, 0, 0, 0, 67, 260,
		1, 0, 0, 0, 69, 262, 1, 0, 0, 0, 71, 264, 1, 0, 0, 0, 73, 266, 1, 0, 0,
		0, 75, 269, 1, 0, 0, 0, 77, 271, 1, 0, 0, 0, 79, 274, 1, 0, 0, 0, 81, 276,
		1, 0, 0, 0, 83, 279, 1, 0, 0, 0, 85, 282, 1, 0, 0, 0, 87, 285, 1, 0, 0,
		0, 89, 288, 1, 0, 0, 0, 91, 293, 1, 0, 0, 0, 93, 299, 1, 0, 0, 0, 95, 307,
		1, 0, 0, 0, 97, 313, 1, 0, 0, 0, 99, 317, 1, 0, 0, 0, 101, 324, 1, 0, 0,
		0, 103, 330, 1, 0, 0, 0, 105, 335, 1, 0, 0, 0, 107, 345, 1, 0, 0, 0, 109,
		354, 1, 0, 0, 0, 111, 361, 1, 0, 0, 0, 113, 365, 1, 0, 0, 0, 115, 370,
		1, 0, 0, 0, 117, 374, 1, 0, 0, 0, 119, 382, 1, 0, 0, 0, 121, 392, 1, 0,
		0, 0, 123, 403, 1, 0, 0, 0, 125, 408, 1, 0, 0, 0, 127, 128, 5, 61, 0, 0,
		128, 2, 1, 0, 0, 0, 129, 130, 5, 58, 0, 0, 130, 4, 1, 0, 0, 0, 131, 132,
		5, 63, 0, 0, 132, 6, 1, 0, 0, 0, 133, 134, 5, 43, 0, 0, 134, 135, 5, 61,
		0, 0, 135, 8, 1, 0, 0, 0, 136, 137, 5, 45, 0, 0, 137, 138, 5, 61, 0, 0,
		138, 10, 1, 0, 0, 0, 139, 140, 5, 91, 0, 0, 140, 12, 1, 0, 0, 0, 141, 142,
		5, 93, 0, 0, 142, 14, 1, 0, 0, 0, 143, 144, 5, 46, 0, 0, 144, 16, 1, 0,
		0, 0, 145, 146, 5, 97, 0, 0, 146, 147, 5, 112, 0, 0, 147, 148, 5, 112,
		0, 0, 148, 149, 5, 101, 0, 0, 149, 150, 5, 110, 0, 0, 150, 151, 5, 100,
		0, 0, 151, 18, 1, 0, 0, 0, 152, 153, 5, 40, 0, 0, 153, 20, 1, 0, 0, 0,
		154, 155, 5, 41, 0, 0, 155, 22, 1, 0, 0, 0, 156, 157, 5, 114, 0, 0, 157,
		158, 5, 101, 0, 0, 158, 159, 5, 109, 0, 0, 159, 160, 5, 111, 0, 0, 160,
		161, 5, 118, 0, 0, 161, 162, 5, 101, 0, 0, 162, 163, 5, 76, 0, 0, 163,
		164, 5, 97, 0, 0, 164, 165, 5, 115, 0, 0, 165, 166, 5, 116, 0, 0, 166,
		24, 1, 0, 0, 0, 167, 168, 5, 114, 0, 0, 168, 169, 5, 101, 0, 0, 169, 170,
		5, 109, 0, 0, 170, 171, 5, 111, 0, 0, 171, 172, 5, 118, 0, 0, 172, 173,
		5, 101, 0, 0, 173, 26, 1, 0, 0, 0, 174, 175, 5, 97, 0, 0, 175, 176, 5,
		116, 0, 0, 176, 28, 1, 0, 0, 0, 177, 178, 5, 44, 0, 0, 178, 30, 1, 0, 0,
		0, 179, 180, 5, 112, 0, 0, 180, 181, 5, 114, 0, 0, 181, 182, 5, 105, 0,
		0, 182, 183, 5, 110, 0, 0, 183, 184, 5, 116, 0, 0, 184, 185, 5, 108, 0,
		0, 185, 186, 5, 110, 0, 0, 186, 32, 1, 0, 0, 0, 187, 188, 5, 112, 0, 0,
		188, 189, 5, 114, 0, 0, 189, 190, 5, 105, 0, 0, 190, 191, 5, 110, 0, 0,
		191, 192, 5, 116, 0, 0, 192, 34, 1, 0, 0, 0, 193, 194, 5, 105, 0, 0, 194,
		195, 5, 102, 0, 0, 195, 36, 1, 0, 0, 0, 196, 197, 5, 123, 0, 0, 197, 38,
		1, 0, 0, 0, 198, 199, 5, 125, 0, 0, 199, 40, 1, 0, 0, 0, 200, 201, 5, 101,
		0, 0, 201, 202, 5, 108, 0, 0, 202, 203, 5, 115, 0, 0, 203, 204, 5, 101,
		0, 0, 204, 42, 1, 0, 0, 0, 205, 206, 5, 115, 0, 0, 206, 207, 5, 119, 0,
		0, 207, 208, 5, 105, 0, 0, 208, 209, 5, 116, 0, 0, 209, 210, 5, 99, 0,
		0, 210, 211, 5, 104, 0, 0, 211, 44, 1, 0, 0, 0, 212, 213, 5, 99, 0, 0,
		213, 214, 5, 97, 0, 0, 214, 215, 5, 115, 0, 0, 215, 216, 5, 101, 0, 0,
		216, 46, 1, 0, 0, 0, 217, 218, 5, 98, 0, 0, 218, 219, 5, 114, 0, 0, 219,
		220, 5, 101, 0, 0, 220, 221, 5, 97, 0, 0, 221, 222, 5, 107, 0, 0, 222,
		48, 1, 0, 0, 0, 223, 224, 5, 100, 0, 0, 224, 225, 5, 101, 0, 0, 225, 226,
		5, 102, 0, 0, 226, 227, 5, 97, 0, 0, 227, 228, 5, 117, 0, 0, 228, 229,
		5, 108, 0, 0, 229, 230, 5, 116, 0, 0, 230, 50, 1, 0, 0, 0, 231, 232, 5,
		119, 0, 0, 232, 233, 5, 104, 0, 0, 233, 234, 5, 105, 0, 0, 234, 235, 5,
		108, 0, 0, 235, 236, 5, 101, 0, 0, 236, 52, 1, 0, 0, 0, 237, 238, 5, 102,
		0, 0, 238, 239, 5, 111, 0, 0, 239, 240, 5, 114, 0, 0, 240, 54, 1, 0, 0,
		0, 241, 242, 5, 105, 0, 0, 242, 243, 5, 110, 0, 0, 243, 56, 1, 0, 0, 0,
		244, 245, 5, 46, 0, 0, 245, 246, 5, 46, 0, 0, 246, 247, 5, 46, 0, 0, 247,
		58, 1, 0, 0, 0, 248, 249, 5, 103, 0, 0, 249, 250, 5, 117, 0, 0, 250, 251,
		5, 97, 0, 0, 251, 252, 5, 114, 0, 0, 252, 253, 5, 100, 0, 0, 253, 60, 1,
		0, 0, 0, 254, 255, 5, 33, 0, 0, 255, 62, 1, 0, 0, 0, 256, 257, 5, 37, 0,
		0, 257, 64, 1, 0, 0, 0, 258, 259, 5, 42, 0, 0, 259, 66, 1, 0, 0, 0, 260,
		261, 5, 47, 0, 0, 261, 68, 1, 0, 0, 0, 262, 263, 5, 43, 0, 0, 263, 70,
		1, 0, 0, 0, 264, 265, 5, 45, 0, 0, 265, 72, 1, 0, 0, 0, 266, 267, 5, 62,
		0, 0, 267, 268, 5, 61, 0, 0, 268, 74, 1, 0, 0, 0, 269, 270, 5, 62, 0, 0,
		270, 76, 1, 0, 0, 0, 271, 272, 5, 60, 0, 0, 272, 273, 5, 61, 0, 0, 273,
		78, 1, 0, 0, 0, 274, 275, 5, 60, 0, 0, 275, 80, 1, 0, 0, 0, 276, 277, 5,
		61, 0, 0, 277, 278, 5, 61, 0, 0, 278, 82, 1, 0, 0, 0, 279, 280, 5, 33,
		0, 0, 280, 281, 5, 61, 0, 0, 281, 84, 1, 0, 0, 0, 282, 283, 5, 38, 0, 0,
		283, 284, 5, 38, 0, 0, 284, 86, 1, 0, 0, 0, 285, 286, 5, 124, 0, 0, 286,
		287, 5, 124, 0, 0, 287, 88, 1, 0, 0, 0, 288, 289, 5, 116, 0, 0, 289, 290,
		5, 114, 0, 0, 290, 291, 5, 117, 0, 0, 291, 292, 5, 101, 0, 0, 292, 90,
		1, 0, 0, 0, 293, 294, 5, 102, 0, 0, 294, 295, 5, 97, 0, 0, 295, 296, 5,
		108, 0, 0, 296, 297, 5, 115, 0, 0, 297, 298, 5, 101, 0, 0, 298, 92, 1,
		0, 0, 0, 299, 300, 5, 105, 0, 0, 300, 301, 5, 115, 0, 0, 301, 302, 5, 69,
		0, 0, 302, 303, 5, 109, 0, 0, 303, 304, 5, 112, 0, 0, 304, 305, 5, 116,
		0, 0, 305, 306, 5, 121, 0, 0, 306, 94, 1, 0, 0, 0, 307, 308, 5, 99, 0,
		0, 308, 309, 5, 111, 0, 0, 309, 310, 5, 117, 0, 0, 310, 311, 5, 110, 0,
		0, 311, 312, 5, 116, 0, 0, 312, 96, 1, 0, 0, 0, 313, 314, 5, 73, 0, 0,
		314, 315, 5, 110, 0, 0, 315, 316, 5, 116, 0, 0, 316, 98, 1, 0, 0, 0, 317,
		318, 5, 83, 0, 0, 318, 319, 5, 116, 0, 0, 319, 320, 5, 114, 0, 0, 320,
		321, 5, 105, 0, 0, 321, 322, 5, 110, 0, 0, 322, 323, 5, 103, 0, 0, 323,
		100, 1, 0, 0, 0, 324, 325, 5, 70, 0, 0, 325, 326, 5, 108, 0, 0, 326, 327,
		5, 111, 0, 0, 327, 328, 5, 97, 0, 0, 328, 329, 5, 116, 0, 0, 329, 102,
		1, 0, 0, 0, 330, 331, 5, 66, 0, 0, 331, 332, 5, 111, 0, 0, 332, 333, 5,
		111, 0, 0, 333, 334, 5, 108, 0, 0, 334, 104, 1, 0, 0, 0, 335, 336, 5, 67,
		0, 0, 336, 337, 5, 104, 0, 0, 337, 338, 5, 97, 0, 0, 338, 339, 5, 114,
		0, 0, 339, 340, 5, 97, 0, 0, 340, 341, 5, 99, 0, 0, 341, 342, 5, 116, 0,
		0, 342, 343, 5, 101, 0, 0, 343, 344, 5, 114, 0, 0, 344, 106, 1, 0, 0, 0,
		345, 346, 5, 99, 0, 0, 346, 347, 5, 111, 0, 0, 347, 348, 5, 110, 0, 0,
		348, 349, 5, 116, 0, 0, 349, 350, 5, 105, 0, 0, 350, 351, 5, 110, 0, 0,
		351, 352, 5, 117, 0, 0, 352, 353, 5, 101, 0, 0, 353, 108, 1, 0, 0, 0, 354,
		355, 5, 114, 0, 0, 355, 356, 5, 101, 0, 0, 356, 357, 5, 116, 0, 0, 357,
		358, 5, 117, 0, 0, 358, 359, 5, 114, 0, 0, 359, 360, 5, 110, 0, 0, 360,
		110, 1, 0, 0, 0, 361, 362, 5, 108, 0, 0, 362, 363, 5, 101, 0, 0, 363, 364,
		5, 116, 0, 0, 364, 112, 1, 0, 0, 0, 365, 366, 5, 118, 0, 0, 366, 367, 5,
		97, 0, 0, 367, 368, 5, 114, 0, 0, 368, 114, 1, 0, 0, 0, 369, 371, 7, 0,
		0, 0, 370, 369, 1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0,
		372, 373, 1, 0, 0, 0, 373, 116, 1, 0, 0, 0, 374, 378, 7, 1, 0, 0, 375,
		377, 7, 2, 0, 0, 376, 375, 1, 0, 0, 0, 377, 380, 1, 0, 0, 0, 378, 376,
		1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 118, 1, 0, 0, 0, 380, 378, 1, 0,
		0, 0, 381, 383, 7, 0, 0, 0, 382, 381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0,
		384, 382, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386,
		388, 5, 46, 0, 0, 387, 389, 7, 0, 0, 0, 388, 387, 1, 0, 0, 0, 389, 390,
		1, 0, 0, 0, 390, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 120, 1, 0,
		0, 0, 392, 398, 5, 34, 0, 0, 393, 397, 8, 3, 0, 0, 394, 395, 5, 34, 0,
		0, 395, 397, 5, 34, 0, 0, 396, 393, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 397,
		400, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 401,
		1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 401, 402, 5, 34, 0, 0, 402, 122, 1, 0,
		0, 0, 403, 404, 5, 34, 0, 0, 404, 405, 8, 4, 0, 0, 405, 406, 5, 34, 0,
		0, 406, 124, 1, 0, 0, 0, 407, 409, 7, 5, 0, 0, 408, 407, 1, 0, 0, 0, 409,
		410, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 412,
		1, 0, 0, 0, 412, 413, 6, 62, 0, 0, 413, 126, 1, 0, 0, 0, 8, 0, 372, 378,
		384, 390, 396, 398, 410, 1, 6, 0, 0,
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

// ControlLexerInit initializes any static state used to implement ControlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewControlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ControlLexerInit() {
	staticData := &ControlLexerLexerStaticData
	staticData.once.Do(controllexerLexerInit)
}

// NewControlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewControlLexer(input antlr.CharStream) *ControlLexer {
	ControlLexerInit()
	l := new(ControlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ControlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Control.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ControlLexer tokens.
const (
	ControlLexerT__0   = 1
	ControlLexerT__1   = 2
	ControlLexerT__2   = 3
	ControlLexerT__3   = 4
	ControlLexerT__4   = 5
	ControlLexerT__5   = 6
	ControlLexerT__6   = 7
	ControlLexerT__7   = 8
	ControlLexerT__8   = 9
	ControlLexerT__9   = 10
	ControlLexerT__10  = 11
	ControlLexerT__11  = 12
	ControlLexerT__12  = 13
	ControlLexerT__13  = 14
	ControlLexerT__14  = 15
	ControlLexerT__15  = 16
	ControlLexerT__16  = 17
	ControlLexerT__17  = 18
	ControlLexerT__18  = 19
	ControlLexerT__19  = 20
	ControlLexerT__20  = 21
	ControlLexerT__21  = 22
	ControlLexerT__22  = 23
	ControlLexerT__23  = 24
	ControlLexerT__24  = 25
	ControlLexerT__25  = 26
	ControlLexerT__26  = 27
	ControlLexerT__27  = 28
	ControlLexerT__28  = 29
	ControlLexerT__29  = 30
	ControlLexerT__30  = 31
	ControlLexerT__31  = 32
	ControlLexerT__32  = 33
	ControlLexerT__33  = 34
	ControlLexerT__34  = 35
	ControlLexerT__35  = 36
	ControlLexerT__36  = 37
	ControlLexerT__37  = 38
	ControlLexerT__38  = 39
	ControlLexerT__39  = 40
	ControlLexerT__40  = 41
	ControlLexerT__41  = 42
	ControlLexerT__42  = 43
	ControlLexerT__43  = 44
	ControlLexerT__44  = 45
	ControlLexerT__45  = 46
	ControlLexerT__46  = 47
	ControlLexerT__47  = 48
	ControlLexerT__48  = 49
	ControlLexerT__49  = 50
	ControlLexerT__50  = 51
	ControlLexerT__51  = 52
	ControlLexerT__52  = 53
	ControlLexerT__53  = 54
	ControlLexerT__54  = 55
	ControlLexerT__55  = 56
	ControlLexerT__56  = 57
	ControlLexerINT    = 58
	ControlLexerID     = 59
	ControlLexerFLOAT  = 60
	ControlLexerSTRING = 61
	ControlLexerCHAR   = 62
	ControlLexerWS     = 63
)