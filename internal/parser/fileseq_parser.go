// Code generated from fileseq.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // fileseq
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

type fileseqParser struct {
	*antlr.BaseParser
}

var FileseqParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fileseqParserInit() {
	staticData := &FileseqParserStaticData
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
		"input", "sequence", "patternOnly", "singleFrame", "plainFile", "directory",
		"basenameChar", "plainBasenameChar", "dirSegment", "basename", "patBasename",
		"frameBasename", "plainBasename", "frameRange", "frameNum", "padding",
		"extension",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 170, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 3, 0, 47, 8, 0, 1, 1, 1, 1, 3, 1, 51, 8, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 60, 8, 1, 1, 1, 5, 1, 63, 8, 1, 10,
		1, 12, 1, 66, 9, 1, 1, 2, 1, 2, 3, 2, 70, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		75, 8, 2, 1, 2, 5, 2, 78, 8, 2, 10, 2, 12, 2, 81, 9, 2, 1, 3, 1, 3, 3,
		3, 85, 8, 3, 1, 3, 1, 3, 4, 3, 89, 8, 3, 11, 3, 12, 3, 90, 1, 4, 1, 4,
		3, 4, 95, 8, 4, 1, 4, 5, 4, 98, 8, 4, 10, 4, 12, 4, 101, 9, 4, 1, 5, 1,
		5, 3, 5, 105, 8, 5, 3, 5, 107, 8, 5, 1, 5, 1, 5, 1, 5, 5, 5, 112, 8, 5,
		10, 5, 12, 5, 115, 9, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 4, 8, 122, 8, 8,
		11, 8, 12, 8, 123, 1, 9, 4, 9, 127, 8, 9, 11, 9, 12, 9, 128, 1, 10, 4,
		10, 132, 8, 10, 11, 10, 12, 10, 133, 1, 11, 4, 11, 137, 8, 11, 11, 11,
		12, 11, 138, 1, 12, 4, 12, 142, 8, 12, 11, 12, 12, 12, 143, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 4, 15, 155, 8, 15, 11,
		15, 12, 15, 156, 3, 15, 159, 8, 15, 1, 16, 1, 16, 1, 16, 3, 16, 164, 8,
		16, 1, 16, 1, 16, 3, 16, 168, 8, 16, 1, 16, 0, 0, 17, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 4, 2, 0, 7, 10, 12, 17,
		3, 0, 5, 6, 8, 9, 12, 17, 2, 0, 8, 10, 13, 13, 1, 0, 5, 6, 182, 0, 46,
		1, 0, 0, 0, 2, 48, 1, 0, 0, 0, 4, 67, 1, 0, 0, 0, 6, 82, 1, 0, 0, 0, 8,
		92, 1, 0, 0, 0, 10, 106, 1, 0, 0, 0, 12, 116, 1, 0, 0, 0, 14, 118, 1, 0,
		0, 0, 16, 121, 1, 0, 0, 0, 18, 126, 1, 0, 0, 0, 20, 131, 1, 0, 0, 0, 22,
		136, 1, 0, 0, 0, 24, 141, 1, 0, 0, 0, 26, 145, 1, 0, 0, 0, 28, 147, 1,
		0, 0, 0, 30, 158, 1, 0, 0, 0, 32, 167, 1, 0, 0, 0, 34, 35, 3, 2, 1, 0,
		35, 36, 5, 0, 0, 1, 36, 47, 1, 0, 0, 0, 37, 38, 3, 4, 2, 0, 38, 39, 5,
		0, 0, 1, 39, 47, 1, 0, 0, 0, 40, 41, 3, 6, 3, 0, 41, 42, 5, 0, 0, 1, 42,
		47, 1, 0, 0, 0, 43, 44, 3, 8, 4, 0, 44, 45, 5, 0, 0, 1, 45, 47, 1, 0, 0,
		0, 46, 34, 1, 0, 0, 0, 46, 37, 1, 0, 0, 0, 46, 40, 1, 0, 0, 0, 46, 43,
		1, 0, 0, 0, 47, 1, 1, 0, 0, 0, 48, 50, 3, 10, 5, 0, 49, 51, 3, 18, 9, 0,
		50, 49, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 3,
		26, 13, 0, 53, 59, 3, 30, 15, 0, 54, 55, 3, 26, 13, 0, 55, 56, 3, 30, 15,
		0, 56, 60, 1, 0, 0, 0, 57, 58, 5, 12, 0, 0, 58, 60, 3, 30, 15, 0, 59, 54,
		1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 64, 1, 0, 0, 0,
		61, 63, 3, 32, 16, 0, 62, 61, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1,
		0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 3, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67,
		69, 3, 10, 5, 0, 68, 70, 3, 20, 10, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0,
		0, 0, 70, 71, 1, 0, 0, 0, 71, 74, 3, 30, 15, 0, 72, 73, 5, 12, 0, 0, 73,
		75, 3, 30, 15, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 79, 1, 0,
		0, 0, 76, 78, 3, 32, 16, 0, 77, 76, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79,
		77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 5, 1, 0, 0, 0, 81, 79, 1, 0, 0,
		0, 82, 84, 3, 10, 5, 0, 83, 85, 3, 22, 11, 0, 84, 83, 1, 0, 0, 0, 84, 85,
		1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 88, 3, 28, 14, 0, 87, 89, 3, 32, 16,
		0, 88, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91,
		1, 0, 0, 0, 91, 7, 1, 0, 0, 0, 92, 94, 3, 10, 5, 0, 93, 95, 3, 24, 12,
		0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 99, 1, 0, 0, 0, 96, 98,
		3, 32, 16, 0, 97, 96, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0,
		0, 99, 100, 1, 0, 0, 0, 100, 9, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 104,
		5, 11, 0, 0, 103, 105, 5, 11, 0, 0, 104, 103, 1, 0, 0, 0, 104, 105, 1,
		0, 0, 0, 105, 107, 1, 0, 0, 0, 106, 102, 1, 0, 0, 0, 106, 107, 1, 0, 0,
		0, 107, 113, 1, 0, 0, 0, 108, 109, 3, 16, 8, 0, 109, 110, 5, 11, 0, 0,
		110, 112, 1, 0, 0, 0, 111, 108, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113,
		111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 11, 1, 0, 0, 0, 115, 113, 1,
		0, 0, 0, 116, 117, 7, 0, 0, 0, 117, 13, 1, 0, 0, 0, 118, 119, 7, 1, 0,
		0, 119, 15, 1, 0, 0, 0, 120, 122, 3, 12, 6, 0, 121, 120, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 17, 1,
		0, 0, 0, 125, 127, 3, 12, 6, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0,
		0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 19, 1, 0, 0, 0, 130,
		132, 3, 12, 6, 0, 131, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 131,
		1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 21, 1, 0, 0, 0, 135, 137, 3, 12,
		6, 0, 136, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0,
		138, 139, 1, 0, 0, 0, 139, 23, 1, 0, 0, 0, 140, 142, 3, 14, 7, 0, 141,
		140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144,
		1, 0, 0, 0, 144, 25, 1, 0, 0, 0, 145, 146, 7, 2, 0, 0, 146, 27, 1, 0, 0,
		0, 147, 148, 5, 10, 0, 0, 148, 29, 1, 0, 0, 0, 149, 159, 5, 1, 0, 0, 150,
		159, 5, 2, 0, 0, 151, 159, 5, 3, 0, 0, 152, 159, 5, 4, 0, 0, 153, 155,
		7, 3, 0, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 154, 1, 0,
		0, 0, 156, 157, 1, 0, 0, 0, 157, 159, 1, 0, 0, 0, 158, 149, 1, 0, 0, 0,
		158, 150, 1, 0, 0, 0, 158, 151, 1, 0, 0, 0, 158, 152, 1, 0, 0, 0, 158,
		154, 1, 0, 0, 0, 159, 31, 1, 0, 0, 0, 160, 163, 5, 7, 0, 0, 161, 162, 5,
		15, 0, 0, 162, 164, 5, 13, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0,
		0, 0, 164, 168, 1, 0, 0, 0, 165, 168, 5, 10, 0, 0, 166, 168, 5, 14, 0,
		0, 167, 160, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168,
		33, 1, 0, 0, 0, 23, 46, 50, 59, 64, 69, 74, 79, 84, 90, 94, 99, 104, 106,
		113, 123, 128, 133, 138, 143, 156, 158, 163, 167,
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

// fileseqParserInit initializes any static state used to implement fileseqParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewfileseqParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FileseqParserInit() {
	staticData := &FileseqParserStaticData
	staticData.once.Do(fileseqParserInit)
}

