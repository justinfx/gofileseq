package main

import (
	"crypto/rand"
	"encoding/binary"
	"sync"
	"sync/atomic"

	"github.com/justinfx/gofileseq"
)

var (
	sFrameSets frameSetMap
	sFileSeqs  fileSeqMap
)

func init() {
	sFrameSets = frameSetMap{
		lock: new(sync.RWMutex),
		m:    make(map[FrameSetId]*frameSetRef),
	}

	sFileSeqs = fileSeqMap{
		lock: new(sync.RWMutex),
		m:    make(map[FileSeqId]*fileSeqRef),
	}
}

func uuid() uint64 {
	b := make([]byte, 16)
	rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return binary.BigEndian.Uint64(b)
}

type frameSetRef struct {
	fileseq.FrameSet
	refs uint32
}

type frameSetMap struct {
	lock *sync.RWMutex
	m    map[FrameSetId]*frameSetRef
}

func (m frameSetMap) Len() int {
	m.lock.RLock()
	l := len(m.m)
	m.lock.RUnlock()
	return l
}

func (m frameSetMap) Add(fset fileseq.FrameSet) FrameSetId {
	id := FrameSetId(uuid())
	// fmt.Printf("frameset Add %v as %v\n", fset.String(), id)
	m.lock.Lock()
	m.m[id] = &frameSetRef{fset, 1}
	m.lock.Unlock()
	return id
}

func (m frameSetMap) Incref(id FrameSetId) {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()

	if !ok {
		return
	}

	atomic.AddUint32(&ref.refs, 1)
	// fmt.Printf("Incref %v to %d\n", ref, refs)
}

func (m frameSetMap) Decref(id FrameSetId) {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()

	if !ok {
		return
	}

	refs := atomic.AddUint32(&ref.refs, ^uint32(0))
	// fmt.Printf("Decref %v to %d\n", ref, refs)
	if refs != 0 {
		return
	}

	m.lock.Lock()
	if atomic.LoadUint32(&ref.refs) == 0 {
		// fmt.Printf("Deleting %v\n", ref)
		delete(m.m, id)
	}
	m.lock.Unlock()
}

func (m frameSetMap) Get(id FrameSetId) (*frameSetRef, bool) {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()
	return ref, ok
}

type fileSeqRef struct {
	*fileseq.FileSequence
	refs uint32
}

type fileSeqMap struct {
	lock *sync.RWMutex
	m    map[FileSeqId]*fileSeqRef
}

func (m fileSeqMap) Len() int {
	m.lock.RLock()
	l := len(m.m)
	m.lock.RUnlock()
	return l
}

func (m fileSeqMap) Add(seq *fileseq.FileSequence) FileSeqId {
	id := FileSeqId(uuid())
	// fmt.Printf("fileseq Add %v as %v\n", seq.String(), id)
	m.lock.Lock()
	m.m[id] = &fileSeqRef{seq, 1}
	m.lock.Unlock()
	return id
}

func (m fileSeqMap) Incref(id FileSeqId) {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()

	if !ok {
		return
	}

	atomic.AddUint32(&ref.refs, 1)
	// fmt.Printf("Incref %v to %d\n", ref, refs)
}

func (m fileSeqMap) Decref(id FileSeqId) {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()

	if !ok {
		return
	}

	refs := atomic.AddUint32(&ref.refs, ^uint32(0))
	// fmt.Printf("Decref %v to %d\n", ref, refs)
	if refs != 0 {
		return
	}

	m.lock.Lock()
	if atomic.LoadUint32(&ref.refs) == 0 {
		// fmt.Printf("Deleting %v\n", ref)
		delete(m.m, id)
	}
	m.lock.Unlock()
}

func (m fileSeqMap) Get(id FileSeqId) (*fileSeqRef, bool) {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()
	return ref, ok
}
