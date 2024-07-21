
<原文开始>
// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 2019 gf 作者（https://github.com/gogf/gf）。保留所有权利。
//
// 此源代码形式受麻省理工学院（MIT）许可证的条款约束。
// 如果未随此文件一起分发MIT许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:47e609239e0cb2bc
# <翻译结束>


<原文开始>
// The test scenario index of this test case (exact matching field) is a keyword in the Dameng database and cannot exist as a field name.
// If the data structure previously migrated from mysql has an index (completely matching field), it will also be allowed.
// However, when processing the index (completely matching field), the adapter will automatically add security character
// In principle, such problems will not occur if you directly use Dameng database initialization instead of migrating the data structure from mysql.
// If so, the adapter has also taken care of it.
<原文结束>

# <翻译开始>
// 该测试用例的测试场景索引（完全匹配字段）在达梦数据库中是一个保留关键字，不能作为字段名存在。
// 如果之前从mysql迁移过来的数据结构中有一个索引（完全匹配字段），也会被允许。
// 但是，在处理索引（完全匹配字段）时，适配器会自动添加安全字符。
// 原则上，如果你直接使用达梦数据库初始化，而不是从mysql迁移数据结构，就不会出现此类问题。
// 即使如此，适配器也已经考虑到了这种情况。
// md5:bdfc451ff291c639
# <翻译结束>


<原文开始>
// model.Where("account_name like ?", "%"+"list"+"%")
<原文结束>

# <翻译开始>
// 使用model的Where方法，传入一个SQL条件："account_name"字段包含"%list%"字符串。 md5:008bda1d9a70b4f2
# <翻译结束>


<原文开始>
// g.Model.insert not lost default not null coloumn
<原文结束>

# <翻译开始>
// g.Model.insert 不会丢失默认非空列. md5:475cfebfed134e1b
# <翻译结束>


<原文开始>
// _, err := db.Schema(TestDBName).Model(table).Data(data).Insert()
<原文结束>

# <翻译开始>
// _, err := db.Schema(TestDBName).Model(table).Data(data).Insert()
// 使用TestDBName数据库的模式，根据table模型和data数据执行插入操作，返回一个表示是否成功的空值和错误信息。 md5:665c442bb4e1be49
# <翻译结束>


<原文开始>
		// TODO Question2
		// this is DM bug.
		// t.Assert(one["CREATED_TIME"].GTime().String(), timeStr)
<原文结束>

# <翻译开始>
		// 待办事项：问题2
		// 这是DM（可能是某个项目或模块的缩写）的bug。
		// 断言one["CREATED_TIME"]的GTime转换后字符串与timeStr相等。
		// md5:6c078020ce38de99
# <翻译结束>

