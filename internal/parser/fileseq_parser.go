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
		"SPECIAL_CHAR", "NUM", "WORD", "DASH", "WS",
	}
	staticData.RuleNames = []string{
		"input", "sequence", "patternOnly", "singleFrame", "plainFile", "directory",
		"dirSegment", "sequenceBasename", "patternBasename", "singleFrameBasename",
		"plainBasename", "frameRange", "frameNum", "padding", "extension",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 139, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 43, 8,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 50, 8, 1, 11, 1, 12, 1, 51, 1, 2,
		1, 2, 1, 2, 1, 2, 4, 2, 58, 8, 2, 11, 2, 12, 2, 59, 1, 3, 1, 3, 1, 3, 1,
		3, 5, 3, 66, 8, 3, 10, 3, 12, 3, 69, 9, 3, 1, 4, 1, 4, 3, 4, 73, 8, 4,
		1, 4, 5, 4, 76, 8, 4, 10, 4, 12, 4, 79, 9, 4, 1, 5, 3, 5, 82, 8, 5, 1,
		5, 1, 5, 1, 5, 5, 5, 87, 8, 5, 10, 5, 12, 5, 90, 9, 5, 1, 6, 4, 6, 93,
		8, 6, 11, 6, 12, 6, 94, 1, 7, 4, 7, 98, 8, 7, 11, 7, 12, 7, 99, 1, 8, 4,
		8, 103, 8, 8, 11, 8, 12, 8, 104, 1, 9, 4, 9, 108, 8, 9, 11, 9, 12, 9, 109,
		1, 10, 4, 10, 113, 8, 10, 11, 10, 12, 10, 114, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 4, 13, 126, 8, 13, 11, 13, 12, 13,
		127, 1, 13, 4, 13, 131, 8, 13, 11, 13, 12, 13, 132, 3, 13, 135, 8, 13,
		1, 14, 1, 14, 1, 14, 0, 0, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 0, 4, 2, 0, 8, 10, 12, 15, 3, 0, 7, 7, 10, 10, 12, 15, 2, 0,
		10, 10, 12, 15, 2, 0, 8, 10, 13, 13, 145, 0, 42, 1, 0, 0, 0, 2, 44, 1,
		0, 0, 0, 4, 53, 1, 0, 0, 0, 6, 61, 1, 0, 0, 0, 8, 70, 1, 0, 0, 0, 10, 81,
		1, 0, 0, 0, 12, 92, 1, 0, 0, 0, 14, 97, 1, 0, 0, 0, 16, 102, 1, 0, 0, 0,
		18, 107, 1, 0, 0, 0, 20, 112, 1, 0, 0, 0, 22, 116, 1, 0, 0, 0, 24, 118,
		1, 0, 0, 0, 26, 134, 1, 0, 0, 0, 28, 136, 1, 0, 0, 0, 30, 31, 3, 2, 1,
		0, 31, 32, 5, 0, 0, 1, 32, 43, 1, 0, 0, 0, 33, 34, 3, 4, 2, 0, 34, 35,
		5, 0, 0, 1, 35, 43, 1, 0, 0, 0, 36, 37, 3, 6, 3, 0, 37, 38, 5, 0, 0, 1,
		38, 43, 1, 0, 0, 0, 39, 40, 3, 8, 4, 0, 40, 41, 5, 0, 0, 1, 41, 43, 1,
		0, 0, 0, 42, 30, 1, 0, 0, 0, 42, 33, 1, 0, 0, 0, 42, 36, 1, 0, 0, 0, 42,
		39, 1, 0, 0, 0, 43, 1, 1, 0, 0, 0, 44, 45, 3, 10, 5, 0, 45, 46, 3, 14,
		7, 0, 46, 47, 3, 22, 11, 0, 47, 49, 3, 26, 13, 0, 48, 50, 3, 28, 14, 0,
		49, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1,
		0, 0, 0, 52, 3, 1, 0, 0, 0, 53, 54, 3, 10, 5, 0, 54, 55, 3, 16, 8, 0, 55,
		57, 3, 26, 13, 0, 56, 58, 3, 28, 14, 0, 57, 56, 1, 0, 0, 0, 58, 59, 1,
		0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 5, 1, 0, 0, 0, 61,
		62, 3, 10, 5, 0, 62, 63, 3, 18, 9, 0, 63, 67, 3, 24, 12, 0, 64, 66, 3,
		28, 14, 0, 65, 64, 1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0,
		67, 68, 1, 0, 0, 0, 68, 7, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 72, 3, 10,
		5, 0, 71, 73, 3, 20, 10, 0, 72, 71, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73,
		77, 1, 0, 0, 0, 74, 76, 3, 28, 14, 0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0,
		0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 9, 1, 0, 0, 0, 79, 77,
		1, 0, 0, 0, 80, 82, 5, 11, 0, 0, 81, 80, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0,
		82, 88, 1, 0, 0, 0, 83, 84, 3, 12, 6, 0, 84, 85, 5, 11, 0, 0, 85, 87, 1,
		0, 0, 0, 86, 83, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88,
		89, 1, 0, 0, 0, 89, 11, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 93, 7, 0, 0,
		0, 92, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 95, 13, 1, 0, 0, 0, 96, 98, 7, 1, 0, 0, 97, 96, 1, 0, 0, 0,
		98, 99, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 15, 1,
		0, 0, 0, 101, 103, 7, 1, 0, 0, 102, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0,
		0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 17, 1, 0, 0, 0, 106,
		108, 7, 1, 0, 0, 107, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 107,
		1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 19, 1, 0, 0, 0, 111, 113, 7, 2,
		0, 0, 112, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0,
		114, 115, 1, 0, 0, 0, 115, 21, 1, 0, 0, 0, 116, 117, 7, 3, 0, 0, 117, 23,
		1, 0, 0, 0, 118, 119, 5, 10, 0, 0, 119, 25, 1, 0, 0, 0, 120, 135, 5, 1,
		0, 0, 121, 135, 5, 2, 0, 0, 122, 135, 5, 3, 0, 0, 123, 135, 5, 4, 0, 0,
		124, 126, 5, 5, 0, 0, 125, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127,
		125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 135, 1, 0, 0, 0, 129, 131,
		5, 6, 0, 0, 130, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 130, 1, 0,
		0, 0, 132, 133, 1, 0, 0, 0, 133, 135, 1, 0, 0, 0, 134, 120, 1, 0, 0, 0,
		134, 121, 1, 0, 0, 0, 134, 122, 1, 0, 0, 0, 134, 123, 1, 0, 0, 0, 134,
		125, 1, 0, 0, 0, 134, 130, 1, 0, 0, 0, 135, 27, 1, 0, 0, 0, 136, 137, 5,
		7, 0, 0, 137, 29, 1, 0, 0, 0, 16, 42, 51, 59, 67, 72, 77, 81, 88, 94, 99,
		104, 109, 114, 127, 132, 134,
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
)

