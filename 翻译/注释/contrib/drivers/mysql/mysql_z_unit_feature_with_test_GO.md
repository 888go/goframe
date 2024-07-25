
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
// With part attribute: UserDetail.
<原文结束>

# <翻译开始>
	// 带有部分属性：UserDetail.. md5:5d4187f92cc37f3c
# <翻译结束>


<原文开始>
// With part attribute: UserScores.
<原文结束>

# <翻译开始>
	// 配置部分属性：UserScores。 md5:c8e42122566fe2d2
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
	// 使用gtest.C进行测试，参数为t（*gtest.T）
	// 定义一个User类型的指针变量user
	// 使用db的ORM方式，执行查询语句（获取tableUser表中id为3的所有数据），并将结果赋值给user
	// 断言错误(err)为nil
	// 断言user的ID字段值为3
	// 断言user的UserDetail字段不为nil
	// 断言user的UserDetail UserID字段值为3
	// 断言user的UserDetail Address字段值为`address_3`
	// 断言user的UserScores切片长度为5
	// 断言user的UserScores切片的第一个元素UserID字段值为3
	// 断言user的UserScores切片的第一个元素Score字段值为1
	// 断言user的UserScores切片的第5个元素UserID字段值为3
	// 断言user的UserScores切片的第5个元素Score字段值为5 md5:1ebf51134a7a3187
# <翻译结束>

