
// Generated from fileseq.g4 by ANTLR 4.13.1


#include "fileseqVisitor.h"

#include "fileseqParser.h"


using namespace antlrcpp;
using namespace fileseq;

using namespace antlr4;

namespace {

struct FileseqParserStaticData final {
  FileseqParserStaticData(std::vector<std::string> ruleNames,
                        std::vector<std::string> literalNames,
                        std::vector<std::string> symbolicNames)
      : ruleNames(std::move(ruleNames)), literalNames(std::move(literalNames)),
        symbolicNames(std::move(symbolicNames)),
        vocabulary(this->literalNames, this->symbolicNames) {}

  FileseqParserStaticData(const FileseqParserStaticData&) = delete;
  FileseqParserStaticData(FileseqParserStaticData&&) = delete;
  FileseqParserStaticData& operator=(const FileseqParserStaticData&) = delete;
  FileseqParserStaticData& operator=(FileseqParserStaticData&&) = delete;

  std::vector<antlr4::dfa::DFA> decisionToDFA;
  antlr4::atn::PredictionContextCache sharedContextCache;
  const std::vector<std::string> ruleNames;
  const std::vector<std::string> literalNames;
  const std::vector<std::string> symbolicNames;
  const antlr4::dfa::Vocabulary vocabulary;
  antlr4::atn::SerializedATNView serializedATN;
  std::unique_ptr<antlr4::atn::ATN> atn;
};

::antlr4::internal::OnceFlag fileseqParserOnceFlag;
#if ANTLR4_USE_THREAD_LOCAL_CACHE
static thread_local
#endif
FileseqParserStaticData *fileseqParserStaticData = nullptr;

void fileseqParserInitialize() {
#if ANTLR4_USE_THREAD_LOCAL_CACHE
  if (fileseqParserStaticData != nullptr) {
    return;
  }
#else
  assert(fileseqParserStaticData == nullptr);
#endif
  auto staticData = std::make_unique<FileseqParserStaticData>(
    std::vector<std::string>{
      "input", "sequence", "patternOnly", "singleFrame", "plainFile", "directory", 
      "dirSegment", "sequenceBasename", "patternBasename", "singleFrameBasename", 
      "plainBasename", "frameRange", "frameNum", "padding", "extension"
    },
    std::vector<std::string>{
      "", "'<UDIM>'", "'%(UDIM)d'", "", "", "'#'", "'@'", "", "", "", "", 
      "", "", "", "", "'-'"
    },
    std::vector<std::string>{
      "", "UDIM_ANGLE", "UDIM_PAREN", "PRINTF_PAD", "HOUDINI_PAD", "HASH", 
      "AT", "EXTENSION", "DOT_FRAME_RANGE", "FRAME_RANGE", "DOT_NUM", "SLASH", 
      "SPECIAL_CHAR", "NUM", "WORD", "DASH", "WS", "OTHER_CHAR"
    }
  );
  static const int32_t serializedATNSegment[] = {
  	4,1,17,152,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,
  	7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,
  	14,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,3,0,43,8,0,1,1,1,1,
  	3,1,47,8,1,1,1,1,1,1,1,5,1,52,8,1,10,1,12,1,55,9,1,1,2,1,2,3,2,59,8,2,
  	1,2,1,2,5,2,63,8,2,10,2,12,2,66,9,2,1,3,1,3,1,3,1,3,5,3,72,8,3,10,3,12,
  	3,75,9,3,1,4,1,4,3,4,79,8,4,1,4,5,4,82,8,4,10,4,12,4,85,9,4,1,5,3,5,88,
  	8,5,1,5,1,5,1,5,5,5,93,8,5,10,5,12,5,96,9,5,1,6,4,6,99,8,6,11,6,12,6,
  	100,1,7,4,7,104,8,7,11,7,12,7,105,1,8,4,8,109,8,8,11,8,12,8,110,1,9,4,
  	9,114,8,9,11,9,12,9,115,1,10,4,10,119,8,10,11,10,12,10,120,1,11,1,11,
  	1,12,1,12,1,13,1,13,1,13,1,13,1,13,4,13,132,8,13,11,13,12,13,133,1,13,
  	4,13,137,8,13,11,13,12,13,138,3,13,141,8,13,1,14,1,14,1,14,3,14,146,8,
  	14,1,14,1,14,3,14,150,8,14,1,14,0,0,15,0,2,4,6,8,10,12,14,16,18,20,22,
  	24,26,28,0,4,2,0,8,10,12,17,2,0,7,10,12,17,2,0,8,9,12,17,2,0,8,10,13,
  	13,163,0,42,1,0,0,0,2,44,1,0,0,0,4,56,1,0,0,0,6,67,1,0,0,0,8,76,1,0,0,
  	0,10,87,1,0,0,0,12,98,1,0,0,0,14,103,1,0,0,0,16,108,1,0,0,0,18,113,1,
  	0,0,0,20,118,1,0,0,0,22,122,1,0,0,0,24,124,1,0,0,0,26,140,1,0,0,0,28,
  	149,1,0,0,0,30,31,3,2,1,0,31,32,5,0,0,1,32,43,1,0,0,0,33,34,3,4,2,0,34,
  	35,5,0,0,1,35,43,1,0,0,0,36,37,3,6,3,0,37,38,5,0,0,1,38,43,1,0,0,0,39,
  	40,3,8,4,0,40,41,5,0,0,1,41,43,1,0,0,0,42,30,1,0,0,0,42,33,1,0,0,0,42,
  	36,1,0,0,0,42,39,1,0,0,0,43,1,1,0,0,0,44,46,3,10,5,0,45,47,3,14,7,0,46,
  	45,1,0,0,0,46,47,1,0,0,0,47,48,1,0,0,0,48,49,3,22,11,0,49,53,3,26,13,
  	0,50,52,3,28,14,0,51,50,1,0,0,0,52,55,1,0,0,0,53,51,1,0,0,0,53,54,1,0,
  	0,0,54,3,1,0,0,0,55,53,1,0,0,0,56,58,3,10,5,0,57,59,3,16,8,0,58,57,1,
  	0,0,0,58,59,1,0,0,0,59,60,1,0,0,0,60,64,3,26,13,0,61,63,3,28,14,0,62,
  	61,1,0,0,0,63,66,1,0,0,0,64,62,1,0,0,0,64,65,1,0,0,0,65,5,1,0,0,0,66,
  	64,1,0,0,0,67,68,3,10,5,0,68,69,3,18,9,0,69,73,3,24,12,0,70,72,3,28,14,
  	0,71,70,1,0,0,0,72,75,1,0,0,0,73,71,1,0,0,0,73,74,1,0,0,0,74,7,1,0,0,
  	0,75,73,1,0,0,0,76,78,3,10,5,0,77,79,3,20,10,0,78,77,1,0,0,0,78,79,1,
  	0,0,0,79,83,1,0,0,0,80,82,3,28,14,0,81,80,1,0,0,0,82,85,1,0,0,0,83,81,
  	1,0,0,0,83,84,1,0,0,0,84,9,1,0,0,0,85,83,1,0,0,0,86,88,5,11,0,0,87,86,
  	1,0,0,0,87,88,1,0,0,0,88,94,1,0,0,0,89,90,3,12,6,0,90,91,5,11,0,0,91,
  	93,1,0,0,0,92,89,1,0,0,0,93,96,1,0,0,0,94,92,1,0,0,0,94,95,1,0,0,0,95,
  	11,1,0,0,0,96,94,1,0,0,0,97,99,7,0,0,0,98,97,1,0,0,0,99,100,1,0,0,0,100,
  	98,1,0,0,0,100,101,1,0,0,0,101,13,1,0,0,0,102,104,7,1,0,0,103,102,1,0,
  	0,0,104,105,1,0,0,0,105,103,1,0,0,0,105,106,1,0,0,0,106,15,1,0,0,0,107,
  	109,7,1,0,0,108,107,1,0,0,0,109,110,1,0,0,0,110,108,1,0,0,0,110,111,1,
  	0,0,0,111,17,1,0,0,0,112,114,7,1,0,0,113,112,1,0,0,0,114,115,1,0,0,0,
  	115,113,1,0,0,0,115,116,1,0,0,0,116,19,1,0,0,0,117,119,7,2,0,0,118,117,
  	1,0,0,0,119,120,1,0,0,0,120,118,1,0,0,0,120,121,1,0,0,0,121,21,1,0,0,
  	0,122,123,7,3,0,0,123,23,1,0,0,0,124,125,5,10,0,0,125,25,1,0,0,0,126,
  	141,5,1,0,0,127,141,5,2,0,0,128,141,5,3,0,0,129,141,5,4,0,0,130,132,5,
  	5,0,0,131,130,1,0,0,0,132,133,1,0,0,0,133,131,1,0,0,0,133,134,1,0,0,0,
  	134,141,1,0,0,0,135,137,5,6,0,0,136,135,1,0,0,0,137,138,1,0,0,0,138,136,
  	1,0,0,0,138,139,1,0,0,0,139,141,1,0,0,0,140,126,1,0,0,0,140,127,1,0,0,
  	0,140,128,1,0,0,0,140,129,1,0,0,0,140,131,1,0,0,0,140,136,1,0,0,0,141,
  	27,1,0,0,0,142,145,5,7,0,0,143,144,5,15,0,0,144,146,5,13,0,0,145,143,
  	1,0,0,0,145,146,1,0,0,0,146,150,1,0,0,0,147,150,5,10,0,0,148,150,5,14,
  	0,0,149,142,1,0,0,0,149,147,1,0,0,0,149,148,1,0,0,0,150,29,1,0,0,0,20,
  	42,46,53,58,64,73,78,83,87,94,100,105,110,115,120,133,138,140,145,149
  };
  staticData->serializedATN = antlr4::atn::SerializedATNView(serializedATNSegment, sizeof(serializedATNSegment) / sizeof(serializedATNSegment[0]));

  antlr4::atn::ATNDeserializer deserializer;
  staticData->atn = deserializer.deserialize(staticData->serializedATN);

  const size_t count = staticData->atn->getNumberOfDecisions();
  staticData->decisionToDFA.reserve(count);
  for (size_t i = 0; i < count; i++) { 
    staticData->decisionToDFA.emplace_back(staticData->atn->getDecisionState(i), i);
  }
  fileseqParserStaticData = staticData.release();
}

}

