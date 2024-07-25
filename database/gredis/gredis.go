// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 包gredis提供了对Redis服务器的便捷客户端。
//
// Redis客户端。
//
// Redis命令官方文档：https://redis.io/commands
//
// Redis中文文档：http://redisdoc.com/ md5:fd856764d3114fd3
package gredis

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// AdapterFunc 是创建 Redis 适配器的函数。 md5:806d31217d679afd
type AdapterFunc func(config *Config) Adapter

var (
	// defaultAdapterFunc 是默认的适配器函数，用于创建 Redis 适配器。 md5:e518e287cdac2bb2
	defaultAdapterFunc AdapterFunc = func(config *Config) Adapter {
		return nil
	}
)

// New 创建并返回一个redis客户端。
// 它创建了一个默认的go-redis适配器。 md5:3f355ab0e775862a
func New(config ...*Config) (*Redis, error) {
	var (
		usedConfig  *Config
		usedAdapter Adapter
	)
	if len(config) > 0 && config[0] != nil {
		// 使用Go Redis实现的Redis客户端，根据给定配置进行适配。 md5:3df311ce390d85c6
		usedConfig = config[0]
		usedAdapter = defaultAdapterFunc(config[0])
	} else if configFromGlobal, ok := GetConfig(); ok {
		// 使用Go Redis实现的Redis客户端，它实现了包配置中的适配器。 md5:91dc0454a671c4de
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

// NewWithAdapter 使用给定的适配器创建并返回一个Redis客户端。 md5:ab7dc6695935087f
func NewWithAdapter(adapter Adapter) (*Redis, error) {
	if adapter == nil {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `adapter cannot be nil`)
	}
	redis := &Redis{localAdapter: adapter}
	return redis.initGroup(), nil
}

// RegisterAdapterFunc 注册创建默认 Redis 适配器的函数。 md5:2390765dfc7002cb
func RegisterAdapterFunc(adapterFunc AdapterFunc) {
	defaultAdapterFunc = adapterFunc
}
