
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
// New creates and returns a file pointer pool with given file path, flag and opening permission.
//
// Note the expiration logic:
// ttl = 0 : not expired;
// ttl < 0 : immediate expired after use;
// ttl > 0 : timeout expired;
// It is not expired in default.
<原文结束>

# <翻译开始>
// New根据给定的文件路径、标志和打开权限创建并返回一个文件指针池。
//
// 注意过期逻辑：
// ttl = 0 : 不过期；
// ttl < 0 : 使用后立即过期；
// ttl > 0 : 超时后过期；
// 默认情况下，它不会过期。
# <翻译结束>


<原文开始>
// newFilePool creates and returns a file pointer pool with given file path, flag and opening permission.
<原文结束>

# <翻译开始>
// newFilePool 根据给定的文件路径、标志和打开权限创建并返回一个文件指针池。
# <翻译结束>


<原文开始>
// File retrieves file item from the file pointer pool and returns it. It creates one if
// the file pointer pool is empty.
// Note that it should be closed when it will never be used. When it's closed, it is not
// really closed the underlying file pointer but put back to the file pointer pool.
<原文结束>

# <翻译开始>
// File 从文件指针池中获取文件项并返回，如果文件指针池为空，则创建一个新的文件项。
// 注意：当文件项不再使用时，应关闭它。当其被关闭时，并非真正关闭底层的文件指针，而是将其放回文件指针池中。
# <翻译结束>


<原文开始>
// Retrieve the state of the new created file.
<原文结束>

# <翻译开始>
// 获取新创建文件的状态。
# <翻译结束>


<原文开始>
// It firstly checks using !p.init.Val() for performance purpose.
<原文结束>

# <翻译开始>
// 为了性能优化，首先使用 !p.init.Val() 进行检查。
# <翻译结束>


<原文开始>
// If the file is removed or renamed, recreates the pool by increasing the pool id.
<原文结束>

# <翻译开始>
// 如果文件被删除或重命名，通过增加pool id重新创建pool。
# <翻译结束>


<原文开始>
// Clears the pool items staying in the pool.
<原文结束>

# <翻译开始>
// 清除池中留存的池项。
# <翻译结束>


<原文开始>
					// It uses another adding to drop the file items between the two adding.
					// Whenever the pool id changes, the pool will be recreated.
<原文结束>

# <翻译开始>
// 它利用另一个添加操作来丢弃两个添加之间的文件项。
// 每当池ID发生变化时，将会重新创建该池。
# <翻译结束>


<原文开始>
// Close closes current file pointer pool.
<原文结束>

# <翻译开始>
// Close 关闭当前文件指针池。
# <翻译结束>


<原文开始>
// It drops the old pool.
<原文结束>

# <翻译开始>
// 它会丢弃旧的连接池。
# <翻译结束>

