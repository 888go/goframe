// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

// 包 errors 为内部使用提供处理错误的功能。. md5:5ebd576fdf2856b2
package errors

import (
	"github.com/gogf/gf/v2/internal/command"
)

// StackMode是在StackModeBrief或StackModeDetail模式下打印堆栈信息的模式。. md5:2ad08dc7990f413c
type StackMode string

const (
// commandEnvKeyForBrief 是用于切换简要错误堆栈的命令环境变量名。
// 已弃用：请使用 commandEnvKeyForStackMode 代替。
// md5:ec191a8df835f7da
	commandEnvKeyForBrief = "gf.gerror.brief"

	// commandEnvKeyForStackMode 是用于切换简短错误堆栈的命令环境名称。. md5:23f1a1bc661e992f
	commandEnvKeyForStackMode = "gf.gerror.stack.mode"
)

const (
	// StackModeBrief 指定只打印所有错误堆栈，不打印框架错误堆栈。. md5:b592fb9cad302a5a
	StackModeBrief StackMode = "brief"

	// StackModeDetail 指定打印详细错误堆栈，包括框架堆栈。. md5:fab5a35426363d26
	StackModeDetail StackMode = "detail"
)

var (
// stackModeConfigured 是配置错误堆栈模式的变量。
// 默认情况下，它处于简略堆栈模式。
// md5:5f27ddfb8b3441b5
	stackModeConfigured = StackModeBrief
)

func init() {
	// Deprecated.
	briefSetting := command.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// 错误堆栈模式通过命令行参数或环境变量进行配置。. md5:a6032170fbced764
	stackModeSetting := command.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackModeSettingMode := StackMode(stackModeSetting)
		switch stackModeSettingMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackModeSettingMode
		}
	}
}

// IsStackModeBrief 返回当前错误堆栈模式是否为简洁模式。. md5:5cccc95d12155d45
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
