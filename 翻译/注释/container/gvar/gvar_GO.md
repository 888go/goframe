
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Package gvar provides an universal variable type, like generics.
<原文结束>

# <翻译开始>
// gvar 包提供了一个通用变量类型，类似于泛型。 md5:edfcd2c00687a1cf
# <翻译结束>


<原文开始>
// Var is an universal variable type implementer.
<原文结束>

# <翻译开始>
// Var 是一个通用变量类型的实现者。 md5:8d1126ac62635ed2
# <翻译结束>


<原文开始>
// Concurrent safe or not.
<原文结束>

# <翻译开始>
// 是否是并发安全的。 md5:b857aa81bf287914
# <翻译结束>


<原文开始>
// New creates and returns a new Var with given `value`.
// The optional parameter `safe` specifies whether Var is used in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个具有给定`value`的新Var。
// 可选参数`safe`指定Var是否用于并发安全，默认为false。 md5:451fb2bb36ca4e4f
# <翻译结束>


<原文开始>
// Copy does a deep copy of current Var and returns a pointer to this Var.
<原文结束>

# <翻译开始>
// Copy 对当前的 Var 进行深拷贝，并返回指向这个新 Var 的指针。 md5:78d7c2be2a0563f7
# <翻译结束>


<原文开始>
// Clone does a shallow copy of current Var and returns a pointer to this Var.
<原文结束>

# <翻译开始>
// Clone 创建当前Var的浅拷贝，并返回指向这个Var的指针。 md5:1f467c25c395f6f1
# <翻译结束>


<原文开始>
// Set sets `value` to `v`, and returns the old value.
<原文结束>

# <翻译开始>
// Set 将 `value` 设置为 `v`，并返回旧值。 md5:ee2b9da700fa7f95
# <翻译结束>


<原文开始>
// Val returns the current value of `v`.
<原文结束>

# <翻译开始>
// Val返回当前的`v`值。 md5:6c5265469db610f7
# <翻译结束>


<原文开始>
// Interface is alias of Val.
<原文结束>

# <翻译开始>
// Interface 是 Val 的别名。 md5:7ddc9573cd7d9927
# <翻译结束>


<原文开始>
// Bytes converts and returns `v` as []byte.
<原文结束>

# <翻译开始>
// Bytes 将 `v` 转换并返回为 []byte。 md5:f6ac565af1bd5f76
# <翻译结束>


<原文开始>
// String converts and returns `v` as string.
<原文结束>

# <翻译开始>
// String 将 `v` 转换为字符串并返回。 md5:773073091c0b6fb0
# <翻译结束>


<原文开始>
// Bool converts and returns `v` as bool.
<原文结束>

# <翻译开始>
// Bool 将 `v` 转换为布尔值并返回。 md5:cb5fceb22f0740d6
# <翻译结束>


<原文开始>
// Int converts and returns `v` as int.
<原文结束>

# <翻译开始>
// Int 将 `v` 转换并返回为 int 类型。 md5:0edb94d8263e3c57
# <翻译结束>


<原文开始>
// Int8 converts and returns `v` as int8.
<原文结束>

# <翻译开始>
// Int8 将 `v` 转换并返回为 int8 类型。 md5:6854263a414a9d3e
# <翻译结束>


<原文开始>
// Int16 converts and returns `v` as int16.
<原文结束>

# <翻译开始>
// Int16 将 `v` 转换为 int16 并返回。 md5:880f0d0288aaaf50
# <翻译结束>


<原文开始>
// Int32 converts and returns `v` as int32.
<原文结束>

# <翻译开始>
// Int32 将 `v` 转换为 int32 并返回。 md5:ba00aec88defc21e
# <翻译结束>


<原文开始>
// Int64 converts and returns `v` as int64.
<原文结束>

# <翻译开始>
// Int64 将 `v` 转换并返回为 int64 类型。 md5:d4d88962698d555e
# <翻译结束>


<原文开始>
// Uint converts and returns `v` as uint.
<原文结束>

# <翻译开始>
// Uint 将 `v` 转换并返回为无符号整数。 md5:5c94bb67c818fb47
# <翻译结束>


<原文开始>
// Uint8 converts and returns `v` as uint8.
<原文结束>

# <翻译开始>
// Uint8 将 `v` 转换为 uint8 并返回。 md5:aa0db1622c86fbf4
# <翻译结束>


<原文开始>
// Uint16 converts and returns `v` as uint16.
<原文结束>

# <翻译开始>
// Uint16 将 `v` 转换为 uint16 并返回。 md5:45ebb672f56f12b0
# <翻译结束>


<原文开始>
// Uint32 converts and returns `v` as uint32.
<原文结束>

# <翻译开始>
// Uint32 将 `v` 转换并返回为 uint32 类型。 md5:b37b73d600b5c94f
# <翻译结束>


<原文开始>
// Uint64 converts and returns `v` as uint64.
<原文结束>

# <翻译开始>
// Uint64 将 `v` 转换并返回为 uint64 类型。 md5:b9d756b5c1231aaa
# <翻译结束>


<原文开始>
// Float32 converts and returns `v` as float32.
<原文结束>

# <翻译开始>
// Float32 将 `v` 转换为 float32 并返回。 md5:10c3ad7673a95ff1
# <翻译结束>


<原文开始>
// Float64 converts and returns `v` as float64.
<原文结束>

# <翻译开始>
// Float64 将 `v` 转换为 float64 并返回。 md5:0dd01006c903cd28
# <翻译结束>


<原文开始>
// Time converts and returns `v` as time.Time.
// The parameter `format` specifies the format of the time string using gtime,
// eg: Y-m-d H:i:s.
<原文结束>

# <翻译开始>
// Time 将 `v` 转换并返回为 time.Time 类型。
// 参数 `format` 使用 gtime 指定时间字符串的格式，
// 例如：Y-m-d H:i:s。 md5:f8b0cb9b11c12546
# <翻译结束>


<原文开始>
// Duration converts and returns `v` as time.Duration.
// If value of `v` is string, then it uses time.ParseDuration for conversion.
<原文结束>

# <翻译开始>
// Duration 将 `v` 转换并返回为 time.Duration 类型。
// 如果 `v` 的值为字符串，那么它会使用 time.ParseDuration 进行转换。 md5:202e87ef6d521c17
# <翻译结束>


<原文开始>
// GTime converts and returns `v` as *gtime.Time.
// The parameter `format` specifies the format of the time string using gtime,
// eg: Y-m-d H:i:s.
<原文结束>

# <翻译开始>
// GTime将`v`转换为*gtime.Time并返回。
// 参数`format`使用gtime指定时间字符串的格式，例如：Y-m-d H:i:s。 md5:0809b54d564e1570
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
# <翻译结束>


<原文开始>
// UnmarshalValue is an interface implement which sets any type of value for Var.
<原文结束>

# <翻译开始>
// UnmarshalValue 是一个接口实现，用于将任何类型的价值设置为 Var。 md5:c6a2fce2313ec90f
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
# <翻译结束>

