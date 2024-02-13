// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类

import (
	"github.com/888go/goframe/internal/json"
)

// SliceMap 是 Maps 的别名。
func SliceMap别名(值 interface{}, option ...MapOption) []map[string]interface{} {
	return X取Map数组(值, option...)
}

// SliceMapDeep 是 MapsDeep 的别名。
// 已弃用：请改用 SliceMap。
func SliceMapDeep别名(any interface{}) []map[string]interface{} {
	return X取Map数组_递归(any)
}

// Maps 将 `value` 转换为 []map[string]interface{} 类型。
// 注意，如果 `value` 为 string 或 []byte 类型，它会自动检查并转换为 json 字符串到 []map。
func X取Map数组(值 interface{}, 选项 ...MapOption) []map[string]interface{} {
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
		array := X取any数组(值)
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

// MapsDeep 递归地将`value`转换为[]map[string]interface{}。
//
// TODO 完全实现对所有类型的递归转换。
// 废弃: 请改用Maps方法。
func X取Map数组_递归(值 interface{}, 值标签 ...string) []map[string]interface{} {
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
		array := X取any数组(值)
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
