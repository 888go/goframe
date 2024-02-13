// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gvar提供了一种通用变量类型，类似于泛型。
package 泛型类

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
func X创建(值 interface{}, 并发安全 ...bool) *Var {
	if len(并发安全) > 0 && 并发安全[0] {
		return &Var{
			value: 安全变量类.NewInterface(值),
			safe:  true,
		}
	}
	return &Var{
		value: 值,
	}
}

// Copy 对当前 Var 进行深度复制，并返回指向新复制得到的 Var 的指针。
func (v *Var) X深拷贝() *Var {
	return X创建(工具类.X深拷贝(v.X取值()), v.safe)
}

// Clone 执行当前 Var 的浅复制，并返回指向此 Var 的指针。
func (v *Var) X浅拷贝() *Var {
	return X创建(v.X取值(), v.safe)
}

// Set将`value`设置为`v`，并返回旧的值。
func (v *Var) X设置值(值 interface{}) (旧值 interface{}) {
	if v.safe {
		if t, ok := v.value.(*安全变量类.Interface); ok {
			旧值 = t.X设置值(值)
			return
		}
	}
	旧值 = v.value
	v.value = 值
	return
}

// Val 返回当前变量 `v` 的值。
func (v *Var) X取值() interface{} {
	if v == nil {
		return nil
	}
	if v.safe {
		if t, ok := v.value.(*安全变量类.Interface); ok {
			return t.X取值()
		}
	}
	return v.value
}

// Interface 是 Val 的别名。
func (v *Var) Interface() interface{} {
	return v.X取值()
}

// Bytes 将 `v` 转换并返回为 []byte 类型。
func (v *Var) X取字节集() []byte {
	return 转换类.X取字节集(v.X取值())
}

// String将`v`转换并以字符串形式返回。
func (v *Var) String() string {
	return 转换类.String(v.X取值())
}

// Bool将`v`转换并作为布尔值返回。
func (v *Var) X取布尔() bool {
	return 转换类.X取布尔(v.X取值())
}

// Int 将 `v` 转换并返回为 int 类型。
func (v *Var) X取整数() int {
	return 转换类.X取整数(v.X取值())
}

// Int8将`v`转换并返回为int8类型。
func (v *Var) X取整数8位() int8 {
	return 转换类.X取整数8位(v.X取值())
}

// Int16将`v`转换并返回为int16类型。
func (v *Var) X取整数16位() int16 {
	return 转换类.X取整数16位(v.X取值())
}

// Int32将`v`转换为int32类型并返回。
func (v *Var) X取整数32位() int32 {
	return 转换类.X取整数32位(v.X取值())
}

// Int64将`v`转换并作为int64类型返回。
func (v *Var) X取整数64位() int64 {
	return 转换类.X取整数64位(v.X取值())
}

// Uint将`v`转换并作为uint类型返回。
func (v *Var) X取正整数() uint {
	return 转换类.X取正整数(v.X取值())
}

// Uint8将`v`转换并作为uint8类型返回。
func (v *Var) X取正整数8位() uint8 {
	return 转换类.X取正整数8位(v.X取值())
}

// Uint16将`v`转换并作为uint16类型返回。
func (v *Var) X取正整数16位() uint16 {
	return 转换类.X取正整数16位(v.X取值())
}

// Uint32将`v`转换并作为uint32类型返回。
func (v *Var) X取正整数32位() uint32 {
	return 转换类.X取正整数32位(v.X取值())
}

// Uint64将`v`转换并作为uint64类型返回。
func (v *Var) X取正整数64位() uint64 {
	return 转换类.X取正整数64位(v.X取值())
}

// Float32将`v`转换为float32类型并返回。
func (v *Var) X取小数32位() float32 {
	return 转换类.X取小数32位(v.X取值())
}

// Float64将`v`转换为float64类型并返回。
func (v *Var) X取小数64位() float64 {
	return 转换类.X取小数64位(v.X取值())
}

// Time将`v`转换并返回为time.Time类型。
// 参数`format`用于指定时间字符串的格式，采用gtime格式规范，
// 例如：Y-m-d H:i:s。
func (v *Var) X取时间类(格式 ...string) time.Time {
	return 转换类.X取时间(v.X取值(), 格式...)
}

// Duration 将 `v` 转换并返回为 time.Duration 类型。
// 如果 `v` 的值为字符串，那么它将使用 time.ParseDuration 进行转换。
func (v *Var) X取时长() time.Duration {
	return 转换类.X取时长(v.X取值())
}

// GTime 将 `v` 转换并返回为 *gtime.Time 类型。
// 参数 `format` 指定了时间字符串的格式，遵循 gtime 的规则，
// 例如：Y-m-d H:i:s。
func (v *Var) X取gtime时间类(格式 ...string) *时间类.Time {
	return 转换类.X取gtime时间类(v.X取值(), 格式...)
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (v Var) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.X取值())
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (v *Var) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.UnmarshalUseNumber(b, &i); err != nil {
		return err
	}
	v.X设置值(i)
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 Var 设置任意类型的值。
func (v *Var) UnmarshalValue(value interface{}) error {
	v.X设置值(value)
	return nil
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (v *Var) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return X创建(deepcopy.Copy(v.X取值()), v.safe)
}
