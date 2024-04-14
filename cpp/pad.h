#ifndef FILESEQ_PAD_H_
#define FILESEQ_PAD_H_

#include <map>
#include <string>
#include <vector>


namespace fileseq {


// Constants defining the style of padding to use
// when converting between padding characters ('#', '##', '@@@')
// and their equivalent numeric padding width
enum PadStyle {
    PadStyleHash1 = 0,
    PadStyleHash4 = 1,
    PadStyleDefault = PadStyleHash4
};


/*!
padFrameRange takes a frame range string and returns a new range
with each number padded out with zeros to a given width
*/
std::string padFrameRange(const std::string &frange, size_t pad);


/*!
getPaddingChars returns the proper padding characters,
given an amount of padding. Uses PadStyleDefault.
*/
std::string getPaddingChars(size_t width);


// Internal
//

namespace internal {

// Fwd Decl
class PaddingMapper;


const PaddingMapper& getPadMapperForStyle(PadStyle style);


// PaddingMapper defines behavior for converting between
// padding characters and their pad width
class PaddingMapper {

public:
    virtual ~PaddingMapper() = default;

    // Return all supported padding characters
    std::string getAllChars() const;

    // Return a width for a string of pad characters
    size_t getPaddingCharsSize(const std::string &chars) const;

    // Return the padding string sequence representing a width
    virtual std::string getPaddingChars(long width) const = 0;

protected:
    typedef std::map<char, size_t> CharSizeMap;

    // Mapping pad characters to their width
    // CharSizeMap* m_charToSize;
    CharSizeMap m_charToSize;

    // The default padding character to use
    char m_defaultChar;
};


class SingleHashPad : public PaddingMapper {

public:
    SingleHashPad();

    virtual std::string getPaddingChars(long width) const;
};


class MultiHashPad : public PaddingMapper {

public:
    MultiHashPad();

    virtual std::string getPaddingChars(long width) const;
};


} // internal


} // fileseq

#endif // FILESEQ_PAD_H_

