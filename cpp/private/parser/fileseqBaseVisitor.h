
// Generated from fileseq.g4 by ANTLR 4.9.3

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

  virtual antlrcpp::Any visitInput(fileseqParser::InputContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitSequence(fileseqParser::SequenceContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitPatternOnly(fileseqParser::PatternOnlyContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitSingleFrame(fileseqParser::SingleFrameContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitPlainFile(fileseqParser::PlainFileContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitDirectory(fileseqParser::DirectoryContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitDirSegment(fileseqParser::DirSegmentContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitSequenceBasename(fileseqParser::SequenceBasenameContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitPatternBasename(fileseqParser::PatternBasenameContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitSingleFrameBasename(fileseqParser::SingleFrameBasenameContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitPlainBasename(fileseqParser::PlainBasenameContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitFrameRange(fileseqParser::FrameRangeContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitFrameNum(fileseqParser::FrameNumContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitPadding(fileseqParser::PaddingContext *ctx) override {
    return visitChildren(ctx);
  }

  virtual antlrcpp::Any visitExtension(fileseqParser::ExtensionContext *ctx) override {
    return visitChildren(ctx);
  }


};

}  // namespace fileseq
