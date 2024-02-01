
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
// RuleRequiredUnless implements `required-unless` rule:
// Required unless all given field and its value are not equal.
//
// Format:  required-unless:field,value,...
// Example: required-unless:id,1,age,18
<原文结束>

# <翻译开始>
// RuleRequiredUnless 实现了 `required-unless` 规则：
// 当且仅当所有给定的字段及其对应值不相等时，该字段才是必填的。
//
// 格式： required-unless:field,value,...
// 示例： required-unless:id,1,age,18
// 这段代码注释描述了一个名为`RuleRequiredUnless`的Go语言规则实现，该规则用于表单验证或者其他数据校验场景。具体而言，某个字段只有在其它指定字段与其对应的值不相等的情况下，才被视为必填项。注释中给出了该规则的格式示例，表明需要按照 "字段名,字段值,..." 的形式来配置所需的条件。例如："required-unless:id,1,age,18" 表示如果 id 不为 1 且 age 不为 18，则当前字段是必需的。
# <翻译结束>


<原文开始>
// It supports multiple field and value pairs.
<原文结束>

# <翻译开始>
// 它支持多个字段和值对。
# <翻译结束>

