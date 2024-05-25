
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
// Option contains the extra options for Pack functions.
<原文结束>

# <翻译开始>
// Option包含Pack函数的额外选项。 md5:1aecf45a2bd621ac
# <翻译结束>


<原文开始>
// The file path prefix for each file item in resource manager.
<原文结束>

# <翻译开始>
// 在资源管理器中每个文件项的文件路径前缀。 md5:54cf09b52af7353f
# <翻译结束>


<原文开始>
// Keep the passed path when packing, usually for relative path.
<原文结束>

# <翻译开始>
// 在打包时保留传递的路径，通常用于相对路径。 md5:78a556a27d461bea
# <翻译结束>


<原文开始>
// Pack packs the path specified by `srcPaths` into bytes.
// The unnecessary parameter `keyPrefix` indicates the prefix for each file
// packed into the result bytes.
//
// Note that parameter `srcPaths` supports multiple paths join with ','.
//
// Deprecated: use PackWithOption instead.
<原文结束>

# <翻译开始>
// Pack 将由 `srcPaths` 指定的路径打包成字节。不必要的参数 `keyPrefix` 表示每个文件打包到结果字节中的前缀。
// 
// 注意，参数 `srcPaths` 支持用逗号分隔多个路径。
// 
// 警告：请使用 PackWithOption 替代此方法。
// md5:bba941587b4a7962
# <翻译结束>


<原文开始>
// PackWithOption packs the path specified by `srcPaths` into bytes.
//
// Note that parameter `srcPaths` supports multiple paths join with ','.
<原文结束>

# <翻译开始>
// PackWithOption 将由 `srcPaths` 指定的路径打包成字节。
// 
// 注意，参数 `srcPaths` 支持使用逗号分隔多个路径。
// md5:15ee3362e7cd91a0
# <翻译结束>


<原文开始>
// Gzip the data bytes to reduce the size.
<原文结束>

# <翻译开始>
// 使用Gzip压缩数据字节以减小大小。 md5:d15c5898ab8d9408
# <翻译结束>


<原文开始>
// PackToFile packs the path specified by `srcPaths` to target file `dstPath`.
// The unnecessary parameter `keyPrefix` indicates the prefix for each file
// packed into the result bytes.
//
// Note that parameter `srcPaths` supports multiple paths join with ','.
//
// Deprecated: use PackToFileWithOption instead.
<原文结束>

# <翻译开始>
// PackToFile 将`srcPaths`指定的路径打包到目标文件`dstPath`中。
// 不必要的参数`keyPrefix`表示打包到结果字节中的每个文件的前缀。
//
// 注意，参数`srcPaths`支持使用','连接的多个路径。
//
// 已弃用：请改用PackToFileWithOption。
// md5:222d6d9ef38edd09
# <翻译结束>


<原文开始>
// PackToFileWithOption packs the path specified by `srcPaths` to target file `dstPath`.
//
// Note that parameter `srcPaths` supports multiple paths join with ','.
<原文结束>

# <翻译开始>
// PackToFileWithOption 将由 `srcPaths` 指定的路径打包到目标文件 `dstPath` 中。
// 
// 注意，参数 `srcPaths` 支持使用逗号分隔多个路径。
// md5:5daf8e107f124634
# <翻译结束>


<原文开始>
// PackToGoFile packs the path specified by `srcPaths` to target go file `goFilePath`
// with given package name `pkgName`.
//
// The unnecessary parameter `keyPrefix` indicates the prefix for each file
// packed into the result bytes.
//
// Note that parameter `srcPaths` supports multiple paths join with ','.
//
// Deprecated: use PackToGoFileWithOption instead.
<原文结束>

# <翻译开始>
// PackToGoFile 将由 `srcPaths` 指定的路径打包成目标 Go 文件 `goFilePath`，并使用给定的包名 `pkgName`。
//
// 参数 `keyPrefix`（可选）表示打包到结果字节中的每个文件的前缀。
//
// 注意，`srcPaths` 参数支持用逗号分隔多个路径。
//
// 警告：请改用 PackToGoFileWithOption。
// md5:99701ca10a176f76
# <翻译结束>


<原文开始>
// PackToGoFileWithOption packs the path specified by `srcPaths` to target go file `goFilePath`
// with given package name `pkgName`.
//
// Note that parameter `srcPaths` supports multiple paths join with ','.
<原文结束>

# <翻译开始>
// PackToGoFileWithOption 将由 `srcPaths` 指定的路径打包到目标Go文件 `goFilePath` 中，
// 使用给定的包名 `pkgName`。
//
// 注意，参数 `srcPaths` 支持使用逗号`,`连接多个路径。
// md5:0e7ba248d1ba0543
# <翻译结束>


<原文开始>
// Unpack unpacks the content specified by `path` to []*File.
<原文结束>

# <翻译开始>
// Unpack 将由 `path` 指定的内容解压缩到 []*File 中。 md5:c88b5e566f58802e
# <翻译结束>


<原文开始>
// UnpackContent unpacks the content to []*File.
<原文结束>

# <翻译开始>
// UnpackContent 将内容解包为 []*File。 md5:a49a123f27175e6d
# <翻译结束>


<原文开始>
		// It here keeps compatible with old version packing string using hex string.
		// TODO remove this support in the future.
<原文结束>

# <翻译开始>
// 这里是为了保持与旧版本使用十六进制字符串打包字符串的兼容性。
// TODO：未来移除这个支持。
// md5:5253278930daad11
# <翻译结束>


<原文开始>
// New version packing string using base64.
<原文结束>

# <翻译开始>
// 使用base64的新版本打包字符串。 md5:c884a25b1e4334ae
# <翻译结束>


<原文开始>
// isBase64 checks and returns whether given content `s` is base64 string.
// It returns true if `s` is base64 string, or false if not.
<原文结束>

# <翻译开始>
// isBase64 检查并返回给定内容 `s` 是否为 Base64 编码的字符串。
// 如果 `s` 是 Base64 字符串，它将返回 true，否则返回 false。
// md5:314047c834f3cf6c
# <翻译结束>


<原文开始>
// isHexStr checks and returns whether given content `s` is hex string.
// It returns true if `s` is hex string, or false if not.
<原文结束>

# <翻译开始>
// isHexStr 检查并返回给定内容 `s` 是否为十六进制字符串。如果 `s` 是十六进制字符串，它将返回 true，否则返回 false。
// md5:ca395ed524f01122
# <翻译结束>


<原文开始>
// hexStrToBytes converts hex string content to []byte.
<原文结束>

# <翻译开始>
// hexStrToBytes 将十六进制字符串内容转换为[]byte。 md5:0b3c7f4ed4b490fb
# <翻译结束>

