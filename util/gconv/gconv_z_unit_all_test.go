// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类_test

import (
	"math"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/gvar"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

type iString interface {
	String() string
}

type S struct {
}

func (s S) String() string {
	return "22222"
}

type iError interface {
	Error() string
}

type S1 struct {
}

func (s1 S1) Error() string {
	return "22222"
}

func Test_Bool_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.AssertEQ(转换类.X取布尔(any), false)
		t.AssertEQ(转换类.X取布尔(false), false)
		t.AssertEQ(转换类.X取布尔(nil), false)
		t.AssertEQ(转换类.X取布尔(0), false)
		t.AssertEQ(转换类.X取布尔("0"), false)
		t.AssertEQ(转换类.X取布尔(""), false)
		t.AssertEQ(转换类.X取布尔("false"), false)
		t.AssertEQ(转换类.X取布尔("off"), false)
		t.AssertEQ(转换类.X取布尔([]byte{}), false)
		t.AssertEQ(转换类.X取布尔([]string{}), false)
		t.AssertEQ(转换类.X取布尔([2]int{1, 2}), true)
		t.AssertEQ(转换类.X取布尔([]interface{}{}), false)
		t.AssertEQ(转换类.X取布尔([]map[int]int{}), false)

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取布尔(countryCapitalMap), true)

		t.AssertEQ(转换类.X取布尔("1"), true)
		t.AssertEQ(转换类.X取布尔("on"), true)
		t.AssertEQ(转换类.X取布尔(1), true)
		t.AssertEQ(转换类.X取布尔(123.456), true)
		t.AssertEQ(转换类.X取布尔(boolStruct{}), true)
		t.AssertEQ(转换类.X取布尔(&boolStruct{}), true)
	})
}

func Test_Int_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.AssertEQ(转换类.X取整数(any), 0)
		t.AssertEQ(转换类.X取整数(false), 0)
		t.AssertEQ(转换类.X取整数(nil), 0)
		t.Assert(转换类.X取整数(nil), 0)
		t.AssertEQ(转换类.X取整数(0), 0)
		t.AssertEQ(转换类.X取整数("0"), 0)
		t.AssertEQ(转换类.X取整数(""), 0)
		t.AssertEQ(转换类.X取整数("false"), 0)
		t.AssertEQ(转换类.X取整数("off"), 0)
		t.AssertEQ(转换类.X取整数([]byte{}), 0)
		t.AssertEQ(转换类.X取整数([]string{}), 0)
		t.AssertEQ(转换类.X取整数([2]int{1, 2}), 0)
		t.AssertEQ(转换类.X取整数([]interface{}{}), 0)
		t.AssertEQ(转换类.X取整数([]map[int]int{}), 0)

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取整数(countryCapitalMap), 0)

		t.AssertEQ(转换类.X取整数("1"), 1)
		t.AssertEQ(转换类.X取整数("on"), 0)
		t.AssertEQ(转换类.X取整数(1), 1)
		t.AssertEQ(转换类.X取整数(123.456), 123)
		t.AssertEQ(转换类.X取整数(boolStruct{}), 0)
		t.AssertEQ(转换类.X取整数(&boolStruct{}), 0)
		t.AssertEQ(转换类.X取整数("NaN"), 0)
	})
}

func Test_Int8_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.Assert(转换类.X取整数8位(any), int8(0))
		t.AssertEQ(转换类.X取整数8位(false), int8(0))
		t.AssertEQ(转换类.X取整数8位(nil), int8(0))
		t.AssertEQ(转换类.X取整数8位(0), int8(0))
		t.AssertEQ(转换类.X取整数8位("0"), int8(0))
		t.AssertEQ(转换类.X取整数8位(""), int8(0))
		t.AssertEQ(转换类.X取整数8位("false"), int8(0))
		t.AssertEQ(转换类.X取整数8位("off"), int8(0))
		t.AssertEQ(转换类.X取整数8位([]byte{}), int8(0))
		t.AssertEQ(转换类.X取整数8位([]string{}), int8(0))
		t.AssertEQ(转换类.X取整数8位([2]int{1, 2}), int8(0))
		t.AssertEQ(转换类.X取整数8位([]interface{}{}), int8(0))
		t.AssertEQ(转换类.X取整数8位([]map[int]int{}), int8(0))

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取整数8位(countryCapitalMap), int8(0))

		t.AssertEQ(转换类.X取整数8位("1"), int8(1))
		t.AssertEQ(转换类.X取整数8位("on"), int8(0))
		t.AssertEQ(转换类.X取整数8位(int8(1)), int8(1))
		t.AssertEQ(转换类.X取整数8位(123.456), int8(123))
		t.AssertEQ(转换类.X取整数8位(boolStruct{}), int8(0))
		t.AssertEQ(转换类.X取整数8位(&boolStruct{}), int8(0))
		t.AssertEQ(转换类.X取整数8位("NaN"), int8(0))

	})
}

func Test_Int16_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.Assert(转换类.X取整数16位(any), int16(0))
		t.AssertEQ(转换类.X取整数16位(false), int16(0))
		t.AssertEQ(转换类.X取整数16位(nil), int16(0))
		t.AssertEQ(转换类.X取整数16位(0), int16(0))
		t.AssertEQ(转换类.X取整数16位("0"), int16(0))
		t.AssertEQ(转换类.X取整数16位(""), int16(0))
		t.AssertEQ(转换类.X取整数16位("false"), int16(0))
		t.AssertEQ(转换类.X取整数16位("off"), int16(0))
		t.AssertEQ(转换类.X取整数16位([]byte{}), int16(0))
		t.AssertEQ(转换类.X取整数16位([]string{}), int16(0))
		t.AssertEQ(转换类.X取整数16位([2]int{1, 2}), int16(0))
		t.AssertEQ(转换类.X取整数16位([]interface{}{}), int16(0))
		t.AssertEQ(转换类.X取整数16位([]map[int]int{}), int16(0))

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取整数16位(countryCapitalMap), int16(0))

		t.AssertEQ(转换类.X取整数16位("1"), int16(1))
		t.AssertEQ(转换类.X取整数16位("on"), int16(0))
		t.AssertEQ(转换类.X取整数16位(int16(1)), int16(1))
		t.AssertEQ(转换类.X取整数16位(123.456), int16(123))
		t.AssertEQ(转换类.X取整数16位(boolStruct{}), int16(0))
		t.AssertEQ(转换类.X取整数16位(&boolStruct{}), int16(0))
		t.AssertEQ(转换类.X取整数16位("NaN"), int16(0))
	})
}

