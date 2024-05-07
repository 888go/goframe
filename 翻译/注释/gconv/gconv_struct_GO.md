
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
// Struct maps the params key-value pairs to the corresponding struct object's attributes.
// The third parameter `mapping` is unnecessary, indicating the mapping rules between the
// custom key name and the attribute name(case-sensitive).
//
// Note:
//  1. The `params` can be any type of map/struct, usually a map.
//  2. The `pointer` should be type of *struct/**struct, which is a pointer to struct object
//     or struct pointer.
//  3. Only the public attributes of struct object can be mapped.
//  4. If `params` is a map, the key of the map `params` can be lowercase.
//     It will automatically convert the first letter of the key to uppercase
//     in mapping procedure to do the matching.
//     It ignores the map key, if it does not match.
<原文结束>

# <翻译开始>
// Struct 将参数键值对映射到相应结构体对象的属性上。
// 第三个参数 `mapping` 是可选的，表示自定义键名与属性名（大小写敏感）之间的映射规则。
// 注意：
// 1. `params` 可以是任何类型的 map 或 struct，通常为 map 类型。
// 2. `pointer` 应为 *struct/**struct 类型，即指向结构体对象或结构体指针的指针。
// 3. 只有结构体对象的公共属性可以被映射。
// 4. 如果 `params` 是一个 map，map 的键可以是小写的。
//    在映射过程中，会自动将键的首字母转换为大写进行匹配。
//    如果不匹配，则忽略该 map 键。
# <翻译结束>


<原文开始>
// StructTag acts as Struct but also with support for priority tag feature, which retrieves the
// specified tags for `params` key-value items to struct attribute names mapping.
// The parameter `priorityTag` supports multiple tags that can be joined with char ','.
<原文结束>

# <翻译开始>
// StructTag 结构体在 Struct 的基础上增加了支持优先级标签功能，该功能用于获取 `params` 键值对中指定的标签，并映射到结构体属性名称。  
// 参数 `priorityTag` 支持多个标签，多个标签之间可以通过字符 ',' 连接。
# <翻译结束>


<原文开始>
// doStructWithJsonCheck checks if given `params` is JSON, it then uses json.Unmarshal doing the converting.
<原文结束>

# <翻译开始>
// doStructWithJsonCheck 检查给定的 `params` 是否为 JSON 格式，如果是，则使用 json.Unmarshal 进行转换。
# <翻译结束>


<原文开始>
// The `params` might be struct that implements interface function Interface, eg: gvar.Var.
<原文结束>

# <翻译开始>
// `params` 可能是一个实现了 Interface 接口函数的结构体，例如：gvar.Var。
# <翻译结束>


<原文开始>
// doStruct is the core internal converting function for any data to struct.
<原文结束>

# <翻译开始>
// doStruct 是用于将任何数据转换为结构体的核心内部函数。
# <翻译结束>


<原文开始>
// If `params` is nil, no conversion.
<原文结束>

# <翻译开始>
// 如果`params`为nil，则不进行转换。
# <翻译结束>


<原文开始>
// Catch the panic, especially the reflection operation panics.
<原文结束>

# <翻译开始>
// 捕获 panic，特别是反射操作引发的 panic。
# <翻译结束>












<原文开始>
// Using IsNil on reflect.Ptr variable is OK.
<原文结束>

# <翻译开始>
// 对reflect.Ptr类型的变量使用IsNil方法是可行的。
# <翻译结束>







<原文开始>
	// If `params` and `pointer` are the same type, the do directly assignment.
	// For performance enhancement purpose.
<原文结束>

# <翻译开始>
// 如果`params`和`pointer`是相同类型，则直接进行赋值操作。
// 为了提升性能。
# <翻译结束>


<原文开始>
		// Eg:
		// UploadFile  => UploadFile
		// *UploadFile => *UploadFile
<原文结束>

# <翻译开始>
// 例如：
// UploadFile  => 上传文件
// *UploadFile => 指向UploadFile的指针
# <翻译结束>


<原文开始>
		// Eg:
		// UploadFile  => *UploadFile
<原文结束>

# <翻译开始>
// 例如：
// UploadFile  => *UploadFile
// （此代码注释翻译为：）
// 示例：
// UploadFile  => 指向UploadFile类型的指针
# <翻译结束>


<原文开始>
		// Eg:
		// *UploadFile  => UploadFile
<原文结束>

# <翻译开始>
// 示例：
// *UploadFile  => 上传文件
# <翻译结束>


<原文开始>
// Normal unmarshalling interfaces checks.
<原文结束>

# <翻译开始>
// 正常的接口解码检查。
# <翻译结束>


<原文开始>
	// It automatically creates struct object if necessary.
	// For example, if `pointer` is **User, then `elem` is *User, which is a pointer to User.
