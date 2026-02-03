package parser

import (
	"testing"
)

func TestParseFileSequence(t *testing.T) {
	tests := []struct {
		input       string
		wantDir     string
		wantBase    string
		wantRange   string
		wantPad     string
		wantExt     string
		wantSeq     bool
		wantSingle  bool
		wantPlain   bool
	}{
		// Basic sequence with dot before frame range
		{
			input:     "/path/file.1-100#.exr",
			wantDir:   "/path/",
			wantBase:  "file.",
			wantRange: "1-100",
			wantPad:   "#",
			wantExt:   ".exr",
			wantSeq:   true,
		},
		// Sequence without dot before frame range
		{
			input:     "/path/file1-100#.exr",
			wantDir:   "/path/",
			wantBase:  "file",
			wantRange: "1-100",
			wantPad:   "#",
			wantExt:   ".exr",
			wantSeq:   true,
		},
		// Single frame file
		{
			input:      "/path/file.100.exr",
			wantDir:    "/path/",
			wantBase:   "file.",
			wantRange:  "100",
			wantPad:    "",
			wantExt:    ".exr",
			wantSingle: true,
		},
		// Plain file
		{
			input:     "/path/file.txt",
			wantDir:   "/path/",
			wantBase:  "file",
			wantExt:   ".txt",
			wantPlain: true,
		},
		// Basename with special chars (no dot before frame range)
		{
			input:     "/path/to/file:1-1x1#.exr",
			wantDir:   "/path/to/",
			wantBase:  "file:",
			wantRange: "1-1x1",
			wantPad:   "#",
			wantExt:   ".exr",
			wantSeq:   true,
		},
		// Basename with colon and underscore (no dot before frame range)
		{
			input:     "/path/to/file:_1-100x10@@.exr",
			wantDir:   "/path/to/",
			wantBase:  "file:_",
			wantRange: "1-100x10",
			wantPad:   "@@",
			wantExt:   ".exr",
			wantSeq:   true,
		},
		// Negative frame range with dot
		{
			input:     "/path/to/file.-10--1x2##.exr",
			wantDir:   "/path/to/",
			wantBase:  "file.",
			wantRange: "-10--1x2",
			wantPad:   "##",
			wantExt:   ".exr",
			wantSeq:   true,
		},
		// Multiple hash padding with dot
		{
			input:     "/path/file.1-100##.exr",
			wantDir:   "/path/",
			wantBase:  "file.",
			wantRange: "1-100",
			wantPad:   "##",
			wantExt:   ".exr",
			wantSeq:   true,
		},
		// Printf padding with dot
		{
			input:     "/path/file.1-100%04d.tif",
			wantDir:   "/path/",
			wantBase:  "file.",
			wantRange: "1-100",
			wantPad:   "%04d",
			wantExt:   ".tif",
			wantSeq:   true,
		},
		// Houdini padding with dot
		{
			input:     "/path/file.1-100$F4.exr",
			wantDir:   "/path/",
			wantBase:  "file.",
			wantRange: "1-100",
			wantPad:   "$F4",
			wantExt:   ".exr",
			wantSeq:   true,
		},
		// UDIM with dot
		{
			input:     "/path/file.1-100<UDIM>.tif",
			wantDir:   "/path/",
			wantBase:  "file.",
			wantRange: "1-100",
			wantPad:   "<UDIM>",
			wantExt:   ".tif",
			wantSeq:   true,
		},
		// Multi-part extension with dot
		{
			input:     "/path/file.1-100#.tar.gz",
			wantDir:   "/path/",
			wantBase:  "file.",
			wantRange: "1-100",
			wantPad:   "#",
			wantExt:   ".tar.gz",
			wantSeq:   true,
		},
		// Comma-separated frames with dot
		{
			input:     "/path/file.1,2,3,5-10,20-30#.exr",
			wantDir:   "/path/",
			wantBase:  "file.",
			wantRange: "1,2,3,5-10,20-30",
			wantPad:   "#",
			wantExt:   ".exr",
			wantSeq:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseFileSequence(tt.input)
			if err != nil {
				t.Fatalf("ParseFileSequence() error = %v", err)
			}

			if result.Directory != tt.wantDir {
				t.Errorf("Directory = %q, want %q", result.Directory, tt.wantDir)
			}
			if result.Basename != tt.wantBase {
				t.Errorf("Basename = %q, want %q", result.Basename, tt.wantBase)
			}
			if result.FrameRange != tt.wantRange {
				t.Errorf("FrameRange = %q, want %q", result.FrameRange, tt.wantRange)
			}
			if result.Padding != tt.wantPad {
				t.Errorf("Padding = %q, want %q", result.Padding, tt.wantPad)
			}
			if result.Extension != tt.wantExt {
				t.Errorf("Extension = %q, want %q", result.Extension, tt.wantExt)
			}
			if result.IsSequence != tt.wantSeq {
				t.Errorf("IsSequence = %v, want %v", result.IsSequence, tt.wantSeq)
			}
			if result.IsSingleFile != tt.wantSingle {
				t.Errorf("IsSingleFile = %v, want %v", result.IsSingleFile, tt.wantSingle)
			}
			if result.IsPlainFile != tt.wantPlain {
				t.Errorf("IsPlainFile = %v, want %v", result.IsPlainFile, tt.wantPlain)
			}
		})
	}
}
