
// Generated from fileseq.g4 by ANTLR 4.13.1

#pragma once


#include "antlr4-runtime.h"


namespace fileseq {


class  fileseqParser : public antlr4::Parser {
public:
  enum {
    UDIM_ANGLE = 1, UDIM_PAREN = 2, PRINTF_PAD = 3, HOUDINI_PAD = 4, HASH = 5, 
    AT = 6, EXTENSION = 7, DOT_FRAME_RANGE = 8, FRAME_RANGE = 9, DOT_NUM = 10, 
    SLASH = 11, SPECIAL_CHAR = 12, NUM = 13, WORD = 14, DASH = 15, WS = 16, 
    OTHER_CHAR = 17
  };

  enum {
    RuleInput = 0, RuleSequence = 1, RulePatternOnly = 2, RuleSingleFrame = 3, 
    RulePlainFile = 4, RuleDirectory = 5, RuleDirSegment = 6, RuleSequenceBasename = 7, 
    RulePatternBasename = 8, RuleSingleFrameBasename = 9, RulePlainBasename = 10, 
    RuleFrameRange = 11, RuleFrameNum = 12, RulePadding = 13, RuleExtension = 14
  };

  explicit fileseqParser(antlr4::TokenStream *input);

  fileseqParser(antlr4::TokenStream *input, const antlr4::atn::ParserATNSimulatorOptions &options);

  ~fileseqParser() override;

  std::string getGrammarFileName() const override;

  const antlr4::atn::ATN& getATN() const override;

  const std::vector<std::string>& getRuleNames() const override;

  const antlr4::dfa::Vocabulary& getVocabulary() const override;

  antlr4::atn::SerializedATNView getSerializedATN() const override;


  class InputContext;
  class SequenceContext;
  class PatternOnlyContext;
  class SingleFrameContext;
  class PlainFileContext;
  class DirectoryContext;
  class DirSegmentContext;
  class SequenceBasenameContext;
  class PatternBasenameContext;
  class SingleFrameBasenameContext;
  class PlainBasenameContext;
  class FrameRangeContext;
  class FrameNumContext;
  class PaddingContext;
  class ExtensionContext; 

  class  InputContext : public antlr4::ParserRuleContext {
  public:
    InputContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    SequenceContext *sequence();
    antlr4::tree::TerminalNode *EOF();
    PatternOnlyContext *patternOnly();
    SingleFrameContext *singleFrame();
    PlainFileContext *plainFile();


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  InputContext* input();

  class  SequenceContext : public antlr4::ParserRuleContext {
  public:
    SequenceContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    DirectoryContext *directory();
    FrameRangeContext *frameRange();
    PaddingContext *padding();
    SequenceBasenameContext *sequenceBasename();
    std::vector<ExtensionContext *> extension();
    ExtensionContext* extension(size_t i);


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  SequenceContext* sequence();

  class  PatternOnlyContext : public antlr4::ParserRuleContext {
  public:
    PatternOnlyContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    DirectoryContext *directory();
    PaddingContext *padding();
    PatternBasenameContext *patternBasename();
    std::vector<ExtensionContext *> extension();
    ExtensionContext* extension(size_t i);


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  PatternOnlyContext* patternOnly();

  class  SingleFrameContext : public antlr4::ParserRuleContext {
  public:
    SingleFrameContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    DirectoryContext *directory();
    SingleFrameBasenameContext *singleFrameBasename();
    FrameNumContext *frameNum();
    std::vector<ExtensionContext *> extension();
    ExtensionContext* extension(size_t i);


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  SingleFrameContext* singleFrame();

  class  PlainFileContext : public antlr4::ParserRuleContext {
  public:
    PlainFileContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    DirectoryContext *directory();
    PlainBasenameContext *plainBasename();
    std::vector<ExtensionContext *> extension();
    ExtensionContext* extension(size_t i);


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  PlainFileContext* plainFile();

  class  DirectoryContext : public antlr4::ParserRuleContext {
  public:
    DirectoryContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    std::vector<antlr4::tree::TerminalNode *> SLASH();
    antlr4::tree::TerminalNode* SLASH(size_t i);
    std::vector<DirSegmentContext *> dirSegment();
    DirSegmentContext* dirSegment(size_t i);


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  DirectoryContext* directory();

