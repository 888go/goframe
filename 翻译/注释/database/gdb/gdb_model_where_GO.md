
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://githum.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://githum.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// callWhereBuilder creates and returns a new Model, and sets its WhereBuilder if current Model is safe.
// It sets the WhereBuilder and returns current Model directly if it is not safe.
<原文结束>

# <翻译开始>
// callWhereBuilder 函数根据当前 Model 是否安全，创建并返回一个新的 Model。
// 如果当前 Model 是安全的，则设置其 WhereBuilder；如果不安全，则直接设置 WhereBuilder 并返回当前 Model。
# <翻译结束>


<原文开始>
// Where sets the condition statement for the builder. The parameter `where` can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// See WhereBuilder.Where.
<原文结束>

# <翻译开始>
// Where 为生成器设置条件语句。参数`where`可以是以下类型：
// string/map/gmap/slice/struct/*struct 等等。请注意，如果多次调用，
// 多个条件将会通过 "AND" 连接符合并到 where 语句中。
// 参见 WhereBuilder.Where 。
# <翻译结束>


<原文开始>
// Wheref builds condition string using fmt.Sprintf and arguments.
// Note that if the number of `args` is more than the placeholder in `format`,
// the extra `args` will be used as the where condition arguments of the Model.
// See WhereBuilder.Wheref.
<原文结束>

# <翻译开始>
// Wheref 通过 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果 `args` 的数量大于 `format` 中的占位符数量，
// 额外的 `args` 将作为 Model 的 where 条件参数使用。
// 参见 WhereBuilder.Wheref。
# <翻译结束>


<原文开始>
// WherePri does the same logic as Model.Where except that if the parameter `where`
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given `where` parameter as "123", the
// WherePri function treats the condition as "id=123", but Model.Where treats the condition
// as string "123".
// See WhereBuilder.WherePri.
<原文结束>

# <翻译开始>
// WherePri 执行的逻辑与 Model.Where 相同，但区别在于：如果参数 `where` 是一个单一条件，如 int、string、float 或 slice 类型，
// 那么它会将这个条件视为主键值。也就是说，如果主键是 "id"，给定的 `where` 参数为 "123"，那么 WherePri 函数会把条件当作 "id=123" 处理，
// 但是 Model.Where 函数会将该条件当作字符串 "123" 处理。
// 请参阅 WhereBuilder.WherePri。
# <翻译结束>


<原文开始>
// WhereLT builds `column < value` statement.
// See WhereBuilder.WhereLT.
<原文结束>

# <翻译开始>
// WhereLT构建`column < value`表达式语句。
// 参见WhereBuilder.WhereLT。
# <翻译结束>


<原文开始>
// WhereLTE builds `column <= value` statement.
// See WhereBuilder.WhereLTE.
<原文结束>

# <翻译开始>
// WhereLTE 用于构建 `column <= value` 条件语句。
// 参见 WhereBuilder.WhereLTE。
# <翻译结束>


<原文开始>
// WhereGT builds `column > value` statement.
// See WhereBuilder.WhereGT.
<原文结束>

# <翻译开始>
// WhereGT构建`column > value`语句。
// 请参阅WhereBuilder.WhereGT。
# <翻译结束>


<原文开始>
// WhereGTE builds `column >= value` statement.
// See WhereBuilder.WhereGTE.
<原文结束>

# <翻译开始>
// WhereGTE 用于构建 `column >= value` 条件语句。
// 参见 WhereBuilder.WhereGTE。
# <翻译结束>


<原文开始>
// WhereBetween builds `column BETWEEN min AND max` statement.
// See WhereBuilder.WhereBetween.
<原文结束>

# <翻译开始>
// WhereBetween 用于构建 `column BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereBetween。
# <翻译结束>


<原文开始>
// WhereLike builds `column LIKE like` statement.
// See WhereBuilder.WhereLike.
<原文结束>

# <翻译开始>
// WhereLike 用于构建 `column LIKE like` 语句。
// 参见 WhereBuilder.WhereLike。
# <翻译结束>


<原文开始>
// WhereIn builds `column IN (in)` statement.
// See WhereBuilder.WhereIn.
<原文结束>

# <翻译开始>
// WhereIn 构建 `column IN (in)` 语句。
// 参见 WhereBuilder.WhereIn。
# <翻译结束>


<原文开始>
// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
// See WhereBuilder.WhereNull.
<原文结束>

# <翻译开始>
// WhereNull 构建 `columns[0] IS NULL AND columns[1] IS NULL ...` 语句。
// 参见 WhereBuilder.WhereNull 方法。
# <翻译结束>


<原文开始>
// WhereNotBetween builds `column NOT BETWEEN min AND max` statement.
// See WhereBuilder.WhereNotBetween.
<原文结束>

# <翻译开始>
// WhereNotBetween 用于构建 `column NOT BETWEEN min AND max` 语句。
// 参见 WhereBuilder.WhereNotBetween。
# <翻译结束>


<原文开始>
// WhereNotLike builds `column NOT LIKE like` statement.
// See WhereBuilder.WhereNotLike.
<原文结束>

# <翻译开始>
// WhereNotLike 用于构建 `column NOT LIKE like` 语句。
// 参见 WhereBuilder.WhereNotLike。
# <翻译结束>


<原文开始>
// WhereNot builds `column != value` statement.
// See WhereBuilder.WhereNot.
<原文结束>

# <翻译开始>
// WhereNot 用于构建 `column != value` 条件语句。
// 请参阅 WhereBuilder.WhereNot。
# <翻译结束>


<原文开始>
// WhereNotIn builds `column NOT IN (in)` statement.
// See WhereBuilder.WhereNotIn.
<原文结束>

# <翻译开始>
// WhereNotIn 用于构建 `column NOT IN (in)` 语句。
// 请参阅 WhereBuilder.WhereNotIn。
# <翻译结束>


<原文开始>
// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
// See WhereBuilder.WhereNotNull.
<原文结束>

# <翻译开始>
// WhereNotNull 用于构建 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。
// 请参考 WhereBuilder.WhereNotNull。
# <翻译结束>

