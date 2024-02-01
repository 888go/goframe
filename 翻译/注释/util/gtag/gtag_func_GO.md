
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
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Set sets tag content for specified name.
// Note that it panics if `name` already exists.
<原文结束>

# <翻译开始>
// Set为指定名称设置标签内容。
// 注意，如果`name`已存在，则会引发panic。
# <翻译结束>


<原文开始>
// SetOver performs as Set, but it overwrites the old value if `name` already exists.
<原文结束>

# <翻译开始>
// SetOver 函数表现如同 Set，但当 `name` 已经存在时，它会覆盖旧的值。
# <翻译结束>


<原文开始>
// Sets sets multiple tag content by map.
<原文结束>

# <翻译开始>
// Sets 通过映射设置多个标签内容。
# <翻译结束>


<原文开始>
// SetsOver performs as Sets, but it overwrites the old value if `name` already exists.
<原文结束>

# <翻译开始>
// SetsOver 函数表现与 Sets 相同，但当 `name` 已经存在时，它会覆盖旧的值。
# <翻译结束>


<原文开始>
// Get retrieves and returns the stored tag content for specified name.
<原文结束>

# <翻译开始>
// Get 方法用于根据指定名称检索并返回存储的标签内容。
# <翻译结束>


<原文开始>
// Parse parses and returns the content by replacing all tag name variable to
// its content for given `content`.
// Eg:
// gtag.Set("demo", "content")
// Parse(`This is {demo}`) -> `This is content`.
<原文结束>

# <翻译开始>
// Parse函数解析并返回内容，将给定`content`中所有标签名称变量替换为它的实际内容。
// 示例：
// gtag.Set("demo", "content")
// Parse(`This is {demo}`) -> `This is content`。
# <翻译结束>

