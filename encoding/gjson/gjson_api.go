// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类

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
func (j *Json) X取泛型类() *泛型类.Var {
	return 泛型类.X创建(j.Interface())
}

// IsNil 检查通过 `j` 指向的值是否为 nil。
func (j *Json) X是否为Nil() bool {
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
func (j *Json) X取值(表达式 string, 默认值 ...interface{}) *泛型类.Var {
	if j == nil {
		return nil
	}
	j.mu.RLock()
	defer j.mu.RUnlock()

	// 如果模式为空，则返回nil。
	if 表达式 == "" {
		return nil
	}

	result := j.getPointerByPattern(表达式)
	if result != nil {
		return 泛型类.X创建(*result)
	}
	if len(默认值) > 0 {
		return 泛型类.X创建(默认值[0])
	}
	return nil
}

// GetJson 通过指定的 `pattern` 获取值，
// 并将其转换为一个非并发安全的 Json 对象。
func (j *Json) X取对象(表达式 string, 默认值 ...interface{}) *Json {
	return X创建(j.X取值(表达式, 默认值...).X取值())
}

// GetJsons 通过指定的 `pattern` 获取值，
// 并将其转换为一个非并发安全的 Json 对象切片。
func (j *Json) X取对象数组(表达式 string, 默认值 ...interface{}) []*Json {
	array := j.X取值(表达式, 默认值...).Array别名()
	if len(array) > 0 {
		jsonSlice := make([]*Json, len(array))
		for i := 0; i < len(array); i++ {
			jsonSlice[i] = X创建(array[i])
		}
		return jsonSlice
	}
	return nil
}

// GetJsonMap 通过指定的 `pattern` 获取值，
// 并将其转换为一个非并发安全的 Json 对象映射。
func (j *Json) X取对象Map(表达式 string, 默认值 ...interface{}) map[string]*Json {
	m := j.X取值(表达式, 默认值...).X取Map()
	if len(m) > 0 {
		jsonMap := make(map[string]*Json, len(m))
		for k, v := range m {
			jsonMap[k] = X创建(v)
		}
		return jsonMap
	}
	return nil
}

// Set 使用指定的`pattern`设置值。
// 它支持通过默认为'.'的字符分隔符进行层级数据访问。
func (j *Json) X设置值(表达式 string, 值 interface{}) error {
	return j.setValue(表达式, 值, false)
}

// MustSet 表现如同 Set，但当发生任何错误时会触发panic（异常）。
func (j *Json) X设置值PANI(表达式 string, 值 interface{}) {
	if err := j.X设置值(表达式, 值); err != nil {
		panic(err)
	}
}

// Remove 删除具有指定 `pattern` 的值。
// 它支持通过默认为 '.' 的字符分隔符进行层级数据访问。
func (j *Json) X删除(表达式 string) error {
	return j.setValue(表达式, nil, true)
}

// MustRemove 的行为与 Remove 相同，但是当发生任何错误时，它会触发 panic（异常）。
func (j *Json) X删除PANI(表达式 string) {
	if err := j.X删除(表达式); err != nil {
		panic(err)
	}
}

// Contains 检查指定的 `pattern` 是否存在对应的值。
func (j *Json) X是否存在(表达式 string) bool {
	return j.X取值(表达式) != nil
}

// Len 通过指定的 `pattern` 返回值的长度/大小。
// 通过 `pattern` 所获取的目标值应为切片或字典类型。
// 若目标值未找到，或者其类型无效，则返回 -1。
func (j *Json) X取长度(表达式 string) int {
	p := j.getPointerByPattern(表达式)
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
func (j *Json) X加入(表达式 string, 值 interface{}) error {
	p := j.getPointerByPattern(表达式)
	if p == nil || *p == nil {
		if 表达式 == "." {
			return j.X设置值("0", 值)
		}
		return j.X设置值(fmt.Sprintf("%s.0", 表达式), 值)
	}
	switch (*p).(type) {
	case []interface{}:
		if 表达式 == "." {
			return j.X设置值(fmt.Sprintf("%d", len((*p).([]interface{}))), 值)
		}
		return j.X设置值(fmt.Sprintf("%s.%d", 表达式, len((*p).([]interface{}))), 值)
	}
	return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, "invalid variable type of %s", 表达式)
}

// MustAppend 的行为与 Append 相同，但当发生任何错误时，它会触发panic（异常）。
func (j *Json) X加入PANI(表达式 string, 值 interface{}) {
	if err := j.X加入(表达式, 值); err != nil {
		panic(err)
	}
}

// Map 将当前的 Json 对象转换为 map[string]interface{} 类型。
// 如果转换失败，则返回 nil。
func (j *Json) X取Map() map[string]interface{} {
	return j.X取泛型类().X取Map()
}

// Array 将当前Json对象转换为 []interface{} 类型的切片。
// 如果转换失败，则返回nil。
func (j *Json) X取数组() []interface{} {
	return j.X取泛型类().Array别名()
}

// Scan 会根据参数 `pointer` 的类型自动调用 Struct 或 Structs 函数来实现转换功能。
func (j *Json) X取结构体指针(结构体指针 interface{}, 名称映射 ...map[string]string) error {
	return j.X取泛型类().X取结构体指针(结构体指针, 名称映射...)
}

// Dump 打印当前Json对象，使其更易于人工阅读。
func (j *Json) X调试输出() {
	if j == nil {
		return
	}
	j.mu.RLock()
	defer j.mu.RUnlock()
	if j.p == nil {
		return
	}
	工具类.X调试输出(*j.p)
}
