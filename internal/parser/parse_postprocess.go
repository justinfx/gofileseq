package parser

import (
	"strings"
)

// This file contains post-processing logic for special cases that are not
// solved in the ANTLR grammar alone.
//
// IMPORTANT: When porting to C++, these same post-processing steps will be needed.
//
// Summary of Special Cases:
//   2. Extension with trailing digits - lexer maximal munch embeds digits in token
//   3. Leading dot in frame range - lexer atomicity includes dot in token text

// ============================================================================
// SPECIAL CASE 2: Extension Token Contains Trailing Digits
// ============================================================================
//
// Problem: Input like "/dir/f.tmp12345@@@@@" where lexer produces EXTENSION(".tmp12345")
//
// Lexer limitation:
//   The EXTENSION token pattern [a-zA-Z][a-zA-Z0-9_]* uses maximal munch, consuming
//   ".tmp12345" as a single atomic token. By the time the parser sees it, the digits
//   are already embedded in the token text.
//
// Why can't this be in grammar:
//   The lexer has already tokenized the input. To fix this in the lexer would require:
//     - Restricting EXTENSION to not end with digits (breaks valid extensions like .mp4, .7zip)
//     - Using lexer modes to change tokenization after seeing padding (complex, language-specific)
//     - Semantic predicates with lookahead (very complex, non-portable)
//
//   Post-processing the token text is simpler and more portable across language implementations.
//
// Solution:
//   When we have padding but no frame range, and basename exists, extract trailing digits
//   from the basename as the frame number. This handles cases where the lexer consumed
//   digits as part of an EXTENSION token.

// extractDigitsFromExtensionToken extracts trailing digits from basename when the lexer
// consumed them as part of an EXTENSION token (e.g., ".tmp12345").
//
// This only applies to pattern-only sequences (padding without explicit frame range),
// where we need to recover the frame number that got absorbed into the extension token.
func extractDigitsFromExtensionToken(result *ParseResult) {
	// Only apply when: padding exists, no frame range yet, basename exists
	if result.Padding != "" && result.FrameRange == "" && result.Basename != "" {
		basename := result.Basename

		// Find trailing digits in basename
		i := len(basename) - 1
		for i >= 0 && basename[i] >= '0' && basename[i] <= '9' {
			i--
		}

		// If we found trailing digits (and there's content before them), extract as frame
		if i < len(basename)-1 && i >= 0 {
			result.FrameRange = basename[i+1:]
			result.Basename = basename[:i+1]
		}
	}
}

// ============================================================================
// SPECIAL CASE 3: Leading Dot in Frame Range Token
// ============================================================================
//
// Problem: Lexer tokens DOT_FRAME_RANGE (".1-100") and DOT_NUM (".100") include the
//          leading dot as part of the token text, but semantically the dot belongs
//          to the basename as a separator.
//
// Lexer design:
//   Having separate DOT_FRAME_RANGE and DOT_NUM tokens provides semantic clarity and
//   avoids token ambiguity (vs having DOT as separate token which conflicts with EXTENSION).
//
//   However, this means the dot is part of the token text: ".1-100" not "." + "1-100"
//
// Why can't this be in grammar:
//   Lexer tokens are atomic - the parser cannot split ".1-100" into "." and "1-100".
//   We would need to:
//     - Remove DOT_FRAME_RANGE/DOT_NUM tokens entirely
//     - Make DOT a separate token (conflicts with EXTENSION: '.' [a-zA-Z]...)
//     - Carefully manage token priority to avoid ".123" ambiguity (DOT+NUM vs DOT_NUM vs extension)
//
//   The current lexer design is cleaner; this simple post-processing is more maintainable.
//
// Solution:
//   When frame range starts with ".", move the dot from frame range to basename.
//   This makes "file" + ".1-100" → "file." + "1-100" for proper reconstruction.
//
// Examples:
//   Input: "/path/file.1-100#.exr"
//     Grammar: basename="file", frameRange=".1-100"
//     After:   basename="file.", frameRange="1-100"
//     String(): "/path/" + "file." + "1-100" + "#" + ".exr" = "/path/file.1-100#.exr" ✓
//
//   Input: "/path/file.100.exr"
//     Grammar: basename="file", frameRange=".100"
//     After:   basename="file.", frameRange="100"
//     String(): "/path/" + "file." + "100" + ".exr" = "/path/file.100.exr" ✓

// moveLeadingDotFromFrameRange moves a leading dot from the frame range to the basename.
//
// This handles the semantic meaning of the dot as a separator between basename and frame
// number, even though the lexer captured it as part of the frame range token for clarity.
func moveLeadingDotFromFrameRange(result *ParseResult) {
	if result.FrameRange != "" && strings.HasPrefix(result.FrameRange, ".") {
		// Move the dot from frame range to basename
		result.Basename += "."
		result.FrameRange = result.FrameRange[1:] // Remove leading dot
	}
}

// splitDecimalStepFromFrameRange restores the v2 meaning of a decimal step
// value, which go and cpp do not support. The lexer keeps "1-5x0.25" as one
// FRAME_RANGE token for python's decimal steps. Here the text up to the last
// dot moves to the basename and the trailing digits become the frame, so
// "file.1-5x0.25#.exr" is basename "file.1-5x0.", frame "25", as in v2.
// Must run after moveLeadingDotFromFrameRange.
func splitDecimalStepFromFrameRange(result *ParseResult) {
	if i := strings.LastIndexByte(result.FrameRange, '.'); i >= 0 {
		result.Basename += result.FrameRange[:i+1]
		result.FrameRange = result.FrameRange[i+1:]
	}
}
