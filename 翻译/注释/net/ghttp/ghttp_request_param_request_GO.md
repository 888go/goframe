
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
// GetRequest retrieves and returns the parameter named `key` passed from the client and
// custom params as interface{}, no matter what HTTP method the client is using. The
// parameter `def` specifies the default value if the `key` does not exist.
//
// GetRequest is one of the most commonly used functions for retrieving parameters.
//
// Note that if there are multiple parameters with the same name, the parameters are
// retrieved and overwrote in order of priority: router < query < body < form < custom.
<原文结束>

# <翻译开始>
// GetRequest 函数用于获取并返回客户端通过任意HTTP方法传递的名为`key`的参数以及作为interface{}类型的自定义参数。参数`def`用于指定当`key`不存在时的默认值。
//
// GetRequest 是用于检索参数的最常用函数之一。
//
// 注意，可获取客户端提交的所有参数，不区分提交方式。如果有多个同名参数，按照以下优先级顺序获取并覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
# <翻译结束>


<原文开始>
// GetRequestMap retrieves and returns all parameters passed from the client and custom params
// as the map, no matter what HTTP method the client is using. The parameter `kvMap` specifies
// the keys retrieving from client parameters, the associated values are the default values
// if the client does not pass the according keys.
//
// GetRequestMap is one of the most commonly used functions for retrieving parameters.
//
// Note that if there are multiple parameters with the same name, the parameters are retrieved
// and overwrote in order of priority: router < query < body < form < custom.
<原文结束>

# <翻译开始>
// GetRequestMap 从客户端获取并返回所有传递的参数以及自定义参数，无论客户端使用何种HTTP方法。参数`kvMap`指定了从客户端参数中检索的键，关联的值是如果客户端未传递相应键时的默认值。
//
// GetRequestMap 是用于检索参数的最常用函数之一。
//
// 注意，可获取客户端提交的所有参数，不区分提交方式。如果有多个同名参数，按照以下优先级顺序获取并覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
# <翻译结束>


<原文开始>
// Check none exist parameters and assign it with default value.
<原文结束>

# <翻译开始>
// 检查不存在的参数，并赋予其默认值。
# <翻译结束>


<原文开始>
// GetRequestMapStrStr retrieve and returns all parameters passed from the client and custom
// params as map[string]string, no matter what HTTP method the client is using. The parameter
// `kvMap` specifies the keys retrieving from client parameters, the associated values are the
// default values if the client does not pass.
<原文结束>

# <翻译开始>
// GetRequestMapStrStr 从客户端获取并返回所有传递的参数以及自定义参数，无论客户端使用何种HTTP方法。
// 参数`kvMap`指定了从客户端参数中检索的键，如果客户端未传递，则关联的值是默认值。返回类型为map[string]string。
// 注意，可获取客户端提交的所有参数，不区分提交方式。如果有多个同名参数，按照以下优先级顺序获取并覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
# <翻译结束>


<原文开始>
// GetRequestMapStrVar retrieve and returns all parameters passed from the client and custom
// params as map[string]*gvar.Var, no matter what HTTP method the client is using. The parameter
// `kvMap` specifies the keys retrieving from client parameters, the associated values are the
// default values if the client does not pass.
<原文结束>

# <翻译开始>
// GetRequestMapStrVar 获取并返回客户端通过任何HTTP方法传递的所有参数，以及自定义参数，
// 并以map[string]*gvar.Var的形式返回。参数`kvMap`指定了从客户端参数中获取的键，
// 相关联的值是当客户端未传递时的默认值。
// 注意，可获取客户端提交的所有参数，不区分提交方式。如果有多个同名参数，按照以下优先级顺序获取并覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
# <翻译结束>


<原文开始>
// GetRequestStruct retrieves all parameters passed from the client and custom params no matter
// what HTTP method the client is using, and converts them to give the struct object. Note that
// the parameter `pointer` is a pointer to the struct object.
// The optional parameter `mapping` is used to specify the key to attribute mapping.
<原文结束>

# <翻译开始>
// GetRequestStruct 从客户端获取所有传递的参数以及自定义参数，无论客户端使用何种HTTP方法，
// 并将它们转换为给定的结构体对象。注意，参数`pointer`是指向结构体对象的指针。
// 可选参数`mapping`用于指定键到属性的映射关系。
// 注意，可获取客户端提交的所有参数，不区分提交方式。如果有多个同名参数，按照以下优先级顺序获取并覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
# <翻译结束>












<原文开始>
// mergeDefaultStructValue merges the request parameters with default values from struct tag definition.
<原文结束>

# <翻译开始>
// mergeDefaultStructValue 将请求参数与来自结构体标签定义的默认值进行合并。
# <翻译结束>







<原文开始>
// mergeInTagStructValue merges the request parameters with header or cookie values from struct `in` tag definition.
<原文结束>

# <翻译开始>
// mergeInTagStructValue 将请求参数与来自 `in` 标签定义的结构体中的头部或cookie值进行合并。
# <翻译结束>







<原文开始>
// Default struct values.
<原文结束>

# <翻译开始>
// 默认结构体值。
# <翻译结束>


<原文开始>
// `in` Tag Struct values.
<原文结束>

# <翻译开始>
// `in` 标签结构体值。
# <翻译结束>


<原文开始>
// provide non strict routing
<原文结束>

# <翻译开始>
// 提供非严格路由
# <翻译结束>

