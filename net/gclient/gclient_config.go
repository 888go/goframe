// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 网页类

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/proxy"

	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/net/gsel"
	"github.com/888go/goframe/net/gsvc"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
)

// X启用浏览器模式 启用客户端的浏览器模式。
// 当浏览器模式启用时，它会自动保存并从客户端向服务器发送cookie内容以及从服务器接收cookie内容。
// md5:00d8775b045e9c09
func (c *Client) X启用浏览器模式(启用 bool) *Client {
	if 启用 {
		jar, _ := cookiejar.New(nil)
		c.Jar = jar
	}
	return c
}

// X设置协议头 为客户端设置自定义的 HTTP 头部对。 md5:adc9509c3dab54ca
func (c *Client) X设置协议头(名称, 值 string) *Client {
	c.header[名称] = 值
	return c
}

// X设置Map协议头 使用映射设置自定义HTTP头。 md5:466373137e3ccd66
func (c *Client) X设置Map协议头(map协议头 map[string]string) *Client {
	for k, v := range map协议头 {
		c.header[k] = v
	}
	return c
}

// X设置UA 设置客户端的 User-Agent 头部。 md5:1ec87db52b5537ba
func (c *Client) X设置UA(UA值 string) *Client {
	c.header[httpHeaderUserAgent] = UA值
	return c
}

// X设置内容类型 为客户端设置HTTP内容类型。 md5:063d3cafd0626b0a
func (c *Client) X设置内容类型(内容类型 string) *Client {
	c.header[httpHeaderContentType] = 内容类型
	return c
}

// X设置原始协议头 使用原始字符串设置自定义HTTP头。 md5:e15c66308baf6cd5
func (c *Client) X设置原始协议头(原始协议头 string) *Client {
	for _, line := range gstr.X分割并忽略空值(原始协议头, "\n") {
		array, _ := gregex.X匹配文本(httpRegexHeaderRaw, line)
		if len(array) >= 3 {
			c.header[array[1]] = array[2]
		}
	}
	return c
}

// X设置cookie 为客户端设置一个 cookie 对。 md5:656700fcca56fb72
func (c *Client) X设置cookie(名称, 值 string) *Client {
	c.cookies[名称] = 值
	return c
}

// X设置CookieMap 使用映射设置Cookie项。 md5:3abd18bc89684efb
func (c *Client) X设置CookieMap(MapCookie map[string]string) *Client {
	for k, v := range MapCookie {
		c.cookies[k] = v
	}
	return c
}

// X设置url前缀 设置请求服务器的URL前缀。 md5:945a0fd6f4acac16
func (c *Client) X设置url前缀(前缀 string) *Client {
	c.prefix = 前缀
	return c
}

// X设置超时 设置客户端的请求超时时间。 md5:ce4f874cd14c1c2d
func (c *Client) X设置超时(时长 time.Duration) *Client {
	c.Client.Timeout = 时长
	return c
}

// X设置账号密码 为客户端设置HTTP基本认证信息。 md5:22c36a5363199cd0
func (c *Client) X设置账号密码(账号, 密码 string) *Client {
	c.authUser = 账号
	c.authPass = 密码
	return c
}

// X设置重试与间隔 设置重试次数和间隔。
// TODO：移除。
// md5:1089293b9f9371f0
func (c *Client) X设置重试与间隔(重试计数 int, 重试间隔时长 time.Duration) *Client {
	c.retryCount = 重试计数
	c.retryInterval = 重试间隔时长
	return c
}

// X设置重定向次数限制 限制跳转次数。 md5:14e010f8e3d003b5
func (c *Client) X设置重定向次数限制(次数 int) *Client {
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= 次数 {
			return http.ErrUseLastResponse
		}
		return nil
	}
	return c
}

// X设置请求参数禁止URL编码 设置标记，表示在发送请求之前不编码参数。 md5:6dd55f5543918206
func (c *Client) X设置请求参数禁止URL编码(禁止编码 bool) *Client {
	c.noUrlEncode = 禁止编码
	return c
}

// X设置代理 为客户端设置代理。
// 当参数 `proxyURL` 为空或格式不正确时，此函数将不会执行任何操作。
// 正确的格式应为 `http://用户名:密码@IP:端口` 或 `socks5://用户名:密码@IP:端口`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。
// md5:aa3f2b21308c7bec
func (c *Client) X设置代理(代理地址 string) {
	if strings.TrimSpace(代理地址) == "" {
		return
	}
	_proxy, err := url.Parse(代理地址)
	if err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
		return
	}
	if _proxy.Scheme == httpProtocolName {
		if v, ok := c.Transport.(*http.Transport); ok {
			v.Proxy = http.ProxyURL(_proxy)
		}
	} else {
		auth := &proxy.Auth{}
		user := _proxy.User.Username()
		if user != "" {
			auth.User = user
			password, hasPassword := _proxy.User.Password()
			if hasPassword && password != "" {
				auth.Password = password
			}
		} else {
			auth = nil
		}
						// 参考源代码，错误始终为nil. md5:43df5b2c264029cb
		dialer, err := proxy.SOCKS5(
			"tcp",
			_proxy.Host,
			auth,
			&net.Dialer{
				Timeout:   c.Client.Timeout,
				KeepAlive: c.Client.Timeout,
			},
		)
		if err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
			return
		}
		if v, ok := c.Transport.(*http.Transport); ok {
			v.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
				return dialer.Dial(network, addr)
			}
		}
				// 设置超时时间为10秒. md5:ee88d389b4a64b4a
	}
}

// X设置证书 设置客户端TLS配置的证书和密钥文件。 md5:48b3322243e8e691
func (c *Client) X设置证书(crt路径, key路径 string) error {
	tlsConfig, err := X创建TLS配置(crt路径, key路径)
	if err != nil {
		return gerror.X多层错误(err, "LoadKeyCrt failed")
	}
	if v, ok := c.Transport.(*http.Transport); ok {
		tlsConfig.InsecureSkipVerify = true
		v.TLSClientConfig = tlsConfig
		return nil
	}
	return gerror.X创建(`cannot set TLSClientConfig for custom Transport of the client`)
}

// X设置TLS配置 设置客户端的TLS配置。 md5:f1882ff235302c92
func (c *Client) X设置TLS配置(TLS配置 *tls.Config) error {
	if v, ok := c.Transport.(*http.Transport); ok {
		v.TLSClientConfig = TLS配置
		return nil
	}
	return gerror.X创建(`cannot set TLSClientConfig for custom Transport of the client`)
}

// SetBuilder 设置客户端的负载均衡构建器。 md5:1f374a9a600309bb
func (c *Client) SetBuilder(builder gsel.Builder) {
	c.builder = builder
}

// SetDiscovery 为客户端设置负载均衡构建器。 md5:0ea9a7eaf5c235e7
func (c *Client) SetDiscovery(discovery gsvc.Discovery) {
	c.discovery = discovery
}
