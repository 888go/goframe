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

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/internal/intlog"
)

// GetVar 发送一个GET请求，获取并转换结果内容为*gvar.Var。
// 客户端会自动内部读取并关闭响应对象。
// 结果*gvar.Var可以方便地转换为你想要的任何类型。
// md5:47d99fdd051eb344
// ff:Get泛型类
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) GetVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodGet, url, data...)
}

// PutVar 发送一个PUT请求，获取并转换结果内容为*gvar.Var类型。
// 客户端会自动内部读取并关闭响应对象。
// 结果*gvar.Var可以方便地转换为你想要的任何类型。
// md5:cb882faeaad7615d
// ff:Put泛型类
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) PutVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodPut, url, data...)
}

// PostVar 发送一个POST请求，获取并转换结果内容为*gvar.Var。
// 客户端内部会自动读取并关闭响应对象。
// 返回的*gvar.Var可以方便地转换为你想要的任何类型。
// md5:51e0af70f2e10306
// ff:Post泛型类
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) PostVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodPost, url, data...)
}

// DeleteVar 发送一个 DELETE 请求，获取并把结果内容转换为 *gvar.Var。
// 客户端会自动内部读取并关闭响应对象。
// 结果 *gvar.Var 可以方便地转换为你需要的任何类型。
// md5:1954cc9f6686cac5
// ff:Delete泛型类
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) DeleteVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodDelete, url, data...)
}

// HeadVar 发送一个HEAD请求，获取并转换结果内容为*gvar.Var。
// 客户端会自动内部读取并关闭响应对象。
// 结果*gvar.Var可以方便地转换为你想要的任何类型。
// md5:d4338600d66a6acf
// ff:Head泛型类
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) HeadVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodHead, url, data...)
}

// PatchVar 发送一个PATCH请求，获取并转换结果内容为*gvar.Var类型。
// 客户端会自动内部读取并关闭响应对象。
// 结果*gvar.Var可以方便地转换为你想要的任何类型。
// md5:9786881ed5fc525d
// ff:Patch泛型类
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) PatchVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodPatch, url, data...)
}

// ConnectVar 发送一个 CONNECT 请求，获取并转换结果内容为 *gvar.Var。
// 客户端会自动内部读取并关闭响应对象。
// 得到的 *gvar.Var 可以方便地转换为你想要的任何类型。
// md5:152bf738c9672ae0
// ff:Connect泛型类
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) ConnectVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodConnect, url, data...)
}

// OptionsVar 发送一个 OPTIONS 请求，获取并转换结果内容为 *gvar.Var。
// 客户端会自动读取并关闭响应对象。
// 结果 *gvar.Var 可以方便地转换为你需要的任何类型。
// md5:9ab65dcef0e9422e
// ff:Options泛型类
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) OptionsVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodOptions, url, data...)
}

// TraceVar 发送一个 TRACE 请求，获取并转换结果内容为 *gvar.Var 类型。
// 客户端会自动内部读取并关闭响应对象。
// 结果的 *gvar.Var 可以方便地转换为你想要的任何类型。
// md5:560c30dddc223d10
// ff:Trace泛型类
// c:
// ctx:上下文
// url:
// data:参数
func (c *Client) TraceVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {
	return c.RequestVar(ctx, http.MethodTrace, url, data...)
}

// RequestVar 使用给定的HTTP方法和数据发送请求，然后将结果转换为*gvar.Var类型。
// 客户端会自动内部读取并关闭响应对象。
// 结果*gvar.Var可以方便地转换为你想要的任何类型。
// md5:4f9fa2909af65cbb
// ff:请求泛型类
// c:
// ctx:上下文
// method:方法
// url:
// data:参数
func (c *Client) RequestVar(ctx context.Context, method string, url string, data ...interface{}) *gvar.Var {
	response, err := c.DoRequest(ctx, method, url, data...)
	if err != nil {
		intlog.Errorf(ctx, `%+v`, err)
		return gvar.New(nil)
	}
	defer func() {
		if err = response.Close(); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
	}()
	return gvar.New(response.ReadAll())
}
