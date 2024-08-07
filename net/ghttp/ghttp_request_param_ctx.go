// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"context"

	gvar "github.com/888go/goframe/container/gvar"
	gctx "github.com/888go/goframe/os/gctx"
)

// X从上下文取请求对象 从上下文中检索并返回Request对象。 md5:c247eac3d031fb2b
func X从上下文取请求对象(上下文 context.Context) *Request {
	if v := 上下文.Value(ctxKeyForRequest); v != nil {
		return v.(*Request)
	}
	return nil
}

// Context别名 是函数 GetCtx 的别名。
// 这个函数重写了 http.Request.Context别名 函数。
// 参考 GetCtx 函数。
// md5:cc1d0847fad835ae
func (r *Request) Context别名() context.Context {
	var ctx = r.Request.Context()
		// 检查并把Request对象注入到context中。 md5:6db244059baa2f77
	if X从上下文取请求对象(ctx) == nil {
				// 将 Request 对象注入到上下文中。 md5:561b491dd0877f18
		ctx = context.WithValue(ctx, ctxKeyForRequest, r)
				// 如果使用默认跟踪提供程序，添加默认跟踪信息。 md5:ffe39a076b5cfe89
		ctx = gctx.X创建并从上下文(ctx)
				// 更新原始HTTP请求的值。 md5:b7f6f8f267437800
		*r.Request = *r.Request.WithContext(ctx)
	}
	return ctx
}

// X取上下文对象 获取并返回请求的上下文。
// 它是函数Context的别名，与SetCtx函数相关联。
// md5:67b84a74829ae7e1
func (r *Request) X取上下文对象() context.Context {
	return r.Context别名()
}

// GetNeverDoneCtx 创建并返回一个永不完成的上下文对象，
// 允许在手动完成上下文之前，使得该上下文可以传播到异步goroutine中，
// 不受HTTP请求结束的影响。
// 
// 这个更改是考虑到开发人员在单个HTTP请求中创建多个goroutine时，常见的上下文传播习惯。
// md5:14245b7febbf75b4
func (r *Request) GetNeverDoneCtx() context.Context {
	return gctx.NeverDone(r.Context别名())
}

// X设置上下文对象 为当前请求设置自定义上下文。 md5:447efd465325d822
func (r *Request) X设置上下文对象(上下文 context.Context) {
	*r.Request = *r.WithContext(上下文)
}

// X取上下文对象值 根据给定的键名获取并返回一个 Var。
// 可选参数 `def` 指定了如果给定的 `key` 在上下文中不存在时，Var 的默认值。
// md5:8e874c6ac730ae7b
func (r *Request) X取上下文对象值(名称 interface{}, 默认值 ...interface{}) *gvar.Var {
	value := r.Context别名().Value(名称)
	if value == nil && len(默认值) > 0 {
		value = 默认值[0]
	}
	return gvar.X创建(value)
}

// X设置上下文对象值 为上下文（context）设置自定义参数，使用键值对形式。 md5:f6782a132f936ef1
func (r *Request) X设置上下文对象值(名称 interface{}, 值 interface{}) {
	var ctx = r.Context别名()
	ctx = context.WithValue(ctx, 名称, 值)
	*r.Request = *r.Request.WithContext(ctx)
}
