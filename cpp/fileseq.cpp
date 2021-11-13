#include "fileseq.h"
#include "pad.h"
#include "private/frameset_p.h"
#include "private/sequence_p.h"
#include "private/strings.h"

#include <algorithm>
#include <cerrno>
#include <cstring>
#include <dirent.h>
#include <iterator>
#include <map>
#include <string>
#include <sstream>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <utility>

namespace fileseq {

std::string framesToFrameRange(const Frames &frames,
                               bool sorted, int zfill) {

    Frames::size_type count = frames.size();
    if (count == 0) {
        return "";
    }

    if (count == 1) {
        return internal::zfill(frames[0], zfill);
    }

    Frames::const_iterator framesIt = frames.begin();
    Frames::const_iterator framesEnd = frames.end();

    Frames sortedFrames;

    if (sorted) {
        // Copy
        sortedFrames = frames;

        // Sort
        std::sort(sortedFrames.begin(), sortedFrames.end());

        framesIt = sortedFrames.begin();
        framesEnd = sortedFrames.end();
    }

    long step = 0;
    bool hasWritten = false;
    Frames::size_type i = 0;

    std::string start, end;
    std::stringstream buf;

    while (framesIt != framesEnd) {

        // Items left
        count = std::distance(framesIt, framesEnd);

        // If we get to the last element, just write it and end
        if (count <= 2) {
            for (i = 0; i < count; ++i) {
                if (hasWritten) {
                    buf << ",";
                }
                buf << internal::zfill(framesIt[i], zfill);
            }
            hasWritten = true;
            break;
        }

        // At this point, we have 3 or more frames to check.
        // Scan the current window of the slice to see how
        // many frames we can consume into a group
        step = framesIt[1] - framesIt[0];

        // Scan and update the pointer index until we get
        // to a point where the step changes
        for (i = 0; i < (count-1); ++i) {

            if ((framesIt[i+1] - framesIt[i]) != step) {
                // We have scanned as many frames as we can
                // for this group. Now write them and stop
                // looping on this window, because the step
                // has changed.
                break;
            }
        }

        // Subsequent groups are comma-separated
        if (hasWritten) {
            buf << ",";
            hasWritten = true;
        }

        // We only have a single frame to write for this group
        if (i == 0) {
            buf << internal::zfill(framesIt[0], zfill);
            framesIt++;
            hasWritten = true;
            continue;
        }

        // First do a check to see if we could have gotten a larger range
        // out of subsequent values with a different step size
        if (i == 1 && count > 3) {
            // Check if the next two pairwise frames have the same step.
            // If so, then it is better than our current grouping.
            if ((framesIt[2] - framesIt[1]) == (framesIt[3] - framesIt[2])) {
                // Just consume the first frame, and allow the next
                // loop to scan the new stepping
                buf << internal::zfill(framesIt[0], zfill);
                framesIt++;
                hasWritten = true;
                continue;
            }
        }

        // Otherwise write out this step range
        start = internal::zfill(framesIt[0], zfill);
        end = internal::zfill(framesIt[i], zfill);
        buf << start << "-" << end;

        if (step > 1) {
            buf << "x" << step;
        }

        hasWritten = true;

        // Advance our iterator to the end of the current window
        // so that we can test the next available number
        framesIt += (i+1);
    }

    return buf.str();
}

bool isFrameRange(const std::string &frange) {
    internal::RangeMatches matches;
    return frameRangeMatches(matches, frange);
}

FileSequence findSequenceOnDisk(const std::string &pattern, Status* ok) {
    return findSequenceOnDisk(pattern, PadStyleDefault, ok);
}

FileSequence findSequenceOnDisk(const std::string &pattern, PadStyle style, Status* ok) {
    if (ok != NULL) ok->clearError();

    Status localOk;
    FileSequence fs(pattern, style, &localOk);

    if (!localOk) {
        // Treat a bad pattern as a non-match
        std::cerr << "fileseq: " << localOk << std::endl;
        return FileSequence();
    }

    FileSequences seqs;
    localOk = findSequencesOnDisk(seqs, fs.dirname());
    if (!localOk) {
        if (ok != NULL) {
            std::ostringstream err;
            err << "Failed to find " << pattern << ": " << localOk;
            ok->setError(err.str());
        }
        return FileSequence();
    }

    const std::string &base = fs.basename();
    const std::string &ext = fs.ext();

    FileSequences::iterator it;
    for (it = seqs.begin(); it != seqs.end(); ++it) {
        // Find the first match and return it
        if (it->basename() == base && it->ext() == ext) {
            it->setPaddingStyle(style);
            return (*it);
        }
    }

    // If we get this far, we didn't find a match
    return FileSequence();
}

Status findSequencesOnDisk(FileSequences &seqs,
                           const std::string &path,
                           bool hiddenFiles,
                           bool singleFiles,
                           PadStyle style) {

    // Flags
    FindSequenceOpts opts = kNoOpt;
    if (hiddenFiles) {
        opts = opts | kOptHiddenFiles;
    }
    if (singleFiles) {
        opts = opts | kOptSingleFiles;
    }

    // Delegate
    return findSequencesOnDisk(seqs, path, opts, style);
}


namespace internal {

// Wrapper to make sure a DIR* is closed properly
// when it goes out of scope
class DirCloser {
public:
    DirCloser(DIR* d) : m_dir(d) {}

