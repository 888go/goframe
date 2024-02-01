
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
// MapOption specifies the option for map converting.
<原文结束>

# <翻译开始>
// MapOption 定义了映射转换的选项。
# <翻译结束>


<原文开始>
	// Deep marks doing Map function recursively, which means if the attribute of given converting value
	// is also a struct/*struct, it automatically calls Map function on this attribute converting it to
	// a map[string]interface{} type variable.
<原文结束>

# <翻译开始>
// Deep 标志表示递归地执行 Map 函数，这意味着如果给定转换值的属性也是一个结构体（struct/*struct），它会自动对该属性调用 Map 函数，将其转换为 map[string]interface{} 类型的变量。
# <翻译结束>


<原文开始>
// OmitEmpty ignores the attributes that has json `omitempty` tag.
<原文结束>

# <翻译开始>
// OmitEmpty 忽略具有 json `omitempty` 标签的属性。
# <翻译结束>


<原文开始>
// Tags specifies the converted map key name by struct tag name.
<原文结束>

# <翻译开始>
// Tags 通过结构体标签名称指定转换后的映射键名称。
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
// Map 将任意变量 `value` 转换为 map[string]interface{} 类型。如果参数 `value` 不是
// map、struct 或 *struct 类型，那么转换将会失败并返回 nil。
//
// 如果 `value` 是 struct 或 *struct 对象，第二个参数 `tags` 指定了优先检测的标签，
// 否则将按照 gconv、json、字段名的顺序进行检测。
# <翻译结束>


<原文开始>
// MapDeep does Map function recursively, which means if the attribute of `value`
// is also a struct/*struct, calls Map function on this attribute converting it to
// a map[string]interface{} type variable.
// Deprecated: used Map instead.
<原文结束>

# <翻译开始>
// MapDeep 递归地执行 Map 函数，这意味着如果 `value` 的属性也是一个结构体（struct/*struct），则对该属性调用 Map 函数，并将其转换为 map[string]interface{} 类型的变量。
// 注意：已弃用，请改用 Map。
# <翻译结束>


<原文开始>
// doMapConvert implements the map converting.
// It automatically checks and converts json string to map if `value` is string/[]byte.
//
// TODO completely implement the recursive converting for all types, especially the map.
<原文结束>

# <翻译开始>
// doMapConvert 实现了映射转换功能。
// 它会自动检查并转换，如果 `value` 为字符串或 []byte 类型，则将其转换为 map。
//
// TODO 完全实现对所有类型的递归转换，特别是对 map 的转换。
# <翻译结束>


<原文开始>
// It redirects to its underlying value if it has implemented interface iVal.
<原文结束>

# <翻译开始>
// 如果实现了接口 iVal，则重定向到其底层值。
# <翻译结束>







<原文开始>
// Assert the common combination of types, and finally it uses reflection.
<原文结束>

# <翻译开始>
// 断言常见类型的组合，并最终使用反射。
# <翻译结束>


<原文开始>
// If it is a JSON string, automatically unmarshal it!
<原文结束>

# <翻译开始>
// 如果这是一个JSON字符串，自动进行反序列化操作！
# <翻译结束>







<原文开始>
// It returns the map directly without any changing.
<原文结束>

# <翻译开始>
// 它直接返回该映射，不作任何更改。
# <翻译结束>


<原文开始>
// Not a common type, it then uses reflection for conversion.
<原文结束>

# <翻译开始>
// 如果不是常见类型，它将使用反射进行转换。
# <翻译结束>


<原文开始>
// If it is a pointer, we should find its real data type.
<原文结束>

# <翻译开始>
// 如果是指针，我们应该找到它指向的实际数据类型。
# <翻译结束>


<原文开始>
		// If `value` is type of array, it converts the value of even number index as its key and
		// the value of odd number index as its corresponding value, for example:
		// []string{"k1","v1","k2","v2"} => map[string]interface{}{"k1":"v1", "k2":"v2"}
		// []string{"k1","v1","k2"}      => map[string]interface{}{"k1":"v1", "k2":nil}
<原文结束>

# <翻译开始>
// 如果`value`的类型是数组，它会将偶数索引位置的值作为键，
// 奇数索引位置的值作为对应的值。例如：
// []string{"k1","v1","k2","v2"} 转换为 map[string]interface{}{"k1":"v1", "k2":"v2"}
// []string{"k1","v1","k2"}      转换为 map[string]interface{}{"k1":"v1", "k2":nil}
// 以下是中文注释：
// ```go
// 若变量`value`的类型为数组，该函数将数组中偶数下标的元素作为键，
// 奇数下标的元素作为对应的值进行转换，例如：
// []string{"k1","v1","k2","v2"} 将被转化为 map[string]interface{}{"k1":"v1", "k2":"v2"}
// []string{"k1","v1","k2"}      将被转化为 map[string]interface{}{"k1":"v1", "k2":nil}
# <翻译结束>


<原文开始>
// It returns directly if it is not root and with no recursive converting.
<原文结束>

# <翻译开始>
// 如果当前不是根目录且不需要递归转换，则直接返回。
# <翻译结束>







