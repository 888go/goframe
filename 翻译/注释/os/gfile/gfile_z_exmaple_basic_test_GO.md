
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
// Check if directory exists
<原文结束>

# <翻译开始>
// 检查目录是否存在. md5:0c502e5e10c3d1bc
# <翻译结束>


<原文开始>
// Check whether the file exists
<原文结束>

# <翻译开始>
// 检查文件是否存在. md5:d1455b2a0aa17f63
# <翻译结束>


<原文开始>
// Creates file with given `path` recursively
<原文结束>

# <翻译开始>
// 递归创建具有给定`path`的文件. md5:587a4af68c8bc5ac
# <翻译结束>


<原文开始>
// Write some content to file
<原文结束>

# <翻译开始>
// 向文件中写入一些内容. md5:856ea5269b5be5ff
# <翻译结束>


<原文开始>
// Reads len(b) bytes from the File
<原文结束>

# <翻译开始>
// 从File中读取len(b)字节. md5:a14d5883b14d9063
# <翻译结束>


<原文开始>
// Open file or directory with READONLY model
<原文结束>

# <翻译开始>
// 使用只读模式打开文件或目录. md5:78e9e881c189899d
# <翻译结束>


<原文开始>
	// Opens file/directory with custom `flag` and `perm`
	// Create if file does not exist,it is created in a readable and writable mode,prem 0777
<原文结束>

# <翻译开始>
// 使用自定义的 `flag` 和 `perm` 打开文件/目录
// 如果文件不存在，则创建一个可读写模式的文件，权限默认为 0777
// md5:77b0a10407d251c5
# <翻译结束>


<原文开始>
	// Opens file/directory with custom `flag`
	// Create if file does not exist,it is created in a readable and writable mode with default `perm` is 0666
<原文结束>

# <翻译开始>
// 使用自定义的`flag`打开文件/目录
// 如果文件不存在，将创建文件，并以可读写模式打开，默认的`perm`权限为0666
// md5:510ad8864d50d6b6
# <翻译结束>


<原文开始>
// Joins string array paths with file separator of current system.
<原文结束>

# <翻译开始>
// 使用当前系统文件分隔符将字符串数组路径连接起来。 md5:729553e2f763ca20
# <翻译结束>


<原文开始>
	// May Output:
	// /tmp/gfile_example_basic_dir/file1
<原文结束>

# <翻译开始>
	// May Output:
	// /tmp/gfile_example_basic_dir/file1
# <翻译结束>


<原文开始>
// Checks whether given `path` exist.
<原文结束>

# <翻译开始>
// 检查给定的`path`是否存在。 md5:801440e91778729a
# <翻译结束>


<原文开始>
// Checks whether given `path` a directory.
<原文结束>

# <翻译开始>
// 检查给定的`path`是否为目录。 md5:5744f2242b1a0948
# <翻译结束>


<原文开始>
// Get absolute path of current working directory.
<原文结束>

# <翻译开始>
// 获取当前工作目录的绝对路径。 md5:02d8656598c3d01b
# <翻译结束>


<原文开始>
	// May Output:
	// xxx/gf/os/gfile
<原文结束>

# <翻译开始>
	// May Output:
	// xxx/gf/os/gfile
# <翻译结束>


<原文开始>
// Get current working directory
<原文结束>

# <翻译开始>
// 获取当前工作目录. md5:87642df8d64a090c
# <翻译结束>


<原文开始>
// Changes the current working directory to the named directory.
<原文结束>

# <翻译开始>
// 将当前工作目录更改为指定的目录。 md5:c7ba95b4405caafe
# <翻译结束>


<原文开始>
	// May Output:
	// xxx/gf/os/gfile
	// /tmp/gfile_example_basic_dir/file1
<原文结束>

# <翻译开始>
	// May Output:
	// xxx/gf/os/gfile
	// /tmp/gfile_example_basic_dir/file1
# <翻译结束>


<原文开始>
// Checks whether given `path` a file, which means it's not a directory.
<原文结束>

# <翻译开始>
// 检查给定的`path`是否为文件，这意味着它不是目录。 md5:cb0ae2363ad14139
# <翻译结束>


<原文开始>
// Get a FileInfo describing the named file.
<原文结束>

# <翻译开始>
// 获取关于指定文件的FileInfo信息。 md5:189ffdaf06730055
# <翻译结束>


<原文开始>
	// May Output:
	// file1
	// false
	// -rwxr-xr-x
	// 2021-12-02 11:01:27.261441694 +0800 CST
	// &{16777220 33261 1 8597857090 501 20 0 [0 0 0 0] {1638414088 192363490} {1638414087 261441694} {1638414087 261441694} {1638413480 485068275} 38 8 4096 0 0 0 [0 0]}
<原文结束>

# <翻译开始>
	// May Output:
	// file1
	// false
	// -rwxr-xr-x
	// 2021-12-02 11:01:27.261441694 +0800 CST
	// &{16777220 33261 1 8597857090 501 20 0 [0 0 0 0] {1638414088 192363490} {1638414087 261441694} {1638414087 261441694} {1638413480 485068275} 38 8 4096 0 0 0 [0 0]}
# <翻译结束>


<原文开始>
	//  Moves `src` to `dst` path.
	// If `dst` already exists and is not a directory, it'll be replaced.
<原文结束>

