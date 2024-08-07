// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gcfg提供了配置的读取、缓存和管理功能。 md5:5ae504d1379cd99a
package 配置类

import (
	"context"

	gvar "github.com/888go/goframe/container/gvar"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/utils"
	genv "github.com/888go/goframe/os/genv"
)

// Config 是配置管理对象。 md5:1ef57338c678e400
type Config struct {
	adapter Adapter
}

const (
	X默认实例名称   = "config" // DefaultName 是实例使用的默认实例名称。 md5:4736f3b4285b6846
	X默认配置文件名称 = "config" // DefaultConfigFile 是默认的配置文件名。 md5:b558e9c92a774f9a
)

// X创建 创建并返回一个 Config 对象，其默认适配器为 AdapterFile。 md5:52cd678118524272
func X创建() (*Config, error) {
	adapterFile, err := NewAdapterFile()
	if err != nil {
		return nil, err
	}
	return &Config{
		adapter: adapterFile,
	}, nil
}

// X创建并按适配器使用给定的适配器创建并返回一个Config对象。 md5:9ddaae0ddb0e0297
func X创建并按适配器(适配器 Adapter) *Config {
	return &Config{
		adapter: 适配器,
	}
}

// X取单例对象 返回一个具有默认设置的 Config 实例。
// 参数 `name` 是该实例的名称。但请注意，如果配置目录中存在文件 "name.toml"，
// 则将其设置为默认配置文件。TOML 文件类型是默认的配置文件类型。
// md5:4164ff567a8c8c31
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

// X设置适配器 设置当前 Config 对象的适配器。 md5:8d00d377baafeb01
func (c *Config) X设置适配器(适配器 Adapter) {
	c.adapter = 适配器
}

// X取适配器 返回当前Config对象的适配器。 md5:46c003ab367518d8
func (c *Config) X取适配器() Adapter {
	return c.adapter
}

// 可用性检查并返回配置服务是否可用。
// 可选参数 `pattern` 指定某些配置资源。
// 
// 如果默认AdapterFile中存在配置文件，则返回true，否则返回false。
// 请注意，此函数不会返回错误，因为它只是简单地检查后端配置服务。
// md5:771d98d194158bc1
func (c *Config) X是否可用(上下文 context.Context, resource ...string) (可用 bool) {
	return c.adapter.Available(上下文, resource...)
}

// X取值 通过指定的`pattern`获取并返回值。
// 如果`pattern`为空字符串或"."，它将返回当前Json对象的所有值。
// 如果根据`pattern`没有找到值，它将返回nil。
//
// 如果没有为`pattern`找到值，它将返回由`def`指定的默认值。
// md5:b10a106fb9d6af41
func (c *Config) X取值(上下文 context.Context, 表达式 string, 默认值 ...interface{}) (*gvar.Var, error) {
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
			return gvar.X创建(默认值[0]), nil
		}
		return nil, nil
	}
	return gvar.X创建(value), nil
}

// X取值并从环境变量 根据模式`pattern`返回配置值。
// 如果配置值不存在，那么它会获取并返回由`key`指定的环境变量值。
// 如果两者都不存在，则返回默认值`def`。
//
// 获取规则：环境变量参数以大写格式表示，例如：GF_PACKAGE_VARIABLE。
// md5:d533293fbfbf6350
func (c *Config) X取值并从环境变量(上下文 context.Context, 表达式 string, 默认值 ...interface{}) (*gvar.Var, error) {
	value, err := c.X取值(上下文, 表达式)
	if err != nil && gerror.X取错误码(err) != gcode.CodeNotFound {
		return nil, err
	}
	if value == nil {
		if v := genv.X取值(utils.FormatEnvKey(表达式)); v != nil {
			return v, nil
		}
		if len(默认值) > 0 {
			return gvar.X创建(默认值[0]), nil
		}
		return nil, nil
	}
	return value, nil
}

// X取值并从启动命令 根据模式 `pattern` 返回配置值。
// 如果找不到配置值，它将检索并返回由 `key` 指定的命令行选项。
// 如果它们都不存在，则返回默认值 `def`。
// 
// 获取规则：命令行参数采用小写格式，例如：gf.package.variable。
// md5:2a77887f42041d88
func (c *Config) X取值并从启动命令(上下文 context.Context, 表达式 string, 默认值 ...interface{}) (*gvar.Var, error) {
	value, err := c.X取值(上下文, 表达式)
	if err != nil && gerror.X取错误码(err) != gcode.CodeNotFound {
		return nil, err
	}
	if value == nil {
		if v := command.GetOpt(utils.FormatCmdKey(表达式)); v != "" {
			return gvar.X创建(v), nil
		}
		if len(默认值) > 0 {
			return gvar.X创建(默认值[0]), nil
		}
		return nil, nil
	}
	return value, nil
}

// X取Map 获取并以映射类型返回所有配置数据。 md5:2a92e8bbe7388f01
func (c *Config) X取Map(上下文 context.Context) (值 map[string]interface{}, 错误 error) {
	return c.adapter.Data(上下文)
}

// X取值PANI 行为类似于函数 Get，但如果发生错误时会引发 panic。 md5:b1d3af83a52fd248
func (c *Config) X取值PANI(上下文 context.Context, 表达式 string, 默认值 ...interface{}) *gvar.Var {
	v, err := c.X取值(上下文, 表达式, 默认值...)
	if err != nil {
		panic(err)
	}
	if v == nil {
		return nil
	}
	return v
}

// X取值并从环境变量PANI 作为 GetWithEnv 函数的行为，但如果发生错误，它会引发恐慌。 md5:9f816c41440b51cf
func (c *Config) X取值并从环境变量PANI(上下文 context.Context, 表达式 string, 默认值 ...interface{}) *gvar.Var {
	v, err := c.X取值并从环境变量(上下文, 表达式, 默认值...)
	if err != nil {
		panic(err)
	}
	return v
}

// X取值并从启动命令PANI_有bug 的行为类似于 GetWithCmd 函数，但如果发生错误，它会直接 panic。 md5:683d24a1f4aceb7b
func (c *Config) X取值并从启动命令PANI_有bug(上下文 context.Context, 表达式 string, 默认值 ...interface{}) *gvar.Var {
	v, err := c.X取值并从启动命令(上下文, 表达式, 默认值...)
	if err != nil {
		panic(err)
	}
	return v
}

// X取MapPANI 行为类似于函数 Data，但如果发生错误则会引发恐慌。 md5:eb72c1ce036d70b6
func (c *Config) X取MapPANI(上下文 context.Context) map[string]interface{} {
	v, err := c.X取Map(上下文)
	if err != nil {
		panic(err)
	}
	return v
}
