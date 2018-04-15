#include "fileseq_p.h"

#include <iostream>

namespace fileseq {

namespace internal {

void printErrorIgnored(const std::string &err) {
    std::cerr << "Unhandled Error: " << err << std::endl;
}

} // internal

} //fileseq
