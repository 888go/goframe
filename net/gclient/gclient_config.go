// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

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
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/net/gsel"
	"github.com/888go/goframe/net/gsvc"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	)
// SetBrowserMode启用客户端的浏览器模式。
// 当浏览器模式被启用后，它会自动保存并从服务器发送、接收cookie内容。
func (c *Client) SetBrowserMode(enabled bool) *Client {
	if enabled {
		jar, _ := cookiejar.New(nil)
		c.Jar = jar
	}
	return c
}

// SetHeader 为客户端设置自定义HTTP头键值对。
func (c *Client) SetHeader(key, value string) *Client {
	c.header[key] = value
	return c
}

// SetHeaderMap 通过映射设置自定义HTTP头。
func (c *Client) SetHeaderMap(m map[string]string) *Client {
	for k, v := range m {
		c.header[k] = v
	}
	return c
}

// SetAgent 设置客户端的 User-Agent 头部信息。
func (c *Client) SetAgent(agent string) *Client {
	c.header[httpHeaderUserAgent] = agent
	return c
}

// SetContentType 为客户端设置HTTP内容类型。
func (c *Client) SetContentType(contentType string) *Client {
	c.header[httpHeaderContentType] = contentType
	return c
}

// SetHeaderRaw 通过原始字符串设置自定义HTTP头。
func (c *Client) SetHeaderRaw(headers string) *Client {
	for _, line := range gstr.SplitAndTrim(headers, "\n") {
		array, _ := gregex.MatchString(httpRegexHeaderRaw, line)
		if len(array) >= 3 {
			c.header[array[1]] = array[2]
		}
	}
	return c
}

// SetCookie 为客户端设置一个cookie对。
func (c *Client) SetCookie(key, value string) *Client {
	c.cookies[key] = value
	return c
}

// SetCookieMap 通过map设置cookie项目。
func (c *Client) SetCookieMap(m map[string]string) *Client {
	for k, v := range m {
		c.cookies[k] = v
	}
	return c
}

// SetPrefix 设置请求服务器 URL 前缀。
func (c *Client) SetPrefix(prefix string) *Client {
	c.prefix = prefix
	return c
}

// SetTimeout 设置客户端的请求超时时间。
func (c *Client) SetTimeout(t time.Duration) *Client {
	c.Client.Timeout = t
	return c
}

// SetBasicAuth为客户端设置HTTP基础认证信息。
func (c *Client) SetBasicAuth(user, pass string) *Client {
	c.authUser = user
	c.authPass = pass
	return c
}

// SetRetry 设置重试次数和间隔。
func (c *Client) SetRetry(retryCount int, retryInterval time.Duration) *Client {
	c.retryCount = retryCount
	c.retryInterval = retryInterval
	return c
}

// SetRedirectLimit 限制跳转次数。
func (c *Client) SetRedirectLimit(redirectLimit int) *Client {
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= redirectLimit {
			return http.ErrUseLastResponse
		}
		return nil
	}
	return c
}

// SetNoUrlEncode 设置标记，表示在发送请求前不对参数进行URL编码。
func (c *Client) SetNoUrlEncode(noUrlEncode bool) *Client {
	c.noUrlEncode = noUrlEncode
	return c
}

// SetProxy 为客户端设置代理。
// 当参数`proxyURL`为空或者格式不正确时，此函数将不做任何操作。
// 正确的格式应如 `http://USER:PASSWORD@IP:PORT` 或 `socks5://USER:PASSWORD@IP:PORT`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。
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

// SetTLSConfig 设置客户端的 TLS 配置。
func (c *Client) SetTLSConfig(tlsConfig *tls.Config) error {
	if v, ok := c.Transport.(*http.Transport); ok {
		v.TLSClientConfig = tlsConfig
		return nil
	}
	return gerror.New(`cannot set TLSClientConfig for custom Transport of the client`)
}

// SetBuilder 为客户端设置负载均衡器生成器。
func (c *Client) SetBuilder(builder gsel.Builder) {
	c.builder = builder
}

// SetDiscovery 设置客户端的负载均衡器生成器。
func (c *Client) SetDiscovery(discovery gsvc.Discovery) {
	c.discovery = discovery
}
