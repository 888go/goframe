// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类

import (
	"context"
	"net/http"
	
	"github.com/888go/goframe/internal/intlog"
)

// GetBytes 发送一个GET请求，获取并以字节形式返回结果内容。
func (c *Client) Get字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodGet, url, 参数...)
}

// PutBytes 发送一个 PUT 请求，获取并以字节形式返回结果内容。
func (c *Client) Put字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodPut, url, 参数...)
}

// PostBytes 发送一个POST请求，获取并以字节形式返回结果内容。
func (c *Client) Post字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodPost, url, 参数...)
}

// DeleteBytes 发送一个 DELETE 请求，获取并以字节形式返回结果内容。
func (c *Client) Delete字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodDelete, url, 参数...)
}

// HeadBytes 发送一个 HEAD 请求，获取并以字节形式返回结果内容。
func (c *Client) Head字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodHead, url, 参数...)
}

// PatchBytes 发送一个 PATCH 请求，获取并以字节形式返回结果内容。
func (c *Client) Patch字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodPatch, url, 参数...)
}

// ConnectBytes 发送一个 CONNECT 请求，获取并以字节形式返回结果内容。
func (c *Client) Connect字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodConnect, url, 参数...)
}

// OptionsBytes 发送一个 OPTIONS 请求，获取并以字节形式返回结果内容。
func (c *Client) Options字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodOptions, url, 参数...)
}

// TraceBytes 发送一个 TRACE 请求，获取并以字节形式返回结果内容。
func (c *Client) Trace字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodTrace, url, 参数...)
}

// RequestBytes 使用给定的HTTP方法和数据发送请求，然后以字节形式返回结果。
// 它会自动内部读取并关闭响应对象。
func (c *Client) X请求字节集(上下文 context.Context, 方法 string, url string, 参数 ...interface{}) []byte {
	response, err := c.X请求响应对象(上下文, 方法, url, 参数...)
	if err != nil {
		intlog.Errorf(上下文, `%+v`, err)
		return nil
	}
	defer func() {
		if err = response.X关闭(); err != nil {
			intlog.Errorf(上下文, `%+v`, err)
		}
	}()
	return response.X取响应字节集()
}
