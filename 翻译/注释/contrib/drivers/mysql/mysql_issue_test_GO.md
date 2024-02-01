
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
// https://github.com/gogf/gf/issues/1934
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个问题编号1934。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题1934。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1570
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库下的第1570号issue。
// 翻译成中文：
// 引用了GitHub上gogf/gf项目中的第1570号问题。
# <翻译结束>







<原文开始>
// Result ScanList with struct elements and pointer attributes.
<原文结束>

# <翻译开始>
// Result ScanList，用于包含结构体元素和指针属性。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1401
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库的第1401个issue（问题）。
// 中文翻译：
// 引用了GitHub上gogf/gf项目的问题1401。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1412
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf仓库的第1412号问题链接
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1002
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容为一个GitHub仓库的issue链接。
// 翻译为：
// 参考GitHub上gogf/gf项目中的第1002号问题。
# <翻译结束>







<原文开始>
// where + gtime.Time arguments.
<原文结束>

# <翻译开始>
// where + gtime.Time 参数
# <翻译结束>


<原文开始>
// where + time.Time arguments, UTC.
<原文结束>

# <翻译开始>
// 函数参数中包含 where + time.Time 类型，时间使用 UTC（协调世界时）。
# <翻译结束>


<原文开始>
	// where + time.Time arguments, +8.
	// gtest.C(t, func(t *gtest.T) {
	//	// Change current timezone to +8 zone.
	//	location, err := time.LoadLocation("Asia/Shanghai")
	//	t.AssertNil(err)
	//	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-27 19:03:32", location)
	//	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-27 19:03:34", location)
	//	{
	//		v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).Value()
	//		t.AssertNil(err)
	//		t.Assert(v.Int(), 1)
	//	}
	//	{
	//		v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).FindValue()
	//		t.AssertNil(err)
	//		t.Assert(v.Int(), 1)
	//	}
	//	{
	//		v, err := db.Model(table).Where("create_time>? and create_time<?", t1, t2).FindValue("id")
	//		t.AssertNil(err)
	//		t.Assert(v.Int(), 1)
	//	}
	// })
<原文结束>

# <翻译开始>
// 此处使用了+8时区的时间参数。
// gtest.C(t, func(t *gtest.T) {
//// 将当前时区更改为+8时区（即中国北京时间）。
// location, err := time.LoadLocation("Asia/Shanghai")
// t.AssertNil(err) // 断言加载时区无错误
// t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-27 19:03:32", location) // 解析字符串为指定时区的time.Time类型
// t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-27 19:03:34", location)
// 
// // 使用create_time字段在t1和t2时间范围内的查询条件，获取id字段值
// {
//     v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).Value()
//     t.AssertNil(err) // 断言查询过程无错误
//     t.Assert(v.Int(), 1) // 断言查询结果转换为整型后为1
// }
// 
// // 使用create_time字段在t1和t2时间范围内的查询条件，通过FindValue方法获取id字段值
// {
//     v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).FindValue()
//     t.AssertNil(err) // 断言查询过程无错误
//     t.Assert(v.Int(), 1) // 断言查询结果转换为整型后为1
// }
// 
// // 使用create_time字段在t1和t2时间范围内的查询条件，通过FindValue方法并指定"id"字段获取值
// {
//     v, err := db.Model(table).Where("create_time>? and create_time<?", t1, t2).FindValue("id")
//     t.AssertNil(err) // 断言查询过程无错误
//     t.Assert(v.Int(), 1) // 断言查询结果转换为整型后为1
// }
// })
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1700
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题 issue #1700。
// 翻译为：
// 参考GitHub上gogf/gf项目的第1700号问题。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1701
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个问题 issue 1701。
// 中文翻译：
// 这是引用了GitHub上gogf/gf项目第1701号问题的链接。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1733
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个问题编号1733。
// 中文翻译：
// 这是引用了GitHub上gogf/gf项目第1733号问题的链接。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2105
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题编号2105。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题#2105
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2231
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库的第2231号议题。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题#2231
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2339
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个问题（issue）：#2339。
// 翻译成中文：
// 这是Go语言代码中的一个注释，它提供了一个链接地址：https://github.com/gogf/gf/issues/2339，该链接指向GitHub上gogf/gf项目的问题编号2339。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2356
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf仓库的第2356号issue
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2338
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf仓库的第2338个issue
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2427
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库下的第2427个issue（问题）。
// 翻译成中文：
// 引用了GitHub上gogf/gf项目中的第2427个问题。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2561
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf仓库的第2561个issue链接
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2439
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf项目的一个问题链接，编号为2439。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2782
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容为一个URL链接，指向GitHub上gogf/gf项目的一个issue（问题）讨论页面，编号为2782。
// 中文翻译：
// 这是Go语言代码中的一个注释，它提供了一个链接至GitHub上gogf/gf项目第2782号问题的讨论页面。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2907
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf仓库的第2907号问题链接
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/3086
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf仓库的第3086个issue链接
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/3204
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题链接：https://github.com/gogf/gf/issues/3204
// 翻译成中文：
// 这指向了GitHub上gogf/gf项目的一个问题，编号为3204
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/3218
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf仓库的第3218个issue
# <翻译结束>


<原文开始>
// Initialize the data.
<原文结束>

# <翻译开始>
// 初始化数据
# <翻译结束>


<原文开始>
// where + string arguments.
<原文结束>

# <翻译开始>
// where + 字符串参数。
# <翻译结束>

