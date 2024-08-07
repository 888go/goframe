// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtag

import (
	"regexp"

	gerror "github.com/888go/goframe/errors/gerror"
)

var (
	data  = make(map[string]string)
	regex = regexp.MustCompile(`\{(.+?)\}`)
)

// X设置值 设置指定名称的标签内容。
// 请注意，如果`name`已经存在，该函数会引发恐慌。
// md5:3b301b4174b60616
func X设置值(name, value string) {
	if _, ok := data[name]; ok {
		panic(gerror.X创建并格式化(`value for tag name "%s" already exists`, name))
	}
	data[name] = value
}

// SetOver 执行 Set 的功能，但如果 `name` 已经存在，它会覆盖旧的值。 md5:906ca9f516be44d0
func SetOver(name, value string) {
	data[name] = value
}

// 通过map设置多个标签的内容。 md5:c02ae9dd9350cf50
func Sets(m map[string]string) {
	for k, v := range m {
		X设置值(k, v)
	}
}

// SetsOver 的行为类似于 Sets，但如果 `name` 已经存在，它会覆盖旧值。 md5:6a87c6587ed9794f
func SetsOver(m map[string]string) {
	for k, v := range m {
		SetOver(k, v)
	}
}

// Get 获取并返回指定名称的存储标签内容。 md5:1a0a007cb18c41fa
func Get(name string) string {
	return data[name]
}

// Parse 通过将所有标签名变量替换为其内容，解析并返回给定的`content`。
// 示例：
// gtag.Set("demo", "content")
// Parse(`This is {demo}`) -> `This is content`。
// md5:b45c5273962c7662
func Parse(content string) string {
	return regex.ReplaceAllStringFunc(content, func(s string) string {
		if v, ok := data[s[1:len(s)-1]]; ok {
			return v
		}
		return s
	})
}
