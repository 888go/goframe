// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtype

import (
	"bytes"
	"sync/atomic"

	"github.com/gogf/gf/v2/util/gconv"
)

// String 是一个用于类型 string 的并发安全操作的结构体。 md5:33de4be4fa63f063
type String struct {
	value atomic.Value
}

// NewString 创建并返回一个针对字符串类型的并发安全对象，
// 初始化值为给定的 `value`。
// md5:3e768b94fd16a8d0
func NewString(value ...string) *String {
	t := &String{}
	if len(value) > 0 {
		t.value.Store(value[0])
	}
	return t
}

// Clone 克隆并返回一个新的用于字符串类型的并发安全对象。 md5:1f8299657e3ed3d3
func (v *String) Clone() *String {
	return NewString(v.Val())
}

// Set 原子地将 `value` 存储到 t.value 中，并返回 t.value 的旧值。 md5:2ce98b05d0290b37
func (v *String) Set(value string) (old string) {
	old = v.Val()
	v.value.Store(value)
	return
}

// Val原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *String) Val() string {
	s := v.value.Load()
	if s != nil {
		return s.(string)
	}
	return ""
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *String) String() string {
	return v.Val()
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v String) MarshalJSON() ([]byte, error) {
	return []byte(`"` + v.Val() + `"`), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *String) UnmarshalJSON(b []byte) error {
	v.Set(string(bytes.Trim(b, `"`)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *String) UnmarshalValue(value interface{}) error {
	v.Set(gconv.String(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *String) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewString(v.Val())
}
