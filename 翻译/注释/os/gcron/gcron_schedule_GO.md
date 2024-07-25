
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
// cronSchedule is the schedule for cron job.
<原文结束>

# <翻译开始>
// cronSchedule 是定时任务的调度计划。 md5:4731e43288725f27
# <翻译结束>


<原文开始>
// Created timestamp in seconds.
<原文结束>

# <翻译开始>
// 创建时间的时间戳，以秒为单位。 md5:4a0001cda2177f41
# <翻译结束>


<原文开始>
// Running interval in seconds.
<原文结束>

# <翻译开始>
// 运行间隔（以秒为单位）。 md5:a62fd57ffa9e26f4
# <翻译结束>


<原文开始>
// The raw cron pattern string that is passed in cron job creation.
<原文结束>

# <翻译开始>
// 在创建cron作业时传递的原始cron模式字符串。 md5:18b07692590ddf66
# <翻译结束>


<原文开始>
// Mark the pattern is standard 5 parts crontab pattern instead 6 parts pattern.
<原文结束>

# <翻译开始>
// 标记该模式是标准的5部分cron表达式模式，而不是6部分模式。 md5:89774325ba9632d2
# <翻译结束>


<原文开始>
// Job can run in these second numbers.
<原文结束>

# <翻译开始>
// 该Job可以在这些秒数内运行。 md5:603e4f208dcc04bf
# <翻译结束>


<原文开始>
// Job can run in these minute numbers.
<原文结束>

# <翻译开始>
// Job可以在这些分钟数运行。 md5:9cc64d9456bc318a
# <翻译结束>


<原文开始>
// Job can run in these hour numbers.
<原文结束>

# <翻译开始>
// Job可以在这些小时数运行。 md5:cf1a7bc2b7ada427
# <翻译结束>


<原文开始>
// Job can run in these day numbers.
<原文结束>

# <翻译开始>
// Job 可以在这些天数中运行。 md5:9be6d3ae1549f6c8
# <翻译结束>


<原文开始>
// Job can run in these week numbers.
<原文结束>

# <翻译开始>
// Job 可以在这些星期数中运行。 md5:e9d2ed887e372b17
# <翻译结束>


<原文开始>
// Job can run in these moth numbers.
<原文结束>

# <翻译开始>
// Job可以在这些月份运行。 md5:e58af4ea6da7e868
# <翻译结束>


<原文开始>
// This field stores the timestamp that meets schedule latest.
<原文结束>

# <翻译开始>
	// 这个字段存储满足计划的最新时间戳。 md5:df6f9fc73fbf03d6
# <翻译结束>


<原文开始>
// Last timestamp number, for timestamp fix in some latency.
<原文结束>

# <翻译开始>
	// 最后一个时间戳编号，用于在某些延迟情况下固定时间戳。 md5:6839316ecd982e4b
# <翻译结束>


<原文开始>
// regular expression for cron pattern, which contains 6 parts of time units.
<原文结束>

# <翻译开始>
	// 正则表达式表示的cron模式，包含6个时间单位部分。 md5:75e472ef39ca5aab
# <翻译结束>


<原文开始>
// Predefined pattern map.
<原文结束>

# <翻译开始>
	// 预定义的模式映射。 md5:dc23a289b509e3b6
# <翻译结束>


<原文开始>
// Short month name to its number.
<原文结束>

# <翻译开始>
	// 短月名到其对应的数字。 md5:44f6938b62580af0
# <翻译结束>


<原文开始>
// Full month name to its number.
<原文结束>

# <翻译开始>
	// 完整的月份名称转换为其对应的数字。 md5:e9b9f99b1f2191d0
# <翻译结束>


<原文开始>
// Short week name to its number.
<原文结束>

# <翻译开始>
	// 短星期名转换为对应的数字。 md5:c8dde2776e296b0a
# <翻译结束>


<原文开始>
// Full week name to its number.
<原文结束>

# <翻译开始>
	// 完整的星期名称到其数字。 md5:05d1a360fc5b25ee
# <翻译结束>


<原文开始>
// newSchedule creates and returns a schedule object for given cron pattern.
<原文结束>

# <翻译开始>
// newSchedule根据给定的cron模式创建并返回一个调度对象。 md5:14dff188c64f1e56
# <翻译结束>


<原文开始>
// Check given `pattern` if the predefined patterns.
<原文结束>

# <翻译开始>
	// 检查给定的`pattern`是否在预定义的模式中。 md5:31badfbc0ed60d2b
# <翻译结束>


<原文开始>
// Handle given `pattern` as common 6 parts pattern.
<原文结束>

# <翻译开始>
	// 处理给定的`pattern`作为常见的6部分模式。 md5:224ce220d8873fe0
# <翻译结束>


<原文开始>
// parsePatternItem parses every item in the pattern and returns the result as map, which is used for indexing.
<原文结束>

# <翻译开始>
// parsePatternItem 解析模式中的每个项目，并将结果作为映射返回，该映射用于索引。 md5:66716855d8c0f694
# <翻译结束>


<原文开始>
// Example: 1-10/2,11-30/3
<原文结束>

# <翻译开始>
	// 这个注释表示一个范围的分组示例。"1-10/2" 表示从1开始到10，每2个数一组；"11-30/3" 表示从11开始到30，每3个数一组。 md5:7074496c7eb487df
# <翻译结束>


<原文开始>
// parseWeekAndMonthNameToInt parses the field value to a number according to its field type.
<原文结束>

# <翻译开始>
// parseWeekAndMonthNameToInt 根据字段类型将字段值解析为数字。 md5:10e98c83dca57c49
# <翻译结束>


<原文开始>
		// Check if it contains letter,
		// it converts the value to number according to predefined map.
<原文结束>

# <翻译开始>
		// 检查是否包含字母，
		// 根据预定义的映射将值转换为数字。
		// md5:d6cf713cc1230de9
# <翻译结束>

