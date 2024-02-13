// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类

import (
	"context"
	"net/http"
)

// GetContent 是一个便捷方法，用于发送 GET 请求，它获取并返回结果内容，并自动关闭响应对象。
func (c *Client) Get文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodGet, url, 参数...))
}

// PutContent 是一个便捷的方法，用于发送 PUT 请求，它会获取并返回结果内容，并自动关闭响应对象。
func (c *Client) Put文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodPut, url, 参数...))
}

// PostContent 是一个便捷方法，用于发送 POST 请求，它会检索并返回结果内容，并自动关闭响应对象。
func (c *Client) Post文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodPost, url, 参数...))
}

// DeleteContent 是一个便捷方法，用于发送 DELETE 请求，该方法会获取并返回结果内容，并自动关闭响应对象。
func (c *Client) Delete文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodDelete, url, 参数...))
}

// HeadContent 是一个便捷方法，用于发送 HEAD 请求。该方法会检索并返回结果内容，并自动关闭响应对象。
func (c *Client) Head文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodHead, url, 参数...))
}

// PatchContent 是一个便捷方法，用于发送 PATCH 请求。它会获取并返回结果内容，并自动关闭响应对象。
func (c *Client) Patch文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodPatch, url, 参数...))
}

// ConnectContent 是一个便捷方法，用于发送 CONNECT 请求，它会检索并返回结果内容，并自动关闭响应对象。
func (c *Client) Connect文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodConnect, url, 参数...))
}

// OptionsContent 是一个便捷方法，用于发送 OPTIONS 请求，该方法会获取并返回结果内容，并自动关闭响应对象。
func (c *Client) Options文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodOptions, url, 参数...))
}

// TraceContent 是一个便捷方法，用于发送 TRACE 请求，该请求会检索并返回结果内容，并自动关闭响应对象。
func (c *Client) Trace文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodTrace, url, 参数...))
}

// RequestContent 是一个便捷方法，用于发送自定义 HTTP 方法请求，
// 它会获取并返回结果内容，并自动关闭响应对象。
func (c *Client) X请求文本(上下文 context.Context, 方法 string, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, 方法, url, 参数...))
}
