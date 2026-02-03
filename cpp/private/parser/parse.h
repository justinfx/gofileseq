#ifndef FILESEQ_PRIVATE_PARSE_H_
#define FILESEQ_PRIVATE_PARSE_H_

#include <string>

namespace fileseq {
namespace internal {

// ParseResult holds the parsed components of a file sequence string
struct ParseResult {
    std::string directory;     // "/path/to/" (includes trailing separator)
    std::string basename;      // "file." (includes trailing dot if present)
    std::string frameRange;    // "1-100" (no leading dot after post-processing)
    std::string padding;       // "#", "@@", "%04d", etc.
    std::string extension;     // ".exr", ".tar.gz" (all parts concatenated)

    bool isSequence;           // Has padding (sequence or pattern-only)
    bool isSingleFile;         // Single frame, no padding
    bool isPlainFile;          // No frame info

    ParseResult()
        : isSequence(false)
        , isSingleFile(false)
        , isPlainFile(false) {}
};

// ParseFileSequence parses a file sequence string and returns the result.
// Returns true if parsing succeeded, false otherwise.
bool parseFileSequence(ParseResult& result, const std::string& path);

} // namespace internal
} // namespace fileseq

#endif // FILESEQ_PRIVATE_PARSE_H_
