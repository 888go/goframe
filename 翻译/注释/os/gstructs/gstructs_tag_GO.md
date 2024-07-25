
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
// ParseTag parses tag string into map.
// For example:
// ParseTag(`v:"required" p:"id" d:"1"`) => map[v:required p:id d:1].
<原文结束>

# <翻译开始>
// ParseTag 将标签字符串解析为映射。
// 例如：
// ParseTag(`v:"required" p:"id" d:"1"`)) => map[v:required p:id d:1]。
// md5:967d381052c3a2d8
# <翻译结束>


<原文开始>
		// Scan to colon. A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
<原文结束>

# <翻译开始>
		// 扫描到冒号。空格、引号或控制字符都是语法错误。
		// 严格来说，控制字符包括范围 [0x7f, 0x9f]，而不仅仅是 [0x00, 0x1f]。但在实践中，我们忽略多字节控制字符，因为检查标签的字节比检查标签的 rune 更简单。
		// md5:2b37f6b6cf4e8415
# <翻译结束>


<原文开始>
// Scan quoted string to find value.
<原文结束>

# <翻译开始>
		// 扫描带引号的字符串以找到值。 md5:022e03f120cb2054
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
// TagFields 从`pointer`获取并返回结构体标签作为[]Field。
//
// 参数`pointer`应为struct/*struct类型。
//
// 请注意：
// 1. 它只从结构体中检索首字母大写的导出属性。
// 2. 应提供参数`priority`，它只检索具有给定标签的字段。
// md5:55390bfc1f5537f2
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
// TagMapName从`pointer`获取并返回结构体标签作为map[tag]attribute。
// 
// 参数`pointer`应为结构体或*struct类型。
// 
// 注意：
// 1. 它仅从结构体中检索首字母大写的导出属性。
// 2. 需要提供参数`priority`，它只检索具有给定标签的字段。
// 3. 如果一个字段没有指定标签，它将使用其字段名称作为结果映射的键。
// md5:0eb7c62c8a6f7e09
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
// TagMapField 从 `pointer` 中获取结构体标签作为 map[tag]Field，然后返回它。
// 参数 `object` 应该是 struct 类型、*struct 类型、struct 切片或 []*struct 类型之一。
// 
// 注意：
// 1. 它只会从结构体中检索首字母大写的导出属性。
// 2. 需要提供参数 `priority`，只检索具有给定标签的字段。
// 3. 如果一个字段没有指定标签，它将使用其字段名称作为结果映射的键。
// md5:ba865b4214b27332
# <翻译结束>


<原文开始>
// If pointer is type of *struct and nil, then automatically create a temporary struct.
<原文结束>

# <翻译开始>
				// 如果指针是*struct类型且为nil，那么会自动创建一个临时的struct。 md5:23b5ebc131739e7d
# <翻译结束>


<原文开始>
// Only retrieve exported attributes.
<原文结束>

# <翻译开始>
		// 只检索导出的属性。 md5:d8185f07060feffb
# <翻译结束>


<原文开始>
// If this is an embedded attribute, it retrieves the tags recursively.
<原文结束>

# <翻译开始>
		// 如果这是一个嵌入属性，它将递归地获取标签。 md5:ed1233074f938682
# <翻译结束>

