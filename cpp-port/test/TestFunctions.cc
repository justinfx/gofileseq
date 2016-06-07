#include "common.h"
#include "../fileseq.h"

#include <gtest/gtest.h>

#include <algorithm>
#include <cstring>
#include <map>

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
    TODO();

    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        // TODO: When PadMapper function types are available to test
    }
}


class TestPaddingCharsSize : public testing::Test {

public:
    struct Case {
        fileseq::PadStyle padStyle;
        std::string padChars;
        int pad;
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
    TODO();

    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        // TODO: When PadMapper function types are available to test
    }
}

void testFindSequencesOnDisk(bool singleFiles) {
    std::map<std::string, bool> cases;

    cases["seqC.-5-2,4-10,20-21,27-30@@.tif"] = false;
    cases["seqB.5-14,16-18,20#.jpg"] = false;
    cases["seqA.1,3-6,8-10#.exr"] = false;

    if (singleFiles) {
        cases["aFile.ext"] = false;
        cases["file_without_ext"] = false;
    }

    fileseq::Status stat;
    fileseq::FileSequences seqs;

    stat = fileseq::findSequencesOnDisk(seqs, "testdata",
                                        false,      // Hiden files
                                        singleFiles // Single files
                                        );

    ASSERT_TRUE(stat) << "Failed to find seqs in location 'testdata': " << stat;
    ASSERT_NE(0, seqs.size()) << "Empty sequence search results";

    std::map<std::string, bool>::iterator mapIt;

    for (size_t i=0; i < seqs.size(); ++i) {
        std::string fullpath = seqs[i].string();
        std::string name = basename(fullpath.c_str());

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
    stat = fileseq::findSequencesOnDisk(seqs, "testdata/versions");

    ASSERT_TRUE(stat) << "Failed to find seqs in location 'testdata/versions': " << stat;
    ASSERT_EQ(1, seqs.size()) << "Did not find expected number of seqs";

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