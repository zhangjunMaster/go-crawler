package scheduler

import "go-crawler/module"

// ModuleArgsSummary 代表组件相关的参数容器的摘要类型。
type ModuleArgsSummary struct {
	DownloaderListSize int `json:"downloader_list_size"`
	AnalyzerListSize   int `json:"analyzer_list_size"`
	PipelineListSize   int `json:"pipeline_list_size"`
}

type DataArgs struct {
	ReqBufferCap         uint32 `json:"req_buf_cap"`
	ReqMaxBufferNumber   uint32 `json:"req_max_buffer_number"`
	RespBufferCap        uint32 `json:"resp_buffer_cap"`
	RespMaxBufferNumber  uint32 `json:"resp_max_buffer_number"`
	ItemBufferCap        uint32 `json:"item_buffer_cap"`
	ItemMaxBufferNumber  uint32 `json:"item_max_buffer_number"`
	ErrorBufferCap       uint32 `json:"error_buffer_cap"`
	ErrorMaxBufferNumber uint32 `json:"error_max_buffer_number"`
}

type RequestArgs struct {
	AcceptedDomains []string `json:"accepted_primary_domains"`
	MaxDepth        uint32   `json:"max_depth"`
}

type ModuleArgs struct {
	Downloaders []module.Downloader
	Analyzers   []module.Analyzer
	Pipelines   []module.Pipeline
}

type Args interface {
	Check() error
}
