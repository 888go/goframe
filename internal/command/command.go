// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。 md5:a114f4bdd106ab31

// 包command提供控制台操作，如选项/参数读取。 md5:940e3926fff20c9a
package command

import (
	"os"
	"regexp"
	"strings"
)

var (
	defaultParsedArgs    = make([]string, 0)
	defaultParsedOptions = make(map[string]string)
	argumentOptionRegex  = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)
)

// Init 进行自定义初始化。 md5:08f8a2052942d9c8
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
	// 使用默认算法解析os.Args。 md5:460cde73efccc8f2
	defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
}

// 使用默认算法解析参数。 md5:d48f7f39a81379e1
func ParseUsingDefaultAlgorithm(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = make([]string, 0)
	parsedOptions = make(map[string]string)
	for i := 0; i < len(args); {
		array := argumentOptionRegex.FindStringSubmatch(args[i])
		if len(array) > 2 {
			if array[2] == "=" {
				parsedOptions[array[1]] = array[3]
			} else if i < len(args)-1 {
				if len(args[i+1]) > 0 && args[i+1][0] == '-' {
					// 这段Go语言的注释翻译成中文是：“这是一个示例命令行：使用gf（一个工具）生成代码，指定-d表示启用调试模式，-n 1表示生成1个示例。”. md5:a66be18aee4c44b9
					parsedOptions[array[1]] = array[3]
				} else {
					// Example: gf gen -n 2
					parsedOptions[array[1]] = args[i+1]
					i += 2
					continue
				}
			} else {
				// Example: gf gen -h
				parsedOptions[array[1]] = array[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}
	return
}

// GetOpt 返回名为 `name` 的选项值。 md5:5de4cb85c231ce6b
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

// GetOptAll 返回所有已解析的选项。 md5:6de4d266d8991786
func GetOptAll() map[string]string {
	Init()
	return defaultParsedOptions
}

// ContainsOpt 检查参数中是否存在名为 `name` 的选项。 md5:32ce4c1cf77651fb
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg 返回索引为 `index` 的参数。 md5:822343a8734fe602
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

// GetArgAll 返回所有解析的参数。 md5:85cc0fd5995d4878
func GetArgAll() []string {
	Init()
	return defaultParsedArgs
}

// GetOptWithEnv 返回指定 `key` 的命令行参数。
// 如果该参数不存在，则返回指定 `key` 的环境变量。
// 如果两者都不存在，它将返回默认值 `def`。
//
// 获取规则：
// 1. 命令行参数采用小写格式，例如：gf.package.variable；
// 2. 环境变量采用大写格式，例如：GF_PACKAGE_VARIABLE。 md5:13bcb9c2795488a1
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
