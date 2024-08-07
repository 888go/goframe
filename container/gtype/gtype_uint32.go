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

// Uint32是一个用于类型uint32并发安全操作的结构体。 md5:916b918898b4b0fa
type Uint32 struct {
	value uint32
}

// NewUint32 创建并返回一个针对 uint32 类型的并发安全对象，初始值为 `value`。
// md5:5cc7d55fe951a549
func NewUint32(value ...uint32) *Uint32 {
	if len(value) > 0 {
		return &Uint32{
			value: value[0],
		}
	}
	return &Uint32{}
}

// Clone 克隆并返回一个新的对于 uint32 类型的并发安全对象。 md5:3dc7263b57b51dd5
func (v *Uint32) Clone() *Uint32 {
	return NewUint32(v.X取值())
}

// X设置值 原子地将 `value` 存储到 t.value 中，并返回 t.value 的旧值。 md5:2ce98b05d0290b37
func (v *Uint32) X设置值(value uint32) (old uint32) {
	return atomic.SwapUint32(&v.value, value)
}

// X取值原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *Uint32) X取值() uint32 {
	return atomic.LoadUint32(&v.value)
}

// Atomically 将 `delta` 增加到 t.value 中，并返回新的值。 md5:73547274aea5fe91
func (v *Uint32) Add(delta uint32) (new uint32) {
	return atomic.AddUint32(&v.value, delta)
}

// Cas 执行针对值的比较并交换操作。 md5:4c2d06b4167bee48
func (v *Uint32) Cas(old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&v.value, old, new)
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *Uint32) String() string {
	return strconv.FormatUint(uint64(v.X取值()), 10)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Uint32) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(v.X取值()), 10)), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Uint32) UnmarshalJSON(b []byte) error {
	v.X设置值(gconv.X取正整数32位(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *Uint32) UnmarshalValue(value interface{}) error {
	v.X设置值(gconv.X取正整数32位(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Uint32) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewUint32(v.X取值())
}
