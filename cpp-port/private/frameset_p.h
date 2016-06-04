#ifndef FILESEQ_FRAMESET_P_H_
#define FILESEQ_FRAMESET_P_H_

#include <string>

namespace fileseq {

class FrameSetData {

public:
    FrameSetData()
        : frange()
        // , rangePtr(NULL)
        {

    }

    std::string frange;
    // InclusiveRanges* rangePtr
};

}

#endif // FILESEQ_FRAMESET_P_H_
