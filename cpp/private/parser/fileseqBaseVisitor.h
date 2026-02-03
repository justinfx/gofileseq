
// Generated from fileseq.g4 by ANTLR 4.13.1

#pragma once


#include "antlr4-runtime.h"
#include "fileseqVisitor.h"


namespace fileseq {

/**
 * This class provides an empty implementation of fileseqVisitor, which can be
 * extended to create a visitor which only needs to handle a subset of the available methods.
 */
class  fileseqBaseVisitor : public fileseqVisitor {
public:

  virtual std::any visitInput(fileseqParser::InputContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitSequence(fileseqParser::SequenceContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitPatternOnly(fileseqParser::PatternOnlyContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitSingleFrame(fileseqParser::SingleFrameContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitPlainFile(fileseqParser::PlainFileContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitDirectory(fileseqParser::DirectoryContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitDirSegment(fileseqParser::DirSegmentContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitSequenceBasename(fileseqParser::SequenceBasenameContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitPatternBasename(fileseqParser::PatternBasenameContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitSingleFrameBasename(fileseqParser::SingleFrameBasenameContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitPlainBasename(fileseqParser::PlainBasenameContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitFrameRange(fileseqParser::FrameRangeContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitFrameNum(fileseqParser::FrameNumContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitPadding(fileseqParser::PaddingContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual std::any visitExtension(fileseqParser::ExtensionContext *ctx) override {
    return visitChildren(ctx);
  }


};

}  // namespace fileseq
