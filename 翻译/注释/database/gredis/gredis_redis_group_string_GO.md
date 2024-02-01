
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
// IGroupString manages redis string operations.
// Implements see redis.GroupString.
<原文结束>

# <翻译开始>
// IGroupString 管理 Redis 字符串操作。
// 实现细节请参考 redis.GroupString。
# <翻译结束>


<原文开始>
// TTLOption provides extra option for TTL related functions.
<原文结束>

# <翻译开始>
// TTLOption 提供了与TTL相关函数的额外选项。
# <翻译结束>


<原文开始>
// EX seconds -- Set the specified expire time, in seconds.
<原文结束>

# <翻译开始>
// EX 秒数 -- 设置指定的过期时间，单位为秒。
# <翻译结束>


<原文开始>
// PX milliseconds -- Set the specified expire time, in milliseconds.
<原文结束>

# <翻译开始>
// PX milliseconds -- 设置指定的过期时间，单位为毫秒。
# <翻译结束>


<原文开始>
// EXAT timestamp-seconds -- Set the specified Unix time at which the key will expire, in seconds.
<原文结束>

# <翻译开始>
// EXAT 时间戳-秒 -- 设置键在指定的 Unix 时间（单位：秒）时过期。
# <翻译结束>


<原文开始>
// PXAT timestamp-milliseconds -- Set the specified Unix time at which the key will expire, in milliseconds.
<原文结束>

# <翻译开始>
// PXAT 时间戳-毫秒 -- 设置键的过期时间，以毫秒为单位，指定 Unix 时间。
# <翻译结束>


<原文开始>
// Retain the time to live associated with the key.
<原文结束>

# <翻译开始>
// 保留与键关联的生存时间（TTL）。
# <翻译结束>


<原文开始>
// SetOption provides extra option for Set function.
<原文结束>

# <翻译开始>
// SetOption为Set函数提供额外的选项。
# <翻译结束>


<原文开始>
// Only set the key if it does not already exist.
<原文结束>

# <翻译开始>
// 如果键尚未存在，则设置该键。
# <翻译结束>


<原文开始>
// Only set the key if it already exists.
<原文结束>

# <翻译开始>
// 只有在键已存在的时候才设置该键。
# <翻译结束>


<原文开始>
	// Return the old string stored at key, or nil if key did not exist.
	// An error is returned and SET aborted if the value stored at key is not a string.
<原文结束>

# <翻译开始>
// 如果键存在，则返回该键存储的旧字符串，否则返回nil。
// 若键中存储的值不是字符串，则会返回错误并中止SET操作。
# <翻译结束>


<原文开始>
// GetEXOption provides extra option for GetEx function.
<原文结束>

# <翻译开始>
// GetEXOption为GetEx函数提供额外的选项。
# <翻译结束>


<原文开始>
// Persist -- Remove the time to live associated with the key.
<原文结束>

# <翻译开始>
// Persist -- 移除与该键关联的生存时间（TTL）。
# <翻译结束>

