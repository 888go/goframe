// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtrace

import (
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/util/gconv"
)

// 2024-01-14 备注,此文件方法不要翻译,  
// gtrace_z_unit_carrier_test.go, "otel.GetTextMapPropagator().Inject(ctx, carrier1)"
// carrier1参数要求TextMapCarrier类型结构体.
//
// Carrier 是 TextMapPropagator 使用的存储介质。
type Carrier map[string]interface{}

// NewCarrier 创建并返回一个 Carrier。
func NewCarrier(data ...map[string]interface{}) Carrier {
	if len(data) > 0 && data[0] != nil {
		return data[0]
	}
	return make(map[string]interface{})
}

// 2024-01-14 备注,此文件方法不要翻译,  
// gtrace_z_unit_carrier_test.go, "otel.GetTextMapPropagator().Inject(ctx, carrier1)"
// carrier1参数要求TextMapCarrier类型结构体.
//
// Get 方法用于获取与传递的键关联的值。
func (c Carrier) Get(k string) string {
	return gconv.String(c[k])
}

// 2024-01-14 备注,此文件方法不要翻译,  
// gtrace_z_unit_carrier_test.go, "otel.GetTextMapPropagator().Inject(ctx, carrier1)"
// carrier1参数要求TextMapCarrier类型结构体.
//
// Set 存储键值对。
func (c Carrier) Set(k, v string) {
	c[k] = v
}

// Keys 列出存储在此 carrier 中的所有键。
func (c Carrier) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}

// MustMarshal .必须返回c的JSON编码
func (c Carrier) MustMarshal() []byte {
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return b
}

// String 将当前Carrier转换并以字符串形式返回。
func (c Carrier) String() string {
	return string(c.MustMarshal())
}

// UnmarshalJSON 实现了 json 包中的 UnmarshalJSON 接口。
func (c Carrier) UnmarshalJSON(b []byte) error {
	carrier := NewCarrier(nil)
	return json.UnmarshalUseNumber(b, carrier)
}
