
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
// GetCore returns the underlying *Core object.
<原文结束>

# <翻译开始>
// GetCore 返回底层的 *Core 对象。
# <翻译结束>


<原文开始>
// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
<原文结束>

# <翻译开始>
// Ctx 是一个链式函数，它创建并返回一个新的 DB 对象，该对象是对当前 DB 对象的浅复制，并且其中包含给定的上下文。
// 注意，返回的这个 DB 对象只能使用一次，所以不要将其赋值给全局或包级别的变量以长期使用。
# <翻译结束>


<原文开始>
// It makes a shallow copy of current db and changes its context for next chaining operation.
<原文结束>

# <翻译开始>
// 它对当前db进行浅复制，并更改其上下文以进行下一个链式操作。
# <翻译结束>


<原文开始>
// It creates a new DB object(NOT NEW CONNECTION), which is commonly a wrapper for object `Core`.
<原文结束>

# <翻译开始>
// 它创建一个新的DB对象（非新连接），这个对象通常是对`Core`对象的一个包装。
# <翻译结束>


<原文开始>
		// It is really a serious error here.
		// Do not let it continue.
<原文结束>

# <翻译开始>
// 这里确实是一个严重的错误。
// 不要让它继续执行。
# <翻译结束>


<原文开始>
// GetCtx returns the context for current DB.
// It returns `context.Background()` is there's no context previously set.
<原文结束>

# <翻译开始>
// GetCtx 返回当前数据库的上下文。
// 如果之前未设置上下文，则返回 `context.Background()`。
# <翻译结束>


<原文开始>
// GetCtxTimeout returns the context and cancel function for specified timeout type.
<原文结束>

# <翻译开始>
// GetCtxTimeout 根据指定的超时类型返回上下文和取消函数。
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
// Close 会等待在服务器上已经开始处理的所有查询完成。
//
// 很少会关闭一个 DB，因为 DB 处理程序旨在长期存在并被多个 goroutine 共享。
# <翻译结束>


<原文开始>
// Master creates and returns a connection from master node if master-slave configured.
// It returns the default connection if master-slave not configured.
<原文结束>

# <翻译开始>
// Master在主从配置的情况下，创建并从主节点返回一个连接。如果未配置主从，则返回默认连接。
# <翻译结束>


<原文开始>
// Slave creates and returns a connection from slave node if master-slave configured.
// It returns the default connection if master-slave not configured.
<原文结束>

# <翻译开始>
// Slave在主从配置的情况下，创建并返回从节点的连接。如果未配置主从，则返回默认连接。
# <翻译结束>


<原文开始>
// GetAll queries and returns data records from database.
<原文结束>

# <翻译开始>
// GetAll 从数据库查询并返回数据记录。
# <翻译结束>


<原文开始>
// DoSelect queries and returns data records from database.
<原文结束>

# <翻译开始>
// DoSelect 从数据库查询并返回数据记录。
# <翻译结束>


<原文开始>
// GetOne queries and returns one record from database.
<原文结束>

# <翻译开始>
// GetOne 从数据库查询并返回一条记录。
# <翻译结束>


<原文开始>
// GetArray queries and returns data values as slice from database.
// Note that if there are multiple columns in the result, it returns just one column values randomly.
<原文结束>

# <翻译开始>
// GetArray 从数据库查询并返回数据值作为切片。
// 注意，如果结果中有多个列，则它会随机返回其中一列的值。
# <翻译结束>


<原文开始>
// doGetStruct queries one record from database and converts it to given struct.
// The parameter `pointer` should be a pointer to struct.
<原文结束>

# <翻译开始>
// doGetStruct 从数据库查询一条记录并将其转换为给定的结构体。
// 参数 `pointer` 应该是指向结构体的指针。
# <翻译结束>


<原文开始>
// doGetStructs queries records from database and converts them to given struct.
// The parameter `pointer` should be type of struct slice: []struct/[]*struct.
<原文结束>

