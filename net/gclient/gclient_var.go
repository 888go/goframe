// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gclient

import (
	"context"
	"net/http"
	
	"github.com/888go/goframe/container/gvar"
)

// GetVar 发送一个GET请求，获取并转换结果内容为指定的指针类型。
// 参数`pointer`可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct等。
func (c *Client) GetVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodGet, url, data...)
}

// PutVar 发送一个 PUT 请求，获取并转换结果内容为指定的指针类型。
// 参数 `pointer` 可以是以下类型：struct、*struct、**struct、[]struct、[]*struct、*[]struct 等等。
func (c *Client) PutVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodPut, url, data...)
}

// PostVar 发送一个POST请求，获取并转换结果内容为指定的指针类型。
// 参数`pointer`可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct等。
func (c *Client) PostVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodPost, url, data...)
}

// DeleteVar 发送一个 DELETE 请求，获取并转换结果内容为指定的指针类型。
// 参数 `pointer` 可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct 等等。
func (c *Client) DeleteVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodDelete, url, data...)
}

// HeadVar 发送一个 HEAD 请求，获取并转换结果内容到指定的指针。
// 参数 `pointer` 可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct 等等。
func (c *Client) HeadVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodHead, url, data...)
}

// PatchVar 发送一个 PATCH 请求，获取并转换结果内容为指定的指针类型。
// 参数 `pointer` 的类型可以是：struct/*struct/**struct/[]struct/[]*struct/*[]struct 等。
func (c *Client) PatchVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodPatch, url, data...)
}

// ConnectVar 发送一个 CONNECT 请求，获取并转换结果内容到指定的指针。
// 参数 `pointer` 的类型可以是：struct/*struct/**struct/[]struct/[]*struct/*[]struct 等。
func (c *Client) ConnectVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodConnect, url, data...)
}

// OptionsVar 发送一个 OPTIONS 请求，获取并转换结果内容到指定的指针。
// 参数 `pointer` 的类型可以是：struct/*struct/[]struct/[]*struct/*[]struct 等。
func (c *Client) OptionsVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodOptions, url, data...)
}

// TraceVar 发送一个 TRACE 请求，获取并转换结果内容到指定的指针。
// 参数 `pointer` 可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct 等等。
func (c *Client) TraceVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodTrace, url, data...)
}

// RequestVar 使用给定的HTTP方法和数据发送请求，检索并将结果转换为指定的指针。
// 它会自动内部读取并关闭响应对象。
// 参数`pointer`可以是以下类型：struct/*struct/**struct/[]struct/[]*struct/*[]struct等。
func (c *Client) RequestVar(ctx context.Context, method string, url string, data ...interface{}) *gvar.Var {
	response, err := c.DoRequest(ctx, method, url, data...)
	if err != nil {
		return gvar.New(nil)
	}
	defer response.Close()
	return gvar.New(response.ReadAll())
}
