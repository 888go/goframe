// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gclient

import (
	"time"
	
	"github.com/gogf/gf/v2/net/gsvc"
)

// Prefix 是一个链式函数，
// 它用于设置该客户端下一次请求的URL前缀。
// 示例：
// Prefix("http://127.0.0.1:8199/api/v1") // 设置URL前缀为v1版本
// Prefix("http://127.0.0.1:8199/api/v2") // 设置URL前缀为v2版本
func (c *Client) Prefix(prefix string) *Client {
	newClient := c.Clone()
	newClient.SetPrefix(prefix)
	return newClient
}

// Header 是一个链式函数，
// 用于为下一个请求设置自定义HTTP头部，参数为映射(map)类型。
func (c *Client) Header(m map[string]string) *Client {
	newClient := c.Clone()
	newClient.SetHeaderMap(m)
	return newClient
}

// HeaderRaw 是一个链式函数，
// 它用于设置下一次请求的自定义HTTP头部，使用原始字符串形式。
func (c *Client) HeaderRaw(headers string) *Client {
	newClient := c.Clone()
	newClient.SetHeaderRaw(headers)
	return newClient
}

// Discovery 是一个链式函数，用于为客户端设置发现机制。
// 你可以使用 `Discovery(nil)` 来禁用当前客户端的发现功能。
func (c *Client) Discovery(discovery gsvc.Discovery) *Client {
	newClient := c.Clone()
	newClient.SetDiscovery(discovery)
	return newClient
}

// Cookie 是一个链式函数，
// 用于为下一次请求通过映射设置cookie项。
func (c *Client) Cookie(m map[string]string) *Client {
	newClient := c.Clone()
	newClient.SetCookieMap(m)
	return newClient
}

// ContentType 是一个链式函数，
// 用于为下一个请求设置HTTP内容类型。
func (c *Client) ContentType(contentType string) *Client {
	newClient := c.Clone()
	newClient.SetContentType(contentType)
	return newClient
}

// ContentJson 是一个链式函数，
// 它用于为下一个请求设置HTTP内容类型为 "application/json"。
//
// 注意，它还会自动检查并以JSON格式对参数进行编码。
func (c *Client) ContentJson() *Client {
	newClient := c.Clone()
	newClient.SetContentType(httpHeaderContentTypeJson)
	return newClient
}

// ContentXml 是一个链式函数，
// 它为下一次请求设置HTTP内容类型为"application/xml"。
//
// 注意，它还会自动检查并以XML格式对参数进行编码。
func (c *Client) ContentXml() *Client {
	newClient := c.Clone()
	newClient.SetContentType(httpHeaderContentTypeXml)
	return newClient
}

// Timeout 是一个链式函数，
// 用于设置下一次请求的超时时间。
func (c *Client) Timeout(t time.Duration) *Client {
	newClient := c.Clone()
	newClient.SetTimeout(t)
	return newClient
}

// BasicAuth 是一个链式函数，
// 用于为下一次请求设置HTTP基本认证信息。
func (c *Client) BasicAuth(user, pass string) *Client {
	newClient := c.Clone()
	newClient.SetBasicAuth(user, pass)
	return newClient
}

// Retry 是一个链式函数，
// 它在下一次请求失败时设置重试次数和间隔。
func (c *Client) Retry(retryCount int, retryInterval time.Duration) *Client {
	newClient := c.Clone()
	newClient.SetRetry(retryCount, retryInterval)
	return newClient
}

// Proxy 是一个链式函数，
// 用于设置下一次请求的代理。
// 请确保传递正确的 `proxyURL`。
// 正确的格式应为 `http://USER:PASSWORD@IP:PORT` 或 `socks5://USER:PASSWORD@IP:PORT`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。
func (c *Client) Proxy(proxyURL string) *Client {
	newClient := c.Clone()
	newClient.SetProxy(proxyURL)
	return newClient
}

// RedirectLimit 是一个链式函数，
// 用于设置请求的重定向限制次数。
func (c *Client) RedirectLimit(redirectLimit int) *Client {
	newClient := c.Clone()
	newClient.SetRedirectLimit(redirectLimit)
	return newClient
}

// NoUrlEncode 设置标记，表示在发送请求前不应对参数进行URL编码。
func (c *Client) NoUrlEncode() *Client {
	newClient := c.Clone()
	newClient.SetNoUrlEncode(true)
	return newClient
}
