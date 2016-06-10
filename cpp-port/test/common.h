#ifndef FILESEQ_CPP_GTEST_COMMON_H
#define FILESEQ_CPP_GTEST_COMMON_H

#include "../fileseq.h"

#define TODO() FAIL() << "Not Implemented"

// Create a std::vector<long> from an long[] array
#define LONG_ARR_TO_VEC(X) (std::vector<long>(X, X+(sizeof(X)/sizeof(X[0]))))

// Create a fileseq::Frames from an fileseq::Frame[] array
#define FRAME_ARR_TO_VEC(X) (fileseq::Frames(X, X+(sizeof(X)/sizeof(X[0]))))

#endif
