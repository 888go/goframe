// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gjson

import (
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gutil"
)

// Interface 返回 json 值。 md5:63120da1c18a889c
func (j *Json) Interface() interface{} {
	if j == nil {
		return nil
	}
	j.mu.RLock()
	defer j.mu.RUnlock()
	if j.p == nil {
		return nil
	}
	return *(j.p)
}

// Var将json值作为*gvar.Var返回。 md5:fcfd99090165a7b5
func (j *Json) Var() *gvar.Var {
	return gvar.New(j.Interface())
}

// IsNil 检查指针 `j` 所指向的值是否为 nil。 md5:669f952cb58d64ce
func (j *Json) IsNil() bool {
	if j == nil {
		return true
	}
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.p == nil || *(j.p) == nil
}

// Get 通过指定的`pattern`获取并返回值。
// 如果`pattern`为"."，它将返回当前Json对象的所有值。
// 如果根据`pattern`没有找到值，它将返回nil。
//
// 我们也可以在`pattern`中通过索引数字来访问切片的元素，例如：
// "list.10", "array.0.name", "array.0.1.id"。
//
// 如果没有为`pattern`找到值，它将返回由`def`指定的默认值。
// md5:f56f76d635296903
func (j *Json) Get(pattern string, def ...interface{}) *gvar.Var {
	if j == nil {
		return nil
	}
	j.mu.RLock()
	defer j.mu.RUnlock()

		// 如果pattern为空，它将返回nil。 md5:8e2a6f56affd353a
	if pattern == "" {
		return nil
	}

	result := j.getPointerByPattern(pattern)
	if result != nil {
		return gvar.New(*result)
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}

// GetJson 通过指定的`pattern`获取值，并将其转换为一个非并发安全的Json对象。
// md5:e0f60541ee6017b5
func (j *Json) GetJson(pattern string, def ...interface{}) *Json {
	return New(j.Get(pattern, def...).Val())
}

// GetJsons 通过指定的`pattern`获取值，并将其转换为一个不并发安全的Json对象切片。
// md5:1bd75964e1b32ed2
func (j *Json) GetJsons(pattern string, def ...interface{}) []*Json {
	array := j.Get(pattern, def...).Array()
	if len(array) > 0 {
		jsonSlice := make([]*Json, len(array))
		for i := 0; i < len(array); i++ {
			jsonSlice[i] = New(array[i])
		}
		return jsonSlice
	}
	return nil
}

// GetJsonMap 通过指定的`pattern`获取值，
// 并将其转换为非并发安全的Json对象映射。
// md5:d549d238d186a4e0
func (j *Json) GetJsonMap(pattern string, def ...interface{}) map[string]*Json {
	m := j.Get(pattern, def...).Map()
	if len(m) > 0 {
		jsonMap := make(map[string]*Json, len(m))
		for k, v := range m {
			jsonMap[k] = New(v)
		}
		return jsonMap
	}
	return nil
}

// Set 使用指定的 `pattern` 设置值。
// 它支持通过字符分隔符（默认为'.'）进行层次数据访问。
// md5:85400f8aa43895d6
func (j *Json) Set(pattern string, value interface{}) error {
	return j.setValue(pattern, value, false)
}

// MustSet 执行与 Set 相同的操作，但如果发生任何错误，它将引发恐慌。 md5:89753cb5f56f60cc
func (j *Json) MustSet(pattern string, value interface{}) {
	if err := j.Set(pattern, value); err != nil {
		panic(err)
	}
}

// Remove 删除具有指定`pattern`的值。它支持通过字符分隔符（默认为`.`）进行层次数据访问。
// md5:a8bd1b8b0e8d7d8e
func (j *Json) Remove(pattern string) error {
	return j.setValue(pattern, nil, true)
}

// MustRemove 的行为与 Remove 相同，但如果发生任何错误，它会直接 panic。 md5:ad4ac7324486398a
func (j *Json) MustRemove(pattern string) {
	if err := j.Remove(pattern); err != nil {
		panic(err)
	}
}

// Contains 检查是否存在指定的 `pattern` 值。 md5:4f248b6aebb74d05
func (j *Json) Contains(pattern string) bool {
	return j.Get(pattern) != nil
}

// Len 返回由指定 `pattern` 定义的值的长度/大小。目标值应该是切片或映射类型。如果找不到目标值或者其类型无效，它将返回 -1。
// md5:f929eb27a0ef1a36
func (j *Json) Len(pattern string) int {
	p := j.getPointerByPattern(pattern)
	if p != nil {
		switch (*p).(type) {
		case map[string]interface{}:
			return len((*p).(map[string]interface{}))
		case []interface{}:
			return len((*p).([]interface{}))
		default:
			return -1
		}
	}
	return -1
}

// Append 将指定的 `pattern` 所引用的值追加到目标值（应该是切片类型）中。
// md5:5b8e7f4c493419ba
func (j *Json) Append(pattern string, value interface{}) error {
	p := j.getPointerByPattern(pattern)
	if p == nil || *p == nil {
		if pattern == "." {
			return j.Set("0", value)
		}
		return j.Set(fmt.Sprintf("%s.0", pattern), value)
	}
	switch (*p).(type) {
	case []interface{}:
		if pattern == "." {
			return j.Set(fmt.Sprintf("%d", len((*p).([]interface{}))), value)
		}
		return j.Set(fmt.Sprintf("%s.%d", pattern, len((*p).([]interface{}))), value)
	}
	return gerror.NewCodef(gcode.CodeInvalidParameter, "invalid variable type of %s", pattern)
}

// MustAppend 的行为与 Append 相同，但如果发生任何错误，它会直接 panic。 md5:3a0acd5a244f264f
func (j *Json) MustAppend(pattern string, value interface{}) {
	if err := j.Append(pattern, value); err != nil {
		panic(err)
	}
}

// Map 将当前的 Json 对象转换为 map[string]interface{}。
// 如果转换失败，它将返回 nil。
// md5:599d2c152000d26b
func (j *Json) Map() map[string]interface{} {
	return j.Var().Map()
}

// Array 将当前Json对象转换为 []interface{} 类型。如果转换失败，它将返回nil。
// md5:8b3042473c46995f
func (j *Json) Array() []interface{} {
	return j.Var().Array()
}

// Scan会根据参数类型自动调用Struct或Structs函数，并通过`pointer`实现转换。
// md5:afdb5ab720fddc3b
func (j *Json) Scan(pointer interface{}, mapping ...map[string]string) error {
	return j.Var().Scan(pointer, mapping...)
}

// Dump 打印当前的Json对象，使其更便于人工阅读。 md5:c8c6bbdb40fa6383
func (j *Json) Dump() {
	if j == nil {
		return
	}
	j.mu.RLock()
	defer j.mu.RUnlock()
	if j.p == nil {
		return
	}
	gutil.Dump(*j.p)
}
