#include "common.h"

#include "../fileseq.h"
#include "../pad.h"

#include <gtest/gtest.h>

#include <algorithm>
#include <cstring>
#include <iostream>
#include <map>
#include <string>


// Fwd decl
std::string basename(const std::string& s);


class TestFramesToFrameRange : public testing::Test {

public:
    struct Case {
        fileseq::Frames frames;
        bool sorted;
        int zfill;
        std::string expected;
    };

protected:
    void SetUp() {
        {
            fileseq::Frame arr[] = {-10};
            Case t = {FRAME_ARR_TO_VEC(arr), false, 4, "-010"};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {1, 2, 3, 4, 5};
            Case t = {FRAME_ARR_TO_VEC(arr), true, 0, "1-5"};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {1, 3, 5, 7, 9};
            Case t = {FRAME_ARR_TO_VEC(arr), true, 0, "1-9x2"};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {1, 2, 3, 4, 5, 20, 30, 40, 41, 42, 43, 44, 50, 60, 62, 64, 66};
            Case t = {FRAME_ARR_TO_VEC(arr), true, 0, "1-5,20-40x10,41-44,50,60-66x2"};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {1, 3, 6, 9, 12, 20, 22, 24, 28, 32};
            Case t = {FRAME_ARR_TO_VEC(arr), true, 0, "1,3-12x3,20-24x2,28,32"};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {1, 11, 21, 31, 41, 51, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 15, 18};
            Case t = {FRAME_ARR_TO_VEC(arr), false, 2, "01-51x10,02-10,12-18x3"};
            m_cases.push_back(t);
        }
    }

    std::vector<Case> m_cases;
};


TEST_F( TestFramesToFrameRange, FramesToFrameRange ) {
    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        EXPECT_EQ(t.expected, fileseq::framesToFrameRange(t.frames, t.sorted, t.zfill));
    }
}


class TestPaddingChars : public testing::Test {

public:
    struct Case {
        fileseq::PadStyle padStyle;
        int pad;
        std::string padChars;
    };

protected:
    void SetUp() {
        using namespace fileseq;

        {Case t = {PadStyleHash1, -1, "#"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, 0, "#"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, 1, "#"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, 2, "##"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, 3, "###"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, 1, "#"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, 5, "#####"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, 2, "##"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, 3, "###"}; m_cases.push_back(t);}

        {Case t = {PadStyleHash4, -1, "@"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, 0, "@"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, 1, "@"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, 2, "@@"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, 3, "@@@"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, 4, "#"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, 5, "@@@@@"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, 8, "##"}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, 12, "###"}; m_cases.push_back(t);}
    }

    std::vector<Case> m_cases;
};


TEST_F( TestPaddingChars, Chars ) {
    using fileseq::internal::PaddingMapper;
    using fileseq::internal::getPadMapperForStyle;

    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        const PaddingMapper &mapper = getPadMapperForStyle(t.padStyle);
        EXPECT_EQ(t.padChars, mapper.getPaddingChars(t.pad));
    }
}


class TestPaddingCharsSize : public testing::Test {

public:
    struct Case {
        fileseq::PadStyle padStyle;
        std::string padChars;
        size_t pad;
    };

protected:
    void SetUp() {
        using namespace fileseq;

        {Case t = {PadStyleHash1, "", 0}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, "#", 1}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, "##", 2}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, "###", 3}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, "####", 4}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, "@", 1}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, "@@", 2}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, "@@@", 3}; m_cases.push_back(t);}
        {Case t = {PadStyleHash1, "@@@@", 4}; m_cases.push_back(t);}

        {Case t = {PadStyleHash4, "", 0}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, "#", 4}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, "##", 8}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, "###", 12}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, "####", 16}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, "@", 1}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, "@@", 2}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, "@@@", 3}; m_cases.push_back(t);}
        {Case t = {PadStyleHash4, "@@@@", 4}; m_cases.push_back(t);}
    }

    std::vector<Case> m_cases;
};


TEST_F( TestPaddingCharsSize, CharsSize ) {
    using fileseq::internal::PaddingMapper;
    using fileseq::internal::getPadMapperForStyle;

    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        const PaddingMapper &mapper = getPadMapperForStyle(t.padStyle);

        EXPECT_EQ(t.pad, mapper.getPaddingCharsSize(t.padChars))
            << "Given chars '"<< t.padChars << "' and pad style "
            << std::string((t.padStyle == fileseq::PadStyleHash1) ? "hash1" : "hash4");
    }
}

