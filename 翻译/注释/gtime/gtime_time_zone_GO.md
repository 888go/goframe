
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
// SetTimeZone sets the time zone for current whole process.
// The parameter `zone` is an area string specifying corresponding time zone,
// eg: Asia/Shanghai.
//
// PLEASE VERY NOTE THAT:
// 1. This should be called before package "time" import.
// 2. This function should be called once.
// 3. Please refer to issue: https://github.com/golang/go/issues/34814
<原文结束>

# <翻译开始>
// SetTimeZone 设置当前整个进程的时区。
// 参数 `zone` 是一个区域字符串，用于指定相应时区，例如：Asia/Shanghai。
//
// **非常重要，请注意**：
// 1. 此函数应在导入 "time" 包之前调用。
// 2. 此函数应仅调用一次。
// 3. 参考相关问题：https://github.com/golang/go/issues/34814
# <翻译结束>


<原文开始>
// It is already set to time.Local.
<原文结束>

# <翻译开始>
// 它已经被设置为 time.Local。
# <翻译结束>


<原文开始>
// Load zone info from specified name.
<原文结束>

# <翻译开始>
// 从指定名称加载时区信息。
# <翻译结束>


<原文开始>
// Update the time.Local for once.
<原文结束>

# <翻译开始>
// 更新一次 time.Local。
# <翻译结束>


<原文开始>
// Update the timezone environment for *nix systems.
<原文结束>

# <翻译开始>
// 更新*nix系统的时间zone环境变量。
# <翻译结束>


<原文开始>
// ToLocation converts current time to specified location.
<原文结束>

# <翻译开始>
// ToLocation将当前时间转换为指定时区的时间。
# <翻译结束>


<原文开始>
// ToZone converts current time to specified zone like: Asia/Shanghai.
<原文结束>

# <翻译开始>
// ToZone 将当前时间转换为指定时区，如：Asia/Shanghai。
# <翻译结束>


<原文开始>
// Local converts the time to local timezone.
<原文结束>

# <翻译开始>
// Local将时间转换为本地时区。
# <翻译结束>

