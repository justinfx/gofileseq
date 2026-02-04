
// Generated from fileseq.g4 by ANTLR 4.9.3

#pragma once


#include "antlr4-runtime.h"


namespace fileseq {


class  fileseqLexer : public antlr4::Lexer {
public:
  enum {
    UDIM_ANGLE = 1, UDIM_PAREN = 2, PRINTF_PAD = 3, HOUDINI_PAD = 4, HASH = 5, 
    AT = 6, EXTENSION = 7, DOT_FRAME_RANGE = 8, FRAME_RANGE = 9, DOT_NUM = 10, 
    SLASH = 11, SPECIAL_CHAR = 12, NUM = 13, WORD = 14, DASH = 15, WS = 16, 
    OTHER_CHAR = 17
  };

  explicit fileseqLexer(antlr4::CharStream *input);
  ~fileseqLexer();

  virtual std::string getGrammarFileName() const override;
  virtual const std::vector<std::string>& getRuleNames() const override;

  virtual const std::vector<std::string>& getChannelNames() const override;
  virtual const std::vector<std::string>& getModeNames() const override;
  virtual const std::vector<std::string>& getTokenNames() const override; // deprecated, use vocabulary instead
  virtual antlr4::dfa::Vocabulary& getVocabulary() const override;

  virtual const std::vector<uint16_t> getSerializedATN() const override;
  virtual const antlr4::atn::ATN& getATN() const override;

private:
  static std::vector<antlr4::dfa::DFA> _decisionToDFA;
  static antlr4::atn::PredictionContextCache _sharedContextCache;
  static std::vector<std::string> _ruleNames;
  static std::vector<std::string> _tokenNames;
  static std::vector<std::string> _channelNames;
  static std::vector<std::string> _modeNames;

  static std::vector<std::string> _literalNames;
  static std::vector<std::string> _symbolicNames;
  static antlr4::dfa::Vocabulary _vocabulary;
  static antlr4::atn::ATN _atn;
  static std::vector<uint16_t> _serializedATN;


  // Individual action functions triggered by action() above.

  // Individual semantic predicate functions triggered by sempred() above.

  struct Initializer {
    Initializer();
  };
  static Initializer _init;
};

}  // namespace fileseq
