
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
// ListItemValues retrieves and returns the elements of all item struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
//
// The parameter `list` supports types like:
// []map[string]interface{}
// []map[string]sub-map
// []struct
// []struct:sub-struct
// Note that the sub-map/sub-struct makes sense only if the optional parameter `subKey` is given.
<原文结束>

# <翻译开始>
// ListItemValues 从所有元素为映射或结构体的切片（list）中，根据给定的键（key）获取并返回对应的值。
// 注意，参数 list 应该是包含 map 或结构体元素的切片，否则将返回一个空切片。
// 
// 参数 list 支持的类型包括：
// []map[string]interface{}
// []map[string]子映射
// []struct
// []struct:子结构体
// 如果提供了可选参数 `subKey`，子映射/子结构体才有意义。
// md5:9523dac525318de2
# <翻译结束>


<原文开始>
// ItemValue retrieves and returns its value of which name/attribute specified by `key`.
// The parameter `item` can be type of map/*map/struct/*struct.
<原文结束>

# <翻译开始>
// ItemValue 获取并返回由`key`指定的名称/属性的值。
// 参数`item`可以是地图(map)、指针地图(*map)、结构体(struct)或指针结构体(*struct)类型。
// md5:ca5bcda09a11157b
# <翻译结束>


<原文开始>
// The `key` must be type of string.
<原文结束>

# <翻译开始>
// `key` 必须是字符串类型。. md5:6ffd36d1a5fc0de1
# <翻译结束>


<原文开始>
// The `mapKey` must be type of string.
<原文结束>

# <翻译开始>
// `mapKey`必须是字符串类型。. md5:d2b6db36f99feed4
# <翻译结束>


<原文开始>
// ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
<原文结束>

# <翻译开始>
// ListItemValuesUnique 获取并返回具有键为`key`的所有结构体/映射的独特元素。
// 请注意，参数`list`应为包含映射或结构体元素的切片类型，否则将返回一个空切片。
// md5:0f361d3ff901d0a1
# <翻译结束>


<原文开始>
// make byte slice comparable
<原文结束>

# <翻译开始>
// 将字节切片设置为可比较. md5:1ca9d6fe5290f517
# <翻译结束>


<原文开始>
// ListToMapByKey converts `list` to a map[string]interface{} of which key is specified by `key`.
// Note that the item value may be type of slice.
<原文结束>

# <翻译开始>
// ListToMapByKey 将 `list` 转换为一个键为 `key` 的 map[string]interface{}。注意，项的值可能为切片类型。
// md5:6509753e629d5dc6
# <翻译结束>

