// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

// Package command 提供控制台操作功能，如选项/参数读取。
package command
import (
	"os"
	"regexp"
	"strings"
	)
var (
	defaultParsedArgs    = make([]string, 0)
	defaultParsedOptions = make(map[string]string)
	argumentRegex        = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)
)

// Init 进行自定义初始化。
func Init(args ...string) {
	if len(args) == 0 {
		if len(defaultParsedArgs) == 0 && len(defaultParsedOptions) == 0 {
			args = os.Args
		} else {
			return
		}
	} else {
		defaultParsedArgs = make([]string, 0)
		defaultParsedOptions = make(map[string]string)
	}
	// 使用默认算法解析os.Args
	defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
}

// 使用默认算法解析参数。
func ParseUsingDefaultAlgorithm(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = make([]string, 0)
	parsedOptions = make(map[string]string)
	for i := 0; i < len(args); {
		array := argumentRegex.FindStringSubmatch(args[i])
		if len(array) > 2 {
			if array[2] == "=" {
				parsedOptions[array[1]] = array[3]
			} else if i < len(args)-1 {
				if len(args[i+1]) > 0 && args[i+1][0] == '-' {
					// 示例：gf gen -d -n 1
// （注释翻译：该行代码为命令行使用示例，表示在gf框架中执行命令生成数据，其中参数-d表示进行数据驱动模式操作，-n 1表示生成1条数据。）
					parsedOptions[array[1]] = array[3]
				} else {
					// Eg: gf gen -n 2
					parsedOptions[array[1]] = args[i+1]
					i += 2
					continue
				}
			} else {
				// Eg: gf gen -h
				parsedOptions[array[1]] = array[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}
	return
}

// GetOpt返回名为`name`的选项值。
func GetOpt(name string, def ...string) string {
	Init()
	if v, ok := defaultParsedOptions[name]; ok {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetOptAll 返回所有已解析的选项。
func GetOptAll() map[string]string {
	Init()
	return defaultParsedOptions
}

// ContainsOpt 检查选项名 `name` 是否存在于参数中。
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg 返回位于 `index` 索引处的参数。
func GetArg(index int, def ...string) string {
	Init()
	if index < len(defaultParsedArgs) {
		return defaultParsedArgs[index]
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetArgAll 返回所有已解析的参数。
func GetArgAll() []string {
	Init()
	return defaultParsedArgs
}

// GetOptWithEnv returns the command line argument of the specified `key`.
// If the argument does not exist, then it returns the environment variable with specified `key`.
// It returns the default value `def` if none of them exists.
//
// Fetching Rules:
// 1. Command line arguments are in lowercase format, eg: gf.package.variable;
// 2. Environment arguments are in uppercase format, eg: GF_PACKAGE_VARIABLE；
func GetOptWithEnv(key string, def ...string) string {
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	} else {
		envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		if r, ok := os.LookupEnv(envKey); ok {
			return r
		} else {
			if len(def) > 0 {
				return def[0]
			}
		}
	}
	return ""
}
