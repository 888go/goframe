
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
// Package gfile provides easy-to-use operations for file system.
<原文结束>

# <翻译开始>
// gfile 包提供了易于使用的文件系统操作。 md5:51d18e994a768bb4
# <翻译结束>


<原文开始>
	// Separator for file system.
	// It here defines the separator as variable
	// to allow it modified by developer if necessary.
<原文结束>

# <翻译开始>
	// 文件系统分隔符。
	// 这里将分隔符定义为变量，
	// 以便开发人员在必要时进行修改。
	// md5:ec0b6e47ec28478f
# <翻译结束>


<原文开始>
// DefaultPermOpen is the default perm for file opening.
<原文结束>

# <翻译开始>
// DefaultPermOpen 是文件打开的默认权限。 md5:dc57341030d46a11
# <翻译结束>


<原文开始>
// DefaultPermCopy is the default perm for file/folder copy.
<原文结束>

# <翻译开始>
// DefaultPermCopy 是用于文件/文件夹复制的默认权限。 md5:fdef6c133c375aa4
# <翻译结束>


<原文开始>
	// The absolute file path for main package.
	// It can be only checked and set once.
<原文结束>

# <翻译开始>
	// 主包的绝对文件路径。
	// 只能检查和设置一次。
	// md5:4a0d292a2835bc99
# <翻译结束>


<原文开始>
	// selfPath is the current running binary path.
	// As it is most commonly used, it is so defined as an internal package variable.
<原文结束>

# <翻译开始>
	// selfPath 是当前运行二进制文件的路径。
	// 由于它最常被使用，因此作为内部包变量进行定义。
	// md5:0e75acfdf2f4b9e7
# <翻译结束>


<原文开始>
// Initialize internal package variable: selfPath.
<原文结束>

# <翻译开始>
// 初始化内部包变量：selfPath。 md5:8d1168ac7361bb54
# <翻译结束>


<原文开始>
// Mkdir creates directories recursively with given `path`.
// The parameter `path` is suggested to be an absolute path instead of relative one.
<原文结束>

# <翻译开始>
// Mkdir 递归创建给定的 `path` 所表示的目录。建议使用绝对路径而非相对路径作为参数。
// md5:e78abb40a45c2886
# <翻译结束>


<原文开始>
// Create creates a file with given `path` recursively.
// The parameter `path` is suggested to be absolute path.
<原文结束>

# <翻译开始>
// Create 递归地创建具有给定`path`的文件。
// 建议参数`path`使用绝对路径。
// md5:163accaf36969b42
# <翻译结束>


<原文开始>
// Open opens file/directory READONLY.
<原文结束>

# <翻译开始>
// Open以只读方式打开文件/目录。 md5:7f50cf0f63b9e34e
# <翻译结束>


<原文开始>
// OpenFile opens file/directory with custom `flag` and `perm`.
// The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.
<原文结束>

# <翻译开始>
// OpenFile 用自定义的 `flag` 和 `perm` 打开文件或目录。参数 `flag` 类似于：O_RDONLY（只读），O_RDWR（读写），O_RDWR|O_CREATE|O_TRUNC 等。
// md5:0cef38d8408ed250
# <翻译结束>


<原文开始>
// OpenWithFlag opens file/directory with default perm and custom `flag`.
// The default `perm` is 0666.
// The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.
<原文结束>

# <翻译开始>
// OpenWithFlag 使用默认权限和自定义`flag`打开文件/目录。默认的`perm`为0666。
// 参数`flag`类似于：O_RDONLY（只读），O_RDWR（读写），O_RDWR|O_CREATE|O_TRUNC（读写并创建或截断），等等。
// md5:2e77d9a0acc43298
# <翻译结束>


<原文开始>
// OpenWithFlagPerm opens file/directory with custom `flag` and `perm`.
// The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.
// The parameter `perm` is like: 0600, 0666, 0777, etc.
<原文结束>

# <翻译开始>
// OpenWithFlagPerm 使用自定义的 `flag` 和 `perm` 来打开文件/目录。
// 参数 `flag` 例如：O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC 等。
// 参数 `perm` 例如：0600, 0666, 0777 等。
// md5:1f44c05cc68654d8
# <翻译结束>


<原文开始>
// Join joins string array paths with file separator of current system.
<原文结束>

# <翻译开始>
// Join 使用当前系统的文件分隔符将字符串数组路径连接起来。 md5:349ebcda51de0442
# <翻译结束>


<原文开始>
// Exists checks whether given `path` exist.
<原文结束>

# <翻译开始>
// Exists检查给定的`path`是否存在。 md5:523f33d374bd2841
# <翻译结束>


<原文开始>
// IsDir checks whether given `path` a directory.
// Note that it returns false if the `path` does not exist.
<原文结束>

# <翻译开始>
// IsDir 检查给定的 `path` 是否为目录。
// 注意，如果 `path` 不存在，它将返回 false。
// md5:c5b2468307c9c9e2
# <翻译结束>


<原文开始>
// Pwd returns absolute path of current working directory.
// Note that it returns an empty string if retrieving current
// working directory failed.
<原文结束>

