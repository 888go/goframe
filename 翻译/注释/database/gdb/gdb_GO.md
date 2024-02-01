
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
// Package gdb provides ORM features for popular relationship databases.
<原文结束>

# <翻译开始>
// Package gdb 提供针对主流关系型数据库的 ORM 功能。
# <翻译结束>


<原文开始>
// DB defines the interfaces for ORM operations.
<原文结束>

# <翻译开始>
// DB 定义了用于ORM操作的接口。
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Model creation.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 模型创建。
// ===========================================================================
// 这段注释是用Go语言编写的，其内容是对代码段的描述。翻译成中文后，其含义如下：
// ===========================================================================
// 模块创建部分。
// ===========================================================================
// 这里“Model creation”译为“模型创建”，表示接下来的代码将用于创建某种程序模型或数据模型。
# <翻译结束>


<原文开始>
	// Model creates and returns a new ORM model from given schema.
	// The parameter `table` can be more than one table names, and also alias name, like:
	// 1. Model names:
	//    Model("user")
	//    Model("user u")
	//    Model("user, user_detail")
	//    Model("user u, user_detail ud")
	// 2. Model name with alias: Model("user", "u")
	// Also see Core.Model.
<原文结束>

# <翻译开始>
// Model 根据给定的模式创建并返回一个新的 ORM 模型。
// 参数 `table` 可以是多个表名，也可以包含别名，例如：
// 1. 表名示例：
//    Model("user") // 用户表
//    Model("user u") // 用户表，并为表设置别名 "u"
//    Model("user, user_detail") // 用户表和用户详情表
//    Model("user u, user_detail ud") // 用户表（别名 u）和用户详情表（别名 ud）
// 2. 带有别名的表名：Model("user", "u")
// 有关更多信息，请参阅 Core.Model。
# <翻译结束>


<原文开始>
// Raw creates and returns a model based on a raw sql not a table.
<原文结束>

# <翻译开始>
// Raw 创建并返回一个基于原始SQL（非表）的模型。
// 通常用于嵌入原始sql语句,如:
// g.Model("user").WhereLT("created_at", gdb.Raw("now()")).All()  // SELECT * FROM `user` WHERE `created_at`<now()
// 参考文档:https://goframe.org/pages/viewpage.action?pageId=111911590&showComments=true
# <翻译结束>


<原文开始>
	// Schema creates and returns a schema.
	// Also see Core.Schema.
<原文结束>

# <翻译开始>
// Schema 创建并返回一个模式（Schema）。
// 另请参阅 Core.Schema。
# <翻译结束>


<原文开始>
	// With creates and returns an ORM model based on metadata of given object.
	// Also see Core.With.
<原文结束>

# <翻译开始>
// With 根据给定对象的元数据创建并返回一个 ORM 模型。
// 也可以参考 Core.With。
# <翻译结束>


<原文开始>
	// Open creates a raw connection object for database with given node configuration.
	// Note that it is not recommended using the function manually.
	// Also see DriverMysql.Open.
<原文结束>

# <翻译开始>
// Open 通过给定的节点配置为数据库创建一个原始连接对象。
// 注意，不建议手动使用此函数。
// 另请参阅 DriverMysql.Open。
# <翻译结束>


<原文开始>
	// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
	// of current DB object and with given context in it.
	// Also see Core.Ctx.
<原文结束>

# <翻译开始>
// Ctx 是一个链式函数，它创建并返回一个新的 DB 对象，该对象是对当前 DB 对象的浅复制，并且其中包含给定的上下文。
// 也可参考 Core.Ctx。
# <翻译结束>


<原文开始>
	// Close closes the database and prevents new queries from starting.
	// Close then waits for all queries that have started processing on the server
	// to finish.
	//
	// It is rare to Close a DB, as the DB handle is meant to be
	// long-lived and shared between many goroutines.
<原文结束>

