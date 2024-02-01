
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
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// GetFieldTypeStr retrieves and returns the field type string for certain field by name.
<原文结束>

# <翻译开始>
// GetFieldTypeStr 通过名称检索并返回特定字段的字段类型字符串。
# <翻译结束>


<原文开始>
// GetFieldType retrieves and returns the field type object for certain field by name.
<原文结束>

# <翻译开始>
// GetFieldType通过名称获取并返回特定字段的字段类型对象。
# <翻译结束>


<原文开始>
// ConvertDataForRecord is a very important function, which does converting for any data that
// will be inserted into table/collection as a record.
//
// The parameter `value` should be type of *map/map/*struct/struct.
// It supports embedded struct definition for struct.
<原文结束>

# <翻译开始>
// ConvertDataForRecord 是一个非常重要的函数，用于将任何要作为记录插入到表/集合中的数据进行转换。
//
// 参数 `value` 应为 *map、map、*struct 或 struct 类型。对于 struct，它支持嵌套的 struct 定义。
# <翻译结束>


<原文开始>
// ConvertValueForField converts value to the type of the record field.
// The parameter `fieldType` is the target record field.
// The parameter `fieldValue` is the value that to be committed to record field.
<原文结束>

# <翻译开始>
// ConvertValueForField 将值转换为目标记录字段的类型。
// 参数`fieldType`是目标记录字段。
// 参数`fieldValue`是要提交到记录字段的值。
# <翻译结束>


<原文开始>
// If `value` implements interface `driver.Valuer`, it then uses the interface for value converting.
<原文结束>

# <翻译开始>
// 如果`value`实现了接口`driver.Valuer`，那么它将使用该接口进行值转换。
# <翻译结束>







<原文开始>
// It should ignore the bytes type.
<原文结束>

# <翻译开始>
// 它应当忽略 bytes 类型。
# <翻译结束>







<原文开始>
		// If the time is zero, it then updates it to nil,
		// which will insert/update the value to database as "null".
<原文结束>

# <翻译开始>
// 如果时间是零值，那么将其更新为nil，
// 这将会把该值以"null"的形式插入/更新到数据库中。
# <翻译结束>


<原文开始>
			// If `value` implements interface iNil,
			// check its IsNil() function, if got ture,
			// which will insert/update the value to database as "null".
<原文结束>

# <翻译开始>
// 如果`value`实现了接口iNil，
// 检查其IsNil()函数，如果得到true，
// 则将该值作为"null"插入/更新到数据库中。
// 在Go语言中，这段代码的注释描述了当变量value实现了名为iNil的接口时，会进一步调用其IsNil()方法进行判断。若该方法返回true，则会在对数据库进行操作时，将这个变量值视为"null"进行插入或更新操作。
# <翻译结束>


<原文开始>
// Use string conversion in default.
<原文结束>

# <翻译开始>
// 默认情况下使用字符串转换
# <翻译结束>







<原文开始>
// CheckLocalTypeForField checks and returns corresponding type for given db type.
<原文结束>

# <翻译开始>
// CheckLocalTypeForField 检查并返回给定数据库类型对应的本地类型。
# <翻译结束>


<原文开始>
// It is suggested using bit(1) as boolean.
<原文结束>

# <翻译开始>
// 建议使用bit(1)作为布尔值。
# <翻译结束>


<原文开始>
// mssql is true|false string.
<原文结束>

# <翻译开始>
// mssql 是一个表示真或假的字符串。
# <翻译结束>


<原文开始>
// Auto-detect field type, using key match.
<原文结束>

# <翻译开始>
// 自动检测字段类型，通过键匹配实现。
# <翻译结束>


<原文开始>
// ConvertValueForLocal converts value to local Golang type of value according field type name from database.
// The parameter `fieldType` is in lower case, like:
// `float(5,2)`, `unsigned double(5,2)`, `decimal(10,2)`, `char(45)`, `varchar(100)`, etc.
<原文结束>

# <翻译开始>
// ConvertValueForLocal 将值根据数据库字段类型名称转换为本地 Golang 类型的值。
// 参数 `fieldType` 为小写形式，例如：
// `float(5,2)`、`unsigned double(5,2)`、`decimal(10,2)`、`char(45)`、`varchar(100)` 等。
# <翻译结束>


<原文开始>
	// If there's no type retrieved, it returns the `fieldValue` directly
	// to use its original data type, as `fieldValue` is type of interface{}.
<原文结束>

# <翻译开始>
// 如果没有检索到类型，它会直接返回 `fieldValue`，使用其原始数据类型，
// 因为 `fieldValue` 是 interface{} 类型。
# <翻译结束>







<原文开始>
// mappingAndFilterData automatically mappings the map key to table field and removes
// all key-value pairs that are not the field of given table.
<原文结束>

# <翻译开始>
// mappingAndFilterData 自动将映射键映射到表字段，并移除所有非给定表字段的键值对。
# <翻译结束>


<原文开始>
// Automatic data key to table field name mapping.
<原文结束>

# <翻译开始>
// 自动将数据键映射到表字段名称。
# <翻译结束>


<原文开始>
	// Data filtering.
	// It deletes all key-value pairs that has incorrect field name.
<原文结束>

# <翻译开始>
// 数据过滤。
// 删除所有具有错误字段名的键值对。
# <翻译结束>


<原文开始>
// Default value converting.
<原文结束>

# <翻译开始>
// 默认值转换
# <翻译结束>


<原文开始>
// Convert the value to JSON.
<原文结束>

# <翻译开始>
// 将值转换为JSON。
# <翻译结束>


<原文开始>
// Date without time.
<原文结束>

# <翻译开始>
// 仅日期，不含时间。
# <翻译结束>

