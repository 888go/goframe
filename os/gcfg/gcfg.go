// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 包gcfg提供了配置的读取、缓存和管理功能。 md5:5ae504d1379cd99a
package gcfg

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/command"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/os/genv"
)

// Config 是配置管理对象。 md5:1ef57338c678e400
type Config struct {
	adapter Adapter
}

const (
	DefaultInstanceName   = "config" // DefaultName 是实例使用的默认实例名称。 md5:4736f3b4285b6846
	DefaultConfigFileName = "config" // DefaultConfigFile 是默认的配置文件名。 md5:b558e9c92a774f9a
)

// New 创建并返回一个 Config 对象，其默认适配器为 AdapterFile。 md5:52cd678118524272
func New() (*Config, error) {
	adapterFile, err := NewAdapterFile()
	if err != nil {
		return nil, err
	}
	return &Config{
		adapter: adapterFile,
	}, nil
}

// NewWithAdapter使用给定的适配器创建并返回一个Config对象。 md5:9ddaae0ddb0e0297
func NewWithAdapter(adapter Adapter) *Config {
	return &Config{
		adapter: adapter,
	}
}

// Instance 返回一个具有默认设置的 Config 实例。
// 参数 `name` 是该实例的名称。但请注意，如果配置目录中存在文件 "name.toml"，
// 则将其设置为默认配置文件。TOML 文件类型是默认的配置文件类型。 md5:4164ff567a8c8c31
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

// SetAdapter 设置当前 Config 对象的适配器。 md5:8d00d377baafeb01
func (c *Config) SetAdapter(adapter Adapter) {
	c.adapter = adapter
}

// GetAdapter 返回当前Config对象的适配器。 md5:46c003ab367518d8
func (c *Config) GetAdapter() Adapter {
	return c.adapter
}

// 可用性检查并返回配置服务是否可用。
// 可选参数 `pattern` 指定某些配置资源。
//
// 如果默认AdapterFile中存在配置文件，则返回true，否则返回false。
// 请注意，此函数不会返回错误，因为它只是简单地检查后端配置服务。 md5:771d98d194158bc1
func (c *Config) Available(ctx context.Context, resource ...string) (ok bool) {
	return c.adapter.Available(ctx, resource...)
}

// Get 通过指定的`pattern`获取并返回值。
// 如果`pattern`为空字符串或"."，它将返回当前Json对象的所有值。
// 如果根据`pattern`没有找到值，它将返回nil。
//
// 如果没有为`pattern`找到值，它将返回由`def`指定的默认值。 md5:b10a106fb9d6af41
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

// GetWithEnv 根据模式`pattern`返回配置值。
// 如果配置值不存在，那么它会获取并返回由`key`指定的环境变量值。
// 如果两者都不存在，则返回默认值`def`。
//
// 获取规则：环境变量参数以大写格式表示，例如：GF_PACKAGE_VARIABLE。 md5:d533293fbfbf6350
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
// 如果找不到配置值，它将检索并返回由 `key` 指定的命令行选项。
// 如果它们都不存在，则返回默认值 `def`。
//
// 获取规则：命令行参数采用小写格式，例如：gf.package.variable。 md5:2a77887f42041d88
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

// Data 获取并以映射类型返回所有配置数据。 md5:2a92e8bbe7388f01
func (c *Config) Data(ctx context.Context) (data map[string]interface{}, err error) {
	return c.adapter.Data(ctx)
}

// MustGet 行为类似于函数 Get，但如果发生错误时会引发 panic。 md5:b1d3af83a52fd248
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

// MustGetWithEnv 作为 GetWithEnv 函数的行为，但如果发生错误，它会引发恐慌。 md5:9f816c41440b51cf
func (c *Config) MustGetWithEnv(ctx context.Context, pattern string, def ...interface{}) *gvar.Var {
	v, err := c.GetWithEnv(ctx, pattern, def...)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetWithCmd 的行为类似于 GetWithCmd 函数，但如果发生错误，它会直接 panic。 md5:683d24a1f4aceb7b
func (c *Config) MustGetWithCmd(ctx context.Context, pattern string, def ...interface{}) *gvar.Var {
	v, err := c.GetWithCmd(ctx, pattern, def...)
	if err != nil {
		panic(err)
	}
	return v
}

// MustData 行为类似于函数 Data，但如果发生错误则会引发恐慌。 md5:eb72c1ce036d70b6
func (c *Config) MustData(ctx context.Context) map[string]interface{} {
	v, err := c.Data(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
