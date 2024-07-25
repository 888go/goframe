
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
// 结构体将参数的键值对映射到对应结构对象的属性。
// 第三个参数 `mapping` 不必要，表示自定义键名和属性名之间的映射规则（区分大小写）。
// 
// 注意：
// 1. `params` 可以是任何类型的 map/struct，通常为 map。
// 2. `pointer` 应该是 *struct/**struct 类型，即指向结构体对象或结构体指针。
// 3. 只有结构体对象的公共属性可以被映射。
// 4. 如果 `params` 是一个 map，其键 `params` 可以是小写。在映射过程中，它会自动将键的首字母转换为大写进行匹配。如果键不匹配，它将忽略该键。
// md5:b39a46da903b06f5
# <翻译结束>


<原文开始>
// StructTag acts as Struct but also with support for priority tag feature, which retrieves the
// specified tags for `params` key-value items to struct attribute names mapping.
// The parameter `priorityTag` supports multiple tags that can be joined with char ','.
<原文结束>

# <翻译开始>
// StructTag 作为 Struct 的功能，但同时也支持优先级标签特性。这个特性用于获取 `params` 键值对中的指定标签，并将其映射到结构体属性名上。
// 参数 `priorityTag` 支持多个标签，这些标签之间可以使用逗号 `,` 进行连接。
// md5:14d47a8c22737303
# <翻译结束>


<原文开始>
// doStruct is the core internal converting function for any data to struct.
<原文结束>

# <翻译开始>
// doStruct 是将任何数据转换为结构体的核心内部函数。 md5:43cdc6b6cc398c7c
# <翻译结束>


<原文开始>
// If `params` is nil, no conversion.
<原文结束>

# <翻译开始>
		// 如果`params`为nil，则不进行转换。 md5:0520708a0e7e1c1d
# <翻译结束>


<原文开始>
// JSON content converting.
<原文结束>

# <翻译开始>
	// JSON内容转换。 md5:8a29b5a7aa430047
# <翻译结束>


<原文开始>
// Catch the panic, especially the reflection operation panics.
<原文结束>

# <翻译开始>
		// 捕获panic，尤其是反射操作引发的panic。 md5:dd183bf8028f513a
# <翻译结束>


<原文开始>
// DO NOT use `params` directly as it might be type `reflect.Value`
<原文结束>

# <翻译开始>
// 不要直接使用`params`，因为它可能是`reflect.Value`类型. md5:f469653f5ba4e08c
# <翻译结束>


<原文开始>
// Using IsNil on reflect.Ptr variable is OK.
<原文结束>

# <翻译开始>
		// 使用 IsNil 检查 reflect.Ptr 类型的变量是可行的。 md5:0ba920ba8a6a19cf
# <翻译结束>


<原文开始>
	// If `params` and `pointer` are the same type, the do directly assignment.
	// For performance enhancement purpose.
<原文结束>

# <翻译开始>
	// 如果`params`和`pointer`是相同类型，直接进行赋值操作。
	// 为了性能优化。
	// md5:87eefbed1426eef0
# <翻译结束>


<原文开始>
// Normal unmarshalling interfaces checks.
<原文结束>

# <翻译开始>
	// 通常的接口解码检查。 md5:838cb73b6b92dc54
# <翻译结束>


<原文开始>
	// It automatically creates struct object if necessary.
	// For example, if `pointer` is **User, then `elem` is *User, which is a pointer to User.
<原文结束>

# <翻译开始>
	// 如果必要，它会自动创建结构体对象。
	// 例如，如果`pointer`是**User（双星号表示指针），那么`elem`就是*User，即User类型的指针。
	// md5:172757349701f610
# <翻译结束>


<原文开始>
// If it is converted failed, it reset the `pointer` to nil.
<原文结束>

# <翻译开始>
					// 如果转换失败，它将`pointer`重置为nil。 md5:52f95bfcfceeefc0
# <翻译结束>


<原文开始>
		// if v, ok := pointerElemReflectValue.Interface().(iUnmarshalValue); ok {
		//	return v.UnmarshalValue(params)
		// }
		// Note that it's `pointerElemReflectValue` here not `pointerReflectValue`.
<原文结束>

