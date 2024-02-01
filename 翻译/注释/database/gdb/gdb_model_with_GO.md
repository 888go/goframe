
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
// With 启用关联查询，通过给定的属性对象指定开启。
// 常考"模型关联-静态关联"文档:https://goframe.org/pages/viewpage.action?pageId=7297190
// 例如，如果给定如下的结构体定义：
//
//	type User struct {
//		 gmeta.Meta `orm:"table:user"` // 定义表名为 user
//		 Id         int           `json:"id"`    // 用户ID
//		 Name       string        `json:"name"`   // 用户名
//		 UserDetail *UserDetail   `orm:"with:uid=id"` // 关联 UserDetail 表，通过 uid 等于 id 进行关联
//		 UserScores []*UserScores `orm:"with:uid=id"` // 关联 UserScores 表，通过 uid 等于 id 进行关联
//	}
//
// 我们可以通过以下方式在属性 `UserDetail` 和 `UserScores` 上启用模型关联操作：
// db.With(User{}.UserDetail).With(User{}.UserScores).Scan(xxx)
// 或者：
// db.With(UserDetail{}).With(UserScores{}).Scan(xxx)
// 或者：
// db.With(UserDetail{}, UserScores{}).Scan(xxx)
# <翻译结束>


<原文开始>
// WithAll enables model association operations on all objects that have "with" tag in the struct.
<原文结束>

# <翻译开始>
// WithAll 开启在所有具有"struct"标签中包含"with"标签的对象上的模型关联操作。
// 常考"模型关联-静态关联"文档:https://goframe.org/pages/viewpage.action?pageId=7297190
# <翻译结束>


<原文开始>
// doWithScanStruct handles model association operations feature for single struct.
<原文结束>

# <翻译开始>
// doWithScanStruct 处理单个结构体的模型关联操作功能。
# <翻译结束>


<原文开始>
// It checks the with array and automatically calls the ScanList to complete association querying.
<原文结束>

# <翻译开始>
// 它会检查with数组，并自动调用ScanList完成关联查询。
# <翻译结束>


<原文开始>
// It does select operation if the field type is in the specified "with" type array.
<原文结束>

# <翻译开始>
// 如果字段类型在指定的“with”类型数组中，则进行选择操作。
# <翻译结束>


<原文开始>
// It just handlers "with" type attribute struct, so it ignores other struct types.
<原文结束>

# <翻译开始>
// 它仅处理“with”类型属性的结构体，因此会忽略其他类型的结构体。
# <翻译结束>


<原文开始>
			// It also supports using only one column name
			// if both tables associates using the same column name.
<原文结束>

# <翻译开始>
// 它还支持仅使用一个列名
// 如果两个表关联时使用相同的列名。
# <翻译结束>


<原文开始>
// Find the value of related attribute from `pointer`.
<原文结束>

# <翻译开始>
// 从`pointer`中查找相关属性的值。
# <翻译结束>


<原文开始>
// It automatically retrieves struct field names from current attribute struct/slice.
<原文结束>

# <翻译开始>
// 它会自动从当前属性结构体/切片中检索结构体字段名称。
# <翻译结束>


<原文开始>
// Recursively with feature checks.
<原文结束>

# <翻译开始>
// 递归并进行特性检查
# <翻译结束>







<原文开始>
// It ignores sql.ErrNoRows in with feature.
<原文结束>

# <翻译开始>
// 它在特性中忽略 sql.ErrNoRows 错误。
# <翻译结束>


<原文开始>
// doWithScanStructs handles model association operations feature for struct slice.
// Also see doWithScanStruct.
<原文结束>

# <翻译开始>
// doWithScanStructs 处理结构体切片的模型关联操作特性。
// 也可参考 doWithScanStruct。
# <翻译结束>


<原文开始>
// It does select operation if the field type is in the specified with type array.
<原文结束>

# <翻译开始>
// 如果字段类型在指定的类型数组中，它将执行选择操作。
# <翻译结束>


<原文开始>
			// It supports using only one column name
			// if both tables associates using the same column name.
<原文结束>

# <翻译开始>
// 如果两个表使用相同的列名关联，则它支持仅使用一个列名。
# <翻译结束>


<原文开始>
// Find the value slice of related attribute from `pointer`.
<原文结束>

# <翻译开始>
// 从`pointer`中找到相关属性的值切片。
# <翻译结束>


<原文开始>
// If related value is empty, it does nothing but just returns.
<原文结束>

# <翻译开始>
// 如果相关值为空，则此函数不做任何操作，仅返回。
# <翻译结束>


<原文开始>
// With cache feature.
<原文结束>

# <翻译开始>
// 带有缓存功能。
# <翻译结束>