func Test_Int32_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.Assert(转换类.X取整数32位(any), int32(0))
		t.AssertEQ(转换类.X取整数32位(false), int32(0))
		t.AssertEQ(转换类.X取整数32位(nil), int32(0))
		t.AssertEQ(转换类.X取整数32位(0), int32(0))
		t.AssertEQ(转换类.X取整数32位("0"), int32(0))
		t.AssertEQ(转换类.X取整数32位(""), int32(0))
		t.AssertEQ(转换类.X取整数32位("false"), int32(0))
		t.AssertEQ(转换类.X取整数32位("off"), int32(0))
		t.AssertEQ(转换类.X取整数32位([]byte{}), int32(0))
		t.AssertEQ(转换类.X取整数32位([]string{}), int32(0))
		t.AssertEQ(转换类.X取整数32位([2]int{1, 2}), int32(0))
		t.AssertEQ(转换类.X取整数32位([]interface{}{}), int32(0))
		t.AssertEQ(转换类.X取整数32位([]map[int]int{}), int32(0))

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取整数32位(countryCapitalMap), int32(0))

		t.AssertEQ(转换类.X取整数32位("1"), int32(1))
		t.AssertEQ(转换类.X取整数32位("on"), int32(0))
		t.AssertEQ(转换类.X取整数32位(int32(1)), int32(1))
		t.AssertEQ(转换类.X取整数32位(123.456), int32(123))
		t.AssertEQ(转换类.X取整数32位(boolStruct{}), int32(0))
		t.AssertEQ(转换类.X取整数32位(&boolStruct{}), int32(0))
		t.AssertEQ(转换类.X取整数32位("NaN"), int32(0))
	})
}

func Test_Int64_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.AssertEQ(转换类.X取整数64位("0x00e"), int64(14))
		t.Assert(转换类.X取整数64位("022"), int64(22))

		t.Assert(转换类.X取整数64位(any), int64(0))
		t.Assert(转换类.X取整数64位(true), 1)
		t.Assert(转换类.X取整数64位("1"), int64(1))
		t.Assert(转换类.X取整数64位("0"), int64(0))
		t.Assert(转换类.X取整数64位("X"), int64(0))
		t.Assert(转换类.X取整数64位("x"), int64(0))
		t.Assert(转换类.X取整数64位(int64(1)), int64(1))
		t.Assert(转换类.X取整数64位(int(0)), int64(0))
		t.Assert(转换类.X取整数64位(int8(0)), int64(0))
		t.Assert(转换类.X取整数64位(int16(0)), int64(0))
		t.Assert(转换类.X取整数64位(int32(0)), int64(0))
		t.Assert(转换类.X取整数64位(uint64(0)), int64(0))
		t.Assert(转换类.X取整数64位(uint32(0)), int64(0))
		t.Assert(转换类.X取整数64位(uint16(0)), int64(0))
		t.Assert(转换类.X取整数64位(uint8(0)), int64(0))
		t.Assert(转换类.X取整数64位(uint(0)), int64(0))
		t.Assert(转换类.X取整数64位(float32(0)), int64(0))

		t.AssertEQ(转换类.X取整数64位(false), int64(0))
		t.AssertEQ(转换类.X取整数64位(nil), int64(0))
		t.AssertEQ(转换类.X取整数64位(0), int64(0))
		t.AssertEQ(转换类.X取整数64位("0"), int64(0))
		t.AssertEQ(转换类.X取整数64位(""), int64(0))
		t.AssertEQ(转换类.X取整数64位("false"), int64(0))
		t.AssertEQ(转换类.X取整数64位("off"), int64(0))
		t.AssertEQ(转换类.X取整数64位([]byte{}), int64(0))
		t.AssertEQ(转换类.X取整数64位([]string{}), int64(0))
		t.AssertEQ(转换类.X取整数64位([2]int{1, 2}), int64(0))
		t.AssertEQ(转换类.X取整数64位([]interface{}{}), int64(0))
		t.AssertEQ(转换类.X取整数64位([]map[int]int{}), int64(0))

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取整数64位(countryCapitalMap), int64(0))

		t.AssertEQ(转换类.X取整数64位("1"), int64(1))
		t.AssertEQ(转换类.X取整数64位("on"), int64(0))
		t.AssertEQ(转换类.X取整数64位(int64(1)), int64(1))
		t.AssertEQ(转换类.X取整数64位(123.456), int64(123))
		t.AssertEQ(转换类.X取整数64位(boolStruct{}), int64(0))
		t.AssertEQ(转换类.X取整数64位(&boolStruct{}), int64(0))
		t.AssertEQ(转换类.X取整数64位("NaN"), int64(0))
	})
}

func Test_Uint_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.AssertEQ(转换类.X取正整数(any), uint(0))
		t.AssertEQ(转换类.X取正整数(false), uint(0))
		t.AssertEQ(转换类.X取正整数(nil), uint(0))
		t.Assert(转换类.X取正整数(nil), uint(0))
		t.AssertEQ(转换类.X取正整数(uint(0)), uint(0))
		t.AssertEQ(转换类.X取正整数("0"), uint(0))
		t.AssertEQ(转换类.X取正整数(""), uint(0))
		t.AssertEQ(转换类.X取正整数("false"), uint(0))
		t.AssertEQ(转换类.X取正整数("off"), uint(0))
		t.AssertEQ(转换类.X取正整数([]byte{}), uint(0))
		t.AssertEQ(转换类.X取正整数([]string{}), uint(0))
		t.AssertEQ(转换类.X取正整数([2]int{1, 2}), uint(0))
		t.AssertEQ(转换类.X取正整数([]interface{}{}), uint(0))
		t.AssertEQ(转换类.X取正整数([]map[int]int{}), uint(0))

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取正整数(countryCapitalMap), uint(0))

		t.AssertEQ(转换类.X取正整数("1"), uint(1))
		t.AssertEQ(转换类.X取正整数("on"), uint(0))
		t.AssertEQ(转换类.X取正整数(1), uint(1))
		t.AssertEQ(转换类.X取正整数(123.456), uint(123))
		t.AssertEQ(转换类.X取正整数(boolStruct{}), uint(0))
		t.AssertEQ(转换类.X取正整数(&boolStruct{}), uint(0))
		t.AssertEQ(转换类.X取正整数("NaN"), uint(0))
	})
}

func Test_Uint8_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.Assert(转换类.X取正整数8位(any), uint8(0))
		t.AssertEQ(转换类.X取正整数8位(uint8(1)), uint8(1))
		t.AssertEQ(转换类.X取正整数8位(false), uint8(0))
		t.AssertEQ(转换类.X取正整数8位(nil), uint8(0))
		t.AssertEQ(转换类.X取正整数8位(0), uint8(0))
		t.AssertEQ(转换类.X取正整数8位("0"), uint8(0))
		t.AssertEQ(转换类.X取正整数8位(""), uint8(0))
		t.AssertEQ(转换类.X取正整数8位("false"), uint8(0))
		t.AssertEQ(转换类.X取正整数8位("off"), uint8(0))
		t.AssertEQ(转换类.X取正整数8位([]byte{}), uint8(0))
		t.AssertEQ(转换类.X取正整数8位([]string{}), uint8(0))
		t.AssertEQ(转换类.X取正整数8位([2]int{1, 2}), uint8(0))
		t.AssertEQ(转换类.X取正整数8位([]interface{}{}), uint8(0))
		t.AssertEQ(转换类.X取正整数8位([]map[int]int{}), uint8(0))

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取正整数8位(countryCapitalMap), uint8(0))

		t.AssertEQ(转换类.X取正整数8位("1"), uint8(1))
		t.AssertEQ(转换类.X取正整数8位("on"), uint8(0))
		t.AssertEQ(转换类.X取正整数8位(int8(1)), uint8(1))
		t.AssertEQ(转换类.X取正整数8位(123.456), uint8(123))
		t.AssertEQ(转换类.X取正整数8位(boolStruct{}), uint8(0))
		t.AssertEQ(转换类.X取正整数8位(&boolStruct{}), uint8(0))
		t.AssertEQ(转换类.X取正整数8位("NaN"), uint8(0))
	})
}

