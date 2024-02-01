// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtype
import (
	"math"
	"strconv"
	"sync/atomic"
	
	"github.com/888go/goframe/util/gconv"
	)
// Float32 是一个结构体，用于对 float32 类型进行并发安全操作。
type Float32 struct {
	value uint32
}

// NewFloat32 创建并返回一个用于 float32 类型的并发安全对象，
// 其初始值为给定的 `value`。
func NewFloat32(value ...float32) *Float32 {
	if len(value) > 0 {
		return &Float32{
			value: math.Float32bits(value[0]),
		}
	}
	return &Float32{}
}

// Clone 克隆并返回一个用于 float32 类型的新并发安全对象。
func (v *Float32) Clone() *Float32 {
	return NewFloat32(v.Val())
}

// Set 方法通过原子操作将`value`存储到t.value中，并返回修改前的t.value的值。
func (v *Float32) Set(value float32) (old float32) {
	return math.Float32frombits(atomic.SwapUint32(&v.value, math.Float32bits(value)))
}

// Val 原子性地加载并返回 t.value。
func (v *Float32) Val() float32 {
	return math.Float32frombits(atomic.LoadUint32(&v.value))
}

// Add 原子性地将 `delta` 加到 t.value 上，并返回新的值。
func (v *Float32) Add(delta float32) (new float32) {
	for {
		old := math.Float32frombits(v.value)
		new = old + delta
		if atomic.CompareAndSwapUint32(
			&v.value,
			math.Float32bits(old),
			math.Float32bits(new),
		) {
			break
		}
	}
	return
}

// Cas 执行值的比较并交换操作。
func (v *Float32) Cas(old, new float32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&v.value, math.Float32bits(old), math.Float32bits(new))
}

// String 实现了 String 接口以便进行字符串打印。
func (v *Float32) String() string {
	return strconv.FormatFloat(float64(v.Val()), 'g', -1, 32)
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Float32) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(v.Val()), 'g', -1, 32)), nil
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Float32) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Float32(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 `v` 设置任意类型的值。
func (v *Float32) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Float32(value))
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Float32) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewFloat32(v.Val())
}
