// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtype

import (
	"sync/atomic"
	
	"github.com/888go/goframe/internal/deepcopy"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/util/gconv"
)

// Interface 是一个结构体，用于对 interface{} 类型进行并发安全操作。
type Interface struct {
	value atomic.Value
}

// NewInterface 创建并返回一个对 interface{} 类型安全的并发对象，
// 其初始值为给定的 `value`。
func NewInterface(value ...interface{}) *Interface {
	t := &Interface{}
	if len(value) > 0 && value[0] != nil {
		t.value.Store(value[0])
	}
	return t
}

// Clone 克隆并返回一个用于 interface{} 类型的新并发安全对象。
func (v *Interface) Clone() *Interface {
	return NewInterface(v.Val())
}

// Set 方法通过原子操作将`value`存储到t.value，并返回修改前的t.value的值。
// 注意：参数`value`不能为空。
func (v *Interface) Set(value interface{}) (old interface{}) {
	old = v.Val()
	v.value.Store(value)
	return
}

// Val 原子性地加载并返回 t.value。
func (v *Interface) Val() interface{} {
	return v.value.Load()
}

// String 实现了 String 接口以便进行字符串打印。
func (v *Interface) String() string {
	return gconv.String(v.Val())
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Interface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Val())
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Interface) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.UnmarshalUseNumber(b, &i); err != nil {
		return err
	}
	v.Set(i)
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 `v` 设置任意类型的值。
func (v *Interface) UnmarshalValue(value interface{}) error {
	v.Set(value)
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Interface) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewInterface(deepcopy.Copy(v.Val()))
}
