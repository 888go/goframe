
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
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31
# <翻译结束>


<原文开始>
// GetCore returns the underlying *Core object.
<原文结束>

# <翻译开始>
// GetCore 返回底层的 *Core 对象。 md5:b7d2ff344b9a6a33
# <翻译结束>


<原文开始>
// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
<原文结束>

# <翻译开始>
// Ctx是一个链式函数，它创建并返回一个新的DB对象，它是当前DB对象的浅拷贝，并且包含了给定的上下文。
// 注意，这个返回的DB对象只能使用一次，所以不要将其分配给全局变量或长期使用的包变量。
// md5:9dfddec16d5df9f5
# <翻译结束>


<原文开始>
// It makes a shallow copy of current db and changes its context for next chaining operation.
<原文结束>

# <翻译开始>
// 它会浅复制当前的数据库，并为下一个链式操作更改其上下文。 md5:bf3a0a0f1f30a496
# <翻译结束>


<原文开始>
// It creates a new DB object(NOT NEW CONNECTION), which is commonly a wrapper for object `Core`.
<原文结束>

# <翻译开始>
// 它创建了一个新的DB对象（不是新连接），通常是对`Core`对象的封装。 md5:6cd230087401c98f
# <翻译结束>


<原文开始>
		// It is really a serious error here.
		// Do not let it continue.
<原文结束>

# <翻译开始>
		// 这里确实是一个严重的错误。
		// 不要让它继续下去。
		// md5:790820a929dc0bfd
# <翻译结束>


<原文开始>
// GetCtx returns the context for current DB.
// It returns `context.Background()` is there's no context previously set.
<原文结束>

# <翻译开始>
// GetCtx 返回当前数据库的上下文。
// 如果之前没有设置上下文，则返回 `context.Background()`。
// md5:9b56f79a5eaa891e
# <翻译结束>


<原文开始>
// GetCtxTimeout returns the context and cancel function for specified timeout type.
<原文结束>

# <翻译开始>
// GetCtxTimeout 返回指定超时类型的上下文和取消函数。 md5:5d0be7078de61c6d
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
// Close 方法关闭数据库并阻止新的查询开始。
// 随后，Close 会等待所有已在服务器上开始处理的查询完成。
//
// 关闭 DB 实例是很少见的操作，因为 DB 处理句柄设计为长生命周期的，
// 并且旨在多个 goroutine 间共享。
// md5:39e5c90e1da0ee5e
# <翻译结束>


<原文开始>
// Master creates and returns a connection from master node if master-slave configured.
// It returns the default connection if master-slave not configured.
<原文结束>

# <翻译开始>
// Master 如果配置了主从节点，则创建并返回一个与主节点的连接。
// 如果未配置主从节点，则返回默认连接。
// md5:0bd77b596cdae9a3
# <翻译结束>


<原文开始>
// Slave creates and returns a connection from slave node if master-slave configured.
// It returns the default connection if master-slave not configured.
<原文结束>

# <翻译开始>
// 如果配置了主从模式，Slave 会创建并返回一个从节点连接。如果没有配置主从模式，则返回默认连接。
// md5:d92640050cf063d3
# <翻译结束>


<原文开始>
// GetAll queries and returns data records from database.
<原文结束>

# <翻译开始>
// GetAll 从数据库中查询并返回数据记录。 md5:dfdcfddaa70ab1d4
# <翻译结束>


<原文开始>
// DoSelect queries and returns data records from database.
<原文结束>

# <翻译开始>
// DoSelect 从数据库查询并返回数据记录。 md5:82b06146b8d539d1
# <翻译结束>


<原文开始>
// GetOne queries and returns one record from database.
<原文结束>

# <翻译开始>
// GetOne 从数据库中查询并返回一条记录。 md5:9552f7e095f58141
# <翻译结束>


<原文开始>
// GetArray queries and returns data values as slice from database.
// Note that if there are multiple columns in the result, it returns just one column values randomly.
<原文结束>

# <翻译开始>
// GetArray 从数据库查询并返回数据值作为切片。
// 注意，如果结果中有多个列，它会随机返回一列的值。
// md5:b81cd4c5e063a6f2
# <翻译结束>


<原文开始>
// doGetStruct queries one record from database and converts it to given struct.
// The parameter `pointer` should be a pointer to struct.
<原文结束>

# <翻译开始>
// doGetStruct 从数据库查询一条记录，并将其转换为给定的结构体。
// 参数 `pointer` 应该是指向结构体的指针。
// md5:9260d4f62deef626
# <翻译结束>


<原文开始>
// doGetStructs queries records from database and converts them to given struct.
// The parameter `pointer` should be type of struct slice: []struct/[]*struct.
<原文结束>

