// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gredis

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/util/gconv"
)

// Config 是 Redis 配置。 md5:5409b3144db1576a
type Config struct {
	// Address 支持单个和集群 Redis 服务器。多个地址使用逗号分隔。例如：192.168.1.1:6379, 192.168.1.2:6379。 md5:21ac53e24210b32c
	Address         string        `json:"address"`
	Db              int           `json:"db"`              // Redis db.
	User            string        `json:"user"`            // Username for AUTH.
	Pass            string        `json:"pass"`            // Password for AUTH.
	SentinelUser    string        `json:"sentinel_user"`   // 防卫者AUTH的用户名。 md5:c85c5044b04f7ec3
	SentinelPass    string        `json:"sentinel_pass"`   // 密码，用于sentinel的AUTH。 md5:517cbd39fc9e4f20
	MinIdle         int           `json:"minIdle"`         // 允许处于空闲状态的连接的最小数量（默认为0）. md5:534a8e485a7c5664
	MaxIdle         int           `json:"maxIdle"`         // 允许的最大空闲连接数（默认为10）. md5:6b33b39ddbb7c42b
	MaxActive       int           `json:"maxActive"`       // 连接数的最大限制（默认为0表示无限制）。 md5:4dbd7ce4d80b4597
	MaxConnLifetime time.Duration `json:"maxConnLifetime"` // 连接的最大生命周期（默认为30秒，不允许设置为0）. md5:1650bf54f8065411
	IdleTimeout     time.Duration `json:"idleTimeout"`     // 连接的最大空闲时间（默认为10秒，不允许设置为0）. md5:1f9346d51eb9e76a
	WaitTimeout     time.Duration `json:"waitTimeout"`     // 超时等待从连接池获取连接的持续时间。 md5:ff75b0772da43843
	DialTimeout     time.Duration `json:"dialTimeout"`     // TCP连接的超时时间。 md5:d900304d9b7c2e58
	ReadTimeout     time.Duration `json:"readTimeout"`     // TCP读超时。如果没有必要，请不要设置它。 md5:44e33a5ef46ceb97
	WriteTimeout    time.Duration `json:"writeTimeout"`    // Write timeout for TCP.
	MasterName      string        `json:"masterName"`      // 用于Redis哨兵模式。 md5:44b4d0f3813a15e6
	TLS             bool          `json:"tls"`             // 指定连接到服务器时是否应使用TLS。 md5:0b36620d5b0321dd
	TLSSkipVerify   bool          `json:"tlsSkipVerify"`   // 连接TLS时禁用服务器名称验证。 md5:3bde87f1295352e9
	TLSConfig       *tls.Config   `json:"-"`               // 使用的TLS配置。如果设置，将协商TLS。 md5:b5a9a25bb2762b0b
	SlaveOnly       bool          `json:"slaveOnly"`       // 将所有命令路由到从属只读节点。 md5:9ba156f404a631f1
	Cluster         bool          `json:"cluster"`         // 指定是否使用集群模式。 md5:7952648d7b1da3f9
	Protocol        int           `json:"protocol"`        // 定义RESP版本（协议2或3）。 md5:dbc1edfd3b1e3b35
}

const (
	DefaultGroupName = "default" // 默认配置组名称。 md5:eb4945d78061d92a
)

var (
	// Configuration groups.
	localConfigMap = gmap.NewStrAnyMap(true)
)

// SetConfig 为指定的组设置全局配置。
// 如果没有传递 `name`，则为默认组名设置配置。
// md5:8d7c1f181c0057f0
func SetConfig(config *Config, name ...string) {
	group := DefaultGroupName
	if len(name) > 0 {
		group = name[0]
	}
	localConfigMap.Set(group, config)

	intlog.Printf(context.TODO(), `SetConfig for group "%s": %+v`, group, config)
}

// SetConfigByMap 使用映射设置指定组的全局配置。
// 如果未传递`name`，则将配置设置为默认组名。
// md5:1d191bb426ab05fb
func SetConfigByMap(m map[string]interface{}, name ...string) error {
	group := DefaultGroupName
	if len(name) > 0 {
		group = name[0]
	}
	config, err := ConfigFromMap(m)
	if err != nil {
		return err
	}
	localConfigMap.Set(group, config)
	return nil
}

// ConfigFromMap 从给定的映射中解析并返回配置。 md5:105a2224aed53bc9
func ConfigFromMap(m map[string]interface{}) (config *Config, err error) {
	config = &Config{}
	if err = gconv.Scan(m, config); err != nil {
		err = gerror.NewCodef(gcode.CodeInvalidConfiguration, `invalid redis configuration: %#v`, m)
	}
	if config.DialTimeout < time.Second {
		config.DialTimeout = config.DialTimeout * time.Second
	}
	if config.WaitTimeout < time.Second {
		config.WaitTimeout = config.WaitTimeout * time.Second
	}
	if config.WriteTimeout < time.Second {
		config.WriteTimeout = config.WriteTimeout * time.Second
	}
	if config.ReadTimeout < time.Second {
		config.ReadTimeout = config.ReadTimeout * time.Second
	}
	if config.IdleTimeout < time.Second {
		config.IdleTimeout = config.IdleTimeout * time.Second
	}
	if config.MaxConnLifetime < time.Second {
		config.MaxConnLifetime = config.MaxConnLifetime * time.Second
	}
	if config.Protocol != 2 && config.Protocol != 3 {
		config.Protocol = 3
	}
	return
}

// GetConfig 返回指定组名的全局配置。如果未传入 `name`，则返回默认组名的配置。
// md5:327a839e91668442
func GetConfig(name ...string) (config *Config, ok bool) {
	group := DefaultGroupName
	if len(name) > 0 {
		group = name[0]
	}
	if v := localConfigMap.Get(group); v != nil {
		return v.(*Config), true
	}
	return &Config{}, false
}

// RemoveConfig 删除指定组的全局配置。
// 如果没有传递 `name`，则删除默认组名的配置。
// md5:8e808827f299122d
func RemoveConfig(name ...string) {
	group := DefaultGroupName
	if len(name) > 0 {
		group = name[0]
	}
	localConfigMap.Remove(group)

	intlog.Printf(context.TODO(), `RemoveConfig: %s`, group)
}

// ClearConfig 删除所有的 Redis 配置。 md5:337bf67372d51962
func ClearConfig() {
	localConfigMap.Clear()
}
