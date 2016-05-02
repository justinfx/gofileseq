package main

/*
#include <inttypes.h>
#include <stdbool.h>
#include <stdlib.h>

typedef struct {
	bool hiddenFiles;
	bool singleFiles;
} FileOption;

*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/justinfx/gofileseq"
)

//export FrameSetId
type FrameSetId uint64

//export FileSeqId
type FileSeqId uint64

//export Error
type Error *C.char

//export FileSequences
type FileSequences *FileSeqId

//export FS_AllocStats
func FS_AllocStats() *C.char {
	s := fmt.Sprintf("FrameSets=%d; FileSeqs=%d",
		sFrameSets.Len(), sFileSeqs.Len())
	// string is freed by caller
	return C.CString(s)
}

/*

FrameSet

*/

//export FrameSet_New
func FrameSet_New(frange *C.char) (FrameSetId, Error) {
	fset, e := fileseq.NewFrameSet(C.GoString(frange))
	if e != nil {
		// err string is freed by caller
		return 0, C.CString(e.Error())
	}

	id := sFrameSets.Add(*fset)
	return id, nil
}

//export FrameSet_Incref
func FrameSet_Incref(id FrameSetId) {
	sFrameSets.Incref(id)
}

//export FrameSet_Decref
func FrameSet_Decref(id FrameSetId) {
	sFrameSets.Decref(id)
}

//export FrameSet_Copy
func FrameSet_Copy(id FrameSetId) FrameSetId {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return 0
	}

	copyFs, _ := fileseq.NewFrameSet(fs.FrameRange())
	copyId := sFrameSets.Add(*copyFs)
	return copyId
}

//export FrameSet_String
func FrameSet_String(id FrameSetId) (str *C.char) {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		str = C.CString("")
	} else {
		str = C.CString(fs.String())
	}
	// CString is freed by caller
	return str
}

//export FrameSet_Len
func FrameSet_Len(id FrameSetId) C.size_t {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return C.size_t(0)
	}
	return C.size_t(fs.Len())
}

//export FrameSet_Index
func FrameSet_Index(id FrameSetId, frame int) C.int {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return C.int(-1)
	}
	return C.int(fs.Index(frame))
}

//export FrameSet_Frame
func FrameSet_Frame(id FrameSetId, index int) (C.int, Error) {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return 0, C.CString("FrameSet was deleted")
	}
	frame, err := fs.Frame(index)
	if err != nil {
		return 0, C.CString(err.Error())
	}

	// err string is freed by caller
	return C.int(frame), nil
}

//export FrameSet_Frames
func FrameSet_Frames(id FrameSetId, out *C.int) C.int {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return C.int(0)
	}

	goframes := fs.Frames()
	size := len(goframes)

	if size == 0 {
		return C.int(0)
	}

	intSize := int(unsafe.Sizeof(C.int(0)))
	startPtr := uintptr(unsafe.Pointer(out))
	for i, frame := range goframes {
		ptr := unsafe.Pointer(startPtr + uintptr(intSize*i))
		*(*C.int)(ptr) = C.int(int32(frame))
	}

	return C.int(size)
}

//export FrameSet_HasFrame
func FrameSet_HasFrame(id FrameSetId, frame int) bool {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return false
	}
	return fs.HasFrame(frame)
}

//export FrameSet_Start
func FrameSet_Start(id FrameSetId) C.int {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return C.int(0)
	}
	return C.int(fs.Start())
}

//export FrameSet_End
func FrameSet_End(id FrameSetId) C.int {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return C.int(0)
	}
	return C.int(fs.End())
}

//export FrameSet_FrameRange
func FrameSet_FrameRange(id FrameSetId) (ret *C.char) {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		ret = C.CString("")
	} else {
		ret = C.CString(fs.FrameRange())
	}
	// caller must free the string
	return ret
}

//export FrameSet_FrameRangePadded
func FrameSet_FrameRangePadded(id FrameSetId, pad int) (ret *C.char) {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		ret = C.CString("")
	} else {
		ret = C.CString(fs.FrameRangePadded(pad))
	}
	// caller must free the string
	return ret
}

//export FrameSet_Invert
func FrameSet_Invert(id FrameSetId) FrameSetId {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return 0
	}
	inverted := fs.Invert()
	invertedId := sFrameSets.Add(*inverted)
	return invertedId
}

//export FrameSet_InvertedFrameRange
func FrameSet_InvertedFrameRange(id FrameSetId, pad int) (ret *C.char) {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		ret = C.CString("")
	} else {
		ret = C.CString(fs.InvertedFrameRange(pad))
	}
	// caller must free the string
	return ret
}

//export FrameSet_Normalize
func FrameSet_Normalize(id FrameSetId) FrameSetId {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return 0
	}
	normalized := fs.Normalize()
	normalizedId := sFrameSets.Add(*normalized)
	return normalizedId
}

