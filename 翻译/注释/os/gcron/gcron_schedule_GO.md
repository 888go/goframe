
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
// cronSchedule is the schedule for cron job.
<原文结束>

# <翻译开始>
// cronSchedule 是cron作业的计划安排。
# <翻译结束>


<原文开始>
// Created timestamp in seconds.
<原文结束>

# <翻译开始>
// 创建时间戳（以秒为单位）
# <翻译结束>


<原文开始>
// Running interval in seconds.
<原文结束>

# <翻译开始>
// 运行间隔时间（单位：秒）
# <翻译结束>


<原文开始>
// The raw cron pattern string.
<原文结束>

# <翻译开始>
// 原始的cron模式字符串。
# <翻译结束>


<原文开始>
// Job can run in these second numbers.
<原文结束>

# <翻译开始>
// Job可以在以下秒数运行。
# <翻译结束>


<原文开始>
// Job can run in these minute numbers.
<原文结束>

# <翻译开始>
// Job可以在以下分钟数运行。
# <翻译结束>


<原文开始>
// Job can run in these hour numbers.
<原文结束>

# <翻译开始>
// Job可以在以下小时数运行。
# <翻译结束>


<原文开始>
// Job can run in these day numbers.
<原文结束>

# <翻译开始>
// Job可以在这些天数中运行。
# <翻译结束>


<原文开始>
// Job can run in these week numbers.
<原文结束>

# <翻译开始>
// Job可以在这些周数中运行。
# <翻译结束>


<原文开始>
// Job can run in these moth numbers.
<原文结束>

# <翻译开始>
// Job 可以在以下月份数字中运行。
# <翻译结束>


<原文开始>
// Last timestamp number, for timestamp fix in some delay.
<原文结束>

# <翻译开始>
// 上次时间戳编号，用于修正某些延迟情况下的时间戳。
# <翻译结束>


<原文开始>
// regular expression for cron pattern, which contains 6 parts of time units.
<原文结束>

# <翻译开始>
// 正则表达式用于cron模式，该模式包含6部分时间单元。
# <翻译结束>


<原文开始>
// Short month name to its number.
<原文结束>

# <翻译开始>
// 短月份名称及其对应的数字。
# <翻译结束>


<原文开始>
// Full month name to its number.
<原文结束>

# <翻译开始>
// 完整的月份名称转为对应的数字。
# <翻译结束>


<原文开始>
// Short week name to its number.
<原文结束>

# <翻译开始>
// 短星期名称及其对应的数字。
# <翻译结束>


<原文开始>
// Full week name to its number.
<原文结束>

# <翻译开始>
// 完整周名称及其对应的周数。
# <翻译结束>


<原文开始>
// newSchedule creates and returns a schedule object for given cron pattern.
<原文结束>

# <翻译开始>
// newSchedule 根据给定的cron模式创建并返回一个schedule对象。
# <翻译结束>


<原文开始>
// Check if the predefined patterns.
<原文结束>

# <翻译开始>
// 检查预定义的模式是否存在
# <翻译结束>


<原文开始>
	// Handle the common cron pattern, like:
	// 0 0 0 1 1 2
<原文结束>

# <翻译开始>
// 处理常见的cron模式，例如：
// 0 0 0 1 1 2
// （注：该段代码注释省略了对cron模式的具体解释，以下是补充说明）
// 上述代码注释提到的"常见的cron模式"在Unix/Linux系统中用于表示定时任务的时间配置，
// 其格式为：分 时 天(月) 月 星期 周
// 示例 "0 0 0 1 1 2" 的含义是：
// 在每月的第一天（1号）的第一个星期二（2）的凌晨0点0分执行定时任务。
# <翻译结束>


<原文开始>
// parsePatternItem parses every item in the pattern and returns the result as map, which is used for indexing.
<原文结束>

# <翻译开始>
// parsePatternItem 解析模式中的每一项，并将结果以映射形式返回，该映射用于索引。
# <翻译结束>


<原文开始>
// parsePatternItemValue parses the field value to a number according to its field type.
<原文结束>

# <翻译开始>
// parsePatternItemValue 根据字段类型将字段值解析为数字。
# <翻译结束>


<原文开始>
		// Check if it contains letter,
		// it converts the value to number according to predefined map.
<原文结束>

# <翻译开始>
// 检查其中是否包含字母，
// 根据预定义的映射将其值转换为数字。
# <翻译结束>


<原文开始>
// checkMeetAndUpdateLastSeconds checks if the given time `t` meets the runnable point for the job.
<原文结束>

# <翻译开始>
// checkMeetAndUpdateLastSeconds 检查给定的时间 `t` 是否满足作业的可执行时间点，并更新最后执行秒数。
# <翻译结束>


<原文开始>
// It checks using normal cron pattern.
<原文结束>

# <翻译开始>
// 它使用标准cron模式进行检查。
# <翻译结束>


<原文开始>
// Next returns the next time this schedule is activated, greater than the given
// time.  If no time can be found to satisfy the schedule, return the zero time.
<原文结束>

# <翻译开始>
// Next 函数返回该计划下一次激活的时间，该时间大于给定的时间。
// 如果找不到满足计划要求的时间，则返回零时间（即时间的零值，表示无效时间）。
# <翻译结束>


<原文开始>
// Start at the earliest possible time (the upcoming second).
<原文结束>

# <翻译开始>
// 从最早可能的时间开始（即即将到来的下一秒）。
# <翻译结束>


<原文开始>
// who will care the job that run in five years later
<原文结束>

# <翻译开始>
// 谁会在意五年后运行的那份工作
# <翻译结束>


<原文开始>
		// Notice if the hour is no longer midnight due to DST.
		// Add an hour if it's 23, subtract an hour if it's 1.
<原文结束>

# <翻译开始>
// 注意由于DST（夏令时）导致的小时数是否不再为午夜。
// 如果是23点则加1小时，如果是1点则减1小时。
# <翻译结束>


<原文开始>
// dayMatches returns true if the schedule's day-of-week and day-of-month
// restrictions are satisfied by the given time.
<原文结束>

# <翻译开始>
// dayMatches 函数返回一个布尔值，如果给定时间满足该计划的周几和每月几日的限制条件，则返回true。
# <翻译结束>


<原文开始>
// Predefined pattern map.
<原文结束>

# <翻译开始>
// 预定义模式映射
# <翻译结束>


<原文开始>
// Like: 1-30, JAN-DEC
<原文结束>

# <翻译开始>
// 类似于：1-30，JAN-DEC
# <翻译结束>


<原文开始>
// It's checking week field.
<原文结束>

# <翻译开始>
// 正在检查周字段。
# <翻译结束>


<原文开始>
// It's checking month field.
<原文结束>

# <翻译开始>
// 正在检查月份字段。
# <翻译结束>


<原文开始>
// It is pure number.
<原文结束>

# <翻译开始>
// 这是一个纯数字。
# <翻译结束>


<原文开始>
// It checks using interval.
<原文结束>

# <翻译开始>
// 它使用间隔进行检查。
# <翻译结束>

