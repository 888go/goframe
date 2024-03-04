
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
// CopyOption is the option for Copy* functions.
<原文结束>

# <翻译开始>
// CopyOption 是用于 Copy* 函数的选项。
# <翻译结束>


<原文开始>
// Auto call file sync after source file content copied to target file.
<原文结束>

# <翻译开始>
// 在源文件内容复制到目标文件后自动调用文件同步
# <翻译结束>


<原文开始>
	// Preserve the mode of the original file to the target file.
	// If true, the Mode attribute will make no sense.
<原文结束>

# <翻译开始>
// 保留原始文件的模式到目标文件。
// 如果为 true，则 Mode 属性将不起作用。
# <翻译结束>


<原文开始>
	// Destination created file mode.
	// The default file mode is DefaultPermCopy if PreserveMode is false.
<原文结束>

# <翻译开始>
// 目标文件创建时的模式
// 若PreserveMode为false，默认的文件模式为DefaultPermCopy
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
// 将文件/目录从`src`复制到`dst`。
//
// 如果`src`是文件，它将调用CopyFile实现复制功能，
// 否则调用CopyDir。
//
// 如果`src`是文件，但`dst`已存在且是一个文件夹，
// 则在`dst`目录下创建一个与`src`同名的文件。
//
// 示例：
// Copy("/tmp/file1", "/tmp/file2") => 将/tmp/file1复制到/tmp/file2
// Copy("/tmp/dir1",  "/tmp/dir2")  => 将/tmp/dir1复制到/tmp/dir2
// Copy("/tmp/file1", "/tmp/dir2")  => 将/tmp/file1复制到/tmp/dir2/file1
// Copy("/tmp/dir1",  "/tmp/file2") => 报错
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
// CopyFile 将名为 `src` 的文件内容复制到名为 `dst` 的文件中。如果目标文件不存在，将会创建该文件。如果目标文件已存在，则其所有内容将被源文件内容替换。文件模式将从源文件复制，并且复制的数据将同步/刷新到稳定的存储设备中。
// 感谢：https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
# <翻译结束>


<原文开始>
// If src and dst are the same path, it does nothing.
<原文结束>

# <翻译开始>
// 如果src和dst是相同的路径，则不做任何操作。
# <翻译结束>







<原文开始>
// CopyDir recursively copies a directory tree, attempting to preserve permissions.
//
// Note that, the Source directory must exist and symlinks are ignored and skipped.
<原文结束>

# <翻译开始>
// CopyDir递归地复制一个目录树，尝试保持原有的权限设置。
//
// 注意：源目录必须存在，并且符号链接会被忽略并跳过。
# <翻译结束>






