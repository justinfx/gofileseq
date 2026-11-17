package parser

import (
	"fmt"
	"strings"
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

// ParseFileSequence parses a file sequence string and returns its components.
//
// The parser is generated from grammar/fileseq.g4 by galop
// (https://github.com/martianoff/galop), which gives the parse ANTLR's LL
// prediction would give, without a runtime prediction pass in most decisions.
func ParseFileSequence(input string) (ParseResult, error) {
	tokens := FileseqLex(input)
	tree, err := ParseInput(tokens)
	// As with ANTLR, a token recognition error fails the parse too, and the
	// error reported is the first one in the input.
	if len(tokens.Errors) > 0 {
		first := tokens.Errors[0]
		if se, ok := err.(*FileseqSyntaxError); ok && before(*se, first) {
			first = *se
		}
		return ParseResult{}, parseError(input, first)
	}
	if err != nil {
		if se, ok := err.(*FileseqSyntaxError); ok {
			return ParseResult{}, parseError(input, *se)
		}
		return ParseResult{}, fmt.Errorf("failed to parse sequence %q: %w", input, err)
	}

	var result ParseResult
	switch {
	case tree.Sequence().Ok():
		err = extractSequence(tree.Sequence(), &result)
	case tree.PatternOnly().Ok():
		err = extractPatternOnly(tree.PatternOnly(), &result)
	case tree.SingleFrame().Ok():
		extractSingleFrame(tree.SingleFrame(), &result)
	case tree.PlainFile().Ok():
		extractPlainFile(tree.PlainFile(), &result)
	}
	if err != nil {
		return ParseResult{}, err
	}
	return result, nil
}

// parseError is the error for a syntax error, in the format the ANTLR-based
// parser used.
func parseError(input string, e FileseqSyntaxError) error {
	return fmt.Errorf("failed to parse sequence %q: %w", input,
		fmt.Errorf("parse error at %d:%d: %s", e.Line, e.Column, e.Msg))
}

// before reports whether a is earlier in the input than b.
func before(a, b FileseqSyntaxError) bool {
	return a.Line < b.Line || a.Line == b.Line && a.Column < b.Column
}

// extractSequence - grammar already separated all components
func extractSequence(ctx SequenceNode, result *ParseResult) error {
	result.IsSequence = true

	// Extract directory (grammar already parsed it)
	if dirCtx := ctx.Directory(); dirCtx.Ok() {
		result.Directory = dirCtx.GetText()
	}

	// Extract basename (grammar already parsed it)
	if basenameCtx := ctx.Basename(); basenameCtx.Ok() {
		result.Basename = basenameCtx.GetText()
	}

	// Check for unsupported subframe syntax (Python-only feature)
	// Subframes use: file.1-5#.10-20@@.exr (dual range) or file.1-5@.#.exr (composite padding)
	if len(ctx.AllFrameRange()) > 1 || len(ctx.AllPadding()) > 1 {
		return fmt.Errorf("subframe syntax not supported: %s", ctx.GetText())
	}

	// Extract frame range (grammar already parsed it)
	if frameRangeCtx := ctx.FrameRange(0); frameRangeCtx.Ok() {
		result.FrameRange = frameRangeCtx.GetText()
	}

	// Post-process: move leading dot from frame range to basename
	// (see parse_postprocess.go Special Case 3 for detailed explanation)
	moveLeadingDotFromFrameRange(result)

	// Extract padding (grammar already parsed it)
	if padCtx := ctx.Padding(0); padCtx.Ok() {
		result.Padding = padCtx.GetText()
	}

	// Extract extensions (grammar already parsed them)
	result.Extension = extractExtensions(ctx.AllExtension())
	return nil
}

// extractSingleFrame - grammar already separated all components
func extractSingleFrame(ctx SingleFrameNode, result *ParseResult) {
	result.IsSingleFile = true

	// Extract directory (grammar already parsed it)
	if dirCtx := ctx.Directory(); dirCtx.Ok() {
		result.Directory = dirCtx.GetText()
	}

	// Extract basename (grammar already parsed it)
	if basenameCtx := ctx.Basename(); basenameCtx.Ok() {
		result.Basename = basenameCtx.GetText()
	}

	// Extract frame number (grammar already parsed it)
	if frameNumCtx := ctx.FrameNum(); frameNumCtx.Ok() {
		result.FrameRange = frameNumCtx.GetText()
	}

	// Post-process: move leading dot from frame range to basename
	// (see parse_postprocess.go Special Case 3 for detailed explanation)
	moveLeadingDotFromFrameRange(result)

	// Extract extensions (grammar already parsed them)
	result.Extension = extractExtensions(ctx.AllExtension())

	// Post-process: handle special case of multiple DOT_NUM tokens
	// (see parse_postprocess.go Special Case 3 for detailed explanation)
	handleMultipleDotNumTokens(result)
}

// extractPatternOnly - pattern-only sequence (padding without frame range)
func extractPatternOnly(ctx PatternOnlyNode, result *ParseResult) error {
	result.IsSequence = true // Pattern-only is still a sequence type

	// Extract directory
	if dirCtx := ctx.Directory(); dirCtx.Ok() {
		result.Directory = dirCtx.GetText()
	}

	// Extract basename
	if basenameCtx := ctx.Basename(); basenameCtx.Ok() {
		result.Basename = basenameCtx.GetText()
	}

	// Check for unsupported subframe syntax (Python-only feature)
	// Pattern-only subframes use: file.#.@@.exr (composite padding)
	if len(ctx.AllPadding()) > 1 {
		return fmt.Errorf("subframe syntax not supported: %s", ctx.GetText())
	}

	// No frame range for pattern-only (initially)

	// Get first padding
	if padCtx := ctx.Padding(0); padCtx.Ok() {
		result.Padding = padCtx.GetText()
	}

	// Extract extensions
	result.Extension = extractExtensions(ctx.AllExtension())

	// Post-process: extract digits that were absorbed into extension token
	// (see parse_postprocess.go for detailed explanation)
	extractDigitsFromExtensionToken(result)
	return nil
}

// extractPlainFile - grammar already separated all components
func extractPlainFile(ctx PlainFileNode, result *ParseResult) {
	result.IsPlainFile = true

	// Extract directory (grammar already parsed it)
	if dirCtx := ctx.Directory(); dirCtx.Ok() {
		result.Directory = dirCtx.GetText()
	}

	// Extract basename (grammar already parsed it)
	if basenameCtx := ctx.PlainBasename(); basenameCtx.Ok() {
		result.Basename = basenameCtx.GetText()
	}

	// Extract extensions (grammar already parsed them)
	result.Extension = extractExtensions(ctx.AllExtension())
}

// extractExtensions concatenates all extension tokens
func extractExtensions(extCtxs []ExtensionNode) string {
	if len(extCtxs) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, extCtx := range extCtxs {
		sb.WriteString(extCtx.GetText())
	}
	return sb.String()
}
