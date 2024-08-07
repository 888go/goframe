// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 安全变量类

import (
	"sync/atomic"

	"github.com/888go/goframe/internal/deepcopy"
	"github.com/888go/goframe/internal/json"
	gconv "github.com/888go/goframe/util/gconv"
)

// Interface 是一个结构体，用于实现类型 interface{} 的并发安全操作。 md5:5655f929d7777a3d
type Interface struct {
	value atomic.Value
}

// NewInterface 创建并返回一个并发安全的对象，用于interface{}类型，初始值为`value`。
// md5:4f93c81a49f5b2f6
func NewInterface(value ...interface{}) *Interface {
	t := &Interface{}
	if len(value) > 0 && value[0] != nil {
		t.value.Store(value[0])
	}
	return t
}

// Clone 为接口类型创建并返回一个新的并发安全的对象。 md5:ea3e89ab199c1ad7
func (v *Interface) Clone() *Interface {
	return NewInterface(v.X取值())
}

// X设置值 原子地将 `value` 赋值给 t.value，并返回 t.value 的旧值。
// 注意：参数 `value` 不能为 nil。
// md5:00adcc3b6d3bb3da
func (v *Interface) X设置值(value interface{}) (old interface{}) {
	old = v.X取值()
	v.value.Store(value)
	return
}

// X取值原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *Interface) X取值() interface{} {
	return v.value.Load()
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *Interface) String() string {
	return gconv.String(v.X取值())
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Interface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.X取值())
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Interface) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.UnmarshalUseNumber(b, &i); err != nil {
		return err
	}
	v.X设置值(i)
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *Interface) UnmarshalValue(value interface{}) error {
	v.X设置值(value)
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Interface) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewInterface(deepcopy.Copy(v.X取值()))
}
