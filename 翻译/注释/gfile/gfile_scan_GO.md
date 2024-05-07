
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
// Max recursive depth for directory scanning.
<原文结束>

# <翻译开始>
// 目录扫描的最大递归深度。
# <翻译结束>


<原文开始>
// ScanDir returns all sub-files with absolute paths of given `path`,
// It scans directory recursively if given parameter `recursive` is true.
//
// The pattern parameter `pattern` supports multiple file name patterns,
// using the ',' symbol to separate multiple patterns.
<原文结束>

# <翻译开始>
// ScanDir 返回给定路径`path`下所有子文件的绝对路径，
// 若给定参数`recursive`为真，则会递归扫描目录。
//
// 参数`pattern`支持多种文件名模式，
// 可以使用逗号 ',' 作为分隔符来指定多个模式。
# <翻译结束>


<原文开始>
// ScanDirFunc returns all sub-files with absolute paths of given `path`,
// It scans directory recursively if given parameter `recursive` is true.
//
// The pattern parameter `pattern` supports multiple file name patterns, using the ','
// symbol to separate multiple patterns.
//
// The parameter `recursive` specifies whether scanning the `path` recursively, which
// means it scans its sub-files and appends the files path to result array if the sub-file
// is also a folder. It is false in default.
//
// The parameter `handler` specifies the callback function handling each sub-file path of
// the `path` and its sub-folders. It ignores the sub-file path if `handler` returns an empty
// string, or else it appends the sub-file path to result slice.
<原文结束>

# <翻译开始>
// ScanDirFunc 函数返回给定 `path` 下所有子文件的绝对路径，
// 如果给定参数 `recursive` 为 true，则会递归扫描目录。
//
// 参数 `pattern` 支持多个文件名模式，使用 ',' 符号分隔多个模式。
//
// 参数 `recursive` 指定是否递归扫描 `path`，这意味着如果子文件也是一个文件夹，它会扫描其下的子文件并将文件路径追加到结果切片中，默认情况下为 false。
//
// 参数 `handler` 指定了处理 `path` 及其子文件夹下每个子文件路径的回调函数。如果 `handler` 返回空字符串，则忽略该子文件路径，否则将其子文件路径追加到结果切片中。
# <翻译结束>


<原文开始>
// ScanDirFile returns all sub-files with absolute paths of given `path`,
// It scans directory recursively if given parameter `recursive` is true.
//
// The pattern parameter `pattern` supports multiple file name patterns,
// using the ',' symbol to separate multiple patterns.
//
// Note that it returns only files, exclusive of directories.
<原文结束>

# <翻译开始>
// ScanDirFile 返回给定 `path` 下所有子文件的绝对路径，
// 如果给定参数 `recursive` 为 true，则会递归扫描目录。
//
// 参数 `pattern` 支持多个文件名模式，
// 使用 `,` 符号来分隔多个模式。
//
// 注意，它只返回文件，不包括目录。
# <翻译结束>


<原文开始>
// ScanDirFileFunc returns all sub-files with absolute paths of given `path`,
// It scans directory recursively if given parameter `recursive` is true.
//
// The pattern parameter `pattern` supports multiple file name patterns, using the ','
// symbol to separate multiple patterns.
//
// The parameter `recursive` specifies whether scanning the `path` recursively, which
// means it scans its sub-files and appends the file paths to result array if the sub-file
// is also a folder. It is false in default.
//
// The parameter `handler` specifies the callback function handling each sub-file path of
// the `path` and its sub-folders. It ignores the sub-file path if `handler` returns an empty
// string, or else it appends the sub-file path to result slice.
//
// Note that the parameter `path` for `handler` is not a directory but a file.
// It returns only files, exclusive of directories.
<原文结束>

# <翻译开始>
// ScanDirFileFunc 函数返回给定路径 `path` 下所有子文件的绝对路径。
// 如果给定参数 `recursive` 为 true，则会递归扫描目录。
//
// 参数 `pattern` 支持多个文件名模式，使用 ',' 符号分隔多个模式。
//
// 参数 `recursive` 指定是否递归扫描 `path`，这意味着如果子文件也是一个文件夹，则会扫描其下级文件并将文件路径添加到结果切片中，默认情况下为 false。
//
// 参数 `handler` 指定了处理 `path` 及其子文件夹下每个子文件路径的回调函数。如果 `handler` 返回空字符串，则忽略该子文件路径，否则将子文件路径追加到结果切片中。
//
// 注意，传给 `handler` 的参数 `path` 不是一个目录而是一个文件。
// 此函数仅返回文件（不包括目录）。
# <翻译结束>


<原文开始>
// doScanDir is an internal method which scans directory and returns the absolute path
// list of files that are not sorted.
//
// The pattern parameter `pattern` supports multiple file name patterns, using the ','
// symbol to separate multiple patterns.
//
// The parameter `recursive` specifies whether scanning the `path` recursively, which
// means it scans its sub-files and appends the files path to result array if the sub-file
// is also a folder. It is false in default.
//
// The parameter `handler` specifies the callback function handling each sub-file path of
// the `path` and its sub-folders. It ignores the sub-file path if `handler` returns an empty
// string, or else it appends the sub-file path to result slice.
<原文结束>

# <翻译开始>
// doScanDir 是一个内部方法，用于扫描目录并返回未排序的绝对路径文件列表。
// 参数 `pattern` 支持多个文件名模式，使用 ',' 符号分隔多个模式。
// 参数 `recursive` 指定是否递归扫描 `path`，这意味着它会扫描其子文件，如果子文件也是一个文件夹，则将子文件路径追加到结果切片中。默认情况下，recursive 为 false。
// 参数 `handler` 指定了处理 `path` 及其子文件夹下每个子文件路径的回调函数。如果 `handler` 返回空字符串，则忽略该子文件路径，否则将其追加到结果切片中。
# <翻译结束>







<原文开始>
// If it meets pattern, then add it to the result list.
<原文结束>

# <翻译开始>
// 如果满足模式，则将其添加到结果列表中。
# <翻译结束>


<原文开始>
// Handler filtering.
<原文结束>

# <翻译开始>
// 处理器筛选功能
# <翻译结束>

