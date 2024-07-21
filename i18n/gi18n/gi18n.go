// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package gi18n 实现国际化和本地化。 md5:c7d7e3f7580f80a2
package gi18n

import "context"

// SetPath 设置存储i18n文件的目录路径。 md5:b39e1d244949dcf8
// ff:
// path:
func SetPath(path string) error {
	return Instance().SetPath(path)
}

// SetLanguage 设置翻译器的语言。 md5:50b09b0bb0944dc1
// ff:
// language:
func SetLanguage(language string) {
	Instance().SetLanguage(language)
}

// SetDelimiters 为翻译器设置分隔符。 md5:f84b046b11204dc7
// ff:
// left:
// right:
func SetDelimiters(left, right string) {
	Instance().SetDelimiters(left, right)
}

// T 是为了方便而对 Translate 的别名。 md5:c07a6fa99a429eb3
// ff:
// ctx:
// content:
func T(ctx context.Context, content string) string {
	return Instance().T(ctx, content)
}

// Tf是TranslateFormat的别名，为了方便起见。 md5:bdb209b24c669f5a
// ff:
// ctx:
// format:
// values:
func Tf(ctx context.Context, format string, values ...interface{}) string {
	return Instance().TranslateFormat(ctx, format, values...)
}

// TranslateFormat 使用配置的语言和给定的 `values` 对 `format` 进行翻译、格式化并返回结果。
// md5:2806a81d6db86c7f
// ff:
// ctx:
// format:
// values:
func TranslateFormat(ctx context.Context, format string, values ...interface{}) string {
	return Instance().TranslateFormat(ctx, format, values...)
}

// Translate 使用配置的语言翻译`content`并返回翻译后的内容。 md5:a39ef8b5e28189db
// ff:
// ctx:
// content:
func Translate(ctx context.Context, content string) string {
	return Instance().Translate(ctx, content)
}

		// GetContent 获取并返回给定键和指定语言的配置内容。
		// 如果未找到，将返回一个空字符串。
		// md5:c64a3a803ac07e38
// ff:
// ctx:
// key:
func GetContent(ctx context.Context, key string) string {
	return Instance().GetContent(ctx, key)
}
