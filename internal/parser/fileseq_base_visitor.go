// Code generated from fileseq.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // fileseq
import "github.com/antlr4-go/antlr/v4"

type BasefileseqVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasefileseqVisitor) VisitInput(ctx *InputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitSequence(ctx *SequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitPatternOnly(ctx *PatternOnlyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitSingleFrame(ctx *SingleFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitPlainFile(ctx *PlainFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitDirectory(ctx *DirectoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitDirSegment(ctx *DirSegmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitSequenceBasename(ctx *SequenceBasenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitPatternBasename(ctx *PatternBasenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitSingleFrameBasename(ctx *SingleFrameBasenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitPlainBasename(ctx *PlainBasenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitFrameRange(ctx *FrameRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitFrameNum(ctx *FrameNumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitPadding(ctx *PaddingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefileseqVisitor) VisitExtension(ctx *ExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}
