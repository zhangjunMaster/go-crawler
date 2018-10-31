package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

func main() {
	var tt uint32
	atomic.AddUint32(&tt, 10)
	atomic.AddUint32(&tt, 1)
	t := atomic.LoadUint32(&tt)
	fmt.Println(t)
	tty, err := strconv.ParseUint("1234", 10, 12)
	fmt.Println(tty, err)
}
