#include "parse.h"
#include "parse_postprocess.h"
#include "fileseqLexer.h"
#include "fileseqParser.h"
#include "fileseqBaseVisitor.h"

#include "antlr4-runtime.h"
#include "atn/ParserATNSimulator.h"
#include "atn/PredictionMode.h"

#include <sstream>

namespace fileseq {
namespace internal {

// Custom visitor that extracts components from the parse tree
class FileSequenceVisitor : public fileseqBaseVisitor {
public:
    ParseResult result;
    std::string error;

    // Visit sequence: dir + basename? + frameRange + padding + extension*
    virtual antlrcpp::Any visitSequence(fileseqParser::SequenceContext *ctx) override {
        result.isSequence = true;

        // Extract directory
        if (ctx->directory()) {
            result.directory = ctx->directory()->getText();
        }

        // Extract basename (optional)
        if (ctx->sequenceBasename()) {
            result.basename = ctx->sequenceBasename()->getText();
        }

        // Check for unsupported subframe syntax (Python-only feature)
        // Subframes use: file.1-5#.10-20@@.exr (dual range) or file.1-5@.#.exr (composite padding)
        if (ctx->frameRange().size() > 1 || ctx->padding().size() > 1) {
            error = "subframe syntax not supported: " + ctx->getText();
            return nullptr;
        }

        // Extract frame range
        if (ctx->frameRange(0)) {
            result.frameRange = ctx->frameRange(0)->getText();
        }

        // Extract padding
        if (ctx->padding(0)) {
            result.padding = ctx->padding(0)->getText();
        }

        // Extract extension (can be multiple parts)
        if (!ctx->extension().empty()) {
            std::ostringstream ext;
            for (auto extCtx : ctx->extension()) {
                ext << extCtx->getText();
            }
            result.extension = ext.str();
        }

        return nullptr;
    }

    // Visit patternOnly: dir + basename? + padding + extension*
    virtual antlrcpp::Any visitPatternOnly(fileseqParser::PatternOnlyContext *ctx) override {
        result.isSequence = true;

        // Extract directory
        if (ctx->directory()) {
            result.directory = ctx->directory()->getText();
        }

        // Extract basename (optional)
        if (ctx->patternBasename()) {
            result.basename = ctx->patternBasename()->getText();
        }

        // Check for unsupported subframe syntax (Python-only feature)
        // Pattern-only subframes use: file.#.@@.exr (composite padding)
        if (ctx->padding().size() > 1) {
            error = "subframe syntax not supported: " + ctx->getText();
            return nullptr;
        }

        // Extract padding
        if (ctx->padding(0)) {
            result.padding = ctx->padding(0)->getText();
        }

        // Extract extension (can be multiple parts)
        if (!ctx->extension().empty()) {
            std::ostringstream ext;
            for (auto extCtx : ctx->extension()) {
                ext << extCtx->getText();
            }
            result.extension = ext.str();
        }

        return nullptr;
    }

    // Visit singleFrame: dir + basename + frameNum + extension*
    virtual antlrcpp::Any visitSingleFrame(fileseqParser::SingleFrameContext *ctx) override {
        result.isSingleFile = true;

        // Extract directory
        if (ctx->directory()) {
            result.directory = ctx->directory()->getText();
        }

        // Extract basename
        if (ctx->singleFrameBasename()) {
            result.basename = ctx->singleFrameBasename()->getText();
        }

        // Extract frame number
        if (ctx->frameNum()) {
            result.frameRange = ctx->frameNum()->getText();
        }

        // Extract extension (can be multiple parts)
        if (!ctx->extension().empty()) {
            std::ostringstream ext;
            for (auto extCtx : ctx->extension()) {
                ext << extCtx->getText();
            }
            result.extension = ext.str();
        }

        return nullptr;
    }

    // Visit plainFile: dir + basename? + extension*
    virtual antlrcpp::Any visitPlainFile(fileseqParser::PlainFileContext *ctx) override {
        result.isPlainFile = true;

        // Extract directory
        if (ctx->directory()) {
            result.directory = ctx->directory()->getText();
        }

        // Extract basename (optional)
        if (ctx->plainBasename()) {
            result.basename = ctx->plainBasename()->getText();
        }

        // Extract extension (can be multiple parts)
        if (!ctx->extension().empty()) {
            std::ostringstream ext;
            for (auto extCtx : ctx->extension()) {
                ext << extCtx->getText();
            }
            result.extension = ext.str();
        }

        return nullptr;
    }
};


bool parseFileSequence(ParseResult& result, const std::string& path) {
    try {
        // Create input stream from string
        antlr4::ANTLRInputStream input(path);

        // Create lexer
        fileseqLexer lexer(&input);

        // Remove default error listeners from lexer
        lexer.removeErrorListeners();

        // Create token stream
        antlr4::CommonTokenStream tokens(&lexer);

        // Create parser
        fileseqParser parser(&tokens);

        // Remove default error listeners from parser
        parser.removeErrorListeners();

        // Try SLL prediction mode first (30x faster for unambiguous grammars)
        parser.getInterpreter<antlr4::atn::ParserATNSimulator>()->setPredictionMode(
            antlr4::atn::PredictionMode::SLL);

        // Parse the input
        fileseqParser::InputContext* tree = parser.input();

        // If SLL failed, retry with LL prediction mode (slower but more thorough)
        if (lexer.getNumberOfSyntaxErrors() > 0 || parser.getNumberOfSyntaxErrors() > 0) {
            // Reset for retry
            input.seek(0);
            lexer.setInputStream(&input);
            tokens.setTokenSource(&lexer);
            parser.setInputStream(&tokens);

            parser.getInterpreter<antlr4::atn::ParserATNSimulator>()->setPredictionMode(
                antlr4::atn::PredictionMode::LL);

            tree = parser.input();

            // Check for parse errors after LL retry
            if (lexer.getNumberOfSyntaxErrors() > 0 || parser.getNumberOfSyntaxErrors() > 0) {
                return false;
            }
        }

        // Create visitor and walk the tree
        FileSequenceVisitor visitor;
        visitor.visit(tree);

        // Check for visitor errors (e.g., unsupported subframe syntax)
        if (!visitor.error.empty()) {
            return false;
        }

        // Get the result
        result = visitor.result;

        // Apply post-processing for the 3 special cases
        postProcessParseResult(result);

        return true;

    } catch (...) {
        // Any exception during parsing is a failure
        return false;
    }
}

} // namespace internal
} // namespace fileseq
