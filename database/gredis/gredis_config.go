// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gredis
import (
	"context"
	"crypto/tls"
	"time"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/util/gconv"
	)
// Config 是 Redis 配置。
type Config struct {
	// Address 地址 支持单个和集群模式的Redis服务器。多个地址使用逗号 ',' 连接。例如：192.168.1.1:6379, 192.168.1.2:6379。
	Address         string        `json:"address"`
	Db              int           `json:"db"`              // Redis db.
	User            string        `json:"user"`            // AUTH的用户名。
	Pass            string        `json:"pass"`            // AUTH的密码。
	MinIdle         int           `json:"minIdle"`         // 最小允许空闲的连接数（默认为0）
	MaxIdle         int           `json:"maxIdle"`         // 最大允许空闲连接数（默认为10）
	MaxActive       int           `json:"maxActive"`       // 连接数的最大限制（默认值为0，表示无限制）。
	MaxConnLifetime time.Duration `json:"maxConnLifetime"` // 连接的最大生命周期（默认为30秒，不允许设置为0）
	IdleTimeout     time.Duration `json:"idleTimeout"`     // 连接的最大空闲时间（默认为10秒，不允许设置为0）
	WaitTimeout     time.Duration `json:"waitTimeout"`     // 从连接池获取连接时超时的持续时间。
	DialTimeout     time.Duration `json:"dialTimeout"`     // 设置TCP连接的超时时间。
	ReadTimeout     time.Duration `json:"readTimeout"`     // TCP读取超时时间。如果不是必需的，请勿设置。
	WriteTimeout    time.Duration `json:"writeTimeout"`    // TCP写入超时时间
	MasterName      string        `json:"masterName"`      // 用于 Redis Sentinel 模式。
	TLS             bool          `json:"tls"`             // 指定在连接到服务器时是否应使用TLS（传输层安全协议）。
	TLSSkipVerify   bool          `json:"tlsSkipVerify"`   // 禁用通过TLS连接时的服务器名称验证。
	TLSConfig       *tls.Config   `json:"-"`               // TLS配置使用。当设置此配置时，将进行TLS协商。
	SlaveOnly       bool          `json:"slaveOnly"`       // 将所有命令路由到从节点（只读模式）。
	Cluster         bool          `json:"cluster"`         // 指定是否使用集群模式。
	Protocol        int           `json:"protocol"`        // 指定 RESP 协议版本（协议 2 或 3）
}

const (
	DefaultGroupName = "default" // 默认配置组名称。
)

var (
	// 配置组。
	localConfigMap = gmap.NewStrAnyMap(true)
)

// SetConfig 为指定的组设置全局配置。
// 如果未传递 `name`，则会为默认组名设置配置。
func SetConfig(config *Config, name ...string) {
	group := DefaultGroupName
	if len(name) > 0 {
		group = name[0]
	}
	localConfigMap.Set(group, config)

	intlog.Printf(context.TODO(), `SetConfig for group "%s": %+v`, group, config)
}

// SetConfigByMap 通过map设置指定组的全局配置。
// 如果未传递 `name`，则设置默认组名的配置。
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

// ConfigFromMap 从给定的 map 中解析并返回配置。
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

// GetConfig 返回具有指定组名的全局配置。
// 如果未传递 `name`，则返回默认组名的配置。
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
// 如果未传入 `name`，则移除默认组名的配置。
func RemoveConfig(name ...string) {
	group := DefaultGroupName
	if len(name) > 0 {
		group = name[0]
	}
	localConfigMap.Remove(group)

	intlog.Printf(context.TODO(), `RemoveConfig: %s`, group)
}

// ClearConfig 清除所有redis配置。
func ClearConfig() {
	localConfigMap.Clear()
}
