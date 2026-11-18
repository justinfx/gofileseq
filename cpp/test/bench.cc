// fileseq_bench reports the time to parse file sequence strings, so parser
// and grammar changes can be compared against a baseline.
//
//   make fileseq_bench && ./fileseq_bench

#include "fileseq.h"
#include "private/parser/parse.h"

#include <chrono>
#include <cstdio>
#include <fstream>
#include <string>
#include <vector>

#define xstr(a) str(a)
#define str(a) #a

namespace {

const int kRepeats = 5;

// Returns the best mean time in microseconds of running fn n times.
template <typename Fn>
double bestMicros(int n, Fn fn) {
    double best = 1e18;
    for (int r = 0; r < kRepeats; ++r) {
        std::chrono::steady_clock::time_point t0 = std::chrono::steady_clock::now();
        for (int i = 0; i < n; ++i) {
            fn();
        }
        std::chrono::duration<double, std::micro> d = std::chrono::steady_clock::now() - t0;
        double us = d.count() / n;
        if (us < best) {
            best = us;
        }
    }
    return best;
}

std::vector<std::string> loadCorpusInputs(const char* path) {
    std::vector<std::string> inputs;
    std::ifstream f(path);
    std::string line;
    while (std::getline(f, line)) {
        if (line.empty() || line[0] == '#') {
            continue;
        }
        size_t a = line.find('\t');
        size_t b = line.find('\t', a + 1);
        size_t c = line.find('\t', b + 1);
        std::string ports = line.substr(0, a);
        if (ports == "all" || ("," + ports + ",").find(",cpp,") != std::string::npos) {
            inputs.push_back(line.substr(b + 1, c == std::string::npos ? std::string::npos : c - b - 1));
        }
    }
    return inputs;
}

} // namespace

int main(int argc, char** argv) {
    const char* names[] = {"sequence", "single", "stepped", "plain"};
    const char* inputs[] = {
        "/mnt/projects/show/shots/sh010/comp/v012/sh010_comp_v012.1001-1240#.exr",
        "/mnt/projects/show/shots/sh010/plates/sh010_plate.1001.dpx",
        "/renders/beauty/beauty.1-100x2,150-200@@@@.tif",
        "/mnt/projects/show/assets/chars/hero/model/hero_model_v003.abc",
    };

    for (int k = 0; k < 4; ++k) {
        const char* input = inputs[k];
        double us = bestMicros(2000, [input]() {
            fileseq::Status stat;
            fileseq::FileSequence s(input, &stat);
        });
        std::printf("FileSequence/%-9s %10.1f us/op\n", names[k], us);
    }

    const char* corpusFile = argc > 1 ? argv[1] : xstr(FILESEQ_CORPUS_FILE);
    std::vector<std::string> corpus = loadCorpusInputs(corpusFile);
    if (corpus.empty()) {
        std::fprintf(stderr, "could not read corpus %s\n", corpusFile);
        return 1;
    }
    double us = bestMicros(20, [&corpus]() {
        for (size_t i = 0; i < corpus.size(); ++i) {
            fileseq::internal::ParseResult r;
            fileseq::internal::parseFileSequence(r, corpus[i]);
        }
    });
    std::printf("ParseCorpus (%zu inputs) %8.1f us/op\n", corpus.size(), us);
    return 0;
}
