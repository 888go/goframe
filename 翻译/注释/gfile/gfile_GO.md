
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
// Package gfile provides easy-to-use operations for file system.
<原文结束>

# <翻译开始>
// Package gfile 提供了对文件系统进行便捷操作的功能。
# <翻译结束>


<原文开始>
	// Separator for file system.
	// It here defines the separator as variable
	// to allow it modified by developer if necessary.
<原文结束>

# <翻译开始>
// 文件系统的分隔符
// 这里将分隔符定义为变量
// 以便在必要时允许开发者进行修改
# <翻译结束>


<原文开始>
// DefaultPermOpen is the default perm for file opening.
<原文结束>

# <翻译开始>
// DefaultPermOpen 是文件打开时的默认权限。
# <翻译结束>


<原文开始>
// DefaultPermCopy is the default perm for file/folder copy.
<原文结束>

# <翻译开始>
// DefaultPermCopy 是文件/文件夹复制时的默认权限。
# <翻译结束>


<原文开始>
	// The absolute file path for main package.
	// It can be only checked and set once.
<原文结束>

# <翻译开始>
// 主包的绝对文件路径。
// 它只能被检查和设置一次。
# <翻译结束>


<原文开始>
	// selfPath is the current running binary path.
	// As it is most commonly used, it is so defined as an internal package variable.
<原文结束>

# <翻译开始>
// selfPath 是当前运行的二进制文件路径。
// 由于它被广泛使用，因此将其定义为内部包变量。
# <翻译结束>


<原文开始>
// Initialize internal package variable: selfPath.
<原文结束>

# <翻译开始>
// 初始化内部包变量：selfPath
# <翻译结束>


<原文开始>
// Mkdir creates directories recursively with given `path`.
// The parameter `path` is suggested to be an absolute path instead of relative one.
<原文结束>

# <翻译开始>
// Mkdir 通过给定的 `path` 参数递归创建目录。
// 建议 `path` 参数使用绝对路径而非相对路径。
# <翻译结束>


<原文开始>
// Create creates a file with given `path` recursively.
// The parameter `path` is suggested to be absolute path.
<原文结束>

# <翻译开始>
// Create 以给定的 `path` 创建文件并递归创建其所在目录。
// 建议参数 `path` 使用绝对路径。
# <翻译结束>


<原文开始>
// Open opens file/directory READONLY.
<原文结束>

# <翻译开始>
// Open 以只读方式打开文件/目录。
# <翻译结束>


<原文开始>
// OpenFile opens file/directory with custom `flag` and `perm`.
// The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.
<原文结束>

# <翻译开始>
// OpenFile函数以自定义的`flag`和`perm`打开文件/目录。
// 参数`flag`类似于：O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC等。
# <翻译结束>


<原文开始>
// OpenWithFlag opens file/directory with default perm and custom `flag`.
// The default `perm` is 0666.
// The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.
<原文结束>

# <翻译开始>
// OpenWithFlag 函数以默认权限和自定义标志 `flag` 打开文件/目录。
// 默认的 `perm` 为 0666。
// 参数 `flag` 如：O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC 等。
# <翻译结束>


<原文开始>
// OpenWithFlagPerm opens file/directory with custom `flag` and `perm`.
// The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.
// The parameter `perm` is like: 0600, 0666, 0777, etc.
<原文结束>

# <翻译开始>
// OpenWithFlagPerm 使用自定义`flag`和`perm`打开文件/目录。
// 参数`flag`类似于：O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC等。
// 参数`perm`类似于：0600, 0666, 0777等。
# <翻译结束>


<原文开始>
// Join joins string array paths with file separator of current system.
<原文结束>

# <翻译开始>
// Join 使用当前系统的文件分隔符连接字符串数组paths。
# <翻译结束>


<原文开始>
// Exists checks whether given `path` exist.
<原文结束>

# <翻译开始>
// Exists 检查给定的 `path` 是否存在。
# <翻译结束>


<原文开始>
// IsDir checks whether given `path` a directory.
// Note that it returns false if the `path` does not exist.
<原文结束>

# <翻译开始>
// IsDir 检查给定的 `path` 是否为一个目录。
// 注意，如果 `path` 不存在，则返回 false。
# <翻译结束>


<原文开始>
// Pwd returns absolute path of current working directory.
// Note that it returns an empty string if retrieving current
// working directory failed.
<原文结束>

# <翻译开始>
// Pwd 返回当前工作目录的绝对路径。
// 注意，如果获取当前工作目录失败，它将返回一个空字符串。
# <翻译结束>


<原文开始>
// Chdir changes the current working directory to the named directory.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Chdir函数将当前工作目录更改为指定的目录名称。
// 如果出现错误，其类型将会是*PathError。
# <翻译结束>


<原文开始>
// IsFile checks whether given `path` a file, which means it's not a directory.
// Note that it returns false if the `path` does not exist.
<原文结束>

