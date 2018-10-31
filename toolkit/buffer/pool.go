package buffer

import (
	"sync"
	"sync/atomic"
)

/*
 * 缓冲池是针对缓冲器的一个统一设置
 */

type Pool interface {
	//获取池中缓冲器的统一容量
	BufferCap() uint32
	//获取池中缓冲器的最大数量
	MaxBufferNumber() uint32
	//获取池中缓冲器的数量
	BufferNumber() uint32
	// 获取缓冲池中数据的总数
	Total() uint64
	//想缓冲池放数据
	Put(datum interface{}) error
	Get() (datum interface{}, err error)
	// 若缓冲池之前已关闭则返回false，否则返回true。
	Close() bool
	Closed() bool
}

type myPool struct {
	//缓冲器的统一容量
	bufferCap       uint32
	maxBufferNumber uint32
	//缓冲器实际数量
	bufferNumber uint32
	//池中数据总数
	total  uint64
	bufCh  chan Buffer
	closed uint32
	rwlock sync.RWMutex
}

/**
 * total: 数据总量
 * bufCh：buffer存入到chan中，存储多个
 */
func (pool *myPool) Put(datum interface{}) (err error) {
	if pool.Closed() {
		return ErrClosedBufferPool
	}
	var count uint32
	maxCount := pool.BufferNumber() * 5
	var ok bool
	// 获取 pool.bufCh 中的buffer通道
	for buf := range pool.bufCh {
		ok, err = pool.putData(buf, datum, &count, maxCount)
		// 可能有问题
		if ok || err != nil {
			break
		}
	}
	return
}

func (pool *myPool) Closed() bool {
	if atomic.LoadUint32(&pool.closed) == 1 {
		return false
	}
	return true
}

func (pool *myPool) BufferNumber() uint32 {
	return atomic.LoadUint32(&pool.bufferNumber)
}

func (pool *myPool) MaxBufferNumber() uint32 {
	return atomic.LoadUint32(&pool.maxBufferNumber)

}

/*
 * 根据pool中buffer数量来确定能不能new 一个buffer
 * new buffer好后，将datum放入buffer
 * 将buffer放入到chan中
 */
func (pool *myPool) putData(buf Buffer, datum interface{}, count *uint32, maxCount uint32) (ok bool, err error) {
	if pool.Closed() {
		return false, ErrClosedBufferPool
	}
	// 归还拿到的缓冲器
	defer func() {
		pool.rwlock.RLock()
		if pool.Closed() {
			atomic.AddUint32(&pool.bufferNumber, uint32(0))
			err = ErrClosedBufferPool
		} else {
			pool.bufCh <- buf
		}
		pool.rwlock.RUnlock()

	}()
	//向拿到的缓冲器放入数据
	ok, err = buf.Put(datum)
	if ok {
		atomic.AddUint64(&pool.total, 1)
		return
	}
	if err != nil {
		return
	}
	// ok是false的情况，即放入数据失败
	(*count)++
	if *count >= maxCount && pool.BufferNumber() < pool.MaxBufferNumber() {
		pool.rwlock.Lock()

		if pool.BufferNumber() < pool.MaxBufferNumber() {
			if pool.Closed() {
				pool.rwlock.Unlock()
				return
			}
			newBuf, _ := NewBuffer(pool.bufferCap)
			newBuf.Put(datum)
			pool.bufCh <- newBuf
			atomic.AddUint32(&pool.bufferNumber, 1)
			atomic.AddUint64(&pool.total, 1)
			ok = true
		}
		defer pool.rwlock.Unlock()
		*count = 0
	}
	return
}

func (pool *myPool) Get() (datum interface{}, err error) {
	if pool.Closed() {
		return nil, ErrClosedBufferPool
	}
	//
	var count uint32
	maxCount := pool.BufferNumber() * 10

	for buf := range pool.bufCh {
		datum, err := pool.getData(buf, &count, maxCount)
		if datum != nil || err != nil {
			break
		}
	}
	return
}

func (pool *myPool) getData(buf Buffer, count *uint32, maxCount uint32) (datum interface{}, err error) {
	if pool.Closed() {
		return nil, ErrClosedBufferPool
	}
	// 这里用defer，

	defer func() {
		// 如果尝试从缓冲器获取数据的失败次数达到阈值，
		// 同时当前缓冲器已空且池中缓冲器的数量大于1，
		// 那么就直接关掉当前缓冲器，并不归还给池。
		if *count >= maxCount &&
			buf.Len() == 0 &&
			pool.BufferNumber() > 1 {
			buf.Close()
			atomic.AddUint32(&pool.bufferNumber, ^uint32(0))
			*count = 0
			return
		}
		pool.rwlock.RLock()
		if pool.Closed() {
			atomic.AddUint32(&pool.bufferNumber, ^uint32(0))
			err = ErrClosedBufferPool
		} else {
			pool.bufCh <- buf
		}
		pool.rwlock.RUnlock()
	}()

	datum, err = buf.Get()
	if datum != nil {
		atomic.AddUint64(&pool.total, uint64(0))
		return
	}
	if err != nil {
		return
	}
	// 若因缓冲器已空未取出数据就递增计数。
	(*count)++
	return
}