func Test_Uint16_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.Assert(转换类.X取正整数16位(any), uint16(0))
		t.AssertEQ(转换类.X取正整数16位(uint16(1)), uint16(1))
		t.AssertEQ(转换类.X取正整数16位(false), uint16(0))
		t.AssertEQ(转换类.X取正整数16位(nil), uint16(0))
		t.AssertEQ(转换类.X取正整数16位(0), uint16(0))
		t.AssertEQ(转换类.X取正整数16位("0"), uint16(0))
		t.AssertEQ(转换类.X取正整数16位(""), uint16(0))
		t.AssertEQ(转换类.X取正整数16位("false"), uint16(0))
		t.AssertEQ(转换类.X取正整数16位("off"), uint16(0))
		t.AssertEQ(转换类.X取正整数16位([]byte{}), uint16(0))
		t.AssertEQ(转换类.X取正整数16位([]string{}), uint16(0))
		t.AssertEQ(转换类.X取正整数16位([2]int{1, 2}), uint16(0))
		t.AssertEQ(转换类.X取正整数16位([]interface{}{}), uint16(0))
		t.AssertEQ(转换类.X取正整数16位([]map[int]int{}), uint16(0))

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取正整数16位(countryCapitalMap), uint16(0))

		t.AssertEQ(转换类.X取正整数16位("1"), uint16(1))
		t.AssertEQ(转换类.X取正整数16位("on"), uint16(0))
		t.AssertEQ(转换类.X取正整数16位(int16(1)), uint16(1))
		t.AssertEQ(转换类.X取正整数16位(123.456), uint16(123))
		t.AssertEQ(转换类.X取正整数16位(boolStruct{}), uint16(0))
		t.AssertEQ(转换类.X取正整数16位(&boolStruct{}), uint16(0))
		t.AssertEQ(转换类.X取正整数16位("NaN"), uint16(0))
	})
}

func Test_Uint32_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.Assert(转换类.X取正整数32位(any), uint32(0))
		t.AssertEQ(转换类.X取正整数32位(uint32(1)), uint32(1))
		t.AssertEQ(转换类.X取正整数32位(false), uint32(0))
		t.AssertEQ(转换类.X取正整数32位(nil), uint32(0))
		t.AssertEQ(转换类.X取正整数32位(0), uint32(0))
		t.AssertEQ(转换类.X取正整数32位("0"), uint32(0))
		t.AssertEQ(转换类.X取正整数32位(""), uint32(0))
		t.AssertEQ(转换类.X取正整数32位("false"), uint32(0))
		t.AssertEQ(转换类.X取正整数32位("off"), uint32(0))
		t.AssertEQ(转换类.X取正整数32位([]byte{}), uint32(0))
		t.AssertEQ(转换类.X取正整数32位([]string{}), uint32(0))
		t.AssertEQ(转换类.X取正整数32位([2]int{1, 2}), uint32(0))
		t.AssertEQ(转换类.X取正整数32位([]interface{}{}), uint32(0))
		t.AssertEQ(转换类.X取正整数32位([]map[int]int{}), uint32(0))

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取正整数32位(countryCapitalMap), uint32(0))

		t.AssertEQ(转换类.X取正整数32位("1"), uint32(1))
		t.AssertEQ(转换类.X取正整数32位("on"), uint32(0))
		t.AssertEQ(转换类.X取正整数32位(int32(1)), uint32(1))
		t.AssertEQ(转换类.X取正整数32位(123.456), uint32(123))
		t.AssertEQ(转换类.X取正整数32位(boolStruct{}), uint32(0))
		t.AssertEQ(转换类.X取正整数32位(&boolStruct{}), uint32(0))
		t.AssertEQ(转换类.X取正整数32位("NaN"), uint32(0))
	})
}

func Test_Uint64_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.AssertEQ(转换类.X取正整数64位("0x00e"), uint64(14))
		t.Assert(转换类.X取正整数64位("022"), uint64(22))

		t.AssertEQ(转换类.X取正整数64位(any), uint64(0))
		t.AssertEQ(转换类.X取正整数64位(true), uint64(1))
		t.Assert(转换类.X取正整数64位("1"), int64(1))
		t.Assert(转换类.X取正整数64位("0"), uint64(0))
		t.Assert(转换类.X取正整数64位("X"), uint64(0))
		t.Assert(转换类.X取正整数64位("x"), uint64(0))
		t.Assert(转换类.X取正整数64位(int64(1)), uint64(1))
		t.Assert(转换类.X取正整数64位(int(0)), uint64(0))
		t.Assert(转换类.X取正整数64位(int8(0)), uint64(0))
		t.Assert(转换类.X取正整数64位(int16(0)), uint64(0))
		t.Assert(转换类.X取正整数64位(int32(0)), uint64(0))
		t.Assert(转换类.X取正整数64位(uint64(0)), uint64(0))
		t.Assert(转换类.X取正整数64位(uint32(0)), uint64(0))
		t.Assert(转换类.X取正整数64位(uint16(0)), uint64(0))
		t.Assert(转换类.X取正整数64位(uint8(0)), uint64(0))
		t.Assert(转换类.X取正整数64位(uint(0)), uint64(0))
		t.Assert(转换类.X取正整数64位(float32(0)), uint64(0))

		t.AssertEQ(转换类.X取正整数64位(false), uint64(0))
		t.AssertEQ(转换类.X取正整数64位(nil), uint64(0))
		t.AssertEQ(转换类.X取正整数64位(0), uint64(0))
		t.AssertEQ(转换类.X取正整数64位("0"), uint64(0))
		t.AssertEQ(转换类.X取正整数64位(""), uint64(0))
		t.AssertEQ(转换类.X取正整数64位("false"), uint64(0))
		t.AssertEQ(转换类.X取正整数64位("off"), uint64(0))
		t.AssertEQ(转换类.X取正整数64位([]byte{}), uint64(0))
		t.AssertEQ(转换类.X取正整数64位([]string{}), uint64(0))
		t.AssertEQ(转换类.X取正整数64位([2]int{1, 2}), uint64(0))
		t.AssertEQ(转换类.X取正整数64位([]interface{}{}), uint64(0))
		t.AssertEQ(转换类.X取正整数64位([]map[int]int{}), uint64(0))

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取正整数64位(countryCapitalMap), uint64(0))

		t.AssertEQ(转换类.X取正整数64位("1"), uint64(1))
		t.AssertEQ(转换类.X取正整数64位("on"), uint64(0))
		t.AssertEQ(转换类.X取正整数64位(int64(1)), uint64(1))
		t.AssertEQ(转换类.X取正整数64位(123.456), uint64(123))
		t.AssertEQ(转换类.X取正整数64位(boolStruct{}), uint64(0))
		t.AssertEQ(转换类.X取正整数64位(&boolStruct{}), uint64(0))
		t.AssertEQ(转换类.X取正整数64位("NaN"), uint64(0))
	})
}

