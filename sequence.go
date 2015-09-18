package fileseq

import (
	"bytes"
	"fmt"
	"os"
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
		if idx > -1 {
			basename, ext = basename[:idx], basename[idx:]
		}

		if dir == "" && basename == "" && ext != "" {
			// Just assume all we have is a file extension

		} else {
			// Try to see if we can at least find a specific frame
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
				ext = parts[3]
			}
		}

	} else {
		frameSet, err = NewFrameSet(parts[2])
		if err != nil {
			frameSet = nil
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
    pad      - the detected padding characters (i.e. # , @@@ , ...).
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

func (s *FileSequence) Split() FileSequences {
	if s.frameSet == nil {
		// Return a copy
		return FileSequences{s.Copy()}
	}

	franges := strings.Split(s.frameSet.FrameRange(), ",")
	if len(franges) == 1 {
		// Return a copy
		return FileSequences{s.Copy()}
	}

	var buf bytes.Buffer

	// Write the dir and base components once
	buf.WriteString(s.dir)
	buf.WriteString(s.basename)

	// Mark the buffer so we can truncate
	size := buf.Len()

	list := make(FileSequences, len(franges))
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
	buf := bytes.NewBufferString(s.dir)
	buf.WriteString(s.basename)
	buf.WriteString(fs)
	buf.WriteString(s.padChar)
	buf.WriteString(s.ext)
	return buf.String()
}

// Copy returns a copy of the FileSequence
func (s *FileSequence) Copy() *FileSequence {
	seq, _ := NewFileSequence(s.String())
	return seq
}

// FileSequences is a slice of FileSequence pointers, which can be sorted
type FileSequences []*FileSequence

func (fs FileSequences) Len() int { return len(fs) }

func (fs FileSequences) Less(i, j int) bool {
	return fs[i].String() < fs[j].String()
}

func (fs FileSequences) Swap(i, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}

// FileOptions indicate how file listings should be performed and returned
type FileOption int

const (
	// Include hidden files in file listings
	HiddenFiles FileOption = iota
	// Include paths that aren't detected as sequences, and are
	// just single files
	SingleFiles
)

// FindSequencesOnDisk searches a given directory path and
// sorts all valid sequence-compatible files into a list of
// FileSequences
//
// By default, only non-hidden sequences of files will be returned.
// Extra FileOption values may be given to control whether things like
// hidden files, or single (non-sequence) files should be included in
// the search results.
//
// If there are any errors reading the directory or the files,
// a non-nil error will be returned.
func FindSequencesOnDisk(path string, opts ...FileOption) (FileSequences, error) {
	return findSequencesOnDisk(path, opts...)
}

// ListFiles is an alias for FindSequencesOnDisk, passing the SingleFiles option.
// It. will include all files in the results. Even those that do not contain
// frame range patterns. It is like an ls, but with collapsed sequences.
func ListFiles(path string) (FileSequences, error) {
	return findSequencesOnDisk(path, SingleFiles)
}

func findSequencesOnDisk(path string, opts ...FileOption) (FileSequences, error) {
	root, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer root.Close()

	infos, err := root.Readdir(-1)
	if err != nil {
		return nil, err
	}

	// Get options
	var singleFiles bool
	var hiddenFiles bool
	for _, opt := range opts {
		switch opt {
		case HiddenFiles:
			hiddenFiles = true
		case SingleFiles:
			singleFiles = true
		}
	}

	seqs := make(map[[2]string][]int)
	padMap := make(map[[2]string]string)

	var files FileSequences
	if singleFiles {
		files = make(FileSequences, 0)
	}

	// Prep a string buffer that we can reuse to constantly
	// build strings
	path = filepath.Clean(path)
	buf := bytes.NewBufferString(path)
	if !strings.HasSuffix(path, string(filepath.Separator)) {
		buf.WriteRune(filepath.Separator)
	}

	size := buf.Len()

	// Read dir and sort files into groups
	var match []string

	for _, info := range infos {
		if info.IsDir() {
			continue
		}

		name := info.Name()

		if !hiddenFiles && strings.HasPrefix(name, ".") {
			continue
		}

		match = singleFrame.FindStringSubmatch(name)
		if len(match) == 0 {
			if singleFiles {
				buf.WriteString(name)

				fs, err := NewFileSequence(buf.String())
				if err != nil {
					return nil, err
				}
				fs.SetFrameSet(nil)
				fs.SetPadding("")
				files = append(files, fs)

				buf.Truncate(size)
			}
			continue
		}

		frame, _ := strconv.Atoi(match[2])
		key := [2]string{match[1], match[3]}
		frames, ok := seqs[key]
		if !ok {
			frames = []int{frame}
			padMap[key] = PaddingChars(len(match[2]))
		} else {
			frames = append(frames, frame)
		}
		seqs[key] = frames
	}

	var i int

	fseqs := make(FileSequences, len(seqs))

	// Convert groups into sequences
	for key, frames := range seqs {
		name, ext := key[0], key[1]
		pad := padMap[key]

		var frange string
		if len(frames) == 1 {
			frange = strconv.Itoa(frames[0])
			if name != "" {
				// Make sure a non-sequence file doesn't accidentally
				// get reparsed as a range.
				pos := 1
				// Check if the parsed number was preceded by a "-",
				// if so, check before that char to see if its a number
				if strings.HasSuffix(name, "-") && len(name) >= 2 {
					pos = 2
				}
				dig := string(name[len(name)-pos])
				// If it is a number, clear the padding char
				if _, err := strconv.ParseUint(dig, 10, 8); err == nil {
					pad = ""
				}
			}
		} else {
			frange = FramesToFrameRange(frames, true, 0)
		}

		buf.WriteString(name)
		buf.WriteString(frange)
		buf.WriteString(pad)
		buf.WriteString(ext)

		fs, err := NewFileSequence(buf.String())
		if err != nil {
			return nil, err
		}
		if pad == "" {
			fs.SetFrameSet(nil)
		}
		fseqs[i] = fs

		buf.Truncate(size)
		i++
	}

	if singleFiles {
		fseqs = append(fseqs, files...)
	}

	return fseqs, nil
}

// FindSequenceOnDisk takes a string that is a compatible/parsible
// FileSequence pattern, and finds a sequence on disk which matches
// the Basename and Extension.
// If no match is found, a nil FileSequence is returned.
// If an error occurs while reading the filesystem, a non-nil error
// is returned.
func FindSequenceOnDisk(pattern string) (*FileSequence, error) {
	fs, err := NewFileSequence(pattern)
	if err != nil {
		// Treat a bad pattern as a non-match
		fmt.Println(err.Error())
		return nil, nil
	}

	seqs, err := FindSequencesOnDisk(fs.Dirname())
	if err != nil {
		return nil, err
	}

	base := fs.Basename()
	ext := fs.Ext()

	for _, seq := range seqs {
		// Find the first match and return it
		if seq.Basename() == base && seq.Ext() == ext {
			return seq, nil
		}
	}
	// If we get this far, we didn't find a match
	return nil, nil
}
