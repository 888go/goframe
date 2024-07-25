
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Model is core struct implementing the DAO for ORM.
<原文结束>

# <翻译开始>
// Model是实现ORM DAO的核心结构体。 md5:7230072d015718fc
# <翻译结束>


<原文开始>
// Underlying DB interface.
<原文结束>

# <翻译开始>
// 底层数据库接口。 md5:5b73fd8ce3fdaf5a
# <翻译结束>


<原文开始>
// Underlying TX interface.
<原文结束>

# <翻译开始>
// 底层的TX接口。 md5:d13e03783f7815aa
# <翻译结束>


<原文开始>
// rawSql is the raw SQL string which marks a raw SQL based Model not a table based Model.
<原文结束>

# <翻译开始>
// rawSql 是原始的SQL字符串，它标志着一个基于原始SQL的模型，而不是基于表的模型。 md5:b83edc253c98f3de
# <翻译结束>


<原文开始>
// Custom database schema.
<原文结束>

# <翻译开始>
// 自定义数据库模式。 md5:173e3cf7ad252f16
# <翻译结束>


<原文开始>
// Mark for operation on master or slave.
<原文结束>

# <翻译开始>
// 用于标记在主节点或从节点上执行的操作。 md5:d59587510f982160
# <翻译结束>


<原文开始>
// Table names when model initialization.
<原文结束>

# <翻译开始>
// 模型初始化时的表格名称。 md5:7569da250a03d8b2
# <翻译结束>


<原文开始>
// Operation table names, which can be more than one table names and aliases, like: "user", "user u", "user u, user_detail ud".
<原文结束>

# <翻译开始>
// 操作表名，可以是多个表名和别名，例如："user"、"user u"、"user u, user_detail ud"。 md5:140ed796dfa7b2e5
# <翻译结束>


<原文开始>
// Operation fields, multiple fields joined using char ','.
<原文结束>

# <翻译开始>
// 操作字段，多个字段通过字符'.'连接。 md5:90a8233be912ab73
# <翻译结束>


<原文开始>
// Excluded operation fields, multiple fields joined using char ','.
<原文结束>

# <翻译开始>
// 排除的操作字段，多个字段使用逗号','连接。 md5:0757072228393ad3
# <翻译结束>


<原文开始>
// Arguments for With feature.
<原文结束>

# <翻译开始>
// With功能的参数。 md5:6da19be4d3cc5337
# <翻译结束>


<原文开始>
// Enable model association operations on all objects that have "with" tag in the struct.
<原文结束>

# <翻译开始>
// 启用在结构体中带有 "with" 标签的所有对象的模型关联操作。 md5:68534968f6dd65cd
# <翻译结束>


<原文开始>
// Extra custom arguments for sql, which are prepended to the arguments before sql committed to underlying driver.
<原文结束>

# <翻译开始>
// 为SQL提供的额外自定义参数，这些参数将在将SQL提交给底层驱动程序之前添加到参数前面。 md5:e6a840d23cdc5b31
# <翻译结束>


<原文开始>
// Condition builder for where operation.
<原文结束>

# <翻译开始>
// 用于构建where操作的条件生成器。 md5:4e7d38dd793619e1
# <翻译结束>


<原文开始>
// Used for "group by" statement.
<原文结束>

# <翻译开始>
// 用于"分组 by"语句。 md5:0054c7d82c75aa83
# <翻译结束>


<原文开始>
// Used for "order by" statement.
<原文结束>

# <翻译开始>
// 用于"ORDER BY"语句。 md5:974c6823a972edbe
# <翻译结束>


<原文开始>
// Used for "having..." statement.
<原文结束>

# <翻译开始>
// 用于 "having..." 语句。 md5:fc87b6be31414f4e
# <翻译结束>


<原文开始>
// Used for "select ... start, limit ..." statement.
<原文结束>

# <翻译开始>
// 用于 "select ... start, limit ..." 语句。 md5:28a92730f0f33ffe
# <翻译结束>


<原文开始>
// Option for extra operation features.
<原文结束>

# <翻译开始>
// 用于额外操作功能的选项。 md5:46fa8be84b899e8f
# <翻译结束>


<原文开始>
// Offset statement for some databases grammar.
<原文结束>

# <翻译开始>
// 用于某些数据库语法的偏移语句。 md5:222cd8b108c2f2fc
# <翻译结束>


<原文开始>
// Partition table partition name.
<原文结束>

# <翻译开始>
// 分区表的分区名称。 md5:f8b787fa2b446be6
# <翻译结束>


<原文开始>
// Data for operation, which can be type of map/[]map/struct/*struct/string, etc.
<原文结束>

# <翻译开始>
// 操作数据，可以是映射类型/切片映射/结构体/结构体指针/字符串等类型。 md5:d9d2d3cef3841513
# <翻译结束>


<原文开始>
// Batch number for batch Insert/Replace/Save operations.
<原文结束>

# <翻译开始>
// 批量操作的批次编号，用于批量插入/替换/保存操作。 md5:72e06e8a06a3dfa8
# <翻译结束>


