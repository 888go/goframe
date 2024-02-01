
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Package gvar provides an universal variable type, like generics.
<原文结束>

# <翻译开始>
// 包gvar提供了一种通用变量类型，类似于泛型。
# <翻译结束>


<原文开始>
// Var is an universal variable type implementer.
<原文结束>

# <翻译开始>
// Var 是一个通用变量类型的实现者。
# <翻译结束>







<原文开始>
// New creates and returns a new Var with given `value`.
// The optional parameter `safe` specifies whether Var is used in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// New 函数用于创建并返回一个具有给定 `value` 的新 Var。
// 可选参数 `safe` 指定了 Var 是否在并发安全环境下使用，默认为 false。
# <翻译结束>


<原文开始>
// Copy does a deep copy of current Var and returns a pointer to this Var.
<原文结束>

# <翻译开始>
// Copy 对当前 Var 进行深度复制，并返回指向新复制得到的 Var 的指针。
# <翻译结束>


<原文开始>
// Clone does a shallow copy of current Var and returns a pointer to this Var.
<原文结束>

# <翻译开始>
// Clone 执行当前 Var 的浅复制，并返回指向此 Var 的指针。
# <翻译结束>


<原文开始>
// Set sets `value` to `v`, and returns the old value.
<原文结束>

# <翻译开始>
// Set将`value`设置为`v`，并返回旧的值。
# <翻译结束>


<原文开始>
// Val returns the current value of `v`.
<原文结束>

# <翻译开始>
// Val 返回当前变量 `v` 的值。
# <翻译结束>







<原文开始>
// Bytes converts and returns `v` as []byte.
<原文结束>

# <翻译开始>
// Bytes 将 `v` 转换并返回为 []byte 类型。
# <翻译结束>


<原文开始>
// String converts and returns `v` as string.
<原文结束>

# <翻译开始>
// String将`v`转换并以字符串形式返回。
# <翻译结束>


<原文开始>
// Bool converts and returns `v` as bool.
<原文结束>

# <翻译开始>
// Bool将`v`转换并作为布尔值返回。
# <翻译结束>


<原文开始>
// Int converts and returns `v` as int.
<原文结束>

# <翻译开始>
// Int 将 `v` 转换并返回为 int 类型。
# <翻译结束>


<原文开始>
// Int8 converts and returns `v` as int8.
<原文结束>

# <翻译开始>
// Int8将`v`转换并返回为int8类型。
# <翻译结束>


<原文开始>
// Int16 converts and returns `v` as int16.
<原文结束>

# <翻译开始>
// Int16将`v`转换并返回为int16类型。
# <翻译结束>


<原文开始>
// Int32 converts and returns `v` as int32.
<原文结束>

# <翻译开始>
// Int32将`v`转换为int32类型并返回。
# <翻译结束>


<原文开始>
// Int64 converts and returns `v` as int64.
<原文结束>

# <翻译开始>
// Int64将`v`转换并作为int64类型返回。
# <翻译结束>


<原文开始>
// Uint converts and returns `v` as uint.
<原文结束>

# <翻译开始>
// Uint将`v`转换并作为uint类型返回。
# <翻译结束>


<原文开始>
// Uint8 converts and returns `v` as uint8.
<原文结束>

# <翻译开始>
// Uint8将`v`转换并作为uint8类型返回。
# <翻译结束>


<原文开始>
// Uint16 converts and returns `v` as uint16.
<原文结束>

# <翻译开始>
// Uint16将`v`转换并作为uint16类型返回。
# <翻译结束>


<原文开始>
// Uint32 converts and returns `v` as uint32.
<原文结束>

# <翻译开始>
// Uint32将`v`转换并作为uint32类型返回。
# <翻译结束>


<原文开始>
// Uint64 converts and returns `v` as uint64.
<原文结束>

# <翻译开始>
// Uint64将`v`转换并作为uint64类型返回。
# <翻译结束>


<原文开始>
// Float32 converts and returns `v` as float32.
<原文结束>

# <翻译开始>
// Float32将`v`转换为float32类型并返回。
# <翻译结束>


<原文开始>
// Float64 converts and returns `v` as float64.
<原文结束>

# <翻译开始>
// Float64将`v`转换为float64类型并返回。
# <翻译结束>


<原文开始>
// Time converts and returns `v` as time.Time.
// The parameter `format` specifies the format of the time string using gtime,
// eg: Y-m-d H:i:s.
<原文结束>

# <翻译开始>
// Time将`v`转换并返回为time.Time类型。
// 参数`format`用于指定时间字符串的格式，采用gtime格式规范，
// 例如：Y-m-d H:i:s。
# <翻译结束>


<原文开始>
// Duration converts and returns `v` as time.Duration.
// If value of `v` is string, then it uses time.ParseDuration for conversion.
<原文结束>

# <翻译开始>
// Duration 将 `v` 转换并返回为 time.Duration 类型。
// 如果 `v` 的值为字符串，那么它将使用 time.ParseDuration 进行转换。
# <翻译结束>


<原文开始>
// GTime converts and returns `v` as *gtime.Time.
// The parameter `format` specifies the format of the time string using gtime,
// eg: Y-m-d H:i:s.
<原文结束>

# <翻译开始>
// GTime 将 `v` 转换并返回为 *gtime.Time 类型。
// 参数 `format` 指定了时间字符串的格式，遵循 gtime 的规则，
// 例如：Y-m-d H:i:s。
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
# <翻译结束>


<原文开始>
// UnmarshalValue is an interface implement which sets any type of value for Var.
<原文结束>

# <翻译开始>
// UnmarshalValue 是一个接口实现，用于为 Var 设置任意类型的值。
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy 实现接口，用于当前类型的深度复制。
# <翻译结束>







<原文开始>
// Concurrent safe or not.
<原文结束>

# <翻译开始>
// 是否支持并发安全
# <翻译结束>


<原文开始>
// Interface is alias of Val.
<原文结束>

# <翻译开始>
// Interface 是 Val 的别名。
# <翻译结束>

