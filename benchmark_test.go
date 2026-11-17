package fileseq

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/justinfx/gofileseq/v3/internal/parser"
)

var benchSeqs = map[string]string{
	"sequence": "/mnt/projects/show/shots/sh010/comp/v012/sh010_comp_v012.1001-1240#.exr",
	"single":   "/mnt/projects/show/shots/sh010/plates/sh010_plate.1001.dpx",
	"stepped":  "/renders/beauty/beauty.1-100x2,150-200@@@@.tif",
	"plain":    "/mnt/projects/show/assets/chars/hero/model/hero_model_v003.abc",
}

var benchKinds = []string{"sequence", "single", "stepped", "plain"}

func BenchmarkNewFileSequence(b *testing.B) {
	for _, k := range benchKinds {
		s := benchSeqs[k]
		b.Run(k, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if _, err := NewFileSequence(s); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkParseFileSequence(b *testing.B) {
	for _, k := range benchKinds {
		s := benchSeqs[k]
		b.Run(k, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if _, err := parser.ParseFileSequence(s); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkNewFileSequenceParallel is NewFileSequence on all cores, as a
// service resolving many sequence strings at once would call it.
func BenchmarkNewFileSequenceParallel(b *testing.B) {
	for _, k := range benchKinds {
		s := benchSeqs[k]
		b.Run(k, func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if _, err := NewFileSequence(s); err != nil {
						b.Fatal(err)
					}
				}
			})
		})
	}
}

// shotListing is a render directory listing: 20 shots with 3 sequences of
// 240 frames each (beauty, depth, normals), plus 200 plain files.
func shotListing() []string {
	var paths []string
	for shot := 10; shot < 210; shot += 10 {
		for _, pass := range []string{"beauty", "depth", "normals"} {
			for f := 1001; f <= 1240; f++ {
				paths = append(paths, fmt.Sprintf("/mnt/renders/sh%03d/sh%03d_%s_v012.%04d.exr", shot, shot, pass, f))
			}
		}
	}
	for i := 0; i < 200; i++ {
		paths = append(paths, fmt.Sprintf("/mnt/renders/notes/review_note_%03d.txt", i))
	}
	return paths
}

func BenchmarkFindSequencesInList(b *testing.B) {
	paths := shotListing()
	for _, c := range []struct {
		name string
		opts []FileOption
	}{{"sequences", nil}, {"with-single-files", []FileOption{SingleFiles}}} {
		b.Run(c.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if _, err := FindSequencesInList(paths, c.opts...); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkFindSequencesOnDisk lists a directory of 20 sequences of 100
// empty frames and 100 plain files.
func BenchmarkFindSequencesOnDisk(b *testing.B) {
	dir := b.TempDir()
	for s := 0; s < 20; s++ {
		for f := 1; f <= 100; f++ {
			if err := os.WriteFile(filepath.Join(dir, fmt.Sprintf("shot%02d_comp.%04d.exr", s, f)), nil, 0o644); err != nil {
				b.Fatal(err)
			}
		}
	}
	for i := 0; i < 100; i++ {
		if err := os.WriteFile(filepath.Join(dir, fmt.Sprintf("note_%03d.txt", i)), nil, 0o644); err != nil {
			b.Fatal(err)
		}
	}
	for _, c := range []struct {
		name string
		opts []FileOption
	}{{"sequences", nil}, {"with-single-files", []FileOption{SingleFiles}}} {
		b.Run(c.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if _, err := FindSequencesOnDisk(dir, c.opts...); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
