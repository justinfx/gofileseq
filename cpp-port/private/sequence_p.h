#ifndef FILESEQ_SEQUENCE_P_H_
#define FILESEQ_SEQUENCE_P_H_

#include <string>

namespace fileseq {

namespace internal {


class FileSequenceData {

public:
    FileSequenceData()
        : str()
        , base()
        , dir()
        , ext()
        , pad()
        , zfill(0) {

    }

    std::string str;
    std::string base;
    std::string dir;
    std::string ext;
    std::string pad;
    int zfill;
};


} // internal
} // fileseq

#endif // FILESEQ_SEQUENCE_P_H_
