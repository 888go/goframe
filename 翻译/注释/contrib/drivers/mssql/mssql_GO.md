
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
// Package mssql implements gdb.Driver, which supports operations for database MSSql.
//
// Note:
// 1. It does not support Save/Replace features.
// 2. It does not support LastInsertId.
<原文结束>

# <翻译开始>
// Package mssql 实现了 gdb.Driver 接口，该接口支持对 MSSql 数据库进行操作。
//
// 注意：
// 1. 不支持 Save/Replace 功能。
// 2. 不支持 LastInsertId 功能。
// 以下是将上述Go语言代码注释翻译成中文：
// ```markdown
// 这个mssql包实现了gdb.Driver接口，主要用于对MSSql数据库的各种操作提供支持。
//
// 注意事项：
// 1. 该实现暂不支持Save/Replace功能。
// 2. 该实现暂不支持LastInsertId方法。
# <翻译结束>


<原文开始>
// Driver is the driver for SQL server database.
<原文结束>

# <翻译开始>
// Driver 是 SQL 服务器数据库的驱动程序。
# <翻译结束>


<原文开始>
// New create and returns a driver that implements gdb.Driver, which supports operations for Mssql.
<原文结束>

# <翻译开始>
// New 创建并返回一个实现 gdb.Driver 接口的驱动程序，该驱动支持对 Mssql 的操作。
# <翻译结束>


<原文开始>
// New creates and returns a database object for SQL server.
// It implements the interface of gdb.Driver for extra database driver installation.
<原文结束>

# <翻译开始>
// New 创建并返回一个用于SQL服务器的数据库对象。
// 它实现了gdb.Driver接口，以便支持额外的数据库驱动安装。
# <翻译结束>


<原文开始>
// Open creates and returns an underlying sql.DB object for mssql.
<原文结束>

# <翻译开始>
// Open 创建并返回一个用于mssql的底层sql.DB对象。
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
// DoFilter deals with the sql string before commits it to underlying sql driver.
<原文结束>

# <翻译开始>
// DoFilter 在将 SQL 字符串提交给底层 SQL 驱动程序之前，对其进行处理。
# <翻译结束>


<原文开始>
// Convert placeholder char '?' to string "@px".
<原文结束>

# <翻译开始>
// 将占位符字符 '?' 转换为字符串 "@px"。
# <翻译结束>


<原文开始>
// parseSql does some replacement of the sql before commits it to underlying driver,
// for support of microsoft sql server.
<原文结束>

# <翻译开始>
// parseSql在将SQL提交给底层驱动程序之前执行一些替换操作，
// 以便支持Microsoft SQL Server。
# <翻译结束>


<原文开始>
// SELECT * FROM USER WHERE ID=1 LIMIT 1
<原文结束>

# <翻译开始>
// 从USER表中选取ID为1的记录，限制返回结果数量为1条
# <翻译结束>


<原文开始>
// SELECT * FROM USER WHERE AGE>18 ORDER BY ID DESC LIMIT 100, 200
<原文结束>

# <翻译开始>
// 从USER表中选取AGE大于18的所有列，并按ID降序排列，然后获取第101至300条记录（LIMIT offset, count语法）
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
// TableFields retrieves and returns the fields' information of specified table of current schema.
//
// Also see DriverMysql.TableFields.
<原文结束>

# <翻译开始>
// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 另请参阅 DriverMysql.TableFields。
# <翻译结束>


<原文开始>
// DoInsert inserts or updates data forF given table.
<原文结束>

# <翻译开始>
// DoInsert 对给定表执行插入或更新数据操作。
# <翻译结束>


<原文开始>
// LIMIT statement checks.
<原文结束>

# <翻译开始>
// LIMIT语句检查。
# <翻译结束>


<原文开始>
// ORDER BY statement checks.
<原文结束>

# <翻译开始>
// ORDER BY 语句检查。
# <翻译结束>

