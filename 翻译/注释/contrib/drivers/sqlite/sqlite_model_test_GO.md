
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
// Fix issue: https://github.com/gogf/gf/issues/819
<原文结束>

# <翻译开始>
// 解决问题：https://github.com/gogf/gf/issues/819
# <翻译结束>


<原文开始>
// batch insert, retrieving last insert auto-increment id.
<原文结束>

# <翻译开始>
// 批量插入数据，并获取最后插入的自增ID。
# <翻译结束>


<原文开始>
	// UPDATE...LIMIT
	// gtest.C(t, func(t *gtest.T) {
	// 	result, err := db.Model(table).Data("nickname", "T100").Where(1).Limit(2).Update()
	// 	t.AssertNil(err)
	// 	n, _ := result.RowsAffected()
	// 	t.Assert(n, 2)
<原文结束>

# <翻译开始>
// UPDATE...LIMIT
// 使用gtest进行单元测试，参数t为测试实例
// 根据条件更新表格中的一条记录，并限制只更新两条数据
// 更新table表中nickname为"T100"且满足Where条件(如id=1)的前两条记录
// 
// gtest.C(t, func(t *gtest.T) {
//   // 执行数据库更新操作，返回结果result和可能发生的错误err
//   result, err := db.Model(table).Data("nickname", "T100").Where(1).Limit(2).Update()
//   // 断言错误应为nil，即无错误发生
//   t.AssertNil(err)
//   // 获取更新影响的行数，_表示忽略可能的错误信息
//   n, _ := result.RowsAffected()
//   // 断言实际影响的行数应为2
//   t.Assert(n, 2)
// }
# <翻译结束>


<原文开始>
	// 	v1, err := db.Model(table).Fields("nickname").Where("id", 10).Value()
	// 	t.AssertNil(err)
	// 	t.Assert(v1.String(), "T100")
<原文结束>

# <翻译开始>
// 使用db.Model方法设置数据表为table，仅查询nickname字段，并根据id为10的条件进行查询，获取查询结果赋值给v1变量，同时返回可能的错误信息并存储在err变量中
// t.AssertNil(err) 表示断言err应为空（即没有错误发生）
// t.Assert(v1.String(), "T100") 表示断言v1变量转换为字符串后与"T100"相等
# <翻译结束>


<原文开始>
	// 	v2, err := db.Model(table).Fields("nickname").Where("id", 8).Value()
	// 	t.AssertNil(err)
	// 	t.Assert(v2.String(), "name_8")
	// })
<原文结束>

# <翻译开始>
// 获取表名为table的数据库模型，指定查询字段为nickname，条件为id为8的记录的值
// 并将查询结果赋值给变量v2，同时返回可能的错误信息err
// 
// v2, err := db.Model(table).Fields("nickname").Where("id", 8).Value()
// 断言错误err为nil，即查询过程中没有出现错误
// t.AssertNil(err)
// 断言v2转换为字符串后的结果为"name_8"
// t.Assert(v2.String(), "name_8")
// }) 表示闭合的测试用例或匿名函数
# <翻译结束>



























<原文开始>
// AllAndCount with normal result
<原文结束>

# <翻译开始>
// AllAndCount 正常结果时的操作
# <翻译结束>












<原文开始>
// AllAndCount with Join return CodeDbOperationError
<原文结束>

# <翻译开始>
// AllAndCount 使用 Join 并返回 CodeDbOperationError
# <翻译结束>


<原文开始>
// Count with cache, check internal ctx data feature.
<原文结束>

# <翻译开始>
// 带缓存计数，检查内部ctx数据特性。
# <翻译结束>


<原文开始>
// Auto creating struct object.
<原文结束>

# <翻译开始>
// 自动创建结构体对象。
# <翻译结束>












<原文开始>
// ScanAndCount with normal struct result
<原文结束>

# <翻译开始>
// ScanAndCount 使用普通结构体作为结果进行扫描并计数
# <翻译结束>


<原文开始>
// ScanAndCount with normal array result
<原文结束>

# <翻译开始>
// ScanAndCount 采用普通数组作为结果进行扫描并计数
# <翻译结束>

















<原文开始>
// ScanAndCount with join return CodeDbOperationError
<原文结束>

# <翻译开始>
// ScanAndCount 使用join并返回CodeDbOperationError
// （根据代码片段，可能的完整含义：）
// ```go
// ScanAndCount 函数在使用join操作进行扫描并计数时，如果发生数据库操作错误，则返回CodeDbOperationError错误码
// 请提供更多上下文信息以便提供更精确的翻译。
# <翻译结束>





































<原文开始>
// struct, automatic mapping and filtering.
<原文结束>

# <翻译开始>
// 结构体，自动映射和过滤。
# <翻译结束>


<原文开始>
	// table := createInitTable()
	// defer dropTable(table)
<原文结束>