<原文开始>
// Filter data and where key-value pairs according to the fields of the table.
<原文结束>

# <翻译开始>
// 根据表格的字段，过滤数据和键值对。 md5:6af1b96126cc53e6
# <翻译结束>


<原文开始>
// Force the query to only return distinct results.
<原文结束>

# <翻译开始>
// 强制查询只返回唯一的结果。 md5:10ef583cb57e7d16
# <翻译结束>


<原文开始>
// Lock for update or in shared lock.
<原文结束>

# <翻译开始>
// 用于更新或共享锁的锁定。 md5:a2e8bcf922a3cd09
# <翻译结束>


<原文开始>
// Enable sql result cache feature, which is mainly for indicating cache duration(especially 0) usage.
<原文结束>

# <翻译开始>
// 启用SQL结果缓存功能，主要用于指示缓存持续时间（尤其是0）的使用。 md5:426f2265f5437a86
# <翻译结束>


<原文开始>
// Cache option for query statement.
<原文结束>

# <翻译开始>
// 查询语句的缓存选项。 md5:0243bce17a4463a8
# <翻译结束>


<原文开始>
// Hook functions for model hook feature.
<原文结束>

# <翻译开始>
// 用于模型钩子功能的钩子函数。 md5:cb10889f174ab53d
# <翻译结束>


<原文开始>
// Disables soft deleting features when select/delete operations.
<原文结束>

# <翻译开始>
// 在进行选择/删除操作时，禁用软删除特性。 md5:b6cc5bc9aefe18bf
# <翻译结束>


<原文开始>
// If true, it clones and returns a new model object whenever operation done; or else it changes the attribute of current model.
<原文结束>

# <翻译开始>
// 如果为真，每次操作后都会克隆并返回一个新的模型对象；否则，它将修改当前模型的属性。 md5:b4a3ad6d8438d2de
# <翻译结束>


<原文开始>
// onDuplicate is used for on Upsert clause.
<原文结束>

# <翻译开始>
// onDuplicate 用于 upsert 子句。 md5:e139824b32378802
# <翻译结束>


<原文开始>
// onDuplicateEx is used for excluding some columns on Upsert clause.
<原文结束>

# <翻译开始>
// onDuplicateEx 用于在 Upsert 子句中排除某些列。 md5:f985786a6831d9ec
# <翻译结束>


<原文开始>
// onConflict is used for conflict keys on Upsert clause.
<原文结束>

# <翻译开始>
// onConflict 用于在 Upsert 子句中处理冲突键。 md5:ec57ba30c97c0bd2
# <翻译结束>


<原文开始>
// Table alias to true table name, usually used in join statements.
<原文结束>

# <翻译开始>
// 表别名到真实表名的映射，通常在JOIN语句中使用。 md5:5951bd1c3aa8b870
# <翻译结束>


<原文开始>
// SoftTimeOption is the option to customize soft time feature for Model.
<原文结束>

# <翻译开始>
// SoftTimeOption 是用于自定义 Model 的软时间功能的选项。 md5:fcc19f5ef8ad45e7
# <翻译结束>


<原文开始>
// ModelHandler is a function that handles given Model and returns a new Model that is custom modified.
<原文结束>

# <翻译开始>
// Handler calls each of `handlers` on current Model and returns a new Model.
// ModelHandler 是一个处理给定 Model 并返回一个自定义修改后的新 Model 的函数。 md5:a02af46ff8fb2568
# <翻译结束>


<原文开始>
// ChunkHandler is a function that is used in function Chunk, which handles given Result and error.
// It returns true if it wants to continue chunking, or else it returns false to stop chunking.
<原文结束>

# <翻译开始>
// ChunkHandler 是一个函数，用于 Chunk 函数中，负责处理给定的结果和错误。
// 如果希望继续分块处理，则返回true；否则返回false以停止分块。 md5:e7b2a1b4761ac415
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
// Model 根据给定的模式创建并返回一个新的ORM模型。
// 参数 `tableNameQueryOrStruct` 可以是多个表名，也可以是别名，例如：
//  1. 模型名称：
//     db.Model("user")
//     db.Model("user u")
//     db.Model("user, user_detail")
//     db.Model("user u, user_detail ud")
//  2. 带别名的模型名称：
//     db.Model("user", "u")
//  3. 带子查询的模型名称：
//     db.Model("? AS a, ? AS b", subQuery1, subQuery2) md5:add855a912a9b6ef
# <翻译结束>


<原文开始>
// Model creation with sub-query.
<原文结束>

# <翻译开始>
	// 使用子查询创建模型。 md5:1c8112f948bca053
# <翻译结束>


<原文开始>
// Raw creates and returns a model based on a raw sql not a table.
// Example:
//
//	db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
<原文结束>

# <翻译开始>
// Raw根据原始SQL（而不是表）创建并返回一个模型。示例：
//
//	db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result) md5:0865d39f2ab854cb
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
// 示例：
//
//	db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
//
// 参见 Core.Raw。 md5:ced75308536ddfff
# <翻译结束>


<原文开始>
// With creates and returns an ORM model based on metadata of given object.
<原文结束>

