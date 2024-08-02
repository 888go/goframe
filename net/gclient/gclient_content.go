// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 网页类

import (
	"context"
	"net/http"
)

// GetContent 是一个方便的方法，用于发送 GET 请求，获取并返回结果内容，并自动关闭响应对象。
// md5:20ecb29898e734f8
func (c *Client) GetContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodGet, url, data...))
}

// PutContent是一个方便的方法，用于发送PUT请求。它会获取并返回结果内容，并自动关闭响应对象。
// md5:7976549c1f9698ad
func (c *Client) PutContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodPut, url, data...))
}

// PostContent是一个方便的方法，用于发送POST请求，它会获取并返回结果内容，并自动关闭响应对象。
// md5:b7dfefd872489f24
func (c *Client) PostContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodPost, url, data...))
}

// DeleteContent 是一个便捷方法，用于发送 DELETE 请求，它会获取并返回
// 结果内容，并自动关闭响应对象。
// md5:06595eb1a948e524
func (c *Client) DeleteContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodDelete, url, data...))
}

// HeadContent 是一个便利的方法，用于发送 HEAD 请求，获取并返回结果内容，并自动关闭响应对象。
// md5:2d973d59e0f31caf
func (c *Client) HeadContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodHead, url, data...))
}

// PatchContent 是一个方便的方法，用于发送 PATCH 请求，它会检索并返回结果内容，并自动关闭响应对象。
// md5:c9044b57a77a7fbf
func (c *Client) PatchContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodPatch, url, data...))
}

// ConnectContent 是一个方便的方法，用于发送CONNECT请求，它会检索并返回结果内容，并自动关闭响应对象。
// md5:8c261a851ce9854d
func (c *Client) ConnectContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodConnect, url, data...))
}

// OptionsContent 是一个便捷方法，用于发送 OPTIONS 请求，它会获取并返回结果内容，
// 并自动关闭响应对象。
// md5:9d3090c51505c963
func (c *Client) OptionsContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodOptions, url, data...))
}

// TraceContent 是一个方便的方法，用于发送TRACE请求，获取并返回结果内容，并自动关闭响应对象。
// md5:b684a0ca261df5ed
func (c *Client) TraceContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodTrace, url, data...))
}

// RequestContent 是一个方便的方法，用于发送自定义HTTP方法请求，它会检索并返回结果内容，并自动关闭响应对象。
// md5:b57bddfc89dd348b
func (c *Client) RequestContent(ctx context.Context, method string, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, method, url, data...))
}
