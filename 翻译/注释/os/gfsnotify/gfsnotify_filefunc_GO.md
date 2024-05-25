
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// ThIs Source Code Form Is subject to the terms of the MIT License.
// If a copy of the MIT was not dIstributed with thIs file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权声明：GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码受MIT许可证条款约束。如果此文件未附带MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:12b80d680e9de440
# <翻译结束>


<原文开始>
// fileDir returns all but the last element of path, typically the path's directory.
// After dropping the final element, Dir calls Clean on the path and trailing
// slashes are removed.
// If the path is empty, Dir returns ".".
// If the path consists entirely of separators, Dir returns a single separator.
// The returned path does not end in a separator unless it is the root directory.
<原文结束>

# <翻译开始>
// fileDir 返回路径除最后一个元素之外的所有内容，通常为路径的目录。
// 在丢弃最后一个元素后，Dir 函数会对路径调用 Clean 方法，并移除尾随斜杠。
// 如果路径为空，Dir 返回"."。
// 如果路径仅由分隔符组成，Dir 返回单个分隔符。
// 返回的路径除非是根目录，否则不会以分隔符结尾。
// md5:c4932e9f21542326
# <翻译结束>


<原文开始>
// fileRealPath converts the given `path` to its absolute path
// and checks if the file path exists.
// If the file does not exist, return an empty string.
<原文结束>

# <翻译开始>
// fileRealPath 将给定的 `path` 转换为其绝对路径，并检查文件路径是否存在。
// 如果文件不存在，返回空字符串。
// md5:e30bc7542e1c332e
# <翻译结束>


<原文开始>
// fileExists checks whether given `path` exist.
<原文结束>

# <翻译开始>
// fileExists 检查给定的 `path` 是否存在。 md5:88ff2b38709b04ab
# <翻译结束>


<原文开始>
// fileIsDir checks whether given `path` a directory.
<原文结束>

# <翻译开始>
// fileIsDir 检查给定的 `path` 是否为目录。 md5:e16eca80a5a42f13
# <翻译结束>


<原文开始>
// fileAllDirs returns all sub-folders including itself of given `path` recursively.
<原文结束>

# <翻译开始>
// fileAllDirs递归地返回给定`path`的所有子文件夹，包括自身。 md5:b6638d72b44fee31
# <翻译结束>


<原文开始>
// fileScanDir returns all sub-files with absolute paths of given `path`,
// It scans directory recursively if given parameter `recursive` is true.
<原文结束>

# <翻译开始>
// fileScanDir 返回给定 `path` 所有子文件的绝对路径，
// 如果参数 `recursive` 为 true，则递归扫描目录。
// md5:871aa4d8d78c9ee4
# <翻译结束>


<原文开始>
// doFileScanDir is an internal method which scans directory
// and returns the absolute path list of files that are not sorted.
//
// The pattern parameter `pattern` supports multiple file name patterns,
// using the ',' symbol to separate multiple patterns.
//
// It scans directory recursively if given parameter `recursive` is true.
<原文结束>

# <翻译开始>
// doFileScanDir 是一个内部方法，用于扫描目录
// 并返回未排序的文件的绝对路径列表。
//
// 模式参数 `pattern` 支持多个文件名模式，
// 使用 `,` 符号分隔多个模式。
//
// 如果给定的参数 `recursive` 为 true，它将递归地扫描目录。
// md5:187616f6d86800cf
# <翻译结束>

