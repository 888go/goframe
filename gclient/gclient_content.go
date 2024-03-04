// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gclient

import (
	"context"
	"net/http"
)

// GetContent 是一个便捷方法，用于发送 GET 请求，它获取并返回结果内容，并自动关闭响应对象。
func (c *Client) GetContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodGet, url, data...))
}

// PutContent 是一个便捷的方法，用于发送 PUT 请求，它会获取并返回结果内容，并自动关闭响应对象。
func (c *Client) PutContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodPut, url, data...))
}

// PostContent 是一个便捷方法，用于发送 POST 请求，它会检索并返回结果内容，并自动关闭响应对象。
func (c *Client) PostContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodPost, url, data...))
}

// DeleteContent 是一个便捷方法，用于发送 DELETE 请求，该方法会获取并返回结果内容，并自动关闭响应对象。
func (c *Client) DeleteContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodDelete, url, data...))
}

// HeadContent 是一个便捷方法，用于发送 HEAD 请求。该方法会检索并返回结果内容，并自动关闭响应对象。
func (c *Client) HeadContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodHead, url, data...))
}

// PatchContent 是一个便捷方法，用于发送 PATCH 请求。它会获取并返回结果内容，并自动关闭响应对象。
func (c *Client) PatchContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodPatch, url, data...))
}

// ConnectContent 是一个便捷方法，用于发送 CONNECT 请求，它会检索并返回结果内容，并自动关闭响应对象。
func (c *Client) ConnectContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodConnect, url, data...))
}

// OptionsContent 是一个便捷方法，用于发送 OPTIONS 请求，该方法会获取并返回结果内容，并自动关闭响应对象。
func (c *Client) OptionsContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodOptions, url, data...))
}

// TraceContent 是一个便捷方法，用于发送 TRACE 请求，该请求会检索并返回结果内容，并自动关闭响应对象。
func (c *Client) TraceContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodTrace, url, data...))
}

// RequestContent 是一个便捷方法，用于发送自定义 HTTP 方法请求，
// 它会获取并返回结果内容，并自动关闭响应对象。
func (c *Client) RequestContent(ctx context.Context, method string, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, method, url, data...))
}
