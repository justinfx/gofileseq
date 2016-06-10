
#include "frameset_p.h"

#include "../types.h"

#include <iomanip>
#include <sstream>
#include <string>
#include <vector>

namespace fileseq {

namespace internal {


bool getRangePatternMatch(RangePatternMatch &match, const std::string &range) {
    // Frame range:  1-10
    static const pcrecpp::RE* s_rxRange1 = new pcrecpp::RE("^(-?\\d+)-(-?\\d+)$");
    // Single frame:  10
    static const pcrecpp::RE* s_rxRange2 = new pcrecpp::RE("^(-?\\d+)$");
    // Complex range:  1-10x2
    static const pcrecpp::RE* s_rxRange3 = new pcrecpp::RE("^(-?\\d+)-(-?\\d+)([:xy])(\\d+)$");

    if ( s_rxRange1->FullMatch(range, &(match.start), &(match.end)) ) {
        match.matches = 2;
        return true;
    }

    if ( s_rxRange2->FullMatch(range, &(match.start)) ) {
        match.matches = 1;
        return true;
    }

    if ( s_rxRange3->FullMatch(
        range,
        &(match.start), &(match.end),
        &(match.stepMod), &(match.step)) ) {

        match.matches = 4;
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


std::string zfill(const std::string &src, size_t z) {
    size_t size = src.size();
    if (size >= z) {
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


std::string zfill(Frame value, size_t z) {
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

    if (s.find_first_of(mods) == std::string::npos) {
        return false;
    }

    return true;
}

} // internal
} // fileseq
