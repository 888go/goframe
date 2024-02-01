
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
		//t.Assert(gstr.Count(content, ".go"), 1)
		//t.Assert(gstr.Contains(content, gfile.Separator), true)
<原文结束>

# <翻译开始>
// t.Assert(gstr.Count(content, ".go"), 1)
// 检查content字符串中".go"出现的次数是否为1，将结果断言为真
// t.Assert(gstr.Contains(content, gfile.Separator), true)
// 断言content字符串中是否包含系统文件分隔符（gfile.Separator），结果应为真
# <翻译结束>


<原文开始>
		//t.Assert(gstr.Count(content, ".go"), 1)
		//t.Assert(gstr.Contains(content, gfile.Separator), false)
<原文结束>

# <翻译开始>
// t.Assert(gstr.Count(content, ".go"), 1) // 断言content字符串中".go"子串出现的次数为1次
// t.Assert(gstr.Contains(content, gfile.Separator), false) // 断言content字符串中不包含gfile.Separator定义的分隔符
# <翻译结束>


<原文开始>
//t.Assert(gstr.Count(content, "Stack"), 1)
<原文结束>

# <翻译开始>
// t.Assert(gstr.Count(content, "Stack"), 1) // 翻译：// t断言content中"Stack"出现的次数为1次
# <翻译结束>


<原文开始>
//t.Assert(gstr.Count(content, "Stack"), 0)
<原文结束>

# <翻译开始>
// t.Assert(gstr.Count(content, "Stack"), 0) // 翻译：// t断言content中"Stack"出现的次数为0
# <翻译结束>

