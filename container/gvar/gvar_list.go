// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gvar

import (
	"github.com/gogf/gf/v2/util/gutil"
)

// ListItemValues retrieves and returns the elements of all item struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.

// ff:取结构数组或Map数组值
// values:值s
// key:名称
func (v *Var) ListItemValues(key interface{}) (values []interface{}) {
	return gutil.ListItemValues(v.Val(), key)
}

// ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.

// ff:取结构数组或Map数组值并去重
// key:名称
func (v *Var) ListItemValuesUnique(key string) []interface{} {
	return gutil.ListItemValuesUnique(v.Val(), key)
}
