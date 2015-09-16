/*
Package ranges provides objects that can represent
simple integer ranges with an optional stepping, as
well as complex non-contiguous ranges.

Simple range:

    1-10
    20-100x3
    -10-50

Complex ranges:

    1-10,20-40x2,30,80-100x3

*/
package ranges

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

// Iterator is an interface type that can
// iterate over int values.
type Iterator interface {
	// The Iterator has another valid value
	// to produce, while IsDone() returns true
	IsDone() bool
	// Next retrieves the next available value,
	// and advances the Iterator. It is only
	// valid to call Next() while IsDone()
	// returns true
	Next() int
}

// InclusiveRange represents a start and end
// value, and a stepping amount between each
// value in the range. The range includes the end value.
// It dynamically acts like a slice of values
// in an integer range, without needing to store
// the discreet values. So it can represent very
// large ranges.
type InclusiveRange struct {
	start int
	end   int
	step  int

	cachedEnd   int
	isEndCached bool

	cachedLen   int
	isLenCached bool
}

// NewInclusiveRange creates a new InclusiveRange instance
func NewInclusiveRange(start, end, step int) *InclusiveRange {
	if step == 0 {
		if start <= end {
			step = 1

		} else {
			step = -1
		}
	}

	r := &InclusiveRange{
		start: start,
		end:   end,
		step:  step,
	}
	return r
}

// String returns a formatted string representation
// of the integer range
func (r *InclusiveRange) String() string {
	var buf bytes.Buffer

	// Always for a single value
	buf.WriteString(strconv.Itoa(r.Start()))

	// If we have a range, express the end value
	if r.End() != r.Start() {
		buf.WriteString(`-`)
		buf.WriteString(strconv.Itoa(r.End()))

		// Express the stepping, if its not 1
		step := r.Step()
		if step > 1 || step < -1 {
			buf.WriteString(`x`)
			buf.WriteString(strconv.Itoa(r.Step()))
		}
	}
	return buf.String()
}

// Start returns the start of the range
func (r *InclusiveRange) Start() int {
	return r.start
}

// End returns the end of the range. This value may
// be different from the end value given when the
// range was first initialized, since it takes into
// account the stepping value. The end value may be
// shifted to the closest valid value within the
// stepped range.
func (r *InclusiveRange) End() int {
	if r.isEndCached {
		return r.cachedEnd
	}

	r.isEndCached = true

	// If we aren't stepping, or we don't have
	// a full range, then just use the end value
	if r.step == 1 || r.step == -1 || r.start == r.end {
		r.cachedEnd = r.end
		return r.cachedEnd
	}

	// If the step is in the wrong direction,
	// compared to the range direction, then
	// just use the start as the end.
	if (r.end < r.start) && r.step < (r.end-r.start) {
		r.cachedEnd = r.start
		return r.cachedEnd

	} else if (r.end > r.start) && r.step > (r.end-r.start) {
		r.cachedEnd = r.start
		return r.cachedEnd
	}

	// Calculate the end, taking into account the stepping
	r.cachedEnd = r.closestInRange(r.end, r.start, r.end, r.step)
	return r.cachedEnd
}

// Step returns the stepping value used in the range
func (r *InclusiveRange) Step() int {
	return r.step
}

// Len returns the number of values in the range
func (r *InclusiveRange) Len() int {
	if r.isLenCached {
		return r.cachedLen
	}

	// Offset by one to include the end value
	diff := math.Abs(float64(r.end-r.start)) + 1
	r.cachedLen = int(math.Ceil(diff / math.Abs(float64(r.step))))
	r.isLenCached = true
	return r.cachedLen
}

// Min returns the smallest value in the range
func (r *InclusiveRange) Min() int {
	start := r.Start()
	end := r.End()
	if start < end {
		return start
	}
	return end
}

