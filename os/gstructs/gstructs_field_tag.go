// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gstructs

import (
	"strings"

	"github.com/gogf/gf/v2/util/gtag"
)

// TagJsonName 返回字段的`json`标签名称字符串。 md5:5c012e1a7ddc21a4
func (f *Field) TagJsonName() string {
	if jsonTag := f.Tag(gtag.Json); jsonTag != "" {
		return strings.Split(jsonTag, ",")[0]
	}
	return ""
}

// TagDefault 返回字段最常用的标签 `default/d` 值。 md5:05a366c7274145a7
func (f *Field) TagDefault() string {
	v := f.Tag(gtag.Default)
	if v == "" {
		v = f.Tag(gtag.DefaultShort)
	}
	return v
}

// TagParam 返回字段最常见的标签 `param/p` 的值。 md5:4a9896d7eafb2571
func (f *Field) TagParam() string {
	v := f.Tag(gtag.Param)
	if v == "" {
		v = f.Tag(gtag.ParamShort)
	}
	return v
}

// TagValid 返回字段最常用的标签 `valid/v` 的值。 md5:94454fb9da424ab5
func (f *Field) TagValid() string {
	v := f.Tag(gtag.Valid)
	if v == "" {
		v = f.Tag(gtag.ValidShort)
	}
	return v
}

// TagDescription 返回字段最常用的标签 `description/des/dc` 的值。 md5:d9c639d81519a9e6
func (f *Field) TagDescription() string {
	v := f.Tag(gtag.Description)
	if v == "" {
		v = f.Tag(gtag.DescriptionShort)
	}
	if v == "" {
		v = f.Tag(gtag.DescriptionShort2)
	}
	return v
}

// TagSummary 返回字段最常使用的标签 `summary/sum/sm` 的值。 md5:73fda199beee8e03
func (f *Field) TagSummary() string {
	v := f.Tag(gtag.Summary)
	if v == "" {
		v = f.Tag(gtag.SummaryShort)
	}
	if v == "" {
		v = f.Tag(gtag.SummaryShort2)
	}
	return v
}

// TagAdditional 返回字段最常见的附加标签`additional/ad`的值。 md5:5fa6809b09e574f9
func (f *Field) TagAdditional() string {
	v := f.Tag(gtag.Additional)
	if v == "" {
		v = f.Tag(gtag.AdditionalShort)
	}
	return v
}

// TagExample 返回字段中最常用的标签 `example/eg` 的值。 md5:db6ca440d0e1c869
func (f *Field) TagExample() string {
	v := f.Tag(gtag.Example)
	if v == "" {
		v = f.Tag(gtag.ExampleShort)
	}
	return v
}

// TagIn 函数返回字段中最常用的标签`in`值。 md5:422dfe5b89aad3c4
func (f *Field) TagIn() string {
	v := f.Tag(gtag.In)
	return v
}

// TagPriorityName 检查并返回与`gtag.StructTagPriority`中的名称项匹配的标签名。
// 如果没有按照`gtag.StructTagPriority`设置标签名，它将返回属性字段的Name。 md5:5323769f60ec004e
func (f *Field) TagPriorityName() string {
	var name = f.Name()
	for _, tagName := range gtag.StructTagPriority {
		if tagValue := f.Tag(tagName); tagValue != "" {
			name = tagValue
			break
		}
	}
	return name
}
