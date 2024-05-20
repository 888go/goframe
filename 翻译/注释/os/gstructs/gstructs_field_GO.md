
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Tag returns the value associated with key in the tag string. If there is no
// such key in the tag, Tag returns the empty string.
<原文结束>

# <翻译开始>
// Tag 函数从标签字符串中返回与给定键关联的值。如果标签中没有该键，Tag 函数将返回空字符串。
// md5:1f7397ec7f558f60
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
// TagLookup 从标签字符串中返回与给定键关联的值。如果键在标签中存在，即使值为空，也会返回。否则，返回的值将是空字符串。ok返回值报告了该值是否明确设置在标签字符串中。如果标签不具备常规格式，Lookup返回的值是未定义的。
// md5:d4bff95e89bd22d0
# <翻译结束>


<原文开始>
// IsEmbedded returns true if the given field is an anonymous field (embedded)
<原文结束>

# <翻译开始>
// IsEmbedded 如果给定的字段是一个匿名字段（嵌入式），则返回true. md5:db717a9b06b1f0f5
# <翻译结束>


<原文开始>
// TagStr returns the tag string of the field.
<原文结束>

# <翻译开始>
// TagStr 返回字段的标签字符串。. md5:d608cb4dcc85989d
# <翻译结束>


<原文开始>
// TagMap returns all the tag of the field along with its value string as map.
<原文结束>

# <翻译开始>
// TagMap 返回字段的所有标签及其对应的值字符串作为映射。. md5:80b1670604d9eef4
# <翻译结束>


<原文开始>
// IsExported returns true if the given field is exported.
<原文结束>

# <翻译开始>
// IsExported 返回给定字段是否被导出。. md5:b863b7d714c969fc
# <翻译结束>


<原文开始>
// Name returns the name of the given field.
<原文结束>

# <翻译开始>
// Name 返回给定字段的名称。. md5:bfd1563575d622f5
# <翻译结束>


<原文开始>
// Type returns the type of the given field.
// Note that this Type is not reflect.Type. If you need reflect.Type, please use Field.Type().Type.
<原文结束>

# <翻译开始>
// Type 返回给定字段的类型。
// 请注意，此Type不是reflect.Type。如果需要reflect.Type，请使用Field.Type().Type。
// md5:27a135d33cbd8f21
# <翻译结束>


<原文开始>
// Kind returns the reflect.Kind for Value of Field `f`.
<原文结束>

# <翻译开始>
// Kind返回Field `f`的Value对应的reflect.Kind。. md5:6c3599f3dff91746
# <翻译结束>


<原文开始>
// OriginalKind retrieves and returns the original reflect.Kind for Value of Field `f`.
<原文结束>

# <翻译开始>
// OriginalKind 获取并返回字段 `f` 的Value对应的原始reflect.Kind。. md5:62d8a3604e2114ec
# <翻译结束>


<原文开始>
// OriginalValue retrieves and returns the original reflect.Value of Field `f`.
<原文结束>

# <翻译开始>
// OriginalValue 获取并返回字段`f`的原始reflect.Value。. md5:0f37794c6e9ea990
# <翻译结束>


<原文开始>
// IsEmpty checks and returns whether the value of this Field is empty.
<原文结束>

# <翻译开始>
// IsEmpty 检查并返回这个字段的值是否为空。. md5:125094bfbb4cc317
# <翻译结束>


<原文开始>
// IsNil checks and returns whether the value of this Field is nil.
<原文结束>

# <翻译开始>
// IsNil 检查并返回此Field的值是否为nil。. md5:6637754b5d35923d
# <翻译结束>

