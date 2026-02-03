
#include "frameset_p.h"

#include "../error.h"
#include "../pad.h"

#include <iomanip>
#include <sstream>
#include <string>
#include <vector>

#include <regex>

namespace fileseq {

namespace internal {

bool getRangePatternMatch(RangePatternMatch &match, const std::string &range) {
    // Frame range:  1-10
    static const char* s_pattern1 = "^(-?\\d+)-(-?\\d+)$";
    // Single frame:  10
    static const char* s_pattern2 = "^(-?\\d+)$";
    // Complex range:  1-10x2
    static const char* s_pattern3 = R"(^(-?\d+)-(-?\d+)([:xy])(-?\d+)$)";

    match.stepMod.clear();
    match.step = 1;

    static const auto flags = std::regex_constants::optimize|std::regex_constants::ECMAScript;
    static const std::regex s_rxRange1(s_pattern1, flags);
    static const std::regex s_rxRange2(s_pattern2, flags);
    static const std::regex s_rxRange3(s_pattern3, flags);

    std::smatch submatch;

    if ( std::regex_match(range, submatch, s_rxRange1) ) {
        match.matches = 2;
        match.start = std::stol(submatch[1].str(), nullptr);
        match.end = std::stol(submatch[2].str(), nullptr);
        return true;
    }

    if ( std::regex_match(range, submatch, s_rxRange2) ) {
        match.matches = 1;
        match.start = std::stol(submatch[1].str(), nullptr);
        match.end = match.start;
        return true;
    }

    if ( std::regex_match(range, submatch, s_rxRange3) ) {
        match.matches = 4;
        match.start = std::stol(submatch[1].str(), nullptr);
        match.end = std::stol(submatch[2].str(), nullptr);
        match.stepMod = submatch[3].str();
        match.step = std::stol(submatch[4].str(), nullptr);
        return true;
    }

    match.matches = 0;
    return false;
}


Status frameRangeMatches(RangeMatches &matches, std::string range) {
    static const PaddingMapper &mapper = getPadMapperForStyle(PadStyleDefault);

    Status stat;

    matches.clear();

    std::string chars = mapper.getAllChars();
    chars.append(" "); // whitespace

    // Remove all padding characters and whitespace from string
    size_t found;
    std::string::const_iterator it;
    for (it = chars.begin(); it != chars.end(); ++it) {

        while (found = range.find(*it), found != std::string::npos) {
            range.erase(found, 1);
        }
    }

    std::string part;
    std::stringstream ss(range);

    RangePatternMatch rxMatch;

    // Split on comma
    while (std::getline(ss, part, ',')) {

        // Build up frames for all comma-sep components
        if (!getRangePatternMatch(rxMatch, part)) {

            std::stringstream err;
            err << "Failed to parse frame range: " << range << " on part " << part;
            stat.setError(err.str());
            return stat;
        }

        matches.push_back(rxMatch);
    }

    return stat;
}


std::string zfill(const std::string &src, int z) {
    size_t size = src.size();
    if (size >= size_t(z)) {
        return src;
    }

    // Size difference that we need to pad
    std::string fill(z-size, '0');

    std::stringstream ss;
    if (size > 0 && src[0] == '-') {
        // Pad negative numbers
        ss << "-" << fill << src.substr(1, std::string::npos);
    } else {
        // Pad positive numbers
        ss << fill << src;
    }
    return ss.str();
}


std::string zfill(Frame value, int z) {
    std::stringstream ss;

    ss << std::setfill('0')
       << std::setw(z)
       << std::internal
       << value;

    return ss.str();
}


bool isRangeModifier(const std::string &s) {
    if (s.size() != 1) {
        return false;
    }

    static const std::string mods("xy:");
    return s.find_first_of(mods) != std::string::npos;
}

} // internal
} // fileseq
