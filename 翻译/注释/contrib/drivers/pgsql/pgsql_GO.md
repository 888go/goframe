
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
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Package pgsql implements gdb.Driver, which supports operations for database PostgreSQL.
//
// Note:
// 1. It does not support Save/Replace features.
// 2. It does not support Insert Ignore features.
<原文结束>

# <翻译开始>
// Package pgsql 实现了 gdb.Driver 接口，该接口支持对 PostgreSQL 数据库进行操作。
// 注意：
// 1. 该包不支持 Save/Replace 功能。
// 2. 该包不支持 Insert Ignore 特性。
# <翻译结束>


<原文开始>
// Driver is the driver for postgresql database.
<原文结束>

# <翻译开始>
// Driver 是用于 PostgreSQL 数据库的驱动。
# <翻译结束>


<原文开始>
// New create and returns a driver that implements gdb.Driver, which supports operations for PostgreSql.
<原文结束>

# <翻译开始>
// New 创建并返回一个实现了 gdb.Driver 的驱动程序，该驱动程序支持针对 PostgreSql 的操作。
# <翻译结束>


<原文开始>
// New creates and returns a database object for postgresql.
// It implements the interface of gdb.Driver for extra database driver installation.
<原文结束>

# <翻译开始>
// New 创建并返回一个用于 PostgreSQL 的数据库对象。
// 它实现了 gdb.Driver 接口，以便进行额外的数据库驱动安装。
# <翻译结束>


<原文开始>
// Open creates and returns an underlying sql.DB object for pgsql.
// https://pkg.go.dev/github.com/lib/pq
<原文结束>

# <翻译开始>
// Open 创建并返回一个用于pgsql的底层sql.DB对象。
// 参考文档：https://pkg.go.dev/github.com/lib/pq
# <翻译结束>


<原文开始>
		// ============================================================================
		// Deprecated from v2.2.0.
		// ============================================================================
<原文结束>

# <翻译开始>
// ============================================================================
// 从 v2.2.0 版本开始已弃用。
// ============================================================================
# <翻译结束>


<原文开始>
// Custom changing the schema in runtime.
<原文结束>

# <翻译开始>
// 在运行时自定义更改架构
# <翻译结束>


<原文开始>
// GetChars returns the security char for this type of database.
<原文结束>

# <翻译开始>
// GetChars 返回该类型数据库的安全字符。
# <翻译结束>


<原文开始>
// CheckLocalTypeForField checks and returns corresponding local golang type for given db type.
<原文结束>

# <翻译开始>
// CheckLocalTypeForField 检查并返回给定数据库类型对应的本地 Go 语言类型。
# <翻译结束>


<原文开始>
// ConvertValueForLocal converts value to local Golang type of value according field type name from database.
// The parameter `fieldType` is in lower case, like:
// `float(5,2)`, `unsigned double(5,2)`, `decimal(10,2)`, `char(45)`, `varchar(100)`, etc.
<原文结束>

# <翻译开始>
// ConvertValueForLocal 将值根据数据库中的字段类型名称转换为本地 Go 语言类型的值。
// 参数 `fieldType` 为小写形式，例如：
// `float(5,2)`, `unsigned double(5,2)`, `decimal(10,2)`, `char(45)`, `varchar(100)` 等。
# <翻译结束>


<原文开始>
// For pgsql, int2 = smallint and int4 = integer.
<原文结束>

# <翻译开始>
// 对于pgsql，int2代表smallint，而int4代表integer。
# <翻译结束>


<原文开始>
// DoFilter deals with the sql string before commits it to underlying sql driver.
<原文结束>

# <翻译开始>
// DoFilter 在将 SQL 字符串提交给底层 SQL 驱动程序之前，对其进行处理。
# <翻译结束>


<原文开始>
// Convert placeholder char '?' to string "$x".
<原文结束>

# <翻译开始>
// 将占位符字符 '?' 转换为字符串 "$x"。
# <翻译结束>


<原文开始>
	// Handle pgsql jsonb feature support, which contains place-holder char '?'.
	// Refer:
	// https://github.com/gogf/gf/issues/1537
	// https://www.postgresql.org/docs/12/functions-json.html
<原文结束>

