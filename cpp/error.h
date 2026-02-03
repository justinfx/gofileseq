#ifndef FILESEQ_ERROR_H_
#define FILESEQ_ERROR_H_

#include <iostream>
#include <ostream>
#include <string>


#ifdef __GNUC__
#define DEPRECATED(func) func __attribute__ ((deprecated))
#elif defined(_MSC_VER)
#define DEPRECATED(func) __declspec(deprecated) func
#else
#pragma message("WARNING: You need to implement DEPRECATED for this compiler")
#define DEPRECATED(func) func
#endif


namespace fileseq {


/*!
Status is used to communicate whether various calls
have succeeded or failed. It can be treated like a
bool to test the success, and converted to a
std::string to receive the error message.
*/
class Status {

public:
    Status() = default;

    //! If Status evaluates to false, it indicates an error state
    operator bool() const { return m_error.empty(); }

    //! Retrieve the string error message
    operator std::string() const { return m_error; }

    Status(const Status& rhs) : m_error(rhs.m_error) {}

    Status& operator=(const Status& rhs) {
        m_error = rhs.m_error;
        return *this;
    }

    void clearError() { m_error.clear(); }
    void setError(const std::string &err) { m_error = err; }
    void setError(const char* err) { m_error = err; }

    friend std::ostream& operator<< (std::ostream& stream, const Status& stat) {
        stream << (std::string)stat;
        return stream;
    }

private:
    std::string m_error;

};


namespace internal {

void handleErrorStatus(const std::string &msg, Status* stat= nullptr);

void setError(const std::string &msg, Status* stat);

} // internal


} // fileseq

#endif // FILESEQ_ERROR_H_