# <翻译开始>
// Close 关闭数据库并阻止新的查询开始。
// Close 之后会等待所有已在服务器上开始处理的查询完成。
//
// 关闭 DB 是罕见的操作，因为 DB 连接句柄设计意图是长期存在且被多个 goroutine 共享。
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Query APIs.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 查询API。
// ===========================================================================
// 这段注释是用英文书写的，翻译成中文后，其内容如下：
// ===========================================================================
// 查询相关的API接口。
// ===========================================================================
// 这里对代码段进行了概括性注释，表明该部分包含查询相关的API（应用程序接口）功能。
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Common APIs for CURD.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 常用的CURD API.
// ===========================================================================
// 这段注释是Go语言代码的一部分，用于描述该部分代码的功能。翻译成中文后，其含义如下：
// ===========================================================================
// 提供常用的创建（Create）、更新（Update）、读取（Read）和删除（Delete）操作的API。
// ===========================================================================
# <翻译结束>












<原文开始>
	// ===========================================================================
	// Internal APIs for CURD, which can be overwritten by custom CURD implements.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 内部CURD API，可以被自定义的CURD实现覆盖。
// ===========================================================================
# <翻译结束>





































<原文开始>
	// ===========================================================================
	// Query APIs for convenience purpose.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 为了方便起见，提供的查询APIs。
// ===========================================================================
# <翻译结束>






















<原文开始>
	// ===========================================================================
	// Master/Slave specification support.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 主从模式支持。
// ===========================================================================
// 这段注释是用于描述Go语言代码中关于主从（Master/Slave）规范或模式的相关实现。主从模式通常是指在分布式系统中，存在一个主节点负责处理写入操作以及数据同步，而从节点则主要用于读取操作和备份数据的场景。
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Ping-Pong.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 乒乓球.
// ===========================================================================
// 这段 Go 语言代码的注释表明这是一个关于“Ping-Pong”的模块或功能，但没有提供具体的代码实现细节。这里的注释翻译成中文后，其含义不变，仍然是对这一部分功能或模块的描述，表示与乒乓球游戏或者网络中的 Ping-Pong（心跳检测）机制相关的代码。
# <翻译结束>












<原文开始>
	// ===========================================================================
	// Transaction.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 事务。
// ===========================================================================
# <翻译结束>







<原文开始>
	// ===========================================================================
	// Configuration methods.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 配置方法。
// ===========================================================================
// 这段 Go 语言代码注释表明接下来的代码是关于配置相关的方法，用于对程序或服务进行配置。
# <翻译结束>

























































<原文开始>
// See Core.SetMaxIdleConnCount.
<原文结束>

# <翻译开始>
// 参见 Core.SetMaxIdleConnCount.
# <翻译结束>


<原文开始>
// See Core.SetMaxOpenConnCount.
<原文结束>

# <翻译开始>
// 参见 Core.SetMaxOpenConnCount.
# <翻译结束>


<原文开始>
// See Core.SetMaxConnLifeTime.
<原文结束>

# <翻译开始>
// 参见 Core.SetMaxConnLifeTime.
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Utility methods.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 工具方法。
// ===========================================================================
// 这段注释是用英语书写的，以下是翻译成中文的版本：
// ===========================================================================
// 实用工具函数集合。
// ===========================================================================
// 这里是对Golang代码中一组工具方法的注释描述，表示这一部分包含一些通用、便捷的辅助函数。
# <翻译结束>







<原文开始>
// See Core.Tables. The driver must implement this function.
<原文结束>

# <翻译开始>
// 查看Core.Tables。驱动程序必须实现这个函数。
# <翻译结束>


<原文开始>
// See Core.TableFields. The driver must implement this function.
<原文结束>

# <翻译开始>
// 查看 Core.TableFields。驱动程序必须实现此函数。
# <翻译结束>


<原文开始>
// See Core.ConvertValueForField
<原文结束>

# <翻译开始>
// 查看 Core.ConvertValueForField
# <翻译结束>


<原文开始>
// See Core.ConvertValueForLocal
<原文结束>