# <翻译开始>
// Pwd 返回当前工作目录的绝对路径。
// 注意，如果获取当前工作目录失败，它将返回一个空字符串。
// md5:90f41f1bfdd61dba
# <翻译结束>


<原文开始>
// Chdir changes the current working directory to the named directory.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Chdir 将当前工作目录更改为指定的目录。
// 如果发生错误，该错误将为 *PathError 类型。
// md5:7bafb79d47f0c3b3
# <翻译结束>


<原文开始>
// IsFile checks whether given `path` a file, which means it's not a directory.
// Note that it returns false if the `path` does not exist.
<原文结束>

# <翻译开始>
// IsFile 检查给定的 `path` 是否为文件，即不是目录。
// 注意，如果 `path` 不存在，它将返回 false。
// md5:38595d733f36d367
# <翻译结束>


<原文开始>
// Stat returns a FileInfo describing the named file.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Stat 返回一个描述命名文件的 FileInfo。如果出现错误，错误类型为 *PathError。
// md5:f4ee45de3278c17f
# <翻译结束>


<原文开始>
// Move renames (moves) `src` to `dst` path.
// If `dst` already exists and is not a directory, it'll be replaced.
<原文结束>

# <翻译开始>
// Move 将路径 `src` 重命名（移动）到 `dst`。
// 如果 `dst` 已经存在且不是一个目录，它将被替换。
// md5:4bc635341db78f64
# <翻译结束>


<原文开始>
// Rename is alias of Move.
// See Move.
<原文结束>

# <翻译开始>
// Rename 是 Move 的别名。
// 参考 Move。
// md5:f235456881c23527
# <翻译结束>


<原文开始>
// DirNames returns sub-file names of given directory `path`.
// Note that the returned names are NOT absolute paths.
<原文结束>

# <翻译开始>
// DirNames 返回给定目录 `path` 的子文件名。请注意，返回的名称不是绝对路径。
// md5:62471f4b0e0bb389
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
// Glob 返回匹配 pattern 的所有文件名，如果没有匹配的文件，则返回 nil。pattern 的语法与 Match 中的相同。pattern 可以描述像 /usr/*/bin/ed（假设分隔符为 '/'）这样的层次名称。
// 
// Glob 忽略读取目录时发生的文件系统错误，如 I/O 错误。可能返回的唯一错误是 ErrBadPattern，当 pattern 格式不正确时。
// md5:0baeeb8710df5d67
# <翻译结束>


<原文开始>
// Remove deletes all file/directory with `path` parameter.
// If parameter `path` is directory, it deletes it recursively.
//
// It does nothing if given `path` does not exist or is empty.
<原文结束>

# <翻译开始>
// Remove 删除具有`path`参数的所有文件/目录。
// 如果参数`path`是一个目录，它会递归地删除该目录。
//
// 如果给定的`path`不存在或为空，它不做任何操作。
// md5:ae9ec60d038f0ebd
# <翻译结束>


<原文开始>
// It does nothing if `path` is empty.
<原文结束>

# <翻译开始>
// 如果`path`为空，该函数不执行任何操作。 md5:5ea1be4b22dde448
# <翻译结束>


<原文开始>
// IsReadable checks whether given `path` is readable.
<原文结束>

# <翻译开始>
// IsReadable 检查给定的 `path` 是否可读。 md5:1b38deb3c4c35233
# <翻译结束>


<原文开始>
// IsWritable checks whether given `path` is writable.
//
// TODO improve performance; use golang.org/x/sys to cross-plat-form
<原文结束>

# <翻译开始>
// IsWritable 检查给定的 `path` 是否可写。
//
// TODO 优化性能；使用 golang.org/x/sys 进行跨平台处理
// md5:2b947cc78310d3f1
# <翻译结束>


<原文开始>
// If it's a directory, create a temporary file to test whether it's writable.
<原文结束>

# <翻译开始>
// 如果是一个目录，则创建一个临时文件来测试是否可写。 md5:171566b92b9fb098
# <翻译结束>


<原文开始>
// If it's a file, check if it can open it.
<原文结束>

# <翻译开始>
// 如果它是一个文件，检查是否可以打开它。 md5:48e1a1f6b6b7d3aa
# <翻译结束>


<原文开始>
// Chmod is alias of os.Chmod.
// See os.Chmod.
<原文结束>

# <翻译开始>
// Chmod是os.Chmod的别名。
// 请参阅os.Chmod。
// md5:edb0528fe01cdccd
# <翻译结束>


<原文开始>
// Abs returns an absolute representation of path.
// If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path. The absolute
// path name for a given file is not guaranteed to be unique.
// Abs calls Clean on the result.
<原文结束>

# <翻译开始>
// Abs返回一个路径的绝对表示。
// 如果路径不是绝对路径，它将与当前工作目录连接起来，使其成为一个绝对路径。对于给定的文件，其绝对路径名称并不保证是唯一的。
// Abs会调用Clean方法处理结果。
// md5:4cb8146c59de22da
# <翻译结束>


<原文开始>
// RealPath converts the given `path` to its absolute path
// and checks if the file path exists.
// If the file does not exist, return an empty string.
<原文结束>

