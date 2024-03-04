
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
// Get send GET request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Get 发送GET请求并返回响应对象。
// 注意：如果响应对象将永远不会被使用，那么它必须被关闭。
# <翻译结束>


<原文开始>
// Put send PUT request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Put 发送PUT请求并返回响应对象。
// 注意：如果响应对象将永不被使用，则必须关闭它。
# <翻译结束>


<原文开始>
// Post sends request using HTTP method POST and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Post 使用HTTP方法POST发送请求，并返回响应对象。
// 注意，如果响应对象将永远不会被使用，则必须关闭它。
# <翻译结束>


<原文开始>
// Delete send DELETE request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Delete 发送DELETE请求并返回响应对象。
// 注意：如果响应对象将永远不会被使用，那么它必须被关闭。
# <翻译结束>


<原文开始>
// Head send HEAD request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Head 发送HEAD请求并返回响应对象。
// 注意，如果响应对象将永远不会被使用，则必须关闭它。
# <翻译结束>


<原文开始>
// Patch send PATCH request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Patch 发送 PATCH 请求并返回响应对象。
// 注意：如果响应对象将永不被使用，那么它必须被关闭。
# <翻译结束>


<原文开始>
// Connect send CONNECT request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Connect 发送 CONNECT 请求并返回响应对象。
// 注意：如果响应对象将永不被使用，则必须关闭它。
# <翻译结束>


<原文开始>
// Options send OPTIONS request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Options 发送 OPTIONS 请求并返回响应对象。
// 注意，如果响应对象将永远不会被使用，则必须关闭它。
# <翻译结束>


<原文开始>
// Trace send TRACE request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Trace 发送 TRACE 请求并返回响应对象。
// 注意：如果响应对象将永远不会被使用，则必须关闭它。
# <翻译结束>


<原文开始>
// PostForm is different from net/http.PostForm.
// It's a wrapper of Post method, which sets the Content-Type as "multipart/form-data;".
// and It will automatically set boundary characters for the request body and Content-Type.
//
// It's Seem like the following case:
//
// Content-Type: multipart/form-data; boundary=----Boundarye4Ghaog6giyQ9ncN
//
// And form data is like:
// ------Boundarye4Ghaog6giyQ9ncN
// Content-Disposition: form-data; name="checkType"
//
// none
//
// It's used for sending form data.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// PostForm 与 net/http.PostForm 不同。
// 它是 Post 方法的一个包装器，会将 Content-Type 设置为 "multipart/form-data;"。
// 并且它会自动为请求体和 Content-Type 设置边界字符。
//
// 其效果类似于以下情况：
//
// Content-Type: multipart/form-data; boundary=----Boundarye4Ghaog6giyQ9ncN
//
// 表单数据格式如下：
// ------Boundarye4Ghaog6giyQ9ncN
// Content-Disposition: form-data; name="checkType"
//
// none
//
// 该方法用于发送表单数据。
// 注意，如果响应对象不再使用，则必须关闭它。
# <翻译结束>


<原文开始>
// DoRequest sends request with given HTTP method and data and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
//
// Note that it uses "multipart/form-data" as its Content-Type if it contains file uploading,
// else it uses "application/x-www-form-urlencoded". It also automatically detects the post
// content for JSON format, and for that it automatically sets the Content-Type as
// "application/json".
<原文结束>

# <翻译开始>
// DoRequest 使用给定的HTTP方法和数据发送请求，并返回响应对象。
// 注意，如果响应对象将不再使用，则必须关闭它。
//
// 注意，如果请求包含文件上传，则它使用"multipart/form-data"作为其Content-Type，
// 否则使用"application/x-www-form-urlencoded"。它还会自动检测POST内容的JSON格式，
// 对于JSON格式的内容，会自动将其Content-Type设置为"application/json"。
# <翻译结束>







<原文开始>
// prepareRequest verifies request parameters, builds and returns http request.
<原文结束>

# <翻译开始>
// 准备请求：验证请求参数，构建并返回HTTP请求。
# <翻译结束>


<原文开始>
				// It appends the parameters to the url
				// if http method is GET and Content-Type is not specified.
<原文结束>

# <翻译开始>
// 如果HTTP方法为GET且未指定Content-Type，则将参数追加到URL中。
# <翻译结束>







<原文开始>
			// Close finishes the multipart message and writes the trailing
			// boundary end line to the output.
<原文结束>

# <翻译开始>
// Close 结束多部分消息，并将尾部边界结束行写入输出。
# <翻译结束>












<原文开始>
// Auto-detecting and setting the post content format: JSON.
<原文结束>

# <翻译开始>
// 自动检测并设置帖子内容格式：JSON。
# <翻译结束>


<原文开始>
// If the parameters passed like "name=value", it then uses form type.
<原文结束>

# <翻译开始>
// 如果传入的参数类似 "name=value" 形式，则使用表单类型。
# <翻译结束>


<原文开始>
	// It's necessary set the req.Host if you want to custom the host value of the request.
	// It uses the "Host" value from header if it's not empty.
<原文结束>

# <翻译开始>
// 如果你想自定义请求的host值，设置req.Host是必要的。
// 如果header中的"Host"值不为空，则会使用该值。
# <翻译结束>







<原文开始>
// callRequest sends request with give http.Request, and returns the responses object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// callRequest 函数使用给定的 http.Request 发送请求，并返回响应对象。
// 注意：如果响应对象将永不被使用，则必须关闭该响应对象。
# <翻译结束>


<原文开始>
	// Dump feature.
	// The request body can be reused for dumping
	// raw HTTP request-response procedure.
<原文结束>

# <翻译开始>
// Dump 功能.
// 请求体可用于转储
// 原始 HTTP 请求-响应过程.
# <翻译结束>


<原文开始>
// The response might not be nil when err != nil.
<原文结束>

# <翻译开始>
// 当err不为nil时，response可能也不会为nil。
# <翻译结束>







<原文开始>
// Client middleware.
<原文结束>

# <翻译开始>
// 客户端中间件。
# <翻译结束>


<原文开始>
// File uploading request.
<原文结束>

# <翻译开始>
// 文件上传请求。
# <翻译结束>


<原文开始>
// Custom Content-Type.
<原文结束>

# <翻译开始>
// 自定义内容类型。
# <翻译结束>


<原文开始>
// HTTP basic authentication.
<原文结束>

# <翻译开始>
// HTTP基础认证
# <翻译结束>

