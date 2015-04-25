package fileseq

import "fmt"

// FrameSet wraps a sequence of frames in container that
// exposes array-like operations and queries, after parsing
// the given frame range string.
type FrameSet struct {
	frange string
	frames []int
	set    map[int]struct{}
}

// Create a new FrameSet from a given frame range string
// Returns an error if the frame range could not be parsed.
func NewFrameSet(frange string) (*FrameSet, error) {
	var err error

	frameSet := &FrameSet{
		frange: frange,
		frames: []int{},
		set:    map[int]struct{}{},
	}

	// Process the frame range and get a slice of match slices
	matches, err := frameRangeMatches(frange)
	if err != nil {
		return nil, err
	}

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
	var parsed []int

	switch len(match) {

	// Single frame match
	case 1:
		f, err := parseInt(match[0])
		if err != nil {
			return err
		}
		parsed = []int{f}

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
		var size, inc int
		if start > end {
			inc = -1
			size = (start + 1) - end
		} else {
			inc = 1
			size = (end + 1) - start
		}

		parsed = make([]int, size, size)
		for i := range parsed {
			parsed[i] = start
			start += inc
		}

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
			parsed = toRange(start, end, chunk)

		case `y`:
			parsed = []int{}
			skip := start
			for _, val := range toRange(start, end, 1) {
				if val == skip {
					skip += chunk
					continue
				}
				parsed = append(parsed, val)
			}

		case `:`:
			parsed = []int{}
			for ; chunk >= 0; chunk-- {
				staggered := toRange(start, end, chunk)
				parsed = append(parsed, staggered...)
			}
		}

	default:
		return fmt.Errorf("Unexpected match []string size: %v", match)
	}

	s.addFrames(parsed)
	return nil
}

// String implements Stringer, by returning the frame
// range string
func (s *FrameSet) String() string {
	return s.FrameRange()
}

// Len returns the number of frames in the set
func (s *FrameSet) Len() int {
	return len(s.frames)
}

// Adds a slice of frame numbers to the internal
// array, only if they have not yet been added.
func (s *FrameSet) addFrames(frames []int) {
	var exists bool
	for _, f := range frames {
		if _, exists = s.set[f]; exists {
			continue
		}
		s.set[f] = struct{}{}
		s.frames = append(s.frames, f)
	}
}

// Index returns the index position of the frame value
// within the frame set.
// If the given frame does not exist, then return -1
func (s *FrameSet) Index(frame int) int {
	for i, v := range s.frames {
		if v == frame {
			return i
		}
	}
	return -1
}

// Frame returns the frame number value for a given index into
// the frame set.
// If the index is outside the bounds of the frame set range,
// then an error is returned.
func (s *FrameSet) Frame(index int) (int, error) {
	if index < 0 || index >= len(s.frames) {
		return 0, fmt.Errorf("index %d is outside the bounds of the frame set 0-%d",
			index, len(s.frames))
	}
	return s.frames[index], nil
}

// Frames returns a slice of the frame numbers that were parsed
// from the original frame range string
func (s *FrameSet) Frames() []int {
	return s.frames
}

// HasFrame returns true if the frameset contains the given
// frame value.
func (s *FrameSet) HasFrame(frame int) bool {
	if s.Index(frame) == -1 {
		return false
	}
	return true
}

// Start returns the first frame in the frame set
func (s *FrameSet) Start() int {
	if len(s.frames) == 0 {
		return 0
	}
	return s.frames[0]
}

// End returns the last frame in the frame set
func (s *FrameSet) End() int {
	if len(s.frames) == 0 {
		return 0
	}
	return s.frames[len(s.frames)-1]
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

// InvertedFrameRange returns a new frame range that represents
// all frames *not* within the current frame range. That is, it
// will create a range that "fills in" the current one.
func (s *FrameSet) InvertedFrameRange(pad int) string {
	size := len(s.frames)
	if size < 2 {
		return ""
	}

	min, max := minMaxFrame(s.frames)
	if max-min <= 1 {
		return ""
	}

	// Compute how big our result set will be
	size = (max - min + 1) - size
	if size < 1 {
		return ""
	}

	frames := make([]int, size, size)
	var idx int
	var exists bool
	for val := min + 1; val < max; val++ {
		if _, exists = s.set[val]; exists {
			continue
		}
		frames[idx] = val
		idx++
	}
	return FramesToFrameRange(frames, false, pad)
}

// Normalize returns a new sorted and compacted FrameSet
func (s *FrameSet) Normalize() *FrameSet {
	ret, _ := NewFrameSet(FramesToFrameRange(s.frames, true, 0))
	return ret
}
