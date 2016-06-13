#include "sequence.h"
#include "private/sequence_p.h"
#include "private/frameset_p.h"
#include "private/strings.h"
#include "error.h"
#include "pad.h"

#include <algorithm>
#include <cerrno>
#include <climits>
#include <cstdlib>
#include <sstream>

namespace fileseq {


FileSequence::FileSequence() : m_seqData(NULL) , m_frameSet() {

}

FileSequence::FileSequence(const std::string &path, Status* ok)
    : m_seqData(new internal::FileSequenceData)
    , m_frameSet() {

    init(path, PadStyleDefault, ok);
}

FileSequence::FileSequence(const std::string &path, PadStyle padStyle, Status* ok)
    : m_seqData(new internal::FileSequenceData)
    , m_frameSet() {

    init(path, padStyle, ok);
}

bool FileSequence::init(const std::string &path, PadStyle padStyle, Status* ok) {
    std::string dir, base, pad, ext;
    FrameSet frameSet;

    using namespace internal;

    // Determine which style of padding to use
    const PaddingMapper &padder = getPadMapperForStyle(padStyle);

    internal::SeqPatternMatch match;
    bool didMatch = internal::getSplitPatternMatch(match, path);

    if (!didMatch) {
        // If no match at this point, we are dealing with a possible
        // path that contains no range. Maybe a single frame or no frame.

        if (path.find_first_of(padder.getAllChars()) != std::string::npos) {
            // String contains padding characters, but we simply could
            // not parse it as a sequence
            internal::setError("Failed to parse sequence: " + path, ok);
            return false;
        }

        // We assume the sequence is just a single file path, containing
        // no frame patterns
        path_split(dir, base, path);

        // See if we can find the file ext
        size_t found = base.find_last_of('.');
        if (found != std::string::npos) {

            const std::string t_base(base);
            base = t_base.substr(0, found);

            // including the leading "."
            ext = t_base.substr(found);
        }

        if (dir.empty() && base.empty() && !ext.empty()) {
            // Just assume all we have is a file extension

        } else {
            // Try to see if we can at least find a specific frame
            // number, a la  .<frame>.ext
            if ( (didMatch = internal::getSingleFrameMatch(match, path)) ) {

                frameSet = FrameSet(match.range);
                if (frameSet.isValid()) {

                    // Reparse the dir/base to not include the trailing frame
                    path_split(dir, base, match.base);

                    // Calculate the padding chars
                    trim_string(match.range);
                    pad = padder.getPaddingChars(match.range.size());
                }

                ext = match.ext;
            }
        }

    } else {
        // We are dealing with a pattern containing a frame range

        frameSet = FrameSet(match.range);
        path_split(dir, base, match.base);
        pad = match.padChars;
        ext = match.ext;
    }

    m_seqData->str.assign(path);
    m_seqData->base.assign(base);
    m_seqData->dir.assign(dir);
    m_seqData->pad.assign(pad);
    m_seqData->ext.assign(ext);
    m_seqData->padStyle = padStyle;

    std::swap( m_frameSet, frameSet );

    // Initializes the proper padding and fill
    setPadding(pad);

    return true;
}

FileSequence::~FileSequence() {
    if (m_seqData) {
        delete m_seqData;
        m_seqData = NULL;
    }
}

FileSequence::FileSequence(const FileSequence& rhs) {
    if (m_seqData == NULL) {
        delete m_seqData;
    }

    m_seqData = new internal::FileSequenceData(*(rhs.m_seqData));
    m_frameSet = rhs.m_frameSet;
}

bool FileSequence::isValid() const {
    // Sanity check
    if (m_seqData == NULL) {
        return false;
    }

    // We at least need a basename or ext to consider
    // ourselves valid
    if (m_seqData->str.empty() || (m_seqData->base.empty() && m_seqData->ext.empty())) {
        return false;
    }

    return true;
}

size_t FileSequence::length() const {
    if (!isValid()) {
        return 0;
    }
    // Sequences should return a min length of 1
    // if they are valid
    size_t len = m_frameSet.length();
    return len > 1 ? len : 1;
}

std::string FileSequence::string() const {
    // TODO: Cache this until setters invalidate the
    // string representation
    std::stringstream ss;

    ss << dirname()
       << basename()
       << frameRange()
       << padding()
       << ext();

    return ss.str();
}

std::string FileSequence::dirname() const {
    return m_seqData->dir;
}

void FileSequence::setDirname(const std::string& dirname) const {
    m_seqData->dir.assign(dirname);

    if (!dirname.empty() && dirname[dirname.size()-1] != kPathSep) {
        // A non empty string does not yet end with the path sep.
        // We will add it.
        m_seqData->dir.append(1, kPathSep);
    }
}

std::string FileSequence::basename() const {
    return m_seqData->base;
}

void FileSequence::setBasename(const std::string& basename) const {
    m_seqData->base.assign(basename);
}

std::string FileSequence::ext() const {
    return m_seqData->ext;
}

void FileSequence::setExt(const std::string& ext) const {
    m_seqData->ext.assign(ext);

    if (!ext.empty() && ext[0] != '.') {
        // A non empty string does not yet start with a '.'
        // We will prepend it.
        m_seqData->ext.insert(0, 1, '.');
    }
}

std::string FileSequence::padding() const {
    return m_seqData->pad;
}

void FileSequence::setPadding(const std::string &padChars) const {
    m_seqData->pad = padChars;

    using namespace fileseq::internal;

    // Update the zfill using the pad mapping
    const PaddingMapper &mapper = getPadMapperForStyle(m_seqData->padStyle);
    m_seqData->zfill = mapper.getPaddingCharsSize(padChars);
}

PadStyle FileSequence::paddingStyle() const {
    return m_seqData->padStyle;;
}

void FileSequence::setPaddingStyle(PadStyle style) const {
    m_seqData->padStyle = style;

    using namespace fileseq::internal;

    // Update the padding chars from the current zfill
    const PaddingMapper &mapper = getPadMapperForStyle(style);
    setPadding(mapper.getPaddingChars(zfill()));
}

Frame FileSequence::start() const {
    return m_frameSet.start();
}

Frame FileSequence::end() const {
    return m_frameSet.end();
}

int FileSequence::zfill() const {
    return m_seqData->zfill;
}

FrameSet FileSequence::frameSet() {
    return m_frameSet;
}

void FileSequence::setFrameSet(const FrameSet &frameSet) {
    m_frameSet = frameSet;
}

std::string FileSequence::frameRange(bool padded) const {
    return m_frameSet.frameRange( padded ? m_seqData->zfill : 0 );
}

void FileSequence::setFrameRange(const std::string &frange, Status* ok) {
    Status localOk;
    FrameSet fs(frange, &localOk);

    if (!localOk) {
        if (ok != NULL) {
            *ok = localOk;
        }
        return;
    }

    if (ok != NULL) ok->clearError();
    m_frameSet = fs;
}

std::string FileSequence::index(size_t idx) const {
    if (!m_frameSet.isValid()) {
        return string();
    }

    Status ok;
    Frame f = m_frameSet.frame(idx, &ok);
    if (!ok) {
        return "";
    }

    return frame(f);
}

std::string FileSequence::frame(Frame frame) const {
    std::string zframe = internal::zfill(frame, zfill());

    std::stringstream ss;
    ss << dirname() << basename() << zframe << ext();
    return ss.str();
}

std::string FileSequence::frame(const std::string &fillPattern) const {
    if (fillPattern.empty()) {
        return "";
    }

    // Try and convert the string to a long
    char* pEnd;
    long f = std::strtol(fillPattern.c_str(), &pEnd, 10);

    // Check if the number was larger than we can handle
    if (errno == ERANGE) {
        return "";
    }

    return frame(Frame(f));
}

std::string FileSequence::invertedFrameRange(bool padded) const {
    int pad = padded ? zfill() : 0;
    return m_frameSet.invertedFrameRange(pad);
}

std::string FileSequence::format(const std::string &fmt, Status* ok) const {
    if (ok != NULL) ok->clearError();
    return "";
}


} // fileseq
