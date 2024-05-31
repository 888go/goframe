// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gi18n implements internationalization and localization.
package gi18n

import "context"

// SetPath sets the directory path storing i18n files.

// ff:
// path:
func SetPath(path string) error {
	return Instance().SetPath(path)
}

// SetLanguage sets the language for translator.

// ff:
// language:
func SetLanguage(language string) {
	Instance().SetLanguage(language)
}

// SetDelimiters sets the delimiters for translator.

// ff:
// right:
// left:
func SetDelimiters(left, right string) {
	Instance().SetDelimiters(left, right)
}

// T is alias of Translate for convenience.

// ff:
// content:
// ctx:
func T(ctx context.Context, content string) string {
	return Instance().T(ctx, content)
}

// Tf is alias of TranslateFormat for convenience.

// ff:
// values:
// format:
// ctx:
func Tf(ctx context.Context, format string, values ...interface{}) string {
	return Instance().TranslateFormat(ctx, format, values...)
}

// TranslateFormat translates, formats and returns the `format` with configured language
// and given `values`.

// ff:
// values:
// format:
// ctx:
func TranslateFormat(ctx context.Context, format string, values ...interface{}) string {
	return Instance().TranslateFormat(ctx, format, values...)
}

// Translate translates `content` with configured language and returns the translated content.

// ff:
// content:
// ctx:
func Translate(ctx context.Context, content string) string {
	return Instance().Translate(ctx, content)
}

// GetContent retrieves and returns the configured content for given key and specified language.
// It returns an empty string if not found.

// ff:
// key:
// ctx:
func GetContent(ctx context.Context, key string) string {
	return Instance().GetContent(ctx, key)
}
