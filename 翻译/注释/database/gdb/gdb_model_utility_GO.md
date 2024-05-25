
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
// QuoteWord checks given string `s` a word,
// if true it quotes `s` with security chars of the database
// and returns the quoted string; or else it returns `s` without any change.
//
// The meaning of a `word` can be considered as a column name.
<原文结束>

# <翻译开始>
// QuoteWord 检查给定的字符串 `s` 是否为一个单词，
// 如果是，它将使用数据库的安全字符对 `s` 进行转义，并返回带引号的字符串；否则，返回原始字符串不做任何更改。
//
// 可以认为一个 `word` 表示列名。
// md5:71291615d7bcffe0
# <翻译结束>


<原文开始>
// TableFields retrieves and returns the fields' information of specified table of current
// schema.
//
// Also see DriverMysql.TableFields.
<原文结束>

# <翻译开始>
// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 参见 DriverMysql.TableFields。
// md5:61e256ba53f813cb
# <翻译结束>


<原文开始>
// getModel creates and returns a cloned model of current model if `safe` is true, or else it returns
// the current model.
<原文结束>

# <翻译开始>
// getModel 如果`safe`为真，则创建并返回当前模型的克隆，否则直接返回当前模型。
// md5:e4ae726aba6b01ab
# <翻译结束>


<原文开始>
// mappingAndFilterToTableFields mappings and changes given field name to really table field name.
// Eg:
// ID        -> id
// NICK_Name -> nickname.
<原文结束>

# <翻译开始>
// mappingAndFilterToTableFields：将给定的字段名映射并转换为实际的表格字段名。
// 例如：
// ID        -> id
// NICK_Name -> nickname.
// md5:35f1e9dc3d13c4f0
# <翻译结束>


<原文开始>
			// Example:
			// user.id, user.name
			// replace(concat_ws(',',lpad(s.id, 6, '0'),s.name),',','') `code`
<原文结束>

# <翻译开始>
// 示例：
// user.id, user.name
// 将逗号分隔的字符串（格式：lpad(s.id, 6, '0')，s.name）替换为`code`
// md5:5ee6374c41194bf3
# <翻译结束>


<原文开始>
				// Example:
				// id, name
<原文结束>

# <翻译开始>
// 示例：
// id, 名称
// md5:f16c15c62075a7aa
# <翻译结束>


<原文开始>
// filterDataForInsertOrUpdate does filter feature with data for inserting/updating operations.
// Note that, it does not filter list item, which is also type of map, for "omit empty" feature.
<原文结束>

# <翻译开始>
// filterDataForInsertOrUpdate 对用于插入/更新操作的数据执行过滤功能。
// 请注意，它不会对列表项（也是一种映射类型）进行“忽略空值”处理。
// md5:ffc8a604eaec8a77
# <翻译结束>


<原文开始>
// doMappingAndFilterForInsertOrUpdateDataMap does the filter features for map.
// Note that, it does not filter list item, which is also type of map, for "omit empty" feature.
<原文结束>

# <翻译开始>
// doMappingAndFilterForInsertOrUpdateDataMap 为映射类型的数据执行过滤功能。
// 注意，它不会对"忽略空"特性下的列表项（也是映射类型）进行过滤。
// md5:93fefbe3176f55de
# <翻译结束>


<原文开始>
// Remove key-value pairs of which the value is nil.
<原文结束>

# <翻译开始>
// 删除值为nil的键值对。 md5:5219c3473c86d38c
# <翻译结束>


<原文开始>
// Remove key-value pairs of which the value is empty.
<原文结束>

# <翻译开始>
// 删除值为空的键值对。 md5:706fbf04684a1301
# <翻译结束>


<原文开始>
// Special type filtering.
<原文结束>

# <翻译开始>
// 特殊类型的过滤。 md5:48598cd9e3395cfc
# <翻译结束>


<原文开始>
// Filter specified fields.
<原文结束>

# <翻译开始>
// 过滤指定字段。 md5:c1817e5f938542f0
# <翻译结束>


<原文开始>
// getLink returns the underlying database link object with configured `linkType` attribute.
// The parameter `master` specifies whether using the master node if master-slave configured.
<原文结束>

# <翻译开始>
// getLink 函数返回配置了 `linkType` 属性的底层数据库链接对象。
// 参数 `master` 指定是否在主从配置中使用主节点。
// md5:e8add2f9371393db
# <翻译结束>


<原文开始>
// getPrimaryKey retrieves and returns the primary key name of the model table.
// It parses m.tables to retrieve the primary table name, supporting m.tables like:
// "user", "user u", "user as u, user_detail as ud".
<原文结束>

# <翻译开始>
// getPrimaryKey 获取并返回模型表的主键名称。
// 它解析 m.tables 以获取主表名，支持如下的 m.tables 写法：
// "user", "user u", "user as u, user_detail as ud"。
// md5:07ea92a426e953d1
# <翻译结束>


<原文开始>
// mergeArguments creates and returns new arguments by merging `m.extraArgs` and given `args`.
<原文结束>

# <翻译开始>
// mergeArguments 将 `m.extraArgs` 和给定的 `args` 合并，创建并返回新的参数。 md5:80f949384113727a
# <翻译结束>

