package 网页类

import (
	"net/http"

	gctx "github.com/888go/goframe/os/gctx"
)

// HandlerFunc 是用于处理中间件的处理器函数. md5:1e0565dffbfa907c
type HandlerFunc = func(c *Client, r *http.Request) (*Response, error)

// clientMiddleware 是用于管理HTTP客户端请求工作流程的插件。 md5:72add0d1b66ac073
type clientMiddleware struct {
	client       *Client       // http client.
	handlers     []HandlerFunc // mdl handlers.
	handlerIndex int           // current handler index.
	resp         *Response     // save resp.
	err          error         // save err.
}

const clientMiddlewareKey gctx.StrKey = "__clientMiddlewareKey"

// Use 向客户端添加一个或多个中间件处理器。 md5:92665269b902692e
func (c *Client) Use(handlers ...HandlerFunc) *Client {
	c.middlewareHandler = append(c.middlewareHandler, handlers...)
	return c
}

// Next 调用下一个中间件。
// 这应该只在 HandlerFunc 中调用。
// md5:70c74664d7d9f919
func (c *Client) Next(req *http.Request) (*Response, error) {
	if v := req.Context().Value(clientMiddlewareKey); v != nil {
		if m, ok := v.(*clientMiddleware); ok {
			return m.Next(req)
		}
	}
	return c.callRequest(req)
}

// Next 调用下一个中间件处理器。 md5:51a6ca6a21a9942e
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
