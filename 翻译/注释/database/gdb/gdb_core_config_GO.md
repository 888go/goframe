
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
// Config is the configuration management object.
<原文结束>

# <翻译开始>
// Config 是配置管理对象。
# <翻译结束>


<原文开始>
// ConfigGroup is a slice of configuration node for specified named group.
<原文结束>

# <翻译开始>
// ConfigGroup 是为指定名称组配置的配置节点切片。
# <翻译结束>


<原文开始>
// ConfigNode is configuration for one node.
<原文结束>

# <翻译开始>
// ConfigNode 是单个节点的配置。
# <翻译结束>


<原文开始>
// Host of server, ip or domain like: 127.0.0.1, localhost
<原文结束>

# <翻译开始>
// 服务器主机，可以是IP地址或域名，例如：127.0.0.1、localhost
# <翻译结束>












<原文开始>
// Default used database name.
<原文结束>

# <翻译开始>
// 默认使用的数据库名称。
# <翻译结束>


<原文开始>
// Database type: mysql, sqlite, mssql, pgsql, oracle.
<原文结束>

# <翻译开始>
// 数据库类型：mysql，sqlite，mssql，pgsql，oracle。
# <翻译结束>


<原文开始>
// (Optional) Custom link information for all configuration in one single string.
<原文结束>

# <翻译开始>
// (可选) 将所有配置的自定义链接信息放在一个单独字符串中。
# <翻译结束>


<原文开始>
// (Optional) Extra configuration according the registered third-party database driver.
<原文结束>

# <翻译开始>
// (可选) 根据已注册的第三方数据库驱动进行额外配置
# <翻译结束>


<原文开始>
// (Optional) Debug mode enables debug information logging and output.
<原文结束>

# <翻译开始>
// (可选) Debug模式启用调试信息日志记录和输出。
# <翻译结束>







<原文开始>
// (Optional) Dry run, which does SELECT but no INSERT/UPDATE/DELETE statements.
<原文结束>

# <翻译开始>
// (可选) 干预运行，仅执行 SELECT 语句但不执行 INSERT/UPDATE/DELETE 语句。
# <翻译结束>


<原文开始>
// (Optional) Sets the time zone for displaying and interpreting time stamps.
<原文结束>

# <翻译开始>
// (可选) 设置显示和解释时间戳时区。
# <翻译结束>


<原文开始>
// (Optional) Max idle connection configuration for underlying connection pool.
<原文结束>

# <翻译开始>
// （可选）为底层连接池配置的最大空闲连接数。
# <翻译结束>


<原文开始>
// (Optional) Max open connection configuration for underlying connection pool.
<原文结束>

# <翻译开始>
// (可选) 用于底层连接池的最大打开连接配置。
# <翻译结束>


<原文开始>
// (Optional) Max amount of time a connection may be idle before being closed.
<原文结束>

# <翻译开始>
// （可选）在连接被关闭之前，允许其空闲的最大时长。
# <翻译结束>


<原文开始>
// (Optional) Max query time for per dql.
<原文结束>

# <翻译开始>
// (可选) 每个DQL的最大查询时间。
# <翻译结束>


<原文开始>
// (Optional) Max exec time for dml.
<原文结束>

# <翻译开始>
// （可选）dml的最大执行时间。
# <翻译结束>


<原文开始>
// (Optional) Max exec time for a transaction.
<原文结束>

# <翻译开始>
// (可选) 事务执行的最大时间。
# <翻译结束>


<原文开始>
// (Optional) Max exec time for prepare operation.
<原文结束>

# <翻译开始>
// （可选）为准备操作设置最大执行时间。
# <翻译结束>


<原文开始>
// (Optional) The field name of table for automatic-filled created datetime.
<原文结束>

# <翻译开始>
// (可选) 自动填充创建日期时间的表字段名称。
# <翻译结束>


<原文开始>
// (Optional) The field name of table for automatic-filled updated datetime.
<原文结束>

# <翻译开始>
// (可选) 自动填充更新日期时间的表格字段名称。
# <翻译结束>


<原文开始>
// (Optional) Disable the automatic time maintaining feature.
<原文结束>

# <翻译开始>
// (可选) 禁用自动时间维护功能。
# <翻译结束>







<原文开始>
// configs specifies internal used configuration object.
<原文结束>

# <翻译开始>
// configs 指定内部使用的配置对象。
# <翻译结束>







<原文开始>
// Default configuration group.
<原文结束>

# <翻译开始>
// 默认配置组。
# <翻译结束>


<原文开始>
// SetConfig sets the global configuration for package.
// It will overwrite the old configuration of package.
<原文结束>

