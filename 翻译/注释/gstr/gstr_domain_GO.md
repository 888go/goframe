
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
// IsSubDomain checks whether `subDomain` is sub-domain of mainDomain.
// It supports '*' in `mainDomain`.
<原文结束>

# <翻译开始>
// IsSubDomain 检查 `subDomain` 是否为 `mainDomain` 的子域名。
// 它支持在 `mainDomain` 中使用 '*'。
# <翻译结束>


<原文开始>
	// Eg:
	// "goframe.org" is not sub-domain of "s.goframe.org".
<原文结束>

# <翻译开始>
// 示例：
// "goframe.org" 不是 "s.goframe.org" 的子域名。
# <翻译结束>


<原文开始>
	// Eg:
	// "s.s.goframe.org" is not sub-domain of "*.goframe.org"
	// but
	// "s.s.goframe.org" is sub-domain of "goframe.org"
<原文结束>

# <翻译开始>
// 示例：
// "s.s.goframe.org" 不是 "*.goframe.org" 的子域名，
// 但是
// "s.s.goframe.org" 是 "goframe.org" 的子域名。
# <翻译结束>

