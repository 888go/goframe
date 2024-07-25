// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstructs

import (
	"reflect"

	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/util/gtag"
)

// Tag 函数从标签字符串中返回与给定键关联的值。如果标签中没有该键，Tag 函数将返回空字符串。
// md5:1f7397ec7f558f60
func (f *Field) Tag(key string) string {
	s := f.Field.Tag.Get(key)
	if s != "" {
		s = gtag.Parse(s)
	}
	return s
}

// TagLookup 从标签字符串中返回与给定键关联的值。如果键在标签中存在，即使值为空，也会返回。否则，返回的值将是空字符串。ok返回值报告了该值是否明确设置在标签字符串中。如果标签不具备常规格式，Lookup返回的值是未定义的。
// md5:d4bff95e89bd22d0
func (f *Field) TagLookup(key string) (value string, ok bool) {
	value, ok = f.Field.Tag.Lookup(key)
	if ok && value != "" {
		value = gtag.Parse(value)
	}
	return
}

// IsEmbedded 如果给定的字段是一个匿名字段（嵌入式），则返回true. md5:db717a9b06b1f0f5
func (f *Field) IsEmbedded() bool {
	return f.Field.Anonymous
}

// TagStr 返回字段的标签字符串。 md5:d608cb4dcc85989d
func (f *Field) TagStr() string {
	return string(f.Field.Tag)
}

// TagMap 返回字段的所有标签及其对应的值字符串作为映射。 md5:80b1670604d9eef4
func (f *Field) TagMap() map[string]string {
	var (
		data = ParseTag(f.TagStr())
	)
	for k, v := range data {
		data[k] = utils.StripSlashes(gtag.Parse(v))
	}
	return data
}

// IsExported 返回给定字段是否被导出。 md5:b863b7d714c969fc
func (f *Field) IsExported() bool {
	return f.Field.PkgPath == ""
}

// Name 返回给定字段的名称。 md5:bfd1563575d622f5
func (f *Field) Name() string {
	return f.Field.Name
}

// Type 返回给定字段的类型。
// 请注意，此Type不是reflect.Type。如果需要reflect.Type，请使用Field.Type().Type。
// md5:27a135d33cbd8f21
func (f *Field) Type() Type {
	return Type{
		Type: f.Field.Type,
	}
}

// Kind返回Field `f`的Value对应的reflect.Kind。 md5:6c3599f3dff91746
func (f *Field) Kind() reflect.Kind {
	return f.Value.Kind()
}

// OriginalKind 获取并返回字段 `f` 的Value对应的原始reflect.Kind。 md5:62d8a3604e2114ec
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

// OriginalValue 获取并返回字段`f`的原始reflect.Value。 md5:0f37794c6e9ea990
func (f *Field) OriginalValue() reflect.Value {
	var (
		reflectValue = f.Value
		reflectType  = reflectValue.Type()
		reflectKind  = reflectType.Kind()
	)

	for reflectKind == reflect.Ptr && !f.IsNil() {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Type().Kind()
	}

	return reflectValue
}

// IsEmpty 检查并返回这个字段的值是否为空。 md5:125094bfbb4cc317
func (f *Field) IsEmpty() bool {
	return empty.IsEmpty(f.Value)
}

// IsNil 检查并返回此Field的值是否为nil。 md5:6637754b5d35923d
func (f *Field) IsNil(traceSource ...bool) bool {
	return empty.IsNil(f.Value, traceSource...)
}
