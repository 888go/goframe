
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
// Scan automatically checks the type of `pointer` and converts `params` to `pointer`.
// It supports `pointer` in type of `*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct` for converting.
//
// TODO change `paramKeyToAttrMap` to `ScanOption` to be more scalable; add `DeepCopy` option for `ScanOption`.
<原文结束>

# <翻译开始>
// Scan 自动检查`pointer`的类型，并将`params`转换为`pointer`。
// 它支持以下类型的`pointer`进行转换：`*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct`。
//
// 待办：将`paramKeyToAttrMap`改为`ScanOption`以提高可扩展性；为`ScanOption`添加`DeepCopy`选项。
// md5:6b1a82a906dd8ec9
# <翻译结束>


<原文开始>
// If `srcValue` is nil, no conversion.
<原文结束>

# <翻译开始>
		// 如果`srcValue`为nil，不进行转换。 md5:72ad5bbbd6b824ce
# <翻译结束>


<原文开始>
// do not use dstPointerReflectValue.Type() as dstPointerReflectValue might be zero.
<原文结束>

# <翻译开始>
		// 不要使用dstPointerReflectValue.Type()，因为dstPointerReflectValue可能为零。 md5:8930f50c4dd8479c
# <翻译结束>


<原文开始>
// pointer kind validation.
<原文结束>

# <翻译开始>
	// 指针类型验证。 md5:4a9fffd5103820ed
# <翻译结束>


<原文开始>
// direct assignment checks!
<原文结束>

# <翻译开始>
	// 直接赋值检查！. md5:fd96cdd5962c2f14
# <翻译结束>


<原文开始>
	// if `srcValue` and `dstPointer` are the same type, the do directly assignment.
	// For performance enhancement purpose.
<原文结束>

# <翻译开始>
	// 如果`srcValue`和`dstPointer`是相同的类型，直接进行赋值。
	// 为了提高性能。
	// md5:5d0efd73f7a58b6f
# <翻译结束>


<原文开始>
	// if `srcValue` and `dstPointer` are the same type, the do directly assignment.
	// for performance enhancement purpose.
<原文结束>

# <翻译开始>
	// 如果`srcValue`和`dstPointer`是相同的类型，直接进行赋值。
	// 为了提高性能。
	// md5:d2daf4894ee19b6e
# <翻译结束>


<原文开始>
	// Example:
	// UploadFile    => UploadFile
	// []UploadFile  => []UploadFile
	// *UploadFile   => *UploadFile
	// *[]UploadFile => *[]UploadFile
	// map           => map
	// []map         => []map
	// *[]map        => *[]map
<原文结束>

# <翻译开始>
	// 示例：
	// 
	// UploadFile    => 上传文件
	// []UploadFile  => 上传文件切片
	// *UploadFile   => 指向上传文件的指针
	// *[]UploadFile => 指向上传文件切片的指针
	// map           => 映射
	// []map         => 映射切片
	// *[]map        => 指向映射切片的指针
	// md5:f6ba941ba3b0269f
# <翻译结束>


<原文开始>
	// Example:
	// UploadFile    => *UploadFile
	// []UploadFile  => *[]UploadFile
	// map           => *map
	// []map         => *[]map
<原文结束>

# <翻译开始>
	// 示例：
	// UploadFile    => *UploadFile
	// []UploadFile  => *[]UploadFile
	// map           => *map
	// []map         => *[]map
	// 
	// 这些注释表示的是Go语言中的指针和数据结构的转换。在Go中，`*`符号用于表示指针类型。这里展示了如何将非指针类型转换为指针类型：
	// 
	// - `UploadFile` 是一个类型，`*UploadFile` 是它的指针类型。
	// - `[]UploadFile` 是 `UploadFile` 类型的切片（数组），`*[]UploadFile` 是这个切片类型的指针。
	// - `map` 表示一个未指定类型的映射，`*map` 则是这个映射类型的指针。
	// - `[]map` 表示一个包含映射的切片，`*[]map` 是这个切片类型的指针。
	// md5:3b743427a52ed67e
# <翻译结束>


<原文开始>
	// Example:
	// *UploadFile    => UploadFile
	// *[]UploadFile  => []UploadFile
	// *map           => map
	// *[]map         => []map
<原文结束>

# <翻译开始>
	// 示例：
	// *UploadFile    => 上传文件
	// *[]UploadFile  => 上传文件的切片
	// *map           => 映射（字典）
	// *[]map         => 映射的切片
	// md5:a787c0f77f0eaa64
# <翻译结束>


<原文开始>
// doConvertWithJsonCheck does json converting check.
// If given `params` is JSON, it then uses json.Unmarshal doing the converting.
<原文结束>

# <翻译开始>
// doConvertWithJsonCheck 做 JSON 转换检查。
// 如果给定的 `params` 是 JSON，那么它会使用 json.Unmarshal 进行转换。
// md5:aa79d041fb48e2db
# <翻译结束>


<原文开始>
// The `params` might be struct that implements interface function Interface, eg: gvar.Var.
<原文结束>

# <翻译开始>
		// `params`可能是一个实现了接口函数Interface的结构体，例如：gvar.Var。 md5:c02e870a76bafdaa
# <翻译结束>

