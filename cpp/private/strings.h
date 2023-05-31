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

// Return whether a string starts with a given prefix
bool starts_with(const std::string &str, const std::string &prefix);


// Return whether a string ends with a given suffix
bool ends_with(const std::string &str, const std::string &suffix);


// Return whether a string contains another substring
bool contains(const std::string &str, const std::string &substr);


// replace all occurrences of "from" with "to" within "str"
void replace_all(std::string &str, const std::string &from, const std::string &to);


// Trim left and right side whitespace from a string
void trim(std::string &str);


// Split a given path string into the dirname and basename output vars.
// If path cannot be split, then dirname will be empty, and basename
// will be set to the original path.
//
// dirname will preserve the trailing slash.
void path_split(std::string &dirname, std::string &basename, const std::string &path);


} // strings

} // fileseq

#endif // FILESEQ_STRINGS_P_H_
