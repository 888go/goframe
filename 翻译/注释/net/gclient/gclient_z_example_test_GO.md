
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Default server for client.
<原文结束>

# <翻译开始>
	// 客户端的默认服务器。 md5:2b3306283554596f
# <翻译结束>


<原文开始>
// Client chaining operations handlers.
<原文结束>

# <翻译开始>
	// 用于客户端链式操作的处理器。 md5:7a613ac703db33dd
# <翻译结束>


<原文开始>
// Other testing handlers.
<原文结束>

# <翻译开始>
	// 其他测试处理程序。 md5:99df94400fbb41dc
# <翻译结束>


<原文开始>
// controls the maximum idle(keep-alive) connections to keep per-host
<原文结束>

# <翻译开始>
	// 控制每个主机的最大闲置（保持活动）连接数. md5:71b53159157ddb6e
# <翻译结束>


<原文开始>
// Post using JSON string.
<原文结束>

# <翻译开始>
	// 使用JSON字符串进行POST操作。 md5:4d52d60dd39bd628
# <翻译结束>


<原文开始>
// Send with string parameter in request body.
<原文结束>

# <翻译开始>
	// 使用字符串参数作为请求体发送。 md5:ba68880cfea93a12
# <翻译结束>


<原文开始>
// Send with map parameter.
<原文结束>

# <翻译开始>
	// 使用map参数发送。 md5:270768ac9382ef2b
# <翻译结束>


<原文开始>
// Send with string parameter along with URL.
<原文结束>

# <翻译开始>
	// 使用字符串参数和URL一起发送。 md5:0fae209daa2970ad
# <翻译结束>


<原文开始>
// ExampleClient_SetProxy an example for `gclient.Client.SetProxy` method.
// please prepare two proxy server before running this example.
// http proxy server listening on `127.0.0.1:1081`
// socks5 proxy server listening on `127.0.0.1:1080`
<原文结束>

# <翻译开始>
// ExampleClient_SetProxy 是 `gclient.Client.SetProxy` 方法的一个示例。
// 在运行这个示例之前，请准备两个代理服务器。
// HTTP 代理服务器监听 `127.0.0.1:1081`
// Socks5 代理服务器监听 `127.0.0.1:1080`
// md5:c26527766902fe18
# <翻译结束>


<原文开始>
// connect to an http proxy server
<原文结束>

# <翻译开始>
	// 连接到一个HTTP代理服务器. md5:62686127819e0038 which needs auth
# <翻译结束>


<原文开始>
		// err is not nil when your proxy server is down.
		// eg. Get "http://127.0.0.1:8999": proxyconnect tcp: dial tcp 127.0.0.1:1087: connect: connection refused
<原文结束>

# <翻译开始>
		// 当您的代理服务器不可用时，err 不为 nil。
		// 例如：获取 "http:		//127.0.0.1:8999" 时：proxyconnect tcp: 拨打 tcp 127.0.0.1:1087: 连接被拒绝
		// md5:51c9b1789e6b5346
# <翻译结束>


<原文开始>
// connect to an http proxy server which needs auth
<原文结束>

# <翻译开始>
// 连接到需要身份验证的HTTP代理服务器. md5:ff9335c208e5b72f
# <翻译结束>


<原文开始>
// connect to a socks5 proxy server
<原文结束>

# <翻译开始>
	// 连接到一个SOCKS5代理服务器. md5:51f0ad95ea53343f which needs auth
# <翻译结束>


<原文开始>
		// err is not nil when your proxy server is down.
		// eg. Get "http://127.0.0.1:8999": socks connect tcp 127.0.0.1:1087->api.ip.sb:443: dial tcp 127.0.0.1:1087: connect: connection refused
<原文结束>

# <翻译开始>
		// 当你的代理服务器不可用时，err不为nil。
		// 例如：Get "http:		//127.0.0.1:8999"：socks connect tcp 127.0.0.1:1087->api.ip.sb:443: 连接 tcp 127.0.0.1:1087：连接拒绝。
		// md5:f6d9173b84667e10
# <翻译结束>


<原文开始>
// connect to a socks5 proxy server which needs auth
<原文结束>

# <翻译开始>
// 连接到需要身份验证的SOCKS5代理服务器. md5:e17dd954ebc24894
# <翻译结束>


<原文开始>
// ExampleClient_Proxy a chain version of example for `gclient.Client.Proxy` method.
// please prepare two proxy server before running this example.
// http proxy server listening on `127.0.0.1:1081`
// socks5 proxy server listening on `127.0.0.1:1080`
// for more details, please refer to ExampleClient_SetProxy
<原文结束>

# <翻译开始>
// ExampleClient_Proxy 是一个`gclient.Client.Proxy`方法的链式版本示例。
// 在运行此示例之前，请准备两个代理服务器。
// HTTP 代理服务器监听 `127.0.0.1:1081`
// SOCKS5 代理服务器监听 `127.0.0.1:1080`
// 更多详细信息，请参考 ExampleClient_SetProxy
// md5:4d9e0da3aa8a180d
# <翻译结束>

