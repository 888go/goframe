
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
// Package gres provides resource management and packing/unpacking feature between files and bytes.
<原文结束>

# <翻译开始>
// Package gres 提供了资源管理功能以及文件和字节之间的打包/解包功能。
# <翻译结束>


<原文开始>
// Add unpacks and adds the `content` into the default resource object.
// The unnecessary parameter `prefix` indicates the prefix
// for each file storing into current resource object.
<原文结束>

# <翻译开始>
// Add 函数将解包并把 `content` 添加到默认资源对象中。
// 不必要的参数 `prefix` 表示存储到当前资源对象中每个文件的前缀。
# <翻译结束>


<原文开始>
// Load loads, unpacks and adds the data from `path` into the default resource object.
// The unnecessary parameter `prefix` indicates the prefix
// for each file storing into current resource object.
<原文结束>

# <翻译开始>
// Load 从`path`加载、解压并将数据添加到默认资源对象中。
// 不必要的参数`prefix`表示存储到当前资源对象时每个文件的前缀。
# <翻译结束>


<原文开始>
// Get returns the file with given path.
<原文结束>

# <翻译开始>
// Get 返回指定路径的文件。
# <翻译结束>


<原文开始>
// GetWithIndex searches file with `path`, if the file is directory
// it then does index files searching under this directory.
//
// GetWithIndex is usually used for http static file service.
<原文结束>

# <翻译开始>
// GetWithIndex 搜索指定 `path` 的文件，如果该文件是一个目录，
// 则进一步在该目录下进行索引文件的搜索。
//
// GetWithIndex 通常用于 HTTP 静态文件服务。
# <翻译结束>


<原文开始>
// GetContent directly returns the content of `path` in default resource object.
<原文结束>

# <翻译开始>
// GetContent 直接返回默认资源对象中 `path` 的内容。
# <翻译结束>


<原文开始>
// Contains checks whether the `path` exists in the default resource object.
<原文结束>

# <翻译开始>
// Contains 检查默认资源对象中是否存在 `path`。
# <翻译结束>


<原文开始>
// IsEmpty checks and returns whether the resource manager is empty.
<原文结束>

# <翻译开始>
// IsEmpty 检查并返回资源管理器是否为空。
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
// ScanDir 返回给定路径下的文件，参数 `path` 应为文件夹类型。
//
// 参数 `pattern` 支持多个文件名模式，
// 使用 ',' 符号分隔多个模式。
//
// 若给定的参数 `recursive` 为 true，则会递归扫描目录。
# <翻译结束>


<原文开始>
// ScanDirFile returns all sub-files with absolute paths of given `path`,
// It scans directory recursively if given parameter `recursive` is true.
//
// Note that it returns only files, exclusive of directories.
<原文结束>

# <翻译开始>
// ScanDirFile 返回给定 `path` 下所有子文件的绝对路径，
// 若给定参数 `recursive` 为 true，则会递归扫描目录。
//
// 注意，该函数仅返回文件，不包括目录。
# <翻译结束>


<原文开始>
// Export exports and saves specified path `src` and all its sub files to specified system path `dst` recursively.
<原文结束>

# <翻译开始>
// Export 函数会递归地导出并保存指定路径 `src` 及其所有子文件到指定的系统路径 `dst`。
# <翻译结束>


<原文开始>
// Dump prints the files of the default resource object.
<原文结束>

# <翻译开始>
// Dump 打印默认资源对象中的文件。
# <翻译结束>


<原文开始>
// Separator for directories.
<原文结束>

# <翻译开始>
// 目录分隔符。
# <翻译结束>


<原文开始>
// Default resource object.
<原文结束>

# <翻译开始>
// 默认资源对象。
# <翻译结束>

