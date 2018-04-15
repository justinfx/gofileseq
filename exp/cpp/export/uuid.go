package main

import "time"

var randSource *Xor64Source

func init() {
	// rand.Seed(time.Now().UnixNano())
	randSource = NewXor64Source(time.Now().UnixNano())
}

// idMaker interface generates a unique uint64 id
type idMaker interface {
	Uint64() uint64
}

// Create a new idMaker
func NewRandSource() idMaker {
	return NewXor64Source(time.Now().UnixNano())
}

// Generate a new unique id using the shared generator
func uuid() uint64 {

	//// Slowest
	// b := make([]byte, 16)
	// rand.Read(b)
	// b[6] = (b[6] & 0x0f) | 0x40
	// b[8] = (b[8] & 0x3f) | 0x80
	// u := binary.BigEndian.Uint64(b)

	//// Better
	// u := uint64(rand.Uint32())<<32 + uint64(rand.Uint32())

	// Best
	u := randSource.Uint64()

	return u
}

// xorshift implements a 64-bit version of Marsaglia's xorshift PRNG as
// described in http://www.jstatsoft.org/v08/i14/paper.
// The generator has a period of 2^64-1 and passes most of the tests in the
// dieharder test suit.
//
// https://gist.githubusercontent.com/anonymous/3908425/raw/0972271d707997619a45da44103d4dddee17b8c5/xorshift.go

// *Xor64Source implements the rand.Source interface from math/rand.
type Xor64Source uint64

// seed0 is used as default seed to initialize the generator.
const seed0 = 88172645463325252

// NewXor64Source returns a pointer to a new Xor64Source seeded with the given
// value.
func NewXor64Source(seed int64) *Xor64Source {
	var s Xor64Source
	s.Seed(seed)
	return &s
}

// xor64 generates the next value of a pseudo-random sequence given a current
// state x.
func xor64(x uint64) uint64 {
	x ^= x << 13
	x ^= x >> 7
	x ^= x << 17
	return x
}

// next advances the generators internal state to the next value and returns
// this value as an uint64.
func (s *Xor64Source) next() uint64 {
	x := xor64(uint64(*s))
	*s = Xor64Source(x)
	return x
}

// Int63 returns a pseudo-random integer in [0,2^63-1) as an int64.
func (s *Xor64Source) Int63() int64 {
	return int64(s.next() >> 1)
}

// Uint64 returns a pseudo-random integer in [1,2^64-1) as an uint64.
func (s *Xor64Source) Uint64() uint64 {
	return s.next()
}

// Seed uses the given value to initialize the generator. If this value is 0, a
// pre-defined seed is used instead, since the xorshift algorithm requires at
// least one bit of the internal state to be set.
func (s *Xor64Source) Seed(seed int64) {
	if seed == 0 {
		seed = seed0
	}
	*s = Xor64Source(seed)
}