// NewfileseqParser produces a new parser instance for the optional input antlr.TokenStream.
func NewfileseqParser(input antlr.TokenStream) *fileseqParser {
	FileseqParserInit()
	this := new(fileseqParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FileseqParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "fileseq.g4"

	return this
}

// fileseqParser tokens.
const (
	fileseqParserEOF             = antlr.TokenEOF
	fileseqParserUDIM_ANGLE      = 1
	fileseqParserUDIM_PAREN      = 2
	fileseqParserPRINTF_PAD      = 3
	fileseqParserHOUDINI_PAD     = 4
	fileseqParserHASH            = 5
	fileseqParserAT              = 6
	fileseqParserEXTENSION       = 7
	fileseqParserDOT_FRAME_RANGE = 8
	fileseqParserFRAME_RANGE     = 9
	fileseqParserDOT_NUM         = 10
	fileseqParserSLASH           = 11
	fileseqParserSPECIAL_CHAR    = 12
	fileseqParserNUM             = 13
	fileseqParserWORD            = 14
	fileseqParserDASH            = 15
	fileseqParserWS              = 16
	fileseqParserOTHER_CHAR      = 17
)

// fileseqParser rules.
const (
	fileseqParserRULE_input             = 0
	fileseqParserRULE_sequence          = 1
	fileseqParserRULE_patternOnly       = 2
	fileseqParserRULE_singleFrame       = 3
	fileseqParserRULE_plainFile         = 4
	fileseqParserRULE_directory         = 5
	fileseqParserRULE_basenameChar      = 6
	fileseqParserRULE_plainBasenameChar = 7
	fileseqParserRULE_dirSegment        = 8
	fileseqParserRULE_basename          = 9
	fileseqParserRULE_patBasename       = 10
	fileseqParserRULE_frameBasename     = 11
	fileseqParserRULE_plainBasename     = 12
	fileseqParserRULE_frameRange        = 13
	fileseqParserRULE_frameNum          = 14
	fileseqParserRULE_padding           = 15
	fileseqParserRULE_extension         = 16
)

// IInputContext is an interface to support dynamic dispatch.
type IInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sequence() ISequenceContext
	EOF() antlr.TerminalNode
	PatternOnly() IPatternOnlyContext
	SingleFrame() ISingleFrameContext
	PlainFile() IPlainFileContext

	// IsInputContext differentiates from other interfaces.
	IsInputContext()
}

type InputContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputContext() *InputContext {
	var p = new(InputContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_input
	return p
}

func InitEmptyInputContext(p *InputContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_input
}

func (*InputContext) IsInputContext() {}

func NewInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputContext {
	var p = new(InputContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_input

	return p
}

func (s *InputContext) GetParser() antlr.Parser { return s.parser }

func (s *InputContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *InputContext) EOF() antlr.TerminalNode {
	return s.GetToken(fileseqParserEOF, 0)
}

func (s *InputContext) PatternOnly() IPatternOnlyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternOnlyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternOnlyContext)
}

func (s *InputContext) SingleFrame() ISingleFrameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleFrameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleFrameContext)
}

func (s *InputContext) PlainFile() IPlainFileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPlainFileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPlainFileContext)
}

func (s *InputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitInput(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) Input() (localctx IInputContext) {
	localctx = NewInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, fileseqParserRULE_input)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Sequence()
		}
		{
			p.SetState(35)
			p.Match(fileseqParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.PatternOnly()
		}
		{
			p.SetState(38)
			p.Match(fileseqParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.SingleFrame()
		}
		{
			p.SetState(41)
			p.Match(fileseqParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(43)
			p.PlainFile()
		}
		{
			p.SetState(44)
			p.Match(fileseqParserEOF)
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

// ISequenceContext is an interface to support dynamic dispatch.
type ISequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Directory() IDirectoryContext
	AllFrameRange() []IFrameRangeContext
	FrameRange(i int) IFrameRangeContext
	AllPadding() []IPaddingContext
	Padding(i int) IPaddingContext
	Basename() IBasenameContext
	SPECIAL_CHAR() antlr.TerminalNode
	AllExtension() []IExtensionContext
	Extension(i int) IExtensionContext

	// IsSequenceContext differentiates from other interfaces.
	IsSequenceContext()
}

type SequenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceContext() *SequenceContext {
	var p = new(SequenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_sequence
	return p
}

func InitEmptySequenceContext(p *SequenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_sequence
}

func (*SequenceContext) IsSequenceContext() {}

func NewSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceContext {
	var p = new(SequenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_sequence

	return p
}

func (s *SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceContext) Directory() IDirectoryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectoryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectoryContext)
}

func (s *SequenceContext) AllFrameRange() []IFrameRangeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFrameRangeContext); ok {
			len++
		}
	}

	tst := make([]IFrameRangeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFrameRangeContext); ok {
			tst[i] = t.(IFrameRangeContext)
			i++
		}
	}

	return tst
}

func (s *SequenceContext) FrameRange(i int) IFrameRangeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrameRangeContext); ok {
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

	return t.(IFrameRangeContext)
}

func (s *SequenceContext) AllPadding() []IPaddingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPaddingContext); ok {
			len++
		}
	}

	tst := make([]IPaddingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPaddingContext); ok {
			tst[i] = t.(IPaddingContext)
			i++
		}
	}

	return tst
}

func (s *SequenceContext) Padding(i int) IPaddingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPaddingContext); ok {
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

	return t.(IPaddingContext)
}

func (s *SequenceContext) Basename() IBasenameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasenameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasenameContext)
}

func (s *SequenceContext) SPECIAL_CHAR() antlr.TerminalNode {
	return s.GetToken(fileseqParserSPECIAL_CHAR, 0)
}

func (s *SequenceContext) AllExtension() []IExtensionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtensionContext); ok {
			len++
		}
	}

	tst := make([]IExtensionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtensionContext); ok {
			tst[i] = t.(IExtensionContext)
			i++
		}
	}

	return tst
}

func (s *SequenceContext) Extension(i int) IExtensionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionContext); ok {
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

	return t.(IExtensionContext)
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) Sequence() (localctx ISequenceContext) {
	localctx = NewSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, fileseqParserRULE_sequence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Directory()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(49)
			p.Basename()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(52)
		p.FrameRange()
	}
	{
		p.SetState(53)
		p.Padding()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(54)
			p.FrameRange()
		}
		{
			p.SetState(55)
			p.Padding()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(57)
			p.Match(fileseqParserSPECIAL_CHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Padding()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17536) != 0 {
		{
			p.SetState(61)
			p.Extension()
		}

		p.SetState(66)
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

// IPatternOnlyContext is an interface to support dynamic dispatch.
type IPatternOnlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Directory() IDirectoryContext
	AllPadding() []IPaddingContext
	Padding(i int) IPaddingContext
	PatBasename() IPatBasenameContext
	SPECIAL_CHAR() antlr.TerminalNode
	AllExtension() []IExtensionContext
	Extension(i int) IExtensionContext

	// IsPatternOnlyContext differentiates from other interfaces.
	IsPatternOnlyContext()
}

type PatternOnlyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternOnlyContext() *PatternOnlyContext {
	var p = new(PatternOnlyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_patternOnly
	return p
}

func InitEmptyPatternOnlyContext(p *PatternOnlyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_patternOnly
}

func (*PatternOnlyContext) IsPatternOnlyContext() {}

func NewPatternOnlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternOnlyContext {
	var p = new(PatternOnlyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_patternOnly

	return p
}

func (s *PatternOnlyContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternOnlyContext) Directory() IDirectoryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectoryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectoryContext)
}

func (s *PatternOnlyContext) AllPadding() []IPaddingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPaddingContext); ok {
			len++
		}
	}

	tst := make([]IPaddingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPaddingContext); ok {
			tst[i] = t.(IPaddingContext)
			i++
		}
	}

	return tst
}