# <翻译开始>
// 创建并初始化表格，将结果赋值给table变量
// table := createInitTable()
// 在函数结束时（defer关键字确保的）执行dropTable函数，传入table作为参数，用于删除已创建的表格
// defer dropTable(table)
# <翻译结束>


<原文开始>
	// DELETE...LIMIT
	// https://github.com/mattn/go-sqlite3/pull/802
	// gtest.C(t, func(t *gtest.T) {
	// 	result, err := db.Model(table).Where(1).Limit(2).Delete()
	// 	t.AssertNil(err)
	// 	n, _ := result.RowsAffected()
	// 	t.Assert(n, 2)
	// })
<原文结束>

# <翻译开始>
// DELETE...LIMIT
// 这段代码引用了 GitHub 上 go-sqlite3 库的一个 PR（#802）
// 使用 gtest 单元测试框架进行测试
// 测试内容如下：
// 针对 table 模型，使用 Where 条件为 1 并限制删除数量为 2 的记录执行删除操作
// gtest.C 函数中进行单元测试的主体部分：
// t 表示当前测试环境
// 调用 db.Model(table).Where(1).Limit(2).Delete() 删除符合条件的记录，将结果和错误信息分别赋值给 result 和 err
// 断言 err 为空（即无错误发生）
// 获取被影响的行数并赋值给 n，这里忽略了可能的错误信息
// 最后断言受影响的行数 n 等于 2，即成功删除了两条记录
# <翻译结束>







<原文开始>
// Select with alias to struct.
<原文结束>

# <翻译开始>
// 使用别名选择到结构体。
# <翻译结束>


<原文开始>
// Select with alias and join statement.
<原文结束>

# <翻译开始>
// 使用别名和连接语句进行选择。
# <翻译结束>






















<原文开始>
	// "id":          i,
	// "passport":    fmt.Sprintf(`user_%d`, i),
	// "password":    fmt.Sprintf(`pass_%d`, i),
	// "nickname":    fmt.Sprintf(`name_%d`, i),
	// "create_time": gtime.NewFromStr(CreateTime).String(),
<原文结束>

# <翻译开始>
// "id":          i, // "id"字段：存储变量i的值
// "passport":    fmt.Sprintf(`user_%d`, i), // "passport"字段：格式化输出字符串，形式为"user_数字"，其中数字为变量i的值
// "password":    fmt.Sprintf(`pass_%d`, i), // "password"字段：格式化输出字符串，形式为"pass_数字"，其中数字为变量i的值
// "nickname":    fmt.Sprintf(`name_%d`, i), // "nickname"字段：格式化输出字符串，形式为"name_数字"，其中数字为变量i的值
// "create_time": gtime.NewFromStr(CreateTime).String(), // "create_time"字段：通过CreateTime字符串创建一个gtime.Time对象，并获取该时间对象的字符串表示形式
// （注：这里的`CreateTime`应是一个符合日期时间格式的字符串）
# <翻译结束>


<原文开始>
// Issue: https://github.com/gogf/gf/issues/1002
<原文结束>

# <翻译开始>
// 问题：https://github.com/gogf/gf/issues/1002
# <翻译结束>







<原文开始>
// where + gtime.Time arguments.
<原文结束>

# <翻译开始>
// where + gtime.Time 参数
# <翻译结束>


<原文开始>
	// TODO
	// where + time.Time arguments, UTC.
	// gtest.C(t, func(t *gtest.T) {
	// 	t1, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:32")
	// 	t2, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:34")
	// 	{
	// 		v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).Value()
	// 		t.AssertNil(err)
	// 		t.Assert(v.Int(), 1)
	// 	}
	// })
<原文结束>

# <翻译开始>
// TODO
// 对包含time.Time参数的where条件进行测试，时间采用UTC格式。
// 使用gtest.C进行单元测试，传入t参数作为测试环境上下文。
// 
// 首先，解析两个时间字符串为time.Time类型：
// t1表示"2020-10-27 11:03:32"
// t2表示"2020-10-27 11:03:34"
// 
// 然后执行如下代码块：
// 通过db.Model(table).Fields("id")设置SQL查询模型和字段（这里为表名为table的表中的id字段），
// 并添加where条件：create_time>? and create_time<?，其中问号占位符分别被t1和t2替换。
// 执行查询并获取查询结果值v及可能的错误err。
// 验证错误err应为nil，即查询无错误发生。
// 验证查询结果值v转换为整数后为1。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1012
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf项目的一个问题链接，具体为第1012号问题
# <翻译结束>


<原文开始>
		// TODO
		// t.Assert(userEntity.CreatedAt.String(), "2020-11-22 11:23:45")
		// t.Assert(userEntity.UpdatedAt.String(), "2020-11-22 12:23:45")
		// t.Assert(gtime.NewFromTime(userEntity.DeletedAt).String(), "2020-11-22 13:23:45")
<原文结束>

