
<原文开始>
// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457
# <翻译结束>


<原文开始>
// Package deepcopy makes deep copies of things using reflection.
//
// This package is maintained from: https://github.com/mohae/deepcopy
<原文结束>

# <翻译开始>
// 包deepcopy使用反射来创建对象的深拷贝。
//
// 此包的维护源代码来自：https://github.com/mohae/deepcopy
// md5:5b2714e1c20589de
# <翻译结束>


<原文开始>
// Interface for delegating copy process to type
<原文结束>

# <翻译开始>
// 用于将复制过程委派给类型的接口. md5:1c1a67064d6703f2
# <翻译结束>


<原文开始>
// Copy creates a deep copy of whatever is passed to it and returns the copy
// in an interface{}.  The returned value will need to be asserted to the
// correct type.
<原文结束>

# <翻译开始>
// Copy 创建一个深度复制的副本，将传递给它的内容返回为一个接口{}。返回的值需要进行类型断言以获取正确的类型。
// md5:e0d2ca231ef6877f
# <翻译结束>


<原文开始>
// Copy by type assertion.
<原文结束>

# <翻译开始>
// 通过类型断言进行复制。 md5:71fbe824c64b5731
# <翻译结束>


<原文开始>
// Make the interface a reflect.Value
<原文结束>

# <翻译开始>
// 将接口转换为reflect.Value类型. md5:9201337756a02f33
# <翻译结束>


<原文开始>
// Make a copy of the same type as the original.
<原文结束>

# <翻译开始>
// 创建与原始类型相同的副本。 md5:5b5c1d4fae5a0eff
# <翻译结束>


<原文开始>
// Recursively copy the original.
<原文结束>

# <翻译开始>
// 递归地复制原始内容。 md5:2beee6938a2c312f
# <翻译结束>


<原文开始>
// Return the copy as an interface.
<原文结束>

# <翻译开始>
// 返回副本作为接口。 md5:50f142c98dfe40dc
# <翻译结束>


<原文开始>
// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
<原文结束>

# <翻译开始>
// copyRecursive 实现了接口的实际复制。目前，它支持的类型有限，根据需要添加更多的支持。
// md5:aa40c1acbba074ce
# <翻译结束>


<原文开始>
// check for implement deepcopy.Interface
<原文结束>

# <翻译开始>
// 检查是否实现了deepcopy.Interface接口. md5:52d685857bbf7b9b
# <翻译结束>


<原文开始>
// handle according to original's Kind
<原文结束>

# <翻译开始>
// 根据原始Kind进行处理. md5:7aba57c6e2dbe8f3
# <翻译结束>


<原文开始>
// Get the actual value being pointed to.
<原文结束>

# <翻译开始>
// 获取指针实际指向的值。 md5:89e7bbb6f609f50b
# <翻译结束>


<原文开始>
// if  it isn't valid, return.
<原文结束>

# <翻译开始>
// 如果它无效，就返回。 md5:925afc12cf8e9320
# <翻译结束>


<原文开始>
// If this is a nil, don't do anything
<原文结束>

# <翻译开始>
// 如果这是一个空值（nil），则什么都不做. md5:35928cbf0bf1c8f7
# <翻译结束>


<原文开始>
// Get the value for the interface, not the pointer.
<原文结束>

# <翻译开始>
// 获取接口的值，而不是指针的值。 md5:b5cb1b00bd69a260
# <翻译结束>


<原文开始>
// Get the value by calling Elem().
<原文结束>

# <翻译开始>
// 通过调用Elem()获取值。 md5:bf65094ef26fb870
# <翻译结束>


<原文开始>
// Go through each field of the struct and copy it.
<原文结束>

# <翻译开始>
// 遍历结构体的每个字段并复制它。 md5:9aa258862f5c519c
# <翻译结束>


<原文开始>
			// The Type's StructField for a given field is checked to see if StructField.PkgPath
			// is set to determine if the field is exported or not because CanSet() returns false
			// for settable fields.  I'm not sure why.  -mohae
<原文结束>

# <翻译开始>
			// 对于给定字段，检查Type的StructField，看StructField.PkgPath是否设置，以确定字段是否已导出。因为CanSet()方法对可设置的字段返回false，我不确定原因是什么。-mohae
			// md5:3b6525bda8d105bc
# <翻译结束>


<原文开始>
// Make a new slice and copy each element.
<原文结束>

# <翻译开始>
// 创建一个新的切片，并复制每个元素。 md5:5c3c5bd2aaf76a00
# <翻译结束>

