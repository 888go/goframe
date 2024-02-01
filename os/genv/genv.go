// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package genv 提供了对系统环境变量的操作。
package genv
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
func All() []string {
	return os.Environ()
}

// Map 返回一个字符串环境表示的映射副本。
func Map() map[string]string {
	return MapFromEnv(os.Environ())
}

// Get 函数创建并返回一个 Var，其值为环境变量中名为 `key` 的变量的值。
// 如果该变量在环境中不存在，则使用给定的 `def` 作为默认值。
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

// Set 函数用于设置名为 `key` 的环境变量的值。
// 如果出现错误，该函数会返回一个错误。
func Set(key, value string) (err error) {
	err = os.Setenv(key, value)
	if err != nil {
		err = gerror.Wrapf(err, `set environment key-value failed with key "%s", value "%s"`, key, value)
	}
	return
}

// SetMap 通过 map 设置环境变量。
func SetMap(m map[string]string) (err error) {
	for k, v := range m {
		if err = Set(k, v); err != nil {
			return err
		}
	}
	return nil
}

// Contains 检查名为 `key` 的环境变量是否存在。
func Contains(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}

// Remove 删除一个或多个环境变量。
func Remove(key ...string) (err error) {
	for _, v := range key {
		if err = os.Unsetenv(v); err != nil {
			err = gerror.Wrapf(err, `delete environment key failed with key "%s"`, v)
			return err
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

// Build 函数用于构建一个映射到环境变量切片的映射。
func Build(m map[string]string) []string {
	array := make([]string, len(m))
	index := 0
	for k, v := range m {
		array[index] = k + "=" + v
		index++
	}
	return array
}

// MapFromEnv 将环境变量从切片转换为映射（map）。
func MapFromEnv(envs []string) map[string]string {
	m := make(map[string]string)
	i := 0
	for _, s := range envs {
		i = strings.IndexByte(s, '=')
		m[s[0:i]] = s[i+1:]
	}
	return m
}

// MapToEnv 将环境变量从映射转换为切片。
func MapToEnv(m map[string]string) []string {
	envs := make([]string, 0)
	for k, v := range m {
		envs = append(envs, fmt.Sprintf(`%s=%s`, k, v))
	}
	return envs
}

// Filter 从给定的环境变量中过滤掉重复项。
func Filter(envs []string) []string {
	return MapToEnv(MapFromEnv(envs))
}
