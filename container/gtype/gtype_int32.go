// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gtype

import (
	"strconv"
	"sync/atomic"

	"github.com/gogf/gf/v2/util/gconv"
)

// Int32是一个用于int32类型并发安全操作的结构体。 md5:54bb3e06eb1184ea
type Int32 struct {
	value int32
}

// NewInt32 创建并返回一个针对int32类型的并发安全对象，
// 初始值为给定的`value`。 md5:c3a543fa77ce812c
func NewInt32(value ...int32) *Int32 {
	if len(value) > 0 {
		return &Int32{
			value: value[0],
		}
	}
	return &Int32{}
}

// Clone 克隆并返回一个新的针对int32类型的并发安全对象。 md5:c474dc7363567e12
func (v *Int32) Clone() *Int32 {
	return NewInt32(v.Val())
}

// Set 原子地将 `value` 存储到 t.value 中，并返回 t.value 的旧值。 md5:2ce98b05d0290b37
func (v *Int32) Set(value int32) (old int32) {
	return atomic.SwapInt32(&v.value, value)
}

// Val原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *Int32) Val() int32 {
	return atomic.LoadInt32(&v.value)
}

// Atomically 将 `delta` 增加到 t.value 中，并返回新的值。 md5:73547274aea5fe91
func (v *Int32) Add(delta int32) (new int32) {
	return atomic.AddInt32(&v.value, delta)
}

// Cas 执行针对值的比较并交换操作。 md5:4c2d06b4167bee48
func (v *Int32) Cas(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&v.value, old, new)
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *Int32) String() string {
	return strconv.Itoa(int(v.Val()))
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Int32) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(v.Val()))), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Int32) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Int32(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *Int32) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Int32(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Int32) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewInt32(v.Val())
}
