
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
// Package gdb provides ORM features for popular relationship databases.
//
// TODO use context.Context as required parameter for all DB operations.
<原文结束>

# <翻译开始>
// 包gdb为流行的关系型数据库提供ORM（对象关系映射）功能。
//
// 待办事项：将context.Context作为所有数据库操作的必需参数。
// md5:ed61b69bd00b7384
# <翻译结束>


<原文开始>
// DB defines the interfaces for ORM operations.
<原文结束>

# <翻译开始>
// DB 定义 ORM 操作的接口。. md5:328f032182d38455
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Model creation.
	// ===========================================================================
<原文结束>

# <翻译开始>
// =============================================================================
// 模型创建。
// =============================================================================
// 这里是对一段Go代码中的注释进行翻译，"Model creation"指的是模型的创建过程。这部分代码可能是用于描述一个函数或部分代码的作用，即它负责构建或初始化某个模型。
// md5:1c8c0b09089a9689
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
// 参数 `table` 可以是多个表名，也可以包括别名，例如：
// 1. 模型名称：
//    Model("user")
//    Model("user u") // u 作为 user 表的别名
//    Model("user, user_detail") // 多个模型名称
//    Model("user u, user_detail ud") // 多个模型名称及别名
// 2. 带别名的模型名称：Model("user", "u") // 第二个参数指定别名
// 参见 Core.Model 了解更多。
// md5:61d3e6d835068122
# <翻译结束>


<原文开始>
// Raw creates and returns a model based on a raw sql not a table.
<原文结束>

# <翻译开始>
// Raw 根据原始SQL（而不是表格）创建并返回一个模型。. md5:96066a9d41296a2a
# <翻译结束>


<原文开始>
	// Schema creates and returns a schema.
	// Also see Core.Schema.
<原文结束>

# <翻译开始>
// Schema 创建并返回一个模式。
// 参见 Core.Schema。
// md5:0f4472ee79f06819
# <翻译结束>


<原文开始>
	// With creates and returns an ORM model based on metadata of given object.
	// Also see Core.With.
<原文结束>

# <翻译开始>
// With根据给定对象的元数据创建并返回一个ORM模型。同时参见Core.With。
// md5:78ab17ce6b00ce6e
# <翻译结束>


<原文开始>
	// Open creates a raw connection object for database with given node configuration.
	// Note that it is not recommended using the function manually.
	// Also see DriverMysql.Open.
<原文结束>

# <翻译开始>
// Open 使用给定的节点配置为数据库创建一个原始连接对象。
// 注意，不建议手动使用此函数。
// 另请参阅 DriverMysql.Open。
// md5:1021f26472df579e
# <翻译结束>


<原文开始>
	// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
	// of current DB object and with given context in it.
	// Also see Core.Ctx.
<原文结束>

# <翻译开始>
// Ctx 是一个链式函数，它创建并返回一个新的 DB，该 DB 是当前 DB 对象的浅拷贝，并在其中设置了给定的上下文。
// 另请参阅 Core.Ctx。
// md5:7eec5fab912764e7
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
// Close 关闭数据库并阻止新的查询开始。然后，Close 等待已经在服务器上开始处理的所有查询完成。
// 
// 通常不会关闭 DB，因为 DB句柄应该是长期存在的，并且在多个 goroutine 之间共享。
// md5:0985fc8e558f83fc
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Query APIs.
	// ===========================================================================
<原文结束>

# <翻译开始>
// =============================================================================
// 查询接口。
// =============================================================================
// 这里是对查询相关的API进行的注释。
// md5:06da8c4c9c8d957b
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Common APIs for CURD.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 用于CURD操作的通用API。
// ===========================================================================
// md5:781fc1b4ac386204
# <翻译结束>


<原文开始>
// See Core.InsertAndGetId.
<原文结束>

