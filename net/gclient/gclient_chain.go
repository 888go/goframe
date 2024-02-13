// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类

import (
	"time"
	
	"github.com/888go/goframe/net/gsvc"
)

// Prefix 是一个链式函数，
// 它用于设置该客户端下一次请求的URL前缀。
// 示例：
// Prefix("http://127.0.0.1:8199/api/v1") // 设置URL前缀为v1版本
// Prefix("http://127.0.0.1:8199/api/v2") // 设置URL前缀为v2版本
func (c *Client) Url前缀(前缀 string) *Client {
	newClient := c.X取副本()
	newClient.X设置url前缀(前缀)
	return newClient
}

// Header 是一个链式函数，
// 用于为下一个请求设置自定义HTTP头部，参数为映射(map)类型。
func (c *Client) X协议头(map协议头 map[string]string) *Client {
	newClient := c.X取副本()
	newClient.X设置Map协议头(map协议头)
	return newClient
}

// HeaderRaw 是一个链式函数，
// 它用于设置下一次请求的自定义HTTP头部，使用原始字符串形式。
func (c *Client) X原始协议头(原始协议头 string) *Client {
	newClient := c.X取副本()
	newClient.X设置原始协议头(原始协议头)
	return newClient
}

// Discovery 是一个链式函数，用于为客户端设置发现机制。
// 你可以使用 `Discovery(nil)` 来禁用当前客户端的发现功能。
func (c *Client) Discovery(discovery gsvc.Discovery) *Client {
	newClient := c.X取副本()
	newClient.SetDiscovery(discovery)
	return newClient
}

// Cookie 是一个链式函数，
// 用于为下一次请求通过映射设置cookie项。
func (c *Client) Cookie(MapCookie map[string]string) *Client {
	newClient := c.X取副本()
	newClient.X设置CookieMap(MapCookie)
	return newClient
}

// ContentType 是一个链式函数，
// 用于为下一个请求设置HTTP内容类型。
func (c *Client) X内容类型(contentType string) *Client {
	newClient := c.X取副本()
	newClient.X设置内容类型(contentType)
	return newClient
}

// ContentJson 是一个链式函数，
// 它用于为下一个请求设置HTTP内容类型为 "application/json"。
//
// 注意，它还会自动检查并以JSON格式对参数进行编码。
func (c *Client) X内容类型json() *Client {
	newClient := c.X取副本()
	newClient.X设置内容类型(httpHeaderContentTypeJson)
	return newClient
}

// ContentXml 是一个链式函数，
// 它为下一次请求设置HTTP内容类型为"application/xml"。
//
// 注意，它还会自动检查并以XML格式对参数进行编码。
func (c *Client) X内容类型xml() *Client {
	newClient := c.X取副本()
	newClient.X设置内容类型(httpHeaderContentTypeXml)
	return newClient
}

// Timeout 是一个链式函数，
// 用于设置下一次请求的超时时间。
func (c *Client) X超时(时长 time.Duration) *Client {
	newClient := c.X取副本()
	newClient.X设置超时(时长)
	return newClient
}

// BasicAuth 是一个链式函数，
// 用于为下一次请求设置HTTP基本认证信息。
func (c *Client) X账号密码(账号, 密码 string) *Client {
	newClient := c.X取副本()
	newClient.X设置账号密码(账号, 密码)
	return newClient
}

// Retry 是一个链式函数，
// 它在下一次请求失败时设置重试次数和间隔。
func (c *Client) X重试与间隔(重试次数 int, 重试间隔时长 time.Duration) *Client {
	newClient := c.X取副本()
	newClient.X设置重试与间隔(重试次数, 重试间隔时长)
	return newClient
}

// Proxy 是一个链式函数，
// 用于设置下一次请求的代理。
// 请确保传递正确的 `proxyURL`。
// 正确的格式应为 `http://USER:PASSWORD@IP:PORT` 或 `socks5://USER:PASSWORD@IP:PORT`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。
func (c *Client) X代理(代理地址 string) *Client {
	newClient := c.X取副本()
	newClient.X设置代理(代理地址)
	return newClient
}

// RedirectLimit 是一个链式函数，
// 用于设置请求的重定向限制次数。
func (c *Client) X重定向次数限制(次数 int) *Client {
	newClient := c.X取副本()
	newClient.X设置重定向次数限制(次数)
	return newClient
}

// NoUrlEncode 设置标记，表示在发送请求前不应对参数进行URL编码。
func (c *Client) X请求参数禁止URL编码() *Client {
	newClient := c.X取副本()
	newClient.X设置请求参数禁止URL编码(true)
	return newClient
}