# <翻译开始>
// doGetStructs 从数据库查询记录并将其转换为给定的结构体。
// 参数 `pointer` 应为结构体切片类型：[]struct 或 []*struct。
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
// 如果参数`pointer`是结构体指针类型，它会内部调用 GetStruct 进行转换。如果参数 `pointer` 是切片类型，则会内部调用 GetStructs 进行转换。
# <翻译结束>


<原文开始>
// GetValue queries and returns the field value from database.
// The sql should query only one field from database, or else it returns only one
// field of the result.
<原文结束>

# <翻译开始>
// GetValue 从数据库查询并返回字段值。
// SQL语句应当只查询数据库中的一个字段，否则它将仅返回结果中的一个字段。
# <翻译结束>


<原文开始>
// GetCount queries and returns the count from database.
<原文结束>

# <翻译开始>
// GetCount 从数据库查询并返回计数。
# <翻译结束>


<原文开始>
	// If the query fields do not contain function "COUNT",
	// it replaces the sql string and adds the "COUNT" function to the fields.
<原文结束>

# <翻译开始>
// 如果查询字段中不包含函数"COUNT"，
// 则替换sql字符串，并在字段中添加"COUNT"函数。
// 这段代码的注释是说，当SQL查询语句中的字段部分未使用“COUNT”函数时，会对原始的sql字符串进行替换处理，将“COUNT”函数添加到字段表达式中。
# <翻译结束>


<原文开始>
// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement.
<原文结束>

# <翻译开始>
// Union 执行 "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." 语句。
# <翻译结束>


<原文开始>
// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement.
<原文结束>

# <翻译开始>
// UnionAll 执行 "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." 语句。
# <翻译结束>


<原文开始>
// PingMaster pings the master node to check authentication or keeps the connection alive.
<原文结束>

# <翻译开始>
// PingMaster 用于向主节点发送心跳以检查身份验证或保持连接存活。
# <翻译结束>


<原文开始>
// PingSlave pings the slave node to check authentication or keeps the connection alive.
<原文结束>

# <翻译开始>
// PingSlave 向从节点发送ping请求，用于检查身份验证或保持连接活跃。
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
// Insert 执行针对该表的 "INSERT INTO ..." 语句。
// 如果表中已经存在一条相同数据的唯一记录，则返回错误。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 当给定的数据为切片时，参数 `batch` 指定了批量操作的数量。
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
// InsertIgnore 执行针对表的 "INSERT IGNORE INTO ..." 语句。
// 如果表中已存在一条相同的数据记录，它将忽略插入操作。
//
// 参数 `data` 可以为 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 当给定数据为切片时，参数 `batch` 指定批量操作的数量。
# <翻译结束>


<原文开始>
// InsertAndGetId performs action Insert and returns the last insert id that automatically generated.
<原文结束>

# <翻译开始>
// InsertAndGetId 执行插入操作，并返回自动生成的最后一个插入ID。
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
// Replace 执行针对该表的 "REPLACE INTO ..." 语句。
// 如果表中已存在一条唯一数据记录，它会先删除这条记录，然后插入一条新的记录。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 若给定的数据是切片类型，它将执行批量替换操作，可选参数
// `batch` 指定了批量操作的数量。
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
// Save 执行针对表的 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句。
// 如果保存数据中存在主键或唯一索引，它将更新记录，否则将在表中插入新的记录。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 如果给定的数据是切片类型，那么它将执行批量保存操作，可选参数
// `batch` 指定了批量操作的数量。
# <翻译结束>












<原文开始>
// DoInsert inserts or updates data forF given table.
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
// DoInsert 插入或更新给定表中的数据。
// 该函数通常用于自定义接口定义，您无需手动调用它。
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型的。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
// 参数 `option` 的取值如下：
// InsertOptionDefault：仅插入，如果数据中存在唯一/主键，则返回错误；
// InsertOptionReplace：如果数据中存在唯一/主键，先从表中删除，再插入新的数据；
// InsertOptionSave：如果数据中存在唯一/主键，则更新该记录，否则插入新记录；
// InsertOptionIgnore：如果数据中存在唯一/主键，则忽略插入操作。
# <翻译结束>


