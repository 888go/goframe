// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package 时间类

import (
	"time"
)

// wrapper is a wrapper for stdlib struct time.Time.
// It's used for overwriting some functions of time.Time, for example: String.
type wrapper struct {
	time.Time
}

// String overwrites the String function of time.Time.
func (t wrapper) String() string {
	if t.IsZero() {
		return ""
	}
	if t.Year() == 0 {
		// Only time.
		return t.Format("15:04:05")
	}
	return t.Format("2006-01-02 15:04:05")
}
