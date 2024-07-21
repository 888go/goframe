
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
// Config is the configuration management object.
<原文结束>

# <翻译开始>
// Config 是配置管理对象。 md5:1ef57338c678e400
# <翻译结束>


<原文开始>
// ConfigGroup is a slice of configuration node for specified named group.
<原文结束>

# <翻译开始>
// ConfigGroup 是为指定命名组的配置节点切片。 md5:fd0679403bacd284
# <翻译结束>


<原文开始>
// ConfigNode is configuration for one node.
<原文结束>

# <翻译开始>
// ConfigNode 是一个节点的配置信息。
// 备注: 此配置结构不做名称翻译, 防止通过map载入配置时, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// md5:c441354f84b63933
# <翻译结束>


<原文开始>
// Host of server, ip or domain like: 127.0.0.1, localhost
<原文结束>

# <翻译开始>
// 服务器的主机，可以是 IP 地址或域名，如：127.0.0.1，localhost. md5:995f8d0f775d1561
# <翻译结束>


<原文开始>
// Authentication username.
<原文结束>

# <翻译开始>
// 身份验证用户名。 md5:bc205abcf7fb6329
# <翻译结束>


<原文开始>
// Authentication password.
<原文结束>

# <翻译开始>
// 认证密码。 md5:325b61eb5935f198
# <翻译结束>


<原文开始>
// Default used database name.
<原文结束>

# <翻译开始>
// 默认使用的数据库名称。 md5:13fed4b7ca845a3b
# <翻译结束>


<原文开始>
// Database type: mysql, mariadb, sqlite, mssql, pgsql, oracle, clickhouse, dm.
<原文结束>

# <翻译开始>
// 数据库类型：mysql，mariadb，sqlite，mssql，pgsql，oracle，clickhouse，dm。 md5:6b9f1786cf5a0d7a
# <翻译结束>


<原文开始>
// (Optional) Custom link information for all configuration in one single string.
<原文结束>

# <翻译开始>
// （可选）所有配置的自定义链接信息，作为一个单个字符串。 md5:cb9784decaebf7c3
# <翻译结束>


<原文开始>
// (Optional) Extra configuration according the registered third-party database driver.
<原文结束>

# <翻译开始>
// （可选）根据注册的第三方数据库驱动程序提供的额外配置。 md5:6ba9ea91183a2b0c
# <翻译结束>


<原文开始>
// (Optional, "master" in default) Node role, used for master-slave mode: master, slave.
<原文结束>

# <翻译开始>
// （可选，默认为"master"）节点角色，用于主从模式：master, slave。 md5:9645d0e7417ebf0c
# <翻译结束>


<原文开始>
// (Optional) Debug mode enables debug information logging and output.
<原文结束>

# <翻译开始>
// （可选）调试模式启用调试信息日志和输出。 md5:e292d7585b9505f9
# <翻译结束>


<原文开始>
// (Optional) Table prefix.
<原文结束>

# <翻译开始>
//（可选）表前缀。 md5:201acb7d8a3cfba7
# <翻译结束>


<原文开始>
// (Optional) Dry run, which does SELECT but no INSERT/UPDATE/DELETE statements.
<原文结束>

# <翻译开始>
// （可选）空跑特性，只执行 SELECT 语句，而不执行 INSERT/UPDATE/DELETE 语句。 md5:3983d4a8bb269d45
# <翻译结束>


<原文开始>
// (Optional) Weight for load balance calculating, it's useless if there's just one node.
<原文结束>

# <翻译开始>
// （可选）用于负载均衡计算的权重，如果只有一个节点则无效。 md5:6be8657f1809396b
# <翻译结束>


<原文开始>
// (Optional, "utf8" in default) Custom charset when operating on database.
<原文结束>

# <翻译开始>
// (可选，默认为 "utf8") 操作数据库时使用的自定义字符集。 md5:e63288ee7f2834e2
# <翻译结束>


<原文开始>
// (Optional, "tcp" in default) See net.Dial for more information which networks are available.
<原文结束>

# <翻译开始>
// （可选， 默认为 "tcp"）有关可用网络的更多信息，请参阅 net.Dial。 md5:96a17fcac4ef394d
# <翻译结束>


<原文开始>
// (Optional) Sets the time zone for displaying and interpreting time stamps.
<原文结束>

# <翻译开始>
// （可选）设置时区，用于显示和解释时间戳。 md5:f1c54d7158bc5e89
# <翻译结束>


<原文开始>
// (Optional) Namespace for some databases. Eg, in pgsql, the `Name` acts as the `catalog`, the `NameSpace` acts as the `schema`.
<原文结束>

