#include "common.h"
#include <gtest/gtest.h>

#include "../fileseq.h"


class TestNewFileSequences : public testing::Test {

public:
    struct Case {
        std::string path;
        std::string outPath;
        fileseq::Frame start;
        fileseq::Frame end;
        int zfill;
        size_t frameCount;
    };

protected:
    void SetUp() {
        {
            Case t = {"/file_path.100.exr", "/file_path.100@@@.exr", 100, 100, 3, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/file_path.0100.exr", "/file_path.0100#.exr", 100, 100, 4, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/f.1-100#.jpeg", "/dir/f.1-100#.jpeg", 1, 100, 4, 100};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/f.1-100@@@.f", "/dir/f.1-100@@@.f", 1, 100, 3, 100};
            m_cases.push_back(t);
        }
        {
            Case t = {
                "/dir/f.1-10,50,60-90x2##.mp4",
                "/dir/f.1-10,50,60-90x2##.mp4",
                1, 90, 8, 27
            };
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/f.exr", "/dir/f.exr", 0, 0, 0, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/f.100", "/dir/f.100@@@", 100, 100, 3, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/f.@@.ext", "/dir/f.@@.ext", 0, 0, 2, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/f100.ext", "/dir/f100@@@.ext", 100, 100, 3, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/f_100.ext", "/dir/f_100@@@.ext", 100, 100, 3, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/no_frames.ext", "/dir/no_frames.ext", 0, 0, 0, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/no_file_extension", "/dir/no_file_extension", 0, 0, 0, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/.hidden", "/dir/.hidden", 0, 0, 0, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/.hidden.100", "/dir/.hidden.100@@@", 100, 100, 3, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/.hidden.100.ext", "/dir/.hidden.100@@@.ext", 100, 100, 3, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {"/dir/.hidden5.1-10#.7zip", "/dir/.hidden5.1-10#.7zip", 1, 10, 4, 10};
            m_cases.push_back(t);
        }
        {
            Case t = {".10000000000", ".10000000000", 0, 0, 0, 1};
            m_cases.push_back(t);
        }
        {
            Case t = {
                ".10000000000.123", ".10000000000@@@@@@@@@@@.123",
                10000000000, 10000000000, 11, 1
            };
            m_cases.push_back(t);
        }
    }

    std::vector<Case> m_cases;
};


TEST_F( TestNewFileSequences, New ) {
    fileseq::Status stat;

    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        fileseq::FileSequence s(t.path, &stat);
        ASSERT_TRUE(stat) << "#" << i << ": Failed to parse " << t.path << " : " << stat;

        EXPECT_EQ(t.outPath, s.string());
        EXPECT_EQ(t.start, s.start());
        EXPECT_EQ(t.end, s.end());
        EXPECT_EQ(t.zfill, s.zfill());
        EXPECT_EQ(t.frameCount, s.length());
    }
}


class TestNewFileSequenceFrameIndex : public testing::Test {

public:
    struct Case {
        std::string path;
        size_t size;
        size_t idx;
        fileseq::Frame frame;
        std::string expected;
    };

protected:
    void SetUp() {
        {
            Case t = {"file.100.exr", 1, 0, 100, "file.100.exr"};
            m_cases.push_back(t);
        }
        {
            Case t = {"file.1-100#.exr", 100, 4, 5, "file.0005.exr"};
            m_cases.push_back(t);
        }
        {
            Case t = {"file.-10-30@@@.exr", 40, 25, 15, "file.015.exr"};
            m_cases.push_back(t);
        }
        {
            Case t = {"file.-10-30@@@.exr", 40, 5, -5, "file.-05.exr"};
            m_cases.push_back(t);
        }
        {
            Case t = {"file.1,3-12x3,20-24x2,28,32##.exr", 10, 8, 28, "file.00000028.exr"};
            m_cases.push_back(t);
        }
    }

    std::vector<Case> m_cases;
};


TEST_F( TestNewFileSequenceFrameIndex, FrameIndex ) {
    fileseq::Status stat;

    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        fileseq::FileSequence s(t.path, &stat);
        ASSERT_TRUE(stat) << "#" << i << ": Failed to parse " << t.path << " : " << stat;

        EXPECT_EQ(t.expected, s.frame(t.frame)) << "Given frame: " << t.frame;
        EXPECT_EQ(t.expected, s.index(t.idx)) << "Given index: " << t.idx;

        for (size_t i=0; i < t.size; ++i) {
            EXPECT_NE("", s.index(i)) << "Got empty path for index: " << i;
        }
    }
}


class TestNewFileSequenceSetDir : public testing::Test {

public:
    struct Case {
        std::string src;
        std::string newDir;
        std::string newBase;
        std::string newExt;
        std::string newRange;
        std::string newPad;
        std::string expected;
    };

protected:
    void SetUp() {
        {
            Case t = {
                "/path/to/file.100.exr",
                "/other", "fileB.", "jpg", "200", "#",
                "/other/fileB.200#.jpg"
            };
            m_cases.push_back(t);
        }
        {
            Case t = {
                "/path/to/file.1-100#.exr",
                "/other/subdir", "fileB", "f", "-10-5,20-30x2", "@@@",
                "/other/subdir/fileB-10-5,20-30x2@@@.f"
            };
            m_cases.push_back(t);
        }
    }

    std::vector<Case> m_cases;
};


TEST_F( TestNewFileSequenceSetDir, SetDir ) {
    fileseq::Status stat;

    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        fileseq::FileSequence s(t.src, &stat);
        ASSERT_TRUE(stat) << "#" << i << ": Failed to parse " << t.src << " : " << stat;

        s.setDirname(t.newDir);
        s.setBasename(t.newBase);
        s.setExt(t.newExt);
        s.setPadding(t.newPad);
        s.setFrameRange(t.newRange);

        EXPECT_EQ(t.expected, s.string());
    }
}


class TestNewFileSequenceFormat : public testing::Test {

public:
    struct Case {
        std::string format;
        std::string expected;
    };

protected:
    void SetUp() {
        fileseq::Status stat;
        m_seq = fileseq::FileSequence("/a/path/to/the/file_foo.1-10,50,70-100x5#.ext", &stat);
        ASSERT_TRUE(stat) << stat;

        {
            Case t = {
                "{{dir}}{{base}}{{frange}}{{pad}}{{ext}}",
                m_seq.string(),
            };
            m_cases.push_back(t);
        }
        {
            Case t = {
                "{{startf}} {{endf}} {{len}} {{zfill}}",
                "1 100 18 4",
            };
            m_cases.push_back(t);
        }
        {
            Case t = {
                "{{base}}{{if inverted}}{{inverted}}{{else}}{{frange}}{{end}}{{ext}}",
                "file_foo.11-49,51-69,71-74,76-79,81-84,86-89,91-94,96-99.ext",
            };
            m_cases.push_back(t);
        }
    }

    fileseq::FileSequence m_seq;
    std::vector<Case> m_cases;
};


TEST_F( TestNewFileSequenceFormat, Format ) {
    fileseq::Status stat;
    for (size_t i=0; i < m_cases.size(); ++i) {
        Case t = m_cases[i];

        std::string actual = m_seq.format(t.format, &stat);
        ASSERT_TRUE(stat) << "Failed to format using: " << t.format;
        EXPECT_EQ(t.expected, actual);
    }
}

