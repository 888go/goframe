// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 安全变量类

import (
	"strconv"
	"sync/atomic"
	
	"github.com/888go/goframe/util/gconv"
)

// Int64 是一个结构体，用于对 int64 类型进行并发安全操作。
type Int64 struct {
	value int64
}

// NewInt64 创建并返回一个适用于 int64 类型的并发安全对象，
// 其初始值为给定的 `value`。
func NewInt64(value ...int64) *Int64 {
	if len(value) > 0 {
		return &Int64{
			value: value[0],
		}
	}
	return &Int64{}
}

// Clone 克隆并返回一个用于 int64 类型的新并发安全对象。
func (v *Int64) Clone() *Int64 {
	return NewInt64(v.X取值())
}

// Set 方法通过原子操作将`value`存储到t.value中，并返回修改前的t.value的值。
func (v *Int64) X设置值(value int64) (old int64) {
	return atomic.SwapInt64(&v.value, value)
}

// Val 原子性地加载并返回 t.value。
func (v *Int64) X取值() int64 {
	return atomic.LoadInt64(&v.value)
}

// Add 原子性地将 `delta` 加到 t.value 上，并返回新的值。
func (v *Int64) Add(delta int64) (new int64) {
	return atomic.AddInt64(&v.value, delta)
}

// Cas 执行值的比较并交换操作。
func (v *Int64) Cas(old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(&v.value, old, new)
}

// String 实现了 String 接口以便进行字符串打印。
func (v *Int64) String() string {
	return strconv.FormatInt(v.X取值(), 10)
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Int64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(v.X取值(), 10)), nil
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Int64) UnmarshalJSON(b []byte) error {
	v.X设置值(转换类.X取整数64位(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 `v` 设置任意类型的值。
func (v *Int64) UnmarshalValue(value interface{}) error {
	v.X设置值(转换类.X取整数64位(value))
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Int64) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewInt64(v.X取值())
}