# <翻译开始>
// 参见Core.InsertAndGetId。. md5:b7dec69920da6e7a
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Internal APIs for CURD, which can be overwritten by custom CURD implements.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// CURD的内部API，可以被自定义的CURD实现覆盖。
// ===========================================================================
// md5:02480feeb95bda1e
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Query APIs for convenience purpose.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 为了方便起见，提供查询API。
// ===========================================================================
// md5:be53a34b0863cf28
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Master/Slave specification support.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 主从规范支持。
// ===========================================================================
// md5:f0ac82262204c704
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Ping-Pong.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 乒乓游戏。
// ===========================================================================
// md5:548138891df7682f
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Transaction.
	// ===========================================================================
<原文结束>

# <翻译开始>
// =============================================================================
// 事务处理。
// =============================================================================
// 这里是对一个名为 "Transaction" 的部分或函数的注释，表示它与交易操作相关。
// md5:98c80ce4a302c379
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
// md5:e4c7270c61398365
# <翻译结束>


<原文开始>
// See Core.SetMaxIdleConnCount.
<原文结束>

# <翻译开始>
// 参见 Core.SetMaxIdleConnCount。. md5:93a41dde27176210
# <翻译结束>


<原文开始>
// See Core.SetMaxOpenConnCount.
<原文结束>

# <翻译开始>
// 参见Core.SetMaxOpenConnCount。. md5:781ba14245ef2d2f
# <翻译结束>


<原文开始>
// See Core.SetMaxConnLifeTime.
<原文结束>

# <翻译开始>
// 请参考Core.SetMaxConnLifeTime。. md5:9886c404ca6b5919
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Utility methods.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 辅助方法。
// ===========================================================================
// md5:0c5a132a773f89c0
# <翻译结束>


<原文开始>
// See Core.Tables. The driver must implement this function.
<原文结束>

# <翻译开始>
// 参见 Core.Tables。驱动程序必须实现此函数。. md5:d7f231f6b59af607
# <翻译结束>


<原文开始>
// See Core.TableFields. The driver must implement this function.
<原文结束>

# <翻译开始>
// 参见 Core.TableFields。驱动程序必须实现此函数。. md5:657c24bb39017da1
# <翻译结束>


<原文开始>
// See Core.ConvertValueForField
<原文结束>

# <翻译开始>
// 参见Core(ConvertValueForField). md5:cd3e4aabe989b5b0
# <翻译结束>


<原文开始>
// See Core.ConvertValueForLocal
<原文结束>

# <翻译开始>
// 参见Core.ConvertValueForLocal. md5:c5ed8f55d002cc9b
# <翻译结束>


<原文开始>
// See Core.CheckLocalTypeForField
<原文结束>

# <翻译开始>
// 参见 Core.CheckLocalTypeForField. md5:9dab404962da3137
# <翻译结束>


<原文开始>
// See Core.DoFormatUpsert
<原文结束>

# <翻译开始>
// 参见Core.DoFormatUpsert. md5:e28a610aead90684
# <翻译结束>


<原文开始>
// TX defines the interfaces for ORM transaction operations.
<原文结束>

# <翻译开始>
// TX 定义 ORM 事务操作的接口。. md5:d71a7d0434928cac
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Nested transaction if necessary.
	// ===========================================================================
<原文结束>

# <翻译开始>
// 如果需要，嵌套事务。
// ===========================================================================
// md5:96e249df6d75bc7f
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Core method.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 核心方法。
// ===========================================================================
// md5:a10911bb5021107c
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Query.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===========================================================================
// 查询
// ===========================================================================
// md5:4612a1ae72dd3cf5
# <翻译结束>


<原文开始>
	// ===========================================================================
	// CURD.
	// ===========================================================================
<原文结束>

# <翻译开始>
// =============================================================================
// CURD (Create, Read, Update, Delete) 操作。
// =============================================================================
// 这是Go语言中的注释，描述了一个与CRUD（创建、读取、更新、删除）操作相关的部分。在软件开发中，CURD通常用于数据库操作的基本操作。
// md5:b9584d9a2373e908
# <翻译结束>


<原文开始>
	// ===========================================================================
	// Save point feature.
	// ===========================================================================
<原文结束>

# <翻译开始>
// ===================================================================================
// 保存点功能。
// ===================================================================================
// md5:54487b34337e4026
# <翻译结束>


