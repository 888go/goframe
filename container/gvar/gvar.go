// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gvar提供了一种通用变量类型，类似于泛型。
package gvar
import (
	"time"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/internal/deepcopy"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
	)
// Var 是一个通用变量类型的实现者。
type Var struct {
	value interface{} // Underlying value.
	safe  bool        // 是否支持并发安全
}

// New 函数用于创建并返回一个具有给定 `value` 的新 Var。
// 可选参数 `safe` 指定了 Var 是否在并发安全环境下使用，默认为 false。
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

// Copy 对当前 Var 进行深度复制，并返回指向新复制得到的 Var 的指针。
func (v *Var) Copy() *Var {
	return New(gutil.Copy(v.Val()), v.safe)
}

// Clone 执行当前 Var 的浅复制，并返回指向此 Var 的指针。
func (v *Var) Clone() *Var {
	return New(v.Val(), v.safe)
}

// Set将`value`设置为`v`，并返回旧的值。
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

// Val 返回当前变量 `v` 的值。
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

// Interface 是 Val 的别名。
func (v *Var) Interface() interface{} {
	return v.Val()
}

// Bytes 将 `v` 转换并返回为 []byte 类型。
func (v *Var) Bytes() []byte {
	return gconv.Bytes(v.Val())
}

// String将`v`转换并以字符串形式返回。
func (v *Var) String() string {
	return gconv.String(v.Val())
}

// Bool将`v`转换并作为布尔值返回。
func (v *Var) Bool() bool {
	return gconv.Bool(v.Val())
}

// Int 将 `v` 转换并返回为 int 类型。
func (v *Var) Int() int {
	return gconv.Int(v.Val())
}

// Int8将`v`转换并返回为int8类型。
func (v *Var) Int8() int8 {
	return gconv.Int8(v.Val())
}

// Int16将`v`转换并返回为int16类型。
func (v *Var) Int16() int16 {
	return gconv.Int16(v.Val())
}

// Int32将`v`转换为int32类型并返回。
func (v *Var) Int32() int32 {
	return gconv.Int32(v.Val())
}

// Int64将`v`转换并作为int64类型返回。
func (v *Var) Int64() int64 {
	return gconv.Int64(v.Val())
}

// Uint将`v`转换并作为uint类型返回。
func (v *Var) Uint() uint {
	return gconv.Uint(v.Val())
}

// Uint8将`v`转换并作为uint8类型返回。
func (v *Var) Uint8() uint8 {
	return gconv.Uint8(v.Val())
}

// Uint16将`v`转换并作为uint16类型返回。
func (v *Var) Uint16() uint16 {
	return gconv.Uint16(v.Val())
}

// Uint32将`v`转换并作为uint32类型返回。
func (v *Var) Uint32() uint32 {
	return gconv.Uint32(v.Val())
}

// Uint64将`v`转换并作为uint64类型返回。
func (v *Var) Uint64() uint64 {
	return gconv.Uint64(v.Val())
}

// Float32将`v`转换为float32类型并返回。
func (v *Var) Float32() float32 {
	return gconv.Float32(v.Val())
}

// Float64将`v`转换为float64类型并返回。
func (v *Var) Float64() float64 {
	return gconv.Float64(v.Val())
}

// Time将`v`转换并返回为time.Time类型。
// 参数`format`用于指定时间字符串的格式，采用gtime格式规范，
// 例如：Y-m-d H:i:s。
func (v *Var) Time(format ...string) time.Time {
	return gconv.Time(v.Val(), format...)
}

// Duration 将 `v` 转换并返回为 time.Duration 类型。
// 如果 `v` 的值为字符串，那么它将使用 time.ParseDuration 进行转换。
func (v *Var) Duration() time.Duration {
	return gconv.Duration(v.Val())
}

// GTime 将 `v` 转换并返回为 *gtime.Time 类型。
// 参数 `format` 指定了时间字符串的格式，遵循 gtime 的规则，
// 例如：Y-m-d H:i:s。
func (v *Var) GTime(format ...string) *gtime.Time {
	return gconv.GTime(v.Val(), format...)
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Var) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Val())
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Var) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.UnmarshalUseNumber(b, &i); err != nil {
		return err
	}
	v.Set(i)
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 Var 设置任意类型的值。
func (v *Var) UnmarshalValue(value interface{}) error {
	v.Set(value)
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Var) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return New(deepcopy.Copy(v.Val()), v.safe)
}
