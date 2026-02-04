
// Generated from fileseq.g4 by ANTLR 4.9.3

#pragma once


#include "antlr4-runtime.h"
#include "fileseqParser.h"


namespace fileseq {

/**
 * This class defines an abstract visitor for a parse tree
 * produced by fileseqParser.
 */
class  fileseqVisitor : public antlr4::tree::AbstractParseTreeVisitor {
public:

  /**
   * Visit parse trees produced by fileseqParser.
   */
    virtual antlrcpp::Any visitInput(fileseqParser::InputContext *context) = 0;

    virtual antlrcpp::Any visitSequence(fileseqParser::SequenceContext *context) = 0;

    virtual antlrcpp::Any visitPatternOnly(fileseqParser::PatternOnlyContext *context) = 0;

    virtual antlrcpp::Any visitSingleFrame(fileseqParser::SingleFrameContext *context) = 0;

    virtual antlrcpp::Any visitPlainFile(fileseqParser::PlainFileContext *context) = 0;

    virtual antlrcpp::Any visitDirectory(fileseqParser::DirectoryContext *context) = 0;

    virtual antlrcpp::Any visitDirSegment(fileseqParser::DirSegmentContext *context) = 0;

    virtual antlrcpp::Any visitSequenceBasename(fileseqParser::SequenceBasenameContext *context) = 0;

    virtual antlrcpp::Any visitPatternBasename(fileseqParser::PatternBasenameContext *context) = 0;

    virtual antlrcpp::Any visitSingleFrameBasename(fileseqParser::SingleFrameBasenameContext *context) = 0;

    virtual antlrcpp::Any visitPlainBasename(fileseqParser::PlainBasenameContext *context) = 0;

    virtual antlrcpp::Any visitFrameRange(fileseqParser::FrameRangeContext *context) = 0;

    virtual antlrcpp::Any visitFrameNum(fileseqParser::FrameNumContext *context) = 0;

    virtual antlrcpp::Any visitPadding(fileseqParser::PaddingContext *context) = 0;

    virtual antlrcpp::Any visitExtension(fileseqParser::ExtensionContext *context) = 0;


};

}  // namespace fileseq
