// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gclient提供了便捷的HTTP客户端功能。
package 网页类

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"
	
	"github.com/888go/goframe"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/net/gsel"
	"github.com/888go/goframe/net/gsvc"
	"github.com/888go/goframe/os/gfile"
)

// Client 是用于管理 HTTP 请求的 HTTP 客户端。
type Client struct {
	http.Client                         // 底层HTTP客户端
	header            map[string]string // 自定义头部映射。
	cookies           map[string]string // 自定义cookie映射。
	prefix            string            // 请求前缀
	authUser          string            // HTTP基本认证：用户
	authPass          string            // HTTP基本认证：通过。
	retryCount        int               // 当请求失败时重试次数。
	noUrlEncode       bool              // 对请求参数不进行URL编码。
	retryInterval     time.Duration     // 当请求失败时的重试间隔。
	middlewareHandler []X中间件函数     // 拦截器处理器
	discovery         gsvc.Discovery    // 服务发现功能
	builder           gsel.Builder      // 请求余额的构建器
}

const (
	httpProtocolName          = `http`
	httpParamFileHolder       = `@file:`
	httpRegexParamJson        = `^[\w\[\]]+=.+`
	httpRegexHeaderRaw        = `^([\w\-]+):\s*(.+)`
	httpHeaderHost            = `Host`
	httpHeaderCookie          = `Cookie`
	httpHeaderUserAgent       = `User-Agent`
	httpHeaderContentType     = `Content-Type`
	httpHeaderContentTypeJson = `application/json`
	httpHeaderContentTypeXml  = `application/xml`
	httpHeaderContentTypeForm = `application/x-www-form-urlencoded`
)

var (
	hostname, _        = os.Hostname()
	defaultClientAgent = fmt.Sprintf(`GClient %s at %s`, gf.VERSION, hostname)
)

// New 创建并返回一个新的HTTP客户端对象。
func X创建() *Client {
	c := &Client{
		Client: http.Client{
			Transport: &http.Transport{
				// 默认情况下，不对服务器的https证书进行验证。
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				DisableKeepAlives: true,
			},
		},
		header:    make(map[string]string),
		cookies:   make(map[string]string),
		builder:   gsel.GetBuilder(),
		discovery: gsvc.GetRegistry(),
	}
	c.header[httpHeaderUserAgent] = defaultClientAgent
	// 它默认为客户端启用 OpenTelemetry。
	c.X中间件(internalMiddlewareTracing, internalMiddlewareDiscovery)
	return c
}

// Clone 深度克隆当前客户端并返回一个新的客户端。
func (c *Client) X取副本() *Client {
	newClient := X创建()
	*newClient = *c
	if len(c.header) > 0 {
		newClient.header = make(map[string]string)
		for k, v := range c.header {
			newClient.header[k] = v
		}
	}
	if len(c.cookies) > 0 {
		newClient.cookies = make(map[string]string)
		for k, v := range c.cookies {
			newClient.cookies[k] = v
		}
	}
	return newClient
}

// LoadKeyCrt 通过给定的证书和密钥文件创建并返回一个 TLS 配置对象。
func X创建TLS配置(crt路径, key路径 string) (*tls.Config, error) {
	crtPath, err := 文件类.X查找(crt路径)
	if err != nil {
		return nil, err
	}
	keyPath, err := 文件类.X查找(key路径)
	if err != nil {
		return nil, err
	}
	crt, err := tls.LoadX509KeyPair(crtPath, keyPath)
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `tls.LoadX509KeyPair failed for certFile "%s", keyFile "%s"`, crtPath, keyPath)
		return nil, err
	}
	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = []tls.Certificate{crt}
	tlsConfig.Time = time.Now
	tlsConfig.Rand = rand.Reader
	return tlsConfig, nil
}
