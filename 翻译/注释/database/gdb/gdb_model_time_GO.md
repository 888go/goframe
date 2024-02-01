
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
// Default field names of table for automatic-filled created datetime.
<原文结束>

# <翻译开始>
// 默认表字段名称，用于自动填充创建日期和时间。
# <翻译结束>


<原文开始>
// Default field names of table for automatic-filled updated datetime.
<原文结束>

# <翻译开始>
// 默认表字段名称，用于自动填充更新的日期和时间。
# <翻译结束>


<原文开始>
// Default field names of table for automatic-filled deleted datetime.
<原文结束>

# <翻译开始>
// 默认表字段名称，用于自动填充删除时间戳。
# <翻译结束>


<原文开始>
// Unscoped disables the auto-update time feature for insert, update and delete options.
<原文结束>

# <翻译开始>
// Unscoped 禁用在插入、更新和删除选项时自动更新时间的特性。
# <翻译结束>


<原文开始>
// getSoftFieldNameCreate checks and returns the field name for record creating time.
// If there's no field name for storing creating time, it returns an empty string.
// It checks the key with or without cases or chars '-'/'_'/'.'/' '.
<原文结束>

# <翻译开始>
// getSoftFieldNameCreate 检查并返回记录创建时间的字段名称。
// 如果没有存储创建时间的字段名称，则返回一个空字符串。
// 它会检查包含或不包含大小写、字符 '-'/'_'/'.'/' ' 的键。
# <翻译结束>


<原文开始>
// It checks whether this feature disabled.
<原文结束>

# <翻译开始>
// 它用于检查该特性是否已禁用。
# <翻译结束>


<原文开始>
// getSoftFieldNameUpdate checks and returns the field name for record updating time.
// If there's no field name for storing updating time, it returns an empty string.
// It checks the key with or without cases or chars '-'/'_'/'.'/' '.
<原文结束>

# <翻译开始>
// getSoftFieldNameUpdate 检查并返回记录更新时间所对应的字段名称。
// 如果没有存储更新时间的字段名称，则返回一个空字符串。
// 它会检查包含或不包含大小写字符、'-'、'_'、'.'/' '等字符的关键字。
# <翻译结束>


<原文开始>
// getSoftFieldNameDelete checks and returns the field name for record deleting time.
// If there's no field name for storing deleting time, it returns an empty string.
// It checks the key with or without cases or chars '-'/'_'/'.'/' '.
<原文结束>

# <翻译开始>
// getSoftFieldNameDelete 检查并返回记录删除时间所使用的字段名。
// 如果没有存储删除时间的字段名，则返回一个空字符串。
// 它会检查包含或不包含大小写、字符 '-'/'_'/'.'/' ' 的键。
# <翻译结束>


<原文开始>
// getSoftFieldName retrieves and returns the field name of the table for possible key.
<原文结束>

# <翻译开始>
// getSoftFieldName 获取并返回表中可能键的字段名称。
# <翻译结束>


<原文开始>
// Ignore the error from TableFields.
<原文结束>

# <翻译开始>
// 忽略 TableFields 函数返回的错误。
# <翻译结束>


<原文开始>
// getConditionForSoftDeleting retrieves and returns the condition string for soft deleting.
// It supports multiple tables string like:
// "user u, user_detail ud"
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid)"
// "user LEFT JOIN user_detail ON(user_detail.uid=user.uid)"
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid) LEFT JOIN user_stats us ON(us.uid=u.uid)".
<原文结束>

# <翻译开始>
// getConditionForSoftDeleting 获取并返回用于软删除的条件字符串。
// 它支持多种表字符串，例如：
// "user u, user_detail ud" // 多个表别名定义
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid)" // 左连接查询语句
// "user LEFT JOIN user_detail ON(user_detail.uid=user.uid)" // 简化的左连接查询语句
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid) LEFT JOIN user_stats us ON(us.uid=u.uid)" // 多表左连接查询语句
// 该函数用于根据给定的多表查询条件，生成适用于软删除操作的SQL条件子句。
# <翻译结束>


<原文开始>
// Multiple joined tables, exclude the sub query sql which contains char '(' and ')'.
<原文结束>

# <翻译开始>
// 多表连接，排除包含 '(' 和 ')' 字符的子查询SQL语句。
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
// getConditionOfTableStringForSoftDeleting 函数的作用正如其名称所描述的那样。
// `s` 参数的例子：
// - `test`.`demo` as b （将`test`数据库中的`demo`表别名为b）
// - `test`.`demo` b （在`test`数据库中引用`demo`表，此处的 b 可能是别名或语法错误）
// - `demo` （假设是在当前默认数据库中引用`demo`表）
// - demo （与上例类似，直接引用`demo`表，未指定数据库）
# <翻译结束>


<原文开始>
// Multiple base tables.
<原文结束>

# <翻译开始>
// 多个基础表。
# <翻译结束>

