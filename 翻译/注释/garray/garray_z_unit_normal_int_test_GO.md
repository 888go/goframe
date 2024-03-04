
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
// Map
// 使用gtest.C进行测试，参数t为gtest.T类型指针
// 定义一个*v类型的指针变量v
// 将g.Map类型的键值对（"name": "john", "array": g.Slice{1, 2, 3}）通过gconv.Struct转换到结构体变量v中
// 断言转换过程中的错误为nil
// 断言转换后v的Name字段值为"john"
// 断言转换后v的Array字段转换为切片后的值等于g.Slice{1, 2, 3}
# <翻译结束>