<原文结束>

# <翻译开始>
// 如果有必要，它会自动创建结构体对象。
// 例如，如果 `pointer` 是 **User 类型，那么 `elem` 就是 *User 类型，即指向 User 的指针。
# <翻译结束>


<原文开始>
// If it is converted failed, it reset the `pointer` to nil.
<原文结束>

# <翻译开始>
// 如果转换失败，则将`pointer`重置为nil。
# <翻译结束>


<原文开始>
		// if v, ok := pointerElemReflectValue.Interface().(iUnmarshalValue); ok {
		//	return v.UnmarshalValue(params)
		// }
		// Note that it's `pointerElemReflectValue` here not `pointerReflectValue`.
<原文结束>

# <翻译开始>
// 如果v, ok := 将pointerElemReflectValue.Interface().(iUnmarshalValue)进行类型断言并赋值；如果ok为真（即转换成功），
// 则返回v.UnmarshalValue(params)的结果
// 注意：这里使用的是`pointerElemReflectValue`而非`pointerReflectValue`
// 这段代码的中文注释翻译如下：
// ```go
// 若能将pointerElemReflectValue.Interface()转换为iUnmarshalValue类型，并将转换结果赋值给v和ok，且ok为真，
// 则调用v的UnmarshalValue方法处理params并返回其结果。
// 需要注意的是，此处使用的是`pointerElemReflectValue`变量，而不是`pointerReflectValue`变量。
# <翻译结束>


<原文开始>
// Retrieve its element, may be struct at last.
<原文结束>

# <翻译开始>
// 获取其元素，最后可能是一个结构体。
# <翻译结束>


<原文开始>
	// paramsMap is the map[string]interface{} type variable for params.
	// DO NOT use MapDeep here.
<原文结束>

# <翻译开始>
// paramsMap 是用于参数的 map[string]interface{} 类型变量。
// 在此处不要使用 MapDeep。
# <翻译结束>


<原文开始>
// Nothing to be done as the parameters are empty.
<原文结束>

# <翻译开始>
// 由于参数为空，无需执行任何操作。
# <翻译结束>


<原文开始>
	// It only performs one converting to the same attribute.
	// doneMap is used to check repeated converting, its key is the real attribute name
	// of the struct.
<原文结束>

# <翻译开始>
// 它只执行同一属性的一次转换。
// doneMap 用于检查重复转换，其键是结构体的实际属性名称。
# <翻译结束>


<原文开始>
	// The key of the attrMap is the attribute name of the struct,
	// and the value is its replaced name for later comparison to improve performance.
<原文结束>

# <翻译开始>
// attrMap 的键是结构体的属性名，
// 值是用于后续比较时的替换名称，目的是为了提升性能。
# <翻译结束>


<原文开始>
		// Attribute name to its symbols-removed name,
		// in order to quick index and comparison in following logic.
<原文结束>

# <翻译开始>
// 将属性名称映射到去除符号后的名称，
// 以便在后续逻辑中快速索引和比较。
# <翻译结束>


<原文开始>
// Only do converting to public attributes.
<原文结束>

# <翻译开始>
// 只对公开属性进行转换
# <翻译结束>


<原文开始>
// Maybe it's struct/*struct embedded.
<原文结束>

# <翻译开始>
// 可能这是一个结构体 /* 或是结构体嵌入.
# <翻译结束>


<原文开始>
// Ignore the interface attribute if it's nil.
<原文结束>

# <翻译开始>
// 如果接口属性为nil，则忽略它。
# <翻译结束>


<原文开始>
	// The key of the `attrToTagCheckNameMap` is the attribute name of the struct,
	// and the value is its replaced tag name for later comparison to improve performance.
<原文结束>

# <翻译开始>
// `attrToTagCheckNameMap` 的键是结构体的属性名称，
// 而值则是用于后续比较时替换的标签名称，目的是为了提高性能。
# <翻译结束>


<原文开始>
		// If there's something else in the tag string,
		// it uses the first part which is split using char ','.
		// Eg:
		// orm:"id, priority"
		// orm:"name, with:uid=id"
<原文结束>

# <翻译开始>
// 如果标签字符串中还有其他内容，
// 它会使用通过逗号（,）分割得到的第一部分。
// 例如：
// orm:"id, priority" // 使用id和priority
// orm:"name, with:uid=id" // 使用name和with:uid=id中的name部分
# <翻译结束>


<原文开始>
		// If tag and attribute values both exist in `paramsMap`,
		// it then uses the tag value overwriting the attribute value in `paramsMap`.
<原文结束>

# <翻译开始>
// 如果tag和attribute值同时存在于`paramsMap`中，
// 则优先使用tag值，并在`paramsMap`中覆盖原有的attribute值。
# <翻译结束>


