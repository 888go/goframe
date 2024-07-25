// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// gmode 包为项目提供发布模式管理功能。
//
// 它使用字符串而非整数来标记模式，便于配置。 md5:716bfe2e364994bd
package gmode

import (
	"github.com/gogf/gf/v2/debug/gdebug"
	"github.com/gogf/gf/v2/internal/command"
	"github.com/gogf/gf/v2/os/gfile"
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
	// 请注意，`currentMode` 不是并发安全的。 md5:71c58ea0b37583d8
	currentMode = NOT_SET
)

// Set 设置当前应用程序的模式。 md5:523a8be23f0521ca
func Set(mode string) {
	currentMode = mode
}

// SetDevelop 将当前应用的模式设置为 DEVELOP。 md5:7d10eaf834b69114
func SetDevelop() {
	Set(DEVELOP)
}

// SetTesting 将当前应用程序的运行模式设置为测试模式。 md5:3e3ee802bec9b04e
func SetTesting() {
	Set(TESTING)
}

// SetStaging 将当前应用的模式设置为 STAGING。 md5:c8e3fac819c1d0b9
func SetStaging() {
	Set(STAGING)
}

// SetProduct 将当前应用设置为PRODUCT模式。 md5:e681c0b16f9b2bf0
func SetProduct() {
	Set(PRODUCT)
}

// Mode 返回当前设置的应用模式。 md5:76410fdca2d2e6a9
func Mode() string {
	// 如果当前模式未设置，则执行此自动检查。 md5:a6a2104d461130ba
	if currentMode == NOT_SET {
		if v := command.GetOptWithEnv(commandEnvKey); v != "" {
			// 从命令行参数或环境设置的模式。 md5:9b483ed8828c68a6
			currentMode = v
		} else {
			// 如果找到源代码，则为开发模式，否则为产品模式。 md5:4feb614f3843b3df
			if gfile.Exists(gdebug.CallerFilePath()) {
				currentMode = DEVELOP
			} else {
				currentMode = PRODUCT
			}
		}
	}
	return currentMode
}

// IsDevelop 检查并返回当前应用是否正在运行在DEVELOP模式下。 md5:577f79eacdd2c47e
func IsDevelop() bool {
	return Mode() == DEVELOP
}

// IsTesting 检查并返回当前应用是否正在以测试模式运行。 md5:3411dababba12269
func IsTesting() bool {
	return Mode() == TESTING
}

// IsStaging 检查并返回当前应用程序是否在 STAGING 模式下运行。 md5:99c92c19f9a8925f
func IsStaging() bool {
	return Mode() == STAGING
}

// IsProduct 检查并返回当前应用程序是否正在PRODUCT模式下运行。 md5:cf7849a0659f26bb
func IsProduct() bool {
	return Mode() == PRODUCT
}
