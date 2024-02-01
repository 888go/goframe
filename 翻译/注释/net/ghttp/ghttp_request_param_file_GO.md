
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
// UploadFile wraps the multipart uploading file with more and convenient features.
<原文结束>

# <翻译开始>
// UploadFile 通过提供更多的便捷功能，对使用multipart方式上传文件进行了封装。
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
# <翻译结束>


<原文开始>
// UploadFiles is an array type of *UploadFile.
<原文结束>

# <翻译开始>
// UploadFiles 是 *UploadFile 类型的数组。
# <翻译结束>


<原文开始>
// Save saves the single uploading file to directory path and returns the saved file name.
//
// The parameter `dirPath` should be a directory path, or it returns error.
//
// Note that it will OVERWRITE the target file if there's already a same name file exist.
<原文结束>

# <翻译开始>
// Save 保存单个上传的文件到指定目录路径，并返回已保存的文件名。
//
// 参数`dirPath`应为一个目录路径，否则将返回错误。
//
// 注意：如果目标位置已经存在同名文件，则会直接覆盖该文件。
# <翻译结束>


<原文开始>
// Save saves all uploading files to specified directory path and returns the saved file names.
//
// The parameter `dirPath` should be a directory path or it returns error.
//
// The parameter `randomlyRename` specifies whether randomly renames all the file names.
<原文结束>

# <翻译开始>
// Save 保存所有上传的文件到指定的目录路径，并返回已保存的文件名。
//
// 参数`dirPath`应为一个目录路径，否则将返回错误。
//
// 参数`randomlyRename`指定了是否随机重命名所有的文件名。
// 以下是详细的中文注释：
// ```go
// Save 函数的作用是将所有正在上传的文件保存到指定的目录路径下，并返回这些文件在保存后的文件名列表。
//
// 参数 `dirPath` 表示目标保存目录的路径，如果该路径不是一个有效的目录，则函数会返回错误信息。
//
// 参数 `randomlyRename` 是一个布尔值，用来指定是否对所有上传的文件进行随机重命名操作。如果设为 true，则在保存文件时将会生成随机文件名；否则，文件将以原有文件名进行保存。
# <翻译结束>


<原文开始>
// GetUploadFile retrieves and returns the uploading file with specified form name.
// This function is used for retrieving single uploading file object, which is
// uploaded using multipart form content type.
//
// It returns nil if retrieving failed or no form file with given name posted.
//
// Note that the `name` is the file field name of the multipart form from client.
<原文结束>

# <翻译开始>
// GetUploadFile 通过指定表单名称获取并返回上传文件。
// 此函数用于检索单个上传文件对象，该对象是通过多部分表单内容类型上传的。
//
// 如果检索失败或没有提交给定名称的表单文件，则返回 nil。
//
// 注意，`name` 是客户端多部分表单中文件字段的名称。
# <翻译结束>


<原文开始>
// GetUploadFiles retrieves and returns multiple uploading files with specified form name.
// This function is used for retrieving multiple uploading file objects, which are
// uploaded using multipart form content type.
//
// It returns nil if retrieving failed or no form file with given name posted.
//
// Note that the `name` is the file field name of the multipart form from client.
<原文结束>

# <翻译开始>
// GetUploadFiles 根据指定表单名称获取并返回多个上传文件。
// 此函数用于检索多个使用 multipart/form-data 类型上传的文件对象。
//
// 如果检索失败或没有接收到客户端通过给定名称上传的表单文件，将返回 nil。
//
// 注意，`name` 是客户端 multipart 表单中文件字段的名称。
# <翻译结束>

