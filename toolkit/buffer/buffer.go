package buffer

import (
	"fmt"
	"sync"
	"sync/atomic"

	"gopcp.v2/chapter6/webcrawler/errors"
)

/**
* chan造成运行时恐慌：1.向close的put 2.重复close chan
* 使用原子操作来改变变量的值
* 使用原子操作改变值时，就需要使用原子操作来获取值
*
 */
/**
 * 缓冲器是扩展通道用的
 */
type Buffer interface {
	Cap() uint32
	Len() uint32
	Put(datum interface{}) (bool, error)
	Get() (interface{}, error)
	Close() bool
	Closed() bool
}

// 对chan的封装
type myBuffer struct {
	ch          chan interface{}
	closed      uint32
	closinglock sync.RWMutex
}

func NewBuffer(size uint32) (Buffer, error) {
	if size == 0 {
		errMsg := fmt.Sprintf("illegal size for Buffer:%d", size)
		return nil, errors.NewIllegalParameterError(errMsg)
	}
	return &myBuffer{
		ch: make(chan interface{}, size),
	}, nil
}

func (buf *myBuffer) Cap() uint32 {
	return uint32(cap(buf.ch))
}

func (buf *myBuffer) Len() uint32 {
	return uint32(len(buf.ch))
}

/*
 * 返回值如果有变量名，则就会有默认值
 * 直接return就可以了
 */
func (buf *myBuffer) Put(datum interface{}) (ok bool, err error) {
	buf.closinglock.RLock()
	defer buf.closinglock.RUnlock()
	if buf.Closed() {
		return false, ErrClosedBuffer
	}
	select {
	case buf.ch <- datum:
		ok = true
	default:
		ok = false
	}
	return
}

// <- chan时，判断是否是false，如果是false则是关闭了
// 使用select防止阻塞
func (buf *myBuffer) Get() (interface{}, error) {
	select {
	case datum, ok := <-buf.ch:
		if !ok {
			return nil, ErrClosedBuffer
		}
		return datum, nil
	default:
		return nil, nil
	}
}

func (buf *myBuffer) Close() bool {
	if atomic.CompareAndSwapUint32(&buf.closed, 0, 1) {
		buf.closinglock.Lock()
		close(buf.ch)
		buf.closinglock.Unlock()
		return true
	}
	return false
}

func (buf *myBuffer) Closed() bool {
	if atomic.LoadUint32(&buf.closed) == 0 {
		return false
	}
	return true
}