func (s *PatternOnlyContext) Padding(i int) IPaddingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPaddingContext); ok {
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

	return t.(IPaddingContext)
}

func (s *PatternOnlyContext) PatBasename() IPatBasenameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatBasenameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatBasenameContext)
}

func (s *PatternOnlyContext) SPECIAL_CHAR() antlr.TerminalNode {
	return s.GetToken(fileseqParserSPECIAL_CHAR, 0)
}

func (s *PatternOnlyContext) AllExtension() []IExtensionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtensionContext); ok {
			len++
		}
	}

	tst := make([]IExtensionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtensionContext); ok {
			tst[i] = t.(IExtensionContext)
			i++
		}
	}

	return tst
}

func (s *PatternOnlyContext) Extension(i int) IExtensionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionContext); ok {
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

	return t.(IExtensionContext)
}

func (s *PatternOnlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternOnlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternOnlyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitPatternOnly(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) PatternOnly() (localctx IPatternOnlyContext) {
	localctx = NewPatternOnlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, fileseqParserRULE_patternOnly)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Directory()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&259968) != 0 {
		{
			p.SetState(68)
			p.PatBasename()
		}

	}
	{
		p.SetState(71)
		p.Padding()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fileseqParserSPECIAL_CHAR {
		{
			p.SetState(72)
			p.Match(fileseqParserSPECIAL_CHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Padding()
		}

	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17536) != 0 {
		{
			p.SetState(76)
			p.Extension()
		}

		p.SetState(81)
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

// ISingleFrameContext is an interface to support dynamic dispatch.
type ISingleFrameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Directory() IDirectoryContext
	FrameNum() IFrameNumContext
	FrameBasename() IFrameBasenameContext
	AllExtension() []IExtensionContext
	Extension(i int) IExtensionContext

	// IsSingleFrameContext differentiates from other interfaces.
	IsSingleFrameContext()
}

type SingleFrameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleFrameContext() *SingleFrameContext {
	var p = new(SingleFrameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_singleFrame
	return p
}

func InitEmptySingleFrameContext(p *SingleFrameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_singleFrame
}

func (*SingleFrameContext) IsSingleFrameContext() {}

func NewSingleFrameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleFrameContext {
	var p = new(SingleFrameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_singleFrame

	return p
}

func (s *SingleFrameContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleFrameContext) Directory() IDirectoryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectoryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectoryContext)
}

func (s *SingleFrameContext) FrameNum() IFrameNumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrameNumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrameNumContext)
}

func (s *SingleFrameContext) FrameBasename() IFrameBasenameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrameBasenameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrameBasenameContext)
}

func (s *SingleFrameContext) AllExtension() []IExtensionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtensionContext); ok {
			len++
		}
	}

	tst := make([]IExtensionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtensionContext); ok {
			tst[i] = t.(IExtensionContext)
			i++
		}
	}

	return tst
}

func (s *SingleFrameContext) Extension(i int) IExtensionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionContext); ok {
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

	return t.(IExtensionContext)
}

func (s *SingleFrameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleFrameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleFrameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitSingleFrame(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) SingleFrame() (localctx ISingleFrameContext) {
	localctx = NewSingleFrameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, fileseqParserRULE_singleFrame)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Directory()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(83)
			p.FrameBasename()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(86)
		p.FrameNum()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17536) != 0) {
		{
			p.SetState(87)
			p.Extension()
		}

		p.SetState(90)
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

// IPlainFileContext is an interface to support dynamic dispatch.
type IPlainFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Directory() IDirectoryContext
	PlainBasename() IPlainBasenameContext
	AllExtension() []IExtensionContext
	Extension(i int) IExtensionContext

	// IsPlainFileContext differentiates from other interfaces.
	IsPlainFileContext()
}

type PlainFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlainFileContext() *PlainFileContext {
	var p = new(PlainFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_plainFile
	return p
}

func InitEmptyPlainFileContext(p *PlainFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_plainFile
}

func (*PlainFileContext) IsPlainFileContext() {}

func NewPlainFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlainFileContext {
	var p = new(PlainFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_plainFile

	return p
}

func (s *PlainFileContext) GetParser() antlr.Parser { return s.parser }

func (s *PlainFileContext) Directory() IDirectoryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectoryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectoryContext)
}

func (s *PlainFileContext) PlainBasename() IPlainBasenameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPlainBasenameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPlainBasenameContext)
}

func (s *PlainFileContext) AllExtension() []IExtensionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtensionContext); ok {
			len++
		}
	}

	tst := make([]IExtensionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtensionContext); ok {
			tst[i] = t.(IExtensionContext)
			i++
		}
	}

	return tst
}

func (s *PlainFileContext) Extension(i int) IExtensionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionContext); ok {
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

	return t.(IExtensionContext)
}

func (s *PlainFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlainFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitPlainFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) PlainFile() (localctx IPlainFileContext) {
	localctx = NewPlainFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, fileseqParserRULE_plainFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Directory()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(93)
			p.PlainBasename()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17536) != 0 {
		{
			p.SetState(96)
			p.Extension()
		}

		p.SetState(101)
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

// IDirectoryContext is an interface to support dynamic dispatch.
type IDirectoryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode
	AllDirSegment() []IDirSegmentContext
	DirSegment(i int) IDirSegmentContext

	// IsDirectoryContext differentiates from other interfaces.
	IsDirectoryContext()
}

type DirectoryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectoryContext() *DirectoryContext {
	var p = new(DirectoryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_directory
	return p
}

func InitEmptyDirectoryContext(p *DirectoryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_directory
}

func (*DirectoryContext) IsDirectoryContext() {}

func NewDirectoryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectoryContext {
	var p = new(DirectoryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_directory

	return p
}

func (s *DirectoryContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectoryContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserSLASH)
}

func (s *DirectoryContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserSLASH, i)
}

func (s *DirectoryContext) AllDirSegment() []IDirSegmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDirSegmentContext); ok {
			len++
		}
	}

	tst := make([]IDirSegmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDirSegmentContext); ok {
			tst[i] = t.(IDirSegmentContext)
			i++
		}
	}

	return tst
}

func (s *DirectoryContext) DirSegment(i int) IDirSegmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirSegmentContext); ok {
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

	return t.(IDirSegmentContext)
}

func (s *DirectoryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectoryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectoryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitDirectory(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) Directory() (localctx IDirectoryContext) {
	localctx = NewDirectoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, fileseqParserRULE_directory)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fileseqParserSLASH {
		{
			p.SetState(102)
			p.Match(fileseqParserSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fileseqParserSLASH {
			{
				p.SetState(103)
				p.Match(fileseqParserSLASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(108)
				p.DirSegment()
			}
			{
				p.SetState(109)
				p.Match(fileseqParserSLASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
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

// IBasenameCharContext is an interface to support dynamic dispatch.
type IBasenameCharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD() antlr.TerminalNode
	NUM() antlr.TerminalNode
	DOT_NUM() antlr.TerminalNode
	DASH() antlr.TerminalNode
	SPECIAL_CHAR() antlr.TerminalNode
	EXTENSION() antlr.TerminalNode
	FRAME_RANGE() antlr.TerminalNode
	DOT_FRAME_RANGE() antlr.TerminalNode
	WS() antlr.TerminalNode
	OTHER_CHAR() antlr.TerminalNode

	// IsBasenameCharContext differentiates from other interfaces.
	IsBasenameCharContext()
}

type BasenameCharContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasenameCharContext() *BasenameCharContext {
	var p = new(BasenameCharContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_basenameChar
	return p
}

func InitEmptyBasenameCharContext(p *BasenameCharContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_basenameChar
}

func (*BasenameCharContext) IsBasenameCharContext() {}

func NewBasenameCharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasenameCharContext {
	var p = new(BasenameCharContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_basenameChar

	return p
}

func (s *BasenameCharContext) GetParser() antlr.Parser { return s.parser }

func (s *BasenameCharContext) WORD() antlr.TerminalNode {
	return s.GetToken(fileseqParserWORD, 0)
}

func (s *BasenameCharContext) NUM() antlr.TerminalNode {
	return s.GetToken(fileseqParserNUM, 0)
}

func (s *BasenameCharContext) DOT_NUM() antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_NUM, 0)
}

func (s *BasenameCharContext) DASH() antlr.TerminalNode {
	return s.GetToken(fileseqParserDASH, 0)
}

func (s *BasenameCharContext) SPECIAL_CHAR() antlr.TerminalNode {
	return s.GetToken(fileseqParserSPECIAL_CHAR, 0)
}

func (s *BasenameCharContext) EXTENSION() antlr.TerminalNode {
	return s.GetToken(fileseqParserEXTENSION, 0)
}

func (s *BasenameCharContext) FRAME_RANGE() antlr.TerminalNode {
	return s.GetToken(fileseqParserFRAME_RANGE, 0)
}

func (s *BasenameCharContext) DOT_FRAME_RANGE() antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_FRAME_RANGE, 0)
}

func (s *BasenameCharContext) WS() antlr.TerminalNode {
	return s.GetToken(fileseqParserWS, 0)
}

func (s *BasenameCharContext) OTHER_CHAR() antlr.TerminalNode {
	return s.GetToken(fileseqParserOTHER_CHAR, 0)
}

func (s *BasenameCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasenameCharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasenameCharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitBasenameChar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) BasenameChar() (localctx IBasenameCharContext) {
	localctx = NewBasenameCharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, fileseqParserRULE_basenameChar)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&259968) != 0) {
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

// IPlainBasenameCharContext is an interface to support dynamic dispatch.
type IPlainBasenameCharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD() antlr.TerminalNode
	NUM() antlr.TerminalNode
	DASH() antlr.TerminalNode
	SPECIAL_CHAR() antlr.TerminalNode
	FRAME_RANGE() antlr.TerminalNode
	DOT_FRAME_RANGE() antlr.TerminalNode
	WS() antlr.TerminalNode
	OTHER_CHAR() antlr.TerminalNode
	HASH() antlr.TerminalNode
	AT() antlr.TerminalNode

	// IsPlainBasenameCharContext differentiates from other interfaces.
	IsPlainBasenameCharContext()
}

type PlainBasenameCharContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlainBasenameCharContext() *PlainBasenameCharContext {
	var p = new(PlainBasenameCharContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_plainBasenameChar
	return p
}

func InitEmptyPlainBasenameCharContext(p *PlainBasenameCharContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_plainBasenameChar
}

func (*PlainBasenameCharContext) IsPlainBasenameCharContext() {}

func NewPlainBasenameCharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlainBasenameCharContext {
	var p = new(PlainBasenameCharContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_plainBasenameChar

	return p
}

func (s *PlainBasenameCharContext) GetParser() antlr.Parser { return s.parser }

func (s *PlainBasenameCharContext) WORD() antlr.TerminalNode {
	return s.GetToken(fileseqParserWORD, 0)
}

func (s *PlainBasenameCharContext) NUM() antlr.TerminalNode {
	return s.GetToken(fileseqParserNUM, 0)
}

func (s *PlainBasenameCharContext) DASH() antlr.TerminalNode {
	return s.GetToken(fileseqParserDASH, 0)
}

func (s *PlainBasenameCharContext) SPECIAL_CHAR() antlr.TerminalNode {
	return s.GetToken(fileseqParserSPECIAL_CHAR, 0)
}

func (s *PlainBasenameCharContext) FRAME_RANGE() antlr.TerminalNode {
	return s.GetToken(fileseqParserFRAME_RANGE, 0)
}

func (s *PlainBasenameCharContext) DOT_FRAME_RANGE() antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_FRAME_RANGE, 0)
}

func (s *PlainBasenameCharContext) WS() antlr.TerminalNode {
	return s.GetToken(fileseqParserWS, 0)
}

func (s *PlainBasenameCharContext) OTHER_CHAR() antlr.TerminalNode {
	return s.GetToken(fileseqParserOTHER_CHAR, 0)
}

func (s *PlainBasenameCharContext) HASH() antlr.TerminalNode {
	return s.GetToken(fileseqParserHASH, 0)
}

func (s *PlainBasenameCharContext) AT() antlr.TerminalNode {
	return s.GetToken(fileseqParserAT, 0)
}

func (s *PlainBasenameCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainBasenameCharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlainBasenameCharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitPlainBasenameChar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) PlainBasenameChar() (localctx IPlainBasenameCharContext) {
	localctx = NewPlainBasenameCharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, fileseqParserRULE_plainBasenameChar)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&258912) != 0) {
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

// IDirSegmentContext is an interface to support dynamic dispatch.
type IDirSegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBasenameChar() []IBasenameCharContext
	BasenameChar(i int) IBasenameCharContext

	// IsDirSegmentContext differentiates from other interfaces.
	IsDirSegmentContext()
}

type DirSegmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirSegmentContext() *DirSegmentContext {
	var p = new(DirSegmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_dirSegment
	return p
}

func InitEmptyDirSegmentContext(p *DirSegmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_dirSegment
}

func (*DirSegmentContext) IsDirSegmentContext() {}

func NewDirSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirSegmentContext {
	var p = new(DirSegmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_dirSegment

	return p
}

func (s *DirSegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *DirSegmentContext) AllBasenameChar() []IBasenameCharContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBasenameCharContext); ok {
			len++
		}
	}

	tst := make([]IBasenameCharContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBasenameCharContext); ok {
			tst[i] = t.(IBasenameCharContext)
			i++
		}
	}

	return tst
}

