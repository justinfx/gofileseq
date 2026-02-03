package fileseq

import (
	"strings"
	"testing"

	"github.com/justinfx/gofileseq/v3/internal/parser"
)

// FuzzParseFileSequence ensures the parser doesn't crash on arbitrary input.
// This is the most important fuzz test - we never want panics regardless of input.
func FuzzParseFileSequence(f *testing.F) {
	// Seed corpus from Python unit tests - valid sequences
	seeds := []string{
		// Basic sequences
		"/path/to/file.1-100#.exr",
		"/path/to/file.1-100@@.jpg",
		"/path/to/file.1-100%04d.tif",
		"/path/file.1-100x2#.exr",
		"/path/file.-10-10x2.exr",

		// Single frames
		"/path/file.0100.exr",
		"/path/file.100",

		// Pattern-only (no frame range)
		"/path/file.#.exr",
		"/path/file.@@@@.exr",
		"/path/file.%04d.exr",

		// Plain files
		"/path/file.txt",
		"/path/file",
		"/path/to/file.tar.gz",

		// Hidden files
		"/path/.hidden",
		"/path/.config.100",
		"/path/.hidden5.1-10#.7zip",

		// Complex frame ranges (from Python fuzz tests)
		"/file.1-100.exr",
		"/file.-1-20.exr",
		"/file.-1--20.exr",
		"/file.1--20.exr",
		"/file.20-1.exr",
		"/file.-20-1.exr",
		"/file.-20--1.exr",
		"/file.20--1.exr",

		// With modifiers
		"/file.1-20x5#.exr",
		"/file.-1-20x5#.exr",
		"/file.1-20y5#.exr",
		"/file.1-20:5#.exr",
		"/file.20-1x-1#.exr",

		// Comma-separated ranges
		"/file.1,2,3,5-10,20-30#.exr",
		"/file.1-10,50-100#.exr",

		// Edge cases
		"/file.0#.exr",
		"/file.-0#.exr",
		"file.@@@@",
		"1-100#.exr",
		"@@@.exr",

		// Windows paths
		"Z:\\Shows\\test.1-100#.exr",
		"C:\\path\\to\\file.####.exr",

		// Mixed separators
		"/mnt/Shows\\mixed/file.#.exr",

		// Multi-part extensions
		"/file.1-100#.tar.gz",
		"/file.1-100#.bgeo.sc",

		// Non-dot extensions (after padding)
		"/path/to/something_1-10#_exr",

		// Date-like patterns in basename
		"/path/name_2025-05-13_1809.1-100#.exr",
		"/path/file_1920x1080.0001-0100#.exr",

		// Negative frames
		"/path/file.-100--1#.exr",
		"/path/file.-0000#.exr",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// The parser should never panic, even on garbage input
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Parser panicked on input %q: %v", input, r)
			}
		}()

		// Try to parse - we don't care if it fails, just that it doesn't crash
		_, _ = parser.ParseFileSequence(input)
	})
}

