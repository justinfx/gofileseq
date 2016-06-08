
#include "ranges.h"
#include "util.h"

#include <cmath>
#include <algorithm>
#include <functional>
#include <sstream>


namespace fileseq {


Range::Range(long start, long end, int step)
    : m_start(start)
    , m_end(end)
    , m_step(step)
    , m_cachedEnd(0)
    , m_isEndCached(false)
    , m_cachedSize(0)
    , m_isSizeCached(false)
{

    if (m_step == 0) {
        m_step = (m_start <= m_end) ? 1 : -1;
    }

}


std::string Range::string() const {
    std::stringstream buf;

    // Always for a single value
    buf << start();

    // If we have a range, express the end value
    if (end() != start()) {
        buf << "-" << end();

        // Express the stepping, if its not 1
        if (step() > 1 || step() < -1) {
            buf << "x" << step();
        }
    }

    return buf.str();
}


long Range::end() const {
    if (m_isEndCached) {
        return m_cachedEnd;
    }

    m_isEndCached = true;

    // If we aren't stepping, or we don't have
    // a full range, then just use the end value
    if (m_step == 1 || m_step == -1 || m_start == m_end) {
        m_cachedEnd = m_end;
        return m_cachedEnd;
    }

    // If the step is in the wrong direction,
    // compared to the range direction, then
    // just use the start as the end.
    if (m_end < m_start && m_step < m_end-m_start) {
        m_cachedEnd = m_start;
        return m_cachedEnd;

    } else if (m_end > m_start && m_step > m_end-m_start) {
        m_cachedEnd = m_start;
        return m_cachedEnd;
    }

    // Calculate the end, taking into account the stepping
    m_cachedEnd = closestInRange(m_end, m_start, m_end, m_step);
    return m_cachedEnd;
}


size_t Range::length() const {
    if (m_isSizeCached) {
        return m_cachedSize;
    }

    // Offset by one to include the end value
    double diff = std::abs(double(m_end-m_start)) + 1;
    m_cachedSize = long(std::ceil(diff / std::abs(double(m_step))));
    m_isSizeCached = true;
    return m_cachedSize;
}


bool Range::contains(long value) const {
    // If we attempt to find the closest value, given
    // the start of the range and the step, we can check
    // if it is still the same number. If it hasn't changed,
    // then it is in the range.
    long closest = closestInRange(value, start(), end(), step());
    return (closest == value);
}


long Range::closestInRange(long value, long start, long end, int step) const {
    // Possibly clamp the value if it is outside the range
    if (end >= start) {
        if (value < start) {
            return start;
        } else if (value > end) {
            return end;
        }

    } else {
        if (value > start) {
            return start;
        } else if (value < end) {
            return end;
        }
    }

    // No calculation needed if there is no stepping
    if (step == 1 || step == -1) {
        return value;
    }

    // Modified the value so that it is a properly stepped
    // increment within the range
    return (((value - start) / step) * step) + start;
}


long Range::value(size_t idx, Status* ok) const {
    if (ok != NULL) ok->clearError();

    // Calculate the value as an offset from the start
    long l_start = start();
    long l_end = end();
    int l_step = step();

    long val = l_start + (l_step * idx);

    if (l_start <= l_end && (val < l_start || val > l_end)) {
        std::stringstream ss;
        ss << "Index " << idx << " exceeds max index of " << length()-1;
        internal::handleErrorStatus(ss.str(), ok);
        return 0;

    } else if (l_end < l_start && (val > l_start || val < l_end)) {
        std::stringstream ss;
        ss << "Index " << idx << " exceeds max index of " << length()-1;
        internal::handleErrorStatus(ss.str(), ok);
        return 0;
    }

    return val;
}


ssize_t Range::index(long value) const {
    long closest = closestInRange(value, start(), end(), step());
    if (closest != value) {
        return -1;
    }
    ssize_t idx = (value - start()) / step();
    if (idx < 0) {
        idx *= -1;
    }
    return idx;
}


RangeIterator::RangeIterator(const Range* range)
    : m_range(range)
    , m_index(-1) {

}


void RangeIterator::reset(const Range* range) {
    m_range = range;
    m_index = -1;
}


bool RangeIterator::isDone() {
    if (m_range == NULL) {
        return true;
    }

    if (m_index >= ssize_t(m_range->length())) {
        // At the end of the range
        return true;
    }

    return false;
}


long RangeIterator::operator*() {
    if (m_index == -1) {
        // Auto-advance from a new iterator state
        m_index = 0;
    }
    return m_range->value(m_index);
}


bool RangeIterator::next() {
    m_index += 1;
    if (isDone()) {
        return false;
    }
    return true;
}


Ranges::~Ranges() {
    // Delete the range list
    while (!m_blocks.empty()) delete m_blocks.back(), m_blocks.pop_back();
}


std::string Ranges::string() const {
    std::stringstream buf;

    bool first = true;
    Blocks::const_iterator it;
    for (it = m_blocks.begin(); it != m_blocks.end(); ++it) {
        if (!first) {
            buf << ",";
        }
        buf << (*it)->string();
        first = false;
    }
    return buf.str();
}


size_t Ranges::length() const {
    size_t total = 0;
    Blocks::const_iterator it;
    for (it = m_blocks.begin(); it != m_blocks.end(); ++it) {
        total += (*it)->length();
    }
    return total;
}


long Ranges::start() const {
    if (m_blocks.empty()) {
        return 0;
    }
    return m_blocks.front()->start();
}


long Ranges::end() const {
    if (m_blocks.empty()) {
        return 0;
    }
    return m_blocks.back()->end();
}


long Ranges::min() const {
    long val = start();
    Blocks::const_iterator it;
    for (it = m_blocks.begin(); it != m_blocks.end(); ++it) {
        val = std::min((*it)->min(), val);
    }
    return val;
}


long Ranges::max() const {
    long val = end();
    Blocks::const_iterator it;
    std::cerr << "looping range: " << string() << " (# blocks: " << m_blocks.size() << ")" << std::endl;
    for (it = m_blocks.begin(); it != m_blocks.end(); ++it) {
        std::cerr << "  block: " << (*it)->string() << std::endl;
        val = std::max((*it)->min(), val);
    }
    return val;
}


bool predIncRangeDone(long val, long end) {
    return val <= end;
}

bool predDecRangeDone(long val, long end) {
    return val >= end;
}

void Ranges::appendUnique(long start, long end, int step) {
    if (step == 0) {
        // Invalid step. Do nothing.
        return;
    }

    long subStart = start;
    long subEnd = start;
    long subStep = step;
    long last = start;
    size_t pending = 0; // Track unique value count

    // Handle loop test for both increasing
    // and decreasing ranges
    bool (*pred) (long, long);

    if (start <= end) {
        if (step < 0) {
            step *= -1;
        }
        pred = &predIncRangeDone;

    } else {
        if (step > 0) {
            step *= -1;
        }
        pred = &predDecRangeDone;
    }

    // Short-circuit if this is the first range being added
    if (m_blocks.empty()) {
        append(start, end, step);
        return;
    }

    // TODO: More intelligent fast-paths for easy-to-identify
    // overlapping ranges. Such as when the existing range is:
    // 1-100x1 and we are appending 50-150x1. Should be easy
    // enough to just know we can Append(101,150,1)

    for ( ; pred(subEnd, end); subEnd += step ) {

        if (!contains(subEnd)) {
            // Is a unique value in the range
            last = subEnd;
            if (pending == 0) {
                subStart = last;
            }
            pending++;
            continue;
        }

        if (pending == 0) {
            // Nothing to add yet
            continue;
        }

        // Current value is already in range.
        // Add previous values
        append(subStart, last, subStep);
        subStart = subEnd + step;
        pending = 0;
    }

    // Flush the remaining values
    if (pending > 0) {
        append(subStart, last, subStep);
    }
}


bool Ranges::contains(long value) const {
    Blocks::const_iterator it;
    for (it = m_blocks.begin(); it != m_blocks.end(); ++it) {
        if ((*it)->contains(value)) {
            return true;
        }
    }
    return false;
}


long Ranges::value(size_t idx, Status* ok) const {
    long val;
    size_t size, n;
    Range* range;
    Status found;
    Blocks::const_iterator it;

    n = 0;

    if (ok != NULL) ok->clearError();

    for (it = m_blocks.begin(); it != m_blocks.end(); ++it) {
        range = (*it);
        size = range->length();

        // modify the original index to be within
        // the range of the current range block, and
        // then check that it is a valid index value
        if ((idx - n) < size) {
            val = range->value(idx - n, &found);
            if (found) {
                return val;
            }
        }

        // record the offset so we can continue to
        // modify the original index value to the local range
        n += size;
    }

    // The previous loop ended in error
    std::stringstream ss;
    ss << "Index " << idx << " exceeds the total range";
    internal::handleErrorStatus(ss.str());

    return 0;
}


ssize_t Ranges::index(long value) const {
    ssize_t idx;
    size_t n = 0;
    Range* range;
    Blocks::const_iterator it;

    for (it = m_blocks.begin(); it != m_blocks.end(); ++it) {
        range = (*it);
        // If the value is within the current block
        // then return the local index, offset by the
        // number of previous values we have tracked
        if ((idx = range->index(value)) >= 0) {
            return idx + n;
        }
        // Update the offset for the values we have seen
        n += range->length();
    }

    // The previous loop ended in error
    return -1;
}


void Ranges::normalized(Ranges &outRanges, bool invert) const {
    long start = 0;
    long end = 0;
    long current = 0;
    int step = 0;
    bool keepValue = false;
    size_t pending = 0;

    // Accumulate temp results here
    Ranges normalizedRange;

    Range totalRange(min(), max(), 1);
    std::cerr << "DEBUG: " << string() << "  ;  " << totalRange.string() << std::endl;
    RangeIterator rangeIt = totalRange.iterValues();

    while (rangeIt.next()) {

        current = (*rangeIt);

        keepValue = contains(current);
        if (!invert) {
            keepValue = (!keepValue);
        }

        // Short-circuit if we encounter a value that
        // is not in the original sequence.
        if (keepValue) {

            // If we haven't accumulated 2+ values to
            // add, just continue now and keep trying
            if (pending < 2) {
                step++;
                continue;
            }

            // If the step has changed from what we have
            // already accumulated, then add what we have
            // and start a new range.
            if ((current + 1 - end) != step) {
                normalizedRange.append(start, end, step);
                step = 1;
                start = current;
                pending = 0;
            }

            continue;
        }

        // If the value we want to keep is a different
        // stepping from the pending values, add what
        // we have and start a new range again.
        if (pending >= 2 && current-end != step) {
            normalizedRange.append(start, end, step);
            pending = 0;
        }

        end = current;

        // Start a new range
        if (pending == 0) {
            start = end;
            step = 1;
        }

        pending++;
        continue;
    }

    // Flush the remaining values
    if (pending > 0) {
        normalizedRange.append(start, end, step);
    }

    // Update the callers object data
    outRanges.m_blocks.swap(normalizedRange.m_blocks);
}


RangesIterator::RangesIterator(const Ranges* ranges)
    : m_ranges(ranges)
    , m_current()
    , m_index(0) {

}


void RangesIterator::reset(const Ranges* ranges) {
    m_ranges = ranges;
    m_current.reset(NULL);
    m_index = 0;
}


bool RangesIterator::ready() {
    if (!isValid()) {
        return false;
    }

    if (!m_current.isValid()) {
        const Range* range = m_ranges->rangeAt(m_index);
        if (range == NULL) {
            return false;
        }
        m_current = range->iterValues();
    }
    return true;
}

bool RangesIterator::isDone() {
    if (!ready()) {
        return true;
    }

    return m_current.isDone();
}


long RangesIterator::operator*() {
    if (!ready()) {
        return m_ranges->end();
    }

    long val = (*m_current);
    return val;
}


bool RangesIterator::next() {
    if (!ready()) {
        return false;
    }

    if (!m_current.next()) {
        // Get the next available range
        m_index++;
        m_current.reset(NULL);
    }

    return ready();
}


} // fileseq
