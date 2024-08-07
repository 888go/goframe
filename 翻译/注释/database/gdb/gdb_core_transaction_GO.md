
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
// TXCore is the struct for transaction management.
<原文结束>

# <翻译开始>
// TXCore是事务管理的结构体。 md5:c2173551528f4399
# <翻译结束>


<原文开始>
// db is the current gdb database manager.
<原文结束>

# <翻译开始>
// db是当前的gdb数据库管理器。 md5:cf7449b67dd32334
# <翻译结束>


<原文开始>
// tx is the raw and underlying transaction manager.
<原文结束>

# <翻译开始>
// tx 是原始且底层的交易管理器。 md5:ca359da6d7cbfd5b
# <翻译结束>


<原文开始>
// ctx is the context for this transaction only.
<原文结束>

# <翻译开始>
// ctx 是此次交易专用的上下文。 md5:029c6f2bb9191f37
# <翻译结束>


<原文开始>
// master is the raw and underlying database manager.
<原文结束>

# <翻译开始>
// master是原始的和底层的数据库管理器。 md5:cf639ffb6a4872a2
# <翻译结束>


<原文开始>
// transactionId is a unique id generated by this object for this transaction.
<原文结束>

# <翻译开始>
// transactionId是此对象为此次交易自动生成的唯一标识符。 md5:1837a379fa0972f8
# <翻译结束>


<原文开始>
// transactionCount marks the times that Begins.
<原文结束>

# <翻译开始>
// transactionCount 标记了Begin操作执行的次数。 md5:b733593df5711420
# <翻译结束>


<原文开始>
// isClosed marks this transaction has already been committed or rolled back.
<原文结束>

# <翻译开始>
// isClosed 标记该事务已经提交或回滚。 md5:4a5014ffe4a470ba
# <翻译结束>


<原文开始>
// Begin starts and returns the transaction object.
// You should call Commit or Rollback functions of the transaction object
// if you no longer use the transaction. Commit or Rollback functions will also
// close the transaction automatically.
<原文结束>

# <翻译开始>
// Begin 启动并返回事务对象。
// 如果不再使用事务，你应该调用事务对象的Commit或Rollback方法。
// Commit或Rollback方法也会自动关闭事务。
// md5:cca0e58680665343
# <翻译结束>


<原文开始>
// Transaction wraps the transaction logic using function `f`.
// It rollbacks the transaction and returns the error from function `f` if
// it returns non-nil error. It commits the transaction and returns nil if
// function `f` returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function `f`
// as it is automatically handled by this function.
<原文结束>

# <翻译开始>
// Transaction 包装了使用函数 `f` 执行的事务逻辑。
// 如果函数 `f` 返回非空错误，它将回滚事务并返回该错误。如果函数 `f` 返回 nil，它将提交事务并返回 nil。
// 
// 注意，在函数 `f` 中不应手动提交或回滚事务，因为这些操作将由这个函数自动处理。
// md5:8906440d4dbbef1f
# <翻译结束>


<原文开始>
// Check transaction object from context.
<原文结束>

# <翻译开始>
	// 从上下文中检查交易对象。 md5:98b621386407ef30
# <翻译结束>


<原文开始>
// Inject transaction object into context.
<原文结束>

# <翻译开始>
		// 将事务对象注入上下文。 md5:f5ae21debffd107d
# <翻译结束>


<原文开始>
// WithTX injects given transaction object into context and returns a new context.
<原文结束>

# <翻译开始>
// WithTX 将给定的事务对象注入到上下文中，并返回一个新的上下文。 md5:b4c3c1077b95f681
# <翻译结束>


<原文开始>
// Check repeat injection from given.
<原文结束>

# <翻译开始>
	// 检查从给定的输入中是否存在重复注入。 md5:46e37fcbcbc508b5
# <翻译结束>


<原文开始>
// Inject transaction object and id into context.
<原文结束>

# <翻译开始>
	// 向上下文中注入交易对象和ID。 md5:b9cf191f9e07b60d
# <翻译结束>


<原文开始>
// TXFromCtx retrieves and returns transaction object from context.
// It is usually used in nested transaction feature, and it returns nil if it is not set previously.
<原文结束>

# <翻译开始>
// TXFromCtx 从上下文中获取并返回事务对象。
// 它通常用于嵌套事务功能，如果之前未设置，则返回nil。
// md5:21e22b68139fc8b6
# <翻译结束>


