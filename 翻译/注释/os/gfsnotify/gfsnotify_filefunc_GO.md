
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// ThIs Source Code Form Is subject to the terms of the MIT License.
// If a copy of the MIT was not dIstributed with thIs file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT协议条款。如果随此文件未分发MIT协议副本，
// 您可以在https://github.com/gogf/gf获取一份。
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
// fileDir 返回路径中除最后一个元素之外的所有元素，通常是路径的目录部分。
// 在去掉最后一个元素后，Dir 会调用 Clean 来清理路径，并移除尾部的斜杠。
// 如果路径为空，Dir 返回 "."
// 如果路径完全由分隔符组成，Dir 返回单个分隔符。
// 返回的路径除非是根目录，否则不会以分隔符结尾。
// 此函数（fileDir）用于提取并返回给定路径的基本目录部分，在处理过程中会对路径进行规范化处理。具体规则如下：
// 1. 删除路径的最后一部分（通常为文件名或子目录名）。
// 2. 调用标准库中的 `Clean` 函数清理路径，去除末尾多余的斜杠。
// 3. 若路径为空，则返回当前目录（`.`）表示。
// 4. 若路径仅包含分隔符，则返回一个分隔符。
// 5. 返回的目录路径不以分隔符结尾，除非该路径指向的是根目录。
# <翻译结束>


<原文开始>
// fileRealPath converts the given `path` to its absolute path
// and checks if the file path exists.
// If the file does not exist, return an empty string.
<原文结束>

# <翻译开始>
// fileRealPath 将给定的 `path` 转换为绝对路径
// 并检查文件路径是否存在。
// 若文件不存在，则返回一个空字符串。
# <翻译结束>


<原文开始>
// fileExists checks whether given `path` exist.
<原文结束>

# <翻译开始>
// fileExists 检查给定的 `path` 是否存在。
# <翻译结束>


<原文开始>
// fileIsDir checks whether given `path` a directory.
<原文结束>

# <翻译开始>
// fileIsDir 检查给定的 `path` 是否为一个目录。
# <翻译结束>


<原文开始>
// fileAllDirs returns all sub-folders including itself of given `path` recursively.
<原文结束>

# <翻译开始>
// fileAllDirs 返回给定 `path`（包括其自身）的所有子目录，递归遍历。
# <翻译结束>


<原文开始>
// fileScanDir returns all sub-files with absolute paths of given `path`,
// It scans directory recursively if given parameter `recursive` is true.
<原文结束>

# <翻译开始>
// fileScanDir 函数返回给定 `path` 下所有子文件的绝对路径，
// 如果给定参数 `recursive` 为 true，则会递归地扫描目录。
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
// doFileScanDir 是一个内部方法，用于扫描目录并返回未排序的文件绝对路径列表。
//
// 参数`pattern`支持多个文件名模式，使用逗号 ',' 作为分隔符来指定多个模式。
//
// 如果给定的参数 `recursive` 为 true，则会递归地扫描目录。
# <翻译结束>

