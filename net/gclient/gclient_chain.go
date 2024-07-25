// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gclient

import (
	"time"

	"github.com/gogf/gf/v2/net/gsvc"
)

// Prefix是一个链式函数，
// 它为该客户端的下一个请求设置URL前缀。例如：
// Prefix("http://127.0.0.1:8199/api/v1")
// Prefix("http://127.0.0.1:8199/api/v2") md5:223b00b33201dec6
func (c *Client) Prefix(prefix string) *Client {
	newClient := c.Clone()
	newClient.SetPrefix(prefix)
	return newClient
}

// Header是一个链式函数，它为下一个请求设置自定义HTTP头，使用映射（map）来存储头信息。 md5:e9d44530a2916792
func (c *Client) Header(m map[string]string) *Client {
	newClient := c.Clone()
	newClient.SetHeaderMap(m)
	return newClient
}

// HeaderRaw 是一个链式函数，
// 用于为下一个请求使用原始字符串设置自定义的HTTP头。 md5:dff868fea39b9738
func (c *Client) HeaderRaw(headers string) *Client {
	newClient := c.Clone()
	newClient.SetHeaderRaw(headers)
	return newClient
}

// Discovery 是一个链式函数，用于设置客户端的发现机制。
// 你可以使用 `Discovery(nil)` 来禁用当前客户端的发现功能。 md5:75d0bac47755ed2a
func (c *Client) Discovery(discovery gsvc.Discovery) *Client {
	newClient := c.Clone()
	newClient.SetDiscovery(discovery)
	return newClient
}

// Cookie是一个链式函数，它使用映射为下一个请求设置Cookie项。 md5:458381c5ded10dae
func (c *Client) Cookie(m map[string]string) *Client {
	newClient := c.Clone()
	newClient.SetCookieMap(m)
	return newClient
}

// ContentType是一个链式函数，用于为下一个请求设置HTTP内容类型。 md5:48b017c4d22a94ee
func (c *Client) ContentType(contentType string) *Client {
	newClient := c.Clone()
	newClient.SetContentType(contentType)
	return newClient
}

// ContentJson 是一个链式函数，
// 用于为下一个请求将HTTP内容类型设置为 "application/json"。
//
// 注意，它还会自动检查并以JSON格式对参数进行编码。 md5:3b265101262bf8be
func (c *Client) ContentJson() *Client {
	newClient := c.Clone()
	newClient.SetContentType(httpHeaderContentTypeJson)
	return newClient
}

// ContentXml是一个链式函数，
// 它将下一个请求的HTTP内容类型设置为"application/xml"。
//
// 请注意，它还会自动检查并把参数编码为XML格式。 md5:b6e6ad51561fa7a6
func (c *Client) ContentXml() *Client {
	newClient := c.Clone()
	newClient.SetContentType(httpHeaderContentTypeXml)
	return newClient
}

// Timeout是一个链式函数，它为下一个请求设置超时时间。 md5:8198ec107ce0113f
func (c *Client) Timeout(t time.Duration) *Client {
	newClient := c.Clone()
	newClient.SetTimeout(t)
	return newClient
}

// BasicAuth是一个链式函数，为下一个请求设置HTTP基本认证信息。 md5:c0ca33e41692898b
func (c *Client) BasicAuth(user, pass string) *Client {
	newClient := c.Clone()
	newClient.SetBasicAuth(user, pass)
	return newClient
}

// Retry 是一个链式函数，
// 用于设置在下次请求失败时的重试次数和间隔时间。
// TODO 待移除。 md5:2206b70379c77ed7
func (c *Client) Retry(retryCount int, retryInterval time.Duration) *Client {
	newClient := c.Clone()
	newClient.SetRetry(retryCount, retryInterval)
	return newClient
}

// Proxy 是一个链式函数，
// 用于为下一个请求设置代理。
// 确保传递正确的 `proxyURL`。
// 正确的格式应为 `http://USER:PASSWORD@IP:PORT` 或 `socks5://USER:PASSWORD@IP:PORT`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。 md5:d8c660f7a12e72ea
func (c *Client) Proxy(proxyURL string) *Client {
	newClient := c.Clone()
	newClient.SetProxy(proxyURL)
	return newClient
}

// RedirectLimit 是一个链式函数，用于设置请求的重定向限制次数。 md5:ecac398510aa92bb
func (c *Client) RedirectLimit(redirectLimit int) *Client {
	newClient := c.Clone()
	newClient.SetRedirectLimit(redirectLimit)
	return newClient
}

// NoUrlEncode 设置一个标记，表示在发送请求之前不进行参数编码。 md5:0f78cc83f0909b0e
func (c *Client) NoUrlEncode() *Client {
	newClient := c.Clone()
	newClient.SetNoUrlEncode(true)
	return newClient
}
