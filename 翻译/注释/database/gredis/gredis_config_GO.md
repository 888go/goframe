
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
// Config is redis configuration.
<原文结束>

# <翻译开始>
// Config 是 Redis 配置。
# <翻译结束>


<原文开始>
// Address It supports single and cluster redis server. Multiple addresses joined with char ','. Eg: 192.168.1.1:6379, 192.168.1.2:6379.
<原文结束>

# <翻译开始>
// Address 地址 支持单个和集群模式的Redis服务器。多个地址使用逗号 ',' 连接。例如：192.168.1.1:6379, 192.168.1.2:6379。
# <翻译结束>












<原文开始>
// Minimum number of connections allowed to be idle (default is 0)
<原文结束>

# <翻译开始>
// 最小允许空闲的连接数（默认为0）
# <翻译结束>


<原文开始>
// Maximum number of connections allowed to be idle (default is 10)
<原文结束>

# <翻译开始>
// 最大允许空闲连接数（默认为10）
# <翻译结束>


<原文开始>
// Maximum number of connections limit (default is 0 means no limit).
<原文结束>

# <翻译开始>
// 连接数的最大限制（默认值为0，表示无限制）。
# <翻译结束>


<原文开始>
// Maximum lifetime of the connection (default is 30 seconds, not allowed to be set to 0)
<原文结束>

# <翻译开始>
// 连接的最大生命周期（默认为30秒，不允许设置为0）
# <翻译结束>


<原文开始>
// Maximum idle time for connection (default is 10 seconds, not allowed to be set to 0)
<原文结束>

# <翻译开始>
// 连接的最大空闲时间（默认为10秒，不允许设置为0）
# <翻译结束>


<原文开始>
// Timed out duration waiting to get a connection from the connection pool.
<原文结束>

# <翻译开始>
// 从连接池获取连接时超时的持续时间。
# <翻译结束>


<原文开始>
// Dial connection timeout for TCP.
<原文结束>

# <翻译开始>
// 设置TCP连接的超时时间。
# <翻译结束>


<原文开始>
// Read timeout for TCP. DO NOT set it if not necessary.
<原文结束>

# <翻译开始>
// TCP读取超时时间。如果不是必需的，请勿设置。
# <翻译结束>







<原文开始>
// Used in Redis Sentinel mode.
<原文结束>

# <翻译开始>
// 用于 Redis Sentinel 模式。
# <翻译结束>


<原文开始>
// Specifies whether TLS should be used when connecting to the server.
<原文结束>

# <翻译开始>
// 指定在连接到服务器时是否应使用TLS（传输层安全协议）。
# <翻译结束>


<原文开始>
// Disables server name verification when connecting over TLS.
<原文结束>

# <翻译开始>
// 禁用通过TLS连接时的服务器名称验证。
# <翻译结束>


<原文开始>
// TLS Config to use. When set TLS will be negotiated.
<原文结束>

# <翻译开始>
// TLS配置使用。当设置此配置时，将进行TLS协商。
# <翻译结束>


<原文开始>
// Route all commands to slave read-only nodes.
<原文结束>

# <翻译开始>
// 将所有命令路由到从节点（只读模式）。
# <翻译结束>


<原文开始>
// Specifies whether cluster mode be used.
<原文结束>

# <翻译开始>
// 指定是否使用集群模式。
# <翻译结束>


<原文开始>
// Specifies the RESP version (Protocol 2 or 3.)
<原文结束>

# <翻译开始>
// 指定 RESP 协议版本（协议 2 或 3）
# <翻译结束>


<原文开始>
// Default configuration group name.
<原文结束>

# <翻译开始>
// 默认配置组名称。
# <翻译结束>







<原文开始>
// SetConfig sets the global configuration for specified group.
// If `name` is not passed, it sets configuration for the default group name.
<原文结束>

# <翻译开始>
// SetConfig 为指定的组设置全局配置。
// 如果未传递 `name`，则会为默认组名设置配置。
# <翻译结束>


<原文开始>
// SetConfigByMap sets the global configuration for specified group with map.
// If `name` is not passed, it sets configuration for the default group name.
<原文结束>

# <翻译开始>
// SetConfigByMap 通过map设置指定组的全局配置。
// 如果未传递 `name`，则设置默认组名的配置。
# <翻译结束>


<原文开始>
// ConfigFromMap parses and returns config from given map.
<原文结束>

# <翻译开始>
// ConfigFromMap 从给定的 map 中解析并返回配置。
# <翻译结束>


<原文开始>
// GetConfig returns the global configuration with specified group name.
// If `name` is not passed, it returns configuration of the default group name.
<原文结束>

# <翻译开始>
// GetConfig 返回具有指定组名的全局配置。
// 如果未传递 `name`，则返回默认组名的配置。
# <翻译结束>


<原文开始>
// RemoveConfig removes the global configuration with specified group.
// If `name` is not passed, it removes configuration of the default group name.
<原文结束>

# <翻译开始>
// RemoveConfig 删除指定组的全局配置。
// 如果未传入 `name`，则移除默认组名的配置。
# <翻译结束>


<原文开始>
// ClearConfig removes all configurations of redis.
<原文结束>

# <翻译开始>
// ClearConfig 清除所有redis配置。
# <翻译结束>


<原文开始>
// Username for AUTH.
<原文结束>

# <翻译开始>
// AUTH的用户名。
# <翻译结束>


<原文开始>
// Password for AUTH.
<原文结束>

# <翻译开始>
// AUTH的密码。
# <翻译结束>


<原文开始>
// Write timeout for TCP.
<原文结束>

# <翻译开始>
// TCP写入超时时间
# <翻译结束>


<原文开始>
// Configuration groups.
<原文结束>

# <翻译开始>
// 配置组。
# <翻译结束>

