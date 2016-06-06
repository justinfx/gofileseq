#include "fileseq.h"

#include <string>

namespace fileseq {

std::string framesToFrameRange(const Frames &frames,
                               bool sorted, int zfill) {
    return "";
}

bool isFrameRange(const std::string &frange) {
    return false;
}

std::string getPaddingChars(int pad) {
    return "";
}

std::string padFrameRange(const std::string &frange, int pad) {
    return "";
}

FileSequence findSequenceOnDisk(const std::string &pattern, Status* ok) {
    return findSequenceOnDisk(pattern, PadStyleDefault, ok);
}

FileSequence findSequenceOnDisk(const std::string &pattern, PadStyle style, Status* ok) {
    return FileSequence("");
}

Status findSequencesOnDisk(FileSequences &seqs,
                           const std::string &path,
                           bool hiddenFiles,
                           bool singleFiles,
                           PadStyle style) {

    Status stat;
    stat.setError("Not implemented");
    return stat;
}

} // fileseq
