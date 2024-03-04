
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
// Interface returns the json value.
<原文结束>

# <翻译开始>
// Interface 返回 JSON 值。
# <翻译结束>


<原文开始>
// Var returns the json value as *gvar.Var.
<原文结束>

# <翻译开始>
// Var 返回json值作为*gvar.Var类型。
# <翻译结束>


<原文开始>
// IsNil checks whether the value pointed by `j` is nil.
<原文结束>

# <翻译开始>
// IsNil 检查通过 `j` 指向的值是否为 nil。
# <翻译结束>


<原文开始>
// Get retrieves and returns value by specified `pattern`.
// It returns all values of current Json object if `pattern` is given ".".
// It returns nil if no value found by `pattern`.
//
// We can also access slice item by its index number in `pattern` like:
// "list.10", "array.0.name", "array.0.1.id".
//
// It returns a default value specified by `def` if value for `pattern` is not found.
<原文结束>

# <翻译开始>
// Get 方法通过指定的`pattern`获取并返回值。
// 如果 `pattern` 为"."，则返回当前Json对象的所有值。
// 如果根据`pattern`未找到值，则返回nil。
//
// 我们还可以通过在`pattern`中指定切片的索引号来访问切片项，例如：
// "list.10", "array.0.name", "array.0.1.id"。
//
// 如果未找到`pattern`对应的值，将返回由`def`指定的默认值。
# <翻译结束>


<原文开始>
// It returns nil if pattern is empty.
<原文结束>

# <翻译开始>
// 如果模式为空，则返回nil。
# <翻译结束>


<原文开始>
// GetJson gets the value by specified `pattern`,
// and converts it to an un-concurrent-safe Json object.
<原文结束>

# <翻译开始>
// GetJson 通过指定的 `pattern` 获取值，
// 并将其转换为一个非并发安全的 Json 对象。
# <翻译结束>


<原文开始>
// GetJsons gets the value by specified `pattern`,
// and converts it to a slice of un-concurrent-safe Json object.
<原文结束>

# <翻译开始>
// GetJsons 通过指定的 `pattern` 获取值，
// 并将其转换为一个非并发安全的 Json 对象切片。
# <翻译结束>


<原文开始>
// GetJsonMap gets the value by specified `pattern`,
// and converts it to a map of un-concurrent-safe Json object.
<原文结束>

# <翻译开始>
// GetJsonMap 通过指定的 `pattern` 获取值，
// 并将其转换为一个非并发安全的 Json 对象映射。
# <翻译结束>


<原文开始>
// Set sets value with specified `pattern`.
// It supports hierarchical data access by char separator, which is '.' in default.
<原文结束>

# <翻译开始>
// Set 使用指定的`pattern`设置值。
// 它支持通过默认为'.'的字符分隔符进行层级数据访问。
# <翻译结束>


<原文开始>
// MustSet performs as Set, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustSet 表现如同 Set，但当发生任何错误时会触发panic（异常）。
# <翻译结束>


<原文开始>
// Remove deletes value with specified `pattern`.
// It supports hierarchical data access by char separator, which is '.' in default.
<原文结束>

# <翻译开始>
// Remove 删除具有指定 `pattern` 的值。
// 它支持通过默认为 '.' 的字符分隔符进行层级数据访问。
# <翻译结束>


<原文开始>
// MustRemove performs as Remove, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustRemove 的行为与 Remove 相同，但是当发生任何错误时，它会触发 panic（异常）。
# <翻译结束>


<原文开始>
// Contains checks whether the value by specified `pattern` exist.
<原文结束>

# <翻译开始>
// Contains 检查指定的 `pattern` 是否存在对应的值。
# <翻译结束>


<原文开始>
// Len returns the length/size of the value by specified `pattern`.
// The target value by `pattern` should be type of slice or map.
// It returns -1 if the target value is not found, or its type is invalid.
<原文结束>

# <翻译开始>
// Len 通过指定的 `pattern` 返回值的长度/大小。
// 通过 `pattern` 所获取的目标值应为切片或字典类型。
// 若目标值未找到，或者其类型无效，则返回 -1。
# <翻译结束>


<原文开始>
// Append appends value to the value by specified `pattern`.
// The target value by `pattern` should be type of slice.
<原文结束>

# <翻译开始>
// Append 函数通过指定的 `pattern` 追加值到目标值中。
// 通过 `pattern` 所获取的目标值应当是切片类型。
# <翻译结束>


<原文开始>
// MustAppend performs as Append, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustAppend 的行为与 Append 相同，但当发生任何错误时，它会触发panic（异常）。
# <翻译结束>


<原文开始>
// Map converts current Json object to map[string]interface{}.
// It returns nil if fails.
<原文结束>

# <翻译开始>
// Map 将当前的 Json 对象转换为 map[string]interface{} 类型。
// 如果转换失败，则返回 nil。
# <翻译结束>


<原文开始>
// Array converts current Json object to []interface{}.
// It returns nil if fails.
<原文结束>

# <翻译开始>
// Array 将当前Json对象转换为 []interface{} 类型的切片。
// 如果转换失败，则返回nil。
# <翻译结束>


<原文开始>
// Scan automatically calls Struct or Structs function according to the type of parameter
// `pointer` to implement the converting.
<原文结束>

# <翻译开始>
// Scan 会根据参数 `pointer` 的类型自动调用 Struct 或 Structs 函数来实现转换功能。
# <翻译结束>


<原文开始>
// Dump prints current Json object with more manually readable.
<原文结束>

# <翻译开始>
// Dump 打印当前Json对象，使其更易于人工阅读。
# <翻译结束>

