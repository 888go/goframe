
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
// UploadFile wraps the multipart uploading file with more and convenient features.
<原文结束>

# <翻译开始>
// UploadFile 包装了multipart上传文件，提供了更多和更方便的功能。. md5:a7173285d087c4aa
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。. md5:43c3b36e60a18f9a
# <翻译结束>


<原文开始>
// UploadFiles is an array type of *UploadFile.
<原文结束>

# <翻译开始>
// UploadFiles 是 *UploadFile 的数组类型。. md5:94b6aef81609f12b
# <翻译结束>


<原文开始>
// Save saves the single uploading file to directory path and returns the saved file name.
//
// The parameter `dirPath` should be a directory path, or it returns error.
//
// Note that it will OVERWRITE the target file if there's already a same name file exist.
<原文结束>

# <翻译开始>
// Save 将单个上传的文件保存到指定的目录路径，并返回保存的文件名。
//
// 参数 `dirPath` 应为一个目录路径，否则会返回错误。
//
// 注意：如果目标位置已经存在同名文件，该函数将覆盖原有的文件。
// md5:ffe3d8f90d14185a
# <翻译结束>


<原文开始>
// Save saves all uploading files to specified directory path and returns the saved file names.
//
// The parameter `dirPath` should be a directory path or it returns error.
//
// The parameter `randomlyRename` specifies whether randomly renames all the file names.
<原文结束>

# <翻译开始>
// Save 将所有上传的文件保存到指定的目录路径，并返回保存的文件名。
//
// 参数 `dirPath` 应该是一个目录路径，否则会返回错误。
//
// 参数 `randomlyRename` 指定是否为所有文件名随机重命名。
// md5:de2b45ea5a89ccad
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
// GetUploadFile 通过指定的表单名称检索并返回上传中的文件。
// 此函数用于检索使用multipart/form-data内容类型上传的单个文件对象。
//
// 如果检索失败或没有给定名称的表单文件被上传，它将返回nil。
//
// 注意，`name` 是客户端multipart表单中文件字段的名称。
// md5:a49268bd7e014ab6
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
// GetUploadFiles 用于检索并返回具有指定表单名称的多个上传文件。
// 此函数用于获取多个上传文件对象，这些对象是使用多部分表单内容类型上传的。
//
// 如果检索失败或没有给定名称的表单文件被上传，则返回nil。
//
// 注意，`name` 是来自客户端的多部分表单中的文件字段名称。
// md5:cbbf4db398137505
# <翻译结束>

