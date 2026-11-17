
// Generated from fileseq.g4 by ANTLR 4.9.3


#include "fileseqVisitor.h"

#include "fileseqParser.h"


using namespace antlrcpp;
using namespace fileseq;
using namespace antlr4;

fileseqParser::fileseqParser(TokenStream *input) : Parser(input) {
  _interpreter = new atn::ParserATNSimulator(this, _atn, _decisionToDFA, _sharedContextCache);
}

fileseqParser::~fileseqParser() {
  delete _interpreter;
}

std::string fileseqParser::getGrammarFileName() const {
  return "fileseq.g4";
}

const std::vector<std::string>& fileseqParser::getRuleNames() const {
  return _ruleNames;
}

dfa::Vocabulary& fileseqParser::getVocabulary() const {
  return _vocabulary;
}


//----------------- InputContext ------------------------------------------------------------------

fileseqParser::InputContext::InputContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

fileseqParser::SequenceContext* fileseqParser::InputContext::sequence() {
  return getRuleContext<fileseqParser::SequenceContext>(0);
}

tree::TerminalNode* fileseqParser::InputContext::EOF() {
  return getToken(fileseqParser::EOF, 0);
}

fileseqParser::PatternOnlyContext* fileseqParser::InputContext::patternOnly() {
  return getRuleContext<fileseqParser::PatternOnlyContext>(0);
}

fileseqParser::SingleFrameContext* fileseqParser::InputContext::singleFrame() {
  return getRuleContext<fileseqParser::SingleFrameContext>(0);
}

fileseqParser::PlainFileContext* fileseqParser::InputContext::plainFile() {
  return getRuleContext<fileseqParser::PlainFileContext>(0);
}


size_t fileseqParser::InputContext::getRuleIndex() const {
  return fileseqParser::RuleInput;
}


antlrcpp::Any fileseqParser::InputContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitInput(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::InputContext* fileseqParser::input() {
  InputContext *_localctx = _tracker.createInstance<InputContext>(_ctx, getState());
  enterRule(_localctx, 0, fileseqParser::RuleInput);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(46);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 0, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(34);
      sequence();
      setState(35);
      match(fileseqParser::EOF);
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(37);
      patternOnly();
      setState(38);
      match(fileseqParser::EOF);
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);
      setState(40);
      singleFrame();
      setState(41);
      match(fileseqParser::EOF);
      break;
    }

    case 4: {
      enterOuterAlt(_localctx, 4);
      setState(43);
      plainFile();
      setState(44);
      match(fileseqParser::EOF);
      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- SequenceContext ------------------------------------------------------------------

fileseqParser::SequenceContext::SequenceContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

fileseqParser::DirectoryContext* fileseqParser::SequenceContext::directory() {
  return getRuleContext<fileseqParser::DirectoryContext>(0);
}

std::vector<fileseqParser::FrameRangeContext *> fileseqParser::SequenceContext::frameRange() {
  return getRuleContexts<fileseqParser::FrameRangeContext>();
}

fileseqParser::FrameRangeContext* fileseqParser::SequenceContext::frameRange(size_t i) {
  return getRuleContext<fileseqParser::FrameRangeContext>(i);
}

std::vector<fileseqParser::PaddingContext *> fileseqParser::SequenceContext::padding() {
  return getRuleContexts<fileseqParser::PaddingContext>();
}

fileseqParser::PaddingContext* fileseqParser::SequenceContext::padding(size_t i) {
  return getRuleContext<fileseqParser::PaddingContext>(i);
}

fileseqParser::BasenameContext* fileseqParser::SequenceContext::basename() {
  return getRuleContext<fileseqParser::BasenameContext>(0);
}

tree::TerminalNode* fileseqParser::SequenceContext::SPECIAL_CHAR() {
  return getToken(fileseqParser::SPECIAL_CHAR, 0);
}

std::vector<fileseqParser::ExtensionContext *> fileseqParser::SequenceContext::extension() {
  return getRuleContexts<fileseqParser::ExtensionContext>();
}

fileseqParser::ExtensionContext* fileseqParser::SequenceContext::extension(size_t i) {
  return getRuleContext<fileseqParser::ExtensionContext>(i);
}


size_t fileseqParser::SequenceContext::getRuleIndex() const {
  return fileseqParser::RuleSequence;
}


antlrcpp::Any fileseqParser::SequenceContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitSequence(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::SequenceContext* fileseqParser::sequence() {
  SequenceContext *_localctx = _tracker.createInstance<SequenceContext>(_ctx, getState());
  enterRule(_localctx, 2, fileseqParser::RuleSequence);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(48);
    directory();
    setState(50);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 1, _ctx)) {
    case 1: {
      setState(49);
      basename();
      break;
    }

    default:
      break;
    }
    setState(52);
    frameRange();
    setState(53);
    padding();
    setState(59);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 2, _ctx)) {
    case 1: {
      setState(54);
      frameRange();
      setState(55);
      padding();
      break;
    }

    case 2: {
      setState(57);
      match(fileseqParser::SPECIAL_CHAR);
      setState(58);
      padding();
      break;
    }

    default:
      break;
    }
    setState(64);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << fileseqParser::EXTENSION)
      | (1ULL << fileseqParser::DOT_NUM)
      | (1ULL << fileseqParser::WORD))) != 0)) {
      setState(61);
      extension();
      setState(66);
      _errHandler->sync(this);
      _la = _input->LA(1);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PatternOnlyContext ------------------------------------------------------------------

fileseqParser::PatternOnlyContext::PatternOnlyContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

fileseqParser::DirectoryContext* fileseqParser::PatternOnlyContext::directory() {
  return getRuleContext<fileseqParser::DirectoryContext>(0);
}

std::vector<fileseqParser::PaddingContext *> fileseqParser::PatternOnlyContext::padding() {
  return getRuleContexts<fileseqParser::PaddingContext>();
}

fileseqParser::PaddingContext* fileseqParser::PatternOnlyContext::padding(size_t i) {
  return getRuleContext<fileseqParser::PaddingContext>(i);
}

fileseqParser::PatBasenameContext* fileseqParser::PatternOnlyContext::patBasename() {
  return getRuleContext<fileseqParser::PatBasenameContext>(0);
}

tree::TerminalNode* fileseqParser::PatternOnlyContext::SPECIAL_CHAR() {
  return getToken(fileseqParser::SPECIAL_CHAR, 0);
}

std::vector<fileseqParser::ExtensionContext *> fileseqParser::PatternOnlyContext::extension() {
  return getRuleContexts<fileseqParser::ExtensionContext>();
}

fileseqParser::ExtensionContext* fileseqParser::PatternOnlyContext::extension(size_t i) {
  return getRuleContext<fileseqParser::ExtensionContext>(i);
}


size_t fileseqParser::PatternOnlyContext::getRuleIndex() const {
  return fileseqParser::RulePatternOnly;
}


antlrcpp::Any fileseqParser::PatternOnlyContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitPatternOnly(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::PatternOnlyContext* fileseqParser::patternOnly() {
  PatternOnlyContext *_localctx = _tracker.createInstance<PatternOnlyContext>(_ctx, getState());
  enterRule(_localctx, 4, fileseqParser::RulePatternOnly);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(67);
    directory();
    setState(69);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << fileseqParser::EXTENSION)
      | (1ULL << fileseqParser::DOT_FRAME_RANGE)
      | (1ULL << fileseqParser::FRAME_RANGE)
      | (1ULL << fileseqParser::DOT_NUM)
      | (1ULL << fileseqParser::SPECIAL_CHAR)
      | (1ULL << fileseqParser::NUM)
      | (1ULL << fileseqParser::WORD)
      | (1ULL << fileseqParser::DASH)
      | (1ULL << fileseqParser::WS)
      | (1ULL << fileseqParser::OTHER_CHAR))) != 0)) {
      setState(68);
      patBasename();
    }
    setState(71);
    padding();
    setState(74);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == fileseqParser::SPECIAL_CHAR) {
      setState(72);
      match(fileseqParser::SPECIAL_CHAR);
      setState(73);
      padding();
    }
    setState(79);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << fileseqParser::EXTENSION)
      | (1ULL << fileseqParser::DOT_NUM)
      | (1ULL << fileseqParser::WORD))) != 0)) {
      setState(76);
      extension();
      setState(81);
      _errHandler->sync(this);
      _la = _input->LA(1);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- SingleFrameContext ------------------------------------------------------------------

fileseqParser::SingleFrameContext::SingleFrameContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

fileseqParser::DirectoryContext* fileseqParser::SingleFrameContext::directory() {
  return getRuleContext<fileseqParser::DirectoryContext>(0);
}

fileseqParser::FrameNumContext* fileseqParser::SingleFrameContext::frameNum() {
  return getRuleContext<fileseqParser::FrameNumContext>(0);
}

fileseqParser::FrameBasenameContext* fileseqParser::SingleFrameContext::frameBasename() {
  return getRuleContext<fileseqParser::FrameBasenameContext>(0);
}

std::vector<fileseqParser::ExtensionContext *> fileseqParser::SingleFrameContext::extension() {
  return getRuleContexts<fileseqParser::ExtensionContext>();
}

fileseqParser::ExtensionContext* fileseqParser::SingleFrameContext::extension(size_t i) {
  return getRuleContext<fileseqParser::ExtensionContext>(i);
}


size_t fileseqParser::SingleFrameContext::getRuleIndex() const {
  return fileseqParser::RuleSingleFrame;
}


antlrcpp::Any fileseqParser::SingleFrameContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitSingleFrame(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::SingleFrameContext* fileseqParser::singleFrame() {
  SingleFrameContext *_localctx = _tracker.createInstance<SingleFrameContext>(_ctx, getState());
  enterRule(_localctx, 6, fileseqParser::RuleSingleFrame);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(82);
    directory();
    setState(84);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 7, _ctx)) {
    case 1: {
      setState(83);
      frameBasename();
      break;
    }

    default:
      break;
    }
    setState(86);
    frameNum();
    setState(88); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(87);
      extension();
      setState(90); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << fileseqParser::EXTENSION)
      | (1ULL << fileseqParser::DOT_NUM)
      | (1ULL << fileseqParser::WORD))) != 0));
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PlainFileContext ------------------------------------------------------------------

