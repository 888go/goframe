// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 安全变量类

import (
	"strconv"
	"sync/atomic"

	gconv "github.com/888go/goframe/util/gconv"
)

// Uint 是一个用于类型 uint 的并发安全操作的结构体。 md5:3b5bf4b2533c3310
type Uint struct {
	value uint64
}

// NewUint 创建并返回一个并发安全的 uint 类型对象，初始值为 `value`。
// md5:7e99bd1d0ac4986d
func NewUint(value ...uint) *Uint {
	if len(value) > 0 {
		return &Uint{
			value: uint64(value[0]),
		}
	}
	return &Uint{}
}

// Clone 为 uint 类型创建并返回一个新的并发安全的对象。 md5:d17e8edf52e037bb
func (v *Uint) Clone() *Uint {
	return NewUint(v.X取值())
}

// X设置值 原子地将 `value` 存储到 t.value 中，并返回 t.value 的旧值。 md5:2ce98b05d0290b37
func (v *Uint) X设置值(value uint) (old uint) {
	return uint(atomic.SwapUint64(&v.value, uint64(value)))
}

// X取值原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *Uint) X取值() uint {
	return uint(atomic.LoadUint64(&v.value))
}

// Atomically 将 `delta` 增加到 t.value 中，并返回新的值。 md5:73547274aea5fe91
func (v *Uint) Add(delta uint) (new uint) {
	return uint(atomic.AddUint64(&v.value, uint64(delta)))
}

// Cas 执行针对值的比较并交换操作。 md5:4c2d06b4167bee48
func (v *Uint) Cas(old, new uint) (swapped bool) {
	return atomic.CompareAndSwapUint64(&v.value, uint64(old), uint64(new))
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *Uint) String() string {
	return strconv.FormatUint(uint64(v.X取值()), 10)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Uint) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(v.X取值()), 10)), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Uint) UnmarshalJSON(b []byte) error {
	v.X设置值(gconv.X取正整数(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *Uint) UnmarshalValue(value interface{}) error {
	v.X设置值(gconv.X取正整数(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Uint) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewUint(v.X取值())
}
