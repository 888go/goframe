// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtype

import (
	"bytes"
	"sync/atomic"

	"github.com/gogf/gf/v2/util/gconv"
)

// Bool 是一个用于并发安全操作布尔类型的结构体。. md5:1d3e571c42d4a013
type Bool struct {
	value int32
}

var (
	bytesTrue  = []byte("true")
	bytesFalse = []byte("false")
)

// NewBool 创建并返回一个针对布尔类型的并发安全对象，
// 初始化值为 `value`。
// md5:d6d603ef4fb898a9
func NewBool(value ...bool) *Bool {
	t := &Bool{}
	if len(value) > 0 {
		if value[0] {
			t.value = 1
		} else {
			t.value = 0
		}
	}
	return t
}

// Clone 克隆并返回一个新的布尔类型的并发安全对象。. md5:097dd9b0b48ac960
func (v *Bool) Clone() *Bool {
	return NewBool(v.Val())
}

// Set 原子地将 `value` 存储到 t.value 中，并返回 t.value 的旧值。. md5:2ce98b05d0290b37
func (v *Bool) Set(value bool) (old bool) {
	if value {
		old = atomic.SwapInt32(&v.value, 1) == 1
	} else {
		old = atomic.SwapInt32(&v.value, 0) == 1
	}
	return
}

// Val原子性地加载并返回t.value。. md5:429a11b89436cc12
func (v *Bool) Val() bool {
	return atomic.LoadInt32(&v.value) > 0
}

// Cas 执行针对值的比较并交换操作。. md5:4c2d06b4167bee48
func (v *Bool) Cas(old, new bool) (swapped bool) {
	var oldInt32, newInt32 int32
	if old {
		oldInt32 = 1
	}
	if new {
		newInt32 = 1
	}
	return atomic.CompareAndSwapInt32(&v.value, oldInt32, newInt32)
}

// String 实现了 String 接口，用于字符串打印。. md5:9f0b8c0bcf2362d3
func (v *Bool) String() string {
	if v.Val() {
		return "true"
	}
	return "false"
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。. md5:43c3b36e60a18f9a
func (v Bool) MarshalJSON() ([]byte, error) {
	if v.Val() {
		return bytesTrue, nil
	}
	return bytesFalse, nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。. md5:f6766b88cf3d63c2
func (v *Bool) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Bool(bytes.Trim(b, `"`)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。. md5:f1b49be4502b95a4
func (v *Bool) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Bool(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。. md5:9cfbcb08109f6ce1
func (v *Bool) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewBool(v.Val())
}