<原文开始>
// StatsItem defines the stats information for a configuration node.
<原文结束>

# <翻译开始>
// StatsItem 定义了配置节点的统计信息。. md5:95acda1876ad44fa
# <翻译结束>


<原文开始>
// Node returns the configuration node info.
<原文结束>

# <翻译开始>
// Node 返回配置节点信息。. md5:868005c0df3fa483
# <翻译结束>


<原文开始>
// Stats returns the connection stat for current node.
<原文结束>

# <翻译开始>
// Stats 返回当前节点的连接状态统计信息。. md5:b497e68c5fce778b
# <翻译结束>


<原文开始>
// Core is the base struct for database management.
<原文结束>

# <翻译开始>
// Core 是数据库管理的基础结构。. md5:9f4e8f0e026a368e
# <翻译结束>


<原文开始>
// Context for chaining operation only. Do not set a default value in Core initialization.
<原文结束>

# <翻译开始>
// 只用于链式操作的上下文。不要在核心初始化时设置默认值。. md5:4d22c5c34f3c1128
# <翻译结束>


<原文开始>
// Configuration group name.
<原文结束>

# <翻译开始>
// 配置组名称。. md5:50dbd64908990986
# <翻译结束>


<原文开始>
// Custom schema for this object.
<原文结束>

# <翻译开始>
// 该对象的自定义架构。. md5:dd54ddfd9c22e232
# <翻译结束>


<原文开始>
// Enable debug mode for the database, which can be changed in runtime.
<原文结束>

# <翻译开始>
// 启用数据库的调试模式，该模式可以在运行时更改。. md5:31b723dcdbabbcb3
# <翻译结束>


<原文开始>
// Cache manager, SQL result cache only.
<原文结束>

# <翻译开始>
// 缓存管理器，仅用于SQL结果缓存。. md5:480937caf34cae3b
# <翻译结束>


<原文开始>
// links caches all created links by node.
<原文结束>

# <翻译开始>
// links 缓存由节点创建的所有链接。. md5:a89aca4e23df0139
# <翻译结束>


<原文开始>
// Logger for logging functionality.
<原文结束>

# <翻译开始>
// 用于记录功能的日志记录器。. md5:ff375387d5036677
# <翻译结束>


<原文开始>
// Dynamic configurations, which can be changed in runtime.
<原文结束>

# <翻译开始>
// 动态配置，这些配置可以在运行时进行更改。. md5:11c382c381ba12fc
# <翻译结束>


<原文开始>
// DoCommitInput is the input parameters for function DoCommit.
<原文结束>

# <翻译开始>
// DoCommitInput是DoCommit函数的输入参数。. md5:151d182ffc05e6f3
# <翻译结束>


<原文开始>
// DoCommitOutput is the output parameters for function DoCommit.
<原文结束>

# <翻译开始>
// DoCommitOutput是DoCommit函数的输出参数。. md5:bb154a9d2f960894
# <翻译结束>


<原文开始>
// Result is the result of exec statement.
<原文结束>

# <翻译开始>
// Result 是执行语句的结果。. md5:92181818237c3bdd
# <翻译结束>


<原文开始>
// Records is the result of query statement.
<原文结束>

# <翻译开始>
// Records 是查询语句的结果。. md5:3ab79979d5bb7a15
# <翻译结束>


<原文开始>
// Stmt is the Statement object result for Prepare.
<原文结束>

# <翻译开始>
// Stmt是Prepare的结果，是一个Statement对象。. md5:f7d8689435820710
# <翻译结束>


<原文开始>
// Tx is the transaction object result for Begin.
<原文结束>

# <翻译开始>
// Tx是Begin操作的结果交易对象。. md5:388468f78948bf40
# <翻译结束>


<原文开始>
// RawResult is the underlying result, which might be sql.Result/*sql.Rows/*sql.Row.
<原文结束>

# <翻译开始>
// RawResult 是底层结果，可能是 sql.Result/*sql.Rows/*sql.Row。. md5:8f6721571bd4ebc3
# <翻译结束>