# <翻译开始>
// 参见 Core.ConvertValueForLocal
# <翻译结束>


<原文开始>
// See Core.CheckLocalTypeForField
<原文结束>

# <翻译开始>
// 查看 Core.CheckLocalTypeForField
# <翻译结束>


<原文开始>
// TX defines the interfaces for ORM transaction operations.
<原文结束>

# <翻译开始>
// TX 定义了用于ORM事务操作的接口。
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Nested transaction if necessary.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 如果有必要，进行嵌套事务。
// ===========================================================================
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Core method.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 核心方法
// ===========================================================================
// 这段代码中的注释是将英文注释翻译成中文，其含义如下：
// ```go
// ===========================================================================
// 核心方法
// ===========================================================================
// 这个注释是对接下来要定义或实现的 Go 语言代码段的描述，表示这部分代码是整个程序或模块的核心功能方法。
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Query.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 查询。
// ===========================================================================
# <翻译结束>


<原文开始>
	// ===========================================================================
	// CURD.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// CURD（增删改查操作）
// ===========================================================================
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Save point feature.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 保存点功能。
// ===========================================================================
# <翻译结束>


<原文开始>
// Core is the base struct for database management.
<原文结束>

# <翻译开始>
// Core是数据库管理的基础结构体。
# <翻译结束>







<原文开始>
// Context for chaining operation only. Do not set a default value in Core initialization.
<原文结束>

# <翻译开始>
// 此上下文仅用于链式操作。请勿在 Core 初始化时设置默认值。
# <翻译结束>







<原文开始>
// Custom schema for this object.
<原文结束>

# <翻译开始>
// 为此对象定制的自定义模式。
# <翻译结束>


<原文开始>
// Enable debug mode for the database, which can be changed in runtime.
<原文结束>

# <翻译开始>
// 启用数据库的调试模式，该模式可以在运行时进行更改。
# <翻译结束>


<原文开始>
// Cache manager, SQL result cache only.
<原文结束>

# <翻译开始>
// 缓存管理器，仅用于SQL结果缓存。
# <翻译结束>


<原文开始>
// links caches all created links by node.
<原文结束>

# <翻译开始>
// links 缓存了所有已创建的按节点链接。
# <翻译结束>


<原文开始>
// Logger for logging functionality.
<原文结束>

# <翻译开始>
// Logger 用于提供日志记录功能。
# <翻译结束>







<原文开始>
// Dynamic configurations, which can be changed in runtime.
<原文结束>

# <翻译开始>
// 动态配置，可以在运行时进行更改。
# <翻译结束>


<原文开始>
// DoCommitInput is the input parameters for function DoCommit.
<原文结束>

# <翻译开始>
// DoCommitInput 是函数 DoCommit 的输入参数。
# <翻译结束>


<原文开始>
// DoCommitOutput is the output parameters for function DoCommit.
<原文结束>

# <翻译开始>
// DoCommitOutput 是函数 DoCommit 的输出参数。
# <翻译结束>


<原文开始>
// Result is the result of exec statement.
<原文结束>

# <翻译开始>
// Result 是执行语句的结果。
# <翻译结束>


<原文开始>
// Records is the result of query statement.
<原文结束>

# <翻译开始>
// Records 是查询语句的结果。
# <翻译结束>


<原文开始>
// Stmt is the Statement object result for Prepare.
<原文结束>

# <翻译开始>
// Stmt是Prepare方法执行后返回的Statement对象结果。
# <翻译结束>


<原文开始>
// Tx is the transaction object result for Begin.
<原文结束>

# <翻译开始>
// Tx是Begin方法返回的事务对象。
# <翻译结束>


<原文开始>
// RawResult is the underlying result, which might be sql.Result/*sql.Rows/*sql.Row.
<原文结束>

# <翻译开始>
// RawResult 是底层结果，它可能是 sql.Result、*sql.Rows 或 *sql.Row。
# <翻译结束>


