package parser

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

// sllInputs are valid inputs that must parse with SLL prediction alone.
// A grammar change that makes any of these fall back to full LL prediction
// is a large performance regression (roughly 30x for range sequences).
var sllInputs = []string{
	"/mnt/projects/show/shots/sh010/comp/v012/sh010_comp_v012.1001-1240#.exr",
	"/mnt/projects/show/shots/sh010/plates/sh010_plate.1001.dpx",
	"/renders/beauty/beauty.1-100x2,150-200@@@@.tif",
	"/mnt/projects/show/assets/chars/hero/model/hero_model_v003.abc",
	"/path/file.1-100#.exr",
	"/path/file1-100#.exr",
	"/path/file.100.exr",
	"/path/file.txt",
	"/path/to/file:1-1x1#.exr",
	"/path/file.-10-100#.exr",
	"/path/file.1-100x0.25#.exr",
	"/path/file.1,2,3,5-10@@.exr",
	"/path/file.#.exr",
	"/path/file.@@@@.exr",
	"/path/file.%04d.exr",
	"/path/file.$F4.exr",
	"/path/tex.<UDIM>.exr",
	"/path/tex.%(UDIM)d.exr",
	"/path/.hidden5.1-10#.7zip",
	"/path/.hidden.100.ext",
	"/path/1-100#.exr",
	"/path/@@@.exr",
	"/path/name_2025-05-13_v1.1-10#.exr",
	"/path/file.1-10#_exr",
	"/path/file.1-10#.tar.gz",
	"//server/share/file.1-10#.exr",
	"C:\\path\\file.1-10#.exr",
	"relative/file.1001-1100#.exr",
	"file.1-10#.exr",
}

func parsesWithSLL(input string) []error {
	lexer := NewfileseqLexer(antlr.NewInputStream(input))
	parser := NewfileseqParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	errs := &errorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errs)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errs)
	parser.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	parser.Input()
	return errs.errors
}

func TestValidInputsParseWithSLL(t *testing.T) {
	for _, input := range sllInputs {
		if errs := parsesWithSLL(input); len(errs) > 0 {
			t.Errorf("%q needed the LL fallback: %v", input, errs[0])
		}
	}
}
