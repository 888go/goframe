
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
// With creates and returns an ORM model based on metadata of given object.
// It also enables model association operations feature on given `object`.
// It can be called multiple times to add one or more objects to model and enable
// their mode association operations feature.
// For example, if given struct definition:
//
//	type User struct {
//		 gmeta.Meta `orm:"table:user"`
//		 Id         int           `json:"id"`
//		 Name       string        `json:"name"`
//		 UserDetail *UserDetail   `orm:"with:uid=id"`
//		 UserScores []*UserScores `orm:"with:uid=id"`
//	}
//
// We can enable model association operations on attribute `UserDetail` and `UserScores` by:
//
//	db.With(User{}.UserDetail).With(User{}.UserScores).Scan(xxx)
//
// Or:
//
//	db.With(UserDetail{}).With(UserScores{}).Scan(xxx)
//
// Or:
//
//	db.With(UserDetail{}, UserScores{}).Scan(xxx)
<原文结束>

# <翻译开始>
// With 创建并返回一个基于给定对象元数据的 ORM 模型。它还为给定的 `object` 启用模型关联操作功能。
// 可以多次调用此函数，以向模型中添加一个或多个对象，并启用它们的模式关联操作功能。
// 例如，如果给出的结构体定义如下：
// 
// ```
// type User struct {
//     gmeta.Meta `orm:"table:user"`
//     Id         int           `json:"id"`
//     Name       string        `json:"name"`
//     UserDetail *UserDetail   `orm:"with:uid=id"`
//     UserScores []*UserScores `orm:"with:uid=id"`
// }
// ```
// 
// 我们可以通过以下方式在 `UserDetail` 和 `UserScores` 属性上启用模型关联操作：
// 
// ```
// db.With(User{}.UserDetail).With(User{}.UserScores).Scan(xxx)
// ```
// 
// 或者：
// 
// ```
// db.With(UserDetail{}).With(UserScores{}).Scan(xxx)
// ```
// 
// 或者：
// 
// ```
// db.With(UserDetail{}, UserScores{}).Scan(xxx)
// ```
// md5:c9498702475d54a9
# <翻译结束>


<原文开始>
// WithAll enables model association operations on all objects that have "with" tag in the struct.
<原文结束>

# <翻译开始>
// WithAll 启用对结构体中带有 "with" 标签的所有对象进行模型关联操作。 md5:83d3591315f0add0
# <翻译结束>


<原文开始>
// doWithScanStruct handles model association operations feature for single struct.
<原文结束>

# <翻译开始>
// doWithScanStruct 处理单个结构体的模型关联操作功能。 md5:64dcc9bfd0382aa8
# <翻译结束>


<原文开始>
// It checks the with array and automatically calls the ScanList to complete association querying.
<原文结束>

# <翻译开始>
// 它会检查with数组，并自动调用ScanList来完成关联查询。 md5:cb83f16b7131ad65
# <翻译结束>


<原文开始>
// It does select operation if the field type is in the specified "with" type array.
<原文结束>

# <翻译开始>
// 如果字段类型在指定的"with"类型数组中，它会执行选择操作。 md5:b425357c98d952c8
# <翻译结束>


<原文开始>
// It just handlers "with" type attribute struct, so it ignores other struct types.
<原文结束>

# <翻译开始>
// 它仅处理带有"type"属性的"with"类型结构体，因此会忽略其他类型的结构体。 md5:c1f385406b699f00
# <翻译结束>


<原文开始>
			// It also supports using only one column name
			// if both tables associates using the same column name.
<原文结束>

# <翻译开始>
			// 它还支持仅使用一个列名
			// 如果两个表使用相同的列名进行关联。
			// md5:c924339d8b4eddbc
# <翻译结束>


<原文开始>
// Find the value of related attribute from `pointer`.
<原文结束>

# <翻译开始>
// 从`pointer`中找到相关的属性值。 md5:b2da611599aed2d2
# <翻译结束>


<原文开始>
// It automatically retrieves struct field names from current attribute struct/slice.
<原文结束>

# <翻译开始>
// 它会自动从当前属性结构体/切片中获取字段名。 md5:09af2856a6801ffd
# <翻译结束>


<原文开始>
// Recursively with feature checks.
<原文结束>

# <翻译开始>
// 递归实现并带有特性检查。 md5:9ddeb46ca8a2b86d
# <翻译结束>


<原文开始>
// It ignores sql.ErrNoRows in with feature.
<原文结束>

# <翻译开始>
// 它在该特性中忽略sql.ErrNoRows错误。 md5:4b82d692c0646927
# <翻译结束>


<原文开始>
// doWithScanStructs handles model association operations feature for struct slice.
// Also see doWithScanStruct.
<原文结束>

# <翻译开始>
// doWithScanStructs 处理结构切片的模型关联操作功能。
// 参见 doWithScanStruct。
// md5:6219b8feabf0e7d9
# <翻译结束>


<原文开始>
// It does select operation if the field type is in the specified with type array.
<原文结束>

# <翻译开始>
// 如果字段类型在指定的数组类型中，它将执行选择操作。 md5:afefe105662c6d79
# <翻译结束>


<原文开始>
			// It supports using only one column name
			// if both tables associates using the same column name.
<原文结束>

# <翻译开始>
			// 它支持仅使用一个列名的情况，
			// 当两个表通过相同的列名关联时。
			// md5:18222f22ecbee1ef
# <翻译结束>


<原文开始>
// Find the value slice of related attribute from `pointer`.
<原文结束>

# <翻译开始>
// 从`pointer`中查找相关属性的值切片。 md5:e729db1e29dfb929
# <翻译结束>


<原文开始>
// If related value is empty, it does nothing but just returns.
<原文结束>

# <翻译开始>
// 如果相关值为空，它什么也不做，只是返回。 md5:e4acb6a4c5d73f8f
# <翻译结束>