# <翻译开始>
// TODO
// t.Assert(userEntity.CreatedAt.String(), "2020-11-22 11:23:45") // 待办：断言userEntity的CreatedAt字段转换为字符串后，其值应为"2020-11-22 11:23:45"
// t.Assert(userEntity.UpdatedAt.String(), "2020-11-22 12:23:45") // 待办：断言userEntity的UpdatedAt字段转换为字符串后，其值应为"2020-11-22 12:23:45"
// t.Assert(gtime.NewFromTime(userEntity.DeletedAt).String(), "2020-11-22 13:23:45") // 待办：断言根据userEntity的DeletedAt字段创建的新时间对象转换为字符串后，其值应为"2020-11-22 13:23:45"
# <翻译结束>







<原文开始>
// https://github.com/gogf/gf/issues/1387
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf仓库中关于第1387号问题的链接
# <翻译结束>


<原文开始>
// Using filter does not affect the outside value inside function.
<原文结束>

# <翻译开始>
// 在函数内部使用filter不会影响外部的值。
# <翻译结束>


<原文开始>
// This is no longer used as the filter feature is automatically enabled from GoFrame v1.16.0.
<原文结束>

# <翻译开始>
// 从GoFrame v1.16.0版本开始，此功能不再使用，因为过滤功能会自动启用。
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1159
<原文结束>

# <翻译开始>
// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库的第1159号问题。
// 中文翻译：
// 参考GitHub上gogf/gf项目的问题1159。
# <翻译结束>
































<原文开始>
// fmt.Println(buffer.String())
<原文结束>

# <翻译开始>
// 打印buffer.String()的输出结果
# <翻译结束>


<原文开始>
// Update + Data(string)
<原文结束>

# <翻译开始>
// 更新 + 数据(字符串)
# <翻译结束>


<原文开始>
// Update + Fields(string)
<原文结束>

# <翻译开始>
// Update + Fields(string)
// 更新 + 字段(string)
# <翻译结束>


<原文开始>
// AllAndCount with all data
<原文结束>

# <翻译开始>
// AllAndCount 获取所有数据及计数
# <翻译结束>


<原文开始>
// AllAndCount with no data
<原文结束>

# <翻译开始>
// AllAndCount 无数据时
# <翻译结束>


<原文开始>
// AllAndCount with page
<原文结束>

# <翻译开始>
// AllAndCount 带分页功能
# <翻译结束>


<原文开始>
// AllAndCount with distinct
<原文结束>

# <翻译开始>
// AllAndCount 带有 distinct 的功能
# <翻译结束>


<原文开始>
// Auto create struct slice.
<原文结束>

# <翻译开始>
// 自动创建结构体切片。
# <翻译结束>


<原文开始>
// ScanAndCount with page
<原文结束>

# <翻译开始>
// ScanAndCount 带分页功能
// （注：由于没有提供完整的代码上下文，这里的翻译可能不够精确。根据现有信息，“ScanAndCount with page”可以理解为这个函数或方法用于扫描数据并结合分页进行计数。）
# <翻译结束>


<原文开始>
// ScanAndCount with distinct
<原文结束>

# <翻译开始>
// ScanAndCount 带有唯一性的扫描并计数
# <翻译结束>


<原文开始>
// map + slice parameter
<原文结束>

# <翻译开始>
// map + slice 参数
# <翻译结束>


<原文开始>
// gmap.Map key operator
<原文结束>

# <翻译开始>
// gmap.Map 键操作器
# <翻译结束>


<原文开始>
// list map key operator
<原文结束>

# <翻译开始>
// 列表映射键操作员
# <翻译结束>


<原文开始>
// tree map key operator
<原文结束>

# <翻译开始>
// 树状映射键操作器
# <翻译结束>


<原文开始>
// complicated where 1
<原文结束>

# <翻译开始>
// 复杂条件 1
# <翻译结束>


<原文开始>
// complicated where 2
<原文结束>

# <翻译开始>
// 复杂条件 2
# <翻译结束>


<原文开始>
// Select with alias.
<原文结束>

# <翻译开始>
// 使用别名进行选择。
# <翻译结束>


<原文开始>
// make cache for id 3
<原文结束>

# <翻译开始>
// 为id 3创建缓存
# <翻译结束>


<原文开始>
// make cache for id 4
<原文结束>

# <翻译开始>
// 为id 4创建缓存
# <翻译结束>


<原文开始>
// Cache feature disabled.
<原文结束>

# <翻译开始>
// 缓存功能已禁用。
# <翻译结束>


<原文开始>
// where + string arguments.
<原文结束>

# <翻译开始>
// where + 字符串参数。
# <翻译结束>


<原文开始>
// fmt.Printf("%+v", err)
<原文结束>

# <翻译开始>
// 使用 %+v 格式化输出错误信息，会包含错误类型和详细堆栈信息
# <翻译结束>

