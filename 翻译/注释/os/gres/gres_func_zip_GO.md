
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
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// ZipPathWriter compresses `paths` to `writer` using zip compressing algorithm.
// The unnecessary parameter `prefix` indicates the path prefix for zip file.
//
// Note that the parameter `paths` can be either a directory or a file, which
// supports multiple paths join with ','.
<原文结束>

# <翻译开始>
// ZipPathWriter 使用zip压缩算法将`paths`压缩到`writer`。
// 不必要的参数`prefix`表示zip文件的路径前缀。
//
// 注意，参数`paths`可以是目录或文件，支持使用','连接的多个路径。
# <翻译结束>


<原文开始>
// doZipPathWriter compresses the file of given `path` and writes the content to `zipWriter`.
// The parameter `exclude` specifies the exclusive file path that is not compressed to `zipWriter`,
// commonly the destination zip file path.
// The unnecessary parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// doZipPathWriter 函数用于压缩指定 `path` 的文件并将压缩内容写入到 `zipWriter`。
// 参数 `exclude` 指定不需要被压缩到 `zipWriter` 中的文件路径，通常是指定的目标 zip 文件路径本身。
// 参数 `prefix`（非必需）表示 zip 文件中的路径前缀。
# <翻译结束>


<原文开始>
			// It keeps the path from file system to zip info in resource manager.
			// Usually for relative path, it makes little sense for absolute path.
<原文结束>

# <翻译开始>
// 它在资源管理器中保留从文件系统到zip信息的路径。
// 通常对于相对路径有意义，但对于绝对路径意义不大。
# <翻译结束>


<原文开始>
		// It here calculates the file name prefix, especially packing the directory.
		// Eg:
		// path: dir1
		// file: dir1/dir2/file
		// file[len(absolutePath):] => /dir2/file
		// gfile.Dir(subFilePath)   => /dir2
<原文结束>

# <翻译开始>
// 这里计算文件名前缀，特别是打包目录。
// 例如：
// 路径：dir1
// 文件：dir1/dir2/file
// file[len(absolutePath):] => /dir2/file （取绝对路径后缀部分）
// gfile.Dir(subFilePath)   => /dir2 （获取子文件路径的目录部分）
# <翻译结束>


<原文开始>
// Normal handling: remove the `absolutePath`(source directory path) for file.
<原文结束>

# <翻译开始>
// 正常处理：从文件中移除`absolutePath`(源目录路径)。
# <翻译结束>


<原文开始>
// Add all directories to zip archive.
<原文结束>

# <翻译开始>
// 将所有目录添加到zip归档中。
# <翻译结束>


<原文开始>
// zipFile compresses the file of given `path` and writes the content to `zw`.
// The parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// zipFile 将给定 `path` 的文件压缩，并将内容写入 `zw`。
// 参数 `prefix` 表示 zip 文件的路径前缀。
# <翻译结束>


<原文开始>
// Zip header containing the info of a zip file.
<原文结束>

# <翻译开始>
// Zip头包含zip文件的信息。
# <翻译结束>


<原文开始>
// Default compression level.
<原文结束>

# <翻译开始>
// 默认压缩级别。
# <翻译结束>

