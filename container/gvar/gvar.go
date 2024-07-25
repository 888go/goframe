// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// gvar 包提供了一个通用变量类型，类似于泛型。 md5:edfcd2c00687a1cf
package gvar

import (
	"time"

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/internal/deepcopy"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// Var 是一个通用变量类型的实现者。 md5:8d1126ac62635ed2
type Var struct {
	value interface{} // Underlying value.
	safe  bool        // 是否是并发安全的。 md5:b857aa81bf287914
}

// New 创建并返回一个具有给定`value`的新Var。
// 可选参数`safe`指定Var是否用于并发安全，默认为false。 md5:451fb2bb36ca4e4f
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

// Copy 对当前的 Var 进行深拷贝，并返回指向这个新 Var 的指针。 md5:78d7c2be2a0563f7
func (v *Var) Copy() *Var {
	return New(gutil.Copy(v.Val()), v.safe)
}

// Clone 创建当前Var的浅拷贝，并返回指向这个Var的指针。 md5:1f467c25c395f6f1
func (v *Var) Clone() *Var {
	return New(v.Val(), v.safe)
}

// Set 将 `value` 设置为 `v`，并返回旧值。 md5:ee2b9da700fa7f95
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

// Val返回当前的`v`值。 md5:6c5265469db610f7
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

// Interface 是 Val 的别名。 md5:7ddc9573cd7d9927
func (v *Var) Interface() interface{} {
	return v.Val()
}

// Bytes 将 `v` 转换并返回为 []byte。 md5:f6ac565af1bd5f76
func (v *Var) Bytes() []byte {
	return gconv.Bytes(v.Val())
}

// String 将 `v` 转换为字符串并返回。 md5:773073091c0b6fb0
func (v *Var) String() string {
	return gconv.String(v.Val())
}

// Bool 将 `v` 转换为布尔值并返回。 md5:cb5fceb22f0740d6
func (v *Var) Bool() bool {
	return gconv.Bool(v.Val())
}

// Int 将 `v` 转换并返回为 int 类型。 md5:0edb94d8263e3c57
func (v *Var) Int() int {
	return gconv.Int(v.Val())
}

// Int8 将 `v` 转换并返回为 int8 类型。 md5:6854263a414a9d3e
func (v *Var) Int8() int8 {
	return gconv.Int8(v.Val())
}

// Int16 将 `v` 转换为 int16 并返回。 md5:880f0d0288aaaf50
func (v *Var) Int16() int16 {
	return gconv.Int16(v.Val())
}

// Int32 将 `v` 转换为 int32 并返回。 md5:ba00aec88defc21e
func (v *Var) Int32() int32 {
	return gconv.Int32(v.Val())
}

// Int64 将 `v` 转换并返回为 int64 类型。 md5:d4d88962698d555e
func (v *Var) Int64() int64 {
	return gconv.Int64(v.Val())
}

// Uint 将 `v` 转换并返回为无符号整数。 md5:5c94bb67c818fb47
func (v *Var) Uint() uint {
	return gconv.Uint(v.Val())
}

// Uint8 将 `v` 转换为 uint8 并返回。 md5:aa0db1622c86fbf4
func (v *Var) Uint8() uint8 {
	return gconv.Uint8(v.Val())
}

// Uint16 将 `v` 转换为 uint16 并返回。 md5:45ebb672f56f12b0
func (v *Var) Uint16() uint16 {
	return gconv.Uint16(v.Val())
}

// Uint32 将 `v` 转换并返回为 uint32 类型。 md5:b37b73d600b5c94f
func (v *Var) Uint32() uint32 {
	return gconv.Uint32(v.Val())
}

// Uint64 将 `v` 转换并返回为 uint64 类型。 md5:b9d756b5c1231aaa
func (v *Var) Uint64() uint64 {
	return gconv.Uint64(v.Val())
}

// Float32 将 `v` 转换为 float32 并返回。 md5:10c3ad7673a95ff1
func (v *Var) Float32() float32 {
	return gconv.Float32(v.Val())
}

// Float64 将 `v` 转换为 float64 并返回。 md5:0dd01006c903cd28
func (v *Var) Float64() float64 {
	return gconv.Float64(v.Val())
}

// Time 将 `v` 转换并返回为 time.Time 类型。
// 参数 `format` 使用 gtime 指定时间字符串的格式，
// 例如：Y-m-d H:i:s。 md5:f8b0cb9b11c12546
func (v *Var) Time(format ...string) time.Time {
	return gconv.Time(v.Val(), format...)
}

// Duration 将 `v` 转换并返回为 time.Duration 类型。
// 如果 `v` 的值为字符串，那么它会使用 time.ParseDuration 进行转换。 md5:202e87ef6d521c17
func (v *Var) Duration() time.Duration {
	return gconv.Duration(v.Val())
}

// GTime将`v`转换为*gtime.Time并返回。
// 参数`format`使用gtime指定时间字符串的格式，例如：Y-m-d H:i:s。 md5:0809b54d564e1570
func (v *Var) GTime(format ...string) *gtime.Time {
	return gconv.GTime(v.Val(), format...)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (v Var) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Val())
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (v *Var) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.UnmarshalUseNumber(b, &i); err != nil {
		return err
	}
	v.Set(i)
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的价值设置为 Var。 md5:c6a2fce2313ec90f
func (v *Var) UnmarshalValue(value interface{}) error {
	v.Set(value)
	return nil
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (v *Var) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return New(deepcopy.Copy(v.Val()), v.safe)
}
