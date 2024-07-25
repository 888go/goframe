// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gclient

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

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/net/gsel"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

// SetBrowserMode 启用客户端的浏览器模式。
// 当浏览器模式启用时，它会自动保存并从客户端向服务器发送cookie内容以及从服务器接收cookie内容。 md5:00d8775b045e9c09
func (c *Client) SetBrowserMode(enabled bool) *Client {
	if enabled {
		jar, _ := cookiejar.New(nil)
		c.Jar = jar
	}
	return c
}

// SetHeader 为客户端设置自定义的 HTTP 头部对。 md5:adc9509c3dab54ca
func (c *Client) SetHeader(key, value string) *Client {
	c.header[key] = value
	return c
}

// SetHeaderMap 使用映射设置自定义HTTP头。 md5:466373137e3ccd66
func (c *Client) SetHeaderMap(m map[string]string) *Client {
	for k, v := range m {
		c.header[k] = v
	}
	return c
}

// SetAgent 设置客户端的 User-Agent 头部。 md5:1ec87db52b5537ba
func (c *Client) SetAgent(agent string) *Client {
	c.header[httpHeaderUserAgent] = agent
	return c
}

// SetContentType 为客户端设置HTTP内容类型。 md5:063d3cafd0626b0a
func (c *Client) SetContentType(contentType string) *Client {
	c.header[httpHeaderContentType] = contentType
	return c
}

// SetHeaderRaw 使用原始字符串设置自定义HTTP头。 md5:e15c66308baf6cd5
func (c *Client) SetHeaderRaw(headers string) *Client {
	for _, line := range gstr.SplitAndTrim(headers, "\n") {
		array, _ := gregex.MatchString(httpRegexHeaderRaw, line)
		if len(array) >= 3 {
			c.header[array[1]] = array[2]
		}
	}
	return c
}

// SetCookie 为客户端设置一个 cookie 对。 md5:656700fcca56fb72
func (c *Client) SetCookie(key, value string) *Client {
	c.cookies[key] = value
	return c
}

// SetCookieMap 使用映射设置Cookie项。 md5:3abd18bc89684efb
func (c *Client) SetCookieMap(m map[string]string) *Client {
	for k, v := range m {
		c.cookies[k] = v
	}
	return c
}

// SetPrefix 设置请求服务器的URL前缀。 md5:945a0fd6f4acac16
func (c *Client) SetPrefix(prefix string) *Client {
	c.prefix = prefix
	return c
}

// SetTimeout 设置客户端的请求超时时间。 md5:ce4f874cd14c1c2d
func (c *Client) SetTimeout(t time.Duration) *Client {
	c.Client.Timeout = t
	return c
}

// SetBasicAuth 为客户端设置HTTP基本认证信息。 md5:22c36a5363199cd0
func (c *Client) SetBasicAuth(user, pass string) *Client {
	c.authUser = user
	c.authPass = pass
	return c
}

// SetRetry 设置重试次数和间隔。
// TODO：移除。 md5:1089293b9f9371f0
func (c *Client) SetRetry(retryCount int, retryInterval time.Duration) *Client {
	c.retryCount = retryCount
	c.retryInterval = retryInterval
	return c
}

// SetRedirectLimit 限制跳转次数。 md5:14e010f8e3d003b5
func (c *Client) SetRedirectLimit(redirectLimit int) *Client {
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= redirectLimit {
			return http.ErrUseLastResponse
		}
		return nil
	}
	return c
}

// SetNoUrlEncode 设置标记，表示在发送请求之前不编码参数。 md5:6dd55f5543918206
func (c *Client) SetNoUrlEncode(noUrlEncode bool) *Client {
	c.noUrlEncode = noUrlEncode
	return c
}

// SetProxy 为客户端设置代理。
// 当参数 `proxyURL` 为空或格式不正确时，此函数将不会执行任何操作。
// 正确的格式应为 `http://用户名:密码@IP:端口` 或 `socks5://用户名:密码@IP:端口`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。 md5:aa3f2b21308c7bec
func (c *Client) SetProxy(proxyURL string) {
	if strings.TrimSpace(proxyURL) == "" {
		return
	}
	_proxy, err := url.Parse(proxyURL)
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

// SetTLSKeyCrt 设置客户端TLS配置的证书和密钥文件。 md5:48b3322243e8e691
func (c *Client) SetTLSKeyCrt(crtFile, keyFile string) error {
	tlsConfig, err := LoadKeyCrt(crtFile, keyFile)
	if err != nil {
		return gerror.Wrap(err, "LoadKeyCrt failed")
	}
	if v, ok := c.Transport.(*http.Transport); ok {
		tlsConfig.InsecureSkipVerify = true
		v.TLSClientConfig = tlsConfig
		return nil
	}
	return gerror.New(`cannot set TLSClientConfig for custom Transport of the client`)
}

// SetTLSConfig 设置客户端的TLS配置。 md5:f1882ff235302c92
func (c *Client) SetTLSConfig(tlsConfig *tls.Config) error {
	if v, ok := c.Transport.(*http.Transport); ok {
		v.TLSClientConfig = tlsConfig
		return nil
	}
	return gerror.New(`cannot set TLSClientConfig for custom Transport of the client`)
}

// SetBuilder 设置客户端的负载均衡构建器。 md5:1f374a9a600309bb
func (c *Client) SetBuilder(builder gsel.Builder) {
	c.builder = builder
}

// SetDiscovery 为客户端设置负载均衡构建器。 md5:0ea9a7eaf5c235e7
func (c *Client) SetDiscovery(discovery gsvc.Discovery) {
	c.discovery = discovery
}
