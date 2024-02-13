// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

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
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/net/gsel"
	"github.com/888go/goframe/net/gsvc"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
)

// SetBrowserMode启用客户端的浏览器模式。
// 当浏览器模式被启用后，它会自动保存并从服务器发送、接收cookie内容。
func (c *Client) X启用浏览器模式(启用 bool) *Client {
	if 启用 {
		jar, _ := cookiejar.New(nil)
		c.Jar = jar
	}
	return c
}

// SetHeader 为客户端设置自定义HTTP头键值对。
func (c *Client) X设置协议头(名称, 值 string) *Client {
	c.header[名称] = 值
	return c
}

// SetHeaderMap 通过映射设置自定义HTTP头。
func (c *Client) X设置Map协议头(map协议头 map[string]string) *Client {
	for k, v := range map协议头 {
		c.header[k] = v
	}
	return c
}

// SetAgent 设置客户端的 User-Agent 头部信息。
func (c *Client) X设置UA(UA值 string) *Client {
	c.header[httpHeaderUserAgent] = UA值
	return c
}

// SetContentType 为客户端设置HTTP内容类型。
func (c *Client) X设置内容类型(内容类型 string) *Client {
	c.header[httpHeaderContentType] = 内容类型
	return c
}

// SetHeaderRaw 通过原始字符串设置自定义HTTP头。
func (c *Client) X设置原始协议头(原始协议头 string) *Client {
	for _, line := range 文本类.X分割并忽略空值(原始协议头, "\n") {
		array, _ := 正则类.X匹配文本(httpRegexHeaderRaw, line)
		if len(array) >= 3 {
			c.header[array[1]] = array[2]
		}
	}
	return c
}

// SetCookie 为客户端设置一个cookie对。
func (c *Client) X设置cookie(名称, 值 string) *Client {
	c.cookies[名称] = 值
	return c
}

// SetCookieMap 通过map设置cookie项目。
func (c *Client) X设置CookieMap(MapCookie map[string]string) *Client {
	for k, v := range MapCookie {
		c.cookies[k] = v
	}
	return c
}

// SetPrefix 设置请求服务器 URL 前缀。
func (c *Client) X设置url前缀(前缀 string) *Client {
	c.prefix = 前缀
	return c
}

// SetTimeout 设置客户端的请求超时时间。
func (c *Client) X设置超时(时长 time.Duration) *Client {
	c.Client.Timeout = 时长
	return c
}

// SetBasicAuth为客户端设置HTTP基础认证信息。
func (c *Client) X设置账号密码(账号, 密码 string) *Client {
	c.authUser = 账号
	c.authPass = 密码
	return c
}

// SetRetry 设置重试次数和间隔。
func (c *Client) X设置重试与间隔(重试计数 int, 重试间隔时长 time.Duration) *Client {
	c.retryCount = 重试计数
	c.retryInterval = 重试间隔时长
	return c
}

// SetRedirectLimit 限制跳转次数。
func (c *Client) X设置重定向次数限制(次数 int) *Client {
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= 次数 {
			return http.ErrUseLastResponse
		}
		return nil
	}
	return c
}

// SetNoUrlEncode 设置标记，表示在发送请求前不对参数进行URL编码。
func (c *Client) X设置请求参数禁止URL编码(禁止编码 bool) *Client {
	c.noUrlEncode = 禁止编码
	return c
}

// SetProxy 为客户端设置代理。
// 当参数`proxyURL`为空或者格式不正确时，此函数将不做任何操作。
// 正确的格式应如 `http://USER:PASSWORD@IP:PORT` 或 `socks5://USER:PASSWORD@IP:PORT`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。
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
		// 参考源代码，error 值始终为 nil
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
		// c.SetTimeout(10*time.Second) // 设置c的超时时间为10秒钟
	}
}

// SetTLSKeyCrt 用于设置客户端TLS配置所需的证书和密钥文件。
func (c *Client) X设置证书(crt路径, key路径 string) error {
	tlsConfig, err := X创建TLS配置(crt路径, key路径)
	if err != nil {
		return 错误类.X多层错误(err, "LoadKeyCrt failed")
	}
	if v, ok := c.Transport.(*http.Transport); ok {
		tlsConfig.InsecureSkipVerify = true
		v.TLSClientConfig = tlsConfig
		return nil
	}
	return 错误类.X创建(`cannot set TLSClientConfig for custom Transport of the client`)
}

// SetTLSConfig 设置客户端的 TLS 配置。
func (c *Client) X设置TLS配置(TLS配置 *tls.Config) error {
	if v, ok := c.Transport.(*http.Transport); ok {
		v.TLSClientConfig = TLS配置
		return nil
	}
	return 错误类.X创建(`cannot set TLSClientConfig for custom Transport of the client`)
}

// SetBuilder 为客户端设置负载均衡器生成器。
func (c *Client) SetBuilder(builder gsel.Builder) {
	c.builder = builder
}

// SetDiscovery 设置客户端的负载均衡器生成器。
func (c *Client) SetDiscovery(discovery gsvc.Discovery) {
	c.discovery = discovery
}
