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
// Byte 是一个结构体，用于对 byte 类型进行并发安全操作。
type Byte struct {
	value int32
}

// NewByte创建并返回一个用于byte类型的并发安全对象，
// 其初始值为给定的`value`。
func NewByte(value ...byte) *Byte {
	if len(value) > 0 {
		return &Byte{
			value: int32(value[0]),
		}
	}
	return &Byte{}
}

// Clone 克隆并返回一个新的适用于 byte 类型的并发安全对象副本。
func (v *Byte) Clone() *Byte {
	return NewByte(v.Val())
}

// Set 方法通过原子操作将`value`存储到t.value中，并返回修改前的t.value的值。
func (v *Byte) Set(value byte) (old byte) {
	return byte(atomic.SwapInt32(&v.value, int32(value)))
}

// Val 原子性地加载并返回 t.value。
func (v *Byte) Val() byte {
	return byte(atomic.LoadInt32(&v.value))
}

// Add 原子性地将 `delta` 加到 t.value 上，并返回新的值。
func (v *Byte) Add(delta byte) (new byte) {
	return byte(atomic.AddInt32(&v.value, int32(delta)))
}

// Cas 执行值的比较并交换操作。
func (v *Byte) Cas(old, new byte) (swapped bool) {
	return atomic.CompareAndSwapInt32(&v.value, int32(old), int32(new))
}

// String 实现了 String 接口以便进行字符串打印。
func (v *Byte) String() string {
	return strconv.FormatUint(uint64(v.Val()), 10)
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Byte) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(v.Val()), 10)), nil
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Byte) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Uint8(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 `v` 设置任意类型的值。
func (v *Byte) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Byte(value))
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Byte) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewByte(v.Val())
}
