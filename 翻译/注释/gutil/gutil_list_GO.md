
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
// ListItemValues 通过键 `key` 获取并返回所有项（item）结构体或映射中的元素。
// 注意，参数 `list` 应为包含映射或结构体元素的切片类型，否则将返回一个空切片。
//
// 参数 `list` 支持以下类型：
// []map[string]interface{}
// []map[string]子映射
// []struct
// []struct:子结构体
// 注意，只有当提供可选参数 `subKey` 时，子映射/子结构体才有意义。
# <翻译结束>


<原文开始>
// ItemValue retrieves and returns its value of which name/attribute specified by `key`.
// The parameter `item` can be type of map/*map/struct/*struct.
<原文结束>

# <翻译开始>
// ItemValue 通过 `key` 参数指定的名称/属性获取并返回其对应的值。
// 参数 `item` 可以是 map/*map/struct/*struct 类型。
# <翻译结束>


<原文开始>
// The `key` must be type of string.
<原文结束>

# <翻译开始>
// `key`必须为字符串类型。
# <翻译结束>


<原文开始>
// The `mapKey` must be type of string.
<原文结束>

# <翻译开始>
// `mapKey`必须为字符串类型。
# <翻译结束>


<原文开始>
// ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
<原文结束>

# <翻译开始>
// ListItemValuesUnique 通过键 `key` 获取并返回所有结构体或映射中的唯一元素。
// 注意，参数 `list` 应为包含映射或结构体元素的切片类型，
// 否则将返回一个空切片。
# <翻译结束>







<原文开始>
// ListToMapByKey converts `list` to a map[string]interface{} of which key is specified by `key`.
// Note that the item value may be type of slice.
<原文结束>

# <翻译开始>
// ListToMapByKey 将 `list` 转换为一个 map[string]interface{}，其中键由 `key` 指定。
// 注意，项的值可能为 slice 类型。
# <翻译结束>


<原文开始>
// make byte slice comparable
<原文结束>

# <翻译开始>
// 使字节切片可比较
# <翻译结束>

