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

	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/intlog"
)

// Get泛型类 发送一个GET请求，获取并转换结果内容为*gvar.Var。
// 客户端会自动内部读取并关闭响应对象。
// 结果*gvar.Var可以方便地转换为你想要的任何类型。
// md5:47d99fdd051eb344
func (c *Client) Get泛型类(上下文 context.Context, url string, 参数 ...interface{}) *gvar.Var {
	return c.X请求泛型类(上下文, http.MethodGet, url, 参数...)
}

// Put泛型类 发送一个PUT请求，获取并转换结果内容为*gvar.Var类型。
// 客户端会自动内部读取并关闭响应对象。
// 结果*gvar.Var可以方便地转换为你想要的任何类型。
// md5:cb882faeaad7615d
func (c *Client) Put泛型类(上下文 context.Context, url string, 参数 ...interface{}) *gvar.Var {
	return c.X请求泛型类(上下文, http.MethodPut, url, 参数...)
}

// Post泛型类 发送一个POST请求，获取并转换结果内容为*gvar.Var。
// 客户端内部会自动读取并关闭响应对象。
// 返回的*gvar.Var可以方便地转换为你想要的任何类型。
// md5:51e0af70f2e10306
func (c *Client) Post泛型类(上下文 context.Context, url string, 参数 ...interface{}) *gvar.Var {
	return c.X请求泛型类(上下文, http.MethodPost, url, 参数...)
}

// Delete泛型类 发送一个 DELETE 请求，获取并把结果内容转换为 *gvar.Var。
// 客户端会自动内部读取并关闭响应对象。
// 结果 *gvar.Var 可以方便地转换为你需要的任何类型。
// md5:1954cc9f6686cac5
func (c *Client) Delete泛型类(上下文 context.Context, url string, 参数 ...interface{}) *gvar.Var {
	return c.X请求泛型类(上下文, http.MethodDelete, url, 参数...)
}

// Head泛型类 发送一个HEAD请求，获取并转换结果内容为*gvar.Var。
// 客户端会自动内部读取并关闭响应对象。
// 结果*gvar.Var可以方便地转换为你想要的任何类型。
// md5:d4338600d66a6acf
func (c *Client) Head泛型类(上下文 context.Context, url string, 参数 ...interface{}) *gvar.Var {
	return c.X请求泛型类(上下文, http.MethodHead, url, 参数...)
}

// Patch泛型类 发送一个PATCH请求，获取并转换结果内容为*gvar.Var类型。
// 客户端会自动内部读取并关闭响应对象。
// 结果*gvar.Var可以方便地转换为你想要的任何类型。
// md5:9786881ed5fc525d
func (c *Client) Patch泛型类(上下文 context.Context, url string, 参数 ...interface{}) *gvar.Var {
	return c.X请求泛型类(上下文, http.MethodPatch, url, 参数...)
}

// Connect泛型类 发送一个 CONNECT 请求，获取并转换结果内容为 *gvar.Var。
// 客户端会自动内部读取并关闭响应对象。
// 得到的 *gvar.Var 可以方便地转换为你想要的任何类型。
// md5:152bf738c9672ae0
func (c *Client) Connect泛型类(上下文 context.Context, url string, 参数 ...interface{}) *gvar.Var {
	return c.X请求泛型类(上下文, http.MethodConnect, url, 参数...)
}

// Options泛型类 发送一个 OPTIONS 请求，获取并转换结果内容为 *gvar.Var。
// 客户端会自动读取并关闭响应对象。
// 结果 *gvar.Var 可以方便地转换为你需要的任何类型。
// md5:9ab65dcef0e9422e
func (c *Client) Options泛型类(上下文 context.Context, url string, 参数 ...interface{}) *gvar.Var {
	return c.X请求泛型类(上下文, http.MethodOptions, url, 参数...)
}

// Trace泛型类 发送一个 TRACE 请求，获取并转换结果内容为 *gvar.Var 类型。
// 客户端会自动内部读取并关闭响应对象。
// 结果的 *gvar.Var 可以方便地转换为你想要的任何类型。
// md5:560c30dddc223d10
func (c *Client) Trace泛型类(上下文 context.Context, url string, 参数 ...interface{}) *gvar.Var {
	return c.X请求泛型类(上下文, http.MethodTrace, url, 参数...)
}

// X请求泛型类 使用给定的HTTP方法和数据发送请求，然后将结果转换为*gvar.Var类型。
// 客户端会自动内部读取并关闭响应对象。
// 结果*gvar.Var可以方便地转换为你想要的任何类型。
// md5:4f9fa2909af65cbb
func (c *Client) X请求泛型类(上下文 context.Context, 方法 string, url string, 参数 ...interface{}) *gvar.Var {
	response, err := c.X请求响应对象(上下文, 方法, url, 参数...)
	if err != nil {
		intlog.Errorf(上下文, `%+v`, err)
		return gvar.X创建(nil)
	}
	defer func() {
		if err = response.X关闭(); err != nil {
			intlog.Errorf(上下文, `%+v`, err)
		}
	}()
	return gvar.X创建(response.X取响应字节集())
}
