
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
// ZipPathWriter compresses `paths` to `writer` using zip compressing algorithm.
// The unnecessary parameter `prefix` indicates the path prefix for zip file.
//
// Note that the parameter `paths` can be either a directory or a file, which
// supports multiple paths join with ','.
<原文结束>

# <翻译开始>
// ZipPathWriter 使用zip压缩算法将`paths`压缩到`writer`中。
// 不需要的参数`prefix`表示zip文件中的路径前缀。
//
// 注意，参数`paths`可以是目录或文件，支持使用逗号','连接多个路径。
// md5:d392a5d80ec973d9
# <翻译结束>


<原文开始>
// doZipPathWriter compresses the file of given `path` and writes the content to `zipWriter`.
// The parameter `exclude` specifies the exclusive file path that is not compressed to `zipWriter`,
// commonly the destination zip file path.
// The unnecessary parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// doZipPathWriter 将给定路径`path`的文件压缩，并将内容写入`zipWriter`。
// 参数`exclude`指定了不被压缩到`zipWriter`的排除文件路径，通常为目标zip文件路径。
// 参数`prefix`表示zip文件的路径前缀，可选。
// md5:46c9d23dcfa03c25
# <翻译结束>


<原文开始>
			// It keeps the path from file system to zip info in resource manager.
			// Usually for relative path, it makes little sense for absolute path.
<原文结束>

# <翻译开始>
// 它在资源管理器中保存从文件系统到zip信息的路径。通常对于相对路径，绝对路径意义不大。
// md5:bba8ee186d063506
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
// 它在这里计算文件名前缀，特别是打包目录。
// 例如：
// 路径：dir1
// 文件：dir1/dir2/file
// file[字符串长度(absolutePath)：] => /dir2/file
// gfile.Dir(subFilePath) => /dir2
// md5:80c4920a234839ce
# <翻译结束>


<原文开始>
// Normal handling: remove the `absolutePath`(source directory path) for file.
<原文结束>

# <翻译开始>
// 正常处理：移除文件的`absolutePath`（源目录路径）。 md5:66bfc67471cf5f63
# <翻译结束>


<原文开始>
// Add all directories to zip archive.
<原文结束>

# <翻译开始>
// 将所有目录添加到zip归档中。 md5:f8910528d8dda79d
# <翻译结束>


<原文开始>
// zipFile compresses the file of given `path` and writes the content to `zw`.
// The parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// zipFile 压缩给定路径 `path` 的文件，并将内容写入 `zw`。参数 `prefix` 表示zip文件的路径前缀。
// md5:bb4064703bf6d8ad
# <翻译结束>


<原文开始>
// Default compression level.
<原文结束>

# <翻译开始>
// 默认压缩级别。 md5:27fa604e26eb1270
# <翻译结束>


<原文开始>
// Zip header containing the info of a zip file.
<原文结束>

# <翻译开始>
// 包含ZIP文件信息的ZIP头。 md5:df2d788fe836a2e5
# <翻译结束>

