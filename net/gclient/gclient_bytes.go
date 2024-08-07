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

	"github.com/888go/goframe/internal/intlog"
)

// Get字节集 发送一个GET请求，并获取结果内容，将其作为字节切片返回。 md5:1814e36a86c5ea59
func (c *Client) Get字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodGet, url, 参数...)
}

// Put字节集 发送一个PUT请求，获取并返回结果内容为字节。 md5:50c5c9707e1b037e
func (c *Client) Put字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodPut, url, 参数...)
}

// Post字节集 发送一个POST请求，获取并返回结果内容作为字节。 md5:2d211f331f5cac73
func (c *Client) Post字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodPost, url, 参数...)
}

// Delete字节集 发送一个 DELETE 请求，获取并返回结果内容作为字节。 md5:6ae662cb66a0df6a
func (c *Client) Delete字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodDelete, url, 参数...)
}

// Head字节集 发送一个HEAD请求，获取并以字节形式返回结果内容。 md5:8f0602e1d7086700
func (c *Client) Head字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodHead, url, 参数...)
}

// Patch字节集 发送一个PATCH请求，获取并返回结果内容为字节。 md5:821d9f04def319cf
func (c *Client) Patch字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodPatch, url, 参数...)
}

// Connect字节集 发送一个CONNECT请求，获取并返回结果内容作为字节。 md5:32db7a6df27fca73
func (c *Client) Connect字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodConnect, url, 参数...)
}

// Options字节集 发送一个OPTIONS请求，获取并返回结果内容作为字节。 md5:34e2f77f73d41f0d
func (c *Client) Options字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodOptions, url, 参数...)
}

// Trace字节集发送一个TRACE请求，并获取结果内容，以字节形式返回。 md5:ab95bf40e1213375
func (c *Client) Trace字节集(上下文 context.Context, url string, 参数 ...interface{}) []byte {
	return c.X请求字节集(上下文, http.MethodTrace, url, 参数...)
}

// X请求字节集 使用给定的HTTP方法和数据发送请求，并将结果作为字节返回。它会自动读取并关闭响应对象。
// md5:9c3a91009734fc7a
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
