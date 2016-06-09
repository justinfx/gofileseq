#include "fileseq.h"
#include "private/frameset_p.h"

#include <algorithm>
#include <iterator>
#include <string>
#include <sstream>

namespace fileseq {

std::string framesToFrameRange(const Frames &frames,
                               bool sorted, int zfill) {

    size_t count = frames.size();
    if (count == 0) {
        return "";
    }

    if (count == 1) {
        return internal::zfill(frames[0], zfill);
    }

    const Frames* framesPtr = &frames;

    if (sorted) {
        Frames sortedFrames = frames;
        std::sort(sortedFrames.begin(), sortedFrames.end());
        framesPtr = &sortedFrames;
    }

    size_t i = 0;
    long step = 0;
    bool hasWritten = false;
    std::string start, end;
    std::stringstream buf;

    Frames::const_iterator framesIt = framesPtr->begin();
    Frames::const_iterator framesEnd = framesPtr->end();

    while (framesIt < framesEnd) {

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
        for (i = 0; i < (count-1); ++i) {
            // We have scanned as many frames as we can
            // for this group. Now write them and stop
            // looping on this window
            if ((framesIt[i+1] - framesIt[i]) != step) {
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
        framesIt++;
        hasWritten = true;
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
    return FileSequence("");
}

Status findSequencesOnDisk(FileSequences &seqs,
                           const std::string &path,
                           bool hiddenFiles,
                           bool singleFiles,
                           PadStyle style) {

    Status stat;
    stat.setError("Not implemented");
    return stat;
}

} // fileseq
