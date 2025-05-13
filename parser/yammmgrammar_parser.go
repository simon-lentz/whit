// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // YammmGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type YammmGrammarParser struct {
	*antlr.BaseParser
}

var yammmgrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func yammmgrammarParserInit() {
	staticData := &yammmgrammarParserStaticData
	staticData.literalNames = []string{
		"", "'schema'", "'abstract'", "'part'", "'type'", "'extends'", "'primary'",
		"'required'", "'one'", "'many'", "'Integer'", "'Float'", "'Boolean'",
		"'String'", "'Enum'", "'Pattern'", "'Timestamp'", "'Spacevector'", "'Date'",
		"'UUID'", "'in'", "'datatype'", "'includes'", "'{'", "'}'", "'['", "']'",
		"'('", "')'", "':'", "','", "'='", "'-->'", "'*->'", "'->'", "'/'",
		"'_'", "'*'", "'@'", "'!'", "'+'", "'-'", "'||'", "'&&'", "'=='", "'!='",
		"'=~'", "'!~'", "'?'", "'>'", "'>='", "'<'", "'<='", "'$'", "'|'", "'.'",
		"'%'", "'^'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "LPAR",
		"RPAR", "COLON", "COMMA", "EQUALS", "ASSOC", "COMP", "ARROW", "SLASH",
		"USCORE", "STAR", "AT", "EXCLAMATION", "PLUS", "MINUS", "OR", "AND",
		"EQUAL", "NOTEQUAL", "MATCH", "NOTMATCH", "QMARK", "GT", "GTE", "LT",
		"LTE", "DOLLAR", "PIPE", "PERIOD", "PERCENT", "HAT", "STRING", "DOC_COMMENT",
		"SL_COMMENT", "REGEXP", "WS", "VARIABLE", "INTEGER", "FLOAT", "BOOLEAN",
		"UC_WORD", "LC_WORD", "ANY_OTHER",
	}
	staticData.ruleNames = []string{
		"schema", "schema_name", "type", "datatype", "type_name", "plural_name",
		"extends_types", "type_body", "property", "rel_property", "property_name",
		"data_type_ref", "alias", "association", "composition", "any_name",
		"multiplicity", "relation_body", "built_in", "integerT", "floatT", "boolT",
		"stringT", "enumT", "patternT", "timestampT", "spacevectorT", "dateT",
		"uuidT", "datatypeKeyword", "invariant", "expr", "arguments", "parameters",
		"literal", "lc_keyword",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 69, 448, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 1, 0,
		1, 0, 5, 0, 76, 8, 0, 10, 0, 12, 0, 79, 9, 0, 1, 0, 1, 0, 1, 1, 3, 1, 84,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 2, 90, 8, 2, 1, 2, 1, 2, 3, 2, 94, 8,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 99, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 3, 3,
		106, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 5, 6, 121, 8, 6, 10, 6, 12, 6, 124, 9, 6, 1, 6, 3, 6,
		127, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 133, 8, 7, 10, 7, 12, 7, 136,
		9, 7, 1, 8, 3, 8, 139, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 145, 8, 8, 1,
		9, 3, 9, 148, 8, 9, 1, 9, 1, 9, 1, 9, 3, 9, 153, 8, 9, 1, 10, 1, 10, 3,
		10, 157, 8, 10, 1, 11, 1, 11, 3, 11, 161, 8, 11, 1, 12, 1, 12, 1, 13, 3,
		13, 166, 8, 13, 1, 13, 1, 13, 1, 13, 3, 13, 171, 8, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 3, 13, 177, 8, 13, 3, 13, 179, 8, 13, 1, 13, 1, 13, 3, 13, 183,
		8, 13, 1, 13, 3, 13, 186, 8, 13, 1, 14, 3, 14, 189, 8, 14, 1, 14, 1, 14,
		1, 14, 3, 14, 194, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 200, 8, 14,
		3, 14, 202, 8, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 210,
		8, 16, 1, 16, 1, 16, 1, 16, 3, 16, 215, 8, 16, 1, 16, 3, 16, 218, 8, 16,
		1, 16, 1, 16, 1, 17, 4, 17, 223, 8, 17, 11, 17, 12, 17, 224, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 237,
		8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 245, 8, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 253, 8, 20, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 263, 8, 22, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 4, 23, 270, 8, 23, 11, 23, 12, 23, 271, 1, 23,
		3, 23, 275, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3,
		24, 284, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 292, 8,
		25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 5, 31, 315, 8, 31, 10, 31, 12, 31, 318, 9, 31, 1, 31, 3, 31, 321, 8,
		31, 3, 31, 323, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 339, 8, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		5, 31, 373, 8, 31, 10, 31, 12, 31, 376, 9, 31, 1, 31, 3, 31, 379, 8, 31,
		3, 31, 381, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 388, 8, 31,
		1, 31, 3, 31, 391, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 397, 8, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 405, 8, 31, 1, 31, 1,
		31, 5, 31, 409, 8, 31, 10, 31, 12, 31, 412, 9, 31, 1, 32, 1, 32, 1, 32,
		1, 32, 5, 32, 418, 8, 32, 10, 32, 12, 32, 421, 9, 32, 3, 32, 423, 8, 32,
		1, 32, 3, 32, 426, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 5,
		33, 434, 8, 33, 10, 33, 12, 33, 437, 9, 33, 1, 33, 3, 33, 440, 8, 33, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 0, 1, 62, 36, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 0, 13, 1, 0, 67,
		68, 1, 0, 8, 9, 2, 0, 36, 36, 64, 64, 2, 0, 36, 36, 64, 65, 1, 0, 10, 19,
		3, 0, 35, 35, 37, 37, 56, 56, 1, 0, 40, 41, 1, 0, 49, 52, 1, 0, 46, 47,
		1, 0, 44, 45, 2, 0, 42, 42, 57, 57, 3, 0, 58, 58, 61, 61, 64, 66, 3, 0,
		1, 2, 4, 9, 21, 22, 499, 0, 72, 1, 0, 0, 0, 2, 83, 1, 0, 0, 0, 4, 89, 1,
		0, 0, 0, 6, 105, 1, 0, 0, 0, 8, 112, 1, 0, 0, 0, 10, 114, 1, 0, 0, 0, 12,
		116, 1, 0, 0, 0, 14, 134, 1, 0, 0, 0, 16, 138, 1, 0, 0, 0, 18, 147, 1,
		0, 0, 0, 20, 156, 1, 0, 0, 0, 22, 160, 1, 0, 0, 0, 24, 162, 1, 0, 0, 0,
		26, 165, 1, 0, 0, 0, 28, 188, 1, 0, 0, 0, 30, 203, 1, 0, 0, 0, 32, 205,
		1, 0, 0, 0, 34, 222, 1, 0, 0, 0, 36, 236, 1, 0, 0, 0, 38, 238, 1, 0, 0,
		0, 40, 246, 1, 0, 0, 0, 42, 254, 1, 0, 0, 0, 44, 256, 1, 0, 0, 0, 46, 264,
		1, 0, 0, 0, 48, 278, 1, 0, 0, 0, 50, 287, 1, 0, 0, 0, 52, 293, 1, 0, 0,
		0, 54, 298, 1, 0, 0, 0, 56, 300, 1, 0, 0, 0, 58, 302, 1, 0, 0, 0, 60, 304,
		1, 0, 0, 0, 62, 338, 1, 0, 0, 0, 64, 413, 1, 0, 0, 0, 66, 429, 1, 0, 0,
		0, 68, 443, 1, 0, 0, 0, 70, 445, 1, 0, 0, 0, 72, 77, 3, 2, 1, 0, 73, 76,
		3, 4, 2, 0, 74, 76, 3, 6, 3, 0, 75, 73, 1, 0, 0, 0, 75, 74, 1, 0, 0, 0,
		76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 80, 1,
		0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 5, 0, 0, 1, 81, 1, 1, 0, 0, 0, 82,
		84, 5, 59, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 1, 0,
		0, 0, 85, 86, 5, 1, 0, 0, 86, 87, 5, 58, 0, 0, 87, 3, 1, 0, 0, 0, 88, 90,
		5, 59, 0, 0, 89, 88, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0,
		91, 94, 5, 2, 0, 0, 92, 94, 5, 3, 0, 0, 93, 91, 1, 0, 0, 0, 93, 92, 1,
		0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 5, 4, 0, 0, 96,
		98, 3, 8, 4, 0, 97, 99, 3, 12, 6, 0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0,
		0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 5, 23, 0, 0, 101, 102, 3, 14, 7, 0,
		102, 103, 5, 24, 0, 0, 103, 5, 1, 0, 0, 0, 104, 106, 5, 59, 0, 0, 105,
		104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108,
		5, 4, 0, 0, 108, 109, 3, 8, 4, 0, 109, 110, 5, 31, 0, 0, 110, 111, 3, 36,
		18, 0, 111, 7, 1, 0, 0, 0, 112, 113, 5, 67, 0, 0, 113, 9, 1, 0, 0, 0, 114,
		115, 3, 8, 4, 0, 115, 11, 1, 0, 0, 0, 116, 117, 5, 5, 0, 0, 117, 122, 3,
		8, 4, 0, 118, 119, 5, 30, 0, 0, 119, 121, 3, 8, 4, 0, 120, 118, 1, 0, 0,
		0, 121, 124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123,
		126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 127, 5, 30, 0, 0, 126, 125,
		1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 13, 1, 0, 0, 0, 128, 133, 3, 16,
		8, 0, 129, 133, 3, 26, 13, 0, 130, 133, 3, 28, 14, 0, 131, 133, 3, 60,
		30, 0, 132, 128, 1, 0, 0, 0, 132, 129, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0,
		132, 131, 1, 0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134,
		135, 1, 0, 0, 0, 135, 15, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 137, 139, 5,
		59, 0, 0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 1, 0, 0,
		0, 140, 141, 3, 20, 10, 0, 141, 144, 3, 22, 11, 0, 142, 145, 5, 6, 0, 0,
		143, 145, 5, 7, 0, 0, 144, 142, 1, 0, 0, 0, 144, 143, 1, 0, 0, 0, 144,
		145, 1, 0, 0, 0, 145, 17, 1, 0, 0, 0, 146, 148, 5, 59, 0, 0, 147, 146,
		1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 3, 20,
		10, 0, 150, 152, 3, 22, 11, 0, 151, 153, 5, 7, 0, 0, 152, 151, 1, 0, 0,
		0, 152, 153, 1, 0, 0, 0, 153, 19, 1, 0, 0, 0, 154, 157, 5, 68, 0, 0, 155,
		157, 3, 70, 35, 0, 156, 154, 1, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157, 21,
		1, 0, 0, 0, 158, 161, 3, 36, 18, 0, 159, 161, 3, 24, 12, 0, 160, 158, 1,
		0, 0, 0, 160, 159, 1, 0, 0, 0, 161, 23, 1, 0, 0, 0, 162, 163, 5, 67, 0,
		0, 163, 25, 1, 0, 0, 0, 164, 166, 5, 59, 0, 0, 165, 164, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 5, 32, 0, 0, 168, 170,
		3, 30, 15, 0, 169, 171, 3, 32, 16, 0, 170, 169, 1, 0, 0, 0, 170, 171, 1,
		0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 178, 3, 8, 4, 0, 173, 174, 5, 35, 0,
		0, 174, 176, 3, 30, 15, 0, 175, 177, 3, 32, 16, 0, 176, 175, 1, 0, 0, 0,
		176, 177, 1, 0, 0, 0, 177, 179, 1, 0, 0, 0, 178, 173, 1, 0, 0, 0, 178,
		179, 1, 0, 0, 0, 179, 185, 1, 0, 0, 0, 180, 182, 5, 23, 0, 0, 181, 183,
		3, 34, 17, 0, 182, 181, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 1,
		0, 0, 0, 184, 186, 5, 24, 0, 0, 185, 180, 1, 0, 0, 0, 185, 186, 1, 0, 0,
		0, 186, 27, 1, 0, 0, 0, 187, 189, 5, 59, 0, 0, 188, 187, 1, 0, 0, 0, 188,
		189, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 5, 33, 0, 0, 191, 193,
		3, 30, 15, 0, 192, 194, 3, 32, 16, 0, 193, 192, 1, 0, 0, 0, 193, 194, 1,
		0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 201, 3, 8, 4, 0, 196, 197, 5, 35, 0,
		0, 197, 199, 3, 30, 15, 0, 198, 200, 3, 32, 16, 0, 199, 198, 1, 0, 0, 0,
		199, 200, 1, 0, 0, 0, 200, 202, 1, 0, 0, 0, 201, 196, 1, 0, 0, 0, 201,
		202, 1, 0, 0, 0, 202, 29, 1, 0, 0, 0, 203, 204, 7, 0, 0, 0, 204, 31, 1,
		0, 0, 0, 205, 217, 5, 27, 0, 0, 206, 209, 5, 36, 0, 0, 207, 208, 5, 29,
		0, 0, 208, 210, 7, 1, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0,
		210, 218, 1, 0, 0, 0, 211, 214, 5, 8, 0, 0, 212, 213, 5, 29, 0, 0, 213,
		215, 7, 1, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 218,
		1, 0, 0, 0, 216, 218, 5, 9, 0, 0, 217, 206, 1, 0, 0, 0, 217, 211, 1, 0,
		0, 0, 217, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220, 5, 28, 0, 0,
		220, 33, 1, 0, 0, 0, 221, 223, 3, 18, 9, 0, 222, 221, 1, 0, 0, 0, 223,
		224, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 35, 1,
		0, 0, 0, 226, 237, 3, 38, 19, 0, 227, 237, 3, 40, 20, 0, 228, 237, 3, 42,
		21, 0, 229, 237, 3, 44, 22, 0, 230, 237, 3, 46, 23, 0, 231, 237, 3, 48,
		24, 0, 232, 237, 3, 50, 25, 0, 233, 237, 3, 54, 27, 0, 234, 237, 3, 56,
		28, 0, 235, 237, 3, 52, 26, 0, 236, 226, 1, 0, 0, 0, 236, 227, 1, 0, 0,
		0, 236, 228, 1, 0, 0, 0, 236, 229, 1, 0, 0, 0, 236, 230, 1, 0, 0, 0, 236,
		231, 1, 0, 0, 0, 236, 232, 1, 0, 0, 0, 236, 233, 1, 0, 0, 0, 236, 234,
		1, 0, 0, 0, 236, 235, 1, 0, 0, 0, 237, 37, 1, 0, 0, 0, 238, 244, 5, 10,
		0, 0, 239, 240, 5, 25, 0, 0, 240, 241, 7, 2, 0, 0, 241, 242, 5, 30, 0,
		0, 242, 243, 7, 2, 0, 0, 243, 245, 5, 26, 0, 0, 244, 239, 1, 0, 0, 0, 244,
		245, 1, 0, 0, 0, 245, 39, 1, 0, 0, 0, 246, 252, 5, 11, 0, 0, 247, 248,
		5, 25, 0, 0, 248, 249, 7, 3, 0, 0, 249, 250, 5, 30, 0, 0, 250, 251, 7,
		3, 0, 0, 251, 253, 5, 26, 0, 0, 252, 247, 1, 0, 0, 0, 252, 253, 1, 0, 0,
		0, 253, 41, 1, 0, 0, 0, 254, 255, 5, 12, 0, 0, 255, 43, 1, 0, 0, 0, 256,
		262, 5, 13, 0, 0, 257, 258, 5, 25, 0, 0, 258, 259, 7, 2, 0, 0, 259, 260,
		5, 30, 0, 0, 260, 261, 7, 2, 0, 0, 261, 263, 5, 26, 0, 0, 262, 257, 1,
		0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 45, 1, 0, 0, 0, 264, 265, 5, 14, 0,
		0, 265, 266, 5, 25, 0, 0, 266, 269, 5, 58, 0, 0, 267, 268, 5, 30, 0, 0,
		268, 270, 5, 58, 0, 0, 269, 267, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271,
		269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 274, 1, 0, 0, 0, 273, 275,
		5, 30, 0, 0, 274, 273, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 276, 1, 0,
		0, 0, 276, 277, 5, 26, 0, 0, 277, 47, 1, 0, 0, 0, 278, 279, 5, 15, 0, 0,
		279, 280, 5, 25, 0, 0, 280, 283, 5, 58, 0, 0, 281, 282, 5, 30, 0, 0, 282,
		284, 5, 58, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 285,
		1, 0, 0, 0, 285, 286, 5, 26, 0, 0, 286, 49, 1, 0, 0, 0, 287, 291, 5, 16,
		0, 0, 288, 289, 5, 25, 0, 0, 289, 290, 5, 58, 0, 0, 290, 292, 5, 26, 0,
		0, 291, 288, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 51, 1, 0, 0, 0, 293,
		294, 5, 17, 0, 0, 294, 295, 5, 25, 0, 0, 295, 296, 5, 64, 0, 0, 296, 297,
		5, 26, 0, 0, 297, 53, 1, 0, 0, 0, 298, 299, 5, 18, 0, 0, 299, 55, 1, 0,
		0, 0, 300, 301, 5, 19, 0, 0, 301, 57, 1, 0, 0, 0, 302, 303, 7, 4, 0, 0,
		303, 59, 1, 0, 0, 0, 304, 305, 5, 39, 0, 0, 305, 306, 5, 58, 0, 0, 306,
		307, 3, 62, 31, 0, 307, 61, 1, 0, 0, 0, 308, 309, 6, 31, -1, 0, 309, 339,
		3, 68, 34, 0, 310, 322, 5, 25, 0, 0, 311, 316, 3, 62, 31, 0, 312, 313,
		5, 30, 0, 0, 313, 315, 3, 62, 31, 0, 314, 312, 1, 0, 0, 0, 315, 318, 1,
		0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 320, 1, 0, 0,
		0, 318, 316, 1, 0, 0, 0, 319, 321, 5, 30, 0, 0, 320, 319, 1, 0, 0, 0, 320,
		321, 1, 0, 0, 0, 321, 323, 1, 0, 0, 0, 322, 311, 1, 0, 0, 0, 322, 323,
		1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 339, 5, 26, 0, 0, 325, 326, 5, 41,
		0, 0, 326, 339, 3, 62, 31, 20, 327, 328, 5, 39, 0, 0, 328, 339, 3, 62,
		31, 16, 329, 330, 5, 27, 0, 0, 330, 331, 3, 62, 31, 0, 331, 332, 5, 28,
		0, 0, 332, 339, 1, 0, 0, 0, 333, 339, 5, 63, 0, 0, 334, 339, 3, 20, 10,
		0, 335, 339, 3, 58, 29, 0, 336, 339, 5, 67, 0, 0, 337, 339, 5, 36, 0, 0,
		338, 308, 1, 0, 0, 0, 338, 310, 1, 0, 0, 0, 338, 325, 1, 0, 0, 0, 338,
		327, 1, 0, 0, 0, 338, 329, 1, 0, 0, 0, 338, 333, 1, 0, 0, 0, 338, 334,
		1, 0, 0, 0, 338, 335, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 338, 337, 1, 0,
		0, 0, 339, 410, 1, 0, 0, 0, 340, 341, 10, 17, 0, 0, 341, 342, 5, 55, 0,
		0, 342, 409, 3, 62, 31, 18, 343, 344, 10, 15, 0, 0, 344, 345, 7, 5, 0,
		0, 345, 409, 3, 62, 31, 16, 346, 347, 10, 14, 0, 0, 347, 348, 7, 6, 0,
		0, 348, 409, 3, 62, 31, 15, 349, 350, 10, 13, 0, 0, 350, 351, 7, 7, 0,
		0, 351, 409, 3, 62, 31, 14, 352, 353, 10, 12, 0, 0, 353, 354, 5, 20, 0,
		0, 354, 409, 3, 62, 31, 13, 355, 356, 10, 11, 0, 0, 356, 357, 7, 8, 0,
		0, 357, 409, 3, 62, 31, 12, 358, 359, 10, 10, 0, 0, 359, 360, 7, 9, 0,
		0, 360, 409, 3, 62, 31, 11, 361, 362, 10, 9, 0, 0, 362, 363, 5, 43, 0,
		0, 363, 409, 3, 62, 31, 10, 364, 365, 10, 8, 0, 0, 365, 366, 7, 10, 0,
		0, 366, 409, 3, 62, 31, 9, 367, 368, 10, 19, 0, 0, 368, 380, 5, 25, 0,
		0, 369, 374, 3, 62, 31, 0, 370, 371, 5, 30, 0, 0, 371, 373, 3, 62, 31,
		0, 372, 370, 1, 0, 0, 0, 373, 376, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 374,
		375, 1, 0, 0, 0, 375, 378, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 377, 379,
		5, 30, 0, 0, 378, 377, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 381, 1, 0,
		0, 0, 380, 369, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0,
		382, 409, 5, 26, 0, 0, 383, 384, 10, 18, 0, 0, 384, 385, 5, 34, 0, 0, 385,
		387, 7, 0, 0, 0, 386, 388, 3, 64, 32, 0, 387, 386, 1, 0, 0, 0, 387, 388,
		1, 0, 0, 0, 388, 390, 1, 0, 0, 0, 389, 391, 3, 66, 33, 0, 390, 389, 1,
		0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 396, 1, 0, 0, 0, 392, 393, 5, 23, 0,
		0, 393, 394, 3, 62, 31, 0, 394, 395, 5, 24, 0, 0, 395, 397, 1, 0, 0, 0,
		396, 392, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 409, 1, 0, 0, 0, 398,
		399, 10, 7, 0, 0, 399, 400, 5, 48, 0, 0, 400, 401, 5, 23, 0, 0, 401, 404,
		3, 62, 31, 0, 402, 403, 5, 29, 0, 0, 403, 405, 3, 62, 31, 0, 404, 402,
		1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 407, 5, 24,
		0, 0, 407, 409, 1, 0, 0, 0, 408, 340, 1, 0, 0, 0, 408, 343, 1, 0, 0, 0,
		408, 346, 1, 0, 0, 0, 408, 349, 1, 0, 0, 0, 408, 352, 1, 0, 0, 0, 408,
		355, 1, 0, 0, 0, 408, 358, 1, 0, 0, 0, 408, 361, 1, 0, 0, 0, 408, 364,
		1, 0, 0, 0, 408, 367, 1, 0, 0, 0, 408, 383, 1, 0, 0, 0, 408, 398, 1, 0,
		0, 0, 409, 412, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0,
		411, 63, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 413, 422, 5, 27, 0, 0, 414,
		419, 3, 62, 31, 0, 415, 416, 5, 30, 0, 0, 416, 418, 3, 62, 31, 0, 417,
		415, 1, 0, 0, 0, 418, 421, 1, 0, 0, 0, 419, 417, 1, 0, 0, 0, 419, 420,
		1, 0, 0, 0, 420, 423, 1, 0, 0, 0, 421, 419, 1, 0, 0, 0, 422, 414, 1, 0,
		0, 0, 422, 423, 1, 0, 0, 0, 423, 425, 1, 0, 0, 0, 424, 426, 5, 30, 0, 0,
		425, 424, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427,
		428, 5, 28, 0, 0, 428, 65, 1, 0, 0, 0, 429, 430, 5, 54, 0, 0, 430, 435,
		5, 63, 0, 0, 431, 432, 5, 30, 0, 0, 432, 434, 5, 63, 0, 0, 433, 431, 1,
		0, 0, 0, 434, 437, 1, 0, 0, 0, 435, 433, 1, 0, 0, 0, 435, 436, 1, 0, 0,
		0, 436, 439, 1, 0, 0, 0, 437, 435, 1, 0, 0, 0, 438, 440, 5, 30, 0, 0, 439,
		438, 1, 0, 0, 0, 439, 440, 1, 0, 0, 0, 440, 441, 1, 0, 0, 0, 441, 442,
		5, 54, 0, 0, 442, 67, 1, 0, 0, 0, 443, 444, 7, 11, 0, 0, 444, 69, 1, 0,
		0, 0, 445, 446, 7, 12, 0, 0, 446, 71, 1, 0, 0, 0, 57, 75, 77, 83, 89, 93,
		98, 105, 122, 126, 132, 134, 138, 144, 147, 152, 156, 160, 165, 170, 176,
		178, 182, 185, 188, 193, 199, 201, 209, 214, 217, 224, 236, 244, 252, 262,
		271, 274, 283, 291, 316, 320, 322, 338, 374, 378, 380, 387, 390, 396, 404,
		408, 410, 419, 422, 425, 435, 439,
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

// YammmGrammarParserInit initializes any static state used to implement YammmGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewYammmGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func YammmGrammarParserInit() {
	staticData := &yammmgrammarParserStaticData
	staticData.once.Do(yammmgrammarParserInit)
}

// NewYammmGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewYammmGrammarParser(input antlr.TokenStream) *YammmGrammarParser {
	YammmGrammarParserInit()
	this := new(YammmGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &yammmgrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// YammmGrammarParser tokens.
const (
	YammmGrammarParserEOF         = antlr.TokenEOF
	YammmGrammarParserT__0        = 1
	YammmGrammarParserT__1        = 2
	YammmGrammarParserT__2        = 3
	YammmGrammarParserT__3        = 4
	YammmGrammarParserT__4        = 5
	YammmGrammarParserT__5        = 6
	YammmGrammarParserT__6        = 7
	YammmGrammarParserT__7        = 8
	YammmGrammarParserT__8        = 9
	YammmGrammarParserT__9        = 10
	YammmGrammarParserT__10       = 11
	YammmGrammarParserT__11       = 12
	YammmGrammarParserT__12       = 13
	YammmGrammarParserT__13       = 14
	YammmGrammarParserT__14       = 15
	YammmGrammarParserT__15       = 16
	YammmGrammarParserT__16       = 17
	YammmGrammarParserT__17       = 18
	YammmGrammarParserT__18       = 19
	YammmGrammarParserT__19       = 20
	YammmGrammarParserT__20       = 21
	YammmGrammarParserT__21       = 22
	YammmGrammarParserLBRACE      = 23
	YammmGrammarParserRBRACE      = 24
	YammmGrammarParserLBRACK      = 25
	YammmGrammarParserRBRACK      = 26
	YammmGrammarParserLPAR        = 27
	YammmGrammarParserRPAR        = 28
	YammmGrammarParserCOLON       = 29
	YammmGrammarParserCOMMA       = 30
	YammmGrammarParserEQUALS      = 31
	YammmGrammarParserASSOC       = 32
	YammmGrammarParserCOMP        = 33
	YammmGrammarParserARROW       = 34
	YammmGrammarParserSLASH       = 35
	YammmGrammarParserUSCORE      = 36
	YammmGrammarParserSTAR        = 37
	YammmGrammarParserAT          = 38
	YammmGrammarParserEXCLAMATION = 39
	YammmGrammarParserPLUS        = 40
	YammmGrammarParserMINUS       = 41
	YammmGrammarParserOR          = 42
	YammmGrammarParserAND         = 43
	YammmGrammarParserEQUAL       = 44
	YammmGrammarParserNOTEQUAL    = 45
	YammmGrammarParserMATCH       = 46
	YammmGrammarParserNOTMATCH    = 47
	YammmGrammarParserQMARK       = 48
	YammmGrammarParserGT          = 49
	YammmGrammarParserGTE         = 50
	YammmGrammarParserLT          = 51
	YammmGrammarParserLTE         = 52
	YammmGrammarParserDOLLAR      = 53
	YammmGrammarParserPIPE        = 54
	YammmGrammarParserPERIOD      = 55
	YammmGrammarParserPERCENT     = 56
	YammmGrammarParserHAT         = 57
	YammmGrammarParserSTRING      = 58
	YammmGrammarParserDOC_COMMENT = 59
	YammmGrammarParserSL_COMMENT  = 60
	YammmGrammarParserREGEXP      = 61
	YammmGrammarParserWS          = 62
	YammmGrammarParserVARIABLE    = 63
	YammmGrammarParserINTEGER     = 64
	YammmGrammarParserFLOAT       = 65
	YammmGrammarParserBOOLEAN     = 66
	YammmGrammarParserUC_WORD     = 67
	YammmGrammarParserLC_WORD     = 68
	YammmGrammarParserANY_OTHER   = 69
)

// YammmGrammarParser rules.
const (
	YammmGrammarParserRULE_schema          = 0
	YammmGrammarParserRULE_schema_name     = 1
	YammmGrammarParserRULE_type            = 2
	YammmGrammarParserRULE_datatype        = 3
	YammmGrammarParserRULE_type_name       = 4
	YammmGrammarParserRULE_plural_name     = 5
	YammmGrammarParserRULE_extends_types   = 6
	YammmGrammarParserRULE_type_body       = 7
	YammmGrammarParserRULE_property        = 8
	YammmGrammarParserRULE_rel_property    = 9
	YammmGrammarParserRULE_property_name   = 10
	YammmGrammarParserRULE_data_type_ref   = 11
	YammmGrammarParserRULE_alias           = 12
	YammmGrammarParserRULE_association     = 13
	YammmGrammarParserRULE_composition     = 14
	YammmGrammarParserRULE_any_name        = 15
	YammmGrammarParserRULE_multiplicity    = 16
	YammmGrammarParserRULE_relation_body   = 17
	YammmGrammarParserRULE_built_in        = 18
	YammmGrammarParserRULE_integerT        = 19
	YammmGrammarParserRULE_floatT          = 20
	YammmGrammarParserRULE_boolT           = 21
	YammmGrammarParserRULE_stringT         = 22
	YammmGrammarParserRULE_enumT           = 23
	YammmGrammarParserRULE_patternT        = 24
	YammmGrammarParserRULE_timestampT      = 25
	YammmGrammarParserRULE_spacevectorT    = 26
	YammmGrammarParserRULE_dateT           = 27
	YammmGrammarParserRULE_uuidT           = 28
	YammmGrammarParserRULE_datatypeKeyword = 29
	YammmGrammarParserRULE_invariant       = 30
	YammmGrammarParserRULE_expr            = 31
	YammmGrammarParserRULE_arguments       = 32
	YammmGrammarParserRULE_parameters      = 33
	YammmGrammarParserRULE_literal         = 34
	YammmGrammarParserRULE_lc_keyword      = 35
)

// ISchemaContext is an interface to support dynamic dispatch.
type ISchemaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaContext differentiates from other interfaces.
	IsSchemaContext()
}

type SchemaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaContext() *SchemaContext {
	var p = new(SchemaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_schema
	return p
}

func (*SchemaContext) IsSchemaContext() {}

func NewSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaContext {
	var p = new(SchemaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_schema

	return p
}

func (s *SchemaContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaContext) Schema_name() ISchema_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchema_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchema_nameContext)
}

func (s *SchemaContext) EOF() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserEOF, 0)
}

