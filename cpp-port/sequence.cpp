#include "sequence.h"
#include "private/sequence_p.h"

namespace fileseq {

FileSequence::FileSequence() : m_seqData(NULL) , m_frameSet() {

}

FileSequence::FileSequence(const std::string &path, Status* ok)
    : m_seqData(NULL)
    , m_frameSet() {

    init(path, PadStyleDefault, ok);
}

FileSequence::FileSequence(const std::string &path, PadStyle padStyle, Status* ok)
    : m_seqData(NULL)
    , m_frameSet() {

    init(path, padStyle, ok);
}

void FileSequence::init(const std::string &path, PadStyle padStyle, Status* ok) {

}

FileSequence::~FileSequence() {
    if (m_seqData) {
        delete m_seqData;
        m_seqData = NULL;
    }
}

bool FileSequence::isValid() const {
    return m_seqData != NULL;
}

size_t FileSequence::length() const {
    return m_frameSet.length();
}

std::string FileSequence::string() const {
    return m_seqData ? m_seqData->str : "";
}

std::string FileSequence::format(const std::string& fmt, Status* ok) const {
    return "";
}

std::string FileSequence::dirname() const {
    return m_seqData ? m_seqData->dir : "";
}

void FileSequence::setDirname(const std::string& dirname) const {

}

std::string FileSequence::basename() const {
    return m_seqData ? m_seqData->base : "";
}

void FileSequence::setBasename(const std::string& basename) const {

}

std::string FileSequence::ext() const {
    return m_seqData ? m_seqData->ext : "";
}

void FileSequence::setExt(const std::string& ext) const {

}

std::string FileSequence::padding() const {
    return m_seqData ? m_seqData->pad : "";
}

void FileSequence::setPadding(const std::string &padChars) const {

}

PadStyle FileSequence::paddingStyle() const {
    return PadStyleDefault;
}

void FileSequence::setPaddingStyle(PadStyle style) const {

}

int FileSequence::start() const {
    return m_frameSet.start();
}

int FileSequence::end() const {
    return m_frameSet.end();
}

int FileSequence::zfill() const {
    return 0;
}

std::string FileSequence::frameRange(bool padded) const {
    return "";
}

void FileSequence::setFrameRange(const std::string &frange, Status* ok) {

}

std::string FileSequence::invertedFrameRange(bool padded) const {
    return "";
}

std::string FileSequence::index(int idx) const {
    return "";
}

std::string FileSequence::frame(int frame) const {
    return "";
}

std::string FileSequence::frame(const std::string &fillPattern) const {
    return "";
}

FrameSet FileSequence::frameSet() {
    return m_frameSet;
}

void FileSequence::setFrameSet(const FrameSet &frameSet) {
    m_frameSet = frameSet;
}

} // fileseq
