grammar fileseq;

// ============================================================================
// Parser Rules
// ============================================================================

input
    : sequence EOF
    | patternOnly EOF
    | singleFrame EOF
    | plainFile EOF
    ;

// Sequence with padding: /path/file.1-100#.exr or /path/file1-100#.exr
// Also handles hidden files: /path/.hidden5.1-10#.7zip
// Basename and extension are optional (allows patterns like /path/1-100#.ext)
sequence
    : directory sequenceBasename? frameRange padding extension*
    ;

// Pattern-only sequence (padding without frame range): /path/file.@@.ext
// Basename and extension are optional (allows patterns like /path/@@@.ext)
patternOnly
    : directory patternBasename? padding extension*
    ;

// Single frame: /path/file.100.exr or /path/file.100 (extension optional)
// Also handles hidden files: /path/.hidden.100
singleFrame
    : directory singleFrameBasename frameNum extension*
    ;

// Plain file: /path/file.txt or /path/file or /path/.hidden (no frame pattern)
plainFile
    : directory plainBasename? extension*
    ;

// Directory: optional leading slash + segments ending with slash
directory
    : SLASH? (dirSegment SLASH)*
    ;

// Directory segments can contain anything including frame-range-like patterns
dirSegment
    : (WORD | NUM | DASH | SPECIAL_CHAR | FRAME_RANGE | DOT_FRAME_RANGE | DOT_NUM)+
    ;

// Basename for sequences: can include EXTENSION (for hidden files)
// Also includes FRAME_RANGE tokens for date-like patterns (e.g., "name_2025-05-13_")
sequenceBasename
    : (WORD | NUM | DOT_NUM | DASH | SPECIAL_CHAR | EXTENSION | FRAME_RANGE | DOT_FRAME_RANGE)+
    ;

// Basename for pattern-only: same as sequence
patternBasename
    : (WORD | NUM | DOT_NUM | DASH | SPECIAL_CHAR | EXTENSION | FRAME_RANGE | DOT_FRAME_RANGE)+
    ;

// Basename for single frames: can include EXTENSION (for hidden files like .hidden.100)
// Also includes FRAME_RANGE for date-like patterns
singleFrameBasename
    : (WORD | NUM | DOT_NUM | DASH | SPECIAL_CHAR | EXTENSION | FRAME_RANGE | DOT_FRAME_RANGE)+
    ;

// Basename for plain files: does NOT include EXTENSION or DOT_NUM
// (so both regular and digit-only extensions can be consumed by extension rule)
// But DOES include FRAME_RANGE tokens (for filenames like "name_2025-05-13.ext")
plainBasename
    : (WORD | NUM | DASH | SPECIAL_CHAR | FRAME_RANGE | DOT_FRAME_RANGE)+
    ;

// Frame range: may or may not have leading dot
// Also includes single frame numbers (for single-frame sequences with padding)
frameRange
    : DOT_FRAME_RANGE
    | FRAME_RANGE
    | DOT_NUM      // Single frame with dot: .100
    | NUM          // Single frame without dot: 100
    ;

// Single frame number with leading dot: .100 or .-10
frameNum
    : DOT_NUM
    ;

padding
    : UDIM_ANGLE
    | UDIM_PAREN
    | PRINTF_PAD
    | HOUDINI_PAD
    | HASH+
    | AT+
    ;

// Extension can be:
// - EXTENSION tokens (.tar, .gz, .exr)
// - DOT_NUM for digit-only extensions (.123, .10000000000)
// - WORD for non-dot extensions after padding (_exr, _extra)
// - Followed by optional DASH and NUM (for extensions like .tar.gz-1)
extension
    : EXTENSION (DASH NUM)?
    | DOT_NUM
    | WORD
    ;

// ============================================================================
// Lexer Rules - ORDER MATTERS FOR PRIORITY
// ============================================================================

// Padding markers - HIGHEST PRIORITY
UDIM_ANGLE: '<UDIM>';
UDIM_PAREN: '%(UDIM)d';
PRINTF_PAD: '%' [0-9]* 'd';
HOUDINI_PAD: '$F' [0-9]*;
HASH: '#';
AT: '@';

// Extension: dot + pattern containing at least one letter
EXTENSION: '.' ([a-zA-Z_] | [0-9]* [a-zA-Z] [a-zA-Z0-9_]*);

// Frame range with leading dot (must have comma, colon, or dash after first number)
// Matches: .1-100, .-10-100, .1,2,3, .1-10x2, .1,2,3,5-10,20-30
DOT_FRAME_RANGE: '.' '-'? [0-9]+ [,:-] [0-9xy:,-]*;

// Frame range without leading dot (must have comma, colon, or dash after first number)
// Matches: 1-100, -10-100, 1,2,3, 1-10x2
FRAME_RANGE: '-'? [0-9]+ [,:-] [0-9xy:,-]*;

// Frame number with dot: .100 or .-10 (single frame, no range delimiter)
DOT_NUM: '.' '-'? [0-9]+;

// Slash separator
SLASH: '/' | '\\';

// Special characters allowed in basenames
SPECIAL_CHAR: [:,.];

// Number sequence (for basenames containing numbers)
NUM: [0-9]+;

// Words (letters and underscores, no digits or dashes)
WORD: [a-zA-Z_]+;

// Dash as separate token
DASH: '-';

// Skip whitespace
WS: [ \t\r\n]+ -> skip;
