
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
// 注意，参数 `structSlicePointer` 应该是 *[]struct/*[]*struct 类型。
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
// 使用示例 2：嵌入属性结构体关系：
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
// 示例代码中的 "uid" 是结果表字段名，而 "Uid" 是相关结构体属性名，而不是绑定目标的属性名。
// 在示例代码中，它是 "Entity" 实体的 "User" 的属性名 "Uid"。它会根据给定的 `relation` 参数自动计算 HasOne/HasMany 关系。
//
// 可参考示例或单元测试用例以更清晰地理解此函数的工作方式。
// md5:d6997acc67d472c4
# <翻译结束>


<原文开始>
// Necessary checks for parameters.
<原文结束>

# <翻译开始>
	// 对参数进行必要的检查。 md5:00bddba1a043bfdd
# <翻译结束>


<原文开始>
// Find the element struct type of the slice.
<原文结束>

# <翻译开始>
	// 找到切片的元素结构类型。 md5:a55c378f6fa3b326
# <翻译结束>


<原文开始>
// Find the target field by given name.
<原文结束>

# <翻译开始>
	// 通过给定的名称查找目标字段。 md5:8fe292d32e17dba0
# <翻译结束>


<原文开始>
// Find the attribute struct type for ORM fields filtering.
<原文结束>

# <翻译开始>
	// 查找用于ORM字段过滤的属性结构体类型。 md5:1b98a4f65808a146
# <翻译结束>


<原文开始>
// doScanList converts `result` to struct slice which contains other complex struct attributes recursively.
// The parameter `model` is used for recursively scanning purpose, which means, it can scan the attribute struct/structs recursively,
// but it needs the Model for database accessing.
// Note that the parameter `structSlicePointer` should be type of *[]struct/*[]*struct.
<原文结束>

# <翻译开始>
// doScanList 将 `result` 转换为包含嵌套复杂结构体属性的切片。参数 `model` 用于递归扫描，即它可以递归地扫描结构体/结构体的属性，但需要数据库访问模型。
// 注意参数 `structSlicePointer` 应该是 *[]struct 或 *[]*struct 类型。
// md5:b32c3ddd7d2b8656
# <翻译结束>


<原文开始>
// The pointed slice is not empty.
<原文结束>

# <翻译开始>
		// 指向的切片不为空。 md5:1348d4b6d686b8f3
# <翻译结束>


<原文开始>
			// It here checks if it has struct item, which is already initialized.
			// It then returns error to warn the developer its empty and no conversion.
<原文结束>

# <翻译开始>
			// 这里检查是否具有已初始化的结构体项。
			// 然后返回错误以警告开发者其为空且无法进行转换。
			// md5:cd5f133a393c1157
# <翻译结束>


<原文开始>
// Do nothing for empty struct slice.
<原文结束>

# <翻译开始>
		// 对于空的结构体切片，什么也不做。 md5:f65a6d24cd42ca62
# <翻译结束>


<原文开始>
// Eg: relationKV: id:uid  -> id
<原文结束>

# <翻译开始>
// 例如：relationKV：id：uid -> id. md5:3732472417ccbf22
# <翻译结束>


<原文开始>
// Eg: relationKV: id:uid  -> uid
<原文结束>

# <翻译开始>
// 例如：relationKV：id：uid -> uid. md5:dda263df86dc03a1
# <翻译结束>


<原文开始>
		// The relation key string of table field name and attribute name
		// can be joined with char '=' or ':'.
<原文结束>

# <翻译开始>
		// 表字段名与属性名之间的关联键字符串
		// 可以使用字符'='或':'进行连接。
		// md5:a3dd08343df8a7ac
# <翻译结束>


<原文开始>
// Compatible with old splitting char ':'.
<原文结束>

# <翻译开始>
			// 与旧的分隔字符':'兼容。 md5:21a764d3ea1e081b
# <翻译结束>


<原文开始>
// The relation names are the same.
<原文结束>

# <翻译开始>
			// 关系名称是相同的。 md5:1075b6495b26357b
# <翻译结束>


<原文开始>
			// Defined table field to relation attribute name.
			// Like:
			// uid:Uid
			// uid:UserId
<原文结束>

# <翻译开始>
			// 定义表字段到关系属性名。
			// 例如：
			// uid:Uid
			// uid:UserId
			// md5:029253159bee75d1
# <翻译结束>


<原文开始>
// Note that the value might be type of slice.
<原文结束>

# <翻译开始>
			// 请注意，该值可能是切片类型。 md5:079de568e97881a6
# <翻译结束>


<原文开始>
// Bind to target attribute.
<原文结束>

# <翻译开始>
	// 将其绑定到目标属性。 md5:6248a034de9b08e4
# <翻译结束>


<原文开始>
// Bind to relation conditions.
<原文结束>

# <翻译开始>
	// 绑定关系条件。 md5:1d13e1ebe0b47bd2
# <翻译结束>


<原文开始>
// The FieldByName should be called on non-pointer reflect.Value.
<原文结束>

# <翻译开始>
		// 应该在非指针的reflect.Value上调用FieldByName。 md5:1343ff0ec0419e1f
# <翻译结束>


<原文开始>
				// The element is nil, then create one and set it to the slice.
				// The "reflect.New(itemType.Elem())" creates a new element and returns the address of it.
				// For example:
				// reflect.New(itemType.Elem())        => *Entity
				// reflect.New(itemType.Elem()).Elem() => Entity
<原文结束>

# <翻译开始>
				// 如果元素为nil，则创建一个并将其设置到切片中。
				// "reflect.New(itemType.Elem())" 用于创建一个新的元素，并返回该元素的地址。
				// 例如：
				// reflect.New(itemType.Elem())        => *实体
				// reflect.New(itemType.Elem()).Elem() => 实体
				// md5:0897d7c0e7467f9d
# <翻译结束>


<原文开始>
// Attribute value of current slice element.
<原文结束>

# <翻译开始>
			// 当前切片元素的属性值。 md5:b46440a93bb1ddaa
# <翻译结束>


<原文开始>
// Check and find possible bind to attribute name.
<原文结束>

# <翻译开始>
		// 检查并尝试找到可能与属性名绑定的位置。 md5:b1e1f2121b3b5f92
# <翻译结束>


<原文开始>
// Maybe the attribute does not exist yet.
<原文结束>

# <翻译开始>
					// 可能属性还不存在。 md5:d7992076e8a1e5fe
# <翻译结束>


<原文开始>
// There's no relational data.
<原文结束>

# <翻译开始>
					// 没有关联数据。 md5:4f76ca1525fb5005
# <翻译结束>

