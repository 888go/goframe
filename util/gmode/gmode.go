// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gmode 为项目提供发布模式管理功能。
//
// 它使用字符串而非整数来标记模式，这种方式对于配置更为方便。
package gmode
import (
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/os/gfile"
	)
const (
	NOT_SET       = "not-set"
	DEVELOP       = "develop"
	TESTING       = "testing"
	STAGING       = "staging"
	PRODUCT       = "product"
	commandEnvKey = "gf.gmode"
)

var (
	// 注意，`currentMode` 不是线程安全的。
	currentMode = NOT_SET
)

// Set 设置当前应用程序的模式。
func Set(mode string) {
	currentMode = mode
}

// SetDevelop 将当前应用程序的当前模式设置为 DEVELOP 模式。
func SetDevelop() {
	Set(DEVELOP)
}

// SetTesting 设置当前应用程序的当前模式为 TESTING。
func SetTesting() {
	Set(TESTING)
}

// SetStaging 将当前应用的当前模式设置为 STAGING。
func SetStaging() {
	Set(STAGING)
}

// SetProduct 将当前应用的模式设置为PRODUCT模式。
func SetProduct() {
	Set(PRODUCT)
}

// Mode 返回当前设置的应用程序模式。
func Mode() string {
	// 如果当前模式未设置，则执行此自动检查。
	if currentMode == NOT_SET {
		if v := command.GetOptWithEnv(commandEnvKey); v != "" {
			// Mode 由命令行参数或环境变量配置。
			currentMode = v
		} else {
			// 如果找到了源代码，那么处于开发模式；否则，处于产品模式。
			if gfile.Exists(gdebug.CallerFilePath()) {
				currentMode = DEVELOP
			} else {
				currentMode = PRODUCT
			}
		}
	}
	return currentMode
}

// IsDevelop 检查并返回当前应用程序是否在开发模式下运行。
func IsDevelop() bool {
	return Mode() == DEVELOP
}

// IsTesting 检查并返回当前应用程序是否处于测试模式运行。
func IsTesting() bool {
	return Mode() == TESTING
}

// IsStaging 检查并返回当前应用程序是否在 STAGING（暂存）模式下运行。
func IsStaging() bool {
	return Mode() == STAGING
}

// IsProduct 检查并返回当前应用程序是否在PRODUCT模式下运行。
func IsProduct() bool {
	return Mode() == PRODUCT
}
