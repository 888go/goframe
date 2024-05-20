
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
// Set sets tag content for specified name.
// Note that it panics if `name` already exists.
<原文结束>

# <翻译开始>
// Set 设置指定名称的标签内容。
// 请注意，如果`name`已经存在，该函数会引发恐慌。
// md5:3b301b4174b60616
# <翻译结束>


<原文开始>
// SetOver performs as Set, but it overwrites the old value if `name` already exists.
<原文结束>

# <翻译开始>
// SetOver 执行 Set 的功能，但如果 `name` 已经存在，它会覆盖旧的值。. md5:906ca9f516be44d0
# <翻译结束>


<原文开始>
// Sets sets multiple tag content by map.
<原文结束>

# <翻译开始>
// 通过map设置多个标签的内容。. md5:c02ae9dd9350cf50
# <翻译结束>


<原文开始>
// SetsOver performs as Sets, but it overwrites the old value if `name` already exists.
<原文结束>

# <翻译开始>
// SetsOver 的行为类似于 Sets，但如果 `name` 已经存在，它会覆盖旧值。. md5:6a87c6587ed9794f
# <翻译结束>


<原文开始>
// Get retrieves and returns the stored tag content for specified name.
<原文结束>

# <翻译开始>
// Get 获取并返回指定名称的存储标签内容。. md5:1a0a007cb18c41fa
# <翻译结束>


<原文开始>
// Parse parses and returns the content by replacing all tag name variable to
// its content for given `content`.
// Eg:
// gtag.Set("demo", "content")
// Parse(`This is {demo}`) -> `This is content`.
<原文结束>

# <翻译开始>
// Parse 通过将所有标签名变量替换为其内容，解析并返回给定的`content`。
// 示例：
// gtag.Set("demo", "content")
// Parse(`This is {demo}`) -> `This is content`。
// md5:b45c5273962c7662
# <翻译结束>

