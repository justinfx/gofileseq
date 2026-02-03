#include "sequence.h"
#include "private/sequence_p.h"
#include "private/frameset_p.h"
#include "private/strings.h"
#include "private/parser/parse.h"
#include "error.h"
#include "pad.h"

#include <algorithm>
#include <cerrno>
#include <cstdlib>
#include <sstream>

namespace fileseq {


FileSequence::FileSequence() : m_seqData(nullptr) , m_frameSet() {

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
    using namespace internal;

    // Determine which style of padding to use
    const PaddingMapper &padder = getPadMapperForStyle(padStyle);

    // Parse using ANTLR grammar-based parser (single call)
    ParseResult result;
    if (!parseFileSequence(result, path)) {
        internal::setError("Failed to parse sequence: " + path, ok);
        return false;
    }

    std::string dir, base, pad, ext;
    FrameSet frameSet;

    // Use parser results directly - already split and processed
    dir = result.directory;
    base = result.basename;
    ext = result.extension;

    if (result.isSequence) {
        // Full sequence with padding
        pad = result.padding;
        if (!result.frameRange.empty()) {
            frameSet = FrameSet(result.frameRange);
        }
    } else if (result.isSingleFile) {
        // Single frame file - auto-generate padding from frame number width
        if (!result.frameRange.empty()) {
            frameSet = FrameSet(result.frameRange);
            if (frameSet.isValid()) {
                std::string trimmed = result.frameRange;
                fileseq::strings::trim(trimmed);
                pad = padder.getPaddingChars(trimmed.size()); // NOLINT(*-narrowing-conversions)
            }
        }
    }
    // For plain files (result.isPlainFile), frameSet and pad remain empty

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

FileSequence::~FileSequence() {} // NOLINT(*-use-equals-default)

FileSequence::FileSequence(const FileSequence& rhs) {
    if (rhs.m_seqData) {
        m_seqData = std::unique_ptr<internal::FileSequenceData>(
                new internal::FileSequenceData(*rhs.m_seqData));
    } else {
        m_seqData = std::unique_ptr<internal::FileSequenceData>(
                new internal::FileSequenceData());
    }
    m_frameSet = rhs.m_frameSet;
}

FileSequence &FileSequence::operator=(FileSequence rhs) {
    // Swap with copied arg
    swap(*this, rhs);
    return *this;
}

void swap(FileSequence &first, FileSequence &second) {
    using std::swap;

    swap(first.m_seqData, second.m_seqData);
    swap(first.m_frameSet, second.m_frameSet);
}

bool FileSequence::isValid() const {
    // Sanity check
    if (m_seqData == nullptr) {
        return false;
    }

    if (m_seqData->str.empty()) {
        return false;
    }

    // We at least need a basename or ext to consider
    // ourselves valid
    return !(m_seqData->base.empty() && m_seqData->ext.empty()) || m_frameSet.isValid();

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
       << (m_frameSet.isValid() ? padding() : "")
       << ext();

    return ss.str();
}

std::string FileSequence::dirname() const {
    if (!isValid()) return "";
    return m_seqData->dir;
}

void FileSequence::setDirname(const std::string& dirname) const {
    m_seqData->dir.assign(dirname);

    if (!dirname.empty() && dirname[dirname.size()-1] != fileseq::strings::kPathSep) {
        // A non-empty string does not yet end with the path sep.
        // We will add it.
        m_seqData->dir.append(1, fileseq::strings::kPathSep);
    }
}

std::string FileSequence::basename() const {
    if (!isValid()) return "";
    return m_seqData->base;
}

void FileSequence::setBasename(const std::string& basename) const {
    m_seqData->base.assign(basename);
}

std::string FileSequence::ext() const {
    if (!isValid()) return "";
    return m_seqData->ext;
}

void FileSequence::setExt(const std::string& ext) const {
    m_seqData->ext.assign(ext);

    if (!ext.empty() && ext[0] != '.') {
        // A non-empty string does not yet start with a '.'
        // We will prepend it.
        m_seqData->ext.insert(0, 1, '.');
    }
}

std::string FileSequence::padding() const {
    if (!isValid()) return "";
    return m_seqData->pad;
}

void FileSequence::setPadding(const std::string &padChars) const {
    m_seqData->pad = padChars;

    using namespace fileseq::internal;

    // Update the zfill using the pad mapping
    const PaddingMapper &mapper = getPadMapperForStyle(m_seqData->padStyle);
    m_seqData->zfill = mapper.getPaddingCharsSize(padChars); // NOLINT(*-narrowing-conversions)
}

PadStyle FileSequence::paddingStyle() const {
    if (!isValid()) return PadStyleDefault;
    return m_seqData->padStyle;
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
    if (!isValid()) return 0;
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
        if (ok != nullptr) {
            *ok = localOk;
        }
        return;
    }

    if (ok != nullptr) ok->clearError();
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

    return frame(f);
}

std::string FileSequence::invertedFrameRange(bool padded) const {
    int pad = padded ? zfill() : 0;
    return m_frameSet.invertedFrameRange(pad);
}


// Template strings for the format() operation
static const std::string s_tpl_dir = "{{dir}}";
static const std::string s_tpl_base = "{{base}}";
static const std::string s_tpl_ext = "{{ext}}";
static const std::string s_tpl_startf = "{{startf}}";
static const std::string s_tpl_endf = "{{endf}}";
static const std::string s_tpl_len = "{{len}}";
static const std::string s_tpl_pad = "{{pad}}";
static const std::string s_tpl_zfill = "{{zfill}}";
static const std::string s_tpl_frange = "{{frange}}";
static const std::string s_tpl_inverted = "{{inverted}}";


std::string FileSequence::format(const std::string &fmt, Status* ok) const {
    if (ok != nullptr) ok->clearError();

    if (!isValid() || fmt.empty()) {
        return string();
    }

    // Start with a copy
    std::string out(fmt);

    using namespace fileseq::strings;

    // Do the cheap ones first, since we already have the values
    replace_all(out, s_tpl_dir, m_seqData->dir);
    replace_all(out, s_tpl_base, m_seqData->base);
    replace_all(out, s_tpl_ext, m_seqData->ext);
    replace_all(out, s_tpl_pad, m_seqData->pad);

    // Do the replacements that are more expensive, so we need
    // to check first if the pattern exists
    std::ostringstream ss;

    if (contains(out, s_tpl_frange)) {
        replace_all(out, s_tpl_frange, frameRange());
    }

    if (contains(out, s_tpl_inverted)) {
        replace_all(out, s_tpl_inverted, invertedFrameRange());
    }

    if (contains(out, s_tpl_zfill)) {
        ss.str("");
        ss << m_seqData->zfill;
        replace_all(out, s_tpl_zfill, ss.str());
    }

    if (contains(out, s_tpl_startf)) {
        ss.str("");
        ss << m_frameSet.start();
        replace_all(out, s_tpl_startf, ss.str());
    }

    if (contains(out, s_tpl_endf)) {
        ss.str("");
        ss << m_frameSet.end();
        replace_all(out, s_tpl_endf, ss.str());
    }

    if (contains(out, s_tpl_len)) {
        ss.str("");
        ss << m_frameSet.length();
        replace_all(out, s_tpl_len, ss.str());
    }

    return out;
}


} // fileseq
