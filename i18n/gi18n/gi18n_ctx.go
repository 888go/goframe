// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gi18n 实现了国际化和本地化功能。
package gi18n
import (
	"context"
	
	"github.com/888go/goframe/os/gctx"
	)
const (
	ctxLanguage gctx.StrKey = "I18nLanguage"
)

// WithLanguage 将语言设置追加到上下文中并返回一个新的上下文。
func WithLanguage(ctx context.Context, language string) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	return context.WithValue(ctx, ctxLanguage, language)
}

// LanguageFromCtx 从context中获取并返回语言名称。
// 如果此前未设置，则返回一个空字符串。
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