fileseqParser::PlainFileContext::PlainFileContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

fileseqParser::DirectoryContext* fileseqParser::PlainFileContext::directory() {
  return getRuleContext<fileseqParser::DirectoryContext>(0);
}

fileseqParser::PlainBasenameContext* fileseqParser::PlainFileContext::plainBasename() {
  return getRuleContext<fileseqParser::PlainBasenameContext>(0);
}

std::vector<fileseqParser::ExtensionContext *> fileseqParser::PlainFileContext::extension() {
  return getRuleContexts<fileseqParser::ExtensionContext>();
}

fileseqParser::ExtensionContext* fileseqParser::PlainFileContext::extension(size_t i) {
  return getRuleContext<fileseqParser::ExtensionContext>(i);
}


size_t fileseqParser::PlainFileContext::getRuleIndex() const {
  return fileseqParser::RulePlainFile;
}


antlrcpp::Any fileseqParser::PlainFileContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitPlainFile(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::PlainFileContext* fileseqParser::plainFile() {
  PlainFileContext *_localctx = _tracker.createInstance<PlainFileContext>(_ctx, getState());
  enterRule(_localctx, 8, fileseqParser::RulePlainFile);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(92);
    directory();
    setState(94);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 9, _ctx)) {
    case 1: {
      setState(93);
      plainBasename();
      break;
    }

    default:
      break;
    }
    setState(99);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << fileseqParser::EXTENSION)
      | (1ULL << fileseqParser::DOT_NUM)
      | (1ULL << fileseqParser::WORD))) != 0)) {
      setState(96);
      extension();
      setState(101);
      _errHandler->sync(this);
      _la = _input->LA(1);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- DirectoryContext ------------------------------------------------------------------

fileseqParser::DirectoryContext::DirectoryContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<tree::TerminalNode *> fileseqParser::DirectoryContext::SLASH() {
  return getTokens(fileseqParser::SLASH);
}

tree::TerminalNode* fileseqParser::DirectoryContext::SLASH(size_t i) {
  return getToken(fileseqParser::SLASH, i);
}

std::vector<fileseqParser::DirSegmentContext *> fileseqParser::DirectoryContext::dirSegment() {
  return getRuleContexts<fileseqParser::DirSegmentContext>();
}

fileseqParser::DirSegmentContext* fileseqParser::DirectoryContext::dirSegment(size_t i) {
  return getRuleContext<fileseqParser::DirSegmentContext>(i);
}


size_t fileseqParser::DirectoryContext::getRuleIndex() const {
  return fileseqParser::RuleDirectory;
}


antlrcpp::Any fileseqParser::DirectoryContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitDirectory(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::DirectoryContext* fileseqParser::directory() {
  DirectoryContext *_localctx = _tracker.createInstance<DirectoryContext>(_ctx, getState());
  enterRule(_localctx, 10, fileseqParser::RuleDirectory);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(106);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == fileseqParser::SLASH) {
      setState(102);
      match(fileseqParser::SLASH);
      setState(104);
      _errHandler->sync(this);

      _la = _input->LA(1);
      if (_la == fileseqParser::SLASH) {
        setState(103);
        match(fileseqParser::SLASH);
      }
    }
    setState(113);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 13, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        setState(108);
        dirSegment();
        setState(109);
        match(fileseqParser::SLASH); 
      }
      setState(115);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 13, _ctx);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- BasenameCharContext ------------------------------------------------------------------

fileseqParser::BasenameCharContext::BasenameCharContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* fileseqParser::BasenameCharContext::WORD() {
  return getToken(fileseqParser::WORD, 0);
}

tree::TerminalNode* fileseqParser::BasenameCharContext::NUM() {
  return getToken(fileseqParser::NUM, 0);
}