  class  DirSegmentContext : public antlr4::ParserRuleContext {
  public:
    DirSegmentContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    std::vector<antlr4::tree::TerminalNode *> WORD();
    antlr4::tree::TerminalNode* WORD(size_t i);
    std::vector<antlr4::tree::TerminalNode *> NUM();
    antlr4::tree::TerminalNode* NUM(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DASH();
    antlr4::tree::TerminalNode* DASH(size_t i);
    std::vector<antlr4::tree::TerminalNode *> SPECIAL_CHAR();
    antlr4::tree::TerminalNode* SPECIAL_CHAR(size_t i);
    std::vector<antlr4::tree::TerminalNode *> FRAME_RANGE();
    antlr4::tree::TerminalNode* FRAME_RANGE(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DOT_FRAME_RANGE();
    antlr4::tree::TerminalNode* DOT_FRAME_RANGE(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DOT_NUM();
    antlr4::tree::TerminalNode* DOT_NUM(size_t i);
    std::vector<antlr4::tree::TerminalNode *> WS();
    antlr4::tree::TerminalNode* WS(size_t i);
    std::vector<antlr4::tree::TerminalNode *> OTHER_CHAR();
    antlr4::tree::TerminalNode* OTHER_CHAR(size_t i);


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  DirSegmentContext* dirSegment();

  class  SequenceBasenameContext : public antlr4::ParserRuleContext {
  public:
    SequenceBasenameContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    std::vector<antlr4::tree::TerminalNode *> WORD();
    antlr4::tree::TerminalNode* WORD(size_t i);
    std::vector<antlr4::tree::TerminalNode *> NUM();
    antlr4::tree::TerminalNode* NUM(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DOT_NUM();
    antlr4::tree::TerminalNode* DOT_NUM(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DASH();
    antlr4::tree::TerminalNode* DASH(size_t i);
    std::vector<antlr4::tree::TerminalNode *> SPECIAL_CHAR();
    antlr4::tree::TerminalNode* SPECIAL_CHAR(size_t i);
    std::vector<antlr4::tree::TerminalNode *> EXTENSION();
    antlr4::tree::TerminalNode* EXTENSION(size_t i);
    std::vector<antlr4::tree::TerminalNode *> FRAME_RANGE();
    antlr4::tree::TerminalNode* FRAME_RANGE(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DOT_FRAME_RANGE();
    antlr4::tree::TerminalNode* DOT_FRAME_RANGE(size_t i);
    std::vector<antlr4::tree::TerminalNode *> WS();
    antlr4::tree::TerminalNode* WS(size_t i);
    std::vector<antlr4::tree::TerminalNode *> OTHER_CHAR();
    antlr4::tree::TerminalNode* OTHER_CHAR(size_t i);


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  SequenceBasenameContext* sequenceBasename();

  class  PatternBasenameContext : public antlr4::ParserRuleContext {
  public:
    PatternBasenameContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    std::vector<antlr4::tree::TerminalNode *> WORD();
    antlr4::tree::TerminalNode* WORD(size_t i);
    std::vector<antlr4::tree::TerminalNode *> NUM();
    antlr4::tree::TerminalNode* NUM(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DOT_NUM();
    antlr4::tree::TerminalNode* DOT_NUM(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DASH();
    antlr4::tree::TerminalNode* DASH(size_t i);
    std::vector<antlr4::tree::TerminalNode *> SPECIAL_CHAR();
    antlr4::tree::TerminalNode* SPECIAL_CHAR(size_t i);
    std::vector<antlr4::tree::TerminalNode *> EXTENSION();
    antlr4::tree::TerminalNode* EXTENSION(size_t i);
    std::vector<antlr4::tree::TerminalNode *> FRAME_RANGE();
    antlr4::tree::TerminalNode* FRAME_RANGE(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DOT_FRAME_RANGE();
    antlr4::tree::TerminalNode* DOT_FRAME_RANGE(size_t i);
    std::vector<antlr4::tree::TerminalNode *> WS();
    antlr4::tree::TerminalNode* WS(size_t i);
    std::vector<antlr4::tree::TerminalNode *> OTHER_CHAR();
    antlr4::tree::TerminalNode* OTHER_CHAR(size_t i);


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  PatternBasenameContext* patternBasename();

  class  SingleFrameBasenameContext : public antlr4::ParserRuleContext {
  public:
    SingleFrameBasenameContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    std::vector<antlr4::tree::TerminalNode *> WORD();
    antlr4::tree::TerminalNode* WORD(size_t i);
    std::vector<antlr4::tree::TerminalNode *> NUM();
    antlr4::tree::TerminalNode* NUM(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DOT_NUM();
    antlr4::tree::TerminalNode* DOT_NUM(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DASH();
    antlr4::tree::TerminalNode* DASH(size_t i);
    std::vector<antlr4::tree::TerminalNode *> SPECIAL_CHAR();
    antlr4::tree::TerminalNode* SPECIAL_CHAR(size_t i);
    std::vector<antlr4::tree::TerminalNode *> EXTENSION();
    antlr4::tree::TerminalNode* EXTENSION(size_t i);
    std::vector<antlr4::tree::TerminalNode *> FRAME_RANGE();
    antlr4::tree::TerminalNode* FRAME_RANGE(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DOT_FRAME_RANGE();
    antlr4::tree::TerminalNode* DOT_FRAME_RANGE(size_t i);
    std::vector<antlr4::tree::TerminalNode *> WS();
    antlr4::tree::TerminalNode* WS(size_t i);
    std::vector<antlr4::tree::TerminalNode *> OTHER_CHAR();
    antlr4::tree::TerminalNode* OTHER_CHAR(size_t i);


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  SingleFrameBasenameContext* singleFrameBasename();

  class  PlainBasenameContext : public antlr4::ParserRuleContext {
  public:
    PlainBasenameContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    std::vector<antlr4::tree::TerminalNode *> WORD();
    antlr4::tree::TerminalNode* WORD(size_t i);
    std::vector<antlr4::tree::TerminalNode *> NUM();
    antlr4::tree::TerminalNode* NUM(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DASH();
    antlr4::tree::TerminalNode* DASH(size_t i);
    std::vector<antlr4::tree::TerminalNode *> SPECIAL_CHAR();
    antlr4::tree::TerminalNode* SPECIAL_CHAR(size_t i);
    std::vector<antlr4::tree::TerminalNode *> FRAME_RANGE();
    antlr4::tree::TerminalNode* FRAME_RANGE(size_t i);
    std::vector<antlr4::tree::TerminalNode *> DOT_FRAME_RANGE();
    antlr4::tree::TerminalNode* DOT_FRAME_RANGE(size_t i);
    std::vector<antlr4::tree::TerminalNode *> WS();
    antlr4::tree::TerminalNode* WS(size_t i);
    std::vector<antlr4::tree::TerminalNode *> OTHER_CHAR();
    antlr4::tree::TerminalNode* OTHER_CHAR(size_t i);


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  PlainBasenameContext* plainBasename();

  class  FrameRangeContext : public antlr4::ParserRuleContext {
  public:
    FrameRangeContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *DOT_FRAME_RANGE();
    antlr4::tree::TerminalNode *FRAME_RANGE();
    antlr4::tree::TerminalNode *DOT_NUM();
    antlr4::tree::TerminalNode *NUM();


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  FrameRangeContext* frameRange();

  class  FrameNumContext : public antlr4::ParserRuleContext {
  public:
    FrameNumContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *DOT_NUM();


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  FrameNumContext* frameNum();

  class  PaddingContext : public antlr4::ParserRuleContext {
  public:
    PaddingContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *UDIM_ANGLE();
    antlr4::tree::TerminalNode *UDIM_PAREN();
    antlr4::tree::TerminalNode *PRINTF_PAD();
    antlr4::tree::TerminalNode *HOUDINI_PAD();
    std::vector<antlr4::tree::TerminalNode *> HASH();
    antlr4::tree::TerminalNode* HASH(size_t i);
    std::vector<antlr4::tree::TerminalNode *> AT();
    antlr4::tree::TerminalNode* AT(size_t i);


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  PaddingContext* padding();

  class  ExtensionContext : public antlr4::ParserRuleContext {
  public:
    ExtensionContext(antlr4::ParserRuleContext *parent, size_t invokingState);
    virtual size_t getRuleIndex() const override;
    antlr4::tree::TerminalNode *EXTENSION();
    antlr4::tree::TerminalNode *DASH();
    antlr4::tree::TerminalNode *NUM();
    antlr4::tree::TerminalNode *DOT_NUM();
    antlr4::tree::TerminalNode *WORD();


    virtual std::any accept(antlr4::tree::ParseTreeVisitor *visitor) override;
   
  };

  ExtensionContext* extension();


  // By default the static state used to implement the parser is lazily initialized during the first
  // call to the constructor. You can call this function if you wish to initialize the static state
  // ahead of time.
  static void initialize();

private:
};

}  // namespace fileseq