# <翻译开始>
// With 根据给定对象的元数据创建并返回一个ORM模型。 md5:18604e26c0c946fb
# <翻译结束>


<原文开始>
// Partition sets Partition name.
// Example:
// dao.User.Ctx(ctx).Partition（"p1","p2","p3").All()
<原文结束>

# <翻译开始>
// 分区设置分区名称。
// 例子：
// dao.User.Ctx(ctx).Partition("p1", "p2", "p3").All()
//
// 这段Go代码的注释表示：使用`Partition`方法对数据进行分区操作，传入多个分区名称（如："p1", "p2", "p3"），然后在查询时指定这些分区。`Ctx(ctx)`表示使用上下文`ctx`进行操作。`All()`是获取所有满足条件的数据。 md5:f133a577ba31c05f
# <翻译结束>


<原文开始>
// Model acts like Core.Model except it operates on transaction.
// See Core.Model.
<原文结束>

# <翻译开始>
// Model 类似于 Core.Model，但它是基于事务操作的。
// 请参阅 Core.Model。 md5:2c5866afc2e5dd90
# <翻译结束>


<原文开始>
// With acts like Core.With except it operates on transaction.
// See Core.With.
<原文结束>

# <翻译开始>
// With 的行为类似于 Core.With，但它是在事务上操作。
// 参见 Core.With。 md5:37000d6ea41561fc
# <翻译结束>


<原文开始>
// Ctx sets the context for current operation.
<原文结束>

# <翻译开始>
// Ctx 设置当前操作的上下文。 md5:77d589f34a65753b
# <翻译结束>


<原文开始>
// GetCtx returns the context for current Model.
// It returns `context.Background()` is there's no context previously set.
<原文结束>

# <翻译开始>
// GetCtx返回当前Model的上下文。
// 如果之前没有设置上下文，则返回`context.Background()`。 md5:48edd9b438a38523
# <翻译结束>


<原文开始>
// As sets an alias name for current table.
<原文结束>

# <翻译开始>
// As 设置当前表的别名名称。 md5:c28e3f79c6fe2e48
# <翻译结束>


<原文开始>
// DB sets/changes the db object for current operation.
<原文结束>

# <翻译开始>
// DB 为当前操作设置或更改 db 对象。 md5:1761cc3b00f1d6bb
# <翻译结束>


<原文开始>
// TX sets/changes the transaction for current operation.
<原文结束>

# <翻译开始>
// TX 设置或更改当前操作的事务。 md5:7171a26d8d2d8431
# <翻译结束>


<原文开始>
// Schema sets the schema for current operation.
<原文结束>

# <翻译开始>
// Schema 设置当前操作的模式。 md5:723e31c5f24ff604
# <翻译结束>


<原文开始>
// Clone creates and returns a new model which is a Clone of current model.
// Note that it uses deep-copy for the Clone.
<原文结束>

# <翻译开始>
// Clone 创建并返回一个新的模型，它是当前模型的克隆。请注意，它使用深拷贝进行克隆。 md5:27e973f2f4fb42b3
# <翻译结束>


<原文开始>
// WhereBuilder copy, note the attribute pointer.
<原文结束>

# <翻译开始>
	// WhereBuilder 的复制方法，注意属性是指针。 md5:c9aa75059eb72059
# <翻译结束>


<原文开始>
// Shallow copy slice attributes.
<原文结束>

# <翻译开始>
	// 浅复制切片属性。 md5:d03df5f661b330b7
# <翻译结束>


<原文开始>
// Master marks the following operation on master node.
<原文结束>

# <翻译开始>
// Master 在主节点上标记以下操作。 md5:86cff0c5fb8d6d5d
# <翻译结束>


<原文开始>
// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
<原文结束>

# <翻译开始>
// Slave 在 slave 节点上标记以下操作。
// 请注意，只有在配置了 slave 节点的情况下，此注释才有意义。 md5:3d6dbca5bafb9cdf
# <翻译结束>


<原文开始>
// Safe marks this model safe or unsafe. If safe is true, it clones and returns a new model object
// whenever the operation done, or else it changes the attribute of current model.
<原文结束>

# <翻译开始>
// Safe 标记此模型为安全或不安全。如果 safe 为 true，那么在执行完操作后，它会克隆并返回一个新的模型对象；
// 否则，它将直接修改当前模型的属性。 md5:56aecad30556ca98
# <翻译结束>


<原文开始>
// Args sets custom arguments for model operation.
<原文结束>

# <翻译开始>
// Args 为模型操作设置自定义参数。 md5:6cf507acdf0e2401
# <翻译结束>


<原文开始>
// Handler calls each of `handlers` on current Model and returns a new Model.
// ModelHandler is a function that handles given Model and returns a new Model that is custom modified.
<原文结束>

# <翻译开始>
// Handler 对当前 Model 上的每个 `handlers` 进行调用，并返回一个新的经过自定义修改的 Model。
// ModelHandler 是一个处理给定 Model 并返回一个新的定制修改的 Model 的函数。
// md5:1ddef37520be2b6b
# <翻译结束>

