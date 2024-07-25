
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。 md5:a114f4bdd106ab31
# <翻译结束>


<原文开始>
// GetDB returns the underlying DB.
<原文结束>

# <翻译开始>
// GetDB 返回底层的DB。 md5:5ebeb6e695bd2a8a
# <翻译结束>


<原文开始>
// GetLink creates and returns the underlying database link object with transaction checks.
// The parameter `master` specifies whether using the master node if master-slave configured.
<原文结束>

# <翻译开始>
// GetLink 创建并返回底层数据库链接对象，并进行事务检查。
// 参数 `master` 指定在配置了主从的情况下是否使用主节点。 md5:51315fe7b2e9a929
# <翻译结束>


<原文开始>
// MasterLink acts like function Master but with additional `schema` parameter specifying
// the schema for the connection. It is defined for internal usage.
// Also see Master.
<原文结束>

# <翻译开始>
// MasterLink 类似于函数 Master，但增加了 `schema` 参数，用于指定连接的模式。
// 这个函数主要用于内部使用。同时也参考 Master 函数。 md5:ae74b996555aea95
# <翻译结束>


<原文开始>
// SlaveLink acts like function Slave but with additional `schema` parameter specifying
// the schema for the connection. It is defined for internal usage.
// Also see Slave.
<原文结束>

# <翻译开始>
// SlaveLink 行为类似于 Slave 函数，但增加了 `schema` 参数，用于指定连接的模式。它主要用于内部使用。
// 参阅 Slave。 md5:8a8929395882c04a
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
// 可以认为一个 `word` 表示列名。 md5:71291615d7bcffe0
# <翻译结束>


<原文开始>
// QuoteString quotes string with quote chars. Strings like:
// "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".
//
// The meaning of a `string` can be considered as part of a statement string including columns.
<原文结束>

# <翻译开始>
// QuoteString 使用引号字符对字符串进行引用。例如这样的字符串：
// "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".
//
// `string` 的含义可以理解为作为包含列名的语句字符串的一部分。 md5:09c5263950e9ed1a
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
// QuotePrefixTableName 为表名添加前缀字符串并包围引号。
// 它可以处理如下形式的表字符串：
// "user", "user u",
// "user,user_detail",
// "user u, user_detail ut",
// "user as u, user_detail as ut".
//
// 请注意，此函数会自动检查表前缀是否已经添加，
// 如果是，则不对表名做任何处理，否则会在表名前添加前缀。 md5:46ab3c3833cc0124
# <翻译结束>


<原文开始>
// GetChars returns the security char for current database.
// It does nothing in default.
<原文结束>

# <翻译开始>
// GetChars 返回当前数据库的安全字符。在默认情况下，它不执行任何操作。 md5:681b4cc93b5adecd
# <翻译结束>


<原文开始>
// Tables retrieves and returns the tables of current schema.
// It's mainly used in cli tool chain for automatically generating the models.
<原文结束>

# <翻译开始>
// Tables 获取并返回当前模式下的表格列表。
// 主要用于命令行工具链，用于自动生成模型。 md5:bce161ba95454bf5
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
// TableFields 获取并返回当前模式指定表的字段信息。
//
// 参数 `link` 是可选的，如果为 nil，则自动获取一个原始 SQL 连接，用于执行必要的 SQL 查询。
//
// 它返回一个包含字段名及其对应字段的映射。由于映射是无序的，TableField 结构体有一个 "Index" 字段，标记其在字段中的顺序。
//
// 该方法使用缓存功能来提高性能，直到进程重启，缓存永不过期。 md5:c844572d5210b35e
# <翻译结束>


<原文开始>
// ClearTableFields removes certain cached table fields of current configuration group.
<原文结束>

# <翻译开始>
// ClearTableFields 清除当前配置组中特定的缓存表字段。 md5:061271b8a4f298a0
# <翻译结束>


<原文开始>
// ClearTableFieldsAll removes all cached table fields of current configuration group.
<原文结束>

# <翻译开始>
// ClearTableFieldsAll 删除当前配置组中所有缓存的表字段。 md5:2b2f2ebba86cfda6
# <翻译结束>


<原文开始>
// ClearCache removes cached sql result of certain table.
<原文结束>

# <翻译开始>
// ClearCache 删除特定表的缓存SQL结果。 md5:5849435c2385500b
# <翻译结束>


<原文开始>
// ClearCacheAll removes all cached sql result from cache
<原文结束>

# <翻译开始>
// ClearCacheAll 从缓存中移除所有已缓存的SQL结果. md5:1cafe85ca7b9f62d
# <翻译结束>


<原文开始>
// HasField determine whether the field exists in the table.
<原文结束>

# <翻译开始>
// HasField 用于判断该字段是否存在于表中。 md5:e26ad0ecb292096b
# <翻译结束>


<原文开始>
// guessPrimaryTableName parses and returns the primary table name.
<原文结束>

# <翻译开始>
// guessPrimaryTableName 解析并返回主表名称。 md5:d6aaf3f09d0afaaa
# <翻译结束>

