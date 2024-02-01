
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
// Package mysql implements gdb.Driver, which supports operations for database MySQL.
<原文结束>

# <翻译开始>
// Package mysql 实现了 gdb.Driver 接口，该接口支持对 MySQL 数据库的相关操作。
# <翻译结束>


<原文开始>
// Driver is the driver for mysql database.
<原文结束>

# <翻译开始>
// Driver 是 MySQL 数据库的驱动程序。
# <翻译结束>


<原文开始>
// New create and returns a driver that implements gdb.Driver, which supports operations for MySQL.
<原文结束>

# <翻译开始>
// New 创建并返回一个实现 gdb.Driver 接口的驱动程序，该驱动程序支持针对 MySQL 的操作。
# <翻译结束>


<原文开始>
// New creates and returns a database object for mysql.
// It implements the interface of gdb.Driver for extra database driver installation.
<原文结束>

# <翻译开始>
// New 创建并返回一个用于 mysql 的数据库对象。
// 它实现了 gdb.Driver 接口，以便进行额外的数据库驱动安装。
# <翻译结束>


<原文开始>
// Open creates and returns an underlying sql.DB object for mysql.
// Note that it converts time.Time argument to local timezone in default.
<原文结束>

# <翻译开始>
// Open 创建并返回一个用于 mysql 的底层 sql.DB 对象。
// 注意，它默认会将 time.Time 类型参数转换为本地时区。
# <翻译结束>


<原文开始>
// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
<原文结束>

# <翻译开始>
// [用户名[:密码]@][协议[(地址)]]/数据库名[?参数1=值1&...&参数N=值N]
// 这段注释是对Go语言中符合MySQL连接格式的字符串进行描述，具体含义如下：
// - `[username[:password]@]`：可选的用户名和密码部分，用于登录数据库。冒号（:）分隔用户名和密码。
// - `[protocol[(address)]]`：指定数据库连接协议以及服务器地址，例如 `tcp(` 或 `unix(` 等，其中括号内的 `address` 为服务器地址或socket路径。
// - `/dbname`：必填项，表示要连接的数据库名称。
// - `[?param1=value1&...&paramN=valueN]`：可选的查询参数部分，通常用于设置额外的连接选项，如 `charset=utf8mb4`、`parseTime=true` 等，多个参数之间用 `&` 符号分隔。
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
// 自定义在运行时更改架构
# <翻译结束>


<原文开始>
// TODO: Do not set charset when charset is not specified (in v2.5.0)
<原文结束>

# <翻译开始>
// TODO: 当未指定字符集时（在v2.5.0版本中），不要设置字符集
# <翻译结束>


<原文开始>
// GetChars returns the security char for this type of database.
<原文结束>

# <翻译开始>
// GetChars 返回此类型数据库的安全字符。
# <翻译结束>


<原文开始>
// DoFilter handles the sql before posts it to database.
<原文结束>

# <翻译开始>
// DoFilter 在将 SQL 发送给数据库之前处理 SQL。
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
// As a map is unsorted, the TableField struct has a "Index" field marks its sequence in
// the fields.
//
// It's using cache feature to enhance the performance, which is never expired util the
// process restarts.
<原文结束>

# <翻译开始>
// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 参数`link`是可选的，如果给定为nil，它会自动获取一个原始sql连接作为链接执行必要的sql查询。
//
// 注意，它返回一个包含字段名及其对应字段信息的map。由于map是无序的，TableField结构体中有一个"Index"字段标记其在所有字段中的顺序。
//
// 为了提高性能，该方法使用了缓存功能，缓存有效期直到进程重启才会失效。
# <翻译结束>

