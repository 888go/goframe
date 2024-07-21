
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
// New creates and returns a file pointer pool with given file path, flag and opening permission.
//
// Note the expiration logic:
// ttl = 0 : not expired;
// ttl < 0 : immediate expired after use;
// ttl > 0 : timeout expired;
// It is not expired in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个具有给定文件路径、标志和打开权限的文件指针池。
// 
// 注意过期逻辑：
// ttl = 0：不过期；
// ttl < 0：使用后立即过期；
// ttl > 0：超时过期；
// 默认情况下，它不过期。
// md5:521d6eb77a70a063
# <翻译结束>


<原文开始>
// newFilePool creates and returns a file pointer pool with given file path, flag and opening permission.
<原文结束>

# <翻译开始>
// newFilePool 创建并返回一个具有给定文件路径、标志和打开权限的文件指针池。 md5:62fda6ff96d41c0f
# <翻译结束>


<原文开始>
// File retrieves file item from the file pointer pool and returns it. It creates one if
// the file pointer pool is empty.
// Note that it should be closed when it will never be used. When it's closed, it is not
// really closed the underlying file pointer but put back to the file pointer pool.
<原文结束>

# <翻译开始>
// File 从文件指针池中获取文件项并返回。如果文件指针池为空，它将创建一个。
// 注意，当文件不再会被使用时，应当关闭它。当它被“关闭”时，并不是真正关闭底层的文件指针，而是将其放回文件指针池中。
// md5:b6c4b6a3ade746fc
# <翻译结束>


<原文开始>
// Retrieve the state of the new created file.
<原文结束>

# <翻译开始>
// 获取新创建的文件的状态。 md5:dbe21999357cbc52
# <翻译结束>


<原文开始>
// It firstly checks using !p.init.Val() for performance purpose.
<原文结束>

# <翻译开始>
// 为了提高性能，它首先使用！p.init.Val()进行检查。 md5:bd8c9ebe349c994a
# <翻译结束>


<原文开始>
// If the file is removed or renamed, recreates the pool by increasing the pool id.
<原文结束>

# <翻译开始>
// 如果文件被删除或重命名，通过增加池ID来重新创建池。 md5:e825bec9648178de
# <翻译结束>


<原文开始>
// Clears the pool items staying in the pool.
<原文结束>

# <翻译开始>
// 清除池中残留的项目。 md5:630859bb0da3cfb4
# <翻译结束>


<原文开始>
					// It uses another adding to drop the file items between the two adding.
					// Whenever the pool id changes, the pool will be recreated.
<原文结束>

# <翻译开始>
					// 它使用另一个添加操作来移除两个添加操作之间的文件项。
					// 每当池ID改变时，池将被重新创建。
					// md5:d5f8fd9aa698b70a
# <翻译结束>


<原文开始>
// Close closes current file pointer pool.
<原文结束>

# <翻译开始>
// Close关闭当前文件指针池。 md5:01a922bcbbea5a0f
# <翻译结束>

