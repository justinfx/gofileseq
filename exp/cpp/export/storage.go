package main

import (
	"sync"
	"sync/atomic"

	"github.com/justinfx/gofileseq/v2"
)

var (
	sFrameSets *frameSetMap
	sFileSeqs  *fileSeqMap
)

func init() {
	sFrameSets = &frameSetMap{
		lock: new(sync.RWMutex),
		m:    make(map[FrameSetId]*frameSetRef),
		rand: NewRandSource(),
	}

	sFileSeqs = &fileSeqMap{
		lock: new(sync.RWMutex),
		m:    make(map[FileSeqId]*fileSeqRef),
		rand: NewRandSource(),
	}
}

type frameSetRef struct {
	fileseq.FrameSet
	refs uint32
}

type frameSetMap struct {
	lock *sync.RWMutex
	m    map[FrameSetId]*frameSetRef
	rand idMaker
}

func (m *frameSetMap) Len() int {
	m.lock.RLock()
	l := len(m.m)
	m.lock.RUnlock()
	return l
}

func (m *frameSetMap) Add(fset fileseq.FrameSet) FrameSetId {
	// fmt.Printf("frameset Add %v as %v\n", fset.String(), id)
	m.lock.Lock()
	id := FrameSetId(m.rand.Uint64())
	m.m[id] = &frameSetRef{fset, 1}
	m.lock.Unlock()
	return id
}

func (m *frameSetMap) Incref(id FrameSetId) {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()

	if !ok {
		return
	}

	atomic.AddUint32(&ref.refs, 1)
	// fmt.Printf("Incref %v to %d\n", ref, refs)
}

func (m *frameSetMap) Decref(id FrameSetId) {
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

func (m *frameSetMap) Get(id FrameSetId) (*frameSetRef, bool) {
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
	rand idMaker
}

func (m *fileSeqMap) Len() int {
	m.lock.RLock()
	l := len(m.m)
	m.lock.RUnlock()
	return l
}

func (m *fileSeqMap) Add(seq *fileseq.FileSequence) FileSeqId {
	// fmt.Printf("fileseq Add %v as %v\n", seq.String(), id)
	m.lock.Lock()
	id := FileSeqId(m.rand.Uint64())
	m.m[id] = &fileSeqRef{seq, 1}
	m.lock.Unlock()
	return id
}

func (m *fileSeqMap) Incref(id FileSeqId) {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()

	if !ok {
		return
	}

	atomic.AddUint32(&ref.refs, 1)
	// fmt.Printf("Incref %v to %d\n", ref, refs)
}

func (m *fileSeqMap) Decref(id FileSeqId) {
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

func (m *fileSeqMap) Get(id FileSeqId) (*fileSeqRef, bool) {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()
	return ref, ok
}
