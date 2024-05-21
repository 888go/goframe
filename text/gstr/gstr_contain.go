// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gstr

import "strings"

// Contains reports whether `substr` is within `str`, case-sensitively.

// ff:是否包含
// substr:欲寻找的文本
// str:文本
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// ContainsI reports whether substr is within str, case-insensitively.

// ff:是否包含并忽略大小写
// substr:欲寻找的文本
// str:文本
func ContainsI(str, substr string) bool {
	return PosI(str, substr) != -1
}

// ContainsAny reports whether any Unicode code points in `chars` are within `s`.

// ff:是否包含Any
// chars:欲寻找的文本
// s:文本
func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}