# <翻译开始>
// （可选）某些数据库的命名空间。例如，在pgsql中，`Name` 用作 `catalog`，而 `NameSpace` 用作 `schema`。 md5:8dc2fa65d951a94e
# <翻译结束>


<原文开始>
// (Optional) Max idle connection configuration for underlying connection pool.
<原文结束>

# <翻译开始>
//（可选）底层连接池的最大空闲连接配置。 md5:b1f10cf73af643a1
# <翻译结束>


<原文开始>
// (Optional) Max open connection configuration for underlying connection pool.
<原文结束>

# <翻译开始>
//（可选）底层连接池的最大打开连接配置。 md5:bd8ec40915479f50
# <翻译结束>


<原文开始>
// (Optional) Max amount of time a connection may be idle before being closed.
<原文结束>

# <翻译开始>
// （可选）连接在被关闭之前允许的最大空闲时间。 md5:b078efc5d928800a
# <翻译结束>


<原文开始>
// (Optional) Max query time for per dql.
<原文结束>

# <翻译开始>
// （可选）每个DQL查询的最大时间。 md5:4c57cb456fb2ee8b
# <翻译结束>


<原文开始>
// (Optional) Max exec time for dml.
<原文结束>

# <翻译开始>
// (可选) DML的最大执行时间。 md5:481b9bbde2c49552
# <翻译结束>


<原文开始>
// (Optional) Max exec time for a transaction.
<原文结束>

# <翻译开始>
//（可选）事务的最大执行时间。 md5:8e1cd855a33e2c6b
# <翻译结束>


<原文开始>
// (Optional) Max exec time for prepare operation.
<原文结束>

# <翻译开始>
// （可选）准备操作的最大执行时间。 md5:14739995bdfab318
# <翻译结束>


<原文开始>
// (Optional) The field name of table for automatic-filled created datetime.
<原文结束>

# <翻译开始>
// （可选）自动填充创建时间的字段名。 md5:85fd52001481b733
# <翻译结束>


<原文开始>
// (Optional) The field name of table for automatic-filled updated datetime.
<原文结束>

# <翻译开始>
//（可选）自动填充更新日期时间的表字段名称。 md5:5d4aa50ffafb9f8e
# <翻译结束>


<原文开始>
// (Optional) Disable the automatic time maintaining feature.
<原文结束>

# <翻译开始>
// （可选）禁用自动时间保持功能。 md5:082a8741b2590acc
# <翻译结束>


<原文开始>
// configs specifies internal used configuration object.
<原文结束>

# <翻译开始>
// configs 指定用于内部使用的配置对象。 md5:32dbb902d70edfe2
# <翻译结束>


<原文开始>
// Default configuration group.
<原文结束>

# <翻译开始>
// 默认配置组。 md5:1592850319e0d8a7
# <翻译结束>


<原文开始>
// SetConfig sets the global configuration for package.
// It will overwrite the old configuration of package.
<原文结束>

# <翻译开始>
// SetConfig 设置包的全局配置。
// 它将覆盖包原有的配置。
// md5:e9d794925a260f3e
# <翻译结束>


<原文开始>
// SetConfigGroup sets the configuration for given group.
<原文结束>

# <翻译开始>
// SetConfigGroup 设置给定组的配置。 md5:3ca6c2845a6dcd23
# <翻译结束>


<原文开始>
// AddConfigNode adds one node configuration to configuration of given group.
<原文结束>

# <翻译开始>
// AddConfigNode 向给定组的配置中添加一个节点配置。 md5:3f2a775dc2be575b
# <翻译结束>


<原文开始>
// parseConfigNode parses `Link` configuration syntax.
<原文结束>

# <翻译开始>
// parseConfigNode 解析 `Link` 配置语法。 md5:ca390415077cad45
# <翻译结束>


<原文开始>
// AddDefaultConfigNode adds one node configuration to configuration of default group.
<原文结束>

# <翻译开始>
// AddDefaultConfigNode 在默认组的配置中添加一个节点配置。 md5:0b566ab59e6984e4
# <翻译结束>


<原文开始>
// AddDefaultConfigGroup adds multiple node configurations to configuration of default group.
<原文结束>

# <翻译开始>
// AddDefaultConfigGroup 将多个节点配置添加到默认组的配置中。 md5:b289b432b1f9a27f
# <翻译结束>


<原文开始>
// GetConfig retrieves and returns the configuration of given group.
<原文结束>

# <翻译开始>
// GetConfig 获取并返回给定组的配置。 md5:e4487cb50b45e5f4
# <翻译结束>


<原文开始>
// SetDefaultGroup sets the group name for default configuration.
<原文结束>

# <翻译开始>
// SetDefaultGroup 设置默认配置的组名。 md5:e7734b91e5838c18
# <翻译结束>


