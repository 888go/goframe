
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
	// Sequence tag: []sequence tag
	// Sequence has order for error results.
<原文结束>

# <翻译开始>
// 序列标签: []序列标签
// 序列对错误结果有顺序要求。
# <翻译结束>


<原文开始>
					// If length of custom messages is lesser than length of rules,
					// the rest rules use the default error messages.
<原文结束>

# <翻译开始>
// 如果自定义消息的长度小于规则的长度，
// 剩余的规则将使用默认错误消息。
# <翻译结束>


<原文开始>
// No sequence rules: map[field]rule
<原文结束>

# <翻译开始>
// 无序规则：map[field]rule
// （这段代码注释表明，该处定义了一个无序的映射关系，其中键（key）为field，值（value）为rule。在Go语言中，"map[field]rule"代表一个映射类型，其键和值分别为field类型和rule类型，且这个映射中的元素没有特定顺序。）
# <翻译结束>


<原文开始>
	// It checks the struct recursively if its attribute is an embedded struct.
	// Ignore inputParamMap, assoc, rules and messages from parent.
<原文结束>

# <翻译开始>
// 它递归地检查结构体，如果其属性是一个嵌入式结构体。
// 忽略来自父级的inputParamMap、assoc、rules和messages。
# <翻译结束>


<原文开始>
// The following logic is the same as some of CheckStruct but without sequence support.
<原文结束>

# <翻译开始>
// 下面的逻辑与 CheckStruct 的部分功能相同，但不支持顺序检查。
# <翻译结束>


<原文开始>
// It checks each rule and its value in loop.
<原文结束>

# <翻译开始>
// 它在循环中检查每一条规则及其对应的值。
# <翻译结束>


<原文开始>
			// ===========================================================
			// Only in map and struct validations:
			// If value is nil or empty string and has no required* rules,
			// it clears the error message.
			// ===========================================================
<原文结束>

# <翻译开始>
// ===========================================================
// 仅在map和结构体验证中：
// 如果值为nil或空字符串且没有required*规则，
// 它将清除错误消息。
// ===========================================================
# <翻译结束>


<原文开始>
// map[RuleKey]ErrorMsg.
<原文结束>

# <翻译开始>
// map[规则键]错误信息。
# <翻译结束>

