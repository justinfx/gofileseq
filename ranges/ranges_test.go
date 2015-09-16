package ranges

import (
	"reflect"
	"testing"
)

func TestRangeInit(t *testing.T) {
	table := []struct {
		start, end, step int
		rEnd, rLen       int
		rStr             string
	}{
		{1, 10, 1,
			10, 10, "1-10"}, // 0
		{10, 1, -1,
			1, 10, "10-1"}, // 1
		{-10, 20, 1,
			20, 31, "-10-20"}, // 2
		{20, -10, -1,
			-10, 31, "20--10"}, // 3

		{1, 10, 3,
			10, 4, "1-10x3"}, // 4
		{10, -10, -3,
			-8, 7, "10--8x-3"}, // 5

		{-100, -33, 7,
			-37, 10, "-100--37x7"}, // 6
		{-33, -100, -7,
			-96, 10, "-33--96x-7"}, // 7

		{1, 10, 0,
			10, 10, "1-10"}, // 8
		{10, 1, 0,
			1, 10, "10-1"}, // 9
	}

	var actual int
	for i, tt := range table {
		r := NewInclusiveRange(tt.start, tt.end, tt.step)
		actual = r.Start()
		if actual != tt.start {
			t.Errorf("Test %d: Expected start %d; got %d", i, tt.start, actual)
		}
		actual = r.End()
		if actual != tt.rEnd {
			t.Errorf("Test %d: Expected end %d; got %d", i, tt.rEnd, actual)
		}
		actual = r.Step()
		expected := tt.step
		if expected == 0 {
			if r.End() >= r.Start() {
				expected = 1
			} else {
				expected = -1
			}
		}
		if actual != expected {
			t.Errorf("Test %d: Expected step %d; got %d", i, expected, actual)
		}
		actual = r.Len()
		if actual != tt.rLen {
			t.Errorf("Test %d: Expected len %d; got %d", i, tt.rLen, actual)
		}
		actualStr := r.String()
		if actualStr != tt.rStr {
			t.Errorf("Test %d: Expected %q; got %q", i, tt.rStr, actualStr)
		}
	}
}

func TestRangeLookups(t *testing.T) {
	table := []struct {
		start, end, step int
		indexToValue     map[int]int
		valueToIndex     map[int]int
	}{
		// 0
		{1, 10, 1,
			map[int]int{0: 1, 5: 6, 9: 10},
			map[int]int{1: 0, 6: 5, 10: 9, -1: -1, 11: -1}},
		// 1
		{10, -10, -3,
			map[int]int{0: 10, 3: 1, 6: -8},
			map[int]int{10: 0, 1: 3, -8: 6, 3: -1, -7: -1}},
	}

	for i, tt := range table {
		r := NewInclusiveRange(tt.start, tt.end, tt.step)

		for idx, expected := range tt.indexToValue {
			actual, err := r.Value(idx)
			if err != nil {
				t.Errorf("Test %d: Failed Value(idx): %s", err.Error())
				continue
			}
			if actual != expected {
				t.Errorf("Test %d: Expected value %d for index %d; got %d",
					i, expected, idx, actual)
			}
		}

		for val, expected := range tt.valueToIndex {
			actual := r.Index(val)
			if actual != expected {
				t.Errorf("Test %d: Expected index %d for value %d; got %d",
					i, expected, val, actual)
			}
		}
	}
}

func TestRangeIterValues(t *testing.T) {
	table := []struct {
		start, end, step int
		expected         []int
	}{
		// 0
		{-5, 5, 1,
			[]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
		// 1
		{-10, 10, 3,
			[]int{-10, -7, -4, -1, 2, 5, 8}},
		// 2
		{20, -30, -7,
			[]int{20, 13, 6, -1, -8, -15, -22, -29}},
	}

	for i, tt := range table {
		r := NewInclusiveRange(tt.start, tt.end, tt.step)
		var actual []int
		for it := r.IterValues(); !it.IsDone(); {
			actual = append(actual, it.Next())
		}
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Test %d: got %v; expected %v", i, actual, tt.expected)
		}
	}
}

func TestRangesInit(t *testing.T) {
	table := []struct {
		params     [][]int
		rEnd, rLen int
		rStr       string
	}{
		// 0
		{
			[][]int{},
			0, 0, "",
		},
		// 1
		{
			[][]int{{1, 10, 1}, {13, 14, 1}, {17, 17, 1}, {1, 30, 3}},
			28, 18, "1-10,13-14,17,16-28x3",
		},
		// 2
		{
			[][]int{{-5, 5, 2}, {10, 10, 1}, {-10, 10, 1}},
			9, 21, "-5-5x2,10,-10--6,-4,-2,0,2,4,6-9",
		},
	}

	var actual int
	for i, tt := range table {
		var r InclusiveRanges
		for _, p := range tt.params {
			r.AppendUnique(p[0], p[1], p[2])
		}
		var start int
		if len(tt.params) > 0 {
			start = tt.params[0][0]
		}

		actual = r.Start()
		if actual != start {
			t.Errorf("Test %d: Expected start %d; got %d", i, tt.params[0], actual)
		}
		actual = r.End()
		if actual != tt.rEnd {
			t.Errorf("Test %d: Expected end %d; got %d", i, tt.rEnd, actual)
		}
		actual = r.Len()
		if actual != tt.rLen {
			t.Errorf("Test %d: Expected len %d; got %d", i, tt.rLen, actual)
		}
		actualStr := r.String()
		if actualStr != tt.rStr {
			t.Errorf("Test %d: Expected %q; got %q", i, tt.rStr, actualStr)
		}
	}
}

