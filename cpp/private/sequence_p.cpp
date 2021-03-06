
#include "sequence_p.h"

#include <iostream>
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
    //     /film/shot/renders/hero_bty.%02d.exr
    //     /film/shot/renders/hero_bty.$F02.exr
    static const char* s_pattern =
            R"(^(.*?))" // dir and basename
            R"(([\d-][:xy\d,-]*)?)" // optional frame range
            R"(([#@]+|%\d*d|\$F\d*))" // padding: chars, printf, houdini
            R"(((?:\.[a-zA-Z0-9]+)*)$)"; // extension

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


size_t getPadSize(const std::string &pad, PadSyntax syntax) {
    // printf:  %04d
    static const char *s_printf_pattern = R"(^%(\d*)d$)";
    // houdini:  $F, $F4, $F04
    static const char *s_houdini_pattern = R"(^\$F(\d*)$)";

    if (pad.empty()) {
        return 0;
    }

    size_t val = 0;

#if HAVE_REGEX == 1
    static const auto flags = std::regex_constants::optimize|std::regex_constants::ECMAScript;
    static const std::regex s_printf_rx(s_printf_pattern, flags);
    static const std::regex s_houdini_rx(s_houdini_pattern, flags);

    const std::regex* rx;
    switch (syntax) {
        case PadSyntaxPrintf:
            rx = &s_printf_rx;
            break;
        case PadSyntaxHoudini:
            rx = &s_houdini_rx;
            break;
        default:
            std::cerr << "warning: unhandled fileseq pad format "
                         "syntax getPadSize(syntax=" << syntax << ")\n";
            return val;
    }

    std::smatch submatch;
    if (!std::regex_match(pad, submatch, *rx)) {
        return val;
    }
    val = !submatch[1].str().empty() ? std::stol(submatch[1].str()) : 1;

#else
    static const pcrecpp::RE* s_printf_rx = new pcrecpp::RE(s_printf_pattern);
    static const pcrecpp::RE* s_houdini_rx = new pcrecpp::RE(s_houdini_pattern);

    const pcrecpp::RE* rx;
    switch (syntax) {
        case PadSyntaxPrintf:
            rx = s_printf_rx;
            break;
        case PadSyntaxHoudini:
            rx = s_houdini_rx;
            break;
        default:
            std::cerr << "warning: unhandled fileseq pad format "
                         "syntax getPadSize(syntax=" << syntax << ")\n";
            return val;
    }

    std::string str;
    if (!(rx->FullMatch(pad, &str))) {
        return val;
    }
    val = !str.empty() ? std::stol(str) : 1;

#endif

    return std::max(val, size_t(1));
}


} // internal

} // fileseq