func (s *SchemaContext) AllType_() []ITypeContext {
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

func (s *SchemaContext) Type_(i int) ITypeContext {
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

func (s *SchemaContext) AllDatatype() []IDatatypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatatypeContext); ok {
			len++
		}
	}

	tst := make([]IDatatypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatatypeContext); ok {
			tst[i] = t.(IDatatypeContext)
			i++
		}
	}

	return tst
}

func (s *SchemaContext) Datatype(i int) IDatatypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatatypeContext); ok {
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

	return t.(IDatatypeContext)
}

func (s *SchemaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterSchema(s)
	}
}

func (s *SchemaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitSchema(s)
	}
}

func (s *SchemaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitSchema(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Schema() (localctx ISchemaContext) {
	this := p
	_ = this

	localctx = NewSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, YammmGrammarParserRULE_schema)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Schema_name()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&576460752303423516) != 0 {
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(73)
				p.Type_()
			}

		case 2:
			{
				p.SetState(74)
				p.Datatype()
			}

		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(YammmGrammarParserEOF)
	}

	return localctx
}

// ISchema_nameContext is an interface to support dynamic dispatch.
type ISchema_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchema_nameContext differentiates from other interfaces.
	IsSchema_nameContext()
}

type Schema_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchema_nameContext() *Schema_nameContext {
	var p = new(Schema_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_schema_name
	return p
}

func (*Schema_nameContext) IsSchema_nameContext() {}

func NewSchema_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Schema_nameContext {
	var p = new(Schema_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_schema_name

	return p
}

func (s *Schema_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Schema_nameContext) STRING() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserSTRING, 0)
}

