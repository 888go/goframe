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

// SliceMap别名 是 Maps 的别名。 md5:af436f87335633ac
func SliceMap别名(值 interface{}, option ...MapOption) []map[string]interface{} {
	return X取Map切片(值, option...)
}

// SliceMapDeep别名 是 MapsDeep 的别名。
// 不推荐使用：请改用 SliceMap。
// md5:e577a1254364096d
func SliceMapDeep别名(any interface{}) []map[string]interface{} {
	return X取Map切片_递归(any)
}

// X取Map切片 将 `value` 转换为 []map[string]interface{} 类型。
// 注意，如果 `value` 是字符串或[]byte类型，它会自动检查并转换为json字符串到 []map。
// md5:cff297515bb58eac
func X取Map切片(值 interface{}, 选项 ...MapOption) []map[string]interface{} {
	if 值 == nil {
		return nil
	}
	switch r := 值.(type) {
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
		array := X取any切片(值)
		if len(array) == 0 {
			return nil
		}
		list := make([]map[string]interface{}, len(array))
		for k, v := range array {
			list[k] = X取Map(v, 选项...)
		}
		return list
	}
}

// X取Map切片_递归 将 `value` 递归地转换为 []map[string]interface{} 类型。
//
// TODO：完全实现所有类型的递归转换。
// 已弃用：推荐使用 Maps 替代。
// md5:4fca16e80380d63d
func X取Map切片_递归(值 interface{}, 值标签 ...string) []map[string]interface{} {
	if 值 == nil {
		return nil
	}
	switch r := 值.(type) {
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
			list[k] = X取Map_递归(v, 值标签...)
		}
		return list

	default:
		array := X取any切片(值)
		if len(array) == 0 {
			return nil
		}
		list := make([]map[string]interface{}, len(array))
		for k, v := range array {
			list[k] = X取Map_递归(v, 值标签...)
		}
		return list
	}
}