<原文开始>
// To convert value base on custom parameter key to attribute name map.
<原文结束>

# <翻译开始>
// 根据自定义参数键到属性名映射来转换值。
# <翻译结束>


<原文开始>
// Already done all attributes value assignment nothing to do next.
<原文结束>

# <翻译开始>
// 已经完成了所有属性值的赋值，接下来无事可做。
# <翻译结束>


<原文开始>
// To convert value base on precise attribute name.
<原文结束>

# <翻译开始>
// 根据精确属性名称转换值。
# <翻译结束>


<原文开始>
// To convert value base on parameter map.
<原文结束>

# <翻译开始>
// 根据参数映射转换值。
# <翻译结束>


<原文开始>
// If the attribute name is already checked converting, then skip it.
<原文结束>

# <翻译开始>
// 如果属性名称已经经过转换检查，那么跳过它。
# <翻译结束>


<原文开始>
// It ignores the attribute names if it is specified in the `paramKeyToAttrMap`.
<原文结束>

# <翻译开始>
// 如果属性名在`paramKeyToAttrMap`中指定，它将忽略这些属性名。
# <翻译结束>


<原文开始>
// The value by precise attribute name.
<原文结束>

# <翻译开始>
// 通过精确属性名称获取的值。
# <翻译结束>


<原文开始>
// If the attribute name is in custom paramKeyToAttrMap, it then ignores this converting.
<原文结束>

# <翻译开始>
// 如果属性名存在于自定义的paramKeyToAttrMap中，则忽略该转换操作。
# <翻译结束>


<原文开始>
		// It firstly considers `paramName` as accurate tag name,
		// and retrieve attribute name from `tagToAttrNameMap` .
<原文结束>

# <翻译开始>
// 它首先将`paramName`视为准确的标签名称，
// 然后从`tagToAttrNameMap`中检索属性名称。
# <翻译结束>


<原文开始>
			// Loop to find the matched attribute name with or without
			// string cases and chars like '-'/'_'/'.'/' '.
<原文结束>

# <翻译开始>
// 循环查找匹配的属性名，支持大小写不敏感以及包含'-'/'_'/'.'/' '等字符的情况
# <翻译结束>


<原文开始>
			// Matching the parameters to struct tag names.
			// The `attrKey` is the attribute name of the struct.
<原文结束>

# <翻译开始>
// 将参数与结构体标签名称进行匹配。
// `attrKey` 是该结构体的属性名称。
# <翻译结束>


<原文开始>
// Matching the parameters to struct attributes.
<原文结束>

# <翻译开始>
// 将参数与结构体属性进行匹配。
# <翻译结束>


<原文开始>
				// Eg:
				// UserName  eq user_name
				// User-Name eq username
				// username  eq userName
				// etc.
<原文结束>

# <翻译开始>
// 示例：
// UserName 等价于 user_name
// User-Name 等价于 username
// username 等价于 userName
// 等等。
// 这段Go代码注释描述了不同形式的字符串表示，它们在某种上下文中被视为等价：
// - `UserName` 和 `user_name` 是等价的；
// - `User-Name` 和 `username` 也是等价的；
// - `username` 和 `userName` 同样视为等价。
// 这通常出现在将驼峰命名（camelCase）和下划线命名（snake_case）互相转换的场景中。
# <翻译结束>


<原文开始>
// No matching, it gives up this attribute converting.
<原文结束>

# <翻译开始>
// 没有找到匹配项，放弃该属性的转换。
# <翻译结束>


<原文开始>
// bindVarToStructAttr sets value to struct object attribute by name.
<原文结束>

# <翻译开始>
// bindVarToStructAttr 通过名称将值设置到结构体对象的属性中。
# <翻译结束>


<原文开始>
// CanSet checks whether attribute is public accessible.
<原文结束>

# <翻译开始>
// CanSet 检查属性是否可公开访问。
# <翻译结束>







<原文开始>
		// Try to call custom converter.
		// Issue: https://github.com/gogf/gf/issues/3099
<原文结束>

# <翻译开始>
// 尝试调用自定义转换器。
// 问题：https://github.com/gogf/gf/issues/3099
# <翻译结束>


<原文开始>
		// Special handling for certain types:
		// - Overwrite the default type converting logic of stdlib for time.Time/*time.Time.
<原文结束>

# <翻译开始>
// 特殊处理某些类型：
// - 覆盖标准库中time.Time类型的默认转换逻辑。
# <翻译结束>


<原文开始>
		// Hold the time zone consistent in recursive
		// Issue: https://github.com/gogf/gf/issues/2980
<原文结束>