<原文开始>
// Value holder string array, like: (?,?,?)
<原文结束>

# <翻译开始>
// 值持有者字符串数组，例如：(?,?,?)
# <翻译结束>


<原文开始>
// Values that will be committed to underlying database driver.
<原文结束>

# <翻译开始>
// 这些值将会被提交到底层数据库驱动中。
# <翻译结束>


<原文开始>
	// ============================================================================================
	// Group the list by fields. Different fields to different list.
	// It here uses ListMap to keep sequence for data inserting.
	// ============================================================================================
<原文结束>

# <翻译开始>
// ============================================================================================
// 根据字段对列表进行分组。不同的字段将数据分到不同的列表中。
// 这里使用ListMap来保持数据插入时的顺序。
// ============================================================================================
# <翻译结束>


<原文开始>
// Prepare the batch result pointer.
<原文结束>

# <翻译开始>
// 准备批量结果指针。
# <翻译结束>


<原文开始>
		// Note that the map type is unordered,
		// so it should use slice+key to retrieve the value.
<原文结束>

# <翻译开始>
// 请注意，map类型是无序的，
// 所以应当使用切片+键来获取值。
# <翻译结束>


<原文开始>
// Batch package checks: It meets the batch number, or it is the last element.
<原文结束>

# <翻译开始>
// 批量校验包：满足批量数量，或者已是最后一个元素。
# <翻译结束>


<原文开始>
// If it's SAVE operation, do not automatically update the creating time.
<原文结束>

# <翻译开始>
// 如果是保存操作，不自动更新创建时间。
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
// Update 执行针对该表的 "UPDATE ... " 语句。
//
// 参数 `data` 可以为 string/map/gmap/struct/*struct 等类型。
// 例如："uid=10000"、"uid", 10000、g.Map{"uid": 10000, "name":"john"}
//
// 参数 `condition` 可以为 string/map/gmap/slice/struct/*struct 等类型，通常与参数 `args` 一起使用。
// 例如：
// "uid=10000"
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}
// 注：这里的 `g.Map` 和 `g.Slice` 是一种特定的 Go 语言数据结构（可能是自定义类型），分别代表映射和切片。
# <翻译结束>


<原文开始>
// DoUpdate does "UPDATE ... " statement for the table.
// This function is usually used for custom interface definition, you do not need to call it manually.
<原文结束>

# <翻译开始>
// DoUpdate 执行针对该表的 "UPDATE ... " 语句。
// 该函数通常用于自定义接口定义，您无需手动调用它。
# <翻译结束>


<原文开始>
// Sort the data keys in sequence of table fields.
<原文结束>

# <翻译开始>
// 按照表格字段的顺序对数据键进行排序。
# <翻译结束>


<原文开始>
// If no link passed, it then uses the master link.
<原文结束>

# <翻译开始>
// 如果没有传递链接，则使用主链接。
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
// Delete 执行针对该表的 "DELETE FROM ... " 语句。
//
// 参数 `condition` 可以是字符串、映射（map/gmap）、切片、结构体或指针类型等。
// 常常与参数 `args` 一起使用。例如：
// "uid=10000"，
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}
// 中文注释：
// Delete 方法用于执行对该数据表执行 "DELETE FROM ... " SQL语句。
//
// 参数 `condition` 支持多种类型，如字符串、字典(map/gmap)、数组、结构体、结构体指针等。
// 通常会结合参数 `args` 使用，例如以下示例：
// "uid=10000"，（条件为 uid 等于 10000）
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%" （条件为 money 大于 99999 并且 name 字段匹配 "vip_%"）
// "status IN (?)", g.Slice{1,2,3} （条件为 status 字段在数组 [1,2,3] 中）
// "age IN(?,?)", 18, 50 （条件为 age 字段在范围 18 到 50 内）
// User{ Id : 1, UserName : "john"} （根据结构体定义的字段作为条件）
# <翻译结束>


