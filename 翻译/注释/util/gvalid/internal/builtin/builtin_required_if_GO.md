
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
// RuleRequiredIf implements `required-if` rule:
// Required unless all given field and its value are equal.
//
// Format:  required-if:field,value,...
// Example: required-if: id,1,age,18
<原文结束>

# <翻译开始>
// RuleRequiredIf 实现了 `required-if` 规则：
// 当且仅当所有给定的字段及其对应的值相等时，该字段才是必填项。
//
// 格式： required-if:field,value,...
// 示例： required-if: id,1,age,18
// 这段代码注释描述了一个名为`RuleRequiredIf`的Go语言实现，它遵循一个自定义验证规则——`required-if`。这个规则表示某个字段只有在其他指定字段具有特定值时才需要（即为必填）。具体来说，规则的格式是在字符串中以逗号分隔指定字段名和它们对应的值，例如："required-if: id,1,age,18"意味着如果id不等于1或age不等于18，则当前字段是必需填写的。
# <翻译结束>


<原文开始>
// It supports multiple field and value pairs.
<原文结束>

# <翻译开始>
// 它支持多个字段和值对。
# <翻译结束>

