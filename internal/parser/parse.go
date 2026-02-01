package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// ParseResult holds the components extracted from parsing a file sequence
type ParseResult struct {
	// Full path components
	Directory    string
	Basename     string
	FrameRange   string
	Padding      string
	Extension    string

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

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Errorf("parse error at %d:%d: %s", line, column, msg))
}

// ParseFileSequence parses a file sequence string and returns its components
func ParseFileSequence(input string) (ParseResult, error) {
	// Create the lexer
	inputStream := antlr.NewInputStream(input)
	lexer := NewfileseqLexer(inputStream)

	// Create the token stream
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	parser := NewfileseqParser(stream)

	// Add custom error listener
	parser.RemoveErrorListeners()
	errListener := &errorListener{}
	parser.AddErrorListener(errListener)

	// Parse the input
	tree := parser.Input()

	// Check for parse errors
	if len(errListener.errors) > 0 {
		return ParseResult{}, fmt.Errorf("failed to parse sequence %q: %w", input, errListener.errors[0])
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
		text := frameRangeCtx.GetText()
		if strings.HasPrefix(text, ".") {
			// Frame range had a leading dot - move it to basename
			v.result.Basename += "."
			v.result.FrameRange = text[1:]
		} else {
			v.result.FrameRange = text
		}
	}

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
		text := frameNumCtx.GetText()
		if strings.HasPrefix(text, ".") {
			// Frame number had a leading dot - move it to basename
			v.result.Basename += "."
			v.result.FrameRange = text[1:]
		} else {
			v.result.FrameRange = text
		}
	}

	// Extract extensions (grammar already parsed them)
	v.result.Extension = v.extractExtensions(ctx.AllExtension())

	// SPECIAL CASE: Multiple DOT_NUM tokens (e.g., ".10000000000.123")
	// If basename contains a DOT_NUM pattern and frameNum is also a DOT_NUM,
	// and extension is empty, we should:
	// - Treat the FIRST number (in basename) as the frame
	// - Treat the LAST number (current frame) as a digit-only extension
	if v.result.Extension == "" && v.result.FrameRange != "" && v.result.Basename != "" {
		basename := v.result.Basename
		// Remove trailing dot if present (we added it when processing frameNum)
		if strings.HasSuffix(basename, ".") {
			basename = basename[:len(basename)-1]
		}

		// Check if basename ends with a pattern like ".digits" or is just ".digits"
		if dotIdx := strings.LastIndex(basename, "."); dotIdx >= 0 {
			suffix := basename[dotIdx+1:]
			if _, err := strconv.ParseInt(suffix, 10, 64); err == nil {
				// Basename ends with .number and we have a frame that's also a number
				// Swap them: first number becomes frame, second (current frame) becomes extension
				v.result.Extension = "." + v.result.FrameRange
				v.result.FrameRange = suffix
				v.result.Basename = basename[:dotIdx+1] // Keep through the dot
			}
		}
	}

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

	// SPECIAL CASE: Extension + digits + padding (e.g., "/dir/f.tmp12345@@@@@")
	// If basename ends with digits and we have padding but no frame range,
	// extract the trailing digits as the frame number.
	// This handles the case where the lexer consumed ".tmp12345" as a single EXTENSION token.
	if v.result.Padding != "" && v.result.FrameRange == "" && v.result.Basename != "" {
		// Find trailing digits in basename
		basename := v.result.Basename
		i := len(basename) - 1
		for i >= 0 && basename[i] >= '0' && basename[i] <= '9' {
			i--
		}
		// If we found trailing digits, extract them as the frame
		if i < len(basename)-1 && i >= 0 {
			v.result.FrameRange = basename[i+1:]
			v.result.Basename = basename[:i+1]
		}
	}

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
