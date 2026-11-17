package parser

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// corpusFile is shared with the Python fileseq repository and must be kept
// identical there, alongside fileseq.g4.
const corpusFile = "../../grammar/parse_cases.tsv"

type corpusCase struct {
	line  int
	kind  string
	input string
	want  ParseResult
}

func loadCorpus(t testing.TB, port string) []corpusCase {
	t.Helper()
	f, err := os.Open(corpusFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var cases []corpusCase
	sc := bufio.NewScanner(f)
	for n := 1; sc.Scan(); n++ {
		line := sc.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		cols := strings.Split(line, "\t")
		if len(cols) != 3 && len(cols) != 8 {
			t.Fatalf("%s:%d: want 3 or 8 columns, got %d", corpusFile, n, len(cols))
		}
		if cols[0] != "all" && !strings.Contains(","+cols[0]+",", ","+port+",") {
			continue
		}
		c := corpusCase{line: n, kind: cols[1], input: cols[2]}
		if len(cols) == 8 {
			c.want = ParseResult{
				Directory:    cols[3],
				Basename:     cols[4],
				FrameRange:   cols[5],
				Padding:      cols[6],
				Extension:    cols[7],
				IsSequence:   c.kind == "seq",
				IsSingleFile: c.kind == "single",
				IsPlainFile:  c.kind == "plain",
			}
		}
		cases = append(cases, c)
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
	return cases
}

func TestParseCorpus(t *testing.T) {
	for _, c := range loadCorpus(t, "go") {
		got, err := ParseFileSequence(c.input)
		if c.kind == "error" {
			if err == nil {
				t.Errorf("line %d: %q: expected an error, got %+v", c.line, c.input, got)
			}
			continue
		}
		if err != nil {
			t.Errorf("line %d: %q: %v", c.line, c.input, err)
			continue
		}
		if got != c.want {
			t.Errorf("line %d: %q:\n got %+v\nwant %+v", c.line, c.input, got, c.want)
		}
	}
}

// TestParseCorpusWithSLL guards against grammar changes that push valid input
// onto the much slower LL fallback.
func TestParseCorpusWithSLL(t *testing.T) {
	for _, c := range loadCorpus(t, "go") {
		if c.kind == "error" {
			continue
		}
		if errs := parsesWithSLL(c.input); len(errs) > 0 {
			t.Errorf("line %d: %q needed the LL fallback: %v", c.line, c.input, errs[0])
		}
	}
}

func BenchmarkParseCorpus(b *testing.B) {
	cases := loadCorpus(b, "go")
	b.ReportAllocs()
	for b.Loop() {
		for _, c := range cases {
			_, _ = ParseFileSequence(c.input)
		}
	}
}
