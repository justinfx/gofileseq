package main

import (
	"crypto/rand"
	"encoding/binary"
	"sync"

	"github.com/justinfx/gofileseq"
)

var (
	sFrameSets frameSetMap
)

func init() {
	sFrameSets = frameSetMap{
		lock: new(sync.RWMutex),
		m:    make(map[Id]fileseq.FrameSet),
	}
}

type frameSetMap struct {
	lock *sync.RWMutex
	m    map[Id]fileseq.FrameSet
}

func (m frameSetMap) Add(fset fileseq.FrameSet) Id {
	id := Id(uuid())
	m.lock.Lock()
	m.m[id] = fset
	m.lock.Unlock()
	return id
}

func (m frameSetMap) Remove(id Id) {
	m.lock.Lock()
	delete(m.m, id)
	m.lock.Unlock()
}

func (m frameSetMap) Get(id Id) (fileseq.FrameSet, bool) {
	m.lock.RLock()
	fset, ok := m.m[id]
	m.lock.RUnlock()
	return fset, ok
}

func uuid() uint64 {
	b := make([]byte, 16)
	rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return binary.BigEndian.Uint64(b)
}
