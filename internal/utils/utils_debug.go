// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package utils
import (
	"github.com/888go/goframe/internal/command"
	)
const (
	// Debug键，用于检查是否处于调试模式。
	commandEnvKeyForDebugKey = "gf.debug"
)

var (
	// isDebugEnabled 标记 GoFrame 是否启用了调试模式。
	isDebugEnabled = false
)

func init() {
	// 已配置调试。
	value := command.GetOptWithEnv(commandEnvKeyForDebugKey)
	if value == "" || value == "0" || value == "false" {
		isDebugEnabled = false
	} else {
		isDebugEnabled = true
	}
}

// IsDebugEnabled 检查并返回是否启用了调试模式。
// 当命令行参数 "gf.debug" 或环境变量 "GF_DEBUG" 被设置时，调试模式将被启用。
func IsDebugEnabled() bool {
	return isDebugEnabled
}

// SetDebugEnabled 用于开启或关闭内部调试信息。
func SetDebugEnabled(enabled bool) {
	isDebugEnabled = enabled
}
