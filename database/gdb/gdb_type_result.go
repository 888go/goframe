// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"database/sql"
	"math"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/util/gconv"
)

// IsEmpty 检查并返回 `r` 是否为空。
func (r Result) X是否为空() bool {
	return r == nil || r.X取数量() == 0
}

// Len 返回结果列表的长度。
func (r Result) X取数量() int {
	return len(r)
}

// Size 是函数 Len 的别名。
func (r Result) Size别名() int {
	return r.X取数量()
}

// Chunk 将一个 Result 切分成多个 Result，
// 每个数组的大小由 `size` 决定。
// 最后一块切片可能包含少于 size 个元素。
func (r Result) X分割(数量 int) []Result {
	if 数量 < 1 {
		return nil
	}
	length := len(r)
	chunks := int(math.Ceil(float64(length) / float64(数量)))
	var n []Result
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * 数量
		if end > length {
			end = length
		}
		n = append(n, r[i*数量:end])
		i++
	}
	return n
}

// Json 将 `r` 转换为 JSON 格式的内容。
func (r Result) X取json() string {
	content, _ := json类.X创建(r.X取Map数组()).X取json文本()
	return content
}

// Xml 将`r`转换为XML格式的内容。
func (r Result) X取xml(根标记 ...string) string {
	content, _ := json类.X创建(r.X取Map数组()).X取xml文本(根标记...)
	return content
}

// List 将 `r` 转换为一个 List。
func (r Result) X取Map数组() Map数组 {
	list := make(Map数组, len(r))
	for k, v := range r {
		list[k] = v.X取Map()
	}
	return list
}

// Array 根据指定列字段获取并返回其值作为一个切片。
// 当列字段只有一个时，参数 `field` 可选。
// 如果未给出参数 `field`，则默认的 `field` 为 `Result` 中第一项的第一个字段名。
func (r Result) X取字段数组(字段名称 ...string) []Value {
	array := make([]Value, len(r))
	if len(r) == 0 {
		return array
	}
	key := ""
	if len(字段名称) > 0 && 字段名称[0] != "" {
		key = 字段名称[0]
	} else {
		for k := range r[0] {
			key = k
			break
		}
	}
	for k, v := range r {
		array[k] = v[key]
	}
	return array
}

// MapKeyValue 将 `r` 转换为一个 map[string]Value，其中键由 `key` 指定。
// 注意，项值可能为切片类型。
func (r Result) X取字段Map泛型类(字段名称 string) map[string]Value {
	var (
		s              string
		m              = make(map[string]Value)
		tempMap        = make(map[string][]interface{})
		hasMultiValues bool
	)
	for _, item := range r {
		if k, ok := item[字段名称]; ok {
			s = k.String()
			tempMap[s] = append(tempMap[s], item)
			if len(tempMap[s]) > 1 {
				hasMultiValues = true
			}
		}
	}
	for k, v := range tempMap {
		if hasMultiValues {
			m[k] = 泛型类.X创建(v)
		} else {
			m[k] = 泛型类.X创建(v[0])
		}
	}
	return m
}

// MapKeyStr 将 `r` 转换为一个 map[string]Map 类型的映射，其中的键由 `key` 指定。
func (r Result) X取字段MapStr(字段名称 string) map[string]Map {
	m := make(map[string]Map)
	for _, item := range r {
		if v, ok := item[字段名称]; ok {
			m[v.String()] = item.X取Map()
		}
	}
	return m
}

// MapKeyInt 将 `r` 转换为一个映射 map[int]Map，其中键由 `key` 指定。
// （注：这里可能需要上下文信息，对于 `Map` 类型没有明确说明，所以翻译时假设它是一个已知的类型名。如果 `Map` 是自定义类型或有特殊含义，请替换为实际含义。）
func (r Result) X取字段MapInt(字段名称 string) map[int]Map {
	m := make(map[int]Map)
	for _, item := range r {
		if v, ok := item[字段名称]; ok {
			m[v.X取整数()] = item.X取Map()
		}
	}
	return m
}

// MapKeyUint 将`r`转换为一个map[uint]Map类型，其中键由`key`指定。
func (r Result) X取字段MapUint(字段名称 string) map[uint]Map {
	m := make(map[uint]Map)
	for _, item := range r {
		if v, ok := item[字段名称]; ok {
			m[v.X取正整数()] = item.X取Map()
		}
	}
	return m
}

// RecordKeyStr 将 `r` 转换为一个 map[string]Record 类型的映射，其中键由 `key` 指定。
func (r Result) RecordKeyStr(key string) map[string]Record {
	m := make(map[string]Record)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.String()] = item
		}
	}
	return m
}

// RecordKeyInt 将 `r` 转换为一个 map[int]Record 类型的映射，其中键由 `key` 指定。
func (r Result) RecordKeyInt(key string) map[int]Record {
	m := make(map[int]Record)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.X取整数()] = item
		}
	}
	return m
}

// RecordKeyUint 将 `r` 转换为一个 map[uint]Record 类型的映射，其中键由 `key` 指定。
func (r Result) RecordKeyUint(key string) map[uint]Record {
	m := make(map[uint]Record)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.X取正整数()] = item
		}
	}
	return m
}

// Structs 将 `r` 转换为结构体切片。
// 注意参数 `pointer` 应该是指向结构体切片的指针类型，即 *[]struct 或 *[]*struct。
func (r Result) X取数组结构体指针(结构体指针 interface{}) (错误 error) {
	// 如果结果为空且目标指针不为空，则返回错误。
	if r.X是否为空() {
		if !empty.IsEmpty(结构体指针, true) {
			return sql.ErrNoRows
		}
		return nil
	}
	return 转换类.StructsTag(r, 结构体指针, OrmTagForStruct)
}
