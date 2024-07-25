
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Max recursive depth for directory scanning.
<原文结束>

# <翻译开始>
	// 扫描目录时的最大递归深度。 md5:6443b3b221d62366
# <翻译结束>


<原文开始>
// ScanDir returns all sub-files with absolute paths of given `path`,
// It scans directory recursively if given parameter `recursive` is true.
//
// The pattern parameter `pattern` supports multiple file name patterns,
// using the ',' symbol to separate multiple patterns.
<原文结束>

# <翻译开始>
// ScanDir 返回给定`path`下的所有子文件的绝对路径，
// 如果给定的参数`recursive`为true，则递归扫描目录。
//
// 模式参数`pattern`支持多个文件名模式，
// 使用`,`符号分隔多个模式。 md5:1f662f1008f0113e
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
// ScanDirFunc 返回给定`path`下的所有子文件的绝对路径。
// 如果参数`recursive`为真，它将递归扫描目录。
//
// 参数`pattern`支持多个文件名模式，使用逗号分隔多个模式。
//
// 参数`recursive`指定是否递归扫描`path`。默认情况下，它是false，表示不递归。
//
// 参数`handler`指定了处理`path`及其子目录下每个子文件路径的回调函数。如果`handler`返回空字符串，将忽略子文件路径，否则将子文件路径添加到结果切片中。 md5:93774b4b752cee08
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
// ScanDirFile 返回给定 `path` 所有子文件的绝对路径，
// 如果 `recursive` 参数为真，它会递归扫描目录。
//
// `pattern` 参数支持多个文件名模式，使用逗号 `,` 来分隔多个模式。
//
// 注意，它只返回文件，不包括目录。 md5:1d9c6ada055eaa05
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
// ScanDirFileFunc 返回给定 `path` 的所有子文件的绝对路径，
// 如果参数 `recursive` 为 true，则会递归扫描目录。
//
// `pattern` 参数支持多个文件名模式，使用逗号（',') 分隔多个模式。
//
// 参数 `recursive` 指定是否递归扫描 `path`，即如果子文件也是一个文件夹，它将把子文件路径添加到结果数组中。默认情况下为 false。
//
// 参数 `handler` 指定处理 `path` 和其子文件夹每个子文件路径的回调函数。如果 `handler` 返回空字符串，那么忽略该子文件路径；否则，将子文件路径添加到结果切片中。
//
// 注意，`handler` 中的参数 `path` 不是目录，而是文件。它只返回文件，不包括目录。 md5:036965ff87c95b63
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
// doScanDir 是一个内部方法，用于扫描目录并返回未排序的文件绝对路径列表。
//
// 模式参数 `pattern` 支持多个文件名模式，使用 ',' 符号分隔多个模式。
//
// 参数 `recursive` 指定是否递归扫描 `path`，即如果子文件也是一个文件夹，
// 则扫描其子文件并将文件路径追加到结果数组中。默认为 false。
//
// 参数 `handler` 指定一个回调函数，用于处理 `path` 及其子目录下的每个子文件路径。
// 如果 `handler` 返回空字符串，则忽略该子文件路径；否则，将子文件路径追加到结果切片中。 md5:5f6bc88fb2ff75fe
# <翻译结束>


<原文开始>
// If it meets pattern, then add it to the result list.
<原文结束>

# <翻译开始>
		// 如果满足模式，将其添加到结果列表中。 md5:11ed1569cf70af04
# <翻译结束>

