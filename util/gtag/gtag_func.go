// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtag
import (
	"regexp"
	
	"github.com/888go/goframe/errors/gerror"
	)
var (
	data  = make(map[string]string)
	regex = regexp.MustCompile(`\{(.+?)\}`)
)

// Set为指定名称设置标签内容。
// 注意，如果`name`已存在，则会引发panic。
func Set(name, value string) {
	if _, ok := data[name]; ok {
		panic(gerror.Newf(`value for tag name "%s" already exists`, name))
	}
	data[name] = value
}

// SetOver 函数表现如同 Set，但当 `name` 已经存在时，它会覆盖旧的值。
func SetOver(name, value string) {
	data[name] = value
}

// Sets 通过映射设置多个标签内容。
func Sets(m map[string]string) {
	for k, v := range m {
		Set(k, v)
	}
}

// SetsOver 函数表现与 Sets 相同，但当 `name` 已经存在时，它会覆盖旧的值。
func SetsOver(m map[string]string) {
	for k, v := range m {
		SetOver(k, v)
	}
}

// Get 方法用于根据指定名称检索并返回存储的标签内容。
func Get(name string) string {
	return data[name]
}

// Parse函数解析并返回内容，将给定`content`中所有标签名称变量替换为它的实际内容。
// 示例：
// gtag.Set("demo", "content")
// Parse(`This is {demo}`) -> `This is content`。
func Parse(content string) string {
	return regex.ReplaceAllStringFunc(content, func(s string) string {
		if v, ok := data[s[1:len(s)-1]]; ok {
			return v
		}
		return s
	})
}
