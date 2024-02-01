
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
// Open creates and returns a file item with given file path, flag and opening permission.
// It automatically creates an associated file pointer pool internally when it's called first time.
// It retrieves a file item from the file pointer pool after then.
<原文结束>

# <翻译开始>
// Open 函数通过给定的文件路径、标志和打开权限创建并返回一个文件项。
// 当首次调用时，它会自动内部创建一个相关的文件指针池。
// 此后，它将从该文件指针池中获取文件项。
# <翻译结束>


<原文开始>
	// DO NOT search the path here wasting performance!
	// Leave following codes just for warning you.
	//
	// path, err = gfile.Search(path)
	// if err != nil {
	//	return nil, err
	// }
<原文结束>

# <翻译开始>
// **不要**在这里搜索路径以免浪费性能！
// 保留以下代码只是为了给您一个警告。
//
// path, err = gfile.Search(path)
// if err != nil {
//     return nil, err
// }
# <翻译结束>


<原文开始>
// Get returns a file item with given file path, flag and opening permission.
// It retrieves a file item from the file pointer pool after then.
<原文结束>

# <翻译开始>
// Get 函数通过给定的文件路径、标志和打开权限获取一个文件项。
// 然后，它从文件指针池中检索一个文件项。
# <翻译结束>


<原文开始>
// Stat returns the FileInfo structure describing file.
<原文结束>

# <翻译开始>
// Stat返回描述文件的FileInfo结构体。
# <翻译结束>


<原文开始>
// Close puts the file pointer back to the file pointer pool.
<原文结束>

# <翻译开始>
// Close将文件指针放回文件指针池。
# <翻译结束>