fileseqParser::fileseqParser(TokenStream *input) : fileseqParser(input, antlr4::atn::ParserATNSimulatorOptions()) {}

fileseqParser::fileseqParser(TokenStream *input, const antlr4::atn::ParserATNSimulatorOptions &options) : Parser(input) {
  fileseqParser::initialize();
  _interpreter = new atn::ParserATNSimulator(this, *fileseqParserStaticData->atn, fileseqParserStaticData->decisionToDFA, fileseqParserStaticData->sharedContextCache, options);
}

fileseqParser::~fileseqParser() {
  delete _interpreter;
}

const atn::ATN& fileseqParser::getATN() const {
  return *fileseqParserStaticData->atn;
}

std::string fileseqParser::getGrammarFileName() const {
  return "fileseq.g4";
}

const std::vector<std::string>& fileseqParser::getRuleNames() const {
  return fileseqParserStaticData->ruleNames;
}

const dfa::Vocabulary& fileseqParser::getVocabulary() const {
  return fileseqParserStaticData->vocabulary;
}

antlr4::atn::SerializedATNView fileseqParser::getSerializedATN() const {
  return fileseqParserStaticData->serializedATN;
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


std::any fileseqParser::InputContext::accept(tree::ParseTreeVisitor *visitor) {
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
    setState(42);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 0, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(30);
      sequence();
      setState(31);
      match(fileseqParser::EOF);
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(33);
      patternOnly();
      setState(34);
      match(fileseqParser::EOF);
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);
      setState(36);
      singleFrame();
      setState(37);
      match(fileseqParser::EOF);
      break;
    }

    case 4: {
      enterOuterAlt(_localctx, 4);
      setState(39);
      plainFile();
      setState(40);
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

fileseqParser::FrameRangeContext* fileseqParser::SequenceContext::frameRange() {
  return getRuleContext<fileseqParser::FrameRangeContext>(0);
}

fileseqParser::PaddingContext* fileseqParser::SequenceContext::padding() {
  return getRuleContext<fileseqParser::PaddingContext>(0);
}

fileseqParser::SequenceBasenameContext* fileseqParser::SequenceContext::sequenceBasename() {
  return getRuleContext<fileseqParser::SequenceBasenameContext>(0);
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


std::any fileseqParser::SequenceContext::accept(tree::ParseTreeVisitor *visitor) {
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
    setState(44);
    directory();
    setState(46);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 1, _ctx)) {
    case 1: {
      setState(45);
      sequenceBasename();
      break;
    }

    default:
      break;
    }
    setState(48);
    frameRange();
    setState(49);
    padding();
    setState(53);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 17536) != 0)) {
      setState(50);
      extension();
      setState(55);
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

fileseqParser::PaddingContext* fileseqParser::PatternOnlyContext::padding() {
  return getRuleContext<fileseqParser::PaddingContext>(0);
}

fileseqParser::PatternBasenameContext* fileseqParser::PatternOnlyContext::patternBasename() {
  return getRuleContext<fileseqParser::PatternBasenameContext>(0);
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


std::any fileseqParser::PatternOnlyContext::accept(tree::ParseTreeVisitor *visitor) {
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
    setState(56);
    directory();
    setState(58);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 259968) != 0)) {
      setState(57);
      patternBasename();
    }
    setState(60);
    padding();
    setState(64);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 17536) != 0)) {
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

//----------------- SingleFrameContext ------------------------------------------------------------------

fileseqParser::SingleFrameContext::SingleFrameContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

fileseqParser::DirectoryContext* fileseqParser::SingleFrameContext::directory() {
  return getRuleContext<fileseqParser::DirectoryContext>(0);
}

fileseqParser::SingleFrameBasenameContext* fileseqParser::SingleFrameContext::singleFrameBasename() {
  return getRuleContext<fileseqParser::SingleFrameBasenameContext>(0);
}

fileseqParser::FrameNumContext* fileseqParser::SingleFrameContext::frameNum() {
  return getRuleContext<fileseqParser::FrameNumContext>(0);
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


std::any fileseqParser::SingleFrameContext::accept(tree::ParseTreeVisitor *visitor) {
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
    setState(67);
    directory();
    setState(68);
    singleFrameBasename();
    setState(69);
    frameNum();
    setState(73);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 17536) != 0)) {
      setState(70);
      extension();
      setState(75);
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


std::any fileseqParser::PlainFileContext::accept(tree::ParseTreeVisitor *visitor) {
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
    setState(76);
    directory();
    setState(78);
    _errHandler->sync(this);

    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 6, _ctx)) {
    case 1: {
      setState(77);
      plainBasename();
      break;
    }

    default:
      break;
    }
    setState(83);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 17536) != 0)) {
      setState(80);
      extension();
      setState(85);
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


std::any fileseqParser::DirectoryContext::accept(tree::ParseTreeVisitor *visitor) {
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
    setState(87);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == fileseqParser::SLASH) {
      setState(86);
      match(fileseqParser::SLASH);
    }
    setState(94);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 9, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        setState(89);
        dirSegment();
        setState(90);
        match(fileseqParser::SLASH); 
      }
      setState(96);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 9, _ctx);
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

