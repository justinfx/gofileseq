#include "frameset.h"
#include "fileseq_p.h"

#include <iostream>

namespace fileseq {

FrameSet::FrameSet(const std::string &frange)
    : m_valid(false)
    , m_id(0) 
    , m_lastError() {

    internal::FrameSet_New_return fs = internal::FrameSet_New(
                const_cast<char*>(frange.c_str()));

    if (fs.r1 != NULL) {
        std::string err(fs.r1);
        free(fs.r1);
        m_lastError = "Failed to create Frames(" + frange + "): " + err; 
        return;
    }

    m_id = fs.r0;
    m_valid = true;
}

FrameSet::FrameSet(uint64_t id)
    : m_valid(id != 0)
    , m_id(id) 
    , m_lastError()
{
}

FrameSet::~FrameSet() {
    if (m_valid) {
        m_valid = false;
        internal::FrameSet_Decref(m_id);
    }
}

FrameSet::FrameSet(const FrameSet& rhs)
    : m_valid(false)
    , m_id(0)
    , m_lastError()
{
    m_id = internal::FrameSet_Copy(rhs.m_id);
    if (m_id != 0) {
        m_valid = true;
    }
}

FrameSet& FrameSet::operator=(const FrameSet& rhs) {
    // Decref the previous value
    if (m_id != 0) {
        internal::FrameSet_Decref(m_id);
        m_valid = false;
    }

    // Take on the reference from the other
    m_id = rhs.m_id;
    if (m_id != 0) {
        m_valid = true;
        internal::FrameSet_Incref(rhs.m_id);
    }

    return *this;
}

std::string FrameSet::string() const {
    internal::StringProxy str = internal::FrameSet_String(m_id);
    return str;
}

size_t FrameSet::length() const {
    return internal::FrameSet_Len(m_id);
};

int FrameSet::index(int frame) const {
    return internal::FrameSet_Index(m_id, frame);
};

int FrameSet::frame(int index, bool* ok) const {
    m_lastError.clear();

    internal::FrameSet_Frame_return fs = internal::FrameSet_Frame(m_id, index);
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
    size_t n = internal::FrameSet_Frames(m_id, &frames[0]);
    frames.resize(n);
};

bool FrameSet::hasFrame(int frame) const {
    return internal::FrameSet_HasFrame(m_id, frame);
};

int FrameSet::start() const {
    return internal::FrameSet_Start(m_id);
};

int FrameSet::end() const {
    return internal::FrameSet_End(m_id);
};

std::string FrameSet::frameRange(int pad) const {
    internal::StringProxy str;
    if (pad >= 2) {
        str = internal::FrameSet_FrameRangePadded(m_id, pad);
    } else {
        str = internal::FrameSet_FrameRange(m_id);
    }
    return str;
};

FrameSet FrameSet::inverted() const {
    internal::GoUint64 invertedId = internal::FrameSet_Invert(m_id);
    return FrameSet(invertedId);
};

std::string FrameSet::invertedFrameRange(int pad) const {
    internal::StringProxy ret = internal::FrameSet_InvertedFrameRange(m_id, pad);
    return ret;
};

FrameSet FrameSet::normalized() const {
    internal::GoUint64 normalizedId = internal::FrameSet_Normalize(m_id);
    return FrameSet(normalizedId);
};

} // fileseq