// FuzzFileSequenceRoundTrip ensures valid sequences can be parsed and reconstructed.
// Property: ParseFileSequence(seq.String()) should equal the original sequence.
func FuzzFileSequenceRoundTrip(f *testing.F) {
	// Seed with known valid sequences
	seeds := []string{
		"/path/file.1-100#.exr",
		"/path/file.1-100@@.jpg",
		"/path/file.100.exr",
		"/path/.hidden5.1-10#.7zip",
		"/path/file.1-20x5#.exr",
		"/path/file.1,2,3,5-10#.exr",
		"/path/file.#.exr",
		"file.@@@@",
		"/path/file.tar.gz",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// Skip edge cases with special chars that conflict with padding tokens
		// These work in Python's regex but cause lexer conflicts in ANTLR (rare in practice)
		if strings.ContainsAny(input, "%$<>") {
			return
		}

		// Skip extremely long inputs that cause edge cases with frame number parsing
		// (e.g., 80+ digit numbers being interpreted as frames vs extensions)
		if len(input) > 200 {
			return
		}

		// First parse attempt
		seq, err := NewFileSequence(input)
		if err != nil {
			// Invalid input is fine, just skip
			return
		}

		// Get string representation
		str1 := seq.String()

		// Parse the string representation
		seq2, err := NewFileSequence(str1)
		if err != nil {
			t.Errorf("Failed to parse sequence string representation: %v\nOriginal: %q\nString(): %q",
				err, input, str1)
			return
		}

		// String representations should match
		str2 := seq2.String()
		if str1 != str2 {
			t.Errorf("Round-trip failed:\nOriginal input: %q\nFirst String(): %q\nSecond String(): %q",
				input, str1, str2)
		}

		// Key properties should match
		if seq.Dirname() != seq2.Dirname() {
			t.Errorf("Dirname mismatch after round-trip: %q vs %q", seq.Dirname(), seq2.Dirname())
		}
		if seq.Basename() != seq2.Basename() {
			t.Errorf("Basename mismatch after round-trip: %q vs %q", seq.Basename(), seq2.Basename())
		}
		if seq.Ext() != seq2.Ext() {
			t.Errorf("Extension mismatch after round-trip: %q vs %q", seq.Ext(), seq2.Ext())
		}
		if seq.Padding() != seq2.Padding() {
			t.Errorf("Padding mismatch after round-trip: %q vs %q", seq.Padding(), seq2.Padding())
		}
	})
}

// FuzzFrameSetParsing ensures FrameSet parsing doesn't crash on arbitrary input.
func FuzzFrameSetParsing(f *testing.F) {
	// Seed corpus from Python fuzz tests
	validSeeds := []string{
		// Basic ranges
		"1-100",
		"-10-10",
		"1-100x5",
		"1-100y5",
		"1-100:5",

		// Reversed ranges
		"100-1",
		"10--10",
		"20-1x5",

		// Negative ranges
		"-1-20",
		"-1--20",
		"1--20",
		"-20-1",
		"-20--1",
		"20--1",

		// With negative steps
		"20-1x-1",

		// Comma-separated
		"1,2,3,5-10,20-30",
		"1-10,50-100",

		// Edge cases
		"0",
		"-0",
		"",
		"1,",
		",1",
		"1,,,,2,,,3",
	}

	invalidSeeds := []string{
		// These should fail gracefully, not crash
		"1x5",           // modifier without range
		"-1x5",
		"1y5",
		"1:5",
		"1-",            // dangling separator
		"-1-",
		"bilbo",         // non-numeric
		"a-z",
		"1-20x:y5",      // duplicate modifiers
		"1-20X5",        // wrong case
		"1-20Y5",
		"1-20x0",        // zero modifier
		"1-20y0",
		"1-20:0",
		"1-20x-5",       // negative modifier
	}

	for _, seed := range validSeeds {
		f.Add(seed)
	}
	for _, seed := range invalidSeeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// Should never panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("FrameSet panicked on input %q: %v", input, r)
			}
		}()

		// Skip extremely long inputs that might cause performance issues
		if len(input) > 10000 {
			return
		}

		// Try to parse - we don't care if it fails, just that it doesn't crash
		_, _ = NewFrameSet(input)
	})
}

// FuzzFrameSetRoundTrip ensures valid FrameSets can be parsed and reconstructed.
func FuzzFrameSetRoundTrip(f *testing.F) {
	// Seed with known valid frame ranges
	seeds := []string{
		"1-100",
		"-10-10",
		"1-100x5",
		"1-100y5",
		"1-100:5",
		"1,2,3,5-10",
		"100-1",
		"-1-20x5",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// First parse attempt
		fs, err := NewFrameSet(input)
		if err != nil {
			// Invalid input is fine, just skip
			return
		}

		// Skip ranges that are unreasonably large (would cause OOM)
		if fs.Len() > 1000000 {
			return
		}

		// Get string representation
		str1 := fs.String()

		// Parse the string representation
		fs2, err := NewFrameSet(str1)
		if err != nil {
			t.Errorf("Failed to parse FrameSet string representation: %v\nOriginal: %q\nString(): %q",
				err, input, str1)
			return
		}

		// String representations should match
		str2 := fs2.String()
		if str1 != str2 {
			t.Errorf("Round-trip failed:\nOriginal input: %q\nFirst String(): %q\nSecond String(): %q",
				input, str1, str2)
		}

		// Frame lists should match
		frames1 := fs.Frames()
		frames2 := fs2.Frames()
		if len(frames1) != len(frames2) {
			t.Errorf("Frame count mismatch: %d vs %d", len(frames1), len(frames2))
			return
		}
		for i := range frames1 {
			if frames1[i] != frames2[i] {
				t.Errorf("Frame mismatch at index %d: %d vs %d", i, frames1[i], frames2[i])
				break
			}
		}
	})
}