func Test_Float32_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.Assert(转换类.X取小数32位(any), float32(0))
		t.AssertEQ(转换类.X取小数32位(false), float32(0))
		t.AssertEQ(转换类.X取小数32位(nil), float32(0))
		t.AssertEQ(转换类.X取小数32位(0), float32(0))
		t.AssertEQ(转换类.X取小数32位("0"), float32(0))
		t.AssertEQ(转换类.X取小数32位(""), float32(0))
		t.AssertEQ(转换类.X取小数32位("false"), float32(0))
		t.AssertEQ(转换类.X取小数32位("off"), float32(0))
		t.AssertEQ(转换类.X取小数32位([]byte{}), float32(0))
		t.AssertEQ(转换类.X取小数32位([]string{}), float32(0))
		t.AssertEQ(转换类.X取小数32位([2]int{1, 2}), float32(0))
		t.AssertEQ(转换类.X取小数32位([]interface{}{}), float32(0))
		t.AssertEQ(转换类.X取小数32位([]map[int]int{}), float32(0))
		t.AssertEQ(转换类.X取小数32位(泛型类.X创建(float32(0))), float32(0))

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取小数32位(countryCapitalMap), float32(0))

		t.AssertEQ(转换类.X取小数32位("1"), float32(1))
		t.AssertEQ(转换类.X取小数32位("on"), float32(0))
		t.AssertEQ(转换类.X取小数32位(float32(1)), float32(1))
		t.AssertEQ(转换类.X取小数32位(123.456), float32(123.456))
		t.AssertEQ(转换类.X取小数32位(boolStruct{}), float32(0))
		t.AssertEQ(转换类.X取小数32位(&boolStruct{}), float32(0))
		t.AssertEQ(转换类.X取小数32位("NaN"), float32(math.NaN()))
	})
}

func Test_Float64_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.Assert(转换类.X取小数64位(any), float64(0))
		t.AssertEQ(转换类.X取小数64位(false), float64(0))
		t.AssertEQ(转换类.X取小数64位(nil), float64(0))
		t.AssertEQ(转换类.X取小数64位(0), float64(0))
		t.AssertEQ(转换类.X取小数64位("0"), float64(0))
		t.AssertEQ(转换类.X取小数64位(""), float64(0))
		t.AssertEQ(转换类.X取小数64位("false"), float64(0))
		t.AssertEQ(转换类.X取小数64位("off"), float64(0))
		t.AssertEQ(转换类.X取小数64位([]byte{}), float64(0))
		t.AssertEQ(转换类.X取小数64位([]string{}), float64(0))
		t.AssertEQ(转换类.X取小数64位([2]int{1, 2}), float64(0))
		t.AssertEQ(转换类.X取小数64位([]interface{}{}), float64(0))
		t.AssertEQ(转换类.X取小数64位([]map[int]int{}), float64(0))
		t.AssertEQ(转换类.X取小数64位(泛型类.X创建(float64(0))), float64(0))

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.X取小数64位(countryCapitalMap), float64(0))

		t.AssertEQ(转换类.X取小数64位("1"), float64(1))
		t.AssertEQ(转换类.X取小数64位("on"), float64(0))
		t.AssertEQ(转换类.X取小数64位(float64(1)), float64(1))
		t.AssertEQ(转换类.X取小数64位(123.456), float64(123.456))
		t.AssertEQ(转换类.X取小数64位(boolStruct{}), float64(0))
		t.AssertEQ(转换类.X取小数64位(&boolStruct{}), float64(0))
		t.AssertEQ(转换类.X取小数64位("NaN"), float64(math.NaN()))
	})
}

func Test_String_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var s []rune
		t.AssertEQ(转换类.String(s), "")
		var any interface{} = nil
		t.AssertEQ(转换类.String(any), "")
		t.AssertEQ(转换类.String("1"), "1")
		t.AssertEQ(转换类.String("0"), string("0"))
		t.Assert(转换类.String("X"), string("X"))
		t.Assert(转换类.String("x"), string("x"))
		t.Assert(转换类.String(int64(1)), uint64(1))
		t.Assert(转换类.String(int(0)), string("0"))
		t.Assert(转换类.String(int8(0)), string("0"))
		t.Assert(转换类.String(int16(0)), string("0"))
		t.Assert(转换类.String(int32(0)), string("0"))
		t.Assert(转换类.String(uint64(0)), string("0"))
		t.Assert(转换类.String(uint32(0)), string("0"))
		t.Assert(转换类.String(uint16(0)), string("0"))
		t.Assert(转换类.String(uint8(0)), string("0"))
		t.Assert(转换类.String(uint(0)), string("0"))
		t.Assert(转换类.String(float32(0)), string("0"))
		t.AssertEQ(转换类.String(true), "true")
		t.AssertEQ(转换类.String(false), "false")
		t.AssertEQ(转换类.String(nil), "")
		t.AssertEQ(转换类.String(0), string("0"))
		t.AssertEQ(转换类.String("0"), string("0"))
		t.AssertEQ(转换类.String(""), "")
		t.AssertEQ(转换类.String("false"), "false")
		t.AssertEQ(转换类.String("off"), string("off"))
		t.AssertEQ(转换类.String([]byte{}), "")
		t.AssertEQ(转换类.String([]string{}), "[]")
		t.AssertEQ(转换类.String([2]int{1, 2}), "[1,2]")
		t.AssertEQ(转换类.String([]interface{}{}), "[]")
		t.AssertEQ(转换类.String(map[int]int{}), "{}")

		var countryCapitalMap = make(map[string]string)
		/* map插入key - value对,各个国家对应的首都 */
		countryCapitalMap["France"] = "巴黎"
		countryCapitalMap["Italy"] = "罗马"
		countryCapitalMap["Japan"] = "东京"
		countryCapitalMap["India "] = "新德里"
		t.AssertEQ(转换类.String(countryCapitalMap), `{"France":"巴黎","India ":"新德里","Italy":"罗马","Japan":"东京"}`)
		t.AssertEQ(转换类.String(int64(1)), "1")
		t.AssertEQ(转换类.String(123.456), "123.456")
		t.AssertEQ(转换类.String(boolStruct{}), "{}")
		t.AssertEQ(转换类.String(&boolStruct{}), "{}")

		var info = new(S)
		t.AssertEQ(转换类.String(info), "22222")
		var errInfo = new(S1)
		t.AssertEQ(转换类.String(errInfo), "22222")
	})
}

func Test_Runes_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取字符数组("www"), []int32{119, 119, 119})
		var s []rune
		t.AssertEQ(转换类.X取字符数组(s), nil)
	})
}

func Test_Rune_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取字符("www"), int32(0))
		t.AssertEQ(转换类.X取字符(int32(0)), int32(0))
		var s []rune
		t.AssertEQ(转换类.X取字符(s), int32(0))
	})
}

func Test_Bytes_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取字节集(nil), nil)
		t.AssertEQ(转换类.X取字节集(int32(0)), []uint8{0, 0, 0, 0})
		t.AssertEQ(转换类.X取字节集("s"), []uint8{115})
		t.AssertEQ(转换类.X取字节集([]byte("s")), []uint8{115})
		t.AssertEQ(转换类.X取字节集(泛型类.X创建([]byte("s"))), []uint8{115})
	})
}

func Test_Byte_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取字节(uint8(0)), uint8(0))
		t.AssertEQ(转换类.X取字节("s"), uint8(0))
		t.AssertEQ(转换类.X取字节([]byte("s")), uint8(115))
	})
}

