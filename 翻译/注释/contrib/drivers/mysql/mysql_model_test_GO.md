
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
	// db.SetDebug(true)
	// defer db.SetDebug(false)
<原文结束>

# <翻译开始>
// 设置数据库调试模式为开启状态
// db.SetDebug(true)
// 在函数结束时，确保关闭数据库调试模式
// defer db.SetDebug(false)
# <翻译结束>





































<原文开始>
// struct, automatic mapping and filtering.
<原文结束>

# <翻译开始>
// 结构体，自动映射和过滤。
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
	// "create_time": gtime.NewFromStr("2018-10-24 10:00:00").String(),
<原文结束>

# <翻译开始>
// "id":          i, // "id"字段：i的值
// "passport":    fmt.Sprintf(`user_%d`, i), // "passport"字段：格式化输出字符串，形如"user_1"，其中%d用i的值替换
// "password":    fmt.Sprintf(`pass_%d`, i), // "password"字段：格式化输出字符串，形如"pass_1"，其中%d用i的值替换
// "nickname":    fmt.Sprintf(`name_%d`, i), // "nickname"字段：格式化输出字符串，形如"name_1"，其中%d用i的值替换
// "create_time": gtime.NewFromStr("2018-10-24 10:00:00").String(), // "create_time"字段：创建一个时间对象，使用"2018-10-24 10:00:00"字符串初始化，并将其转换为字符串表示形式
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1012
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf项目的一个问题链接，具体为第1012号问题
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
// 从GoFrame v1.16.0开始，此功能不再使用，因为过滤特性会自动启用。
// func Test_Model_Insert_KeyFieldNameMapping_Error(t *testing.T) {
// 	// 创建测试表
// 	table := createTable()
// 	// 在测试结束后删除测试表
// 	defer dropTable(table)
//
// 	// 使用gtest进行单元测试
// 	gtest.C(t, func(t *gtest.T) {
//  	// 定义User结构体
//  	type User struct {
//  		Id             int
//  		Passport       string
//  		Password       string
//  		Nickname       string
//  		CreateTime     string
//  		NoneExistField string // 不存在的字段
//  	}
//  	// 初始化用户数据
//  	data := User{
//  		Id:         1,
//  		Passport:   "user_1",
//  		Password:   "pass_1",
//  		Nickname:   "name_1",
//  		CreateTime: "2020-10-10 12:00:01",
//  	}
//  	// 尝试将数据插入到指定表中
//  	result, err := db.Model(table).Data(data).Insert()
//  	// 断言错误不为空
//  	t.AssertNE(err, nil)
//  })
// }
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
// If the result is empty, it returns error.
<原文结束>

# <翻译开始>
// 如果结果为空，则返回错误。
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
// Auto create struct slice.
<原文结束>

# <翻译开始>
// 自动创建结构体切片。
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
// fmt.Printf("%+v", err)
<原文结束>

# <翻译开始>
// 使用 %+v 格式化输出错误信息，会包含错误类型和详细堆栈信息
# <翻译结束>

