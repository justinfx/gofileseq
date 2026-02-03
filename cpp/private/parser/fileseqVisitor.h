
// Generated from fileseq.g4 by ANTLR 4.13.1

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
    virtual std::any visitInput(fileseqParser::InputContext *context) = 0;

    virtual std::any visitSequence(fileseqParser::SequenceContext *context) = 0;

    virtual std::any visitPatternOnly(fileseqParser::PatternOnlyContext *context) = 0;

    virtual std::any visitSingleFrame(fileseqParser::SingleFrameContext *context) = 0;

    virtual std::any visitPlainFile(fileseqParser::PlainFileContext *context) = 0;

    virtual std::any visitDirectory(fileseqParser::DirectoryContext *context) = 0;

    virtual std::any visitDirSegment(fileseqParser::DirSegmentContext *context) = 0;

    virtual std::any visitSequenceBasename(fileseqParser::SequenceBasenameContext *context) = 0;

    virtual std::any visitPatternBasename(fileseqParser::PatternBasenameContext *context) = 0;

    virtual std::any visitSingleFrameBasename(fileseqParser::SingleFrameBasenameContext *context) = 0;

    virtual std::any visitPlainBasename(fileseqParser::PlainBasenameContext *context) = 0;

    virtual std::any visitFrameRange(fileseqParser::FrameRangeContext *context) = 0;

    virtual std::any visitFrameNum(fileseqParser::FrameNumContext *context) = 0;

    virtual std::any visitPadding(fileseqParser::PaddingContext *context) = 0;

    virtual std::any visitExtension(fileseqParser::ExtensionContext *context) = 0;


};

}  // namespace fileseq
