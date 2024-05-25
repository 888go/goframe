
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
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// CacheOption is options for model cache control in query.
<原文结束>

# <翻译开始>
// CacheOption是查询中用于模型缓存控制的选项。 md5:8a833b8335c45455
# <翻译结束>


<原文开始>
	// Duration is the TTL for the cache.
	// If the parameter `Duration` < 0, which means it clear the cache with given `Name`.
	// If the parameter `Duration` = 0, which means it never expires.
	// If the parameter `Duration` > 0, which means it expires after `Duration`.
<原文结束>

# <翻译开始>
// Duration 是缓存的过期时间。
// 如果参数 `Duration` 小于 0，表示使用给定的 `Name` 清除缓存。
// 如果参数 `Duration` 等于 0，表示永不过期。
// 如果参数 `Duration` 大于 0，表示在 `Duration` 秒后过期。
// md5:28707300732ac411
# <翻译结束>


<原文开始>
	// Name is an optional unique name for the cache.
	// The Name is used to bind a name to the cache, which means you can later control the cache
	// like changing the `duration` or clearing the cache with specified Name.
<原文结束>

# <翻译开始>
// Name 是一个可选的唯一名称，用于标识缓存。
// 通过 Name 可以将一个名称与缓存绑定，这意味着您之后可以根据指定的名称来控制该缓存，
// 例如更改缓存的 `持续时间` 或者清除指定名称的缓存。
// md5:8c2eeafa42d36067
# <翻译结束>


<原文开始>
	// Force caches the query result whatever the result is nil or not.
	// It is used to avoid Cache Penetration.
<原文结束>

# <翻译开始>
// 强制缓存查询结果，无论结果是否为nil。
// 这用于防止缓存穿透。
// md5:78fc7d8520d64954
# <翻译结束>


<原文开始>
// selectCacheItem is the cache item for SELECT statement result.
<原文结束>

# <翻译开始>
// selectCacheItem是用于SELECT语句结果的缓存项。 md5:73fb34eaa64ea7d1
# <翻译结束>


<原文开始>
// Sql result of SELECT statement.
<原文结束>

# <翻译开始>
// SELECT语句的SQL结果。 md5:1f098617a374fffc
# <翻译结束>


<原文开始>
// The first column name of result, for Value/Count functions.
<原文结束>

# <翻译开始>
// 结果的第一列名称，用于Value/Count函数。 md5:2c091aca88ae5aa3
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
// Cache 为模型设置缓存功能。它将 SQL 的结果缓存起来，这意味着如果有相同的 SQL 请求，
// 它会直接从缓存中读取并返回结果，而不会真正提交并执行到数据库中。
//
// 注意，如果模型在事务中执行 SELECT 语句，缓存功能将被禁用。
// md5:5d7ea513a485f3ad
# <翻译结束>


<原文开始>
// checkAndRemoveSelectCache checks and removes the cache in insert/update/delete statement if
// cache feature is enabled.
<原文结束>

# <翻译开始>
// checkAndRemoveSelectCache 检查并移除插入/更新/删除语句中的缓存，如果启用了缓存功能。
// md5:7247a2e1e2e19e4b
# <翻译结束>


<原文开始>
// Special handler for Value/Count operations result.
<原文结束>

# <翻译开始>
// 对Value/Count操作结果的特殊处理器。 md5:beba69dc2347fa3a
# <翻译结束>


<原文开始>
// In case of Cache Penetration.
<原文结束>

# <翻译开始>
// 针对缓存穿透的情况。 md5:1464372279e61a7d
# <翻译结束>

