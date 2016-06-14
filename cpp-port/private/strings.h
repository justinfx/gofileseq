#ifndef FILESEQ_STRINGS_P_H_
#define FILESEQ_STRINGS_P_H_

#include <string>

namespace fileseq {

namespace strings {

// Platform-specific file path sep
const char kPathSep =
#if defined _WIN32 || defined __CYGWIN__
                      '\\';
#else
                      '/';
#endif


// Return whether a string contains another substring
bool contains(const std::string &str, const std::string &substr) {
    return (str.find(substr) != std::string::npos);
}


// replace all occurances of "from" with "to" within "str"
void replace_all(std::string &str, const std::string &from, const std::string &to) {
    size_t start_pos = 0;
    while((start_pos = str.find(from, start_pos)) != std::string::npos) {
        str.replace(start_pos, from.length(), to);
        start_pos += to.length(); // Handles case where 'to' is a substring of 'from'
    }
}


// Trim left and right side whitespace from a string
void trim(std::string &str) {
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


} // strings

} // fileseq

#endif // FILESEQ_STRINGS_P_H_
