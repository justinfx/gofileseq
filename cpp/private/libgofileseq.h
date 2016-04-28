/* Created by "go tool cgo" - DO NOT EDIT. */

/* package github.com/justinfx/gofileseq/cpp/private */

/* Start of preamble from import "C" comments.  */


#line 3 "/Users/justin/src/go/src/github.com/justinfx/gofileseq/cpp/private/export.go"

#include <stdlib.h>



/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

typedef struct { const char *p; GoInt n; } GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


/* Return type for FrameSet_New */
struct FrameSet_New_return {
	GoUint64 r0; /* id */
	char* r1; /* err */
};

extern struct FrameSet_New_return FrameSet_New(char* p0);

extern void FrameSet_Delete(GoUint64 p0);

extern GoUint64 FrameSet_Copy(GoUint64 p0);

extern char* FrameSet_String(GoUint64 p0);

extern size_t FrameSet_Len(GoUint64 p0);

extern int FrameSet_Index(GoUint64 p0, GoInt p1);

/* Return type for FrameSet_Frame */
struct FrameSet_Frame_return {
	GoInt r0; /* frame */
	char* r1; /* err */
};

extern struct FrameSet_Frame_return FrameSet_Frame(GoUint64 p0, GoInt p1);

extern int FrameSet_Frames(GoUint64 p0, int* p1);

extern GoUint8 FrameSet_HasFrame(GoUint64 p0, GoInt p1);

extern int FrameSet_Start(GoUint64 p0);

extern int FrameSet_End(GoUint64 p0);

extern char* FrameSet_FrameRange(GoUint64 p0);

extern char* FrameSet_FrameRangePadded(GoUint64 p0, GoInt p1);

extern GoUint64 FrameSet_Invert(GoUint64 p0);

extern char* FrameSet_InvertedFrameRange(GoUint64 p0, GoInt p1);

extern GoUint64 FrameSet_Normalize(GoUint64 p0);

#ifdef __cplusplus
}
#endif
