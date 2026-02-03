#include "parse.h"
#include "parse_postprocess.h"
#include "fileseqLexer.h"
#include "fileseqParser.h"
#include "fileseqBaseVisitor.h"

#include "antlr4-runtime.h"

#include <sstream>

namespace fileseq {
namespace internal {

// Custom visitor that extracts components from the parse tree
class FileSequenceVisitor : public fileseqBaseVisitor {
public:
    ParseResult result;

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

        // Extract frame range
        if (ctx->frameRange()) {
            result.frameRange = ctx->frameRange()->getText();
        }

        // Extract padding
        if (ctx->padding()) {
            result.padding = ctx->padding()->getText();
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

        // Extract padding
        if (ctx->padding()) {
            result.padding = ctx->padding()->getText();
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

        // Parse the input
        fileseqParser::InputContext* tree = parser.input();

        // Check for parse errors (from both lexer and parser)
        if (lexer.getNumberOfSyntaxErrors() > 0 || parser.getNumberOfSyntaxErrors() > 0) {
            return false;
        }

        // Create visitor and walk the tree
        FileSequenceVisitor visitor;
        visitor.visit(tree);

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