// fileseqParser rules.
const (
	fileseqParserRULE_input               = 0
	fileseqParserRULE_sequence            = 1
	fileseqParserRULE_patternOnly         = 2
	fileseqParserRULE_singleFrame         = 3
	fileseqParserRULE_plainFile           = 4
	fileseqParserRULE_directory           = 5
	fileseqParserRULE_dirSegment          = 6
	fileseqParserRULE_sequenceBasename    = 7
	fileseqParserRULE_patternBasename     = 8
	fileseqParserRULE_singleFrameBasename = 9
	fileseqParserRULE_plainBasename       = 10
	fileseqParserRULE_frameRange          = 11
	fileseqParserRULE_frameNum            = 12
	fileseqParserRULE_padding             = 13
	fileseqParserRULE_extension           = 14
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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Sequence()
		}
		{
			p.SetState(31)
			p.Match(fileseqParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.PatternOnly()
		}
		{
			p.SetState(34)
			p.Match(fileseqParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.SingleFrame()
		}
		{
			p.SetState(37)
			p.Match(fileseqParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.PlainFile()
		}
		{
			p.SetState(40)
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
	SequenceBasename() ISequenceBasenameContext
	FrameRange() IFrameRangeContext
	Padding() IPaddingContext
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

func (s *SequenceContext) SequenceBasename() ISequenceBasenameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceBasenameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceBasenameContext)
}

func (s *SequenceContext) FrameRange() IFrameRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrameRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrameRangeContext)
}

func (s *SequenceContext) Padding() IPaddingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPaddingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPaddingContext)
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
		p.SetState(44)
		p.Directory()
	}
	{
		p.SetState(45)
		p.SequenceBasename()
	}
	{
		p.SetState(46)
		p.FrameRange()
	}
	{
		p.SetState(47)
		p.Padding()
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == fileseqParserEXTENSION {
		{
			p.SetState(48)
			p.Extension()
		}

		p.SetState(51)
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
	PatternBasename() IPatternBasenameContext
	Padding() IPaddingContext
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

func (s *PatternOnlyContext) PatternBasename() IPatternBasenameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternBasenameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternBasenameContext)
}