<原文开始>
// Driver is the interface for integrating sql drivers into package gdb.
<原文结束>

# <翻译开始>
// Driver 是用于将 SQL 驱动程序集成到 gdb 包的接口。
# <翻译结束>


<原文开始>
// New creates and returns a database object for specified database server.
<原文结束>

# <翻译开始>
// New 创建并返回指定数据库服务器的数据库对象。
# <翻译结束>


<原文开始>
// Link is a common database function wrapper interface.
// Note that, any operation using `Link` will have no SQL logging.
<原文结束>

# <翻译开始>
// Link 是一个通用的数据库函数包装器接口。
// 注意，使用 `Link` 进行的任何操作将不会有 SQL 日志记录。
# <翻译结束>


<原文开始>
// Sql is the sql recording struct.
<原文结束>

# <翻译开始>
// Sql 是用于记录 SQL 的结构体。
# <翻译结束>












<原文开始>
// Formatted sql which contains arguments in the sql.
<原文结束>

# <翻译开始>
// 格式化后的SQL，其中包含在SQL中的参数。
# <翻译结束>


<原文开始>
// Start execution timestamp in milliseconds.
<原文结束>

# <翻译开始>
// 开始执行的时间戳（毫秒）。
# <翻译结束>


<原文开始>
// End execution timestamp in milliseconds.
<原文结束>

# <翻译开始>
// 结束执行的时间戳（毫秒）。
# <翻译结束>


<原文开始>
// Group is the group name of the configuration that the sql is executed from.
<原文结束>

# <翻译开始>
// Group 是执行 SQL 时所使用的配置组名称。
# <翻译结束>


<原文开始>
// Schema is the schema name of the configuration that the sql is executed from.
<原文结束>

# <翻译开始>
// Schema 是执行 SQL 的配置的架构名称。
# <翻译结束>


<原文开始>
// IsTransaction marks whether this sql is executed in transaction.
<原文结束>

# <翻译开始>
// IsTransaction 标记了这个SQL语句是否在事务中执行。
# <翻译结束>


<原文开始>
// RowsAffected marks retrieved or affected number with current sql statement.
<原文结束>

# <翻译开始>
// RowsAffected 标记了当前SQL语句执行后获取或影响的行数。
# <翻译结束>


<原文开始>
// DoInsertOption is the input struct for function DoInsert.
<原文结束>

# <翻译开始>
// DoInsertOption 是函数 DoInsert 的输入结构体。
# <翻译结束>


<原文开始>
// Insert operation in constant value.
<原文结束>

# <翻译开始>
// 在常数值中执行插入操作。
# <翻译结束>


<原文开始>
// Batch count for batch inserting.
<原文结束>

# <翻译开始>
// 批量插入的批次数量
# <翻译结束>


<原文开始>
// TableField is the struct for table field.
<原文结束>

# <翻译开始>
// TableField 是用于表示表格字段的结构体。
# <翻译结束>


<原文开始>
// For ordering purpose as map is unordered.
<原文结束>

# <翻译开始>
// 用于排序目的，因为映射（map）是无序的。
# <翻译结束>







<原文开始>
// Default value for the field.
<原文结束>

# <翻译开始>
// 字段的默认值。
# <翻译结束>


<原文开始>
// Extra information. Eg: auto_increment.
<原文结束>

# <翻译开始>
// 额外信息。例如：自动增长。
# <翻译结束>


<原文开始>
// Counter  is the type for update count.
<原文结束>

# <翻译开始>
// Counter 是用于更新计数的类型。
# <翻译结束>


<原文开始>
// Raw is a raw sql that will not be treated as argument but as a direct sql part.
<原文结束>

# <翻译开始>
// Raw 是一个原始SQL语句，它不会被视为参数处理，而是直接作为SQL部分。
// 通常用于嵌入原始sql语句,如:
// g.Model("user").WhereLT("created_at", gdb.Raw("now()")).All()  // SELECT * FROM `user` WHERE `created_at`<now()
// 参考文档:https://goframe.org/pages/viewpage.action?pageId=111911590&showComments=true
# <翻译结束>