std::vector<tree::TerminalNode *> fileseqParser::DirSegmentContext::WORD() {
  return getTokens(fileseqParser::WORD);
}

tree::TerminalNode* fileseqParser::DirSegmentContext::WORD(size_t i) {
  return getToken(fileseqParser::WORD, i);
}

std::vector<tree::TerminalNode *> fileseqParser::DirSegmentContext::NUM() {
  return getTokens(fileseqParser::NUM);
}

tree::TerminalNode* fileseqParser::DirSegmentContext::NUM(size_t i) {
  return getToken(fileseqParser::NUM, i);
}

std::vector<tree::TerminalNode *> fileseqParser::DirSegmentContext::DASH() {
  return getTokens(fileseqParser::DASH);
}

tree::TerminalNode* fileseqParser::DirSegmentContext::DASH(size_t i) {
  return getToken(fileseqParser::DASH, i);
}

std::vector<tree::TerminalNode *> fileseqParser::DirSegmentContext::SPECIAL_CHAR() {
  return getTokens(fileseqParser::SPECIAL_CHAR);
}

tree::TerminalNode* fileseqParser::DirSegmentContext::SPECIAL_CHAR(size_t i) {
  return getToken(fileseqParser::SPECIAL_CHAR, i);
}

std::vector<tree::TerminalNode *> fileseqParser::DirSegmentContext::FRAME_RANGE() {
  return getTokens(fileseqParser::FRAME_RANGE);
}