    ~DirCloser() {
        if (m_dir != NULL) {
            closedir(m_dir);
        }
    }

private:
    DIR* m_dir;
};


// Struct containing the ongoing parsing info
// about a sequence while reading a directory
struct SeqInfo {
    SeqInfo() : minWidth(0), padding(), frames() {}

    size_t minWidth;
    std::string padding;
    Frames frames;
};

// Key for sequence results map: basename, extension
typedef std::pair< const std::string, const std::string > SeqKey;
// Map type to track sequence results in a directory
typedef std::map< SeqKey, SeqInfo* > SeqsMap;


// Wrapper to make sure a SeqsMap is cleaned up properly
// when it goes out of scope
class SeqsMapCloser {
public:
    SeqsMapCloser(SeqsMap* m) : m_map(m) {}

    ~SeqsMapCloser() {
        if (m_map == NULL) {
            return;
        }

        SeqsMap::iterator it;
        for (it = m_map->begin(); it != m_map->end(); ++it) {
            if (it->second != NULL) {
                delete it->second;
                it->second = NULL;
            }
        }
    }

private:
    SeqsMap* m_map;
};


} // internal


Status findSequencesOnDisk(FileSequences &seqs,
                           const std::string &path,
                           FindSequenceOpts opts,
                           PadStyle style) {

    Status status;

    using namespace internal;

    // Options
    bool hiddenFiles = opts & kOptHiddenFiles;
    bool singleFiles = opts & kOptSingleFiles;

    // Open the directory
    DIR* dir;
    if ( (dir = opendir(path.c_str())) == NULL ) {
        status.setError(std::strerror(errno));
        return status;
    }

    DirCloser dirCloser(dir); // close when scope ends

    const PaddingMapper &padder = getPadMapperForStyle(style);

    // Track members of file sequences as we scan
    SeqsMap seqsMap;
    SeqsMapCloser seqsMapCloser(&seqsMap);

    // Track individual files (if option allows)
    FileSequences files;

    // Get our root search path and make sure it ends with
    // a path sep character
    std::string root = path;
    if ( !root.empty() && (root[root.size()-1] != strings::kPathSep) ) {
        root.append( 1, strings::kPathSep );
    }

    // Prep a string buffer that we can reuse to constantly build strings
    std::ostringstream buf(root, std::ios_base::out | std::ios_base::ate);

    Frame frame;
    size_t frameWidth = 0;
    std::string name;
    FileSequence fs;
    SeqInfo* seqInfo;
    SeqPatternMatch match;
    SeqsMap::iterator seqFound;

    struct stat f_stat;
    struct dirent* d_ent;

    errno = 0;

    // Read dir and sort files into groups
    while ( (d_ent = readdir(dir)) != NULL ) {

        // Is it a directory?
        if (d_ent->d_type == DT_DIR) {
            // Skip directories
            continue;
        }

        // Get the string name of the file
        name.assign(&(d_ent->d_name[0]));

        // Is it a hidden file?
        if (hiddenFiles && name[0] == '.' ) {
            continue;
        }

        // Is it a symlink?
        if (d_ent->d_type == DT_LNK) {

            // See if we are pointing to a directory
            buf << name;
            if ( stat(buf.str().c_str(), &f_stat) == -1 ) {

                std::stringstream err;
                err << "Error reading symlink " << buf.str()
                    << ": " << std::strerror(errno);

                // TODO: Maybe don't bail on the entire loop
                // if we can't read the symlink.
                // Report the error by some means (logger, ...)
                // and continue collecting

                status.setError(err.str());
                return status;
            }

            buf.str(root); // reset

            if (S_ISDIR(f_stat.st_mode)) {
                // Also skip symlinks that point to directories
                continue;
            }

        }

        // Are we matching a single frame (no range)?
        bool ok = getSingleFrameMatch(match, name, /*require_frame */ false);

        if (ok && (match.range.empty() || (match.base.empty() && match.ext.empty()))) {
            ok = false;
        }

        if (!ok) {
            // Only keep it if we wanted single files
            if (singleFiles) {
                buf << name;

                fs = FileSequence(buf.str(), style, &status);
                if (!status) {
                    return status;
                }
                // Preserve the parsed base/frame/ext
                fs.setBasename(match.base);
                fs.setExt(match.ext);
                if (match.range.empty()) {
                    fs.setFrameSet(FrameSet());
                    fs.setPadding("");
                } else {
                    fs.setFrameRange(match.range);
                }
                files.push_back(fs);

                buf.str(root); // reset
            }

            continue;
        }

        char* pEnd;
        frame = Frame(std::strtol(match.range.c_str(), &pEnd, 10));

        frameWidth = match.range.size();

        SeqKey key = std::make_pair(match.base, match.ext);
        seqFound = seqsMap.find(key);

        if (seqFound == seqsMap.end()) {

            // Doesn't exist yet. Create new entry
            seqInfo = new SeqInfo;
            seqInfo->minWidth = frameWidth;
            seqInfo->padding = padder.getPaddingChars(frameWidth);
            seqsMap[key] = seqInfo;


        } else {

            // Exists
            seqInfo = seqFound->second;

            // Keep trying to find the smallest padding width
            // in the sequence
            if (frameWidth < seqInfo->minWidth) {
                seqInfo->minWidth = frameWidth;
                seqInfo->padding = padder.getPaddingChars(frameWidth);
            }
        }

        seqInfo->frames.push_back(frame);

        // Got what we needed. On to the next file, if available
        errno = 0;
    }

    // Any errors reading the directory listing?
    if ( errno != 0 ) {
        status.setError( std::strerror(errno) );
        return status;
    }


    // Now that the directory listing has been read,
    // do a pass on the collected groups of files to
    // convert them into sequence results

    seqs.reserve(seqsMap.size());

    size_t pos = 0;
    char digit;
    std::string ext, pad, frange;

    SeqsMap::const_iterator seqsIt;
    for (seqsIt = seqsMap.begin(); seqsIt != seqsMap.end(); ++seqsIt) {

        seqInfo = seqsIt->second;

        name = seqsIt->first.first;
        ext = seqsIt->first.second;
        pad = seqInfo->padding;

        if (seqInfo->frames.size() == 1) {

            buf.str("");
            buf << seqInfo->frames[0];
            frange = buf.str();

            if (!name.empty()) {
                // Make sure a non-sequence file doesn't accidentally
                // get reparsed as a range.
                pos = 1;
                // Check if the parsed number was preceded by a "-",
                // if so, check before that char to see if its a number
                if (name[name.size()-1] == '-' && name.size() >= 2) {
                    pos = 2;
                }

                digit = name[name.size()-pos];
                char* p;
                std::strtoul(&digit, &p, 10);
                if (!(*p)) {
                    // If it is a number, clear the padding char
                    pad.clear();
                }
            }

        } else {

            frange = framesToFrameRange(seqInfo->frames, true, 0);
        }

        // Build the sequence string
        buf.str(root); // reset
        buf << name << frange << pad << ext;

        fs = FileSequence(buf.str(), style, &status);
        // TODO: Maybe don't bail on the entire loop
        // if we can't process this particular sequence.
        // Report the error by some means (logger, ...)
        // and continue collecting
        if (!status) {
            return status;
        }

        if (pad.empty()) {
            fs.setFrameSet(FrameSet());
        }

        // Save the FileSequence
        seqs.push_back(fs);
    }

    if (singleFiles) {
        // Copy the single file patterns
        seqs.insert(seqs.begin(), files.begin(), files.end());
    }

    return status;
}


} // fileseq
