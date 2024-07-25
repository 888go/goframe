
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
// Batch sets the batch operation number for the model.
<原文结束>

# <翻译开始>
// Batch 为模型设置批处理操作的数量。 md5:7ae8528d1f8ac604
# <翻译结束>


<原文开始>
// Data sets the operation data for the model.
// The parameter `data` can be type of string/map/gmap/slice/struct/*struct, etc.
// Note that, it uses shallow value copying for `data` if `data` is type of map/slice
// to avoid changing it inside function.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data("uid=? AND name=?", 10000, "john")
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"}).
<原文结束>

# <翻译开始>
// Data 设置模型的操作数据。
// 参数 `data` 可以为字符串/映射/gmap/切片/结构体/**结构体指针**等类型。
// 注意，如果`data`是映射或切片类型，它将使用浅复制以避免在函数内部改变原数据。
// 例如：
// Data("uid=10000")
// Data("uid", 10000)
// Data("uid=? AND name=?", 10000, "john")
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"}}) md5:116cf94880dfa535
# <翻译结束>


<原文开始>
					// If the `data` parameter is a DO struct,
					// it then adds `OmitNilData` option for this condition,
					// which will filter all nil parameters in `data`.
<原文结束>

# <翻译开始>
					// 如果`data`参数是一个DO结构体，
					// 则为这个条件添加`OmitNilData`选项，
					// 这将过滤掉`data`中的所有空值参数。 md5:c978d65b6ea1129a
# <翻译结束>


<原文开始>
// OnConflict sets the primary key or index when columns conflicts occurs.
// It's not necessary for MySQL driver.
<原文结束>

# <翻译开始>
// OnConflict在列冲突时设置主键或索引。对于MySQL驱动程序来说，这通常是不必要的。 md5:30314cb75360b0e6
# <翻译结束>


<原文开始>
// OnDuplicate sets the operations when columns conflicts occurs.
// In MySQL, this is used for "ON DUPLICATE KEY UPDATE" statement.
// In PgSQL, this is used for "ON CONFLICT (id) DO UPDATE SET" statement.
// The parameter `onDuplicate` can be type of string/Raw/*Raw/map/slice.
// Example:
//
// OnDuplicate("nickname, age")
// OnDuplicate("nickname", "age")
//
//	OnDuplicate(g.Map{
//		  "nickname": gdb.Raw("CONCAT('name_', VALUES(`nickname`))"),
//	})
//
//	OnDuplicate(g.Map{
//		  "nickname": "passport",
//	}).
<原文结束>

# <翻译开始>
// OnDuplicate 设置在列发生冲突时执行的操作。
// 在MySQL中，这用于 "ON DUPLICATE KEY UPDATE" 语句。
// 在PgSQL中，这用于 "ON CONFLICT (id) DO UPDATE SET" 语句。
// 参数 `onDuplicate` 可以是字符串/Raw/*Raw/映射/切片类型。
// 示例：
//
// OnDuplicate("nickname, age")
// OnDuplicate("nickname", "age")
//
//	OnDuplicate(g.Map{
//		  "nickname": gdb.Raw("CONCAT('name_', VALUES(`nickname`))"),
//	})
//
//	OnDuplicate(g.Map{
//		  "nickname": "passport",
//	}) md5:fa9214f9681b4e5d
# <翻译结束>


<原文开始>
// OnDuplicateEx sets the excluding columns for operations when columns conflict occurs.
// In MySQL, this is used for "ON DUPLICATE KEY UPDATE" statement.
// In PgSQL, this is used for "ON CONFLICT (id) DO UPDATE SET" statement.
// The parameter `onDuplicateEx` can be type of string/map/slice.
// Example:
//
// OnDuplicateEx("passport, password")
// OnDuplicateEx("passport", "password")
//
//	OnDuplicateEx(g.Map{
//		  "passport": "",
//		  "password": "",
//	}).
<原文结束>