# <翻译开始>
// doGetStructs 从数据库查询记录并将其转换为给定的结构体。
// 参数 `pointer` 应为结构体切片类型：[]struct/[]*struct。
// md5:4ce864edda9b9231
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
	// If the query fields do not contain function "COUNT",
	// it replaces the sql string and adds the "COUNT" function to the fields.
<原文结束>

# <翻译开始>
	// 如果查询字段中不包含"COUNT"函数，
	// 则替换SQL字符串，并在字段中添加"COUNT"函数。
	// md5:624b6da82fb9facd
# <翻译结束>


<原文开始>
// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement.
<原文结束>

# <翻译开始>
// Union 执行 "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." 语句。 md5:6a2f9809c172cb31
# <翻译结束>


<原文开始>
// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement.
<原文结束>

# <翻译开始>
// UnionAll 生成 "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ... " 语句。 md5:5a15c8720fcb152f
# <翻译结束>


<原文开始>
// PingMaster pings the master node to check authentication or keeps the connection alive.
<原文结束>

# <翻译开始>
// PingMaster 向主节点发送请求以检查身份验证或保持连接活动。 md5:47a7df55cbee8583
# <翻译结束>


<原文开始>
// PingSlave pings the slave node to check authentication or keeps the connection alive.
<原文结束>

# <翻译开始>
// PingSlave 调用ping命令检查从节点的认证或者维持连接。 md5:62272b38d874eda6
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
// Sort the fields in order.
<原文结束>

# <翻译开始>
// 按顺序对字段进行排序。 md5:3edaf791b6d06284
# <翻译结束>


<原文开始>
// DoInsert inserts or updates data for given table.
// This function is usually used for custom interface definition, you do not need call it manually.
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// Eg:
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// The parameter `option` values are as follows:
// InsertOptionDefault: just insert, if there's unique/primary key in the data, it returns error;
// InsertOptionReplace: if there's unique/primary key in the data, it deletes it from table and inserts a new one;
// InsertOptionSave:    if there's unique/primary key in the data, it updates it or else inserts a new one;
// InsertOptionIgnore:  if there's unique/primary key in the data, it ignores the inserting;
<原文结束>

# <翻译开始>
// DoInsert 函数用于插入或更新给定表的数据。
// 这个函数通常用于自定义接口定义，不需要手动调用。
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"}})
//
// 参数 `option` 的值如下：
// InsertOptionDefault：仅插入，如果数据中包含唯一键或主键，会返回错误；
// InsertOptionReplace：如果数据中包含唯一键或主键，会先从表中删除再插入新的记录；
// InsertOptionSave：   如果数据中包含唯一键或主键，会进行更新，否则插入新记录；
// InsertOptionIgnore： 如果数据中包含唯一键或主键，忽略插入操作。
// md5:e9554638335c9c80
# <翻译结束>


<原文开始>
// Value holder string array, like: (?,?,?)
<原文结束>

# <翻译开始>
// 值持有字符串数组，例如：(?,?,?). md5:4dd91c222c15917f
# <翻译结束>


<原文开始>
// Values that will be committed to underlying database driver.
<原文结束>

# <翻译开始>
// 将被提交给底层数据库驱动程序的值。 md5:d30c8d96f40663c3
# <翻译结束>


<原文开始>
// onDuplicateStr is used in "ON DUPLICATE KEY UPDATE" statement.
<原文结束>

# <翻译开始>
// onDuplicateStr 用于 "ON DUPLICATE KEY UPDATE" 语句。 md5:7056b1b5ea46e69e
# <翻译结束>


<原文开始>
	// ============================================================================================
	// Group the list by fields. Different fields to different list.
	// It here uses ListMap to keep sequence for data inserting.
	// ============================================================================================
<原文结束>

# <翻译开始>
	// ============================================================================================
	// 按照字段对列表进行分组。不同的字段将数据分配到不同的列表中。
	// 此处使用ListMap来保持数据插入时的顺序。
	// ============================================================================================
	// md5:f3b3fbc2fd4a59f8
# <翻译结束>


<原文开始>
// Prepare the batch result pointer.
<原文结束>

# <翻译开始>
// 准备批量结果指针。 md5:dfc8aa8bb292f9d5
# <翻译结束>


<原文开始>
// Upsert clause only takes effect on Save operation.
<原文结束>

# <翻译开始>
// Upsert 子句只在 Save 操作中生效。 md5:c556e85b127111f7
# <翻译结束>


<原文开始>
		// Note that the map type is unordered,
		// so it should use slice+key to retrieve the value.
<原文结束>

