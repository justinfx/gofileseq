#include "frameset.h"
#include "types.h"
#include "private/frameset_p.h"

#include <iostream>

namespace fileseq {

FrameSet::FrameSet(const std::string &frange, Status* ok) : m_frameData(NULL) {

}

FrameSet::FrameSet() : m_frameData(NULL) {

}

FrameSet::~FrameSet() {
    if ( m_frameData ) {
        delete m_frameData;
        m_frameData = NULL;
    }
}

// FrameSet::FrameSet(const FrameSet& rhs) : m_valid(false) {

// }

// FrameSet& FrameSet::operator=(const FrameSet& rhs) {

// }

bool FrameSet::isValid() const {
    return m_frameData != NULL;
}

std::string FrameSet::string() const {
    return "";
}

size_t FrameSet::length() const {
    return 0;
};

int FrameSet::index(int frame) const {
    return 0;
};

int FrameSet::frame(int index, Status* ok) const {
    return 0;
};

void FrameSet::frames(std::vector<int> &frames) const {
    frames.clear();

    size_t len = length();
    if (len <= 0) {
        return;
    }

    frames.resize(len);
    // TODO
};

bool FrameSet::hasFrame(int frame) const {
    return false;
};

int FrameSet::start() const {
    return 0;
};

int FrameSet::end() const {
    return 0;
};

std::string FrameSet::frameRange(int pad) const {
    return "";
};

FrameSet FrameSet::inverted() const {
    return FrameSet();
};

std::string FrameSet::invertedFrameRange(int pad) const {
    return "";
};

FrameSet FrameSet::normalized() const {
    return FrameSet();
};

} // fileseq
