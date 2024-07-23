// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包genv提供了对系统环境变量的操作。 md5:9605f9d2a2186f5b
package genv

import (
	"fmt"
	"os"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/command"
	"github.com/gogf/gf/v2/internal/utils"
)

// All返回一个字符串切片的副本，表示环境，形式为"key=value"。
// md5:723df5605f199f2b
func All() []string {
	return os.Environ()
}

// Map 返回一个副本，该副本将环境表示为映射（map）形式的字符串。 md5:9477b6d266100b3d
func Map() map[string]string {
	return MapFromEnv(os.Environ())
}

// Get 根据给定的`key`创建并返回一个具有环境变量值的Var。如果环境中不存在该变量，则使用给定的`def`作为默认值。
// md5:1c5c61ffd2aa5106
func Get(key string, def ...interface{}) *gvar.Var {
	v, ok := os.LookupEnv(key)
	if !ok {
		if len(def) > 0 {
			return gvar.New(def[0])
		}
		return nil
	}
	return gvar.New(v)
}

// Set 设置环境变量的值，该变量由 `key` 指定。如果发生任何错误，它将返回一个错误。
// md5:3d9ca695de9bb4ad
func Set(key, value string) (err error) {
	err = os.Setenv(key, value)
	if err != nil {
		err = gerror.Wrapf(err, `set environment key-value failed with key "%s", value "%s"`, key, value)
	}
	return
}

// SetMap 使用映射设置环境变量。 md5:78d0cfffe3bc8311
func SetMap(m map[string]string) (err error) {
	for k, v := range m {
		if err = Set(k, v); err != nil {
			return err
		}
	}
	return nil
}

// Contains 检查名为 `key` 的环境变量是否存在。 md5:76124e3be6d217ff
func Contains(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}

// Remove 删除一个或多个环境变量。 md5:546a01a7df799055
func Remove(key ...string) (err error) {
	for _, v := range key {
		if err = os.Unsetenv(v); err != nil {
			err = gerror.Wrapf(err, `delete environment key failed with key "%s"`, v)
			return err
		}
	}
	return nil
}

// GetWithCmd 返回指定的环境变量值 `key`。
// 如果环境变量不存在，它将从命令行选项中检索并返回值。如果两者都不存在，它将返回默认值 `def`。
// 
// 获取规则：
// 1. 环境变量参数使用大写格式，例如：GF_<包名>_<变量名>；
// 2. 命令行参数使用小写格式，例如：gf.<包名>.<变量名>；
// md5:1bba2e845d6ee0d6
func GetWithCmd(key string, def ...interface{}) *gvar.Var {
	envKey := utils.FormatEnvKey(key)
	if v := os.Getenv(envKey); v != "" {
		return gvar.New(v)
	}
	cmdKey := utils.FormatCmdKey(key)
	if v := command.GetOpt(cmdKey); v != "" {
		return gvar.New(v)
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}

// Build 构建一个映射到环境变量切片的map。 md5:f58dc9490f9468a7
func Build(m map[string]string) []string {
	array := make([]string, len(m))
	index := 0
	for k, v := range m {
		array[index] = k + "=" + v
		index++
	}
	return array
}

// MapFromEnv 将环境变量从切片转换为映射。 md5:1c7b8b3cbc6a6d0d
func MapFromEnv(envs []string) map[string]string {
	m := make(map[string]string)
	i := 0
	for _, s := range envs {
		i = strings.IndexByte(s, '=')
		m[s[0:i]] = s[i+1:]
	}
	return m
}

// MapToEnv 将环境变量从映射转换为切片。 md5:3cef9db0baccea9f
func MapToEnv(m map[string]string) []string {
	envs := make([]string, 0)
	for k, v := range m {
		envs = append(envs, fmt.Sprintf(`%s=%s`, k, v))
	}
	return envs
}

// Filter 从给定的环境变量中过滤重复项。 md5:7b495d60bfff573e
func Filter(envs []string) []string {
	return MapToEnv(MapFromEnv(envs))
}
