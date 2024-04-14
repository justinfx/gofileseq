#include "error.h"


namespace fileseq {

namespace internal {


void handleErrorStatus(const std::string &msg, Status* stat) {
    if (stat != nullptr) {
        stat->setError(msg);
    } else {
        std::cerr << "fileseq: Unchecked error: " << msg << std::endl;
    }
}


void setError(const std::string &msg, Status* stat) {
    if (stat != nullptr) {
        stat->setError(msg);
    }
}

} // internal

} // fileseq

