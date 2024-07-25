
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
//github.com/gogf/gf/issues/819. md5:205f368062ae50a5
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
// 由于过滤特性从 GoFrame v1.16.0 版本起自动启用，此功能不再使用。
// 
//	func Test_DB_Insert_KeyFieldNameMapping_Error(t *testing.T) {
//		// 创建一个表用于测试
//		table := createTable()
//		defer dropTable(table) // 测试结束后删除该表
// 		
//		// 使用 gtest 包进行测试
//		gtest.C(t, func(t *gtest.T) {
//			// 定义一个 User 结构体来表示用户信息
//			type User struct {
//				Id             int    // 用户ID
//				Passport       string // 用户通行证
//				Password       string // 密码
//				Nickname       string // 昵称
//				CreateTime     string // 创建时间
//				NoneExistField string // 一个不存在于数据库中的字段
//			}
//			
//			// 准备一条用户数据
//			data := User{
//				Id:         1,                      // 设置用户ID
//				Passport:   "user_1",               // 设置通行证
//				Password:   "pass_1",               // 设置密码
//				Nickname:   "name_1",               // 设置昵称
//				CreateTime: "2020-10-10 12:00:01", // 设置创建时间
//			}
//			
//			// 尝试将数据插入数据库
//			_, err := db.Insert(ctx, table, data)
//			
//			// 断言：期望错误不为nil，即插入操作应因字段映射问题而失败
//			t.AssertNE(err, nil)
//		})
//	}
// 
// 上述代码是一个测试用例，旨在验证当尝试插入含有数据库中不存在的字段（NoneExistField）的结构体时，`db.Insert` 方法是否会正确返回错误。但从 GoFrame v1.16.0 起，这个特定的测试逻辑已不再适用，因为框架自动处理了这类问题。
// md5:589fbdf4cfdac41e
# <翻译结束>