func (s *DirSegmentContext) BasenameChar(i int) IBasenameCharContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasenameCharContext); ok {
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

	return t.(IBasenameCharContext)
}

func (s *DirSegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirSegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirSegmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitDirSegment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) DirSegment() (localctx IDirSegmentContext) {
	localctx = NewDirSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, fileseqParserRULE_dirSegment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&259968) != 0) {
		{
			p.SetState(120)
			p.BasenameChar()
		}

		p.SetState(123)
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

// IBasenameContext is an interface to support dynamic dispatch.
type IBasenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBasenameChar() []IBasenameCharContext
	BasenameChar(i int) IBasenameCharContext

	// IsBasenameContext differentiates from other interfaces.
	IsBasenameContext()
}

type BasenameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasenameContext() *BasenameContext {
	var p = new(BasenameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_basename
	return p
}

func InitEmptyBasenameContext(p *BasenameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_basename
}

func (*BasenameContext) IsBasenameContext() {}

func NewBasenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasenameContext {
	var p = new(BasenameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_basename

	return p
}

func (s *BasenameContext) GetParser() antlr.Parser { return s.parser }

func (s *BasenameContext) AllBasenameChar() []IBasenameCharContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBasenameCharContext); ok {
			len++
		}
	}

	tst := make([]IBasenameCharContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBasenameCharContext); ok {
			tst[i] = t.(IBasenameCharContext)
			i++
		}
	}

	return tst
}

func (s *BasenameContext) BasenameChar(i int) IBasenameCharContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasenameCharContext); ok {
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

	return t.(IBasenameCharContext)
}

func (s *BasenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasenameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitBasename(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) Basename() (localctx IBasenameContext) {
	localctx = NewBasenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, fileseqParserRULE_basename)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(125)
				p.BasenameChar()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
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

// IPatBasenameContext is an interface to support dynamic dispatch.
type IPatBasenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBasenameChar() []IBasenameCharContext
	BasenameChar(i int) IBasenameCharContext

	// IsPatBasenameContext differentiates from other interfaces.
	IsPatBasenameContext()
}

type PatBasenameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatBasenameContext() *PatBasenameContext {
	var p = new(PatBasenameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_patBasename
	return p
}

func InitEmptyPatBasenameContext(p *PatBasenameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_patBasename
}

func (*PatBasenameContext) IsPatBasenameContext() {}

func NewPatBasenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatBasenameContext {
	var p = new(PatBasenameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_patBasename

	return p
}

func (s *PatBasenameContext) GetParser() antlr.Parser { return s.parser }

func (s *PatBasenameContext) AllBasenameChar() []IBasenameCharContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBasenameCharContext); ok {
			len++
		}
	}

	tst := make([]IBasenameCharContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBasenameCharContext); ok {
			tst[i] = t.(IBasenameCharContext)
			i++
		}
	}

	return tst
}

