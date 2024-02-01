
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
// QuoteWord checks given string `s` a word,
// if true it quotes `s` with security chars of the database
// and returns the quoted string; or else it returns `s` without any change.
//
// The meaning of a `word` can be considered as a column name.
<原文结束>

# <翻译开始>
// 2024-01-09 改成内部方法,此方法属于底层, 几乎用不到.
// QuoteWord 检查给定字符串 `s` 是否为一个单词，
// 如果是，它会使用数据库的安全字符对 `s` 进行引用，并返回引述后的字符串；
// 否则，它将直接返回未经修改的 `s`。
//
// 这里的“单词”可以理解为列名。
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
// 另请参阅 DriverMysql.TableFields。
# <翻译结束>


<原文开始>
// getModel creates and returns a cloned model of current model if `safe` is true, or else it returns
// the current model.
<原文结束>

# <翻译开始>
// getModel函数如果`safe`为真，则创建并返回当前模型的一个克隆副本，否则直接返回当前模型。
# <翻译结束>


<原文开始>
// mappingAndFilterToTableFields mappings and changes given field name to really table field name.
// Eg:
// ID        -> id
// NICK_Name -> nickname.
<原文结束>

# <翻译开始>
// mappingAndFilterToTableFields 将给定的字段名映射并转换为实际的数据库表字段名。
// 例如：
// ID        -> id
// NICK_Name -> nickname.
# <翻译结束>







<原文开始>
// filterDataForInsertOrUpdate does filter feature with data for inserting/updating operations.
// Note that, it does not filter list item, which is also type of map, for "omit empty" feature.
<原文结束>

# <翻译开始>
// filterDataForInsertOrUpdate 对用于插入/更新操作的数据执行过滤功能。
// 注意，它不针对“忽略空值”特性过滤列表项（该列表项也是映射类型）。
# <翻译结束>


<原文开始>
// doMappingAndFilterForInsertOrUpdateDataMap does the filter features for map.
// Note that, it does not filter list item, which is also type of map, for "omit empty" feature.
<原文结束>

# <翻译开始>
// doMappingAndFilterForInsertOrUpdateDataMap 对 map 执行映射和过滤功能。
// 注意，对于“忽略空值”特性，它不会过滤列表项（其类型也为 map）。
# <翻译结束>


<原文开始>
// Remove key-value pairs of which the value is nil.
<原文结束>

# <翻译开始>
// 删除值为nil的键值对。
# <翻译结束>


<原文开始>
// Remove key-value pairs of which the value is empty.
<原文结束>

# <翻译开始>
// 删除值为空的键值对。
# <翻译结束>

















<原文开始>
// getLink returns the underlying database link object with configured `linkType` attribute.
// The parameter `master` specifies whether using the master node if master-slave configured.
<原文结束>

# <翻译开始>
// getLink 返回配置了 `linkType` 属性的基础数据库连接对象。
// 参数 `master` 指定在主从配置时是否使用主节点。
# <翻译结束>


<原文开始>
// getPrimaryKey retrieves and returns the primary key name of the model table.
// It parses m.tables to retrieve the primary table name, supporting m.tables like:
// "user", "user u", "user as u, user_detail as ud".
<原文结束>

# <翻译开始>
// getPrimaryKey 获取并返回模型表的主键名称。
// 它通过解析 m.tables 来检索主表名，支持如下形式的 m.tables：
// "user", "user u", "user as u, user_detail as ud"。
# <翻译结束>


<原文开始>
// mergeArguments creates and returns new arguments by merging `m.extraArgs` and given `args`.
<原文结束>

# <翻译开始>
// mergeArguments通过合并`m.extraArgs`和给定的`args`创建并返回新的参数。
# <翻译结束>


<原文开始>
// Eg: user.id, user.name
<原文结束>

# <翻译开始>
// 示例：user.id, user.name
# <翻译结束>


<原文开始>
// Special type filtering.
<原文结束>

# <翻译开始>
// 特殊类型过滤
# <翻译结束>


<原文开始>
// Keep specified fields.
<原文结束>

# <翻译开始>
// 保留指定字段。
# <翻译结束>


<原文开始>
// Filter specified fields.
<原文结束>

# <翻译开始>
// 过滤指定字段。
# <翻译结束>

