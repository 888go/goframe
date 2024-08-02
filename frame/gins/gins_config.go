// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gins

import (
	gcfg "github.com/888go/goframe/os/gcfg"
)

// Config returns an instance of View with default settings.
// The parameter `name` is the name for the instance.
func Config(name ...string) *gcfg.Config {
	return gcfg.Instance(name...)
}
