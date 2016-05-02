#ifndef _GOFILESEQ_CPP_FILESEQ_H_
#define _GOFILESEQ_CPP_FILESEQ_H_

#include "frameset.h"
#include "sequence.h"
#include "types.h"

#include <vector>

namespace fileseq {

// FindSequenceOnDisk takes a string that is a compatible/parsible
// FileSequence pattern, and finds a sequence on disk which matches
// the basename and extension.
// If no match is found, an invalid FileSequence is returned.
// If an error occurs while reading the filesystem, it can be
// captured by passing a fileseq::Status pointer.
FileSequence findSequenceOnDisk(const std::string &pattern,
                                Status* ok=NULL);

// FindSequencesOnDisk searches a given directory path and sorts all
// valid sequence-compatible files into a list of FileSequence matches.
// By default, only non-hidden sequences of files will be returned.
// Extra options may be given to control whether things like hidden
// files, or single (non-sequence) files should be included in the
// search results.
Status findSequencesOnDisk(FileSequences &seqs,
                           const std::string &path,
                           bool hiddenFiles=false,
                           bool singleFiles=false);


// Debug
std::string allocStats();

} // fileseq

#endif // _GOFILESEQ_CPP_FILESEQ_H_
