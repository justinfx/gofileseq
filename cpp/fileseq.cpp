#include "fileseq.h"
#include "fileseq_p.h"

#include <string>

namespace fileseq {

FileSequence findSequenceOnDisk(const std::string &pattern, std::string* err) {
    internal::FindSequenceOnDisk_return ret;
    ret = internal::FindSequenceOnDisk(const_cast<char*>(pattern.c_str()));

    if (ret.r1 != NULL) {
        if (err != NULL) {
            err->assign(ret.r1);
        };
        free(ret.r1);
        return FileSequence(0);
    }

    return FileSequence(ret.r0);
}


std::string findSequencesOnDisk(FileSequences &seqs,
                                const std::string &path,
                                bool hiddenFiles,
                                bool singleFiles) {

    internal::FileOption opts;
    opts.hiddenFiles = hiddenFiles;
    opts.singleFiles = singleFiles;

    internal::FindSequencesOnDisk_return ret;
    ret = internal::FindSequencesOnDisk(const_cast<char*>(path.c_str()), opts);

    std::string err;

    if (ret.r2 != NULL) {
        err.assign(ret.r2);
        free(ret.r2);

        if (ret.r0 != NULL) {
            free(ret.r0);
        }
        return err;
    }

    seqs.clear();
    seqs.resize(ret.r1, FileSequence(0));

    uint64_t* start = ret.r0;
    for (uint64_t i = 0; i < ret.r1; ++i) {
        seqs[i] = FileSequence(*start);
        ++start;
//        std::cout << "cpp::DEBUG:: " << seqs.at(i).string() << std::endl;
    }

    free(ret.r0);

    return err;
}


/*
  DEBUG
*/

std::string allocStats() {
    internal::StringProxy stats = internal::FS_AllocStats();
    return stats;
}

} // fileseq