# <翻译开始>
		// 如果v, ok := pointerElemReflectValue.Interface().(iUnmarshalValue); ok {
		// 	return v.UnmarshalValue(params)
		// }
		// 请注意，这里是`pointerElemReflectValue`而不是`pointerReflectValue`。
		// md5:722eb6b1c6132d70
# <翻译结束>


<原文开始>
// Retrieve its element, may be struct at last.
<原文结束>

# <翻译开始>
		// 获取其元素，可能是最后的结构体。 md5:4a887dcf759fad9d
# <翻译结束>


<原文开始>
	// paramsMap is the map[string]interface{} type variable for params.
	// DO NOT use MapDeep here.
<原文结束>

# <翻译开始>
	// paramsMap 是一个类型为 map[string]interface{} 的变量，用于存储参数。
	// 不要在這裡使用 MapDeep。
	// md5:96735ea71b035d62
# <翻译结束>


<原文开始>
// Nothing to be done as the parameters are empty.
<原文结束>

# <翻译开始>
	// 由于参数为空，无需进行任何操作。 md5:958747d8f67e1e73
# <翻译结束>


<原文开始>
// Holds the info for subsequent converting.
<原文结束>

# <翻译开始>
	// 用于后续转换的信息存储。 md5:5cb67597e7ff966f
# <翻译结束>


<原文开始>
// Found value by tag name or field name from input.
<原文结束>

# <翻译开始>
// 从输入中通过标签名或字段名找到值。 md5:ec1aa28d82e3ec74
# <翻译结束>


<原文开始>
// The associated reflection field index.
<原文结束>

# <翻译开始>
// 关联的反射字段索引。 md5:f8941ca77d57dd95
# <翻译结束>


<原文开始>
// Field name or tag name for field tag by priority tags.
<原文结束>

# <翻译开始>
// 按优先级标签的字段名称或字段标签名称。 md5:2a10a58537738182
# <翻译结束>


<原文开始>
// Only do converting to public attributes.
<原文结束>

# <翻译开始>
		// 只转换为公共属性。 md5:4fc00fe51391895a
# <翻译结束>


<原文开始>
// Maybe it's struct/*struct embedded.
<原文结束>

# <翻译开始>
		// 也许它嵌入了struct。 md5:e77a8f08191e1bd2
# <翻译结束>


<原文开始>
			// type Name struct {
			//    LastName  string `json:"lastName"`
			//    FirstName string `json:"firstName"`
			// }
			//
			// type User struct {
			//     Name `json:"name"`
			//     // ...
			// }
			//
			// It is only recorded if the name has a fieldTag
<原文结束>

# <翻译开始>
			// 定义一个名为Name的结构体，其中包含两个字段：LastName和FirstName，它们都有`json`标签进行标记
			// ```
			// type Name struct {
			//    LastName  string `json:"lastName"`
			//    FirstName string `json:"firstName"`
			// }
			// ```
			// 
			// 定义一个User结构体，其中包含一个嵌套的Name结构体，并使用`json:"name"`对整个嵌套结构进行标记
			// ```
			// type User struct {
			//     Name `json:"name"`
			//     			// ...
			// }
			// ```
			// 
			// 只有当Name结构体中包含fieldTag（字段标签）时，才会记录这些信息
			// md5:d42e389449351045
# <翻译结束>


<原文开始>
// Ignore the interface attribute if it's nil.
<原文结束>

# <翻译开始>
			// 如果接口属性为nil，则忽略它。 md5:5bbafbaa5b14794d
# <翻译结束>


<原文开始>
// Use the native elemFieldName name as the fieldTag
<原文结束>

# <翻译开始>
			// 使用原生的elemFieldName名称作为字段标签. md5:80bfd9b406ef430f
# <翻译结束>


<原文开始>
// Nothing to be converted.
<原文结束>

# <翻译开始>
	// 没有需要转换的内容。 md5:68441f55873cce91
# <翻译结束>


<原文开始>
// Search the parameter value for the field.
<原文结束>

# <翻译开始>
	// 在参数值中搜索该字段。 md5:761e9f220df7696c
# <翻译结束>


<原文开始>
	// Firstly, search according to custom mapping rules.
	// If a possible direct assignment is found, reduce the number of subsequent map searches.
