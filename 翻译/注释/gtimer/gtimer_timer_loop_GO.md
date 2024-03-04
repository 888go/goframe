
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
// loop starts the ticker using a standalone goroutine.
<原文结束>

# <翻译开始>
// loop 使用独立的 goroutine 启动ticker。
# <翻译结束>












<原文开始>
// proceed function proceeds the timer job checking and running logic.
<原文结束>

# <翻译开始>
// proceed 函数执行定时任务的检查和运行逻辑。
# <翻译结束>


<原文开始>
// It checks if it meets the ticks' requirement.
<原文结束>

# <翻译开始>
// 它检查是否满足滴答（ticks）的要求。
# <翻译结束>


<原文开始>
// It pushes the job back if current ticks does not meet its running ticks requirement.
<原文结束>

# <翻译开始>
// 如果当前的ticks数未达到其运行所需的ticks要求，则将该任务重新推回。
# <翻译结束>


<原文开始>
// It checks the job running requirements and then does asynchronous running.
<原文结束>

# <翻译开始>
// 它检查作业运行需求，然后进行异步运行。
# <翻译结束>


<原文开始>
// Status check: push back or ignore it.
<原文结束>

# <翻译开始>
// 状态检查：将其推回或忽略。
# <翻译结束>


<原文开始>
// It pushes the job back to queue for next running.
<原文结束>

# <翻译开始>
// 它将任务推回到队列中以便下次运行。
# <翻译结束>


<原文开始>
// Check the timer status.
<原文结束>

# <翻译开始>
// 检查定时器状态。
# <翻译结束>

