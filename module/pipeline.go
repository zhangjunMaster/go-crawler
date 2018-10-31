package module

/*
 * 在interface中定义大量的func类型
 */
type Pipeline interface {
	Module
	// 返回当前管道处理条目的函数
	ItemProcess() []ProcessItem
	// Send会向条目处理管道发送条目
	Send(item Item) []error
	// 只要有一个管道出错就报错
	FailFast() bool
	//设置是否快速失败
	SetFailFast(failFast bool)
}

type ProcessItem func(item Item) (result Item, err error)