# <翻译开始>
// 将`src`移动到`dst`路径。
// 如果`dst`已经存在且不是目录，它将被替换。
// md5:3401f06a2c8ccd49
# <翻译结束>


<原文开始>
	//  renames (moves) `src` to `dst` path.
	// If `dst` already exists and is not a directory, it'll be replaced.
<原文结束>

# <翻译开始>
// 将`src`路径重命名（移动）到`dst`。
// 如果`dst`已经存在且不是目录，它将被替换。
// md5:b028a167dc2de1d0
# <翻译结束>


<原文开始>
// Get sub-file names of given directory `path`.
<原文结束>

# <翻译开始>
// 获取给定目录`path`下的子文件名。 md5:a7ba80d33218bf78
# <翻译结束>


<原文开始>
	// May Output:
	// [file1]
<原文结束>

# <翻译开始>
	// May Output:
	// [file1]
# <翻译结束>


<原文开始>
	// Get sub-file names of given directory `path`.
	// Only show file name
<原文结束>

# <翻译开始>
// 获取给定目录`path`下的子文件名。
// 只显示文件名
// md5:438be3da9f43a720
# <翻译结束>


<原文开始>
// Show full path of the file
<原文结束>

# <翻译开始>
// 显示文件的完整路径. md5:d246b83579c32f8a
# <翻译结束>


<原文开始>
	// May Output:
	// [gfile_z_example_basic_test.go]
	// [xxx/gf/os/gfile/gfile_z_example_basic_test.go]
<原文结束>

# <翻译开始>
	// May Output:
	// [gfile_z_example_basic_test.go]
	// [xxx/gf/os/gfile/gfile_z_example_basic_test.go]
# <翻译结束>


<原文开始>
// Checks whether given `path` is readable.
<原文结束>

# <翻译开始>
// 检查给定的`path`是否可读。 md5:fda74ad537c20ca3
# <翻译结束>


<原文开始>
// Checks whether given `path` is writable.
<原文结束>

# <翻译开始>
// 检查给定的`path`是否可写。 md5:cbf170ef62b28ee0
# <翻译结束>


<原文开始>
// Get an absolute representation of path.
<原文结束>

# <翻译开始>
// 获取path的绝对表示形式。 md5:9e6cadaac30f8871
# <翻译结束>


<原文开始>
// fetch an absolute representation of path.
<原文结束>

# <翻译开始>
// 获取path的绝对表示形式。 md5:cca3127b33ff195c
# <翻译结束>


<原文开始>
// Get absolute file path of current running process
<原文结束>

# <翻译开始>
// 获取当前运行进程的绝对文件路径. md5:976eb91d29aba4fd
# <翻译结束>


<原文开始>
	// May Output:
	// xxx/___github_com_gogf_gf_v2_os_gfile__ExampleSelfPath
<原文结束>

# <翻译开始>
	// May Output:
	// xxx/___github_com_gogf_gf_v2_os_gfile__ExampleSelfPath
# <翻译结束>


<原文开始>
// Get file name of current running process
<原文结束>

# <翻译开始>
// 获取当前正在运行进程的文件名. md5:d2f55580550d36cc
# <翻译结束>


<原文开始>
	// May Output:
	// ___github_com_gogf_gf_v2_os_gfile__ExampleSelfName
<原文结束>

# <翻译开始>
	// May Output:
	// ___github_com_gogf_gf_v2_os_gfile__ExampleSelfName
# <翻译结束>


<原文开始>
// Get absolute directory path of current running process
<原文结束>

# <翻译开始>
// 获取当前运行进程的绝对目录路径. md5:f0b7c37862a2865b
# <翻译结束>


<原文开始>
	// May Output:
	// /private/var/folders/p6/gc_9mm3j229c0mjrjp01gqn80000gn/T
<原文结束>

# <翻译开始>
	// May Output:
	// /private/var/folders/p6/gc_9mm3j229c0mjrjp01gqn80000gn/T
# <翻译结束>


<原文开始>
// Get the last element of path, which contains file extension.
<原文结束>

# <翻译开始>
// 获取路径中的最后一个元素，该元素包含文件扩展名。 md5:4868d5ea79029f54
# <翻译结束>


<原文开始>
// Get the last element of path without file extension.
<原文结束>

# <翻译开始>
// 获取路径中不包括文件扩展名的最后一个元素。 md5:8291b4d785e21395
# <翻译结束>


<原文开始>
// Get all but the last element of path, typically the path's directory.
<原文结束>

# <翻译开始>
// 获取路径中除最后一个元素外的所有部分，通常是路径的目录部分。 md5:21ab4b575c298060
# <翻译结束>


<原文开始>
// Check whether the `path` is empty
<原文结束>

# <翻译开始>
// 检查`path`是否为空. md5:87c020da5f9bc2aa
# <翻译结束>


<原文开始>
// Get the file name extension used by path.
<原文结束>

# <翻译开始>
// 获取path所使用的文件扩展名。 md5:5a50317e9cb8596e
# <翻译结束>


<原文开始>
// Get the file name extension used by path but the result does not contains symbol '.'.
<原文结束>

# <翻译开始>
// 从路径中获取文件扩展名，但结果不包含'.'符号。 md5:0a63ac6fbba1d676
# <翻译结束>


<原文开始>
// deletes all file/directory with `path` parameter.
<原文结束>

# <翻译开始>
// 使用`path`参数删除所有文件/目录。 md5:8d2699993a255ec6
# <翻译结束>

