#include "common.h"
#include <gtest/gtest.h>

#include "../fileseq.h"

#include <string>
#include <vector>


class TestFrameSetRanges : public testing::Test {

public:
    struct Case {
        std::string range;
        fileseq::Frames frames;
    };

protected:
    void SetUp() {
        {
            fileseq::Frame arr[] = {10};
            Case t = {"10", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {1, 2, 3, 4, 5};
            Case t = {"1-5", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {1, 6};
            Case t = {"1-10x5", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {1, 3, 5, 7, 9};
            Case t = {"1-10x2", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {1, 5, 9, 4, 7, 10, 3, 2, 6, 8};
            Case t = {"1-10:4", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {2, 4, 6, 8, 10};
            Case t = {"1-10y2", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {2, 3, 4, 6, 7, 8, 10, 11, 12, 14, 15};
            Case t = {"1-15y4", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {1, 2, 3, 20, 30, 31, 32, 33};
            Case t = {"1-3,20,30-33", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {1, 3, 5, 7, 9, 50, 60, 61, 62, 70, 72, 74};
            Case t = {" 1 - 10x2, 50, 60 - 62, 70 - 74x2 ", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {10, 9, 8, 7, 6, 5, 4, 3, 2, 1};
            Case t = {"10-1", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {10, 9, 8, 7, 6, 5, 4, 3, 2, 1};
            Case t = {"10-1x-1", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
        {
            fileseq::Frame arr[] = {10, 9, 8, 7, 6, 5, 4, 3, 2, 1};
            Case t = {"10-1x1", FRAME_ARR_TO_VEC(arr)};
            m_cases.push_back(t);
        }
    }

    std::vector<Case> m_cases;
};


TEST_F( TestFrameSetRanges, NewFrameSet ) {
    fileseq::Status stat;

    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        ASSERT_TRUE(fileseq::isFrameRange(t.range))
            << "#" << i << ": did not validate as a frame range: " << t.range;

        fileseq::FrameSet s(t.range, &stat);
        ASSERT_TRUE(stat)
            << "#" << i << ": Failed to parse " << t.range << " : " << stat;

        ASSERT_TRUE(s.isValid())
            << "#" << i << ": Expected valid FrameSet for '" << t.range << "'";

        EXPECT_NE("", s.frameRange())
            << "#" << i << ": Got an empty frange field on FrameSet";

        EXPECT_EQ(t.frames.size(), s.length());

        fileseq::Frames actual;
        s.frames(actual);
        EXPECT_EQ(t.frames, actual);
    }
};

TEST( TestFrameSet, Invalid ) {
    std::vector<TestFrameSetRanges::Case> cases;
    fileseq::Frames empty;

    {
        TestFrameSetRanges::Case t = {"", empty};
        cases.push_back(t);
    }
    {
        TestFrameSetRanges::Case t = {"abc", empty};
        cases.push_back(t);
    }
    {
        TestFrameSetRanges::Case t = {"1-abc", empty};
        cases.push_back(t);
    }
    {
        TestFrameSetRanges::Case t = {"    ", empty};
        cases.push_back(t);
    }

    for (size_t i=0; i < cases.size(); ++i) {
        TestFrameSetRanges::Case t = cases[i];

        fileseq::FrameSet s(t.range);
        EXPECT_FALSE(s.isValid())
            << "#" << i << ": Expected invalid FrameSet for '" << t.range << "'";
    }
}

TEST_F( TestFrameSetRanges, Lookup ) {
    fileseq::Status stat;

    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        fileseq::FrameSet s(t.range, &stat);
        ASSERT_TRUE(stat)
            << "#" << i << ": Failed to parse " << t.range << " : " << stat;

        ASSERT_EQ(t.frames.front(), s.start()) << "Start frame did not match";
        ASSERT_EQ(t.frames.back(), s.end()) << "End frame did not match";

        for (size_t idx=0; idx < t.frames.size(); ++idx) {
            fileseq::Frame frame = t.frames[idx];

            ASSERT_TRUE(s.hasFrame(frame))
                << "Expected frame " << frame
                << " to be in frame range " << t.range;

            ASSERT_EQ(idx, s.index(frame))
                << "Index of frame " << frame
                << " did not match in range " << t.range;

            fileseq::Frame actualFrame = s.frame(idx, &stat);
            ASSERT_TRUE(stat) << stat;
            ASSERT_EQ(frame, actualFrame);
        }
    }
};


class TestFrameSetPadding : public testing::Test {

public:
    struct Case {
        std::string input;
        int pad;
        std::string output;
    };

protected:
    void SetUp() {
        {Case t = {"1", -1, "1"}; m_cases.push_back(t);}
        {Case t = {"1", 0, "1"}; m_cases.push_back(t);}
        {Case t = {"1", 1, "1"}; m_cases.push_back(t);}
        {Case t = {"1", 4, "0001"}; m_cases.push_back(t);}
        {Case t = {"-1", 4, "-001"}; m_cases.push_back(t);}
        {Case t = {"-123", 4, "-123"}; m_cases.push_back(t);}
        {Case t = {"-123", 5, "-0123"}; m_cases.push_back(t);}
        {Case t = {"1-100", 4, "0001-0100"}; m_cases.push_back(t);}
        {Case t = {"-10-100", 4, "-010-0100"}; m_cases.push_back(t);}
        {Case t = {"-10--5", 4, "-010--005"}; m_cases.push_back(t);}
        {Case t = {"1,20,300,4000", 4, "0001,0020,0300,4000"}; m_cases.push_back(t);}
        {Case t = {"-10,-5,1,5,20-40,50-300x5", 4, "-010,-005,0001,0005,0020-0040,0050-0300x5"}; m_cases.push_back(t);}
    }

    std::vector<Case> m_cases;
};


TEST_F( TestFrameSetPadding, Padding ) {
    fileseq::Status stat;
    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        fileseq::FrameSet s(t.input, &stat);
        ASSERT_TRUE(stat)
            << "#" << i << ": Failed to parse " << t.input << " : " << stat;

        EXPECT_EQ(t.output, s.frameRange(t.pad))
            << "Given pad " << t.pad << ", and input " << t.input;
    }
};


class TestFrameSetNormalize : public testing::Test {

public:
    struct Case {
        std::string input;
        std::string output;
    };

protected:
    void SetUp() {
        {Case t = {"1,2,3,4,5", "1-5"}; m_cases.push_back(t);}
        {Case t = {"5,1,4,2,3", "1-5"}; m_cases.push_back(t);}
        {Case t = {"1-5,6,7,8,9,10,11-20", "1-20"}; m_cases.push_back(t);}
    }

    std::vector<Case> m_cases;
};

TEST_F( TestFrameSetNormalize, Normalize ) {
    fileseq::Status stat;
    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        fileseq::FrameSet s(t.input, &stat);
        ASSERT_TRUE(stat) << "#" << i << ": Failed to parse " << t.input << " : " << stat;

        EXPECT_EQ(t.output, s.normalized().frameRange()) << "Given range: " << t.input;
    }
};


class TestFrameSetInverted : public testing::Test {

public:
    struct Case {
        std::string range;
        std::string expected;
        int zfill;
    };

protected:
    void SetUp() {
        {Case t = {"1-5,7-10", "6", 0}; m_cases.push_back(t);}
        {Case t = {"1-5,7-10", "0006", 4}; m_cases.push_back(t);}
        {Case t = {"-10--5,5-10", "-4-4", 0}; m_cases.push_back(t);}
        {Case t = {"-10--5,5-10", "-004-0004", 4}; m_cases.push_back(t);}
        {Case t = {"1-10x2", "2-8x2", 0}; m_cases.push_back(t);}
        {Case t = {"1-10x2", "002-008x2", 3}; m_cases.push_back(t);}
        {Case t = {"1-10y2", "3-9x2", 0}; m_cases.push_back(t);}
        {Case t = {"1-10y2", "00003-00009x2", 5}; m_cases.push_back(t);}
        {Case t = {"1-10:4", "", 0}; m_cases.push_back(t);}
        {
            Case t = {" 1 - 10x2, 50, 60 - 62, 70 - 74x2 ", "2-10x2,11-49,51-59,63-69,71-73x2", 0};
            m_cases.push_back(t);
        }
    }

    std::vector<Case> m_cases;
};


TEST_F( TestFrameSetInverted, Inverted ) {
    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        fileseq::FrameSet s(t.range);

        EXPECT_EQ(t.expected, s.invertedFrameRange(t.zfill))
            << "Given invertedFrameRange(pad=" << t.zfill << ") on range " << t.range;
    }
};

