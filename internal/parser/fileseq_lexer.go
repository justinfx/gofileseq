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
		"SPECIAL_CHAR", "NUM", "WORD", "DASH", "WS",
	}
	staticData.RuleNames = []string{
		"UDIM_ANGLE", "UDIM_PAREN", "PRINTF_PAD", "HOUDINI_PAD", "HASH", "AT",
		"EXTENSION", "DOT_FRAME_RANGE", "FRAME_RANGE", "DOT_NUM", "SLASH", "SPECIAL_CHAR",
		"NUM", "WORD", "DASH", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 151, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 52, 8, 2, 10, 2, 12, 2,
		55, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 63, 8, 3, 10, 3, 12,
		3, 66, 9, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 75, 8, 6,
		10, 6, 12, 6, 78, 9, 6, 1, 6, 1, 6, 5, 6, 82, 8, 6, 10, 6, 12, 6, 85, 9,
		6, 3, 6, 87, 8, 6, 1, 7, 1, 7, 3, 7, 91, 8, 7, 1, 7, 4, 7, 94, 8, 7, 11,
		7, 12, 7, 95, 1, 7, 1, 7, 5, 7, 100, 8, 7, 10, 7, 12, 7, 103, 9, 7, 1,
		8, 3, 8, 106, 8, 8, 1, 8, 4, 8, 109, 8, 8, 11, 8, 12, 8, 110, 1, 8, 1,
		8, 5, 8, 115, 8, 8, 10, 8, 12, 8, 118, 9, 8, 1, 9, 1, 9, 3, 9, 122, 8,
		9, 1, 9, 4, 9, 125, 8, 9, 11, 9, 12, 9, 126, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 12, 4, 12, 134, 8, 12, 11, 12, 12, 12, 135, 1, 13, 4, 13, 139, 8, 13,
		11, 13, 12, 13, 140, 1, 14, 1, 14, 1, 15, 4, 15, 146, 8, 15, 11, 15, 12,
		15, 147, 1, 15, 1, 15, 0, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		1, 0, 9, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 2, 0, 65, 90, 97,
		122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 44, 45, 58, 58, 3, 0,
		44, 45, 48, 58, 120, 121, 2, 0, 47, 47, 92, 92, 3, 0, 44, 44, 46, 46, 58,
		58, 3, 0, 9, 10, 13, 13, 32, 32, 166, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 1, 33, 1, 0, 0, 0, 3, 40,
		1, 0, 0, 0, 5, 49, 1, 0, 0, 0, 7, 58, 1, 0, 0, 0, 9, 67, 1, 0, 0, 0, 11,
		69, 1, 0, 0, 0, 13, 71, 1, 0, 0, 0, 15, 88, 1, 0, 0, 0, 17, 105, 1, 0,
		0, 0, 19, 119, 1, 0, 0, 0, 21, 128, 1, 0, 0, 0, 23, 130, 1, 0, 0, 0, 25,
		133, 1, 0, 0, 0, 27, 138, 1, 0, 0, 0, 29, 142, 1, 0, 0, 0, 31, 145, 1,
		0, 0, 0, 33, 34, 5, 60, 0, 0, 34, 35, 5, 85, 0, 0, 35, 36, 5, 68, 0, 0,
		36, 37, 5, 73, 0, 0, 37, 38, 5, 77, 0, 0, 38, 39, 5, 62, 0, 0, 39, 2, 1,
		0, 0, 0, 40, 41, 5, 37, 0, 0, 41, 42, 5, 40, 0, 0, 42, 43, 5, 85, 0, 0,
		43, 44, 5, 68, 0, 0, 44, 45, 5, 73, 0, 0, 45, 46, 5, 77, 0, 0, 46, 47,
		5, 41, 0, 0, 47, 48, 5, 100, 0, 0, 48, 4, 1, 0, 0, 0, 49, 53, 5, 37, 0,
		0, 50, 52, 7, 0, 0, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51,
		1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0,
		56, 57, 5, 100, 0, 0, 57, 6, 1, 0, 0, 0, 58, 59, 5, 36, 0, 0, 59, 60, 5,
		70, 0, 0, 60, 64, 1, 0, 0, 0, 61, 63, 7, 0, 0, 0, 62, 61, 1, 0, 0, 0, 63,
		66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 8, 1, 0, 0,
		0, 66, 64, 1, 0, 0, 0, 67, 68, 5, 35, 0, 0, 68, 10, 1, 0, 0, 0, 69, 70,
		5, 64, 0, 0, 70, 12, 1, 0, 0, 0, 71, 86, 5, 46, 0, 0, 72, 87, 7, 1, 0,
		0, 73, 75, 7, 0, 0, 0, 74, 73, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74,
		1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 79, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0,
		79, 83, 7, 2, 0, 0, 80, 82, 7, 3, 0, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1,
		0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85,
		83, 1, 0, 0, 0, 86, 72, 1, 0, 0, 0, 86, 76, 1, 0, 0, 0, 87, 14, 1, 0, 0,
		0, 88, 90, 5, 46, 0, 0, 89, 91, 5, 45, 0, 0, 90, 89, 1, 0, 0, 0, 90, 91,
		1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 94, 7, 0, 0, 0, 93, 92, 1, 0, 0, 0,
		94, 95, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1,
		0, 0, 0, 97, 101, 7, 4, 0, 0, 98, 100, 7, 5, 0, 0, 99, 98, 1, 0, 0, 0,
		100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 16,
		1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 106, 5, 45, 0, 0, 105, 104, 1, 0,
		0, 0, 105, 106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107, 109, 7, 0, 0, 0,
		108, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110,
		111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 116, 7, 4, 0, 0, 113, 115,
		7, 5, 0, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0,
		0, 0, 116, 117, 1, 0, 0, 0, 117, 18, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0,
		119, 121, 5, 46, 0, 0, 120, 122, 5, 45, 0, 0, 121, 120, 1, 0, 0, 0, 121,
		122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0, 123, 125, 7, 0, 0, 0, 124, 123,
		1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0,
		0, 0, 127, 20, 1, 0, 0, 0, 128, 129, 7, 6, 0, 0, 129, 22, 1, 0, 0, 0, 130,
		131, 7, 7, 0, 0, 131, 24, 1, 0, 0, 0, 132, 134, 7, 0, 0, 0, 133, 132, 1,
		0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0,
		0, 136, 26, 1, 0, 0, 0, 137, 139, 7, 1, 0, 0, 138, 137, 1, 0, 0, 0, 139,
		140, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 28, 1,
		0, 0, 0, 142, 143, 5, 45, 0, 0, 143, 30, 1, 0, 0, 0, 144, 146, 7, 8, 0,
		0, 145, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147,
		148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 6, 15, 0, 0, 150, 32,
		1, 0, 0, 0, 17, 0, 53, 64, 76, 83, 86, 90, 95, 101, 105, 110, 116, 121,
		126, 135, 140, 147, 1, 6, 0, 0,
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
)
