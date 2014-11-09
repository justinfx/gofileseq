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
	"bytes"
	"fmt"
	"regexp"
	"sort"
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

// IsFrameRange returns true if the given string is a valid frame
// range format.  Any padding characters, such as '#' and '@' are ignored.
func IsFrameRange(frange string) bool {
	_, err := frameRangeMatches(frange)
	if err == nil {
		return true
	}
	return false
}

// FramesToFrameRange takes a slice of frame numbers and
// compresses them into a frame range string.
//
// If sorted == true, pre-sort the frames instead of respecting
// their current order in the range.
//
// If zfill > 1, then pad out each number with "0" to the given
// total width.
func FramesToFrameRange(frames []int, sorted bool, zfill int) string {
	count := len(frames)
	if count == 0 {
		return ""
	}

	if count == 1 {
		return zfillInt(frames[0], zfill)
	}

	if sorted {
		sort.Ints(frames)
	}

	var i, frame, step int
	var start, end string
	var buf bytes.Buffer

	// Keep looping until all frames are consumed
	for len(frames) > 0 {
		count = len(frames)
		// If we get to the last element, just write it
		// and end
		if count <= 2 {
			for i, frame = range frames {
				if buf.Len() > 0 {
					buf.WriteString(",")
				}
				buf.WriteString(zfillInt(frame, zfill))
			}
			break
		}
		// At this point, we have 3 or more frames to check.
		// Scan the current window of the slice to see how
		// many frames we can consume into a group
		step = frames[1] - frames[0]
		for i = 0; i < len(frames)-1; i++ {
			// We have scanned as many frames as we can
			// for this group. Now write them and stop
			// looping on this window
			if (frames[i+1] - frames[i]) != step {
				break
			}
		}

		// Subsequent groups are comma-separated
		if buf.Len() > 0 {
			buf.WriteString(",")
		}

		// We only have a single frame to write for this group
		if i == 0 {
			buf.WriteString(zfillInt(frames[0], zfill))
			frames = frames[1:]
			continue
		}

		// First do a check to see if we could have gotten a larger range
		// out of subsequent values with a different step size
		if i == 1 && count > 3 {
			// Check if the next two pairwise frames have the same step.
			// If so, then it is better than our current grouping.
			if (frames[2] - frames[1]) == (frames[3] - frames[2]) {
				// Just consume the first frame, and allow the next
				// loop to scan the new stepping
				buf.WriteString(zfillInt(frames[0], zfill))
				frames = frames[1:]
				continue
			}
		}

		// Otherwise write out this step range
		start = zfillInt(frames[0], zfill)
		end = zfillInt(frames[i], zfill)
		buf.WriteString(fmt.Sprintf("%s-%s", start, end))
		if step > 1 {
			buf.WriteString(fmt.Sprintf("x%d", step))
		}
		frames = frames[i+1:]
	}

	return buf.String()
}

// PadFrameRange takes a frame range string and returns a
// new range with each number padded out with zeros to a given width
func PadFrameRange(frange string, pad int) string {
	// We don't need to do anything if they gave us
	// an invalid pad number
	if pad < 2 {
		return frange
	}

	size := strings.Count(frange, ",") + 1
	parts := make([]string, size, size)

	for i, part := range strings.Split(frange, ",") {

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
				parts[i] = zfillString(matched[0], pad)
			case 2:
				parts[i] = fmt.Sprintf("%s-%s",
					zfillString(matched[0], pad),
					zfillString(matched[1], pad))
			case 4:
				parts[i] = fmt.Sprintf("%s-%s%s%s",
					zfillString(matched[0], pad),
					zfillString(matched[1], pad),
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

// PaddingChars returns the proper padding characters,
// given an amount of padding.
func PaddingChars(pad int) string {
	switch {
	case pad <= 0:
		return "@"
	case pad%4 == 0:
		return strings.Repeat("#", pad/4)
	default:
		return strings.Repeat("@", pad)
	}
}

// frameRangeMatches breaks down the string frame range
// into groups of range matches, for further processing.
func frameRangeMatches(frange string) ([][]string, error) {
	for k := range padding {
		frange = strings.Replace(frange, k, "", -1)
	}

	var (
		matched bool
		match   []string
		rx      *regexp.Regexp
	)

	frange = stripSpace.ReplaceAllLiteralString(frange, "")

	// For each comma-sep component, we will parse a frame range
	parts := strings.Split(frange, ",")
	size := len(parts)
	matches := make([][]string, size, size)

	for i, part := range parts {

		matched = false

		// Build up frames for all comma-sep components
		for _, rx = range rangePatterns {
			if match = rx.FindStringSubmatch(part); match == nil {
				continue
			}
			matched = true
			matches[i] = match[1:]
		}

		// If any component of the comma-sep frame range fails to
		// parse, we bail out
		if !matched {
			err := fmt.Errorf("Failed to parse frame range: %s on part %q", frange, part)
			return nil, err
		}
	}

	return matches, nil
}

// Left pads a string to a given with, using "0".
// If the string begins with a negative "-" character, then
// padding is inserted between the "-" and the remaining characters.
func zfillString(src string, z int) string {
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

// Left pads an int to a given with, using "0".
// If the string begins with a negative "-" character, then
// padding is inserted between the "-" and the remaining characters.
func zfillInt(src int, z int) string {
	if z < 2 {
		return strconv.Itoa(src)
	}
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", z), src)
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
	val, err := strconv.Atoi(s)
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

// Return the min/max frames from an unsorted list
func minMaxFrame(frames []int) (int, int) {
	srcframes := make([]int, len(frames), len(frames))
	copy(srcframes, frames)
	sort.Ints(srcframes)
	min, max := srcframes[0], srcframes[len(srcframes)-1]
	return min, max
}
