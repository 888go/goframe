// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

import (
	"github.com/888go/goframe/internal/json"
)

// SliceMap 是 Maps 的别名。 md5:af436f87335633ac
func SliceMap(any interface{}, option ...MapOption) []map[string]interface{} {
	return Maps(any, option...)
}

// SliceMapDeep 是 MapsDeep 的别名。
// 不推荐使用：请改用 SliceMap。
// md5:e577a1254364096d
func SliceMapDeep(any interface{}) []map[string]interface{} {
	return MapsDeep(any)
}

// Maps 将 `value` 转换为 []map[string]interface{} 类型。
// 注意，如果 `value` 是字符串或[]byte类型，它会自动检查并转换为json字符串到 []map。
// md5:cff297515bb58eac
func Maps(value interface{}, option ...MapOption) []map[string]interface{} {
	if value == nil {
		return nil
	}
	switch r := value.(type) {
	case string:
		list := make([]map[string]interface{}, 0)
		if len(r) > 0 && r[0] == '[' && r[len(r)-1] == ']' {
			if err := json.UnmarshalUseNumber([]byte(r), &list); err != nil {
				return nil
			}
			return list
		} else {
			return nil
		}

	case []byte:
		list := make([]map[string]interface{}, 0)
		if len(r) > 0 && r[0] == '[' && r[len(r)-1] == ']' {
			if err := json.UnmarshalUseNumber(r, &list); err != nil {
				return nil
			}
			return list
		} else {
			return nil
		}

	case []map[string]interface{}:
		return r

	default:
		array := Interfaces(value)
		if len(array) == 0 {
			return nil
		}
		list := make([]map[string]interface{}, len(array))
		for k, v := range array {
			list[k] = Map(v, option...)
		}
		return list
	}
}

// MapsDeep 将 `value` 递归地转换为 []map[string]interface{} 类型。
//
// TODO：完全实现所有类型的递归转换。
// 已弃用：推荐使用 Maps 替代。
// md5:4fca16e80380d63d
func MapsDeep(value interface{}, tags ...string) []map[string]interface{} {
	if value == nil {
		return nil
	}
	switch r := value.(type) {
	case string:
		list := make([]map[string]interface{}, 0)
		if len(r) > 0 && r[0] == '[' && r[len(r)-1] == ']' {
			if err := json.UnmarshalUseNumber([]byte(r), &list); err != nil {
				return nil
			}
			return list
		} else {
			return nil
		}

	case []byte:
		list := make([]map[string]interface{}, 0)
		if len(r) > 0 && r[0] == '[' && r[len(r)-1] == ']' {
			if err := json.UnmarshalUseNumber(r, &list); err != nil {
				return nil
			}
			return list
		} else {
			return nil
		}

	case []map[string]interface{}:
		list := make([]map[string]interface{}, len(r))
		for k, v := range r {
			list[k] = MapDeep(v, tags...)
		}
		return list

	default:
		array := Interfaces(value)
		if len(array) == 0 {
			return nil
		}
		list := make([]map[string]interface{}, len(array))
		for k, v := range array {
			list[k] = MapDeep(v, tags...)
		}
		return list
	}
}
