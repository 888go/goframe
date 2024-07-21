
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
// middleware is the plugin for request workflow management.
<原文结束>

# <翻译开始>
// middleware是用于请求工作流程管理的插件。 md5:49f6a68042876e4b
# <翻译结束>


<原文开始>
// Is the request served, which is used for checking response status 404.
<原文结束>

# <翻译开始>
// 请求是否已处理，用于检查响应状态码404。 md5:d6a59089bf4bca74
# <翻译结束>


<原文开始>
// The request object pointer.
<原文结束>

# <翻译开始>
// 请求对象的指针。 md5:d2b2b51d27020540
# <翻译结束>


<原文开始>
// Index number for executing sequence purpose for handler items.
<原文结束>

# <翻译开始>
// 用于处理项执行顺序的索引号。 md5:98c1f7684660f958
# <翻译结束>


<原文开始>
// Index number for executing sequence purpose for bound middleware of handler item.
<原文结束>

# <翻译开始>
// 用于处理项绑定中间件的执行顺序编号。 md5:ea6ccd0cbc909909
# <翻译结束>


<原文开始>
// Next calls the next workflow handler.
// It's an important function controlling the workflow of the server request execution.
<原文结束>

# <翻译开始>
// Next 调用下一个工作流处理器。
// 这是一个重要的函数，用于控制服务器请求执行的工作流程。
// md5:9993825368a59675
# <翻译结束>


<原文开始>
// Check whether the request is excited.
<原文结束>

# <翻译开始>
// 检查请求是否已激活。 md5:87f631160593048d
# <翻译结束>


<原文开始>
// Filter the HOOK handlers, which are designed to be called in another standalone procedure.
<原文结束>

# <翻译开始>
// 过滤HOOK处理器，这些处理器设计用于在另一个独立的程序中被调用。 md5:f285ad394cb72a16
# <翻译结束>


<原文开始>
// Current router switching.
<原文结束>

# <翻译开始>
// 当前路由器切换。 md5:6e5ea35720c091cd
# <翻译结束>


<原文开始>
// Router values switching.
<原文结束>

# <翻译开始>
// 路由器值切换。 md5:8a03c212d83f07ec
# <翻译结束>


<原文开始>
// Execute bound middleware array of the item if it's not empty.
<原文结束>

# <翻译开始>
// 如果项的绑定中间件数组不为空，则执行它。 md5:c2326e1bb33e0423
# <翻译结束>


<原文开始>
// Global middleware array.
<原文结束>

# <翻译开始>
// 全局中间件数组。 md5:1b4420db0d912753
# <翻译结束>


<原文开始>
				// It does not continue calling next middleware after another middleware done.
				// There should be a "Next" function to be called in the middleware in order to manage the workflow.
<原文结束>

# <翻译开始>
				// 它不会在其他中间件完成之后继续调用下一个中间件。为了管理工作流程，中间件应该有一个名为 "Next" 的函数可供调用。
				// md5:0a1a7642101f1bb9
# <翻译结束>


<原文开始>
// It's already an error that has stack info.
<原文结束>

# <翻译开始>
// 它已经是一个带有堆栈信息的错误。 md5:ec045ebe21bca18d
# <翻译结束>


<原文开始>
				// Create a new error with stack info.
				// Note that there's a skip pointing the start stacktrace
				// of the real error point.
<原文结束>

# <翻译开始>
				// 创建一个带有堆栈信息的新错误。
				// 注意，skip 参数指定了从哪个调用栈开始追踪真正的错误点。
				// md5:e23da1f0a4a0c90f
# <翻译结束>


<原文开始>
// Check the http status code after all handlers and middleware done.
<原文结束>

# <翻译开始>
// 在所有处理程序和中间件完成后检查HTTP状态码。 md5:6db5e4718ab69458
# <翻译结束>