<原文开始>
// Value is the field value type.
<原文结束>

# <翻译开始>
// Value 是字段值类型。
# <翻译结束>


<原文开始>
// Record is the row record of the table.
<原文结束>

# <翻译开始>
// Record 是表格的行记录。
# <翻译结束>


<原文开始>
// Result is the row record array.
<原文结束>

# <翻译开始>
// Result 是行记录数组。
# <翻译结束>


<原文开始>
// Map is alias of map[string]interface{}, which is the most common usage map type.
<原文结束>

# <翻译开始>
// Map 是 map[string]interface{} 的别名，这是最常用的映射类型。
# <翻译结束>







<原文开始>
// DoCommit marks it will be committed to underlying driver or not.
<原文结束>

# <翻译开始>
// DoCommit 标记是否将提交到底层驱动。
# <翻译结束>


<原文开始>
// Max idle connection count in pool.
<原文结束>

# <翻译开始>
// 在连接池中的最大空闲连接数。
# <翻译结束>


<原文开始>
// Max open connection count in pool. Default is no limit.
<原文结束>

# <翻译开始>
// 在连接池中的最大打开连接数。默认是没有限制。
# <翻译结束>


<原文开始>
// Max lifetime for per connection in pool in seconds.
<原文结束>

# <翻译开始>
// 连接池中每个连接的最大生命周期，单位为秒。
# <翻译结束>


<原文开始>
// type:[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
<原文结束>

# <翻译开始>
// 定义数据库连接格式：
// type: [用户名[:密码]@][协议[(地址)]]/数据库名[?参数1=值1&...&参数N=值N]
// 其中各部分的含义：
// - type：表示数据库类型，如mysql、postgres等
// - username: 可选，用于登录数据库的用户名
// - password: 可选，对应用户名的密码，通常会进行编码处理
// - protocol: 数据库访问协议，如tcp、unix等
// - address: 协议对应的服务器地址或socket路径
// - dbname: 需要连接的数据库名称
// - param1=value1,...,paramN=valueN: 可选，一系列键值对形式的连接参数，例如charset=utf8、sslmode=disable等
# <翻译结束>


<原文开始>
// instances is the management map for instances.
<原文结束>

# <翻译开始>
// instances 是用于实例管理的映射（map）。
# <翻译结束>


<原文开始>
// driverMap manages all custom registered driver.
<原文结束>

# <翻译开始>
// driverMap 管理所有已注册的自定义驱动。
# <翻译结束>


<原文开始>
	// lastOperatorRegPattern is the regular expression pattern for a string
	// which has operator at its tail.
<原文结束>

# <翻译开始>
// lastOperatorRegPattern 是正则表达式模式，用于表示字符串尾部包含操作符的字符串。
# <翻译结束>


<原文开始>
	// regularFieldNameRegPattern is the regular expression pattern for a string
	// which is a regular field name of table.
<原文结束>

# <翻译开始>
// regularFieldNameRegPattern 是用于表示表中常规字段名的字符串的正则表达式模式。
# <翻译结束>


<原文开始>
	// regularFieldNameWithoutDotRegPattern is similar to regularFieldNameRegPattern but not allows '.'.
	// Note that, although some databases allow char '.' in the field name, but it here does not allow '.'
	// in the field name as it conflicts with "db.table.field" pattern in SOME situations.
<原文结束>

# <翻译开始>
// regularFieldNameWithoutDotRegPattern 与 regularFieldNameRegPattern 类似，但不允许包含“.”字符。
// 注意，尽管某些数据库允许字段名中包含字符'.'，但在这里不允许可字段名中出现'.'，因为在某些情况下它会与 "db.table.field" 的模式产生冲突。
# <翻译结束>


<原文开始>
	// allDryRun sets dry-run feature for all database connections.
	// It is commonly used for command options for convenience.
