
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
// MapOption specifies the option for map converting.
<原文结束>

# <翻译开始>
// MapOption 定义了映射转换的选项。 md5:8dc53d6fdc486bf8
# <翻译结束>


<原文开始>
	// Deep marks doing Map function recursively, which means if the attribute of given converting value
	// is also a struct/*struct, it automatically calls Map function on this attribute converting it to
	// a map[string]interface{} type variable.
<原文结束>

# <翻译开始>
	// Deep 标记表示递归地执行 Map 函数，这意味着如果给定转换值的属性也是一个结构体（struct），它会自动对这个属性调用 Map 函数，将其转换为 map[string]interface{} 类型变量。
	// md5:3653359965fb222d
# <翻译结束>


<原文开始>
// OmitEmpty ignores the attributes that has json `omitempty` tag.
<原文结束>

# <翻译开始>
	// OmitEmpty 忽略具有 json `omitempty` 标签的属性。 md5:ce80b66cfe17a0ba
# <翻译结束>


<原文开始>
// Tags specifies the converted map key name by struct tag name.
<原文结束>

# <翻译开始>
	// Tags 指定了通过结构体标签名转换后的映射键名称。 md5:b08e40ad043d7120
# <翻译结束>


<原文开始>
// Map converts any variable `value` to map[string]interface{}. If the parameter `value` is not a
// map/struct/*struct type, then the conversion will fail and returns nil.
//
// If `value` is a struct/*struct object, the second parameter `tags` specifies the most priority
// tags that will be detected, otherwise it detects the tags in order of:
// gconv, json, field name.
<原文结束>

# <翻译开始>
// Map 将任何变量 `value` 转换为 map[string]interface{}。如果参数 `value` 不是
// 类型为 map/struct/*struct，转换将会失败并返回 nil。
//
// 如果 `value` 是一个 struct/*struct 对象，第二个参数 `tags` 指定了具有最高优先级的
// 将被检测的标签，否则它会按照以下顺序检测标签：
// gconv, json, 字段名称。
// md5:34498665a6393f82
# <翻译结束>


<原文开始>
// MapDeep does Map function recursively, which means if the attribute of `value`
// is also a struct/*struct, calls Map function on this attribute converting it to
// a map[string]interface{} type variable.
// Deprecated: used Map instead.
<原文结束>

# <翻译开始>
// MapDeep递归地执行Map函数，这意味着如果`value`的属性也是一个`struct/*struct`，则会在这个属性上调用Map函数，并将其转换为map[string]interface{}类型的变量。
// 警告：建议使用Map替代。
// md5:dc0620a4d15b4389
# <翻译结束>


<原文开始>
// doMapConvert implements the map converting.
// It automatically checks and converts json string to map if `value` is string/[]byte.
//
// TODO completely implement the recursive converting for all types, especially the map.
<原文结束>

# <翻译开始>
// doMapConvert 实现了映射转换。
// 如果 `value` 是字符串或[]byte，它会自动检查并将其转换为map。
// 
// TODO 完全实现所有类型的递归转换，特别是map。
// md5:f55eadf34b47fad4
# <翻译结束>


<原文开始>
// It redirects to its underlying value if it has implemented interface iVal.
<原文结束>

# <翻译开始>
	// 如果它已经实现了iVal接口，那么它会重定向到其底层值。 md5:fb13fb87762a52a2
# <翻译结束>


<原文开始>
// Assert the common combination of types, and finally it uses reflection.
<原文结束>

# <翻译开始>
	// 断言常见的类型组合，并最终使用反射。 md5:28d02793b273a6c1
# <翻译结束>


<原文开始>
// If it is a JSON string, automatically unmarshal it!
<原文结束>

# <翻译开始>
		// 如果它是一个JSON字符串，自动反序列化它！. md5:2da2afc6ee11f379
# <翻译结束>


<原文开始>
// It returns the map directly without any changing.
<原文结束>

# <翻译开始>
			// 它直接返回映射，不做任何更改。 md5:fa0b6b4232562112
# <翻译结束>


<原文开始>
// Not a common type, it then uses reflection for conversion.
<原文结束>

# <翻译开始>
		// 并非常见类型，因此它使用反射来进行转换。 md5:a4126e9dfe7a56bd
# <翻译结束>


<原文开始>
// If it is a pointer, we should find its real data type.
<原文结束>

# <翻译开始>
	// 如果它是一个指针，我们应该找到其实际的数据类型。 md5:db4733e40015c40e
# <翻译结束>


<原文开始>
		// If `value` is type of array, it converts the value of even number index as its key and
		// the value of odd number index as its corresponding value, for example:
		// []string{"k1","v1","k2","v2"} => map[string]interface{}{"k1":"v1", "k2":"v2"}
		// []string{"k1","v1","k2"}      => map[string]interface{}{"k1":"v1", "k2":nil}
<原文结束>

# <翻译开始>
		// 如果`value`是数组类型，它将偶数索引的值作为键，奇数索引的值作为对应的值。例如：
		// `[]string{"k1","v1","k2","v2"}` => `map[string]interface{}{"k1":"v1", "k2":"v2"}`
		// `[]string{"k1","v1","k2"}`       => `map[string]interface{}{"k1":"v1", "k2":nil}`
		// md5:5e90ff5bc08f2638
# <翻译结束>


