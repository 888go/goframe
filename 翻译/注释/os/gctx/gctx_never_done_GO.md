
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
// neverDoneCtx never done.
<原文结束>

# <翻译开始>
// neverDoneCtx 永远不会完成。 md5:9cd0926cf01acafc
# <翻译结束>


<原文开始>
// Done forbids the context done from parent context.
<原文结束>

# <翻译开始>
// Done 禁止从父上下文中关闭 done。 md5:6ee3971853766306
# <翻译结束>


<原文开始>
// Deadline forbids the context deadline from parent context.
<原文结束>

# <翻译开始>
// Deadline 禁止从父上下文中继承截止期限。 md5:b0a8033fcfdd0483
# <翻译结束>


<原文开始>
// Err forbids the context done from parent context.
<原文结束>

# <翻译开始>
// Err 限制了从父上下文中完成的上下文。 md5:605f4a25a7f54817
# <翻译结束>


<原文开始>
// NeverDone wraps and returns a new context object that will be never done,
// which forbids the context manually done, to make the context can be propagated
// to asynchronous goroutines.
//
// Note that, it does not affect the closing (canceling) of the parent context,
// as it is a wrapper for its parent, which only affects the next context handling.
<原文结束>

# <翻译开始>
// NeverDone 包装并返回一个永远不会完成的新上下文对象，这禁止手动完成上下文，使得上下文可以传递给异步的 Goroutine。
//
// 请注意，这不会影响父上下文的关闭（取消），因为它只是父上下文的包装器，只影响下一个上下文处理。
// md5:38b63a322c5449a9
# <翻译结束>

