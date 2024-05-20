
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
// gracefulServer wraps the net/http.Server with graceful reload/restart feature.
<原文结束>

# <翻译开始>
// gracefulServer 是一个包装了 net/http.Server 的结构，添加了优雅的重新加载/重启功能。. md5:8d812c91a33cd2a2
# <翻译结束>


<原文开始>
// File descriptor for passing to the child process when graceful reload.
<原文结束>

# <翻译开始>
// 用于在优雅重启时传递给子进程的文件描述符。. md5:72ea9b448b106b41
# <翻译结束>


<原文开始>
// Listening address like:":80", ":8080".
<原文结束>

# <翻译开始>
// 监听地址，例如":80"，":8080"。. md5:c746ec22043cf3e0
# <翻译结束>


<原文开始>
// Underlying http.Server.
<原文结束>

# <翻译开始>
// 底层的http.Server。. md5:3b44f2da7272f7f3
# <翻译结束>


<原文开始>
// Underlying net.Listener.
<原文结束>

# <翻译开始>
// 底层的net.Listener。. md5:95d2f6c4d9084a5b
# <翻译结束>


<原文开始>
// Concurrent safety mutex for `rawListener`.
<原文结束>

# <翻译开始>
// 为`rawListener`提供并发安全的互斥锁。. md5:7b358a2cf029baae
# <翻译结束>


<原文开始>
// Status of current server. Using `gtype` to ensure concurrent safety.
<原文结束>

# <翻译开始>
// 当前服务器的状态。使用 `gtype` 确保并发安全。. md5:d11344d5afa40f3a
# <翻译结束>


<原文开始>
// newGracefulServer creates and returns a graceful http server with a given address.
// The optional parameter `fd` specifies the file descriptor which is passed from parent server.
<原文结束>

# <翻译开始>
// newGracefulServer 创建并返回一个给定地址的优雅HTTP服务器。
// 可选参数 `fd` 指定了从父服务器传递过来的文件描述符。
// md5:e7000c344ed0446f
# <翻译结束>


<原文开始>
// Change port to address like: 80 -> :80
<原文结束>

# <翻译开始>
// 将端口转换为地址形式，如：80 -> :80. md5:71e59572a00dec96
# <翻译结束>


<原文开始>
// newHttpServer creates and returns an underlying http.Server with a given address.
<原文结束>

# <翻译开始>
// newHttpServer 创建并返回一个带有给定地址的底层 http.Server。. md5:12a45a5b95a4e7c3
# <翻译结束>


<原文开始>
// Fd retrieves and returns the file descriptor of the current server.
// It is available ony in *nix like operating systems like linux, unix, darwin.
<原文结束>

# <翻译开始>
// Fd获取并返回当前服务器的文件描述符。它只在*nix类操作系统中可用，如Linux、Unix和Darwin。
// md5:40546fed24d791cd
# <翻译结束>


<原文开始>
// CreateListener creates listener on configured address.
<原文结束>

# <翻译开始>
// CreateListener 在配置的地址上创建监听器。. md5:89f8795cf6b796f9
# <翻译结束>


<原文开始>
// CreateListenerTLS creates listener on configured address with HTTPS.
// The parameter `certFile` and `keyFile` specify the necessary certification and key files for HTTPS.
// The optional parameter `tlsConfig` specifies the custom TLS configuration.
<原文结束>

# <翻译开始>
// CreateListenerTLS 在配置的地址上创建使用HTTPS的监听器。
// 参数`certFile`和`keyFile`指定了HTTPS所需的证书和密钥文件。
// 可选参数`tlsConfig`指定自定义的TLS配置。
// md5:04f46f61853037ca
# <翻译结束>


<原文开始>
// Serve starts the serving with blocking way.
<原文结束>

# <翻译开始>
// Serve以阻塞方式启动服务。. md5:230e5731ffa3d482
# <翻译结束>


<原文开始>
// GetListenedAddress retrieves and returns the address string which are listened by current server.
<原文结束>

# <翻译开始>
// GetListenedAddress 获取并返回当前服务器所监听的地址字符串。. md5:51d352ffec9dc329
# <翻译结束>


<原文开始>
// GetListenedPort retrieves and returns one port which is listened to by current server.
// Note that this method is only available if the server is listening on one port.
<原文结束>

# <翻译开始>
// GetListenedPort 获取并返回当前服务器正在监听的其中一个端口。
// 注意，如果服务器只监听一个端口，则此方法才可用。
// md5:2fe5eae2317fe8f9
# <翻译结束>


<原文开始>
// getProto retrieves and returns the proto string of current server.
<原文结束>

# <翻译开始>
// getProto 获取并返回当前服务器的proto字符串。. md5:7860227f594f2ca9
# <翻译结束>


<原文开始>
// getNetListener retrieves and returns the wrapped net.Listener.
<原文结束>

# <翻译开始>
// getNetListener 获取并返回包装的net.Listener。. md5:36d0b8cf9a591408
# <翻译结束>


<原文开始>
// shutdown shuts down the server gracefully.
<原文结束>

# <翻译开始>
// shutdown 停止服务器，优雅地关闭。. md5:6befce727da40eb9
# <翻译结束>


<原文开始>
// setRawListener sets `rawListener` with given net.Listener.
<原文结束>

# <翻译开始>
// 设置RawListener，将给定的net.Listener设置为`rawListener`。. md5:0fe9b7938ed0a876
# <翻译结束>


<原文开始>
// setRawListener returns the `rawListener` of current server.
<原文结束>

# <翻译开始>
// setRawListener 返回当前服务器的 `rawListener`。. md5:e7b9cd54708d26f8
# <翻译结束>


<原文开始>
// close shuts down the server forcibly.
<原文结束>

# <翻译开始>
// close 强制关闭服务器。. md5:46634188c0dbdf78
# <翻译结束>