# <翻译开始>
// RealPath 将给定的`path`转换为其绝对路径
// 并检查文件路径是否存在。
// 如果文件不存在，则返回空字符串。
// md5:125663d904f8d81e
# <翻译结束>


<原文开始>
// SelfPath returns absolute file path of current running process(binary).
<原文结束>

# <翻译开始>
// SelfPath 返回当前运行进程（二进制文件）的绝对文件路径。 md5:87c861104977f515
# <翻译结束>


<原文开始>
// SelfName returns file name of current running process(binary).
<原文结束>

# <翻译开始>
// SelfName 返回当前运行进程（二进制文件）的文件名。 md5:1dea5b20c2c13ef6
# <翻译结束>


<原文开始>
// SelfDir returns absolute directory path of current running process(binary).
<原文结束>

# <翻译开始>
// SelfDir返回当前运行进程（二进制文件）的绝对目录路径。 md5:36d8d88a7947606c
# <翻译结束>


<原文开始>
// Basename returns the last element of path, which contains file extension.
// Trailing path separators are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of separators, Basename returns a single separator.
//
// Example:
// Basename("/var/www/file.js") -> file.js
// Basename("file.js")          -> file.js
<原文结束>

# <翻译开始>
// Basename 返回路径中的最后一个元素，该元素包含文件扩展名。
// 在提取最后一个元素之前，会移除尾随的路径分隔符。
// 如果路径为空，Base 返回 "."。
// 如果路径完全由分隔符组成，Basename 返回一个单个的分隔符。
//
// 示例：
// Basename("/var/www/file.js") -> file.js
// Basename("file.js")          -> file.js
// md5:0601675e20751381
# <翻译结束>


<原文开始>
// Name returns the last element of path without file extension.
//
// Example:
// Name("/var/www/file.js") -> file
// Name("file.js")          -> file
<原文结束>

# <翻译开始>
// Name 返回路径中不包含文件扩展名的最后一个元素。
//
// 示例：
// Name("/var/www/file.js") -> file
// Name("file.js")          -> file
// md5:231670418efd9216
# <翻译结束>


<原文开始>
// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element, Dir calls Clean on the path and trailing
// slashes are removed.
// If the `path` is empty, Dir returns ".".
// If the `path` is ".", Dir treats the path as current working directory.
// If the `path` consists entirely of separators, Dir returns a single separator.
// The returned path does not end in a separator unless it is the root directory.
//
// Example:
// Dir("/var/www/file.js") -> "/var/www"
// Dir("file.js")          -> "."
<原文结束>

# <翻译开始>
// Dir 函数返回 path 中除最后一个元素之外的所有内容，通常为路径的目录。在丢弃最后一个元素后，Dir 对路径调用 Clean 函数，并移除尾随的斜杠。
// 如果 `path` 为空，Dir 返回"."。
// 如果 `path` 为"."，Dir 将路径视为当前工作目录。
// 如果 `path` 仅由分隔符组成，Dir 返回一个单独的分隔符。
// 返回的路径除非是根目录，否则不会以分隔符结尾。
// 
// 示例：
// Dir("/var/www/file.js") -> "/var/www"
// Dir("file.js")          -> "."
// md5:03710913db229986
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
// 如果 `path` 是一个文件夹，它会检查其下是否存在任何文件。
// 如果 `path` 是一个文件，它会检查文件大小是否为零。
// 
// 注意，如果 `path` 不存在，它将返回 true。
// md5:1b96bff377b05eac
# <翻译结束>


<原文开始>
// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is
// no dot.
// Note: the result contains symbol '.'.
//
// Example:
// Ext("main.go")  => .go
// Ext("api.json") => .json
<原文结束>

# <翻译开始>
// Ext 返回路径使用的文件名扩展名。
// 扩展名是路径最后一个元素中从最后一个点开始的后缀；
// 如果没有点，则扩展名为空。
// 注意：结果包含符号'.'。
//
// 示例：
// Ext("main.go")  => .go
// Ext("api.json") => .json
// md5:63a13ee69ce09cec
# <翻译结束>


<原文开始>
// ExtName is like function Ext, which returns the file name extension used by path,
// but the result does not contain symbol '.'.
//
// Example:
// ExtName("main.go")  => go
// ExtName("api.json") => json
<原文结束>

# <翻译开始>
// ExtName 类似于函数 Ext，它返回路径中使用的文件扩展名，
// 但结果不包含符号'.'。
//
// 示例：
// ExtName("main.go")  => "go"
// ExtName("api.json") => "json"
// md5:d508af455375f787
# <翻译结束>


<原文开始>
// Temp retrieves and returns the temporary directory of current system.
//
// The optional parameter `names` specifies the sub-folders/sub-files,
// which will be joined with current system separator and returned with the path.
<原文结束>

# <翻译开始>
// Temp获取并返回当前系统的临时目录。
//
// 可选参数`names`指定了要与当前系统分隔符连接的子文件夹/子文件，将与路径一起返回。
// md5:8db9471945246517
# <翻译结束>

