
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
// SetQuery sets custom query value with key-value pairs.
<原文结束>

# <翻译开始>
// SetQuery 用于设置自定义查询值，通过键值对的方式。
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
// GetQuery 从查询字符串和请求体中获取并返回指定名称`key`的参数。如果`key`在查询中不存在且提供了`def`，则返回`def`；否则返回nil。
//
// 注意，如果有多个同名参数，将以优先级顺序获取并覆盖：query > body。
// 从GET方式传递过来的参数，包括Query String及Body参数解析。
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
// GetQueryMap 从客户端通过 HTTP GET 方法传递的所有参数中获取并以 map 形式返回。参数 `kvMap` 指定了要从客户端参数中检索的键，关联的值是如果客户端未传递时的默认值。
//
// 注意，如果有多个同名参数，则按照 query > body 的优先级顺序获取并覆盖这些参数。
// 从GET方式传递过来的参数，包括Query String及Body参数解析。
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
// GetQueryMapStrStr 从客户端通过HTTP GET方法获取并返回所有传递的参数，以
// map[string]string 的形式。参数 `kvMap` 指定了要从客户端参数中检索的键，
// 关联的值是如果客户端未传递时的默认值。
// 从GET方式传递过来的参数，包括Query String及Body参数解析。
# <翻译结束>


<原文开始>
// GetQueryMapStrVar retrieves and returns all parameters passed from the client using the HTTP GET method
// as map[string]*gvar.Var. The parameter `kvMap` specifies the keys
// retrieving from client parameters, the associated values are the default values if the client
// does not pass.
<原文结束>

# <翻译开始>
// GetQueryMapStrVar 从客户端通过HTTP GET方法获取并返回所有传递的参数，以map[string]*gvar.Var的形式。参数`kvMap`指定了从客户端参数中检索的键，如果客户端未传递，则关联值为默认值。
// 从GET方式传递过来的参数，包括Query String及Body参数解析。
# <翻译结束>


<原文开始>
// GetQueryStruct retrieves all parameters passed from the client using the HTTP GET method
// and converts them to a given struct object. Note that the parameter `pointer` is a pointer
// to the struct object. The optional parameter `mapping` is used to specify the key to
// attribute mapping.
<原文结束>

# <翻译开始>
// GetQueryStruct 通过HTTP GET方法获取客户端传递的所有参数，并将它们转换为给定的结构体对象。注意，参数`pointer`是指向该结构体对象的指针。
// 可选参数`mapping`用于指定键到属性的映射关系。
// 从GET方式传递过来的参数，包括Query String及Body参数解析。
# <翻译结束>

