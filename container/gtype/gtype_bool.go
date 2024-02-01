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
// Bool 是一个结构体，用于对 bool 类型进行并发安全操作。
type Bool struct {
	value int32
}

var (
	bytesTrue  = []byte("true")
	bytesFalse = []byte("false")
)

// NewBool 创建并返回一个用于 bool 类型的并发安全对象，
// 初始值为给定的 `value`。
func NewBool(value ...bool) *Bool {
	t := &Bool{}
	if len(value) > 0 {
		if value[0] {
			t.value = 1
		} else {
			t.value = 0
		}
	}
	return t
}

// Clone 克隆并返回一个用于布尔类型的新的线程安全对象。
func (v *Bool) Clone() *Bool {
	return NewBool(v.Val())
}

// Set 方法通过原子操作将`value`存储到t.value中，并返回修改前的t.value的值。
func (v *Bool) Set(value bool) (old bool) {
	if value {
		old = atomic.SwapInt32(&v.value, 1) == 1
	} else {
		old = atomic.SwapInt32(&v.value, 0) == 1
	}
	return
}

// Val 原子性地加载并返回 t.value。
func (v *Bool) Val() bool {
	return atomic.LoadInt32(&v.value) > 0
}

// Cas 执行值的比较并交换操作。
func (v *Bool) Cas(old, new bool) (swapped bool) {
	var oldInt32, newInt32 int32
	if old {
		oldInt32 = 1
	}
	if new {
		newInt32 = 1
	}
	return atomic.CompareAndSwapInt32(&v.value, oldInt32, newInt32)
}

// String 实现了 String 接口以便进行字符串打印。
func (v *Bool) String() string {
	if v.Val() {
		return "true"
	}
	return "false"
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Bool) MarshalJSON() ([]byte, error) {
	if v.Val() {
		return bytesTrue, nil
	}
	return bytesFalse, nil
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Bool) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Bool(bytes.Trim(b, `"`)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 `v` 设置任意类型的值。
func (v *Bool) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Bool(value))
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Bool) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewBool(v.Val())
}
