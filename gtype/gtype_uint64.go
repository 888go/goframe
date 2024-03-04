// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtype

import (
	"strconv"
	"sync/atomic"
	
	"github.com/gogf/gf/v2/util/gconv"
)

// Uint64 是一个结构体，用于对 uint64 类型进行并发安全操作。
type Uint64 struct {
	value uint64
}

// NewUint64 创建并返回一个用于 uint64 类型的并发安全对象，
// 其初始值为给定的 `value`。
func NewUint64(value ...uint64) *Uint64 {
	if len(value) > 0 {
		return &Uint64{
			value: value[0],
		}
	}
	return &Uint64{}
}

// Clone 克隆并返回一个用于 uint64 类型的新并发安全对象。
func (v *Uint64) Clone() *Uint64 {
	return NewUint64(v.Val())
}

// Set 方法通过原子操作将`value`存储到t.value中，并返回修改前的t.value的值。
func (v *Uint64) Set(value uint64) (old uint64) {
	return atomic.SwapUint64(&v.value, value)
}

// Val 原子性地加载并返回 t.value。
func (v *Uint64) Val() uint64 {
	return atomic.LoadUint64(&v.value)
}

// Add 原子性地将 `delta` 加到 t.value 上，并返回新的值。
func (v *Uint64) Add(delta uint64) (new uint64) {
	return atomic.AddUint64(&v.value, delta)
}

// Cas 执行值的比较并交换操作。
func (v *Uint64) Cas(old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&v.value, old, new)
}

// String 实现了 String 接口以便进行字符串打印。
func (v *Uint64) String() string {
	return strconv.FormatUint(v.Val(), 10)
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Uint64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(v.Val(), 10)), nil
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Uint64) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Uint64(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 `v` 设置任意类型的值。
func (v *Uint64) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Uint64(value))
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Uint64) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewUint64(v.Val())
}
