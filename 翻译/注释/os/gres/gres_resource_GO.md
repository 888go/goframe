
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
// New creates and returns a new resource object.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的资源对象。. md5:8c594601bad2dd64
# <翻译结束>


<原文开始>
// Add unpacks and adds the `content` into current resource object.
// The unnecessary parameter `prefix` indicates the prefix
// for each file storing into current resource object.
<原文结束>

# <翻译开始>
// Add 解包并把`content`添加到当前资源对象中。不必要的参数`prefix`表示每个文件存储在当前资源对象中的前缀。
// md5:93345d9770c1e7fa
# <翻译结束>


<原文开始>
// Load loads, unpacks and adds the data from `path` into current resource object.
// The unnecessary parameter `prefix` indicates the prefix
// for each file storing into current resource object.
<原文结束>

# <翻译开始>
// Load 从`path`加载、解包并将数据添加到当前资源对象中。不必要的参数`prefix`表示将每个文件存储到当前资源对象中的前缀。
// md5:ab3e52fa479e7de6
# <翻译结束>


<原文开始>
// Get returns the file with given path.
<原文结束>

# <翻译开始>
// Get返回给定路径的文件。. md5:f4989a4832cde2d2
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
// Necessary for double char '/' replacement in prefix.
<原文结束>

# <翻译开始>
// 用于在前缀中替换双字符 '/'。. md5:2ab9f670789bab70
# <翻译结束>


<原文开始>
// GetContent directly returns the content of `path`.
<原文结束>

# <翻译开始>
// GetContent 直接返回 `path` 的内容。. md5:50cf0f721b7b89a5
# <翻译结束>


<原文开始>
// Contains checks whether the `path` exists in current resource object.
<原文结束>

# <翻译开始>
// Contains 检查路径 `path` 是否存在于当前资源对象中。. md5:9beb2e9c06e1e221
# <翻译结束>


<原文开始>
// IsEmpty checks and returns whether the resource manager is empty.
<原文结束>

# <翻译开始>
// IsEmpty 检查资源管理器是否为空，并返回结果。. md5:3aaae27781ad4e8c
# <翻译结束>


<原文开始>
// ScanDir returns the files under the given path, the parameter `path` should be a folder type.
//
// The pattern parameter `pattern` supports multiple file name patterns,
// using the ',' symbol to separate multiple patterns.
//
// It scans directory recursively if given parameter `recursive` is true.
//
// Note that the returned files does not contain given parameter `path`.
<原文结束>

# <翻译开始>
// ScanDir 在给定路径下返回文件，参数 `path` 应该是一个文件夹类型。
//
// `pattern` 参数支持多个文件名模式，使用逗号 `,` 来分隔多个模式。
//
// 如果 `recursive` 参数为 true，它会递归扫描目录。
//
// 注意，返回的文件不包含给定的 `path`。
// md5:c7e8c1023db3f55f
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
// doScanDir is an internal method which scans directory
// and returns the absolute path list of files that are not sorted.
//
// The pattern parameter `pattern` supports multiple file name patterns,
// using the ',' symbol to separate multiple patterns.
//
// It scans directory recursively if given parameter `recursive` is true.
<原文结束>

# <翻译开始>
// doScanDir 是一个内部方法，用于扫描目录
// 并返回未排序的文件的绝对路径列表。
//
// 模式参数 `pattern` 支持多个文件名模式，
// 使用 ',' 符号来分隔多个模式。
//
// 如果给定的参数 `recursive` 为 true，则会递归地扫描目录。
// md5:9e5185e985fd2bb6
# <翻译结束>


<原文开始>
// Used for type checking for first entry.
<原文结束>

# <翻译开始>
// 用于检查第一个条目的类型。. md5:da747d2102d6a47c
# <翻译结束>


<原文开始>
// To avoid of, eg: /i18n and /i18n-dir
<原文结束>

# <翻译开始>
// 为了避免，例如：/i18n 和 /i18n-dir. md5:ab3565cb1db7bc63
# <翻译结束>


<原文开始>
// ExportOption is the option for function Export.
<原文结束>

# <翻译开始>
// ExportOption 是 Export 函数的选项。. md5:12a5d99e83d743f7
# <翻译结束>


<原文开始>
// Remove the prefix of file name from resource.
<原文结束>

# <翻译开始>
// 从资源中移除文件名的前缀。. md5:ff1e0af55baecf64
# <翻译结束>


<原文开始>
// Export exports and saves specified path `srcPath` and all its sub files to specified system path `dstPath` recursively.
<原文结束>

# <翻译开始>
// Export 将指定的路径 `srcPath` 及其所有子文件递归导出并保存到指定的系统路径 `dstPath`。. md5:271f4d0f27211419
# <翻译结束>


<原文开始>
// Dump prints the files of current resource object.
<原文结束>

# <翻译开始>
// Dump 打印当前资源对象的文件。. md5:4533063269cc5df2
# <翻译结束>

