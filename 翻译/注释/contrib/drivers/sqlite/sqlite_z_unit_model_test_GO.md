
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
// Fix issue: https://github.com/gogf/gf/issues/819
<原文结束>

# <翻译开始>
// 解决问题：https://github.com/gogf/gf/issues/819. md5:205f368062ae50a5
# <翻译结束>


<原文开始>
// batch insert, retrieving last insert auto-increment id.
<原文结束>

# <翻译开始>
// 批量插入，并获取最后插入的自增ID。. md5:b6507323b980f454
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
// 更新...限制
// 使用gtest进行测试，传入t作为测试上下文
// 执行如下操作：
// 根据模型table，设置数据字段"nickname"为"T100"，
// 并在满足条件1的情况下，限制更新操作影响的行数为2。
// 获取更新操作的结果与错误信息。
// 断言：期望错误为nil。
// 计算并获取更新影响的行数，忽略此操作可能产生的错误。
// 断言：期望更新影响的行数为2。
// md5:cfae918cd0afb1ea
# <翻译结束>


<原文开始>
	// 	v1, err := db.Model(table).Fields("nickname").Where("id", 10).Value()
	// 	t.AssertNil(err)
	// 	t.Assert(v1.String(), "T100")
<原文结束>

# <翻译开始>
// 通过$db$查询$table$表中id为10的nickname字段值，赋值给v1，预期可能产生错误err
// t.AssertNil(err)：断言错误err为nil，即无错误发生
// t.Assert(v1.String(), "T100")：断言v1转换为字符串后的值等于"T100"
// md5:a2bbef8eea48f43a
# <翻译结束>


<原文开始>
	// 	v2, err := db.Model(table).Fields("nickname").Where("id", 8).Value()
	// 	t.AssertNil(err)
	// 	t.Assert(v2.String(), "name_8")
	// })
<原文结束>

# <翻译开始>
// 使用$db$操作数据库，根据模型$table$获取nickname字段，查询id为8的记录，并获取其值。
// 验证错误是否为nil。
// 验证获取到的值（v2）是否等于"name_8"。
// }
// md5:0005058975deac4b
# <翻译结束>


<原文开始>
// Update + Fields(string)
<原文结束>

# <翻译开始>
// 更新 + Fields(字符串). md5:df4e16d13da67d5e
# <翻译结束>


<原文开始>
// AllAndCount with all data
<原文结束>

# <翻译开始>
// 使用所有数据的AllAndCount. md5:04233fbd8b956565
# <翻译结束>


<原文开始>
// AllAndCount with no data
<原文结束>

# <翻译开始>
// AllAndCount 无数据情况. md5:78116cd399301bd7
# <翻译结束>


<原文开始>
// AllAndCount with normal result
<原文结束>

# <翻译开始>
// AllAndCount 返回正常结果. md5:d132fb7fcbc86207
# <翻译结束>


<原文开始>
// AllAndCount with distinct
<原文结束>

# <翻译开始>
// 所有唯一项并计数. md5:ecb27c1ddcd9a325
# <翻译结束>


<原文开始>
// AllAndCount with Join return CodeDbOperationError
<原文结束>

# <翻译开始>
// AllAndCount 与 Join 方法返回 CodeDbOperationError. md5:e59618ae9d29f9f5
# <翻译结束>


<原文开始>
// Count with cache, check internal ctx data feature.
<原文结束>

# <翻译开始>
// 使用缓存计数，检查内部上下文数据特性。. md5:fa8263fd899afcec
# <翻译结束>


<原文开始>
// Auto creating struct object.
<原文结束>

# <翻译开始>
// 自动创建结构体对象。. md5:4b196dfc1321dc30
# <翻译结束>


<原文开始>
// Auto create struct slice.
<原文结束>

# <翻译开始>
// 自动创建结构体切片。. md5:78598f0d7f20b815
# <翻译结束>


<原文开始>
// fmt.Println(buffer.String())
<原文结束>

# <翻译开始>
// fmt.Println(buffer.String()) 翻译成中文为：
// 打印出buffer的内容字符串。. md5:3d49298f0e6d7a25
# <翻译结束>


<原文开始>
// ScanAndCount with normal struct result
<原文结束>

# <翻译开始>
// 使用普通结构体结果的ScanAndCount. md5:941b5fec0e73797f
# <翻译结束>


<原文开始>
// ScanAndCount with normal array result
<原文结束>

# <翻译开始>
// ScanAndCount 使用常规数组作为结果. md5:640a035a18ac03db
# <翻译结束>


<原文开始>
// ScanAndCount with distinct
<原文结束>

# <翻译开始>
// 使用distinct进行扫描和计数. md5:5afa1e02dbecba67
# <翻译结束>


<原文开始>
// ScanAndCount with join return CodeDbOperationError
<原文结束>

# <翻译开始>
// 使用连接执行ScanAndCount，返回CodeDbOperationError. md5:28f0d53619e4ce12
# <翻译结束>


<原文开始>
// struct, automatic mapping and filtering.
<原文结束>

# <翻译开始>
// 结构体，自动映射和过滤。. md5:8edea55227b914af
# <翻译结束>


<原文开始>
	// table := createInitTable()
	// defer dropTable(table)
<原文结束>

# <翻译开始>
// table := createInitTable() // 创建初始化表
// defer dropTable(table)    // 延迟执行，删除表
// md5:b569b2401cb8568d
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
// 参考: https://github.com/mattn/go-sqlite3/pull/802
// gtest.C(t, func(t *gtest.T) {
// 	删除结果, err := db.Model(table).Where(1).Limit(2).Delete()
// 	t.AssertNil(err)
// 	影响行数, _ := result.RowsAffected()
// 	t.Assert(影响行数, 2)
// })
// md5:63b42e136740eea6
# <翻译结束>


