
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>












<原文开始>
// Client chaining operations handlers.
<原文结束>

# <翻译开始>
// 客户端链式操作处理程序。
# <翻译结束>







<原文开始>
// controls the maximum idle(keep-alive) connections to keep per-host
<原文结束>

# <翻译开始>
// 控制每个主机保持的最大空闲(keep-alive)连接数
# <翻译结束>












<原文开始>
// Send with string parameter in request body.
<原文结束>

# <翻译开始>
// 使用字符串参数作为请求体发送。
# <翻译结束>







<原文开始>
// Send with string parameter along with URL.
<原文结束>

# <翻译开始>
// 使用字符串参数并通过URL发送。
# <翻译结束>


<原文开始>
// ExampleClient_SetProxy an example for `gclient.Client.SetProxy` method.
// please prepare two proxy server before running this example.
// http proxy server listening on `127.0.0.1:1081`
// socks5 proxy server listening on `127.0.0.1:1080`
<原文结束>

# <翻译开始>
// ExampleClient_SetProxy 是 `gclient.Client.SetProxy` 方法的示例。
// 请在运行此示例前准备好两个代理服务器。
// HTTP 代理服务器监听 `127.0.0.1:1081`
// SOCKS5 代理服务器监听 `127.0.0.1:1080`
# <翻译结束>


<原文开始>
// connect to an http proxy server
<原文结束>

# <翻译开始>
// 连接到一个HTTP代理服务器
# <翻译结束>


<原文开始>
		// err is not nil when your proxy server is down.
		// eg. Get "http://127.0.0.1:8999": proxyconnect tcp: dial tcp 127.0.0.1:1087: connect: connection refused
<原文结束>

# <翻译开始>
// 当你的代理服务器无法访问时，err将不为nil。
// 例如：尝试获取"http://127.0.0.1:8999"时，出现如下错误：
//       通过代理连接tcp：拨号tcp到127.0.0.1:1087时出错：连接被拒绝
// （注：原文中的 "proxyconnect tcp: dial tcp" 是golang中通过代理连接远程地址时可能出现的错误信息，意指在建立TCP连接至代理服务器的过程中出现了问题，具体表现为连接被拒绝。）
# <翻译结束>


<原文开始>
// connect to an http proxy server which needs auth
<原文结束>

# <翻译开始>
// 连接到需要身份验证的HTTP代理服务器
# <翻译结束>


<原文开始>
// connect to a socks5 proxy server
<原文结束>

# <翻译开始>
// 连接到一个SOCKS5代理服务器
# <翻译结束>


<原文开始>
		// err is not nil when your proxy server is down.
		// eg. Get "http://127.0.0.1:8999": socks connect tcp 127.0.0.1:1087->api.ip.sb:443: dial tcp 127.0.0.1:1087: connect: connection refused
<原文结束>

# <翻译开始>
// 当您的代理服务器不可用时，err将不为nil。
// 例如：获取"http://127.0.0.1:8999"时出错，错误信息为：
// "socks连接tcp 127.0.0.1:1087->api.ip.sb:443时发生错误：拨号tcp 127.0.0.1:1087失败，原因：连接被拒绝"
// 这段Go语言代码的注释翻译成中文是：
// ```go
// 如果您的代理服务器宕机，err将不会是nil（即会返回一个非空错误）。
// 比如：尝试获取"http://127.0.0.1:8999"时，
// 出现如下错误："通过SOCKS协议连接到tcp 127.0.0.1:1087并试图转发到api.ip.sb:443的过程中发生错误，
// 具体原因为：尝试连接到tcp 127.0.0.1:1087时被拒绝，连接无法建立。"
# <翻译结束>


<原文开始>
// connect to a socks5 proxy server which needs auth
<原文结束>

# <翻译开始>
// 连接到需要身份验证的SOCKS5代理服务器
# <翻译结束>


<原文开始>
// ExampleClientChain_Proxy a chain version of example for `gclient.Client.Proxy` method.
// please prepare two proxy server before running this example.
// http proxy server listening on `127.0.0.1:1081`
// socks5 proxy server listening on `127.0.0.1:1080`
// for more details, please refer to ExampleClient_SetProxy
<原文结束>

# <翻译开始>
// ExampleClientChain_Proxy 是 `gclient.Client.Proxy` 方法的链式版本示例。
// 在运行此示例前，请确保已准备好两个代理服务器。
// HTTP 代理服务器监听地址为 `127.0.0.1:1081`
// SOCKS5 代理服务器监听地址为 `127.0.0.1:1080`
// 更多详情，请参考 ExampleClient_SetProxy 示例。
# <翻译结束>







<原文开始>
// Default server for client.
<原文结束>

# <翻译开始>
// 默认的客户端服务器。
# <翻译结束>


<原文开始>
// HTTP method handlers.
<原文结束>

# <翻译开始>
// HTTP方法处理器。
# <翻译结束>


<原文开始>
// Other testing handlers.
<原文结束>

# <翻译开始>
// 其他测试处理程序。
# <翻译结束>


<原文开始>
// Post using JSON string.
<原文结束>

# <翻译开始>
// 使用JSON字符串进行POST请求。
# <翻译结束>


<原文开始>
// Post using JSON map.
<原文结束>

# <翻译开始>
// 使用JSON映射进行POST请求。
# <翻译结束>


<原文开始>
// Send with map parameter.
<原文结束>

# <翻译开始>
// 使用map参数发送。
# <翻译结束>


<原文开始>
// Add Client URI Prefix
<原文结束>

# <翻译开始>
// 添加客户端URI前缀
# <翻译结束>