# <翻译开始>
// 在递归中保持时区的一致性
// 问题：https://github.com/gogf/gf/issues/2980
# <翻译结束>












<原文开始>
// bindVarToReflectValueWithInterfaceCheck does bind using common interfaces checks.
<原文结束>

# <翻译开始>
// bindVarToReflectValueWithInterfaceCheck 通过通用接口检查进行绑定。
# <翻译结束>


<原文开始>
// Not a pointer, but can token address, that makes it can be unmarshalled.
<原文结束>

# <翻译开始>
// 不是指针类型，但可以获取其地址，因此它可以被反序列化。
# <翻译结束>


<原文开始>
// If it is not a valid JSON string, it then adds char `"` on its both sides to make it is.
<原文结束>

# <翻译开始>
// 如果这不是一个有效的JSON字符串，那么会在其两边添加字符`"`以使其成为有效的JSON字符串。
# <翻译结束>


<原文开始>
// bindVarToReflectValue sets `value` to reflect value object `structFieldValue`.
<原文结束>

# <翻译开始>
// bindVarToReflectValue 将 `value` 绑定到反射值对象 `structFieldValue`。
# <翻译结束>


<原文开始>
// Converting using `Set` interface implements, for some types.
<原文结束>

# <翻译开始>
// 通过实现`Set`接口进行转换，针对某些类型。
# <翻译结束>


<原文开始>
// Converting using reflection by kind.
<原文结束>

# <翻译开始>
// 通过 kind 使用反射进行转换。
# <翻译结束>


<原文开始>
// Recursively converting for struct attribute.
<原文结束>

# <翻译开始>
// 递归地转换结构体属性。
# <翻译结束>


<原文开始>
// Note there's reflect conversion mechanism here.
<原文结束>

# <翻译开始>
// 注意这里存在反射转换机制。
# <翻译结束>


<原文开始>
	// Note that the slice element might be type of struct,
	// so it uses Struct function doing the converting internally.
<原文结束>

# <翻译开始>
// 注意，切片元素可能为结构体类型，
// 因此内部使用Struct函数进行转换。
# <翻译结束>


<原文开始>
// Before it sets the `elem` to array, do pointer converting if necessary.
<原文结束>

# <翻译开始>
// 在将`elem`设置到切片之前，如有必要，进行指针转换。
# <翻译结束>







<原文开始>
// Try to find the original type kind of the slice element.
<原文结束>

# <翻译开始>
// 尝试查找切片元素的原始类型种类。
# <翻译结束>


<原文开始>
// Empty string cannot be assigned to string slice.
<原文结束>

# <翻译开始>
// 空字符串不能被赋值给字符串切片。
# <翻译结束>


<原文开始>
// Nil or empty pointer, it creates a new one.
<原文结束>

# <翻译开始>
// 如果是空指针或为空，它会创建一个新的。
# <翻译结束>


<原文开始>
// Not empty pointer, it assigns values to it.
<原文结束>

# <翻译开始>
// 非空指针，用于给它赋值。
# <翻译结束>


<原文开始>
// It mainly and specially handles the interface of nil value.
<原文结束>

# <翻译开始>
// 它主要用于并特别处理接口的 nil 值情况。
# <翻译结束>


<原文开始>
		// It here uses reflect converting `value` to type of the attribute and assigns
		// the result value to the attribute. It might fail and panic if the usual Go
		// conversion rules do not allow conversion.
<原文结束>

# <翻译开始>
// 这里使用 reflect 将 `value` 转换为属性的类型，并将转换后的结果值赋给该属性。如果按照 Go 语言通常的转换规则无法进行转换，则可能会导致失败并引发 panic。
# <翻译结束>


<原文开始>
// DO NOT use `params` directly as it might be type `reflect.Value`
<原文结束>

# <翻译开始>
// **请勿**直接使用 `params`，因为它可能为 `reflect.Value` 类型
# <翻译结束>

















<原文开始>
// JSON content converting.
<原文结束>

# <翻译开始>
// JSON内容转换
# <翻译结束>


<原文开始>
// The pointed element.
<原文结束>

# <翻译开始>
// 指向的元素。
# <翻译结束>


<原文开始>
// custom convert try first
<原文结束>

# <翻译开始>
// 首先尝试自定义转换
# <翻译结束>


<原文开始>
// Directly converting.
<原文结束>

# <翻译开始>
// 直接转换
# <翻译结束>


<原文开始>
// Common interface check.
<原文结束>

# <翻译开始>
// 常用接口检查。
# <翻译结束>


<原文开始>
// Default converting.
<原文结束>

# <翻译开始>
// 默认转换。
# <翻译结束>


<原文开始>
// Value is empty string.
<原文结束>

# <翻译开始>
// Value为空字符串。
# <翻译结束>

