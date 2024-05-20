// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package gi18n 实现国际化和本地化。. md5:c7d7e3f7580f80a2
package gi18n

import (
	"context"

	"github.com/gogf/gf/v2/os/gctx"
)

const (
	ctxLanguage gctx.StrKey = "I18nLanguage"
)

// WithLanguage 向上下文中追加语言设置并返回一个新的上下文。. md5:c09b915c27b37312
func WithLanguage(ctx context.Context, language string) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	return context.WithValue(ctx, ctxLanguage, language)
}

// LanguageFromCtx 从上下文中获取并返回语言名称。
// 如果之前未设置，则返回空字符串。
// md5:f62999632f76669e
func LanguageFromCtx(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	v := ctx.Value(ctxLanguage)
	if v != nil {
		return v.(string)
	}
	return ""
}
