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

// Int64 是一个用于并发安全操作整型 int64 的结构体。 md5:563e61ea81d21a66
type Int64 struct {
	value int64
}

// NewInt64 创建并返回一个针对 int64 类型的并发安全对象，初始值为 `value`。
// md5:5685b29efcb68e8a
func NewInt64(value ...int64) *Int64 {
	if len(value) > 0 {
		return &Int64{
			value: value[0],
		}
	}
	return &Int64{}
}

// Clone 克隆并返回一个新的对于int64类型而言线程安全的对象。 md5:2e6afaa5b059c165
func (v *Int64) Clone() *Int64 {
	return NewInt64(v.Val())
}

// Set 原子地将 `value` 存储到 t.value 中，并返回 t.value 的旧值。 md5:2ce98b05d0290b37
func (v *Int64) Set(value int64) (old int64) {
	return atomic.SwapInt64(&v.value, value)
}

// Val原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *Int64) Val() int64 {
	return atomic.LoadInt64(&v.value)
}

// Atomically 将 `delta` 增加到 t.value 中，并返回新的值。 md5:73547274aea5fe91
func (v *Int64) Add(delta int64) (new int64) {
	return atomic.AddInt64(&v.value, delta)
}

// Cas 执行针对值的比较并交换操作。 md5:4c2d06b4167bee48
func (v *Int64) Cas(old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(&v.value, old, new)
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *Int64) String() string {
	return strconv.FormatInt(v.Val(), 10)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Int64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(v.Val(), 10)), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Int64) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Int64(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *Int64) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Int64(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Int64) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewInt64(v.Val())
}
