// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtest

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/debug/gdebug"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	pathFilterKey = "/test/gtest/gtest"
)

// C 创建一个单元测试用例。
// 参数 `t` 是标准库 (*testing.T) 的 testing.T 指针。
// 参数 `f` 是用于单元测试的闭包函数。
// md5:0a3ae380343ea962
func C(t *testing.T, f func(t *T)) {
	defer func() {
		if err := recover(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n%s", err, gdebug.StackWithFilter([]string{pathFilterKey}))
			t.Fail()
		}
	}()
	f(&T{t})
}

// Assert 检查 `value` 和 `expect` 是否相等。. md5:eaeea7c4fe0d764e
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

// AssertEQ 检查 `value` 和 `expect` 是否相等，包括它们的类型。. md5:31097fa6b823a25a
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

// AssertNE 检查 `value` 和 `expect` 是否不相等。. md5:418e91b330bc944f
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

// AssertNQ 检查 `value` 和 `expect` 是否不相等，包括它们的类型。. md5:bb13af00897290db
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
// 注意，只有字符串、整数和浮点数类型能使用 AssertGT 进行比较，
// 其他类型是无效的。
// md5:647270894818c6c7
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
// 请注意，只有字符串、整数和浮点数类型可以使用 AssertGE 进行比较，其他类型是无效的。
// md5:3227e007891ed72e
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
// 注意，只有字符串、整数和浮点类型可以通过 AssertLT 进行比较，其他类型无效。
// md5:784a9db44c03122b
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

// AssertLE 检查 `value` 是否小于或等于 `expect`。
// 请注意，只有字符串、整数和浮点类型可以通过 AssertLTE 进行比较，其他类型的值是无效的。
// md5:bca4df91bef4e152
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
// `expect` 应该是一个切片，
// 但是 `value` 可以是切片或基本类型变量。
// TODO: 添加对 map 的支持。
// TODO: gconv.Strings(0) 不应该转换为 `[0]`。
// md5:d8391e0c6cba6480
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
	case reflect.String:
		var (
			valueStr  = gconv.String(value)
			expectStr = gconv.String(expect)
		)
		passed = gstr.Contains(expectStr, valueStr)
	default:
		panic(fmt.Sprintf(`[ASSERT] INVALID EXPECT VALUE TYPE: %v`, expectKind))
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v IN %v`, value, expect))
	}
}

// AssertNI 检查 `value` 不在 `expect` 中。
// `expect` 应该是一个切片，
// 但是 `value` 可以是切片或基本类型变量。
// TODO 增加对 map 的支持。
// md5:483febd56930eb64
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
	case reflect.String:
		var (
			valueStr  = gconv.String(value)
			expectStr = gconv.String(expect)
		)
		passed = !gstr.Contains(expectStr, valueStr)
	default:
		panic(fmt.Sprintf(`[ASSERT] INVALID EXPECT VALUE TYPE: %v`, expectKind))
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v NOT IN %v`, value, expect))
	}
}

// 使用给定的`message`引发错误恐慌。. md5:6ddb84d91c681d1f
func Error(message ...interface{}) {
	panic(fmt.Sprintf("[ERROR] %s", fmt.Sprint(message...)))
}

// Fatal 将 `message` 打印到 stderr 并退出进程。. md5:15e177961f66ebe7
func Fatal(message ...interface{}) {
	_, _ = fmt.Fprintf(
		os.Stderr, "[FATAL] %s\n%s", fmt.Sprint(message...),
		gdebug.StackWithFilter([]string{pathFilterKey}),
	)
	os.Exit(1)
}

// compareMap 比较两个地图，如果它们相等则返回nil，否则返回错误。. md5:fd402375a76c3a4a
func compareMap(value, expect interface{}) error {
	var (
		rvValue  = reflect.ValueOf(value)
		rvExpect = reflect.ValueOf(expect)
	)
	if rvExpect.Kind() == reflect.Map {
		if rvValue.Kind() == reflect.Map {
			if rvExpect.Len() == rvValue.Len() {
// 将两个接口映射转换为相同类型以便进行比较。
// 直接使用rvValue.MapIndex(key).Interface() 当键的类型不一致时，会导致恐慌。
// md5:ae85735772c34002
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

// AssertNil 断言 `value` 为 nil。. md5:94a00206ff503e10
func AssertNil(value interface{}) {
	if empty.IsNil(value) {
		return
	}
	if err, ok := value.(error); ok {
		panic(fmt.Sprintf(`%+v`, err))
	}
	Assert(value, nil)
}

// DataPath获取并返回当前包的测试数据路径，仅用于单元测试。
// 可选参数`names`指定了子文件夹/子文件，将与当前系统的分隔符连接，并与路径一起返回。
// md5:55efb430c9f8a73f
func DataPath(names ...string) string {
	_, path, _ := gdebug.CallerWithFilter([]string{pathFilterKey})
	path = filepath.Dir(path) + string(filepath.Separator) + "testdata"
	for _, name := range names {
		path += string(filepath.Separator) + name
	}
	return path
}

// DataContent 从当前包的特定测试数据路径中检索并返回文件内容. md5:26224495ddbd389e
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
