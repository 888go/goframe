// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包genv提供了对系统环境变量的操作。 md5:9605f9d2a2186f5b
package 环境变量类

import (
	"fmt"
	"os"
	"strings"

	gvar "github.com/888go/goframe/container/gvar"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/internal/utils"
)

// X取全部返回一个字符串切片的副本，表示环境，形式为"key=value"。
// md5:723df5605f199f2b
func X取全部() []string {
	return os.Environ()
}

// X取Map 返回一个副本，该副本将环境表示为映射（map）形式的字符串。 md5:9477b6d266100b3d
func X取Map() map[string]string {
	return X切片到Map(os.Environ())
}

// X取值 根据给定的`key`创建并返回一个具有环境变量值的Var。如果环境中不存在该变量，则使用给定的`def`作为默认值。
// md5:1c5c61ffd2aa5106
func X取值(名称 string, 默认值 ...interface{}) *gvar.Var {
	v, ok := os.LookupEnv(名称)
	if !ok {
		if len(默认值) > 0 {
			return gvar.X创建(默认值[0])
		}
		return nil
	}
	return gvar.X创建(v)
}

// X设置值 设置环境变量的值，该变量由 `key` 指定。如果发生任何错误，它将返回一个错误。
// md5:3d9ca695de9bb4ad
func X设置值(key, value string) (err error) {
	err = os.Setenv(key, value)
	if err != nil {
		err = gerror.X多层错误并格式化(err, `set environment key-value failed with key "%s", value "%s"`, key, value)
	}
	return
}

// X设置Map值 使用映射设置环境变量。 md5:78d0cfffe3bc8311
func X设置Map值(m map[string]string) (错误 error) {
	for k, v := range m {
		if 错误 = X设置值(k, v); 错误 != nil {
			return 错误
		}
	}
	return nil
}

// X是否存在 检查名为 `key` 的环境变量是否存在。 md5:76124e3be6d217ff
func X是否存在(名称 string) bool {
	_, ok := os.LookupEnv(名称)
	return ok
}

// X删除 删除一个或多个环境变量。 md5:546a01a7df799055
func X删除(名称 ...string) (错误 error) {
	for _, v := range 名称 {
		if 错误 = os.Unsetenv(v); 错误 != nil {
			错误 = gerror.X多层错误并格式化(错误, `delete environment key failed with key "%s"`, v)
			return 错误
		}
	}
	return nil
}

// X取值或命令行 返回指定的环境变量值 `key`。
// 如果环境变量不存在，它将从命令行选项中检索并返回值。如果两者都不存在，它将返回默认值 `def`。
// 
// 获取规则：
// 1. 环境变量参数使用大写格式，例如：GF_<包名>_<变量名>；
// 2. 命令行参数使用小写格式，例如：gf.<包名>.<变量名>；
// md5:1bba2e845d6ee0d6
func X取值或命令行(名称 string, 默认值 ...interface{}) *gvar.Var {
	envKey := utils.FormatEnvKey(名称)
	if v := os.Getenv(envKey); v != "" {
		return gvar.X创建(v)
	}
	cmdKey := utils.FormatCmdKey(名称)
	if v := command.GetOpt(cmdKey); v != "" {
		return gvar.X创建(v)
	}
	if len(默认值) > 0 {
		return gvar.X创建(默认值[0])
	}
	return nil
}

// Map到切片 构建一个映射到环境变量切片的map。 md5:f58dc9490f9468a7
func Map到切片(m map[string]string) []string {
	array := make([]string, len(m))
	index := 0
	for k, v := range m {
		array[index] = k + "=" + v
		index++
	}
	return array
}

// X切片到Map 将环境变量从切片转换为映射。 md5:1c7b8b3cbc6a6d0d
func X切片到Map(切片 []string) map[string]string {
	m := make(map[string]string)
	i := 0
	for _, s := range 切片 {
		i = strings.IndexByte(s, '=')
		m[s[0:i]] = s[i+1:]
	}
	return m
}

// MapToEnv别名 将环境变量从映射转换为切片。 md5:3cef9db0baccea9f
func MapToEnv别名(m map[string]string) []string {
	envs := make([]string, 0)
	for k, v := range m {
		envs = append(envs, fmt.Sprintf(`%s=%s`, k, v))
	}
	return envs
}

// X切片去重 从给定的环境变量中过滤重复项。 md5:7b495d60bfff573e
func X切片去重(切片 []string) []string {
	return MapToEnv别名(X切片到Map(切片))
}
