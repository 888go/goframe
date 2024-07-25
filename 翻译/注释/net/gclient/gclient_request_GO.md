
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
// Get send GET request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Get 发送GET请求并返回响应对象。
// 请注意，如果响应对象将永远不会被使用，必须关闭它。 md5:bf82e1e2c38506f6
# <翻译结束>


<原文开始>
// Put send PUT request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// 发送PUT请求并返回响应对象。
// 注意，如果响应对象将永远不会使用，必须关闭它。 md5:44e5f3e5edebbb91
# <翻译结束>


<原文开始>
// Post sends request using HTTP method POST and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Post 使用HTTP POST方法发送请求，并返回响应对象。
// 注意，如果响应对象将永远不会使用，必须关闭它。 md5:9ba8d1283ba032cb
# <翻译结束>


<原文开始>
// Delete send DELETE request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Delete 发送DELETE请求并返回响应对象。
// 注意，如果响应对象将永远不会被使用，那么必须关闭它。 md5:4dde007718fff7a6
# <翻译结束>


<原文开始>
// Head send HEAD request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Head 发送HEAD请求并返回响应对象。
// 请注意，如果响应对象不会被使用，必须关闭它。 md5:400dd3a80c3a0ccb
# <翻译结束>


<原文开始>
// Patch send PATCH request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Patch 发送PATCH请求并返回响应对象。
// 注意，如果响应对象将永远不会使用，必须关闭它。 md5:4e530560a87457a1
# <翻译结束>


<原文开始>
// Connect send CONNECT request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Connect 发送CONNECT请求并返回响应对象。
// 注意，如果响应对象永远不会使用，必须关闭它。 md5:cb5555f2c2a7a29d
# <翻译结束>


<原文开始>
// Options send OPTIONS request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Options 发送OPTIONS请求并返回响应对象。
// 注意，如果响应对象将永远不会被使用，那么必须关闭它。 md5:3a2d4fbe5e9f5e31
# <翻译结束>


<原文开始>
// Trace send TRACE request and returns the response object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// Trace 发送TRACE请求并返回响应对象。
// 请注意，如果响应对象将永远不会被使用，必须关闭它。 md5:82834b110d843156
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
// PostForm 与 net/http.PostForm 不同。它是一个 Post 方法的包装器，会将 Content-Type 设置为 "multipart/form-data;"。
// 它会自动为请求体和 Content-Type 设置边界字符。
//
// 类似于下面的情况：
//
// Content-Type: multipart/form-data; boundary=----Boundarye4Ghaog6giyQ9ncN
//
// 表单数据如下：
// ------Boundarye4Ghaog6giyQ9ncN
// Content-Disposition: form-data; name="checkType"
//
// none
//
// 它用于发送表单数据。请注意，如果响应对象永远不会使用，必须关闭它。 md5:bd2237aaca8f2a89
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
// DoRequest 发送给定HTTP方法和数据的请求，并返回响应对象。
// 注意，如果响应对象永远不会使用，必须关闭它。
//
// 如果包含文件上传，它将使用"multipart/form-data"作为Content-Type，否则使用"application/x-www-form-urlencoded"。它还能自动检测POST内容的JSON格式，并为此自动设置Content-Type为"application/json"。 md5:09c1fd65446e9a2e
# <翻译结束>


<原文开始>
// prepareRequest verifies request parameters, builds and returns http request.
<原文结束>

# <翻译开始>
// prepareRequest 验证请求参数，构建并返回http请求。 md5:e955238a4d45cf59
# <翻译结束>


<原文开始>
				// It appends the parameters to the url
				// if http method is GET and Content-Type is not specified.
<原文结束>

# <翻译开始>
				// 如果HTTP方法为GET且未指定Content-Type时，它将参数追加到URL中。 md5:a6325a5bd7f8b355
# <翻译结束>


<原文开始>
// File uploading request.
<原文结束>

# <翻译开始>
			// 上传文件请求。 md5:7975fe0b1475ea53
# <翻译结束>


<原文开始>
			// Close finishes the multipart message and writes the trailing
			// boundary end line to the output.
<原文结束>

# <翻译开始>
			// Close完成multipart消息并写入尾部边界结束行到输出。 md5:0a89f03d075fd9ee
# <翻译结束>


<原文开始>
// Auto-detecting and setting the post content format: JSON.
<原文结束>

# <翻译开始>
						// 自动检测并设置帖子内容格式：JSON。 md5:735d9fcd3200585a
# <翻译结束>


<原文开始>
// If the parameters passed like "name=value", it then uses form type.
<原文结束>

# <翻译开始>
						// 如果传递的参数形如 "name=value"，则使用表单类型。 md5:2f5188c0993569a1
# <翻译结束>


<原文开始>
	// It's necessary set the req.Host if you want to custom the host value of the request.
	// It uses the "Host" value from header if it's not empty.
<原文结束>

# <翻译开始>
	// 如果你想要自定义请求的主机值，那么设置 `req.Host` 是必要的。
	// 如果 `Host` 头部不为空，它会使用头部的 "Host" 值。 md5:e71cb70a52453d4c
# <翻译结束>


<原文开始>
// HTTP basic authentication.
<原文结束>

# <翻译开始>
	// HTTP基本身份验证。 md5:f6fdca448f00def0
# <翻译结束>


<原文开始>
// callRequest sends request with give http.Request, and returns the responses object.
// Note that the response object MUST be closed if it'll never be used.
<原文结束>

# <翻译开始>
// callRequest 使用给定的http.Request发送请求，并返回响应对象。
// 注意，如果响应对象将永远不会被使用，那么必须关闭它。 md5:d6e9d1e1953e082b
# <翻译结束>


<原文开始>
	// Dump feature.
	// The request body can be reused for dumping
	// raw HTTP request-response procedure.
<原文结束>

# <翻译开始>
	// Dump 功能。
	// 请求体可以用于转储
	// 原始HTTP请求-响应过程。 md5:57f6d6cec0adad22
# <翻译结束>


<原文开始>
// The response might not be nil when err != nil.
<原文结束>

# <翻译开始>
			// 当err不为nil时，response可能不为nil。 md5:30e2b1a262fbd8ac
# <翻译结束>

