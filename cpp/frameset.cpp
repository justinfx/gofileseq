#include "frameset.h"

#include <iostream>

namespace fileseq {

FrameSet::FrameSet(const std::string &frange)
    : m_valid(false)
    , m_id(0) 
    , m_lastError() {

    FrameSet_New_return fs = FrameSet_New(const_cast<char*>(frange.c_str()));
    if (fs.r1 != NULL) {
        std::string err(fs.r1);
        free(fs.r1);
        m_lastError = "Failed to create Frames(" + frange + "): " + err; 
        return;
    }

    m_id = fs.r0;
    m_valid = true;
}

FrameSet::FrameSet(GoInt64 id)
    : m_valid(true)
    , m_id(id) 
    , m_lastError()
{

}

FrameSet::~FrameSet() {
    if (m_valid) {
        m_valid = false;
        FrameSet_Delete(m_id);      
    }
}

FrameSet::FrameSet(const FrameSet& rhs)
    : m_valid(false)
    , m_id(0)
    , m_lastError()
{
    m_id = FrameSet_Copy(rhs.m_id);
    if (m_id != 0) {
        m_valid = true;
    }
}

FrameSet& FrameSet::operator=(const FrameSet& rhs) {
    m_id = FrameSet_Copy(rhs.m_id);
    if (m_id != 0) {
        m_valid = true;
    }
    return *this;
}

std::string FrameSet::string() const {
    char* ptr = FrameSet_String(m_id);
    std::string str(ptr);
    free(ptr);
    return str;
}

size_t FrameSet::length() const {
    return FrameSet_Len(m_id);
};

int FrameSet::index(int frame) const {
    return FrameSet_Index(m_id, frame);
};

int FrameSet::frame(int index, bool* ok) const {
    m_lastError.clear();

    FrameSet_Frame_return fs = FrameSet_Frame(m_id, index);
    if (fs.r1 != NULL) {
        if (ok != NULL) {
            *ok = false;
        }
        m_lastError.assign(fs.r1);
        free(fs.r1); // we own the error
        return 0;
    }
    if (ok != NULL) {
        *ok = true;
    }
    return fs.r0;
};

void FrameSet::frames(std::vector<int> &frames) const {
    frames.clear();

    size_t len = length();
    if (len <= 0) {
        return;
    }

    frames.resize(len);
    size_t n = FrameSet_Frames(m_id, &frames[0]);
    frames.resize(n);
};

bool FrameSet::hasFrame(int frame) const {
    return FrameSet_HasFrame(m_id, frame);
};

int FrameSet::start() const {
    return FrameSet_Start(m_id);
};

int FrameSet::end() const {
    return FrameSet_End(m_id);
};

std::string FrameSet::frameRange(int pad) const {
    char* str;
    if (pad >= 2) {
        str = FrameSet_FrameRangePadded(m_id, pad);
    } else {
        str = FrameSet_FrameRange(m_id);
    }
    std::string ret(str);
    free(str);
    return ret;
};

FrameSet FrameSet::inverted() const {
    GoUint64 invertedId = FrameSet_Invert(m_id);
    return FrameSet(invertedId);
};

std::string FrameSet::invertedFrameRange(int pad) const {
    char* str = FrameSet_InvertedFrameRange(m_id, pad);
    std::string ret(str);
    free(str);
    return ret;
};

FrameSet FrameSet::normalized() const {
    GoUint64 normalizedId = FrameSet_Normalize(m_id);
    return FrameSet(normalizedId);
};

} // fileseq