<原文开始>
// GetDefaultGroup returns the { name of default configuration.
<原文结束>

# <翻译开始>
// GetDefaultGroup 返回默认配置的{名称}。 md5:59cc62505c297d57
# <翻译结束>


<原文开始>
// IsConfigured checks and returns whether the database configured.
// It returns true if any configuration exists.
<原文结束>

# <翻译开始>
// IsConfigured 检查并返回数据库是否已配置。
// 如果存在任何配置，它将返回 true。
// md5:1232e7ebd0a7ce10
# <翻译结束>


<原文开始>
// SetLogger sets the logger for orm.
<原文结束>

# <翻译开始>
// SetLogger为ORM设置日志记录器。 md5:a70ca86920e39e54
# <翻译结束>


<原文开始>
// GetLogger returns the (logger) of the orm.
<原文结束>

# <翻译开始>
// GetLogger 返回 ORM 的日志记录器。 md5:8fc6f96186fd98c6
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
// SetMaxIdleConnCount 设置空闲连接池中的最大连接数。
//
// 如果 MaxOpenConns 大于 0 但小于新的 MaxIdleConns，那么新的 MaxIdleConns 将被调整为与 MaxOpenConns 的限制相匹配。
//
// 如果 n 小于或等于 0，则不保留任何空闲连接。
//
// 当前默认的最大空闲连接数为 2。这可能会在未来的版本中改变。
// md5:7d6f4079c0bfc25f
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
// SetMaxOpenConnCount 设置到数据库的最大打开连接数。
//
// 如果 MaxIdleConns 大于 0，并且新的 MaxOpenConns 小于 MaxIdleConns，那么 MaxIdleConns 将被调整为与新的 MaxOpenConns 限制相匹配。
//
// 如果 n 小于或等于 0，则没有对打开连接数的限制。默认值为 0（无限制）。
// md5:e8cfc3ecf7f5887e
# <翻译结束>


<原文开始>
// SetMaxConnLifeTime sets the maximum amount of time a connection may be reused.
//
// Expired connections may be closed lazily before reuse.
//
// If d <= 0, connections are not closed due to a connection's age.
<原文结束>

# <翻译开始>
// SetMaxConnLifeTime 设置连接可被重用的最大时间。
//
// 过期的连接可能会在被重用前被惰性关闭。
//
// 如果 d <= 0，则连接不会因为超时而被关闭。
// md5:f8d0da250f6387ba
# <翻译结束>


<原文开始>
// GetConfig returns the current used node configuration.
<原文结束>

# <翻译开始>
// GetConfig 返回当前使用的节点配置。 md5:c953d82ac4cddf35
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
		// 它会检查并从当前数据库返回配置，
		// 如果当前数据库和上下文中的config.Name（名称）之间存在不同的模式，
		// 例如，在嵌套事务场景中，上下文会传递给整个逻辑处理过程，
		// 但来自上下文的config.Name可能仍然是最初事务对象的原始值。
		// md5:b5980190888563ed
# <翻译结束>


<原文开始>
// SetDebug enables/disables the debug mode.
<原文结束>

# <翻译开始>
// SetDebug 启用/禁用调试模式。 md5:44a23ae9ad388bd8
# <翻译结束>


<原文开始>
// GetDebug returns the debug value.
<原文结束>

# <翻译开始>
// GetDebug 返回调试值。 md5:9bffedbe7bd8f1cf
# <翻译结束>


<原文开始>
// GetCache returns the internal cache object.
<原文结束>

# <翻译开始>
// GetCache 返回内部的缓存对象。 md5:3d83a15ed3d14950
# <翻译结束>


<原文开始>
// GetGroup returns the group string configured.
<原文结束>

# <翻译开始>
// GetGroup 返回配置的组字符串。 md5:fb5b1b36ae36c283
# <翻译结束>


<原文开始>
// SetDryRun enables/disables the DryRun feature.
<原文结束>

# <翻译开始>
// SetDryRun 启用/禁用DryRun功能。 md5:359f8392ba799c27
# <翻译结束>


<原文开始>
// GetDryRun returns the DryRun value.
<原文结束>

# <翻译开始>
// GetDryRun 返回DryRun的值。 md5:7e133dad1f0ee7ba
# <翻译结束>


<原文开始>
// GetPrefix returns the table prefix string configured.
<原文结束>

# <翻译开始>
// GetPrefix 返回配置的表前缀字符串。 md5:637396955caa18c4
# <翻译结束>


<原文开始>
// GetSchema returns the schema configured.
<原文结束>

# <翻译开始>
// GetSchema 返回已配置的模式。 md5:89a8c016a19c9022
# <翻译结束>

