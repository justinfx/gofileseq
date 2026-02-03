
// Generated from fileseq.g4 by ANTLR 4.13.1

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

  ~fileseqLexer() override;


  std::string getGrammarFileName() const override;

  const std::vector<std::string>& getRuleNames() const override;

  const std::vector<std::string>& getChannelNames() const override;

  const std::vector<std::string>& getModeNames() const override;

  const antlr4::dfa::Vocabulary& getVocabulary() const override;

  antlr4::atn::SerializedATNView getSerializedATN() const override;

  const antlr4::atn::ATN& getATN() const override;

  // By default the static state used to implement the lexer is lazily initialized during the first
  // call to the constructor. You can call this function if you wish to initialize the static state
  // ahead of time.
  static void initialize();

private:

  // Individual action functions triggered by action() above.

  // Individual semantic predicate functions triggered by sempred() above.

};

}  // namespace fileseq
