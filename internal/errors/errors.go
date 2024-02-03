// 版权所有 GoFrame gf 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。

// Package errors 提供用于内部使用目的的错误处理功能。
package errors

import (
	"github.com/888go/goframe/internal/command"
)

// StackMode 是打印堆栈信息的模式，可在 StackModeBrief（简要模式）或 StackModeDetail（详细模式）中选择。
type StackMode string

const (
// commandEnvKeyForBrief 是用于简短错误堆栈切换键的命令环境名称。
// 已弃用：请改用 commandEnvKeyForStackMode。
	commandEnvKeyForBrief = "gf.gerror.brief"

	// commandEnvKeyForStackMode 是用于简略错误堆栈模式切换键的命令环境名称。
	commandEnvKeyForStackMode = "gf.gerror.stack.mode"
)

const (
	// StackModeBrief 指定错误堆栈打印时，不包含框架内部的错误堆栈信息。
	StackModeBrief StackMode = "brief"

	// StackModeDetail 指定打印所有错误堆栈，包括框架堆栈的详细错误堆栈。
	StackModeDetail StackMode = "detail"
)

var (
// stackModeConfigured 是配置的错误堆栈模式变量。
// 默认情况下，它是简洁堆栈模式。
	stackModeConfigured = StackModeBrief
)

func init() {
	// Deprecated.
	briefSetting := command.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// 错误堆栈模式通过命令行参数或环境变量进行配置。
	stackModeSetting := command.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackModeSettingMode := StackMode(stackModeSetting)
		switch stackModeSettingMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackModeSettingMode
		}
	}
}

// IsStackModeBrief 返回当前错误堆栈模式是否处于简洁模式。
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
