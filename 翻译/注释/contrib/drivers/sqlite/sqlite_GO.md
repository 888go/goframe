
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
// Package sqlite implements gdb.Driver, which supports operations for database SQLite.
//
// Note:
// 1. It does not support Save features.
<原文结束>

# <翻译开始>
// Package sqlite 实现了 gdb.Driver 接口，该接口支持对 SQLite 数据库的操作。
//
// 注意：
// 1. 此包不支持 Save 功能。
# <翻译结束>


<原文开始>
// Driver is the driver for sqlite database.
<原文结束>

# <翻译开始>
// Driver 是用于 SQLite 数据库的驱动程序。
# <翻译结束>


<原文开始>
// New create and returns a driver that implements gdb.Driver, which supports operations for SQLite.
<原文结束>

# <翻译开始>
// New 创建并返回一个实现了 gdb.Driver 接口的驱动器，该驱动器支持对 SQLite 的操作。
# <翻译结束>


<原文开始>
// New creates and returns a database object for sqlite.
// It implements the interface of gdb.Driver for extra database driver installation.
<原文结束>

# <翻译开始>
// New 创建并返回一个用于 sqlite 的数据库对象。
// 它实现了 gdb.Driver 接口，以便进行额外的数据库驱动安装。
# <翻译结束>


<原文开始>
// Open creates and returns an underlying sql.DB object for sqlite.
// https://github.com/glebarez/go-sqlite
<原文结束>

# <翻译开始>
// Open创建并返回一个用于sqlite的底层sql.DB对象。
// 参考链接：https://github.com/glebarez/go-sqlite
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
// It searches the source file to locate its absolute path..
<原文结束>

# <翻译开始>
// 它在源文件中搜索以定位其绝对路径。
# <翻译结束>


<原文开始>
	// Multiple PRAGMAs can be specified, e.g.:
	// path/to/some.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)
<原文结束>

# <翻译开始>
// 多个PRAGMA指令可以通过如下方式指定：
// path/to/some.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)
// （翻译成中文）
// 可以通过以下方式同时指定多个PRAGMA参数：
// path/to/some.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)
// 其中，"busy_timeout"设置为5000毫秒，"journal_mode"设置为WAL模式。
# <翻译结束>


<原文开始>
// GetChars returns the security char for this type of database.
<原文结束>

# <翻译开始>
// GetChars 返回此类型数据库的安全字符。
# <翻译结束>


<原文开始>
// DoFilter deals with the sql string before commits it to underlying sql driver.
<原文结束>

# <翻译开始>
// DoFilter 在将SQL字符串提交给底层SQL驱动程序之前对其进行处理。
# <翻译结束>


<原文开始>
// Special insert/ignore operation for sqlite.
<原文结束>

# <翻译开始>
// 特殊的插入/忽略操作，用于SQLite.
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
// TableFields retrieves and returns the fields' information of specified table of current schema.
//
// Also see DriverMysql.TableFields.
<原文结束>

# <翻译开始>
// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 另请参阅 DriverMysql.TableFields。
# <翻译结束>

