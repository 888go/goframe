// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtest
import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/util/gconv"
	)
const (
	pathFilterKey = "/test/gtest/gtest"
)

// C 创建一个单元测试用例。
// 参数 `t` 是指向标准库 testing.T 的指针（*testing.T）。
// 参数 `f` 是用于单元测试用例的闭包函数。
func C(t *testing.T, f func(t *T)) {
	defer func() {
		if err := recover(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n%s", err, gdebug.StackWithFilter([]string{pathFilterKey}))
			t.Fail()
		}
	}()
	f(&T{t})
}

// Assert 检查 `value` 和 `expect` 是否相等。
func Assert(value, expect interface{}) {
	rvExpect := reflect.ValueOf(expect)
	if empty.IsNil(value) {
		value = nil
	}
	if rvExpect.Kind() == reflect.Map {
		if err := compareMap(value, expect); err != nil {
			panic(err)
		}
		return
	}
	var (
		strValue  = gconv.String(value)
		strExpect = gconv.String(expect)
	)
	if strValue != strExpect {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v == %v`, strValue, strExpect))
	}
}

// AssertEQ 检查 `value` 和 `expect` 是否相等，包括它们的 TYPE（类型）。
func AssertEQ(value, expect interface{}) {
	// Value assert.
	rvExpect := reflect.ValueOf(expect)
	if empty.IsNil(value) {
		value = nil
	}
	if rvExpect.Kind() == reflect.Map {
		if err := compareMap(value, expect); err != nil {
			panic(err)
		}
		return
	}
	strValue := gconv.String(value)
	strExpect := gconv.String(expect)
	if strValue != strExpect {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v == %v`, strValue, strExpect))
	}
	// Type assert.
	t1 := reflect.TypeOf(value)
	t2 := reflect.TypeOf(expect)
	if t1 != t2 {
		panic(fmt.Sprintf(`[ASSERT] EXPECT TYPE %v[%v] == %v[%v]`, strValue, t1, strExpect, t2))
	}
}

// AssertNE 检查 `value` 和 `expect` 是否不相等。
func AssertNE(value, expect interface{}) {
	rvExpect := reflect.ValueOf(expect)
	if empty.IsNil(value) {
		value = nil
	}
	if rvExpect.Kind() == reflect.Map {
		if err := compareMap(value, expect); err == nil {
			panic(fmt.Sprintf(`[ASSERT] EXPECT %v != %v`, value, expect))
		}
		return
	}
	var (
		strValue  = gconv.String(value)
		strExpect = gconv.String(expect)
	)
	if strValue == strExpect {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v != %v`, strValue, strExpect))
	}
}

// AssertNQ 检查 `value` 和 `expect` 是否不相等，包括它们的类型。
func AssertNQ(value, expect interface{}) {
	// Type assert.
	t1 := reflect.TypeOf(value)
	t2 := reflect.TypeOf(expect)
	if t1 == t2 {
		panic(
			fmt.Sprintf(
				`[ASSERT] EXPECT TYPE %v[%v] != %v[%v]`,
				gconv.String(value), t1, gconv.String(expect), t2,
			),
		)
	}
	// Value assert.
	AssertNE(value, expect)
}

// AssertGT 检查 `value` 是否大于 `expect`。
// 注意，只有字符串、整数和浮点类型可以通过 AssertGT 进行比较，
// 其他类型是无效的。
func AssertGT(value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = gconv.String(value) > gconv.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = gconv.Int(value) > gconv.Int(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = gconv.Uint(value) > gconv.Uint(expect)

	case reflect.Float32, reflect.Float64:
		passed = gconv.Float64(value) > gconv.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v > %v`, value, expect))
	}
}

// AssertGE 检查 `value` 是否大于或等于 `expect`。
// 注意，只有字符串、整数和浮点类型可以使用 AssertGTE 进行比较，其他类型无效。
func AssertGE(value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = gconv.String(value) >= gconv.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = gconv.Int64(value) >= gconv.Int64(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = gconv.Uint64(value) >= gconv.Uint64(expect)

	case reflect.Float32, reflect.Float64:
		passed = gconv.Float64(value) >= gconv.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(
			`[ASSERT] EXPECT %v(%v) >= %v(%v)`,
			value, reflect.ValueOf(value).Kind(),
			expect, reflect.ValueOf(expect).Kind(),
		))
	}
}

// AssertLT 检查 `value` 是否小于等于 `expect`。
// 注意，只有字符串、整数和浮点类型可以使用 AssertLT 进行比较，
// 其他类型无效。
func AssertLT(value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = gconv.String(value) < gconv.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = gconv.Int(value) < gconv.Int(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = gconv.Uint(value) < gconv.Uint(expect)

	case reflect.Float32, reflect.Float64:
		passed = gconv.Float64(value) < gconv.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v < %v`, value, expect))
	}
}

// AssertLE 检查 `value` 是否小于等于 `expect`。
// 注意，只有字符串、整数和浮点类型可以通过 AssertLTE 进行比较，
// 其他类型是无效的。
func AssertLE(value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = gconv.String(value) <= gconv.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = gconv.Int(value) <= gconv.Int(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = gconv.Uint(value) <= gconv.Uint(expect)

	case reflect.Float32, reflect.Float64:
		passed = gconv.Float64(value) <= gconv.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v <= %v`, value, expect))
	}
}

