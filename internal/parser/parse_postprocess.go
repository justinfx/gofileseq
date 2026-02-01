package parser

import (
	"strconv"
	"strings"
)

// This file contains post-processing logic for 2 special cases that CANNOT be solved
// in the ANTLR grammar alone. These cases require semantic decision-making beyond what
// context-free grammars can express.
//
// IMPORTANT: When porting to C++, these same post-processing steps will be needed.

// ============================================================================
// SPECIAL CASE 1: Multiple DOT_NUM Tokens
// ============================================================================
//
// Problem: Input like ".10000000000.123" has two DOT_NUM tokens.
//
// Grammar limitation:
//   The grammar matches greedily: singleFrameBasename consumes first DOT_NUM,
//   frameNum consumes second DOT_NUM. Both parses are syntactically valid:
//     Parse A: basename=".10000000000." frame="123" ext=""
//     Parse B: basename="." frame="10000000000" ext=".123"
//
// Why can't this be in grammar:
//   Context-free grammars cannot express semantic preferences like "prefer the
//   first/larger number as frame". Both parses are valid syntax; we need semantic
//   knowledge (based on Python's behavior) to choose Parse B.
//
// Solution:
//   When basename ends with DOT_NUM pattern and frameNum is also DOT_NUM, and
//   extension is empty, swap them to prefer first number as frame.

// handleMultipleDotNumTokens swaps basename and frame when we detect the pattern
// of two consecutive DOT_NUM tokens with no regular extension.
//
// This implements Python's semantic rule: "when two dot-number sequences appear,
// treat the first as the frame and second as a digit-only extension."
func handleMultipleDotNumTokens(result *ParseResult) {
	// Only apply when: extension empty, frame exists, basename exists
	if result.Extension == "" && result.FrameRange != "" && result.Basename != "" {
		basename := result.Basename

		// Remove trailing dot if present (added by normal frameNum processing)
		if strings.HasSuffix(basename, ".") {
			basename = basename[:len(basename)-1]
		}

		// Check if basename ends with pattern ".digits"
		if dotIdx := strings.LastIndex(basename, "."); dotIdx >= 0 {
			suffix := basename[dotIdx+1:]
			// Verify suffix is all digits
			if _, err := strconv.ParseInt(suffix, 10, 64); err == nil {
				// Swap: first number → frame, second number → extension
				result.Extension = "." + result.FrameRange
				result.FrameRange = suffix
				result.Basename = basename[:dotIdx+1] // Keep through the dot
			}
		}
	}
}

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
