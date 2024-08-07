// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"net/http"
	"time"

	gvar "github.com/888go/goframe/container/gvar"
)

// 用于HTTP_COOKIE管理的Cookie。 md5:5990685f9449a6df
type Cookie struct {
	data     map[string]*cookieItem // 底层cookie项。 md5:b89c165cdb180382
	server   *X服务                // Belonged HTTP server
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

// X取cookie对象 通过给定的请求创建或检索一个cookie对象。
// 如果已经存在与给定请求匹配的cookie对象，它将检索并返回该对象。
// 如果不存在与给定请求匹配的cookie对象，它将创建一个新的cookie对象并返回。
// md5:5b2b3a376a2f6162
func X取cookie对象(r *Request) *Cookie {
	if r.Cookie != nil {
		return r.Cookie
	}
	return &Cookie{
		request: r,
		server:  r.X服务,
	}
}

// init 函数对cookie对象进行惰性初始化。 md5:6a771fc762ea2458
func (c *Cookie) init() {
	if c.data != nil {
		return
	}
	c.data = make(map[string]*cookieItem)
	c.response = c.request.X响应
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

// X取Map将Cookie项作为map[string]string返回。 md5:6d265855ff6356fa
func (c *Cookie) X取Map() map[string]string {
	c.init()
	m := make(map[string]string)
	for k, v := range c.data {
		m[k] = v.Value
	}
	return m
}

// X是否已过期 检查给定的键是否存在于cookie中且未过期。 md5:acb79130bbbf78e6
func (c *Cookie) X是否已过期(名称 string) bool {
	c.init()
	if r, ok := c.data[名称]; ok {
		if r.Expires.IsZero() || r.Expires.After(time.Now()) {
			return true
		}
	}
	return false
}

// X设置值 使用默认域名、路径和过期时间来设置cookie项。 md5:66be80c8c6c07dc1
func (c *Cookie) X设置值(key, value string) {
	c.X设置cookie(
		key,
		value,
		c.request.X服务.X取Cookie域名(),
		c.request.X服务.X取Cookie路径(),
		c.request.X服务.X取Cookie最大存活时长(),
		CookieOptions{
			SameSite: c.request.X服务.X取CookieSameSite(),
			Secure:   c.request.X服务.X取Cookie安全(),
			HttpOnly: c.request.X服务.X取CookieHttpOnly(),
		},
	)
}

// X设置cookie 设置具有给定域名、路径和过期时间的cookie项。
// 可选参数`options`指定了额外的安全配置，通常为空。
// md5:2afc45e40597ce0d
func (c *Cookie) X设置cookie(名称, 值, 域名, 路径 string, 最大存活时长 time.Duration, 安全配置项 ...CookieOptions) {
	c.init()
	config := CookieOptions{}
	if len(安全配置项) > 0 {
		config = 安全配置项[0]
	}
	httpCookie := &http.Cookie{
		Name:     名称,
		Value:    值,
		Path:     路径,
		Domain:   域名,
		HttpOnly: config.HttpOnly,
		SameSite: config.SameSite,
		Secure:   config.Secure,
	}
	if 最大存活时长 != 0 {
		httpCookie.Expires = time.Now().Add(最大存活时长)
	}
	c.data[名称] = &cookieItem{
		Cookie: httpCookie,
	}
}

// X设置httpcookie 使用*http.Cookie设置cookie。 md5:de525635cedd10e4
func (c *Cookie) X设置httpcookie(httpCookie *http.Cookie) {
	c.init()
	c.data[httpCookie.Name] = &cookieItem{
		Cookie: httpCookie,
	}
}

// X取SessionId 从cookie中检索并返回session ID。 md5:4b41d62448c6f135
func (c *Cookie) X取SessionId() string {
	return c.X取值(c.server.X取SessionID名称()).String()
}

// X设置SessionId到Cookie 在cookie中设置会话ID。 md5:56899c22c78f2267
func (c *Cookie) X设置SessionId到Cookie(id string) {
	c.X设置cookie(
		c.server.X取SessionID名称(),
		id,
		c.request.X服务.X取Cookie域名(),
		c.request.X服务.X取Cookie路径(),
		c.server.X取SessionCookie存活时长(),
		CookieOptions{
			SameSite: c.request.X服务.X取CookieSameSite(),
			Secure:   c.request.X服务.X取Cookie安全(),
			HttpOnly: c.request.X服务.X取CookieHttpOnly(),
		},
	)
}

// X取值 获取并返回具有指定键的值。
// 如果指定的键不存在，并且提供了默认值`def`，则返回`def`。
// md5:f137dd1311660ee4
func (c *Cookie) X取值(名称 string, 默认值 ...string) *gvar.Var {
	c.init()
	if r, ok := c.data[名称]; ok {
		if r.Expires.IsZero() || r.Expires.After(time.Now()) {
			return gvar.X创建(r.Value)
		}
	}
	if len(默认值) > 0 {
		return gvar.X创建(默认值[0])
	}
	return nil
}

// X删除值 从使用默认域名和路径的cookie中删除指定键及其值。实际上，它告诉HTTP客户端该cookie已过期，下次不要将其发送到服务器。
// md5:4c429c6b33ce790c
func (c *Cookie) X删除值(名称 string) {
	c.X设置cookie(
		名称,
		"",
		c.request.X服务.X取Cookie域名(),
		c.request.X服务.X取Cookie路径(),
		-24*time.Hour,
	)
}

// X删除cookie 使用给定的域名和路径从cookie中删除指定的键及其值。实际上，它告诉HTTP客户端该cookie已过期，下次不要将其发送到服务器。
// md5:a437da6ed4ded047
func (c *Cookie) X删除cookie(名称, 域名, 路径 string) {
	c.X设置cookie(名称, "", 域名, 路径, -24*time.Hour)
}

// X输出 将cookie项输出到客户端。 md5:77b1f132a9910559
func (c *Cookie) X输出() {
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
