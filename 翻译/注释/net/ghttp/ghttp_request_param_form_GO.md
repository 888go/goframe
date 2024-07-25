
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
// SetForm sets custom form value with key-value pairs.
<原文结束>

# <翻译开始>
// SetForm 使用键值对设置自定义表单值。 md5:eca1a8c094c9ff19
# <翻译结束>


<原文开始>
// GetForm retrieves and returns parameter `key` from form.
// It returns `def` if `key` does not exist in the form and `def` is given, or else it returns nil.
<原文结束>

# <翻译开始>
// GetForm 从表单中检索并返回键为 `key` 的参数。如果表单中不存在 `key`，并且提供了默认值 `def`，则返回 `def`；否则返回 `nil`。
// md5:f4a13744025f01b8
# <翻译结束>


<原文开始>
// GetFormMap retrieves and returns all form parameters passed from client as map.
// The parameter `kvMap` specifies the keys retrieving from client parameters,
// the associated values are the default values if the client does not pass.
<原文结束>

# <翻译开始>
// GetFormMap 从客户端获取并返回所有的表单参数，以map形式。参数`kvMap`指定了从客户端参数中检索的键，如果客户端未传递，则关联的值为默认值。
// md5:bc80893a54c1e60c
# <翻译结束>


<原文开始>
// GetFormMapStrStr retrieves and returns all form parameters passed from client as map[string]string.
// The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values
// are the default values if the client does not pass.
<原文结束>

# <翻译开始>
// GetFormMapStrStr 获取并以map[string]string的形式返回客户端传递的所有表单参数。
// 参数 `kvMap` 指定了从客户端参数中提取的键，如果客户端未传递，则关联的值是默认值。
// md5:09a548d91ee42cff
# <翻译结束>


<原文开始>
// GetFormMapStrVar retrieves and returns all form parameters passed from client as map[string]*gvar.Var.
// The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values
// are the default values if the client does not pass.
<原文结束>

# <翻译开始>
// GetFormMapStrVar 从客户端传递的所有表单参数中获取并返回一个 map[string]*gvar.Var。
// 参数 `kvMap` 指定了要从客户端参数中检索的键，对应的值是如果客户端未传递时的默认值。
// md5:0e9cf1899de0705b
# <翻译结束>


<原文开始>
// GetFormStruct retrieves all form parameters passed from client and converts them to
// given struct object. Note that the parameter `pointer` is a pointer to the struct object.
// The optional parameter `mapping` is used to specify the key to attribute mapping.
<原文结束>

# <翻译开始>
// GetFormStruct 从客户端获取所有传递的表单参数，并将它们转换为给定的结构体对象。需要注意的是，参数 `pointer` 是指向结构体对象的指针。可选参数 `mapping` 用于指定键到属性的映射。
// md5:36ac7f24ad6e766e
# <翻译结束>

