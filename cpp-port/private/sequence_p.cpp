
#include "sequence_p.h"

#include <pcrecpp.h>

#include <string>
#include <vector>

namespace fileseq {

namespace internal {


bool getSplitPatternMatch(SeqPatternMatch &match, const std::string &path) {
    // Regular expression for matching a file sequence string.
    // Example:
    //     /film/shot/renders/hero_bty.1-100#.exr
    //     /film/shot/renders/hero_bty.@@.exr
    static const std::string pattern =
        "^(?P<name>.*?)(?P<range>[\\d-][:xy\\d,-]*)?(?P<pad>[#@]+)(?P<ext>(?:\\.[a-zA-Z0-9]+)*)$";
    static const pcrecpp::RE* rx = new pcrecpp::RE(pattern);

    return rx->FullMatch(path, &(match.base), &(match.range),
                               &(match.padChars), &(match.ext));

}


bool getSingleFrameMatch(SeqPatternMatch &match, const std::string &path) {
    // Regular expression for matching a single frame string.
    // Example:
    //     /film/shot/renders/hero_bty.100.exr
    static const std::string pattern =
        "^(?P<name>.*?)(?P<frame>-?\\d+)(?P<ext>(?:\\.[a-zA-Z0-9]+)*)$";
    static const pcrecpp::RE* rx = new pcrecpp::RE(pattern);

    // Not using this field
    match.padChars.clear();

    return rx->FullMatch(path, &(match.base), &(match.range), &(match.ext));
}


} // internal

} // fileseq
