// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

// 包gmap提供了最常用的映射容器，同时支持并发安全/非安全切换功能。
package gmap

type (
	Map     = AnyAnyMap // Map 是 AnyAnyMap 的别名。
	HashMap = AnyAnyMap // HashMap 是 AnyAnyMap 的别名。
)

// New 创建并返回一个空的哈希表。
// 参数`safe`用于指定是否使用线程安全的map，默认为false。
func New(safe ...bool) *Map {
	return NewAnyAnyMap(safe...)
}

// NewFrom 创建并返回一个由给定的 `data` 地图生成的哈希映射。
// 注意，参数 `data` 中的地图将被设置为底层数据地图（无深度复制），
// 当在外部修改该映射时，可能会存在一些并发安全问题。
// 参数 `safe` 用于指定是否在并发环境下使用安全的树结构，默认情况下为 false。
func NewFrom(data map[interface{}]interface{}, safe ...bool) *Map {
	return NewAnyAnyMapFrom(data, safe...)
}

// NewHashMap 创建并返回一个空的哈希映射。
// 参数 `safe` 用于指定是否使用线程安全的 map，其默认值为 false。
func NewHashMap(safe ...bool) *Map {
	return NewAnyAnyMap(safe...)
}

// NewHashMapFrom 通过给定的 map `data` 创建并返回一个哈希映射。
// 注意，参数 `data` 中的地图将被设置为底层数据映射（非深度拷贝），如果在外部修改此映射，可能会存在一些并发安全问题。
// 参数 `safe` 用于指定是否在并发安全场景下使用树结构，默认情况下为 false。
func NewHashMapFrom(data map[interface{}]interface{}, safe ...bool) *Map {
	return NewAnyAnyMapFrom(data, safe...)
}
