// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"sync"
	"time"

	gcache "github.com/888go/goframe/os/gcache"
	glog "github.com/888go/goframe/os/glog"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// Config 是配置管理对象。 md5:1ef57338c678e400
type Config map[string]ConfigGroup

// ConfigGroup 是为指定命名组的配置节点切片。 md5:fd0679403bacd284
type ConfigGroup []ConfigNode

// ConfigNode 是一个节点的配置信息。
// 备注: 此配置结构不做名称翻译, 防止通过map载入配置时, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// ConfigNode 是一个节点的配置信息。
// 备注: 此配置结构不做名称翻译, 防止通过map载入配置时, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// ConfigNode 是一个节点的配置信息。
// 备注: 此配置结构不做名称翻译, 防止通过map载入配置时, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// ConfigNode 是一个节点的配置信息。
// 备注: 此配置结构不做名称翻译, 防止通过map载入配置时, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// ConfigNode 是一个节点的配置信息。
// 备注: 此配置结构不做名称翻译, 防止通过map载入配置时, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// ConfigNode 是一个节点的配置信息。
// 备注: 此配置结构不做名称翻译, 防止通过map载入配置时, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// ConfigNode 是一个节点的配置信息。
// 备注: 此配置结构不做名称翻译, 防止通过map载入配置时, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// ConfigNode 是一个节点的配置信息。
// 备注: 此配置结构不做名称翻译, 防止通过map载入配置时, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// md5:c441354f84b63933
type ConfigNode struct {
	Host                 string        `json:"host"`                 // 服务器的主机，可以是 IP 地址或域名，如：127.0.0.1，localhost. md5:995f8d0f775d1561
	Port                 string        `json:"port"`                 // Port, it's commonly 3306.
	User                 string        `json:"user"`                 // 身份验证用户名。 md5:bc205abcf7fb6329
	Pass                 string        `json:"pass"`                 // 认证密码。 md5:325b61eb5935f198
	Name                 string        `json:"name"`                 // 默认使用的数据库名称。 md5:13fed4b7ca845a3b
	Type                 string        `json:"type"`                 // 数据库类型：mysql，mariadb，sqlite，mssql，pgsql，oracle，clickhouse，dm。 md5:6b9f1786cf5a0d7a
	Link                 string        `json:"link"`                 // （可选）所有配置的自定义链接信息，作为一个单个字符串。 md5:cb9784decaebf7c3
	Extra                string        `json:"extra"`                // （可选）根据注册的第三方数据库驱动程序提供的额外配置。 md5:6ba9ea91183a2b0c
	Role                 string        `json:"role"`                 // （可选，默认为"master"）节点角色，用于主从模式：master, slave。 md5:9645d0e7417ebf0c
	Debug                bool          `json:"debug"`                // （可选）调试模式启用调试信息日志和输出。 md5:e292d7585b9505f9
	Prefix               string        `json:"prefix"`               //（可选）表前缀。 md5:201acb7d8a3cfba7
	DryRun               bool          `json:"dryRun"`               // （可选）空跑特性，只执行 SELECT 语句，而不执行 INSERT/UPDATE/DELETE 语句。 md5:3983d4a8bb269d45
	Weight               int           `json:"weight"`               // （可选）用于负载均衡计算的权重，如果只有一个节点则无效。 md5:6be8657f1809396b
	Charset              string        `json:"charset"`              // (可选，默认为 "utf8") 操作数据库时使用的自定义字符集。 md5:e63288ee7f2834e2
	Protocol             string        `json:"protocol"`             // （可选， 默认为 "tcp"）有关可用网络的更多信息，请参阅 net.Dial。 md5:96a17fcac4ef394d
	Timezone             string        `json:"timezone"`             // （可选）设置时区，用于显示和解释时间戳。 md5:f1c54d7158bc5e89
	Namespace            string        `json:"namespace"`            // （可选）某些数据库的命名空间。例如，在pgsql中，`Name` 用作 `catalog`，而 `NameSpace` 用作 `schema`。 md5:8dc2fa65d951a94e
	MaxIdleConnCount     int           `json:"maxIdle"`              //（可选）底层连接池的最大空闲连接配置。 md5:b1f10cf73af643a1
	MaxOpenConnCount     int           `json:"maxOpen"`              //（可选）底层连接池的最大打开连接配置。 md5:bd8ec40915479f50
	MaxConnLifeTime      time.Duration `json:"maxLifeTime"`          // （可选）连接在被关闭之前允许的最大空闲时间。 md5:b078efc5d928800a
	QueryTimeout         time.Duration `json:"queryTimeout"`         // （可选）每个DQL查询的最大时间。 md5:4c57cb456fb2ee8b
	ExecTimeout          time.Duration `json:"execTimeout"`          // (可选) DML的最大执行时间。 md5:481b9bbde2c49552
	TranTimeout          time.Duration `json:"tranTimeout"`          //（可选）事务的最大执行时间。 md5:8e1cd855a33e2c6b
	PrepareTimeout       time.Duration `json:"prepareTimeout"`       // （可选）准备操作的最大执行时间。 md5:14739995bdfab318
	CreatedAt            string        `json:"createdAt"`            // （可选）自动填充创建时间的字段名。 md5:85fd52001481b733
	UpdatedAt            string        `json:"updatedAt"`            //（可选）自动填充更新日期时间的表字段名称。 md5:5d4aa50ffafb9f8e
	DeletedAt            string        `json:"deletedAt"`            //（可选）自动填充更新日期时间的表字段名称。 md5:5d4aa50ffafb9f8e
	TimeMaintainDisabled bool          `json:"timeMaintainDisabled"` // （可选）禁用自动时间保持功能。 md5:082a8741b2590acc
}

