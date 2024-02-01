
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
// Cookie for HTTP COOKIE management.
<原文结束>

# <翻译开始>
// Cookie 用于 HTTP COOKIE 管理。
# <翻译结束>






















<原文开始>
// CookieOptions provides security config for cookies
<原文结束>

# <翻译开始>
// CookieOptions 提供了用于 cookies 的安全配置选项
# <翻译结束>

















<原文开始>
// cookieItem is the item stored in Cookie.
<原文结束>

# <翻译开始>
// cookieItem 是存储在 Cookie 中的项目。
# <翻译结束>


<原文开始>
// Mark this cookie received from the client.
<原文结束>

# <翻译开始>
// 标记此cookie是从客户端接收到的。
# <翻译结束>


<原文开始>
// GetCookie creates or retrieves a cookie object with given request.
// It retrieves and returns an existing cookie object if it already exists with given request.
// It creates and returns a new cookie object if it does not exist with given request.
<原文结束>

# <翻译开始>
// GetCookie 函数通过给定的请求创建或检索一个 cookie 对象。
// 若已存在与给定请求相关的 cookie，则检索并返回该存在的 cookie 对象。
// 若不存在与给定请求相关的 cookie，则创建并返回一个新的 cookie 对象。
# <翻译结束>


<原文开始>
// init does lazy initialization for the cookie object.
<原文结束>

# <翻译开始>
// init 对cookie对象执行延迟初始化。
# <翻译结束>


<原文开始>
	// DO NOT ADD ANY DEFAULT COOKIE DOMAIN!
	// if c.request.Server.GetCookieDomain() == "" {
	//	c.request.Server.GetCookieDomain() = c.request.GetHost()
	// }
<原文结束>

# <翻译开始>
// **请勿添加任何默认的Cookie域！**
// 如果 c.request.Server.GetCookieDomain() 为空，即返回值为""，则执行以下操作：
// 将 c.request.Server.GetCookieDomain() 的值设置为 c.request.GetHost() 的返回值。
# <翻译结束>


<原文开始>
// Map returns the cookie items as map[string]string.
<原文结束>

# <翻译开始>
// Map 将 cookie 项以 map[string]string 的形式返回。
# <翻译结束>


<原文开始>
// Contains checks if given key exists and not expire in cookie.
<原文结束>

# <翻译开始>
// Contains 检查给定的键是否存在且在 cookie 中未过期。
# <翻译结束>


<原文开始>
// Set sets cookie item with default domain, path and expiration age.
<原文结束>

# <翻译开始>
// Set 使用默认的域名、路径和过期时间设置cookie项。
# <翻译结束>


<原文开始>
// SetCookie sets cookie item with given domain, path and expiration age.
// The optional parameter `options` specifies extra security configurations,
// which is usually empty.
<原文结束>

# <翻译开始>
// SetCookie 用于给指定的域名、路径设置cookie项，并设置其过期时间。
// 可选参数 `options` 指定了额外的安全配置，通常为空。
# <翻译结束>


<原文开始>
// SetHttpCookie sets cookie with *http.Cookie.
<原文结束>

# <翻译开始>
// SetHttpCookie 通过 *http.Cookie 设置 cookie。
# <翻译结束>


<原文开始>
// GetSessionId retrieves and returns the session id from cookie.
<原文结束>

# <翻译开始>
// GetSessionId 从cookie中获取并返回会话id。
# <翻译结束>


<原文开始>
// SetSessionId sets session id in the cookie.
<原文结束>

# <翻译开始>
// SetSessionId 将会话ID设置到cookie中。
# <翻译结束>


<原文开始>
// Get retrieves and returns the value with specified key.
// It returns `def` if specified key does not exist and `def` is given.
<原文结束>

# <翻译开始>
// Get 方法用于获取并返回指定键的值。
// 如果指定的键不存在且提供了默认值 `def`，则返回 `def`。
# <翻译结束>


<原文开始>
// Remove deletes specified key and its value from cookie using default domain and path.
// It actually tells the http client that the cookie is expired, do not send it to server next time.
<原文结束>

# <翻译开始>
// Remove 函数通过使用默认的域名和路径，从cookie中删除指定的键及其对应的值。
// 实质上，它告知http客户端该cookie已过期，下次不要将其发送到服务器。
# <翻译结束>


<原文开始>
// RemoveCookie deletes specified key and its value from cookie using given domain and path.
// It actually tells the http client that the cookie is expired, do not send it to server next time.
<原文结束>

# <翻译开始>
// RemoveCookie 通过给定的域名和路径，从cookie中删除指定的键及其对应的值。
// 实际上，它会告知http客户端该cookie已过期，下次不要将其发送到服务器。
# <翻译结束>


<原文开始>
// Flush outputs the cookie items to the client.
<原文结束>

# <翻译开始>
// Flush 将cookie项输出到客户端。
# <翻译结束>


<原文开始>
// Underlying cookie items.
<原文结束>

# <翻译开始>
// 基础的cookie项。
# <翻译结束>


<原文开始>
// Belonged HTTP server
<原文结束>

# <翻译开始>
// 所属HTTP服务器
# <翻译结束>


<原文开始>
// Belonged HTTP request.
<原文结束>

# <翻译开始>
// 属于HTTP请求。
# <翻译结束>


<原文开始>
// Belonged HTTP response.
<原文结束>

# <翻译开始>
// 属于HTTP响应。
# <翻译结束>


<原文开始>
// cookie SameSite property
<原文结束>

# <翻译开始>
// cookie 的 SameSite 属性
# <翻译结束>


<原文开始>
// cookie Secure property
<原文结束>

# <翻译开始>
// cookie Secure 属性
# <翻译结束>


<原文开始>
// cookie HttpOnly property
<原文结束>

# <翻译开始>
// cookie HttpOnly 属性
# <翻译结束>