<原文开始>
// It returns directly if it is not root and with no recursive converting.
<原文结束>

# <翻译开始>
// 如果它不是根并且没有递归转换，它会直接返回。 md5:70679f33f48f5a89
# <翻译结束>


<原文开始>
// Current operation value.
<原文结束>

# <翻译开始>
// 当前操作值。 md5:2dcb5cbb4a76dbe7
# <翻译结束>


<原文开始>
// The type from top function entry.
<原文结束>

# <翻译开始>
// 从顶级函数入口的类型。 md5:6fd96f3dbc57d815
# <翻译结束>


<原文开始>
// Whether convert recursively for `current` operation.
<原文结束>

# <翻译开始>
// 是否为当前操作`current`进行递归转换。 md5:d915897a37c59c4a
# <翻译结束>


<原文开始>
// Must return map instead of Value when empty.
<原文结束>

# <翻译开始>
// 当空时，必须返回map而不是Value。 md5:e49001a917ef93fb
# <翻译结束>


<原文开始>
// quick check for nil value.
<原文结束>

# <翻译开始>
					// 快速检查值是否为nil。 md5:93138802a95bcbf7
# <翻译结束>


<原文开始>
					// in case of:
					// exception recovered: reflect: call of reflect.Value.Interface on zero Value
<原文结束>

# <翻译开始>
					// 在出现以下情况时：
					// 异常恢复：reflect: 对零值调用reflect.Value.Interface
					// md5:e32f0249911d4dde
# <翻译结束>


<原文开始>
// Map converting interface check.
<原文结束>

# <翻译开始>
		// 转换接口检查的映射。 md5:e4adcda9bbeec1fc
# <翻译结束>


<原文开始>
// Value copy, in case of concurrent safety.
<原文结束>

# <翻译开始>
			// 为了并发安全，进行值复制。 md5:57f6f9976b1be5ca
# <翻译结束>


<原文开始>
// Using reflect for converting.
<原文结束>

# <翻译开始>
		// 使用反射进行转换。 md5:ffcf9b71ef0563af
# <翻译结束>


<原文开始>
// mapKey may be the tag name or the struct attribute name.
<原文结束>

# <翻译开始>
// mapKey 可能是标签名或结构体属性名。 md5:7798f495f1f4211d
# <翻译结束>


<原文开始>
// Only convert the public attributes.
<原文结束>

# <翻译开始>
			// 只转换公共属性。 md5:090d3eafbff3ac6e
# <翻译结束>


<原文开始>
// Support json tag feature: -, omitempty
<原文结束>

# <翻译开始>
				// 支持json标签特性：-，omitempty. md5:89511416feac7bb4
# <翻译结束>


<原文开始>
// Do map converting recursively.
<原文结束>

# <翻译开始>
				// 递归地进行映射转换。 md5:1676b5bed955fd64
# <翻译结束>


<原文开始>
					// Embedded struct and has no fields, just ignores it.
					// Eg: gmeta.Meta
<原文结束>

# <翻译开始>
					// 内嵌结构体且没有字段，将忽略它。
					// 例如：gmeta.Meta
					// md5:8505cb87a6269724
# <翻译结束>


<原文开始>
						// DO NOT use rvAttrField.Interface() here,
						// as it might be changed from pointer to struct.
<原文结束>

# <翻译开始>
						// 不要在这里使用rvAttrField.Interface()，因为它可能会从指针转换为结构体。
						// md5:5cd6517f328dfd1c
# <翻译结束>


<原文开始>
						// It means this attribute field has no tag.
						// Overwrite the attribute with sub-struct attribute fields.
<原文结束>

# <翻译开始>
						// 这意味着这个属性字段没有标签。
						// 使用子结构体的属性字段覆盖该属性。
						// md5:525f64e84a599d2d
# <翻译结束>


<原文开始>
// It means this attribute field has desired tag.
<原文结束>

# <翻译开始>
					// 这意味着该属性字段具有期望的标签。 md5:e6252ec8be3f90cb
# <翻译结束>


<原文开始>
// The struct attribute is type of slice.
<原文结束>

# <翻译开始>
				// 该结构体属性是切片类型。 md5:e1a646d8191abc2f
# <翻译结束>


<原文开始>
// No recursive map value converting
<原文结束>

# <翻译开始>
				// 不进行递归地将映射值转换. md5:fd213a1b3835dd97
# <翻译结束>


<原文开始>
// The given value is type of slice.
<原文结束>

# <翻译开始>
	// 给定的值是切片类型。 md5:fb5a502257cf9a01
# <翻译结束>


<原文开始>
// MapStrStr converts `value` to map[string]string.
// Note that there might be data copy for this map type converting.
<原文结束>

# <翻译开始>
// MapStrStr 将 `value` 转换为 map[string]string 类型。
// 注意，对于这种映射类型转换，可能会有数据复制的情况发生。
// md5:a1ec9ce0d856cd1e
# <翻译结束>


<原文开始>
// MapStrStrDeep converts `value` to map[string]string recursively.
// Note that there might be data copy for this map type converting.
// Deprecated: used MapStrStr instead.
<原文结束>

# <翻译开始>
// MapStrStrDeep 递归地将`value`转换为map[string]string。
// 请注意，这种映射类型的转换可能会涉及数据复制。
// 已弃用：请使用MapStrStr代替。
// md5:79528a85e8ff4c82
# <翻译结束>

