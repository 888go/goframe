// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gins

import (
	"github.com/gogf/gf/v2/i18n/gi18n"
)

// I18n returns an instance of gi18n.Manager.
// The parameter `name` is the name for the instance.

// ff:
// name:
func I18n(name ...string) *gi18n.Manager {
	return gi18n.Instance(name...)
}
