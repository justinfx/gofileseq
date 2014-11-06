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
	"sort"
	"strconv"
	"strings"
)

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
		return strconv.Itoa(frames[0])
	}

	if sorted {
		sort.Ints(frames)
	}

	zf := zfillString

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
				buf.WriteString(zf(strconv.Itoa(frame), zfill))
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
			buf.WriteString(zf(strconv.Itoa(frames[0]), zfill))
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
				buf.WriteString(zf(strconv.Itoa(frames[0]), zfill))
				frames = frames[1:]
				continue
			}
		}

		// Otherwise write out this step range
		start = zf(strconv.Itoa(frames[0]), zfill)
		end = zf(strconv.Itoa(frames[i]), zfill)
		buf.WriteString(fmt.Sprintf("%s-%s", start, end))
		if step > 1 {
			buf.WriteString(fmt.Sprintf("x%d", step))
		}
		frames = frames[i+1:]
	}

	return buf.String()
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
