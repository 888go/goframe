
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// SliceCopy does a shallow copy of slice `data` for most commonly used slice type
// []interface{}.
<原文结束>

# <翻译开始>
// SliceCopy 对于最常用的切片类型（[]interface{}）执行浅拷贝操作。 md5:d119afb140e8324a
# <翻译结束>


<原文开始>
// SliceInsertBefore inserts the `values` to the front of `index` and returns a new slice.
<原文结束>

# <翻译开始>
// SliceInsertBefore 将 `values` 插入到 `index` 位置的前面，并返回一个新的切片。 md5:d1cbe1d61df82fb6
# <翻译结束>


<原文开始>
// SliceInsertAfter inserts the `values` to the back of `index` and returns a new slice.
<原文结束>

# <翻译开始>
// SliceInsertAfter 在 `index` 后方插入 `values`，并返回一个新的切片。 md5:bb561266d0b2b921
# <翻译结束>


<原文开始>
// SliceDelete deletes an element at `index` and returns the new slice.
// It does nothing if the given `index` is invalid.
<原文结束>

# <翻译开始>
// SliceDelete 从索引 `index` 处删除一个元素，并返回新的切片。如果给定的 `index` 不有效，它不会做任何操作。 md5:f57a8afe207e8169
# <翻译结束>


<原文开始>
// Determine array boundaries when deleting to improve deletion efficiency.
<原文结束>

# <翻译开始>
	// 在删除时确定数组边界，以提高删除效率。 md5:bc969ee880edf699
# <翻译结束>


<原文开始>
	// If it is a non-boundary delete,
	// it will involve the creation of an array,
	// then the deletion is less efficient.
<原文结束>

# <翻译开始>
	// 如果是一个非边界删除，
	// 它将涉及创建一个数组，
	// 那么删除操作效率较低。 md5:6a664196d66bc968
# <翻译结束>


<原文开始>
// SliceToMap converts slice type variable `slice` to `map[string]interface{}`.
// Note that if the length of `slice` is not an even number, it returns nil.
// Eg:
// ["K1", "v1", "K2", "v2"] => {"K1": "v1", "K2": "v2"}
// ["K1", "v1", "K2"]       => nil
<原文结束>

# <翻译开始>
// SliceToMap 将切片类型变量 `slice` 转换为 `map[string]interface{}` 类型。注意，如果 `slice` 的长度不是偶数，它将返回nil。
// 例如：
// ["K1", "v1", "K2", "v2"] => {"K1": "v1", "K2": "v2"}
// ["K1", "v1", "K2"]       => nil md5:f7b4384607d6bfbe
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
// 切片中指定列的值将作为返回映射的键。
// 例如：
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K1") => {"v1": {"K1": "v1", "K2": 1}, "v2": {"K1": "v2", "K2": 2}}
// SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K2") => {1: {"K1": "v1", "K2": 1}, 2: {"K1": "v2", "K2": 2}} md5:bf178f79b93bd1b2
# <翻译结束>