<原文开始>
// Driver is the interface for integrating sql drivers into package gdb.
<原文结束>

# <翻译开始>
// Driver 是将 sql 驱动程序集成到 gdb 包的接口。. md5:739e8c3911355df2
# <翻译结束>


<原文开始>
// New creates and returns a database object for specified database server.
<原文结束>

# <翻译开始>
// New 为指定的数据库服务器创建并返回一个数据库对象。. md5:27bb5dc9ab2ddbdf
# <翻译结束>


<原文开始>
// Link is a common database function wrapper interface.
// Note that, any operation using `Link` will have no SQL logging.
<原文结束>

# <翻译开始>
// Link 是一个常见的数据库函数包装接口。
// 注意，使用 `Link` 进行的任何操作都不会有 SQL 日志记录。
// md5:d84360a9ae77a1de
# <翻译结束>


<原文开始>
// Sql is the sql recording struct.
<原文结束>

# <翻译开始>
// Sql 是用于记录 SQL 的结构体。. md5:00be6a573786ce1b
# <翻译结束>


<原文开始>
// SQL string(may contain reserved char '?').
<原文结束>

# <翻译开始>
// SQL字符串（可能包含保留字符'?'）。. md5:8c5af3ca9410a41d
# <翻译结束>


<原文开始>
// Arguments for this sql.
<原文结束>

# <翻译开始>
// 这个SQL的参数。. md5:b4c3922763e4978a
# <翻译结束>


<原文开始>
// Formatted sql which contains arguments in the sql.
<原文结束>

# <翻译开始>
// 包含在SQL中的参数格式化的SQL。. md5:555ebca4fe98838b
# <翻译结束>


<原文开始>
// Start execution timestamp in milliseconds.
<原文结束>

# <翻译开始>
// 开始执行的时间戳（以毫秒为单位）。. md5:90bcb553c20da17f
# <翻译结束>


<原文开始>
// End execution timestamp in milliseconds.
<原文结束>

# <翻译开始>
// 结束执行的时间戳，以毫秒为单位。. md5:a69dc8d121bd65c6
# <翻译结束>


<原文开始>
// Group is the group name of the configuration that the sql is executed from.
<原文结束>

# <翻译开始>
// Group是执行SQL的配置组的名称。. md5:ae6f0deafeed211b
# <翻译结束>


<原文开始>
// Schema is the schema name of the configuration that the sql is executed from.
<原文结束>

# <翻译开始>
// Schema是执行SQL时的配置模式名称。. md5:ac146287e6f60f27
# <翻译结束>


<原文开始>
// IsTransaction marks whether this sql is executed in transaction.
<原文结束>

# <翻译开始>
// IsTransaction 标记此 SQL 语句是否在事务中执行。. md5:df029c1ffc72fbf7
# <翻译结束>


<原文开始>
// RowsAffected marks retrieved or affected number with current sql statement.
<原文结束>

# <翻译开始>
// RowsAffected 标记了当前 SQL 语句所影响或检索到的数量。. md5:24de0b6ae028d942
# <翻译结束>


<原文开始>
// DoInsertOption is the input struct for function DoInsert.
<原文结束>

# <翻译开始>
// DoInsertOption是用于DoInsert函数的输入结构体。. md5:de18cdcf6449aa9a
# <翻译结束>


<原文开始>
// Custom string for `on duplicated` statement.
<原文结束>

# <翻译开始>
// 自定义的`on duplicated`语句字符串。. md5:1076eb09195a063c
# <翻译结束>


<原文开始>
// Custom key-value map from `OnDuplicateEx` function for `on duplicated` statement.
<原文结束>

# <翻译开始>
// 从`OnDuplicateEx`函数中自定义的键值映射，用于`on duplicated`语句。. md5:11d4cf3d337093bb
# <翻译结束>


<原文开始>
// Custom conflict key of upsert clause, if the database needs it.
<原文结束>

# <翻译开始>
// 自定义的更新或插入语句中的冲突键，如果数据库需要的话。. md5:9f554ee78fb66a28
# <翻译结束>