func TestRangesLookup(t *testing.T) {
	table := []struct {
		params       [][]int
		indexToValue map[int]int
		valueToIndex map[int]int
	}{
		// 0
		{
			[][]int{{1, 10, 1}, {13, 14, 1}, {17, 17, 1}, {1, 30, 3}},
			map[int]int{0: 1, 10: 13, 12: 17, 17: 28},
			map[int]int{1: 0, 13: 10, 17: 12, 28: 17, -1: -1, 15: -1, 30: -1},
		},
		// 1
		{
			[][]int{{-5, 5, 2}, {10, 10, 1}, {-10, 10, 1}},
			map[int]int{0: -5, 5: 5, 6: 10, 11: -6, 20: 9},
			map[int]int{-5: 0, 5: 5, 10: 6, -6: 11, 9: 20, -11: -1, 11: -1},
		},
	}

	for i, tt := range table {
		var r InclusiveRanges
		for _, p := range tt.params {
			r.AppendUnique(p[0], p[1], p[2])
		}

		for idx, expected := range tt.indexToValue {
			actual, err := r.Value(idx)
			if err != nil {
				t.Errorf("Test %d: Failed Value(idx): %s", i, err.Error())
				continue
			}
			if actual != expected {
				t.Errorf("Test %d: Expected value %d for index %d; got %d",
					i, expected, idx, actual)
			}
		}

		for val, expected := range tt.valueToIndex {
			actual := r.Index(val)
			if actual != expected {
				t.Errorf("Test %d: Expected index %d for value %d; got %d",
					i, expected, val, actual)
			}
		}
	}
}

func TestRangesIterValues(t *testing.T) {
	table := []struct {
		params   [][]int
		expected []int
	}{
		{ // 0
			[][]int{{-5, 5, 2}, {10, 10, 1}, {-10, 10, 1}},
			[]int{
				-5, -3, -1, 1, 3, 5, 10, -10, -9, -8,
				-7, -6, -4, -2, 0, 2, 4, 6, 7, 8, 9,
			},
		},
		{ // 1
			[][]int{{-3, 3, 1}, {-5, 5, 1}},
			[]int{-3, -2, -1, 0, 1, 2, 3, -5, -4, 4, 5},
		},
		{ // 2
			[][]int{{10, 1, -2}, {-5, 10, 2}},
			[]int{10, 8, 6, 4, 2, -5, -3, -1, 1, 3, 5, 7, 9},
		},
	}

	for i, tt := range table {
		var r InclusiveRanges
		for _, p := range tt.params {
			r.AppendUnique(p[0], p[1], p[2])
		}
		var actual []int
		for it := r.IterValues(); !it.IsDone(); {
			actual = append(actual, it.Next())
		}
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Test %d: got %v; expected %v", i, actual, tt.expected)
		}
	}
}

func TestRangesInverted(t *testing.T) {
	table := []struct {
		params   [][]int
		expected string
	}{
		{ // 0
			[][]int{{1, 10, 1}},
			"",
		},
		{ // 1
			[][]int{{1, 10, 2}},
			"2-8x2",
		},
		{ // 2
			[][]int{{-6, 0, 2}, {2, 4, 1}, {10, 16, 3}},
			"-5-1x2,5-9,11-12,14-15",
		},
		{ // 3
			[][]int{{16, 10, 3}, {4, 2, 1}, {0, -6, 2}},
			"-5-1x2,5-9,11-12,14-15",
		},
		{ // 5
			[][]int{{4, 2, 1}, {0, -6, 2}, {16, 10, 3}, {2, 4, 1}},
			"-5-1x2,5-9,11-12,14-15",
		},
		{ // 6
			[][]int{{1, 10, 2}, {50, 50, 1}, {60, 62, 2}, {70, 74, 2}},
			"2-10x2,11-49,51-59,61-63x2,64-69,71-73x2",
		},
	}

	for i, tt := range table {
		var r InclusiveRanges
		for _, p := range tt.params {
			r.AppendUnique(p[0], p[1], p[2])
		}
		actual := r.Inverted().String()
		if actual != tt.expected {
			t.Errorf("Test %d: got %q; expected %q", i, actual, tt.expected)
		}
	}
}
