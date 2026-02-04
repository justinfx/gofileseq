package parser

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// ParseResult holds the components extracted from parsing a file sequence
type ParseResult struct {
	// Full path components
	Directory  string
	Basename   string
	FrameRange string
	Padding    string
	Extension  string

	// Parse type flags
	IsSequence   bool
	IsSingleFile bool
	IsPlainFile  bool
}

// errorListener collects parse errors
type errorListener struct {
	*antlr.DefaultErrorListener
	errors []error
}

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {

	l.errors = append(l.errors, fmt.Errorf("parse error at %d:%d: %s", line, column, msg))
}

// ParseFileSequence parses a file sequence string and returns its components.
// Uses SLL prediction mode for better performance (30x faster), falling back to LL if needed.
func ParseFileSequence(input string) (ParseResult, error) {
	// Create the lexer
	inputStream := antlr.NewInputStream(input)
	lexer := NewfileseqLexer(inputStream)

	// Create the token stream
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	parser := NewfileseqParser(stream)

	// Add custom error listener to both lexer and parser
	errListener := &errorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errListener)

	// Try SLL prediction mode first (much faster for unambiguous grammars)
	parser.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	tree := parser.Input()

	// If SLL failed, retry with LL prediction mode (slower but more thorough)
	if len(errListener.errors) > 0 {
		// Reset for retry
		errListener.errors = errListener.errors[:0]
		inputStream = antlr.NewInputStream(input)
		lexer.SetInputStream(inputStream)
		stream.SetTokenSource(lexer)
		parser.SetInputStream(stream)

		parser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLL)
		tree = parser.Input()

		// Check for parse errors after LL retry
		if len(errListener.errors) > 0 {
			return ParseResult{}, fmt.Errorf("failed to parse sequence %q: %w", input, errListener.errors[0])
		}
	}

	// Create and run the visitor
	visitor := &fileSeqVisitor{}
	tree.Accept(visitor)

	return visitor.result, nil
}

// fileSeqVisitor implements the visitor pattern to extract components
// This is TRIVIAL - just extracting what the grammar already parsed
type fileSeqVisitor struct {
	BasefileseqVisitor
	result ParseResult
}

// VisitInput dispatches to the correct child visitor
func (v *fileSeqVisitor) VisitInput(ctx *InputContext) interface{} {
	if seqCtx := ctx.Sequence(); seqCtx != nil {
		return v.VisitSequence(seqCtx.(*SequenceContext))
	}
	if poCtx := ctx.PatternOnly(); poCtx != nil {
		return v.VisitPatternOnly(poCtx.(*PatternOnlyContext))
	}
	if sfCtx := ctx.SingleFrame(); sfCtx != nil {
		return v.VisitSingleFrame(sfCtx.(*SingleFrameContext))
	}
	if pfCtx := ctx.PlainFile(); pfCtx != nil {
		return v.VisitPlainFile(pfCtx.(*PlainFileContext))
	}
	return nil
}

// VisitSequence - grammar already separated all components
func (v *fileSeqVisitor) VisitSequence(ctx *SequenceContext) interface{} {
	v.result.IsSequence = true

	// Extract directory (grammar already parsed it)
	if dirCtx := ctx.Directory(); dirCtx != nil {
		v.result.Directory = dirCtx.GetText()
	}

	// Extract basename (grammar already parsed it)
	if basenameCtx := ctx.SequenceBasename(); basenameCtx != nil {
		v.result.Basename = basenameCtx.GetText()
	}

	// Extract frame range (grammar already parsed it)
	if frameRangeCtx := ctx.FrameRange(); frameRangeCtx != nil {
		v.result.FrameRange = frameRangeCtx.GetText()
	}

	// Post-process: move leading dot from frame range to basename
	// (see parse_postprocess.go Special Case 3 for detailed explanation)
	moveLeadingDotFromFrameRange(&v.result)

	// Extract padding (grammar already parsed it)
	if padCtx := ctx.Padding(); padCtx != nil {
		v.result.Padding = padCtx.GetText()
	}

	// Extract extensions (grammar already parsed them)
	v.result.Extension = v.extractExtensions(ctx.AllExtension())

	return nil
}

// VisitSingleFrame - grammar already separated all components
func (v *fileSeqVisitor) VisitSingleFrame(ctx *SingleFrameContext) interface{} {
	v.result.IsSingleFile = true

	// Extract directory (grammar already parsed it)
	if dirCtx := ctx.Directory(); dirCtx != nil {
		v.result.Directory = dirCtx.GetText()
	}

	// Extract basename (grammar already parsed it)
	if basenameCtx := ctx.SingleFrameBasename(); basenameCtx != nil {
		v.result.Basename = basenameCtx.GetText()
	}

	// Extract frame number (grammar already parsed it)
	if frameNumCtx := ctx.FrameNum(); frameNumCtx != nil {
		v.result.FrameRange = frameNumCtx.GetText()
	}

	// Post-process: move leading dot from frame range to basename
	// (see parse_postprocess.go Special Case 3 for detailed explanation)
	moveLeadingDotFromFrameRange(&v.result)

	// Extract extensions (grammar already parsed them)
	v.result.Extension = v.extractExtensions(ctx.AllExtension())

	// Post-process: handle special case of multiple DOT_NUM tokens
	// (see parse_postprocess.go Special Case 3 for detailed explanation)
	handleMultipleDotNumTokens(&v.result)

	return nil
}

// VisitPatternOnly - pattern-only sequence (padding without frame range)
func (v *fileSeqVisitor) VisitPatternOnly(ctx *PatternOnlyContext) interface{} {
	v.result.IsSequence = true // Pattern-only is still a sequence type

	// Extract directory
	if dirCtx := ctx.Directory(); dirCtx != nil {
		v.result.Directory = dirCtx.GetText()
	}

	// Extract basename
	if basenameCtx := ctx.PatternBasename(); basenameCtx != nil {
		v.result.Basename = basenameCtx.GetText()
	}

	// No frame range for pattern-only (initially)

	// Extract padding
	if padCtx := ctx.Padding(); padCtx != nil {
		v.result.Padding = padCtx.GetText()
	}

	// Extract extensions
	v.result.Extension = v.extractExtensions(ctx.AllExtension())

	// Post-process: extract digits that were absorbed into extension token
	// (see parse_postprocess.go for detailed explanation)
	extractDigitsFromExtensionToken(&v.result)

	return nil
}

// VisitPlainFile - grammar already separated all components
func (v *fileSeqVisitor) VisitPlainFile(ctx *PlainFileContext) interface{} {
	v.result.IsPlainFile = true

	// Extract directory (grammar already parsed it)
	if dirCtx := ctx.Directory(); dirCtx != nil {
		v.result.Directory = dirCtx.GetText()
	}

	// Extract basename (grammar already parsed it)
	if basenameCtx := ctx.PlainBasename(); basenameCtx != nil {
		v.result.Basename = basenameCtx.GetText()
	}

	// Extract extensions (grammar already parsed them)
	v.result.Extension = v.extractExtensions(ctx.AllExtension())

	return nil
}

// extractExtensions concatenates all extension tokens
func (v *fileSeqVisitor) extractExtensions(extCtxs []IExtensionContext) string {
	if len(extCtxs) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, extCtx := range extCtxs {
		sb.WriteString(extCtx.GetText())
	}
	return sb.String()
}
