// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb
import (
	"sync"
	"time"
	
	"github.com/888go/goframe/os/gcache"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	)
// Config 是配置管理对象。
type Config map[string]ConfigGroup

// ConfigGroup 是为指定名称组配置的配置节点切片。
type ConfigGroup []ConfigNode

// ConfigNode 是单个节点的配置。
type ConfigNode struct {
	Host                 string        `json:"host"`                 // 服务器主机，可以是IP地址或域名，例如：127.0.0.1、localhost
	Port                 string        `json:"port"`                 // Port, it's commonly 3306.
	User                 string        `json:"user"`                 // 认证用户名。
	Pass                 string        `json:"pass"`                 // 认证密码
	Name                 string        `json:"name"`                 // 默认使用的数据库名称。
	Type                 string        `json:"type"`                 // 数据库类型：mysql，sqlite，mssql，pgsql，oracle。
	Link                 string        `json:"link"`                 // (可选) 将所有配置的自定义链接信息放在一个单独字符串中。
	Extra                string        `json:"extra"`                // (可选) 根据已注册的第三方数据库驱动进行额外配置
	Role                 string        `json:"role"`                 // （可选，默认为 "master"）节点角色，用于主从模式：master（主节点），slave（从节点）。
	Debug                bool          `json:"debug"`                // (可选) Debug模式启用调试信息日志记录和输出。
	Prefix               string        `json:"prefix"`               // （可选）表前缀。
	DryRun               bool          `json:"dryRun"`               // (可选) 干预运行，仅执行 SELECT 语句但不执行 INSERT/UPDATE/DELETE 语句。
	Weight               int           `json:"weight"`               // (可选) 用于负载均衡计算的权重，如果只有一个节点则该参数无效。
	Charset              string        `json:"charset"`              // （可选，默认为"utf8"）在操作数据库时自定义的字符集。
	Protocol             string        `json:"protocol"`             // （可选，默认为 "tcp"）有关可用网络的更多信息，请参阅 net.Dial。
	Timezone             string        `json:"timezone"`             // (可选) 设置显示和解释时间戳时区。
	Namespace            string        `json:"namespace"`            // （可选）为某些数据库提供命名空间。例如，在pgsql中，`Name`充当`catalog`的角色，而`NameSpace`充当`schema`的角色。
	MaxIdleConnCount     int           `json:"maxIdle"`              // （可选）为底层连接池配置的最大空闲连接数。
	MaxOpenConnCount     int           `json:"maxOpen"`              // (可选) 用于底层连接池的最大打开连接配置。
	MaxConnLifeTime      time.Duration `json:"maxLifeTime"`          // （可选）在连接被关闭之前，允许其空闲的最大时长。
	QueryTimeout         time.Duration `json:"queryTimeout"`         // (可选) 每个DQL的最大查询时间。
	ExecTimeout          time.Duration `json:"execTimeout"`          // （可选）dml的最大执行时间。
	TranTimeout          time.Duration `json:"tranTimeout"`          // (可选) 事务执行的最大时间。
	PrepareTimeout       time.Duration `json:"prepareTimeout"`       // （可选）为准备操作设置最大执行时间。
	CreatedAt            string        `json:"createdAt"`            // (可选) 自动填充创建日期时间的表字段名称。
	UpdatedAt            string        `json:"updatedAt"`            // (可选) 自动填充更新日期时间的表格字段名称。
	DeletedAt            string        `json:"deletedAt"`            // (可选) 自动填充更新日期时间的表格字段名称。
	TimeMaintainDisabled bool          `json:"timeMaintainDisabled"` // (可选) 禁用自动时间维护功能。
}

const (
	DefaultGroupName = "default" // 默认分组名称。
)

// configs 指定内部使用的配置对象。
var configs struct {
	sync.RWMutex
	config Config // 所有配置。
	group  string // 默认配置组。
}

func init() {
	configs.config = make(Config)
	configs.group = DefaultGroupName
}

// SetConfig 设置包的全局配置。
// 它将覆盖包的旧配置。
func SetConfig(config Config) {
	defer instances.Clear()
	configs.Lock()
	defer configs.Unlock()
	for k, nodes := range config {
		for i, node := range nodes {
			nodes[i] = parseConfigNode(node)
		}
		config[k] = nodes
	}
	configs.config = config
}

// SetConfigGroup 为给定的组设置配置。
func SetConfigGroup(group string, nodes ConfigGroup) {
	defer instances.Clear()
	configs.Lock()
	defer configs.Unlock()
	for i, node := range nodes {
		nodes[i] = parseConfigNode(node)
	}
	configs.config[group] = nodes
}

// AddConfigNode 向给定组的配置中添加一个节点配置。
func AddConfigNode(group string, node ConfigNode) {
	defer instances.Clear()
	configs.Lock()
	defer configs.Unlock()
	configs.config[group] = append(configs.config[group], parseConfigNode(node))
}

// parseConfigNode 解析 `Link` 配置语法。
func parseConfigNode(node ConfigNode) ConfigNode {
	if node.Link != "" {
		node = *parseConfigNodeLink(&node)
	}
	if node.Link != "" && node.Type == "" {
		match, _ := gregex.MatchString(`([a-z]+):(.+)`, node.Link)
		if len(match) == 3 {
			node.Type = gstr.Trim(match[1])
			node.Link = gstr.Trim(match[2])
		}
	}
	return node
}

