
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
// Model is core struct implementing the DAO for ORM.
<原文结束>

# <翻译开始>
// Model 是核心结构体，实现了 ORM 的 DAO（数据访问对象）。
# <翻译结束>












<原文开始>
// rawSql is the raw SQL string which marks a raw SQL based Model not a table based Model.
<原文结束>

# <翻译开始>
// rawSql 是原始SQL字符串，用于标记基于原始SQL的模型，而非基于表的模型。
# <翻译结束>







<原文开始>
// Mark for operation on master or slave.
<原文结束>

# <翻译开始>
// 标记用于在主服务器或从服务器上执行操作。
# <翻译结束>


<原文开始>
// Table names when model initialization.
<原文结束>

# <翻译开始>
// 在模型初始化时的表格名称。
# <翻译结束>


<原文开始>
// Arguments for With feature.
<原文结束>

# <翻译开始>
// With功能的参数。
# <翻译结束>


<原文开始>
// Extra custom arguments for sql, which are prepended to the arguments before sql committed to underlying driver.
<原文结束>

# <翻译开始>
// 在SQL提交给底层驱动之前，额外自定义的SQL参数，这些参数将被追加到原有参数之前。
# <翻译结束>


<原文开始>
// Condition builder for where operation.
<原文结束>

# <翻译开始>
// 条件构造器，用于where操作。
# <翻译结束>


<原文开始>
// Option for extra operation features.
<原文结束>

# <翻译开始>
// 用于额外操作功能的选项。
# <翻译结束>


<原文开始>
// Offset statement for some databases grammar.
<原文结束>

# <翻译开始>
// 为某些数据库语法提供的偏移量语句。
# <翻译结束>


<原文开始>
// Partition table partition name.
<原文结束>

# <翻译开始>
// 分区表分区名称。
# <翻译结束>


<原文开始>
// Data for operation, which can be type of map/[]map/struct/*struct/string, etc.
<原文结束>

# <翻译开始>
// Data 用于操作的数据，其类型可以是 map/[]map/struct/*struct/string 等等。
# <翻译结束>


<原文开始>
// Batch number for batch Insert/Replace/Save operations.
<原文结束>

# <翻译开始>
// 批量插入/替换/保存操作的批次号。
# <翻译结束>


<原文开始>
// Filter data and where key-value pairs according to the fields of the table.
<原文结束>

# <翻译开始>
// 根据表格字段过滤并筛选出符合条件的键值对数据。
# <翻译结束>


<原文开始>
// Force the query to only return distinct results.
<原文结束>

# <翻译开始>
// 强制查询只返回不重复的结果。
# <翻译结束>


<原文开始>
// Lock for update or in shared lock.
<原文结束>

# <翻译开始>
// 加锁以便进行更新或共享锁操作。
# <翻译结束>


<原文开始>
// Enable sql result cache feature, which is mainly for indicating cache duration(especially 0) usage.
<原文结束>

# <翻译开始>
// 启用SQL结果缓存功能，主要用于指示缓存持续时间（尤其是0）的使用情况。
# <翻译结束>


<原文开始>
// Cache option for query statement.
<原文结束>

# <翻译开始>
// 查询语句的缓存选项。
# <翻译结束>


<原文开始>
// Hook functions for model hook feature.
<原文结束>

# <翻译开始>
// 钩子函数，用于模型钩子功能。
# <翻译结束>


<原文开始>
// Disables soft deleting features when select/delete operations.
<原文结束>

# <翻译开始>
// 禁用在选择/删除操作时的软删除功能。
# <翻译结束>


<原文开始>
// If true, it clones and returns a new model object whenever operation done; or else it changes the attribute of current model.
<原文结束>

# <翻译开始>
// 如果为true，则在操作完成后克隆并返回一个新的模型对象；否则，它会改变当前模型的属性。
# <翻译结束>


<原文开始>
// Table alias to true table name, usually used in join statements.
<原文结束>

# <翻译开始>
// 表别名到真实表名的映射，通常用于连接语句中。
# <翻译结束>


<原文开始>
// ModelHandler is a function that handles given Model and returns a new Model that is custom modified.
<原文结束>

