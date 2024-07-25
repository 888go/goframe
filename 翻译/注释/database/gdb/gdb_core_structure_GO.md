
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
// GetFieldTypeStr retrieves and returns the field type string for certain field by name.
<原文结束>

# <翻译开始>
// GetFieldTypeStr 通过名称检索并返回指定字段的字段类型字符串。 md5:aeb8d310c854c45a
# <翻译结束>


<原文开始>
// GetFieldType retrieves and returns the field type object for certain field by name.
<原文结束>

# <翻译开始>
// GetFieldType 通过字段名获取并返回该字段的类型对象。 md5:eeebff59dbaf1064
# <翻译结束>


<原文开始>
// ConvertDataForRecord is a very important function, which does converting for any data that
// will be inserted into table/collection as a record.
//
// The parameter `value` should be type of *map/map/*struct/struct.
// It supports embedded struct definition for struct.
<原文结束>

# <翻译开始>
// ConvertDataForRecord 是一个非常重要的函数，用于将任何数据转换为
// 以便将其作为记录插入到表或集合中。
//
// 参数 `value` 应为 *map/map/*struct/struct 类型。
// 对于结构体，它支持嵌入式结构体定义。 md5:27b867ec3a1c3c1d
# <翻译结束>


<原文开始>
// ConvertValueForField converts value to the type of the record field.
// The parameter `fieldType` is the target record field.
// The parameter `fieldValue` is the value that to be committed to record field.
<原文结束>

# <翻译开始>
// ConvertValueForField 将值转换为记录字段的类型。
// 参数 `fieldType` 是目标记录字段。
// 参数 `fieldValue` 是要写入记录字段的值。 md5:196c02c9f6cf3380
# <翻译结束>


<原文开始>
// If `value` implements interface `driver.Valuer`, it then uses the interface for value converting.
<原文结束>

# <翻译开始>
	// 如果`value`实现了`driver.Valuer`接口，那么它会使用该接口进行值的转换。 md5:ba72a317b79988a2
# <翻译结束>


<原文开始>
// Default value converting.
<原文结束>

# <翻译开始>
	// 默认值转换。 md5:99c30401ccc62832
# <翻译结束>


<原文开始>
// It should ignore the bytes type.
<原文结束>

# <翻译开始>
		// 它应该忽略字节类型。 md5:48e040f1decb1a5e
# <翻译结束>


<原文开始>
// Convert the value to JSON.
<原文结束>

# <翻译开始>
				// 将值转换为JSON。 md5:f4977a0ae972c910
# <翻译结束>


<原文开始>
		// If the time is zero, it then updates it to nil,
		// which will insert/update the value to database as "null".
<原文结束>

# <翻译开始>
		// 如果时间是零值，它将更新为nil，
		// 这样在数据库中插入或更新的值将会是"null"。 md5:058aebae61025f37
# <翻译结束>


<原文开始>
			// If `value` implements interface iNil,
			// check its IsNil() function, if got ture,
			// which will insert/update the value to database as "null".
<原文结束>

# <翻译开始>
			// 如果`value`实现了iNil接口，
			// 检查其IsNil()函数，如果返回true，
			// 将把该值插入/更新到数据库中作为"null"。 md5:b2415061d93829e6
# <翻译结束>


<原文开始>
// Use string conversion in default.
<原文结束>

# <翻译开始>
				// 默认使用字符串转换。 md5:36cba4c54f848f87
# <翻译结束>


<原文开始>
// CheckLocalTypeForField checks and returns corresponding type for given db type.
<原文结束>

# <翻译开始>
// CheckLocalTypeForField 检查并返回与给定数据库类型相对应的本地类型。 md5:d3191e6393b7e531
# <翻译结束>


<原文开始>
// It is suggested using bit(1) as boolean.
<原文结束>

# <翻译开始>
		// 建议使用 bit(1) 作为布尔值。 md5:5be00c9e8395ea93
# <翻译结束>


<原文开始>
// mssql is true|false string.
<原文结束>

# <翻译开始>
		// mssql 是一个true|false类型的字符串。 md5:6d4dbdb95d9adfa1
# <翻译结束>


<原文开始>
// Auto-detect field type, using key match.
<原文结束>

# <翻译开始>
		// 自动检测字段类型，通过键匹配。 md5:138e4aeac8d26d8a
# <翻译结束>


<原文开始>
// ConvertValueForLocal converts value to local Golang type of value according field type name from database.
// The parameter `fieldType` is in lower case, like:
// `float(5,2)`, `unsigned double(5,2)`, `decimal(10,2)`, `char(45)`, `varchar(100)`, etc.
<原文结束>

# <翻译开始>
// ConvertValueForLocal 根据从数据库中获取的字段类型名称，将值转换为Go语言中的本地类型。
// 参数 `fieldType` 为小写格式，例如：
// `float(5,2)`，`unsigned double(5,2)`，`decimal(10,2)`，`char(45)`，`varchar(100)` 等。 md5:7e1ede2b68158e31
# <翻译结束>


<原文开始>
	// If there's no type retrieved, it returns the `fieldValue` directly
	// to use its original data type, as `fieldValue` is type of interface{}.
<原文结束>

# <翻译开始>
	// 如果没有获取到类型，则直接返回`fieldValue`，
	// 利用其原始数据类型，因为`fieldValue`是`interface{}`类型的。 md5:62cf4d391c9da4f2
# <翻译结束>


<原文开始>
// mappingAndFilterData automatically mappings the map key to table field and removes
// all key-value pairs that are not the field of given table.
<原文结束>

# <翻译开始>
// mappingAndFilterData 自动将映射键映射到表格字段，并删除所有不是给定表格字段的键值对。 md5:27fc8e27d4d4a389
# <翻译结束>


<原文开始>
// Automatic data key to table field name mapping.
<原文结束>

# <翻译开始>
	// 自动将数据键映射到表格字段名。 md5:bdc9aa8a688bb975
# <翻译结束>


<原文开始>
	// Data filtering.
	// It deletes all key-value pairs that has incorrect field name.
<原文结束>

# <翻译开始>
	// 数据过滤。
	// 它会删除所有具有错误字段名的键值对。 md5:24aafcb1699db80c
# <翻译结束>

