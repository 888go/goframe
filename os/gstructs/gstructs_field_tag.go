// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstructs
import (
	"strings"
	
	"github.com/888go/goframe/util/gtag"
	)
// TagJsonName 返回该字段的 `json` 标签名称字符串。
func (f *Field) TagJsonName() string {
	if jsonTag := f.Tag(gtag.Json); jsonTag != "" {
		return strings.Split(jsonTag, ",")[0]
	}
	return ""
}

// TagDefault 返回该字段最常用的标签 `default/d` 的值。
func (f *Field) TagDefault() string {
	v := f.Tag(gtag.Default)
	if v == "" {
		v = f.Tag(gtag.DefaultShort)
	}
	return v
}

// TagParam 返回该字段最常用的标签 `param/p` 的值。
func (f *Field) TagParam() string {
	v := f.Tag(gtag.Param)
	if v == "" {
		v = f.Tag(gtag.ParamShort)
	}
	return v
}

// TagValid 返回该字段最常用的标签 `valid/v` 的值。
func (f *Field) TagValid() string {
	v := f.Tag(gtag.Valid)
	if v == "" {
		v = f.Tag(gtag.ValidShort)
	}
	return v
}

// TagDescription 返回字段中最常用的标签 `description/des/dc` 的值。
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

// TagSummary 返回该字段使用最频繁的标签值，标签键为 `summary`、`sum` 或 `sm`。
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

// TagAdditional 返回该字段最常用的标签 "additional/ad" 的值。
func (f *Field) TagAdditional() string {
	v := f.Tag(gtag.Additional)
	if v == "" {
		v = f.Tag(gtag.AdditionalShort)
	}
	return v
}

// TagExample 返回该字段最常用的标签 `example/eg` 的值。
func (f *Field) TagExample() string {
	v := f.Tag(gtag.Example)
	if v == "" {
		v = f.Tag(gtag.ExampleShort)
	}
	return v
}

// TagIn 返回该字段最常用的标签 "in" 值。
func (f *Field) TagIn() string {
	v := f.Tag(gtag.In)
	return v
}
