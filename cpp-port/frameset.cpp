#include "frameset.h"
#include "types.h"

#include "private/frameset_p.h"

#include <iostream>
#include <sstream>

namespace fileseq {


FrameSet::FrameSet() : m_frameData(NULL) {

}


FrameSet::FrameSet(const std::string &frange, Status* ok)
    : m_frameData(new internal::FrameSetData(frange))
{
    Status stat;
    internal::RangeMatches matches;

    // Process the frame range and get a list of matches
    if (!(stat = internal::frameRangeMatches(matches, frange))) {
        if (ok != NULL) {
            *ok = stat;
        }
        setInvalid();
        return;
    }

    // Process each list of matches and add it to the frame set
    for (size_t i=0; i < matches.size(); ++i) {

        // If handleMatch succeeds, it will populate the FrameSet
        // with more parts of the range
        if (handleMatch(&matches[i], &stat), !stat) {

            // We could not process a match for some reason
            if (ok != NULL) {
                *ok = stat;
            }
            setInvalid();
            return;
        }
    }

    // A valid frameset
    if (ok != NULL) ok->clearError();

    // Save the original frame range string
    m_frameData->frameRange = frange;
}


FrameSet::~FrameSet() {
    setInvalid();
}

void FrameSet::setInvalid() {
    if ( m_frameData ) {
        delete m_frameData;
        m_frameData = NULL;
    }
}


void FrameSet::handleMatch(const internal::RangePatternMatch* match, Status* ok) {
    const int num = match->matches;

    ok->clearError();

    Frame start, end;
    size_t step = 1;

    switch (num) {

    // Single frame match
    case 1:
        start = match->start;
        m_frameData->ranges.appendUnique(start, start, 1);
        return;

    // Simple frame range
    case 2:
        start = match->start;
        end = match->end;

        // Handle descending frame ranges, like 10-1
        step = (start > end) ? -1 : 1;

        m_frameData->ranges.appendUnique(start, end, step);
        return;

    // Complex frame range
    case 4:
        start = match->start;
        end = match->end;
        step = match->step;

        if (step == 0) {
            ok->setError("Got invalid step value 0");
            return;
        }

        if (!internal::isRangeModifier(match->stepMod)) {
            std::stringstream ss;
            ss << match->stepMod << " is not one of the valid modifiers 'xy:'";
            ok->setError(ss.str());
        }

        char mod = match->stepMod[0];

        switch (mod) {

        case 'x':
            m_frameData->ranges.appendUnique(start, end, step);
            break;

        case 'y': {
            // TODO: Add proper support for adding inverse of range.
            // This approach will add excessive amounts of singe
            // range elements. They could be compressed into chunks
            Frame skip = start;
            Range aRange(start, end, 1);
            RangeIterator it = aRange.iterValues();
            Frame val;
            while (it.next()) {
                val = (*it);
                if (val == skip) {
                    skip += step;
                    continue;
                }
                m_frameData->ranges.appendUnique(val, val, 1);
            }
            break;
        }

        case ':':
            while (step > 0) {
                m_frameData->ranges.appendUnique(start, end, step);
                step--;
            }
            break;

        }

        return;
    }

    // If we get here, we  have not handled the match
    std::stringstream ss;
    ss << "Unexpected match size: " << num;
    ok->setError(ss.str());
}


// FrameSet::FrameSet(const FrameSet& rhs) : m_valid(false) {

// }

// FrameSet& FrameSet::operator=(const FrameSet& rhs) {

// }

bool FrameSet::isValid() const {
    return m_frameData != NULL;
}

std::string FrameSet::string() const {
    return isValid() ? m_frameData->frameRange : "";
}

size_t FrameSet::length() const {
    return isValid() ? m_frameData->ranges.length() : 0;
};

size_t FrameSet::index(Frame frame) const {
    return isValid() ? m_frameData->ranges.index(frame) : 0;
};

Frame FrameSet::frame(size_t index, Status* ok) const {
    if (!isValid()) {
        return 0;
    }

    return m_frameData->ranges.value(index, ok);
};

void FrameSet::frames(Frames &frames) const {
    frames.clear();

    if (!isValid()) {
        return;
    }

    size_t len = length();
    if (len <= 0) {
        return;
    }

    frames.reserve(len);

    RangesIterator it = m_frameData->ranges.iterValues();
    while (it.next()) {
        frames.push_back(*it);
    }
};

RangesIterator FrameSet::iterFrames() const {
    if (!isValid()) {
        return RangesIterator();
    }

    return m_frameData->ranges.iterValues();
}

bool FrameSet::hasFrame(Frame frame) const {
    return isValid() ? m_frameData->ranges.contains(frame) : false;
};

Frame FrameSet::start() const {
    return isValid() ? m_frameData->ranges.start() : 0;
};

Frame FrameSet::end() const {
    return isValid() ? m_frameData->ranges.end() : 0;
};

std::string FrameSet::frameRange(int pad) const {
    if (!isValid()) {
        return "";
    }

    if (pad < 2) {
        return m_frameData->frameRange;
    }

    return padFrameRange(m_frameData->frameRange, pad);
};

FrameSet FrameSet::inverted() const {
    FrameSet newFrameSet;

    if (!isValid()) {
        return newFrameSet;
    }

    // get a new range that is inverted
    Ranges ranges;
    m_frameData->ranges.inverted(ranges);

    newFrameSet.m_frameData = new internal::FrameSetData();
    newFrameSet.m_frameData->ranges = ranges;
    newFrameSet.m_frameData->frameRange = ranges.string();

    return newFrameSet;
};

std::string FrameSet::invertedFrameRange(int pad) const {
    if (!isValid()) {
        return "";
    }

    Ranges ranges;
    m_frameData->ranges.inverted(ranges);

    std::string frange = ranges.string();

    if (pad > 1) {
        frange = padFrameRange(frange, pad);
    }

    return frange;
};

FrameSet FrameSet::normalized() const {
    FrameSet newFrameSet;

    if (!isValid()) {
        return newFrameSet;
    }

    // get a new range that is normalized
    Ranges ranges;
    m_frameData->ranges.normalized(ranges);

    newFrameSet.m_frameData = new internal::FrameSetData();
    newFrameSet.m_frameData->ranges = ranges;
    newFrameSet.m_frameData->frameRange = ranges.string();

    return newFrameSet;
};

} // fileseq
