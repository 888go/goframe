// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

import (
	"net/http"
	"time"
	
	"github.com/888go/goframe/container/gvar"
)

// Cookie 用于 HTTP COOKIE 管理。
type Cookie struct {
	data     map[string]*cookieItem // 基础的cookie项。
	server   *Server                // 所属HTTP服务器
	request  *Request               // 属于HTTP请求。
	response *Response              // 属于HTTP响应。
}

// CookieOptions 提供了用于 cookies 的安全配置选项
type CookieOptions struct {
	SameSite http.SameSite // cookie 的 SameSite 属性
	Secure   bool          // cookie Secure 属性
	HttpOnly bool          // cookie HttpOnly 属性
}

// cookieItem 是存储在 Cookie 中的项目。
type cookieItem struct {
	*http.Cookie      // 基础的cookie项。
	FromClient   bool // 标记此cookie是从客户端接收到的。
}

// GetCookie 函数通过给定的请求创建或检索一个 cookie 对象。
// 若已存在与给定请求相关的 cookie，则检索并返回该存在的 cookie 对象。
// 若不存在与给定请求相关的 cookie，则创建并返回一个新的 cookie 对象。
func GetCookie(r *Request) *Cookie {
	if r.Cookie != nil {
		return r.Cookie
	}
	return &Cookie{
		request: r,
		server:  r.Server,
	}
}

// init 对cookie对象执行延迟初始化。
func (c *Cookie) init() {
	if c.data != nil {
		return
	}
	c.data = make(map[string]*cookieItem)
	c.response = c.request.Response
// **请勿添加任何默认的Cookie域！**
// 如果 c.request.Server.GetCookieDomain() 为空，即返回值为""，则执行以下操作：
// 将 c.request.Server.GetCookieDomain() 的值设置为 c.request.GetHost() 的返回值。
	for _, v := range c.request.Cookies() {
		c.data[v.Name] = &cookieItem{
			Cookie:     v,
			FromClient: true,
		}
	}
}

// Map 将 cookie 项以 map[string]string 的形式返回。
func (c *Cookie) Map() map[string]string {
	c.init()
	m := make(map[string]string)
	for k, v := range c.data {
		m[k] = v.Value
	}
	return m
}

// Contains 检查给定的键是否存在且在 cookie 中未过期。
func (c *Cookie) Contains(key string) bool {
	c.init()
	if r, ok := c.data[key]; ok {
		if r.Expires.IsZero() || r.Expires.After(time.Now()) {
			return true
		}
	}
	return false
}

// Set 使用默认的域名、路径和过期时间设置cookie项。
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

// SetCookie 用于给指定的域名、路径设置cookie项，并设置其过期时间。
// 可选参数 `options` 指定了额外的安全配置，通常为空。
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

// SetHttpCookie 通过 *http.Cookie 设置 cookie。
func (c *Cookie) SetHttpCookie(httpCookie *http.Cookie) {
	c.init()
	c.data[httpCookie.Name] = &cookieItem{
		Cookie: httpCookie,
	}
}

// GetSessionId 从cookie中获取并返回会话id。
func (c *Cookie) GetSessionId() string {
	return c.Get(c.server.GetSessionIdName()).String()
}

// SetSessionId 将会话ID设置到cookie中。
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

// Get 方法用于获取并返回指定键的值。
// 如果指定的键不存在且提供了默认值 `def`，则返回 `def`。
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

// Remove 函数通过使用默认的域名和路径，从cookie中删除指定的键及其对应的值。
// 实质上，它告知http客户端该cookie已过期，下次不要将其发送到服务器。
func (c *Cookie) Remove(key string) {
	c.SetCookie(
		key,
		"",
		c.request.Server.GetCookieDomain(),
		c.request.Server.GetCookiePath(),
		-24*time.Hour,
	)
}

// RemoveCookie 通过给定的域名和路径，从cookie中删除指定的键及其对应的值。
// 实际上，它会告知http客户端该cookie已过期，下次不要将其发送到服务器。
func (c *Cookie) RemoveCookie(key, domain, path string) {
	c.SetCookie(key, "", domain, path, -24*time.Hour)
}

// Flush 将cookie项输出到客户端。
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
