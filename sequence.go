package fileseq

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

// A FileSequence represents a path to a sequence of files,
// and expresses a valid frames range (one that can be parsed
// by FrameSet. It also includes one or more padding characters
// to dictate how much padding the actual file numbers have.
//
// Valid padding characters:
//     @ - 1 pad width (@@@@ is equal to 4 padding)
//     # - 4 pad width (## is equal to 8 padding)
//
// Example paths and padding:
//     /path/to/single_image.0100.jpg
//     /path/to/image_foo.1-10x2#.jpg   (i.e. 0001)
//     /path/to/image_foo.1-10x2@.jpg   (i.e. 1)
//     /path/to/image_foo.1-10x2@@@.jpg (i.e. 001)
//
type FileSequence struct {
	basename string
	dir      string
	ext      string
	padChar  string
	zfill    int
	frameSet *FrameSet
}

// NewFileSequence returns a FileSequence from a string
// sequence path. The path should be a valid sequence, expressing
// a frame range, or a single file path.
// If error is non-nil, then the given sequence string could not
// be successfully parsed.
//
// Example paths:
//     /path/to/image_foo.1-10x2#.jpg
//     /path/to/single_image.0100.jpg
//
func NewFileSequence(sequence string) (*FileSequence, error) {
	var dir, basename, pad, ext string
	var frameSet *FrameSet
	var err error

	parts := splitPattern.FindStringSubmatch(sequence)

	if len(parts) == 0 {
		for pad := range padding {
			if strings.Contains(sequence, pad) {
				return nil, fmt.Errorf("Failed to parse sequence: %s", sequence)
			}
		}
		// We assume the sequence is just a single file path, containing
		// no frame patterns
		dir, basename = filepath.Split(sequence)
		idx := strings.LastIndex(basename, ".")
		if idx == -1 {
			return nil, fmt.Errorf("Sequence missing extension: %s", sequence)
		}
		basename, ext = basename[:idx], basename[idx:]

		// Try to see if we can at least find a specific frames
		// number, a la  .<frame>.ext
		parts = singleFrame.FindStringSubmatch(sequence)
		if len(parts) > 0 {
			frameSet, err = NewFrameSet(parts[2])
			if err != nil {
				frameSet = nil
			} else {
				// Reparse the dir/basename to not include the trailing frame
				dir, basename = filepath.Split(parts[1])
			}
		}

	} else {
		frameSet, err = NewFrameSet(parts[2])
		if err != nil {
			return nil, fmt.Errorf("Failed to parse frame range from sequence: %s", sequence)
		}

		dir, basename = filepath.Split(parts[1])
		pad = parts[3]
		ext = parts[4]
	}

	seq := &FileSequence{basename, dir, ext, pad, 0, frameSet}
	seq.SetPadding(pad)

	return seq, nil
}