<原文结束>

# <翻译开始>
	// 首先，根据自定义的映射规则进行搜索。
	// 如果找到了可能的直接赋值关系，减少后续映射搜索的数量。
	// md5:50dd567944f99367
# <翻译结束>


<原文开始>
// Prevent setting of non-existent fields
<原文结束>

# <翻译开始>
		// 防止设置不存在的字段. md5:408a34ea9e6a0539
# <翻译结束>


<原文开始>
// Prevent non-existent values from being set.
<原文结束>

# <翻译开始>
			// 防止不存在的值被设置。 md5:16a6e1bcb81b8eb9
# <翻译结束>


<原文开始>
// Indicates that those values have been used and cannot be reused.
<原文结束>

# <翻译开始>
		// 表示这些值已被使用，不能重复使用。 md5:66845c8e5a8adbe8
# <翻译结束>


<原文开始>
// If it is not empty, the tag or elemFieldName name matches
<原文结束>

# <翻译开始>
		// 如果非空，标签或elemFieldName的名称匹配. md5:dcf5990abe97052c
# <翻译结束>


<原文开始>
// If value is nil, a fuzzy match is used for search the key and value for converting.
<原文结束>

# <翻译开始>
		// 如果value为nil，搜索时会使用模糊匹配来转换键和值。 md5:30209602b5ceef13
# <翻译结束>


<原文开始>
			// If there's something else in the tag string,
			// it uses the first part which is split using char ','.
			// Example:
			// orm:"id, priority"
			// orm:"name, with:uid=id"
<原文结束>

# <翻译开始>
			// 如果标签字符串中还有其他内容，
			// 它会使用以逗号','分隔的第一部分。
			// 例如：
			// `orm:"id, priority"`
			// `orm:"name, with:uid=id"` 
			// 
			// 这段注释说明了一个Go语言中的ORM（对象关系映射）相关代码。它解释了当解析一个包含多个属性的标签字符串时，程序会选择以逗号分隔的第一个属性作为主要处理的部分。如果标签格式为`attribute1, attribute2`，则只会使用`attribute1`。另一个例子展示了如何在`name`属性中使用额外的条件，即`with:uid=id`。
			// md5:fab9db8addb2ccc4
# <翻译结束>


<原文开始>
// fuzzy matching rule:
// to match field name and param key in case-insensitive and without symbols.
<原文结束>

# <翻译开始>
// 模糊匹配规则：
// 不区分大小写，不考虑符号地匹配字段名和参数键。
// md5:22c4645c8af23d0d
# <翻译结束>


<原文开始>
// bindVarToStructAttrWithFieldIndex sets value to struct object attribute by name.
<原文结束>

# <翻译开始>
// bindVarToStructAttrWithFieldIndex 通过名称将值设置给结构体对象的属性。 md5:884feed9b741e07a
# <翻译结束>


<原文开始>
// CanSet checks whether attribute is public accessible.
<原文结束>

# <翻译开始>
	// CanSet 检查该属性是否可以公开访问。 md5:fafe4f3a8bd7621f
# <翻译结束>


<原文开始>
		// Try to call custom converter.
		// Issue: https://github.com/gogf/gf/issues/3099
<原文结束>

# <翻译开始>
		// 尝试调用自定义转换器。
		// 问题：https:		//github.com/gogf/gf/issues/3099
		// md5:e874679d6ecc39f0
# <翻译结束>


<原文开始>
		// Special handling for certain types:
		// - Overwrite the default type converting logic of stdlib for time.Time/*time.Time.
<原文结束>

# <翻译开始>
		// 对某些类型进行特殊处理：
		// - 重写stdlib中time.Time类型的默认类型转换逻辑。
		// md5:39ca7f7684bdc13c
# <翻译结束>


<原文开始>
		// Hold the time zone consistent in recursive
		// Issue: https://github.com/gogf/gf/issues/2980
<原文结束>

# <翻译开始>
		// 在递归中保持时区一致
		// 问题：https:		//github.com/gogf/gf/issues/2980
		// md5:1d09e937a28bf051
# <翻译结束>


<原文开始>
// Common interface check.
<原文结束>

