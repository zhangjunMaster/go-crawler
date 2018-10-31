package scheduler

import "go-crawler/module"

type SchedSummary interface {
	Struct() SummaryStruct
	String() string
}

type SummaryStruct struct {
	RequestArgs     RequestArgs             `json:"request_args"`
	DataArgs        DataArgs                `json:"data_args"`
	ModuleArgs      ModuleArgsSummary       `json:"module_args"`
	Status          string                  `json:"status"`
	Downloaders     []module.SummaryStruct  `json:"downloaders"`
	Analyzer        []module.SummaryStruct  `json:"analyzer"`
	Pipelines       []module.SummaryStruct  `json:"pipelines"`
	ReqBufferPool   BufferPoolSummaryStruct `json:"req_buffer_pool"`
	RespBufferPool  BufferPoolSummaryStruct `json:"resp_buffer_pool"`
	ItemBufferPool  BufferPoolSummaryStruct `json:"item_buffer_pool"`
	ErrorBufferPool BufferPoolSummaryStruct `json:"error_buffer_pool"`
	NumURL          uint64                  `json:"url_number"`
}

// BufferPoolSummaryStruct 代表缓冲池的摘要类型。
type BufferPoolSummaryStruct struct {
	BufferCap       uint32 `json:"buffer_cap"`
	MaxBufferNumber uint32 `json:"max_buffer_number"`
	BufferNumber    uint32 `json:"buffer_number"`
	Total           uint64 `json:"total"`
}
