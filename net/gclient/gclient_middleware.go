package gclient

import (
	"net/http"
	
	"github.com/888go/goframe/os/gctx"
)

// HandlerFunc 中间件处理函数
type HandlerFunc = func(c *Client, r *http.Request) (*Response, error)

// clientMiddleware 是用于 HTTP 客户端请求流程管理的插件。
type clientMiddleware struct {
	client       *Client       // http client.
	handlers     []HandlerFunc // mdl handlers.
	handlerIndex int           // 当前处理器索引。
	resp         *Response     // save resp.
	err          error         // save err.
}

const clientMiddlewareKey gctx.StrKey = "__clientMiddlewareKey"

// Use 向客户端添加一个或多个中间件处理器。
func (c *Client) Use(handlers ...HandlerFunc) *Client {
	c.middlewareHandler = append(c.middlewareHandler, handlers...)
	return c
}

// Next 调用下一个中间件。
// 该函数仅应在 HandlerFunc 中调用。
func (c *Client) Next(req *http.Request) (*Response, error) {
	if v := req.Context().Value(clientMiddlewareKey); v != nil {
		if m, ok := v.(*clientMiddleware); ok {
			return m.Next(req)
		}
	}
	return c.callRequest(req)
}

// Next调用下一个中间件处理器。
func (m *clientMiddleware) Next(req *http.Request) (resp *Response, err error) {
	if m.err != nil {
		return m.resp, m.err
	}
	if m.handlerIndex < len(m.handlers) {
		m.handlerIndex++
		m.resp, m.err = m.handlers[m.handlerIndex](m.client, req)
	}
	return m.resp, m.err
}
