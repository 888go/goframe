
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
// CopyOption is the option for Copy* functions.
<原文结束>

# <翻译开始>
// CopyOption 是 Copy* 函数的选项。 md5:1863c87f867e036e
# <翻译结束>


<原文开始>
// Auto call file sync after source file content copied to target file.
<原文结束>

# <翻译开始>
	// 在源文件内容复制到目标文件后，自动调用文件同步。 md5:ef1f9250b5fdabe3
# <翻译结束>


<原文开始>
	// Preserve the mode of the original file to the target file.
	// If true, the Mode attribute will make no sense.
<原文结束>

# <翻译开始>
	// 保留源文件的模式到目标文件。如果为true，Mode属性将没有意义。
	// md5:681b0704991c814c
# <翻译结束>


<原文开始>
	// Destination created file mode.
	// The default file mode is DefaultPermCopy if PreserveMode is false.
<原文结束>

# <翻译开始>
	// 创建目标文件的模式。
	// 如果PreserveMode为false，默认的文件模式是DefaultPermCopy。
	// md5:e495278ff0787785
# <翻译结束>


<原文开始>
// Copy file/directory from `src` to `dst`.
//
// If `src` is file, it calls CopyFile to implements copy feature,
// or else it calls CopyDir.
//
// If `src` is file, but `dst` already exists and is a folder,
// it then creates a same name file of `src` in folder `dst`.
//
// Eg:
// Copy("/tmp/file1", "/tmp/file2") => /tmp/file1 copied to /tmp/file2
// Copy("/tmp/dir1",  "/tmp/dir2")  => /tmp/dir1  copied to /tmp/dir2
// Copy("/tmp/file1", "/tmp/dir2")  => /tmp/file1 copied to /tmp/dir2/file1
// Copy("/tmp/dir1",  "/tmp/file2") => error
<原文结束>

# <翻译开始>
// 将源`src`文件/目录复制到目标`dst`。
//
// 如果`src`是文件，它会调用CopyFile来实现复制功能，
// 否则，它会调用CopyDir。
//
// 如果`src`是文件，但`dst`已经存在并且是一个文件夹，
// 那么它会在`dst`文件夹中创建一个与`src`同名的文件。
//
// 例如：
// Copy("/tmp/file1", "/tmp/file2") => 将/tmp/file1复制到/tmp/file2
// Copy("/tmp/dir1",  "/tmp/dir2")  => 将/tmp/dir1复制到/tmp/dir2
// Copy("/tmp/file1", "/tmp/dir2")  => 将/tmp/file1复制到/tmp/dir2/file1
// Copy("/tmp/dir1",  "/tmp/file2") => 出错
// md5:51c6598025f6b135
# <翻译结束>


<原文开始>
// CopyFile copies the contents of the file named `src` to the file named
// by `dst`. The file will be created if it does not exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file. The file mode will be copied from the source and
// the copied data is synced/flushed to stable storage.
// Thanks: https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
<原文结束>

# <翻译开始>
// CopyFile 将名为 `src` 的文件的内容复制到由 `dst` 指定的文件中。如果目标文件不存在，它将被创建。如果目标文件已存在，其所有内容将被源文件的内容替换。文件权限将从源文件复制，并且复制的数据会被同步/刷新到稳定的存储中。
// 谢谢：https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
// md5:e2fc3c25ff06fa5b
# <翻译结束>


<原文开始>
// If src and dst are the same path, it does nothing.
<原文结束>

# <翻译开始>
	// 如果src和dst是相同的路径，它不会做任何事情。 md5:1ad6359456a4bebc
# <翻译结束>


<原文开始>
// CopyDir recursively copies a directory tree, attempting to preserve permissions.
//
// Note that, the Source directory must exist and symlinks are ignored and skipped.
<原文结束>

# <翻译开始>
// CopyDir 递归地复制目录树，尝试保留权限。
//
// 注意，源目录必须存在，并且符号链接将被忽略和跳过。
// md5:4dd9167e563fa997
# <翻译结束>