/*

FileSequence

*/

//export FileSequence_New
func FileSequence_New(frange *C.char) (FileSeqId, Error) {
	fs, e := fileseq.NewFileSequence(C.GoString(frange))
	if e != nil {
		// err string is freed by caller
		return 0, C.CString(e.Error())
	}

	id := sFileSeqs.Add(fs)
	return id, nil
}

//export FileSequence_Incref
func FileSequence_Incref(id FileSeqId) {
	sFileSeqs.Incref(id)
}

//export FileSequence_Decref
func FileSequence_Decref(id FileSeqId) {
	sFileSeqs.Decref(id)
}

//export FileSequence_Copy
func FileSequence_Copy(id FileSeqId) FileSeqId {
	fs, ok := sFileSeqs.Get(id)
	if !ok {
		return 0
	}

	copyFs, _ := fileseq.NewFileSequence(fs.String())
	copyId := sFileSeqs.Add(copyFs)
	return copyId
}

//export FileSequence_Format
func FileSequence_Format(id FileSeqId, tpl *C.char) (*C.char, Error) {
	fs, ok := sFileSeqs.Get(id)
	if !ok {
		return C.CString(""), nil
	}

	str, err := fs.Format(C.GoString(tpl))
	if err != nil {
		return nil, C.CString(err.Error())
	}

	// caller must free the string
	return C.CString(str), nil
}

//export FileSequence_Dirname
func FileSequence_Dirname(id FileSeqId) *C.char {
	fs, ok := sFileSeqs.Get(id)
	// caller must free string
	if !ok {
		return C.CString("")
	}
	return C.CString(fs.Dirname())
}

//export FileSequence_Basename
func FileSequence_Basename(id FileSeqId) *C.char {
	fs, ok := sFileSeqs.Get(id)
	// caller must free string
	if !ok {
		return C.CString("")
	}
	return C.CString(fs.Basename())
}

//export FileSequence_Start
func FileSequence_Start(id FileSeqId) C.int {
	fs, ok := sFileSeqs.Get(id)
	if !ok {
		return 0
	}
	return C.int(fs.Start())
}

//export FileSequence_End
func FileSequence_End(id FileSeqId) C.int {
	fs, ok := sFileSeqs.Get(id)
	if !ok {
		return 0
	}
	return C.int(fs.End())
}

//export FileSequence_ZFill
func FileSequence_ZFill(id FileSeqId) C.int {
	fs, ok := sFileSeqs.Get(id)
	if !ok {
		return 0
	}
	return C.int(fs.ZFill())
}

//export FileSequence_FrameRange
func FileSequence_FrameRange(id FileSeqId) *C.char {
	fs, ok := sFileSeqs.Get(id)
	// caller must free string
	if !ok {
		return C.CString("")
	}
	return C.CString(fs.FrameRange())
}

//export FileSequence_FrameRangePadded
func FileSequence_FrameRangePadded(id FileSeqId) *C.char {
	fs, ok := sFileSeqs.Get(id)
	// caller must free string
	if !ok {
		return C.CString("")
	}
	return C.CString(fs.FrameRangePadded())
}

//export FileSequence_InvertedFrameRange
func FileSequence_InvertedFrameRange(id FileSeqId) *C.char {
	fs, ok := sFileSeqs.Get(id)
	// caller must free string
	if !ok {
		return C.CString("")
	}
	return C.CString(fs.InvertedFrameRange())
}

//export FileSequence_InvertedFrameRangePadded
func FileSequence_InvertedFrameRangePadded(id FileSeqId) *C.char {
	fs, ok := sFileSeqs.Get(id)
	// caller must free string
	if !ok {
		return C.CString("")
	}
	return C.CString(fs.InvertedFrameRangePadded())
}

//export FileSequence_FrameSet
func FileSequence_FrameSet(id FileSeqId) (fsetId FrameSetId) {
	fsetId = 0
	fs, ok := sFileSeqs.Get(id)
	if !ok {
		return fsetId
	}

	fset := fs.FrameSet()
	if fset == nil {
		return 0
	}

	fsetId = sFrameSets.Add(*fset)
	return fsetId
}

//export FileSequence_Ext
func FileSequence_Ext(id FileSeqId) *C.char {
	fs, ok := sFileSeqs.Get(id)
	// caller must free string
	if !ok {
		return C.CString("")
	}
	return C.CString(fs.Ext())
}

//export FileSequence_Frame_Int
func FileSequence_Frame_Int(id FileSeqId, frame int) *C.char {
	fs, ok := sFileSeqs.Get(id)
	// caller must free string
	if !ok {
		return C.CString("")
	}
	frameStr, _ := fs.Frame(frame)
	return C.CString(frameStr)
}

