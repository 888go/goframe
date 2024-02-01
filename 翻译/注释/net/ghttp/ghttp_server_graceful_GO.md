
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
// gracefulServer wraps the net/http.Server with graceful reload/restart feature.
<原文结束>

# <翻译开始>
// gracefulServer 将 net/http.Server 包装起来，为其提供优雅的重新加载/重启功能。
# <翻译结束>


<原文开始>
// File descriptor for passing to the child process when graceful reload.
<原文结束>

# <翻译开始>
// 用于在优雅重启时传递给子进程的文件描述符。
# <翻译结束>

















<原文开始>
// newGracefulServer creates and returns a graceful http server with a given address.
// The optional parameter `fd` specifies the file descriptor which is passed from parent server.
<原文结束>

# <翻译开始>
// newGracefulServer 根据给定的地址创建并返回一个优雅的 HTTP 服务器。
// 可选参数 `fd` 指定了从父服务器传递过来的文件描述符。
# <翻译结束>


<原文开始>
// Change port to address like: 80 -> :80
<原文结束>

# <翻译开始>
// 将端口更改为地址形式，例如：80 -> :80
# <翻译结束>


<原文开始>
// newHttpServer creates and returns an underlying http.Server with a given address.
<原文结束>

# <翻译开始>
// newHttpServer 根据给定的地址创建并返回一个底层的 http.Server。
# <翻译结束>


<原文开始>
// Fd retrieves and returns the file descriptor of the current server.
// It is available ony in *nix like operating systems like linux, unix, darwin.
<原文结束>

# <翻译开始>
// Fd 获取并返回当前服务器的文件描述符。
// 该功能仅在类*nix操作系统中可用，如linux、unix、darwin。
# <翻译结束>


<原文开始>
// CreateListener creates listener on configured address.
<原文结束>

# <翻译开始>
// CreateListener 在配置的地址上创建监听器。
# <翻译结束>


<原文开始>
// CreateListenerTLS creates listener on configured address with HTTPS.
// The parameter `certFile` and `keyFile` specify the necessary certification and key files for HTTPS.
// The optional parameter `tlsConfig` specifies the custom TLS configuration.
<原文结束>

# <翻译开始>
// CreateListenerTLS 在配置的地址上创建 HTTPS 侦听器。
// 参数 `certFile` 和 `keyFile` 指定用于 HTTPS 的必要证书和密钥文件。
// 可选参数 `tlsConfig` 指定自定义 TLS 配置。
# <翻译结束>


<原文开始>
// Serve starts the serving with blocking way.
<原文结束>

# <翻译开始>
// Serve 以阻塞方式启动服务。
# <翻译结束>


<原文开始>
// GetListenedAddress retrieves and returns the address string which are listened by current server.
<原文结束>

# <翻译开始>
// GetListenedAddress 获取并返回当前服务器监听的地址字符串。
# <翻译结束>


<原文开始>
// GetListenedPort retrieves and returns one port which is listened to by current server.
// Note that this method is only available if the server is listening on one port.
<原文结束>

# <翻译开始>
// GetListenedPort 获取并返回当前服务器监听的端口。
// 注意：只有当服务器正在监听单个端口时，此方法才可用。
# <翻译结束>


<原文开始>
// getProto retrieves and returns the proto string of current server.
<原文结束>

# <翻译开始>
// getProto 获取并返回当前服务器的协议字符串。
# <翻译结束>


<原文开始>
// getNetListener retrieves and returns the wrapped net.Listener.
<原文结束>

# <翻译开始>
// getNetListener 获取并返回封装后的 net.Listener。
# <翻译结束>


<原文开始>
// shutdown shuts down the server gracefully.
<原文结束>

# <翻译开始>
// shutdown优雅地关闭服务器。
# <翻译结束>


<原文开始>
// setRawListener sets `rawListener` with given net.Listener.
<原文结束>

# <翻译开始>
// setRawListener 将给定的 net.Listener 设置到 `rawListener`。
# <翻译结束>


<原文开始>
// setRawListener returns the `rawListener` of current server.
<原文结束>

# <翻译开始>
// setRawListener 返回当前服务器的 `原始监听器`。
# <翻译结束>


<原文开始>
// close shuts down the server forcibly.
<原文结束>

# <翻译开始>
// close 强制关闭服务器。
# <翻译结束>







<原文开始>
// Listening address like:":80", ":8080".
<原文结束>

# <翻译开始>
// 监听地址格式如":80"、":8080"。
# <翻译结束>


<原文开始>
// Concurrent safety mutex for `rawListener`.
<原文结束>

# <翻译开始>
// `rawListener`的并发安全互斥锁。
# <翻译结束>


<原文开始>
// Status of current server. Using `gtype` to ensure concurrent safety.
<原文结束>

# <翻译开始>
// 当前服务器状态，使用 `gtype` 以确保并发安全性。
# <翻译结束>


<原文开始>
// Underlying http.Server.
<原文结束>

# <翻译开始>
// 底层的 http.Server.
# <翻译结束>


<原文开始>
// Underlying net.Listener.
<原文结束>

# <翻译开始>
// 基础的 net.Listener.
# <翻译结束>


<原文开始>
// Wrapped net.Listener.
<原文结束>

# <翻译开始>
// 包装过的 net.Listener。
# <翻译结束>

