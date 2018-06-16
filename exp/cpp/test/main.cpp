#include <cassert>
#include <ctime>
#include <cstdio>
#include <iostream>
#include <vector>

// #include <fileseq/fileseq.h>
#include "../fileseq.h"

using namespace std;

void test_frameset() {
    fileseq::Status ok;
    fileseq::FrameSet fs("1-20x2", &ok);

    std::cout << "valid? " << fs.isValid() << std::endl;
    if (!ok) {
        std::cout << "  Error:" << ok << std::endl;
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

    std::cout << "frame 1: " << fs.frame(1, &ok);
        if (!ok) std::cout << " (error was: " << ok << ")";
        std::cout << std::endl;

    std::cout << "frame 10: " << fs.frame(9, &ok);
        if (!ok) std::cout << " (error was: " << ok << ")";
        std::cout << std::endl;

    std::cout << "frame 20: " << fs.frame(20, &ok);
        if (!ok) std::cout << " (error was: " << ok << ")";
        std::cout << std::endl;

    std::vector<int> frames;

    fileseq::FrameSet fsEmpty("", &ok);
    fsEmpty.frames(frames);
    std::cout << "frames of empty FrameSet: " << frames.size()
              << ", err: " << ok << std::endl;

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

    fs = fileseq::FrameSet("1,1,2,2,3,8,4,20,15,17,16", &ok);
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
    fileseq::Status ok;

    fileseq::FileSequence fs("/path/to/file_thing.10-30x2#.ext", &ok);
    assert (ok);

    std::cout << "Valid? " << fs.isValid()
              << "\nString: " << fs
              << "\nDirname: " << fs.dirname()
              << "\nBasename: " << fs.basename()
              << "\nExt: " << fs.ext()
              << "\nStart: " << fs.start() << ", End: " << fs.end()
              << "\nZfill: " << fs.zfill()
              << "\nZfill: " << fs.padding()
              << "\nFrameRange: " << fs.frameRange()
              << "\nFrameRange Pad: " << fs.frameRange(true)
              << "\nInverted: " << fs.invertedFrameRange()
              << "\nInverted Pad: " << fs.invertedFrameRange(true)
              << "\nFrame: " << fs.frame(10)
              << "\nFrame: " << fs.frame("#")
              << std::endl;

    std::string fmt = fs.format("{{dir}}::{{base}}::{{frange}}::{{pad}}::{{ext}}", &ok);
    assert (ok);
    std::cout << "format: " << fmt << (ok ? "" : "\n" + (std::string)ok) << std::endl;

    fmt = fs.format("{{bad}}::{{ext}}", &ok);
    assert (!ok);
    std::cout << "format: " << fmt << (ok ? "" : "\n"+ (std::string)ok) << std::endl;

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
    fs.setPadding("##@@@");
    fs.setFrameRange("-20--10", &ok);
    assert (ok);
    std::cout << "String after setters: " << fs << " (ok? " << bool(ok) << ")"
              << "\n  FrameSet: " << fs.frameSet()
              << std::endl;
    fs.setFrameRange("abc", &ok);
    assert (!ok);
    std::cout << "expect an error after setting bad frame range: " << ok
              << "\n  FrameSet: " << fs.frameSet()
              << std::endl;

    frameSet = fileseq::FrameSet("100-120x2,50,60,65-85x4").normalized();
    fs.setFrameSet(frameSet);
    std::cout << "String after setting FrameSet: " << fs << std::endl;

    // Test padding
    fs = fileseq::FileSequence("/path/to/file_thing.10-30x2##.ext", &ok);
    assert (ok);
    assert (fs.paddingStyle() == fileseq::PadStyleHash4);
    assert (fs.zfill() == 8);
    assert (fs.padding() == "##");
    fs.setPaddingStyle(fileseq::PadStyleHash1);
    assert (fs.paddingStyle() == fileseq::PadStyleHash1);
    assert (fs.zfill() == 8);
    assert (fs.padding() == "########");

    fs = fileseq::FileSequence("/path/to/file_thing.10-30x2##.ext", fileseq::PadStyleHash1, &ok);
    assert (ok);
    assert (fs.paddingStyle() == fileseq::PadStyleHash1);
    assert (fs.zfill() == 2);
    assert (fs.padding() == "##");
    fs.setPaddingStyle(fileseq::PadStyleHash4);
    assert (fs.paddingStyle() == fileseq::PadStyleHash4);
    assert (fs.zfill() == 2);
    assert (fs.padding() == "@@");
}

void test_find_seqs() {
    fileseq::Status ok;

    std::string path = "../test/testdata/seqB.#.jpg";
    std::string expect = "../test/testdata/seqB.5-14,16-18,20#.jpg";

    fileseq::FileSequence fs = fileseq::findSequenceOnDisk(path, &ok);
    if (!ok) {
        std::cerr << "findSequenceOnDisk failed: " << ok << std::endl;
        assert (ok);
    }
    assert (fs.string() == expect);

    std::string pathBad = "../../asdlkasdkls";
    fs = fileseq::findSequenceOnDisk(pathBad, &ok);
    std::cout << "isValid (we expect false)? " << fs.isValid()
              << (ok ? "" : "\n  error: ") << ok
              << std::endl;
    assert (!fs.isValid());

    path = "../test/testdata/seqB.####.jpg";
    expect = "../test/testdata/seqB.5-14,16-18,20####.jpg";
    std::cout << "findSequenceOnDisk for " << path << std::endl;
    fs = fileseq::findSequenceOnDisk(path, fileseq::PadStyleHash1, &ok);
    if (!ok) std::cout << "  Error: " << ok << std::endl;
    assert (ok);
    assert (fs.string() == expect);

    path = "../test/testdata";
    std::cout << "findSequencesOnDisk for " << path << std::endl;
    fileseq::FileSequences seqs;
    seqs.push_back(fileseq::FileSequence("existing.1-100#.ext"));
    ok = fileseq::findSequencesOnDisk(seqs, path, true, true);
    assert (ok);
    if (!ok) {
        std::cout << "  Error: " << ok << std::endl;
    } else {
        for (size_t i=0; i < seqs.size(); ++i) {
            std::cout << "  seq: " << seqs.at(i) << std::endl;
        }
    }

}

void test_functions() {
    assert (fileseq::getPaddingChars(1) == "@");
    assert (fileseq::getPaddingChars(2) == "@@");
    assert (fileseq::getPaddingChars(4) == "#");
    assert (fileseq::getPaddingChars(6) == "@@@@@@");

    assert (fileseq::isFrameRange("1"));
    assert (fileseq::isFrameRange("1-100x2"));
    assert (fileseq::isFrameRange("1-100x2,300,400-500x5#"));
    assert (!fileseq::isFrameRange("foo"));
    assert (!fileseq::isFrameRange("a-b"));
    assert (!fileseq::isFrameRange("1-b"));

    assert (fileseq::padFrameRange("1-100", 4) == "0001-0100");
    assert (fileseq::padFrameRange("10-100", 3) == "010-100");

    std::vector<int> frames;

    for (int i=0; i <= 100; ++i) {
        if (i % 2 == 0) {
            frames.push_back(i);
        }
    }

    std::string frange = fileseq::framesToFrameRange(frames);
    std::cout << "got frange: " << frange << std::endl;
    assert (frange == "0-100x2");
}

int main()
{
    test_functions();
    test_frameset();
    test_fileseqs();
    test_find_seqs();

    std::cout << "\nFinal allocs: " << fileseq::allocStats() << std::endl;

    // const int n = 100000;

    // std::string str;
    // int num;

    // clock_t t;
    // t = clock();

    // for (int i=0; i < n; ++i) {
    //     fileseq::FileSequence fs("/path/to/file_name.1-100x2#.ext");
    //     str = fs.string();
    //     str = fs.frameRange();
    //     num = fs.start();
    //     num = fs.end();
    // }

    // t = clock() - t;
    // printf("%0.3f seconds\n", ((double)t)/CLOCKS_PER_SEC);

    // (void*)&num;

    return 0;
}

