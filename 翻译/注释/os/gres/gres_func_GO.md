
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
// Option contains the extra options for Pack functions.
<原文结束>

# <翻译开始>
// Option 包含 Pack 函数的额外选项。
# <翻译结束>


<原文开始>
// The file path prefix for each file item in resource manager.
<原文结束>

# <翻译开始>
// 在资源管理器中，每个文件项的文件路径前缀。
# <翻译结束>


<原文开始>
// Keep the passed path when packing, usually for relative path.
<原文结束>

# <翻译开始>
// 在打包时保留传递的路径，通常用于相对路径。
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
// Pack 将由 `srcPaths` 指定的路径打包成字节形式。
// 不必要的参数 `keyPrefix` 表示每个文件被打包到结果字节中时的前缀。
//
// 注意，参数 `srcPaths` 支持使用 ',' 连接的多个路径。
//
// 已弃用：请改用 PackWithOption。
# <翻译结束>


<原文开始>
// PackWithOption packs the path specified by `srcPaths` into bytes.
//
// Note that parameter `srcPaths` supports multiple paths join with ','.
<原文结束>

# <翻译开始>
// PackWithOption 函数将由 `srcPaths` 指定的路径打包成字节形式。
//
// 注意，参数 `srcPaths` 支持使用 ',' 连接的多个路径。
# <翻译结束>


<原文开始>
// Gzip the data bytes to reduce the size.
<原文结束>

# <翻译开始>
// 使用Gzip压缩数据字节以减少其大小。
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
// PackToFile 将由`srcPaths`指定的路径打包到目标文件`dstPath`中。
// 不必要的参数`keyPrefix`表示每个被打包到结果字节中的文件前缀。
//
// 注意，参数`srcPaths`支持使用','连接的多个路径。
//
// 已弃用：请改用PackToFileWithOption。
# <翻译结束>


<原文开始>
// PackToFileWithOption packs the path specified by `srcPaths` to target file `dstPath`.
//
// Note that parameter `srcPaths` supports multiple paths join with ','.
<原文结束>

# <翻译开始>
// PackToFileWithOption 将由 `srcPaths` 指定的路径打包到目标文件 `dstPath`。
//
// 注意，参数 `srcPaths` 支持通过 ',' 连接的多个路径。
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
// PackToGoFile 将由 `srcPaths` 指定的路径打包到目标 Go 文件 `goFilePath`，
// 并使用给定的包名 `pkgName`。
//
// 非必需参数 `keyPrefix` 表示每个打包进结果字节流文件的前缀。
//
// 注意，参数 `srcPaths` 支持以 ',' 连接的多个路径。
//
// 已弃用：请改用 PackToGoFileWithOption。
# <翻译结束>


<原文开始>
// PackToGoFileWithOption packs the path specified by `srcPaths` to target go file `goFilePath`
// with given package name `pkgName`.
//
// Note that parameter `srcPaths` supports multiple paths join with ','.
<原文结束>

# <翻译开始>
// PackToGoFileWithOption 根据指定的 `srcPaths` 路径将文件打包到目标 Go 文件 `goFilePath`，
// 同时使用给定的包名 `pkgName`。
//
// 注意，参数 `srcPaths` 支持通过 ',' 连接多个路径。
# <翻译结束>


<原文开始>
// Unpack unpacks the content specified by `path` to []*File.
<原文结束>

# <翻译开始>
// Unpack 将由 `path` 指定的内容解包为 []*File 类型的切片。
# <翻译结束>


<原文开始>
// UnpackContent unpacks the content to []*File.
<原文结束>

# <翻译开始>
// UnpackContent 解析内容到 []*File 类型的切片。
# <翻译结束>


<原文开始>
		// It here keeps compatible with old version packing string using hex string.
		// TODO remove this support in the future.
<原文结束>

# <翻译开始>
// 这里保留了使用十六进制字符串进行旧版本打包字符串的兼容性。
// TODO：未来将删除对此种支持。
# <翻译结束>


<原文开始>
// New version packing string using base64.
<原文结束>

# <翻译开始>
// 新版本使用base64对字符串进行打包。
# <翻译结束>


<原文开始>
// isBase64 checks and returns whether given content `s` is base64 string.
// It returns true if `s` is base64 string, or false if not.
<原文结束>

# <翻译开始>
// isBase64 检查并返回给定内容 `s` 是否为 base64 字符串。
// 如果 `s` 是 base64 字符串，则返回 true；否则返回 false。
# <翻译结束>


<原文开始>
// isHexStr checks and returns whether given content `s` is hex string.
// It returns true if `s` is hex string, or false if not.
<原文结束>

# <翻译开始>
// isHexStr 检查并返回给定内容 `s` 是否为十六进制字符串。
// 如果 `s` 是十六进制字符串，则返回 true；否则返回 false。
# <翻译结束>


<原文开始>
// hexStrToBytes converts hex string content to []byte.
<原文结束>

# <翻译开始>
// hexStrToBytes 将十六进制字符串内容转换为 []byte 类型。
# <翻译结束>

