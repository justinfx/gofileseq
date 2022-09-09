package fileseq

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/justinfx/gofileseq/v2/exp/peg"
)

type PadStyle int

const (
	// Constants defining the style of padding to use
	// when converting between padding characters ('#', '##', '@@@')
	// and their equivalent numeric padding width
	PadStyleHash1   PadStyle = 0 // '#' char == 1
	PadStyleHash4   PadStyle = 1 // '#' char == 4
	PadStyleDefault          = PadStyleHash4
)

var (
	defaultPadding paddingMapper
	padders        map[PadStyle]paddingMapper
)

func init() {
	padders = map[PadStyle]paddingMapper{
		PadStyleHash1: newSingleHashPad(),
		PadStyleHash4: newMultiHashPad(),
	}

	defaultPadding = padders[PadStyleDefault]
}

// paddingMapper defines behavior for converting between
// padding characters and their pad width
type paddingMapper interface {
	// Return all supported padding characters
	AllChars() []string
	// Return the padding string sequence representing a width
	PaddingChars(int) string
	// Return a width for a string of pad characters
	PaddingCharsSize(string) int
}

// multiHashPad is a paddingMapper that uses a single # to
// represent a width of 4 (i.e. Katana)
type multiHashPad struct {
	*paddingMap
}

func newMultiHashPad() multiHashPad {
	padMap := &paddingMap{
		charToSize:  map[string]int{"#": 4, "@": 1},
		defaultChar: "@",
	}

	return multiHashPad{padMap}
}

func (m multiHashPad) PaddingChars(pad int) string {
	switch {
	case pad <= 0:
		return m.defaultChar
	case pad%4 == 0:
		return strings.Repeat("#", pad/4)
	default:
		return strings.Repeat("@", pad)
	}
}

// singleHashPad is a paddingMapper that uses a single # to
// represent a width of 1 (i.e. Nuke)
type singleHashPad struct {
	*paddingMap
}

func newSingleHashPad() singleHashPad {
	padMap := &paddingMap{
		charToSize:  map[string]int{"#": 1, "@": 1},
		defaultChar: "#",
	}

	return singleHashPad{padMap}
}

func (m singleHashPad) PaddingChars(pad int) string {
	if pad <= 0 {
		return m.defaultChar
	}
	return strings.Repeat(m.defaultChar, pad)
}

type paddingMap struct {
	charToSize  map[string]int
	cachedKeys  []string
	defaultChar string
}

func (m *paddingMap) AllChars() []string {
	if m.cachedKeys == nil {
		m.cachedKeys = make([]string, len(m.charToSize))
		i := 0
		for k := range m.charToSize {
			m.cachedKeys[i] = k
			i++
		}
	}
	return m.cachedKeys
}

func (m *paddingMap) PaddingCharsSize(chars string) int {
	if len(chars) == 0 {
		return 0
	}

	// udim is a fixed padding size
	if udimPattern.MatchString(chars) {
		return 4
	}

	// check for alternate padding syntax
	var match []string
	for _, rx := range []*regexp.Regexp{printfPattern, houdiniPattern} {
		if match = rx.FindStringSubmatch(chars); len(match) == 0 {
			continue
		}
		intVal, err := strconv.Atoi(match[1])
		if err != nil || intVal < 1 {
			return 1
		}
		return intVal
	}

	// standard pad chars
	var size int
	for _, char := range chars {
		size += m.charToSize[string(char)]
	}
	return size
}

// PaddingChars returns the proper padding characters,
// given an amount of padding. Uses PadStyleDefault.
func PaddingChars(pad int) string {
	return defaultPadding.PaddingChars(pad)
}

// PadFrameRange takes a frame range string and returns a
// new range with each number padded out with zeros to a given width
func PadFrameRange(frange string, pad int) string {
	// We don't need to do anything if they gave us
	// an invalid pad number
	if pad < 2 {
		return frange
	}

	parts, err := peg.ParseRange(frange)
	if err != nil {
		// invalid
		return frange
	}

	var builder strings.Builder

	for i, part := range parts {
		switch len(part.Cmpts) {
		case 1:
			part.Range = zfillInt(part.Cmpts[0], pad)
		case 2:
			part.Range = fmt.Sprintf("%s-%s",
				zfillInt(part.Cmpts[0], pad),
				zfillInt(part.Cmpts[1], pad))
		case 3:
			part.Range = fmt.Sprintf("%s-%s%s%d",
				zfillInt(part.Cmpts[0], pad),
				zfillInt(part.Cmpts[1], pad),
				part.StepType, part.Cmpts[2])
		default:
			// No match. Try the next pattern
			continue
		}
		if i > 0 {
			builder.WriteString(",")
		}
		builder.WriteString(part.Range)
	}
	return builder.String()
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
