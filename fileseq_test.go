package fileseq

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

// Test table for comparing string frame ranges
// to the parsed frame set
var stringRangeTable = []struct {
	frange   string
	expected []int
}{
	{"10", []int{10}},
	{"1-5", []int{1, 2, 3, 4, 5}},
	{"1-10x5", []int{1, 6}},
	{"1-10x2", []int{1, 3, 5, 7, 9}},
	{"1-10:4", []int{1, 5, 9, 4, 7, 10, 3, 2, 6, 8}},
	{"1-10y2", []int{2, 4, 6, 8, 10}},
	{"1-15y4", []int{2, 3, 4, 6, 7, 8, 10, 11, 12, 14, 15}},
	{"1-3,20,30-33", []int{1, 2, 3, 20, 30, 31, 32, 33}},
	// Complex + Whitespace test
	{" 1 - 10x2, 50, 60 - 62, 70 - 74x2 ", []int{1, 3, 5, 7, 9, 50, 60, 61, 62, 70, 72, 74}},
	// Descending test
	{"10-1", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	{"10-1x-1", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	{"10-1x1", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
}

func TestNewFrameSet(t *testing.T) {
	for i, tt := range stringRangeTable {
		isRange := IsFrameRange(tt.frange)
		if !isRange && tt.expected != nil {
			t.Fatalf("#%d: %q did not validate as a frame range", i, tt.frange)
		}

		if isRange && tt.expected == nil {
			t.Fatalf("#%d: %q validated as a frame range when we expected it to fail", i, tt.frange)
		}

		s, err := NewFrameSet(tt.frange)
		if err != nil {
			t.Fatalf("#%d: Failed to parse %q: %s", i, tt.frange, err.Error())
		}
		if s.frange == "" {
			t.Fatalf("#%d: Got an empty frange field on FrameSet", i)
		}

		size := s.Len()
		if size != len(tt.expected) {
			t.Errorf("#%d: Expected Len to be %d, got %d", i, len(tt.expected), size)
		}

		if !reflect.DeepEqual(s.Frames(), tt.expected) {
			t.Errorf("#%d: While parsing %q, got %v, expected %v", i, tt.frange, s.Frames(), tt.expected)
		}
	}
}

func TestFrameSetLookup(t *testing.T) {
	for _, tt := range stringRangeTable {
		s, err := NewFrameSet(tt.frange)
		if err != nil {
			t.Fatalf("Failed to parse %q: %s", tt.frange, err.Error())
		}

		start := tt.expected[0]
		actual := s.Start()
		if actual != start {
			t.Fatalf("Expected Start frame to be %d, got %d", start, actual)
		}

		end := tt.expected[len(tt.expected)-1]
		actual = s.End()

		if actual != end {
			t.Fatalf("Expected End frame to be %d, got %d", end, actual)
		}

		for idx, val := range s.Frames() {
			if !s.HasFrame(val) {
				t.Errorf("Expected frame %d to be in frames %v", val, s.Frames())
			}
			actualIdx := s.Index(val)
			if actualIdx != idx {
				t.Errorf("Expected index for val %d to be %d, got %d", val, idx, actualIdx)
			}
			actualVal, err := s.Frame(idx)
			if err != nil {
				t.Fatalf("Did not find valid frame for idx %d in frames %v", idx, s.Frames())
			}
			if actualVal != val {
				t.Errorf("Expected value for index %d to be %d, got %d", idx, val, actualVal)
			}
		}
	}
}

func TestFrameSetPadding(t *testing.T) {
	var table = []struct {
		input  string
		pad    int
		output string
	}{
		{"1", -1, "1"},
		{"1", 0, "1"},
		{"1", 1, "1"},
		{"1", 4, "0001"},
		{"-1", 4, "-001"},
		{"-123", 4, "-123"},
		{"-123", 5, "-0123"},
		{"1-100", 4, "0001-0100"},
		{"-10-100", 4, "-010-0100"},
		{"-10--5", 4, "-010--005"},
		{"1,20,300,4000", 4, "0001,0020,0300,4000"},
		{"-10,-5,1,5,20-40,50-300x5", 4, "-010,-005,0001,0005,0020-0040,0050-0300x5"},
	}

	for _, tt := range table {
		s, err := NewFrameSet(tt.input)
		if err != nil {
			t.Fatalf("Failed to parse %q: %s", tt.input, err.Error())
		}
		actual := s.FrameRangePadded(tt.pad)
		if actual != tt.output {
			t.Errorf("Expected %s for pad %d, got %s", tt.output, tt.pad, actual)
		}
	}
}

func TestFrameSetNormalize(t *testing.T) {
	var table = []struct {
		input  string
		output string
	}{
		{"1,2,3,4,5", "1-5"},
		{"5,1,4,2,3", "1-5"},
		{"1-5,6,7,8,9,10,11-20", "1-20"},
	}

	for _, tt := range table {
		s, err := NewFrameSet(tt.input)
		if err != nil {
			t.Fatalf("Failed to parse %q: %s", tt.input, err.Error())
		}
		actual := s.Normalize().FrameRange()
		if actual != tt.output {
			t.Errorf("Expected %s, got %s", tt.output, actual)
		}
	}
}

func TestFrameSetInverted(t *testing.T) {
	var table = []struct {
		frange, expected string
		zfill            int
	}{
		{"1-5,7-10", "6", 0},
		{"1-5,7-10", "0006", 4},
		{"-10--5,5-10", "-4-4", 0},
		{"-10--5,5-10", "-004-0004", 4},
		{"1-10x2", "2-8x2", 0},
		{"1-10x2", "002-008x2", 3},
		{"1-10y2", "3-9x2", 0},
		{"1-10y2", "00003-00009x2", 5},
		{"1-10:4", "", 0},
		{" 1 - 10x2, 50, 60 - 62, 70 - 74x2 ", "2-10x2,11-49,51-59,63-69,71-73x2", 0},
	}

	for _, tt := range table {
		s, _ := NewFrameSet(tt.frange)
		actual := s.InvertedFrameRange(tt.zfill)
		if actual != tt.expected {
			t.Errorf("Expected range %q, got %q", tt.expected, actual)
		}
	}
}

func TestFrameSetSplitPattern(t *testing.T) {
	var table = []struct {
		path  string
		base  string
		start int
		end   int
		pad   string
	}{
		{"/path/to/file%s1-1x1#.exr", "file%s", 1, 1, "#"},
		{"/path/to/file%s_1-100x10@@.exr", "file%s_", 1, 91, "@@"},
		{"/path/to/file%s.-10--1x2##.exr", "file%s.", -10, -2, "##"},
		{"/path/to/file%s1,2,3,5-10,20-30#.exr", "file%s", 1, 30, "#"},

		{"/path/to/file%s1-1x1#.tar.gz", "file%s", 1, 1, "#"},
		{"/path/to/file%s_1-100x10@@.tar.gz", "file%s_", 1, 91, "@@"},
		{"/path/to/file%s.-10--1x2##.tar.gz", "file%s.", -10, -2, "##"},
		{"/path/to/file%s1,2,3,5-10,20-30#.tar.gz", "file%s", 1, 30, "#"},

		{"/path/to/file%s_1-100x10%%02d.exr", "file%s_", 1, 91, "%02d"},
		{"/path/to/file%s_1-100x10$F2.exr", "file%s_", 1, 91, "$F2"},
		{"/path/to/file%s_1-100x10$F02.exr", "file%s_", 1, 91, "$F02"},
	}

	// ref: splitPattern range directive chars
	for i, tt := range table {

		for _, char := range ":xy," {

			expected := fmt.Sprintf(tt.path, string(char))
			base := fmt.Sprintf(tt.base, string(char))

			t.Run(fmt.Sprintf("%d_%s", i, base), func(t *testing.T) {
				fs, err := NewFileSequence(expected)
				if err != nil {
					t.Fatal(err.Error())
				}

				if fs.String() != expected {
					t.Fatalf("Expected String %q, got %q", expected, fs.String())
				}
				if fs.Basename() != base {
					t.Fatalf("Expected basename %q, got %q", base, fs.Basename())
				}
				if fs.Start() != tt.start {
					t.Fatalf("Expected Start %v, got %v", tt.start, fs.Start())
				}
				if fs.End() != tt.end {
					t.Fatalf("Expected End %v, got %v", tt.end, fs.End())
				}
				if fs.Padding() != tt.pad {
					t.Fatalf("Expected Padding %v, got %v", tt.pad, fs.Padding())
				}
			})
		}
	}
}

func TestNewFileSequence(t *testing.T) {
	var table = []struct {
		path, outPath     string
		start, end, zfill int
		frameCount        int
		ext               string
	}{
		{"/file_path.100.exr",
			"/file_path.100@@@.exr", 100, 100, 3,
			1, ".exr"},
		{"/file_path.0100.exr",
			"/file_path.0100#.exr", 100, 100, 4,
			1, ".exr"},
		{"/dir/f.1-100#.jpeg",
			"/dir/f.1-100#.jpeg", 1, 100, 4,
			100, ".jpeg"},
		{"/dir/f.1-100@@@.f",
			"/dir/f.1-100@@@.f", 1, 100, 3,
			100, ".f"},
		{"/dir/f.1-10,50,60-90x2##.mp4",
			"/dir/f.1-10,50,60-90x2##.mp4", 1, 90, 8,
			27, ".mp4"},
		{"/dir/f.exr",
			"/dir/f.exr", 0, 0, 0,
			1, ".exr"},
		{"/dir/f.100",
			"/dir/f.100@@@", 100, 100, 3,
			1, ""},
		{"/dir/f.@@.ext",
			"/dir/f.@@.ext", 0, 0, 2,
			1, ".ext"},
		{"/dir/f100.ext",
			"/dir/f100@@@.ext", 100, 100, 3,
			1, ".ext"},
		{"/dir/f_100.ext",
			"/dir/f_100@@@.ext", 100, 100, 3,
			1, ".ext"},
		{"/dir/no_frames.ext",
			"/dir/no_frames.ext", 0, 0, 0,
			1, ".ext"},
		{"/dir/no_file_extension",
			"/dir/no_file_extension", 0, 0, 0,
			1, ""},
		{"/dir/.hidden",
			"/dir/.hidden", 0, 0, 0,
			1, ".hidden"},
		{"/dir/.hidden.100",
			"/dir/.hidden.100@@@", 100, 100, 3,
			1, ""},
		{"/dir/.hidden.100.ext",
			"/dir/.hidden.100@@@.ext", 100, 100, 3,
			1, ".ext"},
		{"/dir/.hidden5.1-10#.7zip",
			"/dir/.hidden5.1-10#.7zip", 1, 10, 4,
			10, ".7zip"},
		{"/dir/complex_extension.1-10,50,60-90x2#.tar.gz",
			"/dir/complex_extension.1-10,50,60-90x2#.tar.gz", 1, 90, 4,
			27, ".tar.gz"},
		{"/dir/many_numbers8_30.v123.1-10,50,60-90x2#.tar.gz",
			"/dir/many_numbers8_30.v123.1-10,50,60-90x2#.tar.gz", 1, 90, 4,
			27, ".tar.gz"},
		{"/dir/f.tar.gz",
			"/dir/f.tar.gz", 0, 0, 0,
			1, ".gz"},
		{"/dir/f.@@.tar.gz",
			"/dir/f.@@.tar.gz", 0, 0, 2,
			1, ".tar.gz"},
		{"/dir/f100.tar.gz",
			"/dir/f100@@@.tar.gz", 100, 100, 3,
			1, ".tar.gz"},
		{"/dir/f10_30.v123.100#.tar.gz",
			"/dir/f10_30.v123.100#.tar.gz", 100, 100, 4,
			1, ".tar.gz"},
		{"/dir/f_100.tar.gz",
			"/dir/f_100@@@.tar.gz", 100, 100, 3,
			1, ".tar.gz"},
		{"/dir/no_frames.tar.gz",
			"/dir/no_frames.tar.gz", 0, 0, 0,
			1, ".gz"},
		{".10000000000",
			".10000000000", 0, 0, 0,
			1, ".10000000000"},
		{".10000000000.123",
			".10000000000@@@@@@@@@@@.123", 10000000000, 10000000000, 11,
			1, ".123"},
		{"/dir/f.1-100%04d.jpeg",
			"/dir/f.1-100%04d.jpeg", 1, 100, 4,
			100, ".jpeg"},
		{"/dir/f.1-100$F04.jpeg",
			"/dir/f.1-100$F04.jpeg", 1, 100, 4,
			100, ".jpeg"},
		{"/dir/f.tmp12345@@@@@",
			"/dir/f.tmp12345@@@@@", 12345, 12345, 5,
			1, ""},
	}
	for _, tt := range table {
		seq, err := NewFileSequence(tt.path)
		if err != nil {
			t.Fatal(err.Error())
		}
		actual := seq.String()
		if actual != tt.outPath {
			t.Errorf("Expected %s; got %s (for path: %s)", tt.outPath, actual, tt.path)
		}
		if seq.Start() != tt.start {
			t.Errorf("Expected %q start to be %d, got %d", tt.path, tt.start, seq.Start())
		}
		if seq.End() != tt.end {
			t.Errorf("Expected %q end to be %d, got %d", tt.path, tt.end, seq.End())
		}
		if seq.ZFill() != tt.zfill {
			t.Errorf("Expected %q zfill to be %d, got %d", tt.path, tt.zfill, seq.ZFill())
		}
		if seq.Len() != tt.frameCount {
			t.Errorf("Expected %q frame count to be %d, got %d", tt.path, tt.frameCount, seq.Len())
		}
		if seq.Ext() != tt.ext {
			t.Errorf("Expected %q ext to be %q, got %q", tt.path, tt.ext, seq.Ext())
		}
	}
}

func TestFileSequenceSplit(t *testing.T) {
	var table = []struct {
		path    string
		franges []string
	}{
		{"/file_path.100.exr", []string{"100"}},
		{"/file_path.-005.exr", []string{"-005"}},
		{"/dir/f.-1-100#.jpeg", []string{"-1-100"}},
		{"/dir/f.-1-100%04d.jpeg", []string{"-1-100"}},
		{"/dir/f.-1-100$F04.jpeg", []string{"-1-100"}},
		{"/dir/f.1-10,50,60-90x2##.exr", []string{"1-10", "50", "60-90x2"}},
		{"/dir/f.1-10,50,60-90x2##.tar.gz", []string{"1-10", "50", "60-90x2"}},
		{"/dir/f.exr", []string{""}},
	}
	for _, tt := range table {
		seq, err := NewFileSequence(tt.path)
		if err != nil {
			t.Fatal(err.Error())
		}
		actual := []string{}
		for _, spl := range seq.Split() {
			actual = append(actual, spl.FrameRange())
		}
		if !reflect.DeepEqual(tt.franges, actual) {
			t.Errorf("Expected ranges %v ; Got %v", tt.franges, actual)
		}
	}
}

func TestFileSequenceFrameIndex(t *testing.T) {
	var table = []struct {
		path     string
		size     int
		idx      int
		frame    int
		expected string
	}{
		{"file.100.exr", 1, 0, 100, "file.100.exr"},
		{"file.1-100#.exr", 100, 4, 5, "file.0005.exr"},
		{"file.-10-30@@@.exr", 40, 25, 15, "file.015.exr"},
		{"file.-10-30@@@.exr", 40, 5, -5, "file.-05.exr"},
		{"file.1,3-12x3,20-24x2,28,32##.exr", 10, 8, 28, "file.00000028.exr"},
	}
	for _, tt := range table {
		seq, err := NewFileSequence(tt.path)
		if err != nil {
			t.Fatal(err.Error())
		}
		// Frame test
		actual, err := seq.Frame(tt.frame)
		if err != nil {
			t.Fatalf("Unexpected error in Frame(%d) for %q: %s", tt.frame, tt.path, err.Error())
		}
		if actual != tt.expected {
			t.Errorf("Expected %q for frame %d ; got %q", tt.expected, tt.frame, actual)
		}
		// Index test
		actual = seq.Index(tt.idx)
		if actual != tt.expected {
			t.Errorf("Expected %q for index %d ; got %q", tt.expected, tt.idx, actual)
		}
		// Frame count test
		for i := 0; i < tt.size; i++ {
			path := seq.Index(i)
			if path == "" {
				t.Fatalf("Got an empty path string for index %d in %q", i, tt.path)
			}
		}

	}
}

func TestFileSequenceSetDir(t *testing.T) {
	var table = []struct {
		src, newDir, newBase, newExt, newFrange, newPad, expected string
	}{
		{
			src:       "/path/to/file.100.exr",
			newDir:    "/other",
			newBase:   "fileB.",
			newExt:    "jpg",
			newFrange: "200",
			newPad:    "#",
			expected:  "/other/fileB.200#.jpg",
		},
		{
			src:       "/path/to/file.1-100#.exr",
			newDir:    "/other/subdir",
			newBase:   "fileB",
			newExt:    "f",
			newFrange: "-10-5,20-30x2",
			newPad:    "@@@",
			expected:  "/other/subdir/fileB-10-5,20-30x2@@@.f",
		},
		{
			src:       "/path/to/file.1-100#.exr",
			newDir:    "/other/subdir",
			newBase:   "fileB",
			newExt:    "f",
			newFrange: "-10-5,20-30x2",
			newPad:    "%03d",
			expected:  "/other/subdir/fileB-10-5,20-30x2%03d.f",
		},
		{
			src:       "/path/to/file.1-100#.exr",
			newDir:    "/other/subdir",
			newBase:   "fileB",
			newExt:    "f",
			newFrange: "-10-5,20-30x2",
			newPad:    "$F03",
			expected:  "/other/subdir/fileB-10-5,20-30x2$F03.f",
		},
	}

	for _, tt := range table {
		seq, err := NewFileSequence(tt.src)
		if err != nil {
			t.Fatal(err.Error())
		}
		seq.SetDirname(tt.newDir)
		seq.SetBasename(tt.newBase)
		seq.SetExt(tt.newExt)
		seq.SetPadding(tt.newPad)
		seq.SetFrameRange(tt.newFrange)
		actual := filepath.ToSlash(seq.String())
		if actual != tt.expected {
			t.Errorf("Expected %q ; got %q", tt.expected, actual)
		}
	}
}

func TestFileSequenceFormat(t *testing.T) {
	seq, err := NewFileSequence("/a/path/to/the/file_foo.1-10,50,70-100x5#.ext")
	if err != nil {
		t.Fatal(err.Error())
	}
	var table = []struct {
		format   string
		expected string
	}{
		{ // 0
			"{{dir}}{{base}}{{frange}}{{pad}}{{ext}}",
			seq.String(),
		},
		{ // 1
			"{{startf}} {{endf}} {{len}} {{zfill}}",
			"1 100 18 4",
		},
		{ // 2
			"{{base}}{{if inverted}}{{inverted}}{{else}}{{frange}}{{end}}{{ext}}",
			"file_foo.11-49,51-69,71-74,76-79,81-84,86-89,91-94,96-99.ext",
		},
	}

	for i, tt := range table {
		actual, err := seq.Format(tt.format)
		if err != nil {
			t.Fatalf("Test %d: Error formatting %q: %s", i, tt.format, err.Error())
		}
		if actual != tt.expected {
			t.Errorf("Test %d: Expected %q ; got %q", i, tt.expected, actual)
		}
	}
}

func TestFramesToFrameRange(t *testing.T) {
	var table = []struct {
		frames   []int
		sorted   bool
		zfill    int
		expected string
	}{
		{[]int{-10}, false, 4, "-010"},
		{[]int{1, 2, 3, 4, 5}, true, 0, "1-5"},
		{[]int{1, 3, 5, 7, 9}, true, 0, "1-9x2"},
		{
			[]int{1, 2, 3, 4, 5, 20, 30, 40, 41, 42, 43, 44, 50, 60, 62, 64, 66},
			true, 0, "1-5,20-40x10,41-44,50,60-66x2",
		},
		{
			[]int{1, 3, 6, 9, 12, 20, 22, 24, 28, 32},
			true, 0, "1,3-12x3,20-24x2,28,32",
		},
		{
			[]int{1, 11, 21, 31, 41, 51, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 15, 18},
			false, 2, "01-51x10,02-10,12-18x3",
		},
	}

	for _, tt := range table {
		actual := FramesToFrameRange(tt.frames, tt.sorted, tt.zfill)
		if actual != tt.expected {
			t.Errorf("Expected range %q, got %q", tt.expected, actual)
		}
	}
}

func TestPaddingChars(t *testing.T) {
	type TestVals struct {
		pad      int
		padChars string
	}

	var table = []struct {
		padMapper paddingMapper
		testVals  []TestVals
	}{
		{
			padders[PadStyleHash1], []TestVals{
				{-1, "#"}, {0, "#"}, {1, "#"}, {2, "##"}, {3, "###"},
				{1, "#"}, {5, "#####"}, {2, "##"}, {3, "###"},
			},
		},
		{
			padders[PadStyleHash4], []TestVals{
				{-1, "@"}, {0, "@"}, {1, "@"}, {2, "@@"}, {3, "@@@"},
				{4, "#"}, {5, "@@@@@"}, {8, "##"}, {12, "###"},
			},
		},
	}

	for _, tg := range table {
		for _, tt := range tg.testVals {
			actual := tg.padMapper.PaddingChars(tt.pad)
			if actual != tt.padChars {
				t.Errorf("Expected %s for padding %d, got %s", tt.padChars, tt.pad, actual)
			}
		}
	}
}

func TestPaddingCharsSize(t *testing.T) {
	type TestVals struct {
		padChars string
		pad      int
	}

	var table = []struct {
		padMapper paddingMapper
		testVals  []TestVals
	}{
		{
			padders[PadStyleHash1], []TestVals{
				{"", 0},
				{"#", 1},
				{"##", 2},
				{"###", 3},
				{"####", 4},
				{"@", 1},
				{"@@", 2},
				{"@@@", 3},
				{"@@@@", 4},
				{"%02d", 2},
				{"%04d", 4},
				{"$F", 1},
				{"$F2", 2},
				{"$F04", 4},
			},
		},
		{
			padders[PadStyleHash4], []TestVals{
				{"", 0},
				{"#", 4},
				{"##", 8},
				{"###", 12},
				{"####", 16},
				{"@", 1},
				{"@@", 2},
				{"@@@", 3},
				{"@@@@", 4},
				{"%02d", 2},
				{"%04d", 4},
				{"$F2", 2},
				{"$F04", 4},
			},
		},
	}

	for _, tg := range table {
		for _, tt := range tg.testVals {
			actual := tg.padMapper.PaddingCharsSize(tt.padChars)
			if actual != tt.pad {
				t.Errorf("Expected %d for padding %s, got %d", tt.pad, tt.padChars, actual)
			}
		}
	}
}

func TestToRange(t *testing.T) {
	var table = []struct {
		expected         []int
		start, end, step int
	}{
		{[]int{1}, 1, 1, 1},
		{[]int{1, 2, 3, 4, 5}, 1, 5, 1},
		{[]int{1, 3, 5}, 1, 6, 2},
		{[]int{1, 6}, 1, 10, 5},
		{[]int{-5, -4, -3, -2, -1, -0, 1, 2, 3}, -5, 3, 1},
		{[]int{-5, -4, -3, -2, -1, -0, 1, 2, 3}, -5, 3, -5},
		{[]int{5, 4, 3, 2, 1}, 5, 1, 1},
		{[]int{6, 4, 2}, 6, 1, 2},
		{[]int{20, 15, 10, 5}, 20, 4, 5},
		{[]int{-5, -10, -15, -20}, -5, -20, 5},
	}

	for i, tt := range table {
		actual := toRange(tt.start, tt.end, tt.step)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("#%d: Got %v, expected %v", i, actual, tt.expected)
		}
	}
}

func TestFindSequencesOnDisk(t *testing.T) {
	table := map[string][]string{
		"testdata": {
			"seqD.2-10@.gif",
			"seqC.-5-2,4-10,20-21,27-30@@.tif",
			"seqB.5-14,16-18,20#.jpg",
			"seqA.1,3-6,8-10#.exr",
			"complex.5-7#.tar.gz",
		},
		"testdata/mixed": {
			"seq.-1-5@@.ext",
			"seq.-1-5#.ext",
			"seq.-1-5,1001@@@@@.ext",
		},
	}

	for dir, expected := range table {
		seqs, err := FindSequencesOnDisk(dir)
		if err != nil {
			t.Fatal(err.Error())
		}

		actual := make([]string, 0, len(seqs))
		for _, seq := range seqs {
			actual = append(actual, filepath.Base(seq.String()))
		}

		sort.Strings(actual)
		sort.Strings(expected)

		if len(actual) != len(expected) {
			t.Logf("\nExpect: %v\nActual: %v", expected, actual)
			t.Errorf("For dir %q, expected %d seqs ; got %d", dir, len(expected), len(actual))
			continue
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("For dir %q\n  expected %v\n  got %v", dir, expected, actual)
		}
	}
}

func TestListFiles(t *testing.T) {
	expected := map[string]int{
		"seqD.2-10@.gif":                   0,
		"seqC.-5-2,4-10,20-21,27-30@@.tif": 0,
		"seqB.5-14,16-18,20#.jpg":          0,
		"seqA.1,3-6,8-10#.exr":             0,
		"complex.5-7#.tar.gz":              0,
		"aFile.ext":                        0,
		"file_without_ext":                 0,
	}

	seqs, err := ListFiles("testdata")
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(seqs) != len(expected) {
		t.Fatalf("Expected %d seqs ; got %d", len(expected), len(seqs))
	}

	for _, s := range seqs {
		name := filepath.Base(s.String())
		if _, ok := expected[name]; !ok {
			t.Fatalf("Parsed seq %q not in expected list", name)
		}
		expected[name]++
	}

	for name, count := range expected {
		if count != 1 {
			t.Errorf("Got # of matchs %d instead of 1, for seq/file %q", count, name)
		}
	}
}

func TestFindSequencesInList(t *testing.T) {
	table := []struct {
		Name   string
		Paths  []string
		Opts   []FileOption
		Expect []string
	}{
		{
			Name:   "empty",
			Paths:  []string{},
			Expect: []string{},
		},
		{
			Name:   "single seq",
			Paths:  []string{"/path/to/file.0001.ext", "/path/to/file.0002.ext", "/path/to/file.0003.ext"},
			Expect: []string{"/path/to/file.1-3#.ext"},
		},
		{
			Name: "multi seq",
			Paths: []string{
				"/path/to/foo.0001.ext", "/path/to/foo.0002.ext", "/path/to/foo.0003.ext",
				"/path/to/a/bar.020.ext", "/path/to/a/bar.030.ext", "/path/to/a/bar.040.ext",
			},
			Expect: []string{
				"/path/to/foo.1-3#.ext",
				"/path/to/a/bar.20-40x10@@@.ext",
			},
		},
		{
			Name: "no single files",
			Paths: []string{
				"/path/to/single/file.ext",
				"/path/to/foo.0001.ext", "/path/to/foo.0002.ext", "/path/to/foo.0003.ext",
				"/path/to/bar.020.ext", "/path/to/bar.030.ext", "/path/to/bar.040.ext",
				"/path/to/single/.another.ext",
			},
			Expect: []string{
				"/path/to/foo.1-3#.ext",
				"/path/to/bar.20-40x10@@@.ext",
			},
		},
		{
			Name: "with single files",
			Paths: []string{
				"/path/to/single/file.ext",
				"/path/to/foo.0001.ext", "/path/to/foo.0002.ext", "/path/to/foo.0003.ext",
				"/path/to/a/bar.020.ext", "/path/to/a/bar.030.ext", "/path/to/a/bar.040.ext",
				"/path/to/single/.another.ext",
			},
			Opts: []FileOption{SingleFiles, HiddenFiles},
			Expect: []string{
				"/path/to/single/file.ext",
				"/path/to/foo.1-3#.ext",
				"/path/to/a/bar.20-40x10@@@.ext",
				"/path/to/single/.another.ext",
			},
		},
		{
			Name: "pad hash 1",
			Paths: []string{
				"/path/to/foo.0001.ext", "/path/to/foo.0002.ext", "/path/to/foo.0003.ext",
				"/path/to/bar.1.ext", "/path/to/bar.2.ext", "/path/to/bar.3.ext",
			},
			Opts: []FileOption{FileOptPadStyleHash1},
			Expect: []string{
				"/path/to/foo.1-3####.ext",
				"/path/to/bar.1-3#.ext",
			},
		},
		{
			Name: "ambiguous single files",
			Paths: []string{
				"/path/to/complex_name_15-14-56-14.ext",
				"/path/to/single/123",
				"/path/to/single/123.ext",
				"/path/to/single/file0001.ext",
				"/path/to/single/file.ext12345",
				"/path/to/single/file.ext12345z",
			},
			Opts: []FileOption{SingleFiles},
			Expect: []string{
				"/path/to/complex_name_15-14-56-14@@@.ext",
				"/path/to/single/123@@@",
				"/path/to/single/123@@@.ext",
				"/path/to/single/file1#.ext",
				"/path/to/single/file.ext12345",
				"/path/to/single/file.ext12345z",
			},
		},
		{
			Name: "ambiguous single files disabled",
			Paths: []string{
				"/path/to/complex_name_15-14-56-14.ext",
				"/path/to/single/123",
				"/path/to/single/123.ext",
				"/path/to/single/file0001.ext",
				"/path/to/single/file.ext12345",
				"/path/to/single/file.ext12345z",
			},
			Expect: []string{
				"/path/to/complex_name_15-14-56-14@@@.ext",
				"/path/to/single/123@@@.ext",
				"/path/to/single/file1#.ext",
			},
		},
	}

	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			seqs, err := FindSequencesInList(tt.Paths, tt.Opts...)
			if err != nil {
				t.Fatal(err)
			}

			actual := make([]string, 0, len(seqs))
			for _, seq := range seqs {
				actual = append(actual, seq.String())
			}

			cleanPaths(tt.Expect)
			sort.Strings(tt.Expect)
			sort.Strings(actual)

			if len(actual) != len(tt.Expect) {
				t.Errorf("\nExpected %d seqs ; got %d", len(tt.Expect), len(actual))
			}

			if !reflect.DeepEqual(actual, tt.Expect) {
				t.Errorf("\nExpect %v\n   Got %v", tt.Expect, actual)
			}
		})
	}
}

