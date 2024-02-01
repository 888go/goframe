
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
// With part attribute: UserDetail.
<原文结束>

# <翻译开始>
// 带有部分属性: UserDetail.
# <翻译结束>


<原文开始>
// With part attribute: UserScores.
<原文结束>

# <翻译开始>
// 带有部分属性: UserScores.
# <翻译结束>


<原文开始>
	// gtest.C(t, func(t *gtest.T) {
	//	var user *User
	//	err := db.Model(tableUser).WithAll().Where("id", 3).Scan(&user)
	//	t.AssertNil(err)
	//	t.Assert(user.ID, 3)
	//	t.AssertNE(user.UserDetail, nil)
	//	t.Assert(user.UserDetail.UserID, 3)
	//	t.Assert(user.UserDetail.Address, `address_3`)
	//	t.Assert(len(user.UserScores), 5)
	//	t.Assert(user.UserScores[0].UserID, 3)
	//	t.Assert(user.UserScores[0].Score, 1)
	//	t.Assert(user.UserScores[4].UserID, 3)
	//	t.Assert(user.UserScores[4].Score, 5)
	// })
<原文结束>

# <翻译开始>
// gtest.C(t, func(t *gtest.T) { // 使用gtest框架对代码进行单元测试
//	var user *User // 声明一个指向User类型的指针变量user
//	err := db.Model(tableUser).WithAll().Where("id", 3).Scan(&user) // 根据id为3查询tableUser表中的数据到user变量中
//	t.AssertNil(err) // 断言查询过程中无错误发生，即err应为nil
//	t.Assert(user.ID, 3) // 断言查询结果中user的ID属性为3
//	t.AssertNE(user.UserDetail, nil) // 断言user的UserDetail属性不为空（非nil）
//	t.Assert(user.UserDetail.UserID, 3) // 断言user的UserDetail结构体中的UserID属性为3
//	t.Assert(user.UserDetail.Address, `address_3`) // 断言user的UserDetail结构体中的Address属性为"address_3"
//	t.Assert(len(user.UserScores), 5) // 断言user的UserScores切片长度为5
//	t.Assert(user.UserScores[0].UserID, 3) // 断言user的UserScores切片中第一个元素的UserID属性为3
//	t.Assert(user.UserScores[0].Score, 1) // 断言user的UserScores切片中第一个元素的Score属性为1
//	t.Assert(user.UserScores[4].UserID, 3) // 断言user的UserScores切片中最后一个元素的UserID属性为3
//	t.Assert(user.UserScores[4].Score, 5) // 断言user的UserScores切片中最后一个元素的Score属性为5
// }) // 结束gtest.C()函数的测试用例定义
# <翻译结束>







<原文开始>
// Initialize the data.
<原文结束>

# <翻译开始>
// 初始化数据
# <翻译结束>