tree::TerminalNode* fileseqParser::DirSegmentContext::FRAME_RANGE(size_t i) {
  return getToken(fileseqParser::FRAME_RANGE, i);
}

std::vector<tree::TerminalNode *> fileseqParser::DirSegmentContext::DOT_FRAME_RANGE() {
  return getTokens(fileseqParser::DOT_FRAME_RANGE);
}

tree::TerminalNode* fileseqParser::DirSegmentContext::DOT_FRAME_RANGE(size_t i) {
  return getToken(fileseqParser::DOT_FRAME_RANGE, i);
}

std::vector<tree::TerminalNode *> fileseqParser::DirSegmentContext::DOT_NUM() {
  return getTokens(fileseqParser::DOT_NUM);
}

tree::TerminalNode* fileseqParser::DirSegmentContext::DOT_NUM(size_t i) {
  return getToken(fileseqParser::DOT_NUM, i);
}

std::vector<tree::TerminalNode *> fileseqParser::DirSegmentContext::WS() {
  return getTokens(fileseqParser::WS);
}

tree::TerminalNode* fileseqParser::DirSegmentContext::WS(size_t i) {
  return getToken(fileseqParser::WS, i);
}

std::vector<tree::TerminalNode *> fileseqParser::DirSegmentContext::OTHER_CHAR() {
  return getTokens(fileseqParser::OTHER_CHAR);
}

tree::TerminalNode* fileseqParser::DirSegmentContext::OTHER_CHAR(size_t i) {
  return getToken(fileseqParser::OTHER_CHAR, i);
}


size_t fileseqParser::DirSegmentContext::getRuleIndex() const {
  return fileseqParser::RuleDirSegment;
}


std::any fileseqParser::DirSegmentContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitDirSegment(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::DirSegmentContext* fileseqParser::dirSegment() {
  DirSegmentContext *_localctx = _tracker.createInstance<DirSegmentContext>(_ctx, getState());
  enterRule(_localctx, 12, fileseqParser::RuleDirSegment);
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
    setState(98); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(97);
      _la = _input->LA(1);
      if (!((((_la & ~ 0x3fULL) == 0) &&
        ((1ULL << _la) & 259840) != 0))) {
      _errHandler->recoverInline(this);
      }
      else {
        _errHandler->reportMatch(this);
        consume();
      }
      setState(100); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 259840) != 0));
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- SequenceBasenameContext ------------------------------------------------------------------

