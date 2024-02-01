
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//
# <翻译结束>


<原文开始>
// GetDB returns the underlying DB.
<原文结束>

# <翻译开始>
// GetDB 返回底层的 DB（数据库）
# <翻译结束>


<原文开始>
// GetLink creates and returns the underlying database link object with transaction checks.
// The parameter `master` specifies whether using the master node if master-slave configured.
<原文结束>

# <翻译开始>
// GetLink 创建并返回底层的数据库连接对象，同时进行事务检查。
// 参数 `master` 指定在主从配置的情况下是否使用主节点。
# <翻译结束>


<原文开始>
// MasterLink acts like function Master but with additional `schema` parameter specifying
// the schema for the connection. It is defined for internal usage.
// Also see Master.
<原文结束>

# <翻译开始>
// MasterLink 表现得像函数 Master，但额外添加了一个 `schema` 参数用于指定连接的模式。它被定义为内部使用。
// 有关更多信息，请参阅 Master。
# <翻译结束>


<原文开始>
// SlaveLink acts like function Slave but with additional `schema` parameter specifying
// the schema for the connection. It is defined for internal usage.
// Also see Slave.
<原文结束>

# <翻译开始>
// SlaveLink 表现得像函数 Slave，但额外添加了一个 `schema` 参数用于指定连接的模式。它被定义为内部使用。
// 有关更多信息，请参阅 Slave。
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
// QuoteString quotes string with quote chars. Strings like:
// "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".
//
// The meaning of a `string` can be considered as part of a statement string including columns.
<原文结束>

# <翻译开始>
// QuoteString 用引号字符对字符串进行引用。例如以下字符串：
// "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".
//
// 可以将 `string` 的含义视为包含列部分的语句字符串中的一部分。
# <翻译结束>


<原文开始>
// QuotePrefixTableName adds prefix string and quotes chars for the table.
// It handles table string like:
// "user", "user u",
// "user,user_detail",
// "user u, user_detail ut",
// "user as u, user_detail as ut".
//
// Note that, this will automatically checks the table prefix whether already added,
// if true it does nothing to the table name, or else adds the prefix to the table name.
<原文结束>

# <翻译开始>
// QuotePrefixTableName 为表名添加前缀字符串和引用字符。
// 它处理诸如以下格式的表名：
// "user", "user u",
// "user,user_detail",
// "user u, user_detail ut",
// "user as u, user_detail as ut"。
//
// 注意，此函数会自动检查表名是否已添加了前缀，
// 如果已经添加，则不对表名做任何处理；否则，将前缀添加到表名中。
# <翻译结束>


<原文开始>
// GetChars returns the security char for current database.
// It does nothing in default.
<原文结束>

# <翻译开始>
// GetChars 返回当前数据库的安全字符。
// 默认情况下，此方法不做任何操作。
# <翻译结束>


<原文开始>
// Tables retrieves and returns the tables of current schema.
// It's mainly used in cli tool chain for automatically generating the models.
<原文结束>

# <翻译开始>
// Tables 获取并返回当前模式的表。
// 它主要用于cli工具链中，用于自动生成模型。
# <翻译结束>


<原文开始>
// TableFields retrieves and returns the fields' information of specified table of current
// schema.
//
// The parameter `link` is optional, if given nil it automatically retrieves a raw sql connection
// as its link to proceed necessary sql query.
//
// Note that it returns a map containing the field name and its corresponding fields.
// As a map is unsorted, the TableField struct has an "Index" field marks its sequence in
// the fields.
//
// It's using cache feature to enhance the performance, which is never expired util the
// process restarts.
<原文结束>

# <翻译开始>
// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 参数 `link` 是可选的，如果给出 nil，则会自动获取一个原始的 SQL 连接作为其链接以执行必要的 SQL 查询。
//
// 注意，它返回一个包含字段名及其对应字段信息的映射。由于映射是无序的，TableField 结构体有一个 "Index" 字段用于标记其在所有字段中的顺序。
//
// 为了提高性能，该函数使用了缓存特性，缓存有效期直到进程重启才会过期。
# <翻译结束>


<原文开始>
// ClearTableFields removes certain cached table fields of current configuration group.
<原文结束>

# <翻译开始>
// ClearTableFields 删除当前配置组中特定的缓存表字段。
# <翻译结束>


<原文开始>
// ClearTableFieldsAll removes all cached table fields of current configuration group.
<原文结束>

# <翻译开始>
// ClearTableFieldsAll 清除当前配置组中所有已缓存的表字段。
# <翻译结束>


<原文开始>
// ClearCache removes cached sql result of certain table.
<原文结束>

# <翻译开始>
// ClearCache 清除特定表的缓存SQL结果。
# <翻译结束>


<原文开始>
// ClearCacheAll removes all cached sql result from cache
<原文结束>

# <翻译开始>
// ClearCacheAll 从缓存中移除所有已缓存的SQL查询结果
# <翻译结束>


<原文开始>
// HasField determine whether the field exists in the table.
<原文结束>

# <翻译开始>
// HasField 判断字段是否在表中存在。
# <翻译结束>


<原文开始>
// guessPrimaryTableName parses and returns the primary table name.
<原文结束>

# <翻译开始>
// guessPrimaryTableName 解析并返回主表名称。
# <翻译结束>

