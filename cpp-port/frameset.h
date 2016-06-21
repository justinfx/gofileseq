#ifndef FILESEQ_FRAMESET_H_
#define FILESEQ_FRAMESET_H_

#include <inttypes.h>
#include <ostream>
#include <string>
#include <vector>

namespace fileseq {

// Fordward decl
class Status;
class FileSequence;
class FrameSetData;

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
    explicit FrameSet(const std::string &frange, Status* ok=NULL);

    /*! Default Constructor - Invalid FrameSet */
    FrameSet();

    virtual ~FrameSet();

    // FrameSet(const FrameSet& rhs);
    // FrameSet& operator=(const FrameSet& rhs);

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
    long index(long frame) const;

    /*! Frame returns the frame number value for a given index into
    the frame set. If the index is outside the bounds of the frame
    set range, then an error is returned as a Status
    */
    long frame(long index, Status* ok=NULL) const;

    /*!
    Frames returns a slice of the frame numbers that were parsed from the
    original frame range string.
    Warning: This allocates a slice containing number of elements equal
    to the Len() of the range. TODO: Support frame iteration.
    */
    void frames(std::vector<long> &frames) const;

    //! HasFrame returns true if the FrameSet contains the given frame value.
    bool hasFrame(long frame) const;

    //! The first frame of the range
    long start() const;

    //! The last frame in the range
    long end() const;

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
    FrameSetData* m_frameData;

private:
    friend class FileSequence;
};

} // FILESEQ_FRAMESET_H_

#endif
