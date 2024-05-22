// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gvar provides an universal variable type, like generics.
package gvar//bm:泛型类

import (
	"time"

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/internal/deepcopy"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// Var is an universal variable type implementer.
type Var struct {
	value interface{} // Underlying value.
	safe  bool        // Concurrent safe or not.
}

// New creates and returns a new Var with given `value`.
// The optional parameter `safe` specifies whether Var is used in concurrent-safety,
// which is false in default.

// ff:创建
// safe:并发安全
// value:值
func New(value interface{}, safe ...bool) *Var {
	if len(safe) > 0 && safe[0] {
		return &Var{
			value: gtype.NewInterface(value),
			safe:  true,
		}
	}
	return &Var{
		value: value,
	}
}

// Copy does a deep copy of current Var and returns a pointer to this Var.

// ff:深拷贝
func (v *Var) Copy() *Var {
	return New(gutil.Copy(v.Val()), v.safe)
}

// Clone does a shallow copy of current Var and returns a pointer to this Var.

// ff:浅拷贝
func (v *Var) Clone() *Var {
	return New(v.Val(), v.safe)
}

// Set sets `value` to `v`, and returns the old value.

// ff:设置值
// yx:true
// old:
// value:
func (v *Var) Set(value interface{}) (old interface{}) {
	if v.safe {
		if t, ok := v.value.(*gtype.Interface); ok {
			old = t.Set(value)
			return
		}
	}
	old = v.value
	v.value = value
	return
}

// Val returns the current value of `v`.

// ff:取值
// yx:true
func (v *Var) Val() interface{} {
	if v == nil {
		return nil
	}
	if v.safe {
		if t, ok := v.value.(*gtype.Interface); ok {
			return t.Val()
		}
	}
	return v.value
}

// Interface is alias of Val.

// ff:
func (v *Var) Interface() interface{} {
	return v.Val()
}

// Bytes converts and returns `v` as []byte.

// ff:取字节集
// yx:true
func (v *Var) Bytes() []byte {
	return gconv.Bytes(v.Val())
}

// String converts and returns `v` as string.

// ff:
func (v *Var) String() string {
	return gconv.String(v.Val())
}

// Bool converts and returns `v` as bool.

// ff:取布尔
// yx:true
func (v *Var) Bool() bool {
	return gconv.Bool(v.Val())
}

// Int converts and returns `v` as int.

// ff:取整数
func (v *Var) Int() int {
	return gconv.Int(v.Val())
}

// Int8 converts and returns `v` as int8.

// ff:取整数8位
func (v *Var) Int8() int8 {
	return gconv.Int8(v.Val())
}

// Int16 converts and returns `v` as int16.

// ff:取整数16位
func (v *Var) Int16() int16 {
	return gconv.Int16(v.Val())
}

// Int32 converts and returns `v` as int32.

// ff:取整数32位
func (v *Var) Int32() int32 {
	return gconv.Int32(v.Val())
}

// Int64 converts and returns `v` as int64.

// ff:取整数64位
// yx:true
func (v *Var) Int64() int64 {
	return gconv.Int64(v.Val())
}

// Uint converts and returns `v` as uint.

// ff:取正整数
func (v *Var) Uint() uint {
	return gconv.Uint(v.Val())
}

// Uint8 converts and returns `v` as uint8.

// ff:取正整数8位
func (v *Var) Uint8() uint8 {
	return gconv.Uint8(v.Val())
}

// Uint16 converts and returns `v` as uint16.

// ff:取正整数16位
func (v *Var) Uint16() uint16 {
	return gconv.Uint16(v.Val())
}

// Uint32 converts and returns `v` as uint32.

// ff:取正整数32位
func (v *Var) Uint32() uint32 {
	return gconv.Uint32(v.Val())
}

// Uint64 converts and returns `v` as uint64.

// ff:取正整数64位
// yx:true
func (v *Var) Uint64() uint64 {
	return gconv.Uint64(v.Val())
}

// Float32 converts and returns `v` as float32.

// ff:取小数32位
// yx:true
func (v *Var) Float32() float32 {
	return gconv.Float32(v.Val())
}

// Float64 converts and returns `v` as float64.

// ff:取小数64位
// yx:true
func (v *Var) Float64() float64 {
	return gconv.Float64(v.Val())
}

// Time converts and returns `v` as time.Time.
// The parameter `format` specifies the format of the time string using gtime,
// eg: Y-m-d H:i:s.

// ff:取时间类
// format:格式
func (v *Var) Time(format ...string) time.Time {
	return gconv.Time(v.Val(), format...)
}

// Duration converts and returns `v` as time.Duration.
// If value of `v` is string, then it uses time.ParseDuration for conversion.

// ff:取时长
func (v *Var) Duration() time.Duration {
	return gconv.Duration(v.Val())
}

// GTime converts and returns `v` as *gtime.Time.
// The parameter `format` specifies the format of the time string using gtime,
// eg: Y-m-d H:i:s.

// ff:取gtime时间类
// yx:true
// format:
func (v *Var) GTime(format ...string) *gtime.Time {
	return gconv.GTime(v.Val(), format...)
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.

// ff:
func (v Var) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Val())
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

// ff:
// b:
func (v *Var) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.UnmarshalUseNumber(b, &i); err != nil {
		return err
	}
	v.Set(i)
	return nil
}

// UnmarshalValue is an interface implement which sets any type of value for Var.

// ff:
// value:
func (v *Var) UnmarshalValue(value interface{}) error {
	v.Set(value)
	return nil
}

// DeepCopy implements interface for deep copy of current type.

// ff:
func (v *Var) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return New(deepcopy.Copy(v.Val()), v.safe)
}
