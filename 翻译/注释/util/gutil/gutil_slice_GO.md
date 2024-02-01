
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
// SliceCopy does a shallow copy of slice `data` for most commonly used slice type
// []interface{}.
<原文结束>

# <翻译开始>
// SliceCopy 对于最常用的切片类型 []interface{}，执行浅拷贝操作，复制 `data` 切片。
# <翻译结束>


<原文开始>
// SliceInsertBefore inserts the `values` to the front of `index` and returns a new slice.
<原文结束>

# <翻译开始>
// SliceInsertBefore 将 `values` 插入到 `index` 位置之前，并返回一个新的切片。
# <翻译结束>


<原文开始>
// SliceInsertAfter inserts the `values` to the back of `index` and returns a new slice.
<原文结束>

# <翻译开始>
// SliceInsertAfter 在 `index` 后面插入 `values`，并返回一个新的切片。
# <翻译结束>


<原文开始>
// SliceDelete deletes an element at `index` and returns the new slice.
// It does nothing if the given `index` is invalid.
<原文结束>

# <翻译开始>
// SliceDelete 在`index`处删除一个元素并返回新的切片。
// 如果给定的`index`无效，则不做任何操作。
# <翻译结束>


<原文开始>
// Determine array boundaries when deleting to improve deletion efficiency.
<原文结束>

# <翻译开始>
// 确定删除时的数组边界以提高删除效率
# <翻译结束>


<原文开始>
	// If it is a non-boundary delete,
	// it will involve the creation of an array,
	// then the deletion is less efficient.
<原文结束>

# <翻译开始>
// 如果这是一个非边界删除，
// 那么它将涉及创建一个数组，
// 因此，删除操作效率较低。
# <翻译结束>


<原文开始>
// SliceToMap converts slice type variable `slice` to `map[string]interface{}`.
// Note that if the length of `slice` is not an even number, it returns nil.
// Eg:
// ["K1", "v1", "K2", "v2"] => {"K1": "v1", "K2": "v2"}
// ["K1", "v1", "K2"]       => nil
<原文结束>

# <翻译开始>
// SliceToMap 将切片类型变量 `slice` 转换为 `map[string]interface{}` 类型。
// 注意，如果 `slice` 的长度不是偶数，则返回 nil。
// 示例：
// ["K1", "v1", "K2", "v2"] => {"K1": "v1", "K2": "v2"}
// ["K1", "v1", "K2"]       => nil
# <翻译结束>


<原文开始>
// SliceToMapWithColumnAsKey converts slice type variable `slice` to `map[interface{}]interface{}`
// The value of specified column use as the key for returned map.
// Eg:
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K1") => {"v1": {"K1": "v1", "K2": 1}, "v2": {"K1": "v2", "K2": 2}}
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K2") => {1: {"K1": "v1", "K2": 1}, 2: {"K1": "v2", "K2": 2}}
<原文结束>

# <翻译开始>
// SliceToMapWithColumnAsKey 将切片类型变量 `slice` 转换为 `map[interface{}]interface{}`
// 指定列的值作为返回映射中的键。
// 示例：
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K1") => {"v1": {"K1": "v1", "K2": 1}, "v2": {"K1": "v2", "K2": 2}}
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K2") => {1: {"K1": "v1", "K2": 1}, 2: {"K1": "v2", "K2": 2}}
// 这段Go语言代码注释翻译成中文后，其内容如下：
// SliceToMapWithColumnAsKey 函数将 slice 类型变量转换为 map 类型变量，其中 map 的键是指定列的值。
// 例如：
// 当调用 SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K1") 时，返回结果为 {"v1": {"K1": "v1", "K2": 1}, "v2": {"K1": "v2", "K2": 2}}
// 当调用 SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K2") 时，返回结果为 {1: {"K1": "v1", "K2": 1}, 2: {"K1": "v2", "K2": 2}}
# <翻译结束>

