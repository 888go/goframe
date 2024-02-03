// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstructs

import (
	"reflect"
	
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/util/gtag"
)

// Tag 函数返回在标签字符串中与 key 关联的值。如果该标签中不存在相应的 key，则 Tag 函数返回空字符串。
func (f *Field) Tag(key string) string {
	s := f.Field.Tag.Get(key)
	if s != "" {
		s = gtag.Parse(s)
	}
	return s
}

// TagLookup在tag字符串中查找与key关联的值。
// 如果key在tag中存在，返回对应的值（可能为空）。
// 否则，返回值将是空字符串。
// ok返回值报告该值是否在tag字符串中显式设置。如果tag不具有常规格式，
// Lookup返回的值未指定。
func (f *Field) TagLookup(key string) (value string, ok bool) {
	value, ok = f.Field.Tag.Lookup(key)
	if ok && value != "" {
		value = gtag.Parse(value)
	}
	return
}

// IsEmbedded 返回 true 如果给定的字段是一个匿名字段（嵌入式）
func (f *Field) IsEmbedded() bool {
	return f.Field.Anonymous
}

// TagStr 返回该字段的标签字符串。
func (f *Field) TagStr() string {
	return string(f.Field.Tag)
}

// TagMap 返回该字段的所有标签及其对应的值字符串，以映射(map)形式表示。
func (f *Field) TagMap() map[string]string {
	var (
		data = ParseTag(f.TagStr())
	)
	for k, v := range data {
		data[k] = utils.StripSlashes(gtag.Parse(v))
	}
	return data
}

// IsExported 返回 true 如果给定的字段是导出的。
func (f *Field) IsExported() bool {
	return f.Field.PkgPath == ""
}

// Name 返回给定字段的名称。
func (f *Field) Name() string {
	return f.Field.Name
}

// Type 返回给定字段的类型。
// 注意，此处的 Type 并非 reflect.Type。如果你需要 reflect.Type，请使用 Field.Type().Type。
func (f *Field) Type() Type {
	return Type{
		Type: f.Field.Type,
	}
}

// Kind 返回Field `f`值的reflect.Kind。
func (f *Field) Kind() reflect.Kind {
	return f.Value.Kind()
}

// OriginalKind 从Field `f`的Value中获取并返回其原始的 reflect.Kind 类型。
func (f *Field) OriginalKind() reflect.Kind {
	var (
		reflectType = f.Value.Type()
		reflectKind = reflectType.Kind()
	)
	for reflectKind == reflect.Ptr {
		reflectType = reflectType.Elem()
		reflectKind = reflectType.Kind()
	}
	return reflectKind
}

// IsEmpty 检查并返回该 Field 的值是否为空。
func (f *Field) IsEmpty() bool {
	return empty.IsEmpty(f.Value)
}

// IsNil 检查并返回该 Field 的值是否为 nil。
func (f *Field) IsNil(traceSource ...bool) bool {
	return empty.IsNil(f.Value, traceSource...)
}
