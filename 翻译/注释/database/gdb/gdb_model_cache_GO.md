
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
// CacheOption is options for model cache control in query.
<原文结束>

# <翻译开始>
// CacheOption 是用于查询中模型缓存控制的选项。
# <翻译结束>


<原文开始>
	// Duration is the TTL for the cache.
	// If the parameter `Duration` < 0, which means it clear the cache with given `Name`.
	// If the parameter `Duration` = 0, which means it never expires.
	// If the parameter `Duration` > 0, which means it expires after `Duration`.
<原文结束>

# <翻译开始>
// Duration 是缓存的生存时间（TTL）。
// 如果参数 `Duration` < 0，表示按照给定的 `Name` 清除缓存。
// 如果参数 `Duration` = 0，表示缓存永不过期。
// 如果参数 `Duration` > 0，表示在 `Duration` 时间后缓存过期。
# <翻译结束>


<原文开始>
	// Name is an optional unique name for the cache.
	// The Name is used to bind a name to the cache, which means you can later control the cache
	// like changing the `duration` or clearing the cache with specified Name.
<原文结束>

# <翻译开始>
// Name 是缓存的一个可选的唯一名称。
// 名称用于将名称绑定到缓存，这意味着您可以在之后通过名称控制缓存，
// 例如：更改 `duration` 或清除指定名称的缓存。
# <翻译结束>


<原文开始>
	// Force caches the query result whatever the result is nil or not.
	// It is used to avoid Cache Penetration.
<原文结束>

# <翻译开始>
// Force无论查询结果是否为nil，都会缓存该查询结果。
// 它用于避免缓存穿透。
# <翻译结束>


<原文开始>
// selectCacheItem is the cache item for SELECT statement result.
<原文结束>

# <翻译开始>
// selectCacheItem 是用于 SELECT 语句结果的缓存项。
# <翻译结束>







<原文开始>
// The first column name of result, for Value/Count functions.
<原文结束>

# <翻译开始>
// 结果中的第一列名称，用于Value/Count函数。
# <翻译结束>


<原文开始>
// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// Note that, the cache feature is disabled if the model is performing select statement
// on a transaction.
<原文结束>

# <翻译开始>
// Cache 为模型设置缓存功能。它会缓存SQL查询的结果，这意味着
// 如果存在相同的SQL请求，它将直接从缓存读取并返回结果，
// 而不是提交并执行到数据库中。
//
// 注意，如果模型在事务中执行选择语句时，缓存功能是禁用的。
# <翻译结束>


<原文开始>
// checkAndRemoveSelectCache checks and removes the cache in insert/update/delete statement if
// cache feature is enabled.
<原文结束>

# <翻译开始>
// checkAndRemoveSelectCache 在缓存功能启用的情况下，检查并移除在插入/更新/删除语句中的缓存。
# <翻译结束>







<原文开始>
// Other cache, it needs conversion.
<原文结束>

# <翻译开始>
// 其他缓存，需要进行转换。
# <翻译结束>


<原文开始>
// Special handler for Value/Count operations result.
<原文结束>

# <翻译开始>
// 特殊处理 Value/Count 操作结果的处理器。
# <翻译结束>


<原文开始>
// In case of Cache Penetration.
<原文结束>

# <翻译开始>
// 在发生缓存穿透的情况下。
# <翻译结束>


<原文开始>
// Sql result of SELECT statement.
<原文结束>

# <翻译开始>
// Sql result of SELECT statement. （SQL语句中SELECT查询的结果。）
# <翻译结束>

