#include <iostream>
#include <vector>

#include "fileseq.h"

using namespace std;

void test_frameset() {
    fileseq::FrameSet fs("1-20x2");

    std::cout << "valid? " << fs.isValid() << std::endl;
    if (!fs.isValid()) {
        std::cout << "  Error:" << fs.getLastError() << std::endl;
    }
    std::cout << "string: " << fs.string() << std::endl;
    std::cout << "length: " << fs.length() << std::endl;
    std::cout << "Start: " << fs.start()
              << ", End: " << fs.end() << std::endl;
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

    fileseq::FrameSet fsCopy(fs);
    fileseq::FrameSet fsCopy2 = fsCopy;
    std::cout << "fsCopy: " << fsCopy.string() << std::endl;
    std::cout << "fsCopy2: " << fsCopy2.string() << std::endl;
}

void test_fileseqs() {
    fileseq::FileSequence fs("/path/to/file_thing.10-30x2#.ext");

    std::cout << "Valid? " << fs.isValid()
              << "\nString: " << fs
              << "\nDirname: " << fs.dirname()
              << "\nBasename: " << fs.basename()
              << "\nExt: " << fs.ext()
              << "\nStart: " << fs.start() << ", End: " << fs.end()
              << "\nZfill: " << fs.zfill()
              << "\nFrameRange: " << fs.frameRange()
              << "\nFrameRange Pad: " << fs.frameRange(true)
              << "\nInverted: " << fs.invertedFrameRange()
              << "\nInverted Pad: " << fs.invertedFrameRange(true)
              << "\nFrame: " << fs.frame(10)
              << "\nFrame: " << fs.frame("#")
              << std::endl;

    bool ok = true;
    std::string fmt = fs.format("{{dir}}::{{base}}::{{frange}}::{{pad}}::{{ext}}", &ok);
    std::cout << "format: " << fmt << (ok ? "" : "\n"+fs.getLastError()) << std::endl;

    fmt = fs.format("{{bad}}::{{ext}}", &ok);
    std::cout << "format: " << fmt << (ok ? "" : "\n"+fs.getLastError()) << std::endl;

    std::cout << "indexes:" << std::endl;
    for (size_t i = 0; i < fs.length() + 1; ++i ) {
        std::cout << "  frame at " << i << "=" << fs.index(i) << std::endl;
    }

    fileseq::FrameSet frameSet = fs.frameSet();
    std::cout << "FrameSet: " << frameSet.string()
              << " (valid? " << frameSet.isValid() << ")" << std::endl;

    fs.setDirname("/a");
    fs.setBasename("base.");
    fs.setExt(".ext2");
    ok = fs.setFrameRange("-20--10");
    std::cout << "String after setters: " << fs << " (ok? " << ok << ")"
              << "\n  FrameSet: " << fs.frameSet()
              << std::endl;
    ok = fs.setFrameRange("abc");
    std::cout << "expect an error after setting bad frame range: " << fs.getLastError()
              << "\n  FrameSet: " << fs.frameSet()
              << std::endl;

    frameSet = fileseq::FrameSet("100-120x2,50,60,65-85x4").normalized();
    fs.setFrameSet(frameSet);
    std::cout << "String after setting FrameSet: " << fs << std::endl;
}

void test_find_seqs() {
    std::string path = "../testdata/seqB.#.jpg";
    std::string pathBad = "asdlkasdkls---___--__adkl.######.jpg";

    std::string err;

    fileseq::FileSequence fs = fileseq::findSequenceOnDisk(path, &err);
    if (!err.empty()) {
        std::cerr << "findSequenceOnDisk failed: " << err << std::endl;
    } else {
        std::cout << "findSequenceOnDisk (valid? " << fs.isValid() << "): "
                  << fs << std::endl;
    }

    fs = fileseq::findSequenceOnDisk(pathBad, &err);
    std::cout << "Got expected error for bad path? " << (!err.empty())
              << "\n  " << err
              << std::endl;

    std::string path2 = "../testdata";
    std::cout << "findSequencesOnDisk for " << path2 << std::endl;
    fileseq::FileSequences seqs;
    seqs.push_back(fileseq::FileSequence("existing.1-100#.ext"));
    std::string err2 = fileseq::findSequencesOnDisk(seqs, path2, true, true);
    if (!err2.empty()) {
        std::cout << "  Error: " << err2 << std::endl;
    } else {
        for (size_t i=0; i < seqs.size(); ++i) {
            std::cout << "  seq: " << seqs.at(i) << std::endl;
        }
    }

}

int main()
{
    test_frameset();
    test_fileseqs();
    test_find_seqs();

    std::cout << "Final allocs: " << fileseq::allocStats() << std::endl;

    return 0;
}

