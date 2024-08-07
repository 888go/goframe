// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"database/sql"
	"math"

	gvar "github.com/888go/goframe/container/gvar"
	gjson "github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/internal/empty"
	gconv "github.com/888go/goframe/util/gconv"
)

// X是否为空 检查 `r` 是否为空，然后返回结果。 md5:4ee28a47e769cceb
func (r Result) X是否为空() bool {
	return r == nil || r.X取数量() == 0
}

// X取数量 返回结果列表的长度。 md5:9abccfc01a850f4f
func (r Result) X取数量() int {
	return len(r)
}

// Size别名 是函数 Len 的别名。 md5:4cfc93cb64eff9b5
func (r Result) Size别名() int {
	return r.X取数量()
}

// X分割 将一个 Result 分割成多个 Result，
// 每个数组的大小由 `size` 决定。
// 最后一块可能包含少于 size 个元素。
// md5:e1e9bbb7e5ba1969
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

// X取json 将 `r` 转换为JSON格式的内容。 md5:60a0626b0a333d14
func (r Result) X取json() string {
	content, _ := gjson.X创建(r.X取Map切片()).X取json文本()
	return content
}

// X取xml 将 `r` 转换为 XML 格式的内容。 md5:31a335fedb874d26
func (r Result) X取xml(根标记 ...string) string {
	content, _ := gjson.X创建(r.X取Map切片()).X取xml文本(根标记...)
	return content
}

// X取Map切片 将 `r` 转换为一个 X取Map切片。 md5:ee79a42f10af264e
func (r Result) X取Map切片() Map切片 {
	list := make(Map切片, len(r))
	for k, v := range r {
		list[k] = v.X取Map()
	}
	return list
}

// X取字段切片 用于获取并返回指定列的值作为切片。
// 参数 `field` 是可选的，如果列字段只有一个。如果未给定 `field` 参数，其默认值为 `Result` 中第一条项的第一个字段名。
// md5:f3e0a3bab6043d80
func (r Result) X取字段切片(字段名称 ...string) []Value {
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

// X取字段Map泛型类 将 `r` 转换为一个 map[string]Value，其中的键由 `key` 指定。
// 注意，项目值可能为切片类型。
// md5:0c805cb25cfa56ff
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
			m[k] = gvar.X创建(v)
		} else {
			m[k] = gvar.X创建(v[0])
		}
	}
	return m
}

// X取字段MapStr 将 `r` 转换为一个键为指定字符串的 map[string]Map。 md5:43bb233c139ab262
func (r Result) X取字段MapStr(字段名称 string) map[string]Map {
	m := make(map[string]Map)
	for _, item := range r {
		if v, ok := item[字段名称]; ok {
			m[v.String()] = item.X取Map()
		}
	}
	return m
}

// X取字段MapInt 将 `r` 转换为一个 map[int]Map，其中的键由 `key` 指定。 md5:2c63f4f80e8a0e1b
func (r Result) X取字段MapInt(字段名称 string) map[int]Map {
	m := make(map[int]Map)
	for _, item := range r {
		if v, ok := item[字段名称]; ok {
			m[v.X取整数()] = item.X取Map()
		}
	}
	return m
}

// X取字段MapUint 将 `r` 转换为一个 map，其中的键是通过 `key` 指定的 uint 类型。 md5:0597073b149b7e00
func (r Result) X取字段MapUint(字段名称 string) map[uint]Map {
	m := make(map[uint]Map)
	for _, item := range r {
		if v, ok := item[字段名称]; ok {
			m[v.X取正整数()] = item.X取Map()
		}
	}
	return m
}

// RecordKeyStr 将 `r` 转换为一个 map[string]Record，其中的键由 `key` 指定。 md5:6eaa1193e5507d8a
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
func (r Result) RecordKeyInt(key string) map[int]Record {
	m := make(map[int]Record)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.X取整数()] = item
		}
	}
	return m
}

// RecordKeyUint 将 `r` 转换为一个以指定的 `key` 作为键的 [uint]Record 映射。 md5:26ce469215f5d9c2
func (r Result) RecordKeyUint(key string) map[uint]Record {
	m := make(map[uint]Record)
	for _, item := range r {
		if v, ok := item[key]; ok {
			m[v.X取正整数()] = item
		}
	}
	return m
}

// X取切片结构体指针 将 `r` 转换为结构体切片。
// 注意参数 `pointer` 的类型应该是 `*[]struct` 或 `*[]*struct`。
// md5:fef766b4997dca03
func (r Result) X取切片结构体指针(结构体指针 interface{}) (错误 error) {
		// 如果结果为空且目标指针不为空，则返回错误。 md5:74dffcb96270ed89
	if r.X是否为空() {
		if !empty.IsEmpty(结构体指针, true) {
			return sql.ErrNoRows
		}
		return nil
	}
	return gconv.StructsTag(r, 结构体指针, OrmTagForStruct)
}
