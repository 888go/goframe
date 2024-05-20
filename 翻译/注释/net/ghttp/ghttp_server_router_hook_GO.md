
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
// BindHookHandler registers handler for specified hook.
<原文结束>

# <翻译开始>
// BindHookHandler 为指定的钩子注册处理程序。. md5:325d65ceb75f1b33
# <翻译结束>


<原文开始>
// doBindHookHandlerInput is the input for BindHookHandler.
<原文结束>

# <翻译开始>
// doBindHookHandlerInput是BindHookHandler的输入参数。. md5:1a2c964079a2227c
# <翻译结束>


<原文开始>
// doBindHookHandler is the internal handler for BindHookHandler.
<原文结束>

# <翻译开始>
// doBindHookHandler是BindHookHandler的内部处理程序。. md5:5393ffad5084d597
# <翻译结束>


<原文开始>
// BindHookHandlerByMap registers handler for specified hook.
<原文结束>

# <翻译开始>
// BindHookHandlerByMap 为指定的钩子注册处理器。. md5:38d4b79317ac1b3f
# <翻译结束>


<原文开始>
// callHookHandler calls the hook handler by their registered sequences.
<原文结束>

# <翻译开始>
// callHookHandler 按照注册的顺序调用钩子处理器。. md5:4e1a8b2998b73ddb
# <翻译结束>


<原文开始>
// Backup the old router variable map.
<原文结束>

# <翻译开始>
// 备份旧的路由器变量映射。. md5:4a1427ee4ccef0a6
# <翻译结束>


<原文开始>
			// DO NOT USE the router of the hook handler,
			// which can overwrite the router of serving handler.
			// r.Router = item.handler.router
<原文结束>

# <翻译开始>
// 不要使用钩子处理器的路由器，
// 因为它可能会覆盖服务处理器的路由器。
// r.Router = item.handler.router
// md5:9c797403c522d44d
# <翻译结束>


<原文开始>
// Restore the old router variable map.
<原文结束>

# <翻译开始>
// 恢复旧的路由器变量映射。. md5:6ae23d30567bb237
# <翻译结束>


<原文开始>
// getHookHandlers retrieves and returns the hook handlers of specified hook.
<原文结束>

# <翻译开始>
// getHookHandlers 获取并返回指定钩子的处理函数。. md5:f19f77b15aa76d7a
# <翻译结束>


<原文开始>
// niceCallHookHandler nicely calls the hook handler function,
// which means it automatically catches and returns the possible panic error to
// avoid goroutine crash.
<原文结束>

# <翻译开始>
// niceCallHookHandler 美好地调用钩子处理函数，
// 即它会自动捕获并返回可能的恐慌错误，以防止goroutine崩溃。
// md5:915bcff9c5f9cc4e
# <翻译结束>