<原文开始>
// DoDelete does "DELETE FROM ... " statement for the table.
// This function is usually used for custom interface definition, you do not need call it manually.
<原文结束>

# <翻译开始>
// DoDelete 执行针对表的 "DELETE FROM ..." 语句。
// 该函数通常用于自定义接口定义，无需手动调用。
# <翻译结束>


<原文开始>
// FilteredLink retrieves and returns filtered `linkInfo` that can be using for
// logging or tracing purpose.
<原文结束>

# <翻译开始>
// FilteredLink 获取并返回可用于日志记录或跟踪目的的已过滤`linkInfo`。
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// It just returns the pointer address.
//
// Note that this interface implements mainly for workaround for a json infinite loop bug
// of Golang version < v1.14.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口的MarshalJSON方法。
// 它仅仅是返回指针地址。
//
// 注意，这个接口主要为了应对 Go 语言版本小于 v1.14 时存在的一个 json 无限循环bug而实现的。
# <翻译结束>


<原文开始>
// writeSqlToLogger outputs the Sql object to logger.
// It is enabled only if configuration "debug" is true.
<原文结束>

# <翻译开始>
// writeSqlToLogger 将 Sql 对象输出到日志器。
// 仅当配置项 "debug" 为 true 时，此功能才被启用。
# <翻译结束>


<原文开始>
// HasTable determine whether the table name exists in the database.
<原文结束>

# <翻译开始>
// HasTable 判断给定的表名是否存在于数据库中。
# <翻译结束>


<原文开始>
// GetTablesWithCache retrieves and returns the table names of current database with cache.
<原文结束>

# <翻译开始>
// GetTablesWithCache 使用缓存获取并返回当前数据库的表名。
# <翻译结束>


<原文开始>
// isSoftCreatedFieldName checks and returns whether given field name is an automatic-filled created time.
<原文结束>

# <翻译开始>
// isSoftCreatedFieldName 检查并返回给定的字段名是否为自动填充的创建时间。
# <翻译结束>


<原文开始>
// FormatSqlBeforeExecuting formats the sql string and its arguments before executing.
// The internal handleArguments function might be called twice during the SQL procedure,
// but do not worry about it, it's safe and efficient.
<原文结束>

# <翻译开始>
// FormatSqlBeforeExecuting 在执行SQL之前，对SQL字符串及其参数进行格式化处理。
// 在SQL过程中，内部函数handleArguments可能被调用两次，
// 但请不用担心，这是安全且高效的。
# <翻译结束>


<原文开始>
	// DO NOT do this as there may be multiple lines and comments in the sql.
	// sql = gstr.Trim(sql)
	// sql = gstr.Replace(sql, "\n", " ")
	// sql, _ = gregex.ReplaceString(`\s{2,}`, ` `, sql)
<原文结束>

# <翻译开始>
// **不要**这样做，因为SQL中可能包含多行和注释。
// 删除sql的首尾空格
// 将sql中的换行符("\n")替换为空格
// 使用正则表达式将sql中连续出现2个或以上空格的情况替换为单个空格，并返回处理后的sql（_用于忽略错误信息）
# <翻译结束>


<原文开始>
// onDuplicateStr is used in "ON DUPLICATE KEY UPDATE" statement.
<原文结束>

# <翻译开始>
// onDuplicateStr 用于 "ON DUPLICATE KEY UPDATE" 语句中。
# <翻译结束>


<原文开始>
// Sort the fields in order.
<原文结束>

# <翻译开始>
// 按照顺序对字段进行排序。
# <翻译结束>


<原文开始>
// Sort the input fields.
<原文结束>

# <翻译开始>
// 对输入字段进行排序。
# <翻译结束>

