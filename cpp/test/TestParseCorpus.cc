#include <gtest/gtest.h>

#include "private/parser/parse.h"

#include <fstream>
#include <string>
#include <vector>

#define xstr(a) str(a)
#define str(a) #a

namespace {

// Shared with the Python fileseq repository and must be kept identical there,
// alongside fileseq.g4. The path is absolute, so the tests can run from any
// directory.
const char* const kCorpusFile = xstr(FILESEQ_CORPUS_FILE);

struct CorpusCase {
    int line;
    std::vector<std::string> cols;
};

std::vector<std::string> splitTabs(const std::string& line) {
    std::vector<std::string> cols;
    size_t start = 0;
    while (true) {
        size_t end = line.find('\t', start);
        cols.push_back(line.substr(start, end - start));
        if (end == std::string::npos) {
            return cols;
        }
        start = end + 1;
    }
}

bool forPort(const std::string& ports, const std::string& port) {
    if (ports == "all") {
        return true;
    }
    std::string padded = "," + ports + ",";
    return padded.find("," + port + ",") != std::string::npos;
}

std::vector<CorpusCase> loadCorpus() {
    std::vector<CorpusCase> cases;
    std::ifstream f(kCorpusFile);
    std::string line;
    for (int n = 1; std::getline(f, line); ++n) {
        if (line.empty() || line[0] == '#') {
            continue;
        }
        CorpusCase c = {n, splitTabs(line)};
        if (forPort(c.cols[0], "cpp")) {
            cases.push_back(c);
        }
    }
    return cases;
}

} // namespace


TEST( TestParseCorpus, Corpus ) {
    std::vector<CorpusCase> cases = loadCorpus();
    ASSERT_FALSE(cases.empty()) << "could not read " << kCorpusFile;

    for (size_t i = 0; i < cases.size(); ++i) {
        const CorpusCase& c = cases[i];
        ASSERT_TRUE(c.cols.size() == 3 || c.cols.size() == 8)
            << kCorpusFile << ":" << c.line << ": want 3 or 8 columns";

        const std::string& kind = c.cols[1];
        const std::string& input = c.cols[2];
        fileseq::internal::ParseResult r;
        bool ok = fileseq::internal::parseFileSequence(r, input);

        if (kind == "error") {
            EXPECT_FALSE(ok) << "line " << c.line << ": expected an error for " << input;
            continue;
        }
        if (!ok) {
            ADD_FAILURE() << "line " << c.line << ": failed to parse " << input;
            continue;
        }
        std::string gotKind = r.isSequence ? "seq" : r.isSingleFile ? "single" : "plain";
        EXPECT_EQ(kind, gotKind) << "line " << c.line << ": " << input;
        EXPECT_EQ(c.cols[3], r.directory) << "line " << c.line << ": " << input;
        EXPECT_EQ(c.cols[4], r.basename) << "line " << c.line << ": " << input;
        EXPECT_EQ(c.cols[5], r.frameRange) << "line " << c.line << ": " << input;
        EXPECT_EQ(c.cols[6], r.padding) << "line " << c.line << ": " << input;
        EXPECT_EQ(c.cols[7], r.extension) << "line " << c.line << ": " << input;
    }
}
