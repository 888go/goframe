
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
// Package gres provides resource management and packing/unpacking feature between files and bytes.
<原文结束>

# <翻译开始>
// 包gres提供了文件和字节之间资源管理及打包/解包的功能。 md5:29e79f40a11fe941
# <翻译结束>


<原文开始>
// Separator for directories.
<原文结束>

# <翻译开始>
// 目录分隔符。 md5:a4832545f002edfe
# <翻译结束>


<原文开始>
// Default resource object.
<原文结束>

# <翻译开始>
// 默认资源对象。 md5:f02aee71ab2f8fc2
# <翻译结束>


<原文开始>
// Add unpacks and adds the `content` into the default resource object.
// The unnecessary parameter `prefix` indicates the prefix
// for each file storing into current resource object.
<原文结束>

# <翻译开始>
// Add 方法将 'content' 解包并添加到默认资源对象中。
// 不需要的参数 'prefix' 表示存储在当前资源对象中的每个文件的前缀。
// md5:3b5da05501708d4d
# <翻译结束>


<原文开始>
// Load loads, unpacks and adds the data from `path` into the default resource object.
// The unnecessary parameter `prefix` indicates the prefix
// for each file storing into current resource object.
<原文结束>

# <翻译开始>
// Load 从 `path` 加载、解包并将数据添加到默认资源对象中。
// 不必要的参数 `prefix` 表示存储到当前资源对象中的每个文件的前缀。
// md5:901d4f7e4d8bf0cf
# <翻译结束>


<原文开始>
// Get returns the file with given path.
<原文结束>

# <翻译开始>
// Get返回给定路径的文件。 md5:f4989a4832cde2d2
# <翻译结束>


<原文开始>
// GetWithIndex searches file with `path`, if the file is directory
// it then does index files searching under this directory.
//
// GetWithIndex is usually used for http static file service.
<原文结束>

# <翻译开始>
// GetWithIndex 在给定路径`path`下搜索文件。如果找到的是一个目录，它会在这个目录下索引文件进行搜索。
//
// GetWithIndex 通常用于HTTP静态文件服务中。
// md5:bfb61cc8920b4633
# <翻译结束>


<原文开始>
// GetContent directly returns the content of `path` in default resource object.
<原文结束>

# <翻译开始>
// GetContent 直接返回默认资源对象中 `path` 的内容。 md5:6446043ef668020c
# <翻译结束>


<原文开始>
// Contains checks whether the `path` exists in the default resource object.
<原文结束>

# <翻译开始>
// Contains 检查默认资源对象中是否存在 `path`。 md5:f69f9f792a33a089
# <翻译结束>


<原文开始>
// IsEmpty checks and returns whether the resource manager is empty.
<原文结束>

# <翻译开始>
// IsEmpty 检查资源管理器是否为空，并返回结果。 md5:3aaae27781ad4e8c
# <翻译结束>


<原文开始>
// ScanDir returns the files under the given path, the parameter `path` should be a folder type.
//
// The pattern parameter `pattern` supports multiple file name patterns,
// using the ',' symbol to separate multiple patterns.
//
// It scans directory recursively if given parameter `recursive` is true.
<原文结束>

# <翻译开始>
// ScanDir 在给定路径下返回文件，参数 `path` 应该是一个文件夹类型。
//
// `pattern` 参数支持多个文件名模式，使用逗号 `,` 来分隔多个模式。
//
// 如果 `recursive` 参数为真，它会递归扫描目录。
// md5:4726ded4e00ca75f
# <翻译结束>


<原文开始>
// ScanDirFile returns all sub-files with absolute paths of given `path`,
// It scans directory recursively if given parameter `recursive` is true.
//
// Note that it returns only files, exclusive of directories.
<原文结束>

# <翻译开始>
// ScanDirFile 返回给定`path`下的所有子文件的绝对路径，
// 如果给定的参数`recursive`为true，它会递归地扫描目录。
//
// 注意，它只返回文件，不包括目录。
// md5:0f3154c32271652b
# <翻译结束>


<原文开始>
// Export exports and saves specified path `src` and all its sub files to specified system path `dst` recursively.
<原文结束>

# <翻译开始>
// Export 将指定的路径 `src` 及其所有子文件递归导出并保存到指定的系统路径 `dst`。 md5:944ad6e86342817b
# <翻译结束>


<原文开始>
// Dump prints the files of the default resource object.
<原文结束>

# <翻译开始>
// Dump 打印默认资源对象的文件。 md5:fc090361befff87e
# <翻译结束>

