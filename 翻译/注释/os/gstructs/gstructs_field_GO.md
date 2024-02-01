
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Tag returns the value associated with key in the tag string. If there is no
// such key in the tag, Tag returns the empty string.
<原文结束>

# <翻译开始>
// Tag 函数返回在标签字符串中与 key 关联的值。如果该标签中不存在相应的 key，则 Tag 函数返回空字符串。
# <翻译结束>


<原文开始>
// TagLookup returns the value associated with key in the tag string.
// If the key is present in the tag the value (which may be empty)
// is returned. Otherwise, the returned value will be the empty string.
// The ok return value reports whether the value was explicitly set in
// the tag string. If the tag does not have the conventional format,
// the value returned by Lookup is unspecified.
<原文结束>

# <翻译开始>
// TagLookup在tag字符串中查找与key关联的值。
// 如果key在tag中存在，返回对应的值（可能为空）。
// 否则，返回值将是空字符串。
// ok返回值报告该值是否在tag字符串中显式设置。如果tag不具有常规格式，
// Lookup返回的值未指定。
# <翻译结束>


<原文开始>
// IsEmbedded returns true if the given field is an anonymous field (embedded)
<原文结束>

# <翻译开始>
// IsEmbedded 返回 true 如果给定的字段是一个匿名字段（嵌入式）
# <翻译结束>


<原文开始>
// TagStr returns the tag string of the field.
<原文结束>

# <翻译开始>
// TagStr 返回该字段的标签字符串。
# <翻译结束>


<原文开始>
// TagMap returns all the tag of the field along with its value string as map.
<原文结束>

# <翻译开始>
// TagMap 返回该字段的所有标签及其对应的值字符串，以映射(map)形式表示。
# <翻译结束>


<原文开始>
// IsExported returns true if the given field is exported.
<原文结束>

# <翻译开始>
// IsExported 返回 true 如果给定的字段是导出的。
# <翻译结束>


<原文开始>
// Name returns the name of the given field.
<原文结束>

# <翻译开始>
// Name 返回给定字段的名称。
# <翻译结束>


<原文开始>
// Type returns the type of the given field.
// Note that this Type is not reflect.Type. If you need reflect.Type, please use Field.Type().Type.
<原文结束>

# <翻译开始>
// Type 返回给定字段的类型。
// 注意，此处的 Type 并非 reflect.Type。如果你需要 reflect.Type，请使用 Field.Type().Type。
# <翻译结束>


<原文开始>
// Kind returns the reflect.Kind for Value of Field `f`.
<原文结束>

# <翻译开始>
// Kind 返回Field `f`值的reflect.Kind。
# <翻译结束>


<原文开始>
// OriginalKind retrieves and returns the original reflect.Kind for Value of Field `f`.
<原文结束>

# <翻译开始>
// OriginalKind 从Field `f`的Value中获取并返回其原始的 reflect.Kind 类型。
# <翻译结束>


<原文开始>
// IsEmpty checks and returns whether the value of this Field is empty.
<原文结束>

# <翻译开始>
// IsEmpty 检查并返回该 Field 的值是否为空。
# <翻译结束>


<原文开始>
// IsNil checks and returns whether the value of this Field is nil.
<原文结束>

# <翻译开始>
// IsNil 检查并返回该 Field 的值是否为 nil。
# <翻译结束>

