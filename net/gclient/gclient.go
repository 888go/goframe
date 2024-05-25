// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gclient 包提供了便捷的 HTTP 客户端功能。. md5:e1b459f6ec089b4e
package gclient

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gsel"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gfile"
)

// Client 是用于HTTP请求管理的HTTP客户端。. md5:4ad1e09a685144a3
type Client struct {
	http.Client                         // 基础的HTTP客户端。. md5:c6441b0ac3b9a383
	header            map[string]string // Custom header map.
	cookies           map[string]string // Custom cookie map.
	prefix            string            // Prefix for request.
	authUser          string            // HTTP基本身份验证：用户。. md5:16a49cf2f34fe020
	authPass          string            // HTTP基本认证：通过。. md5:bb96d4f4a15daaad
	retryCount        int               // 当请求失败时的重试次数。. md5:94508857e0c3610f
	noUrlEncode       bool              // 对请求参数不做URL编码。. md5:e1a507c4ef43df36
	retryInterval     time.Duration     // 请求失败时的重试间隔。. md5:f5c1143b17b02297
	middlewareHandler []HandlerFunc     // Interceptor handlers
	discovery         gsvc.Discovery    // Discovery for service.
	builder           gsel.Builder      // 用于构建请求平衡的构造器。. md5:03b93939823f0270
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

// New 创建并返回一个新的HTTP客户端对象。. md5:704d6059158b6cda
func New() *Client {
	c := &Client{
		Client: http.Client{
			Transport: &http.Transport{
				// 默认情况下，不对服务器的HTTPS证书进行验证。. md5:afd65bbc5be16457
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
	// 它默认启用客户端的OpenTelemetry。. md5:d39de17478071c01
	c.Use(internalMiddlewareObservability, internalMiddlewareDiscovery)
	return c
}

// 克隆当前客户端并深拷贝，返回一个新的客户端实例。. md5:6d6f07ba99eeeabe
func (c *Client) Clone() *Client {
	newClient := New()
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

// LoadKeyCrt 根据给定的证书和密钥文件创建并返回一个 TLS 配置对象。. md5:e31385756c06b0a4
func LoadKeyCrt(crtFile, keyFile string) (*tls.Config, error) {
	crtPath, err := gfile.Search(crtFile)
	if err != nil {
		return nil, err
	}
	keyPath, err := gfile.Search(keyFile)
	if err != nil {
		return nil, err
	}
	crt, err := tls.LoadX509KeyPair(crtPath, keyPath)
	if err != nil {
		err = gerror.Wrapf(err, `tls.LoadX509KeyPair failed for certFile "%s", keyFile "%s"`, crtPath, keyPath)
		return nil, err
	}
	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = []tls.Certificate{crt}
	tlsConfig.Time = time.Now
	tlsConfig.Rand = rand.Reader
	return tlsConfig, nil
}
