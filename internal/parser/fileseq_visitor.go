// Code generated from fileseq.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // fileseq
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by fileseqParser.
type fileseqVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by fileseqParser#input.
	VisitInput(ctx *InputContext) interface{}

	// Visit a parse tree produced by fileseqParser#sequence.
	VisitSequence(ctx *SequenceContext) interface{}

	// Visit a parse tree produced by fileseqParser#patternOnly.
	VisitPatternOnly(ctx *PatternOnlyContext) interface{}

	// Visit a parse tree produced by fileseqParser#singleFrame.
	VisitSingleFrame(ctx *SingleFrameContext) interface{}

	// Visit a parse tree produced by fileseqParser#plainFile.
	VisitPlainFile(ctx *PlainFileContext) interface{}

	// Visit a parse tree produced by fileseqParser#directory.
	VisitDirectory(ctx *DirectoryContext) interface{}

	// Visit a parse tree produced by fileseqParser#dirSegment.
	VisitDirSegment(ctx *DirSegmentContext) interface{}

	// Visit a parse tree produced by fileseqParser#sequenceBasename.
	VisitSequenceBasename(ctx *SequenceBasenameContext) interface{}

	// Visit a parse tree produced by fileseqParser#patternBasename.
	VisitPatternBasename(ctx *PatternBasenameContext) interface{}

	// Visit a parse tree produced by fileseqParser#singleFrameBasename.
	VisitSingleFrameBasename(ctx *SingleFrameBasenameContext) interface{}

	// Visit a parse tree produced by fileseqParser#plainBasename.
	VisitPlainBasename(ctx *PlainBasenameContext) interface{}

	// Visit a parse tree produced by fileseqParser#frameRange.
	VisitFrameRange(ctx *FrameRangeContext) interface{}

	// Visit a parse tree produced by fileseqParser#frameNum.
	VisitFrameNum(ctx *FrameNumContext) interface{}

	// Visit a parse tree produced by fileseqParser#padding.
	VisitPadding(ctx *PaddingContext) interface{}

	// Visit a parse tree produced by fileseqParser#extension.
	VisitExtension(ctx *ExtensionContext) interface{}
}