<原文结束>

# <翻译开始>
// allDryRun 为所有数据库连接设置模拟执行（dry-run）功能。
// 通常为了方便，它被用于命令行选项中。
# <翻译结束>


<原文开始>
// tableFieldsMap caches the table information retrieved from database.
<原文结束>

# <翻译开始>
// tableFieldsMap 缓存从数据库获取的表信息。
# <翻译结束>


<原文开始>
// allDryRun is initialized from environment or command options.
<原文结束>

# <翻译开始>
// allDryRun 从环境变量或命令选项中初始化。
# <翻译结束>


<原文开始>
// Register registers custom database driver to gdb.
<原文结束>

# <翻译开始>
// Register 注册自定义数据库驱动到 gdb。
# <翻译结束>


<原文开始>
// New creates and returns an ORM object with given configuration node.
<原文结束>

# <翻译开始>
// New 根据给定的配置节点创建并返回一个ORM对象。
# <翻译结束>


<原文开始>
// NewByGroup creates and returns an ORM object with global configurations.
// The parameter `name` specifies the configuration group name,
// which is DefaultGroupName in default.
<原文结束>

# <翻译开始>
// NewByGroup 根据全局配置创建并返回一个 ORM 对象。
// 参数 `name` 指定配置组名称，默认为 DefaultGroupName。
# <翻译结束>


<原文开始>
// newDBByConfigNode creates and returns an ORM object with given configuration node and group name.
//
// Very Note:
// The parameter `node` is used for DB creation, not for underlying connection creation.
// So all db type configurations in the same group should be the same.
<原文结束>

# <翻译开始>
// newDBByConfigNode 根据给定的配置节点和组名创建并返回一个ORM对象。
//
// **非常重要**：
// 参数`node`用于数据库创建，而非底层连接创建，
// 因此在同一组内的所有数据库类型配置应当保持一致。
# <翻译结束>


<原文开始>
// Instance returns an instance for DB operations.
// The parameter `name` specifies the configuration group name,
// which is DefaultGroupName in default.
<原文结束>

# <翻译开始>
// Instance 返回一个用于数据库操作的实例。
// 参数 `name` 指定了配置组名称，默认为 DefaultGroupName。
# <翻译结束>


<原文开始>
// getConfigNodeByGroup calculates and returns a configuration node of given group. It
// calculates the value internally using weight algorithm for load balance.
//
// The parameter `master` specifies whether retrieving a master node, or else a slave node
// if master-slave configured.
<原文结束>

# <翻译开始>
// 根据给定组获取配置节点，通过内部计算并返回。它使用权重算法进行负载均衡计算。
//
// 参数`master`指定是否获取主节点，如果不是，则在主从配置情况下获取从节点。
# <翻译结束>


<原文开始>
// Separates master and slave configuration nodes array.
<原文结束>

# <翻译开始>
// 将主节点和从节点配置数组分离。
# <翻译结束>


<原文开始>
// getConfigNodeByWeight calculates the configuration weights and randomly returns a node.
//
// Calculation algorithm brief:
// 1. If we have 2 nodes, and their weights are both 1, then the weight range is [0, 199];
// 2. Node1 weight range is [0, 99], and node2 weight range is [100, 199], ratio is 1:1;
// 3. If the random number is 99, it then chooses and returns node1;.
<原文结束>

# <翻译开始>
// getConfigNodeByWeight 根据权重计算配置，并随机返回一个节点。
//
// 计算算法简述：
// 1. 若我们有2个节点，它们的权重都为1，则权重范围是 [0, 199]；
// 2. 节点1的权重范围是 [0, 99]，节点2的权重范围是 [100, 199]，比例为 1:1；
// 3. 如果随机数是99，则选择并返回节点1。
# <翻译结束>


<原文开始>
	// If total is 0 means all the nodes have no weight attribute configured.
	// It then defaults each node's weight attribute to 1.
<原文结束>