<原文开始>
// The type from top function entry.
<原文结束>

# <翻译开始>
// 从顶级函数入口处的类型。
# <翻译结束>







<原文开始>
// Must return map instead of Value when empty.
<原文结束>

# <翻译开始>
// 当为空时，必须返回 map 而不是 Value。
# <翻译结束>







<原文开始>
					// in case of:
					// exception recovered: reflect: call of reflect.Value.Interface on zero Value
<原文结束>

# <翻译开始>
// 在出现以下情况时：
// 恢复的异常：reflect: 对零值调用了reflect.Value.Interface
# <翻译结束>







<原文开始>
// Value copy, in case of concurrent safety.
<原文结束>

# <翻译开始>
// 值复制，以保证并发安全。
# <翻译结束>


<原文开始>
// Using reflect for converting.
<原文结束>

# <翻译开始>
// 使用reflect进行转换
# <翻译结束>







<原文开始>
// mapKey may be the tag name or the struct attribute name.
<原文结束>

# <翻译开始>
// mapKey 可能是标签名称或结构体属性名称。
# <翻译结束>


<原文开始>
// Only convert the public attributes.
<原文结束>

# <翻译开始>
// 只转换公共属性。
# <翻译结束>


<原文开始>
// Support json tag feature: -, omitempty
<原文结束>

# <翻译开始>
// 支持json标签特性：-, omitempty
// 这行注释说明了Go语言中对于结构体字段的JSON标签功能支持两种特殊标记：
// `-`：表示忽略该字段，即在进行JSON编码（Marshal）时，不会将该字段包含到生成的JSON数据中。
// `omitempty`：表示当该字段值为空（如零值、空字符串、长度为0的数组/切片/映射等）时，在进行JSON编码时不包含此字段。
# <翻译结束>


<原文开始>
// Do map converting recursively.
<原文结束>

# <翻译开始>
// 递归地执行映射转换。
# <翻译结束>


<原文开始>
					// Embedded struct and has no fields, just ignores it.
					// Eg: gmeta.Meta
<原文结束>

# <翻译开始>
// 嵌入式结构体，无字段，仅忽略它。
// 例如：gmeta.Meta
# <翻译结束>


<原文开始>
						// DO NOT use rvAttrField.Interface() here,
						// as it might be changed from pointer to struct.
<原文结束>

# <翻译开始>
// **注意**：在此处不要使用rvAttrField.Interface()，
// 因为它可能从指针变为结构体。
# <翻译结束>


<原文开始>
						// It means this attribute field has no tag.
						// Overwrite the attribute with sub-struct attribute fields.
<原文结束>

# <翻译开始>
// 这意味着该属性字段没有标签。
// 用子结构体的属性字段覆盖该属性。
// 这段Go语言代码注释翻译成中文为：
// ```go
// 这表示该属性字段未设置标签。
// 使用子结构体的属性字段来重写该属性。
# <翻译结束>


<原文开始>
// It means this attribute field has desired tag.
<原文结束>

# <翻译开始>
// 这意味着该属性字段具有期望的标签。
# <翻译结束>


<原文开始>
// The struct attribute is type of slice.
<原文结束>

# <翻译开始>
// 结构体属性的类型为切片。
# <翻译结束>


<原文开始>
// No recursive map value converting
<原文结束>

# <翻译开始>
// 不进行递归的映射值转换
# <翻译结束>


<原文开始>
// The given value is type of slice.
<原文结束>

# <翻译开始>
// 给定的值是切片类型。
# <翻译结束>


<原文开始>
// MapStrStr converts `value` to map[string]string.
// Note that there might be data copy for this map type converting.
<原文结束>

# <翻译开始>
// MapStrStr 将`value`转换为map[string]string类型。
// 注意，这种映射类型的转换可能会导致数据复制。
# <翻译结束>


<原文开始>
// MapStrStrDeep converts `value` to map[string]string recursively.
// Note that there might be data copy for this map type converting.
// Deprecated: used MapStrStr instead.
<原文结束>

# <翻译开始>
// MapStrStrDeep递归地将`value`转换为map[string]string类型。
// 注意，这种映射类型转换可能会涉及数据复制。
// 废弃: 请改用MapStrStr。
# <翻译结束>


<原文开始>
// Whether convert recursively for `current` operation.
<原文结束>

# <翻译开始>
// 是否对`current`操作进行递归转换
# <翻译结束>


<原文开始>
// Map converting interface check.
<原文结束>

# <翻译开始>
// 接口检查映射转换
# <翻译结束>


<原文开始>
// A copy of current map.
<原文结束>

# <翻译开始>
// 当前映射的一份副本。
# <翻译结束>


<原文开始>
// Current operation value.
<原文结束>

# <翻译开始>
// 当前操作值。
# <翻译结束>


<原文开始>
// Map converting option.
<原文结束>

# <翻译开始>
// 映射转换选项。
# <翻译结束>


<原文开始>
// quick check for nil value.
<原文结束>

# <翻译开始>
// 快速检查空值。
# <翻译结束>


<原文开始>
// attribute value type.
<原文结束>

# <翻译开始>
// 属性值类型
# <翻译结束>

