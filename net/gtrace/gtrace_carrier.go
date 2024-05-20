// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtrace

import (
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/util/gconv"
)

// Carrier是TextMapPropagator使用的存储介质。. md5:8d62ccc64614a0c7
type Carrier map[string]interface{}

// NewCarrier 创建并返回一个 Carrier。. md5:e8b512c41d0c9bd7
func NewCarrier(data ...map[string]interface{}) Carrier {
	if len(data) > 0 && data[0] != nil {
		return data[0]
	}
	return make(map[string]interface{})
}

// Get 返回与传入键关联的值。. md5:d20be4af58bc2fa7
func (c Carrier) Get(k string) string {
	return gconv.String(c[k])
}

// Set 存储键值对。. md5:797de4e363035487
func (c Carrier) Set(k, v string) {
	c[k] = v
}

// Keys 列出了存储在该载体中的键。. md5:8dca78c8668c962f
func (c Carrier) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}

// MustMarshal 返回c的JSON编码. md5:58f2f2d8e0370a56
func (c Carrier) MustMarshal() []byte {
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return b
}

// String 将当前Carrier转换并以字符串形式返回。. md5:9a30611e4d817539
func (c Carrier) String() string {
	return string(c.MustMarshal())
}

// UnmarshalJSON 实现了 json 包中的 UnmarshalJSON 接口。. md5:6a50aca077e136ff
func (c Carrier) UnmarshalJSON(b []byte) error {
	carrier := NewCarrier(nil)
	return json.UnmarshalUseNumber(b, carrier)
}