void testFindSequencesOnDisk(bool singleFiles) {
    std::map<std::string, bool> cases;

    cases["seqA.1,3-6,8-10#.exr"] = false;
    cases["seqB.5-14,16-18,20#.jpg"] = false;
    cases["seqC.-5-2,4-10,20-21,27-30@@.tif"] = false;
    cases["seqD.2-10@.gif"] = false;
    cases["complex.5-7#.tar.gz"] = false;

    fileseq::FindSequenceOpts opts = fileseq::kNoOpt;

    if (singleFiles) {
        opts = opts | fileseq::kOptSingleFiles;
        cases["aFile.ext"] = false;
        cases["file_without_ext"] = false;
    }

    fileseq::Status stat;
    fileseq::FileSequences seqs;

    stat = fileseq::findSequencesOnDisk(seqs, "testdata", opts);

    ASSERT_TRUE(stat) << "Failed to find seqs in location 'testdata': " << stat;
    ASSERT_NE(size_t(0), seqs.size()) << "Empty sequence search results";

    std::map<std::string, bool>::iterator mapIt;

    for (size_t i=0; i < seqs.size(); ++i) {
        std::string fullpath = seqs[i].string();
        std::string name = basename(fullpath);

        if ((mapIt = cases.find(name)) == cases.end()) {
            ADD_FAILURE() << "Found unexpected seq not in cases: " << name;
            continue;
        }
        // Found
        mapIt->second = true;
    }

    for (mapIt = cases.begin(); mapIt != cases.end(); ++mapIt) {
        EXPECT_TRUE(mapIt->second) << "Expected seq was not found: " << mapIt->first;
    }
}

TEST( TestFindSequencesOnDisk, WithSingleFiles ) {
    testFindSequencesOnDisk(true);
}

TEST( TestFindSequencesOnDisk, SeqsOnly ) {
    testFindSequencesOnDisk(false);
}

TEST( TestFindSequencesOnDisk, HandleSymlinksOnDisk ) {
    fileseq::Status stat;
    fileseq::FileSequences seqs;
    stat = fileseq::findSequencesOnDisk(seqs, "testdata/versions", fileseq::kNoOpt);

    ASSERT_TRUE(stat) << "Failed to find seqs in location 'testdata/versions': " << stat;
    ASSERT_EQ(size_t(1), seqs.size()) << "Did not find expected number of seqs";

    std::string expected = "testdata/versions/seq.1-10#.ext";
    ASSERT_EQ(expected, seqs[0].string());
}


class TestFindSequenceOnDisk : public testing::Test {

public:
    struct Case {
        fileseq::PadStyle padStyle;
        std::string input;
        std::string expected;
    };

protected:
    void SetUp() {
        using namespace fileseq;

        // PadStyleHash1
        {
            Case t = {PadStyleHash1, "testdata/seqC.@@.tif", "testdata/seqC.-5-2,4-10,20-21,27-30##.tif"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash1, "testdata/seqC.0010.tif", "testdata/seqC.-5-2,4-10,20-21,27-30##.tif"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash1, "testdata/seqB.#.jpg", "testdata/seqB.5-14,16-18,20####.jpg"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash1, "testdata/seqB.16-18#.jpg", "testdata/seqB.5-14,16-18,20####.jpg"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash1, "testdata/seqA.#.exr", "testdata/seqA.1,3-6,8-10####.exr"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash1, "testdata/seqA.@.exr", "testdata/seqA.1,3-6,8-10####.exr"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash1, "testdata/seqA.@.jpg", ""};
            m_cases.push_back(t);
        }

        // PadStyleHash4
        {
            Case t = {PadStyleHash4, "testdata/seqC.@@.tif", "testdata/seqC.-5-2,4-10,20-21,27-30@@.tif"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash4, "testdata/seqC.0010.tif", "testdata/seqC.-5-2,4-10,20-21,27-30@@.tif"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash4, "testdata/seqB.#.jpg", "testdata/seqB.5-14,16-18,20#.jpg"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash4, "testdata/seqB.16-18#.jpg", "testdata/seqB.5-14,16-18,20#.jpg"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash4, "testdata/seqA.#.exr", "testdata/seqA.1,3-6,8-10#.exr"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash4, "testdata/seqA.@.exr", "testdata/seqA.1,3-6,8-10#.exr"};
            m_cases.push_back(t);
        }
        {
            Case t = {PadStyleHash4, "testdata/seqA.@.jpg", ""};
            m_cases.push_back(t);
        }
    }

    std::vector<Case> m_cases;
};


TEST_F( TestFindSequenceOnDisk, FindSeq ) {
    using namespace fileseq;

    Status stat;
    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        FileSequence seq = findSequenceOnDisk(t.input, t.padStyle, &stat);
        ASSERT_TRUE(stat) << "#" << i << ": " << stat;

        if ( !seq.isValid() && !t.expected.empty() ) {
            FAIL() << "Expected " << t.expected << "; got an invalid seq";
        }

        if (!t.expected.empty()) {
            EXPECT_EQ(t.expected, seq.string());
        }
    }
}

// https://www.safaribooksonline.com/library/view/c-cookbook/0596007612/ch10s15.html
std::string basename(const std::string& s) {
   char sep = '/';

#ifdef _WIN32
   sep = '\\';
#endif

   size_t i = s.rfind(sep, s.length());
   if (i != std::string::npos) {
      return(s.substr(i+1, s.length() - i));
   }

   return "";
}

