package fileseq

import (
	"reflect"
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
}

func TestNewFrameSet(t *testing.T) {
	for _, tt := range stringRangeTable {
		s, err := NewFrameSet(tt.frange)
		if err != nil {
			t.Fatalf("Failed to parse %q: %s", tt.frange, err.Error())
		}
		if s.frange == "" {
			t.Fatal("Got an empty frange field on FrameSet")
		}
		size := s.Len()
		if size != len(tt.expected) {
			t.Errorf("Expected Len to be %d, got %d", len(tt.expected), size)
		}
		if !reflect.DeepEqual(s.Frames(), tt.expected) {
			t.Errorf("While parsing %q, got %v, expected %v", tt.frange, s.Frames(), tt.expected)
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
				t.Fatalf("Did not find valid frame for idx %s in frames %v", idx, s.Frames())
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

func TestToRange(t *testing.T) {
	var table = []struct {
		expected         []int
		start, end, step int
	}{
		{[]int{1}, 1, 1, 1},
		{[]int{1, 2, 3, 4, 5}, 1, 5, 1},
		{[]int{1, 3, 5}, 1, 6, 2},
		{[]int{1, 6}, 1, 10, 5},
		{[]int{}, 10, 1, 1},
		{[]int{-5, -4, -3, -2, -1, -0, 1, 2, 3}, -5, 3, 1},
		{[]int{-5, -4, -3, -2, -1, -0, 1, 2, 3}, -5, 3, -5},
	}

	for _, tt := range table {
		actual := toRange(tt.start, tt.end, tt.step)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Got %v, expected %v", actual, tt.expected)
		}
	}
}