fileseqParser::SequenceBasenameContext::SequenceBasenameContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<tree::TerminalNode *> fileseqParser::SequenceBasenameContext::WORD() {
  return getTokens(fileseqParser::WORD);
}

tree::TerminalNode* fileseqParser::SequenceBasenameContext::WORD(size_t i) {
  return getToken(fileseqParser::WORD, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SequenceBasenameContext::NUM() {
  return getTokens(fileseqParser::NUM);
}

tree::TerminalNode* fileseqParser::SequenceBasenameContext::NUM(size_t i) {
  return getToken(fileseqParser::NUM, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SequenceBasenameContext::DOT_NUM() {
  return getTokens(fileseqParser::DOT_NUM);
}

tree::TerminalNode* fileseqParser::SequenceBasenameContext::DOT_NUM(size_t i) {
  return getToken(fileseqParser::DOT_NUM, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SequenceBasenameContext::DASH() {
  return getTokens(fileseqParser::DASH);
}

tree::TerminalNode* fileseqParser::SequenceBasenameContext::DASH(size_t i) {
  return getToken(fileseqParser::DASH, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SequenceBasenameContext::SPECIAL_CHAR() {
  return getTokens(fileseqParser::SPECIAL_CHAR);
}

tree::TerminalNode* fileseqParser::SequenceBasenameContext::SPECIAL_CHAR(size_t i) {
  return getToken(fileseqParser::SPECIAL_CHAR, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SequenceBasenameContext::EXTENSION() {
  return getTokens(fileseqParser::EXTENSION);
}

tree::TerminalNode* fileseqParser::SequenceBasenameContext::EXTENSION(size_t i) {
  return getToken(fileseqParser::EXTENSION, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SequenceBasenameContext::FRAME_RANGE() {
  return getTokens(fileseqParser::FRAME_RANGE);
}

tree::TerminalNode* fileseqParser::SequenceBasenameContext::FRAME_RANGE(size_t i) {
  return getToken(fileseqParser::FRAME_RANGE, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SequenceBasenameContext::DOT_FRAME_RANGE() {
  return getTokens(fileseqParser::DOT_FRAME_RANGE);
}

tree::TerminalNode* fileseqParser::SequenceBasenameContext::DOT_FRAME_RANGE(size_t i) {
  return getToken(fileseqParser::DOT_FRAME_RANGE, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SequenceBasenameContext::WS() {
  return getTokens(fileseqParser::WS);
}

tree::TerminalNode* fileseqParser::SequenceBasenameContext::WS(size_t i) {
  return getToken(fileseqParser::WS, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SequenceBasenameContext::OTHER_CHAR() {
  return getTokens(fileseqParser::OTHER_CHAR);
}

tree::TerminalNode* fileseqParser::SequenceBasenameContext::OTHER_CHAR(size_t i) {
  return getToken(fileseqParser::OTHER_CHAR, i);
}


size_t fileseqParser::SequenceBasenameContext::getRuleIndex() const {
  return fileseqParser::RuleSequenceBasename;
}


std::any fileseqParser::SequenceBasenameContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitSequenceBasename(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::SequenceBasenameContext* fileseqParser::sequenceBasename() {
  SequenceBasenameContext *_localctx = _tracker.createInstance<SequenceBasenameContext>(_ctx, getState());
  enterRule(_localctx, 14, fileseqParser::RuleSequenceBasename);
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
    setState(103); 
    _errHandler->sync(this);
    alt = 1;
    do {
      switch (alt) {
        case 1: {
              setState(102);
              _la = _input->LA(1);
              if (!((((_la & ~ 0x3fULL) == 0) &&
                ((1ULL << _la) & 259968) != 0))) {
              _errHandler->recoverInline(this);
              }
              else {
                _errHandler->reportMatch(this);
                consume();
              }
              break;
            }

      default:
        throw NoViableAltException(this);
      }
      setState(105); 
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 11, _ctx);
    } while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- PatternBasenameContext ------------------------------------------------------------------

fileseqParser::PatternBasenameContext::PatternBasenameContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<tree::TerminalNode *> fileseqParser::PatternBasenameContext::WORD() {
  return getTokens(fileseqParser::WORD);
}

tree::TerminalNode* fileseqParser::PatternBasenameContext::WORD(size_t i) {
  return getToken(fileseqParser::WORD, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PatternBasenameContext::NUM() {
  return getTokens(fileseqParser::NUM);
}

tree::TerminalNode* fileseqParser::PatternBasenameContext::NUM(size_t i) {
  return getToken(fileseqParser::NUM, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PatternBasenameContext::DOT_NUM() {
  return getTokens(fileseqParser::DOT_NUM);
}

tree::TerminalNode* fileseqParser::PatternBasenameContext::DOT_NUM(size_t i) {
  return getToken(fileseqParser::DOT_NUM, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PatternBasenameContext::DASH() {
  return getTokens(fileseqParser::DASH);
}

tree::TerminalNode* fileseqParser::PatternBasenameContext::DASH(size_t i) {
  return getToken(fileseqParser::DASH, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PatternBasenameContext::SPECIAL_CHAR() {
  return getTokens(fileseqParser::SPECIAL_CHAR);
}

tree::TerminalNode* fileseqParser::PatternBasenameContext::SPECIAL_CHAR(size_t i) {
  return getToken(fileseqParser::SPECIAL_CHAR, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PatternBasenameContext::EXTENSION() {
  return getTokens(fileseqParser::EXTENSION);
}

tree::TerminalNode* fileseqParser::PatternBasenameContext::EXTENSION(size_t i) {
  return getToken(fileseqParser::EXTENSION, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PatternBasenameContext::FRAME_RANGE() {
  return getTokens(fileseqParser::FRAME_RANGE);
}

tree::TerminalNode* fileseqParser::PatternBasenameContext::FRAME_RANGE(size_t i) {
  return getToken(fileseqParser::FRAME_RANGE, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PatternBasenameContext::DOT_FRAME_RANGE() {
  return getTokens(fileseqParser::DOT_FRAME_RANGE);
}

tree::TerminalNode* fileseqParser::PatternBasenameContext::DOT_FRAME_RANGE(size_t i) {
  return getToken(fileseqParser::DOT_FRAME_RANGE, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PatternBasenameContext::WS() {
  return getTokens(fileseqParser::WS);
}

tree::TerminalNode* fileseqParser::PatternBasenameContext::WS(size_t i) {
  return getToken(fileseqParser::WS, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PatternBasenameContext::OTHER_CHAR() {
  return getTokens(fileseqParser::OTHER_CHAR);
}

tree::TerminalNode* fileseqParser::PatternBasenameContext::OTHER_CHAR(size_t i) {
  return getToken(fileseqParser::OTHER_CHAR, i);
}


size_t fileseqParser::PatternBasenameContext::getRuleIndex() const {
  return fileseqParser::RulePatternBasename;
}


std::any fileseqParser::PatternBasenameContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitPatternBasename(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::PatternBasenameContext* fileseqParser::patternBasename() {
  PatternBasenameContext *_localctx = _tracker.createInstance<PatternBasenameContext>(_ctx, getState());
  enterRule(_localctx, 16, fileseqParser::RulePatternBasename);
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
    setState(108); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(107);
      _la = _input->LA(1);
      if (!((((_la & ~ 0x3fULL) == 0) &&
        ((1ULL << _la) & 259968) != 0))) {
      _errHandler->recoverInline(this);
      }
      else {
        _errHandler->reportMatch(this);
        consume();
      }
      setState(110); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 259968) != 0));
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- SingleFrameBasenameContext ------------------------------------------------------------------

fileseqParser::SingleFrameBasenameContext::SingleFrameBasenameContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<tree::TerminalNode *> fileseqParser::SingleFrameBasenameContext::WORD() {
  return getTokens(fileseqParser::WORD);
}

tree::TerminalNode* fileseqParser::SingleFrameBasenameContext::WORD(size_t i) {
  return getToken(fileseqParser::WORD, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SingleFrameBasenameContext::NUM() {
  return getTokens(fileseqParser::NUM);
}

tree::TerminalNode* fileseqParser::SingleFrameBasenameContext::NUM(size_t i) {
  return getToken(fileseqParser::NUM, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SingleFrameBasenameContext::DOT_NUM() {
  return getTokens(fileseqParser::DOT_NUM);
}

tree::TerminalNode* fileseqParser::SingleFrameBasenameContext::DOT_NUM(size_t i) {
  return getToken(fileseqParser::DOT_NUM, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SingleFrameBasenameContext::DASH() {
  return getTokens(fileseqParser::DASH);
}

tree::TerminalNode* fileseqParser::SingleFrameBasenameContext::DASH(size_t i) {
  return getToken(fileseqParser::DASH, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SingleFrameBasenameContext::SPECIAL_CHAR() {
  return getTokens(fileseqParser::SPECIAL_CHAR);
}

tree::TerminalNode* fileseqParser::SingleFrameBasenameContext::SPECIAL_CHAR(size_t i) {
  return getToken(fileseqParser::SPECIAL_CHAR, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SingleFrameBasenameContext::EXTENSION() {
  return getTokens(fileseqParser::EXTENSION);
}

tree::TerminalNode* fileseqParser::SingleFrameBasenameContext::EXTENSION(size_t i) {
  return getToken(fileseqParser::EXTENSION, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SingleFrameBasenameContext::FRAME_RANGE() {
  return getTokens(fileseqParser::FRAME_RANGE);
}

tree::TerminalNode* fileseqParser::SingleFrameBasenameContext::FRAME_RANGE(size_t i) {
  return getToken(fileseqParser::FRAME_RANGE, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SingleFrameBasenameContext::DOT_FRAME_RANGE() {
  return getTokens(fileseqParser::DOT_FRAME_RANGE);
}

tree::TerminalNode* fileseqParser::SingleFrameBasenameContext::DOT_FRAME_RANGE(size_t i) {
  return getToken(fileseqParser::DOT_FRAME_RANGE, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SingleFrameBasenameContext::WS() {
  return getTokens(fileseqParser::WS);
}

tree::TerminalNode* fileseqParser::SingleFrameBasenameContext::WS(size_t i) {
  return getToken(fileseqParser::WS, i);
}

std::vector<tree::TerminalNode *> fileseqParser::SingleFrameBasenameContext::OTHER_CHAR() {
  return getTokens(fileseqParser::OTHER_CHAR);
}

tree::TerminalNode* fileseqParser::SingleFrameBasenameContext::OTHER_CHAR(size_t i) {
  return getToken(fileseqParser::OTHER_CHAR, i);
}


size_t fileseqParser::SingleFrameBasenameContext::getRuleIndex() const {
  return fileseqParser::RuleSingleFrameBasename;
}


std::any fileseqParser::SingleFrameBasenameContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitSingleFrameBasename(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::SingleFrameBasenameContext* fileseqParser::singleFrameBasename() {
  SingleFrameBasenameContext *_localctx = _tracker.createInstance<SingleFrameBasenameContext>(_ctx, getState());
  enterRule(_localctx, 18, fileseqParser::RuleSingleFrameBasename);
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
    setState(113); 
    _errHandler->sync(this);
    alt = 1;
    do {
      switch (alt) {
        case 1: {
              setState(112);
              _la = _input->LA(1);
              if (!((((_la & ~ 0x3fULL) == 0) &&
                ((1ULL << _la) & 259968) != 0))) {
              _errHandler->recoverInline(this);
              }
              else {
                _errHandler->reportMatch(this);
                consume();
              }
              break;
            }

      default:
        throw NoViableAltException(this);
      }
      setState(115); 
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 13, _ctx);
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

std::vector<tree::TerminalNode *> fileseqParser::PlainBasenameContext::WORD() {
  return getTokens(fileseqParser::WORD);
}

tree::TerminalNode* fileseqParser::PlainBasenameContext::WORD(size_t i) {
  return getToken(fileseqParser::WORD, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PlainBasenameContext::NUM() {
  return getTokens(fileseqParser::NUM);
}

tree::TerminalNode* fileseqParser::PlainBasenameContext::NUM(size_t i) {
  return getToken(fileseqParser::NUM, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PlainBasenameContext::DASH() {
  return getTokens(fileseqParser::DASH);
}

tree::TerminalNode* fileseqParser::PlainBasenameContext::DASH(size_t i) {
  return getToken(fileseqParser::DASH, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PlainBasenameContext::SPECIAL_CHAR() {
  return getTokens(fileseqParser::SPECIAL_CHAR);
}

tree::TerminalNode* fileseqParser::PlainBasenameContext::SPECIAL_CHAR(size_t i) {
  return getToken(fileseqParser::SPECIAL_CHAR, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PlainBasenameContext::FRAME_RANGE() {
  return getTokens(fileseqParser::FRAME_RANGE);
}

tree::TerminalNode* fileseqParser::PlainBasenameContext::FRAME_RANGE(size_t i) {
  return getToken(fileseqParser::FRAME_RANGE, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PlainBasenameContext::DOT_FRAME_RANGE() {
  return getTokens(fileseqParser::DOT_FRAME_RANGE);
}

tree::TerminalNode* fileseqParser::PlainBasenameContext::DOT_FRAME_RANGE(size_t i) {
  return getToken(fileseqParser::DOT_FRAME_RANGE, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PlainBasenameContext::WS() {
  return getTokens(fileseqParser::WS);
}

tree::TerminalNode* fileseqParser::PlainBasenameContext::WS(size_t i) {
  return getToken(fileseqParser::WS, i);
}

std::vector<tree::TerminalNode *> fileseqParser::PlainBasenameContext::OTHER_CHAR() {
  return getTokens(fileseqParser::OTHER_CHAR);
}

tree::TerminalNode* fileseqParser::PlainBasenameContext::OTHER_CHAR(size_t i) {
  return getToken(fileseqParser::OTHER_CHAR, i);
}


size_t fileseqParser::PlainBasenameContext::getRuleIndex() const {
  return fileseqParser::RulePlainBasename;
}


std::any fileseqParser::PlainBasenameContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitPlainBasename(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::PlainBasenameContext* fileseqParser::plainBasename() {
  PlainBasenameContext *_localctx = _tracker.createInstance<PlainBasenameContext>(_ctx, getState());
  enterRule(_localctx, 20, fileseqParser::RulePlainBasename);
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
    setState(118); 
    _errHandler->sync(this);
    alt = 1;
    do {
      switch (alt) {
        case 1: {
              setState(117);
              _la = _input->LA(1);
              if (!((((_la & ~ 0x3fULL) == 0) &&
                ((1ULL << _la) & 258816) != 0))) {
              _errHandler->recoverInline(this);
              }
              else {
                _errHandler->reportMatch(this);
                consume();
              }
              break;
            }

      default:
        throw NoViableAltException(this);
      }
      setState(120); 
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 14, _ctx);
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


std::any fileseqParser::FrameRangeContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitFrameRange(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::FrameRangeContext* fileseqParser::frameRange() {
  FrameRangeContext *_localctx = _tracker.createInstance<FrameRangeContext>(_ctx, getState());
  enterRule(_localctx, 22, fileseqParser::RuleFrameRange);
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
    setState(122);
    _la = _input->LA(1);
    if (!((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & 9984) != 0))) {
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


std::any fileseqParser::FrameNumContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitFrameNum(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::FrameNumContext* fileseqParser::frameNum() {
  FrameNumContext *_localctx = _tracker.createInstance<FrameNumContext>(_ctx, getState());
  enterRule(_localctx, 24, fileseqParser::RuleFrameNum);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(124);
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


std::any fileseqParser::PaddingContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitPadding(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::PaddingContext* fileseqParser::padding() {
  PaddingContext *_localctx = _tracker.createInstance<PaddingContext>(_ctx, getState());
  enterRule(_localctx, 26, fileseqParser::RulePadding);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(140);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case fileseqParser::UDIM_ANGLE: {
        enterOuterAlt(_localctx, 1);
        setState(126);
        match(fileseqParser::UDIM_ANGLE);
        break;
      }

      case fileseqParser::UDIM_PAREN: {
        enterOuterAlt(_localctx, 2);
        setState(127);
        match(fileseqParser::UDIM_PAREN);
        break;
      }

      case fileseqParser::PRINTF_PAD: {
        enterOuterAlt(_localctx, 3);
        setState(128);
        match(fileseqParser::PRINTF_PAD);
        break;
      }

      case fileseqParser::HOUDINI_PAD: {
        enterOuterAlt(_localctx, 4);
        setState(129);
        match(fileseqParser::HOUDINI_PAD);
        break;
      }

      case fileseqParser::HASH: {
        enterOuterAlt(_localctx, 5);
        setState(131); 
        _errHandler->sync(this);
        _la = _input->LA(1);
        do {
          setState(130);
          match(fileseqParser::HASH);
          setState(133); 
          _errHandler->sync(this);
          _la = _input->LA(1);
        } while (_la == fileseqParser::HASH);
        break;
      }

      case fileseqParser::AT: {
        enterOuterAlt(_localctx, 6);
        setState(136); 
        _errHandler->sync(this);
        _la = _input->LA(1);
        do {
          setState(135);
          match(fileseqParser::AT);
          setState(138); 
          _errHandler->sync(this);
          _la = _input->LA(1);
        } while (_la == fileseqParser::AT);
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


std::any fileseqParser::ExtensionContext::accept(tree::ParseTreeVisitor *visitor) {
  if (auto parserVisitor = dynamic_cast<fileseqVisitor*>(visitor))
    return parserVisitor->visitExtension(this);
  else
    return visitor->visitChildren(this);
}

fileseqParser::ExtensionContext* fileseqParser::extension() {
  ExtensionContext *_localctx = _tracker.createInstance<ExtensionContext>(_ctx, getState());
  enterRule(_localctx, 28, fileseqParser::RuleExtension);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(149);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case fileseqParser::EXTENSION: {
        enterOuterAlt(_localctx, 1);
        setState(142);
        match(fileseqParser::EXTENSION);
        setState(145);
        _errHandler->sync(this);

        _la = _input->LA(1);
        if (_la == fileseqParser::DASH) {
          setState(143);
          match(fileseqParser::DASH);
          setState(144);
          match(fileseqParser::NUM);
        }
        break;
      }

      case fileseqParser::DOT_NUM: {
        enterOuterAlt(_localctx, 2);
        setState(147);
        match(fileseqParser::DOT_NUM);
        break;
      }

      case fileseqParser::WORD: {
        enterOuterAlt(_localctx, 3);
        setState(148);
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

void fileseqParser::initialize() {
#if ANTLR4_USE_THREAD_LOCAL_CACHE
  fileseqParserInitialize();
#else
  ::antlr4::internal::call_once(fileseqParserOnceFlag, fileseqParserInitialize);
#endif
}
