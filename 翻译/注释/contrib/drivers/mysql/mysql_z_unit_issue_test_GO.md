
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1934
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/1934. md5:96f55929c7ed56a0
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1570
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/1570. md5:37966850af641bcc
# <翻译结束>


<原文开始>
// Result ScanList with struct elements and pointer attributes.
<原文结束>

# <翻译开始>
	// Result 使用具有结构体元素和指针属性的ScanList。 md5:b23d106d13859ad5
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1401
<原文结束>

# <翻译开始>
// 这段注释引用的是一个GitHub问题（Issue）的链接，来自gogf（一个Go语言的框架）项目。它表示这是对问题1401的讨论或参考。在中文中，这可能表示：“参见GitHub上的gogf项目问题1401”。 md5:5d32589f093beb22
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1412
<原文结束>

# <翻译开始>
// 这段注释引用的是GitHub上的一个issue（问题或讨论），gf（GoGF）是一个用Go语言编写的Web框架。 issue号1412可能是指该框架中某个特定的问题或者提出的改进请求，具体内容需要查看相关issue的详细描述。 md5:c6f20cc497b1e9a6
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1002
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/1002 问题讨论. md5:2a97dfd9cd049763
# <翻译结束>


<原文开始>
// where + string arguments.
<原文结束>

# <翻译开始>
	// where + 字符串参数。 md5:cb1db92222691d4d
# <翻译结束>


<原文开始>
// where + gtime.Time arguments.
<原文结束>

# <翻译开始>
	// 其中包含 gtime.Time 类型的参数。 md5:3bd9bb993dd2cc53
# <翻译结束>


<原文开始>
// where + time.Time arguments, UTC.
<原文结束>

# <翻译开始>
	// 带有时间.Time参数，使用UTC时区。 md5:80f36eaa256e894c
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
	// 在时间.Time参数中，+8代表时区偏移。
	// gtest.C(t, func(t *gtest.T) {
	//     	// 将当前时区更改为+8时区（东八区）。
	//     location, err := time.LoadLocation("Asia/Shanghai")
	//     t.AssertNil(err) 	// 确认加载时区无错误。
	//     t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-27 19:03:32", location)
	//     t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-27 19:03:34", location)
	//     	// 使用定义的时间段进行查询测试：
	//     {
	//         v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).Value()
	//         t.AssertNil(err) 	// 确认查询无错误。
	//         t.Assert(v.Int(), 1) 	// 断言查询结果的ID为1。
	//     }
	//     {
	//         v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).FindValue()
	//         t.AssertNil(err) 	// 同上，确认查询无错误。
	//         t.Assert(v.Int(), 1) 	// 断言查询结果的ID为1。
	//     }
	//     {
	//         v, err := db.Model(table).Where("create_time>? and create_time<?", t1, t2).FindValue("id")
	//         t.AssertNil(err) 	// 再次确认查询无错误。
	//         t.Assert(v.Int(), 1) 	// 继续断言查询结果的ID为1。
	//     }
	// }) md5:766797023d98820e
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1700
<原文结束>

# <翻译开始>
// 这段注释是引用了GitHub上gf框架的一个问题链接，编号为1700。在Go代码中，这种注释通常用于指向相关的讨论、问题或者需求，以便其他开发者了解代码的背景或上下文。翻译成中文后，其含义不变：. md5:a352b9ef5236ff28
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1701
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/1701。GF（Go Foundation）可能是Go语言的一个项目或者库，而"1701"可能是问题的编号。这个注释可能是在讨论或记录与GF项目相关的问题1701的情况。 md5:cc9c86ac60eeaf58
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1733
<原文结束>

# <翻译开始>
// 这段注释链接指向的是GitHub上的一个Issue（问题报告）页面，来自gogf/gf（一个Go语言的框架）项目。具体来说，它可能是指1733号问题或者与该问题相关的内容。在中文中，这通常表示对某个问题、讨论或改进的引用。 md5:76faec7f21ba3b13
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2105
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/2105. md5:579ab324e61be1fb
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2231
<原文结束>