func (s *Schema_nameContext) DOC_COMMENT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserDOC_COMMENT, 0)
}

func (s *Schema_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Schema_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Schema_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterSchema_name(s)
	}
}

func (s *Schema_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitSchema_name(s)
	}
}

func (s *Schema_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitSchema_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Schema_name() (localctx ISchema_nameContext) {
	this := p
	_ = this

	localctx = NewSchema_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, YammmGrammarParserRULE_schema_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserDOC_COMMENT {
		{
			p.SetState(82)
			p.Match(YammmGrammarParserDOC_COMMENT)
		}

	}
	{
		p.SetState(85)
		p.Match(YammmGrammarParserT__0)
	}
	{
		p.SetState(86)
		p.Match(YammmGrammarParserSTRING)
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIs_abstract returns the is_abstract token.
	GetIs_abstract() antlr.Token

	// GetIs_part returns the is_part token.
	GetIs_part() antlr.Token

	// SetIs_abstract sets the is_abstract token.
	SetIs_abstract(antlr.Token)

	// SetIs_part sets the is_part token.
	SetIs_part(antlr.Token)

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	is_abstract antlr.Token
	is_part     antlr.Token
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) GetIs_abstract() antlr.Token { return s.is_abstract }

func (s *TypeContext) GetIs_part() antlr.Token { return s.is_part }

func (s *TypeContext) SetIs_abstract(v antlr.Token) { s.is_abstract = v }

func (s *TypeContext) SetIs_part(v antlr.Token) { s.is_part = v }

func (s *TypeContext) Type_name() IType_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_nameContext)
}

func (s *TypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACE, 0)
}

func (s *TypeContext) Type_body() IType_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_bodyContext)
}

func (s *TypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACE, 0)
}

func (s *TypeContext) DOC_COMMENT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserDOC_COMMENT, 0)
}

func (s *TypeContext) Extends_types() IExtends_typesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtends_typesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtends_typesContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, YammmGrammarParserRULE_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserDOC_COMMENT {
		{
			p.SetState(88)
			p.Match(YammmGrammarParserDOC_COMMENT)
		}

	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case YammmGrammarParserT__1:
		{
			p.SetState(91)

			var _m = p.Match(YammmGrammarParserT__1)

			localctx.(*TypeContext).is_abstract = _m
		}

	case YammmGrammarParserT__2:
		{
			p.SetState(92)

			var _m = p.Match(YammmGrammarParserT__2)

			localctx.(*TypeContext).is_part = _m
		}

	case YammmGrammarParserT__3:

	default:
	}
	{
		p.SetState(95)
		p.Match(YammmGrammarParserT__3)
	}
	{
		p.SetState(96)
		p.Type_name()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserT__4 {
		{
			p.SetState(97)
			p.Extends_types()
		}

	}
	{
		p.SetState(100)
		p.Match(YammmGrammarParserLBRACE)
	}
	{
		p.SetState(101)
		p.Type_body()
	}
	{
		p.SetState(102)
		p.Match(YammmGrammarParserRBRACE)
	}

	return localctx
}

// IDatatypeContext is an interface to support dynamic dispatch.
type IDatatypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatatypeContext differentiates from other interfaces.
	IsDatatypeContext()
}

type DatatypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatatypeContext() *DatatypeContext {
	var p = new(DatatypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_datatype
	return p
}

func (*DatatypeContext) IsDatatypeContext() {}

func NewDatatypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatatypeContext {
	var p = new(DatatypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_datatype

	return p
}

func (s *DatatypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DatatypeContext) Type_name() IType_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_nameContext)
}

func (s *DatatypeContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserEQUALS, 0)
}

func (s *DatatypeContext) Built_in() IBuilt_inContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuilt_inContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuilt_inContext)
}

func (s *DatatypeContext) DOC_COMMENT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserDOC_COMMENT, 0)
}

func (s *DatatypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatatypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatatypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterDatatype(s)
	}
}

func (s *DatatypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitDatatype(s)
	}
}

func (s *DatatypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitDatatype(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Datatype() (localctx IDatatypeContext) {
	this := p
	_ = this

	localctx = NewDatatypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, YammmGrammarParserRULE_datatype)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserDOC_COMMENT {
		{
			p.SetState(104)
			p.Match(YammmGrammarParserDOC_COMMENT)
		}

	}
	{
		p.SetState(107)
		p.Match(YammmGrammarParserT__3)
	}
	{
		p.SetState(108)
		p.Type_name()
	}
	{
		p.SetState(109)
		p.Match(YammmGrammarParserEQUALS)
	}
	{
		p.SetState(110)
		p.Built_in()
	}

	return localctx
}

// IType_nameContext is an interface to support dynamic dispatch.
type IType_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_nameContext differentiates from other interfaces.
	IsType_nameContext()
}

type Type_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_nameContext() *Type_nameContext {
	var p = new(Type_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_type_name
	return p
}

func (*Type_nameContext) IsType_nameContext() {}

func NewType_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_nameContext {
	var p = new(Type_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_type_name

	return p
}

func (s *Type_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_nameContext) UC_WORD() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserUC_WORD, 0)
}

func (s *Type_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterType_name(s)
	}
}

func (s *Type_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitType_name(s)
	}
}

func (s *Type_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitType_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Type_name() (localctx IType_nameContext) {
	this := p
	_ = this

	localctx = NewType_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, YammmGrammarParserRULE_type_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(YammmGrammarParserUC_WORD)
	}

	return localctx
}

// IPlural_nameContext is an interface to support dynamic dispatch.
type IPlural_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlural_nameContext differentiates from other interfaces.
	IsPlural_nameContext()
}

type Plural_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlural_nameContext() *Plural_nameContext {
	var p = new(Plural_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_plural_name
	return p
}

func (*Plural_nameContext) IsPlural_nameContext() {}

func NewPlural_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Plural_nameContext {
	var p = new(Plural_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_plural_name

	return p
}

func (s *Plural_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Plural_nameContext) Type_name() IType_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_nameContext)
}

func (s *Plural_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Plural_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Plural_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterPlural_name(s)
	}
}

func (s *Plural_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitPlural_name(s)
	}
}

func (s *Plural_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitPlural_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Plural_name() (localctx IPlural_nameContext) {
	this := p
	_ = this

	localctx = NewPlural_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, YammmGrammarParserRULE_plural_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Type_name()
	}

	return localctx
}

// IExtends_typesContext is an interface to support dynamic dispatch.
type IExtends_typesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtends_typesContext differentiates from other interfaces.
	IsExtends_typesContext()
}

type Extends_typesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtends_typesContext() *Extends_typesContext {
	var p = new(Extends_typesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_extends_types
	return p
}

func (*Extends_typesContext) IsExtends_typesContext() {}

func NewExtends_typesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Extends_typesContext {
	var p = new(Extends_typesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_extends_types

	return p
}

func (s *Extends_typesContext) GetParser() antlr.Parser { return s.parser }

func (s *Extends_typesContext) AllType_name() []IType_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_nameContext); ok {
			len++
		}
	}

	tst := make([]IType_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_nameContext); ok {
			tst[i] = t.(IType_nameContext)
			i++
		}
	}

	return tst
}

func (s *Extends_typesContext) Type_name(i int) IType_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_nameContext); ok {
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

	return t.(IType_nameContext)
}

func (s *Extends_typesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserCOMMA)
}

func (s *Extends_typesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOMMA, i)
}

func (s *Extends_typesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Extends_typesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Extends_typesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterExtends_types(s)
	}
}

func (s *Extends_typesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitExtends_types(s)
	}
}

func (s *Extends_typesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitExtends_types(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Extends_types() (localctx IExtends_typesContext) {
	this := p
	_ = this

	localctx = NewExtends_typesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, YammmGrammarParserRULE_extends_types)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(YammmGrammarParserT__4)
	}
	{
		p.SetState(117)
		p.Type_name()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(118)
				p.Match(YammmGrammarParserCOMMA)
			}
			{
				p.SetState(119)
				p.Type_name()
			}

		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserCOMMA {
		{
			p.SetState(125)
			p.Match(YammmGrammarParserCOMMA)
		}

	}

	return localctx
}

// IType_bodyContext is an interface to support dynamic dispatch.
type IType_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_bodyContext differentiates from other interfaces.
	IsType_bodyContext()
}

type Type_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_bodyContext() *Type_bodyContext {
	var p = new(Type_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_type_body
	return p
}

func (*Type_bodyContext) IsType_bodyContext() {}

func NewType_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_bodyContext {
	var p = new(Type_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_type_body

	return p
}

func (s *Type_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_bodyContext) AllProperty() []IPropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropertyContext); ok {
			len++
		}
	}

	tst := make([]IPropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropertyContext); ok {
			tst[i] = t.(IPropertyContext)
			i++
		}
	}

	return tst
}

func (s *Type_bodyContext) Property(i int) IPropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyContext); ok {
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

	return t.(IPropertyContext)
}

func (s *Type_bodyContext) AllAssociation() []IAssociationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssociationContext); ok {
			len++
		}
	}

	tst := make([]IAssociationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssociationContext); ok {
			tst[i] = t.(IAssociationContext)
			i++
		}
	}

	return tst
}

func (s *Type_bodyContext) Association(i int) IAssociationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssociationContext); ok {
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

	return t.(IAssociationContext)
}

func (s *Type_bodyContext) AllComposition() []ICompositionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICompositionContext); ok {
			len++
		}
	}

	tst := make([]ICompositionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICompositionContext); ok {
			tst[i] = t.(ICompositionContext)
			i++
		}
	}

	return tst
}

func (s *Type_bodyContext) Composition(i int) ICompositionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompositionContext); ok {
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

	return t.(ICompositionContext)
}

func (s *Type_bodyContext) AllInvariant() []IInvariantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInvariantContext); ok {
			len++
		}
	}

	tst := make([]IInvariantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInvariantContext); ok {
			tst[i] = t.(IInvariantContext)
			i++
		}
	}

	return tst
}

func (s *Type_bodyContext) Invariant(i int) IInvariantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvariantContext); ok {
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

	return t.(IInvariantContext)
}

func (s *Type_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterType_body(s)
	}
}

func (s *Type_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitType_body(s)
	}
}

func (s *Type_bodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitType_body(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Type_body() (localctx IType_bodyContext) {
	this := p
	_ = this

	localctx = NewType_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, YammmGrammarParserRULE_type_body)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&576461314950431734) != 0 || _la == YammmGrammarParserLC_WORD {
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(128)
				p.Property()
			}

		case 2:
			{
				p.SetState(129)
				p.Association()
			}

		case 3:
			{
				p.SetState(130)
				p.Composition()
			}

		case 4:
			{
				p.SetState(131)
				p.Invariant()
			}

		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIs_primary returns the is_primary token.
	GetIs_primary() antlr.Token

	// GetIs_required returns the is_required token.
	GetIs_required() antlr.Token

	// SetIs_primary sets the is_primary token.
	SetIs_primary(antlr.Token)

	// SetIs_required sets the is_required token.
	SetIs_required(antlr.Token)

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	is_primary  antlr.Token
	is_required antlr.Token
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) GetIs_primary() antlr.Token { return s.is_primary }

func (s *PropertyContext) GetIs_required() antlr.Token { return s.is_required }

func (s *PropertyContext) SetIs_primary(v antlr.Token) { s.is_primary = v }

func (s *PropertyContext) SetIs_required(v antlr.Token) { s.is_required = v }

func (s *PropertyContext) Property_name() IProperty_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProperty_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProperty_nameContext)
}

func (s *PropertyContext) Data_type_ref() IData_type_refContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_type_refContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_type_refContext)
}

