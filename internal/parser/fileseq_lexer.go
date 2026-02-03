// Code generated from fileseq.g4 by ANTLR 4.13.1. DO NOT EDIT.

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

type fileseqLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var FileseqLexerLexerStaticData struct {
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

func fileseqlexerLexerInit() {
	staticData := &FileseqLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'<UDIM>'", "'%(UDIM)d'", "", "", "'#'", "'@'", "", "", "", "",
		"", "", "", "", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "UDIM_ANGLE", "UDIM_PAREN", "PRINTF_PAD", "HOUDINI_PAD", "HASH",
		"AT", "EXTENSION", "DOT_FRAME_RANGE", "FRAME_RANGE", "DOT_NUM", "SLASH",
		"SPECIAL_CHAR", "NUM", "WORD", "DASH", "WS", "OTHER_CHAR",
	}
	staticData.RuleNames = []string{
		"UDIM_ANGLE", "UDIM_PAREN", "PRINTF_PAD", "HOUDINI_PAD", "HASH", "AT",
		"EXTENSION", "DOT_FRAME_RANGE", "FRAME_RANGE", "DOT_NUM", "SLASH", "SPECIAL_CHAR",
		"NUM", "WORD", "DASH", "WS", "OTHER_CHAR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 156, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 54, 8, 2,
		10, 2, 12, 2, 57, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 65, 8,
		3, 10, 3, 12, 3, 68, 9, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5,
		6, 77, 8, 6, 10, 6, 12, 6, 80, 9, 6, 1, 6, 1, 6, 5, 6, 84, 8, 6, 10, 6,
		12, 6, 87, 9, 6, 3, 6, 89, 8, 6, 1, 7, 1, 7, 3, 7, 93, 8, 7, 1, 7, 4, 7,
		96, 8, 7, 11, 7, 12, 7, 97, 1, 7, 1, 7, 5, 7, 102, 8, 7, 10, 7, 12, 7,
		105, 9, 7, 1, 8, 3, 8, 108, 8, 8, 1, 8, 4, 8, 111, 8, 8, 11, 8, 12, 8,
		112, 1, 8, 1, 8, 5, 8, 117, 8, 8, 10, 8, 12, 8, 120, 9, 8, 1, 9, 1, 9,
		3, 9, 124, 8, 9, 1, 9, 4, 9, 127, 8, 9, 11, 9, 12, 9, 128, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 4, 12, 136, 8, 12, 11, 12, 12, 12, 137, 1, 13, 4,
		13, 141, 8, 13, 11, 13, 12, 13, 142, 1, 14, 1, 14, 1, 15, 4, 15, 148, 8,
		15, 11, 15, 12, 15, 149, 1, 16, 4, 16, 153, 8, 16, 11, 16, 12, 16, 154,
		0, 0, 17, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 1, 0, 10, 1,
		0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 2, 0, 65, 90, 97, 122, 4, 0,
		48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 44, 45, 58, 58, 3, 0, 44, 45, 48,
		58, 120, 121, 2, 0, 47, 47, 92, 92, 3, 0, 44, 44, 46, 46, 58, 58, 3, 0,
		9, 10, 13, 13, 32, 32, 11, 0, 9, 10, 13, 13, 32, 32, 35, 35, 44, 58, 60,
		60, 62, 62, 64, 90, 92, 92, 95, 95, 97, 122, 172, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 1, 35, 1, 0, 0, 0, 3, 42, 1, 0, 0, 0, 5, 51, 1, 0, 0, 0, 7, 60, 1,
		0, 0, 0, 9, 69, 1, 0, 0, 0, 11, 71, 1, 0, 0, 0, 13, 73, 1, 0, 0, 0, 15,
		90, 1, 0, 0, 0, 17, 107, 1, 0, 0, 0, 19, 121, 1, 0, 0, 0, 21, 130, 1, 0,
		0, 0, 23, 132, 1, 0, 0, 0, 25, 135, 1, 0, 0, 0, 27, 140, 1, 0, 0, 0, 29,
		144, 1, 0, 0, 0, 31, 147, 1, 0, 0, 0, 33, 152, 1, 0, 0, 0, 35, 36, 5, 60,
		0, 0, 36, 37, 5, 85, 0, 0, 37, 38, 5, 68, 0, 0, 38, 39, 5, 73, 0, 0, 39,
		40, 5, 77, 0, 0, 40, 41, 5, 62, 0, 0, 41, 2, 1, 0, 0, 0, 42, 43, 5, 37,
		0, 0, 43, 44, 5, 40, 0, 0, 44, 45, 5, 85, 0, 0, 45, 46, 5, 68, 0, 0, 46,
		47, 5, 73, 0, 0, 47, 48, 5, 77, 0, 0, 48, 49, 5, 41, 0, 0, 49, 50, 5, 100,
		0, 0, 50, 4, 1, 0, 0, 0, 51, 55, 5, 37, 0, 0, 52, 54, 7, 0, 0, 0, 53, 52,
		1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0,
		56, 58, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 59, 5, 100, 0, 0, 59, 6, 1,
		0, 0, 0, 60, 61, 5, 36, 0, 0, 61, 62, 5, 70, 0, 0, 62, 66, 1, 0, 0, 0,
		63, 65, 7, 0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1,
		0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 8, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69,
		70, 5, 35, 0, 0, 70, 10, 1, 0, 0, 0, 71, 72, 5, 64, 0, 0, 72, 12, 1, 0,
		0, 0, 73, 88, 5, 46, 0, 0, 74, 89, 7, 1, 0, 0, 75, 77, 7, 0, 0, 0, 76,
		75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0,
		0, 79, 81, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 85, 7, 2, 0, 0, 82, 84,
		7, 3, 0, 0, 83, 82, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0,
		85, 86, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 74, 1,
		0, 0, 0, 88, 78, 1, 0, 0, 0, 89, 14, 1, 0, 0, 0, 90, 92, 5, 46, 0, 0, 91,
		93, 5, 45, 0, 0, 92, 91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0,
		0, 0, 94, 96, 7, 0, 0, 0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95,
		1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 103, 7, 4, 0, 0,
		100, 102, 7, 5, 0, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 16, 1, 0, 0, 0, 105, 103, 1,
		0, 0, 0, 106, 108, 5, 45, 0, 0, 107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0,
		0, 108, 110, 1, 0, 0, 0, 109, 111, 7, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111,
		112, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114,
		1, 0, 0, 0, 114, 118, 7, 4, 0, 0, 115, 117, 7, 5, 0, 0, 116, 115, 1, 0,
		0, 0, 117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0,
		119, 18, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 123, 5, 46, 0, 0, 122,
		124, 5, 45, 0, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126,
		1, 0, 0, 0, 125, 127, 7, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0,
		0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 20, 1, 0, 0, 0,
		130, 131, 7, 6, 0, 0, 131, 22, 1, 0, 0, 0, 132, 133, 7, 7, 0, 0, 133, 24,
		1, 0, 0, 0, 134, 136, 7, 0, 0, 0, 135, 134, 1, 0, 0, 0, 136, 137, 1, 0,
		0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 26, 1, 0, 0, 0,
		139, 141, 7, 1, 0, 0, 140, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142,
		140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 28, 1, 0, 0, 0, 144, 145, 5,
		45, 0, 0, 145, 30, 1, 0, 0, 0, 146, 148, 7, 8, 0, 0, 147, 146, 1, 0, 0,
		0, 148, 149, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150,
		32, 1, 0, 0, 0, 151, 153, 8, 9, 0, 0, 152, 151, 1, 0, 0, 0, 153, 154, 1,
		0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 34, 1, 0, 0,
		0, 18, 0, 55, 66, 78, 85, 88, 92, 97, 103, 107, 112, 118, 123, 128, 137,
		142, 149, 154, 0,
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

// fileseqLexerInit initializes any static state used to implement fileseqLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewfileseqLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FileseqLexerInit() {
	staticData := &FileseqLexerLexerStaticData
	staticData.once.Do(fileseqlexerLexerInit)
}

// NewfileseqLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewfileseqLexer(input antlr.CharStream) *fileseqLexer {
	FileseqLexerInit()
	l := new(fileseqLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &FileseqLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "fileseq.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// fileseqLexer tokens.
const (
	fileseqLexerUDIM_ANGLE      = 1
	fileseqLexerUDIM_PAREN      = 2
	fileseqLexerPRINTF_PAD      = 3
	fileseqLexerHOUDINI_PAD     = 4
	fileseqLexerHASH            = 5
	fileseqLexerAT              = 6
	fileseqLexerEXTENSION       = 7
	fileseqLexerDOT_FRAME_RANGE = 8
	fileseqLexerFRAME_RANGE     = 9
	fileseqLexerDOT_NUM         = 10
	fileseqLexerSLASH           = 11
	fileseqLexerSPECIAL_CHAR    = 12
	fileseqLexerNUM             = 13
	fileseqLexerWORD            = 14
	fileseqLexerDASH            = 15
	fileseqLexerWS              = 16
	fileseqLexerOTHER_CHAR      = 17
)
