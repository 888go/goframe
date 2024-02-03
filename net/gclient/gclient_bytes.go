// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gclient

import (
	"context"
	"net/http"
	
	"github.com/888go/goframe/internal/intlog"
)

// GetBytes 发送一个GET请求，获取并以字节形式返回结果内容。
func (c *Client) GetBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodGet, url, data...)
}

// PutBytes 发送一个 PUT 请求，获取并以字节形式返回结果内容。
func (c *Client) PutBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodPut, url, data...)
}

// PostBytes 发送一个POST请求，获取并以字节形式返回结果内容。
func (c *Client) PostBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodPost, url, data...)
}

// DeleteBytes 发送一个 DELETE 请求，获取并以字节形式返回结果内容。
func (c *Client) DeleteBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodDelete, url, data...)
}

// HeadBytes 发送一个 HEAD 请求，获取并以字节形式返回结果内容。
func (c *Client) HeadBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodHead, url, data...)
}

// PatchBytes 发送一个 PATCH 请求，获取并以字节形式返回结果内容。
func (c *Client) PatchBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodPatch, url, data...)
}

// ConnectBytes 发送一个 CONNECT 请求，获取并以字节形式返回结果内容。
func (c *Client) ConnectBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodConnect, url, data...)
}

// OptionsBytes 发送一个 OPTIONS 请求，获取并以字节形式返回结果内容。
func (c *Client) OptionsBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodOptions, url, data...)
}

// TraceBytes 发送一个 TRACE 请求，获取并以字节形式返回结果内容。
func (c *Client) TraceBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodTrace, url, data...)
}

// RequestBytes 使用给定的HTTP方法和数据发送请求，然后以字节形式返回结果。
// 它会自动内部读取并关闭响应对象。
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
