
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
// IsGNUVersion checks and returns whether given `version` is valid GNU version string.
<原文结束>

# <翻译开始>
// IsGNUVersion 检查并返回给定的 `version` 是否为有效的 GNU 版本字符串。 md5:6400dc6a399e4aa3
# <翻译结束>


<原文开始>
// CompareVersion compares `a` and `b` as standard GNU version.
//
// It returns  1 if `a` > `b`.
//
// It returns -1 if `a` < `b`.
//
// It returns  0 if `a` = `b`.
//
// GNU standard version is like:
// v1.0
// 1
// 1.0.0
// v1.0.1
// v2.10.8
// 10.2.0
// etc.
<原文结束>

# <翻译开始>
// CompareVersion 按照GNU版本标准比较 `a` 和 `b`。
//
// 如果 `a` 大于 `b`，返回 1。
//
// 如果 `a` 小于 `b`，返回 -1。
//
// 如果 `a` 等于 `b`，返回 0。
//
// GNU版本标准格式例如：
// v1.0
// 1
// 1.0.0
// v1.0.1
// v2.10.8
// 10.2.0
// 等等。
// md5:2716e579b3f9ba4d
# <翻译结束>


<原文开始>
// CompareVersionGo compares `a` and `b` as standard Golang version.
//
// It returns  1 if `a` > `b`.
//
// It returns -1 if `a` < `b`.
//
// It returns  0 if `a` = `b`.
//
// Golang standard version is like:
// 1.0.0
// v1.0.1
// v2.10.8
// 10.2.0
// v0.0.0-20190626092158-b2ccc519800e
// v1.12.2-0.20200413154443-b17e3a6804fa
// v4.20.0+incompatible
// etc.
//
// Docs: https://go.dev/doc/modules/version-numbers
<原文结束>

# <翻译开始>
// CompareVersionGo 将 `a` 和 `b` 当作标准的 Go 语言版本进行比较。
//
// 如果 `a` 大于 `b`，返回 1。
//
// 如果 `a` 小于 `b`，返回 -1。
//
// 如果 `a` 等于 `b`，返回 0。
//
// 标准的 Go 语言版本格式如下：
// 1.0.0
// v1.0.1
// v2.10.8
// 10.2.0
// v0.0.0-20190626092158-b2ccc519800e
// v1.12.2-0.20200413154443-b17e3a6804fa
// v4.20.0+incompatible
// 等等。
//
// 文档：https://go.dev/doc/modules/version-numbers
// md5:27f202ad306995b3
# <翻译结束>


<原文开始>
// check Major.Minor.Patch first
<原文结束>

# <翻译开始>
	// 首先检查Major.Minor.Patch. md5:098a0c10a68fabae
# <翻译结束>


<原文开始>
		// Specially in Golang:
		// "v1.12.2-0.20200413154443-b17e3a6804fa" < "v1.12.2"
		// "v1.12.3-0.20200413154443-b17e3a6804fa" > "v1.12.2"
<原文结束>

# <翻译开始>
		// 特别是在Go语言中：
		// 特别是：
		// "v1.12.2-0.20200413154443-b17e3a6804fa" < "v1.12.2" 		// 表示 v1.12.2-0.20200413154443-b17e3a6804fa 版本早于 v1.12.2
		// "v1.12.3-0.20200413154443-b17e3a6804fa" > "v1.12.2" 		// 表示 v1.12.3-0.20200413154443-b17e3a6804fa 版本晚于 v1.12.2
		// md5:685fe05f97473463
# <翻译结束>


<原文开始>
	// Specially in Golang:
	// "v4.20.1+incompatible" < "v4.20.1"
<原文结束>

# <翻译开始>
	// 特别是在 Golang 中：
	// "v4.20.1+incompatible" 小于 "v4.20.1"
	// md5:a292bd03375fd35c
# <翻译结束>