/*
Format returns the file sequence as a formatted string according to
the given template.

Utilizes Go text/template format syntax.  Available functions include:
    dir      - the directory name.
    base     - the basename of the sequence (leading up to the frame range).
    ext      - the file extension of the sequence including leading period.
    startf   - the start frame.
    endf     - the end frame.
    len      - the length of the frame range.
    pad      - the detecting amount of padding.
    frange   - the frame range.
    inverted - the inverted frame range. (returns empty string if none)
    zfill    - the int width of the frame padding

Example:

	{{dir}}{{base}}{{frange}}{{pad}}{{ext}}

*/
func (s *FileSequence) Format(tpl string) (string, error) {
	c := map[string]interface{}{
		"dir":      s.Dirname,
		"base":     s.Basename,
		"ext":      s.Ext,
		"startf":   s.Start,
		"endf":     s.End,
		"len":      s.Len,
		"pad":      s.Padding,
		"zfill":    s.ZFill,
		"frange":   s.FrameRange,
		"inverted": s.InvertedFrameRange,
	}

	t, err := template.New("sequence").Funcs(c).Parse(tpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, c)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (s *FileSequence) Split() []*FileSequence {
	if s.frameSet == nil {
		// Return a copy
		return []*FileSequence{s.Copy()}
	}

	franges := strings.Split(s.frameSet.FrameRange(), ",")
	if len(franges) == 1 {
		// Return a copy
		return []*FileSequence{s.Copy()}
	}

	var buf bytes.Buffer

	// Write the dir and base components once
	buf.WriteString(s.dir)
	buf.WriteString(s.basename)

	// Mark the buffer so we can truncate
	size := buf.Len()

	list := make([]*FileSequence, len(franges))
	var seq *FileSequence
	for i, frange := range franges {
		buf.WriteString(frange)
		buf.WriteString(s.padChar)
		buf.WriteString(s.ext)

		seq, _ = NewFileSequence(buf.String())
		list[i] = seq

		buf.Truncate(size)
	}
	return list
}

// Dirname returns the parsed directory component of the sequence string
func (s *FileSequence) Dirname() string {
	return s.dir
}

// Basename returns the parsed basename component of the sequence string.
// This is the file part, just before the frame range component.
func (s *FileSequence) Basename() string {
	return s.basename
}

// Padding returns the parsed padding characters component of the sequence string.
// May return an empty string if the file sequence was really a single file.
// i.e. # or @ or @@@
func (s *FileSequence) Padding() string {
	return s.padChar
}

// Start returns the starting frame of the sequence.
// If a frame range/number could not be parsed, then this
// will always return 0
func (s *FileSequence) Start() int {
	if s.frameSet == nil {
		return 0
	}
	return s.frameSet.Start()
}

// End returns the ending frame of the sequence.
// If a frame range/number could not be parsed, then this
// will always return 0
func (s *FileSequence) End() int {
	if s.frameSet == nil {
		return 0
	}
	return s.frameSet.End()
}

// ZFill returns the number of "0" fill characters used to
// pad the frame numbers to match the actual file paths
func (s *FileSequence) ZFill() int {
	return s.zfill
}

// FrameRange returns the string frame range component,
// parsed from the sequence. If no frame range was parsed,
// then this method will return an empty string.
func (s *FileSequence) FrameRange() string {
	if s.frameSet == nil {
		return ""
	}
	return s.frameSet.FrameRange()
}

// FrameRangePadded returns the string frame range component,
// parsed from the sequence, and padded out by the pad characters.
// If no frame range was parsed, then this method will return an empty string.
func (s *FileSequence) FrameRangePadded() string {
	if s.frameSet == nil {
		return ""
	}
	return s.frameSet.FrameRangePadded(s.zfill)
}

// InvertedFrameRange returns a new frame range that represents
// all frames *not* within the current frame range. That is, it
// will create a range that "fills in" the current one.
func (s *FileSequence) InvertedFrameRange() string {
	if s.frameSet == nil {
		return ""
	}
	return s.frameSet.InvertedFrameRange(0)
}

// InvertedFrameRangePadded returns a new frame range that represents
// all frames *not* within the current frame range. That is, it
// will create a range that "fills in" the current one.
// Frames are padded out to the zfill width.
func (s *FileSequence) InvertedFrameRangePadded() string {
	if s.frameSet == nil {
		return ""
	}
	return s.frameSet.InvertedFrameRange(s.zfill)
}

// FrameSet returns a FrameSet instance that was created when
// the sequence was parsed. If no frame range was parsed from
// the sequence, then this method will return nil
func (s *FileSequence) FrameSet() *FrameSet {
	return s.frameSet
}

// Ext returns the file extension component from the sequence,
// including the leading period character.
func (s *FileSequence) Ext() string {
	return s.ext
}

/*
Frame returns a path to the given frame in the sequence.  Integer or string
digits are treated as a frame number and padding is applied, all other values
are passed though. Accepts ints, strings, []byte, and Stringer types

Example:

	seq.Frame(1)
	>> /foo/bar.0001.exr

	seq.Frame("#")
	>> /foo/bar.#.exr
*/
func (s *FileSequence) Frame(frame interface{}) (string, error) {
	var zframe string
	var isInt bool

	if s.frameSet != nil {
		switch t := frame.(type) {
		case int:
			var i int = t
			zframe = zfillInt(i, s.zfill)
			isInt = true
		case string:
			zframe = t
		case []byte:
			zframe = string(t)
		case fmt.Stringer:
			zframe = t.String()
		default:
			return zframe, fmt.Errorf("%v is not a valid int/string type", t)
		}
		// Only try and zfill the string if it was an int format
		if !isInt {
			if _, err := strconv.Atoi(zframe); err == nil {
				zframe = zfillString(zframe, s.zfill)
			}
		}
	}

	var buf bytes.Buffer
	buf.WriteString(s.dir)
	buf.WriteString(s.basename)
	buf.WriteString(zframe)
	buf.WriteString(s.ext)
	return buf.String(), nil
}

// Internal fast path for Frame() when we know we have an int frame
func (s *FileSequence) frameInt(frame int) string {
	var zframe string
	if s.frameSet != nil {
		zframe = zfillInt(frame, s.zfill)
	}
	var buf bytes.Buffer
	buf.WriteString(s.dir)
	buf.WriteString(s.basename)
	buf.WriteString(zframe)
	buf.WriteString(s.ext)
	return buf.String()
}

// Index returns the path to the file at the given index
// in the sequence. If a frame range was not parsed from
// the sequence, this will always returns the original path.
// If the index is not valid, this will return an empty string.
func (s *FileSequence) Index(idx int) string {
	if s.frameSet == nil {
		return s.String()
	}
	frame, err := s.frameSet.Frame(idx)
	if err != nil {
		return ""
	}
	path, err := s.Frame(frame)
	if err != nil {
		return ""
	}
	return path
}

// Set a new dirname for the sequence
func (s *FileSequence) SetDirname(dir string) {
	if !strings.HasSuffix(dir, string(filepath.Separator)) {
		dir = dir + string(filepath.Separator)
	}
	s.dir = dir
}

// Set a new basename for the sequence
func (s *FileSequence) SetBasename(base string) {
	s.basename = base
}

// Set a new padding characters for the sequence
func (s *FileSequence) SetPadding(padChars string) {
	s.padChar = padChars

	var zfill int
	for _, c := range padChars {
		zfill += padding[string(c)]
	}
	s.zfill = zfill
}

// Set a new ext for the sequence
func (s *FileSequence) SetExt(ext string) {
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	s.ext = ext
}

// Set a new FrameSet for the sequence
func (s *FileSequence) SetFrameSet(frameSet *FrameSet) {
	s.frameSet = frameSet
}

// Set a new FrameSet, by way of providing a string frame range.
// If the frame range cannot be parsed, an error will be returned.
func (s *FileSequence) SetFrameRange(frameRange string) error {
	frameSet, err := NewFrameSet(frameRange)
	if err != nil {
		return err
	}
	s.frameSet = frameSet
	return nil
}

// Len returns the number of frames in the FrameSet.
// If a frame range was not parsed, this will always return 1
func (s *FileSequence) Len() int {
	if s.frameSet == nil {
		return 1
	}
	return s.frameSet.Len()
}

// String returns the formatted sequence
func (s *FileSequence) String() string {
	var fs string
	if s.frameSet != nil {
		fs = s.frameSet.String()
	}
	var buf bytes.Buffer
	buf.WriteString(s.basename)
	buf.WriteString(fs)
	buf.WriteString(s.padChar)
	buf.WriteString(s.ext)
	out := filepath.Join(s.dir, buf.String())
	return out
}

// Copy returns a copy of the FileSequence
func (s *FileSequence) Copy() *FileSequence {
	seq, _ := NewFileSequence(s.String())
	return seq
}