//export FileSequence_Frame_Fill
func FileSequence_Frame_Fill(id FileSeqId, fill *C.char) *C.char {
	fs, ok := sFileSeqs.Get(id)
	// caller must free string
	if !ok {
		return C.CString("")
	}
	gofill := C.GoString(fill)
	frameStr, _ := fs.Frame(gofill)
	return C.CString(frameStr)
}

//export FileSequence_Index
func FileSequence_Index(id FileSeqId, frame int) *C.char {
	fs, ok := sFileSeqs.Get(id)
	// caller must free string
	if !ok {
		return C.CString("")
	}
	return C.CString(fs.Index(frame))
}

//export FileSequence_Len
func FileSequence_Len(id FileSeqId) C.size_t {
	fs, ok := sFileSeqs.Get(id)
	if !ok {
		return 0
	}
	return C.size_t(fs.Len())
}

//export FileSequence_String
func FileSequence_String(id FileSeqId) *C.char {
	fs, ok := sFileSeqs.Get(id)
	// caller must free string
	if !ok {
		return C.CString("")
	}
	return C.CString(fs.String())
}

//export FileSequence_SetDirname
func FileSequence_SetDirname(id FileSeqId, dir *C.char) {
	if fs, ok := sFileSeqs.Get(id); ok {
		str := C.GoString(dir)
		fs.SetDirname(str)
	}
}

//export FileSequence_SetBasename
func FileSequence_SetBasename(id FileSeqId, base *C.char) {
	if fs, ok := sFileSeqs.Get(id); ok {
		str := C.GoString(base)
		fs.SetBasename(str)
	}
}

//export FileSequence_SetPadding
func FileSequence_SetPadding(id FileSeqId, padChars *C.char) {
	if fs, ok := sFileSeqs.Get(id); ok {
		str := C.GoString(padChars)
		fs.SetPadding(str)
	}
}

//export FileSequence_SetExt
func FileSequence_SetExt(id FileSeqId, ext *C.char) {
	if fs, ok := sFileSeqs.Get(id); ok {
		str := C.GoString(ext)
		fs.SetExt(str)
	}
}

//export FileSequence_SetFrameSet
func FileSequence_SetFrameSet(id FileSeqId, fsetId FrameSetId) bool {
	fs, ok := sFileSeqs.Get(id)
	if !ok {
		return false
	}

	fset, ok := sFrameSets.Get(fsetId)
	if !ok {
		return false
	}

	fs.SetFrameSet(&fset.FrameSet)
	return true
}

//export FileSequence_SetFrameRange
func FileSequence_SetFrameRange(id FileSeqId, frameRange *C.char) Error {
	if fs, ok := sFileSeqs.Get(id); ok {
		str := C.GoString(frameRange)
		if err := fs.SetFrameRange(str); err != nil {
			return C.CString(err.Error())
		}
	}
	return nil
}

// TODO
//export FileSequence_Split
// func FileSequence_Split() FileSequences {

//export FindSequenceOnDisk
func FindSequenceOnDisk(pattern *C.char) (FileSeqId, Error) {
	str := C.GoString(pattern)
	// Caller is responsible for freeing error string
	if str == "" {
		return 0, C.CString("empty pattern")
	}

	fs, err := fileseq.FindSequenceOnDisk(str)
	if err != nil {
		return 0, C.CString(err.Error())
	}

	// No Match
	if fs == nil {
		return 0, nil
	}

	id := sFileSeqs.Add(fs)
	return id, nil
}

//export FindSequencesOnDisk
func FindSequencesOnDisk(path *C.char, opts C.FileOption) (FileSequences, uint64, Error) {
	fileOpts := []fileseq.FileOption{}
	if opts.hiddenFiles {
		fileOpts = append(fileOpts, fileseq.HiddenFiles)
	}
	if opts.singleFiles {
		fileOpts = append(fileOpts, fileseq.SingleFiles)
	}

	gopath := C.GoString(path)
	seqs, err := fileseq.FindSequencesOnDisk(gopath, fileOpts...)
	// fmt.Printf("export.go::DEBUG:results: %v\n", seqs)
	if err != nil {
		return nil, 0, C.CString(err.Error())
	}

	num := uint64(len(seqs))
	size := uint64(unsafe.Sizeof(C.uint64_t(0)))
	list := C.malloc(C.size_t(num * size))
	startPtr := uintptr(unsafe.Pointer(list))

	for i, seq := range seqs {
		id := sFileSeqs.Add(seq)
		// fmt.Printf("export.go::DEBUG:id: %v\n", id)
		ptr := unsafe.Pointer(startPtr + uintptr(size*uint64(i)))
		*(*C.uint64_t)(ptr) = C.uint64_t(id)
	}

	return FileSequences(list), num, nil
}

// Required by CGO
//
func main() {

}
