package fileseq

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
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
//
//	@ - 1 pad width (@@@@ is equal to 4 padding)
//	# - 4 pad width (## is equal to 8 padding)
//
// Example paths and padding:
//
//	/path/to/single_image.0100.jpg
//	/path/to/image_foo.1-10x2#.jpg   (i.e. 0001)
//	/path/to/image_foo.1-10x2@.jpg   (i.e. 1)
//	/path/to/image_foo.1-10x2@@@.jpg (i.e. 001)
type FileSequence struct {
	basename  string
	dir       string
	ext       string
	padChar   string
	zfill     int
	frameSet  *FrameSet
	padMapper paddingMapper
}

// NewFileSequence returns a FileSequence from a string
// sequence path. The path should be a valid sequence, expressing
// a frame range, or a single file path.
// If error is non-nil, then the given sequence string could not
// be successfully parsed.
//
// # PadStyleDefault is used as the padding character formatter
//
// Example paths:
//
//	/path/to/image_foo.1-10x2#.jpg
//	/path/to/single_image.0100.jpg
func NewFileSequence(sequence string) (*FileSequence, error) {
	return NewFileSequencePad(sequence, PadStyleDefault)
}

// NewFileSequencePad returns a FileSequence from a string
// sequence path. The path should be a valid sequence, expressing
// a frame range, or a single file path.
// If error is non-nil, then the given sequence string could not
// be successfully parsed.
//
// The sequence uses the style of padding given, in
// order to convert between padding characters and their numeric width.
//
// Example path w/ PadStyleHash1:
//
//	/path/to/image_foo.1-10x2#.jpg => /path/to/image_foo.1.jpg ...
//	/path/to/image_foo.1-10x2@.jpg => /path/to/image_foo.1.jpg ...
//	/path/to/image_foo.1-10x2##.jpg => /path/to/image_foo.01.jpg ...
//
// Example path w/ PadStyleHash4:
//
//	/path/to/image_foo.1-10x2#.jpg => /path/to/image_foo.0001.jpg ...
//	/path/to/image_foo.1-10x2@.jpg => /path/to/image_foo.1.jpg ...
//	/path/to/image_foo.1-10x2##.jpg => /path/to/image_foo.00000001.jpg ...
func NewFileSequencePad(sequence string, style PadStyle) (*FileSequence, error) {
	var dir, basename, pad, ext string
	var frameSet *FrameSet
	var err error

	// Determine which style of padding to use
	padder, ok := padders[style]
	if !ok {
		padder = defaultPadding
	}

	parts := splitPattern.FindStringSubmatch(sequence)

	if len(parts) == 0 {
		// If no match at this point, we are dealing with a possible
		// path that contains no range. Maybe a single frame or no frame.

		for _, pad := range defaultPadding.AllChars() {
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
			parts = singleFramePattern.FindStringSubmatch(sequence)
			if len(parts) > 0 {
				dir, basename = filepath.Split(parts[1])
				frameStr := parts[2]
				ext = parts[3]
				// Being selective about when we decide we have a single file
				// with a valid frame number
				if frameStr != "" && !strings.HasSuffix(parts[1], ".") {
					// edge case: we've got a single versioned file, not a sequence
					basename += frameStr
				} else {
					// we have a sequence with a single frame
					frameSet, err = NewFrameSet(frameStr)
					if err != nil {
						frameSet = nil
					} else {
						// Calculate the padding chars
						pad = padder.PaddingChars(len(strings.TrimSpace(frameStr)))
					}
				}
			}
		}

	} else {
		// We are dealing with a pattern containing a frame range

		frameSet, err = NewFrameSet(parts[2])
		if err != nil {
			frameSet = nil
		}

		dir, basename = filepath.Split(parts[1])
		pad = parts[3]
		ext = parts[4]
	}

	seq := &FileSequence{
		basename:  basename,
		dir:       dir,
		ext:       ext,
		padChar:   pad,
		zfill:     0, // Fill will be calculated after SetPadding()
		frameSet:  frameSet,
		padMapper: padder,
	}

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

// Split a non-contiguous (i.e. comma-separated) sequence into
// a list of sequences that each use consistent stepping
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

	var buf strings.Builder

	// Write the dir and base components once
	buf.WriteString(s.dir)
	buf.WriteString(s.basename)

	list := make(FileSequences, len(franges))
	var seq *FileSequence
	for i, frange := range franges {
		buf.WriteString(frange)
		buf.WriteString(s.padChar)
		buf.WriteString(s.ext)

		seq, _ = NewFileSequence(buf.String())
		list[i] = seq

		buf.Reset()
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

// PaddingStyle returns the style of padding being used
// to convert between characters and their numeric width,
// i.e. # == 4
func (s *FileSequence) PaddingStyle() PadStyle {
	for style, mapper := range padders {
		if mapper == s.padMapper {
			return style
		}
	}
	return PadStyleDefault
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

	var buf strings.Builder
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
	var buf strings.Builder
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
	sep := string(filepath.Separator)
	if strings.Contains(dir, "\\") {
		dir = filepath.FromSlash(dir)
		sep = "\\"
	}
	if !strings.HasSuffix(dir, sep) {
		dir += sep
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
	s.zfill = s.padMapper.PaddingCharsSize(padChars)
}

// Set a new padding style for mapping between characters and
// their numeric width
func (s *FileSequence) SetPaddingStyle(style PadStyle) {
	s.padMapper = padders[style]
	s.SetPadding(s.padMapper.PaddingChars(s.ZFill()))
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

// FileOption indicate how file listings should be performed and returned
type FileOption int

const (
	// HiddenFiles includes hidden files in file listings
	HiddenFiles FileOption = iota
	// SingleFiles includes paths that aren't detected as sequences, and are
	// just single files
	SingleFiles
	// FileOptPadStyleHash1 sets the PadStyle to PadStyleHash1
	FileOptPadStyleHash1
	// FileOptPadStyleHash4 sets the PadStyle to PadStyleHash4
	FileOptPadStyleHash4
	// StrictPadding instructs FindSequenceOnDisk to consider the padding
	// length of the search pattern, and to ignore files with a different padding
	StrictPadding
)

type findSeqOptions struct {
	FileOptions []FileOption
	SeqTemplate *FileSequence
}

func (o *findSeqOptions) GetFileOptions() []FileOption {
	if o == nil {
		return nil
	}
	return o.FileOptions
}

func (o *findSeqOptions) GetSeqTemplate() *FileSequence {
	if o == nil {
		return nil
	}
	return o.SeqTemplate
}

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
	return findSequencesOnDisk(path, &findSeqOptions{FileOptions: opts})
}

// ListFiles is an alias for FindSequencesOnDisk, passing the SingleFiles option.
// It. will include all files in the results. Even those that do not contain
// frame range patterns. It is like an ls, but with collapsed sequences.
func ListFiles(path string) (FileSequences, error) {
	return findSequencesOnDisk(path, &findSeqOptions{FileOptions: []FileOption{SingleFiles}})
}

type fileItem struct {
	DirName  string
	FileName string
}

func findSequencesOnDisk(path string, opts *findSeqOptions) (FileSequences, error) {
	root, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	infos, err := root.Readdir(-1)
	root.Close()
	if err != nil {
		return nil, err
	}

	// Prep a string buffer that we can reuse to constantly
	// build strings
	path = filepath.Clean(path)
	sep := string(filepath.Separator)
	if strings.Contains(path, "\\") {
		path = filepath.FromSlash(path)
		sep = "\\"
	}
	if !strings.HasSuffix(path, sep) {
		path += sep
	}
	buf := bytes.NewBufferString(path)
	size := buf.Len()

	// Read dir and sort files into groups based on their name
	// and file extension
	var fileItems []*fileItem
	for _, info := range infos {
		if info.IsDir() {
			continue
		}

		// Also skip symlinks that point to directories
		if (info.Mode() & os.ModeSymlink) != 0 {
			buf.WriteString(info.Name())
			syminfo, err := os.Stat(buf.String())
			buf.Truncate(size)
			if err != nil {
				return nil, fmt.Errorf("Error reading symlink %q: %s", buf.String(), err)
			}
			if syminfo.IsDir() {
				continue
			}
		}

		fileItems = append(fileItems, &fileItem{path, info.Name()})
	}

	return findSequencesInList(fileItems, opts)
}

// FindSequencesInList takes a list of individual file paths and compresses them into
// sequences based on their dirname, basename, and extension.
//
// By default, only non-hidden sequences of files will be returned.
// Extra FileOption values may be given to control whether things like
// hidden files, or single (non-sequence) files should be included in
// the results.
func FindSequencesInList(paths []string, opts ...FileOption) (FileSequences, error) {
	items := make([]*fileItem, len(paths))
	for i, path := range paths {
		path = filepath.Clean(path)

		// normalize path separator on Windows
		sep := string(filepath.Separator)
		if strings.Contains(path, "\\") {
			path = filepath.FromSlash(path)
			sep = "\\"
		}

		var item fileItem
		item.DirName, item.FileName = filepath.Split(path)

		if !strings.HasSuffix(item.DirName, sep) {
			item.DirName += sep
		}

		items[i] = &item
	}

	return findSequencesInList(items, &findSeqOptions{FileOptions: opts})
}

func findSequencesInList(paths []*fileItem, opts *findSeqOptions) (FileSequences, error) {
	// Get options
	var (
		singleFiles bool
		hiddenFiles bool
		padStyle    = PadStyleDefault
	)

	for _, opt := range opts.GetFileOptions() {
		switch opt {
		case SingleFiles:
			singleFiles = true
		case HiddenFiles:
			hiddenFiles = true
		case FileOptPadStyleHash1:
			padStyle = PadStyleHash1
		case FileOptPadStyleHash4:
			padStyle = PadStyleHash4
		}
	}

	type FrameInfo struct {
		Frame    string
		FrameNum int
		MinWidth int
	}

	type SeqInfo struct {
		Frames   []FrameInfo
		Padding  string
		MinWidth int
	}

	type SeqKey [3]string // dir, base, ext

	frameMinSize := func(frame string) int {
		frameSize := len(frame)
		frameNum, _ := strconv.Atoi(frame)
		frameNumSize := len(strconv.Itoa(frameNum))
		if frameSize == frameNumSize {
			return 1
		}
		return frameSize
	}

	// map key for sequences: basename, padWidth, ext
	seqs := make(map[SeqKey]SeqInfo)

	padder := padders[padStyle]

	var files FileSequences
	if singleFiles {
		files = make(FileSequences, 0)
	}

	var (
		buf      strings.Builder
		seqCount int
		dirName  string
		baseName string
		frameStr string
		ext      string
	)

	using := opts.GetSeqTemplate()
	if using != nil {
		dirName = using.Dirname()
		baseName = using.Basename()
		ext = using.Ext()
	}

	for _, item := range paths {
		if !hiddenFiles && strings.HasPrefix(item.FileName, ".") {
			continue
		}

		ok := true

		if using != nil {
			// when a FileSequence template has been provided, we can
			// directly set expected values from that, instead of having
			// to make assumptions from the path string

			// effectively a "glob" against the using pattern: <basename>*<ext>
			if !(strings.HasPrefix(item.FileName, baseName) && strings.HasSuffix(item.FileName, ext)) {
				continue
			}
			frameStr = item.FileName[len(baseName) : len(item.FileName)-len(ext)]
			_ = frameStr

		} else {
			// otherwise, we need to do some tests on the path and figure
			// out the values

			dirName = item.DirName
			baseName, frameStr, ext = "", "", ""

			match := optionalFramePattern.FindStringSubmatch(item.FileName)
			if len(match) == 0 {
				ok = false
			} else {
				baseName, frameStr, ext = match[1], match[2], match[3]
				// having no frame, or having only a frame, is not considered a sequence
				if frameStr == "" || (baseName == "" && ext == "") {
					ok = false
				}
			}

			if !ok {
				if singleFiles {
					buf.WriteString(dirName)
					buf.WriteString(item.FileName)

					fs, err := NewFileSequencePad(buf.String(), padStyle)
					if err != nil {
						return nil, err
					}
					// Preserve the parsed base/frame/ext
					fs.basename = baseName
					fs.ext = ext
					if frameStr == "" {
						fs.SetFrameSet(nil)
						fs.SetPadding("")
					} else {
						fs.SetFrameRange(frameStr)
					}
					files = append(files, fs)

					buf.Reset()
				}
				continue
			}
		}

		frameNum, _ := strconv.Atoi(frameStr)
		frameWidth := len(frameStr)

		key := SeqKey{dirName, baseName, ext}
		seq, ok := seqs[key]

		seq.Frames = append(seq.Frames, FrameInfo{
			Frame:    frameStr,
			FrameNum: frameNum,
			MinWidth: frameMinSize(frameStr),
		})

		if !ok {
			seqCount++
		}

		if (!ok || frameWidth < seq.MinWidth) && frameWidth > 0 {
			seq.MinWidth = frameWidth
			seq.Padding = padder.PaddingChars(frameWidth)
		}

		seqs[key] = seq
	}

	var frange, pad string
	dirName = ""
	baseName = ""
	ext = ""

	fseqs := make(FileSequences, 0, seqCount)

	appendSeq := func() error {
		buf.WriteString(dirName)
		buf.WriteString(baseName)
		buf.WriteString(frange)
		buf.WriteString(pad)
		buf.WriteString(ext)

		fs, err := NewFileSequencePad(buf.String(), padStyle)
		if err != nil {
			return err
		}

		// Always use the previously parsed basename, range, and ext
		fs.basename = baseName
		fs.ext = ext
		if frange == "" {
			fs.SetFrameSet(nil)
			fs.SetPadding("")
		} else {
			fs.SetFrameRange(frange)
		}

		fseqs = append(fseqs, fs)

		buf.Reset()
		return nil
	}

	// Convert groups into sequences based on padding
	for key, seq := range seqs {

		if len(seq.Frames) == 0 {
			continue
		}

		dirName, baseName, ext = key[0], key[1], key[2]
		pad = seq.Padding

		// Handle single frame sequences
		if len(seq.Frames) == 1 {
			if baseName != "" {
				// Make sure a non-sequence file doesn't accidentally
				// get re-parsed as a range.
				pos := 1
				// Check if the parsed number was preceded by a "-",
				// if so, check before that char to see if its a number
				if strings.HasSuffix(baseName, "-") && len(baseName) >= 2 {
					pos = 2
				}
				dig := string(baseName[len(baseName)-pos])
				// If it is a number, clear the padding char
				if _, err := strconv.ParseUint(dig, 10, 8); err == nil {
					//pad = ""
				}
			}
			// If there is no padding, use the full string format of the frame.
			// Otherwise, use the numeric value and the padding char.
			frange = seq.Frames[0].Frame
			if pad != "" {
				frange = strconv.Itoa(seq.Frames[0].FrameNum)
			}

			if appendErr := appendSeq(); appendErr != nil {
				return nil, appendErr
			}
			continue
		}

		// Multi-frame sequences
		//

		// Sort list by their required padding
		sort.Slice(seq.Frames, func(i, j int) bool {
			return len(seq.Frames[i].Frame) < len(seq.Frames[j].Frame)
		})

		var frames []int
		currentWidth := len(seq.Frames[0].Frame)
		pad = padder.PaddingChars(currentWidth)

		// Walk the sorted frames, trying to group sequences with
		// compatible padding. Start a new sequence each time a new
		// padding has to be used.
		for _, frameInfo := range seq.Frames {

			if len(frameInfo.Frame) != currentWidth && frameInfo.MinWidth > currentWidth {
				// Commit current frame range and start over
				frange = FramesToFrameRange(frames, true, 0)
				if appendErr := appendSeq(); appendErr != nil {
					return nil, appendErr
				}

				frames = nil
				currentWidth = len(frameInfo.Frame)
				pad = padder.PaddingChars(currentWidth)
			}

			frames = append(frames, frameInfo.FrameNum)
		}

		// Append a remaining sequence
		if len(frames) > 0 {
			frange = FramesToFrameRange(frames, true, 0)
			if appendErr := appendSeq(); appendErr != nil {
				return nil, appendErr
			}
			frames = nil
		}
	}

	if singleFiles {
		fseqs = append(fseqs, files...)
	}

	return fseqs, nil
}

// FindSequenceOnDisk takes a string that is a compatible/parsible
// FileSequence pattern, and finds a sequence on disk which matches
// the Basename and Extension.
//
// By default, the patterns frame value and padding characters are ignored and
// replaced by a wildcard '*' glob. If you want to match on a specific frame
// padding length, pass the option StrictPadding
//
// If no match is found, a nil FileSequence is returned.
// If an error occurs while reading the filesystem, a non-nil error
// is returned.
//
// Example:
//
//	// Find matches with any frame value
//	FindSequenceOnDisk("/path/to/seq.#.ext")
//
//	// Find matches specifically having 4-padded frames
//	FindSequenceOnDisk("/path/to/seq.#.ext", StrictPadding)
func FindSequenceOnDisk(pattern string, opts ...FileOption) (*FileSequence, error) {
	return FindSequenceOnDiskPad(pattern, PadStyleDefault, opts...)
}

// FindSequenceOnDiskPad takes a string that is a compatible/parsible
// FileSequence pattern, and finds a sequence on disk which matches
// the Basename and Extension. The returned seq will use the given PadStyle.
//
// By default the patterns frame value and padding characters are ignored and
// replaced by a wildcard '*' glob. If you want to match on a specific frame
// padding length, pass the option StrictPadding
//
// If no match is found, a nil FileSequence is returned.
// If an error occurs while reading the filesystem, a non-nil error
// is returned.
func FindSequenceOnDiskPad(pattern string, padStyle PadStyle, opts ...FileOption) (*FileSequence, error) {
	optsCopy := make([]FileOption, len(opts))
	copy(optsCopy, opts)

	var strictPadding bool

	for _, opt := range opts {
		switch opt {

		case FileOptPadStyleHash1:
			padStyle = PadStyleHash1
			optsCopy = append(optsCopy, FileOptPadStyleHash1)

		case FileOptPadStyleHash4:
			padStyle = PadStyleHash4
			optsCopy = append(optsCopy, FileOptPadStyleHash4)

		case StrictPadding:
			strictPadding = true
		}
	}

	fs, err := NewFileSequencePad(pattern, padStyle)
	if err != nil {
		// Treat a bad pattern as a non-match
		//fmt.Println(err.Error())
		return nil, nil
	}

	seqs, err := findSequencesOnDisk(fs.Dirname(), &findSeqOptions{FileOptions: optsCopy, SeqTemplate: fs})
	if err != nil {
		return nil, fmt.Errorf("failed to find %q: %s", pattern, err.Error())
	}

	base := fs.Basename()
	ext := fs.Ext()
	pad := fs.Padding()
	fill := fs.ZFill()

	for _, seq := range seqs {
		// Find the first match and return it
		if seq.Basename() != base || seq.Ext() != ext {
			continue
		}

		if seq.FrameSet() != nil {
			seq.SetPaddingStyle(padStyle)
		}

		// Strict padding check
		if strictPadding && pad != "" && seq.ZFill() != fill {
			continue
		}

		return seq, nil
	}
	// If we get this far, we didn't find a match
	return nil, nil
}
