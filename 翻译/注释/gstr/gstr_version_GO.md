
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
// IsGNUVersion checks and returns whether given `version` is valid GNU version string.
<原文结束>

# <翻译开始>
// IsGNUVersion 检查并返回给定的 `version` 是否为有效的 GNU 版本字符串。
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
// CompareVersion 按照标准 GNU 版本格式比较 `a` 和 `b`。
//
// 如果 `a` > `b`，则返回 1。
//
// 如果 `a` < `b`，则返回 -1。
//
// 如果 `a` = `b`，则返回 0。
//
// 标准 GNU 版本格式例如：
// v1.0
// 1
// 1.0.0
// v1.0.1
// v2.10.8
// 10.2.0
// 等等。
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
// CompareVersionGo 按照 Golang 标准版本格式比较 `a` 和 `b`。
//
// 如果 `a` > `b`，则返回 1。
//
// 如果 `a` < `b`，则返回 -1。
//
// 如果 `a` = `b`，则返回 0。
//
// Golang 标准版本格式例如：
// 1.0.0
// v1.0.1
// v2.10.8
// 10.2.0
// v0.0.0-20190626092158-b2ccc519800e
// v1.12.2-0.20200413154443-b17e3a6804fa
// v4.20.0+incompatible
// 等等。
//
// 文档参考：https://go.dev/doc/modules/version-numbers
# <翻译结束>


<原文开始>
// check Major.Minor.Patch first
<原文结束>

# <翻译开始>
// 首先检查 Major.Minor.Patch
# <翻译结束>


<原文开始>
		// Specially in Golang:
		// "v1.12.2-0.20200413154443-b17e3a6804fa" < "v1.12.2"
		// "v1.12.3-0.20200413154443-b17e3a6804fa" > "v1.12.2"
<原文结束>

# <翻译开始>
// 特别地在 Golang 中：
// "v1.12.2-0.20200413154443-b17e3a6804fa" 小于 "v1.12.2"
// "v1.12.3-0.20200413154443-b17e3a6804fa" 大于 "v1.12.2"
// 这段代码注释是关于 Golang 中版本字符串比较的特殊规则：
// 在 Golang 中，对于包含预发布版本号（如 "-0.20200413154443-b17e3a6804fa"）的版本字符串，在进行字符串比较时，主版本号、次版本号和补丁版本号部分会被优先比较。当这部分相同时，带有预发布版本号的版本会认为小于不带预发布版本号的版本。
// 因此，尽管 "v1.12.2-0.20200413154443-b17e3a6804fa" 的主要部分与 "v1.12.2" 相同，但由于其附加了预发布标识，所以在比较中它被认为小于 "v1.12.2"。
// 同样，"v1.12.3-0.20200413154443-b17e3a6804fa" 由于其主版本号部分高于 "v1.12.2"，所以即使它也有预发布版本号，依然会在比较中大于 "v1.12.2"。
# <翻译结束>


<原文开始>
	// Specially in Golang:
	// "v4.20.1+incompatible" < "v4.20.1"
<原文结束>

# <翻译开始>
// 特别在 Golang 中：
// "v4.20.1+incompatible" < "v4.20.1"
// 这表示包含 "+incompatible" 后缀的版本号在比较时，会被视为小于不包含此后缀的相同主版本号和次版本号的版本。这是 Go 语言处理依赖版本的一种方式，"+incompatible" 后缀通常用于标记非模块化的包或不符合语义化版本控制规范的包。
# <翻译结束>

