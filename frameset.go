package fileseq

import (
	"fmt"

	"github.com/justinfx/gofileseq/ranges"
)

// FrameSet wraps a sequence of frames in container that
// exposes array-like operations and queries, after parsing
// the given frame range string.
type FrameSet struct {
	frange   string
	rangePtr *ranges.InclusiveRanges
}

// Create a new FrameSet from a given frame range string
// Returns an error if the frame range could not be parsed.
func NewFrameSet(frange string) (*FrameSet, error) {
	// Process the frame range and get a slice of match slices
	matches, err := frameRangeMatches(frange)
	if err != nil {
		return nil, err
	}

	frameSet := &FrameSet{frange, &ranges.InclusiveRanges{}}

	// Process each slice match and add it to the frame set
	for _, match := range matches {
		if err = frameSet.handleMatch(match); err != nil {
			return nil, err
		}
	}

	return frameSet, nil
}

// Process a rangePattern match group
func (s *FrameSet) handleMatch(match []string) error {
	switch len(match) {

	// Single frame match
	case 1:
		f, err := parseInt(match[0])
		if err != nil {
			return err
		}
		s.rangePtr.AppendUnique(f, f, 1)

	// Simple frame range
	case 2:
		start, err := parseInt(match[0])
		if err != nil {
			return err
		}
		end, err := parseInt(match[1])
		if err != nil {
			return err
		}

		// Handle descending frame ranges, like 10-1
		var inc int
		if start > end {
			inc = -1
		} else {
			inc = 1
		}

		s.rangePtr.AppendUnique(start, end, inc)

	// Complex frame range
	case 4:
		var (
			err               error
			mod               string
			start, end, chunk int
		)
		chunk, err = parseInt(match[3])
		if err != nil {
			return err
		}
		if chunk == 0 {
			return fmt.Errorf("Failed to parse part of range %v. "+
				"Encountered invalid 0 value", match[3])
		}
		if start, err = parseInt(match[0]); err != nil {
			return err
		}
		if end, err = parseInt(match[1]); err != nil {
			return err
		}
		if mod = match[2]; !isModifier(mod) {
			return fmt.Errorf("%q is not one of the valid modifier 'xy:'", mod)
		}

		switch mod {
		case `x`:
			s.rangePtr.AppendUnique(start, end, chunk)

		case `y`:
			// TODO: Add proper support for adding inverse of range.
			// This approach will add excessive amounts of singe
			// range elements. They could be compressed into chunks
			skip := start
			aRange := ranges.NewInclusiveRange(start, end, 1)
			var val int
			for it := aRange.IterValues(); !it.IsDone(); {
				val = it.Next()
				if val == skip {
					skip += chunk
					continue
				}
				s.rangePtr.AppendUnique(val, val, 1)
			}

		case `:`:
			for ; chunk > 0; chunk-- {
				s.rangePtr.AppendUnique(start, end, chunk)
			}
		}

	default:
		return fmt.Errorf("Unexpected match []string size: %v", match)
	}

	return nil
}

// String implements Stringer, by returning the frame
// range string
func (s *FrameSet) String() string {
	frange := s.FrameRange()
	return frange
}

// Len returns the number of frames in the set
func (s *FrameSet) Len() int {
	return s.rangePtr.Len()
}

// Index returns the index position of the frame value
// within the frame set.
// If the given frame does not exist, then return -1
func (s *FrameSet) Index(frame int) int {
	return s.rangePtr.Index(frame)
}

// Frame returns the frame number value for a given index into
// the frame set.
// If the index is outside the bounds of the frame set range,
// then an error is returned.
func (s *FrameSet) Frame(index int) (int, error) {
	return s.rangePtr.Value(index)
}

// Frames returns a slice of the frame numbers that were parsed
// from the original frame range string.
//
// Warning: This allocates a slice containing number of elements
// equal to the Len() of the range. If the sequence is massive,
// it could hit slice limits. Better to use IterFrames()
func (s *FrameSet) Frames() []int {
	size := s.rangePtr.Len()
	frames := make([]int, size, size)
	i := 0
	for it := s.rangePtr.IterValues(); !it.IsDone(); {
		frames[i] = it.Next()
		i++
	}
	return frames
}

// HasFrame returns true if the frameset contains the given
// frame value.
func (s *FrameSet) HasFrame(frame int) bool {
	return s.rangePtr.Contains(frame)
}

// Start returns the first frame in the frame set
func (s *FrameSet) Start() int {
	return s.rangePtr.Start()
}

// End returns the last frame in the frame set
func (s *FrameSet) End() int {
	return s.rangePtr.End()
}

// FrameRange returns the range string that was used
// to initialize the FrameSet
func (s *FrameSet) FrameRange() string {
	return s.frange
}

// FrameRangePadded returns the range string that was used
// to initialize the FrameSet, with each number padded out
// with zeros to a given width
func (s *FrameSet) FrameRangePadded(pad int) string {
	return PadFrameRange(s.frange, pad)
}

// Invert returns a new FrameSet that represents
// all frames *not* within the current FrameSet. That is, it
// will create a range that "fills in" the current one.
func (s *FrameSet) Invert() *FrameSet {
	ptr := s.rangePtr.Inverted()
	return &FrameSet{ptr.String(), ptr}
}

// InvertedFrameRange returns a new frame range string that represents
// all frames *not* within the current FrameSet. That is, it
// will create a range that "fills in" the current one.
func (s *FrameSet) InvertedFrameRange(pad int) string {
	frange := s.rangePtr.Inverted().String()
	if pad > 1 {
		frange = PadFrameRange(frange, pad)
	}
	return frange
}

// Normalize returns a new sorted and compacted FrameSet
func (s *FrameSet) Normalize() *FrameSet {
	ptr := s.rangePtr.Normalized()
	return &FrameSet{ptr.String(), ptr}
}
