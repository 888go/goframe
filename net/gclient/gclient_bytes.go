// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gclient

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/internal/intlog"
)

// GetBytes sends a GET request, retrieves and returns the result content as bytes.
// ff:Get字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) GetBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodGet, url, data...)
}

// PutBytes sends a PUT request, retrieves and returns the result content as bytes.
// ff:Put字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) PutBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodPut, url, data...)
}

// PostBytes sends a POST request, retrieves and returns the result content as bytes.
// ff:Post字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) PostBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodPost, url, data...)
}

// DeleteBytes sends a DELETE request, retrieves and returns the result content as bytes.
// ff:Delete字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) DeleteBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodDelete, url, data...)
}

// HeadBytes sends a HEAD request, retrieves and returns the result content as bytes.
// ff:Head字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) HeadBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodHead, url, data...)
}

// PatchBytes sends a PATCH request, retrieves and returns the result content as bytes.
// ff:Patch字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) PatchBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodPatch, url, data...)
}

// ConnectBytes sends a CONNECT request, retrieves and returns the result content as bytes.
// ff:Connect字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) ConnectBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodConnect, url, data...)
}

// OptionsBytes sends an OPTIONS request, retrieves and returns the result content as bytes.
// ff:Options字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) OptionsBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodOptions, url, data...)
}

// TraceBytes sends a TRACE request, retrieves and returns the result content as bytes.
// ff:Trace字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) TraceBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodTrace, url, data...)
}

// RequestBytes sends request using given HTTP method and data, retrieves returns the result
// as bytes. It reads and closes the response object internally automatically.
// ff:请求字节集
// c:
// ctx:上下文
// method:方法
// url:
// data:参数
func (c *Client) RequestBytes(ctx context.Context, method string, url string, data ...interface{}) []byte {
	response, err := c.DoRequest(ctx, method, url, data...)
	if err != nil {
		intlog.Errorf(ctx, `%+v`, err)
		return nil
	}
	defer func() {
		if err = response.Close(); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
	}()
	return response.ReadAll()
}