tree::TerminalNode* fileseqParser::BasenameCharContext::DOT_NUM() {
  return getToken(fileseqParser::DOT_NUM, 0);
}

tree::TerminalNode* fileseqParser::BasenameCharContext::DASH() {
  return getToken(fileseqParser::DASH, 0);
}

tree::TerminalNode* fileseqParser::BasenameCharContext::SPECIAL_CHAR() {
  return getToken(fileseqParser::SPECIAL_CHAR, 0);
}

tree::TerminalNode* fileseqParser::BasenameCharContext::EXTENSION() {
  return getToken(fileseqParser::EXTENSION, 0);
}

tree::TerminalNode* fileseqParser::BasenameCharContext::FRAME_RANGE() {
  return getToken(fileseqParser::FRAME_RANGE, 0);
}

tree::TerminalNode* fileseqParser::BasenameCharContext::DOT_FRAME_RANGE() {
  return getToken(fileseqParser::DOT_FRAME_RANGE, 0);
}

tree::TerminalNode* fileseqParser::BasenameCharContext::WS() {
  return getToken(fileseqParser::WS, 0);
}

tree::TerminalNode* fileseqParser::BasenameCharContext::OTHER_CHAR() {
  return getToken(fileseqParser::OTHER_CHAR, 0);
}


size_t fileseqParser::BasenameCharContext::getRuleIndex() const {
  return fileseqParser::RuleBasenameChar;
}


antlrcpp::Any fileseqParser::BasenameCharContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitBasenameChar(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::BasenameCharContext* fileseqParser::basenameChar() {
  BasenameCharContext *_localctx = _tracker.createInstance<BasenameCharContext>(_ctx, getState());
  enterRule(_localctx, 12, fileseqParser::RuleBasenameChar);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(116);
    _la = _input->LA(1);
    if (!((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << fileseqParser::EXTENSION)
      | (1ULL << fileseqParser::DOT_FRAME_RANGE)
      | (1ULL << fileseqParser::FRAME_RANGE)
      | (1ULL << fileseqParser::DOT_NUM)
      | (1ULL << fileseqParser::SPECIAL_CHAR)
      | (1ULL << fileseqParser::NUM)
      | (1ULL << fileseqParser::WORD)
      | (1ULL << fileseqParser::DASH)
      | (1ULL << fileseqParser::WS)
      | (1ULL << fileseqParser::OTHER_CHAR))) != 0))) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PlainBasenameCharContext ------------------------------------------------------------------

fileseqParser::PlainBasenameCharContext::PlainBasenameCharContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* fileseqParser::PlainBasenameCharContext::WORD() {
  return getToken(fileseqParser::WORD, 0);
}

tree::TerminalNode* fileseqParser::PlainBasenameCharContext::NUM() {
  return getToken(fileseqParser::NUM, 0);
}

tree::TerminalNode* fileseqParser::PlainBasenameCharContext::DASH() {
  return getToken(fileseqParser::DASH, 0);
}

tree::TerminalNode* fileseqParser::PlainBasenameCharContext::SPECIAL_CHAR() {
  return getToken(fileseqParser::SPECIAL_CHAR, 0);
}

tree::TerminalNode* fileseqParser::PlainBasenameCharContext::FRAME_RANGE() {
  return getToken(fileseqParser::FRAME_RANGE, 0);
}

tree::TerminalNode* fileseqParser::PlainBasenameCharContext::DOT_FRAME_RANGE() {
  return getToken(fileseqParser::DOT_FRAME_RANGE, 0);
}

tree::TerminalNode* fileseqParser::PlainBasenameCharContext::WS() {
  return getToken(fileseqParser::WS, 0);
}

tree::TerminalNode* fileseqParser::PlainBasenameCharContext::OTHER_CHAR() {
  return getToken(fileseqParser::OTHER_CHAR, 0);
}

tree::TerminalNode* fileseqParser::PlainBasenameCharContext::HASH() {
  return getToken(fileseqParser::HASH, 0);
}

tree::TerminalNode* fileseqParser::PlainBasenameCharContext::AT() {
  return getToken(fileseqParser::AT, 0);
}


size_t fileseqParser::PlainBasenameCharContext::getRuleIndex() const {
  return fileseqParser::RulePlainBasenameChar;
}


antlrcpp::Any fileseqParser::PlainBasenameCharContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitPlainBasenameChar(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::PlainBasenameCharContext* fileseqParser::plainBasenameChar() {
  PlainBasenameCharContext *_localctx = _tracker.createInstance<PlainBasenameCharContext>(_ctx, getState());
  enterRule(_localctx, 14, fileseqParser::RulePlainBasenameChar);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(118);
    _la = _input->LA(1);
    if (!((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << fileseqParser::HASH)
      | (1ULL << fileseqParser::AT)
      | (1ULL << fileseqParser::DOT_FRAME_RANGE)
      | (1ULL << fileseqParser::FRAME_RANGE)
      | (1ULL << fileseqParser::SPECIAL_CHAR)
      | (1ULL << fileseqParser::NUM)
      | (1ULL << fileseqParser::WORD)
      | (1ULL << fileseqParser::DASH)
      | (1ULL << fileseqParser::WS)
      | (1ULL << fileseqParser::OTHER_CHAR))) != 0))) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- DirSegmentContext ------------------------------------------------------------------

fileseqParser::DirSegmentContext::DirSegmentContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<fileseqParser::BasenameCharContext *> fileseqParser::DirSegmentContext::basenameChar() {
  return getRuleContexts<fileseqParser::BasenameCharContext>();
}

fileseqParser::BasenameCharContext* fileseqParser::DirSegmentContext::basenameChar(size_t i) {
  return getRuleContext<fileseqParser::BasenameCharContext>(i);
}


size_t fileseqParser::DirSegmentContext::getRuleIndex() const {
  return fileseqParser::RuleDirSegment;
}


antlrcpp::Any fileseqParser::DirSegmentContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitDirSegment(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::DirSegmentContext* fileseqParser::dirSegment() {
  DirSegmentContext *_localctx = _tracker.createInstance<DirSegmentContext>(_ctx, getState());
  enterRule(_localctx, 16, fileseqParser::RuleDirSegment);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(121); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(120);
      basenameChar();
      setState(123); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << fileseqParser::EXTENSION)
      | (1ULL << fileseqParser::DOT_FRAME_RANGE)
      | (1ULL << fileseqParser::FRAME_RANGE)
      | (1ULL << fileseqParser::DOT_NUM)
      | (1ULL << fileseqParser::SPECIAL_CHAR)
      | (1ULL << fileseqParser::NUM)
      | (1ULL << fileseqParser::WORD)
      | (1ULL << fileseqParser::DASH)
      | (1ULL << fileseqParser::WS)
      | (1ULL << fileseqParser::OTHER_CHAR))) != 0));
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- BasenameContext ------------------------------------------------------------------

fileseqParser::BasenameContext::BasenameContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<fileseqParser::BasenameCharContext *> fileseqParser::BasenameContext::basenameChar() {
  return getRuleContexts<fileseqParser::BasenameCharContext>();
}

