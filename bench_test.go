package fileseq

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func benchFrameSetCreate(b *testing.B, n int) {
	r := fmt.Sprintf("1-%dx2", n)
	_, err := NewFrameSet(r)
	if err != nil {
		b.Fatal(err.Error())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewFrameSet(r)
	}
}

func BenchmarkFrameSetCreate100x2(b *testing.B)      { benchFrameSetCreate(b, 100) }
func BenchmarkFrameSetCreate100000x2(b *testing.B)   { benchFrameSetCreate(b, 100000) }
func BenchmarkFrameSetCreate10000000x2(b *testing.B) { benchFrameSetCreate(b, 10000000) }

func benchFrameSetValue(b *testing.B, n int) {
	r := fmt.Sprintf("1-%dx2", n)
	fs, err := NewFrameSet(r)
	if err != nil {
		b.Fatal(err.Error())
	}
	idx := fs.Len() - 1

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fs.Frame(idx)
	}
}

func BenchmarkFrameSetValue100x2(b *testing.B)      { benchFrameSetValue(b, 100) }
func BenchmarkFrameSetValue100000x2(b *testing.B)   { benchFrameSetValue(b, 100000) }
func BenchmarkFrameSetValue10000000x2(b *testing.B) { benchFrameSetValue(b, 10000000) }

func benchFrameSetIndex(b *testing.B, n int) {
	r := fmt.Sprintf("1-%dx2", n)
	fs, err := NewFrameSet(r)
	if err != nil {
		b.Fatal(err.Error())
	}
	val, err := fs.Frame(fs.Len() - 1)
	if err != nil {
		b.Fatal(err.Error())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fs.Index(val)
	}
}

func BenchmarkFrameSetIndex100x2(b *testing.B)      { benchFrameSetIndex(b, 100) }
func BenchmarkFrameSetIndex100000x2(b *testing.B)   { benchFrameSetIndex(b, 100000) }
func BenchmarkFrameSetIndex10000000x2(b *testing.B) { benchFrameSetIndex(b, 10000000) }

func benchFrameSetContains(b *testing.B, n int) {
	r := fmt.Sprintf("1-%dx2", n)
	fs, err := NewFrameSet(r)
	if err != nil {
		b.Fatal(err.Error())
	}
	start := fs.Start()
	end := fs.End()
	mid := end / 2

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fs.HasFrame(start)
		fs.HasFrame(mid)
		fs.HasFrame(end)
	}
}

func BenchmarkFrameSetContains100x2(b *testing.B)      { benchFrameSetContains(b, 100) }
func BenchmarkFrameSetContains100000x2(b *testing.B)   { benchFrameSetContains(b, 100000) }
func BenchmarkFrameSetContains10000000x2(b *testing.B) { benchFrameSetContains(b, 10000000) }

func benchFrameSetInvertedSingleRange(b *testing.B, n int) {
	fs, err := NewFrameSet(fmt.Sprintf("%d-1x2", n))
	if err != nil {
		b.Fatal(err.Error())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fs.InvertedFrameRange(4)
	}
}

func BenchmarkFrameSetInvertedSingleRange100x2(b *testing.B) {
	benchFrameSetInvertedSingleRange(b, 100)
}
func BenchmarkFrameSetInvertedSingleRange10000x2(b *testing.B) {
	benchFrameSetInvertedSingleRange(b, 10000)
}
func BenchmarkFrameSetInvertedSingleRange100000x2(b *testing.B) {
	benchFrameSetInvertedSingleRange(b, 100000)
}

func benchFrameSetInvertedCommaVals(b *testing.B, n int) {
	var vals []string
	for i := 1; i <= n; i++ {
		vals = append(vals, strconv.Itoa(i))
		i++
	}
	r := strings.Join(vals, `,`)
	fs, err := NewFrameSet(r)
	if err != nil {
		b.Fatal(err.Error())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fs.InvertedFrameRange(4)
	}
}

func BenchmarkFrameSetInvertedCommaVals10x2(b *testing.B) {
	benchFrameSetInvertedCommaVals(b, 10)
}
func BenchmarkFrameSetInvertedCommaVals1000x2(b *testing.B) {
	benchFrameSetInvertedCommaVals(b, 1000)
}
func BenchmarkFrameSetInvertedCommaVals10000x2(b *testing.B) {
	benchFrameSetInvertedCommaVals(b, 10000)
}

func benchFrameSetNormalizedSingleRange(b *testing.B, n int) {
	fs, err := NewFrameSet(fmt.Sprintf("%d-1x2", n))
	if err != nil {
		b.Fatal(err.Error())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fs.Normalize()
	}
}

func BenchmarkFrameSetNormalizedSingleRange100x2(b *testing.B) {
	benchFrameSetNormalizedSingleRange(b, 100)
}
func BenchmarkFrameSetNormalizedSingleRange10000x2(b *testing.B) {
	benchFrameSetNormalizedSingleRange(b, 10000)
}
func BenchmarkFrameSetNormalizedSingleRange100000x2(b *testing.B) {
	benchFrameSetNormalizedSingleRange(b, 100000)
}

func benchFrameSetNormalizedCommaVals(b *testing.B, n int) {
	var vals []string
	for i := 1; i <= n; i++ {
		vals = append(vals, strconv.Itoa(i))
		i++
	}
	r := strings.Join(vals, `,`)
	fs, err := NewFrameSet(r)
	if err != nil {
		b.Fatal(err.Error())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fs.Normalize()
	}
}

func BenchmarkFrameSetNormalizedCommaVals10x2(b *testing.B) {
	benchFrameSetNormalizedCommaVals(b, 10)
}
func BenchmarkFrameSetNormalizedCommaVals1000x2(b *testing.B) {
	benchFrameSetNormalizedCommaVals(b, 1000)
}
func BenchmarkFrameSetNormalizedCommaVals10000x2(b *testing.B) {
	benchFrameSetNormalizedCommaVals(b, 10000)
}