<原文开始>
// Insert operation in constant value.
<原文结束>

# <翻译开始>
// 在常量值中插入操作。. md5:e5ca1b47e1d66f7a
# <翻译结束>


<原文开始>
// Batch count for batch inserting.
<原文结束>

# <翻译开始>
// 批量插入的批次计数。. md5:015bd9ee24bd1f5c
# <翻译结束>


<原文开始>
// TableField is the struct for table field.
<原文结束>

# <翻译开始>
// TableField 是用于表示表字段的结构体。. md5:dad00a23ddbc4525
# <翻译结束>


<原文开始>
// For ordering purpose as map is unordered.
<原文结束>

# <翻译开始>
// 用于排序，因为map是无序的。. md5:2c2b51c0f42d0aa5
# <翻译结束>


<原文开始>
// Field type. Eg: 'int(10) unsigned', 'varchar(64)'.
<原文结束>

# <翻译开始>
// 字段类型。例如：'int(10) unsigned'，'varchar(64)'。. md5:c5cb4af28fd84cc4
# <翻译结束>


<原文开始>
// Field can be null or not.
<原文结束>

# <翻译开始>
// 字段可以为null，也可以不为null。. md5:eecc03ab53cc06c9
# <翻译结束>


<原文开始>
// Default value for the field.
<原文结束>

# <翻译开始>
// 字段的默认值。. md5:e9e6c4fce349ba5e
# <翻译结束>


<原文开始>
// Extra information. Eg: auto_increment.
<原文结束>

# <翻译开始>
//额外的信息。例如：auto_increment。. md5:706ba1f1653042e0
# <翻译结束>


<原文开始>
// Counter  is the type for update count.
<原文结束>

# <翻译开始>
// Counter 是更新计数的类型。. md5:d05ba3f2911b8013
# <翻译结束>


<原文开始>
// Raw is a raw sql that will not be treated as argument but as a direct sql part.
<原文结束>

# <翻译开始>
// Raw是一个原始SQL，它不会被当作参数对待，而是直接作为SQL的一部分。. md5:cd876e2c863a974f
# <翻译结束>


<原文开始>
// Value is the field value type.
<原文结束>

# <翻译开始>
// Value 是字段值的类型。. md5:f2d488e083da124a
# <翻译结束>


<原文开始>
// Record is the row record of the table.
<原文结束>

# <翻译开始>
// Record 是表格的行记录。. md5:f13ea17db2caaf51
# <翻译结束>


<原文开始>
// Result is the row record array.
<原文结束>

# <翻译开始>
// Result 是行记录数组。. md5:cfc50f7c3d051e4e
# <翻译结束>


<原文开始>
// Map is alias of map[string]interface{}, which is the most common usage map type.
<原文结束>

# <翻译开始>
// Map是map[string]interface{}的别名，这是最常用的映射类型。. md5:d30ae3cb84b9285e
# <翻译结束>


<原文开始>
// List is type of map array.
<原文结束>

# <翻译开始>
// List 是映射数组的类型。. md5:a6dda10906f1d599
# <翻译结束>


<原文开始>
// DoCommit marks it will be committed to underlying driver or not.
<原文结束>

# <翻译开始>
// DoCommit 标记是否将提交给底层驱动器。. md5:f117644c40f63234
# <翻译结束>


<原文开始>
// Max idle connection count in pool.
<原文结束>

# <翻译开始>
// 连接池中的最大空闲连接数。. md5:2d7d30a1c51e849a
# <翻译结束>


<原文开始>
// Max open connection count in pool. Default is no limit.
<原文结束>

# <翻译开始>
// 连接池中的最大打开连接数。默认无限制。. md5:aef415d0d3363e03
# <翻译结束>


<原文开始>
// Max lifetime for per connection in pool in seconds.
<原文结束>

# <翻译开始>
// 连接池中每个连接的最大生存时间，单位为秒。. md5:5b2ee66fdff9b7f6
# <翻译结束>


<原文开始>
// type:[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
<原文结束>

