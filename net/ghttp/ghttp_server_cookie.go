// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"net/http"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
)

// 用于HTTP_COOKIE管理的Cookie。 md5:5990685f9449a6df
type Cookie struct {
	data     map[string]*cookieItem // 底层cookie项。 md5:b89c165cdb180382
	server   *Server                // Belonged HTTP server
	request  *Request               // Belonged HTTP request.
	response *Response              // 属于的HTTP响应。 md5:7081ed19a38e6318
}

// CookieOptions 提供了 cookie 的安全配置. md5:538f164732c20846
type CookieOptions struct {
	SameSite http.SameSite // cookie 的 SameSite 属性. md5:4365fbb8775da6c6
	Secure   bool          // cookie Secure property
	HttpOnly bool          // Cookie 的 HttpOnly 属性. md5:77d9da6c97c6713d
}

// cookieItem 是存储在 Cookie 中的项。 md5:70e35f7ab478a99b
type cookieItem struct {
	*http.Cookie      // 底层cookie项。 md5:b89c165cdb180382
	FromClient   bool // 标记从客户端接收的此Cookie。 md5:fa3cc5dd0aefa33c
}

// GetCookie 通过给定的请求创建或检索一个cookie对象。
// 如果已经存在与给定请求匹配的cookie对象，它将检索并返回该对象。
// 如果不存在与给定请求匹配的cookie对象，它将创建一个新的cookie对象并返回。
// md5:5b2b3a376a2f6162
// ff:取cookie对象
// r:
func GetCookie(r *Request) *Cookie {
	if r.Cookie != nil {
		return r.Cookie
	}
	return &Cookie{
		request: r,
		server:  r.Server,
	}
}

// init 函数对cookie对象进行惰性初始化。 md5:6a771fc762ea2458
func (c *Cookie) init() {
	if c.data != nil {
		return
	}
	c.data = make(map[string]*cookieItem)
	c.response = c.request.Response
	// 请勿在此添加任何默认的Cookie域名！
	// 如果 c.request.Server.GetCookieDomain() 等于空字符串 {
	//     c.request.Server.GetCookieDomain() = c.request.GetHost()
	// }
	// md5:dd77fb4cfe35c3cc
	for _, v := range c.request.Cookies() {
		c.data[v.Name] = &cookieItem{
			Cookie:     v,
			FromClient: true,
		}
	}
}

// Map将Cookie项作为map[string]string返回。 md5:6d265855ff6356fa
// ff:取Map
// c:
func (c *Cookie) Map() map[string]string {
	c.init()
	m := make(map[string]string)
	for k, v := range c.data {
		m[k] = v.Value
	}
	return m
}

// Contains 检查给定的键是否存在于cookie中且未过期。 md5:acb79130bbbf78e6
// ff:是否已过期
// c:
// key:名称
func (c *Cookie) Contains(key string) bool {
	c.init()
	if r, ok := c.data[key]; ok {
		if r.Expires.IsZero() || r.Expires.After(time.Now()) {
			return true
		}
	}
	return false
}

// Set 使用默认域名、路径和过期时间来设置cookie项。 md5:66be80c8c6c07dc1
// yx:true
// ff:设置值
// c:
// key:
// value:
func (c *Cookie) Set(key, value string) {
	c.SetCookie(
		key,
		value,
		c.request.Server.GetCookieDomain(),
		c.request.Server.GetCookiePath(),
		c.request.Server.GetCookieMaxAge(),
		CookieOptions{
			SameSite: c.request.Server.GetCookieSameSite(),
			Secure:   c.request.Server.GetCookieSecure(),
			HttpOnly: c.request.Server.GetCookieHttpOnly(),
		},
	)
}

// SetCookie 设置具有给定域名、路径和过期时间的cookie项。
// 可选参数`options`指定了额外的安全配置，通常为空。
// md5:2afc45e40597ce0d
// ff:设置cookie
// c:
// key:名称
// value:值
// domain:域名
// path:路径
// maxAge:最大存活时长
// options:安全配置项
func (c *Cookie) SetCookie(key, value, domain, path string, maxAge time.Duration, options ...CookieOptions) {
	c.init()
	config := CookieOptions{}
	if len(options) > 0 {
		config = options[0]
	}
	httpCookie := &http.Cookie{
		Name:     key,
		Value:    value,
		Path:     path,
		Domain:   domain,
		HttpOnly: config.HttpOnly,
		SameSite: config.SameSite,
		Secure:   config.Secure,
	}
	if maxAge != 0 {
		httpCookie.Expires = time.Now().Add(maxAge)
	}
	c.data[key] = &cookieItem{
		Cookie: httpCookie,
	}
}

// SetHttpCookie 使用*http.Cookie设置cookie。 md5:de525635cedd10e4
// ff:设置httpcookie
// c:
// httpCookie:
func (c *Cookie) SetHttpCookie(httpCookie *http.Cookie) {
	c.init()
	c.data[httpCookie.Name] = &cookieItem{
		Cookie: httpCookie,
	}
}

// GetSessionId 从cookie中检索并返回session ID。 md5:4b41d62448c6f135
// ff:取SessionId
// c:
func (c *Cookie) GetSessionId() string {
	return c.Get(c.server.GetSessionIdName()).String()
}

// SetSessionId 在cookie中设置会话ID。 md5:56899c22c78f2267
// ff:设置SessionId到Cookie
// c:
// id:
func (c *Cookie) SetSessionId(id string) {
	c.SetCookie(
		c.server.GetSessionIdName(),
		id,
		c.request.Server.GetCookieDomain(),
		c.request.Server.GetCookiePath(),
		c.server.GetSessionCookieMaxAge(),
		CookieOptions{
			SameSite: c.request.Server.GetCookieSameSite(),
			Secure:   c.request.Server.GetCookieSecure(),
			HttpOnly: c.request.Server.GetCookieHttpOnly(),
		},
	)
}

// Get 获取并返回具有指定键的值。
// 如果指定的键不存在，并且提供了默认值`def`，则返回`def`。
// md5:f137dd1311660ee4
// ff:取值
// c:
// key:名称
// def:默认值
func (c *Cookie) Get(key string, def ...string) *gvar.Var {
	c.init()
	if r, ok := c.data[key]; ok {
		if r.Expires.IsZero() || r.Expires.After(time.Now()) {
			return gvar.New(r.Value)
		}
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}

// Remove 从使用默认域名和路径的cookie中删除指定键及其值。实际上，它告诉HTTP客户端该cookie已过期，下次不要将其发送到服务器。
// md5:4c429c6b33ce790c
// ff:删除值
// c:
// key:名称
func (c *Cookie) Remove(key string) {
	c.SetCookie(
		key,
		"",
		c.request.Server.GetCookieDomain(),
		c.request.Server.GetCookiePath(),
		-24*time.Hour,
	)
}

// RemoveCookie 使用给定的域名和路径从cookie中删除指定的键及其值。实际上，它告诉HTTP客户端该cookie已过期，下次不要将其发送到服务器。
// md5:a437da6ed4ded047
// ff:删除cookie
// c:
// key:名称
// domain:域名
// path:路径
func (c *Cookie) RemoveCookie(key, domain, path string) {
	c.SetCookie(key, "", domain, path, -24*time.Hour)
}

// Flush 将cookie项输出到客户端。 md5:77b1f132a9910559
// ff:输出
// c:
func (c *Cookie) Flush() {
	if len(c.data) == 0 {
		return
	}
	for _, v := range c.data {
		if v.FromClient {
			continue
		}
		http.SetCookie(c.response.Writer, v.Cookie)
	}
}
