
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
// New creates and returns a new resource object.
<原文结束>

# <翻译开始>
// New 创建并返回一个新的资源对象。
# <翻译结束>


<原文开始>
// Add unpacks and adds the `content` into current resource object.
// The unnecessary parameter `prefix` indicates the prefix
// for each file storing into current resource object.
<原文结束>

# <翻译开始>
// Add 方法对`content`进行解包并将其添加到当前资源对象中。
// 不必要的参数`prefix`表示存储到当前资源对象时，每个文件的前缀。
# <翻译结束>


<原文开始>
// Load loads, unpacks and adds the data from `path` into current resource object.
// The unnecessary parameter `prefix` indicates the prefix
// for each file storing into current resource object.
<原文结束>

# <翻译开始>
// Load 从`path`加载、解压并将数据添加到当前资源对象中。
// 不必要的参数`prefix`表示存储到当前资源对象中的每个文件的前缀。
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
// Necessary for double char '/' replacement in prefix.
<原文结束>

# <翻译开始>
// 用于在前缀中替换双字符 '/'
# <翻译结束>


<原文开始>
// GetContent directly returns the content of `path`.
<原文结束>

# <翻译开始>
// GetContent直接返回`path`的内容。
# <翻译结束>


<原文开始>
// Contains checks whether the `path` exists in current resource object.
<原文结束>

# <翻译开始>
// Contains 检查当前资源对象中是否存在 `path`。
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
//
// Note that the returned files does not contain given parameter `path`.
<原文结束>

# <翻译开始>
// ScanDir 返回给定路径下的文件，参数`path`应为文件夹类型。
//
// 参数`pattern`支持多个文件名模式，
// 使用','符号来分隔多个模式。
//
// 如果给定参数`recursive`为真，则会递归扫描目录。
//
// 注意，返回的文件列表中不包含给定的参数`path`所代表的目录自身。
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
// doScanDir is an internal method which scans directory
// and returns the absolute path list of files that are not sorted.
//
// The pattern parameter `pattern` supports multiple file name patterns,
// using the ',' symbol to separate multiple patterns.
//
// It scans directory recursively if given parameter `recursive` is true.
<原文结束>

# <翻译开始>
// doScanDir 是一个内部方法，用于扫描目录并返回一个未排序的文件绝对路径列表。
//
// 参数`pattern`支持多个文件名模式，使用','符号分隔多个模式。
//
// 如果给定参数`recursive`为真，则会递归地扫描目录。
# <翻译结束>


<原文开始>
// Used for type checking for first entry.
<原文结束>

# <翻译开始>
// 用于对第一个条目的类型检查。
# <翻译结束>


<原文开始>
// To avoid of, eg: /i18n and /i18n-dir
<原文结束>

# <翻译开始>
// 为避免出现诸如/i18n和/i18n-dir这样的情况
# <翻译结束>


<原文开始>
// ExportOption is the option for function Export.
<原文结束>

# <翻译开始>
// ExportOption 是函数 Export 的选项。
# <翻译结束>


<原文开始>
// Remove the prefix of file name from resource.
<原文结束>

# <翻译开始>
// 从资源中移除文件名前缀
# <翻译结束>


<原文开始>
// Export exports and saves specified path `srcPath` and all its sub files to specified system path `dstPath` recursively.
<原文结束>

# <翻译开始>
// Export 递归地导出并保存指定路径`srcPath`及其所有子文件到指定系统路径`dstPath`。
# <翻译结束>


<原文开始>
// Dump prints the files of current resource object.
<原文结束>

# <翻译开始>
// Dump 打印当前资源对象的文件。
# <翻译结束>

