
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
    //     /film/shot/renders/hero_bty.1-100#.tar.gz
    //     /film/shot/renders/hero_bty.@@.exr
    //     /film/shot/renders/hero_bty.%02d.exr
    //     /film/shot/renders/hero_bty.$F02.exr
    //     /film/shot/renders/hero_bty.<UDIM>.exr
    //     /film/shot/renders/hero_bty.%(UDIM)d.exr
    static const char* s_pattern =
            R"(^(.*?))" // dir and basename
            R"(([\d-][:xy\d,-]*)?)" // optional frame range
            R"(([#@]+|%\d*d|\$F\d*|<UDIM>|%\(UDIM\)d))" // padding: chars, printf, houdini, udim
            R"(((?:\.\w*[a-zA-Z]\w)*(?:\.[^.]+)?)$)"; // extension

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


bool getSingleFrameMatch(SeqPatternMatch &match, const std::string &path, bool require_frame) {
    using std::string;
    static const string s_extension =
        // multiple extension parts:
        "("
            R"((?:\.\w*[a-zA-Z]\w)*)"  // optional leading alnum ext prefix (.foo.1bar)
            R"((?:\.[^.]+)?)"          // ext suffix
        ")";

	// Regular expression pattern for matching single file path names containing a frame.
	// Example:
	//     /film/shot/renders/hero_bty.100.exr
    static const string s_singleFramePattern =
        string("^") +
            "(.*?)" +       // file name
            R"((-?\d+))" +  // frame
            s_extension +   // ext
        "$";

    // Regular expression pattern for matching single file path names where the
	// frame may be optional.
	// Example:
	//     /film/shot/renders/hero_bty.exr
    static const string s_optionalFramePattern =
        string("^") +
            "(.*?)" +        // file name
            R"((-?\d+)?)" +  // optional frame
            s_extension +    // ext
        "$";

    // Not using this field
    match.padChars.clear();

#if HAVE_REGEX == 1
    static const std::regex s_rxSingleFrame(s_singleFramePattern.c_str(),
            std::regex_constants::optimize|std::regex_constants::ECMAScript);

    static const std::regex s_rxOptionalFrame(s_optionalFramePattern.c_str(),
        std::regex_constants::optimize|std::regex_constants::ECMAScript);

    const std::regex &rx = require_frame ? s_rxSingleFrame : s_rxOptionalFrame;
    std::smatch submatch;

    if (!std::regex_match(path, submatch, rx)) {
        return false;
    }

    match.base = submatch[1].str();
    match.range = submatch[2].str();
    match.ext = submatch[3].str();
    return true;

#else
    static const pcrecpp::RE* s_rxSingleFrame = new pcrecpp::RE(s_singleFramePattern.c_str());
    static const pcrecpp::RE* s_rxOptionalFrame = new pcrecpp::RE(s_optionalFramePattern.c_str());

    const pcrecpp::RE* &rx = require_frame ? s_rxSingleFrame : s_rxOptionalFrame;
    return rx->FullMatch(path, &(match.base), &(match.range), &(match.ext));

#endif
}


size_t getPadSize(const std::string &pad, PadSyntax syntax) {
    // printf:  %04d
    static const char *s_printf_pattern = R"(^%(\d*)d$)";
    // houdini:  $F, $F4, $F04
    static const char *s_houdini_pattern = R"(^\$F(\d*)$)";
    // UDIM:  <UDIM>, %(UDIM)d
    static const char *s_udim_pattern = R"(^(<UDIM>|%\(UDIM\)d)$)";
    static const size_t udim_size = 4;

    if (pad.empty()) {
        return 0;
    }

    size_t val = 0;

#if HAVE_REGEX == 1
    static const auto flags = std::regex_constants::optimize|std::regex_constants::ECMAScript;
    static const std::regex s_printf_rx(s_printf_pattern, flags);
    static const std::regex s_houdini_rx(s_houdini_pattern, flags);
    static const std::regex s_udim_rx(s_udim_pattern, flags);

    const std::regex* rx;
    switch (syntax) {
        case PadSyntaxPrintf:
            rx = &s_printf_rx;
            break;
        case PadSyntaxHoudini:
            rx = &s_houdini_rx;
            break;
        case PadSyntaxUdim:
            rx = &s_udim_rx;
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
    // udim has a fixed size
    if (syntax == PadSyntaxUdim) {
        return udim_size;
    }
    val = !submatch[1].str().empty() ? std::stol(submatch[1].str()) : 1;

#else
    static const pcrecpp::RE* s_printf_rx = new pcrecpp::RE(s_printf_pattern);
    static const pcrecpp::RE* s_houdini_rx = new pcrecpp::RE(s_houdini_pattern);
    static const pcrecpp::RE* s_udim_rx = new pcrecpp::RE(s_udim_pattern);

    const pcrecpp::RE* rx;
    switch (syntax) {
        case PadSyntaxPrintf:
            rx = s_printf_rx;
            break;
        case PadSyntaxHoudini:
            rx = s_houdini_rx;
            break;
        case PadSyntaxUdim:
            rx = &s_udim_rx;
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
    // udim has a fixed size
    if (syntax == PadSyntaxUdim) {
        return udim_size;
    }
    val = !str.empty() ? std::stol(str) : 1;

#endif

    return std::max(val, size_t(1));
}


} // internal

} // fileseq