# <翻译开始>
// ModelHandler 是一个函数，用于处理给定的 Model，并返回一个经过自定义修改的新 Model。
# <翻译结束>


<原文开始>
// ChunkHandler is a function that is used in function Chunk, which handles given Result and error.
// It returns true if it wants to continue chunking, or else it returns false to stop chunking.
<原文结束>

# <翻译开始>
// ChunkHandler 是一个在 Chunk 函数中使用的函数，用于处理给定的 Result 和错误。
// 如果希望继续分块处理，则返回 true；否则返回 false 以停止分块处理。
# <翻译结束>


<原文开始>
// Model creates and returns a new ORM model from given schema.
// The parameter `tableNameQueryOrStruct` can be more than one table names, and also alias name, like:
//  1. Model names:
//     db.Model("user")
//     db.Model("user u")
//     db.Model("user, user_detail")
//     db.Model("user u, user_detail ud")
//  2. Model name with alias:
//     db.Model("user", "u")
//  3. Model name with sub-query:
//     db.Model("? AS a, ? AS b", subQuery1, subQuery2)
<原文结束>

# <翻译开始>
// Model 根据给定的模式创建并返回一个新的 ORM 模型。
// 参数 `tableNameQueryOrStruct` 可以是多个表名，也可以包含别名，如下所示：
// 1. 表名示例：
//     db.Model("user")                 // 单个表名
//     db.Model("user u")               // 带别名的表名
//     db.Model("user, user_detail")    // 多个表名
//     db.Model("user u, user_detail ud") // 多个带别名的表名
// 2. 包含别名的表名示例：
//     db.Model("user", "u")         // 表名和对应的别名
// 3. 使用子查询作为表名的示例：
//     db.Model("? AS a, ? AS b", subQuery1, subQuery2) // 使用子查询表达式作为模型，并为子查询结果设置别名
# <翻译结束>


<原文开始>
// Model creation with sub-query.
<原文结束>

# <翻译开始>
// 使用子查询创建模型
# <翻译结束>







<原文开始>
// Raw creates and returns a model based on a raw sql not a table.
// Example:
//
//	db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
<原文结束>

# <翻译开始>
// Raw 根据原始 SQL（非表）创建并返回一个模型。
// 通常用于嵌入原始sql语句,如:
// g.Model("user").WhereLT("created_at", gdb.Raw("now()")).All()  // SELECT * FROM `user` WHERE `created_at`<now()
// 参考文档:https://goframe.org/pages/viewpage.action?pageId=111911590&showComments=true

// 也可以直接直接执行原始sql,示例：
// db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
// 上述代码表示，通过执行原始SQL语句（根据"name"为"john"的条件查询user表中所有列），并使用Scan方法将查询结果绑定到result变量中。
# <翻译结束>


<原文开始>
// Raw sets current model as a raw sql model.
// Example:
//
//	db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
//
// See Core.Raw.
<原文结束>

# <翻译开始>
// Raw 将当前模型设置为原始SQL模型。
// 通常用于嵌入原始sql语句,如:
// g.Model("user").WhereLT("created_at", gdb.Raw("now()")).All()  // SELECT * FROM `user` WHERE `created_at`<now()
// 参考文档:https://goframe.org/pages/viewpage.action?pageId=111911590&showComments=true
//
// 也可以直接直接执行原始sql,示例:
// db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
// 请参阅Core.Raw。
# <翻译结束>


<原文开始>
// With creates and returns an ORM model based on metadata of given object.
<原文结束>

# <翻译开始>
// With 根据给定对象的元数据创建并返回一个ORM模型。
//
// 原注释未提及with使用方法, 以下摘自Model对象示例,仅供参考.
// With 启用关联查询，通过给定的属性对象指定开启。
// 常考"模型关联-静态关联"文档:https://goframe.org/pages/viewpage.action?pageId=7297190
// 例如，如果给定如下的结构体定义：
//
//	type User struct {
//		 gmeta.Meta `orm:"table:user"` // 定义表名为 user
//		 Id         int           `json:"id"`    // 用户ID
//		 Name       string        `json:"name"`   // 用户名
//		 UserDetail *UserDetail   `orm:"with:uid=id"` // 关联 UserDetail 表，通过 uid 等于 id 进行关联
//		 UserScores []*UserScores `orm:"with:uid=id"` // 关联 UserScores 表，通过 uid 等于 id 进行关联
//	}
//
// 我们可以通过以下方式在属性 `UserDetail` 和 `UserScores` 上启用模型关联操作：
// db.With(User{}.UserDetail).With(User{}.UserScores).Scan(xxx)
// 或者：
// db.With(UserDetail{}).With(UserScores{}).Scan(xxx)
// 或者：
// db.With(UserDetail{}, UserScores{}).Scan(xxx)
# <翻译结束>


