
#include "strings.h"

namespace fileseq {

namespace strings {


bool starts_with(const std::string &str, const std::string &prefix) {
	if (str.length() < prefix.length()) {
		return false;
	}
	return str.rfind(prefix, 0) == 0;
}


bool ends_with(const std::string &str, const std::string &suffix) {
	if (str.length() < suffix.length()) {
		return false;
	}
	return str.rfind(suffix) == str.size() - suffix.size();
}


bool contains(const std::string &str, const std::string &substr) {
    return (str.find(substr) != std::string::npos);
}


void replace_all(std::string &str, const std::string &from, const std::string &to) {
    size_t start_pos = 0;
    while((start_pos = str.find(from, start_pos)) != std::string::npos) {
        str.replace(start_pos, from.length(), to);
        start_pos += to.length(); // Handles case where 'to' is a substring of 'from'
    }
}


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

