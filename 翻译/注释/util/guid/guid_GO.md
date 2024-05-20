
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
// Package guid provides simple and high performance unique id generation functionality.
<原文结束>

# <翻译开始>
// guid包提供了简单且高性能的唯一ID生成功能。. md5:22d1fe7516a2dff2
# <翻译结束>


<原文开始>
// Random chars string(36 bytes).
<原文结束>

# <翻译开始>
// 随机字符字符串（36个字节）。. md5:0e81d3c5a56e7cf2
# <翻译结束>


<原文开始>
// Sequence for unique purpose of current process.
<原文结束>

# <翻译开始>
// 当前进程独有的序列。. md5:2e6129c144b94c7b
# <翻译结束>


<原文开始>
// MAC addresses hash result in 7 bytes.
<原文结束>

# <翻译开始>
// MAC地址的哈希结果为7字节。. md5:99aa7e69b289dd55
# <翻译结束>


<原文开始>
// init initializes several fixed local variable.
<原文结束>

# <翻译开始>
// init 函数用于初始化几个固定的局部变量。. md5:3e44426e20423c37
# <翻译结束>


<原文开始>
// S creates and returns a global unique string in 32 bytes that meets most common
// usages without strict UUID algorithm. It returns a unique string using default
// unique algorithm if no `data` is given.
//
// The specified `data` can be no more than 2 parts. No matter how long each of the
// `data` size is, each of them will be hashed into 7 bytes as part of the result.
// If given `data` parts is less than 2, the leftover size of the result bytes will
// be token by random string.
//
// The returned string is composed with:
// 1. Default:    MACHash(7) + PID(4) + TimestampNano(12) + Sequence(3) + RandomString(6)
// 2. CustomData: DataHash(7/14) + TimestampNano(12) + Sequence(3) + RandomString(3/10)
//
// Note that：
//  1. The returned length is fixed to 32 bytes for performance purpose.
//  2. The custom parameter `data` composed should have unique attribute in your
//     business scenario.
<原文结束>

# <翻译开始>
// S 函数创建并返回一个32字节的全局唯一字符串，它满足大多数常见的使用需求，但不严格遵循UUID算法。如果没有提供`data`，则返回默认的唯一字符串。
//
// 指定的`data`最多可以有2个部分。无论每个`data`的长度多长，它们都将被哈希成7字节作为结果的一部分。如果给定的`data`部分少于2个，结果字节的剩余部分将由随机字符串填充。
//
// 返回的字符串由以下组成：
// 1. 默认：MAC哈希(7) + 进程ID(4) + 时间戳纳秒(12) + 序列号(3) + 随机字符串(6)
// 2. 自定义数据：Data哈希(7/14) + 时间戳纳秒(12) + 序列号(3) + 随机字符串(3/10)
//
// 注意：
//  1. 为了性能考虑，返回的长度固定为32字节。
//  2. 自定义参数`data`组合的内容在你的业务场景中应具有唯一性。
// md5:b09b2d34d56e1344
# <翻译结束>


<原文开始>
// Ignore empty data item bytes.
<原文结束>

# <翻译开始>
// 忽略空数据项字节。. md5:653aa2fb92e185e8
# <翻译结束>


<原文开始>
// getSequence increases and returns the sequence string in 3 bytes.
// The sequence is less than "zzz"(46655).
<原文结束>

# <翻译开始>
// getSequence 递增并返回一个以3个字节表示的序列字符串。
// 序列小于"zzz"(46655)。
// md5:742b11b09412718d
# <翻译结束>


<原文开始>
// getRandomStr randomly picks and returns `n` count of chars from randomStrBase.
<原文结束>

# <翻译开始>
// getRandomStr 从 randomStrBase 中随机选取并返回 `n` 个字符。. md5:fbef2b139ac9b42f
# <翻译结束>


<原文开始>
// getDataHashStr creates and returns hash bytes in 7 bytes with given data bytes.
<原文结束>

# <翻译开始>
// getDataHashStr 根据给定的数据字节创建并返回7字节的哈希字节。. md5:8947e4208efac1a4
# <翻译结束>

