// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

// Package gcmd 提供控制台操作功能，例如选项/参数读取和命令执行。
package gcmd

import (
	"os"
	
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/888go/goframe/gcmd/internal/command"
	"github.com/888go/goframe/gcmd/internal/utils"
	"github.com/gogf/gf/v2/os/gctx"
)

const (
	CtxKeyParser    gctx.StrKey = `CtxKeyParser`
	CtxKeyCommand   gctx.StrKey = `CtxKeyCommand`
	CtxKeyArguments gctx.StrKey = `CtxKeyArguments`
)

const (
	helpOptionName        = "help"
	helpOptionNameShort   = "h"
	maxLineChars          = 120
	tracingInstrumentName = "github.com/gogf/gf/v2/os/gcmd.Command"
	tagNameName           = "name"
	tagNameShort          = "short"
)

// Init 进行自定义初始化。
func Init(args ...string) {
	command.Init(args...)
}

// GetOpt 函数返回名为 `name` 的选项值，类型为 gvar.Var。
func GetOpt(name string, def ...string) *gvar.Var {
	if v := command.GetOpt(name, def...); v != "" {
		return gvar.New(v)
	}
	if command.ContainsOpt(name) {
		return gvar.New("")
	}
	return nil
}

// GetOptAll 返回所有已解析的选项。
func GetOptAll() map[string]string {
	return command.GetOptAll()
}

// GetArg 返回位于`index`处的参数作为gvar.Var类型。
func GetArg(index int, def ...string) *gvar.Var {
	if v := command.GetArg(index, def...); v != "" {
		return gvar.New(v)
	}
	return nil
}

// GetArgAll 返回所有已解析的参数。
func GetArgAll() []string {
	return command.GetArgAll()
}

// GetOptWithEnv returns the command line argument of the specified `key`.
// If the argument does not exist, then it returns the environment variable with specified `key`.
// It returns the default value `def` if none of them exists.
//
// Fetching Rules:
// 1. Command line arguments are in lowercase format, eg: gf.`package name`.<variable name>;
// 2. Environment arguments are in uppercase format, eg: GF_`package name`_<variable name>；
func GetOptWithEnv(key string, def ...interface{}) *gvar.Var {
	cmdKey := utils.FormatCmdKey(key)
	if command.ContainsOpt(cmdKey) {
		return gvar.New(GetOpt(cmdKey))
	} else {
		envKey := utils.FormatEnvKey(key)
		if r, ok := os.LookupEnv(envKey); ok {
			return gvar.New(r)
		} else {
			if len(def) > 0 {
				return gvar.New(def[0])
			}
		}
	}
	return nil
}

// BuildOptions 将选项构建为字符串。
func BuildOptions(m map[string]string, prefix ...string) string {
	options := ""
	leadStr := "-"
	if len(prefix) > 0 {
		leadStr = prefix[0]
	}
	for k, v := range m {
		if len(options) > 0 {
			options += " "
		}
		options += leadStr + k
		if v != "" {
			options += "=" + v
		}
	}
	return options
}
