// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gcfg 提供了配置的读取、缓存和管理功能。
package 配置类

import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/genv"
)

// Config 是配置管理对象。
type Config struct {
	adapter Adapter
}

const (
	X默认实例名称   = "config" // DefaultName 是用于实例使用的默认实例名称。
	X默认配置文件名称 = "config" // DefaultConfigFile 是默认的配置文件名称。
)

// New 创建并返回一个 Config 对象，其默认适配器为 AdapterFile。
func X创建() (*Config, error) {
	adapterFile, err := NewAdapterFile()
	if err != nil {
		return nil, err
	}
	return &Config{
		adapter: adapterFile,
	}, nil
}

// NewWithAdapter 使用给定的适配器创建并返回一个Config对象。
func X创建并按适配器(适配器 Adapter) *Config {
	return &Config{
		adapter: 适配器,
	}
}

// Instance 返回一个使用默认设置的 Config 实例。
// 参数 `name` 是该实例的名称。但请注意，如果配置目录中存在名为 "name.toml" 的文件，
// 则将其设置为默认配置文件。toml 文件类型是默认的配置文件类型。
func X取单例对象(名称 ...string) *Config {
	var instanceName = X默认实例名称
	if len(名称) > 0 && 名称[0] != "" {
		instanceName = 名称[0]
	}
	return localInstances.X取值或设置值_函数带锁(instanceName, func() interface{} {
		adapterFile, err := NewAdapterFile()
		if err != nil {
			intlog.Errorf(context.Background(), `%+v`, err)
			return nil
		}
		if instanceName != X默认实例名称 {
			adapterFile.SetFileName(instanceName)
		}
		return X创建并按适配器(adapterFile)
	}).(*Config)
}

// SetAdapter 设置当前 Config 对象的适配器。
func (c *Config) X设置适配器(适配器 Adapter) {
	c.adapter = 适配器
}

// GetAdapter 返回当前 Config 对象的适配器。
func (c *Config) X取适配器() Adapter {
	return c.adapter
}

// Available 检查并返回配置服务是否可用。
// 可选参数 `pattern` 指定了特定的配置资源。
//
// 如果配置文件存在于默认的 AdapterFile 中，则返回 true，否则返回 false。
// 注意，此函数不会返回错误，因为它只是简单地检查后端配置服务是否存在。
func (c *Config) X是否可用(上下文 context.Context, resource ...string) (可用 bool) {
	return c.adapter.Available(上下文, resource...)
}

// Get 方法通过指定的`pattern`获取并返回值。
// 如果给定的`pattern`为空字符串或"."，则返回当前Json对象的所有值。
// 若未找到通过`pattern`匹配到的值，则返回nil。
//
// 当通过`pattern`未能找到对应的值时，将返回由`def`指定的默认值。
func (c *Config) X取值(上下文 context.Context, 表达式 string, 默认值 ...interface{}) (*泛型类.Var, error) {
	var (
		err   error
		value interface{}
	)
	value, err = c.adapter.Get(上下文, 表达式)
	if err != nil {
		return nil, err
	}
	if value == nil {
		if len(默认值) > 0 {
			return 泛型类.X创建(默认值[0]), nil
		}
		return nil, nil
	}
	return 泛型类.X创建(value), nil
}

// GetWithEnv 返回通过模式`pattern`指定的配置值。
// 如果配置值不存在，则获取并返回由`key`指定的环境变量值。
// 若两者都不存在，则返回默认值 `def`。
//
// 获取规则：环境变量参数采用大写格式，例如：GF_PACKAGE_VARIABLE。
func (c *Config) X取值并从环境变量(上下文 context.Context, 表达式 string, 默认值 ...interface{}) (*泛型类.Var, error) {
	value, err := c.X取值(上下文, 表达式)
	if err != nil && 错误类.X取错误码(err) != 错误码类.CodeNotFound {
		return nil, err
	}
	if value == nil {
		if v := 环境变量类.X取值(utils.FormatEnvKey(表达式)); v != nil {
			return v, nil
		}
		if len(默认值) > 0 {
			return 泛型类.X创建(默认值[0]), nil
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
func (c *Config) X取值并从启动命令(上下文 context.Context, 表达式 string, 默认值 ...interface{}) (*泛型类.Var, error) {
	value, err := c.X取值(上下文, 表达式)
	if err != nil && 错误类.X取错误码(err) != 错误码类.CodeNotFound {
		return nil, err
	}
	if value == nil {
		if v := command.GetOpt(utils.FormatCmdKey(表达式)); v != "" {
			return 泛型类.X创建(v), nil
		}
		if len(默认值) > 0 {
			return 泛型类.X创建(默认值[0]), nil
		}
		return nil, nil
	}
	return value, nil
}

// Data 函数获取并以 map 类型返回所有配置数据。
func (c *Config) X取Map(上下文 context.Context) (值 map[string]interface{}, 错误 error) {
	return c.adapter.Data(上下文)
}

// MustGet 行为类似于函数 Get，但在发生错误时会触发 panic。
func (c *Config) X取值PANI(上下文 context.Context, 表达式 string, 默认值 ...interface{}) *泛型类.Var {
	v, err := c.X取值(上下文, 表达式, 默认值...)
	if err != nil {
		panic(err)
	}
	if v == nil {
		return nil
	}
	return v
}

// MustGetWithEnv 行为类似于函数 GetWithEnv，但是当发生错误时它会触发panic。
func (c *Config) X取值并从环境变量PANI(上下文 context.Context, 表达式 string, 默认值 ...interface{}) *泛型类.Var {
	v, err := c.X取值并从环境变量(上下文, 表达式, 默认值...)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetWithCmd 的行为与函数 GetWithCmd 相同，但当出现错误时它会触发 panic。
func (c *Config) X取值并从启动命令PANI_有bug(上下文 context.Context, 表达式 string, 默认值 ...interface{}) *泛型类.Var {
	v, err := c.X取值并从启动命令(上下文, 表达式, 默认值...)
	if err != nil {
		panic(err)
	}
	return v
}

// MustData 的行为与函数 Data 相同，但是当发生错误时，它会引发 panic（异常）。
func (c *Config) X取MapPANI(上下文 context.Context) map[string]interface{} {
	v, err := c.X取Map(上下文)
	if err != nil {
		panic(err)
	}
	return v
}
