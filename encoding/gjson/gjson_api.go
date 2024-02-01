// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gjson
import (
	"fmt"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/util/gutil"
	)
// Interface 返回 JSON 值。
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

// Var 返回json值作为*gvar.Var类型。
func (j *Json) Var() *gvar.Var {
	return gvar.New(j.Interface())
}

// IsNil 检查通过 `j` 指向的值是否为 nil。
func (j *Json) IsNil() bool {
	if j == nil {
		return true
	}
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.p == nil || *(j.p) == nil
}

// Get 方法通过指定的`pattern`获取并返回值。
// 如果 `pattern` 为"."，则返回当前Json对象的所有值。
// 如果根据`pattern`未找到值，则返回nil。
//
// 我们还可以通过在`pattern`中指定切片的索引号来访问切片项，例如：
// "list.10", "array.0.name", "array.0.1.id"。
//
// 如果未找到`pattern`对应的值，将返回由`def`指定的默认值。
func (j *Json) Get(pattern string, def ...interface{}) *gvar.Var {
	if j == nil {
		return nil
	}
	j.mu.RLock()
	defer j.mu.RUnlock()

	// 如果模式为空，则返回nil。
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

// GetJson 通过指定的 `pattern` 获取值，
// 并将其转换为一个非并发安全的 Json 对象。
func (j *Json) GetJson(pattern string, def ...interface{}) *Json {
	return New(j.Get(pattern, def...).Val())
}

// GetJsons 通过指定的 `pattern` 获取值，
// 并将其转换为一个非并发安全的 Json 对象切片。
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

// GetJsonMap 通过指定的 `pattern` 获取值，
// 并将其转换为一个非并发安全的 Json 对象映射。
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

// Set 使用指定的`pattern`设置值。
// 它支持通过默认为'.'的字符分隔符进行层级数据访问。
func (j *Json) Set(pattern string, value interface{}) error {
	return j.setValue(pattern, value, false)
}

// MustSet 表现如同 Set，但当发生任何错误时会触发panic（异常）。
func (j *Json) MustSet(pattern string, value interface{}) {
	if err := j.Set(pattern, value); err != nil {
		panic(err)
	}
}

// Remove 删除具有指定 `pattern` 的值。
// 它支持通过默认为 '.' 的字符分隔符进行层级数据访问。
func (j *Json) Remove(pattern string) error {
	return j.setValue(pattern, nil, true)
}

// MustRemove 的行为与 Remove 相同，但是当发生任何错误时，它会触发 panic（异常）。
func (j *Json) MustRemove(pattern string) {
	if err := j.Remove(pattern); err != nil {
		panic(err)
	}
}

// Contains 检查指定的 `pattern` 是否存在对应的值。
func (j *Json) Contains(pattern string) bool {
	return j.Get(pattern) != nil
}

// Len 通过指定的 `pattern` 返回值的长度/大小。
// 通过 `pattern` 所获取的目标值应为切片或字典类型。
// 若目标值未找到，或者其类型无效，则返回 -1。
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

// Append 函数通过指定的 `pattern` 追加值到目标值中。
// 通过 `pattern` 所获取的目标值应当是切片类型。
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

// MustAppend 的行为与 Append 相同，但当发生任何错误时，它会触发panic（异常）。
func (j *Json) MustAppend(pattern string, value interface{}) {
	if err := j.Append(pattern, value); err != nil {
		panic(err)
	}
}

// Map 将当前的 Json 对象转换为 map[string]interface{} 类型。
// 如果转换失败，则返回 nil。
func (j *Json) Map() map[string]interface{} {
	return j.Var().Map()
}

// Array 将当前Json对象转换为 []interface{} 类型的切片。
// 如果转换失败，则返回nil。
func (j *Json) Array() []interface{} {
	return j.Var().Array()
}

// Scan 会根据参数 `pointer` 的类型自动调用 Struct 或 Structs 函数来实现转换功能。
func (j *Json) Scan(pointer interface{}, mapping ...map[string]string) error {
	return j.Var().Scan(pointer, mapping...)
}

// Dump 打印当前Json对象，使其更易于人工阅读。
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