fileseqParser::BasenameCharContext* fileseqParser::BasenameContext::basenameChar(size_t i) {
  return getRuleContext<fileseqParser::BasenameCharContext>(i);
}


size_t fileseqParser::BasenameContext::getRuleIndex() const {
  return fileseqParser::RuleBasename;
}


antlrcpp::Any fileseqParser::BasenameContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitBasename(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::BasenameContext* fileseqParser::basename() {
  BasenameContext *_localctx = _tracker.createInstance<BasenameContext>(_ctx, getState());
  enterRule(_localctx, 18, fileseqParser::RuleBasename);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(126); 
    _errHandler->sync(this);
    alt = 1;
    do {
      switch (alt) {
        case 1: {
              setState(125);
              basenameChar();
              break;
            }

      default:
        throw NoViableAltException(this);
      }
      setState(128); 
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 15, _ctx);
    } while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PatBasenameContext ------------------------------------------------------------------

fileseqParser::PatBasenameContext::PatBasenameContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<fileseqParser::BasenameCharContext *> fileseqParser::PatBasenameContext::basenameChar() {
  return getRuleContexts<fileseqParser::BasenameCharContext>();
}

fileseqParser::BasenameCharContext* fileseqParser::PatBasenameContext::basenameChar(size_t i) {
  return getRuleContext<fileseqParser::BasenameCharContext>(i);
}


size_t fileseqParser::PatBasenameContext::getRuleIndex() const {
  return fileseqParser::RulePatBasename;
}


antlrcpp::Any fileseqParser::PatBasenameContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitPatBasename(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::PatBasenameContext* fileseqParser::patBasename() {
  PatBasenameContext *_localctx = _tracker.createInstance<PatBasenameContext>(_ctx, getState());
  enterRule(_localctx, 20, fileseqParser::RulePatBasename);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(131); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(130);
      basenameChar();
      setState(133); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << fileseqParser::EXTENSION)
      | (1ULL << fileseqParser::DOT_FRAME_RANGE)
      | (1ULL << fileseqParser::FRAME_RANGE)
      | (1ULL << fileseqParser::DOT_NUM)
      | (1ULL << fileseqParser::SPECIAL_CHAR)
      | (1ULL << fileseqParser::NUM)
      | (1ULL << fileseqParser::WORD)
      | (1ULL << fileseqParser::DASH)
      | (1ULL << fileseqParser::WS)
      | (1ULL << fileseqParser::OTHER_CHAR))) != 0));
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- FrameBasenameContext ------------------------------------------------------------------

fileseqParser::FrameBasenameContext::FrameBasenameContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<fileseqParser::BasenameCharContext *> fileseqParser::FrameBasenameContext::basenameChar() {
  return getRuleContexts<fileseqParser::BasenameCharContext>();
}

fileseqParser::BasenameCharContext* fileseqParser::FrameBasenameContext::basenameChar(size_t i) {
  return getRuleContext<fileseqParser::BasenameCharContext>(i);
}


size_t fileseqParser::FrameBasenameContext::getRuleIndex() const {
  return fileseqParser::RuleFrameBasename;
}


