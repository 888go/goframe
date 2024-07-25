
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
// doWhereType sets the condition statement for the model. The parameter `where` can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
<原文结束>

# <翻译开始>
// doWhereType 设置模型的条件语句。参数 `where` 可以是类型为 string、map、gmap、slice、struct 或其派生结构等。需要注意的是，如果多次调用此函数，多个条件将使用 "AND" 连接到 where 语句中。 md5:92ee322b44569cba
# <翻译结束>


<原文开始>
// doWherefType builds condition string using fmt.Sprintf and arguments.
// Note that if the number of `args` is more than the placeholder in `format`,
// the extra `args` will be used as the where condition arguments of the Model.
<原文结束>

# <翻译开始>
// doWherefType 使用 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果 `args` 的数量多于 `format` 中的占位符，多余的 `args` 将作为 Model 的 where 条件参数使用。 md5:67cfb01201c57037
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
// Where 方法为构建器设置条件语句。参数 `where` 可以是多种类型，包括
// 字符串、映射、gmap（可能是自定义的映射类型）、切片、结构体、指针到结构体等。需要注意的是，
// 如果该方法被调用了多次，多个条件将会使用 "AND" 连接起来组成 WHERE 子句。
// 例如：
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"}) // 使用结构体作为查询条件 md5:38a2f7ff889346c5
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
// Wheref 使用 fmt.Sprintf 和参数构建条件字符串。
// 注意，如果`args`的数目超过`format`中的占位符，
// 多余的`args`将作为Model的WHERE条件参数。
// 例如：
// Wheref(`amount<? and status=%s`, "paid", 100)  => WHERE `amount`<100 and status='paid'
// Wheref(`amount<%d and status=%s`, 100, "paid") => WHERE `amount`<100 and status='paid' md5:e4748efd7332202a
# <翻译结束>


<原文开始>
// WherePri does the same logic as Model.Where except that if the parameter `where`
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given `where` parameter as "123", the
// WherePri function treats the condition as "id=123", but Model.Where treats the condition
// as string "123".
<原文结束>

# <翻译开始>
// WherePri 的逻辑与 Model.Where 相同，但当参数 `where` 是单个条件，如 int、string、float 或 slice 时，它将该条件视为主键值。也就是说，如果主键是 "id" 并且给定的 `where` 参数为 "123"，WherePri 函数会将条件解析为 "id=123"，而 Model.Where 则会将条件视为字符串 "123"。 md5:2545fa57bcbd235c
# <翻译结束>


<原文开始>
// WhereLT builds `column < value` statement.
<原文结束>

# <翻译开始>
// WhereLT构建`column < value`语句。 md5:438e43e951037408
# <翻译结束>


<原文开始>
// WhereLTE builds `column <= value` statement.
<原文结束>

# <翻译开始>
// WhereLTE 构建 `column <= value` 的语句。 md5:149d7bc478d211fd
# <翻译结束>


<原文开始>
// WhereGT builds `column > value` statement.
<原文结束>

# <翻译开始>
// WhereGT 构建 `column > value` 语句。 md5:41527fa039c8a299
# <翻译结束>


<原文开始>
// WhereGTE builds `column >= value` statement.
<原文结束>

# <翻译开始>
// WhereGTE构建`column >= value`语句。 md5:fff159ae64237621
# <翻译结束>


<原文开始>
// WhereBetween builds `column BETWEEN min AND max` statement.
<原文结束>

# <翻译开始>
// WhereBetween构建`column BETWEEN min AND max`语句。 md5:cdb9b4a2942f3b60
# <翻译结束>


<原文开始>
// WhereLike builds `column LIKE like` statement.
<原文结束>

# <翻译开始>
// WhereLike 构建 `column LIKE like` 语句。 md5:5cf0790f9754307f
# <翻译结束>


<原文开始>
// WhereIn builds `column IN (in)` statement.
<原文结束>

# <翻译开始>
// WhereIn 构建 `column IN (in)` 语句。 md5:08648a50bb84e2ee
# <翻译结束>


<原文开始>
// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
<原文结束>

# <翻译开始>
// WhereNull 构建 `columns[0] IS NULL AND columns[1] IS NULL ...` 语句。 md5:9341218ae0c32357
# <翻译结束>


<原文开始>
// WhereNotBetween builds `column NOT BETWEEN min AND max` statement.
<原文结束>

# <翻译开始>
// WhereNotBetween构建`column NOT BETWEEN min AND max`语句。 md5:ac5d20d314a9fa0c
# <翻译结束>


<原文开始>
// WhereNotLike builds `column NOT LIKE like` statement.
<原文结束>

# <翻译开始>
// WhereNotLike 构建 `column NOT LIKE like` 的 SQL 语句。 md5:683105cb42e27e3b
# <翻译结束>


<原文开始>
// WhereNot builds `column != value` statement.
<原文结束>

# <翻译开始>
// WhereNot 构建 `column != value` 语句。 md5:d409867c3e8a9641
# <翻译结束>


<原文开始>
// WhereNotIn builds `column NOT IN (in)` statement.
<原文结束>

# <翻译开始>
// WhereNotIn构建`column NOT IN (in)`语句。 md5:658ffbae4d294fa4
# <翻译结束>


<原文开始>
// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
<原文结束>

# <翻译开始>
// WhereNotNull 构建 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。 md5:2444d6e2f6bcbf2d
# <翻译结束>

