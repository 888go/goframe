
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。 md5:a114f4bdd106ab31
# <翻译结束>


<原文开始>
// CORSOptions is the options for CORS feature.
// See https://www.w3.org/TR/cors/ .
<原文结束>

# <翻译开始>
// CORSOptions 是 CORS 功能的选项。
// 参见 https://www.w3.org/TR/cors/ 。 md5:86678849c932cd8d
# <翻译结束>


<原文开始>
// Used for allowing requests from custom domains
<原文结束>

# <翻译开始>
// 用于允许来自自定义域名的请求. md5:3050713aeb6de06c
# <翻译结束>


<原文开始>
// Access-Control-Allow-Origin
<原文结束>

# <翻译开始>
// Access-Control-Allow-Origin 是一个HTTP响应头字段，用于指定浏览器在进行跨域请求时可以访问资源的源（Origin）。 md5:64b1bac364c85a72
# <翻译结束>


<原文开始>
// Access-Control-Allow-Credentials
<原文结束>

# <翻译开始>
// 允许凭证访问控制. md5:9413744affc62151
# <翻译结束>


<原文开始>
// Access-Control-Expose-Headers
<原文结束>

# <翻译开始>
// Access-Control-Expose-Headers：这是HTTP响应头的一部分，用于指定哪些响应头可以被CORS（跨源资源共享）策略允许从服务器传递到客户端。 md5:edb32baedc37800f
# <翻译结束>


<原文开始>
// Access-Control-Allow-Methods
<原文结束>

# <翻译开始>
// 跨域资源共享允许的方法. md5:c78ddd1745514f4a
# <翻译结束>


<原文开始>
// Access-Control-Allow-Headers
<原文结束>

# <翻译开始>
// Access-Control-Allow-Headers：允许的头部字段. md5:9812fd3132d19ca6
# <翻译结束>


<原文开始>
	// defaultAllowHeaders is the default allowed headers for CORS.
	// It defined another map for better header key searching performance.
<原文结束>

# <翻译开始>
	// defaultAllowHeaders 是CORS默认允许的头信息。
	// 为了提高头部键搜索性能，我们定义了另一个映射。 md5:e6a13ea98879b3e6
# <翻译结束>


<原文开始>
// DefaultCORSOptions returns the default CORS options,
// which allows any cross-domain request.
<原文结束>

# <翻译开始>
// DefaultCORSOptions 返回默认的 CORS 选项，它允许任何跨域请求。 md5:ed45ce5e88088eac
# <翻译结束>


<原文开始>
// Allow all client's custom headers in default.
<原文结束>

# <翻译开始>
	// 默认允许客户端的所有自定义头部。 md5:5aa0a6d974ed81b6
# <翻译结束>


<原文开始>
// Allow all anywhere origin in default.
<原文结束>

# <翻译开始>
	// 默认允许所有来源的访问。 md5:bd5e36856694e82f
# <翻译结束>


<原文开始>
// CORS sets custom CORS options.
// See https://www.w3.org/TR/cors/ .
<原文结束>

# <翻译开始>
// CORS 设置自定义CORS选项。
// 参见 https://www.w3.org/TR/cors/ 。 md5:5ace1c84086a260a
# <翻译结束>


<原文开始>
	// No continue service handling if it's OPTIONS request.
	// Note that there's special checks in previous router searching,
	// so if it goes to here it means there's already serving handler exist.
<原文结束>

# <翻译开始>
	// 如果请求是OPTIONS类型，不继续服务处理。
	// 注意，之前的路由器搜索中已经有特殊检查，
	// 所以如果到达这里，意味着已经存在正在处理的处理器。 md5:178e6bee651f512f
# <翻译结束>


<原文开始>
// CORSAllowedOrigin CORSAllowed checks whether the current request origin is allowed cross-domain.
<原文结束>

# <翻译开始>
// CORSAllowedOrigin CORSAllowed 检查当前请求的来源是否被允许进行跨域。 md5:599a140b617c5c1c
# <翻译结束>


<原文开始>
// CORSDefault sets CORS with default CORS options,
// which allows any cross-domain request.
<原文结束>

# <翻译开始>
// CORSDefault 使用默认的 CORS 选项设置 CORS，
// 允许任何跨域请求。 md5:2808119e534c338a
# <翻译结束>