<原文开始>
// Select with alias to struct.
<原文结束>

# <翻译开始>
// 用别名选择到结构体。. md5:86d27c7f5b555a89
# <翻译结束>


<原文开始>
// Select with alias and join statement.
<原文结束>

# <翻译开始>
// 使用别名和连接语句进行选择。. md5:5ae27281997ad29c
# <翻译结束>


<原文开始>
// Cache feature disabled.
<原文结束>

# <翻译开始>
// 缓存功能已禁用。. md5:96110ddd3191b243
# <翻译结束>


<原文开始>
	// "id":          i,
	// "passport":    fmt.Sprintf(`user_%d`, i),
	// "password":    fmt.Sprintf(`pass_%d`, i),
	// "nickname":    fmt.Sprintf(`name_%d`, i),
	// "create_time": gtime.NewFromStr(CreateTime).String(),
<原文结束>

# <翻译开始>
// "id":          i, // 用户ID
// "passport":    fmt.Sprintf("user_%d", i), // 通行证（格式为"user_编号")
// "password":    fmt.Sprintf("pass_%d", i), // 密码（格式为"pass_编号")
// "nickname":    fmt.Sprintf("name_%d", i), // 昵称（格式为"name_编号")
// "create_time": gtime.NewFromStr(CreateTime).String(), // 创建时间（将CreateTime字符串转换为gtime格式并转为字符串）
// md5:ddd0764dc67c4e9f
# <翻译结束>


<原文开始>
// Issue: https://github.com/gogf/gf/issues/1002
<原文结束>

# <翻译开始>
// 问题：https://github.com/gogf/gf/issues/1002. md5:2b9ad829e9523427
# <翻译结束>


<原文开始>
// where + string arguments.
<原文结束>

# <翻译开始>
// where + 字符串参数。. md5:cb1db92222691d4d
# <翻译结束>


<原文开始>
// where + gtime.Time arguments.
<原文结束>

# <翻译开始>
// 其中包含 gtime.Time 类型的参数。. md5:3bd9bb993dd2cc53
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
// 待办事项
// 在这里使用 + time.Time 参数，采用 UTC 时间。
// gtest.C(t, func(t *gtest.T) {
//   t1, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:32") // 解析时间字符串为 t1
//   t2, _ := time.Parse("2006-01-02 15:04:05", "2020-10-27 11:03:34") // 解析时间字符串为 t2
//   {
//     v, err := db.Model(table).Fields("id").Where("create_time>? and create_time<?", t1, t2).Value() // 查询创建时间在 t1 和 t2 之间记录的 id
//     t.AssertNil(err) // 断言 err 为空，即查询无错误
//     t.Assert(v.Int(), 1) // 断言查询结果的整数值为 1
//   }
// })
// md5:6089a1ebb4983ace
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1012
<原文结束>

# <翻译开始>
// https://github.com/gogf/gf/issues/1012
// 
// 这段注释指的是在GitHub上的一个gf项目（Golang Fast Foundation，一个Go语言的优秀库）中的Issue 1012。"Issue"通常在GitHub上表示一个问题、错误报告或者改进的请求。所以，这个注释可能是在指有关gf库的一个已知问题或者开发者希望解决的问题，链接指向了该问题的具体页面。. md5:d21c0bba53139335
# <翻译结束>


<原文开始>
		// TODO
		// t.Assert(userEntity.CreatedAt.String(), "2020-11-22 11:23:45")
		// t.Assert(userEntity.UpdatedAt.String(), "2020-11-22 12:23:45")
		// t.Assert(gtime.NewFromTime(userEntity.DeletedAt).String(), "2020-11-22 13:23:45")
<原文结束>

# <翻译开始>
// TODO
// t.Assert(userEntity.CreatedAt.String(), "2020-11-22 11:23:45") // 断言用户实体的创建时间字符串为 "2020-11-22 11:23:45"
// t.Assert(userEntity.UpdatedAt.String(), "2020-11-22 12:23:45") // 断言用户实体的更新时间字符串为 "2020-11-22 12:23:45"
// t.Assert(gtime.NewFromTime(userEntity.DeletedAt).String(), "2020-11-22 13:23:45") // 断言用户实体的删除时间（转换为gtime类型）字符串为 "2020-11-22 13:23:45"
// md5:8ad9ae5f1d9029d0
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1387
<原文结束>

# <翻译开始>
// https://github.com/gogf/gf/issues/1387
// 
// 这段注释链接指向的是GitHub上的一个Issue，GF（Go Foundation）是一个Go语言的库或框架。"1387"可能是Issue的编号。具体的内容需要查看该链接才能得知，大致意思是关于GF项目在1387号问题上的讨论、报告了一个错误或者提出了一个特性请求。. md5:7c877c3e7a856cb1
# <翻译结束>


<原文开始>
// Using filter does not affect the outside value inside function.
<原文结束>

# <翻译开始>
// 使用过滤器不会影响函数内部的外部值。. md5:857585fd480ebfc6
# <翻译结束>


<原文开始>
// This is no longer used as the filter feature is automatically enabled from GoFrame v1.16.0.
<原文结束>

# <翻译开始>
// 从GoFrame v1.16.0开始，此功能不再使用，因为过滤功能已自动启用。. md5:a491426db314e6d6
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1159
<原文结束>

# <翻译开始>
// https://github.com/gogf/gf/issues/1159
// 
// 这段注释链接指向的是GitHub上的一个 issues（问题或讨论），来自gogf（GoGF）项目。它表示这个注释与 issue #1159 相关，可能是对某个特定问题、错误报告、功能请求或者讨论的引用。具体的内容需要查看该issue页面以获取详细信息。. md5:ef2c3285217b52b1
# <翻译结束>

