package errors

import (
	"bytes"
	"fmt"
)

type ErrorType string

/**
* 面向接口编程
* 业务逻辑线提取出来，作为接口，业务具体实现通过该接口的实现类来完成
* error接口只有	func Error() string
 */

type CrawlerError interface {
	Type() ErrorType
	Error() string
}

/*
 * 使用struct实现interface
 */
type myCrawlerError struct {
	errType    ErrorType
	errMsg     string
	fullErrMsg string
}

func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawlerError{errType: errType, errMsg: errMsg}
}

/**
* 实现 CrawlerError 同时实现了 Error接口
 */
func (ce *myCrawlerError) Type() ErrorType {
	return ce.errType
}

func (ce *myCrawlerError) Error() string {
	if ce.fullErrMsg == "" {
		return ce.genFullErrMsg()
	}
	return ce.fullErrMsg
}

/**
 * struct 的私有实例方法,不暴露给外面生成 fullErrMsg
 */
func (ce *myCrawlerError) genFullErrMsg() string {
	var buffer bytes.Buffer
	buffer.WriteString("crawler err:")
	buffer.WriteString(ce.errMsg)
	return fmt.Sprintf("%s", buffer.String())

}