func (s *PropertyContext) DOC_COMMENT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserDOC_COMMENT, 0)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (s *PropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Property() (localctx IPropertyContext) {
	this := p
	_ = this

	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, YammmGrammarParserRULE_property)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserDOC_COMMENT {
		{
			p.SetState(137)
			p.Match(YammmGrammarParserDOC_COMMENT)
		}

	}
	{
		p.SetState(140)
		p.Property_name()
	}
	{
		p.SetState(141)
		p.Data_type_ref()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(142)

			var _m = p.Match(YammmGrammarParserT__5)

			localctx.(*PropertyContext).is_primary = _m
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(143)

			var _m = p.Match(YammmGrammarParserT__6)

			localctx.(*PropertyContext).is_required = _m
		}

	}

	return localctx
}

// IRel_propertyContext is an interface to support dynamic dispatch.
type IRel_propertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIs_required returns the is_required token.
	GetIs_required() antlr.Token

	// SetIs_required sets the is_required token.
	SetIs_required(antlr.Token)

	// IsRel_propertyContext differentiates from other interfaces.
	IsRel_propertyContext()
}

type Rel_propertyContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	is_required antlr.Token
}

func NewEmptyRel_propertyContext() *Rel_propertyContext {
	var p = new(Rel_propertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_rel_property
	return p
}

func (*Rel_propertyContext) IsRel_propertyContext() {}

func NewRel_propertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rel_propertyContext {
	var p = new(Rel_propertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_rel_property

	return p
}

func (s *Rel_propertyContext) GetParser() antlr.Parser { return s.parser }

func (s *Rel_propertyContext) GetIs_required() antlr.Token { return s.is_required }

func (s *Rel_propertyContext) SetIs_required(v antlr.Token) { s.is_required = v }

func (s *Rel_propertyContext) Property_name() IProperty_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProperty_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProperty_nameContext)
}

func (s *Rel_propertyContext) Data_type_ref() IData_type_refContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_type_refContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_type_refContext)
}

func (s *Rel_propertyContext) DOC_COMMENT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserDOC_COMMENT, 0)
}

func (s *Rel_propertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rel_propertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rel_propertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterRel_property(s)
	}
}

func (s *Rel_propertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitRel_property(s)
	}
}

func (s *Rel_propertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitRel_property(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Rel_property() (localctx IRel_propertyContext) {
	this := p
	_ = this

	localctx = NewRel_propertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, YammmGrammarParserRULE_rel_property)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserDOC_COMMENT {
		{
			p.SetState(146)
			p.Match(YammmGrammarParserDOC_COMMENT)
		}

	}
	{
		p.SetState(149)
		p.Property_name()
	}
	{
		p.SetState(150)
		p.Data_type_ref()
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(151)

			var _m = p.Match(YammmGrammarParserT__6)

			localctx.(*Rel_propertyContext).is_required = _m
		}

	}

	return localctx
}

// IProperty_nameContext is an interface to support dynamic dispatch.
type IProperty_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProperty_nameContext differentiates from other interfaces.
	IsProperty_nameContext()
}

type Property_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_nameContext() *Property_nameContext {
	var p = new(Property_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_property_name
	return p
}

func (*Property_nameContext) IsProperty_nameContext() {}

func NewProperty_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_nameContext {
	var p = new(Property_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_property_name

	return p
}

func (s *Property_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Property_nameContext) LC_WORD() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLC_WORD, 0)
}

func (s *Property_nameContext) Lc_keyword() ILc_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILc_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILc_keywordContext)
}

func (s *Property_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterProperty_name(s)
	}
}

func (s *Property_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitProperty_name(s)
	}
}

func (s *Property_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitProperty_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Property_name() (localctx IProperty_nameContext) {
	this := p
	_ = this

	localctx = NewProperty_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, YammmGrammarParserRULE_property_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(156)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case YammmGrammarParserLC_WORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.Match(YammmGrammarParserLC_WORD)
		}

	case YammmGrammarParserT__0, YammmGrammarParserT__1, YammmGrammarParserT__3, YammmGrammarParserT__4, YammmGrammarParserT__5, YammmGrammarParserT__6, YammmGrammarParserT__7, YammmGrammarParserT__8, YammmGrammarParserT__20, YammmGrammarParserT__21:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Lc_keyword()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IData_type_refContext is an interface to support dynamic dispatch.
type IData_type_refContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsData_type_refContext differentiates from other interfaces.
	IsData_type_refContext()
}

type Data_type_refContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_type_refContext() *Data_type_refContext {
	var p = new(Data_type_refContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_data_type_ref
	return p
}

func (*Data_type_refContext) IsData_type_refContext() {}

func NewData_type_refContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_type_refContext {
	var p = new(Data_type_refContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_data_type_ref

	return p
}

func (s *Data_type_refContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_type_refContext) Built_in() IBuilt_inContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuilt_inContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuilt_inContext)
}

func (s *Data_type_refContext) Alias() IAliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *Data_type_refContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_type_refContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_type_refContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterData_type_ref(s)
	}
}

func (s *Data_type_refContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitData_type_ref(s)
	}
}

func (s *Data_type_refContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitData_type_ref(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Data_type_ref() (localctx IData_type_refContext) {
	this := p
	_ = this

	localctx = NewData_type_refContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, YammmGrammarParserRULE_data_type_ref)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case YammmGrammarParserT__9, YammmGrammarParserT__10, YammmGrammarParserT__11, YammmGrammarParserT__12, YammmGrammarParserT__13, YammmGrammarParserT__14, YammmGrammarParserT__15, YammmGrammarParserT__16, YammmGrammarParserT__17, YammmGrammarParserT__18:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Built_in()
		}

	case YammmGrammarParserUC_WORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Alias()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) UC_WORD() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserUC_WORD, 0)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (s *AliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Alias() (localctx IAliasContext) {
	this := p
	_ = this

	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, YammmGrammarParserRULE_alias)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(YammmGrammarParserUC_WORD)
	}

	return localctx
}

// IAssociationContext is an interface to support dynamic dispatch.
type IAssociationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetThisName returns the thisName rule contexts.
	GetThisName() IAny_nameContext

	// GetThisMp returns the thisMp rule contexts.
	GetThisMp() IMultiplicityContext

	// GetToType returns the toType rule contexts.
	GetToType() IType_nameContext

	// GetReverse_name returns the reverse_name rule contexts.
	GetReverse_name() IAny_nameContext

	// GetReverseMp returns the reverseMp rule contexts.
	GetReverseMp() IMultiplicityContext

	// SetThisName sets the thisName rule contexts.
	SetThisName(IAny_nameContext)

	// SetThisMp sets the thisMp rule contexts.
	SetThisMp(IMultiplicityContext)

	// SetToType sets the toType rule contexts.
	SetToType(IType_nameContext)

	// SetReverse_name sets the reverse_name rule contexts.
	SetReverse_name(IAny_nameContext)

	// SetReverseMp sets the reverseMp rule contexts.
	SetReverseMp(IMultiplicityContext)

	// IsAssociationContext differentiates from other interfaces.
	IsAssociationContext()
}

type AssociationContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	thisName     IAny_nameContext
	thisMp       IMultiplicityContext
	toType       IType_nameContext
	reverse_name IAny_nameContext
	reverseMp    IMultiplicityContext
}

func NewEmptyAssociationContext() *AssociationContext {
	var p = new(AssociationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_association
	return p
}

func (*AssociationContext) IsAssociationContext() {}

func NewAssociationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssociationContext {
	var p = new(AssociationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_association

	return p
}

func (s *AssociationContext) GetParser() antlr.Parser { return s.parser }

func (s *AssociationContext) GetThisName() IAny_nameContext { return s.thisName }

func (s *AssociationContext) GetThisMp() IMultiplicityContext { return s.thisMp }

func (s *AssociationContext) GetToType() IType_nameContext { return s.toType }

func (s *AssociationContext) GetReverse_name() IAny_nameContext { return s.reverse_name }

func (s *AssociationContext) GetReverseMp() IMultiplicityContext { return s.reverseMp }

func (s *AssociationContext) SetThisName(v IAny_nameContext) { s.thisName = v }

func (s *AssociationContext) SetThisMp(v IMultiplicityContext) { s.thisMp = v }

func (s *AssociationContext) SetToType(v IType_nameContext) { s.toType = v }

func (s *AssociationContext) SetReverse_name(v IAny_nameContext) { s.reverse_name = v }

func (s *AssociationContext) SetReverseMp(v IMultiplicityContext) { s.reverseMp = v }

func (s *AssociationContext) ASSOC() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserASSOC, 0)
}

func (s *AssociationContext) AllAny_name() []IAny_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAny_nameContext); ok {
			len++
		}
	}

	tst := make([]IAny_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAny_nameContext); ok {
			tst[i] = t.(IAny_nameContext)
			i++
		}
	}

	return tst
}

func (s *AssociationContext) Any_name(i int) IAny_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAny_nameContext); ok {
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

	return t.(IAny_nameContext)
}

func (s *AssociationContext) Type_name() IType_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_nameContext)
}

func (s *AssociationContext) DOC_COMMENT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserDOC_COMMENT, 0)
}

func (s *AssociationContext) SLASH() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserSLASH, 0)
}

func (s *AssociationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACE, 0)
}

func (s *AssociationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACE, 0)
}

func (s *AssociationContext) AllMultiplicity() []IMultiplicityContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicityContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicityContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicityContext); ok {
			tst[i] = t.(IMultiplicityContext)
			i++
		}
	}

	return tst
}

func (s *AssociationContext) Multiplicity(i int) IMultiplicityContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicityContext); ok {
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

	return t.(IMultiplicityContext)
}

func (s *AssociationContext) Relation_body() IRelation_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelation_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelation_bodyContext)
}

func (s *AssociationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssociationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssociationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterAssociation(s)
	}
}

func (s *AssociationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitAssociation(s)
	}
}

func (s *AssociationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitAssociation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Association() (localctx IAssociationContext) {
	this := p
	_ = this

	localctx = NewAssociationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, YammmGrammarParserRULE_association)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserDOC_COMMENT {
		{
			p.SetState(164)
			p.Match(YammmGrammarParserDOC_COMMENT)
		}

	}
	{
		p.SetState(167)
		p.Match(YammmGrammarParserASSOC)
	}
	{
		p.SetState(168)

		var _x = p.Any_name()

		localctx.(*AssociationContext).thisName = _x
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserLPAR {
		{
			p.SetState(169)

			var _x = p.Multiplicity()

			localctx.(*AssociationContext).thisMp = _x
		}

	}
	{
		p.SetState(172)

		var _x = p.Type_name()

		localctx.(*AssociationContext).toType = _x
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserSLASH {
		{
			p.SetState(173)
			p.Match(YammmGrammarParserSLASH)
		}
		{
			p.SetState(174)

			var _x = p.Any_name()

			localctx.(*AssociationContext).reverse_name = _x
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == YammmGrammarParserLPAR {
			{
				p.SetState(175)

				var _x = p.Multiplicity()

				localctx.(*AssociationContext).reverseMp = _x
			}

		}

	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserLBRACE {
		{
			p.SetState(180)
			p.Match(YammmGrammarParserLBRACE)
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&576460752309715958) != 0 || _la == YammmGrammarParserLC_WORD {
			{
				p.SetState(181)
				p.Relation_body()
			}

		}
		{
			p.SetState(184)
			p.Match(YammmGrammarParserRBRACE)
		}

	}

	return localctx
}

// ICompositionContext is an interface to support dynamic dispatch.
type ICompositionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetThisName returns the thisName rule contexts.
	GetThisName() IAny_nameContext

	// GetThisMp returns the thisMp rule contexts.
	GetThisMp() IMultiplicityContext

	// GetToType returns the toType rule contexts.
	GetToType() IType_nameContext

	// GetReverse_name returns the reverse_name rule contexts.
	GetReverse_name() IAny_nameContext

	// GetReverseMp returns the reverseMp rule contexts.
	GetReverseMp() IMultiplicityContext

	// SetThisName sets the thisName rule contexts.
	SetThisName(IAny_nameContext)

	// SetThisMp sets the thisMp rule contexts.
	SetThisMp(IMultiplicityContext)

	// SetToType sets the toType rule contexts.
	SetToType(IType_nameContext)

	// SetReverse_name sets the reverse_name rule contexts.
	SetReverse_name(IAny_nameContext)

	// SetReverseMp sets the reverseMp rule contexts.
	SetReverseMp(IMultiplicityContext)

	// IsCompositionContext differentiates from other interfaces.
	IsCompositionContext()
}

type CompositionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	thisName     IAny_nameContext
	thisMp       IMultiplicityContext
	toType       IType_nameContext
	reverse_name IAny_nameContext
	reverseMp    IMultiplicityContext
}

func NewEmptyCompositionContext() *CompositionContext {
	var p = new(CompositionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_composition
	return p
}

func (*CompositionContext) IsCompositionContext() {}

func NewCompositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompositionContext {
	var p = new(CompositionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_composition

	return p
}

func (s *CompositionContext) GetParser() antlr.Parser { return s.parser }

func (s *CompositionContext) GetThisName() IAny_nameContext { return s.thisName }

func (s *CompositionContext) GetThisMp() IMultiplicityContext { return s.thisMp }

func (s *CompositionContext) GetToType() IType_nameContext { return s.toType }

func (s *CompositionContext) GetReverse_name() IAny_nameContext { return s.reverse_name }

func (s *CompositionContext) GetReverseMp() IMultiplicityContext { return s.reverseMp }

func (s *CompositionContext) SetThisName(v IAny_nameContext) { s.thisName = v }

func (s *CompositionContext) SetThisMp(v IMultiplicityContext) { s.thisMp = v }

func (s *CompositionContext) SetToType(v IType_nameContext) { s.toType = v }

func (s *CompositionContext) SetReverse_name(v IAny_nameContext) { s.reverse_name = v }

func (s *CompositionContext) SetReverseMp(v IMultiplicityContext) { s.reverseMp = v }

func (s *CompositionContext) COMP() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOMP, 0)
}

func (s *CompositionContext) AllAny_name() []IAny_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAny_nameContext); ok {
			len++
		}
	}

	tst := make([]IAny_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAny_nameContext); ok {
			tst[i] = t.(IAny_nameContext)
			i++
		}
	}

	return tst
}

func (s *CompositionContext) Any_name(i int) IAny_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAny_nameContext); ok {
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

	return t.(IAny_nameContext)
}

func (s *CompositionContext) Type_name() IType_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_nameContext)
}

func (s *CompositionContext) DOC_COMMENT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserDOC_COMMENT, 0)
}

func (s *CompositionContext) SLASH() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserSLASH, 0)
}

func (s *CompositionContext) AllMultiplicity() []IMultiplicityContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicityContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicityContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicityContext); ok {
			tst[i] = t.(IMultiplicityContext)
			i++
		}
	}

	return tst
}

func (s *CompositionContext) Multiplicity(i int) IMultiplicityContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicityContext); ok {
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

	return t.(IMultiplicityContext)
}

func (s *CompositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterComposition(s)
	}
}

func (s *CompositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitComposition(s)
	}
}

