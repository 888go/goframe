
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
	// Map
	// gtest.C(t, func(t *gtest.T) {
	//	var v *V
	//	err := gconv.Struct(g.Map{
	//		"name":  "john",
	//		"array": g.Slice{1, 2, 3},
	//	}, &v)
	//	t.AssertNil(err)
	//	t.Assert(v.Name, "john")
	//	t.Assert(v.Array.Slice(), g.Slice{1, 2, 3})
	// })
<原文结束>

# <翻译开始>
	// 映射（Map）
	// 使用gtest编写测试用例(t *gtest.T)：
	// 创建一个V类型的指针变量v
	// 尝试将映射g.Map{
	//   "name": "john",
	//   "array": [1, 2, 3]
	// } 转换为结构体，并赋值给v
	// 验证转换是否成功，如果失败则设置t.AssertNil(err)
	// 验证v的Name字段是否等于"john"
	// 验证v的Array.Slice()方法返回的切片是否等于[1, 2, 3]
	// md5:1684dcec1caa154e
# <翻译结束>

