// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gmode provides release mode management for project.
//
// It uses string to mark the mode instead of integer, which is convenient for configuration.
package gmode//bm:环境类

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
	// Note that `currentMode` is not concurrent safe.
	currentMode = NOT_SET
)

// Set sets the mode for current application.

// ff:设置值
// yx:true
// mode:
func Set(mode string) {
	currentMode = mode
}

// SetDevelop sets current mode DEVELOP for current application.

// ff:
func SetDevelop() {
	Set(DEVELOP)
}

// SetTesting sets current mode TESTING for current application.

// ff:
func SetTesting() {
	Set(TESTING)
}

// SetStaging sets current mode STAGING for current application.

// ff:
func SetStaging() {
	Set(STAGING)
}

// SetProduct sets current mode PRODUCT for current application.

// ff:
func SetProduct() {
	Set(PRODUCT)
}

// Mode returns current application mode set.

// ff:
func Mode() string {
	// If current mode is not set, do this auto check.
	if currentMode == NOT_SET {
		if v := command.GetOptWithEnv(commandEnvKey); v != "" {
			// Mode configured from command argument of environment.
			currentMode = v
		} else {
			// If there are source codes found, it's in develop mode, or else in product mode.
			if gfile.Exists(gdebug.CallerFilePath()) {
				currentMode = DEVELOP
			} else {
				currentMode = PRODUCT
			}
		}
	}
	return currentMode
}

// IsDevelop checks and returns whether current application is running in DEVELOP mode.

// ff:
func IsDevelop() bool {
	return Mode() == DEVELOP
}

// IsTesting checks and returns whether current application is running in TESTING mode.

// ff:
func IsTesting() bool {
	return Mode() == TESTING
}

// IsStaging checks and returns whether current application is running in STAGING mode.

// ff:
func IsStaging() bool {
	return Mode() == STAGING
}

// IsProduct checks and returns whether current application is running in PRODUCT mode.

// ff:
func IsProduct() bool {
	return Mode() == PRODUCT
}
