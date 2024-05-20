
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
// GetRequest 从客户端获取并返回名为 `key` 的参数，以及作为接口传递的自定义参数。无论客户端使用何种HTTP方法。如果`key`不存在，`def`参数指定了默认值。
// 
// GetRequest 是最常用的用于检索参数的函数之一。
// 
// 注意，如果有多个同名参数，将按照以下优先级顺序进行获取和覆盖：路由器 < 查询参数 < 身份验证 < 表单数据 < 自定义参数。
// md5:a008e7f428967448
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
// GetRequestMap 从客户端获取并返回所有传递的参数及自定义参数，无论客户端使用的是哪种HTTP方法。参数 `kvMap` 指定了从客户端参数中提取的键，关联的值是在客户端未传递相应键时的默认值。
//
// GetRequestMap 是最常用于检索参数的函数之一。
//
// 注意，如果有多个同名参数，参数将按照优先级顺序被获取及覆盖：路由参数 < 查询参数 < 请求体参数 < 表单参数 < 自定义参数。
// md5:b01ba4caf2092f12
# <翻译结束>


<原文开始>
// Check none exist parameters and assign it with default value.
<原文结束>

# <翻译开始>
// 检查不存在的参数并为其分配默认值。. md5:2c9c16dac85c432c
# <翻译结束>


<原文开始>
// GetRequestMapStrStr retrieve and returns all parameters passed from the client and custom
// params as map[string]string, no matter what HTTP method the client is using. The parameter
// `kvMap` specifies the keys retrieving from client parameters, the associated values are the
// default values if the client does not pass.
<原文结束>

# <翻译开始>
// GetRequestMapStrStr 从客户端和自定义参数中获取并返回所有传递的参数，无论客户端使用何种HTTP方法。参数`kvMap`指定了从客户端参数中检索的键，关联的值是客户端未传递时的默认值。
// md5:18e353330403d45b
# <翻译结束>


<原文开始>
// GetRequestMapStrVar retrieve and returns all parameters passed from the client and custom
// params as map[string]*gvar.Var, no matter what HTTP method the client is using. The parameter
// `kvMap` specifies the keys retrieving from client parameters, the associated values are the
// default values if the client does not pass.
<原文结束>

# <翻译开始>
// GetRequestMapStrVar 从客户端和自定义参数中检索并返回所有传递的参数，作为map[string]*gvar.Var。无论客户端使用何种HTTP方法，都会进行检索。参数`kvMap`指定了从客户端参数中获取的键，关联的值是客户端未传递时的默认值。
// md5:1063c291381a5048
# <翻译结束>


<原文开始>
// GetRequestStruct retrieves all parameters passed from the client and custom params no matter
// what HTTP method the client is using, and converts them to give the struct object. Note that
// the parameter `pointer` is a pointer to the struct object.
// The optional parameter `mapping` is used to specify the key to attribute mapping.
<原文结束>

# <翻译开始>
// GetRequestStruct 无论客户端使用何种HTTP方法，都会获取客户端传递的所有参数和自定义参数，
// 并将它们转换为结构体对象。注意，参数`pointer`是一个指向结构体对象的指针。
// 可选参数`mapping`用于指定键到属性的映射。
// md5:a117b2c0722fc3fe
# <翻译结束>


<原文开始>
// `in` Tag Struct values.
<原文结束>

# <翻译开始>
// `in` 标签结构体值。. md5:225b15f233b09df1
# <翻译结束>


<原文开始>
// mergeDefaultStructValue merges the request parameters with default values from struct tag definition.
<原文结束>

# <翻译开始>
// mergeDefaultStructValue 将请求参数与结构体标签定义中的默认值合并。. md5:0a73ebb7f647201a
# <翻译结束>


<原文开始>
// provide non strict routing
<原文结束>

# <翻译开始>
// 提供非严格的路由. md5:c3f73d5de1159867
# <翻译结束>


<原文开始>
// mergeInTagStructValue merges the request parameters with header or cookie values from struct `in` tag definition.
<原文结束>

# <翻译开始>
// mergeInTagStructValue 将请求参数与根据结构体`in`标签定义的头或cookie值合并。. md5:a6444655a59f403d
# <翻译结束>