# <翻译开始>
// SetConfig 设置包的全局配置。
// 它将覆盖包的旧配置。
# <翻译结束>


<原文开始>
// SetConfigGroup sets the configuration for given group.
<原文结束>

# <翻译开始>
// SetConfigGroup 为给定的组设置配置。
# <翻译结束>


<原文开始>
// AddConfigNode adds one node configuration to configuration of given group.
<原文结束>

# <翻译开始>
// AddConfigNode 向给定组的配置中添加一个节点配置。
# <翻译结束>


<原文开始>
// parseConfigNode parses `Link` configuration syntax.
<原文结束>

# <翻译开始>
// parseConfigNode 解析 `Link` 配置语法。
# <翻译结束>


<原文开始>
// AddDefaultConfigNode adds one node configuration to configuration of default group.
<原文结束>

# <翻译开始>
// AddDefaultConfigNode 将一个节点配置添加到默认组的配置中。
# <翻译结束>


<原文开始>
// AddDefaultConfigGroup adds multiple node configurations to configuration of default group.
<原文结束>

# <翻译开始>
// AddDefaultConfigGroup 向默认组的配置中添加多个节点配置。
# <翻译结束>


<原文开始>
// GetConfig retrieves and returns the configuration of given group.
<原文结束>

# <翻译开始>
// GetConfig 获取并返回给定组的配置。
# <翻译结束>


<原文开始>
// SetDefaultGroup sets the group name for default configuration.
<原文结束>

# <翻译开始>
// SetDefaultGroup 设置默认配置的组名称。
# <翻译结束>


<原文开始>
// GetDefaultGroup returns the { name of default configuration.
<原文结束>

# <翻译开始>
// GetDefaultGroup 返回默认配置的名称。
# <翻译结束>


<原文开始>
// IsConfigured checks and returns whether the database configured.
// It returns true if any configuration exists.
<原文结束>

# <翻译开始>
// IsConfigured 检查并返回数据库是否已配置。
// 如果存在任何配置信息，则返回 true。
# <翻译结束>


<原文开始>
// SetLogger sets the logger for orm.
<原文结束>

# <翻译开始>
// SetLogger 设置orm的记录器。
# <翻译结束>


<原文开始>
// GetLogger returns the (logger) of the orm.
<原文结束>

# <翻译开始>
// GetLogger 返回 orm 的（日志器）
# <翻译结束>


<原文开始>
// SetMaxIdleConnCount sets the maximum number of connections in the idle
// connection pool.
//
// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
//
// If n <= 0, no idle connections are retained.
//
// The default max idle connections is currently 2. This may change in
// a future release.
<原文结束>

# <翻译开始>
// 设置空闲连接池中的最大连接数。
//
// 如果MaxOpenConns大于0但小于新的MaxIdleConns，则新的MaxIdleConns将减小以匹配MaxOpenConns限制。
//
// 若n <= 0，则不保留任何空闲连接。
//
// 当前默认的最大空闲连接数为2，这在未来版本中可能会发生变化。
// 以下是逐行翻译：
// ```go
// SetMaxIdleConnCount 用于设置闲置连接池中允许的最大连接数量。
//
// 如果 MaxOpenConns 大于0但小于新设置的 MaxIdleConns 值，
// 那么新的 MaxIdleConns 将会被调整以匹配 MaxOpenConns 的限制。
//
// 如果传入的参数 n 小于等于0，则不会保留任何空闲连接。
//
// 目前默认的最大空闲连接数是2，在未来版本中这个数值可能会有所更改。
# <翻译结束>


<原文开始>
// SetMaxOpenConnCount sets the maximum number of open connections to the database.
//
// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
// MaxIdleConns, then MaxIdleConns will be reduced to match the new
// MaxOpenConns limit.
//
// If n <= 0, then there is no limit on the number of open connections.
// The default is 0 (unlimited).
<原文结束>

# <翻译开始>
// SetMaxOpenConnCount 设置与数据库的最大连接数。
//
// 如果 MaxIdleConns 大于0且新的 MaxOpenConns 值小于 MaxIdleConns，则 MaxIdleConns 会相应减少以匹配新的 MaxOpenConns 限制。
//
// 若 n <= 0，则表示对打开的连接数没有限制。
// 默认值为 0（无限制）。
# <翻译结束>


<原文开始>
// SetMaxConnLifeTime sets the maximum amount of time a connection may be reused.
//
// Expired connections may be closed lazily before reuse.
//
// If d <= 0, connections are not closed due to a connection's age.
<原文结束>

