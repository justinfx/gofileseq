#ifndef FILESEQ_FRAMESET_H_
#define FILESEQ_FRAMESET_H_

#include "ranges/ranges.h"

#include <algorithm>
#include <inttypes.h>
#include <ostream>
#include <string>
#include <vector>


namespace fileseq {

// Forward decl
class Status;
class FileSequence;


namespace internal {

// Forward decl
class FrameSetData;
struct RangePatternMatch;

} // internal


// typedefs
typedef long Frame;
typedef std::vector<Frame> Frames;


/*!
FrameSet wraps a sequence of frames in container that exposes
array-like operations and queries, after parsing the given frame
range string.
*/
class FrameSet {

public:
    /*!
    Create a new FrameSet from a given frame range string.
    If the frame range could not be parsed, the FrameSet.isValid()
    and the Status will be set.
    */
    explicit FrameSet(const std::string &frange, Status* ok=nullptr);

    /*! Default Constructor - Invalid FrameSet */
    FrameSet();

    // Destructor
    virtual ~FrameSet();

    // Copy constructor
    FrameSet(const FrameSet& rhs);

    // Assignment
    FrameSet& operator=(FrameSet rhs) {
        // Swap with copied arg
        swap(*this, rhs);
        return *this;
    }

    // Swap functionality
    friend void swap(FrameSet &first, FrameSet &second) {
        using std::swap;

        swap(first.m_frameData, second.m_frameData);
    }

    //! Return whether the FrameSet is properly parsed and valid as a range.
    bool isValid() const;

    //! The string representation of the frame range
    std::string string() const;

    operator std::string() const { return string(); }

    friend std::ostream& operator<< (std::ostream& stream, const FrameSet& fs) {
        stream << fs.string();
        return stream;
    }

    //! The number of frames in the range
    size_t length() const;

    /*!
    Index returns the index position of the frame value within the frame set.
    If the given frame does not exist, then return -1
    */
    size_t index(Frame frame) const;

    /*! Frame returns the frame number value for a given index into
    the frame set. If the index is outside the bounds of the frame
    set range, then an error is returned as a Status
    */
    Frame frame(size_t index, Status* ok=nullptr) const;

    /*!
    Frames returns a slice of the frame numbers that were parsed from the
    original frame range string.

    Warning: This allocates a slice containing number of elements equal
    to the length() of the range. It is better to use iterFrames() for
    large ranges.
    */
    void frames(Frames &frames) const;

    /*!
    Returns an iterator that can loop over all frame numbers that were
    parsed from the original frame range string. This is more efficient
    than calling frames() for larger sequences.

    Examples:
        RangesIterator it = frameSet.iterFrames();
        while(it.next()) {
            Frame f = *it;
        }

    */
    RangesIterator iterFrames() const;

    //! HasFrame returns true if the FrameSet contains the given frame value.
    bool hasFrame(Frame frame) const;

    //! The first frame of the range
    Frame start() const;

    //! The last frame in the range
    Frame end() const;

    /*!
    FrameRange returns the range string that was used to initialize the FrameSet.
    If a pad value is given, return a range string with each number padded out
    with zeros to a given width
    */
    std::string frameRange(int pad=0) const;

    /*!
    Invert returns a new FrameSet that represents all frames *not* within the
    current FrameSet. That is, it will create a range that "fills in" the current one.
    */
    FrameSet inverted() const;

    /*!
    InvertedFrameRange returns a new frame range string that represents all
    frames *not* within the current FrameSet. That is, it will create a range
    that "fills in" the current one.
    */
    std::string invertedFrameRange(int pad) const;

    //! Normalize returns a new sorted and compacted FrameSet
    FrameSet normalized() const;

private:
    // Process a rangePattern match group
    void handleMatch(const internal::RangePatternMatch* match, Status* ok);

    // Reset the FrameSet to an invalid state;
    void setInvalid();

    internal::FrameSetData* m_frameData;

private:
    friend class FileSequence;
};

} // FILESEQ_FRAMESET_H_

#endif
