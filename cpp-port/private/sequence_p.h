#ifndef FILESEQ_SEQUENCE_P_H_
#define FILESEQ_SEQUENCE_P_H_

#include "../pad.h"

#include <string>

namespace fileseq {


namespace internal {


// Data structure holding information about
// the parsed data from a sequence string
struct SeqPatternMatch {
    std::string base;
    std::string range;
    std::string padChars;
    std::string ext;
};

// Try to match a valid sequence pattern, containing a frame range, and
// fill the match instance with the parsed results.
// Returns true if the match was successful.
bool getSplitPatternMatch(SeqPatternMatch &match, const std::string &path);

// Try to match a valid single file pattern, containing no frame range, and
// fill the match instance with the parsed results.
// Returns true if the match was successful.
bool getSingleFrameMatch(SeqPatternMatch &match, const std::string &path);

// Private data container for FileSequence
class FileSequenceData {

public:
    FileSequenceData()
        : str()
        , base()
        , dir()
        , ext()
        , pad()
        , zfill(0)
        , padStyle(PadStyleDefault) {

    }

    std::string str;
    std::string base;
    std::string dir;
    std::string ext;
    std::string pad;
    int zfill;
    PadStyle padStyle;

};


} // internal
} // fileseq

#endif // FILESEQ_SEQUENCE_P_H_
