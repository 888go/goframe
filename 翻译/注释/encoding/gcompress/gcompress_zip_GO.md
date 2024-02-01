
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
// ZipPath compresses `fileOrFolderPaths` to `dstFilePath` using zip compressing algorithm.
//
// The parameter `paths` can be either a directory or a file, which
// supports multiple paths join with ','.
// The unnecessary parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// ZipPath 使用zip压缩算法将`fileOrFolderPaths`压缩到`dstFilePath`。
//
// 参数`paths`可以是目录或文件，支持使用','连接的多个路径。
// 可选参数`prefix`表示zip文件中的路径前缀。
# <翻译结束>


<原文开始>
// ZipPathWriter compresses `fileOrFolderPaths` to `writer` using zip compressing algorithm.
//
// Note that the parameter `fileOrFolderPaths` can be either a directory or a file, which
// supports multiple paths join with ','.
// The unnecessary parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// ZipPathWriter 使用zip压缩算法将`fileOrFolderPaths`压缩到`writer`。
//
// 注意参数`fileOrFolderPaths`可以是目录或文件，支持使用','连接的多个路径。
// 可选参数`prefix`表示zip文件中的路径前缀。
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
// 注意，参数`fileOrFolderPaths`可以是目录或文件，支持使用','连接多个路径。
// 可选参数`prefix`表示zip文件中的路径前缀。
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
// doZipPathWriter 将给定的 `fileOrFolderPaths` 进行压缩，并将内容写入 `zipWriter`。
//
// 参数 `fileOrFolderPath` 可以是单个文件或文件夹路径。
// 参数 `exclude` 指定了不被压缩到 `zipWriter` 中的排除文件路径，通常是指定的目标 zip 文件路径。
// 非必需参数 `prefix` 表示 zip 文件的路径前缀。
# <翻译结束>


<原文开始>
// UnZipFile decompresses `archive` to `dstFolderPath` using zip compressing algorithm.
//
// The parameter `dstFolderPath` should be a directory.
// The optional parameter `zippedPrefix` specifies the unzipped path of `zippedFilePath`,
// which can be used to specify part of the archive file to unzip.
<原文结束>

# <翻译开始>
// UnZipFile 使用zip压缩算法将`archive`解压到`dstFolderPath`。
//
// 参数`dstFolderPath`应为一个目录。
// 可选参数`zippedPrefix`用于指定`zippedFilePath`解压后的路径前缀，
// 该参数可用于指定只解压归档文件中的部分内容。
# <翻译结束>


<原文开始>
// UnZipContent decompresses `zippedContent` to `dstFolderPath` using zip compressing algorithm.
//
// The parameter `dstFolderPath` should be a directory.
// The parameter `zippedPrefix` specifies the unzipped path of `zippedContent`,
// which can be used to specify part of the archive file to unzip.
<原文结束>

# <翻译开始>
// UnZipContent 使用zip压缩算法将`zippedContent`解压到`dstFolderPath`。
//
// 参数`dstFolderPath`应该是一个目录。
// 参数`zippedPrefix`指定了`zippedContent`解压后的路径，
// 可用于指定要解压的归档文件的部分。
// 进一步细化翻译：
// ```go
// UnZipContent 函数负责使用ZIP压缩算法将压缩内容 `zippedContent` 解压到目标文件夹 `dstFolderPath`。
//
// 参数 `dstFolderPath` 必须是一个存在的目录，解压后的文件将存放在此目录下。
// 参数 `zippedPrefix` 指定 `zippedContent` 中待解压内容的相对路径前缀，
// 通过此参数可以选择性地解压归档文件中的特定部分。
# <翻译结束>


<原文开始>
// The fileReader is closed in function doCopyForUnZipFileWithReader.
<原文结束>

# <翻译开始>
// 文件读取器在函数doCopyForUnZipFileWithReader中关闭。
# <翻译结束>


<原文开始>
// zipFile compresses the file of given `filePath` and writes the content to `zw`.
// The parameter `prefix` indicates the path prefix for zip file.
<原文结束>

# <翻译开始>
// zipFile 将指定 `filePath` 的文件进行压缩，并将压缩内容写入 `zw`。
// 参数 `prefix` 表示压缩文件路径的前缀。
# <翻译结束>