# <翻译开始>
// IsFile 检查给定的 `path` 是否为文件，也就是说它不是一个目录。
// 注意，如果 `path` 不存在，则返回 false。
# <翻译结束>


<原文开始>
// Stat returns a FileInfo describing the named file.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Stat返回一个FileInfo，用于描述指定名称的文件。
// 如果出现错误，其类型将会是*PathError。
# <翻译结束>


<原文开始>
// Move renames (moves) `src` to `dst` path.
// If `dst` already exists and is not a directory, it'll be replaced.
<原文结束>

# <翻译开始>
// Move 将`src`重命名为（移动到）`dst`路径。
// 如果`dst`已存在且不是一个目录，它将会被替换。
# <翻译结束>


<原文开始>
// Rename is alias of Move.
// See Move.
<原文结束>

# <翻译开始>
// Rename 是 Move 的别名。
// 请参阅 Move。
# <翻译结束>


<原文开始>
// DirNames returns sub-file names of given directory `path`.
// Note that the returned names are NOT absolute paths.
<原文结束>

# <翻译开始>
// DirNames 返回给定目录 `path` 下的子文件名。
// 注意，返回的名称不是绝对路径。
# <翻译结束>


<原文开始>
// Glob returns the names of all files matching pattern or nil
// if there is no matching file. The syntax of patterns is the same
// as in Match. The pattern may describe hierarchical names such as
// /usr/*/bin/ed (assuming the Separator is '/').
//
// Glob ignores file system errors such as I/O errors reading directories.
// The only possible returned error is ErrBadPattern, when pattern
// is malformed.
<原文结束>

# <翻译开始>
// Glob返回所有与pattern匹配的文件名，如果没有匹配的文件，则返回nil。模式语法与Match函数中相同。
// 模式可以描述层级式的文件名，例如（假设分隔符为'/'）/usr/*/bin/ed。
//
// Glob会忽略文件系统错误，如读取目录时的I/O错误。唯一可能返回的错误是ErrBadPattern，仅当模式格式不正确时发生。
# <翻译结束>


<原文开始>
// Remove deletes all file/directory with `path` parameter.
// If parameter `path` is directory, it deletes it recursively.
//
// It does nothing if given `path` does not exist or is empty.
<原文结束>

# <翻译开始>
// Remove 函数用于删除参数 `path` 指定的文件或目录。
// 若参数 `path` 为目录，该函数会递归地删除整个目录及其包含的所有文件和子目录。
//
// 如果给定的 `path` 不存在或者为空，则该函数不做任何操作。
# <翻译结束>


<原文开始>
// It does nothing if `path` is empty.
<原文结束>

# <翻译开始>
// 如果`path`为空，则此操作无任何效果。
# <翻译结束>


<原文开始>
// IsReadable checks whether given `path` is readable.
<原文结束>

# <翻译开始>
// IsReadable 检查给定的 `path` 是否可读。
# <翻译结束>


<原文开始>
// IsWritable checks whether given `path` is writable.
//
// TODO improve performance; use golang.org/x/sys to cross-plat-form
<原文结束>

# <翻译开始>
// IsWritable 检查给定的 `path` 是否可写。
//
// TODO 提高性能；使用 golang.org/x/sys 以实现跨平台
# <翻译结束>


<原文开始>
// If it's a directory, create a temporary file to test whether it's writable.
<原文结束>

# <翻译开始>
// 如果是目录，则创建一个临时文件以测试其是否可写。
# <翻译结束>


<原文开始>
// If it's a file, check if it can open it.
<原文结束>

# <翻译开始>
// 如果它是一个文件，检查是否可以打开它。
# <翻译结束>


<原文开始>
// Chmod is alias of os.Chmod.
// See os.Chmod.
<原文结束>

# <翻译开始>
// Chmod 是 os.Chmod 的别名。
// 请参阅 os.Chmod。
# <翻译结束>


<原文开始>
// Abs returns an absolute representation of path.
// If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path. The absolute
// path name for a given file is not guaranteed to be unique.
// Abs calls Clean on the result.
<原文结束>

# <翻译开始>
// Abs 返回 path 的绝对路径表示。
// 如果 path 不是绝对路径，则将其与当前工作目录连接，
// 以转换为绝对路径。对于给定的文件，其绝对路径名称不保证唯一性。
// Abs 会对结果调用 Clean 函数。
// 这段代码是关于 Go 语言（Golang）中的 `path/filepath` 包中 Abs 函数的注释翻译：
// ```go
// Abs 函数返回路径 path 的绝对路径形式。
// 如果路径 path 不是绝对路径，该函数会将其与当前工作目录拼接，
// 从而生成一个绝对路径。需要注意的是，即使对于同一个文件，其计算出的绝对路径也未必是唯一的。
// Abs 函数还会对处理后的结果调用 Clean 函数进行规范化。
# <翻译结束>


