// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

// 包gcmd提供控制台操作，如读取选项/参数和运行命令。 md5:bb72337a704c599f
package cmd类

import (
	"os"

	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/internal/utils"
	gctx "github.com/888go/goframe/os/gctx"
)

const (
	CtxKeyParser         gctx.StrKey = `CtxKeyParser`
	CtxKeyCommand        gctx.StrKey = `CtxKeyCommand`
	CtxKeyArgumentsIndex gctx.StrKey = `CtxKeyArgumentsIndex`
)

const (
	helpOptionName        = "help"
	helpOptionNameShort   = "h"
	maxLineChars          = 120
	tracingInstrumentName = "github.com/gogf/gf/v2/os/gcmd.Command"
	tagNameName           = "name"
	tagNameShort          = "short"
)

// Init 进行自定义初始化。 md5:08f8a2052942d9c8
func Init(args ...string) {
	command.Init(args...)
}

// GetOpt 作为gvar.Var返回名为`name`的选项值。 md5:1859b868ee779be0
func GetOpt(name string, def ...string) *gvar.Var {
	if v := command.GetOpt(name, def...); v != "" {
		return gvar.New(v)
	}
	if command.ContainsOpt(name) {
		return gvar.New("")
	}
	return nil
}

// GetOptAll 返回所有已解析的选项。 md5:6de4d266d8991786
func GetOptAll() map[string]string {
	return command.GetOptAll()
}

// GetArg 作为gvar.Var返回索引为`index`的参数。 md5:12ea2f8d74c6370d
func GetArg(index int, def ...string) *gvar.Var {
	if v := command.GetArg(index, def...); v != "" {
		return gvar.New(v)
	}
	return nil
}

// GetArgAll 返回所有解析的参数。 md5:85cc0fd5995d4878
func GetArgAll() []string {
	return command.GetArgAll()
}

// GetOptWithEnv 返回指定 `key` 的命令行参数。
// 如果该参数不存在，则返回指定 `key` 的环境变量。
// 如果两者都不存在，它将返回默认值 `def`。
//
// 获取规则：
// 1. 命令行参数采用小写格式，例如：gf.`包名`.<变量名>;
// 2. 环境变量采用大写格式，例如：GF_`包名`_<变量名>。
// md5:e3d5c0c773430740
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

// BuildOptions 将选项构建为字符串。 md5:c722b017f3a50346
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