func Test_Convert_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.AssertEQ(转换类.X按名称转换(any, "string"), "")
		t.AssertEQ(转换类.X按名称转换("1", "string"), "1")
		t.Assert(转换类.X按名称转换(int64(1), "int64"), int64(1))
		t.Assert(转换类.X按名称转换(int(0), "int"), int(0))
		t.Assert(转换类.X按名称转换(int8(0), "int8"), int8(0))
		t.Assert(转换类.X按名称转换(int16(0), "int16"), int16(0))
		t.Assert(转换类.X按名称转换(int32(0), "int32"), int32(0))
		t.Assert(转换类.X按名称转换(uint64(0), "uint64"), uint64(0))
		t.Assert(转换类.X按名称转换(uint32(0), "uint32"), uint32(0))
		t.Assert(转换类.X按名称转换(uint16(0), "uint16"), uint16(0))
		t.Assert(转换类.X按名称转换(uint8(0), "uint8"), uint8(0))
		t.Assert(转换类.X按名称转换(uint(0), "uint"), uint(0))
		t.Assert(转换类.X按名称转换(float32(0), "float32"), float32(0))
		t.Assert(转换类.X按名称转换(float64(0), "float64"), float64(0))
		t.AssertEQ(转换类.X按名称转换(true, "bool"), true)
		t.AssertEQ(转换类.X按名称转换([]byte{}, "[]byte"), []uint8{})
		t.AssertEQ(转换类.X按名称转换([]string{}, "[]string"), []string{})
		t.AssertEQ(转换类.X按名称转换([2]int{1, 2}, "[]int"), []int{1, 2})
		t.AssertEQ(转换类.X按名称转换([2]uint8{1, 2}, "[]uint8"), []uint8{1, 2})
		t.AssertEQ(转换类.X按名称转换("1989-01-02", "Time", "Y-m-d"), 转换类.X取时间("1989-01-02", "Y-m-d"))
		t.AssertEQ(转换类.X按名称转换(1989, "Time"), 转换类.X取时间("1970-01-01 08:33:09 +0800 CST"))
		t.AssertEQ(转换类.X按名称转换(时间类.X创建并按当前时间(), "gtime.Time", 1), *时间类.X创建())
		t.AssertEQ(转换类.X按名称转换(1989, "gtime.Time"), *转换类.X取gtime时间类("1970-01-01 08:33:09 +0800 CST"))
		t.AssertEQ(转换类.X按名称转换(时间类.X创建并按当前时间(), "*gtime.Time", 1), 时间类.X创建())
		t.AssertEQ(转换类.X按名称转换(时间类.X创建并按当前时间(), "GTime", 1), *时间类.X创建())
		t.AssertEQ(转换类.X按名称转换(1989, "*gtime.Time"), 转换类.X取gtime时间类(1989))
		t.AssertEQ(转换类.X按名称转换(1989, "Duration"), time.Duration(int64(1989)))
		t.AssertEQ(转换类.X按名称转换("1989", "Duration"), time.Duration(int64(1989)))
		t.AssertEQ(转换类.X按名称转换("1989", ""), "1989")

		var intNum int = 1
		t.Assert(转换类.X按名称转换(&intNum, "*int"), int(1))
		var int8Num int8 = 1
		t.Assert(转换类.X按名称转换(int8Num, "*int8"), int(1))
		t.Assert(转换类.X按名称转换(&int8Num, "*int8"), int(1))
		var int16Num int16 = 1
		t.Assert(转换类.X按名称转换(int16Num, "*int16"), int(1))
		t.Assert(转换类.X按名称转换(&int16Num, "*int16"), int(1))
		var int32Num int32 = 1
		t.Assert(转换类.X按名称转换(int32Num, "*int32"), int(1))
		t.Assert(转换类.X按名称转换(&int32Num, "*int32"), int(1))
		var int64Num int64 = 1
		t.Assert(转换类.X按名称转换(int64Num, "*int64"), int(1))
		t.Assert(转换类.X按名称转换(&int64Num, "*int64"), int(1))

		var uintNum uint = 1
		t.Assert(转换类.X按名称转换(&uintNum, "*uint"), int(1))
		var uint8Num uint8 = 1
		t.Assert(转换类.X按名称转换(uint8Num, "*uint8"), int(1))
		t.Assert(转换类.X按名称转换(&uint8Num, "*uint8"), int(1))
		var uint16Num uint16 = 1
		t.Assert(转换类.X按名称转换(uint16Num, "*uint16"), int(1))
		t.Assert(转换类.X按名称转换(&uint16Num, "*uint16"), int(1))
		var uint32Num uint32 = 1
		t.Assert(转换类.X按名称转换(uint32Num, "*uint32"), int(1))
		t.Assert(转换类.X按名称转换(&uint32Num, "*uint32"), int(1))
		var uint64Num uint64 = 1
		t.Assert(转换类.X按名称转换(uint64Num, "*uint64"), int(1))
		t.Assert(转换类.X按名称转换(&uint64Num, "*uint64"), int(1))

		var float32Num float32 = 1.1
		t.Assert(转换类.X按名称转换(float32Num, "*float32"), float32(1.1))
		t.Assert(转换类.X按名称转换(&float32Num, "*float32"), float32(1.1))

		var float64Num float64 = 1.1
		t.Assert(转换类.X按名称转换(float64Num, "*float64"), float64(1.1))
		t.Assert(转换类.X按名称转换(&float64Num, "*float64"), float64(1.1))

		var boolValue bool = true
		t.Assert(转换类.X按名称转换(boolValue, "*bool"), true)
		t.Assert(转换类.X按名称转换(&boolValue, "*bool"), true)

		var stringValue string = "1"
		t.Assert(转换类.X按名称转换(stringValue, "*string"), "1")
		t.Assert(转换类.X按名称转换(&stringValue, "*string"), "1")

		var durationValue time.Duration = 1989
		var expectDurationValue = time.Duration(int64(1989))
		t.AssertEQ(转换类.X按名称转换(&durationValue, "*time.Duration"), &expectDurationValue)
		t.AssertEQ(转换类.X按名称转换(durationValue, "*time.Duration"), &expectDurationValue)

		var string_interface_map = map[string]interface{}{"k1": 1}
		var string_int_map = map[string]int{"k1": 1}
		var string_string_map = map[string]string{"k1": "1"}
		t.AssertEQ(转换类.X按名称转换(string_int_map, "map[string]string"), string_string_map)
		t.AssertEQ(转换类.X按名称转换(string_int_map, "map[string]interface{}"), string_interface_map)
	})
}