# <翻译开始>
// 处理pgsql对jsonb特性的支持，其中包含占位符字符 '?'。
// 参考：
// https://github.com/gogf/gf/issues/1537
// https://www.postgresql.org/docs/12/functions-json.html
// 这段Go语言代码的注释翻译成中文后，其含义为：
// 该处用于处理PostgreSQL中对jsonb类型功能的支持，这些功能可能包含问号（'?'）作为占位符字符。
// 参考文档：
// GitHub上gf框架的issue #1537
// PostgreSQL官方文档中关于12版本的JSON函数介绍
# <翻译结束>


<原文开始>
// Tables retrieves and returns the tables of current schema.
// It's mainly used in cli tool chain for automatically generating the models.
<原文结束>

# <翻译开始>
// Tables 获取并返回当前模式的表格。
// 它主要用于cli工具链中，用于自动生成模型。
# <翻译结束>


<原文开始>
// DO NOT use `usedSchema` as parameter for function `SlaveLink`.
<原文结束>

# <翻译开始>
// **请勿**将 `usedSchema` 作为参数传递给函数 `SlaveLink`。
# <翻译结束>


<原文开始>
// version checks and returns the database version.
<原文结束>

# <翻译开始>
// version 检查并返回数据库版本。
# <翻译结束>


<原文开始>
// TableFields retrieves and returns the fields' information of specified table of current schema.
<原文结束>

# <翻译开始>
// TableFields 获取并返回当前模式下指定表的字段信息。
# <翻译结束>


<原文开始>
// TODO duplicated `id` result?
<原文结束>

# <翻译开始>
// TODO 是否存在重复的`id`结果？
# <翻译结束>


<原文开始>
// DoInsert inserts or updates data forF given table.
<原文结束>

# <翻译开始>
// DoInsert 对给定表执行插入或更新数据操作。
# <翻译结束>


<原文开始>
// DoExec commits the sql string and its arguments to underlying driver
// through given link object and returns the execution result.
<原文结束>

# <翻译开始>
// DoExec 通过给定的link对象，将sql字符串及其参数提交到底层驱动，并返回执行结果。
# <翻译结束>


<原文开始>
// Check whether the default method needs to be used
<原文结束>

# <翻译开始>
// 检查是否需要使用默认方法
# <翻译结束>


<原文开始>
// Firstly, check and retrieve transaction link from context.
<原文结束>

# <翻译开始>
// 首先，从上下文检查并检索事务链接。
# <翻译结束>


<原文开始>
// Or else it creates one from master node.
<原文结束>

# <翻译开始>
// 或者从主节点创建一个。
# <翻译结束>


<原文开始>
// If current link is not transaction link, it checks and retrieves transaction from context.
<原文结束>

# <翻译开始>
// 如果当前链接不是事务链接，则检查并从上下文中检索事务。
# <翻译结束>


<原文开始>
// Check if it is an insert operation with primary key.
<原文结束>

# <翻译开始>
// 检查是否为主键插入操作。
# <翻译结束>


<原文开始>
// check if it is an insert operation.
<原文结束>

# <翻译开始>
// 检查是否为插入操作。
# <翻译结束>


<原文开始>
// Only the insert operation with primary key can execute the following code
<原文结束>

# <翻译开始>
// 仅当使用主键执行插入操作时，才能执行以下代码
# <翻译结束>


<原文开始>
// For pgsql, int2 = smallint.
<原文结束>

# <翻译开始>
// 对于pgsql，int2等于smallint。
# <翻译结束>


<原文开始>
// For pgsql, int4 = integer
<原文结束>

# <翻译开始>
// 对于pgsql，int4等于integer（整数类型）
# <翻译结束>


<原文开始>
// For pgsql, int8 = bigint
<原文结束>

# <翻译开始>
// 对于pgsql，int8 等同于 bigint
# <翻译结束>


<原文开始>
// For pgsql, int8 = bigint.
<原文结束>

# <翻译开始>
// 对于pgsql，int8 等同于 bigint。
# <翻译结束>


<原文开始>
// Filter duplicated fields.
<原文结束>

# <翻译开始>
// 过滤重复字段。
# <翻译结束>


<原文开始>
// Transaction checks.
<原文结束>

# <翻译开始>
// 事务检查。
# <翻译结束>


<原文开始>
// use default DoExec
<原文结束>

# <翻译开始>
// 使用默认的DoExec
# <翻译结束>

