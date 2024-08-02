// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 安全变量类

import (
	"math"
	"strconv"
	"sync/atomic"

	gconv "github.com/888go/goframe/util/gconv"
)

// Float32是一个用于并发安全操作float32类型的结构体。 md5:df0c1aaf5f1f5645
type Float32 struct {
	value uint32
}

// NewFloat32 创建并返回一个针对float32类型的并发安全对象，
// 初始值为给定的`value`。
// md5:a2e96663c9c91d0d
func NewFloat32(value ...float32) *Float32 {
	if len(value) > 0 {
		return &Float32{
			value: math.Float32bits(value[0]),
		}
	}
	return &Float32{}
}

// Clone 创建并返回一个新的并发安全的float32类型对象。 md5:5848ca2b0b7eef06
func (v *Float32) Clone() *Float32 {
	return NewFloat32(v.Val())
}

// Set 原子地将 `value` 存储到 t.value 中，并返回 t.value 的旧值。 md5:2ce98b05d0290b37
func (v *Float32) Set(value float32) (old float32) {
	return math.Float32frombits(atomic.SwapUint32(&v.value, math.Float32bits(value)))
}

// Val原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *Float32) Val() float32 {
	return math.Float32frombits(atomic.LoadUint32(&v.value))
}

// Atomically 将 `delta` 增加到 t.value 中，并返回新的值。 md5:73547274aea5fe91
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

// Cas 执行针对值的比较并交换操作。 md5:4c2d06b4167bee48
func (v *Float32) Cas(old, new float32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&v.value, math.Float32bits(old), math.Float32bits(new))
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *Float32) String() string {
	return strconv.FormatFloat(float64(v.Val()), 'g', -1, 32)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Float32) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(v.Val()), 'g', -1, 32)), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Float32) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Float32(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *Float32) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Float32(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Float32) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewFloat32(v.Val())
}
