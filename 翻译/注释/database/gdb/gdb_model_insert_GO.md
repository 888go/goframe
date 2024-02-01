
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
// Batch sets the batch operation number for the model.
<原文结束>

# <翻译开始>
// Batch 设置模型的批量操作数量。
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
// Data 设置模型操作的数据。
// 参数 `data` 可以为 string、map、gmap、slice、struct、*struct 等类型。
// 注意，如果 `data` 为 map 或 slice 类型，会采用浅值复制的方式来避免在函数内部对原数据进行修改。
// 示例：
// Data("uid=10000")
// Data("uid", 10000)
// Data("uid=? AND name=?", 10000, "john")
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"}})
# <翻译结束>


<原文开始>
					// If the `data` parameter is a DO struct,
					// it then adds `OmitNilData` option for this condition,
					// which will filter all nil parameters in `data`.
<原文结束>

# <翻译开始>
// 如果`data`参数是一个DO结构体，
// 则为此条件添加`OmitNilData`选项，
// 这将会过滤掉`data`中所有为nil的参数。
# <翻译结束>


<原文开始>
				// If the `data` parameter is a DO struct,
				// it then adds `OmitNilData` option for this condition,
				// which will filter all nil parameters in `data`.
<原文结束>

# <翻译开始>
// 如果`data`参数是一个DO结构体，
// 则为此条件添加`OmitNilData`选项，
// 这将会过滤掉`data`中所有的nil参数。
# <翻译结束>


<原文开始>
// OnDuplicate sets the operations when columns conflicts occurs.
// In MySQL, this is used for "ON DUPLICATE KEY UPDATE" statement.
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
// OnDuplicate 设置在列冲突发生时的操作。
// 在 MySQL 中，此方法用于“ON DUPLICATE KEY UPDATE”语句。
// 参数 `onDuplicate` 可以为 string/Raw/*Raw/map/slice 类型。
// 示例：
//
// OnDuplicate("nickname, age") // 设置当主键重复时更新nickname和age字段
// OnDuplicate("nickname", "age") // 同上，以逗号分隔多个字段名
//
//	OnDuplicate(g.Map{
//		  "nickname": gdb.Raw("CONCAT('name_', VALUES(`nickname`))"), // 使用原始SQL表达式更新nickname字段
//	})
//
//	OnDuplicate(g.Map{
//		  "nickname": "passport", // 当主键重复时，将nickname字段的值设置为passport字段的值
//	}).
# <翻译结束>


<原文开始>
// OnDuplicateEx sets the excluding columns for operations when columns conflict occurs.
// In MySQL, this is used for "ON DUPLICATE KEY UPDATE" statement.
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
// OnDuplicateEx 设置在列冲突发生时操作的排除列。
// 在 MySQL 中，此函数用于 "ON DUPLICATE KEY UPDATE" 语句。
// 参数 `onDuplicateEx` 可以是字符串、映射或切片类型。
// 示例：
//
// OnDuplicateEx("passport, password") // 传入一个包含列名的字符串
// OnDuplicateEx("passport", "password") // 分别指定列名参数
//
//	OnDuplicateEx(g.Map{
//		  "passport": "",
//		  "password": "",
//	}) // 通过映射传入选定列名和其对应的更新值（此处为空字符串）
# <翻译结束>


<原文开始>
// Insert does "INSERT INTO ..." statement for the model.
// The optional parameter `data` is the same as the parameter of Model.Data function,
// see Model.Data.
<原文结束>

# <翻译开始>
// Insert 执行针对模型的 "INSERT INTO ..." 语句。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 请参考 Model.Data。
# <翻译结束>


<原文开始>
// InsertAndGetId performs action Insert and returns the last insert id that automatically generated.
<原文结束>

# <翻译开始>
// InsertAndGetId 执行插入操作，并返回自动生成的最后一个插入ID。
# <翻译结束>


<原文开始>
// InsertIgnore does "INSERT IGNORE INTO ..." statement for the model.
// The optional parameter `data` is the same as the parameter of Model.Data function,
// see Model.Data.
<原文结束>

# <翻译开始>
// InsertIgnore 执行针对模型的 "INSERT IGNORE INTO ..." 语句。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 请参阅 Model.Data。
# <翻译结束>


<原文开始>
// Replace does "REPLACE INTO ..." statement for the model.
// The optional parameter `data` is the same as the parameter of Model.Data function,
// see Model.Data.
<原文结束>

# <翻译开始>
// Replace 执行针对模型的 "REPLACE INTO ..." 语句。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 详情请参阅 Model.Data。
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
// Save 执行 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句，针对给定的 model。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 请参阅 Model.Data。
//
// 如果保存的数据中存在主键或唯一索引，则更新记录，
// 否则会在表中插入一条新记录。
# <翻译结束>


<原文开始>
// doInsertWithOption inserts data with option parameter.
<原文结束>

# <翻译开始>
// doInsertWithOption 使用选项参数插入数据。
# <翻译结束>


<原文开始>
// It converts any data to List type for inserting.
<原文结束>

# <翻译开始>
// 它将任何数据转换为 List 类型以便插入。
# <翻译结束>


<原文开始>
		// It uses gconv.Map here to simply fo the type converting from interface{} to map[string]interface{},
		// as there's another MapOrStructToMapDeep in next logic to do the deep converting.
<原文结束>

# <翻译开始>
// 这里使用gconv.Map简化从interface{}到map[string]interface{}的类型转换，
// 因为接下来的逻辑中会使用MapOrStructToMapDeep进行深度转换。
# <翻译结束>


<原文开始>
// If it's slice type, it then converts it to List type.
<原文结束>

# <翻译开始>
// 如果它是切片类型，那么将其转换为 List 类型。
# <翻译结束>


<原文开始>
// Automatic handling for creating/updating time.
<原文结束>

# <翻译开始>
// 自动处理创建/更新时间。
# <翻译结束>


<原文开始>
// Format DoInsertOption, especially for "ON DUPLICATE KEY UPDATE" statement.
<原文结束>

# <翻译开始>
// 格式化DoInsertOption，特别是用于“ON DUPLICATE KEY UPDATE”语句。
# <翻译结束>

