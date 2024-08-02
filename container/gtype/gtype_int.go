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

// Int 是一个结构体，用于实现类型int的并发安全操作。 md5:73f0c87f02f7764f
type Int struct {
	value int64
}

// NewInt 创建并返回一个并发安全的 int 类型的对象，其初始值为 `value`。
// md5:2bb0b2973897c335
func NewInt(value ...int) *Int {
	if len(value) > 0 {
		return &Int{
			value: int64(value[0]),
		}
	}
	return &Int{}
}

// Clone 为整型类型创建并返回一个新的并发安全的对象克隆。 md5:170a4dc1f40b5178
func (v *Int) Clone() *Int {
	return NewInt(v.Val())
}

// Set 原子地将 `value` 存储到 t.value 中，并返回 t.value 的旧值。 md5:2ce98b05d0290b37
func (v *Int) Set(value int) (old int) {
	return int(atomic.SwapInt64(&v.value, int64(value)))
}

// Val原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *Int) Val() int {
	return int(atomic.LoadInt64(&v.value))
}

// Atomically 将 `delta` 增加到 t.value 中，并返回新的值。 md5:73547274aea5fe91
func (v *Int) Add(delta int) (new int) {
	return int(atomic.AddInt64(&v.value, int64(delta)))
}

// Cas 执行针对值的比较并交换操作。 md5:4c2d06b4167bee48
func (v *Int) Cas(old, new int) (swapped bool) {
	return atomic.CompareAndSwapInt64(&v.value, int64(old), int64(new))
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *Int) String() string {
	return strconv.Itoa(v.Val())
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Int) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(v.Val())), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Int) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Int(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *Int) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Int(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Int) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewInt(v.Val())
}
