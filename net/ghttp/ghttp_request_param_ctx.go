// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/os/gctx"
)

// RequestFromCtx 从 context 中检索并返回 Request 对象。
func X从上下文取请求对象(上下文 context.Context) *Request {
	if v := 上下文.Value(ctxKeyForRequest); v != nil {
		return v.(*Request)
	}
	return nil
}

// Context 是 GetCtx 函数的别名。
// 此函数覆盖了 http.Request.Context 函数的功能。
// 请参阅 GetCtx。
func (r *Request) Context别名() context.Context {
	var ctx = r.Request.Context()
	// 检查并把Request对象注入到上下文中。
	if X从上下文取请求对象(ctx) == nil {
		// 将Request对象注入到context中。
		ctx = context.WithValue(ctx, ctxKeyForRequest, r)
		// 如果使用默认追踪提供者，则添加默认追踪信息。
		ctx = 上下文类.X创建并从上下文(ctx)
		// 更新原始HTTP请求的值。
		*r.Request = *r.Request.WithContext(ctx)
	}
	return ctx
}

// GetCtx 从请求中检索并返回其上下文。
// 它是函数 Context 的别名，为了与函数 SetCtx 保持关联性。
func (r *Request) X取上下文对象() context.Context {
	return r.Context别名()
}

// GetNeverDoneCtx 创建并返回一个永不完成的上下文对象，
// 该对象禁止手动设置完成状态，以保证上下文可以传递到异步goroutine中，
// 因此不会受到HTTP请求结束的影响。
//
// 这个改动是为了适应开发者在处理单个HTTP请求时创建多个goroutine进行上下文传播的常见使用习惯。
func (r *Request) GetNeverDoneCtx() context.Context {
	return 上下文类.NeverDone(r.Context别名())
}

// SetCtx 为当前请求设置自定义上下文。
func (r *Request) X设置上下文对象(上下文 context.Context) {
	*r.Request = *r.WithContext(上下文)
}

// GetCtxVar 通过给定的键名检索并返回一个 Var。
// 可选参数 `def` 指定了如果给定 `key` 在上下文中不存在时，Var 的默认值。
func (r *Request) X取上下文对象值(名称 interface{}, 默认值 ...interface{}) *泛型类.Var {
	value := r.Context别名().Value(名称)
	if value == nil && len(默认值) > 0 {
		value = 默认值[0]
	}
	return 泛型类.X创建(value)
}

// SetCtxVar 通过键值对设置自定义参数到上下文中。
func (r *Request) X设置上下文对象值(名称 interface{}, 值 interface{}) {
	var ctx = r.Context别名()
	ctx = context.WithValue(ctx, 名称, 值)
	*r.Request = *r.Request.WithContext(ctx)
}