<原文开始>
// Model acts like Core.Model except it operates on transaction.
// See Core.Model.
<原文结束>

# <翻译开始>
// Model 类似于 Core.Model，但其在事务上进行操作。
// 请参阅 Core.Model。
# <翻译结束>


<原文开始>
// With acts like Core.With except it operates on transaction.
// See Core.With.
<原文结束>

# <翻译开始>
// With 类似于 Core.With，但其操作针对事务。
// 请参阅 Core.With。
//
// 原注释未提及with使用方法, 以下摘自Model对象示例,仅供参考.
// With 启用关联查询，通过给定的属性对象指定开启。
// 常考"模型关联-静态关联"文档:https://goframe.org/pages/viewpage.action?pageId=7297190
// 例如，如果给定如下的结构体定义：
//	type User struct {
//		 gmeta.Meta `orm:"table:user"` // 定义表名为 user
//		 Id         int           `json:"id"`    // 用户ID
//		 Name       string        `json:"name"`   // 用户名
//		 UserDetail *UserDetail   `orm:"with:uid=id"` // 关联 UserDetail 表，通过 uid 等于 id 进行关联
//		 UserScores []*UserScores `orm:"with:uid=id"` // 关联 UserScores 表，通过 uid 等于 id 进行关联
//	}
//
// 我们可以通过以下方式在属性 `UserDetail` 和 `UserScores` 上启用模型关联操作：
// db.With(User{}.UserDetail).With(User{}.UserScores).Scan(xxx)
// 或者：
// db.With(UserDetail{}).With(UserScores{}).Scan(xxx)
// 或者：
// db.With(UserDetail{}, UserScores{}).Scan(xxx)
# <翻译结束>


<原文开始>
// Ctx sets the context for current operation.
<原文结束>

# <翻译开始>
// Ctx 设置当前操作的上下文。
# <翻译结束>


<原文开始>
// GetCtx returns the context for current Model.
// It returns `context.Background()` is there's no context previously set.
<原文结束>

# <翻译开始>
// GetCtx 返回当前 Model 的上下文。
// 若此前未设置过上下文，则返回 `context.Background()`。
# <翻译结束>


<原文开始>
// As sets an alias name for current table.
<原文结束>

# <翻译开始>
// As 为当前表设置别名名称。
# <翻译结束>












<原文开始>
// DB sets/changes the db object for current operation.
<原文结束>

# <翻译开始>
// DB 设置/更改当前操作的数据库对象。
# <翻译结束>


<原文开始>
// TX sets/changes the transaction for current operation.
<原文结束>

# <翻译开始>
// TX 设置/更改当前操作的事务。
# <翻译结束>


<原文开始>
// Schema sets the schema for current operation.
<原文结束>

# <翻译开始>
// 设置当前操作的模式。
# <翻译结束>


<原文开始>
// Clone creates and returns a new model which is a Clone of current model.
// Note that it uses deep-copy for the Clone.
<原文结束>

# <翻译开始>
// Clone 创建并返回一个新的模型，该模型是当前模型的克隆版本。
// 注意，它使用深度复制进行克隆。
# <翻译结束>







<原文开始>
// WhereBuilder copy, note the attribute pointer.
<原文结束>

# <翻译开始>
// WhereBuilder 复制，注意属性指针。
# <翻译结束>


<原文开始>
// Shallow copy slice attributes.
<原文结束>

# <翻译开始>
// 浅复制切片属性。
# <翻译结束>


<原文开始>
// Master marks the following operation on master node.
<原文结束>

# <翻译开始>
// Master 标识以下操作将在主节点上执行。
# <翻译结束>


<原文开始>
// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
<原文结束>

