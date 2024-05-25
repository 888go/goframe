
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
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Interface returns the json value.
<原文结束>

# <翻译开始>
// Interface 返回 json 值。 md5:63120da1c18a889c
# <翻译结束>


<原文开始>
// Var returns the json value as *gvar.Var.
<原文结束>

# <翻译开始>
// Var将json值作为*gvar.Var返回。 md5:fcfd99090165a7b5
# <翻译结束>


<原文开始>
// IsNil checks whether the value pointed by `j` is nil.
<原文结束>

# <翻译开始>
// IsNil 检查指针 `j` 所指向的值是否为 nil。 md5:669f952cb58d64ce
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
// Get 通过指定的`pattern`获取并返回值。
// 如果`pattern`为"."，它将返回当前Json对象的所有值。
// 如果根据`pattern`没有找到值，它将返回nil。
//
// 我们也可以在`pattern`中通过索引数字来访问切片的元素，例如：
// "list.10", "array.0.name", "array.0.1.id"。
//
// 如果没有为`pattern`找到值，它将返回由`def`指定的默认值。
// md5:f56f76d635296903
# <翻译结束>


<原文开始>
// It returns nil if pattern is empty.
<原文结束>

# <翻译开始>
// 如果pattern为空，它将返回nil。 md5:8e2a6f56affd353a
# <翻译结束>


<原文开始>
// GetJson gets the value by specified `pattern`,
// and converts it to an un-concurrent-safe Json object.
<原文结束>

# <翻译开始>
// GetJson 通过指定的`pattern`获取值，并将其转换为一个非并发安全的Json对象。
// md5:e0f60541ee6017b5
# <翻译结束>


<原文开始>
// GetJsons gets the value by specified `pattern`,
// and converts it to a slice of un-concurrent-safe Json object.
<原文结束>

# <翻译开始>
// GetJsons 通过指定的`pattern`获取值，并将其转换为一个不并发安全的Json对象切片。
// md5:1bd75964e1b32ed2
# <翻译结束>


<原文开始>
// GetJsonMap gets the value by specified `pattern`,
// and converts it to a map of un-concurrent-safe Json object.
<原文结束>

# <翻译开始>
// GetJsonMap 通过指定的`pattern`获取值，
// 并将其转换为非并发安全的Json对象映射。
// md5:d549d238d186a4e0
# <翻译结束>


<原文开始>
// Set sets value with specified `pattern`.
// It supports hierarchical data access by char separator, which is '.' in default.
<原文结束>

# <翻译开始>
// Set 使用指定的 `pattern` 设置值。
// 它支持通过字符分隔符（默认为'.'）进行层次数据访问。
// md5:85400f8aa43895d6
# <翻译结束>


<原文开始>
// MustSet performs as Set, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustSet 执行与 Set 相同的操作，但如果发生任何错误，它将引发恐慌。 md5:89753cb5f56f60cc
# <翻译结束>


<原文开始>
// Remove deletes value with specified `pattern`.
// It supports hierarchical data access by char separator, which is '.' in default.
<原文结束>

# <翻译开始>
// Remove 删除具有指定`pattern`的值。它支持通过字符分隔符（默认为`.`）进行层次数据访问。
// md5:a8bd1b8b0e8d7d8e
# <翻译结束>


<原文开始>
// MustRemove performs as Remove, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustRemove 的行为与 Remove 相同，但如果发生任何错误，它会直接 panic。 md5:ad4ac7324486398a
# <翻译结束>


<原文开始>
// Contains checks whether the value by specified `pattern` exist.
<原文结束>

# <翻译开始>
// Contains 检查是否存在指定的 `pattern` 值。 md5:4f248b6aebb74d05
# <翻译结束>


<原文开始>
// Len returns the length/size of the value by specified `pattern`.
// The target value by `pattern` should be type of slice or map.
// It returns -1 if the target value is not found, or its type is invalid.
<原文结束>

# <翻译开始>
// Len 返回由指定 `pattern` 定义的值的长度/大小。目标值应该是切片或映射类型。如果找不到目标值或者其类型无效，它将返回 -1。
// md5:f929eb27a0ef1a36
# <翻译结束>


<原文开始>
// Append appends value to the value by specified `pattern`.
// The target value by `pattern` should be type of slice.
<原文结束>

# <翻译开始>
// Append 将指定的 `pattern` 所引用的值追加到目标值（应该是切片类型）中。
// md5:5b8e7f4c493419ba
# <翻译结束>


<原文开始>
// MustAppend performs as Append, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustAppend 的行为与 Append 相同，但如果发生任何错误，它会直接 panic。 md5:3a0acd5a244f264f
# <翻译结束>


<原文开始>
// Map converts current Json object to map[string]interface{}.
// It returns nil if fails.
<原文结束>

# <翻译开始>
// Map 将当前的 Json 对象转换为 map[string]interface{}。
// 如果转换失败，它将返回 nil。
// md5:599d2c152000d26b
# <翻译结束>


<原文开始>
// Array converts current Json object to []interface{}.
// It returns nil if fails.
<原文结束>

# <翻译开始>
// Array 将当前Json对象转换为 []interface{} 类型。如果转换失败，它将返回nil。
// md5:8b3042473c46995f
# <翻译结束>


<原文开始>
// Scan automatically calls Struct or Structs function according to the type of parameter
// `pointer` to implement the converting.
<原文结束>

# <翻译开始>
// Scan会根据参数类型自动调用Struct或Structs函数，并通过`pointer`实现转换。
// md5:afdb5ab720fddc3b
# <翻译结束>


<原文开始>
// Dump prints current Json object with more manually readable.
<原文结束>

# <翻译开始>
// Dump 打印当前的Json对象，使其更便于人工阅读。 md5:c8c6bbdb40fa6383
# <翻译结束>