func Test_Slice_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		value := 123.456
		t.AssertEQ(转换类.X取整数数组(value), []int{123})
		t.AssertEQ(转换类.X取整数数组(nil), nil)
		t.AssertEQ(转换类.X取整数数组([]string{"1", "2"}), []int{1, 2})
		t.AssertEQ(转换类.X取整数数组([]int{}), []int{})
		t.AssertEQ(转换类.X取整数数组([]int8{1, 2}), []int{1, 2})
		t.AssertEQ(转换类.X取整数数组([]int16{1, 2}), []int{1, 2})
		t.AssertEQ(转换类.X取整数数组([]int32{1, 2}), []int{1, 2})
		t.AssertEQ(转换类.X取整数数组([]int64{1, 2}), []int{1, 2})
		t.AssertEQ(转换类.X取整数数组([]uint{1}), []int{1})
		t.AssertEQ(转换类.X取整数数组([]uint8{1, 2}), []int{1, 2})
		t.AssertEQ(转换类.X取整数数组([]uint16{1, 2}), []int{1, 2})
		t.AssertEQ(转换类.X取整数数组([]uint32{1, 2}), []int{1, 2})
		t.AssertEQ(转换类.X取整数数组([]uint64{1, 2}), []int{1, 2})
		t.AssertEQ(转换类.X取整数数组([]bool{true}), []int{1})
		t.AssertEQ(转换类.X取整数数组([]float32{1, 2}), []int{1, 2})
		t.AssertEQ(转换类.X取整数数组([]float64{1, 2}), []int{1, 2})
		var inter []interface{} = make([]interface{}, 2)
		t.AssertEQ(转换类.X取整数数组(inter), []int{0, 0})

		t.AssertEQ(转换类.X取文本数组(value), []string{"123.456"})
		t.AssertEQ(转换类.X取文本数组(nil), nil)
		t.AssertEQ(转换类.X取文本数组([]string{"1", "2"}), []string{"1", "2"})
		t.AssertEQ(转换类.X取文本数组([]int{1}), []string{"1"})
		t.AssertEQ(转换类.X取文本数组([]int8{1, 2}), []string{"1", "2"})
		t.AssertEQ(转换类.X取文本数组([]int16{1, 2}), []string{"1", "2"})
		t.AssertEQ(转换类.X取文本数组([]int32{1, 2}), []string{"1", "2"})
		t.AssertEQ(转换类.X取文本数组([]int64{1, 2}), []string{"1", "2"})
		t.AssertEQ(转换类.X取文本数组([]uint{1}), []string{"1"})
		t.AssertEQ(转换类.X取文本数组([]uint8{1, 2}), []string{"1", "2"})
		t.AssertEQ(转换类.X取文本数组([]uint16{1, 2}), []string{"1", "2"})
		t.AssertEQ(转换类.X取文本数组([]uint32{1, 2}), []string{"1", "2"})
		t.AssertEQ(转换类.X取文本数组([]uint64{1, 2}), []string{"1", "2"})
		t.AssertEQ(转换类.X取文本数组([]bool{true}), []string{"true"})
		t.AssertEQ(转换类.X取文本数组([]float32{1, 2}), []string{"1", "2"})
		t.AssertEQ(转换类.X取文本数组([]float64{1, 2}), []string{"1", "2"})
		var strer = make([]interface{}, 2)
		t.AssertEQ(转换类.X取文本数组(strer), []string{"", ""})

		t.AssertEQ(转换类.X取小数数组(value), []float64{123.456})
		t.AssertEQ(转换类.X取小数数组(nil), nil)
		t.AssertEQ(转换类.X取小数数组([]string{"1", "2"}), []float64{1, 2})
		t.AssertEQ(转换类.X取小数数组([]int{1}), []float64{1})
		t.AssertEQ(转换类.X取小数数组([]int8{1, 2}), []float64{1, 2})
		t.AssertEQ(转换类.X取小数数组([]int16{1, 2}), []float64{1, 2})
		t.AssertEQ(转换类.X取小数数组([]int32{1, 2}), []float64{1, 2})
		t.AssertEQ(转换类.X取小数数组([]int64{1, 2}), []float64{1, 2})
		t.AssertEQ(转换类.X取小数数组([]uint{1}), []float64{1})
		t.AssertEQ(转换类.X取小数数组([]uint8{1, 2}), []float64{1, 2})
		t.AssertEQ(转换类.X取小数数组([]uint16{1, 2}), []float64{1, 2})
		t.AssertEQ(转换类.X取小数数组([]uint32{1, 2}), []float64{1, 2})
		t.AssertEQ(转换类.X取小数数组([]uint64{1, 2}), []float64{1, 2})
		t.AssertEQ(转换类.X取小数数组([]bool{true}), []float64{0})
		t.AssertEQ(转换类.X取小数数组([]float32{1, 2}), []float64{1, 2})
		t.AssertEQ(转换类.X取小数数组([]float64{1, 2}), []float64{1, 2})
		var floer = make([]interface{}, 2)
		t.AssertEQ(转换类.X取小数数组(floer), []float64{0, 0})

		t.AssertEQ(转换类.X取any数组(value), []interface{}{123.456})
		t.AssertEQ(转换类.X取any数组(nil), nil)
		t.AssertEQ(转换类.X取any数组([]interface{}{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]string{"1"}), []interface{}{"1"})
		t.AssertEQ(转换类.X取any数组([]int{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]int8{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]int16{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]int32{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]int64{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]uint{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]uint8{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]uint16{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]uint32{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]uint64{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]bool{true}), []interface{}{true})
		t.AssertEQ(转换类.X取any数组([]float32{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([]float64{1}), []interface{}{1})
		t.AssertEQ(转换类.X取any数组([1]int{1}), []interface{}{1})

		type interSlice []int
		slices := interSlice{1}
		t.AssertEQ(转换类.X取any数组(slices), []interface{}{1})

		t.AssertEQ(转换类.X取Map数组(nil), nil)
		t.AssertEQ(转换类.X取Map数组([]map[string]interface{}{{"a": "1"}}), []map[string]interface{}{{"a": "1"}})
		t.AssertEQ(转换类.X取Map数组(1223), []map[string]interface{}{nil})
		t.AssertEQ(转换类.X取Map数组([]int{}), nil)
	})
}

// 私有属性不会进行转换
func Test_Slice_PrivateAttribute_All(t *testing.T) {
	type User struct {
		Id   int           `json:"id"`
		name string        `json:"name"`
		Ad   []interface{} `json:"ad"`
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := &User{1, "john", []interface{}{2}}
		array := 转换类.X取any数组(user)
		t.Assert(len(array), 1)
		t.Assert(array[0].(*User).Id, 1)
		t.Assert(array[0].(*User).name, "john")
		t.Assert(array[0].(*User).Ad, []interface{}{2})
	})
}

