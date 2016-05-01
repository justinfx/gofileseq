#ifndef _GOFILESEQ_CPP_FILESEQ_P_H_
#define _GOFILESEQ_CPP_FILESEQ_P_H_

#include <iostream>
#include <string>

namespace fileseq {

namespace internal {

#include "libgofileseq.h"

// StringProxy manages a char* returned from
// the internal API, which needs to be freed
// after it has been converted to a std::string
class StringProxy {

public:
    StringProxy()
        : m_data(NULL)
        , m_str() {}

    StringProxy(char* managed)
        : m_data(managed)
        , m_str() {

        if (managed != NULL) {
            m_str.assign(managed);
        }
    }

    ~StringProxy() {
        if (m_data != NULL) {
            free(m_data);
            m_data = NULL;
        }
    }

    StringProxy(const StringProxy& rhs)
        : m_data(NULL)
        , m_str(rhs.m_str) {}

    StringProxy& operator=(const StringProxy& rhs) {
        if (&rhs == this) {
            return *this;
        }
        if (m_data != NULL) {
            free(m_data);
        }
        m_data = NULL;
        m_str = rhs.m_str;
        return *this;
    }

    operator std::string() const { return m_str; }

private:
    char* m_data;
    std::string m_str;
};

} // internal

} // fileseq

#endif // _GOFILESEQ_CPP_FILESEQ_P_H_

