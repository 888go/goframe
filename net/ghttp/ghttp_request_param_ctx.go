// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gctx"
)

// RequestFromCtx 从上下文中检索并返回Request对象。 md5:c247eac3d031fb2b
func RequestFromCtx(ctx context.Context) *Request {
	if v := ctx.Value(ctxKeyForRequest); v != nil {
		return v.(*Request)
	}
	return nil
}

// Context 是函数 GetCtx 的别名。
// 这个函数重写了 http.Request.Context 函数。
// 参考 GetCtx 函数。
// md5:cc1d0847fad835ae
func (r *Request) Context() context.Context {
	var ctx = r.Request.Context()
		// 检查并把Request对象注入到context中。 md5:6db244059baa2f77
	if RequestFromCtx(ctx) == nil {
				// 将 Request 对象注入到上下文中。 md5:561b491dd0877f18
		ctx = context.WithValue(ctx, ctxKeyForRequest, r)
				// 如果使用默认跟踪提供程序，添加默认跟踪信息。 md5:ffe39a076b5cfe89
		ctx = gctx.WithCtx(ctx)
				// 更新原始HTTP请求的值。 md5:b7f6f8f267437800
		*r.Request = *r.Request.WithContext(ctx)
	}
	return ctx
}

// GetCtx 获取并返回请求的上下文。
// 它是函数Context的别名，与SetCtx函数相关联。
// md5:67b84a74829ae7e1
func (r *Request) GetCtx() context.Context {
	return r.Context()
}

// GetNeverDoneCtx 创建并返回一个永不完成的上下文对象，
// 允许在手动完成上下文之前，使得该上下文可以传播到异步goroutine中，
// 不受HTTP请求结束的影响。
// 
// 这个更改是考虑到开发人员在单个HTTP请求中创建多个goroutine时，常见的上下文传播习惯。
// md5:14245b7febbf75b4
func (r *Request) GetNeverDoneCtx() context.Context {
	return gctx.NeverDone(r.Context())
}

// SetCtx 为当前请求设置自定义上下文。 md5:447efd465325d822
func (r *Request) SetCtx(ctx context.Context) {
	*r.Request = *r.WithContext(ctx)
}

// GetCtxVar 根据给定的键名获取并返回一个 Var。
// 可选参数 `def` 指定了如果给定的 `key` 在上下文中不存在时，Var 的默认值。
// md5:8e874c6ac730ae7b
func (r *Request) GetCtxVar(key interface{}, def ...interface{}) *gvar.Var {
	value := r.Context().Value(key)
	if value == nil && len(def) > 0 {
		value = def[0]
	}
	return gvar.New(value)
}

// SetCtxVar 为上下文（context）设置自定义参数，使用键值对形式。 md5:f6782a132f936ef1
func (r *Request) SetCtxVar(key interface{}, value interface{}) {
	var ctx = r.Context()
	ctx = context.WithValue(ctx, key, value)
	*r.Request = *r.Request.WithContext(ctx)
}
