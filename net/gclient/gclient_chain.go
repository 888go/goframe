// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 网页类

import (
	"time"

	"github.com/888go/goframe/net/gsvc"
)

// Url前缀是一个链式函数，
// 它为该客户端的下一个请求设置URL前缀。例如：
// Url前缀("http://127.0.0.1:8199/api/v1")
// Url前缀("http://127.0.0.1:8199/api/v2")
// md5:223b00b33201dec6
func (c *Client) Url前缀(前缀 string) *Client {
	newClient := c.X取副本()
	newClient.X设置url前缀(前缀)
	return newClient
}

// X协议头是一个链式函数，它为下一个请求设置自定义HTTP头，使用映射（map）来存储头信息。
// md5:e9d44530a2916792
func (c *Client) X协议头(map协议头 map[string]string) *Client {
	newClient := c.X取副本()
	newClient.X设置Map协议头(map协议头)
	return newClient
}

// X原始协议头 是一个链式函数，
// 用于为下一个请求使用原始字符串设置自定义的HTTP头。
// md5:dff868fea39b9738
func (c *Client) X原始协议头(原始协议头 string) *Client {
	newClient := c.X取副本()
	newClient.X设置原始协议头(原始协议头)
	return newClient
}

// Discovery 是一个链式函数，用于设置客户端的发现机制。
// 你可以使用 `Discovery(nil)` 来禁用当前客户端的发现功能。
// md5:75d0bac47755ed2a
func (c *Client) Discovery(discovery gsvc.Discovery) *Client {
	newClient := c.X取副本()
	newClient.SetDiscovery(discovery)
	return newClient
}

// Cookie是一个链式函数，它使用映射为下一个请求设置Cookie项。
// md5:458381c5ded10dae
func (c *Client) Cookie(MapCookie map[string]string) *Client {
	newClient := c.X取副本()
	newClient.X设置CookieMap(MapCookie)
	return newClient
}

// X内容类型是一个链式函数，用于为下一个请求设置HTTP内容类型。
// md5:48b017c4d22a94ee
func (c *Client) X内容类型(contentType string) *Client {
	newClient := c.X取副本()
	newClient.X设置内容类型(contentType)
	return newClient
}

// X内容类型json 是一个链式函数，
// 用于为下一个请求将HTTP内容类型设置为 "application/json"。
//
// 注意，它还会自动检查并以JSON格式对参数进行编码。
// md5:3b265101262bf8be
func (c *Client) X内容类型json() *Client {
	newClient := c.X取副本()
	newClient.X设置内容类型(httpHeaderContentTypeJson)
	return newClient
}

// X内容类型xml是一个链式函数，
// 它将下一个请求的HTTP内容类型设置为"application/xml"。
//
// 请注意，它还会自动检查并把参数编码为XML格式。
// md5:b6e6ad51561fa7a6
func (c *Client) X内容类型xml() *Client {
	newClient := c.X取副本()
	newClient.X设置内容类型(httpHeaderContentTypeXml)
	return newClient
}

// X超时是一个链式函数，它为下一个请求设置超时时间。
// md5:8198ec107ce0113f
func (c *Client) X超时(时长 time.Duration) *Client {
	newClient := c.X取副本()
	newClient.X设置超时(时长)
	return newClient
}

// X账号密码是一个链式函数，为下一个请求设置HTTP基本认证信息。
// md5:c0ca33e41692898b
func (c *Client) X账号密码(账号, 密码 string) *Client {
	newClient := c.X取副本()
	newClient.X设置账号密码(账号, 密码)
	return newClient
}

// X重试与间隔 是一个链式函数，
// 用于设置在下次请求失败时的重试次数和间隔时间。
// TODO 待移除。
// md5:2206b70379c77ed7
func (c *Client) X重试与间隔(重试次数 int, 重试间隔时长 time.Duration) *Client {
	newClient := c.X取副本()
	newClient.X设置重试与间隔(重试次数, 重试间隔时长)
	return newClient
}

// X代理 是一个链式函数，
// 用于为下一个请求设置代理。
// 确保传递正确的 `proxyURL`。
// 正确的格式应为 `http://USER:PASSWORD@IP:PORT` 或 `socks5://USER:PASSWORD@IP:PORT`。
// 目前仅支持 `http` 和 `socks5` 类型的代理。
// md5:d8c660f7a12e72ea
func (c *Client) X代理(代理地址 string) *Client {
	newClient := c.X取副本()
	newClient.X设置代理(代理地址)
	return newClient
}

// X重定向次数限制 是一个链式函数，用于设置请求的重定向限制次数。
// md5:ecac398510aa92bb
func (c *Client) X重定向次数限制(次数 int) *Client {
	newClient := c.X取副本()
	newClient.X设置重定向次数限制(次数)
	return newClient
}

// X请求参数禁止URL编码 设置一个标记，表示在发送请求之前不进行参数编码。 md5:0f78cc83f0909b0e
func (c *Client) X请求参数禁止URL编码() *Client {
	newClient := c.X取副本()
	newClient.X设置请求参数禁止URL编码(true)
	return newClient
}
