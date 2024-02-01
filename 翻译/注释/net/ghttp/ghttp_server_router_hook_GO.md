
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
// BindHookHandler registers handler for specified hook.
<原文结束>

# <翻译开始>
// BindHookHandler 为指定的钩子注册处理器。
# <翻译结束>


<原文开始>
// doBindHookHandlerInput is the input for BindHookHandler.
<原文结束>

# <翻译开始>
// doBindHookHandlerInput 是 BindHookHandler 的输入参数。
# <翻译结束>


<原文开始>
// doBindHookHandler is the internal handler for BindHookHandler.
<原文结束>

# <翻译开始>
// doBindHookHandler 是 BindHookHandler 的内部处理程序。
# <翻译结束>


<原文开始>
// BindHookHandlerByMap registers handler for specified hook.
<原文结束>

# <翻译开始>
// BindHookHandlerByMap 为指定的钩子注册处理器。
# <翻译结束>


<原文开始>
// callHookHandler calls the hook handler by their registered sequences.
<原文结束>

# <翻译开始>
// callHookHandler 按照已注册的顺序调用钩子处理器
# <翻译结束>


<原文开始>
// Backup the old router variable map.
<原文结束>

# <翻译开始>
// 备份旧的路由器变量映射。
# <翻译结束>


<原文开始>
			// DO NOT USE the router of the hook handler,
			// which can overwrite the router of serving handler.
			// r.Router = item.handler.router
<原文结束>

# <翻译开始>
// **不要在钩子处理器中使用路由器**，
// 这可能会覆盖服务处理器的路由器。
// r.Router = item.handler.router
# <翻译结束>


<原文开始>
// Restore the old router variable map.
<原文结束>

# <翻译开始>
// 恢复旧的路由器变量映射。
# <翻译结束>


<原文开始>
// getHookHandlers retrieves and returns the hook handlers of specified hook.
<原文结束>

# <翻译开始>
// getHookHandlers 获取并返回指定钩子的钩子处理程序。
# <翻译结束>


<原文开始>
// niceCallHookHandler nicely calls the hook handler function,
// which means it automatically catches and returns the possible panic error to
// avoid goroutine crash.
<原文结束>

# <翻译开始>
// niceCallHookHandler 优雅地调用钩子处理器函数，
// 这意味着它会自动捕获并返回可能的 panic 错误，以避免 goroutine 中止。
# <翻译结束>

