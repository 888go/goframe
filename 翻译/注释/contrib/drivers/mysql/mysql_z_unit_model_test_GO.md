
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
// Fix issue: https://github.com/gogf/gf/issues/819
<原文结束>

# <翻译开始>
// github.com/gogf/gf/issues/819. md5:205f368062ae50a5
# <翻译结束>


<原文开始>
// batch insert, retrieving last insert auto-increment id.
<原文结束>

# <翻译开始>
	// 批量插入，并获取最后插入的自增ID。 md5:b6507323b980f454
# <翻译结束>


<原文开始>
// Update + Fields(string)
<原文结束>

# <翻译开始>
	// 更新 + Fields(字符串). md5:df4e16d13da67d5e
# <翻译结束>


<原文开始>
// Count with cache, check internal ctx data feature.
<原文结束>

# <翻译开始>
	// 使用缓存计数，检查内部上下文数据特性。 md5:fa8263fd899afcec
# <翻译结束>


<原文开始>
// Auto creating struct object.
<原文结束>

# <翻译开始>
	// 自动创建结构体对象。 md5:4b196dfc1321dc30
# <翻译结束>


<原文开始>
// Auto create struct slice.
<原文结束>

# <翻译开始>
	// 自动创建结构体切片。 md5:78598f0d7f20b815
# <翻译结束>


<原文开始>
// fmt.Println(buffer.String())
<原文结束>

# <翻译开始>
		// 打印出buffer的内容字符串。 md5:3d49298f0e6d7a25
# <翻译结束>


<原文开始>
	// db.SetDebug(true)
	// defer db.SetDebug(false)
<原文结束>

# <翻译开始>
	// 将数据库设置为调试模式
	// 使用defer语句确保在函数返回前将数据库的调试模式重置为false md5:b9225b2fca692b91
# <翻译结束>


<原文开始>
// struct, automatic mapping and filtering.
<原文结束>

# <翻译开始>
	// 结构体，自动映射和过滤。 md5:8edea55227b914af
# <翻译结束>


<原文开始>
// Select with alias to struct.
<原文结束>

# <翻译开始>
	// 用别名选择到结构体。 md5:86d27c7f5b555a89
# <翻译结束>


<原文开始>
// Select with alias and join statement.
<原文结束>

# <翻译开始>
	// 使用别名和连接语句进行选择。 md5:5ae27281997ad29c
# <翻译结束>


<原文开始>
// Cache feature disabled.
<原文结束>

# <翻译开始>
			// 缓存功能已禁用。 md5:96110ddd3191b243
# <翻译结束>


<原文开始>
	// "id":          i,
	// "passport":    fmt.Sprintf(`user_%d`, i),
	// "password":    fmt.Sprintf(`pass_%d`, i),
	// "nickname":    fmt.Sprintf(`name_%d`, i),
	// "create_time": gtime.NewFromStr("2018-10-24 10:00:00").String(),
<原文结束>

# <翻译开始>
	// "id":          i, 	// ID: i
	// "passport":    fmt.Sprintf(`user_%d`, i), 	// 用户名: "user_" + i
	// "password":    fmt.Sprintf(`pass_%d`, i), 	// 密码: "pass_" + i
	// "nickname":    fmt.Sprintf(`name_%d`, i), 	// 昵称: "name_" + i
	// "create_time": gtime.NewFromStr("2018-10-24 10:00:00").String(), 	// 创建时间: "2018-10-24 10:00:00" 的字符串表示 md5:62b0a78a146bb60c
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1012
<原文结束>

# <翻译开始>
// 这段注释指的是在GitHub上的一个gf项目（Golang Fast Foundation，一个Go语言的优秀库）中的Issue 1012。"Issue"通常在GitHub上表示一个问题、错误报告或者改进的请求。所以，这个注释可能是在指有关gf库的一个已知问题或者开发者希望解决的问题，链接指向了该问题的具体页面。 md5:d21c0bba53139335
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1387
<原文结束>

# <翻译开始>
// 这段注释链接指向的是GitHub上的一个Issue，GF（Go Foundation）是一个Go语言的库或框架。"1387"可能是Issue的编号。具体的内容需要查看该链接才能得知，大致意思是关于GF项目在1387号问题上的讨论、报告了一个错误或者提出了一个特性请求。 md5:7c877c3e7a856cb1
# <翻译结束>


<原文开始>
// Using filter does not affect the outside value inside function.
<原文结束>

# <翻译开始>
// 使用过滤器不会影响函数内部的外部值。 md5:857585fd480ebfc6
# <翻译结束>


<原文开始>
// This is no longer used as the filter feature is automatically enabled from GoFrame v1.16.0.
// func Test_Model_Insert_KeyFieldNameMapping_Error(t *testing.T) {
//	table := createTable()
//	defer dropTable(table)
//
//	gtest.C(t, func(t *gtest.T) {
//		type User struct {
//			Id             int
//			Passport       string
//			Password       string
//			Nickname       string
//			CreateTime     string
//			NoneExistField string
//		}
//		data := User{
//			Id:         1,
//			Passport:   "user_1",
//			Password:   "pass_1",
//			Nickname:   "name_1",
//			CreateTime: "2020-10-10 12:00:01",
//		}
//		_, err := db.Model(table).Data(data).Insert()
//		t.AssertNE(err, nil)
//	})
// }
<原文结束>

# <翻译开始>
// 该代码已不再使用，因为从GoFrame v1.16.0开始，过滤功能会自动启用。
// 测试函数Test_Model_Insert_KeyFieldNameMapping_Error用于检查在模型插入时键字段映射错误的情况：
// 创建表
// defer 释放表（在测试结束后删除）
//
// 使用gtest进行测试：
// 定义User结构体，包含Id、Passport、Password、Nickname、CreateTime和NoneExistField字段
// 初始化一个User实例
// 使用db.Model方法，传入表名和数据，尝试插入记录
// 断言插入操作返回的错误不为nil，预期会出现错误，因为存在不存在的字段
//
// 注意：这个测试用例期望在执行时抛出错误，因为"NoneExistField"字段不在数据库表中。 md5:8afc06ac33d4aa16
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1159
<原文结束>

# <翻译开始>
// 这段注释链接指向的是GitHub上的一个 issues（问题或讨论），来自gogf（GoGF）项目。它表示这个注释与 issue #1159 相关，可能是对某个特定问题、错误报告、功能请求或者讨论的引用。具体的内容需要查看该issue页面以获取详细信息。 md5:ef2c3285217b52b1
# <翻译结束>


<原文开始>
// If the result is empty, it returns error.
<原文结束>

# <翻译开始>
	// 如果结果为空，它将返回错误。 md5:e3f68d57e41236a6
# <翻译结束>

