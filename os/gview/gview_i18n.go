// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gview

import (
	"context"

	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	i18nLanguageVariableName = "I18nLanguage"
)

// i18nTranslate 使用i18n功能翻译内容。 md5:7e06981e1220b9b8
func (view *View) i18nTranslate(ctx context.Context, content string, variables Params) string {
	if view.config.I18nManager != nil {
				// 兼容旧版本。 md5:7b67a6d4fcad7e51
		if language, ok := variables[i18nLanguageVariableName]; ok {
			ctx = gi18n.WithLanguage(ctx, gconv.String(language))
		}
		return view.config.I18nManager.T(ctx, content)
	}
	return content
}

// setI18nLanguageFromCtx 从上下文（context）中获取语言名称，并将其设置到模板变量映射中。 md5:7f03916003c7fc5c
func (view *View) setI18nLanguageFromCtx(ctx context.Context, variables map[string]interface{}) {
	if _, ok := variables[i18nLanguageVariableName]; !ok {
		if language := gi18n.LanguageFromCtx(ctx); language != "" {
			variables[i18nLanguageVariableName] = language
		}
	}
}
