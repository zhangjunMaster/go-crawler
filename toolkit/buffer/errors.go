package buffer

import "errors"

var ErrClosedBuffer = errors.New("closed buffer")
var ErrClosedBufferPool = errors.New("closed buffer pool")
