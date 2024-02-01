
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
// ScanList converts `r` to struct slice which contains other complex struct attributes.
// Note that the parameter `structSlicePointer` should be type of *[]struct/*[]*struct.
//
// Usage example 1: Normal attribute struct relation:
//
//	type EntityUser struct {
//		   Uid  int
//		   Name string
//	}
//
//	type EntityUserDetail struct {
//		   Uid     int
//		   Address string
//	}
//
//	type EntityUserScores struct {
//		   Id     int
//		   Uid    int
//		   Score  int
//		   Course string
//	}
//
//	type Entity struct {
//	    User       *EntityUser
//		   UserDetail *EntityUserDetail
//		   UserScores []*EntityUserScores
//	}
//
// var users []*Entity
// ScanList(&users, "User")
// ScanList(&users, "User", "uid")
// ScanList(&users, "UserDetail", "User", "uid:Uid")
// ScanList(&users, "UserScores", "User", "uid:Uid")
// ScanList(&users, "UserScores", "User", "uid")
//
// Usage example 2: Embedded attribute struct relation:
//
//	type EntityUser struct {
//		   Uid  int
//		   Name string
//	}
//
//	type EntityUserDetail struct {
//		   Uid     int
//		   Address string
//	}
//
//	type EntityUserScores struct {
//		   Id    int
//		   Uid   int
//		   Score int
//	}
//
//	type Entity struct {
//		   EntityUser
//		   UserDetail EntityUserDetail
//		   UserScores []EntityUserScores
//	}
//
// var users []*Entity
// ScanList(&users)
// ScanList(&users, "UserDetail", "uid")
// ScanList(&users, "UserScores", "uid")
//
// The parameters "User/UserDetail/UserScores" in the example codes specify the target attribute struct
// that current result will be bound to.
//
// The "uid" in the example codes is the table field name of the result, and the "Uid" is the relational
// struct attribute name - not the attribute name of the bound to target. In the example codes, it's attribute
// name "Uid" of "User" of entity "Entity". It automatically calculates the HasOne/HasMany relationship with
// given `relation` parameter.
//
// See the example or unit testing cases for clear understanding for this function.
<原文结束>

# <翻译开始>
// ScanList 将 `r` 转换为包含其他复杂结构体属性的结构体切片。
// 注意，参数 `structSlicePointer` 应该是 *[]struct 或 *[]*struct 类型。
// 参考关联模型: https://goframe.org/pages/viewpage.action?pageId=1114326
//
// 使用示例 1：普通属性结构体关系：
//
//	type EntityUser struct {
//		Uid  int
//		Name string
//	}
//
//	type EntityUserDetail struct {
//		Uid     int
//		Address string
//	}
//
//	type EntityUserScores struct {
//		Id     int
//		Uid    int
//		Score  int
//		Course string
//	}
//
//	type Entity struct {
//		User       *EntityUser
//		UserDetail *EntityUserDetail
//		UserScores []*EntityUserScores
//	}
//
//	var users []*Entity
//	ScanList(&users, "User")
//	ScanList(&users, "User", "uid")
//	ScanList(&users, "UserDetail", "User", "uid:Uid")
//	ScanList(&users, "UserScores", "User", "uid:Uid")
//	ScanList(&users, "UserScores", "User", "uid")
//
// 使用示例 2：嵌入式属性结构体关系：
//
//	type EntityUser struct {
//		Uid  int
//		Name string
//	}
//
//	type EntityUserDetail struct {
//		Uid     int
//		Address string
//	}
//
//	type EntityUserScores struct {
//		Id    int
//		Uid   int
//		Score int
//	}
//
//	type Entity struct {
//		EntityUser
//		UserDetail EntityUserDetail
//		UserScores []EntityUserScores
//	}
//
//	var users []*Entity
//	ScanList(&users)
//	ScanList(&users, "UserDetail", "uid")
//	ScanList(&users, "UserScores", "uid")
//
// 示例代码中的 "User/UserDetail/UserScores" 参数指定了当前结果将绑定的目标属性结构体。
//
// 示例代码中的 "uid" 是结果表字段名，而 "Uid" 是关联结构体属性名——而不是目标绑定的属性名。在示例中，它是实体 "Entity" 中 "User" 的属性名 "Uid"。它会根据给定的 `relation` 参数自动计算 HasOne/HasMany 关系。
//
// 为了清晰理解此函数，请参考示例或单元测试用例。
# <翻译结束>


<原文开始>
// Necessary checks for parameters.
<原文结束>

# <翻译开始>
// 对参数进行必要的检查。
# <翻译结束>


<原文开始>
// Find the element struct type of the slice.
<原文结束>

# <翻译开始>
// 查找切片的元素结构体类型。
# <翻译结束>


<原文开始>
// Find the target field by given name.
<原文结束>

# <翻译开始>
// 根据给定名称查找目标字段。
# <翻译结束>


<原文开始>
// Find the attribute struct type for ORM fields filtering.
<原文结束>

# <翻译开始>
// 查找用于ORM字段过滤的属性结构体类型。
# <翻译结束>


<原文开始>
// doScanList converts `result` to struct slice which contains other complex struct attributes recursively.
// The parameter `model` is used for recursively scanning purpose, which means, it can scan the attribute struct/structs recursively,
// but it needs the Model for database accessing.
// Note that the parameter `structSlicePointer` should be type of *[]struct/*[]*struct.
<原文结束>

