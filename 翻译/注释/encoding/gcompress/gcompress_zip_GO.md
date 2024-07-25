
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
// ZipPath compresses `fileOrFolderPaths` to `dstFilePath` using zip compressing algorithm.
//
// The parameter `paths` can be either a directory or a file, which
// supports multiple paths join with ','.
// The unnecessary parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// ZipPath 使用zip压缩算法将`fileOrFolderPaths`压缩到`dstFilePath`。
//
// 参数`paths`可以是目录或文件，支持使用`,`连接多个路径。参数`prefix`（可选）表示zip文件的路径前缀。
// md5:6754e1656d2dfc22
# <翻译结束>


<原文开始>
// ZipPathWriter compresses `fileOrFolderPaths` to `writer` using zip compressing algorithm.
//
// Note that the parameter `fileOrFolderPaths` can be either a directory or a file, which
// supports multiple paths join with ','.
// The unnecessary parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// ZipPathWriter 使用zip压缩算法将`fileOrFolderPaths`压缩到`writer`中。
// 
// 注意，参数`fileOrFolderPaths`可以是目录或文件，支持使用','连接多个路径。
// 参数`prefix`（可选）表示zip文件的路径前缀。
// md5:0e6a4ca6fdf7a9d7
# <翻译结束>


<原文开始>
// ZipPathContent compresses `fileOrFolderPaths` to []byte using zip compressing algorithm.
//
// Note that the parameter `fileOrFolderPaths` can be either a directory or a file, which
// supports multiple paths join with ','.
// The unnecessary parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// ZipPathContent 使用zip压缩算法将`fileOrFolderPaths`压缩为[]byte。
//
// 注意，参数`fileOrFolderPaths`可以是目录或文件，支持使用逗号','连接多个路径。
// 不强制要求的参数`prefix`表示zip文件中的路径前缀。
// md5:6700858e8ecb32a5
# <翻译结束>


<原文开始>
// doZipPathWriter compresses given `fileOrFolderPaths` and writes the content to `zipWriter`.
//
// The parameter `fileOrFolderPath` can be either a single file or folder path.
// The parameter `exclude` specifies the exclusive file path that is not compressed to `zipWriter`,
// commonly the destination zip file path.
// The unnecessary parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// doZipPathWriter 将给定的 `fileOrFolderPaths` 压缩，并将内容写入 `zipWriter`。
//
// 参数 `fileOrFolderPath` 可以是一个单一的文件或文件夹路径。
// 参数 `exclude` 指定了不应被压缩到 `zipWriter` 的排除文件路径，通常为目标zip文件路径。
// 参数 `prefix` 是用于zip文件的路径前缀，一般不需要。
// md5:491b5e660bfd8ac9
# <翻译结束>


<原文开始>
// UnZipFile decompresses `archive` to `dstFolderPath` using zip compressing algorithm.
//
// The parameter `dstFolderPath` should be a directory.
// The optional parameter `zippedPrefix` specifies the unzipped path of `zippedFilePath`,
// which can be used to specify part of the archive file to unzip.
<原文结束>

# <翻译开始>
// UnZipFile 使用 ZIP 压缩算法将 `archive` 解压缩到 `dstFolderPath`。
//
// 参数 `dstFolderPath` 应该是一个目录。可选参数 `zippedPrefix` 指定了 `zippedFilePath` 的解压缩路径部分，可以用来指定要解压缩的归档文件的一部分。
// md5:4ef9114de36ab1d8
# <翻译结束>


<原文开始>
// UnZipContent decompresses `zippedContent` to `dstFolderPath` using zip compressing algorithm.
//
// The parameter `dstFolderPath` should be a directory.
// The parameter `zippedPrefix` specifies the unzipped path of `zippedContent`,
// which can be used to specify part of the archive file to unzip.
<原文结束>

# <翻译开始>
// UnZipContent 使用zip压缩算法将`zippedContent`解压缩到`dstFolderPath`。
//
// 参数`dstFolderPath`应该是一个目录。参数`zippedPrefix`指定了`zippedContent`的解压路径，可以用来指定要解压的归档文件的一部分。
// md5:808f21381d5e3681
# <翻译结束>


<原文开始>
// The fileReader is closed in function doCopyForUnZipFileWithReader.
<原文结束>

# <翻译开始>
		// 文件读取器在函数doCopyForUnZipFileWithReader中被关闭。 md5:bdbed60d16aa0ca2
# <翻译结束>


<原文开始>
// zipFile compresses the file of given `filePath` and writes the content to `zw`.
// The parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// zipFile 将给定的 `filePath` 文件压缩，并将内容写入 `zw`。
// 参数 `prefix` 用于表示在压缩文件中的路径前缀。
// md5:69f2856c4cb49f38
# <翻译结束>