// Max returns the highest value in the range
func (r *InclusiveRange) Max() int {
	start := r.Start()
	end := r.End()
	if start > end {
		return start
	}
	return end
}

// Contains returns true if the given value is a valid
// value within the value range.
func (r *InclusiveRange) Contains(value int) bool {
	// If we attempt to find the closest value, given
	// the start of the range and the step, we can check
	// if it is still the same number. If it hasn't changed,
	// then it is in the range.
	closest := r.closestInRange(value, r.start, r.End(), r.step)
	return closest == value
}

// closestInRange finds the closest valid value within the range,
// to a given value. Values outside the range are clipped to either
// the range min or max.
func (*InclusiveRange) closestInRange(value, start, end, step int) int {
	// Possibly clamp the value if it is outside the range
	if end >= start {
		if value < start {
			return start
		} else if value > end {
			return end
		}

	} else {
		if value > start {
			return start
		} else if value < end {
			return end
		}
	}

	// No calculation needed if there is no stepping
	if step == 1 || step == -1 {
		return value
	}

	// Modified the value so that it is a properly stepped
	// increment within the range
	return (((value - start) / step) * step) + start
}

// Value returns the value at the given index in
// the range. If the index is invalid/exceeds the valid
// range, an error will be returned.
func (r *InclusiveRange) Value(idx int) (int, error) {
	if idx < 0 {
		return 0, fmt.Errorf("Index %d is not >= 0", idx)
	}

	// Calculate the value as an offset from the start
	start := r.Start()
	end := r.End()
	step := r.Step()

	val := start + (step * idx)

	if start <= end && (val < start || val > end) {
		return 0, fmt.Errorf("Index %d exceeds max index of %d", idx, r.Len()-1)

	} else if end < start && (val > start || val < end) {
		return 0, fmt.Errorf("Index %d exceeds max index of %d", idx, r.Len()-1)
	}

	return val, nil
}

// Index returns the 0-based index of the first occurrence
// given value, within the range.
// If the value does not exist in the range, a
// value of -1 will be returned
func (f *InclusiveRange) Index(value int) int {
	closest := f.closestInRange(value, f.start, f.End(), f.step)
	if closest != value {
		return -1
	}
	idx := (value - f.start) / f.step
	if idx < 0 {
		idx *= -1
	}
	return idx
}

// IterValues returns an iterator that will loop through
// every value in the range.
// Iterator is valid while IsDone() returns false.
// Next value on valid iterator can be retrieved via Next()
func (f *InclusiveRange) IterValues() Iterator {
	return &inclusiveRangeIt{ptr: f}
}

// tracks the iteration state across an inclusiveRange instance
type inclusiveRangeIt struct {
	ptr *InclusiveRange
	pos int
}

// IsDone returns false while the iterator is valid
// and can produce another value, via a call to Next()
func (it *inclusiveRangeIt) IsDone() bool {
	if it.pos >= it.ptr.Len() {
		return true
	}
	return false
}

// Next() returns the current value in the
// iteration, and advances the interator. Next()
// will produce a valid result as long as a call
// to isDone() returns false.
func (it *inclusiveRangeIt) Next() int {
	val, err := it.ptr.Value(it.pos)
	it.pos++
	if err != nil {
		return it.ptr.End()
	}
	return val
}

// InclusiveRanges is able to represent multiple non-contiguous
// ranges, with non-uniform stepping, ie. 1-10,20-30x2,100,120
// Because range values are evaluated dynamically as
// offsets within the start/stop, they can represent
// very large ranges.
// The zero value is a ready to use instance.
type InclusiveRanges struct {
	blocks []*InclusiveRange
}

// String returns the formatted representation of
// the combination of all internal InclusiveRange instances
func (l *InclusiveRanges) String() string {
	var buf bytes.Buffer
	for i, b := range l.blocks {
		if i > 0 {
			buf.WriteString(`,`)
		}
		buf.WriteString(b.String())
	}
	return buf.String()
}

