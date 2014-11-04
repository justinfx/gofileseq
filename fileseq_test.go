package fileseq

import (
	"reflect"
	"testing"
)

func TestNewFrameSet(t *testing.T) {
	var table = []struct {
		expected []int
		frange   string
	}{
		{[]int{10}, "10"},
		{[]int{1, 2, 3, 4, 5}, "1-5"},
		{[]int{1, 2, 3, 20, 30, 31, 32, 33}, "1-3,20,30-33"},
		{[]int{1, 6}, "1-10x5"},
		{[]int{1, 3, 5, 7, 9}, "1-10x2"},
		{[]int{1, 3, 5, 7, 9, 50, 60, 61, 62, 70, 72, 74}, "1-10x2,50,60-62,70-74x2"},
		{[]int{1, 5, 9, 4, 7, 10, 3, 2, 6, 8}, "1-10:4"},
		{[]int{2, 4, 6, 8, 10}, "1-10y2"},
		{[]int{2, 3, 4, 6, 7, 8, 10, 11, 12, 14, 15}, "1-15y4"},
	}

	for _, tt := range table {
		s, err := NewFrameSet(tt.frange)
		if err != nil {
			t.Fatalf("Failed to parse %q: %s", tt.frange, err.Error())
		}
		if !reflect.DeepEqual(s.frames, tt.expected) {
			t.Errorf("While parsing %q, got %v, expected %v", tt.frange, s.frames, tt.expected)
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
