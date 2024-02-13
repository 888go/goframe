// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package ghtml 提供了用于处理 HTML 内容的有用 API。
package html类

import (
	"html"
	"reflect"
	"strings"
	
	strip "github.com/grokify/html-strip-tags-go"
)

// StripTags 从内容中剥离 HTML 标签，仅返回文本内容。
// 参考来源：http://php.net/manual/zh/function.strip-tags.php
func X删除HTML标记(html文本 string) string {
	return strip.StripTags(html文本)
}

// Entities 将内容中的所有HTML字符进行编码。
// 参考来源: http://php.net/manual/zh/function.htmlentities.php
func X编码(html文本 string) string {
	return html.EscapeString(html文本)
}

// EntitiesDecode 将内容中的所有HTML字符进行解码。
// 参考：http://php.net/manual/zh/function.html-entity-decode.php
func X解码(文本 string) string {
	return html.UnescapeString(文本)
}

// SpecialChars 为内容编码一些特殊字符，这些特殊字符包括：
// "&", "<", ">", `"`, "'".
// 参考：http://php.net/manual/zh/function.htmlspecialchars.php
func X编码特殊字符(文本 string) string {
	return strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		`"`, "&#34;",
		"'", "&#39;",
	).Replace(文本)
}

// SpecialCharsDecode 对内容中的某些特殊字符进行解码，这些特殊字符包括：
// "&"（和号）、"<"（小于号）、">"（大于号）、`"`（双引号）、"'"（单引号）。
// 参考文档：http://php.net/manual/zh/function.htmlspecialchars-decode.php
func X解码特殊字符(文本 string) string {
	return strings.NewReplacer(
		"&amp;", "&",
		"&lt;", "<",
		"&gt;", ">",
		"&#34;", `"`,
		"&#39;", "'",
	).Replace(文本)
}

// SpecialCharsMapOrStruct 自动对 map 或 struct 中的字符串值/属性进行编码。
func X编码Map(map或Struct interface{}) error {
	var (
		reflectValue = reflect.ValueOf(map或Struct)
		reflectKind  = reflectValue.Kind()
	)
	for reflectValue.IsValid() && (reflectKind == reflect.Ptr || reflectKind == reflect.Interface) {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Map:
		var (
			mapKeys  = reflectValue.MapKeys()
			mapValue reflect.Value
		)
		for _, key := range mapKeys {
			mapValue = reflectValue.MapIndex(key)
			switch mapValue.Kind() {
			case reflect.String:
				reflectValue.SetMapIndex(key, reflect.ValueOf(X编码特殊字符(mapValue.String())))
			case reflect.Interface:
				if mapValue.Elem().Kind() == reflect.String {
					reflectValue.SetMapIndex(key, reflect.ValueOf(X编码特殊字符(mapValue.Elem().String())))
				}
			}
		}

	case reflect.Struct:
		var (
			fieldValue reflect.Value
		)
		for i := 0; i < reflectValue.NumField(); i++ {
			fieldValue = reflectValue.Field(i)
			switch fieldValue.Kind() {
			case reflect.String:
				fieldValue.Set(reflect.ValueOf(X编码特殊字符(fieldValue.String())))
			}
		}
	}
	return nil
}
