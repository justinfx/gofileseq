#ifndef _GOFILESEQ_CPP_FRAMESET_H_
#define _GOFILESEQ_CPP_FRAMESET_H_

#include <string>
#include <vector>

namespace fileseq {

#include "private/libgofileseq.h"

class FrameSet {

public:
	explicit FrameSet(const std::string &frange);
	~FrameSet();

    FrameSet(const FrameSet& rhs);
    FrameSet& operator=(const FrameSet& rhs);

	bool isValid() const { return m_valid; }
    std::string getLastError() const { return m_lastError; }

	std::string string() const;
	size_t length() const;
	int index(int frame) const;
    int frame(int index, bool* ok=NULL) const;
    void frames(std::vector<int> &frames) const;
	bool hasFrame(int frame) const;
	int start() const;
	int end() const;
	std::string frameRange(int pad=0) const;
	FrameSet inverted() const;
	std::string invertedFrameRange(int pad) const;
	FrameSet normalized() const;

private:
    FrameSet(GoInt64 id);

	bool m_valid;
	GoInt64 m_id;

    mutable std::string m_lastError;
};

} //

#endif
