
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
// ParseTag parses tag string into map.
// For example:
// ParseTag(`v:"required" p:"id" d:"1"`) => map[v:required p:id d:1].
<原文结束>

# <翻译开始>
// ParseTag函数用于解析标签字符串并转换为映射（map）。
// 例如：
// ParseTag(`v:"required" p:"id" d:"1"`) => map[v:required p:id d:1].
# <翻译结束>


<原文开始>
		// Scan to colon. A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
<原文结束>

# <翻译开始>
// 扫描到冒号。空格、引号或控制字符都是语法错误。
// 严格来讲，控制字符包括范围 [0x7f, 0x9f]，而不只是 [0x00, 0x1f]，
// 但在实际操作中，我们忽略了多字节的控制字符，
// 因为检查标签字节比检查标签符文更为简单。
# <翻译结束>


<原文开始>
// Scan quoted string to find value.
<原文结束>

# <翻译开始>
// 扫描带引号的字符串以查找值。
# <翻译结束>


<原文开始>
// TagFields retrieves and returns struct tags as []Field from `pointer`.
//
// The parameter `pointer` should be type of struct/*struct.
//
// Note that,
// 1. It only retrieves the exported attributes with first letter upper-case from struct.
// 2. The parameter `priority` should be given, it only retrieves fields that has given tag.
<原文结束>

# <翻译开始>
// TagFields 从`pointer`中获取并返回以[]Field形式的结构体标签。
//
// 参数`pointer`应为struct或*struct类型。
//
// 注意：
// 1. 它仅从结构体中获取首字母大写的导出属性（即公开字段）。
// 2. 参数`priority`应提供，它只检索具有给定标签的字段。
# <翻译结束>


<原文开始>
// TagMapName retrieves and returns struct tags as map[tag]attribute from `pointer`.
//
// The parameter `pointer` should be type of struct/*struct.
//
// Note that,
// 1. It only retrieves the exported attributes with first letter upper-case from struct.
// 2. The parameter `priority` should be given, it only retrieves fields that has given tag.
// 3. If one field has no specified tag, it uses its field name as result map key.
<原文结束>

# <翻译开始>
// TagMapName 从`pointer`中获取并返回以map[tag]attribute形式的结构体标签。
//
// 参数`pointer`应为struct或*struct类型。
//
// 注意：
// 1. 它只从结构体中获取首字母大写的导出属性。
// 2. 应提供参数`priority`，它只检索具有给定标签的字段。
// 3. 如果某个字段没有指定标签，则使用其字段名作为结果映射键。
# <翻译结束>


<原文开始>
// TagMapField retrieves struct tags as map[tag]Field from `pointer`, and returns it.
// The parameter `object` should be either type of struct/*struct/[]struct/[]*struct.
//
// Note that,
// 1. It only retrieves the exported attributes with first letter upper-case from struct.
// 2. The parameter `priority` should be given, it only retrieves fields that has given tag.
// 3. If one field has no specified tag, it uses its field name as result map key.
<原文结束>

# <翻译开始>
// TagMapField 从`pointer`中获取结构体标签并以map[tag]Field的形式返回。参数`object`应为struct/*struct/[]struct/[]*struct类型。
//
// 注意：
// 1. 它只检索结构体中首字母大写的导出属性（即公开字段）。
// 2. 参数`priority`必须给出，它只检索具有该给定标签的字段。
// 3. 如果某个字段没有指定标签，则使用其字段名称作为结果映射键。
# <翻译结束>


<原文开始>
// If pointer is type of *struct and nil, then automatically create a temporary struct.
<原文结束>

# <翻译开始>
// 如果指针是结构体类型且为nil，则自动创建一个临时结构体。
# <翻译结束>


<原文开始>
// Only retrieve exported attributes.
<原文结束>

# <翻译开始>
// 仅获取导出的属性。
# <翻译结束>


<原文开始>
// If this is an embedded attribute, it retrieves the tags recursively.
<原文结束>

# <翻译开始>
// 如果这是一个嵌入式属性，它会递归地获取标签。
# <翻译结束>


<原文开始>
// Skip leading space.
<原文结束>

# <翻译开始>
// 跳过前面的空格。
# <翻译结束>


<原文开始>
// Filter repeated tag.
<原文结束>

# <翻译开始>
// 过滤重复的标签。
# <翻译结束>