# <翻译开始>
// Slave 标记在从属节点上执行的后续操作。
// 注意，只有在配置了从属节点时才有意义。
# <翻译结束>


<原文开始>
// Safe marks this model safe or unsafe. If safe is true, it clones and returns a new model object
// whenever the operation done, or else it changes the attribute of current model.
<原文结束>

# <翻译开始>
// Safe 用于标记该模型为安全或不安全。如果 safe 为 true，则在每次操作完成后都会克隆并返回一个新的模型对象；否则，它会改变当前模型的属性。
# <翻译结束>


<原文开始>
// Args sets custom arguments for model operation.
<原文结束>

# <翻译开始>
// Args 设置模型操作的自定义参数。
# <翻译结束>


<原文开始>
// Handler calls each of `handlers` on current Model and returns a new Model.
// ModelHandler is a function that handles given Model and returns a new Model that is custom modified.
<原文结束>

# <翻译开始>
// Handler 函数对当前 Model 调用 `handlers` 中的每个处理函数，并返回一个新的已修改过的 Model。
// ModelHandler 是一个函数，它接收并处理给定的 Model，并返回一个根据需求自定义修改后的新 Model。
// ```go
// Handler 函数会依次调用传入的所有 `handlers` 对当前 Model 进行操作，并最终返回一个经过各处理函数更新后的新 Model 实例。
// 其中，ModelHandler 类型是一个函数，该函数接受一个 Model 参数，并基于此参数进行处理，最后返回一个经过定制修改后的新的 Model 实例。
# <翻译结束>


<原文开始>
// Operation table names, which can be more than one table names and aliases, like: "user", "user u", "user u, user_detail ud".
<原文结束>

# <翻译开始>
// 操作表名，可以是多个表名及别名，例如："user"、"user u"、"user u, user_detail ud"。
# <翻译结束>


<原文开始>
// Operation fields, multiple fields joined using char ','.
<原文结束>

# <翻译开始>
// 操作字段，多个字段使用字符 ',' 连接。
# <翻译结束>


<原文开始>
// Excluded operation fields, multiple fields joined using char ','.
<原文结束>

# <翻译开始>
// 排除的操作字段，多个字段使用逗号（char ',')连接。
# <翻译结束>


<原文开始>
// Enable model association operations on all objects that have "with" tag in the struct.
<原文结束>

# <翻译开始>
// 在结构体中具有"with"标签的所有对象上启用模型关联操作。
# <翻译结束>


<原文开始>
// Used for "group by" statement.
<原文结束>

# <翻译开始>
// 用于“group by”语句。
# <翻译结束>


<原文开始>
// Used for "order by" statement.
<原文结束>

# <翻译开始>
// 用于 "order by" 语句。
# <翻译结束>


<原文开始>
// Used for "having..." statement.
<原文结束>

# <翻译开始>
// 用于 "having..." 语句。
# <翻译结束>


<原文开始>
// Used for "select ... start, limit ..." statement.
<原文结束>

# <翻译开始>
// 用于 "select ... start, limit ..." 语句。
# <翻译结束>


<原文开始>
// onDuplicate is used for ON "DUPLICATE KEY UPDATE" statement.
<原文结束>

# <翻译开始>
// onDuplicate 用于 ON "DUPLICATE KEY UPDATE" 语句。
# <翻译结束>


<原文开始>
// onDuplicateEx is used for excluding some columns ON "DUPLICATE KEY UPDATE" statement.
<原文结束>

# <翻译开始>
// onDuplicateEx 用于在 "DUPLICATE KEY UPDATE" 语句中排除某些列。
# <翻译结束>


<原文开始>
// Underlying DB interface.
<原文结束>

# <翻译开始>
// 基础数据库接口。
# <翻译结束>


<原文开始>
// Underlying TX interface.
<原文结束>

# <翻译开始>
// 基础的TX接口。
# <翻译结束>


<原文开始>
// Custom database schema.
<原文结束>

# <翻译开始>
// 自定义数据库模式
# <翻译结束>


<原文开始>
// Normal model creation.
<原文结束>

# <翻译开始>
// 正常模型创建。
# <翻译结束>


<原文开始>
// Basic attributes copy.
<原文结束>

# <翻译开始>
// 基础属性复制
# <翻译结束>

