#ifndef FILESEQ_STRINGS_P_H_
#define FILESEQ_STRINGS_P_H_

#include <string>

namespace fileseq {


// Platform-specific file path sep
const char kPathSep =
#if defined _WIN32 || defined __CYGWIN__
                      '\\';
#else
                      '/';
#endif


// Trim left and right side whitespace from a string
void trim_string(std::string &str) {
    static const std::string trim_chars(" \t\n\r");

    size_t pos;

    pos = str.find_last_not_of(trim_chars);
    if ( pos != std::string::npos ) {
        str.substr(0, pos+1).swap(str);
    }

    pos = str.find_first_not_of(trim_chars);
    if ( pos != std::string::npos ) {
        str.substr(pos).swap(str);
    }
}


// Split a given path string into the dirname and basename output vars.
// If path cannot be split, then dirname will be empty, and basename
// will be set to the original path.
//
// dirname will preserve the trailing slash.
void path_split(std::string &dirname, std::string &basename, const std::string &path) {
    size_t found = path.find_last_of("/\\");

    if (found == std::string::npos) {
        dirname.clear();
        basename = path;
        return;
    }

    dirname = path.substr(0, found);
    basename = path.substr(found+1);

    if ( dirname[dirname.size()-1] != kPathSep ) {
        // Preserve the trailing slash
        dirname.append(1, kPathSep);
    }
}


} // fileseq

#endif // FILESEQ_STRINGS_P_H_
