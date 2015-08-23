package fileseq

import (
	"fmt"
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

func BenchmarkFrameSetCreate100x2(b *testing.B)        { benchFrameSetCreate(b, 100) }
func BenchmarkFrameSetCreate100000x2(b *testing.B)     { benchFrameSetCreate(b, 100000) }
func BenchmarkFrameSetCreate10000000x2(b *testing.B)   { benchFrameSetCreate(b, 10000000) }
func BenchmarkFrameSetValue100x2(b *testing.B)         { benchFrameSetValue(b, 100) }
func BenchmarkFrameSetValue100000x2(b *testing.B)      { benchFrameSetValue(b, 100000) }
func BenchmarkFrameSetValue10000000x2(b *testing.B)    { benchFrameSetValue(b, 10000000) }
func BenchmarkFrameSetIndex100x2(b *testing.B)         { benchFrameSetIndex(b, 100) }
func BenchmarkFrameSetIndex100000x2(b *testing.B)      { benchFrameSetIndex(b, 100000) }
func BenchmarkFrameSetIndex10000000x2(b *testing.B)    { benchFrameSetIndex(b, 10000000) }
func BenchmarkFrameSetContains100x2(b *testing.B)      { benchFrameSetContains(b, 100) }
func BenchmarkFrameSetContains100000x2(b *testing.B)   { benchFrameSetContains(b, 100000) }
func BenchmarkFrameSetContains10000000x2(b *testing.B) { benchFrameSetContains(b, 10000000) }
