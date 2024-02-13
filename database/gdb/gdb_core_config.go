// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

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
type X配置 map[string]X配置组

// ConfigGroup 是为指定名称组配置的配置节点切片。
type X配置组 []X配置项

// ConfigNode 是单个节点的配置。
type X配置项 struct {
	X地址                 string        `json:"host"`                 // 服务器主机，可以是IP地址或域名，例如：127.0.0.1、localhost
	X端口                 string        `json:"port"`                 // Port, it's commonly 3306.
	X账号                 string        `json:"user"`                 // 认证用户名。
	X密码                 string        `json:"pass"`                 // 认证密码
	X名称                 string        `json:"name"`                 // 默认使用的数据库名称。
	X类型                 string        `json:"type"`                 // 数据库类型：mysql，sqlite，mssql，pgsql，oracle。
	X自定义链接信息                 string        `json:"link"`                 // (可选) 将所有配置的自定义链接信息放在一个单独字符串中。
	X额外                string        `json:"extra"`                // (可选) 根据已注册的第三方数据库驱动进行额外配置
	X节点角色                 string        `json:"role"`                 // （可选，默认为 "master"）节点角色，用于主从模式：master（主节点），slave（从节点）。
	X调试模式                bool          `json:"debug"`                // (可选) Debug模式启用调试信息日志记录和输出。
	X表前缀               string        `json:"prefix"`               // （可选）表前缀。
	X空跑特性               bool          `json:"dryRun"`               // (可选) 干预运行，仅执行 SELECT 语句但不执行 INSERT/UPDATE/DELETE 语句。
	X负载均衡权重               int           `json:"weight"`               // (可选) 用于负载均衡计算的权重，如果只有一个节点则该参数无效。
	X字符集              string        `json:"charset"`              // （可选，默认为"utf8"）在操作数据库时自定义的字符集。
	X协议             string        `json:"protocol"`             // （可选，默认为 "tcp"）有关可用网络的更多信息，请参阅 net.Dial。
	X时区             string        `json:"timezone"`             // (可选) 设置显示和解释时间戳时区。
	X命名空间            string        `json:"namespace"`            // （可选）为某些数据库提供命名空间。例如，在pgsql中，`Name`充当`catalog`的角色，而`NameSpace`充当`schema`的角色。
	X最大闲置连接数     int           `json:"maxIdle"`              // （可选）为底层连接池配置的最大空闲连接数。
	X最大打开连接数     int           `json:"maxOpen"`              // (可选) 用于底层连接池的最大打开连接配置。
	X最大空闲时长      time.Duration `json:"maxLifeTime"`          // （可选）在连接被关闭之前，允许其空闲的最大时长。
	X查询超时时长         time.Duration `json:"queryTimeout"`         // (可选) 每个DQL的最大查询时间。
	X执行超时时长          time.Duration `json:"execTimeout"`          // （可选）dml的最大执行时间。
	X事务超时时长          time.Duration `json:"tranTimeout"`          // (可选) 事务执行的最大时间。
	X预准备SQL超时时长       time.Duration `json:"prepareTimeout"`       // （可选）为准备操作设置最大执行时间。
	X创建时间字段名            string        `json:"createdAt"`            // (可选) 自动填充创建日期时间的表字段名称。
	X更新时间字段名            string        `json:"updatedAt"`            // (可选) 自动填充更新日期时间的表格字段名称。
	X软删除时间字段名            string        `json:"deletedAt"`            // (可选) 自动填充更新日期时间的表格字段名称。
	X禁用时间自动更新特性 bool          `json:"timeMaintainDisabled"` // (可选) 禁用自动时间维护功能。
}

const (
	DefaultGroupName = "default" // 默认分组名称。
)

// configs 指定内部使用的配置对象。
var configs struct {
	sync.RWMutex
	config X配置 // 所有配置。
	group  string // 默认配置组。
}

func init() {
	configs.config = make(X配置)
	configs.group = DefaultGroupName
}

