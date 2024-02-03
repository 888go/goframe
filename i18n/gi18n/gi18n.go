// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gi18n implements internationalization and localization.
package gi18n

import (
	"context"
)

// SetPath 设置存储 i18n 文件的目录路径。
func SetPath(path string) error {
	return Instance().SetPath(path)
}

// SetLanguage 设置翻译器的语言。
func SetLanguage(language string) {
	Instance().SetLanguage(language)
}

// SetDelimiters 设置翻译器的分隔符。
func SetDelimiters(left, right string) {
	Instance().SetDelimiters(left, right)
}

// T 是 Translate 的别名，用于提供便利。
func T(ctx context.Context, content string) string {
	return Instance().T(ctx, content)
}

// Tf 是 TranslateFormat 的别名，用于提供便利。
func Tf(ctx context.Context, format string, values ...interface{}) string {
	return Instance().TranslateFormat(ctx, format, values...)
}

// TranslateFormat 将根据配置的语言和给定的 `values` 对 `format` 进行翻译、格式化并返回结果。
func TranslateFormat(ctx context.Context, format string, values ...interface{}) string {
	return Instance().TranslateFormat(ctx, format, values...)
}

// Translate 使用配置的语言对`content`进行翻译并返回翻译后的内容。
func Translate(ctx context.Context, content string) string {
	return Instance().Translate(ctx, content)
}

// GetContent 函数根据给定的键和指定的语言获取并返回配置的内容。
// 如果未找到，则返回一个空字符串。
func GetContent(ctx context.Context, key string) string {
	return Instance().GetContent(ctx, key)
}
