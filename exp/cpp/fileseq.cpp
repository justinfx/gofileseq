#include "fileseq.h"
#include "private/fileseq_p.h"

#include <string>

namespace fileseq {

std::string framesToFrameRange(const std::vector<int> &frames,
                               bool sorted, int zfill) {

    internal::StringProxy str = internal::FramesToFrameRange(
        const_cast<int*>(&frames[0]), frames.size(), sorted, zfill);

    return str;
}

bool isFrameRange(const std::string &frange) {
    return internal::IsFrameRange(const_cast<char*>(frange.c_str()));
}

std::string getPaddingChars(int pad) {
    internal::StringProxy str = internal::PaddingChars(pad);
    return str;
}

std::string padFrameRange(const std::string &frange, int pad) {
    internal::StringProxy str = internal::PadFrameRange(
        const_cast<char*>(frange.c_str()), pad);

    return str;
}

FileSequence findSequenceOnDisk(const std::string &pattern, Status* ok) {
    return findSequenceOnDisk(pattern, PadStyleDefault, ok);
}

FileSequence findSequenceOnDisk(const std::string &pattern, PadStyle style, Status* ok) {

    internal::FindSequenceOnDiskPad_return ret;
    ret = internal::FindSequenceOnDiskPad(const_cast<char*>(pattern.c_str()), int(style));

    if (ret.r1 != NULL) {
        if (ok != NULL) {
            ok->setError(ret.r1);
        } else {
            internal::printErrorIgnored(ret.r1);
        }
        free(ret.r1);
        return FileSequence(0);
    }

    if (ok != NULL) {
        ok->clearError();
    }

    return FileSequence(ret.r0);
}

Status findSequencesOnDisk(FileSequences &seqs,
                           const std::string &path,
                           bool hiddenFiles,
                           bool singleFiles,
                           PadStyle style) {

    internal::FileOption opts;
    opts.hiddenFiles = hiddenFiles;
    opts.singleFiles = singleFiles;
    opts.padStyle = int(style);

    internal::FindSequencesOnDisk_return ret;
    ret = internal::FindSequencesOnDisk(const_cast<char*>(path.c_str()), opts);

    Status stat;

    if (ret.r2 != NULL) {
        stat.setError(ret.r2);
        free(ret.r2);

        if (ret.r0 != NULL) {
            free(ret.r0);
        }
        return stat;
    }

    seqs.clear();
    seqs.resize(ret.r1, FileSequence(0));

    internal::GoUint64* start = ret.r0;
    for (internal::GoUint64 i = 0; i < ret.r1; ++i) {
        seqs[i] = FileSequence(*start);
        ++start;
    }

    free(ret.r0);

    return stat;
}


/*
  DEBUG
*/

std::string allocStats() {
    internal::StringProxy stats = internal::FS_AllocStats();
    return stats;
}

} // fileseq
