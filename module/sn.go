package module

import (
	"math"
	"sync"
)

type SNGenertor interface {
	Start() uint64
	Max() uint64
	Next() uint64
	// CycleCount 用于获取循环计数。
	CycleCount() uint64
	// Get 用于获得一个序列号并准备下一个序列号。
	Get() uint64
}

type mySNGenertor struct {
	start      uint64
	max        uint64
	next       uint64
	cycleCount uint64
	lock       sync.RWMutex
}

func NewSNGenertor(start uint64, max uint64) SNGenertor {
	if max == 0 {
		max = math.MaxUint64
	}
	return &mySNGenertor{
		start: start,
		max:   max,
		next:  start,
	}
}

func (gen *mySNGenertor) Start() uint64 {
	return gen.start
}

func (gen *mySNGenertor) Max() uint64 {
	return gen.max
}

func (gen *mySNGenertor) Next() uint64 {
	return gen.next
}

func (gen *mySNGenertor) CycleCount() uint64 {
	gen.lock.RLock()
	defer gen.lock.RUnlock()
	return gen.cycleCount
}

func (gen *mySNGenertor) Get() uint64 {
	gen.lock.Lock()
	defer gen.lock.Unlock()
	id := gen.next
	if id == gen.max {
		// 一个循环，当达到最大时，从最小的再来
		gen.next = gen.start
		gen.cycleCount++
	} else {
		gen.next++
	}
	return id
}
