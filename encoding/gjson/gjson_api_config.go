// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gjson

// SetSplitChar sets the separator char for hierarchical data access.
// ff:设置参数分隔符
// j:
// char:分隔符
func (j *Json) SetSplitChar(char byte) {
	j.mu.Lock()
	j.c = char
	j.mu.Unlock()
}

// SetViolenceCheck enables/disables violence check for hierarchical data access.
// ff:设置分层冲突检查
// j:
// enabled:启用
func (j *Json) SetViolenceCheck(enabled bool) {
	j.mu.Lock()
	j.vc = enabled
	j.mu.Unlock()
}
