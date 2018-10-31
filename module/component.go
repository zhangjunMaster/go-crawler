package module

import "net/http"

/**
 * private Request
 * newRequest, property
 */
type Request struct {
	httpReq *http.Request
	depth   uint32
}

func newRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq, depth}
}

func (r *Request) HTTPReq() *http.Request {
	return r.httpReq
}

func (r *Request) Depth() uint32 {
	return r.depth
}

/**
 * response
 */
type Response struct {
	httpResp *http.Response
	depth    uint32
}

func newResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp, depth}
}

func (r *Response) HTTPResp() *http.Response {
	return r.httpResp
}

func (r *Response) Depth() uint32 {
	return r.depth
}

/**
 * crawler item 条目
 */
type Item map[string]interface{}

/**
 * 判断数据是否合法
 */
type Data interface {
	valid() bool
}

func (req *Request) valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

func (resp *Response) valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

func (item Item) valid() bool {
	return item != nil
}
