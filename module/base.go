package module

import "net/http"

// 对基础的string，但是又意义的都是取用别名
// 先是其他的type
// struct interface

type Counts struct {
	// 调用计数
	CalledCount uint64
	// 接受计数
	AcceptedCount uint64
	// 完成计数
	CompletedCount uint64

	HandlingCount uint64
}

type SummaryStruct struct {
	ID        MID         `json:"id"`
	Called    uint64      `json:"called"`
	Accepted  uint64      `json:"accepted"`
	Completed uint64      `json:"completed"`
	Handling  uint64      `json:"handling"`
	Extra     interface{} `json:"extra"`
}

/**
 * 基础接口
 * 下载，条目，调度 子接口继承基础接口
 * 然后根据基础接口去写内部属性的接口
 */
type Module interface {
	//获取组件id
	ID() MID
	//获取地址
	Addr() string
	//获取组件评分
	Score() uint64
	// 设置当前组件评分
	SetScore(score uint64)
	// 获取评分计算器
	ScoreCalculator()
	// 获取当前组件调用的次数
	CalledCount() uint64
	// 组件接受的调用次数
	AcceptedCount() uint64
	// 组件完成的调用次数
	CompletedCount() uint64
	// 当前组件正在处理的调用次数
	HandlingNumber() uint64
	//一次性获取所有计数
	Counts() Counts
	// 获取组件摘要
	Summary() SummaryStruct
}

/**
 * ID是复杂的构成方式
 * 生成ID的接口
 */
type SNGenerator interface {
	// 获取预设的最小序列号
	Start() uint64
	Max() uint64
	Next() uint64
	// 获取循环计数
	CycleCount() uint64
	// 获取一个序列号，并计算下一个序列号
	Get() uint64
}

type CalculateScore func(counts Counts) uint64

type Downloader interface {
	Module
	Download(req *Request) (*Response, error)
}

type Analyzer interface {
	Module
	//返回解析函数列表
	RespParsers() []ParseResponse
	Analyze() ([]Data, error)
}

//解析函数
type ParseResponse func(httpResp *http.Response, respDepth uint32) ([]Data, error)
