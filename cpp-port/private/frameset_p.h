#ifndef FILESEQ_FRAMESET_P_H_
#define FILESEQ_FRAMESET_P_H_

#include "../ranges/ranges.h"

#include <string>


namespace fileseq {


class FrameSetData {

public:
    FrameSetData() : frameRange(), ranges() {}

    std::string frameRange;
    Ranges ranges;
};


} // fileseq

#endif // FILESEQ_FRAMESET_P_H_