func (s *PatBasenameContext) BasenameChar(i int) IBasenameCharContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasenameCharContext); ok {
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

	return t.(IBasenameCharContext)
}

func (s *PatBasenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatBasenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatBasenameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitPatBasename(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) PatBasename() (localctx IPatBasenameContext) {
	localctx = NewPatBasenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, fileseqParserRULE_patBasename)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&259968) != 0) {
		{
			p.SetState(130)
			p.BasenameChar()
		}

		p.SetState(133)
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

// IFrameBasenameContext is an interface to support dynamic dispatch.
type IFrameBasenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBasenameChar() []IBasenameCharContext
	BasenameChar(i int) IBasenameCharContext

	// IsFrameBasenameContext differentiates from other interfaces.
	IsFrameBasenameContext()
}

type FrameBasenameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrameBasenameContext() *FrameBasenameContext {
	var p = new(FrameBasenameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_frameBasename
	return p
}

func InitEmptyFrameBasenameContext(p *FrameBasenameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_frameBasename
}

func (*FrameBasenameContext) IsFrameBasenameContext() {}

func NewFrameBasenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FrameBasenameContext {
	var p = new(FrameBasenameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_frameBasename

	return p
}

func (s *FrameBasenameContext) GetParser() antlr.Parser { return s.parser }

func (s *FrameBasenameContext) AllBasenameChar() []IBasenameCharContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBasenameCharContext); ok {
			len++
		}
	}

	tst := make([]IBasenameCharContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBasenameCharContext); ok {
			tst[i] = t.(IBasenameCharContext)
			i++
		}
	}

	return tst
}

func (s *FrameBasenameContext) BasenameChar(i int) IBasenameCharContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasenameCharContext); ok {
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

	return t.(IBasenameCharContext)
}

func (s *FrameBasenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FrameBasenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FrameBasenameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitFrameBasename(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) FrameBasename() (localctx IFrameBasenameContext) {
	localctx = NewFrameBasenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, fileseqParserRULE_frameBasename)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(135)
				p.BasenameChar()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
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

// IPlainBasenameContext is an interface to support dynamic dispatch.
type IPlainBasenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPlainBasenameChar() []IPlainBasenameCharContext
	PlainBasenameChar(i int) IPlainBasenameCharContext

	// IsPlainBasenameContext differentiates from other interfaces.
	IsPlainBasenameContext()
}

type PlainBasenameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlainBasenameContext() *PlainBasenameContext {
	var p = new(PlainBasenameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_plainBasename
	return p
}

func InitEmptyPlainBasenameContext(p *PlainBasenameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_plainBasename
}

func (*PlainBasenameContext) IsPlainBasenameContext() {}

func NewPlainBasenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlainBasenameContext {
	var p = new(PlainBasenameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_plainBasename

	return p
}

func (s *PlainBasenameContext) GetParser() antlr.Parser { return s.parser }

func (s *PlainBasenameContext) AllPlainBasenameChar() []IPlainBasenameCharContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPlainBasenameCharContext); ok {
			len++
		}
	}

	tst := make([]IPlainBasenameCharContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPlainBasenameCharContext); ok {
			tst[i] = t.(IPlainBasenameCharContext)
			i++
		}
	}

	return tst
}

func (s *PlainBasenameContext) PlainBasenameChar(i int) IPlainBasenameCharContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPlainBasenameCharContext); ok {
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

	return t.(IPlainBasenameCharContext)
}

func (s *PlainBasenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainBasenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlainBasenameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitPlainBasename(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) PlainBasename() (localctx IPlainBasenameContext) {
	localctx = NewPlainBasenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, fileseqParserRULE_plainBasename)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(140)
				p.PlainBasenameChar()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
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

// IFrameRangeContext is an interface to support dynamic dispatch.
type IFrameRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT_FRAME_RANGE() antlr.TerminalNode
	FRAME_RANGE() antlr.TerminalNode
	DOT_NUM() antlr.TerminalNode
	NUM() antlr.TerminalNode

	// IsFrameRangeContext differentiates from other interfaces.
	IsFrameRangeContext()
}

type FrameRangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrameRangeContext() *FrameRangeContext {
	var p = new(FrameRangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_frameRange
	return p
}

func InitEmptyFrameRangeContext(p *FrameRangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_frameRange
}

func (*FrameRangeContext) IsFrameRangeContext() {}

func NewFrameRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FrameRangeContext {
	var p = new(FrameRangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_frameRange

	return p
}

func (s *FrameRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *FrameRangeContext) DOT_FRAME_RANGE() antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_FRAME_RANGE, 0)
}

func (s *FrameRangeContext) FRAME_RANGE() antlr.TerminalNode {
	return s.GetToken(fileseqParserFRAME_RANGE, 0)
}

func (s *FrameRangeContext) DOT_NUM() antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_NUM, 0)
}

func (s *FrameRangeContext) NUM() antlr.TerminalNode {
	return s.GetToken(fileseqParserNUM, 0)
}

func (s *FrameRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FrameRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FrameRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitFrameRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) FrameRange() (localctx IFrameRangeContext) {
	localctx = NewFrameRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, fileseqParserRULE_frameRange)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9984) != 0) {
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

// IFrameNumContext is an interface to support dynamic dispatch.
type IFrameNumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT_NUM() antlr.TerminalNode

	// IsFrameNumContext differentiates from other interfaces.
	IsFrameNumContext()
}

type FrameNumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrameNumContext() *FrameNumContext {
	var p = new(FrameNumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_frameNum
	return p
}

func InitEmptyFrameNumContext(p *FrameNumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_frameNum
}

func (*FrameNumContext) IsFrameNumContext() {}

func NewFrameNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FrameNumContext {
	var p = new(FrameNumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_frameNum

	return p
}

func (s *FrameNumContext) GetParser() antlr.Parser { return s.parser }

func (s *FrameNumContext) DOT_NUM() antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_NUM, 0)
}

func (s *FrameNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FrameNumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FrameNumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitFrameNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) FrameNum() (localctx IFrameNumContext) {
	localctx = NewFrameNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, fileseqParserRULE_frameNum)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(fileseqParserDOT_NUM)
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

// IPaddingContext is an interface to support dynamic dispatch.
type IPaddingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UDIM_ANGLE() antlr.TerminalNode
	UDIM_PAREN() antlr.TerminalNode
	PRINTF_PAD() antlr.TerminalNode
	HOUDINI_PAD() antlr.TerminalNode
	AllHASH() []antlr.TerminalNode
	HASH(i int) antlr.TerminalNode
	AllAT() []antlr.TerminalNode
	AT(i int) antlr.TerminalNode

	// IsPaddingContext differentiates from other interfaces.
	IsPaddingContext()
}

type PaddingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPaddingContext() *PaddingContext {
	var p = new(PaddingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_padding
	return p
}

func InitEmptyPaddingContext(p *PaddingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_padding
}

func (*PaddingContext) IsPaddingContext() {}

func NewPaddingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PaddingContext {
	var p = new(PaddingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_padding

	return p
}

func (s *PaddingContext) GetParser() antlr.Parser { return s.parser }

func (s *PaddingContext) UDIM_ANGLE() antlr.TerminalNode {
	return s.GetToken(fileseqParserUDIM_ANGLE, 0)
}

func (s *PaddingContext) UDIM_PAREN() antlr.TerminalNode {
	return s.GetToken(fileseqParserUDIM_PAREN, 0)
}

func (s *PaddingContext) PRINTF_PAD() antlr.TerminalNode {
	return s.GetToken(fileseqParserPRINTF_PAD, 0)
}

func (s *PaddingContext) HOUDINI_PAD() antlr.TerminalNode {
	return s.GetToken(fileseqParserHOUDINI_PAD, 0)
}

func (s *PaddingContext) AllHASH() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserHASH)
}

func (s *PaddingContext) HASH(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserHASH, i)
}

func (s *PaddingContext) AllAT() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserAT)
}

func (s *PaddingContext) AT(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserAT, i)
}

func (s *PaddingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PaddingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PaddingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitPadding(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) Padding() (localctx IPaddingContext) {
	localctx = NewPaddingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, fileseqParserRULE_padding)
	var _la int

	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fileseqParserUDIM_ANGLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Match(fileseqParserUDIM_ANGLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fileseqParserUDIM_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Match(fileseqParserUDIM_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fileseqParserPRINTF_PAD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.Match(fileseqParserPRINTF_PAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fileseqParserHOUDINI_PAD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(152)
			p.Match(fileseqParserHOUDINI_PAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fileseqParserHASH, fileseqParserAT:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == fileseqParserHASH || _la == fileseqParserAT {
			{
				p.SetState(153)
				_la = p.GetTokenStream().LA(1)

				if !(_la == fileseqParserHASH || _la == fileseqParserAT) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IExtensionContext is an interface to support dynamic dispatch.
type IExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTENSION() antlr.TerminalNode
	DASH() antlr.TerminalNode
	NUM() antlr.TerminalNode
	DOT_NUM() antlr.TerminalNode
	WORD() antlr.TerminalNode

	// IsExtensionContext differentiates from other interfaces.
	IsExtensionContext()
}

type ExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionContext() *ExtensionContext {
	var p = new(ExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_extension
	return p
}

func InitEmptyExtensionContext(p *ExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_extension
}

func (*ExtensionContext) IsExtensionContext() {}

func NewExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionContext {
	var p = new(ExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_extension

	return p
}

func (s *ExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionContext) EXTENSION() antlr.TerminalNode {
	return s.GetToken(fileseqParserEXTENSION, 0)
}

func (s *ExtensionContext) DASH() antlr.TerminalNode {
	return s.GetToken(fileseqParserDASH, 0)
}

func (s *ExtensionContext) NUM() antlr.TerminalNode {
	return s.GetToken(fileseqParserNUM, 0)
}

func (s *ExtensionContext) DOT_NUM() antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_NUM, 0)
}

func (s *ExtensionContext) WORD() antlr.TerminalNode {
	return s.GetToken(fileseqParserWORD, 0)
}

func (s *ExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) Extension() (localctx IExtensionContext) {
	localctx = NewExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, fileseqParserRULE_extension)
	var _la int

	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fileseqParserEXTENSION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Match(fileseqParserEXTENSION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == fileseqParserDASH {
			{
				p.SetState(161)
				p.Match(fileseqParserDASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(162)
				p.Match(fileseqParserNUM)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case fileseqParserDOT_NUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(fileseqParserDOT_NUM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fileseqParserWORD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.Match(fileseqParserWORD)
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