antlrcpp::Any fileseqParser::FrameBasenameContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitFrameBasename(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::FrameBasenameContext* fileseqParser::frameBasename() {
  FrameBasenameContext *_localctx = _tracker.createInstance<FrameBasenameContext>(_ctx, getState());
  enterRule(_localctx, 22, fileseqParser::RuleFrameBasename);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(136); 
    _errHandler->sync(this);
    alt = 1;
    do {
      switch (alt) {
        case 1: {
              setState(135);
              basenameChar();
              break;
            }

      default:
        throw NoViableAltException(this);
      }
      setState(138); 
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 17, _ctx);
    } while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PlainBasenameContext ------------------------------------------------------------------

fileseqParser::PlainBasenameContext::PlainBasenameContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<fileseqParser::PlainBasenameCharContext *> fileseqParser::PlainBasenameContext::plainBasenameChar() {
  return getRuleContexts<fileseqParser::PlainBasenameCharContext>();
}

fileseqParser::PlainBasenameCharContext* fileseqParser::PlainBasenameContext::plainBasenameChar(size_t i) {
  return getRuleContext<fileseqParser::PlainBasenameCharContext>(i);
}


size_t fileseqParser::PlainBasenameContext::getRuleIndex() const {
  return fileseqParser::RulePlainBasename;
}


antlrcpp::Any fileseqParser::PlainBasenameContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitPlainBasename(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::PlainBasenameContext* fileseqParser::plainBasename() {
  PlainBasenameContext *_localctx = _tracker.createInstance<PlainBasenameContext>(_ctx, getState());
  enterRule(_localctx, 24, fileseqParser::RulePlainBasename);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(141); 
    _errHandler->sync(this);
    alt = 1;
    do {
      switch (alt) {
        case 1: {
              setState(140);
              plainBasenameChar();
              break;
            }

      default:
        throw NoViableAltException(this);
      }
      setState(143); 
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 18, _ctx);
    } while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- FrameRangeContext ------------------------------------------------------------------

fileseqParser::FrameRangeContext::FrameRangeContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* fileseqParser::FrameRangeContext::DOT_FRAME_RANGE() {
  return getToken(fileseqParser::DOT_FRAME_RANGE, 0);
}

tree::TerminalNode* fileseqParser::FrameRangeContext::FRAME_RANGE() {
  return getToken(fileseqParser::FRAME_RANGE, 0);
}

tree::TerminalNode* fileseqParser::FrameRangeContext::DOT_NUM() {
  return getToken(fileseqParser::DOT_NUM, 0);
}

tree::TerminalNode* fileseqParser::FrameRangeContext::NUM() {
  return getToken(fileseqParser::NUM, 0);
}


size_t fileseqParser::FrameRangeContext::getRuleIndex() const {
  return fileseqParser::RuleFrameRange;
}


antlrcpp::Any fileseqParser::FrameRangeContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitFrameRange(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::FrameRangeContext* fileseqParser::frameRange() {
  FrameRangeContext *_localctx = _tracker.createInstance<FrameRangeContext>(_ctx, getState());
  enterRule(_localctx, 26, fileseqParser::RuleFrameRange);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(145);
    _la = _input->LA(1);
    if (!((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << fileseqParser::DOT_FRAME_RANGE)
      | (1ULL << fileseqParser::FRAME_RANGE)
      | (1ULL << fileseqParser::DOT_NUM)
      | (1ULL << fileseqParser::NUM))) != 0))) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- FrameNumContext ------------------------------------------------------------------

fileseqParser::FrameNumContext::FrameNumContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* fileseqParser::FrameNumContext::DOT_NUM() {
  return getToken(fileseqParser::DOT_NUM, 0);
}


size_t fileseqParser::FrameNumContext::getRuleIndex() const {
  return fileseqParser::RuleFrameNum;
}


antlrcpp::Any fileseqParser::FrameNumContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitFrameNum(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::FrameNumContext* fileseqParser::frameNum() {
  FrameNumContext *_localctx = _tracker.createInstance<FrameNumContext>(_ctx, getState());
  enterRule(_localctx, 28, fileseqParser::RuleFrameNum);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(147);
    match(fileseqParser::DOT_NUM);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PaddingContext ------------------------------------------------------------------

fileseqParser::PaddingContext::PaddingContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* fileseqParser::PaddingContext::UDIM_ANGLE() {
  return getToken(fileseqParser::UDIM_ANGLE, 0);
}

tree::TerminalNode* fileseqParser::PaddingContext::UDIM_PAREN() {
  return getToken(fileseqParser::UDIM_PAREN, 0);
}

tree::TerminalNode* fileseqParser::PaddingContext::PRINTF_PAD() {
  return getToken(fileseqParser::PRINTF_PAD, 0);
}

tree::TerminalNode* fileseqParser::PaddingContext::HOUDINI_PAD() {
  return getToken(fileseqParser::HOUDINI_PAD, 0);
}

std::vector<tree::TerminalNode *> fileseqParser::PaddingContext::HASH() {
  return getTokens(fileseqParser::HASH);
}

tree::TerminalNode* fileseqParser::PaddingContext::HASH(size_t i) {
  return getToken(fileseqParser::HASH, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PaddingContext::AT() {
  return getTokens(fileseqParser::AT);
}

tree::TerminalNode* fileseqParser::PaddingContext::AT(size_t i) {
  return getToken(fileseqParser::AT, i);
}


size_t fileseqParser::PaddingContext::getRuleIndex() const {
  return fileseqParser::RulePadding;
}


antlrcpp::Any fileseqParser::PaddingContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitPadding(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::PaddingContext* fileseqParser::padding() {
  PaddingContext *_localctx = _tracker.createInstance<PaddingContext>(_ctx, getState());
  enterRule(_localctx, 30, fileseqParser::RulePadding);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(158);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case fileseqParser::UDIM_ANGLE: {
        enterOuterAlt(_localctx, 1);
        setState(149);
        match(fileseqParser::UDIM_ANGLE);
        break;
      }

      case fileseqParser::UDIM_PAREN: {
        enterOuterAlt(_localctx, 2);
        setState(150);
        match(fileseqParser::UDIM_PAREN);
        break;
      }

      case fileseqParser::PRINTF_PAD: {
        enterOuterAlt(_localctx, 3);
        setState(151);
        match(fileseqParser::PRINTF_PAD);
        break;
      }

      case fileseqParser::HOUDINI_PAD: {
        enterOuterAlt(_localctx, 4);
        setState(152);
        match(fileseqParser::HOUDINI_PAD);
        break;
      }

      case fileseqParser::HASH:
      case fileseqParser::AT: {
        enterOuterAlt(_localctx, 5);
        setState(154); 
        _errHandler->sync(this);
        _la = _input->LA(1);
        do {
          setState(153);
          _la = _input->LA(1);
          if (!(_la == fileseqParser::HASH

          || _la == fileseqParser::AT)) {
          _errHandler->recoverInline(this);
          }
          else {
            _errHandler->reportMatch(this);
            consume();
          }
          setState(156); 
          _errHandler->sync(this);
          _la = _input->LA(1);
        } while (_la == fileseqParser::HASH

        || _la == fileseqParser::AT);
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ExtensionContext ------------------------------------------------------------------

fileseqParser::ExtensionContext::ExtensionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* fileseqParser::ExtensionContext::EXTENSION() {
  return getToken(fileseqParser::EXTENSION, 0);
}

tree::TerminalNode* fileseqParser::ExtensionContext::DASH() {
  return getToken(fileseqParser::DASH, 0);
}

tree::TerminalNode* fileseqParser::ExtensionContext::NUM() {
  return getToken(fileseqParser::NUM, 0);
}

tree::TerminalNode* fileseqParser::ExtensionContext::DOT_NUM() {
  return getToken(fileseqParser::DOT_NUM, 0);
}

tree::TerminalNode* fileseqParser::ExtensionContext::WORD() {
  return getToken(fileseqParser::WORD, 0);
}


size_t fileseqParser::ExtensionContext::getRuleIndex() const {
  return fileseqParser::RuleExtension;
}


antlrcpp::Any fileseqParser::ExtensionContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitExtension(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::ExtensionContext* fileseqParser::extension() {
  ExtensionContext *_localctx = _tracker.createInstance<ExtensionContext>(_ctx, getState());
  enterRule(_localctx, 32, fileseqParser::RuleExtension);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(167);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case fileseqParser::EXTENSION: {
        enterOuterAlt(_localctx, 1);
        setState(160);
        match(fileseqParser::EXTENSION);
        setState(163);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == fileseqParser::DASH) {
          setState(161);
          match(fileseqParser::DASH);
          setState(162);
          match(fileseqParser::NUM);
        }
        break;
      }

      case fileseqParser::DOT_NUM: {
        enterOuterAlt(_localctx, 2);
        setState(165);
        match(fileseqParser::DOT_NUM);
        break;
      }

      case fileseqParser::WORD: {
        enterOuterAlt(_localctx, 3);
        setState(166);
        match(fileseqParser::WORD);
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

// Static vars and initialization.
std::vector<dfa::DFA> fileseqParser::_decisionToDFA;
atn::PredictionContextCache fileseqParser::_sharedContextCache;

// We own the ATN which in turn owns the ATN states.
atn::ATN fileseqParser::_atn;
std::vector<uint16_t> fileseqParser::_serializedATN;

std::vector<std::string> fileseqParser::_ruleNames = {
  "input", "sequence", "patternOnly", "singleFrame", "plainFile", "directory", 
  "basenameChar", "plainBasenameChar", "dirSegment", "basename", "patBasename", 
  "frameBasename", "plainBasename", "frameRange", "frameNum", "padding", 
  "extension"
};

std::vector<std::string> fileseqParser::_literalNames = {
  "", "'<UDIM>'", "'%(UDIM)d'", "", "", "'#'", "'@'", "", "", "", "", "", 
  "", "", "", "'-'"
};

std::vector<std::string> fileseqParser::_symbolicNames = {
  "", "UDIM_ANGLE", "UDIM_PAREN", "PRINTF_PAD", "HOUDINI_PAD", "HASH", "AT", 
  "EXTENSION", "DOT_FRAME_RANGE", "FRAME_RANGE", "DOT_NUM", "SLASH", "SPECIAL_CHAR", 
  "NUM", "WORD", "DASH", "WS", "OTHER_CHAR"
};

dfa::Vocabulary fileseqParser::_vocabulary(_literalNames, _symbolicNames);

std::vector<std::string> fileseqParser::_tokenNames;

fileseqParser::Initializer::Initializer() {
	for (size_t i = 0; i < _symbolicNames.size(); ++i) {
		std::string name = _vocabulary.getLiteralName(i);
		if (name.empty()) {
			name = _vocabulary.getSymbolicName(i);
		}

		if (name.empty()) {
			_tokenNames.push_back("<INVALID>");
		} else {
      _tokenNames.push_back(name);
    }
	}

  static const uint16_t serializedATNSegment0[] = {
    0x3, 0x608b, 0xa72a, 0x8133, 0xb9ed, 0x417c, 0x3be7, 0x7786, 0x5964, 
       0x3, 0x13, 0xac, 0x4, 0x2, 0x9, 0x2, 0x4, 0x3, 0x9, 0x3, 0x4, 0x4, 
       0x9, 0x4, 0x4, 0x5, 0x9, 0x5, 0x4, 0x6, 0x9, 0x6, 0x4, 0x7, 0x9, 
       0x7, 0x4, 0x8, 0x9, 0x8, 0x4, 0x9, 0x9, 0x9, 0x4, 0xa, 0x9, 0xa, 
       0x4, 0xb, 0x9, 0xb, 0x4, 0xc, 0x9, 0xc, 0x4, 0xd, 0x9, 0xd, 0x4, 
       0xe, 0x9, 0xe, 0x4, 0xf, 0x9, 0xf, 0x4, 0x10, 0x9, 0x10, 0x4, 0x11, 
       0x9, 0x11, 0x4, 0x12, 0x9, 0x12, 0x3, 0x2, 0x3, 0x2, 0x3, 0x2, 0x3, 
       0x2, 0x3, 0x2, 0x3, 0x2, 0x3, 0x2, 0x3, 0x2, 0x3, 0x2, 0x3, 0x2, 
       0x3, 0x2, 0x3, 0x2, 0x5, 0x2, 0x31, 0xa, 0x2, 0x3, 0x3, 0x3, 0x3, 
       0x5, 0x3, 0x35, 0xa, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 
       0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x5, 0x3, 0x3e, 0xa, 0x3, 0x3, 0x3, 
       0x7, 0x3, 0x41, 0xa, 0x3, 0xc, 0x3, 0xe, 0x3, 0x44, 0xb, 0x3, 0x3, 
       0x4, 0x3, 0x4, 0x5, 0x4, 0x48, 0xa, 0x4, 0x3, 0x4, 0x3, 0x4, 0x3, 
       0x4, 0x5, 0x4, 0x4d, 0xa, 0x4, 0x3, 0x4, 0x7, 0x4, 0x50, 0xa, 0x4, 
       0xc, 0x4, 0xe, 0x4, 0x53, 0xb, 0x4, 0x3, 0x5, 0x3, 0x5, 0x5, 0x5, 
       0x57, 0xa, 0x5, 0x3, 0x5, 0x3, 0x5, 0x6, 0x5, 0x5b, 0xa, 0x5, 0xd, 
       0x5, 0xe, 0x5, 0x5c, 0x3, 0x6, 0x3, 0x6, 0x5, 0x6, 0x61, 0xa, 0x6, 
       0x3, 0x6, 0x7, 0x6, 0x64, 0xa, 0x6, 0xc, 0x6, 0xe, 0x6, 0x67, 0xb, 
       0x6, 0x3, 0x7, 0x3, 0x7, 0x5, 0x7, 0x6b, 0xa, 0x7, 0x5, 0x7, 0x6d, 
       0xa, 0x7, 0x3, 0x7, 0x3, 0x7, 0x3, 0x7, 0x7, 0x7, 0x72, 0xa, 0x7, 
       0xc, 0x7, 0xe, 0x7, 0x75, 0xb, 0x7, 0x3, 0x8, 0x3, 0x8, 0x3, 0x9, 
       0x3, 0x9, 0x3, 0xa, 0x6, 0xa, 0x7c, 0xa, 0xa, 0xd, 0xa, 0xe, 0xa, 
       0x7d, 0x3, 0xb, 0x6, 0xb, 0x81, 0xa, 0xb, 0xd, 0xb, 0xe, 0xb, 0x82, 
       0x3, 0xc, 0x6, 0xc, 0x86, 0xa, 0xc, 0xd, 0xc, 0xe, 0xc, 0x87, 0x3, 
       0xd, 0x6, 0xd, 0x8b, 0xa, 0xd, 0xd, 0xd, 0xe, 0xd, 0x8c, 0x3, 0xe, 
       0x6, 0xe, 0x90, 0xa, 0xe, 0xd, 0xe, 0xe, 0xe, 0x91, 0x3, 0xf, 0x3, 
       0xf, 0x3, 0x10, 0x3, 0x10, 0x3, 0x11, 0x3, 0x11, 0x3, 0x11, 0x3, 
       0x11, 0x3, 0x11, 0x6, 0x11, 0x9d, 0xa, 0x11, 0xd, 0x11, 0xe, 0x11, 
       0x9e, 0x5, 0x11, 0xa1, 0xa, 0x11, 0x3, 0x12, 0x3, 0x12, 0x3, 0x12, 
       0x5, 0x12, 0xa6, 0xa, 0x12, 0x3, 0x12, 0x3, 0x12, 0x5, 0x12, 0xaa, 
       0xa, 0x12, 0x3, 0x12, 0x2, 0x2, 0x13, 0x2, 0x4, 0x6, 0x8, 0xa, 0xc, 
       0xe, 0x10, 0x12, 0x14, 0x16, 0x18, 0x1a, 0x1c, 0x1e, 0x20, 0x22, 
       0x2, 0x6, 0x4, 0x2, 0x9, 0xc, 0xe, 0x13, 0x5, 0x2, 0x7, 0x8, 0xa, 
       0xb, 0xe, 0x13, 0x4, 0x2, 0xa, 0xc, 0xf, 0xf, 0x3, 0x2, 0x7, 0x8, 
       0x2, 0xb8, 0x2, 0x30, 0x3, 0x2, 0x2, 0x2, 0x4, 0x32, 0x3, 0x2, 0x2, 
       0x2, 0x6, 0x45, 0x3, 0x2, 0x2, 0x2, 0x8, 0x54, 0x3, 0x2, 0x2, 0x2, 
       0xa, 0x5e, 0x3, 0x2, 0x2, 0x2, 0xc, 0x6c, 0x3, 0x2, 0x2, 0x2, 0xe, 
       0x76, 0x3, 0x2, 0x2, 0x2, 0x10, 0x78, 0x3, 0x2, 0x2, 0x2, 0x12, 0x7b, 
       0x3, 0x2, 0x2, 0x2, 0x14, 0x80, 0x3, 0x2, 0x2, 0x2, 0x16, 0x85, 0x3, 
       0x2, 0x2, 0x2, 0x18, 0x8a, 0x3, 0x2, 0x2, 0x2, 0x1a, 0x8f, 0x3, 0x2, 
       0x2, 0x2, 0x1c, 0x93, 0x3, 0x2, 0x2, 0x2, 0x1e, 0x95, 0x3, 0x2, 0x2, 
       0x2, 0x20, 0xa0, 0x3, 0x2, 0x2, 0x2, 0x22, 0xa9, 0x3, 0x2, 0x2, 0x2, 
       0x24, 0x25, 0x5, 0x4, 0x3, 0x2, 0x25, 0x26, 0x7, 0x2, 0x2, 0x3, 0x26, 
       0x31, 0x3, 0x2, 0x2, 0x2, 0x27, 0x28, 0x5, 0x6, 0x4, 0x2, 0x28, 0x29, 
       0x7, 0x2, 0x2, 0x3, 0x29, 0x31, 0x3, 0x2, 0x2, 0x2, 0x2a, 0x2b, 0x5, 
       0x8, 0x5, 0x2, 0x2b, 0x2c, 0x7, 0x2, 0x2, 0x3, 0x2c, 0x31, 0x3, 0x2, 
       0x2, 0x2, 0x2d, 0x2e, 0x5, 0xa, 0x6, 0x2, 0x2e, 0x2f, 0x7, 0x2, 0x2, 
       0x3, 0x2f, 0x31, 0x3, 0x2, 0x2, 0x2, 0x30, 0x24, 0x3, 0x2, 0x2, 0x2, 
       0x30, 0x27, 0x3, 0x2, 0x2, 0x2, 0x30, 0x2a, 0x3, 0x2, 0x2, 0x2, 0x30, 
       0x2d, 0x3, 0x2, 0x2, 0x2, 0x31, 0x3, 0x3, 0x2, 0x2, 0x2, 0x32, 0x34, 
       0x5, 0xc, 0x7, 0x2, 0x33, 0x35, 0x5, 0x14, 0xb, 0x2, 0x34, 0x33, 
       0x3, 0x2, 0x2, 0x2, 0x34, 0x35, 0x3, 0x2, 0x2, 0x2, 0x35, 0x36, 0x3, 
       0x2, 0x2, 0x2, 0x36, 0x37, 0x5, 0x1c, 0xf, 0x2, 0x37, 0x3d, 0x5, 
       0x20, 0x11, 0x2, 0x38, 0x39, 0x5, 0x1c, 0xf, 0x2, 0x39, 0x3a, 0x5, 
       0x20, 0x11, 0x2, 0x3a, 0x3e, 0x3, 0x2, 0x2, 0x2, 0x3b, 0x3c, 0x7, 
       0xe, 0x2, 0x2, 0x3c, 0x3e, 0x5, 0x20, 0x11, 0x2, 0x3d, 0x38, 0x3, 
       0x2, 0x2, 0x2, 0x3d, 0x3b, 0x3, 0x2, 0x2, 0x2, 0x3d, 0x3e, 0x3, 0x2, 
       0x2, 0x2, 0x3e, 0x42, 0x3, 0x2, 0x2, 0x2, 0x3f, 0x41, 0x5, 0x22, 
       0x12, 0x2, 0x40, 0x3f, 0x3, 0x2, 0x2, 0x2, 0x41, 0x44, 0x3, 0x2, 
       0x2, 0x2, 0x42, 0x40, 0x3, 0x2, 0x2, 0x2, 0x42, 0x43, 0x3, 0x2, 0x2, 
       0x2, 0x43, 0x5, 0x3, 0x2, 0x2, 0x2, 0x44, 0x42, 0x3, 0x2, 0x2, 0x2, 
       0x45, 0x47, 0x5, 0xc, 0x7, 0x2, 0x46, 0x48, 0x5, 0x16, 0xc, 0x2, 
       0x47, 0x46, 0x3, 0x2, 0x2, 0x2, 0x47, 0x48, 0x3, 0x2, 0x2, 0x2, 0x48, 
       0x49, 0x3, 0x2, 0x2, 0x2, 0x49, 0x4c, 0x5, 0x20, 0x11, 0x2, 0x4a, 
       0x4b, 0x7, 0xe, 0x2, 0x2, 0x4b, 0x4d, 0x5, 0x20, 0x11, 0x2, 0x4c, 
       0x4a, 0x3, 0x2, 0x2, 0x2, 0x4c, 0x4d, 0x3, 0x2, 0x2, 0x2, 0x4d, 0x51, 
       0x3, 0x2, 0x2, 0x2, 0x4e, 0x50, 0x5, 0x22, 0x12, 0x2, 0x4f, 0x4e, 
       0x3, 0x2, 0x2, 0x2, 0x50, 0x53, 0x3, 0x2, 0x2, 0x2, 0x51, 0x4f, 0x3, 
       0x2, 0x2, 0x2, 0x51, 0x52, 0x3, 0x2, 0x2, 0x2, 0x52, 0x7, 0x3, 0x2, 
       0x2, 0x2, 0x53, 0x51, 0x3, 0x2, 0x2, 0x2, 0x54, 0x56, 0x5, 0xc, 0x7, 
       0x2, 0x55, 0x57, 0x5, 0x18, 0xd, 0x2, 0x56, 0x55, 0x3, 0x2, 0x2, 
       0x2, 0x56, 0x57, 0x3, 0x2, 0x2, 0x2, 0x57, 0x58, 0x3, 0x2, 0x2, 0x2, 
       0x58, 0x5a, 0x5, 0x1e, 0x10, 0x2, 0x59, 0x5b, 0x5, 0x22, 0x12, 0x2, 
       0x5a, 0x59, 0x3, 0x2, 0x2, 0x2, 0x5b, 0x5c, 0x3, 0x2, 0x2, 0x2, 0x5c, 
       0x5a, 0x3, 0x2, 0x2, 0x2, 0x5c, 0x5d, 0x3, 0x2, 0x2, 0x2, 0x5d, 0x9, 
       0x3, 0x2, 0x2, 0x2, 0x5e, 0x60, 0x5, 0xc, 0x7, 0x2, 0x5f, 0x61, 0x5, 
       0x1a, 0xe, 0x2, 0x60, 0x5f, 0x3, 0x2, 0x2, 0x2, 0x60, 0x61, 0x3, 
       0x2, 0x2, 0x2, 0x61, 0x65, 0x3, 0x2, 0x2, 0x2, 0x62, 0x64, 0x5, 0x22, 
       0x12, 0x2, 0x63, 0x62, 0x3, 0x2, 0x2, 0x2, 0x64, 0x67, 0x3, 0x2, 
       0x2, 0x2, 0x65, 0x63, 0x3, 0x2, 0x2, 0x2, 0x65, 0x66, 0x3, 0x2, 0x2, 
       0x2, 0x66, 0xb, 0x3, 0x2, 0x2, 0x2, 0x67, 0x65, 0x3, 0x2, 0x2, 0x2, 
       0x68, 0x6a, 0x7, 0xd, 0x2, 0x2, 0x69, 0x6b, 0x7, 0xd, 0x2, 0x2, 0x6a, 
       0x69, 0x3, 0x2, 0x2, 0x2, 0x6a, 0x6b, 0x3, 0x2, 0x2, 0x2, 0x6b, 0x6d, 
       0x3, 0x2, 0x2, 0x2, 0x6c, 0x68, 0x3, 0x2, 0x2, 0x2, 0x6c, 0x6d, 0x3, 
       0x2, 0x2, 0x2, 0x6d, 0x73, 0x3, 0x2, 0x2, 0x2, 0x6e, 0x6f, 0x5, 0x12, 
       0xa, 0x2, 0x6f, 0x70, 0x7, 0xd, 0x2, 0x2, 0x70, 0x72, 0x3, 0x2, 0x2, 
       0x2, 0x71, 0x6e, 0x3, 0x2, 0x2, 0x2, 0x72, 0x75, 0x3, 0x2, 0x2, 0x2, 
       0x73, 0x71, 0x3, 0x2, 0x2, 0x2, 0x73, 0x74, 0x3, 0x2, 0x2, 0x2, 0x74, 
       0xd, 0x3, 0x2, 0x2, 0x2, 0x75, 0x73, 0x3, 0x2, 0x2, 0x2, 0x76, 0x77, 
       0x9, 0x2, 0x2, 0x2, 0x77, 0xf, 0x3, 0x2, 0x2, 0x2, 0x78, 0x79, 0x9, 
       0x3, 0x2, 0x2, 0x79, 0x11, 0x3, 0x2, 0x2, 0x2, 0x7a, 0x7c, 0x5, 0xe, 
       0x8, 0x2, 0x7b, 0x7a, 0x3, 0x2, 0x2, 0x2, 0x7c, 0x7d, 0x3, 0x2, 0x2, 
       0x2, 0x7d, 0x7b, 0x3, 0x2, 0x2, 0x2, 0x7d, 0x7e, 0x3, 0x2, 0x2, 0x2, 
       0x7e, 0x13, 0x3, 0x2, 0x2, 0x2, 0x7f, 0x81, 0x5, 0xe, 0x8, 0x2, 0x80, 
       0x7f, 0x3, 0x2, 0x2, 0x2, 0x81, 0x82, 0x3, 0x2, 0x2, 0x2, 0x82, 0x80, 
       0x3, 0x2, 0x2, 0x2, 0x82, 0x83, 0x3, 0x2, 0x2, 0x2, 0x83, 0x15, 0x3, 
       0x2, 0x2, 0x2, 0x84, 0x86, 0x5, 0xe, 0x8, 0x2, 0x85, 0x84, 0x3, 0x2, 
       0x2, 0x2, 0x86, 0x87, 0x3, 0x2, 0x2, 0x2, 0x87, 0x85, 0x3, 0x2, 0x2, 
       0x2, 0x87, 0x88, 0x3, 0x2, 0x2, 0x2, 0x88, 0x17, 0x3, 0x2, 0x2, 0x2, 
       0x89, 0x8b, 0x5, 0xe, 0x8, 0x2, 0x8a, 0x89, 0x3, 0x2, 0x2, 0x2, 0x8b, 
       0x8c, 0x3, 0x2, 0x2, 0x2, 0x8c, 0x8a, 0x3, 0x2, 0x2, 0x2, 0x8c, 0x8d, 
       0x3, 0x2, 0x2, 0x2, 0x8d, 0x19, 0x3, 0x2, 0x2, 0x2, 0x8e, 0x90, 0x5, 
       0x10, 0x9, 0x2, 0x8f, 0x8e, 0x3, 0x2, 0x2, 0x2, 0x90, 0x91, 0x3, 
       0x2, 0x2, 0x2, 0x91, 0x8f, 0x3, 0x2, 0x2, 0x2, 0x91, 0x92, 0x3, 0x2, 
       0x2, 0x2, 0x92, 0x1b, 0x3, 0x2, 0x2, 0x2, 0x93, 0x94, 0x9, 0x4, 0x2, 
       0x2, 0x94, 0x1d, 0x3, 0x2, 0x2, 0x2, 0x95, 0x96, 0x7, 0xc, 0x2, 0x2, 
       0x96, 0x1f, 0x3, 0x2, 0x2, 0x2, 0x97, 0xa1, 0x7, 0x3, 0x2, 0x2, 0x98, 
       0xa1, 0x7, 0x4, 0x2, 0x2, 0x99, 0xa1, 0x7, 0x5, 0x2, 0x2, 0x9a, 0xa1, 
       0x7, 0x6, 0x2, 0x2, 0x9b, 0x9d, 0x9, 0x5, 0x2, 0x2, 0x9c, 0x9b, 0x3, 
       0x2, 0x2, 0x2, 0x9d, 0x9e, 0x3, 0x2, 0x2, 0x2, 0x9e, 0x9c, 0x3, 0x2, 
       0x2, 0x2, 0x9e, 0x9f, 0x3, 0x2, 0x2, 0x2, 0x9f, 0xa1, 0x3, 0x2, 0x2, 
       0x2, 0xa0, 0x97, 0x3, 0x2, 0x2, 0x2, 0xa0, 0x98, 0x3, 0x2, 0x2, 0x2, 
       0xa0, 0x99, 0x3, 0x2, 0x2, 0x2, 0xa0, 0x9a, 0x3, 0x2, 0x2, 0x2, 0xa0, 
       0x9c, 0x3, 0x2, 0x2, 0x2, 0xa1, 0x21, 0x3, 0x2, 0x2, 0x2, 0xa2, 0xa5, 
       0x7, 0x9, 0x2, 0x2, 0xa3, 0xa4, 0x7, 0x11, 0x2, 0x2, 0xa4, 0xa6, 
       0x7, 0xf, 0x2, 0x2, 0xa5, 0xa3, 0x3, 0x2, 0x2, 0x2, 0xa5, 0xa6, 0x3, 
       0x2, 0x2, 0x2, 0xa6, 0xaa, 0x3, 0x2, 0x2, 0x2, 0xa7, 0xaa, 0x7, 0xc, 
       0x2, 0x2, 0xa8, 0xaa, 0x7, 0x10, 0x2, 0x2, 0xa9, 0xa2, 0x3, 0x2, 
       0x2, 0x2, 0xa9, 0xa7, 0x3, 0x2, 0x2, 0x2, 0xa9, 0xa8, 0x3, 0x2, 0x2, 
       0x2, 0xaa, 0x23, 0x3, 0x2, 0x2, 0x2, 0x19, 0x30, 0x34, 0x3d, 0x42, 
       0x47, 0x4c, 0x51, 0x56, 0x5c, 0x60, 0x65, 0x6a, 0x6c, 0x73, 0x7d, 
       0x82, 0x87, 0x8c, 0x91, 0x9e, 0xa0, 0xa5, 0xa9, 
  };

  _serializedATN.insert(_serializedATN.end(), serializedATNSegment0,
    serializedATNSegment0 + sizeof(serializedATNSegment0) / sizeof(serializedATNSegment0[0]));


  atn::ATNDeserializer deserializer;
  _atn = deserializer.deserialize(_serializedATN);

  size_t count = _atn.getNumberOfDecisions();
  _decisionToDFA.reserve(count);
  for (size_t i = 0; i < count; i++) { 
    _decisionToDFA.emplace_back(_atn.getDecisionState(i), i);
  }
}

fileseqParser::Initializer fileseqParser::_init;