// AddDefaultConfigNode 将一个节点配置添加到默认组的配置中。
func AddDefaultConfigNode(node ConfigNode) {
	AddConfigNode(DefaultGroupName, node)
}

// AddDefaultConfigGroup 向默认组的配置中添加多个节点配置。
func AddDefaultConfigGroup(nodes ConfigGroup) {
	SetConfigGroup(DefaultGroupName, nodes)
}

// GetConfig 获取并返回给定组的配置。
func GetConfig(group string) ConfigGroup {
	configs.RLock()
	defer configs.RUnlock()
	return configs.config[group]
}

// SetDefaultGroup 设置默认配置的组名称。
func SetDefaultGroup(name string) {
	defer instances.Clear()
	configs.Lock()
	defer configs.Unlock()
	configs.group = name
}

// GetDefaultGroup 返回默认配置的名称。
func GetDefaultGroup() string {
	defer instances.Clear()
	configs.RLock()
	defer configs.RUnlock()
	return configs.group
}

// IsConfigured 检查并返回数据库是否已配置。
// 如果存在任何配置信息，则返回 true。
func IsConfigured() bool {
	configs.RLock()
	defer configs.RUnlock()
	return len(configs.config) > 0
}

// SetLogger 设置orm的记录器。
func (c *Core) SetLogger(logger glog.ILogger) {
	c.logger = logger
}

// GetLogger 返回 orm 的（日志器）
func (c *Core) GetLogger() glog.ILogger {
	return c.logger
}

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
func (c *Core) SetMaxIdleConnCount(n int) {
	c.dynamicConfig.MaxIdleConnCount = n
}

// SetMaxOpenConnCount 设置与数据库的最大连接数。
//
// 如果 MaxIdleConns 大于0且新的 MaxOpenConns 值小于 MaxIdleConns，则 MaxIdleConns 会相应减少以匹配新的 MaxOpenConns 限制。
//
// 若 n <= 0，则表示对打开的连接数没有限制。
// 默认值为 0（无限制）。
func (c *Core) SetMaxOpenConnCount(n int) {
	c.dynamicConfig.MaxOpenConnCount = n
}

// SetMaxConnLifeTime 设置单个连接可重用的最大时长。
//
// 到期的连接在重用前可能被延迟关闭。
//
// 如果 d <= 0，则不会因为连接存在时间过长而关闭连接。
func (c *Core) SetMaxConnLifeTime(d time.Duration) {
	c.dynamicConfig.MaxConnLifeTime = d
}

// GetConfig 返回当前正在使用的节点配置。
func (c *Core) GetConfig() *ConfigNode {
	internalData := c.GetInternalCtxDataFromCtx(c.db.GetCtx())
	if internalData != nil && internalData.ConfigNode != nil {
// 注意：
// 该处会检查并返回当前数据库的配置信息，
// 如果当前数据库与从context获取到的config.Name中的模式（schema）不同时，
// 比如在嵌套事务场景中，context会在整个逻辑过程中被传递，
// 但context中的config.Name可能仍然是最初第一个事务对象中的原始值。
		if c.config.Name == internalData.ConfigNode.Name {
			return internalData.ConfigNode
		}
	}
	return c.config
}

// SetDebug用于开启或关闭调试模式。
func (c *Core) SetDebug(debug bool) {
	c.debug.Set(debug)
}

// GetDebug 返回调试值。
func (c *Core) GetDebug() bool {
	return c.debug.Val()
}

// GetCache 返回内部缓存对象。
func (c *Core) GetCache() *gcache.Cache {
	return c.cache
}

// GetGroup 返回已配置的组字符串。
func (c *Core) GetGroup() string {
	return c.group
}

// SetDryRun 用于启用/禁用 DryRun 功能。
func (c *Core) SetDryRun(enabled bool) {
	c.config.DryRun = enabled
}

// GetDryRun 返回 DryRun 的值。
func (c *Core) GetDryRun() bool {
	return c.config.DryRun || allDryRun
}

// GetPrefix 返回已配置的表前缀字符串。
func (c *Core) GetPrefix() string {
	return c.config.Prefix
}

// GetSchema 返回已配置的架构。
func (c *Core) GetSchema() string {
	schema := c.schema
	if schema == "" {
		schema = c.db.GetConfig().Name
	}
	return schema
}

func parseConfigNodeLink(node *ConfigNode) *ConfigNode {
	var match []string
	if node.Link != "" {
		match, _ = gregex.MatchString(linkPattern, node.Link)
		if len(match) > 5 {
			node.Type = match[1]
			node.User = match[2]
			node.Pass = match[3]
			node.Protocol = match[4]
			array := gstr.Split(match[5], ":")
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
		if m, _ := gstr.Parse(node.Extra); len(m) > 0 {
			_ = gconv.Struct(m, &node)
		}
	}
	// 默认值检查。
	if node.Charset == "" {
		node.Charset = defaultCharset
	}
	if node.Protocol == "" {
		node.Protocol = defaultProtocol
	}
	return node
}