# <翻译开始>
		// 公共接口检查。 md5:0e7cc3af409e672f
# <翻译结束>


<原文开始>
// bindVarToReflectValueWithInterfaceCheck does bind using common interfaces checks.
<原文结束>

# <翻译开始>
// bindVarToReflectValueWithInterfaceCheck 使用通用接口检查进行绑定。 md5:ede209e9eacebf79
# <翻译结束>


<原文开始>
// Not a pointer, but can token address, that makes it can be unmarshalled.
<原文结束>

# <翻译开始>
		// 不是指针，但可以处理地址，因此它可以被反序列化。 md5:52a739dbed72b8c0
# <翻译结束>


<原文开始>
// If it is not a valid JSON string, it then adds char `"` on its both sides to make it is.
<原文结束>

# <翻译开始>
			// 如果它不是一个有效的JSON字符串，那么就在它的两边添加字符 `"` 以使其成为有效JSON字符串。 md5:d6a38f1500604604
# <翻译结束>


<原文开始>
// bindVarToReflectValue sets `value` to reflect value object `structFieldValue`.
<原文结束>

# <翻译开始>
// bindVarToReflectValue 将 `value` 设置为反射值对象 `structFieldValue`。 md5:c78b60ec569060eb
# <翻译结束>


<原文开始>
// Converting using `Set` interface implements, for some types.
<原文结束>

# <翻译开始>
	// 使用`Set`接口实现转换，对于某些类型。 md5:51e8e3ad23771259
# <翻译结束>


<原文开始>
// Converting using reflection by kind.
<原文结束>

# <翻译开始>
	// 使用反射按类型进行转换。 md5:e3c406f111505fd2
# <翻译结束>


<原文开始>
// Recursively converting for struct attribute.
<原文结束>

# <翻译开始>
		// 递归转换结构体属性。 md5:ae6513ef6e56f654
# <翻译结束>


<原文开始>
// Note there's reflect conversion mechanism here.
<原文结束>

# <翻译开始>
			// 请注意这里存在反射转换机制。 md5:84599bf48af19237
# <翻译结束>


<原文开始>
	// Note that the slice element might be type of struct,
	// so it uses Struct function doing the converting internally.
<原文结束>

# <翻译开始>
	// 注意，切片元素的类型可能是结构体，
	// 因此它内部使用了一个名为Struct的函数来进行转换。
	// md5:b8519d4d1a736c40
# <翻译结束>


<原文开始>
// Before it sets the `elem` to array, do pointer converting if necessary.
<原文结束>

# <翻译开始>
				// 在将`elem`设置为数组之前，如果必要的话进行指针转换。 md5:1466632fc1d552e6
# <翻译结束>


<原文开始>
// Try to find the original type kind of the slice element.
<原文结束>

# <翻译开始>
					// 尝试找到切片元素的原始类型类别。 md5:903e45eb4bc9a592
# <翻译结束>


<原文开始>
// Empty string cannot be assigned to string slice.
<原文结束>

# <翻译开始>
						// 空字符串不能赋值给字符串切片。 md5:7015d8a83525c473
# <翻译结束>


<原文开始>
// Nil or empty pointer, it creates a new one.
<原文结束>

# <翻译开始>
			// 如果是空指针或空列表，它会创建一个新的。 md5:a005c5e6ed40f663
# <翻译结束>


<原文开始>
// Not empty pointer, it assigns values to it.
<原文结束>

# <翻译开始>
			// 非空指针，它会给它赋值。 md5:2bd4c15a81dcbdcf
# <翻译结束>


<原文开始>
// It mainly and specially handles the interface of nil value.
<原文结束>

# <翻译开始>
	// 它主要且特别地处理了nil值的接口。 md5:0c8e2dd31d82d96e
# <翻译结束>


<原文开始>
		// It here uses reflect converting `value` to type of the attribute and assigns
		// the result value to the attribute. It might fail and panic if the usual Go
		// conversion rules do not allow conversion.
<原文结束>

# <翻译开始>
		// 此处使用反射将`value`转换为属性的类型，然后将结果值赋给该属性。
		// 如果常规的Go转换规则不允许转换，此操作可能会失败并引发恐慌。
		// md5:931b86f723a12b7c
# <翻译结束>

