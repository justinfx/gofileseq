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
)

func init() {
	padding = map[string]int{
		"#": 4,
		"@": 1,
	}

	// Regular expression patterns for matching frame set strings.
	// Examples:
	//     1-100
	//     100
	//     1-100x5
	rangePatterns = []*regexp.Regexp{
		regexp.MustCompile(`^(\-?[0-9]+)\-(\-?[0-9]+)$`),
		regexp.MustCompile(`^(\-?[0-9]+)$`),
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

	frameSet := &FrameSet{}
	frameSet.frames = []int{}
	frameSet.set = map[int]struct{}{}

	for _, part := range strings.Split(frange, ",") {

		matched = false
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
		if !matched {
			err = fmt.Errorf("Failed to parse frame range: %s on part %q", frange, part)
			return nil, err
		}
	}

	return frameSet, nil
}

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
		end++
		size := end - start
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
