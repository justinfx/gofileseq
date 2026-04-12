
#include "ranges.h"
#include "../error.h"

#include <cmath>
#include <algorithm>
#include <sstream>


namespace fileseq {


// RangeInterval is a simple [lo, hi] interval used internally
// for coverage and overlap calculations in appendUnique.
struct RangeInterval {
    long lo;
    long hi;
};


Range::Range(long start, long end, long step)
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


long Range::closestInRange(long value, long start, long end, long step) const {
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
    long l_step = step();

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
    for (auto* b : m_blocks) {
        delete b;
    }
}


Ranges::Ranges(const Ranges& other) : m_blocks() {
    m_blocks.reserve(other.m_blocks.size());
    for (const auto* b : other.m_blocks) {
        m_blocks.push_back(new Range(*b));
    }
}


std::string Ranges::string() const {
    std::stringstream buf;
    bool first = true;
    for (const auto* b : m_blocks) {
        if (!first) {
            buf << ",";
        }
        buf << b->string();
        first = false;
    }
    return buf.str();
}


size_t Ranges::length() const {
    size_t total = 0;
    for (const auto* b : m_blocks) {
        total += b->length();
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
    for (const auto* b : m_blocks) {
        val = std::min(b->min(), val);
    }
    return val;
}


long Ranges::max() const {
    long val = end();
    for (const auto* b : m_blocks) {
        val = std::max(b->max(), val);
    }
    return val;
}


// allBlocksContiguous returns true if all blocks have a step of 1 or -1.
static bool allBlocksContiguous(const std::vector<Range*>& blocks) {
    for (const auto* b : blocks) {
        long s = b->step();
        if (s != 1 && s != -1) {
            return false;
        }
    }
    return true;
}


// mergedCoverage returns a sorted, non-overlapping list of coverage intervals
// built from the given blocks.
static std::vector<RangeInterval> mergedCoverage(const std::vector<Range*>& blocks) {
    std::vector<RangeInterval> intervals;
    intervals.reserve(blocks.size());

    // Because all blocks have a step of 1 or -1, min and max describe their
    // full coverage with no gaps, so a simple [min, max] interval is enough.
    for (const auto* b : blocks) {
        intervals.push_back({b->min(), b->max()});
    }

    // Sort so we can merge in a single left-to-right pass.
    std::sort(intervals.begin(), intervals.end(),
        [](const RangeInterval& a, const RangeInterval& b) {
            return a.lo < b.lo;
        });

    // Walk the sorted list, extending the last interval when the next one
    // overlaps or is adjacent, otherwise starting a new one.
    std::vector<RangeInterval> merged;
    merged.reserve(intervals.size());
    for (const auto& iv : intervals) {
        if (merged.empty() || iv.lo > merged.back().hi + 1) {
            merged.push_back(iv);
        } else if (iv.hi > merged.back().hi) {
            merged.back().hi = iv.hi;
        }
    }
    return merged;
}


// gapsInRange returns the portions of [lo, hi] not covered by the given
// coverage intervals, in ascending order.
static std::vector<RangeInterval> gapsInRange(long lo, long hi,
                                               const std::vector<RangeInterval>& coverage) {
    std::vector<RangeInterval> gaps;
    long cur = lo;
    for (const auto& iv : coverage) {
        if (iv.hi < cur) {
            continue; // entirely before our position
        }
        if (iv.lo > hi) {
            break; // entirely beyond our range
        }
        if (iv.lo > cur) {
            gaps.push_back({cur, iv.lo - 1});
        }
        cur = iv.hi + 1;
    }
    if (cur <= hi) {
        gaps.push_back({cur, hi});
    }
    return gaps;
}


// appendUniqueContiguous appends only the portions of [lo, hi] not already
// covered by existing blocks, using interval arithmetic rather than per-value
// iteration. Only called when all existing blocks and the new range have a
// step of 1 or -1.
static void appendUniqueContiguous(Ranges* self,
                                    const std::vector<RangeInterval>& coverage,
                                    long lo, long hi, long step) {
    const auto gaps = gapsInRange(lo, hi, coverage);

    // For a descending range, append from high to low so that the
    // block order matches the direction of iteration.
    if (step < 0) {
        for (auto it = gaps.rbegin(); it != gaps.rend(); ++it) {
            self->append(it->hi, it->lo, step);
        }
        return;
    }
    for (const auto& g : gaps) {
        self->append(g.lo, g.hi, step);
    }
}


void Ranges::appendUnique(long start, long end, long step) {
    if (step == 0) {
        return;
    }

    // Normalize step direction to match the start/end direction.
    if (start <= end && step < 0) {
        step = -step;
    } else if (start > end && step > 0) {
        step = -step;
    }

    // Short-circuit if this is the first range being added.
    if (m_blocks.empty()) {
        append(start, end, step);
        return;
    }

    // Value bounds for the new range (direction-independent).
    long lo = std::min(start, end);
    long hi = std::max(start, end);

    // Fast path: new range has no overlap with existing ranges at all.
    if (lo > max() || hi < min()) {
        append(start, end, step);
        return;
    }

    long absStep = std::abs(step);

    // Fast path for contiguous new ranges (step of 1 or -1) when all existing
    // blocks are also contiguous: compute non-overlapping sub-ranges via
    // interval arithmetic, with no per-value iteration required.
    if (absStep == 1 && allBlocksContiguous(m_blocks)) {
        appendUniqueContiguous(this, mergedCoverage(m_blocks), lo, hi, step);
        return;
    }

    // General slow path: iterate each value in the new range and check
    // uniqueness individually.
    //
    // Snapshot the current block count before the loop so that any sub-ranges
    // appended mid-iteration are not included in subsequent uniqueness checks.
    const size_t originalSize = m_blocks.size();
    auto containsInOriginal = [&](long value) -> bool {
        for (size_t i = 0; i < originalSize; ++i) {
            if (m_blocks[i]->contains(value)) {
                return true;
            }
        }
        return false;
    };

    long subStart = start;
    long subEnd = start;
    long last = start;
    size_t pending = 0;

    bool (*pred)(long, long) = (start <= end)
        ? [](long val, long end) -> bool { return val <= end; }
        : [](long val, long end) -> bool { return val >= end; };

    for ( ; pred(subEnd, end); subEnd += step ) {

        if (!containsInOriginal(subEnd)) {
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
        append(subStart, last, step);
        subStart = subEnd + step;
        pending = 0;
    }

    // Flush the remaining values
    if (pending > 0) {
        append(subStart, last, step);
    }
}


bool Ranges::contains(long value) const {
    for (const auto* b : m_blocks) {
        if (b->contains(value)) {
            return true;
        }
    }
    return false;
}


long Ranges::value(size_t idx, Status* ok) const {
    size_t n = 0;

    if (ok != NULL) ok->clearError();

    for (const auto* b : m_blocks) {
        size_t size = b->length();

        // modify the original index to be within
        // the range of the current range block, and
        // then check that it is a valid index value
        if ((idx - n) < size) {
            Status found;
            long val = b->value(idx - n, &found);
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
    size_t n = 0;
    for (const auto* b : m_blocks) {
        // If the value is within the current block
        // then return the local index, offset by the
        // number of previous values we have tracked
        ssize_t idx = b->index(value);
        if (idx >= 0) {
            return idx + n;
        }
        // Update the offset for the values we have seen
        n += b->length();
    }

    // The previous loop ended in error
    return -1;
}


void Ranges::normalized(Ranges &outRanges, bool invert) const {
    long start = 0;
    long end = 0;
    long current = 0;
    long step = 0;
    bool keepValue = false;
    size_t pending = 0;

    // Accumulate temp results here
    Ranges normalizedRange;

    Range totalRange(min(), max(), 1);
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
