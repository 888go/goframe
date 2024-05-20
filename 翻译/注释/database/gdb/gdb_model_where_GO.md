
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://githum.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码文件受MIT许可协议条款的约束。
// 如果您没有随此文件分发MIT许可证的副本，
// 您可以从https://gitee.com/gogf/gf获取。
// md5:358e3ba76264232a
# <翻译结束>


<原文开始>
// callWhereBuilder creates and returns a new Model, and sets its WhereBuilder if current Model is safe.
// It sets the WhereBuilder and returns current Model directly if it is not safe.
<原文结束>

# <翻译开始>
// callWhereBuilder 创建并返回一个新的Model实例，如果当前Model是安全的，则设置其WhereBuilder。
// 如果当前Model不是安全的，它将直接设置WhereBuilder并返回当前Model。
// md5:f847d7aad8140312
# <翻译结束>


<原文开始>
// Where sets the condition statement for the builder. The parameter `where` can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// See WhereBuilder.Where.
<原文结束>

# <翻译开始>
// Where 为构建器设置条件语句。参数 `where` 可以是
// 字符串/映射/gmap/切片/结构体/结构体指针等类型。注意，如果调用多次，
// 多个条件将使用 "AND" 连接成 WHERE 语句。
// 参见 WhereBuilder.Where。
// md5:d18e25934a430281
# <翻译结束>


<原文开始>
// Wheref builds condition string using fmt.Sprintf and arguments.
// Note that if the number of `args` is more than the placeholder in `format`,
// the extra `args` will be used as the where condition arguments of the Model.
// See WhereBuilder.Wheref.
<原文结束>

# <翻译开始>
// Wheref 使用 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果 `args` 的数量多于 `format` 中的占位符，
// 多余的 `args` 将用作 Model 的 WHERE 条件参数。
// 参见 WhereBuilder.Wheref。
// md5:ecb69c9051fee97d
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
// WherePri 的逻辑与 Model.Where 相同，但当参数 `where` 是一个单一的条件，如 int、string、float 或 slice 时，它会将条件视为主键值。也就是说，如果主键是 "id"，并且给定的 `where` 参数为 "123"，WherePri 函数会将条件解析为 "id=123"，而 Model.Where 则会将条件视为字符串 "123"。
// 参阅 WhereBuilder.WherePri。
// md5:13dc66b105f841e9
# <翻译结束>


<原文开始>
// WhereLT builds `column < value` statement.
// See WhereBuilder.WhereLT.
<原文结束>

# <翻译开始>
// WhereLT 用于构建 `column < value` 的语句。
// 参见 WhereBuilder.WhereLT。
// md5:16aa57cc63797f44
# <翻译结束>


<原文开始>
// WhereLTE builds `column <= value` statement.
// See WhereBuilder.WhereLTE.
<原文结束>

# <翻译开始>
// WhereLTE 构建 `column <= value` 条件语句。
// 参考 WhereBuilder.WhereLTE。
// md5:85440b32a52d1c2c
# <翻译结束>


<原文开始>
// WhereGT builds `column > value` statement.
// See WhereBuilder.WhereGT.
<原文结束>

# <翻译开始>
// WhereGT构建`column > value`语句。
// 参见WhereBuilder.WhereGT。
// md5:900e4d7769d650c2
# <翻译结束>


<原文开始>
// WhereGTE builds `column >= value` statement.
// See WhereBuilder.WhereGTE.
<原文结束>

# <翻译开始>
// WhereGTE构建`column >= value`语句。
// 参见WhereBuilder.WhereGTE。
// md5:cd051232a4d3707c
# <翻译结束>


<原文开始>
// WhereBetween builds `column BETWEEN min AND max` statement.
// See WhereBuilder.WhereBetween.
<原文结束>

# <翻译开始>
// WhereBetween 用于构建 `column BETWEEN min AND max` 的语句。
// 参见 WhereBuilder.WhereBetween 的使用。
// md5:88a499f60e180ae2
# <翻译结束>


<原文开始>
// WhereLike builds `column LIKE like` statement.
// See WhereBuilder.WhereLike.
<原文结束>

# <翻译开始>
// WhereLike 构建 `column LIKE like` 语句。
// 参考 WhereBuilder.WhereLike。
// md5:0d0b14277dfc3be8
# <翻译结束>


<原文开始>
// WhereIn builds `column IN (in)` statement.
// See WhereBuilder.WhereIn.
<原文结束>

# <翻译开始>
// WhereIn 构建 `column IN (in)` 语句。
// 参考 WhereBuilder.WhereIn。
// md5:b4b77eb17cf9b671
# <翻译结束>


<原文开始>
// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
// See WhereBuilder.WhereNull.
<原文结束>

# <翻译开始>
// WhereNull构建`columns[0] IS NULL AND columns[1] IS NULL ...`语句。
// 参见WhereBuilder.WhereNull。
// md5:af598d8379efcab6
# <翻译结束>


<原文开始>
// WhereNotBetween builds `column NOT BETWEEN min AND max` statement.
// See WhereBuilder.WhereNotBetween.
<原文结束>

# <翻译开始>
// WhereNotBetween 用于构建 `column NOT BETWEEN min AND max` 的SQL语句。
// 参见WhereBuilder.WhereNotBetween的用法。
// md5:be0e739db028ee79
# <翻译结束>


<原文开始>
// WhereNotLike builds `column NOT LIKE like` statement.
// See WhereBuilder.WhereNotLike.
<原文结束>

# <翻译开始>
// WhereNotLike 构建 `column NOT LIKE like` 语句。
// 参考 WhereBuilder.WhereNotLike。
// md5:1be6fd9ae98cc213
# <翻译结束>


<原文开始>
// WhereNot builds `column != value` statement.
// See WhereBuilder.WhereNot.
<原文结束>

# <翻译开始>
// WhereNot构建`column != value`语句。
// 参见WhereBuilder.WhereNot。
// md5:973b26be7332c7b2
# <翻译结束>


<原文开始>
// WhereNotIn builds `column NOT IN (in)` statement.
// See WhereBuilder.WhereNotIn.
<原文结束>

# <翻译开始>
// WhereNotIn构建`column NOT IN (in)`语句。
// 请参阅WhereBuilder.WhereNotIn。
// md5:4fbfcfa2b85a83d2
# <翻译结束>


<原文开始>
// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
// See WhereBuilder.WhereNotNull.
<原文结束>

# <翻译开始>
// WhereNotNull 构建 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。
// 参见 WhereBuilder.WhereNotNull 的用法。
// md5:05b9b4179d41a28b
# <翻译结束>

