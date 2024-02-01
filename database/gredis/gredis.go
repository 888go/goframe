// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gredis 提供了一个方便的 Redis 服务器客户端。
//
// Redis 客户端。
//
// Redis 官方命令参考: https://redis.io/commands
//
// Redis 中文文档: http://redisdoc.com/
package gredis
import (
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	)
// AdapterFunc 是用于创建 Redis 适配器的函数。
type AdapterFunc func(config *Config) Adapter

var (
	// defaultAdapterFunc 是默认适配器函数，用于创建 Redis 适配器。
	defaultAdapterFunc AdapterFunc = func(config *Config) Adapter {
		return nil
	}
)

// New 创建并返回一个 Redis 客户端。
// 它创建了一个 go-redis 的默认 Redis 适配器。
func New(config ...*Config) (*Redis, error) {
	var (
		usedConfig  *Config
		usedAdapter Adapter
	)
	if len(config) > 0 && config[0] != nil {
		// 根据给定配置实现 Redis 客户端，使用 Go 语言的 Redis 库作为适配器。
		usedConfig = config[0]
		usedAdapter = defaultAdapterFunc(config[0])
	} else if configFromGlobal, ok := GetConfig(); ok {
		// 使用go redis实现Redis客户端，该客户端适配了包配置中的相关设置。
		usedConfig = configFromGlobal
		usedAdapter = defaultAdapterFunc(configFromGlobal)
	}
	if usedConfig == nil {
		return nil, gerror.NewCode(
			gcode.CodeInvalidConfiguration,
			`no configuration found for creating Redis client`,
		)
	}
	if usedAdapter == nil {
		return nil, gerror.NewCode(
			gcode.CodeNecessaryPackageNotImport,
			errorNilAdapter,
		)
	}
	redis := &Redis{
		config:       usedConfig,
		localAdapter: usedAdapter,
	}
	return redis.initGroup(), nil
}

// NewWithAdapter 根据给定的适配器创建并返回一个redis客户端。
func NewWithAdapter(adapter Adapter) (*Redis, error) {
	if adapter == nil {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `adapter cannot be nil`)
	}
	redis := &Redis{localAdapter: adapter}
	return redis.initGroup(), nil
}

// RegisterAdapterFunc 注册一个默认的创建 redis 适配器的函数。
func RegisterAdapterFunc(adapterFunc AdapterFunc) {
	defaultAdapterFunc = adapterFunc
}
