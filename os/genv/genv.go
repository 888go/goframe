// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package genv 提供了对系统环境变量的操作。
package 环境变量类

import (
	"fmt"
	"os"
	"strings"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/internal/utils"
)

// All 函数返回一个表示环境变量的字符串副本，其形式为 "key=value"。
func X取全部() []string {
	return os.Environ()
}

// Map 返回一个字符串环境表示的映射副本。
func X取Map() map[string]string {
	return X数组到Map(os.Environ())
}

// Get 函数创建并返回一个 Var，其值为环境变量中名为 `key` 的变量的值。
// 如果该变量在环境中不存在，则使用给定的 `def` 作为默认值。
func X取值(名称 string, 默认值 ...interface{}) *泛型类.Var {
	v, ok := os.LookupEnv(名称)
	if !ok {
		if len(默认值) > 0 {
			return 泛型类.X创建(默认值[0])
		}
		return nil
	}
	return 泛型类.X创建(v)
}

// Set 函数用于设置名为 `key` 的环境变量的值。
// 如果出现错误，该函数会返回一个错误。
func X设置值(名称, 值 string) (错误 error) {
	错误 = os.Setenv(名称, 值)
	if 错误 != nil {
		错误 = 错误类.X多层错误并格式化(错误, `set environment key-value failed with key "%s", value "%s"`, 名称, 值)
	}
	return
}

// SetMap 通过 map 设置环境变量。
func X设置Map值(m map[string]string) (错误 error) {
	for k, v := range m {
		if 错误 = X设置值(k, v); 错误 != nil {
			return 错误
		}
	}
	return nil
}

// Contains 检查名为 `key` 的环境变量是否存在。
func X是否存在(名称 string) bool {
	_, ok := os.LookupEnv(名称)
	return ok
}

// Remove 删除一个或多个环境变量。
func X删除(名称 ...string) (错误 error) {
	for _, v := range 名称 {
		if 错误 = os.Unsetenv(v); 错误 != nil {
			错误 = 错误类.X多层错误并格式化(错误, `delete environment key failed with key "%s"`, v)
			return 错误
		}
	}
	return nil
}

// GetWithCmd returns the environment value specified `key`.
// If the environment value does not exist, then it retrieves and returns the value from command line options.
// It returns the default value `def` if none of them exists.
//
// Fetching Rules:
// 1. Environment arguments are in uppercase format, eg: GF_<package name>_<variable name>；
// 2. Command line arguments are in lowercase format, eg: gf.<package name>.<variable name>;
func X取值或命令行(名称 string, 默认值 ...interface{}) *泛型类.Var {
	envKey := utils.FormatEnvKey(名称)
	if v := os.Getenv(envKey); v != "" {
		return 泛型类.X创建(v)
	}
	cmdKey := utils.FormatCmdKey(名称)
	if v := command.GetOpt(cmdKey); v != "" {
		return 泛型类.X创建(v)
	}
	if len(默认值) > 0 {
		return 泛型类.X创建(默认值[0])
	}
	return nil
}

// Build 函数用于构建一个映射到环境变量切片的映射。
func Map到数组(m map[string]string) []string {
	array := make([]string, len(m))
	index := 0
	for k, v := range m {
		array[index] = k + "=" + v
		index++
	}
	return array
}

// MapFromEnv 将环境变量从切片转换为映射（map）。
func X数组到Map(数组 []string) map[string]string {
	m := make(map[string]string)
	i := 0
	for _, s := range 数组 {
		i = strings.IndexByte(s, '=')
		m[s[0:i]] = s[i+1:]
	}
	return m
}

// MapToEnv 将环境变量从映射转换为切片。
func MapToEnv别名(m map[string]string) []string {
	envs := make([]string, 0)
	for k, v := range m {
		envs = append(envs, fmt.Sprintf(`%s=%s`, k, v))
	}
	return envs
}

// Filter 从给定的环境变量中过滤掉重复项。
func X数组去重(数组 []string) []string {
	return MapToEnv别名(X数组到Map(数组))
}