<原文开始>
// transactionKeyForContext forms and returns a string for storing transaction object of certain database group into context.
<原文结束>

# <翻译开始>
// transactionKeyForContext 为特定数据库组生成并返回一个字符串，用于将事务对象存储在上下文中。 md5:1dc9cbe3d8e29f02
# <翻译结束>


<原文开始>
// transactionKeyForNestedPoint forms and returns the transaction key at current save point.
<原文结束>

# <翻译开始>
// transactionKeyForNestedPoint 在当前保存点处生成并返回事务键。 md5:ca81c7094d96c9fc
# <翻译结束>


<原文开始>
// Ctx sets the context for current transaction.
<原文结束>

# <翻译开始>
// Ctx为当前事务设置上下文。 md5:da0e9ba442dc74f9
# <翻译结束>


<原文开始>
// GetCtx returns the context for current transaction.
<原文结束>

# <翻译开始>
// GetCtx 返回当前事务的上下文。 md5:e3cb35516cebab84
# <翻译结束>


<原文开始>
// GetDB returns the DB for current transaction.
<原文结束>

# <翻译开始>
// GetDB 返回当前事务的DB。 md5:26a64f5fed9954b6
# <翻译结束>


<原文开始>
// GetSqlTX returns the underlying transaction object for current transaction.
<原文结束>

# <翻译开始>
// GetSqlTX 返回当前事务的底层事务对象。 md5:31b14245dcb30833
# <翻译结束>


<原文开始>
// Commit commits current transaction.
// Note that it releases previous saved transaction point if it's in a nested transaction procedure,
// or else it commits the hole transaction.
<原文结束>

# <翻译开始>
// Commit 提交当前事务。
// 注意，如果处于嵌套事务过程中，它会释放之前的保存事务点，
// 否则，它将提交整个事务。
// md5:9ca50fd58870ed9e
# <翻译结束>


<原文开始>
// Rollback aborts current transaction.
// Note that it aborts current transaction if it's in a nested transaction procedure,
// or else it aborts the hole transaction.
<原文结束>

# <翻译开始>
// Rollback 会回滚当前事务。
// 注意，在嵌套事务过程中，它会回滚当前的事务；否则，它将回滚整个事务。
// md5:0c483721f8447f53
# <翻译结束>


<原文开始>
// IsClosed checks and returns this transaction has already been committed or rolled back.
<原文结束>

# <翻译开始>
// IsClosed检查并返回此事务是否已经提交或回滚。 md5:cecc2f01ef3e3556
# <翻译结束>


<原文开始>
// Begin starts a nested transaction procedure.
<原文结束>

# <翻译开始>
// Begin 启动一个嵌套事务过程。 md5:1b04e48800ebefdd
# <翻译结束>


<原文开始>
// SavePoint performs `SAVEPOINT xxx` SQL statement that saves transaction at current point.
// The parameter `point` specifies the point name that will be saved to server.
<原文结束>

# <翻译开始>
// SavePoint 执行 `SAVEPOINT xxx` SQL 语句，该语句在当前点保存事务。
// 参数 `point` 指定将被保存到服务器的保存点名称。
// md5:f4061450298afabd
# <翻译结束>


<原文开始>
// RollbackTo performs `ROLLBACK TO SAVEPOINT xxx` SQL statement that rollbacks to specified saved transaction.
// The parameter `point` specifies the point name that was saved previously.
<原文结束>

# <翻译开始>
// RollbackTo 执行 `ROLLBACK TO SAVEPOINT xxx` SQL语句，回滚到指定的保存点事务。
// 参数 `point` 指定了之前保存的保存点名称。
// md5:e347c163ad8fefa7
# <翻译结束>


<原文开始>
// Query does query operation on transaction.
// See Core.Query.
<原文结束>

# <翻译开始>
// Query 在事务上执行查询操作。
// 请参阅Core.Query。
// md5:0d7503cceb0dc1d6
# <翻译结束>


<原文开始>
// Exec does none query operation on transaction.
// See Core.Exec.
<原文结束>

# <翻译开始>
// Exec 在事务上执行非查询操作。
// 参见Core.Exec。
// md5:043edf012223f310
# <翻译结束>