<原文开始>
// RealPath converts the given `path` to its absolute path
// and checks if the file path exists.
// If the file does not exist, return an empty string.
<原文结束>

# <翻译开始>
// RealPath 将给定的 `path` 转换为绝对路径
// 并检查文件路径是否存在。
// 如果文件不存在，则返回一个空字符串。
# <翻译结束>


<原文开始>
// SelfPath returns absolute file path of current running process(binary).
<原文结束>

# <翻译开始>
// SelfPath 返回当前运行进程（二进制文件）的绝对文件路径。
# <翻译结束>


<原文开始>
// SelfName returns file name of current running process(binary).
<原文结束>

# <翻译开始>
// SelfName 返回当前运行进程（二进制文件）的文件名。
# <翻译结束>


<原文开始>
// SelfDir returns absolute directory path of current running process(binary).
<原文结束>

# <翻译开始>
// SelfDir 返回当前运行进程（二进制文件）的绝对目录路径。
# <翻译结束>


<原文开始>
// Basename returns the last element of path, which contains file extension.
// Trailing path separators are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of separators, Basename returns a single separator.
// Example:
// /var/www/file.js -> file.js
// file.js          -> file.js
<原文结束>

# <翻译开始>
// Basename 返回路径的最后一个元素，其中包含文件扩展名。
// 在提取最后一个元素之前会移除尾部的路径分隔符。
// 如果路径为空，Base 返回 "."。
// 如果路径完全由分隔符组成，Basename 将返回一个单个分隔符。
// 示例：
// /var/www/file.js -> file.js
// file.js          -> file.js
# <翻译结束>


<原文开始>
// Name returns the last element of path without file extension.
// Example:
// /var/www/file.js -> file
// file.js          -> file
<原文结束>

# <翻译开始>
// Name函数返回路径中最后一个元素的文件名部分，不包括文件扩展名。
// 示例：
// /var/www/file.js -> file
// file.js          -> file
# <翻译结束>


<原文开始>
// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element, Dir calls Clean on the path and trailing
// slashes are removed.
// If the `path` is empty, Dir returns ".".
// If the `path` is ".", Dir treats the path as current working directory.
// If the `path` consists entirely of separators, Dir returns a single separator.
// The returned path does not end in a separator unless it is the root directory.
<原文结束>

# <翻译开始>
// Dir 返回路径path去掉最后一个元素后的部分，通常为路径的目录部分。
// 在去掉最后一个元素之后，Dir会对路径进行Clean操作，并移除尾部的斜杠。
// 如果 `path` 为空，Dir 返回"."。
// 如果 `path` 为".", Dir 将路径视为当前工作目录。
// 如果 `path` 完全由分隔符组成，Dir 返回一个单独的分隔符。
// 返回的路径除非是根目录，否则不会以分隔符结尾。
# <翻译结束>


<原文开始>
// IsEmpty checks whether the given `path` is empty.
// If `path` is a folder, it checks if there's any file under it.
// If `path` is a file, it checks if the file size is zero.
//
// Note that it returns true if `path` does not exist.
<原文结束>

# <翻译开始>
// IsEmpty 检查给定的 `path` 是否为空。
// 如果 `path` 是一个文件夹，它会检查该文件夹下是否存在任何文件。
// 如果 `path` 是一个文件，它会检查该文件的大小是否为零。
//
// 注意，如果 `path` 不存在，此函数也将返回 true。
# <翻译结束>


<原文开始>
// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is
// no dot.
// Note: the result contains symbol '.'.
// Eg:
// main.go  => .go
// api.json => .json
<原文结束>

# <翻译开始>
// Ext 返回路径 path 使用的文件名扩展名。
// 扩展名是从路径中最后一个元素的最后一个点开始的后缀；
// 如果没有点，则为空。
// 注意：结果中包含符号 '.'。
// 示例：
// main.go  => .go
// api.json => .json
# <翻译结束>


<原文开始>
// ExtName is like function Ext, which returns the file name extension used by path,
// but the result does not contain symbol '.'.
// Eg:
// main.go  => go
// api.json => json
<原文结束>

# <翻译开始>
// ExtName 类似于函数 Ext，它返回路径中使用的文件名扩展名，
// 但是结果中不包含符号'.'。
// 示例：
// main.go  => go
// api.json => json
# <翻译结束>


<原文开始>
// Temp retrieves and returns the temporary directory of current system.
//
// The optional parameter `names` specifies the sub-folders/sub-files,
// which will be joined with current system separator and returned with the path.
<原文结束>

# <翻译开始>
// Temp 函数获取并返回当前系统的临时目录路径。
//
// 可选参数 `names` 指定的是子文件夹或子文件名，
// 这些名称会与当前系统的路径分隔符拼接，并将最终生成的完整路径返回。
# <翻译结束>

