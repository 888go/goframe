
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
// Order sets the "ORDER BY" statement for the model.
//
// Eg:
// Order("id desc")
// Order("id", "desc").
// Order("id desc,name asc")
// Order("id desc").Order("name asc")
// Order(gdb.Raw("field(id, 3,1,2)")).
<原文结束>

# <翻译开始>
// Order 为模型设置 "ORDER BY" 语句。
//
// 示例：
// Order("id desc") // 按 id 倒序排序
// Order("id", "desc") // 等同于 Order("id desc")
// Order("id desc,name asc") // 先按 id 倒序，再按 name 正序排序
// Order("id desc").Order("name asc") // 分别对 id 和 name 进行倒序和正序排序
// Order(gdb.Raw("field(id, 3,1,2)")) // 使用原生表达式进行排序，如MySQL中的 field 函数指定排序字段的顺序
# <翻译结束>


<原文开始>
// OrderAsc sets the "ORDER BY xxx ASC" statement for the model.
<原文结束>

# <翻译开始>
// OrderAsc 为模型设置 "ORDER BY xxx ASC" 语句。
# <翻译结束>


<原文开始>
// OrderDesc sets the "ORDER BY xxx DESC" statement for the model.
<原文结束>

# <翻译开始>
// OrderDesc 为模型设置 "ORDER BY xxx DESC" 语句。
# <翻译结束>


<原文开始>
// OrderRandom sets the "ORDER BY RANDOM()" statement for the model.
<原文结束>

# <翻译开始>
// OrderRandom 为模型设置 "ORDER BY RANDOM()" 语句。
# <翻译结束>


<原文开始>
// Group sets the "GROUP BY" statement for the model.
<原文结束>

# <翻译开始>
// Group 设置模型的 "GROUP BY" 语句。
# <翻译结束>