# <翻译开始>
// 如果total为0，表示所有节点都没有配置权重属性。
// 此时，默认将每个节点的权重属性设为1。
# <翻译结束>


<原文开始>
// Exclude the right border value.
<原文结束>

# <翻译开始>
// 排除右侧边界值。
# <翻译结束>


<原文开始>
			// ====================================================
			// Return a COPY of the ConfigNode.
			// ====================================================
<原文结束>

# <翻译开始>
// ====================================================
// 返回ConfigNode的一个副本。
// ====================================================
# <翻译结束>


<原文开始>
// getSqlDb retrieves and returns an underlying database connection object.
// The parameter `master` specifies whether retrieves master node connection if
// master-slave nodes are configured.
<原文结束>

# <翻译开始>
// getSqlDb 获取并返回底层数据库连接对象。
// 参数 `master` 指定在配置了主从节点的情况下，是否获取主节点的连接。
# <翻译结束>












<原文开始>
// Update the configuration object in internal data.
<原文结束>

# <翻译开始>
// 更新内部数据中的配置对象。
# <翻译结束>


<原文开始>
// Cache the underlying connection pool object by node.
<原文结束>

# <翻译开始>
// 通过节点缓存底层连接池对象。
# <翻译结束>


<原文开始>
// It reads from instance map.
<原文结束>

# <翻译开始>
// 它从实例映射中读取。
# <翻译结束>























































































<原文开始>
// SQL string(may contain reserved char '?').
<原文结束>

# <翻译开始>
// SQL 字符串（可能包含保留字符 '?'）。
# <翻译结束>







<原文开始>
// Custom string for `on duplicated` statement.
<原文结束>

# <翻译开始>
// 自定义用于`on duplicated`语句的字符串。
# <翻译结束>


<原文开始>
// Custom key-value map from `OnDuplicateEx` function for `on duplicated` statement.
<原文结束>

# <翻译开始>
// `OnDuplicateEx`函数为`on duplicated`语句提供的自定义键值对映射
# <翻译结束>


<原文开始>
// Field type. Eg: 'int(10) unsigned', 'varchar(64)'.
<原文结束>

# <翻译开始>
// 字段类型。例如：'int(10) unsigned', 'varchar(64)'。
// 这段注释是对Go语言代码中某个表示字段类型的变量或常量的解释，该字段在数据库表结构设计中使用，比如MySQL等关系型数据库。'int(10) unsigned' 表示一个无符号整数类型，长度为10位；'varchar(64)' 则表示变长字符串类型，最大长度为64个字符。
# <翻译结束>


<原文开始>
// See Core.InsertIgnore.
<原文结束>

# <翻译开始>
// 参见 Core.InsertIgnore。
# <翻译结束>


<原文开始>
// See Core.InsertAndGetId.
<原文结束>

# <翻译开始>
// 参见 Core.InsertAndGetId.
# <翻译结束>


<原文开始>
// See Core.DoSelect.
<原文结束>

# <翻译开始>
// 参见 Core.DoSelect。
# <翻译结束>


<原文开始>
// See Core.DoInsert.
<原文结束>

# <翻译开始>
// 参见 Core.DoInsert。
# <翻译结束>


<原文开始>
// See Core.DoUpdate.
<原文结束>

# <翻译开始>
// 参见 Core.DoUpdate.
# <翻译结束>


<原文开始>
// See Core.DoDelete.
<原文结束>

# <翻译开始>
// 参见 Core.DoDelete。
# <翻译结束>


<原文开始>
// See Core.DoFilter.
<原文结束>

# <翻译开始>
// 参见 Core.DoFilter。
# <翻译结束>


<原文开始>
// See Core.DoCommit.
<原文结束>

# <翻译开始>
// 参见 Core.DoCommit。
# <翻译结束>


<原文开始>
// See Core.DoPrepare.
<原文结束>

# <翻译开始>
// 参见 Core.DoPrepare。
# <翻译结束>


<原文开始>
// See Core.GetValue.
<原文结束>

