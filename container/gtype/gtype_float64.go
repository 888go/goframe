// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtype

import (
	"math"
	"strconv"
	"sync/atomic"

	"github.com/gogf/gf/v2/util/gconv"
)

// Float64 是一个用于并发安全操作 float64 类型的结构体。 md5:9134cad59fd8776a
type Float64 struct {
	value uint64
}

// NewFloat64 创建并返回一个针对 float64 类型的并发安全对象，初始值为 `value`。
// md5:ef46288a8eea7230
func NewFloat64(value ...float64) *Float64 {
	if len(value) > 0 {
		return &Float64{
			value: math.Float64bits(value[0]),
		}
	}
	return &Float64{}
}

// Clone 克隆并返回一个新的针对 float64 类型的并发安全对象。 md5:5e22df1240b19bf5
func (v *Float64) Clone() *Float64 {
	return NewFloat64(v.Val())
}

// Set 原子地将 `value` 存储到 t.value 中，并返回 t.value 的旧值。 md5:2ce98b05d0290b37
func (v *Float64) Set(value float64) (old float64) {
	return math.Float64frombits(atomic.SwapUint64(&v.value, math.Float64bits(value)))
}

// Val原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *Float64) Val() float64 {
	return math.Float64frombits(atomic.LoadUint64(&v.value))
}

// Atomically 将 `delta` 增加到 t.value 中，并返回新的值。 md5:73547274aea5fe91
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

// Cas 执行针对值的比较并交换操作。 md5:4c2d06b4167bee48
func (v *Float64) Cas(old, new float64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&v.value, math.Float64bits(old), math.Float64bits(new))
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *Float64) String() string {
	return strconv.FormatFloat(v.Val(), 'g', -1, 64)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Float64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(v.Val(), 'g', -1, 64)), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Float64) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Float64(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *Float64) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Float64(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Float64) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewFloat64(v.Val())
}
