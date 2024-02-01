
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
// Done forbids the context done from parent context.
<原文结束>

# <翻译开始>
// Done 从父上下文中禁止使用已完成的上下文。
# <翻译结束>


<原文开始>
// Deadline forbids the context deadline from parent context.
<原文结束>

# <翻译开始>
// Deadline 禁止从父上下文中继承截止时间（deadline），即设置当前上下文不具有来自父上下文的截止时间。
# <翻译结束>


<原文开始>
// Err forbids the context done from parent context.
<原文结束>

# <翻译开始>
// Err用于禁止从父上下文中获取完成信号。
# <翻译结束>


<原文开始>
// NeverDone wraps and returns a new context object that will be never done,
// which forbids the context manually done, to make the context can be propagated to asynchronous goroutines.
<原文结束>

# <翻译开始>
// NeverDone 包装并返回一个新的上下文对象，该对象将永不完成，
// 这样可以禁止手动完成上下文，以使上下文能够传播到异步goroutine中。
# <翻译结束>


<原文开始>
// neverDoneCtx never done.
<原文结束>

# <翻译开始>
// neverDoneCtx 永不结束
# <翻译结束>

