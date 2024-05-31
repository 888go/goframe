// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gclient

import (
	"context"
	"net/http"
)

// GetContent is a convenience method for sending GET request, which retrieves and returns
// the result content and automatically closes response object.

// ff:Get文本
// data:参数
// url:
// ctx:上下文
func (c *Client) GetContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodGet, url, data...))
}

// PutContent is a convenience method for sending PUT request, which retrieves and returns
// the result content and automatically closes response object.

// ff:Put文本
// data:参数
// url:
// ctx:上下文
func (c *Client) PutContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodPut, url, data...))
}

// PostContent is a convenience method for sending POST request, which retrieves and returns
// the result content and automatically closes response object.

// ff:Post文本
// data:参数
// url:
// ctx:上下文
func (c *Client) PostContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodPost, url, data...))
}

// DeleteContent is a convenience method for sending DELETE request, which retrieves and returns
// the result content and automatically closes response object.

// ff:Delete文本
// data:参数
// url:
// ctx:上下文
func (c *Client) DeleteContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodDelete, url, data...))
}

// HeadContent is a convenience method for sending HEAD request, which retrieves and returns
// the result content and automatically closes response object.

// ff:Head文本
// data:参数
// url:
// ctx:上下文
func (c *Client) HeadContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodHead, url, data...))
}

// PatchContent is a convenience method for sending PATCH request, which retrieves and returns
// the result content and automatically closes response object.

// ff:Patch文本
// data:参数
// url:
// ctx:上下文
func (c *Client) PatchContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodPatch, url, data...))
}

// ConnectContent is a convenience method for sending CONNECT request, which retrieves and returns
// the result content and automatically closes response object.

// ff:Connect文本
// data:参数
// url:
// ctx:上下文
func (c *Client) ConnectContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodConnect, url, data...))
}

// OptionsContent is a convenience method for sending OPTIONS request, which retrieves and returns
// the result content and automatically closes response object.

// ff:Options文本
// data:参数
// url:
// ctx:上下文
func (c *Client) OptionsContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodOptions, url, data...))
}

// TraceContent is a convenience method for sending TRACE request, which retrieves and returns
// the result content and automatically closes response object.

// ff:Trace文本
// data:参数
// url:
// ctx:上下文
func (c *Client) TraceContent(ctx context.Context, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, http.MethodTrace, url, data...))
}

// RequestContent is a convenience method for sending custom http method request, which
// retrieves and returns the result content and automatically closes response object.

// ff:请求文本
// data:参数
// url:
// method:方法
// ctx:上下文
func (c *Client) RequestContent(ctx context.Context, method string, url string, data ...interface{}) string {
	return string(c.RequestBytes(ctx, method, url, data...))
}
