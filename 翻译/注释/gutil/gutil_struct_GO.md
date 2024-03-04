
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
// StructToSlice converts struct to slice of which all keys and values are its items.
// Eg: {"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]
<原文结束>

# <翻译开始>
// StructToSlice 将结构体转换为键值对构成的切片。
// 例如：{"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]
# <翻译结束>


<原文开始>
		// Note that, it uses the gconv tag name instead of the attribute name if
		// the gconv tag is fined in the struct attributes.
<原文结束>

# <翻译开始>
// 注意，如果在结构体属性中找到gconv标签，则它使用gconv标签名称而非属性名称。
# <翻译结束>


<原文开始>
// FillStructWithDefault fills  attributes of pointed struct with tag value from `default/d` tag .
// The parameter `structPtr` should be either type of *struct/[]*struct.
<原文结束>

# <翻译开始>
// FillStructWithDefault 用 `default/d` 标签的值填充指针指向的结构体属性。
// 参数 `structPtr` 应为 *struct 或 []*struct 类型。
# <翻译结束>

