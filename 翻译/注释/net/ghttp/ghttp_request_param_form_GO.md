
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
// SetForm sets custom form value with key-value pairs.
<原文结束>

# <翻译开始>
// SetForm 设置自定义表单值，通过键值对的方式。
# <翻译结束>


<原文开始>
// GetForm retrieves and returns parameter `key` from form.
// It returns `def` if `key` does not exist in the form and `def` is given, or else it returns nil.
<原文结束>

# <翻译开始>
// GetForm 从表单中检索并返回参数 `key`。
// 如果 `key` 在表单中不存在，且提供了 `def`，则返回 `def`，否则返回 nil。
# <翻译结束>


<原文开始>
// GetFormMap retrieves and returns all form parameters passed from client as map.
// The parameter `kvMap` specifies the keys retrieving from client parameters,
// the associated values are the default values if the client does not pass.
<原文结束>

# <翻译开始>
// GetFormMap 从客户端获取并返回所有表单参数，以map形式返回。
// 参数 `kvMap` 指定了从客户端参数中检索的键，如果客户端未传递，则关联的值是默认值。
# <翻译结束>


<原文开始>
// GetFormMapStrStr retrieves and returns all form parameters passed from client as map[string]string.
// The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values
// are the default values if the client does not pass.
<原文结束>

# <翻译开始>
// GetFormMapStrStr 从客户端获取并返回所有以 map[string]string 形式传递的表单参数。
// 参数 `kvMap` 指定了从客户端参数中检索的键，如果客户端未传递，则关联的值是默认值。
# <翻译结束>


<原文开始>
// GetFormMapStrVar retrieves and returns all form parameters passed from client as map[string]*gvar.Var.
// The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values
// are the default values if the client does not pass.
<原文结束>

# <翻译开始>
// GetFormMapStrVar 从客户端获取并返回所有以map[string]*gvar.Var形式传递的表单参数。
// 参数`kvMap`指定了从客户端参数中检索的键，如果客户端未传递，则关联的值是默认值。
# <翻译结束>


<原文开始>
// GetFormStruct retrieves all form parameters passed from client and converts them to
// given struct object. Note that the parameter `pointer` is a pointer to the struct object.
// The optional parameter `mapping` is used to specify the key to attribute mapping.
<原文结束>

# <翻译开始>
// GetFormStruct 从客户端获取所有传递的表单参数，并将其转换为给定的结构体对象。
// 注意，参数 `pointer` 是指向结构体对象的指针。
// 可选参数 `mapping` 用于指定键到属性的映射关系。
# <翻译结束>