const (
	DefaultGroupName = "default" // Default group name.
)

// configs 指定用于内部使用的配置对象。 md5:32dbb902d70edfe2
var configs struct {
	sync.RWMutex
	config Config // All configurations.
	group  string // 默认配置组。 md5:1592850319e0d8a7
}

func init() {
	configs.config = make(Config)
	configs.group = DefaultGroupName
}

// X设置全局配置 设置包的全局配置。
// 它将覆盖包原有的配置。
// md5:e9d794925a260f3e
func X设置全局配置(配置 Config) {
	defer instances.X清空()
	configs.Lock()
	defer configs.Unlock()
	for k, nodes := range 配置 {
		for i, node := range nodes {
			nodes[i] = parseConfigNode(node)
		}
		配置[k] = nodes
	}
	configs.config = 配置
}

// X设置组配置 设置给定组的配置。 md5:3ca6c2845a6dcd23
func X设置组配置(配置组名称 string, 配置 ConfigGroup) {
	defer instances.X清空()
	configs.Lock()
	defer configs.Unlock()
	for i, node := range 配置 {
		配置[i] = parseConfigNode(node)
	}
	configs.config[配置组名称] = 配置
}

// X添加配置组节点 向给定组的配置中添加一个节点配置。 md5:3f2a775dc2be575b
func X添加配置组节点(配置组名称 string, 配置 ConfigNode) {
	defer instances.X清空()
	configs.Lock()
	defer configs.Unlock()
	configs.config[配置组名称] = append(configs.config[配置组名称], parseConfigNode(配置))
}

// parseConfigNode 解析 `Link` 配置语法。 md5:ca390415077cad45
func parseConfigNode(node ConfigNode) ConfigNode {
	if node.Link != "" {
		node = *parseConfigNodeLink(&node)
	}
	if node.Link != "" && node.Type == "" {
		match, _ := gregex.X匹配文本(`([a-z]+):(.+)`, node.Link)
		if len(match) == 3 {
			node.Type = gstr.X过滤首尾符并含空白(match[1])
			node.Link = gstr.X过滤首尾符并含空白(match[2])
		}
	}
	return node
}

// X添加默认配置组节点 在默认组的配置中添加一个节点配置。 md5:0b566ab59e6984e4
func X添加默认配置组节点(配置 ConfigNode) {
	X添加配置组节点(DefaultGroupName, 配置)
}

// X添加默认配置组 将多个节点配置添加到默认组的配置中。 md5:b289b432b1f9a27f
func X添加默认配置组(配置组 ConfigGroup) {
	X设置组配置(DefaultGroupName, 配置组)
}

// X取配置组配置 获取并返回给定组的配置。 md5:e4487cb50b45e5f4
func X取配置组配置(配置组名称 string) ConfigGroup {
	configs.RLock()
	defer configs.RUnlock()
	return configs.config[配置组名称]
}

// X设置默认组名称 设置默认配置的组名。 md5:e7734b91e5838c18
func X设置默认组名称(配置组名称 string) {
	defer instances.X清空()
	configs.Lock()
	defer configs.Unlock()
	configs.group = 配置组名称
}

// X获取默认组名称 返回默认配置的{名称}。 md5:59cc62505c297d57
func X获取默认组名称() string {
	defer instances.X清空()
	configs.RLock()
	defer configs.RUnlock()
	return configs.group
}

// X是否已配置数据库 检查并返回数据库是否已配置。
// 如果存在任何配置，它将返回 true。
// md5:1232e7ebd0a7ce10
func X是否已配置数据库() bool {
	configs.RLock()
	defer configs.RUnlock()
	return len(configs.config) > 0
}

// X设置日志记录器为ORM设置日志记录器。 md5:a70ca86920e39e54
func (c *Core) X设置日志记录器(日志记录器 glog.ILogger) {
	c.logger = 日志记录器
}

// X取日志记录器 返回 ORM 的日志记录器。 md5:8fc6f96186fd98c6
func (c *Core) X取日志记录器() glog.ILogger {
	return c.logger
}