# <翻译开始>
		// 注意，映射类型是无序的，
		// 因此应该使用切片和键来检索值。
		// md5:2495d5e730dae78f
# <翻译结束>


<原文开始>
// Batch package checks: It meets the batch number, or it is the last element.
<原文结束>

# <翻译开始>
// 批量包检查：它满足批量数量，或者它是最后一个元素。 md5:a2ef8b869c6d8888
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
// DoUpdate does "UPDATE ... " statement for the table.
// This function is usually used for custom interface definition, you do not need to call it manually.
<原文结束>

# <翻译开始>
// DoUpdate 执行针对表的 "UPDATE ... " 语句。
// 这个函数通常用于自定义接口定义，一般不需要手动调用。
// md5:6d7e08b57dd59a0b
# <翻译结束>


<原文开始>
// Sort the data keys in sequence of table fields.
<原文结束>

# <翻译开始>
// 按照表格字段的顺序对数据键进行排序。 md5:edcdc64a514af6fa
# <翻译结束>


<原文开始>
// If no link passed, it then uses the master link.
<原文结束>

# <翻译开始>
// 如果没有传递链接，那么它就使用主链接。 md5:02e931534071446b
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
// DoDelete does "DELETE FROM ... " statement for the table.
// This function is usually used for custom interface definition, you do not need call it manually.
<原文结束>

# <翻译开始>
// DoDelete 对表执行 "DELETE FROM ..." 语句。
// 此函数通常用于自定义接口定义，无需手动调用。
// md5:f902004d44b55d73
# <翻译结束>


<原文开始>
// FilteredLink retrieves and returns filtered `linkInfo` that can be using for
// logging or tracing purpose.
<原文结束>

# <翻译开始>
// FilteredLink获取并返回经过过滤的`linkInfo`，这些信息可用于日志记录或跟踪目的。
// md5:5d3d4d2f55af0347
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// It just returns the pointer address.
//
// Note that this interface implements mainly for workaround for a json infinite loop bug
// of Golang version < v1.14.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口的MarshalJSON方法。它只是返回指针地址。
// 
// 注意，这个接口主要是为了解决 Golang 版本小于 v1.14 时的json无限循环bug而实现的。
// md5:1b2346be8e04b5fa
# <翻译结束>


<原文开始>
// writeSqlToLogger outputs the Sql object to logger.
// It is enabled only if configuration "debug" is true.
<原文结束>

# <翻译开始>
// writeSqlToLogger 将Sql对象输出到日志记录器。
// 仅当配置"debug"为true时，此功能才启用。
// md5:ad16123093791e59
# <翻译结束>


<原文开始>
// HasTable determine whether the table name exists in the database.
<原文结束>

# <翻译开始>
// HasTable 判断数据库中是否存在指定的表名。 md5:64f8bb54ba260c03
# <翻译结束>


<原文开始>
// GetTablesWithCache retrieves and returns the table names of current database with cache.
<原文结束>

# <翻译开始>
// GetTablesWithCache 使用缓存检索并返回当前数据库中的表名。 md5:9abf0a08a0dbc629
# <翻译结束>


<原文开始>
// IsSoftCreatedFieldName checks and returns whether given field name is an automatic-filled created time.
<原文结束>

# <翻译开始>
// IsSoftCreatedFieldName 检查并返回给定字段名是否为自动填充的创建时间。 md5:f4c7129bbccec8aa
# <翻译结束>


<原文开始>
// FormatSqlBeforeExecuting formats the sql string and its arguments before executing.
// The internal handleArguments function might be called twice during the SQL procedure,
// but do not worry about it, it's safe and efficient.
<原文结束>

# <翻译开始>
// FormatSqlBeforeExecuting 在执行SQL之前格式化SQL字符串及其参数。
// 在SQL执行过程中，内部的handleArguments函数可能会被调用两次，
// 但请不必担心，这是安全且高效的。
// md5:73af1c35794cea21
# <翻译结束>


<原文开始>
	// DO NOT do this as there may be multiple lines and comments in the sql.
	// sql = gstr.Trim(sql)
	// sql = gstr.Replace(sql, "\n", " ")
	// sql, _ = gregex.ReplaceString(`\s{2,}`, ` `, sql)
<原文结束>

# <翻译开始>
	// 不要这样做，因为SQL语句中可能包含多行和注释。
	// sql = gstr.Trim(sql) 	// 删除sql字符串两侧的空白
	// sql = gstr.Replace(sql, "\n", " ") 	// 将换行符替换为单个空格
	// sql, _ = gregex.ReplaceString(`\s{2,}`, ` `, sql) 	// 替换连续两个或更多空格为单个空格
	// md5:907309db612b16e7
# <翻译结束>

