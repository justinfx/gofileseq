/*
package fileseq is a library for parsing file sequence strings commonly
used in VFX and animation applications.

Frame Range Shorthand

Support for:

	Standard: 1-10
	Comma Delimted: 1-10,10-20
	Chunked: 1-100x5
	Filled: 1-100y5
	Staggered: 1-100:3 (1-100x3, 1-100x2, 1-100)
	Negative frame numbers: -10-100
	Padding: #=4 padded, @=single pad

*/
package fileseq

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	padding       map[string]int
	rangePatterns []*regexp.Regexp
	splitPattern  *regexp.Regexp
	stripSpace    *regexp.Regexp
)

func init() {
	padding = map[string]int{
		"#": 4,
		"@": 1,
	}

	stripSpace = regexp.MustCompile(`(\s+)`)

	// Regular expression patterns for matching frame set strings.
	// Examples:
	//     1-100
	//     100
	//     1-100x5
	rangePatterns = []*regexp.Regexp{
		// Frame range:  1-10
		regexp.MustCompile(`^(\-?[0-9]+)\-(\-?[0-9]+)$`),
		// Single frame:  10
		regexp.MustCompile(`^(\-?[0-9]+)$`),
		// Complex range:  1-10x2
		regexp.MustCompile(`^(\-?[0-9]+)\-(\-?[0-9]+)([:xy]{1})([0-9]+)$`),
	}

	// Regular expression for matching a file sequence string.
	// Example:
	//     /film/shot/renders/hero_bty.1-100#.exr
	splitPattern = regexp.MustCompile(`([\:xy\-0-9,]*)([\#\@]+)`)
}

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
	var (
		match   []string
		rx      *regexp.Regexp
		err     error
		matched bool
	)

	frange = stripSpace.ReplaceAllLiteralString(frange, "")

	frameSet := &FrameSet{
		frange: frange,
		frames: []int{},
		set:    map[int]struct{}{},
	}

	// For each comma-sep component, we will parse a frame range
	for _, part := range strings.Split(frange, ",") {

		matched = false

		// Build up frames for all comma-sep components
		for _, rx = range rangePatterns {
			if match = rx.FindStringSubmatch(part); match == nil {
				continue
			}
			matched = true
			if err = frameSet.handleMatch(match[1:]); err != nil {
				return nil, err
			}
			break
		}

		// If any component of the comma-sep frame range fails to
		// parse, we bail out
		if !matched {
			err = fmt.Errorf("Failed to parse frame range: %s on part %q", frange, part)
			return nil, err
		}
	}

	return frameSet, nil
}

// Process a rangePattern match group
func (s *FrameSet) handleMatch(match []string) error {
	var parsed []int

	switch len(match) {

	case 1:
		f, err := parseInt(match[0])
		if err != nil {
			return err
		}
		parsed = []int{f}

	case 2:
		start, err := parseInt(match[0])
		if err != nil {
			return err
		}
		end, err := parseInt(match[1])
		if err != nil {
			return err
		}
		size := (end + 1) - start
		parsed = make([]int, size, size)
		for i := range parsed {
			parsed[i] = start
			start++
		}

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
	// We don't need to do anything if they gave us
	// an invalid pad number
	if pad < 2 {
		return s.frange
	}

	size := strings.Count(s.frange, ",") + 1
	parts := make([]string, size, size)

	for i, part := range strings.Split(s.frange, ",") {

		didMatch := false

		for _, rx := range rangePatterns {
			matched := rx.FindStringSubmatch(part)
			if len(matched) == 0 {
				continue
			}
			matched = matched[1:]
			size = len(matched)
			switch size {
			case 1:
				parts[i] = zfill(matched[0], pad)
			case 2:
				parts[i] = fmt.Sprintf("%s-%s", zfill(matched[0], pad), zfill(matched[1], pad))
			case 4:
				parts[i] = fmt.Sprintf("%s-%s%s%s",
					zfill(matched[0], pad), zfill(matched[1], pad),
					matched[2], matched[3])
			default:
				// No match. Try the next pattern
				continue
			}
			// If we got here, we matched a case and can stop
			// checking the rest of the patterns
			didMatch = true
			break
		}
		// If we didn't match one of our expected patterns
		// then just take the original part and add it unmodified
		if !didMatch {
			parts = append(parts, part)
		}
	}
	return strings.Join(parts, ",")
}

// Left pads a string to a given with, using "0".
// If the string begins with a negative "-" character, then
// padding is inserted between the "-" and the remaining characters.
func zfill(src string, z int) string {
	size := len(src)
	if size >= z {
		return src
	}

	fill := strings.Repeat("0", z-size)
	if strings.HasPrefix(src, "-") {
		return fmt.Sprintf("-%s%s", fill, src[1:])
	}
	return fmt.Sprintf("%s%s", fill, src)
}

// Expands a start, end, and stepping value
// into the full range of int values.
func toRange(start, end, step int) []int {
	nums := []int{}
	if end < start {
		return nums
	}
	if step < 1 {
		step = 1
	}
	for i := start; i <= end; {
		nums = append(nums, i)
		i += step
	}
	return nums
}

// Parse an int from a specific part of a frame
// range string component
func parseInt(s string) (int, error) {
	val, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("Failed to parse int from part of range %q", s)
	}
	return int(val), nil
}

// Return whether a string component from a frame
// range string is a valid modifier symbol
func isModifier(s string) bool {
	return len(s) == 1 && strings.ContainsAny(s, "xy:")
}
