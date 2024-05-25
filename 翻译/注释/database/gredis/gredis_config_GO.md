
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
// Config is redis configuration.
<原文结束>

# <翻译开始>
// Config 是 Redis 配置。 md5:5409b3144db1576a
# <翻译结束>


<原文开始>
// Address It supports single and cluster redis server. Multiple addresses joined with char ','. Eg: 192.168.1.1:6379, 192.168.1.2:6379.
<原文结束>

# <翻译开始>
// Address 支持单个和集群 Redis 服务器。多个地址使用逗号分隔。例如：192.168.1.1:6379, 192.168.1.2:6379。 md5:21ac53e24210b32c
# <翻译结束>


<原文开始>
// Username for sentinel AUTH.
<原文结束>

# <翻译开始>
// 防卫者AUTH的用户名。 md5:c85c5044b04f7ec3
# <翻译结束>


<原文开始>
// Password for sentinel AUTH.
<原文结束>

# <翻译开始>
// 密码，用于sentinel的AUTH。 md5:517cbd39fc9e4f20
# <翻译结束>


<原文开始>
// Minimum number of connections allowed to be idle (default is 0)
<原文结束>

# <翻译开始>
// 允许处于空闲状态的连接的最小数量（默认为0）. md5:534a8e485a7c5664
# <翻译结束>


<原文开始>
// Maximum number of connections allowed to be idle (default is 10)
<原文结束>

# <翻译开始>
// 允许的最大空闲连接数（默认为10）. md5:6b33b39ddbb7c42b
# <翻译结束>


<原文开始>
// Maximum number of connections limit (default is 0 means no limit).
<原文结束>

# <翻译开始>
// 连接数的最大限制（默认为0表示无限制）。 md5:4dbd7ce4d80b4597
# <翻译结束>


<原文开始>
// Maximum lifetime of the connection (default is 30 seconds, not allowed to be set to 0)
<原文结束>

# <翻译开始>
// 连接的最大生命周期（默认为30秒，不允许设置为0）. md5:1650bf54f8065411
# <翻译结束>


<原文开始>
// Maximum idle time for connection (default is 10 seconds, not allowed to be set to 0)
<原文结束>

# <翻译开始>
// 连接的最大空闲时间（默认为10秒，不允许设置为0）. md5:1f9346d51eb9e76a
# <翻译结束>


<原文开始>
// Timed out duration waiting to get a connection from the connection pool.
<原文结束>

# <翻译开始>
// 超时等待从连接池获取连接的持续时间。 md5:ff75b0772da43843
# <翻译结束>


<原文开始>
// Dial connection timeout for TCP.
<原文结束>

# <翻译开始>
// TCP连接的超时时间。 md5:d900304d9b7c2e58
# <翻译结束>


<原文开始>
// Read timeout for TCP. DO NOT set it if not necessary.
<原文结束>

# <翻译开始>
// TCP读超时。如果没有必要，请不要设置它。 md5:44e33a5ef46ceb97
# <翻译结束>


<原文开始>
// Used in Redis Sentinel mode.
<原文结束>

# <翻译开始>
// 用于Redis哨兵模式。 md5:44b4d0f3813a15e6
# <翻译结束>


<原文开始>
// Specifies whether TLS should be used when connecting to the server.
<原文结束>

# <翻译开始>
// 指定连接到服务器时是否应使用TLS。 md5:0b36620d5b0321dd
# <翻译结束>


<原文开始>
// Disables server name verification when connecting over TLS.
<原文结束>

# <翻译开始>
// 连接TLS时禁用服务器名称验证。 md5:3bde87f1295352e9
# <翻译结束>


<原文开始>
// TLS Config to use. When set TLS will be negotiated.
<原文结束>

# <翻译开始>
// 使用的TLS配置。如果设置，将协商TLS。 md5:b5a9a25bb2762b0b
# <翻译结束>


<原文开始>
// Route all commands to slave read-only nodes.
<原文结束>

# <翻译开始>
// 将所有命令路由到从属只读节点。 md5:9ba156f404a631f1
# <翻译结束>


<原文开始>
// Specifies whether cluster mode be used.
<原文结束>

# <翻译开始>
// 指定是否使用集群模式。 md5:7952648d7b1da3f9
# <翻译结束>


<原文开始>
// Specifies the RESP version (Protocol 2 or 3.)
<原文结束>

# <翻译开始>
// 定义RESP版本（协议2或3）。 md5:dbc1edfd3b1e3b35
# <翻译结束>


<原文开始>
// Default configuration group name.
<原文结束>

# <翻译开始>
// 默认配置组名称。 md5:eb4945d78061d92a
# <翻译结束>


<原文开始>
// SetConfig sets the global configuration for specified group.
// If `name` is not passed, it sets configuration for the default group name.
<原文结束>

# <翻译开始>
// SetConfig 为指定的组设置全局配置。
// 如果没有传递 `name`，则为默认组名设置配置。
// md5:8d7c1f181c0057f0
# <翻译结束>


<原文开始>
// SetConfigByMap sets the global configuration for specified group with map.
// If `name` is not passed, it sets configuration for the default group name.
<原文结束>

# <翻译开始>
// SetConfigByMap 使用映射设置指定组的全局配置。
// 如果未传递`name`，则将配置设置为默认组名。
// md5:1d191bb426ab05fb
# <翻译结束>


<原文开始>
// ConfigFromMap parses and returns config from given map.
<原文结束>

# <翻译开始>
// ConfigFromMap 从给定的映射中解析并返回配置。 md5:105a2224aed53bc9
# <翻译结束>


<原文开始>
// GetConfig returns the global configuration with specified group name.
// If `name` is not passed, it returns configuration of the default group name.
<原文结束>

# <翻译开始>
// GetConfig 返回指定组名的全局配置。如果未传入 `name`，则返回默认组名的配置。
// md5:327a839e91668442
# <翻译结束>


<原文开始>
// RemoveConfig removes the global configuration with specified group.
// If `name` is not passed, it removes configuration of the default group name.
<原文结束>

# <翻译开始>
// RemoveConfig 删除指定组的全局配置。
// 如果没有传递 `name`，则删除默认组名的配置。
// md5:8e808827f299122d
# <翻译结束>


<原文开始>
// ClearConfig removes all configurations of redis.
<原文结束>

# <翻译开始>
// ClearConfig 删除所有的 Redis 配置。 md5:337bf67372d51962
# <翻译结束>

