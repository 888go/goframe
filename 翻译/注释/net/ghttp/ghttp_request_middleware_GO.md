
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
// middleware is the plugin for request workflow management.
<原文结束>

# <翻译开始>
// middleware 是用于请求工作流程管理的插件。
# <翻译结束>


<原文开始>
// Is the request served, which is used for checking response status 404.
<原文结束>

# <翻译开始>
// Is the request served 是用于检查响应状态404的，即该请求是否已成功处理。
# <翻译结束>


<原文开始>
// The request object pointer.
<原文结束>

# <翻译开始>
// 请求对象指针。
# <翻译结束>


<原文开始>
// Index number for executing sequence purpose for handler items.
<原文结束>

# <翻译开始>
// Index 号用于执行顺序目的，针对处理项。
# <翻译结束>


<原文开始>
// Index number for executing sequence purpose for bound middleware of handler item.
<原文结束>

# <翻译开始>
// Index 数值用于执行顺序的目的，用于处理器项绑定的中间件。
# <翻译结束>


<原文开始>
// Next calls the next workflow handler.
// It's an important function controlling the workflow of the server request execution.
<原文结束>

# <翻译开始>
// Next调用下一个工作流处理器。
// 这是一个重要的函数，用于控制服务器请求执行的工作流程。
# <翻译结束>


<原文开始>
// Check whether the request is excited.
<原文结束>

# <翻译开始>
// 检查请求是否已激发。
# <翻译结束>


<原文开始>
// Filter the HOOK handlers, which are designed to be called in another standalone procedure.
<原文结束>

# <翻译开始>
// 过滤HOOK处理程序，它们设计为在另一个独立的进程中被调用。
# <翻译结束>












<原文开始>
// Execute bound middleware array of the item if it's not empty.
<原文结束>

# <翻译开始>
// 如果item的绑定中间件数组不为空，则执行该数组中的中间件。
# <翻译结束>

















<原文开始>
				// It does not continue calling next middleware after another middleware done.
				// There should be a "Next" function to be called in the middleware in order to manage the workflow.
<原文结束>

# <翻译开始>
// 当某个中间件执行完毕后，它不会继续调用下一个中间件。
// 若要管理工作流程，应在中间件中调用“Next”函数。
# <翻译结束>


<原文开始>
// It's already an error that has stack info.
<原文结束>

# <翻译开始>
// 这已经是一个带有堆栈信息的错误。
# <翻译结束>


<原文开始>
				// Create a new error with stack info.
				// Note that there's a skip pointing the start stacktrace
				// of the real error point.
<原文结束>

# <翻译开始>
// 创建一个包含堆栈信息的新错误。
// 注意，这里有一个skip参数用于指向实际错误点的堆栈跟踪起始位置。
# <翻译结束>


<原文开始>
// Check the http status code after all handlers and middleware done.
<原文结束>

# <翻译开始>
// 在所有处理器和中间件执行完毕后，检查HTTP状态码。
# <翻译结束>


<原文开始>
// Current router switching.
<原文结束>

# <翻译开始>
// 当前路由器切换中。
# <翻译结束>


<原文开始>
// Router values switching.
<原文结束>

# <翻译开始>
// 路由值切换。
# <翻译结束>


<原文开始>
// Global middleware array.
<原文结束>

# <翻译开始>
// 全局中间件数组。
# <翻译结束>

