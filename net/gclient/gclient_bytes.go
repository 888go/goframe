// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gclient

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/internal/intlog"
)

// GetBytes 发送一个GET请求，并获取结果内容，将其作为字节切片返回。 md5:1814e36a86c5ea59
// ff:Get字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) GetBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodGet, url, data...)
}

// PutBytes 发送一个PUT请求，获取并返回结果内容为字节。 md5:50c5c9707e1b037e
// ff:Put字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) PutBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodPut, url, data...)
}

// PostBytes 发送一个POST请求，获取并返回结果内容作为字节。 md5:2d211f331f5cac73
// ff:Post字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) PostBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodPost, url, data...)
}

// DeleteBytes 发送一个 DELETE 请求，获取并返回结果内容作为字节。 md5:6ae662cb66a0df6a
// ff:Delete字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) DeleteBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodDelete, url, data...)
}

// HeadBytes 发送一个HEAD请求，获取并以字节形式返回结果内容。 md5:8f0602e1d7086700
// ff:Head字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) HeadBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodHead, url, data...)
}

// PatchBytes 发送一个PATCH请求，获取并返回结果内容为字节。 md5:821d9f04def319cf
// ff:Patch字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) PatchBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodPatch, url, data...)
}

// ConnectBytes 发送一个CONNECT请求，获取并返回结果内容作为字节。 md5:32db7a6df27fca73
// ff:Connect字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) ConnectBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodConnect, url, data...)
}

// OptionsBytes 发送一个OPTIONS请求，获取并返回结果内容作为字节。 md5:34e2f77f73d41f0d
// ff:Options字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) OptionsBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodOptions, url, data...)
}

// TraceBytes发送一个TRACE请求，并获取结果内容，以字节形式返回。 md5:ab95bf40e1213375
// ff:Trace字节集
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) TraceBytes(ctx context.Context, url string, data ...interface{}) []byte {
	return c.RequestBytes(ctx, http.MethodTrace, url, data...)
}

// RequestBytes 使用给定的HTTP方法和数据发送请求，并将结果作为字节返回。它会自动读取并关闭响应对象。
// md5:9c3a91009734fc7a
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