func Test_Map_Basic_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := map[string]string{
			"k": "v",
		}
		m2 := map[int]string{
			3: "v",
		}
		m3 := map[float64]float32{
			1.22: 3.1,
		}
		t.Assert(转换类.X取Map(m1), g.Map{
			"k": "v",
		})
		t.Assert(转换类.X取Map(m2), g.Map{
			"3": "v",
		})
		t.Assert(转换类.X取Map(m3), g.Map{
			"1.22": "3.1",
		})
		t.AssertEQ(转换类.X取Map(nil), nil)
		t.AssertEQ(转换类.X取Map(map[string]interface{}{"a": 1}), map[string]interface{}{"a": 1})
		t.AssertEQ(转换类.X取Map(map[int]interface{}{1: 1}), map[string]interface{}{"1": 1})
		t.AssertEQ(转换类.X取Map(map[uint]interface{}{1: 1}), map[string]interface{}{"1": 1})
		t.AssertEQ(转换类.X取Map(map[uint]string{1: "1"}), map[string]interface{}{"1": "1"})

		t.AssertEQ(转换类.X取Map(map[interface{}]interface{}{"a": 1}), map[interface{}]interface{}{"a": 1})
		t.AssertEQ(转换类.X取Map(map[interface{}]string{"a": "1"}), map[interface{}]string{"a": "1"})
		t.AssertEQ(转换类.X取Map(map[interface{}]int{"a": 1}), map[interface{}]int{"a": 1})
		t.AssertEQ(转换类.X取Map(map[interface{}]uint{"a": 1}), map[interface{}]uint{"a": 1})
		t.AssertEQ(转换类.X取Map(map[interface{}]float32{"a": 1}), map[interface{}]float32{"a": 1})
		t.AssertEQ(转换类.X取Map(map[interface{}]float64{"a": 1}), map[interface{}]float64{"a": 1})

		t.AssertEQ(转换类.X取Map(map[string]bool{"a": true}), map[string]interface{}{"a": true})
		t.AssertEQ(转换类.X取Map(map[string]int{"a": 1}), map[string]interface{}{"a": 1})
		t.AssertEQ(转换类.X取Map(map[string]uint{"a": 1}), map[string]interface{}{"a": 1})
		t.AssertEQ(转换类.X取Map(map[string]float32{"a": 1}), map[string]interface{}{"a": 1})
		t.AssertEQ(转换类.X取Map(map[string]float64{"a": 1}), map[string]interface{}{"a": 1})

	})
}

func Test_Map_StructWithGconvTag_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid      int
			Name     string
			SiteUrl  string   `gconv:"-"`
			NickName string   `gconv:"nickname,omitempty"`
			Pass1    string   `gconv:"password1"`
			Pass2    string   `gconv:"password2"`
			Ss       []string `gconv:"ss"`
		}
		user1 := User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
			Ss:      []string{"sss", "2222"},
		}
		user2 := &user1
		map1 := 转换类.X取Map(user1)
		map2 := 转换类.X取Map(user2)
		t.Assert(map1["Uid"], 100)
		t.Assert(map1["Name"], "john")
		t.Assert(map1["SiteUrl"], nil)
		t.Assert(map1["NickName"], nil)
		t.Assert(map1["nickname"], nil)
		t.Assert(map1["password1"], "123")
		t.Assert(map1["password2"], "456")
		t.Assert(map2["Uid"], 100)
		t.Assert(map2["Name"], "john")
		t.Assert(map2["SiteUrl"], nil)
		t.Assert(map2["NickName"], nil)
		t.Assert(map2["nickname"], nil)
		t.Assert(map2["password1"], "123")
		t.Assert(map2["password2"], "456")
	})
}

func Test_Map_StructWithJsonTag_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid      int
			Name     string
			SiteUrl  string   `json:"-"`
			NickName string   `json:"nickname, omitempty"`
			Pass1    string   `json:"password1,newpassword"`
			Pass2    string   `json:"password2"`
			Ss       []string `json:"omitempty"`
			ssb, ssa string
		}
		user1 := User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
			Ss:      []string{"sss", "2222"},
			ssb:     "11",
			ssa:     "222",
		}
		user3 := User{
			Uid:      100,
			Name:     "john",
			NickName: "SSS",
			SiteUrl:  "https://goframe.org",
			Pass1:    "123",
			Pass2:    "456",
			Ss:       []string{"sss", "2222"},
			ssb:      "11",
			ssa:      "222",
		}
		user2 := &user1
		_ = 转换类.X取Map(user1, 转换类.MapOption{Tags: []string{"Ss"}})
		map1 := 转换类.X取Map(user1, 转换类.MapOption{Tags: []string{"json", "json2"}})
		map2 := 转换类.X取Map(user2)
		map3 := 转换类.X取Map(user3)
		t.Assert(map1["Uid"], 100)
		t.Assert(map1["Name"], "john")
		t.Assert(map1["SiteUrl"], nil)
		t.Assert(map1["NickName"], nil)
		t.Assert(map1["nickname"], nil)
		t.Assert(map1["password1"], "123")
		t.Assert(map1["password2"], "456")
		t.Assert(map2["Uid"], 100)
		t.Assert(map2["Name"], "john")
		t.Assert(map2["SiteUrl"], nil)
		t.Assert(map2["NickName"], nil)
		t.Assert(map2["nickname"], nil)
		t.Assert(map2["password1"], "123")
		t.Assert(map2["password2"], "456")
		t.Assert(map3["NickName"], nil)
	})
}

func Test_Map_PrivateAttribute_All(t *testing.T) {
	type User struct {
		Id   int
		name string
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := &User{1, "john"}
		t.Assert(转换类.X取Map(user), g.Map{"Id": 1})
	})
}

func Test_Map_StructInherit_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Ids struct {
			Id  int `json:"id"`
			Uid int `json:"uid"`
		}
		type Base struct {
			Ids
			CreateTime string `json:"create_time"`
		}
		type User struct {
			Base
			Passport string  `json:"passport"`
			Password string  `json:"password"`
			Nickname string  `json:"nickname"`
			S        *string `json:"nickname2"`
		}

		user := new(User)
		user.Id = 100
		user.Nickname = "john"
		user.CreateTime = "2019"
		var s = "s"
		user.S = &s

		m := 转换类.X取Map_递归(user)
		t.Assert(m["id"], user.Id)
		t.Assert(m["nickname"], user.Nickname)
		t.Assert(m["create_time"], user.CreateTime)
		t.Assert(m["nickname2"], user.S)
	})
}

func Test_Struct_Basic1_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   int
			Result string
		}

		type Score2 struct {
			Name   int
			Result string
		}

		type User struct {
			Uid      int
			Name     string
			Site_Url string
			NickName string
			Pass1    string `gconv:"password1"`
			Pass2    string `gconv:"password2"`
			As       *Score
			Ass      Score
			Assb     []interface{}
		}
		// 使用默认映射规则绑定属性值到对象
		user := new(User)
		params1 := g.Map{
			"uid":       1,
			"Name":      "john",
			"siteurl":   "https://goframe.org",
			"nick_name": "johng",
			"PASS1":     "123",
			"PASS2":     "456",
			"As":        g.Map{"Name": 1, "Result": "22222"},
			"Ass":       &Score{11, "11"},
			"Assb":      []string{"wwww"},
		}
		_ = 转换类.Struct(nil, user)
		_ = 转换类.Struct(params1, nil)
		_ = 转换类.Struct([]interface{}{nil}, user)
		_ = 转换类.Struct(user, []interface{}{nil})

		var a = []interface{}{nil}
		ab := &a
		_ = 转换类.Struct(params1, *ab)
		var pi *int = nil
		_ = 转换类.Struct(params1, pi)

		_ = 转换类.Struct(params1, user)
		_ = 转换类.Struct(params1, user, map[string]string{"uid": "Names"})
		_ = 转换类.Struct(params1, user, map[string]string{"uid": "as"})

		// 使用struct tag映射绑定属性值到对象
		user = new(User)
		params2 := g.Map{
			"uid":       2,
			"name":      "smith",
			"site-url":  "https://goframe.org",
			"nick name": "johng",
			"password1": "111",
			"password2": "222",
		}
		if err := 转换类.Struct(params2, user); err != nil {
			单元测试类.Error(err)
		}
		t.Assert(user, &User{
			Uid:      2,
			Name:     "smith",
			Site_Url: "https://goframe.org",
			NickName: "johng",
			Pass1:    "111",
			Pass2:    "222",
		})
	})
}

