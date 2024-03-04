// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtype

import (
	"math"
	"strconv"
	"sync/atomic"
	
	"github.com/gogf/gf/v2/util/gconv"
)

// Float64 是一个结构体，用于对 float64 类型进行并发安全操作。
type Float64 struct {
	value uint64
}

// NewFloat64 创建并返回一个用于 float64 类型的并发安全对象，
// 其初始值为给定的 `value`。
func NewFloat64(value ...float64) *Float64 {
	if len(value) > 0 {
		return &Float64{
			value: math.Float64bits(value[0]),
		}
	}
	return &Float64{}
}

// Clone 克隆并返回一个用于 float64 类型的新并发安全对象。
func (v *Float64) Clone() *Float64 {
	return NewFloat64(v.Val())
}

// Set 方法通过原子操作将`value`存储到t.value中，并返回修改前的t.value的值。
func (v *Float64) Set(value float64) (old float64) {
	return math.Float64frombits(atomic.SwapUint64(&v.value, math.Float64bits(value)))
}

// Val 原子性地加载并返回 t.value。
func (v *Float64) Val() float64 {
	return math.Float64frombits(atomic.LoadUint64(&v.value))
}

// Add 原子性地将 `delta` 加到 t.value 上，并返回新的值。
func (v *Float64) Add(delta float64) (new float64) {
	for {
		old := math.Float64frombits(v.value)
		new = old + delta
		if atomic.CompareAndSwapUint64(
			&v.value,
			math.Float64bits(old),
			math.Float64bits(new),
		) {
			break
		}
	}
	return
}

// Cas 执行值的比较并交换操作。
func (v *Float64) Cas(old, new float64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&v.value, math.Float64bits(old), math.Float64bits(new))
}

// String 实现了 String 接口以便进行字符串打印。
func (v *Float64) String() string {
	return strconv.FormatFloat(v.Val(), 'g', -1, 64)
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Float64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(v.Val(), 'g', -1, 64)), nil
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Float64) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Float64(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 `v` 设置任意类型的值。
func (v *Float64) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Float64(value))
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Float64) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewFloat64(v.Val())
}
