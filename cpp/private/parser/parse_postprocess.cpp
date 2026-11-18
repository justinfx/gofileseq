#include "parse_postprocess.h"

#include <cctype>

namespace fileseq {
namespace internal {

void extractDigitsFromExtensionToken(ParseResult& result) {
    // Only apply when: padding exists, no frame range yet, basename exists
    if (!result.padding.empty() && result.frameRange.empty() && !result.basename.empty()) {
        std::string basename = result.basename;

        // Find trailing digits in basename
        size_t i = basename.length();
        while (i > 0 && std::isdigit(static_cast<unsigned char>(basename[i-1]))) {
            --i;
        }

        // If we found trailing digits (and there's content before them), extract as frame
        if (i < basename.length() && i > 0) {
            result.frameRange = basename.substr(i);
            result.basename = basename.substr(0, i);
        }
    }
}


void moveLeadingDotFromFrameRange(ParseResult& result) {
    if (!result.frameRange.empty() && result.frameRange[0] == '.') {
        // Move the dot from frame range to basename
        result.basename += '.';
        result.frameRange = result.frameRange.substr(1); // Remove leading dot
    }
}

void splitDecimalStepFromFrameRange(ParseResult& result) {
    std::string::size_type i = result.frameRange.rfind('.');
    if (i != std::string::npos) {
        result.basename += result.frameRange.substr(0, i + 1);
        result.frameRange = result.frameRange.substr(i + 1);
    }
}

} // namespace internal
} // namespace fileseq