# <翻译开始>
// 参见 Core.GetValue。
# <翻译结束>


<原文开始>
// See Core.GetArray.
<原文结束>

# <翻译开始>
// 参见 Core.GetArray.
# <翻译结束>


<原文开始>
// See Core.GetCount.
<原文结束>

# <翻译开始>
// 参见 Core.GetCount。
# <翻译结束>


<原文开始>
// See Core.UnionAll.
<原文结束>

# <翻译开始>
// 参见 Core.UnionAll。
# <翻译结束>


<原文开始>
// See Core.PingMaster.
<原文结束>

# <翻译开始>
// 参见 Core.PingMaster.
# <翻译结束>


<原文开始>
// See Core.PingSlave.
<原文结束>

# <翻译开始>
// 参见 Core.PingSlave.
# <翻译结束>


<原文开始>
// See Core.Transaction.
<原文结束>

# <翻译开始>
// 参见Core.Transaction.
# <翻译结束>


<原文开始>
// See Core.GetCache.
<原文结束>

# <翻译开始>
// 参见 Core.GetCache。
# <翻译结束>


<原文开始>
// See Core.SetDebug.
<原文结束>

# <翻译开始>
// 参见 Core.SetDebug.
# <翻译结束>


<原文开始>
// See Core.GetDebug.
<原文结束>

# <翻译开始>
// 参见 Core.GetDebug。
# <翻译结束>


<原文开始>
// See Core.GetSchema.
<原文结束>

# <翻译开始>
// 参见 Core.GetSchema。
# <翻译结束>


<原文开始>
// See Core.GetPrefix.
<原文结束>

# <翻译开始>
// 参见 Core.GetPrefix。
# <翻译结束>


<原文开始>
// See Core.GetGroup.
<原文结束>

# <翻译开始>
// 参见 Core.GetGroup。
# <翻译结束>


<原文开始>
// See Core.SetDryRun.
<原文结束>

# <翻译开始>
// 参见 Core.SetDryRun。
# <翻译结束>


<原文开始>
// See Core.GetDryRun.
<原文结束>

# <翻译开始>
// 参见 Core.GetDryRun。
# <翻译结束>


<原文开始>
// See Core.SetLogger.
<原文结束>

# <翻译开始>
// 参见 Core.SetLogger。
# <翻译结束>


<原文开始>
// See Core.GetLogger.
<原文结束>

# <翻译开始>
// 参见 Core.GetLogger。
# <翻译结束>


<原文开始>
// See Core.GetConfig.
<原文结束>

# <翻译开始>
// 参见 Core.GetConfig。
# <翻译结束>


<原文开始>
// See Core.GetChars.
<原文结束>

# <翻译开始>
// 参见 Core.GetChars。
# <翻译结束>


<原文开始>
// DB interface object.
<原文结束>

# <翻译开始>
// DB 接口对象。
# <翻译结束>


<原文开始>
// Configuration group name.
<原文结束>

# <翻译开始>
// 配置组名称。
# <翻译结束>


<原文开始>
// Current config node.
<原文结束>

# <翻译开始>
// 当前配置节点。
# <翻译结束>


<原文开始>
// SQL operation type.
<原文结束>

# <翻译开始>
// SQL操作类型。
# <翻译结束>


<原文开始>
// Arguments for this sql.
<原文结束>

# <翻译开始>
// 此SQL的参数。
# <翻译结束>


<原文开始>
// Field can be null or not.
<原文结束>

# <翻译开始>
// 字段可以为空或非空
# <翻译结束>


<原文开始>
// List is type of map array.
<原文结束>

# <翻译开始>
// List 是映射数组的类型。
# <翻译结束>


<原文开始>
// Value COPY for node.
<原文结束>

# <翻译开始>
// Value COPY 表示节点的复制值。
# <翻译结束>


<原文开始>
// Changes the schema.
<原文结束>

# <翻译开始>
// 修改模式
# <翻译结束>

