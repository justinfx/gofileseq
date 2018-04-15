
#include "sequence_p.h"

#include <string>
#include <vector>

// Conditional support for C++11 compilers that support std::regex
#if HAVE_REGEX == 1
#include <regex>
#else
#include <pcrecpp.h>
#endif

namespace fileseq {

namespace internal {


bool getSplitPatternMatch(SeqPatternMatch &match, const std::string &path) {
    // Regular expression for matching a file sequence string.
    // Example:
    //     /film/shot/renders/hero_bty.1-100#.exr
    //     /film/shot/renders/hero_bty.@@.exr
    static const char* s_pattern = R"(^(.*?)([\d-][:xy\d,-]*)?([#@]+)((?:\.[a-zA-Z0-9]+)*)$)";

    match.base.clear();
    match.range.clear();
    match.padChars.clear();
    match.ext.clear();

#if HAVE_REGEX == 1
    static const std::regex rx (s_pattern,
            std::regex_constants::optimize|std::regex_constants::ECMAScript);

    std::smatch submatch;

    if (!std::regex_match(path, submatch, rx)) {
        return false;
    }

    match.base = submatch[1].str();
    match.range = submatch[2].str();
    match.padChars = submatch[3].str();
    match.ext = submatch[4].str();
    return true;

#else
    static const pcrecpp::RE* rx = new pcrecpp::RE(s_pattern);
    return rx->FullMatch(path,
                         &(match.base), &(match.range),
                         &(match.padChars), &(match.ext));
#endif
}


bool getSingleFrameMatch(SeqPatternMatch &match, const std::string &path) {
    // Regular expression for matching a single frame string.
    // Example:
    //     /film/shot/renders/hero_bty.100.exr
    static const char* s_pattern = R"(^(.*?)(-?\d+)((?:\.\w*[a-zA-Z]\w*)*(?:\.[a-zA-Z0-9]+)?)$)";

    // Not using this field
    match.padChars.clear();

#if HAVE_REGEX == 1
    static const std::regex rx(s_pattern,
            std::regex_constants::optimize|std::regex_constants::ECMAScript);

    std::smatch submatch;

    if (!std::regex_match(path, submatch, rx)) {
        return false;
    }

    match.base = submatch[1].str();
    match.range = submatch[2].str();
    match.ext = submatch[3].str();
    return true;

#else
    static const pcrecpp::RE* rx = new pcrecpp::RE(s_pattern);
    return rx->FullMatch(path, &(match.base), &(match.range), &(match.ext));

#endif
}


} // internal

} // fileseq