# <翻译开始>
// 类型: [用户名[:密码]@][协议[(地址)]]/数据库名[?参数1=值1&...&参数N=值N]. md5:ccc4969581f025ac
# <翻译结束>


<原文开始>
// instances is the management map for instances.
<原文结束>

# <翻译开始>
// instances 是实例的管理映射。. md5:4600091cea2428de
# <翻译结束>


<原文开始>
// driverMap manages all custom registered driver.
<原文结束>

# <翻译开始>
// driverMap管理所有自定义注册的驱动程序。. md5:625ff37f5c3fb23d
# <翻译结束>


<原文开始>
	// lastOperatorRegPattern is the regular expression pattern for a string
	// which has operator at its tail.
<原文结束>

# <翻译开始>
// lastOperatorRegPattern 是一个正则表达式模式，用于匹配字符串末尾带有操作符的字符串。
// md5:6a05c1a2b57c687b
# <翻译结束>


<原文开始>
	// regularFieldNameRegPattern is the regular expression pattern for a string
	// which is a regular field name of table.
<原文结束>

# <翻译开始>
// regularFieldNameRegPattern 是一个正则表达式模式，用于匹配表格的普通字段名称的字符串。
// md5:d18bc9e2bf6112ed
# <翻译结束>


<原文开始>
	// regularFieldNameWithCommaRegPattern is the regular expression pattern for one or more strings
	// which are regular field names of table, multiple field names joined with char ','.
<原文结束>

# <翻译开始>
// regularFieldNameWithCommaRegPattern 是用于匹配一个或多个表的常规字段名的正则表达式模式，这些字段名由字符','连接。
// md5:90a0d75039f03540
# <翻译结束>


<原文开始>
	// regularFieldNameWithoutDotRegPattern is similar to regularFieldNameRegPattern but not allows '.'.
	// Note that, although some databases allow char '.' in the field name, but it here does not allow '.'
	// in the field name as it conflicts with "db.table.field" pattern in SOME situations.
<原文结束>

# <翻译开始>
// regularFieldNameWithoutDotRegPattern 与 regularFieldNameRegPattern 类似，但不允许使用点（.）。
// 注意，虽然一些数据库允许字段名中包含字符 '.', 但在某些情况下，这里不允许在字段名中使用 '.'，因为它与 "db.table.field" 的模式冲突。
// md5:4a7a4427aab61aa8
# <翻译结束>


<原文开始>
	// allDryRun sets dry-run feature for all database connections.
	// It is commonly used for command options for convenience.
<原文结束>

# <翻译开始>
// allDryRun 为所有数据库连接设置了 dry-run 特性。
// 它通常用于命令选项，以便于使用时带来方便。
// md5:038bcc87fc3093b6
# <翻译结束>


<原文开始>
// tableFieldsMap caches the table information retrieved from database.
<原文结束>

# <翻译开始>
// tableFieldsMap 缓存从数据库获取的表信息。. md5:5ae26e45c71e9a09
# <翻译结束>


<原文开始>
// allDryRun is initialized from environment or command options.
<原文结束>

# <翻译开始>
// allDryRun 从环境或命令选项中初始化。. md5:1dffa2ad4982da25
# <翻译结束>


<原文开始>
// Register registers custom database driver to gdb.
<原文结束>

# <翻译开始>
// Register 注册自定义数据库驱动到gdb。. md5:d889e7374da12918
# <翻译结束>


<原文开始>
// New creates and returns an ORM object with given configuration node.
<原文结束>

# <翻译开始>
// New 根据给定的配置节点创建并返回一个ORM对象。. md5:c6039a0817062f9e
# <翻译结束>


<原文开始>
// NewByGroup creates and returns an ORM object with global configurations.
// The parameter `name` specifies the configuration group name,
// which is DefaultGroupName in default.
<原文结束>

# <翻译开始>
// NewByGroup 根据指定的配置组名称创建并返回一个ORM对象，带有全局配置。
// 参数`name`指定了配置组的名称，默认为DefaultGroupName。
// md5:a15dd30e999d29e5
# <翻译结束>