func (s *CompositionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitComposition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Composition() (localctx ICompositionContext) {
	this := p
	_ = this

	localctx = NewCompositionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, YammmGrammarParserRULE_composition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserDOC_COMMENT {
		{
			p.SetState(187)
			p.Match(YammmGrammarParserDOC_COMMENT)
		}

	}
	{
		p.SetState(190)
		p.Match(YammmGrammarParserCOMP)
	}
	{
		p.SetState(191)

		var _x = p.Any_name()

		localctx.(*CompositionContext).thisName = _x
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserLPAR {
		{
			p.SetState(192)

			var _x = p.Multiplicity()

			localctx.(*CompositionContext).thisMp = _x
		}

	}
	{
		p.SetState(195)

		var _x = p.Type_name()

		localctx.(*CompositionContext).toType = _x
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserSLASH {
		{
			p.SetState(196)
			p.Match(YammmGrammarParserSLASH)
		}
		{
			p.SetState(197)

			var _x = p.Any_name()

			localctx.(*CompositionContext).reverse_name = _x
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == YammmGrammarParserLPAR {
			{
				p.SetState(198)

				var _x = p.Multiplicity()

				localctx.(*CompositionContext).reverseMp = _x
			}

		}

	}

	return localctx
}

// IAny_nameContext is an interface to support dynamic dispatch.
type IAny_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAny_nameContext differentiates from other interfaces.
	IsAny_nameContext()
}

type Any_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAny_nameContext() *Any_nameContext {
	var p = new(Any_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_any_name
	return p
}

func (*Any_nameContext) IsAny_nameContext() {}

func NewAny_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Any_nameContext {
	var p = new(Any_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_any_name

	return p
}

func (s *Any_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Any_nameContext) UC_WORD() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserUC_WORD, 0)
}

func (s *Any_nameContext) LC_WORD() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLC_WORD, 0)
}

func (s *Any_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Any_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Any_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterAny_name(s)
	}
}

func (s *Any_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitAny_name(s)
	}
}

func (s *Any_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitAny_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Any_name() (localctx IAny_nameContext) {
	this := p
	_ = this

	localctx = NewAny_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, YammmGrammarParserRULE_any_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		_la = p.GetTokenStream().LA(1)

		if !(_la == YammmGrammarParserUC_WORD || _la == YammmGrammarParserLC_WORD) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMultiplicityContext is an interface to support dynamic dispatch.
type IMultiplicityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicityContext differentiates from other interfaces.
	IsMultiplicityContext()
}

type MultiplicityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicityContext() *MultiplicityContext {
	var p = new(MultiplicityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_multiplicity
	return p
}

func (*MultiplicityContext) IsMultiplicityContext() {}

func NewMultiplicityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicityContext {
	var p = new(MultiplicityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_multiplicity

	return p
}

func (s *MultiplicityContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicityContext) LPAR() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLPAR, 0)
}

func (s *MultiplicityContext) RPAR() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRPAR, 0)
}

func (s *MultiplicityContext) USCORE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserUSCORE, 0)
}

func (s *MultiplicityContext) COLON() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOLON, 0)
}

func (s *MultiplicityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterMultiplicity(s)
	}
}

func (s *MultiplicityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitMultiplicity(s)
	}
}

func (s *MultiplicityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitMultiplicity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Multiplicity() (localctx IMultiplicityContext) {
	this := p
	_ = this

	localctx = NewMultiplicityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, YammmGrammarParserRULE_multiplicity)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(YammmGrammarParserLPAR)
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case YammmGrammarParserUSCORE:
		{
			p.SetState(206)
			p.Match(YammmGrammarParserUSCORE)
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == YammmGrammarParserCOLON {
			{
				p.SetState(207)
				p.Match(YammmGrammarParserCOLON)
			}
			{
				p.SetState(208)
				_la = p.GetTokenStream().LA(1)

				if !(_la == YammmGrammarParserT__7 || _la == YammmGrammarParserT__8) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case YammmGrammarParserT__7:
		{
			p.SetState(211)
			p.Match(YammmGrammarParserT__7)
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == YammmGrammarParserCOLON {
			{
				p.SetState(212)
				p.Match(YammmGrammarParserCOLON)
			}
			{
				p.SetState(213)
				_la = p.GetTokenStream().LA(1)

				if !(_la == YammmGrammarParserT__7 || _la == YammmGrammarParserT__8) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case YammmGrammarParserT__8:
		{
			p.SetState(216)
			p.Match(YammmGrammarParserT__8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(219)
		p.Match(YammmGrammarParserRPAR)
	}

	return localctx
}

// IRelation_bodyContext is an interface to support dynamic dispatch.
type IRelation_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelation_bodyContext differentiates from other interfaces.
	IsRelation_bodyContext()
}

type Relation_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelation_bodyContext() *Relation_bodyContext {
	var p = new(Relation_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_relation_body
	return p
}

func (*Relation_bodyContext) IsRelation_bodyContext() {}

func NewRelation_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Relation_bodyContext {
	var p = new(Relation_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_relation_body

	return p
}

func (s *Relation_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Relation_bodyContext) AllRel_property() []IRel_propertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRel_propertyContext); ok {
			len++
		}
	}

	tst := make([]IRel_propertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRel_propertyContext); ok {
			tst[i] = t.(IRel_propertyContext)
			i++
		}
	}

	return tst
}

func (s *Relation_bodyContext) Rel_property(i int) IRel_propertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRel_propertyContext); ok {
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

	return t.(IRel_propertyContext)
}

func (s *Relation_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Relation_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Relation_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterRelation_body(s)
	}
}

func (s *Relation_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitRelation_body(s)
	}
}

func (s *Relation_bodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitRelation_body(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Relation_body() (localctx IRelation_bodyContext) {
	this := p
	_ = this

	localctx = NewRelation_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, YammmGrammarParserRULE_relation_body)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&576460752309715958) != 0 || _la == YammmGrammarParserLC_WORD {
		{
			p.SetState(221)
			p.Rel_property()
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBuilt_inContext is an interface to support dynamic dispatch.
type IBuilt_inContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuilt_inContext differentiates from other interfaces.
	IsBuilt_inContext()
}

type Built_inContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuilt_inContext() *Built_inContext {
	var p = new(Built_inContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_built_in
	return p
}

func (*Built_inContext) IsBuilt_inContext() {}

func NewBuilt_inContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Built_inContext {
	var p = new(Built_inContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_built_in

	return p
}

func (s *Built_inContext) GetParser() antlr.Parser { return s.parser }

func (s *Built_inContext) IntegerT() IIntegerTContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerTContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerTContext)
}

func (s *Built_inContext) FloatT() IFloatTContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatTContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatTContext)
}

func (s *Built_inContext) BoolT() IBoolTContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolTContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolTContext)
}

func (s *Built_inContext) StringT() IStringTContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringTContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringTContext)
}

func (s *Built_inContext) EnumT() IEnumTContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumTContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumTContext)
}

func (s *Built_inContext) PatternT() IPatternTContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternTContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternTContext)
}

func (s *Built_inContext) TimestampT() ITimestampTContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimestampTContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimestampTContext)
}

func (s *Built_inContext) DateT() IDateTContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateTContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateTContext)
}

func (s *Built_inContext) UuidT() IUuidTContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUuidTContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUuidTContext)
}

func (s *Built_inContext) SpacevectorT() ISpacevectorTContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpacevectorTContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpacevectorTContext)
}

func (s *Built_inContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Built_inContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Built_inContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterBuilt_in(s)
	}
}

func (s *Built_inContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitBuilt_in(s)
	}
}

func (s *Built_inContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitBuilt_in(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Built_in() (localctx IBuilt_inContext) {
	this := p
	_ = this

	localctx = NewBuilt_inContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, YammmGrammarParserRULE_built_in)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(236)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case YammmGrammarParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)
			p.IntegerT()
		}

	case YammmGrammarParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.FloatT()
		}

	case YammmGrammarParserT__11:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(228)
			p.BoolT()
		}

	case YammmGrammarParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(229)
			p.StringT()
		}

	case YammmGrammarParserT__13:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(230)
			p.EnumT()
		}

	case YammmGrammarParserT__14:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(231)
			p.PatternT()
		}

	case YammmGrammarParserT__15:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(232)
			p.TimestampT()
		}

	case YammmGrammarParserT__17:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(233)
			p.DateT()
		}

	case YammmGrammarParserT__18:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(234)
			p.UuidT()
		}

	case YammmGrammarParserT__16:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(235)
			p.SpacevectorT()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerTContext is an interface to support dynamic dispatch.
type IIntegerTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMin returns the min token.
	GetMin() antlr.Token

	// GetMax returns the max token.
	GetMax() antlr.Token

	// SetMin sets the min token.
	SetMin(antlr.Token)

	// SetMax sets the max token.
	SetMax(antlr.Token)

	// IsIntegerTContext differentiates from other interfaces.
	IsIntegerTContext()
}

type IntegerTContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	min    antlr.Token
	max    antlr.Token
}

func NewEmptyIntegerTContext() *IntegerTContext {
	var p = new(IntegerTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_integerT
	return p
}

func (*IntegerTContext) IsIntegerTContext() {}

func NewIntegerTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerTContext {
	var p = new(IntegerTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_integerT

	return p
}

func (s *IntegerTContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerTContext) GetMin() antlr.Token { return s.min }

func (s *IntegerTContext) GetMax() antlr.Token { return s.max }

func (s *IntegerTContext) SetMin(v antlr.Token) { s.min = v }

func (s *IntegerTContext) SetMax(v antlr.Token) { s.max = v }

func (s *IntegerTContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACK, 0)
}

func (s *IntegerTContext) COMMA() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOMMA, 0)
}

func (s *IntegerTContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACK, 0)
}

func (s *IntegerTContext) AllUSCORE() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserUSCORE)
}

func (s *IntegerTContext) USCORE(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserUSCORE, i)
}

func (s *IntegerTContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserINTEGER)
}

func (s *IntegerTContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserINTEGER, i)
}

func (s *IntegerTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterIntegerT(s)
	}
}

func (s *IntegerTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitIntegerT(s)
	}
}

func (s *IntegerTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitIntegerT(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) IntegerT() (localctx IIntegerTContext) {
	this := p
	_ = this

	localctx = NewIntegerTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, YammmGrammarParserRULE_integerT)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(YammmGrammarParserT__9)
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserLBRACK {
		{
			p.SetState(239)
			p.Match(YammmGrammarParserLBRACK)
		}
		{
			p.SetState(240)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*IntegerTContext).min = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == YammmGrammarParserUSCORE || _la == YammmGrammarParserINTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*IntegerTContext).min = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(241)
			p.Match(YammmGrammarParserCOMMA)
		}
		{
			p.SetState(242)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*IntegerTContext).max = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == YammmGrammarParserUSCORE || _la == YammmGrammarParserINTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*IntegerTContext).max = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(243)
			p.Match(YammmGrammarParserRBRACK)
		}

	}

	return localctx
}

// IFloatTContext is an interface to support dynamic dispatch.
type IFloatTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMin returns the min token.
	GetMin() antlr.Token

	// GetMax returns the max token.
	GetMax() antlr.Token

	// SetMin sets the min token.
	SetMin(antlr.Token)

	// SetMax sets the max token.
	SetMax(antlr.Token)

	// IsFloatTContext differentiates from other interfaces.
	IsFloatTContext()
}

type FloatTContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	min    antlr.Token
	max    antlr.Token
}

func NewEmptyFloatTContext() *FloatTContext {
	var p = new(FloatTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_floatT
	return p
}

func (*FloatTContext) IsFloatTContext() {}

func NewFloatTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatTContext {
	var p = new(FloatTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_floatT

	return p
}

func (s *FloatTContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatTContext) GetMin() antlr.Token { return s.min }

func (s *FloatTContext) GetMax() antlr.Token { return s.max }

func (s *FloatTContext) SetMin(v antlr.Token) { s.min = v }

func (s *FloatTContext) SetMax(v antlr.Token) { s.max = v }

func (s *FloatTContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACK, 0)
}

func (s *FloatTContext) COMMA() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOMMA, 0)
}

func (s *FloatTContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACK, 0)
}

func (s *FloatTContext) AllUSCORE() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserUSCORE)
}

func (s *FloatTContext) USCORE(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserUSCORE, i)
}

func (s *FloatTContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserINTEGER)
}

func (s *FloatTContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserINTEGER, i)
}

func (s *FloatTContext) AllFLOAT() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserFLOAT)
}

func (s *FloatTContext) FLOAT(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserFLOAT, i)
}

func (s *FloatTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterFloatT(s)
	}
}

func (s *FloatTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitFloatT(s)
	}
}

func (s *FloatTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitFloatT(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) FloatT() (localctx IFloatTContext) {
	this := p
	_ = this

	localctx = NewFloatTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, YammmGrammarParserRULE_floatT)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(YammmGrammarParserT__10)
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserLBRACK {
		{
			p.SetState(247)
			p.Match(YammmGrammarParserLBRACK)
		}
		{
			p.SetState(248)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*FloatTContext).min = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-36)) & ^0x3f) == 0 && ((int64(1)<<(_la-36))&805306369) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*FloatTContext).min = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(249)
			p.Match(YammmGrammarParserCOMMA)
		}
		{
			p.SetState(250)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*FloatTContext).max = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-36)) & ^0x3f) == 0 && ((int64(1)<<(_la-36))&805306369) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*FloatTContext).max = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(251)
			p.Match(YammmGrammarParserRBRACK)
		}

	}

	return localctx
}

// IBoolTContext is an interface to support dynamic dispatch.
type IBoolTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolTContext differentiates from other interfaces.
	IsBoolTContext()
}

type BoolTContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolTContext() *BoolTContext {
	var p = new(BoolTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_boolT
	return p
}

func (*BoolTContext) IsBoolTContext() {}

func NewBoolTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolTContext {
	var p = new(BoolTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_boolT

	return p
}

func (s *BoolTContext) GetParser() antlr.Parser { return s.parser }
func (s *BoolTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterBoolT(s)
	}
}

func (s *BoolTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitBoolT(s)
	}
}

func (s *BoolTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitBoolT(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) BoolT() (localctx IBoolTContext) {
	this := p
	_ = this

	localctx = NewBoolTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, YammmGrammarParserRULE_boolT)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(YammmGrammarParserT__11)
	}

	return localctx
}

// IStringTContext is an interface to support dynamic dispatch.
type IStringTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMin returns the min token.
	GetMin() antlr.Token

	// GetMax returns the max token.
	GetMax() antlr.Token

	// SetMin sets the min token.
	SetMin(antlr.Token)

	// SetMax sets the max token.
	SetMax(antlr.Token)

	// IsStringTContext differentiates from other interfaces.
	IsStringTContext()
}

type StringTContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	min    antlr.Token
	max    antlr.Token
}

func NewEmptyStringTContext() *StringTContext {
	var p = new(StringTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_stringT
	return p
}

func (*StringTContext) IsStringTContext() {}

func NewStringTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringTContext {
	var p = new(StringTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_stringT

	return p
}

func (s *StringTContext) GetParser() antlr.Parser { return s.parser }

func (s *StringTContext) GetMin() antlr.Token { return s.min }

func (s *StringTContext) GetMax() antlr.Token { return s.max }

func (s *StringTContext) SetMin(v antlr.Token) { s.min = v }

func (s *StringTContext) SetMax(v antlr.Token) { s.max = v }

func (s *StringTContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACK, 0)
}

func (s *StringTContext) COMMA() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOMMA, 0)
}

func (s *StringTContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACK, 0)
}

func (s *StringTContext) AllUSCORE() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserUSCORE)
}

func (s *StringTContext) USCORE(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserUSCORE, i)
}

func (s *StringTContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserINTEGER)
}

func (s *StringTContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserINTEGER, i)
}

func (s *StringTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterStringT(s)
	}
}

func (s *StringTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitStringT(s)
	}
}

func (s *StringTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitStringT(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) StringT() (localctx IStringTContext) {
	this := p
	_ = this

	localctx = NewStringTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, YammmGrammarParserRULE_stringT)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(YammmGrammarParserT__12)
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserLBRACK {
		{
			p.SetState(257)
			p.Match(YammmGrammarParserLBRACK)
		}
		{
			p.SetState(258)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*StringTContext).min = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == YammmGrammarParserUSCORE || _la == YammmGrammarParserINTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*StringTContext).min = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(259)
			p.Match(YammmGrammarParserCOMMA)
		}
		{
			p.SetState(260)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*StringTContext).max = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == YammmGrammarParserUSCORE || _la == YammmGrammarParserINTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*StringTContext).max = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(261)
			p.Match(YammmGrammarParserRBRACK)
		}

	}

	return localctx
}

// IEnumTContext is an interface to support dynamic dispatch.
type IEnumTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTContext differentiates from other interfaces.
	IsEnumTContext()
}

type EnumTContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTContext() *EnumTContext {
	var p = new(EnumTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_enumT
	return p
}

func (*EnumTContext) IsEnumTContext() {}

func NewEnumTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTContext {
	var p = new(EnumTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_enumT

	return p
}

func (s *EnumTContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACK, 0)
}

func (s *EnumTContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserSTRING)
}

func (s *EnumTContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserSTRING, i)
}

func (s *EnumTContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACK, 0)
}

func (s *EnumTContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserCOMMA)
}

func (s *EnumTContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOMMA, i)
}

func (s *EnumTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterEnumT(s)
	}
}

func (s *EnumTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitEnumT(s)
	}
}

func (s *EnumTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitEnumT(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) EnumT() (localctx IEnumTContext) {
	this := p
	_ = this

	localctx = NewEnumTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, YammmGrammarParserRULE_enumT)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(YammmGrammarParserT__13)
	}
	{
		p.SetState(265)
		p.Match(YammmGrammarParserLBRACK)
	}
	{
		p.SetState(266)
		p.Match(YammmGrammarParserSTRING)
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(267)
				p.Match(YammmGrammarParserCOMMA)
			}
			{
				p.SetState(268)
				p.Match(YammmGrammarParserSTRING)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserCOMMA {
		{
			p.SetState(273)
			p.Match(YammmGrammarParserCOMMA)
		}

	}
	{
		p.SetState(276)
		p.Match(YammmGrammarParserRBRACK)
	}

	return localctx
}

// IPatternTContext is an interface to support dynamic dispatch.
type IPatternTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatternTContext differentiates from other interfaces.
	IsPatternTContext()
}

type PatternTContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternTContext() *PatternTContext {
	var p = new(PatternTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_patternT
	return p
}

func (*PatternTContext) IsPatternTContext() {}

func NewPatternTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternTContext {
	var p = new(PatternTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_patternT

	return p
}

func (s *PatternTContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternTContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACK, 0)
}

func (s *PatternTContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserSTRING)
}

func (s *PatternTContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserSTRING, i)
}

func (s *PatternTContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACK, 0)
}

func (s *PatternTContext) COMMA() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOMMA, 0)
}

func (s *PatternTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterPatternT(s)
	}
}

func (s *PatternTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitPatternT(s)
	}
}

func (s *PatternTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitPatternT(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) PatternT() (localctx IPatternTContext) {
	this := p
	_ = this

	localctx = NewPatternTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, YammmGrammarParserRULE_patternT)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(YammmGrammarParserT__14)
	}
	{
		p.SetState(279)
		p.Match(YammmGrammarParserLBRACK)
	}
	{
		p.SetState(280)
		p.Match(YammmGrammarParserSTRING)
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserCOMMA {
		{
			p.SetState(281)
			p.Match(YammmGrammarParserCOMMA)
		}
		{
			p.SetState(282)
			p.Match(YammmGrammarParserSTRING)
		}

	}
	{
		p.SetState(285)
		p.Match(YammmGrammarParserRBRACK)
	}

	return localctx
}

// ITimestampTContext is an interface to support dynamic dispatch.
type ITimestampTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFormat returns the format token.
	GetFormat() antlr.Token

	// SetFormat sets the format token.
	SetFormat(antlr.Token)

	// IsTimestampTContext differentiates from other interfaces.
	IsTimestampTContext()
}

type TimestampTContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	format antlr.Token
}

func NewEmptyTimestampTContext() *TimestampTContext {
	var p = new(TimestampTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_timestampT
	return p
}

func (*TimestampTContext) IsTimestampTContext() {}

func NewTimestampTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimestampTContext {
	var p = new(TimestampTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_timestampT

	return p
}

func (s *TimestampTContext) GetParser() antlr.Parser { return s.parser }

func (s *TimestampTContext) GetFormat() antlr.Token { return s.format }

func (s *TimestampTContext) SetFormat(v antlr.Token) { s.format = v }

func (s *TimestampTContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACK, 0)
}

func (s *TimestampTContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACK, 0)
}

func (s *TimestampTContext) STRING() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserSTRING, 0)
}

func (s *TimestampTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimestampTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterTimestampT(s)
	}
}

func (s *TimestampTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitTimestampT(s)
	}
}

func (s *TimestampTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitTimestampT(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) TimestampT() (localctx ITimestampTContext) {
	this := p
	_ = this

	localctx = NewTimestampTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, YammmGrammarParserRULE_timestampT)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Match(YammmGrammarParserT__15)
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserLBRACK {
		{
			p.SetState(288)
			p.Match(YammmGrammarParserLBRACK)
		}
		{
			p.SetState(289)

			var _m = p.Match(YammmGrammarParserSTRING)

			localctx.(*TimestampTContext).format = _m
		}
		{
			p.SetState(290)
			p.Match(YammmGrammarParserRBRACK)
		}

	}

	return localctx
}

// ISpacevectorTContext is an interface to support dynamic dispatch.
type ISpacevectorTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDimensions returns the dimensions token.
	GetDimensions() antlr.Token

	// SetDimensions sets the dimensions token.
	SetDimensions(antlr.Token)

	// IsSpacevectorTContext differentiates from other interfaces.
	IsSpacevectorTContext()
}

type SpacevectorTContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	dimensions antlr.Token
}

func NewEmptySpacevectorTContext() *SpacevectorTContext {
	var p = new(SpacevectorTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_spacevectorT
	return p
}

func (*SpacevectorTContext) IsSpacevectorTContext() {}

func NewSpacevectorTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpacevectorTContext {
	var p = new(SpacevectorTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_spacevectorT

	return p
}

func (s *SpacevectorTContext) GetParser() antlr.Parser { return s.parser }

func (s *SpacevectorTContext) GetDimensions() antlr.Token { return s.dimensions }

func (s *SpacevectorTContext) SetDimensions(v antlr.Token) { s.dimensions = v }

func (s *SpacevectorTContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACK, 0)
}

func (s *SpacevectorTContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACK, 0)
}

func (s *SpacevectorTContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserINTEGER, 0)
}

func (s *SpacevectorTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpacevectorTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpacevectorTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterSpacevectorT(s)
	}
}

func (s *SpacevectorTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitSpacevectorT(s)
	}
}

func (s *SpacevectorTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitSpacevectorT(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) SpacevectorT() (localctx ISpacevectorTContext) {
	this := p
	_ = this

	localctx = NewSpacevectorTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, YammmGrammarParserRULE_spacevectorT)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(YammmGrammarParserT__16)
	}
	{
		p.SetState(294)
		p.Match(YammmGrammarParserLBRACK)
	}
	{
		p.SetState(295)

		var _m = p.Match(YammmGrammarParserINTEGER)

		localctx.(*SpacevectorTContext).dimensions = _m
	}
	{
		p.SetState(296)
		p.Match(YammmGrammarParserRBRACK)
	}

	return localctx
}

// IDateTContext is an interface to support dynamic dispatch.
type IDateTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateTContext differentiates from other interfaces.
	IsDateTContext()
}

type DateTContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTContext() *DateTContext {
	var p = new(DateTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_dateT
	return p
}

func (*DateTContext) IsDateTContext() {}

func NewDateTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTContext {
	var p = new(DateTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_dateT

	return p
}

func (s *DateTContext) GetParser() antlr.Parser { return s.parser }
func (s *DateTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterDateT(s)
	}
}

func (s *DateTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitDateT(s)
	}
}

func (s *DateTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitDateT(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) DateT() (localctx IDateTContext) {
	this := p
	_ = this

	localctx = NewDateTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, YammmGrammarParserRULE_dateT)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(YammmGrammarParserT__17)
	}

	return localctx
}

// IUuidTContext is an interface to support dynamic dispatch.
type IUuidTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUuidTContext differentiates from other interfaces.
	IsUuidTContext()
}

type UuidTContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUuidTContext() *UuidTContext {
	var p = new(UuidTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_uuidT
	return p
}

func (*UuidTContext) IsUuidTContext() {}

func NewUuidTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UuidTContext {
	var p = new(UuidTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_uuidT

	return p
}

func (s *UuidTContext) GetParser() antlr.Parser { return s.parser }
func (s *UuidTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UuidTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UuidTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterUuidT(s)
	}
}

func (s *UuidTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitUuidT(s)
	}
}

func (s *UuidTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitUuidT(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) UuidT() (localctx IUuidTContext) {
	this := p
	_ = this

	localctx = NewUuidTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, YammmGrammarParserRULE_uuidT)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(YammmGrammarParserT__18)
	}

	return localctx
}

// IDatatypeKeywordContext is an interface to support dynamic dispatch.
type IDatatypeKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatatypeKeywordContext differentiates from other interfaces.
	IsDatatypeKeywordContext()
}

type DatatypeKeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatatypeKeywordContext() *DatatypeKeywordContext {
	var p = new(DatatypeKeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_datatypeKeyword
	return p
}

func (*DatatypeKeywordContext) IsDatatypeKeywordContext() {}

func NewDatatypeKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatatypeKeywordContext {
	var p = new(DatatypeKeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_datatypeKeyword

	return p
}

func (s *DatatypeKeywordContext) GetParser() antlr.Parser { return s.parser }
func (s *DatatypeKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatatypeKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatatypeKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterDatatypeKeyword(s)
	}
}

func (s *DatatypeKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitDatatypeKeyword(s)
	}
}

func (s *DatatypeKeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitDatatypeKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) DatatypeKeyword() (localctx IDatatypeKeywordContext) {
	this := p
	_ = this

	localctx = NewDatatypeKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, YammmGrammarParserRULE_datatypeKeyword)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1047552) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInvariantContext is an interface to support dynamic dispatch.
type IInvariantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMessage returns the message token.
	GetMessage() antlr.Token

	// SetMessage sets the message token.
	SetMessage(antlr.Token)

	// GetConstraint returns the constraint rule contexts.
	GetConstraint() IExprContext

	// SetConstraint sets the constraint rule contexts.
	SetConstraint(IExprContext)

	// IsInvariantContext differentiates from other interfaces.
	IsInvariantContext()
}

type InvariantContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	message    antlr.Token
	constraint IExprContext
}

func NewEmptyInvariantContext() *InvariantContext {
	var p = new(InvariantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_invariant
	return p
}

func (*InvariantContext) IsInvariantContext() {}

func NewInvariantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvariantContext {
	var p = new(InvariantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_invariant

	return p
}

func (s *InvariantContext) GetParser() antlr.Parser { return s.parser }

func (s *InvariantContext) GetMessage() antlr.Token { return s.message }

func (s *InvariantContext) SetMessage(v antlr.Token) { s.message = v }

func (s *InvariantContext) GetConstraint() IExprContext { return s.constraint }

func (s *InvariantContext) SetConstraint(v IExprContext) { s.constraint = v }

func (s *InvariantContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserEXCLAMATION, 0)
}

func (s *InvariantContext) STRING() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserSTRING, 0)
}

func (s *InvariantContext) Expr() IExprContext {
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

func (s *InvariantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvariantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvariantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterInvariant(s)
	}
}

func (s *InvariantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitInvariant(s)
	}
}

func (s *InvariantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitInvariant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Invariant() (localctx IInvariantContext) {
	this := p
	_ = this

	localctx = NewInvariantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, YammmGrammarParserRULE_invariant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(YammmGrammarParserEXCLAMATION)
	}
	{
		p.SetState(305)

		var _m = p.Match(YammmGrammarParserSTRING)

		localctx.(*InvariantContext).message = _m
	}
	{
		p.SetState(306)

		var _x = p.expr(0)

		localctx.(*InvariantContext).constraint = _x
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DatatypeNameContext struct {
	*ExprContext
	left IDatatypeKeywordContext
}

func NewDatatypeNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DatatypeNameContext {
	var p = new(DatatypeNameContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DatatypeNameContext) GetLeft() IDatatypeKeywordContext { return s.left }

func (s *DatatypeNameContext) SetLeft(v IDatatypeKeywordContext) { s.left = v }

func (s *DatatypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatatypeNameContext) DatatypeKeyword() IDatatypeKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatatypeKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatatypeKeywordContext)
}

func (s *DatatypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterDatatypeName(s)
	}
}

func (s *DatatypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitDatatypeName(s)
	}
}

