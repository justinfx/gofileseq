#include "error.h"


namespace fileseq {

namespace internal {


void handleErrorStatus(const std::string &msg, Status* stat) {
    if (stat != NULL) {
        stat->setError(msg);
    } else {
        std::cerr << "fileseq: Unchecked error: " << msg << std::endl;
    }
}


} // internal

} // fileseq