// SetConfig 设置包的全局配置。
// 它将覆盖包的旧配置。
func X设置全局配置(配置 X配置) {
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

// SetConfigGroup 为给定的组设置配置。
func X设置组配置(配置组名称 string, 配置 X配置组) {
	defer instances.X清空()
	configs.Lock()
	defer configs.Unlock()
	for i, node := range 配置 {
		配置[i] = parseConfigNode(node)
	}
	configs.config[配置组名称] = 配置
}

// AddConfigNode 向给定组的配置中添加一个节点配置。
func X添加配置组节点(配置组名称 string, 配置 X配置项) {
	defer instances.X清空()
	configs.Lock()
	defer configs.Unlock()
	configs.config[配置组名称] = append(configs.config[配置组名称], parseConfigNode(配置))
}

// parseConfigNode 解析 `Link` 配置语法。
func parseConfigNode(node X配置项) X配置项 {
	if node.X自定义链接信息 != "" {
		node = *parseConfigNodeLink(&node)
	}
	if node.X自定义链接信息 != "" && node.X类型 == "" {
		match, _ := 正则类.X匹配文本(`([a-z]+):(.+)`, node.X自定义链接信息)
		if len(match) == 3 {
			node.X类型 = 文本类.X过滤首尾符并含空白(match[1])
			node.X自定义链接信息 = 文本类.X过滤首尾符并含空白(match[2])
		}
	}
	return node
}

// AddDefaultConfigNode 将一个节点配置添加到默认组的配置中。
func X添加默认配置组节点(配置 X配置项) {
	X添加配置组节点(DefaultGroupName, 配置)
}

// AddDefaultConfigGroup 向默认组的配置中添加多个节点配置。
func X添加默认配置组(配置组 X配置组) {
	X设置组配置(DefaultGroupName, 配置组)
}

// GetConfig 获取并返回给定组的配置。
func X取配置组配置(配置组名称 string) X配置组 {
	configs.RLock()
	defer configs.RUnlock()
	return configs.config[配置组名称]
}

// SetDefaultGroup 设置默认配置的组名称。
func X设置默认组名称(配置组名称 string) {
	defer instances.X清空()
	configs.Lock()
	defer configs.Unlock()
	configs.group = 配置组名称
}

// GetDefaultGroup 返回默认配置的名称。
func X获取默认组名称() string {
	defer instances.X清空()
	configs.RLock()
	defer configs.RUnlock()
	return configs.group
}

// IsConfigured 检查并返回数据库是否已配置。
// 如果存在任何配置信息，则返回 true。
func X是否已配置数据库() bool {
	configs.RLock()
	defer configs.RUnlock()
	return len(configs.config) > 0
}

// SetLogger 设置orm的记录器。
func (c *Core) X设置日志记录器(日志记录器 日志类.ILogger) {
	c.logger = 日志记录器
}

// GetLogger 返回 orm 的（日志器）
func (c *Core) X取日志记录器() 日志类.ILogger {
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
func (c *Core) X设置最大闲置连接数(连接数 int) {
	c.dynamicConfig.MaxIdleConnCount = 连接数
}

// SetMaxOpenConnCount 设置与数据库的最大连接数。
//
// 如果 MaxIdleConns 大于0且新的 MaxOpenConns 值小于 MaxIdleConns，则 MaxIdleConns 会相应减少以匹配新的 MaxOpenConns 限制。
//
// 若 n <= 0，则表示对打开的连接数没有限制。
// 默认值为 0（无限制）。
func (c *Core) X设置最大打开连接数(连接数 int) {
	c.dynamicConfig.MaxOpenConnCount = 连接数
}

// SetMaxConnLifeTime 设置单个连接可重用的最大时长。
//
// 到期的连接在重用前可能被延迟关闭。
//
// 如果 d <= 0，则不会因为连接存在时间过长而关闭连接。
func (c *Core) X设置最大空闲时长(时长 time.Duration) {
	c.dynamicConfig.MaxConnLifeTime = 时长
}

// GetConfig 返回当前正在使用的节点配置。
func (c *Core) X取当前节点配置() *X配置项 {
	internalData := c.底层_GetInternalCtxDataFromCtx(c.db.X取上下文对象())
	if internalData != nil && internalData.ConfigNode != nil {
// 注意：
// 该处会检查并返回当前数据库的配置信息，
// 如果当前数据库与从context获取到的config.Name中的模式（schema）不同时，
// 比如在嵌套事务场景中，context会在整个逻辑过程中被传递，
// 但context中的config.Name可能仍然是最初第一个事务对象中的原始值。
		if c.config.X名称 == internalData.ConfigNode.X名称 {
			return internalData.ConfigNode
		}
	}
	return c.config
}

// SetDebug用于开启或关闭调试模式。
func (c *Core) X设置调试模式(开启 bool) {
	c.debug.X设置值(开启)
}

// GetDebug 返回调试值。
func (c *Core) X取调试模式() bool {
	return c.debug.X取值()
}

// GetCache 返回内部缓存对象。
func (c *Core) X取缓存对象() *缓存类.Cache {
	return c.cache
}

// GetGroup 返回已配置的组字符串。
func (c *Core) X取配置组名称() string {
	return c.group
}

// SetDryRun 用于启用/禁用 DryRun 功能。
func (c *Core) X设置空跑特性(开启 bool) {
	c.config.X空跑特性 = 开启
}

// GetDryRun 返回 DryRun 的值。
func (c *Core) X取空跑特性() bool {
	return c.config.X空跑特性 || allDryRun
}

// GetPrefix 返回已配置的表前缀字符串。
func (c *Core) X取表前缀() string {
	return c.config.X表前缀
}

// GetSchema 返回已配置的架构。
func (c *Core) X取默认数据库名称() string {
	schema := c.schema
	if schema == "" {
		schema = c.db.X取当前节点配置().X名称
	}
	return schema
}

func parseConfigNodeLink(node *X配置项) *X配置项 {
	var match []string
	if node.X自定义链接信息 != "" {
		match, _ = 正则类.X匹配文本(linkPattern, node.X自定义链接信息)
		if len(match) > 5 {
			node.X类型 = match[1]
			node.X账号 = match[2]
			node.X密码 = match[3]
			node.X协议 = match[4]
			array := 文本类.X分割(match[5], ":")
			if len(array) == 2 && node.X协议 != "file" {
				node.X地址 = array[0]
				node.X端口 = array[1]
				node.X名称 = match[6]
			} else {
				node.X名称 = match[5]
			}
			if len(match) > 6 && match[7] != "" {
				node.X额外 = match[7]
			}
			node.X自定义链接信息 = ""
		}
	}
	if node.X额外 != "" {
		if m, _ := 文本类.X参数解析(node.X额外); len(m) > 0 {
			_ = 转换类.Struct(m, &node)
		}
	}
	// 默认值检查。
	if node.X字符集 == "" {
		node.X字符集 = defaultCharset
	}
	if node.X协议 == "" {
		node.X协议 = defaultProtocol
	}
	return node
}
