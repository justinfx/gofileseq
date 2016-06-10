#ifndef FILESEQ_RANGES_P_H_
#define FILESEQ_RANGES_P_H_

#include <algorithm>
#include <ostream>
#include <string>
#include <vector>

namespace fileseq {

// Fwd Decl
class Range;
class Ranges;
class Status;


// tracks the iteration state across an Range instance
class RangeIterator {

public:
    explicit RangeIterator( const Range* range=NULL );

    // Reset the iterator with a new Range.
    void reset( const Range* range );

    // Dereference current value.
    long operator*();

    // Returns if the iterator has exhausted all
    // values, without consuming another value like
    // a call to next()
    bool isDone();

    // Advance to the next value.
    bool next();

    // Returns whether the interator points to a valid range
    bool isValid() const { return m_range != NULL; }

private:
    const Range* m_range;
    ssize_t m_index;
};


class Range {

public:

    // Range represents a start and end
    // value, and a stepping amount between each
    // value in the range. The range includes the end value.
    // It dynamically acts like a slice of values
    // in an integer range, without needing to store
    // the discreet values. So it can represent very
    // large ranges.
    //
    // If step value == 0, it will automatically default
    // to a propert increment depending on whether start is
    // less than or greather than end.
    Range(long start, long end, int step=0);

    virtual ~Range() {}

    // returns the number of values in the range
    size_t length() const;

    // returns the start of the range
    long start() const { return m_start; }

    // returns the end of the range. This value may
    // be different from the end value given when the
    // range was first initialized, since it takes into
    // account the stepping value. The end value may be
    // shifted to the closest valid value within the
    // stepped range.
    long end() const;

    // returns the stepping value used in the range
    int step() const { return m_step; }

    // returns the smallest value in the range
    long min() const { return (start() < end() ? start() : end()); }

    // returns the highest value in the range
    long max() const { return (start() > end() ? start() : end()); }

    // returns true if the given value is a valid
    // value within the value range.
    bool contains(long value) const;

    // returns the value at the given index in
    // the range. If the index is invalid/exceeds the valid
    // range, and a Status pointer is given, it will be filled
    // with the error state.
    long value(size_t idx, Status* ok=NULL) const;

    // returns the 0-based index of the first occurrence
    // given value, within the range.
    // If the value does not exist in the range, -1 will
    // be returned
    ssize_t index(long value) const;

    // returns an iterator that will loop through
    // every value in the range.
    // Call to RangeIterator.next() populates first
    // value and returns whether there are more values available.
    RangeIterator iterValues() const { return RangeIterator(this); }

    // The string representation of the range
    std::string string() const;

    operator std::string() const { return string(); }

    friend std::ostream& operator<< (std::ostream& stream, const Range& fs) {
        stream << fs.string();
        return stream;
    };

protected:

    // closestInRange finds the closest valid value within the range,
    // to a given value. Values outside the range are clipped to either
    // the range min or max.
    long closestInRange(long value, long start, long end, int step=0) const;

private:
    long m_start;
    long m_end;
    int m_step;

    mutable long m_cachedEnd;
    mutable bool m_isEndCached;

    mutable size_t m_cachedSize;
    mutable bool m_isSizeCached;
};


// tracks the iteration state across an Ranges instance
class RangesIterator {

public:
    explicit RangesIterator( const Ranges* ranges=NULL );

    // Reset the iterator with a new Ranges.
    void reset( const Ranges* ranges );

    // Dereference current value.
    long operator*();

    // Returns if the iterator has exhausted all
    // values, without consuming another value like
    // a call to next()
    bool isDone();

    // Advance to the next value and return whether
    // a value is available
    bool next();

    // Returns whether the interator points to a valid range
    bool isValid() const { return m_ranges != NULL; }

private:
    bool ready();

    const Ranges* m_ranges;
    RangeIterator m_current;
    size_t m_index;
};


// Ranges is able to represent multiple non-contiguous
// Range pointers, with non-uniform stepping, ie. 1-10,20-30x2,100,120
// Because range values are evaluated dynamically as
// offsets within the start/stop, they can represent
// very large ranges.
class Ranges {

public:
    Ranges() : m_blocks() {}

    ~Ranges();

    // Copy constructor
    Ranges(const Ranges& rhs);

    // Assignment
    Ranges& operator=(Ranges rhs) {
        swap(*this, rhs);
        return *this;
    }

    // Swap functionality
    friend void swap(Ranges &first, Ranges &second) {
        using std::swap;

        swap(first.m_blocks, second.m_blocks);
    }

    // returns the formatted representation of
    // the combination of all internal InclusiveRange instances
    std::string string() const;

    // returns the total number of values across all ranges
    size_t length() const;

    // returns the first value of the first range
    long start() const;

    // returns the last value of the last range
    long end() const;

    // returns the smallest value in the total range
    long min() const;

    // returns the highest value in the total range
    long max() const;

    // creates and adds another range of values
    // to the total range list.
    // Values in new range may duplicate values in
    // existing ranges.
    void append(long start, long end, int step=0) {
        m_blocks.push_back(new Range(start, end, step));
    }

    // creates and adds another range of values
    // to the total range list. Only unique values from the
    // given range are appended to the total range.
    void appendUnique(long start, long end, int step=0);

    // returns true if a given value is a valid
    // value within the total range.
    bool contains(long value) const;

    // returns the value at the given index in
    // the total range. If the index is invalid/exceeds the valid
    // range, an error will be returned.
    long value(size_t idx, Status* ok=NULL) const;

    // returns the 0-based index of the first occurrence
    // of the given value, within the range.
    // If the value does not exist in the range, a
    // value of -1 will be returned.
    ssize_t index(long value) const;

    // returns an iterator that will loop through
    // every value in the range.
    // Call to RangesIterator.next() populates first
    // value and returns whether there are more values available.
    RangesIterator iterValues() const { return RangesIterator(this); }

    // fills a new instance with a range containing
    // all values within the start/end that are not in the current range.
    // Original ordering is not preserved. New inverted range will be
    // in an increasing value.
    void inverted(Ranges &out) const { normalized(out, true); }

    // fills a new instance, where all values have
    // been sorted and compacted (where possible)
    void normalized(Ranges &out) const { normalized(out, false); }

private:
    friend class RangesIterator;

    // fills a new instance, where all values have
    // been sorted and compacted (where possible).
    // If invert is true, then return all values within the start/end that
    // are not in the current range.
    void normalized(Ranges &outRanges, bool invert) const;

    // rangeAt returns the underlying InclusiveRange instance
    // that was appended, at a given index
    const Range* rangeAt(size_t idx) const {
        return (idx >= m_blocks.size()) ? NULL : m_blocks[idx];
    }

    typedef std::vector<Range*> Blocks;
    Blocks m_blocks;

};

} // fileseq

#endif // FILESEQ_RANGES_P_H_
