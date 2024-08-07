// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包ghtml提供了处理HTML内容的有用API。 md5:218d4666a789e8d7
package html类

import (
	"html"
	"reflect"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
)

// X删除HTML标记 从内容中移除HTML标签，只返回纯文本。参考：http://php.net/manual/zh/function.strip-tags.php
// md5:a4340113401e2abd
func X删除HTML标记(html文本 string) string {
	return strip.StripTags(html文本)
}

// X编码 函数对内容中的所有HTML字符进行编码。
// 参考链接：http://php.net/manual/zh/function.htmlentities.php
// md5:5ef66bc1e4751683
func X编码(html文本 string) string {
	return html.EscapeString(html文本)
}

// X解码 解码内容中的所有HTML字符实体。
// 参考：http://php.net/manual/zh/function.html-entity-decode.php
// md5:0c31df1bae7c4cff
func X解码(文本 string) string {
	return html.UnescapeString(文本)
}

// X编码特殊字符 用于对内容进行编码，特别字符包括：
// "&", "<", ">", `"`, "'".
// 参考：http://php.net/manual/zh/function.htmlspecialchars.php
// md5:bd5792c08196b110
func X编码特殊字符(文本 string) string {
	return strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		`"`, "&#34;",
		"'", "&#39;",
	).Replace(文本)
}

// X解码特殊字符 用于解码内容中的一些特殊字符，这些特殊字符包括：
// "&", "<", ">", `"`, "'".
// 参考：http://php.net/manual/zh/function.htmlspecialchars-decode.php
// md5:839437a4221bd761
func X解码特殊字符(文本 string) string {
	return strings.NewReplacer(
		"&amp;", "&",
		"&lt;", "<",
		"&gt;", ">",
		"&#34;", `"`,
		"&#39;", "'",
	).Replace(文本)
}

// X编码Map 会自动对映射/结构体中的字符串值/属性进行编码。 md5:a26c73e35955b542
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
