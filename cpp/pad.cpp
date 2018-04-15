#include "pad.h"
#include "private/frameset_p.h"

#include <algorithm>
#include <sstream>
#include <vector>

namespace fileseq {


std::string padFrameRange(const std::string &frange, size_t pad) {
    // We don't need to do anything if they gave us
    // an invalid pad number
    if (pad < 2) {
        return frange;
    }

    size_t size = std::count(frange.begin(), frange.end(), ',') + 1;
    std::vector<std::string> parts;
    parts.reserve(size);

    internal::RangePatternMatch rxMatch;

    std::string part;
    std::stringstream ss(frange);

    std::stringstream rangeFormatter;

    // Split on comma
    while (std::getline(ss, part, ',')) {

        // If we didn't match one of our expected patterns
        // then just take the original part and add it unmodified
        if (!internal::getRangePatternMatch(rxMatch, part)) {
            parts.push_back(part);
            continue;
        }

        if (rxMatch.matches == 0) {
            continue;
        }

        switch (rxMatch.matches) {

        case 1:
            // Single frame match:  10
            parts.push_back(internal::zfill(rxMatch.start, (int)pad));
            break;

        case 2:
            // Simple frame range match: 1-10
            rangeFormatter.str(std::string());
            rangeFormatter << internal::zfill(rxMatch.start, (int)pad)
                           << "-"
                           << internal::zfill(rxMatch.end, (int)pad);

            parts.push_back(rangeFormatter.str());
            break;

        case 4:
            // Frame range with stepping:  1-10x2
            rangeFormatter.str(std::string());
            rangeFormatter << internal::zfill(rxMatch.start, (int)pad)
                           << "-"
                           << internal::zfill(rxMatch.end, (int)pad)
                           << rxMatch.stepMod
                           << rxMatch.step;

            parts.push_back(rangeFormatter.str());
            break;

        }

    }

    // Join matches back into string
    std::stringstream out;
    for (size_t i=0; i < parts.size(); ++i) {
        if (i != 0) {
            out << ",";
        }
        out << parts[i];
    }
    return out.str();
}


std::string getPaddingChars(size_t width) {
    return internal::getPadMapperForStyle(PadStyleDefault).getPaddingChars(width);
}


namespace internal {


const PaddingMapper& getPadMapperForStyle(PadStyle style) {
    static const SingleHashPad s_singleHash = SingleHashPad();
    static const MultiHashPad s_multiHash = MultiHashPad();

    switch (style) {

    case PadStyleHash1:
        return s_singleHash;

    case PadStyleHash4:
        return s_multiHash;

    }

    // Should not get here. But sanity-check default..
    return s_multiHash;
}


// PaddingMapper
//

std::string PaddingMapper::getAllChars() const {
    std::stringstream ss;
    CharSizeMap::const_iterator it;
    for (it = m_charToSize.begin(); it != m_charToSize.end(); ++it) {
        ss << it->first;
    }
    return ss.str();
}

size_t PaddingMapper::getPaddingCharsSize(const std::string &chars) const {
    size_t size = 0;

    CharSizeMap::const_iterator found;
    std::string::const_iterator strIt;

    for (strIt = chars.begin(); strIt != chars.end(); ++strIt) {
        found = m_charToSize.find(*strIt);
        if (found != m_charToSize.end()) {
            size += found->second;
        }
    }

    return size;
}


// SingleHashPad
//

SingleHashPad::SingleHashPad() : PaddingMapper() {
    m_defaultChar = '#';
    m_charToSize['@'] = 1;
    m_charToSize['#'] = 1;
}

std::string SingleHashPad::getPaddingChars(long width) const {
    if (width <= 0) {
        return std::string(1, m_defaultChar);
    }

    std::stringstream ss;
    for (long i=0; i < width; ++i) {
        ss << m_defaultChar;
    }
    return ss.str();
}


// MultiHashPad
//

MultiHashPad::MultiHashPad() : PaddingMapper() {
    m_defaultChar = '@';
    m_charToSize['@'] = 1;
    m_charToSize['#'] = 4;
}

std::string MultiHashPad::getPaddingChars(long width) const {
    if (width <= 0) {
        return std::string(1, m_defaultChar);
    }
    if (width % 4 == 0) {
        return std::string(width/4, '#');
    }
    return std::string(width, '@');
}


} // internal

} // fileseq
