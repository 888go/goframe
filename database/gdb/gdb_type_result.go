// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"database/sql"
	"math"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/util/gconv"
)

// IsEmpty 检查 `r` 是否为空，然后返回结果。 md5:4ee28a47e769cceb
// ff:是否为空
// r:
func (r Result) IsEmpty() bool {
	return r == nil || r.Len() == 0
}

// Len 返回结果列表的长度。 md5:9abccfc01a850f4f
// ff:取数量
// r:
func (r Result) Len() int {
	return len(r)
}

// Size 是函数 Len 的别名。 md5:4cfc93cb64eff9b5
// ff:Size别名
// r:
func (r Result) Size() int {
	return r.Len()
}

// Chunk 将一个 Result 分割成多个 Result，
// 每个数组的大小由 `size` 决定。
// 最后一块可能包含少于 size 个元素。
// md5:e1e9bbb7e5ba1969
// ff:分割
// r:
// size:数量
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

// Json 将 `r` 转换为JSON格式的内容。 md5:60a0626b0a333d14
// ff:取json
// r:
func (r Result) Json() string {
	content, _ := gjson.New(r.List()).ToJsonString()
	return content
}

// Xml 将 `r` 转换为 XML 格式的内容。 md5:31a335fedb874d26
// ff:取xml
// r:
// rootTag:根标记
func (r Result) Xml(rootTag ...string) string {
	content, _ := gjson.New(r.List()).ToXmlString(rootTag...)
	return content
}

// List 将 `r` 转换为一个 List。 md5:ee79a42f10af264e
// ff:取Map切片
// r:
func (r Result) List() List {
	list := make(List, len(r))
	for k, v := range r {
		list[k] = v.Map()
	}
	return list
}

// Array 用于获取并返回指定列的值作为切片。
// 参数 `field` 是可选的，如果列字段只有一个。如果未给定 `field` 参数，其默认值为 `Result` 中第一条项的第一个字段名。
// md5:f3e0a3bab6043d80
// ff:取字段切片
// r:
// field:字段名称
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

// MapKeyValue 将 `r` 转换为一个 map[string]Value，其中的键由 `key` 指定。
// 注意，项目值可能为切片类型。
// md5:0c805cb25cfa56ff
// ff:取字段Map泛型类
// r:
// key:字段名称
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

// MapKeyStr 将 `r` 转换为一个键为指定字符串的 map[string]Map。 md5:43bb233c139ab262
// ff:取字段MapStr
// r:
// key:字段名称
func (r Result) MapKeyStr(key string) map[string]Map {
	m := make(map[string]Map)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.String()] = item.Map()
		}
	}
	return m
}

// MapKeyInt 将 `r` 转换为一个 map[int]Map，其中的键由 `key` 指定。 md5:2c63f4f80e8a0e1b
// ff:取字段MapInt
// r:
// key:字段名称
func (r Result) MapKeyInt(key string) map[int]Map {
	m := make(map[int]Map)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.Int()] = item.Map()
		}
	}
	return m
}

// MapKeyUint 将 `r` 转换为一个 map，其中的键是通过 `key` 指定的 uint 类型。 md5:0597073b149b7e00
// ff:取字段MapUint
// r:
// key:字段名称
func (r Result) MapKeyUint(key string) map[uint]Map {
	m := make(map[uint]Map)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.Uint()] = item.Map()
		}
	}
	return m
}

// RecordKeyStr 将 `r` 转换为一个 map[string]Record，其中的键由 `key` 指定。 md5:6eaa1193e5507d8a
// ff:
// r:
// key:
func (r Result) RecordKeyStr(key string) map[string]Record {
	m := make(map[string]Record)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.String()] = item
		}
	}
	return m
}

// RecordKeyInt 将 `r` 转换为一个映射[int]Record，其中键由 `key` 指定。 md5:0ebe0554d495cbae
// ff:
// r:
// key:
func (r Result) RecordKeyInt(key string) map[int]Record {
	m := make(map[int]Record)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.Int()] = item
		}
	}
	return m
}

// RecordKeyUint 将 `r` 转换为一个以指定的 `key` 作为键的 [uint]Record 映射。 md5:26ce469215f5d9c2
// ff:
// r:
// key:
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
// 注意参数 `pointer` 的类型应该是 `*[]struct` 或 `*[]*struct`。
// md5:fef766b4997dca03
// ff:取切片结构体指针
// r:
// pointer:结构体指针
// err:错误
func (r Result) Structs(pointer interface{}) (err error) {
	// 如果结果为空且目标指针不为空，则返回错误。 md5:74dffcb96270ed89
	if r.IsEmpty() {
		if !empty.IsEmpty(pointer, true) {
			return sql.ErrNoRows
		}
		return nil
	}
	return gconv.StructsTag(r, pointer, OrmTagForStruct)
}