# <翻译开始>
// 这段注释是链接到一个GitHub问题的引用，该问题是关于gf（GoFrame）框架的一个问题或讨论。在GitHub仓库gf的 issues 页面中，编号2231的问题提供了更多的上下文和信息。由于注释本身没有详细内容，所以具体的翻译就是保持原样，表示这是一个与gf框架相关的问题链接。 md5:803083b8650008ce
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2339
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/2339。"gf"可能是某个Go语言的库（golang的gopher框架）的简称，"issues/2339"表示该仓库中编号为2339的问题或者issue。这可能是一个开发者社区中关于gf库的报告、提问或者反馈。 md5:fb506ddf20da598c
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2356
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/2356。"gf"可能是某个项目的代号，"gogf"可能是一个开发者的用户名，"issues/2356"表示该问题是编号为2356的 issue（通常是开发者社区中报告的问题、建议或讨论）。 md5:a688eda9a4ec7d89
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2338
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/2338. md5:a504f30db0e1a70a
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2427
<原文结束>

# <翻译开始>
// 这段注释是引用了GitHub上一个名为gf的项目中的问题编号2427。在Go语言中，这种注释通常用于指向相关问题、讨论或需求的链接，以便其他开发人员了解代码背景或跟踪问题。翻译成中文后，它依然保持原样，因为这是一个网址引用，并无实际需要翻译的内容。 md5:cf1b689a44aec285
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2561
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/2561。gf（Golang Foundation）是一个Go语言的开源框架，而"issues/2561"表示该仓库中编号为2561的问题或讨论。可能是用户在报告问题、请求功能或者讨论某个特定的代码问题。 md5:97cd71d9bf45e151
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2439
<原文结束>

# <翻译开始>
// 这段注释引用的是一个GitHub问题（issues）的链接，来自 "gf"（Go Foundation）项目，编号为2439。它可能是一个关于gf库的问题报告、讨论或者是一个已知问题的链接。具体的内容需要查看该链接才能得知。 md5:e37e02e670c04910
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2782
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/2782. md5:e2d84654d9404496
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2907
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/2907. md5:61d8552a7d7948bb
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/3086
<原文结束>

# <翻译开始>
// 这段注释引用的是GitHub上的一个 issue，gf（Go Foundation）是一个用Go语言编写的开源框架。3086号 issue 可能是关于gf框架的一个已知问题、错误报告、功能请求或者讨论点。具体的内容需要查看该issue的详细描述。 md5:629eedddf9f2ae76
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/3204
<原文结束>

# <翻译开始>
// 这段注释引用的是GitHub上的一个 issue，gf（Go Foundation）是一个用Go语言编写的开源框架。"3204"可能是指issue的编号，表示这个注释是在讨论或参考该框架中的第3204个问题或请求。具体的内容需要查看issue页面以获取详细信息。 md5:36c0adae03298bd3
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/3218
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/3218. md5:ebeb6327a156dd70
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2552
<原文结束>

# <翻译开始>
// 这段注释是指向GitHub上一个名为gf的项目的一个问题链接，问题编号为2552。在Go代码中，这种注释通常用于引用外部资源，如问题、讨论或文档，以便其他开发者了解代码的相关背景或上下文。 md5:23870b69cce8c4de
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2643
<原文结束>

# <翻译开始>
// 这段注释引用的是一个GitHub问题（issues）的链接，来自 "gf"（Go Foundation）项目，编号为2643。它可能是一个关于gf库的问题报告、讨论或者是一个已知问题的链接。具体的内容需要查看该链接才能了解。 md5:e98064ecba25be28
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/3238
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/3238。gf（GoGF）是一个用Go语言编写的高性能Web框架。这个注释可能是开发者在提到他们在gf项目中遇到的问题或提出的一个改进请求，3238号issue可能是一个已知问题的编号或者一个讨论的话题。 md5:98233bbbba37f999
# <翻译结束>

