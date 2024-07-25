
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Fields appends `fieldNamesOrMapStruct` to the operation fields of the model, multiple fields joined using char ','.
// The parameter `fieldNamesOrMapStruct` can be type of string/map/*map/struct/*struct.
//
// Eg:
// Fields("id", "name", "age")
// Fields([]string{"id", "name", "age"})
// Fields(map[string]interface{}{"id":1, "name":"john", "age":18})
// Fields(User{ Id: 1, Name: "john", Age: 18}).
<原文结束>

# <翻译开始>
// `Fields` 方法将 `fieldNamesOrMapStruct` 添加到模型的operation字段中，多个字段使用字符`,`连接。
// 参数 `fieldNamesOrMapStruct` 可以是字符串、映射（map）、*map、结构体或*结构体类型。
//
// 示例：
// Fields("id", "name", "age")
// Fields([]string{"id", "name", "age"})
// Fields(map[string]interface{}{"id":1, "name":"john", "age":18})
// Fields(User{ Id: 1, Name: "john", Age: 18}) md5:21db86fe96babad2
# <翻译结束>


<原文开始>
// FieldsPrefix performs as function Fields but add extra prefix for each field.
<原文结束>

# <翻译开始>
// FieldsPrefix 作为 Fields 函数，但为每个字段添加额外的前缀。 md5:8a672048e8753526
# <翻译结束>


<原文开始>
// FieldsEx appends `fieldNamesOrMapStruct` to the excluded operation fields of the model,
// multiple fields joined using char ','.
// Note that this function supports only single table operations.
// The parameter `fieldNamesOrMapStruct` can be type of string/map/*map/struct/*struct.
//
// Also see Fields.
<原文结束>

# <翻译开始>
// FieldsEx 将 `fieldNamesOrMapStruct` 追加到模型的操作排除字段中，
// 多个字段使用逗号 ',' 连接。
// 注意，此函数仅支持单表操作。
// 参数 `fieldNamesOrMapStruct` 可以是字符串类型、映射类型（map）、映射指针类型（*map）、结构体类型或结构体指针类型（*struct）。
//
// 参见 Fields。 md5:7b8ec243202549a8
# <翻译结束>


<原文开始>
// FieldsExPrefix performs as function FieldsEx but add extra prefix for each field.
<原文结束>

# <翻译开始>
// FieldsExPrefix 函数与 FieldsEx 功能相似，但在每个字段前添加额外的前缀。 md5:66ba7040b83e6e81
# <翻译结束>


<原文开始>
// FieldCount formats and appends commonly used field `COUNT(column)` to the select fields of model.
<原文结束>

# <翻译开始>
// FieldCount 将常用字段 `COUNT(column)` 格式化并添加到模型的 select 字段中。 md5:99439830c058a91f
# <翻译结束>


<原文开始>
// FieldSum formats and appends commonly used field `SUM(column)` to the select fields of model.
<原文结束>

# <翻译开始>
// FieldSum 将常用字段 `SUM(column)` 格式化后添加到模型的 select 字段中。 md5:938249bb2923fa1f
# <翻译结束>


<原文开始>
// FieldMin formats and appends commonly used field `MIN(column)` to the select fields of model.
<原文结束>

# <翻译开始>
// FieldMin 格式化并追加常用的字段 `MIN(column)` 到模型的选择字段中。 md5:fd1204ad66608451
# <翻译结束>


<原文开始>
// FieldMax formats and appends commonly used field `MAX(column)` to the select fields of model.
<原文结束>

# <翻译开始>
// FieldMax 格式化并追加常用的字段 `MAX(column)` 到模型的选择字段中。 md5:77150e433b0d44c4
# <翻译结束>


<原文开始>
// FieldAvg formats and appends commonly used field `AVG(column)` to the select fields of model.
<原文结束>

# <翻译开始>
// FieldAvg 将常用字段 `AVG(column)` 格式化并添加到模型的 select 字段中。 md5:0b09ffae1b0cbabe
# <翻译结束>


<原文开始>
// GetFieldsStr retrieves and returns all fields from the table, joined with char ','.
// The optional parameter `prefix` specifies the prefix for each field, eg: GetFieldsStr("u.").
<原文结束>

# <翻译开始>
// GetFieldsStr 从表中检索并返回所有字段，以逗号分隔。
// 可选参数 `prefix` 指定每个字段的前缀，例如：GetFieldsStr("u.")。 md5:c76f2f45c8680a27
# <翻译结束>


<原文开始>
// GetFieldsExStr retrieves and returns fields which are not in parameter `fields` from the table,
// joined with char ','.
// The parameter `fields` specifies the fields that are excluded.
// The optional parameter `prefix` specifies the prefix for each field, eg: FieldsExStr("id", "u.").
<原文结束>

# <翻译开始>
// GetFieldsExStr 从表中获取并返回那些不在参数`fields`中的字段，这些字段通过逗号','连接。
// 参数`fields`指定了需要排除的字段。
// 可选参数`prefix`为每个字段指定前缀，例如：FieldsExStr("id", "u.")。 md5:57698a0c43f54ec9
# <翻译结束>


<原文开始>
// HasField determine whether the field exists in the table.
<原文结束>

# <翻译开始>
// HasField 用于判断该字段是否存在于表中。 md5:e26ad0ecb292096b
# <翻译结束>


<原文开始>
// getFieldsFrom retrieves, filters and returns fields name from table `table`.
<原文结束>

# <翻译开始>
// getFieldsFrom 从表格`table`中获取、过滤并返回字段名。 md5:9a2c6dffbdfe3d24
# <翻译结束>


<原文开始>
// It needs type asserting.
<原文结束>

# <翻译开始>
	// 需要类型断言。 md5:ec336d143828f70d
# <翻译结束>