<原文开始>
// newDBByConfigNode creates and returns an ORM object with given configuration node and group name.
//
// Very Note:
// The parameter `node` is used for DB creation, not for underlying connection creation.
// So all db type configurations in the same group should be the same.
<原文结束>

# <翻译开始>
// newDBByConfigNode 使用给定的配置节点和组名创建并返回一个ORM对象。
// 
// 非常注意：
// 参数`node`用于数据库的创建，而不是底层连接的创建。因此，同一组中的所有数据库类型配置应该相同。
// md5:b916b78d0af6a875
# <翻译结束>


<原文开始>
// Instance returns an instance for DB operations.
// The parameter `name` specifies the configuration group name,
// which is DefaultGroupName in default.
<原文结束>

# <翻译开始>
// Instance 返回用于数据库操作的实例。
// 参数 `name` 指定配置组名称，默认为 DefaultGroupName。
// md5:06c22232a9c53a60
# <翻译结束>


<原文开始>
// getConfigNodeByGroup calculates and returns a configuration node of given group. It
// calculates the value internally using weight algorithm for load balance.
//
// The parameter `master` specifies whether retrieving a master node, or else a slave node
// if master-slave configured.
<原文结束>

# <翻译开始>
// getConfigNodeByGroup 计算并返回给定组的配置节点。它使用权重算法内部计算值，以实现负载均衡。
//
// 参数 `master` 指定是否获取主节点，如果配置了主从结构，则在非主节点情况下获取从节点。
// md5:0e8709cfd78ceae4
# <翻译结束>


<原文开始>
// Separates master and slave configuration nodes array.
<原文结束>

# <翻译开始>
// 分离主节点和从节点配置数组。. md5:0aea1639f2f64823
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
// getConfigNodeByWeight 计算配置权重并随机返回一个节点。
//
// 算法简述：
// 1. 如果我们有两个节点，它们的权重都是 1，那么权重范围是 [0, 199]；
// 2. 节点1的权重范围是 [0, 99]，节点2的权重范围是 [100, 199]，比例为 1:1；
// 3. 如果随机数是 99，那么它会选择并返回节点1。
// md5:dc1548f9e38ff89b
# <翻译结束>


<原文开始>
	// If total is 0 means all the nodes have no weight attribute configured.
	// It then defaults each node's weight attribute to 1.
<原文结束>

# <翻译开始>
// 如果total为0，表示所有节点都没有配置权重属性。在这种情况下，将为每个节点的权重属性默认设置为1。
// md5:a8625af7b996c9a2
# <翻译结束>


<原文开始>
// Exclude the right border value.
<原文结束>

# <翻译开始>
// 不包括右侧边界值。. md5:660dcac461d09c8d
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
// md5:c9cfb887df88f931
# <翻译结束>


<原文开始>
// getSqlDb retrieves and returns an underlying database connection object.
// The parameter `master` specifies whether retrieves master node connection if
// master-slave nodes are configured.
<原文结束>

# <翻译开始>
// getSqlDb 获取并返回底层的数据库连接对象。
// 参数 `master` 指定是否获取主节点连接，如果配置了主从节点。
// md5:fb885ef2d5264cdc
# <翻译结束>


<原文开始>
// Update the configuration object in internal data.
<原文结束>

# <翻译开始>
// 更新内部数据中的配置对象。. md5:9cbb8bdfb84aa63f
# <翻译结束>


<原文开始>
// Cache the underlying connection pool object by node.
<原文结束>

# <翻译开始>
// 按节点缓存底层连接池对象。. md5:5ba47140febabd5e
# <翻译结束>


<原文开始>
// it here uses node value not pointer as the cache key, in case of oracle ORA-12516 error.
<原文结束>

# <翻译开始>
// 这里使用节点值而不是指针作为缓存键，以防出现Oracle ORA-12516错误。. md5:404d8e507e0c4548
# <翻译结束>


<原文开始>
// It reads from instance map.
<原文结束>

# <翻译开始>
// 从实例映射中读取。. md5:9cd258c405d8d50f
# <翻译结束>

