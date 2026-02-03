#ifndef FILESEQ_PRIVATE_PARSE_POSTPROCESS_H_
#define FILESEQ_PRIVATE_PARSE_POSTPROCESS_H_

#include "parse.h"

namespace fileseq {
namespace internal {

// This file contains post-processing logic for 3 special cases that CANNOT be solved
// in the ANTLR grammar alone. These cases require semantic decision-making or lexer
// limitations that cannot be expressed in pure context-free grammar rules.
//
// IMPORTANT: These same 3 post-processing steps are needed in all language implementations.
//
// Summary of Special Cases:
//   1. Multiple DOT_NUM tokens - requires semantic choice of which number is the frame
//   2. Extension with trailing digits - lexer maximal munch embeds digits in token
//   3. Leading dot in frame range - lexer atomicity includes dot in token text

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
void handleMultipleDotNumTokens(ParseResult& result);

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
void extractDigitsFromExtensionToken(ParseResult& result);

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
void moveLeadingDotFromFrameRange(ParseResult& result);

// Apply all post-processing steps to a ParseResult
void postProcessParseResult(ParseResult& result);

} // namespace internal
} // namespace fileseq

#endif // FILESEQ_PRIVATE_PARSE_POSTPROCESS_H_
