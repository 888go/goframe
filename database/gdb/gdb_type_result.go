// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb
import (
	"database/sql"
	"math"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/util/gconv"
	)
// IsEmpty 检查并返回 `r` 是否为空。
func (r Result) IsEmpty() bool {
	return r == nil || r.Len() == 0
}

// Len 返回结果列表的长度。
func (r Result) Len() int {
	return len(r)
}

// Size 是函数 Len 的别名。
func (r Result) Size() int {
	return r.Len()
}

// Chunk 将一个 Result 切分成多个 Result，
// 每个数组的大小由 `size` 决定。
// 最后一块切片可能包含少于 size 个元素。
func (r Result) Chunk(size int) []Result {
	if size < 1 {
		return nil
	}
	length := len(r)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n []Result
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, r[i*size:end])
		i++
	}
	return n
}

// Json 将 `r` 转换为 JSON 格式的内容。
func (r Result) Json() string {
	content, _ := gjson.New(r.List()).ToJsonString()
	return content
}

// Xml 将`r`转换为XML格式的内容。
func (r Result) Xml(rootTag ...string) string {
	content, _ := gjson.New(r.List()).ToXmlString(rootTag...)
	return content
}

// List 将 `r` 转换为一个 List。
func (r Result) List() List {
	list := make(List, len(r))
	for k, v := range r {
		list[k] = v.Map()
	}
	return list
}

// Array 根据指定列字段获取并返回其值作为一个切片。
// 当列字段只有一个时，参数 `field` 可选。
// 如果未给出参数 `field`，则默认的 `field` 为 `Result` 中第一项的第一个字段名。
func (r Result) Array(field ...string) []Value {
	array := make([]Value, len(r))
	if len(r) == 0 {
		return array
	}
	key := ""
	if len(field) > 0 && field[0] != "" {
		key = field[0]
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
func (r Result) MapKeyValue(key string) map[string]Value {
	var (
		s              string
		m              = make(map[string]Value)
		tempMap        = make(map[string][]interface{})
		hasMultiValues bool
	)
	for _, item := range r {
		if k, ok := item[key]; ok {
			s = k.String()
			tempMap[s] = append(tempMap[s], item)
			if len(tempMap[s]) > 1 {
				hasMultiValues = true
			}
		}
	}
	for k, v := range tempMap {
		if hasMultiValues {
			m[k] = gvar.New(v)
		} else {
			m[k] = gvar.New(v[0])
		}
	}
	return m
}

// MapKeyStr 将 `r` 转换为一个 map[string]Map 类型的映射，其中的键由 `key` 指定。
func (r Result) MapKeyStr(key string) map[string]Map {
	m := make(map[string]Map)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.String()] = item.Map()
		}
	}
	return m
}

// MapKeyInt 将 `r` 转换为一个映射 map[int]Map，其中键由 `key` 指定。
// （注：这里可能需要上下文信息，对于 `Map` 类型没有明确说明，所以翻译时假设它是一个已知的类型名。如果 `Map` 是自定义类型或有特殊含义，请替换为实际含义。）
func (r Result) MapKeyInt(key string) map[int]Map {
	m := make(map[int]Map)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.Int()] = item.Map()
		}
	}
	return m
}

// MapKeyUint 将`r`转换为一个map[uint]Map类型，其中键由`key`指定。
func (r Result) MapKeyUint(key string) map[uint]Map {
	m := make(map[uint]Map)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.Uint()] = item.Map()
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
			m[v.Int()] = item
		}
	}
	return m
}

// RecordKeyUint 将 `r` 转换为一个 map[uint]Record 类型的映射，其中键由 `key` 指定。
func (r Result) RecordKeyUint(key string) map[uint]Record {
	m := make(map[uint]Record)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.Uint()] = item
		}
	}
	return m
}

// Structs 将 `r` 转换为结构体切片。
// 注意参数 `pointer` 应该是指向结构体切片的指针类型，即 *[]struct 或 *[]*struct。
func (r Result) Structs(pointer interface{}) (err error) {
	// 如果结果为空且目标指针不为空，则返回错误。
	if r.IsEmpty() {
		if !empty.IsEmpty(pointer, true) {
			return sql.ErrNoRows
		}
		return nil
	}
	return gconv.StructsTag(r, pointer, OrmTagForStruct)
}
