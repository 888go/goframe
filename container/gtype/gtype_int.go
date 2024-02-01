// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtype
import (
	"strconv"
	"sync/atomic"
	
	"github.com/888go/goframe/util/gconv"
	)
// Int 是一个结构体，用于对 int 类型进行并发安全操作。
type Int struct {
	value int64
}

// NewInt 创建并返回一个用于 int 类型的并发安全对象，
// 其初始值为给定的 `value`。
func NewInt(value ...int) *Int {
	if len(value) > 0 {
		return &Int{
			value: int64(value[0]),
		}
	}
	return &Int{}
}

// Clone 克隆并返回一个新的适用于 int 类型的并发安全对象。
func (v *Int) Clone() *Int {
	return NewInt(v.Val())
}

// Set 方法通过原子操作将`value`存储到t.value中，并返回修改前的t.value的值。
func (v *Int) Set(value int) (old int) {
	return int(atomic.SwapInt64(&v.value, int64(value)))
}

// Val 原子性地加载并返回 t.value。
func (v *Int) Val() int {
	return int(atomic.LoadInt64(&v.value))
}

// Add 原子性地将 `delta` 加到 t.value 上，并返回新的值。
func (v *Int) Add(delta int) (new int) {
	return int(atomic.AddInt64(&v.value, int64(delta)))
}

// Cas 执行值的比较并交换操作。
func (v *Int) Cas(old, new int) (swapped bool) {
	return atomic.CompareAndSwapInt64(&v.value, int64(old), int64(new))
}

// String 实现了 String 接口以便进行字符串打印。
func (v *Int) String() string {
	return strconv.Itoa(v.Val())
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Int) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(v.Val())), nil
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Int) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Int(string(b)))
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 `v` 设置任意类型的值。
func (v *Int) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Int(value))
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Int) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewInt(v.Val())
}
