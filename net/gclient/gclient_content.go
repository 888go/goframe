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

// Get文本 是一个方便的方法，用于发送 GET 请求，获取并返回结果内容，并自动关闭响应对象。
// md5:20ecb29898e734f8
func (c *Client) Get文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodGet, url, 参数...))
}

// Put文本是一个方便的方法，用于发送PUT请求。它会获取并返回结果内容，并自动关闭响应对象。
// md5:7976549c1f9698ad
func (c *Client) Put文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodPut, url, 参数...))
}

// Post文本是一个方便的方法，用于发送POST请求，它会获取并返回结果内容，并自动关闭响应对象。
// md5:b7dfefd872489f24
func (c *Client) Post文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodPost, url, 参数...))
}

// Delete文本 是一个便捷方法，用于发送 DELETE 请求，它会获取并返回
// 结果内容，并自动关闭响应对象。
// md5:06595eb1a948e524
func (c *Client) Delete文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodDelete, url, 参数...))
}

// Head文本 是一个便利的方法，用于发送 HEAD 请求，获取并返回结果内容，并自动关闭响应对象。
// md5:2d973d59e0f31caf
func (c *Client) Head文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodHead, url, 参数...))
}

// Patch文本 是一个方便的方法，用于发送 PATCH 请求，它会检索并返回结果内容，并自动关闭响应对象。
// md5:c9044b57a77a7fbf
func (c *Client) Patch文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodPatch, url, 参数...))
}

// Connect文本 是一个方便的方法，用于发送CONNECT请求，它会检索并返回结果内容，并自动关闭响应对象。
// md5:8c261a851ce9854d
func (c *Client) Connect文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodConnect, url, 参数...))
}

// Options文本 是一个便捷方法，用于发送 OPTIONS 请求，它会获取并返回结果内容，
// 并自动关闭响应对象。
// md5:9d3090c51505c963
func (c *Client) Options文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodOptions, url, 参数...))
}

// Trace文本 是一个方便的方法，用于发送TRACE请求，获取并返回结果内容，并自动关闭响应对象。
// md5:b684a0ca261df5ed
func (c *Client) Trace文本(上下文 context.Context, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, http.MethodTrace, url, 参数...))
}

// X请求文本 是一个方便的方法，用于发送自定义HTTP方法请求，它会检索并返回结果内容，并自动关闭响应对象。
// md5:b57bddfc89dd348b
func (c *Client) X请求文本(上下文 context.Context, 方法 string, url string, 参数 ...interface{}) string {
	return string(c.X请求字节集(上下文, 方法, url, 参数...))
}