func (s *DatatypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitDatatypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlusminusContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewPlusminusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusminusContext {
	var p = new(PlusminusContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PlusminusContext) GetOp() antlr.Token { return s.op }

func (s *PlusminusContext) SetOp(v antlr.Token) { s.op = v }

func (s *PlusminusContext) GetLeft() IExprContext { return s.left }

func (s *PlusminusContext) GetRight() IExprContext { return s.right }

func (s *PlusminusContext) SetLeft(v IExprContext) { s.left = v }

func (s *PlusminusContext) SetRight(v IExprContext) { s.right = v }

func (s *PlusminusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusminusContext) AllExpr() []IExprContext {
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

func (s *PlusminusContext) Expr(i int) IExprContext {
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

func (s *PlusminusContext) PLUS() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserPLUS, 0)
}

func (s *PlusminusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserMINUS, 0)
}

func (s *PlusminusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterPlusminus(s)
	}
}

func (s *PlusminusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitPlusminus(s)
	}
}

func (s *PlusminusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitPlusminus(s)

	default:
		return t.VisitChildren(s)
	}
}

type PeriodContext struct {
	*ExprContext
	left IExprContext
	name IExprContext
}

func NewPeriodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PeriodContext {
	var p = new(PeriodContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PeriodContext) GetLeft() IExprContext { return s.left }

func (s *PeriodContext) GetName() IExprContext { return s.name }

func (s *PeriodContext) SetLeft(v IExprContext) { s.left = v }

func (s *PeriodContext) SetName(v IExprContext) { s.name = v }

func (s *PeriodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PeriodContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserPERIOD, 0)
}

func (s *PeriodContext) AllExpr() []IExprContext {
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

func (s *PeriodContext) Expr(i int) IExprContext {
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

func (s *PeriodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterPeriod(s)
	}
}

func (s *PeriodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitPeriod(s)
	}
}

func (s *PeriodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitPeriod(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareContext {
	var p = new(CompareContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CompareContext) GetOp() antlr.Token { return s.op }

func (s *CompareContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareContext) GetLeft() IExprContext { return s.left }

func (s *CompareContext) GetRight() IExprContext { return s.right }

func (s *CompareContext) SetLeft(v IExprContext) { s.left = v }

func (s *CompareContext) SetRight(v IExprContext) { s.right = v }

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) AllExpr() []IExprContext {
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

func (s *CompareContext) Expr(i int) IExprContext {
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

func (s *CompareContext) GT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserGT, 0)
}

func (s *CompareContext) GTE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserGTE, 0)
}

func (s *CompareContext) LT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLT, 0)
}

func (s *CompareContext) LTE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLTE, 0)
}

func (s *CompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterCompare(s)
	}
}

func (s *CompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitCompare(s)
	}
}

func (s *CompareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitCompare(s)

	default:
		return t.VisitChildren(s)
	}
}

type UminusContext struct {
	*ExprContext
	op    antlr.Token
	right IExprContext
}

func NewUminusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UminusContext {
	var p = new(UminusContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UminusContext) GetOp() antlr.Token { return s.op }

func (s *UminusContext) SetOp(v antlr.Token) { s.op = v }

func (s *UminusContext) GetRight() IExprContext { return s.right }

func (s *UminusContext) SetRight(v IExprContext) { s.right = v }

func (s *UminusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UminusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserMINUS, 0)
}

func (s *UminusContext) Expr() IExprContext {
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

func (s *UminusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterUminus(s)
	}
}

func (s *UminusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitUminus(s)
	}
}

func (s *UminusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitUminus(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrContext) GetOp() antlr.Token { return s.op }

func (s *OrContext) SetOp(v antlr.Token) { s.op = v }

func (s *OrContext) GetLeft() IExprContext { return s.left }

func (s *OrContext) GetRight() IExprContext { return s.right }

func (s *OrContext) SetLeft(v IExprContext) { s.left = v }

func (s *OrContext) SetRight(v IExprContext) { s.right = v }

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpr() []IExprContext {
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

func (s *OrContext) Expr(i int) IExprContext {
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

func (s *OrContext) OR() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserOR, 0)
}

func (s *OrContext) HAT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserHAT, 0)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type InContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InContext {
	var p = new(InContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *InContext) GetOp() antlr.Token { return s.op }

func (s *InContext) SetOp(v antlr.Token) { s.op = v }

func (s *InContext) GetLeft() IExprContext { return s.left }

func (s *InContext) GetRight() IExprContext { return s.right }

func (s *InContext) SetLeft(v IExprContext) { s.left = v }

func (s *InContext) SetRight(v IExprContext) { s.right = v }

func (s *InContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InContext) AllExpr() []IExprContext {
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

func (s *InContext) Expr(i int) IExprContext {
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

func (s *InContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterIn(s)
	}
}

func (s *InContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitIn(s)
	}
}

func (s *InContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitIn(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatchContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewMatchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchContext {
	var p = new(MatchContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MatchContext) GetOp() antlr.Token { return s.op }

func (s *MatchContext) SetOp(v antlr.Token) { s.op = v }

func (s *MatchContext) GetLeft() IExprContext { return s.left }

func (s *MatchContext) GetRight() IExprContext { return s.right }

func (s *MatchContext) SetLeft(v IExprContext) { s.left = v }

func (s *MatchContext) SetRight(v IExprContext) { s.right = v }

func (s *MatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchContext) AllExpr() []IExprContext {
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

func (s *MatchContext) Expr(i int) IExprContext {
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

func (s *MatchContext) MATCH() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserMATCH, 0)
}

func (s *MatchContext) NOTMATCH() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserNOTMATCH, 0)
}

func (s *MatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterMatch(s)
	}
}

func (s *MatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitMatch(s)
	}
}

func (s *MatchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitMatch(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListContext struct {
	*ExprContext
	_expr  IExprContext
	values []IExprContext
}

func NewListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListContext {
	var p = new(ListContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ListContext) Get_expr() IExprContext { return s._expr }

func (s *ListContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListContext) GetValues() []IExprContext { return s.values }

func (s *ListContext) SetValues(v []IExprContext) { s.values = v }

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACK, 0)
}

func (s *ListContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACK, 0)
}

func (s *ListContext) AllExpr() []IExprContext {
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

func (s *ListContext) Expr(i int) IExprContext {
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

func (s *ListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserCOMMA)
}

func (s *ListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOMMA, i)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

type MuldivContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewMuldivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MuldivContext {
	var p = new(MuldivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MuldivContext) GetOp() antlr.Token { return s.op }

func (s *MuldivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MuldivContext) GetLeft() IExprContext { return s.left }

func (s *MuldivContext) GetRight() IExprContext { return s.right }

func (s *MuldivContext) SetLeft(v IExprContext) { s.left = v }

func (s *MuldivContext) SetRight(v IExprContext) { s.right = v }

func (s *MuldivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MuldivContext) AllExpr() []IExprContext {
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

func (s *MuldivContext) Expr(i int) IExprContext {
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

func (s *MuldivContext) STAR() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserSTAR, 0)
}

func (s *MuldivContext) SLASH() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserSLASH, 0)
}

func (s *MuldivContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserPERCENT, 0)
}

func (s *MuldivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterMuldiv(s)
	}
}

func (s *MuldivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitMuldiv(s)
	}
}

func (s *MuldivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitMuldiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type FcallContext struct {
	*ExprContext
	left   IExprContext
	op     antlr.Token
	name   antlr.Token
	args   IArgumentsContext
	params IParametersContext
	body   IExprContext
}

func NewFcallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FcallContext {
	var p = new(FcallContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FcallContext) GetOp() antlr.Token { return s.op }

func (s *FcallContext) GetName() antlr.Token { return s.name }

func (s *FcallContext) SetOp(v antlr.Token) { s.op = v }

func (s *FcallContext) SetName(v antlr.Token) { s.name = v }

func (s *FcallContext) GetLeft() IExprContext { return s.left }

func (s *FcallContext) GetArgs() IArgumentsContext { return s.args }

func (s *FcallContext) GetParams() IParametersContext { return s.params }

func (s *FcallContext) GetBody() IExprContext { return s.body }

func (s *FcallContext) SetLeft(v IExprContext) { s.left = v }

func (s *FcallContext) SetArgs(v IArgumentsContext) { s.args = v }

func (s *FcallContext) SetParams(v IParametersContext) { s.params = v }

func (s *FcallContext) SetBody(v IExprContext) { s.body = v }

func (s *FcallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FcallContext) AllExpr() []IExprContext {
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

func (s *FcallContext) Expr(i int) IExprContext {
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

func (s *FcallContext) ARROW() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserARROW, 0)
}

func (s *FcallContext) LC_WORD() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLC_WORD, 0)
}

func (s *FcallContext) UC_WORD() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserUC_WORD, 0)
}

func (s *FcallContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACE, 0)
}

func (s *FcallContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACE, 0)
}

func (s *FcallContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FcallContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FcallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterFcall(s)
	}
}

func (s *FcallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitFcall(s)
	}
}

func (s *FcallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitFcall(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotContext struct {
	*ExprContext
	op    antlr.Token
	right IExprContext
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotContext) GetOp() antlr.Token { return s.op }

func (s *NotContext) SetOp(v antlr.Token) { s.op = v }

func (s *NotContext) GetRight() IExprContext { return s.right }

func (s *NotContext) SetRight(v IExprContext) { s.right = v }

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserEXCLAMATION, 0)
}

func (s *NotContext) Expr() IExprContext {
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

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitNot(s)
	}
}

func (s *NotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtContext struct {
	*ExprContext
	left  IExprContext
	_expr IExprContext
	right []IExprContext
}

func NewAtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtContext {
	var p = new(AtContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AtContext) GetLeft() IExprContext { return s.left }

func (s *AtContext) Get_expr() IExprContext { return s._expr }

func (s *AtContext) SetLeft(v IExprContext) { s.left = v }

func (s *AtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *AtContext) GetRight() []IExprContext { return s.right }

func (s *AtContext) SetRight(v []IExprContext) { s.right = v }

func (s *AtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACK, 0)
}

func (s *AtContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACK, 0)
}

func (s *AtContext) AllExpr() []IExprContext {
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

func (s *AtContext) Expr(i int) IExprContext {
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

func (s *AtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserCOMMA)
}

func (s *AtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOMMA, i)
}

func (s *AtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterAt(s)
	}
}

func (s *AtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitAt(s)
	}
}

func (s *AtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitAt(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationNameContext struct {
	*ExprContext
	left antlr.Token
}

func NewRelationNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationNameContext {
	var p = new(RelationNameContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RelationNameContext) GetLeft() antlr.Token { return s.left }

func (s *RelationNameContext) SetLeft(v antlr.Token) { s.left = v }

func (s *RelationNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationNameContext) UC_WORD() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserUC_WORD, 0)
}

func (s *RelationNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterRelationName(s)
	}
}

func (s *RelationNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitRelationName(s)
	}
}

func (s *RelationNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitRelationName(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndContext) GetOp() antlr.Token { return s.op }

func (s *AndContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndContext) GetLeft() IExprContext { return s.left }

func (s *AndContext) GetRight() IExprContext { return s.right }

func (s *AndContext) SetLeft(v IExprContext) { s.left = v }

func (s *AndContext) SetRight(v IExprContext) { s.right = v }

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpr() []IExprContext {
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

func (s *AndContext) Expr(i int) IExprContext {
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

func (s *AndContext) AND() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserAND, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableContext struct {
	*ExprContext
	left antlr.Token
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VariableContext) GetLeft() antlr.Token { return s.left }

func (s *VariableContext) SetLeft(v antlr.Token) { s.left = v }

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserVARIABLE, 0)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

type NameContext struct {
	*ExprContext
	left IProperty_nameContext
}

func NewNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameContext {
	var p = new(NameContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NameContext) GetLeft() IProperty_nameContext { return s.left }

func (s *NameContext) SetLeft(v IProperty_nameContext) { s.left = v }

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) Property_name() IProperty_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProperty_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProperty_nameContext)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueContext struct {
	*ExprContext
	left ILiteralContext
}

func NewValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueContext {
	var p = new(ValueContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ValueContext) GetLeft() ILiteralContext { return s.left }

func (s *ValueContext) SetLeft(v ILiteralContext) { s.left = v }

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualityContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewEqualityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityContext {
	var p = new(EqualityContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualityContext) GetOp() antlr.Token { return s.op }

func (s *EqualityContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityContext) GetLeft() IExprContext { return s.left }

func (s *EqualityContext) GetRight() IExprContext { return s.right }

func (s *EqualityContext) SetLeft(v IExprContext) { s.left = v }

func (s *EqualityContext) SetRight(v IExprContext) { s.right = v }

func (s *EqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityContext) AllExpr() []IExprContext {
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

func (s *EqualityContext) Expr(i int) IExprContext {
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

func (s *EqualityContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserEQUAL, 0)
}

func (s *EqualityContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserNOTEQUAL, 0)
}

func (s *EqualityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterEquality(s)
	}
}

func (s *EqualityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitEquality(s)
	}
}

func (s *EqualityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitEquality(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfContext struct {
	*ExprContext
	left        IExprContext
	op          antlr.Token
	trueBranch  IExprContext
	falseBranch IExprContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IfContext) GetOp() antlr.Token { return s.op }

func (s *IfContext) SetOp(v antlr.Token) { s.op = v }

func (s *IfContext) GetLeft() IExprContext { return s.left }

func (s *IfContext) GetTrueBranch() IExprContext { return s.trueBranch }

func (s *IfContext) GetFalseBranch() IExprContext { return s.falseBranch }

func (s *IfContext) SetLeft(v IExprContext) { s.left = v }

func (s *IfContext) SetTrueBranch(v IExprContext) { s.trueBranch = v }

func (s *IfContext) SetFalseBranch(v IExprContext) { s.falseBranch = v }

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLBRACE, 0)
}

func (s *IfContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRBRACE, 0)
}

func (s *IfContext) AllExpr() []IExprContext {
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

func (s *IfContext) Expr(i int) IExprContext {
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

func (s *IfContext) QMARK() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserQMARK, 0)
}

func (s *IfContext) COLON() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOLON, 0)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitIf(s)
	}
}

func (s *IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralNilContext struct {
	*ExprContext
}

func NewLiteralNilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralNilContext {
	var p = new(LiteralNilContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiteralNilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralNilContext) USCORE() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserUSCORE, 0)
}

func (s *LiteralNilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterLiteralNil(s)
	}
}

func (s *LiteralNilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitLiteralNil(s)
	}
}

func (s *LiteralNilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitLiteralNil(s)

	default:
		return t.VisitChildren(s)
	}
}

type GroupContext struct {
	*ExprContext
	left IExprContext
}

func NewGroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupContext {
	var p = new(GroupContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GroupContext) GetLeft() IExprContext { return s.left }

func (s *GroupContext) SetLeft(v IExprContext) { s.left = v }

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) LPAR() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLPAR, 0)
}

func (s *GroupContext) RPAR() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRPAR, 0)
}

