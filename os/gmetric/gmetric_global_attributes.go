// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

import (
	"sync"

	"github.com/gogf/gf/v2/text/gregex"
)

// SetGlobalAttributesOption 将全局属性绑定到特定的仪表。 md5:5cba96bea01d2134
type SetGlobalAttributesOption struct {
	// Instrument 指定乐器名称。 md5:9d3f75d7c5acf5ce
	Instrument string

	// Instrument 指定仪器版本。 md5:08f1ad86326ce5c0
	InstrumentVersion string

// InstrumentPattern 通过正则表达式指定要操作的Instrument名称。
// 示例：
// 1. 如果设置为`.+`，将匹配所有Instrument。
// 2. 如果设置为`github.com/gogf/gf.+`，将匹配所有goframe相关的Instrument。
// md5:a3225129ad31cbb0
	InstrumentPattern string
}

// GetGlobalAttributesOption 将全局属性绑定到特定的仪表。 md5:b534095d7e4c28c6
type GetGlobalAttributesOption struct {
	Instrument        string // Instrument 指定乐器名称。 md5:9d3f75d7c5acf5ce
	InstrumentVersion string // Instrument 指定仪器版本。 md5:08f1ad86326ce5c0
}

type globalAttributeItem struct {
	Attributes
	SetGlobalAttributesOption
}

var (
	globalAttributesMu sync.Mutex
	// globalAttributes 将全局属性存储到一个映射中。 md5:e8b73fe60d039913
	globalAttributes = make([]globalAttributeItem, 0)
)

// SetGlobalAttributes 根据 `SetGlobalAttributesOption` 添加全局属性。如果给定的 `SetGlobalAttributesOption` 为空，它将向所有指标添加全局属性。如果提供了特定的 `SetGlobalAttributesOption`，它将向指定的度量添加全局属性。
// md5:5ba03a1e3d761b95
// ff:
// attrs:
// option:
func SetGlobalAttributes(attrs Attributes, option SetGlobalAttributesOption) {
	globalAttributesMu.Lock()
	defer globalAttributesMu.Unlock()
	globalAttributes = append(
		globalAttributes, globalAttributeItem{
			Attributes:                attrs,
			SetGlobalAttributesOption: option,
		},
	)
}

	// GetGlobalAttributes 通过 `GetGlobalAttributesOption` 获取并返回全局属性。
	// 如果给定的 `GetGlobalAttributesOption` 为空，它将返回所有全局属性。
	// 如果 `GetGlobalAttributesOption` 不为空，它将返回特定仪器的全局属性。
	// md5:8327524dc9d44419
// ff:
// option:
func GetGlobalAttributes(option GetGlobalAttributesOption) Attributes {
	globalAttributesMu.Lock()
	defer globalAttributesMu.Unlock()
	var attributes = make(Attributes, 0)
	for _, attrItem := range globalAttributes {
		// instrument name.
		if attrItem.InstrumentPattern != "" {
			if !gregex.IsMatchString(attrItem.InstrumentPattern, option.Instrument) {
				continue
			}
		} else {
			if (attrItem.Instrument != "" || option.Instrument != "") &&
				attrItem.Instrument != option.Instrument {
				continue
			}
		}
		// instrument version.
		if (attrItem.InstrumentVersion != "" || option.InstrumentVersion != "") &&
			attrItem.InstrumentVersion != option.InstrumentVersion {
			continue
		}
		attributes = append(attributes, attrItem.Attributes...)
	}
	return attributes
}