<原文开始>
// Prepare creates a prepared statement for later queries or executions.
// Multiple queries or executions may be run concurrently from the
// returned statement.
// The caller must call the statement's Close method
// when the statement is no longer needed.
<原文结束>

# <翻译开始>
// Prepare 创建一个预处理语句，以便后续的查询或执行。
// 可以从返回的语句中并发地运行多个查询或执行。
// 调用者必须在不再需要该语句时调用语句的 Close 方法。
// md5:16334dc7db1c37a9
# <翻译结束>


<原文开始>
// GetAll queries and returns data records from database.
<原文结束>

# <翻译开始>
// GetAll 从数据库中查询并返回数据记录。 md5:dfdcfddaa70ab1d4
# <翻译结束>


<原文开始>
// GetOne queries and returns one record from database.
<原文结束>

# <翻译开始>
// GetOne 从数据库中查询并返回一条记录。 md5:9552f7e095f58141
# <翻译结束>


<原文开始>
// GetStruct queries one record from database and converts it to given struct.
// The parameter `pointer` should be a pointer to struct.
<原文结束>

# <翻译开始>
// GetStruct 从数据库中查询一条记录，并将其转换为给定的结构体。
// 参数 `pointer` 应该是指向结构体的指针。
// md5:7ddc0d419d6cd2aa
# <翻译结束>


<原文开始>
// GetStructs queries records from database and converts them to given struct.
// The parameter `pointer` should be type of struct slice: []struct/[]*struct.
<原文结束>

# <翻译开始>
// GetStructs 从数据库查询记录，并将它们转换为给定的结构体。参数 `pointer` 应该是结构体切片的类型：[]struct 或 []*struct。
// md5:af7dfbf46c6660c6
# <翻译结束>


<原文开始>
// GetScan queries one or more records from database and converts them to given struct or
// struct array.
//
// If parameter `pointer` is type of struct pointer, it calls GetStruct internally for
// the conversion. If parameter `pointer` is type of slice, it calls GetStructs internally
// for conversion.
<原文结束>

# <翻译开始>
// GetScan 从数据库查询一个或多个记录，并将它们转换为给定的结构体或结构体数组。
//
// 如果参数 `pointer` 是结构体指针类型，它内部会调用 GetStruct 进行转换。如果参数 `pointer` 是切片类型，它内部会调用 GetStructs 进行转换。
// md5:c1dbdab7a7c29c51
# <翻译结束>


<原文开始>
// GetValue queries and returns the field value from database.
// The sql should query only one field from database, or else it returns only one
// field of the result.
<原文结束>

# <翻译开始>
// GetValue 从数据库查询并返回字段值。
// SQL 应该只查询数据库中的一个字段，否则它将只返回结果中的一个字段。
// md5:96794360fadbc288
# <翻译结束>


<原文开始>
// GetCount queries and returns the count from database.
<原文结束>

# <翻译开始>
// GetCount 从数据库中查询并返回计数。 md5:a8368d39f4a58979
# <翻译结束>


<原文开始>
// Insert does "INSERT INTO ..." statement for the table.
// If there's already one unique record of the data in the table, it returns error.
//
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// Eg:
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// The parameter `batch` specifies the batch operation count when given data is slice.
<原文结束>

# <翻译开始>
// Insert 执行 "INSERT INTO..." 语句来操作表。
// 如果表中已经存在数据的唯一记录，它会返回错误。
//
// 参数 `data` 可以是 map、gmap、struct、*struct、[]map 或 []struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 参数 `batch` 在给定数据为切片时，指定批量操作的次数。
// md5:fd75d343f485b8dc
# <翻译结束>


<原文开始>
// InsertIgnore does "INSERT IGNORE INTO ..." statement for the table.
// If there's already one unique record of the data in the table, it ignores the inserting.
//
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// Eg:
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// The parameter `batch` specifies the batch operation count when given data is slice.
<原文结束>

# <翻译开始>
// InsertIgnore 执行 "INSERT IGNORE INTO ..." 语句针对该表。
// 如果表中已存在该数据的唯一记录，则忽略插入操作。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"})
//
// 当给定的数据为切片时，参数 `batch` 指定批处理操作的计数。
// md5:49f76901041c9819
# <翻译结束>


<原文开始>
// InsertAndGetId performs action Insert and returns the last insert id that automatically generated.
<原文结束>

