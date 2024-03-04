// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gcfg 提供了配置的读取、缓存和管理功能。
package gcfg

import (
	"context"

	"github.com/888go/goframe/gcfg/internal/command"
	"github.com/888go/goframe/gcfg/internal/intlog"
	"github.com/888go/goframe/gcfg/internal/utils"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/genv"
)

// Config 是配置管理对象。
type Config struct {
	adapter Adapter
}

const (
	DefaultInstanceName   = "config" // DefaultName 是用于实例使用的默认实例名称。
	DefaultConfigFileName = "config" // DefaultConfigFile 是默认的配置文件名称。
)

// New 创建并返回一个 Config 对象，其默认适配器为 AdapterFile。
func New() (*Config, error) {
	adapterFile, err := NewAdapterFile()
	if err != nil {
		return nil, err
	}
	return &Config{
		adapter: adapterFile,
	}, nil
}

// NewWithAdapter 使用给定的适配器创建并返回一个Config对象。
func NewWithAdapter(adapter Adapter) *Config {
	return &Config{
		adapter: adapter,
	}
}

// Instance 返回一个使用默认设置的 Config 实例。
// 参数 `name` 是该实例的名称。但请注意，如果配置目录中存在名为 "name.toml" 的文件，
// 则将其设置为默认配置文件。toml 文件类型是默认的配置文件类型。
func Instance(name ...string) *Config {
	var instanceName = DefaultInstanceName
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	return localInstances.GetOrSetFuncLock(instanceName, func() interface{} {
		adapterFile, err := NewAdapterFile()
		if err != nil {
			intlog.Errorf(context.Background(), `%+v`, err)
			return nil
		}
		if instanceName != DefaultInstanceName {
			adapterFile.SetFileName(instanceName)
		}
		return NewWithAdapter(adapterFile)
	}).(*Config)
}

// SetAdapter 设置当前 Config 对象的适配器。
func (c *Config) SetAdapter(adapter Adapter) {
	c.adapter = adapter
}

// GetAdapter 返回当前 Config 对象的适配器。
func (c *Config) GetAdapter() Adapter {
	return c.adapter
}

// Available 检查并返回配置服务是否可用。
// 可选参数 `pattern` 指定了特定的配置资源。
//
// 如果配置文件存在于默认的 AdapterFile 中，则返回 true，否则返回 false。
// 注意，此函数不会返回错误，因为它只是简单地检查后端配置服务是否存在。
func (c *Config) Available(ctx context.Context, resource ...string) (ok bool) {
	return c.adapter.Available(ctx, resource...)
}

// Get 方法通过指定的`pattern`获取并返回值。
// 如果给定的`pattern`为空字符串或"."，则返回当前Json对象的所有值。
// 若未找到通过`pattern`匹配到的值，则返回nil。
//
// 当通过`pattern`未能找到对应的值时，将返回由`def`指定的默认值。
func (c *Config) Get(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error) {
	var (
		err   error
		value interface{}
	)
	value, err = c.adapter.Get(ctx, pattern)
	if err != nil {
		return nil, err
	}
	if value == nil {
		if len(def) > 0 {
			return gvar.New(def[0]), nil
		}
		return nil, nil
	}
	return gvar.New(value), nil
}

// GetWithEnv 返回通过模式`pattern`指定的配置值。
// 如果配置值不存在，则获取并返回由`key`指定的环境变量值。
// 若两者都不存在，则返回默认值 `def`。
//
// 获取规则：环境变量参数采用大写格式，例如：GF_PACKAGE_VARIABLE。
func (c *Config) GetWithEnv(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error) {
	value, err := c.Get(ctx, pattern)
	if err != nil && gerror.Code(err) != gcode.CodeNotFound {
		return nil, err
	}
	if value == nil {
		if v := genv.Get(utils.FormatEnvKey(pattern)); v != nil {
			return v, nil
		}
		if len(def) > 0 {
			return gvar.New(def[0]), nil
		}
		return nil, nil
	}
	return value, nil
}

// GetWithCmd 根据模式 `pattern` 返回配置值。
// 如果配置值不存在，则获取并返回由 `key` 指定的命令行选项。
// 若两者都不存在，则返回默认值 `def`。
//
// 获取规则：命令行参数采用小写格式，例如：gf.package.variable。
func (c *Config) GetWithCmd(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error) {
	value, err := c.Get(ctx, pattern)
	if err != nil && gerror.Code(err) != gcode.CodeNotFound {
		return nil, err
	}
	if value == nil {
		if v := command.GetOpt(utils.FormatCmdKey(pattern)); v != "" {
			return gvar.New(v), nil
		}
		if len(def) > 0 {
			return gvar.New(def[0]), nil
		}
		return nil, nil
	}
	return value, nil
}

// Data 函数获取并以 map 类型返回所有配置数据。
func (c *Config) Data(ctx context.Context) (data map[string]interface{}, err error) {
	return c.adapter.Data(ctx)
}

// MustGet 行为类似于函数 Get，但在发生错误时会触发 panic。
func (c *Config) MustGet(ctx context.Context, pattern string, def ...interface{}) *gvar.Var {
	v, err := c.Get(ctx, pattern, def...)
	if err != nil {
		panic(err)
	}
	if v == nil {
		return nil
	}
	return v
}

// MustGetWithEnv 行为类似于函数 GetWithEnv，但是当发生错误时它会触发panic。
func (c *Config) MustGetWithEnv(ctx context.Context, pattern string, def ...interface{}) *gvar.Var {
	v, err := c.GetWithEnv(ctx, pattern, def...)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetWithCmd 的行为与函数 GetWithCmd 相同，但当出现错误时它会触发 panic。
func (c *Config) MustGetWithCmd(ctx context.Context, pattern string, def ...interface{}) *gvar.Var {
	v, err := c.GetWithCmd(ctx, pattern, def...)
	if err != nil {
		panic(err)
	}
	return v
}

// MustData 的行为与函数 Data 相同，但是当发生错误时，它会引发 panic（异常）。
func (c *Config) MustData(ctx context.Context) map[string]interface{} {
	v, err := c.Data(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
