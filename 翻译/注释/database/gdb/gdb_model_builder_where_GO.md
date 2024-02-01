
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
// doWhereType sets the condition statement for the model. The parameter `where` can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
<原文结束>

# <翻译开始>
// doWhereType 为模型设置条件语句。参数`where`可以是以下类型：
// string（字符串）、map（映射表）、gmap（Golang自定义的映射类型）、slice（切片）、struct（结构体）、*struct（指向结构体的指针）等。
// 注意，如果该方法被调用超过一次，则多次调用时传入的多个条件会通过"AND"运算符连接到一起形成最终的where条件语句。
# <翻译结束>


<原文开始>
// doWherefType builds condition string using fmt.Sprintf and arguments.
// Note that if the number of `args` is more than the placeholder in `format`,
// the extra `args` will be used as the where condition arguments of the Model.
<原文结束>

# <翻译开始>
// doWherefType 使用fmt.Sprintf和参数构建条件字符串。
// 注意，如果`args`的数量大于`format`中的占位符，
// 额外的`args`将被用作Model的where条件参数。
# <翻译结束>


<原文开始>
// Where sets the condition statement for the builder. The parameter `where` can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"}).
<原文结束>

# <翻译开始>
// Where 设置构建器的条件语句。参数`where`可以是以下类型：
// string/map/gmap/slice/struct/*struct 等等。请注意，如果调用该方法超过一次，
// 多个条件将会通过使用 "AND" 连接符合并到 where 语句中。
// 示例：
// Where("uid=10000") // 设置条件：uid为10000
// Where("uid", 10000) // 设置条件：uid为10000
// Where("money>? AND name like ?", 99999, "vip_%"）// 设置条件：money大于99999且name字段以"vip_"开头
// Where("uid", 1).Where("name", "john") // 设置条件：uid为1且name为"john"
// Where("status IN (?)", g.Slice{1,2,3}) // 设置条件：status在[1,2,3]范围内
// Where("age IN(?,?)", 18, 50) // 设置条件：age在[18,50]范围内
// Where(User{ Id : 1, UserName : "john"}) // 设置条件：根据User结构体内容（Id为1，UserName为"john"）
# <翻译结束>


<原文开始>
// Wheref builds condition string using fmt.Sprintf and arguments.
// Note that if the number of `args` is more than the placeholder in `format`,
// the extra `args` will be used as the where condition arguments of the Model.
// Eg:
// Wheref(`amount<? and status=%s`, "paid", 100)  => WHERE `amount`<100 and status='paid'
// Wheref(`amount<%d and status=%s`, 100, "paid") => WHERE `amount`<100 and status='paid'
<原文结束>

# <翻译开始>
// Wheref 通过 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果 `args` 的数量大于 `format` 中的占位符数量，
// 额外的 `args` 将作为 Model 的 where 条件参数使用。
// 示例：
// Wheref(`amount<? and status=%s`, "paid", 100)  => WHERE `amount`<100 and status='paid'
// Wheref(`amount<%d and status=%s`, 100, "paid") => WHERE `amount`<100 and status='paid'
# <翻译结束>


<原文开始>
// WherePri does the same logic as Model.Where except that if the parameter `where`
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given `where` parameter as "123", the
// WherePri function treats the condition as "id=123", but Model.Where treats the condition
// as string "123".
<原文结束>

# <翻译开始>
// WherePri 执行的逻辑与 Model.Where 相同，但有个特殊情况：
// 如果参数 `where` 是单个条件，如 int、string、float 或 slice 类型，
// 那么它会将这个条件视为主键值。也就是说，如果主键是 "id"，给定的 `where` 参数为 "123"，
// 那么 WherePri 函数会将条件处理为 "id=123"；
// 但是 Model.Where 函数则会将条件处理为字符串 "123"。
# <翻译结束>


<原文开始>
// WhereLT builds `column < value` statement.
<原文结束>

# <翻译开始>
// WhereLT 用于构建 `column < value` 语句。
# <翻译结束>


<原文开始>
// WhereLTE builds `column <= value` statement.
<原文结束>

# <翻译开始>
// WhereLTE 用于构建 `column <= value` 的表达式语句。
# <翻译结束>


<原文开始>
// WhereGT builds `column > value` statement.
<原文结束>

# <翻译开始>
// WhereGT 构建 `column > value` 语句。
# <翻译结束>


<原文开始>
// WhereGTE builds `column >= value` statement.
<原文结束>

# <翻译开始>
// WhereGTE 用于构建 `column >= value` 的表达式语句。
# <翻译结束>


<原文开始>
// WhereBetween builds `column BETWEEN min AND max` statement.
<原文结束>

# <翻译开始>
// WhereBetween 用于构建 `column BETWEEN min AND max` 语句。
# <翻译结束>


<原文开始>
// WhereLike builds `column LIKE like` statement.
<原文结束>

# <翻译开始>
// WhereLike 用于构建 `column LIKE like` 语句。
# <翻译结束>


<原文开始>
// WhereIn builds `column IN (in)` statement.
<原文结束>

# <翻译开始>
// WhereIn 构建 `column IN (in)` 语句。
# <翻译结束>


<原文开始>
// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
<原文结束>

# <翻译开始>
// WhereNull 用于构建 `columns[0] IS NULL AND columns[1] IS NULL ...` 语句。
# <翻译结束>


<原文开始>
// WhereNotBetween builds `column NOT BETWEEN min AND max` statement.
<原文结束>

# <翻译开始>
// WhereNotBetween 用于构建 `column NOT BETWEEN min AND max` 语句。
# <翻译结束>


<原文开始>
// WhereNotLike builds `column NOT LIKE like` statement.
<原文结束>

# <翻译开始>
// WhereNotLike 构建 `column NOT LIKE like` 语句。
# <翻译结束>


<原文开始>
// WhereNot builds `column != value` statement.
<原文结束>

# <翻译开始>
// WhereNot 用于构建 `column != value` 的表达式语句。
# <翻译结束>


<原文开始>
// WhereNotIn builds `column NOT IN (in)` statement.
<原文结束>

# <翻译开始>
// WhereNotIn 构造 `column NOT IN (in)` 语句。
# <翻译结束>


<原文开始>
// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
<原文结束>

# <翻译开始>
// WhereNotNull 用于构建 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。
# <翻译结束>

