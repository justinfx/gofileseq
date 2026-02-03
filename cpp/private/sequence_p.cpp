
#include "sequence_p.h"
#include "parser/parse.h"

#include <iostream>
#include <string>
#include <vector>
#include <regex>

namespace fileseq {

namespace internal {


bool getSplitPatternMatch(SeqPatternMatch &match, const std::string &path) {
    // Use ANTLR parser to parse the sequence pattern
    ParseResult result;
    if (!parseFileSequence(result, path)) {
        return false;
    }

    // Only accept sequences with padding (pattern or full sequence)
    if (!result.isSequence) {
        return false;
    }

    // Convert ParseResult to SeqPatternMatch
    match.base = result.directory + result.basename;
    match.range = result.frameRange;
    match.padChars = result.padding;
    match.ext = result.extension;

    return true;
}


bool getSingleFrameMatch(SeqPatternMatch &match, const std::string &path, bool require_frame) {
    // Use ANTLR parser to parse the single frame pattern
    ParseResult result;
    if (!parseFileSequence(result, path)) {
        return false;
    }

    // Accept single frame files or plain files (depending on require_frame)
    if (result.isSequence) {
        return false;  // Don't accept sequences
    }

    if (require_frame && result.isPlainFile) {
        return false;  // Require a frame number
    }

    // Convert ParseResult to SeqPatternMatch
    match.base = result.directory + result.basename;
    match.range = result.frameRange;
    match.padChars = result.padding;
    match.ext = result.extension;

    return true;
}


size_t getPadSize(const std::string &pad, PadSyntax syntax) {
    // Helper function to parse padding size from alternate formats
    // This doesn't use ANTLR - it's just simple regex parsing

    static const auto flags = std::regex_constants::optimize|std::regex_constants::ECMAScript;

    switch (syntax) {
        case PadSyntaxPrintf: {
            // %04d, %4d, %d
            static const std::regex rx(R"(%0*(\d+)d)", flags);
            std::smatch match;
            if (std::regex_match(pad, match, rx)) {
                return std::stoul(match[1].str());
            }
            if (pad == "%d") {
                return 1;
            }
            break;
        }
        case PadSyntaxHoudini: {
            // $F4, $F04, $F
            static const std::regex rx(R"(\$F0*(\d+))", flags);
            std::smatch match;
            if (std::regex_match(pad, match, rx)) {
                return std::stoul(match[1].str());
            }
            if (pad == "$F") {
                return 1;
            }
            break;
        }
        case PadSyntaxUdim: {
            // <UDIM> or %(UDIM)d
            if (pad == "<UDIM>" || pad == "%(UDIM)d") {
                return 4;  // UDIM is always 4 digits
            }
            break;
        }
    }

    return 0;
}


} // internal
} // fileseq