// X设置最大闲置连接数 设置空闲连接池中的最大连接数。
//
// 如果 MaxOpenConns 大于 0 但小于新的 MaxIdleConns，那么新的 MaxIdleConns 将被调整为与 MaxOpenConns 的限制相匹配。
//
// 如果 n 小于或等于 0，则不保留任何空闲连接。
//
// 当前默认的最大空闲连接数为 2。这可能会在未来的版本中改变。
// md5:7d6f4079c0bfc25f
func (c *Core) X设置最大闲置连接数(连接数 int) {
	c.dynamicConfig.MaxIdleConnCount = 连接数
}

// X设置最大打开连接数 设置到数据库的最大打开连接数。
//
// 如果 MaxIdleConns 大于 0，并且新的 MaxOpenConns 小于 MaxIdleConns，那么 MaxIdleConns 将被调整为与新的 MaxOpenConns 限制相匹配。
//
// 如果 n 小于或等于 0，则没有对打开连接数的限制。默认值为 0（无限制）。
// md5:e8cfc3ecf7f5887e
func (c *Core) X设置最大打开连接数(连接数 int) {
	c.dynamicConfig.MaxOpenConnCount = 连接数
}

// X设置最大空闲时长 设置连接可被重用的最大时间。
//
// 过期的连接可能会在被重用前被惰性关闭。
//
// 如果 d <= 0，则连接不会因为超时而被关闭。
// md5:f8d0da250f6387ba
func (c *Core) X设置最大空闲时长(时长 time.Duration) {
	c.dynamicConfig.MaxConnLifeTime = 时长
}

// X取当前节点配置 返回当前使用的节点配置。 md5:c953d82ac4cddf35
func (c *Core) X取当前节点配置() *ConfigNode {
	var configNode = c.getConfigNodeFromCtx(c.db.X取上下文对象())
	if configNode != nil {
		// 注意：
		// 它会检查并从当前数据库返回配置，
		// 如果当前数据库和上下文中的config.Name（名称）之间存在不同的模式，
		// 例如，在嵌套事务场景中，上下文会传递给整个逻辑处理过程，
		// 但来自上下文的config.Name可能仍然是最初事务对象的原始值。
		// md5:b5980190888563ed
		if c.config.Name == configNode.Name {
			return configNode
		}
	}
	return c.config
}

// X设置调试模式 启用/禁用调试模式。 md5:44a23ae9ad388bd8
func (c *Core) X设置调试模式(开启 bool) {
	c.debug.X设置值(开启)
}

// X取调试模式 返回调试值。 md5:9bffedbe7bd8f1cf
func (c *Core) X取调试模式() bool {
	return c.debug.X取值()
}

// X取缓存对象 返回内部的缓存对象。 md5:3d83a15ed3d14950
func (c *Core) X取缓存对象() *gcache.Cache {
	return c.cache
}

// X取配置组名称 返回配置的组字符串。 md5:fb5b1b36ae36c283
func (c *Core) X取配置组名称() string {
	return c.group
}

// X设置空跑特性 启用/禁用DryRun功能。 md5:359f8392ba799c27
func (c *Core) X设置空跑特性(开启 bool) {
	c.config.DryRun = 开启
}

// X取空跑特性 返回DryRun的值。 md5:7e133dad1f0ee7ba
func (c *Core) X取空跑特性() bool {
	return c.config.DryRun || allDryRun
}

// X取表前缀 返回配置的表前缀字符串。 md5:637396955caa18c4
func (c *Core) X取表前缀() string {
	return c.config.Prefix
}

// X取默认数据库名称 返回已配置的模式。 md5:89a8c016a19c9022
func (c *Core) X取默认数据库名称() string {
	schema := c.schema
	if schema == "" {
		schema = c.db.X取当前节点配置().Name
	}
	return schema
}

func parseConfigNodeLink(node *ConfigNode) *ConfigNode {
	var match []string
	if node.Link != "" {
		match, _ = gregex.X匹配文本(linkPattern, node.Link)
		if len(match) > 5 {
			node.Type = match[1]
			node.User = match[2]
			node.Pass = match[3]
			node.Protocol = match[4]
			array := gstr.X分割(match[5], ":")
			if len(array) == 2 && node.Protocol != "file" {
				node.Host = array[0]
				node.Port = array[1]
				node.Name = match[6]
			} else {
				node.Name = match[5]
			}
			if len(match) > 6 && match[7] != "" {
				node.Extra = match[7]
			}
			node.Link = ""
		}
	}
	if node.Extra != "" {
		if m, _ := gstr.X参数解析(node.Extra); len(m) > 0 {
			_ = gconv.Struct(m, &node)
		}
	}
	// Default value checks.
	if node.Charset == "" {
		node.Charset = defaultCharset
	}
	if node.Protocol == "" {
		node.Protocol = defaultProtocol
	}
	return node
}
