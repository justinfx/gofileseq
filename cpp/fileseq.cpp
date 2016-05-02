#include "fileseq.h"
#include "private/fileseq_p.h"

#include <string>

namespace fileseq {

FileSequence findSequenceOnDisk(const std::string &pattern, Status* ok) {

    internal::FindSequenceOnDisk_return ret;
    ret = internal::FindSequenceOnDisk(const_cast<char*>(pattern.c_str()));

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
                           bool singleFiles) {

    internal::FileOption opts;
    opts.hiddenFiles = hiddenFiles;
    opts.singleFiles = singleFiles;

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
