// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gview
import (
	"context"
	
	"github.com/888go/goframe/i18n/gi18n"
	"github.com/888go/goframe/util/gconv"
	)
const (
	i18nLanguageVariableName = "I18nLanguage"
)

// i18nTranslate 使用 i18n（国际化）功能翻译内容。
func (view *View) i18nTranslate(ctx context.Context, content string, variables Params) string {
	if view.config.I18nManager != nil {
		// 与旧版本兼容。
		if language, ok := variables[i18nLanguageVariableName]; ok {
			ctx = gi18n.WithLanguage(ctx, gconv.String(language))
		}
		return view.config.I18nManager.T(ctx, content)
	}
	return content
}

// setI18nLanguageFromCtx 从上下文中检索语言名称，并将其设置到模板变量映射中。
func (view *View) setI18nLanguageFromCtx(ctx context.Context, variables map[string]interface{}) {
	if _, ok := variables[i18nLanguageVariableName]; !ok {
		if language := gi18n.LanguageFromCtx(ctx); language != "" {
			variables[i18nLanguageVariableName] = language
		}
	}
}
