// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtype
import (
	"strconv"
	"sync/atomic"
	
	"github.com/888go/goframe/util/gconv"
	)
// Int32 是一个结构体，用于对 int32 类型进行并发安全操作。
type Int32 struct {
	value int32
}

// NewInt32 创建并返回一个适用于 int32 类型的并发安全对象，
// 其初始值为给定的 `value`。
func NewInt32(value ...int32) *Int32 {
	if len(value) > 0 {
		return &Int32{
			value: value[0],
		}
	}
	return &Int32{}
}

// Clone 克隆并返回一个新的适用于 int32 类型的并发安全对象。
func (v *Int32) Clone() *Int32 {
	return NewInt32(v.Val())
}

// Set 方法通过原子操作将`value`存储到t.value中，并返回修改前的t.value的值。
func (v *Int32) Set(value int32) (old int32) {
	return atomic.SwapInt32(&v.value, value)
}

// Val 原子性地加载并返回 t.value。
func (v *Int32) Val() int32 {
	return atomic.LoadInt32(&v.value)
}

// Add 原子性地将 `delta` 加到 t.value 上，并返回新的值。
func (v *Int32) Add(delta int32) (new int32) {
	return atomic.AddInt32(&v.value, delta)
}

// Cas 执行值的比较并交换操作。
func (v *Int32) Cas(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&v.value, old, new)
}

// String 实现了 String 接口以便进行字符串打印。
func (v *Int32) String() string {
	return strconv.Itoa(int(v.Val()))
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Int32) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(v.Val()))), nil
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Int32) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Int32(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 `v` 设置任意类型的值。
func (v *Int32) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Int32(value))
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Int32) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewInt32(v.Val())
}