# <翻译开始>
// OnDuplicateEx 设置在发生列冲突时排除的列，用于操作。
// 在 MySQL 中，这用于 "ON DUPLICATE KEY UPDATE" 语句。
// 在 PgSQL 中，这用于 "ON CONFLICT (id) DO UPDATE SET" 语句。
// 参数 `onDuplicateEx` 可以是字符串、映射或切片类型。
// 示例：
//
// OnDuplicateEx("passport, password")
// OnDuplicateEx("passport", "password")
//
//	OnDuplicateEx(g.Map{
//		  "passport": "",
//		  "password": "",
//	}) md5:6fa8981bef042b71
# <翻译结束>


<原文开始>
// Insert does "INSERT INTO ..." statement for the model.
// The optional parameter `data` is the same as the parameter of Model.Data function,
// see Model.Data.
<原文结束>

# <翻译开始>
// Insert 为模型执行 "INSERT INTO ..." 语句。
// 可选参数 `data` 等同于 Model.Data 函数的参数，参见 Model.Data。 md5:9a6427cabf3ec194
# <翻译结束>


<原文开始>
// InsertAndGetId performs action Insert and returns the last insert id that automatically generated.
<原文结束>

# <翻译开始>
// InsertAndGetId 执行插入操作，并返回自动生成的最后一个插入id。 md5:8d00b40a35fa48a5
# <翻译结束>


<原文开始>
// InsertIgnore does "INSERT IGNORE INTO ..." statement for the model.
// The optional parameter `data` is the same as the parameter of Model.Data function,
// see Model.Data.
<原文结束>

# <翻译开始>
// InsertIgnore 为模型执行 "INSERT IGNORE INTO..." 语句。
// 可选参数 `data` 和 Model.Data 函数的参数相同，详情请参考 Model.Data。 md5:d6d8007d779bd324
# <翻译结束>


<原文开始>
// Replace does "REPLACE INTO ..." statement for the model.
// The optional parameter `data` is the same as the parameter of Model.Data function,
// see Model.Data.
<原文结束>

# <翻译开始>
// Replace 执行 "REPLACE INTO ..." 语句用于模型。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 请参阅 Model.Data。 md5:d5596c2470b6bcf4
# <翻译结束>


<原文开始>
// Save does "INSERT INTO ... ON DUPLICATE KEY UPDATE..." statement for the model.
// The optional parameter `data` is the same as the parameter of Model.Data function,
// see Model.Data.
//
// It updates the record if there's primary or unique index in the saving data,
// or else it inserts a new record into the table.
<原文结束>

# <翻译开始>
// Save 执行 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句，针对指定的模型。
// 可选参数 `data` 与 Model.Data 函数的参数相同，请参阅 Model.Data。
//
// 如果保存的数据中包含主键或唯一索引，它将更新记录；
// 否则，它会向表中插入一条新记录。 md5:9d87bd779f8f5acd
# <翻译结束>


<原文开始>
// doInsertWithOption inserts data with option parameter.
<原文结束>

# <翻译开始>
// doInsertWithOption 使用option参数插入数据。 md5:49dfb820e896850a
# <翻译结束>


<原文开始>
// m.data was already converted to type List/Map by function Data
<原文结束>

# <翻译开始>
	// m.data 已经通过 Data 函数转换为了 List/Map 类型. md5:cce9527c9f06deb0
# <翻译结束>


<原文开始>
// It converts any data to List type for inserting.
<原文结束>

# <翻译开始>
	// 它将任何数据转换为List类型以便插入。 md5:8e4e33863c8e1d24
# <翻译结束>


<原文开始>
// Automatic handling for creating/updating time.
<原文结束>

# <翻译开始>
	// 自动处理创建/更新时间。 md5:c45a07308954de68
# <翻译结束>


<原文开始>
// Format DoInsertOption, especially for "ON DUPLICATE KEY UPDATE" statement.
<原文结束>

# <翻译开始>
	// 格式化DoInsertOption，特别是针对“ON DUPLICATE KEY UPDATE”语句。 md5:e668e4c647415360
# <翻译结束>