func (s *PatternOnlyContext) Padding() IPaddingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPaddingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPaddingContext)
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
		p.SetState(53)
		p.Directory()
	}
	{
		p.SetState(54)
		p.PatternBasename()
	}
	{
		p.SetState(55)
		p.Padding()
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == fileseqParserEXTENSION {
		{
			p.SetState(56)
			p.Extension()
		}

		p.SetState(59)
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
	SingleFrameBasename() ISingleFrameBasenameContext
	FrameNum() IFrameNumContext
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

func (s *SingleFrameContext) SingleFrameBasename() ISingleFrameBasenameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleFrameBasenameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleFrameBasenameContext)
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
		p.SetState(61)
		p.Directory()
	}
	{
		p.SetState(62)
		p.SingleFrameBasename()
	}
	{
		p.SetState(63)
		p.FrameNum()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fileseqParserEXTENSION {
		{
			p.SetState(64)
			p.Extension()
		}

		p.SetState(69)
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
		p.SetState(70)
		p.Directory()
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62464) != 0 {
		{
			p.SetState(71)
			p.PlainBasename()
		}

	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == fileseqParserEXTENSION {
		{
			p.SetState(74)
			p.Extension()
		}

		p.SetState(79)
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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == fileseqParserSLASH {
		{
			p.SetState(80)
			p.Match(fileseqParserSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(88)
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
			{
				p.SetState(83)
				p.DirSegment()
			}
			{
				p.SetState(84)
				p.Match(fileseqParserSLASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(90)
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
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllNUM() []antlr.TerminalNode
	NUM(i int) antlr.TerminalNode
	AllDASH() []antlr.TerminalNode
	DASH(i int) antlr.TerminalNode
	AllSPECIAL_CHAR() []antlr.TerminalNode
	SPECIAL_CHAR(i int) antlr.TerminalNode
	AllFRAME_RANGE() []antlr.TerminalNode
	FRAME_RANGE(i int) antlr.TerminalNode
	AllDOT_FRAME_RANGE() []antlr.TerminalNode
	DOT_FRAME_RANGE(i int) antlr.TerminalNode
	AllDOT_NUM() []antlr.TerminalNode
	DOT_NUM(i int) antlr.TerminalNode

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

func (s *DirSegmentContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserWORD)
}

func (s *DirSegmentContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserWORD, i)
}

func (s *DirSegmentContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserNUM)
}

func (s *DirSegmentContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserNUM, i)
}

func (s *DirSegmentContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserDASH)
}

func (s *DirSegmentContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserDASH, i)
}

func (s *DirSegmentContext) AllSPECIAL_CHAR() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserSPECIAL_CHAR)
}

func (s *DirSegmentContext) SPECIAL_CHAR(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserSPECIAL_CHAR, i)
}

func (s *DirSegmentContext) AllFRAME_RANGE() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserFRAME_RANGE)
}

func (s *DirSegmentContext) FRAME_RANGE(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserFRAME_RANGE, i)
}

func (s *DirSegmentContext) AllDOT_FRAME_RANGE() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserDOT_FRAME_RANGE)
}

func (s *DirSegmentContext) DOT_FRAME_RANGE(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_FRAME_RANGE, i)
}

func (s *DirSegmentContext) AllDOT_NUM() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserDOT_NUM)
}

func (s *DirSegmentContext) DOT_NUM(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_NUM, i)
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
	p.EnterRule(localctx, 12, fileseqParserRULE_dirSegment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&63232) != 0) {
		{
			p.SetState(91)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&63232) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(94)
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

// ISequenceBasenameContext is an interface to support dynamic dispatch.
type ISequenceBasenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllNUM() []antlr.TerminalNode
	NUM(i int) antlr.TerminalNode
	AllDOT_NUM() []antlr.TerminalNode
	DOT_NUM(i int) antlr.TerminalNode
	AllDASH() []antlr.TerminalNode
	DASH(i int) antlr.TerminalNode
	AllSPECIAL_CHAR() []antlr.TerminalNode
	SPECIAL_CHAR(i int) antlr.TerminalNode
	AllEXTENSION() []antlr.TerminalNode
	EXTENSION(i int) antlr.TerminalNode

	// IsSequenceBasenameContext differentiates from other interfaces.
	IsSequenceBasenameContext()
}

type SequenceBasenameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceBasenameContext() *SequenceBasenameContext {
	var p = new(SequenceBasenameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_sequenceBasename
	return p
}

func InitEmptySequenceBasenameContext(p *SequenceBasenameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_sequenceBasename
}

func (*SequenceBasenameContext) IsSequenceBasenameContext() {}

func NewSequenceBasenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceBasenameContext {
	var p = new(SequenceBasenameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_sequenceBasename

	return p
}

func (s *SequenceBasenameContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceBasenameContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserWORD)
}

func (s *SequenceBasenameContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserWORD, i)
}

func (s *SequenceBasenameContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserNUM)
}

func (s *SequenceBasenameContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserNUM, i)
}

func (s *SequenceBasenameContext) AllDOT_NUM() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserDOT_NUM)
}

func (s *SequenceBasenameContext) DOT_NUM(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_NUM, i)
}

func (s *SequenceBasenameContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserDASH)
}

func (s *SequenceBasenameContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserDASH, i)
}

func (s *SequenceBasenameContext) AllSPECIAL_CHAR() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserSPECIAL_CHAR)
}

func (s *SequenceBasenameContext) SPECIAL_CHAR(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserSPECIAL_CHAR, i)
}

func (s *SequenceBasenameContext) AllEXTENSION() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserEXTENSION)
}

func (s *SequenceBasenameContext) EXTENSION(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserEXTENSION, i)
}

func (s *SequenceBasenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceBasenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceBasenameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitSequenceBasename(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) SequenceBasename() (localctx ISequenceBasenameContext) {
	localctx = NewSequenceBasenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, fileseqParserRULE_sequenceBasename)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(96)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62592) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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

// IPatternBasenameContext is an interface to support dynamic dispatch.
type IPatternBasenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllNUM() []antlr.TerminalNode
	NUM(i int) antlr.TerminalNode
	AllDOT_NUM() []antlr.TerminalNode
	DOT_NUM(i int) antlr.TerminalNode
	AllDASH() []antlr.TerminalNode
	DASH(i int) antlr.TerminalNode
	AllSPECIAL_CHAR() []antlr.TerminalNode
	SPECIAL_CHAR(i int) antlr.TerminalNode
	AllEXTENSION() []antlr.TerminalNode
	EXTENSION(i int) antlr.TerminalNode

	// IsPatternBasenameContext differentiates from other interfaces.
	IsPatternBasenameContext()
}

type PatternBasenameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternBasenameContext() *PatternBasenameContext {
	var p = new(PatternBasenameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_patternBasename
	return p
}

func InitEmptyPatternBasenameContext(p *PatternBasenameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_patternBasename
}

func (*PatternBasenameContext) IsPatternBasenameContext() {}

func NewPatternBasenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternBasenameContext {
	var p = new(PatternBasenameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_patternBasename

	return p
}

func (s *PatternBasenameContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternBasenameContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserWORD)
}

func (s *PatternBasenameContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserWORD, i)
}

func (s *PatternBasenameContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserNUM)
}

func (s *PatternBasenameContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserNUM, i)
}

func (s *PatternBasenameContext) AllDOT_NUM() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserDOT_NUM)
}

func (s *PatternBasenameContext) DOT_NUM(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_NUM, i)
}

func (s *PatternBasenameContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserDASH)
}

func (s *PatternBasenameContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserDASH, i)
}

func (s *PatternBasenameContext) AllSPECIAL_CHAR() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserSPECIAL_CHAR)
}

func (s *PatternBasenameContext) SPECIAL_CHAR(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserSPECIAL_CHAR, i)
}

func (s *PatternBasenameContext) AllEXTENSION() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserEXTENSION)
}

func (s *PatternBasenameContext) EXTENSION(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserEXTENSION, i)
}

func (s *PatternBasenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternBasenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternBasenameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitPatternBasename(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) PatternBasename() (localctx IPatternBasenameContext) {
	localctx = NewPatternBasenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, fileseqParserRULE_patternBasename)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62592) != 0) {
		{
			p.SetState(101)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62592) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(104)
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

// ISingleFrameBasenameContext is an interface to support dynamic dispatch.
type ISingleFrameBasenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllNUM() []antlr.TerminalNode
	NUM(i int) antlr.TerminalNode
	AllDOT_NUM() []antlr.TerminalNode
	DOT_NUM(i int) antlr.TerminalNode
	AllDASH() []antlr.TerminalNode
	DASH(i int) antlr.TerminalNode
	AllSPECIAL_CHAR() []antlr.TerminalNode
	SPECIAL_CHAR(i int) antlr.TerminalNode
	AllEXTENSION() []antlr.TerminalNode
	EXTENSION(i int) antlr.TerminalNode

	// IsSingleFrameBasenameContext differentiates from other interfaces.
	IsSingleFrameBasenameContext()
}

type SingleFrameBasenameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleFrameBasenameContext() *SingleFrameBasenameContext {
	var p = new(SingleFrameBasenameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_singleFrameBasename
	return p
}

func InitEmptySingleFrameBasenameContext(p *SingleFrameBasenameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = fileseqParserRULE_singleFrameBasename
}

func (*SingleFrameBasenameContext) IsSingleFrameBasenameContext() {}

func NewSingleFrameBasenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleFrameBasenameContext {
	var p = new(SingleFrameBasenameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = fileseqParserRULE_singleFrameBasename

	return p
}

func (s *SingleFrameBasenameContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleFrameBasenameContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserWORD)
}

func (s *SingleFrameBasenameContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserWORD, i)
}

func (s *SingleFrameBasenameContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserNUM)
}

func (s *SingleFrameBasenameContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserNUM, i)
}

func (s *SingleFrameBasenameContext) AllDOT_NUM() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserDOT_NUM)
}

func (s *SingleFrameBasenameContext) DOT_NUM(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_NUM, i)
}

func (s *SingleFrameBasenameContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserDASH)
}

func (s *SingleFrameBasenameContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserDASH, i)
}

func (s *SingleFrameBasenameContext) AllSPECIAL_CHAR() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserSPECIAL_CHAR)
}

func (s *SingleFrameBasenameContext) SPECIAL_CHAR(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserSPECIAL_CHAR, i)
}

func (s *SingleFrameBasenameContext) AllEXTENSION() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserEXTENSION)
}

func (s *SingleFrameBasenameContext) EXTENSION(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserEXTENSION, i)
}

func (s *SingleFrameBasenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleFrameBasenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleFrameBasenameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fileseqVisitor:
		return t.VisitSingleFrameBasename(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fileseqParser) SingleFrameBasename() (localctx ISingleFrameBasenameContext) {
	localctx = NewSingleFrameBasenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, fileseqParserRULE_singleFrameBasename)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(106)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62592) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllNUM() []antlr.TerminalNode
	NUM(i int) antlr.TerminalNode
	AllDOT_NUM() []antlr.TerminalNode
	DOT_NUM(i int) antlr.TerminalNode
	AllDASH() []antlr.TerminalNode
	DASH(i int) antlr.TerminalNode
	AllSPECIAL_CHAR() []antlr.TerminalNode
	SPECIAL_CHAR(i int) antlr.TerminalNode

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

func (s *PlainBasenameContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserWORD)
}

func (s *PlainBasenameContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserWORD, i)
}

func (s *PlainBasenameContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserNUM)
}

func (s *PlainBasenameContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserNUM, i)
}

func (s *PlainBasenameContext) AllDOT_NUM() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserDOT_NUM)
}

func (s *PlainBasenameContext) DOT_NUM(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserDOT_NUM, i)
}

func (s *PlainBasenameContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserDASH)
}

func (s *PlainBasenameContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserDASH, i)
}

func (s *PlainBasenameContext) AllSPECIAL_CHAR() []antlr.TerminalNode {
	return s.GetTokens(fileseqParserSPECIAL_CHAR)
}

func (s *PlainBasenameContext) SPECIAL_CHAR(i int) antlr.TerminalNode {
	return s.GetToken(fileseqParserSPECIAL_CHAR, i)
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
	p.EnterRule(localctx, 20, fileseqParserRULE_plainBasename)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62464) != 0) {
		{
			p.SetState(111)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62464) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(114)
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
	p.EnterRule(localctx, 22, fileseqParserRULE_frameRange)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
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
	p.EnterRule(localctx, 24, fileseqParserRULE_frameNum)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
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
	p.EnterRule(localctx, 26, fileseqParserRULE_padding)
	var _la int

	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case fileseqParserUDIM_ANGLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(fileseqParserUDIM_ANGLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fileseqParserUDIM_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.Match(fileseqParserUDIM_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fileseqParserPRINTF_PAD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)
			p.Match(fileseqParserPRINTF_PAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fileseqParserHOUDINI_PAD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)
			p.Match(fileseqParserHOUDINI_PAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case fileseqParserHASH:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == fileseqParserHASH {
			{
				p.SetState(124)
				p.Match(fileseqParserHASH)
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
		}

	case fileseqParserAT:
		p.EnterOuterAlt(localctx, 6)
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == fileseqParserAT {
			{
				p.SetState(129)
				p.Match(fileseqParserAT)
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
	p.EnterRule(localctx, 28, fileseqParserRULE_extension)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(fileseqParserEXTENSION)
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
