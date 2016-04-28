package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"

	"github.com/justinfx/gofileseq"
)

//export Id
type Id uint64

const emptyId Id = Id(0)

//export FrameSet_New
func FrameSet_New(frange *C.char) (id Id, err *C.char) {
	fset, e := fileseq.NewFrameSet(C.GoString(frange))
	if e != nil {
		// err string is freed by caller
		return id, C.CString(e.Error())
	}

	id = sFrameSets.Add(*fset)
	return id, nil
}

//export FrameSet_Delete
func FrameSet_Delete(id Id) {
	sFrameSets.Remove(id)
}

//export FrameSet_Copy
func FrameSet_Copy(id Id) Id {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return emptyId
	}

	copyFs, _ := fileseq.NewFrameSet(fs.FrameRange())
	copyId := sFrameSets.Add(*copyFs)
	return copyId
}

//export FrameSet_String
func FrameSet_String(id Id) (str *C.char) {
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
func FrameSet_Len(id Id) C.size_t {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return C.size_t(0)
	}
	return C.size_t(fs.Len())
}

//export FrameSet_Index
func FrameSet_Index(id Id, frame int) C.int {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return C.int(-1)
	}
	return C.int(fs.Index(frame))
}

//export FrameSet_Frame
func FrameSet_Frame(id Id, index int) (frame int, err *C.char) {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		err = C.CString("FrameSet was deleted")
		return 0, err
	}
	var e error
	frame, e = fs.Frame(index)
	if e != nil {
		err = C.CString(e.Error())
		frame = 0
	}

	// err string is freed by caller
	return frame, err
}

//export FrameSet_Frames
func FrameSet_Frames(id Id, out *C.int) C.int {
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
func FrameSet_HasFrame(id Id, frame int) bool {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return false
	}
	return fs.HasFrame(frame)
}

//export FrameSet_Start
func FrameSet_Start(id Id) C.int {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return C.int(0)
	}
	return C.int(fs.Start())
}

//export FrameSet_End
func FrameSet_End(id Id) C.int {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return C.int(0)
	}
	return C.int(fs.End())
}

//export FrameSet_FrameRange
func FrameSet_FrameRange(id Id) (ret *C.char) {
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
func FrameSet_FrameRangePadded(id Id, pad int) (ret *C.char) {
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
func FrameSet_Invert(id Id) Id {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return emptyId
	}
	inverted := fs.Invert()
	invertedId := sFrameSets.Add(*inverted)
	return invertedId
}

//export FrameSet_InvertedFrameRange
func FrameSet_InvertedFrameRange(id Id, pad int) (ret *C.char) {
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
func FrameSet_Normalize(id Id) Id {
	fs, ok := sFrameSets.Get(id)
	if !ok {
		return emptyId
	}
	normalized := fs.Normalize()
	normalizedId := sFrameSets.Add(*normalized)
	return normalizedId
}

func main() {

}
