
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
// SoftTimeType custom defines the soft time field type.
<原文结束>

# <翻译开始>
// SoftTimeType自定义定义软时间字段类型。 md5:dac7cb3a21ca2d1d
# <翻译结束>


<原文开始>
// (Default)Auto detect the field type by table field type.
<原文结束>

# <翻译开始>
// (默认)根据表字段类型自动检测字段类型。 md5:92b9309fac6f4d09
# <翻译结束>


<原文开始>
// Using datetime as the field value.
<原文结束>

# <翻译开始>
// 使用日期时间作为字段值。 md5:fb9be1cf84e4192e
# <翻译结束>


<原文开始>
// SoftTimeOption is the option to customize soft time feature for Model.
<原文结束>

# <翻译开始>
// SoftTimeOption 是用于自定义 Model 的软时间功能的选项。 md5:fcc19f5ef8ad45e7
# <翻译结束>


<原文开始>
// The value type for soft time field.
<原文结束>

# <翻译开始>
// 软时间字段的值类型。 md5:472088e64d8a928f
# <翻译结束>


<原文开始>
// getSoftFieldNameAndTypeCacheItem is the internal struct for storing create/update/delete fields.
<原文结束>

# <翻译开始>
// getSoftFieldNameAndTypeCacheItem 是用于存储创建/更新/删除字段的内部结构体。 md5:df4233b79c5f6dad
# <翻译结束>


<原文开始>
// Default field names of table for automatic-filled for record creating.
<原文结束>

# <翻译开始>
	// 当创建记录时，用于自动填充的表的默认字段名称。 md5:58c0524feef22203
# <翻译结束>


<原文开始>
// Default field names of table for automatic-filled for record updating.
<原文结束>

# <翻译开始>
	// 用于记录更新时自动填充的表默认字段名称。 md5:dfaf612ced6164b4
# <翻译结束>


<原文开始>
// Default field names of table for automatic-filled for record deleting.
<原文结束>

# <翻译开始>
	// 默认的表字段名，用于自动填充记录删除。 md5:82caa57d9d8aac21
# <翻译结束>


<原文开始>
// SoftTime sets the SoftTimeOption to customize soft time feature for Model.
<原文结束>

# <翻译开始>
// SoftTime 设置 SoftTimeOption 以自定义 Model 的软时间功能。 md5:6c4368abcd89e6b0
# <翻译结束>


<原文开始>
// Unscoped disables the soft time feature for insert, update and delete operations.
<原文结束>

# <翻译开始>
// Unscoped禁用插入、更新和删除操作的软时间特性。 md5:0fc4af29459bd61e
# <翻译结束>


<原文开始>
// GetFieldNameAndTypeForCreate checks and returns the field name for record creating time.
// If there's no field name for storing creating time, it returns an empty string.
// It checks the key with or without cases or chars '-'/'_'/'.'/' '.
<原文结束>

# <翻译开始>
// GetFieldNameAndTypeForCreate 检查并返回用于记录创建时间的字段名。
// 如果没有用于存储创建时间的字段名，它将返回一个空字符串。
// 它会检查键名，无论大小写或包含字符 '-'、'_'、'.'、' '。
// md5:c03150380846ea77
# <翻译结束>


<原文开始>
// It checks whether this feature disabled.
<原文结束>

# <翻译开始>
	// 检查是否禁用了此功能。 md5:413ae315bebe927f
# <翻译结束>


<原文开始>
// GetFieldNameAndTypeForUpdate checks and returns the field name for record updating time.
// If there's no field name for storing updating time, it returns an empty string.
// It checks the key with or without cases or chars '-'/'_'/'.'/' '.
<原文结束>

# <翻译开始>
// GetFieldNameAndTypeForUpdate 检查并返回用于更新时间的字段名。如果没有用于存储更新时间的字段名，它将返回空字符串。它会检查带有或不带大小写、字符 '-'/'_'/'.'/' 的键。
// md5:220eb56737359035
# <翻译结束>


<原文开始>
// GetFieldNameAndTypeForDelete checks and returns the field name for record deleting time.
// If there's no field name for storing deleting time, it returns an empty string.
// It checks the key with or without cases or chars '-'/'_'/'.'/' '.
<原文结束>

# <翻译开始>
// GetFieldNameAndTypeForDelete 检查并返回记录删除时间的字段名。如果没有用于存储删除时间的字段名，它将返回空字符串。它会检查大小写敏感或不敏感，以及使用 '-'、'_'、'.' 或 ' ' 作为分隔符的键。
// md5:f7c6b45838b970b0
# <翻译结束>


<原文开始>
// getSoftFieldName retrieves and returns the field name of the table for possible key.
<原文结束>

# <翻译开始>
// getSoftFieldName 获取并返回表中可能键的字段名。 md5:e32e19240070c456
# <翻译结束>


<原文开始>
// Ignore the error from TableFields.
<原文结束>

# <翻译开始>
			// 忽略TableFields函数的错误。 md5:b488d48f86ec5aea
# <翻译结束>


<原文开始>
// GetWhereConditionForDelete retrieves and returns the condition string for soft deleting.
// It supports multiple tables string like:
// "user u, user_detail ud"
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid)"
// "user LEFT JOIN user_detail ON(user_detail.uid=user.uid)"
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid) LEFT JOIN user_stats us ON(us.uid=u.uid)".
<原文结束>

# <翻译开始>
// GetWhereConditionForDelete 用于检索并返回软删除的条件字符串。它支持多表字符串，例如：
// "user u, user_detail ud" - "用户 u 和 user_detail ud"
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid)" - "用户 u 左连接 user_detail ud，连接条件为 ud.uid 等于 u.uid"
// "user LEFT JOIN user_detail ON(user_detail.uid=user.uid)" - "用户左连接 user_detail，连接条件为 user_detail.uid 等于 user.uid"
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid) LEFT JOIN user_stats us ON(us.uid=u.uid)" - "用户 u 先左连接 user_detail ud，再连接 user_stats us，连接条件为 us.uid 等于 u.uid"
// md5:f2c849c59f2ab188
# <翻译结束>


<原文开始>
// Multiple joined tables, exclude the sub query sql which contains char '(' and ')'.
<原文结束>

# <翻译开始>
		// 多个连接的表，排除包含字符'('和')'的子查询SQL。 md5:a9edf50410c73b2c
# <翻译结束>


<原文开始>
// getConditionOfTableStringForSoftDeleting does something as its name describes.
// Examples for `s`:
// - `test`.`demo` as b
// - `test`.`demo` b
// - `demo`
// - demo
<原文结束>

# <翻译开始>
// getConditionOfTableStringForSoftDeleting 的功能如其名称所述。
// `s` 的示例包括：
// - `test`.`demo` as b
// - `test`.`demo` b
// - `demo`
// - demo
// md5:ffb3e23129e1b6db
# <翻译结束>


<原文开始>
// GetDataByFieldNameAndTypeForDelete creates and returns the placeholder and value for
// specified field name and type in soft-deleting scenario.
<原文结束>

# <翻译开始>
// GetDataByFieldNameAndTypeForDelete 用于在软删除场景下，根据指定的字段名和类型创建并返回占位符和值。
// md5:276be24343264681
# <翻译结束>


<原文开始>
// GetValueByFieldTypeForCreateOrUpdate creates and returns the value for specified field type,
// usually for creating or updating operations.
<原文结束>

# <翻译开始>
// GetValueByFieldTypeForCreateOrUpdate 为创建或更新操作创建并返回指定字段类型的值。
// md5:263c89f2a7abf2da
# <翻译结束>