// Len returns the total number of values across all ranges
func (l *InclusiveRanges) Len() int {
	var totalLen int
	for _, b := range l.blocks {
		totalLen += b.Len()
	}
	return totalLen
}

// Start returns the first value of the first range
func (l *InclusiveRanges) Start() int {
	for _, b := range l.blocks {
		return b.Start()
	}
	return 0
}

// End returns the last value of the last range
func (l *InclusiveRanges) End() int {
	if l.blocks == nil {
		return 0
	}
	return l.blocks[len(l.blocks)-1].End()
}

// Min returns the smallest value in the total range
func (l *InclusiveRanges) Min() int {
	val := l.Start()
	for _, aRange := range l.blocks {
		next := aRange.Min()
		if next < val {
			val = next
		}
	}
	return val
}

// Max returns the highest value in the total range
func (l *InclusiveRanges) Max() int {
	val := l.End()
	for _, aRange := range l.blocks {
		next := aRange.Max()
		if next > val {
			val = next
		}
	}
	return val
}

// NumRanges returns the number of discreet sets
// of ranges that were appended.
func (l *InclusiveRanges) numRanges() int {
	if l.blocks == nil {
		return 0
	}
	return len(l.blocks)
}

// rangeAt returns the underlying InclusiveRange instance
// that was appended, at a given index
func (l *InclusiveRanges) rangeAt(idx int) *InclusiveRange {
	if idx < 0 || idx >= l.numRanges() {
		return nil
	}
	return l.blocks[idx]
}

// Append creates and adds another range of values
// to the total range list.
func (l *InclusiveRanges) Append(start, end, step int) {
	block := NewInclusiveRange(start, end, step)
	l.blocks = append(l.blocks, block)
}

// AppendUnique creates and adds another range of values
// to the total range list. Only unique values from the
// given range are appended to the total range.
func (l *InclusiveRanges) AppendUnique(start, end, step int) {
	if step == 0 {
		return
	}

	subStart := start
	subEnd := start
	subStep := step
	last := start
	pending := 0 // Track unique value count

	// Handle loop test for both increasing
	// and decreasing ranges
	var pred func() bool
	if start <= end {
		if step < 0 {
			step *= -1
		}
		pred = func() bool { return subEnd <= end }
	} else {
		if step > 0 {
			step *= -1
		}
		pred = func() bool { return subEnd >= end }
	}

	for ; pred(); subEnd += step {
		if !l.Contains(subEnd) {
			// Is a unique value in the range
			last = subEnd
			if pending == 0 {
				subStart = last
			}
			pending++
			continue
		}

		if pending == 0 {
			// Nothing to add yet
			continue
		}

		// Current value is already in range.
		// Add previous values
		l.Append(subStart, last, subStep)
		subStart = subEnd + step
		pending = 0
	}

	// Flush the remaining values
	if pending > 0 {
		l.Append(subStart, last, subStep)
	}
}

// Contains returns true if a given value is a valid
// value within the total range.
func (l *InclusiveRanges) Contains(value int) bool {
	for _, b := range l.blocks {
		if b.Contains(value) {
			return true
		}
	}
	return false
}

// Value returns the value at the given index in
// the total range. If the index is invalid/exceeds the valid
// range, an error will be returned.
func (l *InclusiveRanges) Value(idx int) (int, error) {
	if idx < 0 {
		return 0, fmt.Errorf("Index %d is not >= 0", idx)
	}

	var err error
	var val, size, n int

	for _, b := range l.blocks {
		size = b.Len()

		// modify the original index to be within
		// the range of the current range block, and
		// then check that it is a valid index value
		if (idx - n) < size {
			if val, err = b.Value(idx - n); err == nil {
				return val, nil
			}
		}

		// record the offset so we can continue to
		// modify the original index value to the local range
		n += size
	}

	// The previous loop ended in error
	return 0, fmt.Errorf("Index %d exceeds the total range", idx)
}

