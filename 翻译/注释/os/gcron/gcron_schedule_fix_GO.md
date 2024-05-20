
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
// getAndUpdateLastCheckTimestamp checks fixes and returns the last timestamp that have delay fix in some seconds.
<原文结束>

# <翻译开始>
// getAndUpdateLastCheckTimestamp 检查、修复并返回在几秒钟内有延迟修复的最后时间戳。. md5:617d53ed6d0eee3f
# <翻译结束>


<原文开始>
	// Often happens, timer triggers in the same second, but the millisecond is different.
	// Example:
	// lastCheckTimestamp: 2024-03-26 19:47:34.000
	// currentTimestamp:   2024-03-26 19:47:34.999
<原文结束>

# <翻译开始>
// 通常情况下，定时器在同一秒内触发，但毫秒数不同。
// 例如：
// lastCheckTimestamp: 2024-03-26 19:47:34.000
// currentTimestamp:   2024-03-26 19:47:34.999
// md5:7ad3ec347d1a6583
# <翻译结束>


<原文开始>
	// Often happens, no latency.
	// Example:
	// lastCheckTimestamp: 2024-03-26 19:47:34.000
	// currentTimestamp:   2024-03-26 19:47:35.000
<原文结束>

# <翻译开始>
// 经常发生的情况，没有延迟。
// 示例：
// lastCheckTimestamp: 2024年03月26日 19时47分34秒.000
// currentTimestamp:   2024年03月26日 19时47分35秒.000
// md5:1ed300ef7b928611
# <翻译结束>


<原文开始>
	// Latency in 3 seconds, which can be tolerant.
	// Example:
	// lastCheckTimestamp: 2024-03-26 19:47:31.000、2024-03-26 19:47:32.000
	// currentTimestamp:   2024-03-26 19:47:34.000
<原文结束>

# <翻译开始>
// 可容忍的延迟时间为3秒。
// 例如：
// lastCheckTimestamp: 2024-03-26 19:47:31.000、2024-03-26 19:47:32.000
// currentTimestamp:   2024-03-26 19:47:34.000
// md5:21934a048bbfddaf
# <翻译结束>


<原文开始>
// Too much latency, it ignores the fix, the cron job might not be triggered.
<原文结束>

# <翻译开始>
// 延迟太多，它忽略了修复，定时任务可能不会被触发。. md5:5d550b27269fafbf
# <翻译结束>


<原文开始>
// Too much delay, let's update the last timestamp to current one.
<原文结束>

# <翻译开始>
// 延迟时间过长，让我们将最后的timestamp更新为当前时间。. md5:7b051dda466c96cf
# <翻译结束>