func (s *GroupContext) Expr() IExprContext {
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

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (s *GroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *YammmGrammarParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, YammmGrammarParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(338)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case YammmGrammarParserSTRING, YammmGrammarParserREGEXP, YammmGrammarParserINTEGER, YammmGrammarParserFLOAT, YammmGrammarParserBOOLEAN:
		localctx = NewValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(309)

			var _x = p.Literal()

			localctx.(*ValueContext).left = _x
		}

	case YammmGrammarParserLBRACK:
		localctx = NewListContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(310)
			p.Match(YammmGrammarParserLBRACK)
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6629295833815711754) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&31) != 0 {
			{
				p.SetState(311)

				var _x = p.expr(0)

				localctx.(*ListContext)._expr = _x
			}
			localctx.(*ListContext).values = append(localctx.(*ListContext).values, localctx.(*ListContext)._expr)
			p.SetState(316)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(312)
						p.Match(YammmGrammarParserCOMMA)
					}
					{
						p.SetState(313)

						var _x = p.expr(0)

						localctx.(*ListContext)._expr = _x
					}
					localctx.(*ListContext).values = append(localctx.(*ListContext).values, localctx.(*ListContext)._expr)

				}
				p.SetState(318)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
			}
			p.SetState(320)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == YammmGrammarParserCOMMA {
				{
					p.SetState(319)
					p.Match(YammmGrammarParserCOMMA)
				}

			}

		}
		{
			p.SetState(324)
			p.Match(YammmGrammarParserRBRACK)
		}

	case YammmGrammarParserMINUS:
		localctx = NewUminusContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(325)

			var _m = p.Match(YammmGrammarParserMINUS)

			localctx.(*UminusContext).op = _m
		}
		{
			p.SetState(326)

			var _x = p.expr(20)

			localctx.(*UminusContext).right = _x
		}

	case YammmGrammarParserEXCLAMATION:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(327)

			var _m = p.Match(YammmGrammarParserEXCLAMATION)

			localctx.(*NotContext).op = _m
		}
		{
			p.SetState(328)

			var _x = p.expr(16)

			localctx.(*NotContext).right = _x
		}

	case YammmGrammarParserLPAR:
		localctx = NewGroupContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(329)
			p.Match(YammmGrammarParserLPAR)
		}
		{
			p.SetState(330)

			var _x = p.expr(0)

			localctx.(*GroupContext).left = _x
		}
		{
			p.SetState(331)
			p.Match(YammmGrammarParserRPAR)
		}

	case YammmGrammarParserVARIABLE:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(333)

			var _m = p.Match(YammmGrammarParserVARIABLE)

			localctx.(*VariableContext).left = _m
		}

	case YammmGrammarParserT__0, YammmGrammarParserT__1, YammmGrammarParserT__3, YammmGrammarParserT__4, YammmGrammarParserT__5, YammmGrammarParserT__6, YammmGrammarParserT__7, YammmGrammarParserT__8, YammmGrammarParserT__20, YammmGrammarParserT__21, YammmGrammarParserLC_WORD:
		localctx = NewNameContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(334)

			var _x = p.Property_name()

			localctx.(*NameContext).left = _x
		}

	case YammmGrammarParserT__9, YammmGrammarParserT__10, YammmGrammarParserT__11, YammmGrammarParserT__12, YammmGrammarParserT__13, YammmGrammarParserT__14, YammmGrammarParserT__15, YammmGrammarParserT__16, YammmGrammarParserT__17, YammmGrammarParserT__18:
		localctx = NewDatatypeNameContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(335)

			var _x = p.DatatypeKeyword()

			localctx.(*DatatypeNameContext).left = _x
		}

	case YammmGrammarParserUC_WORD:
		localctx = NewRelationNameContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(336)

			var _m = p.Match(YammmGrammarParserUC_WORD)

			localctx.(*RelationNameContext).left = _m
		}

	case YammmGrammarParserUSCORE:
		localctx = NewLiteralNilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(337)
			p.Match(YammmGrammarParserUSCORE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(408)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPeriodContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*PeriodContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(340)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(341)
					p.Match(YammmGrammarParserPERIOD)
				}
				{
					p.SetState(342)

					var _x = p.expr(18)

					localctx.(*PeriodContext).name = _x
				}

			case 2:
				localctx = NewMuldivContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MuldivContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(343)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(344)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MuldivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72057765836619776) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MuldivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(345)

					var _x = p.expr(16)

					localctx.(*MuldivContext).right = _x
				}

			case 3:
				localctx = NewPlusminusContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*PlusminusContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(346)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(347)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*PlusminusContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == YammmGrammarParserPLUS || _la == YammmGrammarParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*PlusminusContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(348)

					var _x = p.expr(15)

					localctx.(*PlusminusContext).right = _x
				}

			case 4:
				localctx = NewCompareContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CompareContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(349)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(350)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8444249301319680) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(351)

					var _x = p.expr(14)

					localctx.(*CompareContext).right = _x
				}

			case 5:
				localctx = NewInContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(352)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(353)

					var _m = p.Match(YammmGrammarParserT__19)

					localctx.(*InContext).op = _m
				}
				{
					p.SetState(354)

					var _x = p.expr(13)

					localctx.(*InContext).right = _x
				}

			case 6:
				localctx = NewMatchContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MatchContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(355)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(356)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MatchContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == YammmGrammarParserMATCH || _la == YammmGrammarParserNOTMATCH) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MatchContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(357)

					var _x = p.expr(12)

					localctx.(*MatchContext).right = _x
				}

			case 7:
				localctx = NewEqualityContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EqualityContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(358)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(359)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == YammmGrammarParserEQUAL || _la == YammmGrammarParserNOTEQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(360)

					var _x = p.expr(11)

					localctx.(*EqualityContext).right = _x
				}

			case 8:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AndContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(361)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(362)

					var _m = p.Match(YammmGrammarParserAND)

					localctx.(*AndContext).op = _m
				}
				{
					p.SetState(363)

					var _x = p.expr(10)

					localctx.(*AndContext).right = _x
				}

			case 9:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OrContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(364)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(365)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OrContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == YammmGrammarParserOR || _la == YammmGrammarParserHAT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OrContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(366)

					var _x = p.expr(9)

					localctx.(*OrContext).right = _x
				}

			case 10:
				localctx = NewAtContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AtContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(367)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(368)
					p.Match(YammmGrammarParserLBRACK)
				}
				p.SetState(380)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6629295833815711754) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&31) != 0 {
					{
						p.SetState(369)

						var _x = p.expr(0)

						localctx.(*AtContext)._expr = _x
					}
					localctx.(*AtContext).right = append(localctx.(*AtContext).right, localctx.(*AtContext)._expr)
					p.SetState(374)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

					for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
						if _alt == 1 {
							{
								p.SetState(370)
								p.Match(YammmGrammarParserCOMMA)
							}
							{
								p.SetState(371)

								var _x = p.expr(0)

								localctx.(*AtContext)._expr = _x
							}
							localctx.(*AtContext).right = append(localctx.(*AtContext).right, localctx.(*AtContext)._expr)

						}
						p.SetState(376)
						p.GetErrorHandler().Sync(p)
						_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
					}
					p.SetState(378)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					if _la == YammmGrammarParserCOMMA {
						{
							p.SetState(377)
							p.Match(YammmGrammarParserCOMMA)
						}

					}

				}
				{
					p.SetState(382)
					p.Match(YammmGrammarParserRBRACK)
				}

			case 11:
				localctx = NewFcallContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*FcallContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(383)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(384)

					var _m = p.Match(YammmGrammarParserARROW)

					localctx.(*FcallContext).op = _m
				}
				{
					p.SetState(385)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*FcallContext).name = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == YammmGrammarParserUC_WORD || _la == YammmGrammarParserLC_WORD) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*FcallContext).name = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(387)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(386)

						var _x = p.Arguments()

						localctx.(*FcallContext).args = _x
					}

				}
				p.SetState(390)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(389)

						var _x = p.Parameters()

						localctx.(*FcallContext).params = _x
					}

				}
				p.SetState(396)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(392)
						p.Match(YammmGrammarParserLBRACE)
					}
					{
						p.SetState(393)

						var _x = p.expr(0)

						localctx.(*FcallContext).body = _x
					}
					{
						p.SetState(394)
						p.Match(YammmGrammarParserRBRACE)
					}

				}

			case 12:
				localctx = NewIfContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*IfContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, YammmGrammarParserRULE_expr)
				p.SetState(398)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(399)

					var _m = p.Match(YammmGrammarParserQMARK)

					localctx.(*IfContext).op = _m
				}
				{
					p.SetState(400)
					p.Match(YammmGrammarParserLBRACE)
				}
				{
					p.SetState(401)

					var _x = p.expr(0)

					localctx.(*IfContext).trueBranch = _x
				}
				p.SetState(404)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == YammmGrammarParserCOLON {
					{
						p.SetState(402)
						p.Match(YammmGrammarParserCOLON)
					}
					{
						p.SetState(403)

						var _x = p.expr(0)

						localctx.(*IfContext).falseBranch = _x
					}

				}
				{
					p.SetState(406)
					p.Match(YammmGrammarParserRBRACE)
				}

			}

		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext())
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetArgs returns the args rule context list.
	GetArgs() []IExprContext

	// SetArgs sets the args rule context list.
	SetArgs([]IExprContext)

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_expr  IExprContext
	args   []IExprContext
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) Get_expr() IExprContext { return s._expr }

func (s *ArgumentsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ArgumentsContext) GetArgs() []IExprContext { return s.args }

func (s *ArgumentsContext) SetArgs(v []IExprContext) { s.args = v }

func (s *ArgumentsContext) LPAR() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserLPAR, 0)
}

func (s *ArgumentsContext) RPAR() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserRPAR, 0)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOMMA, i)
}

func (s *ArgumentsContext) AllExpr() []IExprContext {
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

func (s *ArgumentsContext) Expr(i int) IExprContext {
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

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, YammmGrammarParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
		p.Match(YammmGrammarParserLPAR)
	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6629295833815711754) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&31) != 0 {
		{
			p.SetState(414)

			var _x = p.expr(0)

			localctx.(*ArgumentsContext)._expr = _x
		}
		localctx.(*ArgumentsContext).args = append(localctx.(*ArgumentsContext).args, localctx.(*ArgumentsContext)._expr)
		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(415)
					p.Match(YammmGrammarParserCOMMA)
				}
				{
					p.SetState(416)

					var _x = p.expr(0)

					localctx.(*ArgumentsContext)._expr = _x
				}
				localctx.(*ArgumentsContext).args = append(localctx.(*ArgumentsContext).args, localctx.(*ArgumentsContext)._expr)

			}
			p.SetState(421)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())
		}

	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserCOMMA {
		{
			p.SetState(424)
			p.Match(YammmGrammarParserCOMMA)
		}

	}
	{
		p.SetState(427)
		p.Match(YammmGrammarParserRPAR)
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VARIABLE returns the _VARIABLE token.
	Get_VARIABLE() antlr.Token

	// Set_VARIABLE sets the _VARIABLE token.
	Set_VARIABLE(antlr.Token)

	// GetParams returns the params token list.
	GetParams() []antlr.Token

	// SetParams sets the params token list.
	SetParams([]antlr.Token)

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	_VARIABLE antlr.Token
	params    []antlr.Token
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) Get_VARIABLE() antlr.Token { return s._VARIABLE }

func (s *ParametersContext) Set_VARIABLE(v antlr.Token) { s._VARIABLE = v }

func (s *ParametersContext) GetParams() []antlr.Token { return s.params }

func (s *ParametersContext) SetParams(v []antlr.Token) { s.params = v }

func (s *ParametersContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserPIPE)
}

func (s *ParametersContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserPIPE, i)
}

func (s *ParametersContext) AllVARIABLE() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserVARIABLE)
}

func (s *ParametersContext) VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserVARIABLE, i)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(YammmGrammarParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Parameters() (localctx IParametersContext) {
	this := p
	_ = this

	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, YammmGrammarParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		p.Match(YammmGrammarParserPIPE)
	}
	{
		p.SetState(430)

		var _m = p.Match(YammmGrammarParserVARIABLE)

		localctx.(*ParametersContext)._VARIABLE = _m
	}
	localctx.(*ParametersContext).params = append(localctx.(*ParametersContext).params, localctx.(*ParametersContext)._VARIABLE)
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(431)
				p.Match(YammmGrammarParserCOMMA)
			}
			{
				p.SetState(432)

				var _m = p.Match(YammmGrammarParserVARIABLE)

				localctx.(*ParametersContext)._VARIABLE = _m
			}
			localctx.(*ParametersContext).params = append(localctx.(*ParametersContext).params, localctx.(*ParametersContext)._VARIABLE)

		}
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())
	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == YammmGrammarParserCOMMA {
		{
			p.SetState(438)
			p.Match(YammmGrammarParserCOMMA)
		}

	}
	{
		p.SetState(441)
		p.Match(YammmGrammarParserPIPE)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      antlr.Token
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) GetV() antlr.Token { return s.v }

func (s *LiteralContext) SetV(v antlr.Token) { s.v = v }

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserSTRING, 0)
}

func (s *LiteralContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserBOOLEAN, 0)
}

func (s *LiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserFLOAT, 0)
}

func (s *LiteralContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserINTEGER, 0)
}

func (s *LiteralContext) REGEXP() antlr.TerminalNode {
	return s.GetToken(YammmGrammarParserREGEXP, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, YammmGrammarParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(443)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*LiteralContext).v = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-58)) & ^0x3f) == 0 && ((int64(1)<<(_la-58))&457) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*LiteralContext).v = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILc_keywordContext is an interface to support dynamic dispatch.
type ILc_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLc_keywordContext differentiates from other interfaces.
	IsLc_keywordContext()
}

type Lc_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLc_keywordContext() *Lc_keywordContext {
	var p = new(Lc_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YammmGrammarParserRULE_lc_keyword
	return p
}

func (*Lc_keywordContext) IsLc_keywordContext() {}

func NewLc_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lc_keywordContext {
	var p = new(Lc_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YammmGrammarParserRULE_lc_keyword

	return p
}

func (s *Lc_keywordContext) GetParser() antlr.Parser { return s.parser }
func (s *Lc_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lc_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lc_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.EnterLc_keyword(s)
	}
}

func (s *Lc_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YammmGrammarListener); ok {
		listenerT.ExitLc_keyword(s)
	}
}

func (s *Lc_keywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YammmGrammarVisitor:
		return t.VisitLc_keyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YammmGrammarParser) Lc_keyword() (localctx ILc_keywordContext) {
	this := p
	_ = this

	localctx = NewLc_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, YammmGrammarParserRULE_lc_keyword)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(445)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6292470) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *YammmGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 31:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *YammmGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
