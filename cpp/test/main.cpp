#include <iostream>
#include <vector>

#include "frameset.h"

using namespace std;

void test_frameset() {
    fileseq::FrameSet fs("1-20x2");

    std::cout << "valid? " << fs.isValid() << std::endl;
    if (!fs.isValid()) {
        std::cout << "  Error:" << fs.getLastError() << std::endl;
    }
    std::cout << "string: " << fs.string() << std::endl;
    std::cout << "length: " << fs.length() << std::endl;
    std::cout << "Start: " << fs.start() << ", End: " << fs.end() << std::endl;
    std::cout << "index 1: " << fs.index(1) << std::endl;
    std::cout << "index 6: " << fs.index(6) << std::endl;
    std::cout << "index 15: " << fs.index(15) << std::endl;

    std::cout << "frameRange: " << fs.frameRange() << std::endl;
    std::cout << "frameRange(pad=6): " << fs.frameRange(6) << std::endl;

    bool ok;
    std::cout << "frame 1: " << fs.frame(1, &ok);
        if (!ok) std::cout << " (error was: " << fs.getLastError() << ")";
        std::cout << std::endl;

    std::cout << "frame 10: " << fs.frame(9, &ok);
        if (!ok) std::cout << " (error was: " << fs.getLastError() << ")";
        std::cout << std::endl;

    std::cout << "frame 20: " << fs.frame(20, &ok);
        if (!ok) std::cout << " (error was: " << fs.getLastError() << ")";
        std::cout << std::endl;

    std::vector<int> frames;

    fileseq::FrameSet fsEmpty("");
    fsEmpty.frames(frames);
    std::cout << "frames of empty FrameSet: " << frames.size()
              << ", err: " << fsEmpty.getLastError() << std::endl;

    fs.frames(frames);
    std::cout << "frames (size " << frames.size() << "): ";
    for (size_t i = 0; i < frames.size(); ++i) {
        std::cout << frames[i] << ",";
    }
    std::cout << std::endl;

    std::cout << "inverted padded (3) range: " << fs.invertedFrameRange(3) << std::endl;

    fileseq::FrameSet invert = fs.inverted();
    std::cout << "inverted valid? " << invert.isValid() << std::endl;
    std::cout << "inverted string: " << invert.string() << std::endl;
    std::cout << "inverted frameRange: " << invert.frameRange() << std::endl;

    fs = fileseq::FrameSet("1,1,2,2,3,8,4,20,15,17,16");
    std::cout << "string: " << fs.string() << std::endl;
    fs = fs.normalized();
    std::cout << "normalized valid? " << fs.isValid() << std::endl;
    std::cout << "normalized string: " << fs.string() << std::endl;
}

int main()
{
    test_frameset();
    return 0;
}