// 使用默认映射规则绑定属性值到对象
func Test_Struct_Basic2_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid     int
			Name    string
			SiteUrl string
			Pass1   string
			Pass2   string
		}
		user := new(User)
		params := g.Map{
			"uid":      1,
			"Name":     "john",
			"site_url": "https://goframe.org",
			"PASS1":    "123",
			"PASS2":    "456",
		}
		if err := 转换类.Struct(params, user); err != nil {
			单元测试类.Error(err)
		}
		t.Assert(user, &User{
			Uid:     1,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		})
	})
}

// 带有指针的基础类型属性
func Test_Struct_Basic3_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid  int
			Name *string
		}
		user := new(User)
		params := g.Map{
			"uid":  1,
			"Name": "john",
		}
		if err := 转换类.Struct(params, user); err != nil {
			单元测试类.Error(err)
		}
		t.Assert(user.Uid, 1)
		t.Assert(*user.Name, "john")
	})
}

// slice类型属性的赋值
func Test_Struct_Attr_Slice_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Scores []int
		}
		scores := []interface{}{99, 100, 60, 140}
		user := new(User)
		if err := 转换类.Struct(g.Map{"Scores": scores}, user); err != nil {
			单元测试类.Error(err)
		} else {
			t.Assert(user, &User{
				Scores: []int{99, 100, 60, 140},
			})
		}
	})
}

// 属性为struct对象
func Test_Struct_Attr_Struct_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   string
			Result int
		}
		type User struct {
			Scores Score
		}

		user := new(User)
		scores := map[string]interface{}{
			"Scores": map[string]interface{}{
				"Name":   "john",
				"Result": 100,
			},
		}

		// 嵌套struct转换
		if err := 转换类.Struct(scores, user); err != nil {
			单元测试类.Error(err)
		} else {
			t.Assert(user, &User{
				Scores: Score{
					Name:   "john",
					Result: 100,
				},
			})
		}
	})
}

// 属性为struct对象指针
func Test_Struct_Attr_Struct_Ptr_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   string
			Result int
		}
		type User struct {
			Scores *Score
		}

		user := new(User)
		scores := map[string]interface{}{
			"Scores": map[string]interface{}{
				"Name":   "john",
				"Result": 100,
			},
		}

		// 嵌套struct转换
		if err := 转换类.Struct(scores, user); err != nil {
			单元测试类.Error(err)
		} else {
			t.Assert(user.Scores, &Score{
				Name:   "john",
				Result: 100,
			})
		}
	})
}

// 属性为struct对象slice
func Test_Struct_Attr_Struct_Slice1_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   string
			Result int
		}
		type User struct {
			Scores []Score
		}

		user := new(User)
		scores := map[string]interface{}{
			"Scores": map[string]interface{}{
				"Name":   "john",
				"Result": 100,
			},
		}

		// 嵌套struct转换，属性为slice类型，数值为map类型
		if err := 转换类.Struct(scores, user); err != nil {
			单元测试类.Error(err)
		} else {
			t.Assert(user.Scores, []Score{
				{
					Name:   "john",
					Result: 100,
				},
			})
		}
	})
}

// 属性为struct对象slice
func Test_Struct_Attr_Struct_Slice2_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   string
			Result int
		}
		type User struct {
			Scores []Score
		}

		user := new(User)
		scores := map[string]interface{}{
			"Scores": []interface{}{
				map[string]interface{}{
					"Name":   "john",
					"Result": 100,
				},
				map[string]interface{}{
					"Name":   "smith",
					"Result": 60,
				},
			},
		}

		// 嵌套struct转换，属性为slice类型，数值为slice map类型
		if err := 转换类.Struct(scores, user); err != nil {
			单元测试类.Error(err)
		} else {
			t.Assert(user.Scores, []Score{
				{
					Name:   "john",
					Result: 100,
				},
				{
					Name:   "smith",
					Result: 60,
				},
			})
		}
	})
}

// 属性为struct对象slice ptr
func Test_Struct_Attr_Struct_Slice_Ptr_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Score struct {
			Name   string
			Result int
		}
		type User struct {
			Scores []*Score
		}

		user := new(User)
		scores := map[string]interface{}{
			"Scores": []interface{}{
				map[string]interface{}{
					"Name":   "john",
					"Result": 100,
				},
				map[string]interface{}{
					"Name":   "smith",
					"Result": 60,
				},
			},
		}

		// 嵌套struct转换，属性为slice类型，数值为slice map类型
		if err := 转换类.Struct(scores, user); err != nil {
			单元测试类.Error(err)
		} else {
			t.Assert(len(user.Scores), 2)
			t.Assert(user.Scores[0], &Score{
				Name:   "john",
				Result: 100,
			})
			t.Assert(user.Scores[1], &Score{
				Name:   "smith",
				Result: 60,
			})
		}
	})
}

func Test_Struct_PrivateAttribute_All(t *testing.T) {
	type User struct {
		Id   int
		name string
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		err := 转换类.Struct(g.Map{"id": 1, "name": "john"}, user)
		t.AssertNil(err)
		t.Assert(user.Id, 1)
		t.Assert(user.name, "")
	})
}

func Test_Struct_Embedded_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type Ids struct {
			Id  int `json:"id"`
			Uid int `json:"uid"`
		}
		type Base struct {
			Ids
			CreateTime string `json:"create_time"`
		}
		type User struct {
			Base
			Passport string `json:"passport"`
			Password string `json:"password"`
			Nickname string `json:"nickname"`
		}
		data := g.Map{
			"id":          100,
			"uid":         101,
			"passport":    "t1",
			"password":    "123456",
			"nickname":    "T1",
			"create_time": "2019",
		}
		user := new(User)
		转换类.Struct(data, user)
		t.Assert(user.Id, 100)
		t.Assert(user.Uid, 101)
		t.Assert(user.Nickname, "T1")
		t.Assert(user.CreateTime, "2019")
	})
}

func Test_Struct_Time_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			CreateTime time.Time
		}
		now := time.Now()
		user := new(User)
		转换类.Struct(g.Map{
			"create_time": now,
		}, user)
		t.Assert(user.CreateTime.UTC().String(), now.UTC().String())
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			CreateTime *time.Time
		}
		now := time.Now()
		user := new(User)
		转换类.Struct(g.Map{
			"create_time": &now,
		}, user)
		t.Assert(user.CreateTime.UTC().String(), now.UTC().String())
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			CreateTime *时间类.Time
		}
		now := time.Now()
		user := new(User)
		转换类.Struct(g.Map{
			"create_time": &now,
		}, user)
		t.Assert(user.CreateTime.Time.UTC().String(), now.UTC().String())
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			CreateTime 时间类.Time
		}
		now := time.Now()
		user := new(User)
		转换类.Struct(g.Map{
			"create_time": &now,
		}, user)
		t.Assert(user.CreateTime.Time.UTC().String(), now.UTC().String())
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			CreateTime 时间类.Time
		}
		now := time.Now()
		user := new(User)
		转换类.Struct(g.Map{
			"create_time": now,
		}, user)
		t.Assert(user.CreateTime.Time.UTC().String(), now.UTC().String())
	})
}
