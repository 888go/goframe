// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 安全变量类

import (
	"bytes"
	"encoding/base64"
	"sync/atomic"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/util/gconv"
)

// Bytes 是一个结构体，用于对 []byte 类型进行并发安全操作。
type Bytes struct {
	value atomic.Value
}

// NewBytes 创建并返回一个对 []byte 类型安全的并发对象，
// 其初始值为给定的 `value`。
func NewBytes(value ...[]byte) *Bytes {
	t := &Bytes{}
	if len(value) > 0 {
		t.value.Store(value[0])
	}
	return t
}

// Clone 克隆并返回一个新的 []byte 类型的浅复制对象。
func (v *Bytes) Clone() *Bytes {
	return NewBytes(v.X取值())
}

// Set 方法通过原子操作将`value`存储到t.value，并返回修改前的t.value的值。
// 注意：参数`value`不能为空。
func (v *Bytes) X设置值(value []byte) (old []byte) {
	old = v.X取值()
	v.value.Store(value)
	return
}

// Val 原子性地加载并返回 t.value。
func (v *Bytes) X取值() []byte {
	if s := v.value.Load(); s != nil {
		return s.([]byte)
	}
	return nil
}

// String 实现了 String 接口以便进行字符串打印。
func (v *Bytes) String() string {
	return string(v.X取值())
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Bytes) MarshalJSON() ([]byte, error) {
	val := v.X取值()
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(val)))
	base64.StdEncoding.Encode(dst, val)
	return []byte(`"` + string(dst) + `"`), nil
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Bytes) UnmarshalJSON(b []byte) error {
	var (
		src    = make([]byte, base64.StdEncoding.DecodedLen(len(b)))
		n, err = base64.StdEncoding.Decode(src, bytes.Trim(b, `"`))
	)
	if err != nil {
		err = 错误类.X多层错误(err, `base64.StdEncoding.Decode failed`)
		return err
	}
	v.X设置值(src[:n])
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 `v` 设置任意类型的值。
func (v *Bytes) UnmarshalValue(value interface{}) error {
	v.X设置值(转换类.X取字节集(value))
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Bytes) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	oldBytes := v.X取值()
	newBytes := make([]byte, len(oldBytes))
	copy(newBytes, oldBytes)
	return NewBytes(newBytes)
}
