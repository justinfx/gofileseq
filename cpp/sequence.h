#ifndef _GOFILESEQ_CPP_SEQUENCE_H_
#define _GOFILESEQ_CPP_SEQUENCE_H_

#include "frameset.h"
#include "types.h"

#include <ostream>

namespace fileseq {

/*
A FileSequence represents a path to a sequence of files,
and expresses a valid frames range (one that can be parsed by
FrameSet. It also includes one or more padding characters to
dictate how much padding the actual file numbers have.

Valid padding characters:

    @ - 1 pad width (@@@@ is equal to 4 padding)
    # - 4 pad width (## is equal to 8 padding)

Example paths and padding:

    /path/to/single_image.0100.jpg
    /path/to/image_foo.1-10x2#.jpg   (i.e. 0001)
    /path/to/image_foo.1-10x2@.jpg   (i.e. 1)
    /path/to/image_foo.1-10x2@@@.jpg (i.e. 001)

*/
class FileSequence {

public:
    // Create a new FrameSet from a given frame range string.
    // If the range could not be parsed, Status will evaluate to false
    // and be filled with the error message
    explicit FileSequence(const std::string &frange, Status* ok=NULL);

    ~FileSequence();

    FileSequence(const FileSequence& rhs);

    FileSequence& operator=(const FileSequence& rhs);

    // Return whether the range was properly parsed
    bool isValid() const { return m_valid; }

    // Return the string representation of the file sequence
    std::string string() const;

    operator std::string() const { return string(); }

    friend std::ostream& operator<< (std::ostream& stream, const FileSequence& fs) {
        stream << fs.string();
        return stream;
    }

    // The number of files/frames in the sequence range
    size_t length() const;

    // Return the first frame number in the range
    int start() const;

    // Return the last frame number in the range
    int end() const;

    // zfill returns the number of "0" fill characters used to
    // pad the frame numbers to match the actual file paths
    int zfill() const;

    /*
    Format returns the file sequence as a formatted string according
    to the given template.

    Utilizes a template syntax with a set of available replacement symbols.
    These symbols are expressed in the form:  {{SYMBOL}}

    dir      - the directory name.
    base     - the basename of the sequence (leading up to the frame range).
    ext      - the file extension of the sequence including leading period.
    startf   - the start frame.
    endf     - the end frame.
    len      - the length of the frame range.
    pad      - the detected padding characters (i.e. # , @@@ , ...).
    frange   - the frame range.
    inverted - the inverted frame range. (returns empty string if none)
    zfill    - the int width of the frame padding

    Example:

        {{dir}}{{base}}{{frange}}{{pad}}{{ext}}
    */
    std::string format(const std::string &fmt, Status* ok=NULL) const;

    // Dirname returns the parsed directory component of the sequence string
    std::string dirname() const;

    // Set the dir name component of the sequence string
    void setDirname(const std::string &dirname) const;

    // Basename returns the parsed basename component of the sequence string.
    // This is the file part, just before the frame range component.
    std::string basename() const;

    // Set the basename component of the sequence string.
    // This is the file part, just before the frame range component.
    void setBasename(const std::string &basename) const;

    // Ext returns the file extension component from the sequence, including
    // the leading period character.
    std::string ext() const;

    // Set a new extension component for the sequence, which should include
    // the leading period character.
    void setExt(const std::string &ext) const;

    // Returns the parsed padding characters component of the sequence string.
    // May return an empty string if the file sequence was really a single
    // file. i.e. # or @ or @@@
    std::string padding() const;

    // Set new padding characters for the sequence, i.e. # or @ or @@@
    void setPadding(const std::string &padChars) const;


    // FrameRange returns the string frame range component, parsed from the
    // sequence. If no frame range was parsed, then this method will return
    // an empty string.
    // If padded is true, then pad the frame range using the zfill
    std::string frameRange(bool padded=false) const;

    // Set a new FrameSet, by way of providing a string frame range.
    void setFrameRange(const std::string &frange, Status* ok=NULL);

    // InvertedFrameRange returns a new frame range that represents all
    // frames *not* within the current frame range. That is, it will
    // create a range that "fills in" the current one.
    // If padded is true, then pad out frames using zfill
    std::string invertedFrameRange(bool padded=false) const;

    // Index returns the path to the file at the given index in the sequence.
    // If a frame range was not parsed from the sequence, this will always
    // returns the original path. If the index is not valid, this will
    // return an empty string.
    std::string index(int idx) const;

    /*
    Frame returns a path to the given frame in the sequence.
    Integer or string digits are treated as a frame number and padding is
    applied, all other values are passed though.

    Example:

    seq.Frame(1)
    >> /foo/bar.0001.exr

    seq.Frame("#")
    >> /foo/bar.#.exr
    */
    std::string frame(int frame) const;

    /*
    Frame returns a path to the given frame in the sequence.
    Integer or string digits are treated as a frame number and padding is
    applied, all other values are passed though.

    Example:

    seq.Frame(1)
    >> /foo/bar.0001.exr

    seq.Frame("#")
    >> /foo/bar.#.exr

    */
    std::string frame(const std::string &fillPattern) const;

    // FrameSet returns a FrameSet instance that was created when the
    // sequence was parsed. If no frame range was parsed from the sequence,
    // then this method will return an invalid FrameSet.
    FrameSet frameSet();

    // Set a new FrameSet for this sequence
    void setFrameSet(FrameSet &frameSet);

private:
    explicit FileSequence(uint64_t id);

    bool m_valid;
    uint64_t m_id;
    uint64_t m_fsetId;

    friend FileSequence findSequenceOnDisk(const std::string &pattern, Status* ok);

    friend Status findSequencesOnDisk(FileSequences &seqs,
                                      const std::string &path,
                                      bool hiddenFiles,
                                      bool singleFiles);
};

} // fileseq

#endif // _GOFILESEQ_CPP_SEQUENCE_H_