# <翻译开始>
// doScanList 将 `result` 转换为包含其他复杂结构体属性的结构体切片，并且这个转换过程是递归的。
// 参数 `model` 用于递归扫描目的，这意味着它可以递归地扫描结构体/结构体数组，
// 但需要 Model 以便进行数据库访问。
// 注意参数 `structSlicePointer` 的类型应为 *[]struct 或 *[]*struct。
# <翻译结束>


<原文开始>
// The pointed slice is not empty.
<原文结束>

# <翻译开始>
// 指向的切片不为空。
# <翻译结束>


<原文开始>
			// It here checks if it has struct item, which is already initialized.
			// It then returns error to warn the developer its empty and no conversion.
<原文结束>

# <翻译开始>
// 在这里检查它是否包含已初始化的结构体项。
// 如果为空且无法进行转换，则返回错误以警告开发者。
# <翻译结束>


<原文开始>
// Do nothing for empty struct slice.
<原文结束>

# <翻译开始>
// 对于空的结构体切片，不执行任何操作。
# <翻译结束>












<原文开始>
// Eg: relationKV: id:uid  -> id
<原文结束>

# <翻译开始>
// 示例：relationKV: id:uid -> id
// 这个注释表明了一个键值对的示例，其中关系（relationKV）的键是"id:uid"，对应的值为"id"。在实际应用中，这可能表示一个映射关系，通过用户ID(uid)可以找到对应的ID(id)。
# <翻译结束>


<原文开始>
// Eg: relationKV: id:uid  -> uid
<原文结束>

# <翻译开始>
// 示例：relationKV: id:uid  -> uid
// （注释翻译：这个字段或者变量表示一种键值对关系，其中键是"id:uid"，值是"uid"）
# <翻译结束>


<原文开始>
		// The relation key string of table field name and attribute name
		// can be joined with char '=' or ':'.
<原文结束>

# <翻译开始>
// 表字段名与属性名之间的关联键字符串，可以使用字符'='或':'连接。
# <翻译结束>


<原文开始>
// Compatible with old splitting char ':'.
<原文结束>

# <翻译开始>
// 与旧的分隔符 ':' 兼容。
# <翻译结束>


<原文开始>
// The relation names are the same.
<原文结束>

# <翻译开始>
// 关系名称是相同的。
# <翻译结束>


<原文开始>
			// Defined table field to relation attribute name.
			// Like:
			// uid:Uid
			// uid:UserId
<原文结束>

# <翻译开始>
// 定义表格字段到关联属性名称的映射。
// 例如：
// uid:Uid
// uid:UserId
// 这段代码的作用是将数据库表中的字段名（如uid）映射到程序中使用的关联属性名称，以实现字段名称在代码逻辑中的语义化表达。
# <翻译结束>


<原文开始>
// Note that the value might be type of slice.
<原文结束>

# <翻译开始>
// 注意，该值可能是切片类型。
# <翻译结束>







<原文开始>
// Bind to relation conditions.
<原文结束>

# <翻译开始>
// 绑定到关联条件。
# <翻译结束>


<原文开始>
// The FieldByName should be called on non-pointer reflect.Value.
<原文结束>

# <翻译开始>
// 应在非指针 reflect.Value 上调用 FieldByName。
# <翻译结束>







<原文开始>
				// The element is nil, then create one and set it to the slice.
				// The "reflect.New(itemType.Elem())" creates a new element and returns the address of it.
				// For example:
				// reflect.New(itemType.Elem())        => *Entity
				// reflect.New(itemType.Elem()).Elem() => Entity
<原文结束>

# <翻译开始>
// 如果元素为nil，则创建一个新元素并将其设置到切片中。
// "reflect.New(itemType.Elem())" 用于创建一个新的元素，并返回该元素的地址。
// 例如：
// reflect.New(itemType.Elem())        // => *Entity （返回指向新创建实体类型的指针）
// reflect.New(itemType.Elem()).Elem() // => Entity （获取新创建实体类型的值）
# <翻译结束>


<原文开始>
// Attribute value of current slice element.
<原文结束>

# <翻译开始>
// 当前切片元素的属性值。
# <翻译结束>







<原文开始>
// Check and find possible bind to attribute name.
<原文结束>

# <翻译开始>
// 检查并查找可能绑定到属性名称的地方。
# <翻译结束>







<原文开始>
// Maybe the attribute does not exist yet.
<原文结束>

# <翻译开始>
// 可能该属性尚不存在。
# <翻译结束>


<原文开始>
// There's no relational data.
<原文结束>

# <翻译开始>
// 没有关联数据。
# <翻译结束>


<原文开始>
// Slice element item.
<原文结束>

# <翻译开始>
// 切片元素项。
# <翻译结束>


<原文开始>
// Relation variables.
<原文结束>

# <翻译开始>
// 关系变量。
# <翻译结束>


<原文开始>
// Bind to target attribute.
<原文结束>

# <翻译开始>
// 绑定到目标属性。
# <翻译结束>


<原文开始>
// Current slice element.
<原文结束>

# <翻译开始>
// 当前切片元素。
# <翻译结束>

