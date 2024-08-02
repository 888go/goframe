// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 安全变量类

import (
	"bytes"
	"encoding/base64"
	"sync/atomic"

	gerror "github.com/888go/goframe/errors/gerror"
	gconv "github.com/888go/goframe/util/gconv"
)

// Bytes 是一个用于并发安全操作的[]byte类型的结构体。 md5:784dc0993857ec47
type Bytes struct {
	value atomic.Value
}

// NewBytes 创建并返回一个针对 []byte 类型的并发安全对象，
// 初始化值为给定的 `value`。
// md5:6aea34a99a4d10ee
func NewBytes(value ...[]byte) *Bytes {
	t := &Bytes{}
	if len(value) > 0 {
		t.value.Store(value[0])
	}
	return t
}

// Clone 创建并返回一个[]byte类型的浅拷贝新对象。 md5:408a6650b2b17fbd
func (v *Bytes) Clone() *Bytes {
	return NewBytes(v.Val())
}

// Set 原子地将 `value` 赋值给 t.value，并返回 t.value 的旧值。
// 注意：参数 `value` 不能为 nil。
// md5:00adcc3b6d3bb3da
func (v *Bytes) Set(value []byte) (old []byte) {
	old = v.Val()
	v.value.Store(value)
	return
}

// Val原子性地加载并返回t.value。 md5:429a11b89436cc12
func (v *Bytes) Val() []byte {
	if s := v.value.Load(); s != nil {
		return s.([]byte)
	}
	return nil
}

// String 实现了 String 接口，用于字符串打印。 md5:9f0b8c0bcf2362d3
func (v *Bytes) String() string {
	return string(v.Val())
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Bytes) MarshalJSON() ([]byte, error) {
	val := v.Val()
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(val)))
	base64.StdEncoding.Encode(dst, val)
	return []byte(`"` + string(dst) + `"`), nil
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Bytes) UnmarshalJSON(b []byte) error {
	var (
		src    = make([]byte, base64.StdEncoding.DecodedLen(len(b)))
		n, err = base64.StdEncoding.Decode(src, bytes.Trim(b, `"`))
	)
	if err != nil {
		err = gerror.Wrap(err, `base64.StdEncoding.Decode failed`)
		return err
	}
	v.Set(src[:n])
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 `v`。 md5:f1b49be4502b95a4
func (v *Bytes) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Bytes(value))
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Bytes) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	oldBytes := v.Val()
	newBytes := make([]byte, len(oldBytes))
	copy(newBytes, oldBytes)
	return NewBytes(newBytes)
}
