//go:generate peg -inline -switch  grammar.peg

package peg

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func ParseSeq(seq string) (*ParsedSeq, error) {
	p := &Parser{Buffer: seq}
	if err := p.Init(); err != nil {
		return nil, err
	}
	if err := p.Parse(); err != nil {
		return nil, err
	}
	p.Execute()
	if p.parsedType != seqType || p.seq == nil {
		return nil, fmt.Errorf("invalid sequence format")
	}
	return p.seq, nil
}

func ParseRange(rangeStr string) (ParsedRanges, error) {
	p := &Parser{Buffer: rangeStr}
	if err := p.Init(); err != nil {
		return nil, err
	}
	if err := p.Parse(); err != nil {
		return nil, err
	}
	p.Execute()
	if p.parsedType != rangeType || p.ranges == nil {
		return nil, fmt.Errorf("invalid range format")
	}
	return p.ranges.parsed, nil
}

type parser struct {
	parsedType parsedType
	seq        *ParsedSeq
	ranges     *parsedRangeState
}

func (p *parser) Seq() *ParsedSeq {
	if p.seq == nil {
		p.seq = &ParsedSeq{}
	}
	return p.seq
}

func (p *parser) Ranges() *parsedRangeState {
	if p.ranges == nil {
		p.ranges = &parsedRangeState{}
	}
	return p.ranges
}

type parsedType int

const (
	seqType parsedType = iota
	rangeType
)

type ParsedSeq struct {
	Dir   string
	Base  string
	Range string
	Pad   string
	Ext   string
}

func (s *ParsedSeq) IsRangeSingle() bool {
	return !strings.ContainsAny(s.Range, ",-")
}

type StepType string

const (
	StepChunk    = StepType("x")
	FillChunk    = StepType("y")
	StaggerChunk = StepType(":")
)

type ParsedRange struct {
	Range    string
	Cmpts    []int
	StepType StepType
}

func (r *ParsedRange) addCmpt(text string) int {
	val, _ := strconv.Atoi(text)
	r.Cmpts = append(r.Cmpts, val)
	return val
}

type ParsedRanges []*ParsedRange

func (r ParsedRanges) String() string {
	data, _ := json.Marshal(r)
	return string(data)
}

type parsedRangeState struct {
	parsed  ParsedRanges
	current ParsedRange
}

func (r *parsedRangeState) add() {
	cpy := r.current
	r.parsed = append(r.parsed, &cpy)
	r.current = ParsedRange{}
}
