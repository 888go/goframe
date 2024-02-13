// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类

import (
	"context"
	"net/http"
	
	"github.com/888go/goframe/container/gvar"
)

// GetVar 发送一个GET请求，获取并转换结果内容为指定的指针类型。
// 参数`pointer`可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct等。
func (c *Client) Get泛型类(上下文 context.Context, url string, 参数 ...interface{}) *泛型类.Var {
	return c.X请求泛型类(上下文, http.MethodGet, url, 参数...)
}

// PutVar 发送一个 PUT 请求，获取并转换结果内容为指定的指针类型。
// 参数 `pointer` 可以是以下类型：struct、*struct、**struct、[]struct、[]*struct、*[]struct 等等。
func (c *Client) Put泛型类(上下文 context.Context, url string, 参数 ...interface{}) *泛型类.Var {
	return c.X请求泛型类(上下文, http.MethodPut, url, 参数...)
}

// PostVar 发送一个POST请求，获取并转换结果内容为指定的指针类型。
// 参数`pointer`可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct等。
func (c *Client) Post泛型类(上下文 context.Context, url string, 参数 ...interface{}) *泛型类.Var {
	return c.X请求泛型类(上下文, http.MethodPost, url, 参数...)
}

// DeleteVar 发送一个 DELETE 请求，获取并转换结果内容为指定的指针类型。
// 参数 `pointer` 可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct 等等。
func (c *Client) Delete泛型类(上下文 context.Context, url string, 参数 ...interface{}) *泛型类.Var {
	return c.X请求泛型类(上下文, http.MethodDelete, url, 参数...)
}

// HeadVar 发送一个 HEAD 请求，获取并转换结果内容到指定的指针。
// 参数 `pointer` 可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct 等等。
func (c *Client) Head泛型类(上下文 context.Context, url string, 参数 ...interface{}) *泛型类.Var {
	return c.X请求泛型类(上下文, http.MethodHead, url, 参数...)
}

// PatchVar 发送一个 PATCH 请求，获取并转换结果内容为指定的指针类型。
// 参数 `pointer` 的类型可以是：struct/*struct/**struct/[]struct/[]*struct/*[]struct 等。
func (c *Client) Patch泛型类(上下文 context.Context, url string, 参数 ...interface{}) *泛型类.Var {
	return c.X请求泛型类(上下文, http.MethodPatch, url, 参数...)
}

// ConnectVar 发送一个 CONNECT 请求，获取并转换结果内容到指定的指针。
// 参数 `pointer` 的类型可以是：struct/*struct/**struct/[]struct/[]*struct/*[]struct 等。
func (c *Client) Connect泛型类(上下文 context.Context, url string, 参数 ...interface{}) *泛型类.Var {
	return c.X请求泛型类(上下文, http.MethodConnect, url, 参数...)
}

// OptionsVar 发送一个 OPTIONS 请求，获取并转换结果内容到指定的指针。
// 参数 `pointer` 的类型可以是：struct/*struct/[]struct/[]*struct/*[]struct 等。
func (c *Client) Options泛型类(上下文 context.Context, url string, 参数 ...interface{}) *泛型类.Var {
	return c.X请求泛型类(上下文, http.MethodOptions, url, 参数...)
}

// TraceVar 发送一个 TRACE 请求，获取并转换结果内容到指定的指针。
// 参数 `pointer` 可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct 等等。
func (c *Client) Trace泛型类(上下文 context.Context, url string, 参数 ...interface{}) *泛型类.Var {
	return c.X请求泛型类(上下文, http.MethodTrace, url, 参数...)
}

// RequestVar 使用给定的HTTP方法和数据发送请求，检索并将结果转换为指定的指针。
// 它会自动内部读取并关闭响应对象。
// 参数`pointer`可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct等。
func (c *Client) X请求泛型类(上下文 context.Context, 方法 string, url string, 参数 ...interface{}) *泛型类.Var {
	response, err := c.X请求响应对象(上下文, 方法, url, 参数...)
	if err != nil {
		return 泛型类.X创建(nil)
	}
	defer response.X关闭()
	return 泛型类.X创建(response.X取响应字节集())
}