// Index returns the 0-based index of the first occurrence
// of the given value, within the range.
// If the value does not exist in the range, a
// value of -1 will be returned.
func (l *InclusiveRanges) Index(value int) int {
	var idx, n int

	for _, b := range l.blocks {
		// If the value is within the current block
		// then return the local index, offset by the
		// number of previous values we have tracked
		if idx = b.Index(value); idx >= 0 {
			return idx + n
		}
		// Update the offset for the values we have seen
		n += b.Len()
	}

	// The previous loop ended in error
	return -1
}

// IterValues returns an iterator that will loop through
// every value in the total range.
// Iterator is valid while IsDone() returns false.
// Next value on valid iterator can be retrieved via Next()
func (l *InclusiveRanges) IterValues() Iterator {
	return &inclusiveRangesIt{ptr: l}
}

// Inverted returns a new instance with a range containing
// all values within the start/end that are not in the current range.
// Original ordering is not preserved. New inverted range will be
// in an increasing value.
func (l *InclusiveRanges) Inverted() *InclusiveRanges {
	inverted := &InclusiveRanges{}

	var (
		start    int
		end      int
		step     int
		current  int
		pending  int
		contains bool
	)

	// fmt.Println("Inverting ", l.String())
	totalRange := NewInclusiveRange(l.Min(), l.Max(), 1)

	for it := totalRange.IterValues(); !it.IsDone(); {

		current = it.Next()

		// fmt.Println("Checking:", current, "(pending:", pending, ")")
		contains = l.Contains(current)

		// Short-circuit if we encounter a value that
		// is already in the original sequence.
		if contains {

			// fmt.Println("  Existing value")

			// If we haven't accumulated 2+ values to
			// add, just continue now and keep trying
			if pending < 2 {
				step++
				// fmt.Println("  Increasing step to:", step)
				continue
			}

			// If the step has changed from what we have
			// already accumulated, then add what we have
			// and start a new range.
			if (current + 1 - end) != step {
				// fmt.Println("  Step changed. Adding range:", start, end, step)
				inverted.Append(start, end, step)
				step = 1
				start = current
				pending = 0
			}

			continue
		}

		// fmt.Println("  Unique value")

		// If the value we want to keep is a different
		// stepping from the pending values, add what
		// we have and start a new range again.
		if pending >= 2 && current-end != step {
			// fmt.Println("  Step changed. Adding range:", start, end, step)
			inverted.Append(start, end, step)
			pending = 0
		}

		end = current

		// Start a new range
		if pending == 0 {
			// fmt.Println("  Starting new range")
			start = end
			step = 1
		}

		pending++
		continue
	}

	// Flush the remaining values
	if pending > 0 {
		// fmt.Println("  Flushing and adding remaining range:", start, end, step)
		inverted.Append(start, end, step)
	}

	return inverted
}

// tracks the iteration state across a inclusiveRanges instance
type inclusiveRangesIt struct {
	ptr     *InclusiveRanges
	current Iterator
	pos     int
}

func (it *inclusiveRangesIt) init() bool {
	if it.current == nil {
		block := it.ptr.rangeAt(it.pos)
		if block == nil {
			return false
		}
		it.current = block.IterValues()
	}
	return true
}

// IsDone returns false while the iterator is valid
// and can produce another value, via a call to Next()
func (it *inclusiveRangesIt) IsDone() bool {
	if !it.init() {
		return true
	}

	return it.current.IsDone()
}

// Next() returns the current value in the
// iteration, and advances the interator. Next()
// will produce a valid result as long as a call
// to isDone() returns false.
func (it *inclusiveRangesIt) Next() int {
	if !it.init() {
		return it.ptr.End()
	}

	val := it.current.Next()

	if it.current.IsDone() {
		// Get the next available range
		it.pos++
		it.current = nil
	}

	return val
}
