
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
// Cookie for HTTP COOKIE management.
<原文结束>

# <翻译开始>
// 用于HTTP_COOKIE管理的Cookie。 md5:5990685f9449a6df
# <翻译结束>


<原文开始>
// Underlying cookie items.
<原文结束>

# <翻译开始>
// 底层cookie项。 md5:b89c165cdb180382
# <翻译结束>


<原文开始>
// Belonged HTTP response.
<原文结束>

# <翻译开始>
// 属于的HTTP响应。 md5:7081ed19a38e6318
# <翻译结束>


<原文开始>
// CookieOptions provides security config for cookies
<原文结束>

# <翻译开始>
// CookieOptions 提供了 cookie 的安全配置. md5:538f164732c20846
# <翻译结束>


<原文开始>
// cookie SameSite property
<原文结束>

# <翻译开始>
// cookie 的 SameSite 属性. md5:4365fbb8775da6c6
# <翻译结束>


<原文开始>
// cookie HttpOnly property
<原文结束>

# <翻译开始>
// Cookie 的 HttpOnly 属性. md5:77d9da6c97c6713d
# <翻译结束>


<原文开始>
// cookieItem is the item stored in Cookie.
<原文结束>

# <翻译开始>
// cookieItem 是存储在 Cookie 中的项。 md5:70e35f7ab478a99b
# <翻译结束>


<原文开始>
// Mark this cookie received from the client.
<原文结束>

# <翻译开始>
// 标记从客户端接收的此Cookie。 md5:fa3cc5dd0aefa33c
# <翻译结束>


<原文开始>
// GetCookie creates or retrieves a cookie object with given request.
// It retrieves and returns an existing cookie object if it already exists with given request.
// It creates and returns a new cookie object if it does not exist with given request.
<原文结束>

# <翻译开始>
// GetCookie 通过给定的请求创建或检索一个cookie对象。
// 如果已经存在与给定请求匹配的cookie对象，它将检索并返回该对象。
// 如果不存在与给定请求匹配的cookie对象，它将创建一个新的cookie对象并返回。 md5:5b2b3a376a2f6162
# <翻译结束>


<原文开始>
// init does lazy initialization for the cookie object.
<原文结束>

# <翻译开始>
// init 函数对cookie对象进行惰性初始化。 md5:6a771fc762ea2458
# <翻译结束>


<原文开始>
	// DO NOT ADD ANY DEFAULT COOKIE DOMAIN!
	// if c.request.Server.GetCookieDomain() == "" {
	//	c.request.Server.GetCookieDomain() = c.request.GetHost()
	// }
<原文结束>

# <翻译开始>
	// 请勿在此添加任何默认的Cookie域名！
	// 如果 c.request.Server.GetCookieDomain() 等于空字符串 {
	//     c.request.Server.GetCookieDomain() = c.request.GetHost()
	// } md5:dd77fb4cfe35c3cc
# <翻译结束>


<原文开始>
// Map returns the cookie items as map[string]string.
<原文结束>

# <翻译开始>
// Map将Cookie项作为map[string]string返回。 md5:6d265855ff6356fa
# <翻译结束>


<原文开始>
// Contains checks if given key exists and not expire in cookie.
<原文结束>

# <翻译开始>
// Contains 检查给定的键是否存在于cookie中且未过期。 md5:acb79130bbbf78e6
# <翻译结束>


<原文开始>
// Set sets cookie item with default domain, path and expiration age.
<原文结束>

# <翻译开始>
// Set 使用默认域名、路径和过期时间来设置cookie项。 md5:66be80c8c6c07dc1
# <翻译结束>


<原文开始>
// SetCookie sets cookie item with given domain, path and expiration age.
// The optional parameter `options` specifies extra security configurations,
// which is usually empty.
<原文结束>

# <翻译开始>
// SetCookie 设置具有给定域名、路径和过期时间的cookie项。
// 可选参数`options`指定了额外的安全配置，通常为空。 md5:2afc45e40597ce0d
# <翻译结束>


<原文开始>
// SetHttpCookie sets cookie with *http.Cookie.
<原文结束>

# <翻译开始>
// SetHttpCookie 使用*http.Cookie设置cookie。 md5:de525635cedd10e4
# <翻译结束>


<原文开始>
// GetSessionId retrieves and returns the session id from cookie.
<原文结束>

# <翻译开始>
// GetSessionId 从cookie中检索并返回session ID。 md5:4b41d62448c6f135
# <翻译结束>


<原文开始>
// SetSessionId sets session id in the cookie.
<原文结束>

# <翻译开始>
// SetSessionId 在cookie中设置会话ID。 md5:56899c22c78f2267
# <翻译结束>


<原文开始>
// Get retrieves and returns the value with specified key.
// It returns `def` if specified key does not exist and `def` is given.
<原文结束>

# <翻译开始>
// Get 获取并返回具有指定键的值。
// 如果指定的键不存在，并且提供了默认值`def`，则返回`def`。 md5:f137dd1311660ee4
# <翻译结束>


<原文开始>
// Remove deletes specified key and its value from cookie using default domain and path.
// It actually tells the http client that the cookie is expired, do not send it to server next time.
<原文结束>

# <翻译开始>
// Remove 从使用默认域名和路径的cookie中删除指定键及其值。实际上，它告诉HTTP客户端该cookie已过期，下次不要将其发送到服务器。 md5:4c429c6b33ce790c
# <翻译结束>


<原文开始>
// RemoveCookie deletes specified key and its value from cookie using given domain and path.
// It actually tells the http client that the cookie is expired, do not send it to server next time.
<原文结束>

# <翻译开始>
// RemoveCookie 使用给定的域名和路径从cookie中删除指定的键及其值。实际上，它告诉HTTP客户端该cookie已过期，下次不要将其发送到服务器。 md5:a437da6ed4ded047
# <翻译结束>


<原文开始>
// Flush outputs the cookie items to the client.
<原文结束>

# <翻译开始>
// Flush 将cookie项输出到客户端。 md5:77b1f132a9910559
# <翻译结束>

