
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
// Fields appends `fieldNamesOrMapStruct` to the operation fields of the model, multiple fields joined using char ','.
// The parameter `fieldNamesOrMapStruct` can be type of string/map/*map/struct/*struct.
//
// Eg:
// Fields("id", "name", "age")
// Fields([]string{"id", "name", "age"})
// Fields(map[string]interface{}{"id":1, "name":"john", "age":18})
// Fields(User{ Id: 1, Name: "john", Age: 18}).
<原文结束>

# <翻译开始>
// Fields 用于指定需要操作的表字段，包括查询字段、写入字段、更新字段等过滤；
// 参数 `fieldNamesOrMapStruct` 可以是 string/map/*map/struct/*struct 类型。
// 多个字段之间使用字符 ',' 连接。
//
// 查询过滤示例: 
// g.Model("user").Fields("uid, nickname").Order("uid asc").All()  //SELECT `uid`,`nickname` FROM `user` ORDER BY `uid` asc
//
// 写入过滤示例:
// m := g.Map{
// "uid"      : 10000,
// "nickname" : "John Guo",
// "passport" : "john",
// "password" : "123456",
// }
// g.Model(table).Fields("nickname,passport,password").Data(m).Insert()  //INSERT INTO `user`(`nickname`,`passport`,`password`) VALUES('John Guo','john','123456')
//
// 示例：
// Fields("id", "name", "age")    // 通过字符串直接指定字段名
// Fields([]string{"id", "name", "age"})   // 通过字符串切片指定字段名
// Fields(map[string]interface{}{"id":1, "name":"john", "age":18})  // 通过键值对映射指定字段和值
// Fields(User{ Id: 1, Name: "john", Age: 18})   // 通过结构体实例指定字段和值
# <翻译结束>


<原文开始>
// FieldsPrefix performs as function Fields but add extra prefix for each field.
<原文结束>

# <翻译开始>
// FieldsPrefix 函数的功能与 Fields 相同，但会在每个字段前额外添加一个前缀。
# <翻译结束>


<原文开始>
// FieldsEx appends `fieldNamesOrMapStruct` to the excluded operation fields of the model,
// multiple fields joined using char ','.
// Note that this function supports only single table operations.
// The parameter `fieldNamesOrMapStruct` can be type of string/map/*map/struct/*struct.
//
// Also see Fields.
<原文结束>

# <翻译开始>
// FieldsEx 将`fieldNamesOrMapStruct` 追加到模型的排除操作字段列表中，
// 多个字段之间使用逗号字符 ',' 连接。
// 注意，此函数仅支持单表操作。
// 参数 `fieldNamesOrMapStruct` 可以是 string、map、*map 或 struct、*struct 类型。
// 请同时参考 Fields 函数。
//
// 查询排除过滤例子
// g.Model("user").FieldsEx("passport, password").All()  //SELECT `uid`,`nickname` FROM `user`
//
// 写入排除过滤例子
// m := g.Map{
// "uid"      : 10000,
// "nickname" : "John Guo",
// "passport" : "john",
// "password" : "123456",
// }
// g.Model(table).FieldsEx("uid").Data(m).Insert()  // INSERT INTO `user`(`nickname`,`passport`,`password`) VALUES('John Guo','john','123456')
# <翻译结束>


<原文开始>
// FieldsExPrefix performs as function FieldsEx but add extra prefix for each field.
<原文结束>

# <翻译开始>
// FieldsExPrefix 函数的功能与 FieldsEx 相同，但会在每个字段前额外添加一个前缀。
# <翻译结束>


<原文开始>
// FieldCount formats and appends commonly used field `COUNT(column)` to the select fields of model.
<原文结束>

# <翻译开始>
// FieldCount 格式化并追加计数字段别名到模型的 select 字段中。
// 简单点说就是追加一个计数的别名字段
//
// 追加计数字段例子:
// db.Model(table).Fields("id").FieldCount("id", "total")  // COUNT(`id`) AS `total`
# <翻译结束>


<原文开始>
// FieldSum formats and appends commonly used field `SUM(column)` to the select fields of model.
<原文结束>

# <翻译开始>
// FieldSum 格式化并追加常用字段 `SUM(column)` 到模型的 select 字段中。
// 简单点说就是追加一个求和的别名字段
//
// 追加求和字段例子:
// db.Model(table).Fields("column").FieldSum("column", "total")  // SUM(`column`) AS `total`
# <翻译结束>


<原文开始>
// FieldMin formats and appends commonly used field `MIN(column)` to the select fields of model.
<原文结束>

# <翻译开始>
// FieldMin 格式化并追加常用字段 `MIN(column)` 到模型的 select 字段中。
// 简单点说就是追加一个最小值的别名字段
//
// 追加最小值字段例子:
// db.Model(table).Fields("column").FieldMin("column", "total")  // MIN(`column`) AS `total`
# <翻译结束>


<原文开始>
// FieldMax formats and appends commonly used field `MAX(column)` to the select fields of model.
<原文结束>

# <翻译开始>
// FieldMax 格式化并追加常用字段 `MAX(column)` 到模型的 select 字段中。
// 简单点说就是追加一个最大值的别名字段
//
// 追加最大值字段例子:
// db.Model(table).Fields("column").FieldMax("column", "total")  // MAX(`column`) AS `total`
# <翻译结束>


<原文开始>
// FieldAvg formats and appends commonly used field `AVG(column)` to the select fields of model.
<原文结束>

# <翻译开始>
// FieldAvg 格式化并追加常用字段 `AVG(column)` 到模型的 select 字段中。
// 简单点说就是追加一个平均值的别名字段
//
// 追加平均值字段例子:
// db.Model(table).Fields("column").FieldAvg("column", "total")  // AVG(`column`) AS `total`
# <翻译结束>


<原文开始>
// GetFieldsStr retrieves and returns all fields from the table, joined with char ','.
// The optional parameter `prefix` specifies the prefix for each field, eg: GetFieldsStr("u.").
<原文结束>

# <翻译开始>
// GetFieldsStr 函数从表中检索并返回所有字段，各字段之间用字符 ',' 连接。
// 可选参数 `prefix` 用于指定每个字段的前缀，例如：GetFieldsStr("u.").
# <翻译结束>


<原文开始>
// GetFieldsExStr retrieves and returns fields which are not in parameter `fields` from the table,
// joined with char ','.
// The parameter `fields` specifies the fields that are excluded.
// The optional parameter `prefix` specifies the prefix for each field, eg: FieldsExStr("id", "u.").
<原文结束>

# <翻译开始>
// GetFieldsExStr 从表中检索并返回不在参数 `fields` 中的字段，并使用字符 ',' 连接这些字段。
// 参数 `fields` 指定需要排除的字段。
// 可选参数 `prefix` 用于指定每个字段的前缀，例如：FieldsExStr("id", "u.")。
# <翻译结束>


<原文开始>
// HasField determine whether the field exists in the table.
<原文结束>

# <翻译开始>
// HasField 判断字段是否在表中存在。
# <翻译结束>


<原文开始>
// getFieldsFrom retrieves, filters and returns fields name from table `table`.
<原文结束>

# <翻译开始>
// getFieldsFrom 从表 `table` 中检索、过滤并返回字段名称。
# <翻译结束>







<原文开始>
// It needs type asserting.
<原文结束>

# <翻译开始>
// 需要进行类型断言。
# <翻译结束>

