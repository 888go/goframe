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

// Byte是用于并发安全操作byte类型的结构体。 md5:260dada42dab3948
type Byte struct {
	value int32
}

// NewByte 创建并返回一个针对字节类型的安全并发对象，
// 初始化值为`value`。
// md5:cff320090f7808b6
func NewByte(value ...byte) *Byte {
	if len(value) > 0 {
		return &Byte{
			value: int32(value[0]),
		}
	}
	return &Byte{}
}

// Clone 克隆并返回一个新的并发安全的字节类型对象。 md5:8e0d468a5dfb5e0e
func (v *Byte) Clone() *Byte {
	return NewByte(v.Val())
}

// Set 原子地将 `value` 存储到 t.value 中，并返回 t.value 的旧值。 md5:2ce98b05d0290b37
func (v *Byte) Set(value byte) (old byte) {
	return byte(atomic.SwapInt32(&v.value, int32(value)))
}

// Val原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *Byte) Val() byte {
	return byte(atomic.LoadInt32(&v.value))
}

// Atomically 将 `delta` 增加到 t.value 中，并返回新的值。 md5:73547274aea5fe91
func (v *Byte) Add(delta byte) (new byte) {
	return byte(atomic.AddInt32(&v.value, int32(delta)))
}

// Cas 执行针对值的比较并交换操作。 md5:4c2d06b4167bee48
func (v *Byte) Cas(old, new byte) (swapped bool) {
	return atomic.CompareAndSwapInt32(&v.value, int32(old), int32(new))
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *Byte) String() string {
	return strconv.FormatUint(uint64(v.Val()), 10)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Byte) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(v.Val()), 10)), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Byte) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Uint8(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *Byte) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Byte(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Byte) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewByte(v.Val())
}