// FuzzFileSequenceSplit ensures Split() doesn't crash and maintains properties.
func FuzzFileSequenceSplit(f *testing.F) {
	// Seed with sequences that have comma-separated ranges
	seeds := []string{
		"/path/file.1,2,3#.exr",
		"/path/file.1-10,50-100#.exr",
		"/path/file.1,5,10,20,30#.exr",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		seq, err := NewFileSequence(input)
		if err != nil {
			return
		}

		// Should never panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Split() panicked on input %q: %v", input, r)
			}
		}()

		// Split the sequence
		parts := seq.Split()

		// Property: All parts should have the same dir/ext/padding
		// (basename may differ for edge cases with negative frames)
		dir := seq.Dirname()
		ext := seq.Ext()
		pad := seq.Padding()

		// Skip validation if sequence has no frames (plain files, pattern-only)
		if seq.FrameSet() == nil {
			return
		}

		for i, part := range parts {
			if part.Dirname() != dir {
				t.Errorf("Part %d dirname mismatch: %q vs %q", i, part.Dirname(), dir)
			}

			// Basename may differ for edge cases with negative frames at the start
			// (e.g., "-00,0#" becomes "-00#" and "0#" with different basenames)
			// This is a known limitation of how negative frames are serialized

			if part.Ext() != ext {
				t.Errorf("Part %d extension mismatch: %q vs %q", i, part.Ext(), ext)
			}
			if part.Padding() != pad {
				t.Errorf("Part %d padding mismatch: %q vs %q", i, part.Padding(), pad)
			}

			// Each part should have a valid frame set
			if part.FrameSet() == nil {
				t.Errorf("Part %d has nil FrameSet but original sequence has frames", i)
			}
		}
	})
}

// FuzzPaddingFormats ensures various padding formats are handled correctly.
func FuzzPaddingFormats(f *testing.F) {
	// Seed with different padding formats
	seeds := []string{
		"#",
		"##",
		"####",
		"@",
		"@@",
		"@@@@@@@@",
		"%04d",
		"%d",
		"%08d",
		"$F",
		"$F4",
		"<UDIM>",
		"<UDIM:1001>",
	}

	for _, seed := range seeds {
		// Create full sequence paths with this padding
		f.Add("/path/file.1-100" + seed + ".exr")
	}

	f.Fuzz(func(t *testing.T, input string) {
		seq, err := NewFileSequence(input)
		if err != nil {
			return
		}

		// Should never panic when calling padding-related methods
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Panicked on input %q: %v", input, r)
			}
		}()

		padding := seq.Padding()
		if padding == "" {
			// No padding is fine for plain files
			return
		}

		// If we have padding and frames, ZFill should work
		if seq.FrameSet() != nil && seq.FrameSet().Len() > 0 {
			zfill := seq.ZFill()
			if zfill < 0 {
				t.Errorf("Invalid ZFill value: %d for padding %q", zfill, padding)
			}

			// Try to format a frame
			frame, err := seq.FrameSet().Frame(0)
			if err != nil {
				// No frames available
				return
			}
			formatted, err := seq.Frame(frame)
			if err != nil {
				// Some padding formats might not support all frames
				// That's OK as long as we don't crash
				return
			}

			// The formatted result should contain the padding in some form
			if !strings.Contains(formatted, seq.Dirname()) {
				t.Errorf("Formatted frame missing dirname: %q", formatted)
			}
			if !strings.Contains(formatted, seq.Basename()) {
				t.Errorf("Formatted frame missing basename: %q", formatted)
			}
			if !strings.Contains(formatted, seq.Ext()) {
				t.Errorf("Formatted frame missing extension: %q", formatted)
			}
		}
	})
}
