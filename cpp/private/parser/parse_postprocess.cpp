#include "parse_postprocess.h"

#include <algorithm>
#include <cctype>

namespace fileseq {
namespace internal {

// Helper function to check if a string contains only digits
static bool isAllDigits(const std::string& str) {
    if (str.empty()) {
        return false;
    }
    return std::all_of(str.begin(), str.end(), [](unsigned char c) {
        return std::isdigit(c);
    });
}


void handleMultipleDotNumTokens(ParseResult& result) {
    // Only apply when: extension empty, frame exists, basename exists
    if (result.extension.empty() && !result.frameRange.empty() && !result.basename.empty()) {
        std::string basename = result.basename;

        // Remove trailing dot if present (added by normal frameNum processing)
        if (!basename.empty() && basename.back() == '.') {
            basename.pop_back();
        }

        // Check if basename ends with pattern ".digits"
        size_t dotIdx = basename.find_last_of('.');
        if (dotIdx != std::string::npos) {
            std::string suffix = basename.substr(dotIdx + 1);

            // Verify suffix is all digits
            if (isAllDigits(suffix)) {
                // Swap: first number → frame, second number → extension
                result.extension = "." + result.frameRange;
                result.frameRange = suffix;
                result.basename = basename.substr(0, dotIdx + 1); // Keep through the dot
            }
        }
    }
}


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


void postProcessParseResult(ParseResult& result) {
    // Apply all special case handlers in order
    // Order matters - these are applied sequentially

    // MUST run first - removes leading dot from frame range before other handlers use it
    moveLeadingDotFromFrameRange(result);

    handleMultipleDotNumTokens(result);
    extractDigitsFromExtensionToken(result);
}

} // namespace internal
} // namespace fileseq
