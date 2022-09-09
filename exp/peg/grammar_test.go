package peg

import (
	"fmt"
	"testing"
)

func TestSeq(t *testing.T) {
	//seq := "/path/to/.file.comp.v2.1-10x2,50-60%04d.7zip"
	seq := "/path/to/file:1-1x1#.exr"
	//seq := "/path/to/.file"
	parsed, err := ParseSeq(seq)
	if err != nil {
		t.Fatal(err)
	}
	//parsed.PrintSyntaxTree()
	fmt.Printf("%#v\n", parsed)
}

func TestRange(t *testing.T) {
	seq := "1-10x2,50-60,-66,100-111:3"
	parsed, err := ParseRange(seq)
	if err != nil {
		t.Fatal(err)
	}
	//parsed.PrintSyntaxTree()
	fmt.Printf("%s\n", parsed)
}
