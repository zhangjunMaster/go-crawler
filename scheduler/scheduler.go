package scheduler

import (
	"net/http"
)

type Scheduler interface {
	Init(requestArgs RequestArgs, dataArgs DataArgs, moduleArags ModuleArgs) (err error)
	Start(firdtHttpRequest *http.Request) (err error)
	Stop() (err error)
	Status() Status
	// 获取错误通道，如果为 nil，则说明通道不可用或调度器已停止
	ErrorChan() <-chan error
	Summary() SchedSummary
}
