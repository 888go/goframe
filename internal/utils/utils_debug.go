// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package utils

import (
	"github.com/gogf/gf/v2/internal/command"
)

const (
		// 调试键，用于检查是否处于调试模式。 md5:69f2af5509cc7334
	commandEnvKeyForDebugKey = "gf.debug"
)

var (
		// isDebugEnabled 标记是否启用了GoFrame的调试模式。 md5:50f038bff2bf5f20
	isDebugEnabled = false
)

func init() {
	// Debugging configured.
	value := command.GetOptWithEnv(commandEnvKeyForDebugKey)
	if value == "" || value == "0" || value == "false" {
		isDebugEnabled = false
	} else {
		isDebugEnabled = true
	}
}

// IsDebugEnabled 检查并返回是否启用了调试模式。
// 当命令参数 "gf.debug" 或环境变量 "GF_DEBUG" 被设置时，调试模式启用。
// md5:df7415f68212ff27
func IsDebugEnabled() bool {
	return isDebugEnabled
}

// SetDebugEnabled 启用/禁用内部调试信息。 md5:444b75c65786012e
func SetDebugEnabled(enabled bool) {
	isDebugEnabled = enabled
}
