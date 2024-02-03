// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/os/gctx"
)

// RequestFromCtx 从 context 中检索并返回 Request 对象。
func RequestFromCtx(ctx context.Context) *Request {
	if v := ctx.Value(ctxKeyForRequest); v != nil {
		return v.(*Request)
	}
	return nil
}

// Context 是 GetCtx 函数的别名。
// 此函数覆盖了 http.Request.Context 函数的功能。
// 请参阅 GetCtx。
func (r *Request) Context() context.Context {
	var ctx = r.Request.Context()
	// 检查并把Request对象注入到上下文中。
	if RequestFromCtx(ctx) == nil {
		// 将Request对象注入到context中。
		ctx = context.WithValue(ctx, ctxKeyForRequest, r)
		// 如果使用默认追踪提供者，则添加默认追踪信息。
		ctx = gctx.WithCtx(ctx)
		// 更新原始HTTP请求的值。
		*r.Request = *r.Request.WithContext(ctx)
	}
	return ctx
}

// GetCtx 从请求中检索并返回其上下文。
// 它是函数 Context 的别名，为了与函数 SetCtx 保持关联性。
func (r *Request) GetCtx() context.Context {
	return r.Context()
}

// GetNeverDoneCtx 创建并返回一个永不完成的上下文对象，
// 该对象禁止手动设置完成状态，以保证上下文可以传递到异步goroutine中，
// 因此不会受到HTTP请求结束的影响。
//
// 这个改动是为了适应开发者在处理单个HTTP请求时创建多个goroutine进行上下文传播的常见使用习惯。
func (r *Request) GetNeverDoneCtx() context.Context {
	return gctx.NeverDone(r.Context())
}

// SetCtx 为当前请求设置自定义上下文。
func (r *Request) SetCtx(ctx context.Context) {
	*r.Request = *r.WithContext(ctx)
}

// GetCtxVar 通过给定的键名检索并返回一个 Var。
// 可选参数 `def` 指定了如果给定 `key` 在上下文中不存在时，Var 的默认值。
func (r *Request) GetCtxVar(key interface{}, def ...interface{}) *gvar.Var {
	value := r.Context().Value(key)
	if value == nil && len(def) > 0 {
		value = def[0]
	}
	return gvar.New(value)
}

// SetCtxVar 通过键值对设置自定义参数到上下文中。
func (r *Request) SetCtxVar(key interface{}, value interface{}) {
	var ctx = r.Context()
	ctx = context.WithValue(ctx, key, value)
	*r.Request = *r.Request.WithContext(ctx)
}
