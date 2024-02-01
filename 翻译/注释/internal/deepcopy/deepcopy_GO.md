
<原文开始>
// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame gf 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
# <翻译结束>


<原文开始>
// Package deepcopy makes deep copies of things using reflection.
//
// This package is maintained from: https://github.com/mohae/deepcopy
<原文结束>

# <翻译开始>
// Package deepcopy 提供了通过反射实现深拷贝的功能。
//
// 该包来源于：https://github.com/mohae/deepcopy
# <翻译结束>


<原文开始>
// Interface for delegating copy process to type
<原文结束>

# <翻译开始>
// 为将复制过程委托给类型定义的接口
# <翻译结束>


<原文开始>
// Copy creates a deep copy of whatever is passed to it and returns the copy
// in an interface{}.  The returned value will need to be asserted to the
// correct type.
<原文结束>

# <翻译开始>
// Copy 函数接收任意参数，并创建其深度拷贝，然后以 interface{} 类型返回该拷贝。
// 返回的值需要断言为正确的类型。
# <翻译结束>







<原文开始>
// Make the interface a reflect.Value
<原文结束>

# <翻译开始>
// 将接口转换为 reflect.Value 类型
# <翻译结束>


<原文开始>
// Make a copy of the same type as the original.
<原文结束>

# <翻译开始>
// 创建一个与原始类型相同的副本。
# <翻译结束>


<原文开始>
// Recursively copy the original.
<原文结束>

# <翻译开始>
// 递归地复制原对象。
# <翻译结束>


<原文开始>
// Return the copy as an interface.
<原文结束>

# <翻译开始>
// 返回作为接口的副本。
# <翻译结束>


<原文开始>
// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
<原文结束>

# <翻译开始>
// copyRecursive 实现了接口的实际复制操作。当前，它对所能处理的数据类型的支持有限。根据需要进行添加。
# <翻译结束>


<原文开始>
// check for implement deepcopy.Interface
<原文结束>

# <翻译开始>
// 检查是否实现了 deepcopy.Interface 接口
# <翻译结束>


<原文开始>
// handle according to original's Kind
<原文结束>

# <翻译开始>
// 根据原始的 Kind 进行处理
# <翻译结束>


<原文开始>
// Get the actual value being pointed to.
<原文结束>

# <翻译开始>
// 获取所指向的实际值。
# <翻译结束>


<原文开始>
// if  it isn't valid, return.
<原文结束>

# <翻译开始>
// 如果它无效，则返回。
# <翻译结束>


<原文开始>
// If this is a nil, don't do anything
<原文结束>

# <翻译开始>
// 如果这是一个nil，那么什么也不做
# <翻译结束>


<原文开始>
// Get the value for the interface, not the pointer.
<原文结束>

# <翻译开始>
// 获取接口的值，而不是指针。
# <翻译结束>


<原文开始>
// Get the value by calling Elem().
<原文结束>

# <翻译开始>
// 通过调用 Elem() 获取值。
# <翻译结束>


<原文开始>
// Go through each field of the struct and copy it.
<原文结束>

# <翻译开始>
// 遍历结构体中的每个字段并进行复制。
# <翻译结束>


<原文开始>
			// The Type's StructField for a given field is checked to see if StructField.PkgPath
			// is set to determine if the field is exported or not because CanSet() returns false
			// for settable fields.  I'm not sure why.  -mohae
<原文结束>

# <翻译开始>
// 对于给定的字段，检查Type的StructField以查看StructField.PkgPath是否已设置，从而确定该字段是否为导出字段。这是因为CanSet()对于可设置字段会返回false，我不确定具体原因。——mohae
# <翻译结束>


<原文开始>
// Make a new slice and copy each element.
<原文结束>

# <翻译开始>
// 创建一个新的切片并复制每个元素。
# <翻译结束>


<原文开始>
// Copy by type assertion.
<原文结束>

# <翻译开始>
// 通过类型断言进行复制。
# <翻译结束>

