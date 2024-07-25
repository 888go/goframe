// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtype

import (
	"strconv"
	"sync/atomic"

	"github.com/gogf/gf/v2/util/gconv"
)

// Uint64是一个结构体，用于并发安全的uint64类型操作。 md5:6a07488d07d4b044
type Uint64 struct {
	value uint64
}

// NewUint64 创建并返回一个针对 uint64 类型的并发安全对象，初始值为 `value`。
// md5:6b0c10ffdeecf7b1
func NewUint64(value ...uint64) *Uint64 {
	if len(value) > 0 {
		return &Uint64{
			value: value[0],
		}
	}
	return &Uint64{}
}

// Clone 为 uint64 类型创建并返回一个新的并发安全的对象。 md5:5fdab60d860cae3b
func (v *Uint64) Clone() *Uint64 {
	return NewUint64(v.Val())
}

// Set 原子地将 `value` 存储到 t.value 中，并返回 t.value 的旧值。 md5:2ce98b05d0290b37
func (v *Uint64) Set(value uint64) (old uint64) {
	return atomic.SwapUint64(&v.value, value)
}

// Val原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *Uint64) Val() uint64 {
	return atomic.LoadUint64(&v.value)
}

// Atomically 将 `delta` 增加到 t.value 中，并返回新的值。 md5:73547274aea5fe91
func (v *Uint64) Add(delta uint64) (new uint64) {
	return atomic.AddUint64(&v.value, delta)
}

// Cas 执行针对值的比较并交换操作。 md5:4c2d06b4167bee48
func (v *Uint64) Cas(old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&v.value, old, new)
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *Uint64) String() string {
	return strconv.FormatUint(v.Val(), 10)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Uint64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(v.Val(), 10)), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Uint64) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Uint64(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *Uint64) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Uint64(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Uint64) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewUint64(v.Val())
}
