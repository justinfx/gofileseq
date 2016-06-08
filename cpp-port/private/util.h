#ifndef FILESEQ_UTIL_P_H_
#define FILESEQ_UTIL_P_H_

#include "../types.h"

#include <iostream>
#include <string>

namespace fileseq {

namespace internal {

void handleErrorStatus(const std::string &msg, fileseq::Status *stat=NULL) {
    if (stat != NULL) {
        stat->setError(msg);
    } else {
        std::cerr << "fileseq: Unchecked error: " << msg << std::endl;
    }
}

} // internal

} // fileseq

#endif // FILESEQ_UTIL_P_H_