# <翻译开始>
// InsertAndGetId 执行插入操作，并返回自动生成的最后一个插入id。 md5:8d00b40a35fa48a5
# <翻译结束>


<原文开始>
// Replace does "REPLACE INTO ..." statement for the table.
// If there's already one unique record of the data in the table, it deletes the record
// and inserts a new one.
//
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// Eg:
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// If given data is type of slice, it then does batch replacing, and the optional parameter
// `batch` specifies the batch operation count.
<原文结束>

# <翻译开始>
// Replace 用于执行针对该表的 "REPLACE INTO..." 语句。如果表中已经存在数据的唯一记录，它会删除该记录并插入新的。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。如果给定的数据是切片类型，它将进行批量替换，并可选地通过参数 `batch` 指定批量操作次数。
// md5:69ecd0994eab5bbb
# <翻译结束>


<原文开始>
// Save does "INSERT INTO ... ON DUPLICATE KEY UPDATE..." statement for the table.
// It updates the record if there's primary or unique index in the saving data,
// or else it inserts a new record into the table.
//
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// Eg:
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// If given data is type of slice, it then does batch saving, and the optional parameter
// `batch` specifies the batch operation count.
<原文结束>

# <翻译开始>
// Save 执行 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句来操作表。
// 如果保存的数据中存在主键或唯一索引，它将更新记录；否则，将在表中插入新记录。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 如果给定的数据是切片类型，它将进行批量保存。可选参数 `batch` 指定了批量操作的次数。
// md5:c76721f5e0b01424
# <翻译结束>


<原文开始>
// Update does "UPDATE ... " statement for the table.
//
// The parameter `data` can be type of string/map/gmap/struct/*struct, etc.
// Eg: "uid=10000", "uid", 10000, g.Map{"uid": 10000, "name":"john"}
//
// The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc.
// It is commonly used with parameter `args`.
// Eg:
// "uid=10000",
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}.
<原文结束>

# <翻译开始>
// Update 执行表的 "UPDATE ... " 语句。
//
// 参数 `data` 可以是字符串、映射、gmap、结构体或指向结构体的指针等类型。
// 例如："uid=10000", "uid", 10000, g.Map{"uid": 10000, "name":"john"}
//
// 参数 `condition` 也可以是字符串、映射、gmap、切片或结构体及指向结构体的指针等类型。
// 常与参数 `args` 配合使用。
// 例如：
// "uid=10000",
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}.
// md5:8651eb1bd7e10da0
# <翻译结束>


<原文开始>
// Delete does "DELETE FROM ... " statement for the table.
//
// The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc.
// It is commonly used with parameter `args`.
// Eg:
// "uid=10000",
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}.
<原文结束>

# <翻译开始>
// Delete 执行 "DELETE FROM ... " 语句针对该表。
//
// `condition` 参数可以是字符串、映射、gmap、切片、结构体或指向结构体的指针等多种类型。
// 它常与参数 `args` 一起使用。
// 例如：
// "uid=10000",
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}.
// md5:c6c87830434eba7d
# <翻译结束>


<原文开始>
// QueryContext implements interface function Link.QueryContext.
<原文结束>

# <翻译开始>
// QueryContext实现了Link.QueryContext接口函数。 md5:f42e7710688a27fc
# <翻译结束>


<原文开始>
// ExecContext implements interface function Link.ExecContext.
<原文结束>

# <翻译开始>
// ExecContext 实现了 Link.ExecContext 接口函数。 md5:9bd9a386cc5fc878
# <翻译结束>


<原文开始>
// PrepareContext implements interface function Link.PrepareContext.
<原文结束>

# <翻译开始>
// PrepareContext 实现了接口 Link 的 PrepareContext 函数。 md5:b08e1c50bfb8f8e8
# <翻译结束>


<原文开始>
// IsOnMaster implements interface function Link.IsOnMaster.
<原文结束>

# <翻译开始>
// IsOnMaster 实现接口函数 Link.IsOnMaster。 md5:4fddd5d2ad612d30
# <翻译结束>


<原文开始>
// IsTransaction implements interface function Link.IsTransaction.
<原文结束>

# <翻译开始>
// IsTransaction 实现了 Link 接口中的函数 IsTransaction。 md5:692b7be612be52bd
# <翻译结束>