# <翻译开始>
// SetMaxConnLifeTime 设置单个连接可重用的最大时长。
//
// 到期的连接在重用前可能被延迟关闭。
//
// 如果 d <= 0，则不会因为连接存在时间过长而关闭连接。
# <翻译结束>


<原文开始>
// GetConfig returns the current used node configuration.
<原文结束>

# <翻译开始>
// GetConfig 返回当前正在使用的节点配置。
# <翻译结束>


<原文开始>
		// Note:
		// It so here checks and returns the config from current DB,
		// if different schemas between current DB and config.Name from context,
		// for example, in nested transaction scenario, the context is passed all through the logic procedure,
		// but the config.Name from context may be still the original one from the first transaction object.
<原文结束>

# <翻译开始>
// 注意：
// 该处会检查并返回当前数据库的配置信息，
// 如果当前数据库与从context获取到的config.Name中的模式（schema）不同时，
// 比如在嵌套事务场景中，context会在整个逻辑过程中被传递，
// 但context中的config.Name可能仍然是最初第一个事务对象中的原始值。
# <翻译结束>


<原文开始>
// SetDebug enables/disables the debug mode.
<原文结束>

# <翻译开始>
// SetDebug用于开启或关闭调试模式。
# <翻译结束>


<原文开始>
// GetDebug returns the debug value.
<原文结束>

# <翻译开始>
// GetDebug 返回调试值。
# <翻译结束>


<原文开始>
// GetCache returns the internal cache object.
<原文结束>

# <翻译开始>
// GetCache 返回内部缓存对象。
# <翻译结束>


<原文开始>
// GetGroup returns the group string configured.
<原文结束>

# <翻译开始>
// GetGroup 返回已配置的组字符串。
# <翻译结束>


<原文开始>
// SetDryRun enables/disables the DryRun feature.
<原文结束>

# <翻译开始>
// SetDryRun 用于启用/禁用 DryRun 功能。
# <翻译结束>


<原文开始>
// GetDryRun returns the DryRun value.
<原文结束>

# <翻译开始>
// GetDryRun 返回 DryRun 的值。
# <翻译结束>


<原文开始>
// GetPrefix returns the table prefix string configured.
<原文结束>

# <翻译开始>
// GetPrefix 返回已配置的表前缀字符串。
# <翻译结束>


<原文开始>
// GetSchema returns the schema configured.
<原文结束>

# <翻译开始>
// GetSchema 返回已配置的架构。
# <翻译结束>







<原文开始>
// (Optional, "master" in default) Node role, used for master-slave mode: master, slave.
<原文结束>

# <翻译开始>
// （可选，默认为 "master"）节点角色，用于主从模式：master（主节点），slave（从节点）。
# <翻译结束>


<原文开始>
// (Optional) Weight for load balance calculating, it's useless if there's just one node.
<原文结束>

# <翻译开始>
// (可选) 用于负载均衡计算的权重，如果只有一个节点则该参数无效。
# <翻译结束>


<原文开始>
// (Optional, "utf8" in default) Custom charset when operating on database.
<原文结束>

# <翻译开始>
// （可选，默认为"utf8"）在操作数据库时自定义的字符集。
# <翻译结束>


<原文开始>
// (Optional, "tcp" in default) See net.Dial for more information which networks are available.
<原文结束>

# <翻译开始>
// （可选，默认为 "tcp"）有关可用网络的更多信息，请参阅 net.Dial。
# <翻译结束>


<原文开始>
// (Optional) Namespace for some databases. Eg, in pgsql, the `Name` acts as the `catalog`, the `NameSpace` acts as the `schema`.
<原文结束>

# <翻译开始>
// （可选）为某些数据库提供命名空间。例如，在pgsql中，`Name`充当`catalog`的角色，而`NameSpace`充当`schema`的角色。
# <翻译结束>


<原文开始>
// Authentication username.
<原文结束>

# <翻译开始>
// 认证用户名。
# <翻译结束>


<原文开始>
// Authentication password.
<原文结束>

# <翻译开始>
// 认证密码
# <翻译结束>


<原文开始>
// (Optional) Table prefix.
<原文结束>

# <翻译开始>
// （可选）表前缀。
# <翻译结束>


<原文开始>
// Default group name.
<原文结束>

# <翻译开始>
// 默认分组名称。
# <翻译结束>


<原文开始>
// All configurations.
<原文结束>

# <翻译开始>
// 所有配置。
# <翻译结束>


<原文开始>
// Default value checks.
<原文结束>

# <翻译开始>
// 默认值检查。
# <翻译结束>

