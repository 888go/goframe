
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
// SetQuery sets custom query value with key-value pairs.
<原文结束>

# <翻译开始>
// SetQuery 使用键值对设置自定义查询值。 md5:464e6b634ef97c90
# <翻译结束>


<原文开始>
// GetQuery retrieves and return parameter with the given name `key` from query string
// and request body. It returns `def` if `key` does not exist in the query and `def` is given,
// or else it returns nil.
//
// Note that if there are multiple parameters with the same name, the parameters are retrieved
// and overwrote in order of priority: query > body.
<原文结束>

# <翻译开始>
// GetQuery 从查询字符串和请求体中检索并返回给定名称 `key` 的参数。如果 `key` 不在查询中并且提供了 `def`，则返回 `def`；否则返回 nil。
//
// 注意，如果有多个同名的参数，将按照优先级顺序进行检索和覆盖：查询参数 > 身体参数。 md5:3948868b7e507e93
# <翻译结束>


<原文开始>
// GetQueryMap retrieves and returns all parameters passed from the client using HTTP GET method
// as the map. The parameter `kvMap` specifies the keys retrieving from client parameters,
// the associated values are the default values if the client does not pass.
//
// Note that if there are multiple parameters with the same name, the parameters are retrieved and overwrote
// in order of priority: query > body.
<原文结束>

# <翻译开始>
// GetQueryMap 从客户端通过HTTP GET方法传递的所有参数中检索并返回它们作为映射。参数 `kvMap` 指定了从客户端参数中获取的键，如果客户端未提供，则关联的值为默认值。
//
// 注意，如果有多个具有相同名称的参数，将按照优先级顺序检索和覆盖：查询参数 > 身体（请求体）参数。 md5:72471cd6457be5f2
# <翻译结束>


<原文开始>
// GetQueryMapStrStr retrieves and returns all parameters passed from the client using the HTTP GET method as a
//
//	map[string]string. The parameter `kvMap` specifies the keys
//
// retrieving from client parameters, the associated values are the default values if the client
// does not pass.
<原文结束>

# <翻译开始>
// GetQueryMapStrStr 获取并返回所有通过HTTP GET方法从客户端传递过来的参数，作为一个
//
// map[string]string。参数 `kvMap` 指定了从客户端参数中提取的键
//
// ，关联的值是如果客户端没有传递时的默认值。 md5:b1d5d46b8cc53f3a
# <翻译结束>


<原文开始>
// GetQueryMapStrVar retrieves and returns all parameters passed from the client using the HTTP GET method
// as map[string]*gvar.Var. The parameter `kvMap` specifies the keys
// retrieving from client parameters, the associated values are the default values if the client
// does not pass.
<原文结束>

# <翻译开始>
// GetQueryMapStrVar 从使用 HTTP GET 方法传递的客户端参数中获取并返回所有参数，形式为 map[string]*gvar.Var。参数 `kvMap` 指定了要从客户端参数中获取的键，对应的值是如果客户端未传递时的默认值。 md5:3db7496b4b165e99
# <翻译结束>


<原文开始>
// GetQueryStruct retrieves all parameters passed from the client using the HTTP GET method
// and converts them to a given struct object. Note that the parameter `pointer` is a pointer
// to the struct object. The optional parameter `mapping` is used to specify the key to
// attribute mapping.
<原文结束>

# <翻译开始>
// GetQueryStruct 从客户端通过HTTP GET方法获取所有传递的参数，并将它们转换为给定的结构体对象。请注意，参数`pointer`是指向结构体对象的指针。可选参数`mapping`用于指定键到属性的映射。 md5:7061a83f935b7317
# <翻译结束>

