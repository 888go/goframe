
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
// This is no longer used as the filter feature is automatically enabled from GoFrame v1.16.0.
// func Test_DB_Insert_KeyFieldNameMapping_Error(t *testing.T) {
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
//		_, err := db.Insert(ctx, table, data)
//		t.AssertNE(err, nil)
//	})
// }
<原文结束>

# <翻译开始>
// 从GoFrame v1.16.0版本开始，此功能不再使用，因为过滤器特性已自动启用。
// func Test_DB_Insert_KeyFieldNameMapping_Error(t *testing.T) {
// 	// 创建并获取测试用表
// 	table := createTable()
// 	defer dropTable(table) // 在测试结束后删除测试用表
// 	// 使用gtest框架进行单元测试
// 	gtest.C(t, func(t *gtest.T) {
// 		// 定义User结构体
// 		type User struct {
// 			Id             int
// 			Passport       string
// 			Password       string
// 			Nickname       string
// 			CreateTime     string
// 			NoneExistField string // 不存在于数据库中的字段
// 		}
// 
// 		// 准备待插入的数据
// 		data := User{
// 			Id:         1,
// 			Passport:   "user_1",
// 			Password:   "pass_1",
// 			Nickname:   "name_1",
// 			CreateTime: "2020-10-10 12:00:01",
// 		}
// 
// 		// 尝试将数据插入到数据库中
// 		result, err := db.Insert(ctx, table, data)
// 
// 		// 断言错误不为nil，即预期该操作应出现错误
// 		t.AssertNE(err, nil)
// 	})
// }
// 上述代码是Go语言的一个单元测试函数，用于测试在存在键字段名称映射错误的情况下，执行数据库插入操作是否会返回非空错误。
# <翻译结束>






















<原文开始>
// batch insert struct
<原文结束>

# <翻译开始>
// 批量插入结构体
# <翻译结束>


<原文开始>
// update counter test.
<原文结束>

# <翻译开始>
// 更新计数器测试
# <翻译结束>


<原文开始>
// All types testing.
<原文结束>

# <翻译开始>
// 所有类型测试
# <翻译结束>

