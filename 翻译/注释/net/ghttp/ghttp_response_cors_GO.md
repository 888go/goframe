
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//
# <翻译结束>


<原文开始>
// CORSOptions is the options for CORS feature.
// See https://www.w3.org/TR/cors/ .
<原文结束>

# <翻译开始>
// CORSOptions 是CORS功能的选项。
// 参见 https://www.w3.org/TR/cors/ 。
# <翻译结束>


<原文开始>
// Used for allowing requests from custom domains
<原文结束>

# <翻译开始>
// 用于允许来自自定义域名的请求
# <翻译结束>
































<原文开始>
	// defaultAllowHeaders is the default allowed headers for CORS.
	// It defined another map for better header key searching performance.
<原文结束>

# <翻译开始>
// defaultAllowHeaders 是 CORS 的默认允许头信息。
// 它定义了另一个映射，以便提高查找头部键的性能。
# <翻译结束>


<原文开始>
// DefaultCORSOptions returns the default CORS options,
// which allows any cross-domain request.
<原文结束>

# <翻译开始>
// DefaultCORSOptions 返回默认的 CORS 选项，
// 这些选项允许任何跨域请求。
# <翻译结束>


<原文开始>
// Allow all client's custom headers in default.
<原文结束>

# <翻译开始>
// 默认情况下，允许所有客户端自定义头部。
# <翻译结束>


<原文开始>
// Allow all anywhere origin in default.
<原文结束>

# <翻译开始>
// 默认情况下允许所有来源的任何请求
# <翻译结束>


<原文开始>
// CORS sets custom CORS options.
// See https://www.w3.org/TR/cors/ .
<原文结束>

# <翻译开始>
// CORS 设置自定义 CORS 选项。
// 参见 https://www.w3.org/TR/cors/ 。
# <翻译结束>


<原文开始>
	// No continue service handling if it's OPTIONS request.
	// Note that there's special checks in previous router searching,
	// so if it goes to here it means there's already serving handler exist.
<原文结束>

# <翻译开始>
// 如果请求为OPTIONS，则不进行继续服务处理。
// 注意，之前在路由搜索中有特殊的检查，
// 所以如果执行到这里，意味着已存在正在服务的处理程序。
# <翻译结束>







<原文开始>
// CORSAllowedOrigin CORSAllowed checks whether the current request origin is allowed cross-domain.
<原文结束>

# <翻译开始>
// CORSAllowedOrigin CORSAllowedOrigin函数检查当前请求的来源是否允许跨域。
# <翻译结束>


<原文开始>
// CORSDefault sets CORS with default CORS options,
// which allows any cross-domain request.
<原文结束>

# <翻译开始>
// CORSDefault 使用默认CORS选项设置CORS，
// 这将允许任何跨域请求。
# <翻译结束>


<原文开始>
// Access-Control-Allow-Credentials
<原文结束>

# <翻译开始>
// Access-Control-Allow-Credentials: 允许跨域请求时携带验证凭据（cookies, Authorization 头等）
# <翻译结束>


<原文开始>
// Access-Control-Expose-Headers
<原文结束>

# <翻译开始>
// Access-Control-Expose-Headers：
# <翻译结束>


<原文开始>
// Access-Control-Allow-Methods
<原文结束>

# <翻译开始>
// Access-Control-Allow-Methods：允许跨域请求的方法列表
# <翻译结束>


<原文开始>
// Access-Control-Allow-Headers
<原文结束>

# <翻译开始>
// Access-Control-Allow-Headers: 允许跨域请求中携带的自定义请求头
# <翻译结束>


<原文开始>
// Access-Control-Allow-Origin
<原文结束>

# <翻译开始>
// Access-Control-Allow-Origin：允许跨域请求的源，这是HTTP响应头的一部分，用于指示服务器允许哪些源发起跨域请求。在Go语言中设置该响应头可以实现跨域资源共享（CORS）。
# <翻译结束>


<原文开始>
// Access-Control-Max-Age
<原文结束>

# <翻译开始>
// Access-Control-Max-Age： 
// （该注释表示HTTP响应头中的Access-Control-Max-Age字段，用于指示预检请求（OPTIONS）的结果能够被缓存多久。）
// 设置浏览器对跨域资源共享（CORS）中间结果的最大缓存时间（单位为秒），即预检请求的有效期。
// 当值设为非零时，在此期间内再次进行相同的跨域请求将不再发送预检请求，直接使用第一次预检请求的结果。
# <翻译结束>


<原文开始>
// No continue serving.
<原文结束>

# <翻译开始>
// 不再继续服务。
# <翻译结束>