// AssertIN 检查 `value` 是否在 `expect` 中。
// `expect` 应该是一个切片类型，
// 但 `value` 可以是切片类型或基本类型变量。
// TODO: 添加对 map 类型的支持。
// TODO: gconv.Strings(0) 的结果不是 [0]
func AssertIN(value, expect interface{}) {
	var (
		passed     = true
		expectKind = reflect.ValueOf(expect).Kind()
	)
	switch expectKind {
	case reflect.Slice, reflect.Array:
		expectSlice := gconv.Strings(expect)
		for _, v1 := range gconv.Strings(value) {
			result := false
			for _, v2 := range expectSlice {
				if v1 == v2 {
					result = true
					break
				}
			}
			if !result {
				passed = false
				break
			}
		}
	default:
		panic(fmt.Sprintf(`[ASSERT] INVALID EXPECT VALUE TYPE: %v`, expectKind))
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v IN %v`, value, expect))
	}
}

// AssertNI 检查 `value` 是否不在 `expect` 中。
// `expect` 应该是一个切片，
// 但 `value` 可以是切片或基本类型变量。
// TODO: 添加对 map 的支持。
func AssertNI(value, expect interface{}) {
	var (
		passed     = true
		expectKind = reflect.ValueOf(expect).Kind()
	)
	switch expectKind {
	case reflect.Slice, reflect.Array:
		for _, v1 := range gconv.Strings(value) {
			result := true
			for _, v2 := range gconv.Strings(expect) {
				if v1 == v2 {
					result = false
					break
				}
			}
			if !result {
				passed = false
				break
			}
		}
	default:
		panic(fmt.Sprintf(`[ASSERT] INVALID EXPECT VALUE TYPE: %v`, expectKind))
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v NOT IN %v`, value, expect))
	}
}

// Error 使用给定的`message`引发panic异常。
func Error(message ...interface{}) {
	panic(fmt.Sprintf("[ERROR] %s", fmt.Sprint(message...)))
}

// Fatal将`message`打印到标准错误输出（stderr）并退出进程。
func Fatal(message ...interface{}) {
	_, _ = fmt.Fprintf(
		os.Stderr, "[FATAL] %s\n%s", fmt.Sprint(message...),
		gdebug.StackWithFilter([]string{pathFilterKey}),
	)
	os.Exit(1)
}

// compareMap 比较两个映射，如果它们相等则返回 nil，否则返回错误。
func compareMap(value, expect interface{}) error {
	var (
		rvValue  = reflect.ValueOf(value)
		rvExpect = reflect.ValueOf(expect)
	)
	if rvExpect.Kind() == reflect.Map {
		if rvValue.Kind() == reflect.Map {
			if rvExpect.Len() == rvValue.Len() {
// 将两个接口映射转换为同一类型以便进行比较。
// 若直接使用rvValue.MapIndex(key).Interface()，当键类型不一致时会触发 panic。
				mValue := make(map[string]string)
				mExpect := make(map[string]string)
				ksValue := rvValue.MapKeys()
				ksExpect := rvExpect.MapKeys()
				for _, key := range ksValue {
					mValue[gconv.String(key.Interface())] = gconv.String(rvValue.MapIndex(key).Interface())
				}
				for _, key := range ksExpect {
					mExpect[gconv.String(key.Interface())] = gconv.String(rvExpect.MapIndex(key).Interface())
				}
				for k, v := range mExpect {
					if v != mValue[k] {
						return fmt.Errorf(`[ASSERT] EXPECT VALUE map["%v"]:%v == map["%v"]:%v`+
							"\nGIVEN : %v\nEXPECT: %v", k, mValue[k], k, v, mValue, mExpect)
					}
				}
			} else {
				return fmt.Errorf(`[ASSERT] EXPECT MAP LENGTH %d == %d`, rvValue.Len(), rvExpect.Len())
			}
		} else {
			return fmt.Errorf(`[ASSERT] EXPECT VALUE TO BE A MAP, BUT GIVEN "%s"`, rvValue.Kind())
		}
	}
	return nil
}

// AssertNil 断言 `value` 为 nil。
func AssertNil(value interface{}) {
	if empty.IsNil(value) {
		return
	}
	if err, ok := value.(error); ok {
		panic(fmt.Sprintf(`%+v`, err))
	}
	Assert(value, nil)
}

// DataPath 获取并返回当前包的 testdata 路径，
// 该路径仅用于单元测试用例。
// 可选参数 `names` 指定子文件夹/子文件，
// 这些名称将与当前系统分隔符连接，并与路径一起返回。
func DataPath(names ...string) string {
	_, path, _ := gdebug.CallerWithFilter([]string{pathFilterKey})
	path = filepath.Dir(path) + string(filepath.Separator) + "testdata"
	for _, name := range names {
		path += string(filepath.Separator) + name
	}
	return path
}

// DataContent 函数用于获取并返回当前包中指定testdata路径下的文件内容
func DataContent(names ...string) string {
	path := DataPath(names...)
	if path != "" {
		data, err := os.ReadFile(path)
		if err == nil {
			return string(data)
		}
	}
	return ""
}
