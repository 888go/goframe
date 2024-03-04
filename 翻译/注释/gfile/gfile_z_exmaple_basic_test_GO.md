
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
// Check whether the file exists
<原文结束>

# <翻译开始>
// 检查文件是否存在
# <翻译结束>


<原文开始>
// Creates file with given `path` recursively
<原文结束>

# <翻译开始>
// 根据给定的`path`递归创建文件
# <翻译结束>







<原文开始>
// Reads len(b) bytes from the File
<原文结束>

# <翻译开始>
// 从File中读取len(b)个字节
# <翻译结束>


<原文开始>
// Open file or directory with READONLY model
<原文结束>

# <翻译开始>
// 以只读模式打开文件或目录
# <翻译结束>


<原文开始>
	// Opens file/directory with custom `flag` and `perm`
	// Create if file does not exist,it is created in a readable and writable mode,prem 0777
<原文结束>

# <翻译开始>
// 通过自定义的`flag`和`perm`打开文件/目录
// 如果文件不存在，则创建，创建时默认为可读写模式，权限为0777
# <翻译结束>


<原文开始>
	// Opens file/directory with custom `flag`
	// Create if file does not exist,it is created in a readable and writable mode with default `perm` is 0666
<原文结束>

# <翻译开始>
// 使用自定义`flag`打开文件/目录
// 如果文件不存在，则创建，以默认权限`perm`为0666的可读写模式创建
# <翻译结束>


<原文开始>
// Joins string array paths with file separator of current system.
<原文结束>

# <翻译开始>
// 使用当前系统的文件分隔符连接字符串数组路径。
# <翻译结束>


<原文开始>
// Checks whether given `path` exist.
<原文结束>

# <翻译开始>
// 检查给定的 `path` 是否存在。
# <翻译结束>


<原文开始>
// Checks whether given `path` a directory.
<原文结束>

# <翻译开始>
// 检查给定的`path`是否为一个目录。
# <翻译结束>


<原文开始>
// Get absolute path of current working directory.
<原文结束>

# <翻译开始>
// 获取当前工作目录的绝对路径。
# <翻译结束>


<原文开始>
// Get current working directory
<原文结束>

# <翻译开始>
// 获取当前工作目录
# <翻译结束>


<原文开始>
// Changes the current working directory to the named directory.
<原文结束>

# <翻译开始>
// 将当前工作目录更改为指定的目录。
# <翻译结束>


<原文开始>
// Checks whether given `path` a file, which means it's not a directory.
<原文结束>

# <翻译开始>
// 检查给定的`path`是否为文件，也就是说它不是一个目录。
# <翻译结束>


<原文开始>
// Get a FileInfo describing the named file.
<原文结束>

# <翻译开始>
// 获取描述指定文件的FileInfo对象。
# <翻译结束>


<原文开始>
	//  Moves `src` to `dst` path.
	// If `dst` already exists and is not a directory, it'll be replaced.
<原文结束>

# <翻译开始>
// 将`src`移动到`dst`路径。
// 如果`dst`已存在且不是一个目录，它将会被替换。
# <翻译结束>


<原文开始>
	//  renames (moves) `src` to `dst` path.
	// If `dst` already exists and is not a directory, it'll be replaced.
<原文结束>

# <翻译开始>
// 将`src`重命名为（移动到）`dst`路径。
// 如果`dst`已存在且不是一个目录，则会被替换。
# <翻译结束>


<原文开始>
// Get sub-file names of given directory `path`.
<原文结束>

# <翻译开始>
// 获取给定目录 `path` 下的子文件名列表。
# <翻译结束>


<原文开始>
// Get sub-file names of given directory `path`.
	// Only show file name
<原文结束>

# <翻译开始>
// 获取给定目录 `path` 下的子文件名。
// 仅显示文件名
# <翻译结束>







<原文开始>
// Checks whether given `path` is readable.
<原文结束>

# <翻译开始>
// 检查给定的`path`是否可读。
# <翻译结束>


<原文开始>
// Checks whether given `path` is writable.
<原文结束>

# <翻译开始>
// 检查给定的`path`是否可写。
# <翻译结束>

















<原文开始>
// Get an absolute representation of path.
<原文结束>

# <翻译开始>
// 获取path的绝对表示形式。
# <翻译结束>


<原文开始>
// fetch an absolute representation of path.
<原文结束>

# <翻译开始>
// 获取path的绝对表示形式。
# <翻译结束>


<原文开始>
// Get absolute file path of current running process
<原文结束>

# <翻译开始>
// 获取当前运行进程的绝对文件路径
# <翻译结束>


<原文开始>
// Get file name of current running process
<原文结束>

# <翻译开始>
// 获取当前运行进程的文件名
# <翻译结束>


<原文开始>
// Get absolute directory path of current running process
<原文结束>

# <翻译开始>
// 获取当前运行进程的绝对目录路径
# <翻译结束>


<原文开始>
// Get the last element of path, which contains file extension.
<原文结束>

# <翻译开始>
// 获取路径中的最后一个元素，该元素包含文件扩展名。
# <翻译结束>


<原文开始>
// Get the last element of path without file extension.
<原文结束>

# <翻译开始>
// 获取路径中最后一个元素，不包括文件扩展名。
# <翻译结束>


<原文开始>
// Get all but the last element of path, typically the path's directory.
<原文结束>

# <翻译开始>
// 获取路径中除最后一个元素之外的所有元素，通常是指路径的目录部分。
# <翻译结束>


<原文开始>
// Check whether the `path` is empty
<原文结束>

# <翻译开始>
// 检查 `path` 是否为空
# <翻译结束>


<原文开始>
// Get the file name extension used by path.
<原文结束>

# <翻译开始>
// 获取路径中使用的文件名扩展名。
# <翻译结束>


<原文开始>
// Get the file name extension used by path but the result does not contains symbol '.'.
<原文结束>

# <翻译开始>
// 获取路径path所使用的文件名扩展名，但结果中不包含符号'.'。
# <翻译结束>


<原文开始>
// deletes all file/directory with `path` parameter.
<原文结束>

# <翻译开始>
// 删除具有`path`参数的所有文件/目录。
# <翻译结束>
 
<原文开始>
// Check if directory exists
<原文结束>

# <翻译开始>
// 检查目录是否存在
# <翻译结束>


<原文开始>
// Write some content to file
<原文结束>

# <翻译开始>
// 向文件写入一些内容
# <翻译结束>


<原文开始>
// Show full path of the file
<原文结束>

# <翻译开始>
// 显示文件的完整路径
# <翻译结束>


<原文开始>
// Show original mode
<原文结束>

# <翻译开始>
// 显示原始模式
# <翻译结束>


<原文开始>
// Show the modified mode
<原文结束>

# <翻译开始>
// 显示修改后的模式
# <翻译结束>