func TestFindSequenceOnDisk(t *testing.T) {
	type TestSet struct {
		name   string
		mapper PadStyle
		tests  map[string]string
	}

	table := []TestSet{
		{
			name:   "pad hash 1",
			mapper: PadStyleHash1,
			tests: map[string]string{
				"testdata/seqC.@@.tif":        "testdata/seqC.-5-2,4-10,20-21,27-30##.tif",
				"testdata/seqC.0010.tif":      "",
				"testdata/seqC.10.tif":        "testdata/seqC.-5-2,4-10,20-21,27-30##.tif",
				"testdata/seqB.####.jpg":      "testdata/seqB.5-14,16-18,20####.jpg",
				"testdata/seqB.16-18####.jpg": "testdata/seqB.5-14,16-18,20####.jpg",
				"testdata/seqA.####.exr":      "testdata/seqA.1,3-6,8-10####.exr",
				"testdata/seqA.@@@@.exr":      "testdata/seqA.1,3-6,8-10####.exr",
				"testdata/seqA.@.jpg":         "",

				"testdata/seqC.%02d.tif": "testdata/seqC.-5-2,4-10,20-21,27-30##.tif",
				"testdata/seqC.$F02.tif": "testdata/seqC.-5-2,4-10,20-21,27-30##.tif",

				"testdata/mixed/seq.####.ext":  "testdata/mixed/seq.-1-5####.ext",
				"testdata/mixed/seq.#.ext":     "",
				"testdata/mixed/seq.@@.ext":    "testdata/mixed/seq.-1-5##.ext",
				"testdata/mixed/seq.@@@@@.ext": "testdata/mixed/seq.-1-5,1001#####.ext",
				"testdata/mixed/seq.@.ext":     "",
			},
		},

		{
			name:   "pad hash 4",
			mapper: PadStyleHash4,
			tests: map[string]string{
				"testdata/seqC.@@.tif":     "testdata/seqC.-5-2,4-10,20-21,27-30@@.tif",
				"testdata/seqC.0010.tif":   "",
				"testdata/seqC.10.tif":     "testdata/seqC.-5-2,4-10,20-21,27-30@@.tif",
				"testdata/seqB.#.jpg":      "testdata/seqB.5-14,16-18,20#.jpg",
				"testdata/seqB.16-18#.jpg": "testdata/seqB.5-14,16-18,20#.jpg",
				"testdata/seqA.#.exr":      "testdata/seqA.1,3-6,8-10#.exr",
				"testdata/seqA.@.exr":      "",
				"testdata/seqA.@.jpg":      "",

				"testdata/seqC.%02d.tif": "testdata/seqC.-5-2,4-10,20-21,27-30@@.tif",
				"testdata/seqC.$F02.tif": "testdata/seqC.-5-2,4-10,20-21,27-30@@.tif",

				"testdata/mixed/seq.#.ext":     "testdata/mixed/seq.-1-5#.ext",
				"testdata/mixed/seq.@@.ext":    "testdata/mixed/seq.-1-5@@.ext",
				"testdata/mixed/seq.@@@@@.ext": "testdata/mixed/seq.-1-5,1001@@@@@.ext",
				"testdata/mixed/seq.@.ext":     "",
			},
		},
	}

	for _, tg := range table {
		t.Run(tg.name, func(t *testing.T) {
			for pattern, expected := range tg.tests {
				seq, err := FindSequenceOnDiskPad(pattern, tg.mapper, StrictPadding)
				if err != nil {
					t.Fatal(err.Error())
				}
				if seq == nil && expected != "" {
					t.Fatalf("For input %q, expected %q ; got a nil sequence", pattern, expected)
				}

				var actual string
				if expected != "" {
					actual = filepath.ToSlash(seq.String())
				}

				if actual != expected {
					t.Fatalf("For input %q,\n  expected %q\n  got %q", pattern, expected, actual)
				}
			}
		})
	}
}

func TestHandleSymlinksOnDisk(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping symlink test on windows")
	}

	seqs, err := FindSequencesOnDisk("testdata/versions")
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(seqs) != 1 {
		t.Fatalf("Expected 1 result ; got %d : %v", len(seqs), seqs)
	}

	actual := seqs[0].String()
	expected := "testdata/versions/seq.1-10#.ext"

	if actual != expected {
		t.Fatalf("Expected to find seq %q ; got %q", expected, actual)
	}
}

func cleanPaths(paths []string) {
	for i, path := range paths {
		paths[i] = filepath.Clean(path)
	}
}
