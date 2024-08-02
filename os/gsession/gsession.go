// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gsession implements manager and storage features for sessions.
package session类

import (
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	guid "github.com/888go/goframe/util/guid"
)

var (
	// ErrorDisabled is used for marking certain interface function not used.
	ErrorDisabled = gerror.NewWithOption(gerror.Option{
		Text: "this feature is disabled in this storage",
		Code: gcode.CodeNotSupported,
	})
)

// NewSessionId creates and returns a new and unique session id string,
// which is in 32 bytes.
func NewSessionId() string {
	return guid.S()
}
