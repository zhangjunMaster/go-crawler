package stub

import "go-crawler/module"

type ModuleInternal interface {
	module.Module
	// IncrCalledCount 会把调用计数增1。
	IncrCalledCount()
	// IncrAcceptedCount 会把接受计数增1。
	IncrAcceptedCount()
	// IncrCompletedCount 会把成功完成计数增1。
	IncrCompletedCount()
	// IncrHandlingNumber 会把实时处理数增1。
	IncrHandlingNumber()
	// DecrHandlingNumber 会把实时处理数减1。
	DecrHandlingNumber()
	// Clear 用于清空所有计数。
	Clear()
}
