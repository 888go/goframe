
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
// Open creates and returns a file item with given file path, flag and opening permission.
// It automatically creates an associated file pointer pool internally when it's called first time.
// It retrieves a file item from the file pointer pool after then.
<原文结束>

# <翻译开始>
// Open 函数根据给定的文件路径、标志和打开权限创建并返回一个文件项。当它首次被调用时，它会自动内部创建一个关联的文件指针池。然后，它从文件指针池中获取文件项。
// md5:94bbe2b7d15d2c1f
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
	// 不要在這裡浪費性能搜索路径！
	// 保留以下代码只是为了警告你。
	//
	// path, err := gfile.Search(path)
	// 如果 err 不为 nil，则：
	// 返回 nil 和 err
	// md5:763fc7901895ec01
# <翻译结束>


<原文开始>
// Get returns a file item with given file path, flag and opening permission.
// It retrieves a file item from the file pointer pool after then.
<原文结束>

# <翻译开始>
// Get 根据给定的文件路径、标志和打开权限返回一个文件项。
// 随后，它从文件指针池中检索一个文件项。
// md5:f56943d16a070df7
# <翻译结束>


<原文开始>
// Stat returns the FileInfo structure describing file.
<原文结束>

# <翻译开始>
// Stat 返回描述文件的FileInfo结构。 md5:86e6f3f0a508aa53
# <翻译结束>


<原文开始>
// Close puts the file pointer back to the file pointer pool.
<原文结束>

# <翻译开始>
// Close 将文件指针放回文件指针池。 md5:a47bacf277b7f774
# <翻译结束>

