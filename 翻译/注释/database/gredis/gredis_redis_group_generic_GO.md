
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
// IGroupGeneric manages generic redis operations.
// Implements see redis.GroupGeneric.
<原文结束>

# <翻译开始>
// IGroupGeneric 管理通用的 Redis 操作。
// 实现了 redis.GroupGeneric 接口。
// md5:d6eb4921760b60f4
# <翻译结束>


<原文开始>
// CopyOption provides options for function Copy.
<原文结束>

# <翻译开始>
// CopyOption 为 Copy 函数提供选项。 md5:985df0dc4c62e896
# <翻译结束>


<原文开始>
// DB option allows specifying an alternative logical database index for the destination key.
<原文结束>

# <翻译开始>
// DB 选项允许为目的地键指定一个替代的逻辑数据库索引。 md5:f7752ecd2c09888e
# <翻译结束>


<原文开始>
// REPLACE option removes the destination key before copying the value to it.
<原文结束>

# <翻译开始>
// REPLACE选项在将值复制到目标键之前删除目标键。 md5:7d1daa6e1cf324ab
# <翻译结束>


<原文开始>
// ASYNC: flushes the databases asynchronously
<原文结束>

# <翻译开始>
// ASYNC：异步刷新数据库. md5:8f0fb503842c62dc
# <翻译结束>


<原文开始>
// SYNC: flushes the databases synchronously
<原文结束>

# <翻译开始>
// 同步：同步刷新数据库. md5:c995019017769085
# <翻译结束>


<原文开始>
// ExpireOption provides options for function Expire.
<原文结束>

# <翻译开始>
// ExpireOption 提供了 Expire 函数的选项。 md5:fe605b48792fd395
# <翻译结束>


<原文开始>
// NX -- Set expiry only when the key has no expiry
<原文结束>

# <翻译开始>
// NX -- 只在键没有过期时设置过期时间. md5:753349361957bc17
# <翻译结束>


<原文开始>
// XX -- Set expiry only when the key has an existing expiry
<原文结束>

# <翻译开始>
// XX -- 只在键已存在过期时间时设置过期. md5:005a0b6114104985
# <翻译结束>


<原文开始>
// GT -- Set expiry only when the new expiry is greater than current one
<原文结束>

# <翻译开始>
// GT -- 仅当新过期时间大于当前过期时间时，才设置过期时间. md5:e25f0e8a00a61ecf
# <翻译结束>


<原文开始>
// LT -- Set expiry only when the new expiry is less than current one
<原文结束>

# <翻译开始>
// LT -- 只有当新的过期时间小于当前过期时间时才设置过期时间. md5:7d837833fbcaa3f3
# <翻译结束>


<原文开始>
// ScanOption provides options for function Scan.
<原文结束>

# <翻译开始>
// ScanOption为Scan函数提供了选项。 md5:32efa528c8a65e49
# <翻译结束>


<原文开始>
// Match -- Specifies a glob-style pattern for filtering keys.
<原文结束>

# <翻译开始>
// Match - 定义用于筛选键的通配符风格模式。 md5:8a1fe0030e22d0f9
# <翻译结束>


<原文开始>
// Count -- Suggests the number of keys to return per scan.
<原文结束>

# <翻译开始>
// Count -- 建议每次扫描返回的键的数量。 md5:9090884e4078ad30
# <翻译结束>


<原文开始>
// Type -- Filters keys by their data type. Valid types are "string", "list", "set", "zset", "hash", and "stream".
<原文结束>

# <翻译开始>
// Type -- 根据键的数据类型过滤。有效的类型包括 "string"、"list"、"set"、"zset"、"hash" 和 "stream"。 md5:e1661eb01e6db304
# <翻译结束>


<原文开始>
// doScanOption is the internal representation of ScanOption.
<原文结束>

# <翻译开始>
// doScanOption是ScanOption的内部表示。 md5:3846dba237546aef
# <翻译结束>


<原文开始>
// ToUsedOption converts fields in ScanOption with zero values to nil. Only fields with values are retained.
<原文结束>

# <翻译开始>
// ToUsedOption 将ScanOption中的零值字段转换为nil。只有具有值的字段才会被保留。 md5:42a6307a3e94db33
# <翻译结束>

