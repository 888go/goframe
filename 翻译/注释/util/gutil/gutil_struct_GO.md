
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
// StructToSlice converts struct to slice of which all keys and values are its items.
// Eg: {"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]
<原文结束>

# <翻译开始>
// StructToSlice 将结构体转换为一个键值对作为元素的切片。
// 例如：{"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"] md5:ca8c34ec711fb0de
# <翻译结束>


<原文开始>
		// Note that, it uses the gconv tag name instead of the attribute name if
		// the gconv tag is fined in the struct attributes.
<原文结束>

# <翻译开始>
		// 如果在结构体属性中找到了gconv标签，它将使用gconv标签名而不是属性名。 md5:697077ff458895f0
# <翻译结束>


<原文开始>
// FillStructWithDefault fills  attributes of pointed struct with tag value from `default/d` tag .
// The parameter `structPtr` should be either type of *struct/[]*struct.
<原文结束>

# <翻译开始>
// FillStructWithDefault 使用`default/d`标签的值填充指向的结构体的属性。参数`structPtr`应该是`*struct`或`[]*struct`类型的一种。 md5:5777fe6fdb6efa8a
# <翻译结束>

