package errors

/**
 * 错误常量，类型是 ErrorType
 */
const (
	//下载错误
	ERROR_TYPE_DOWNLOAD ErrorType = "downloader error"
	//分析器错误
	ERROR_TYPE_ANALYZER ErrorType = "analyzer anerroralyzer"
	//条目管理错误
	ERROR_TYPE_PIPELINE ErrorType = "pipeline error"
	//调度错误
	ERROR_TYPE_SCHEDULE ErrorType = "schedule error"
)
