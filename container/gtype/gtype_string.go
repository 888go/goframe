// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtype

import (
	"bytes"
	"sync/atomic"
	
	"github.com/888go/goframe/util/gconv"
)

// String 是一个结构体，用于对字符串类型进行并发安全操作。
type String struct {
	value atomic.Value
}

// NewString 创建并返回一个用于字符串类型的安全并发对象，
// 其初始值为给定的 `value`。
func NewString(value ...string) *String {
	t := &String{}
	if len(value) > 0 {
		t.value.Store(value[0])
	}
	return t
}

// Clone 克隆并返回一个用于字符串类型的新并发安全对象。
func (v *String) Clone() *String {
	return NewString(v.Val())
}

// Set 方法通过原子操作将`value`存储到t.value中，并返回修改前的t.value的值。
func (v *String) Set(value string) (old string) {
	old = v.Val()
	v.value.Store(value)
	return
}

// Val 原子性地加载并返回 t.value。
func (v *String) Val() string {
	s := v.value.Load()
	if s != nil {
		return s.(string)
	}
	return ""
}

// String 实现了 String 接口以便进行字符串打印。
func (v *String) String() string {
	return v.Val()
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v String) MarshalJSON() ([]byte, error) {
	return []byte(`"` + v.Val() + `"`), nil
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *String) UnmarshalJSON(b []byte) error {
	v.Set(string(bytes.Trim(b, `"`)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 `v` 设置任意类型的值。
func (v *String) UnmarshalValue(value interface{}) error {
	v.Set(gconv.String(value))
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *String) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewString(v.Val())
}
