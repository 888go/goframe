// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类_test

import (
	"testing"

	gvar "github.com/888go/goframe/container/gvar"
	gdb "github.com/888go/goframe/database/gdb"
	gjson "github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func Test_Slice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := 123.456
		t.AssertEQ(gconv.X取字节集("123"), []byte("123"))
		t.AssertEQ(gconv.X取字节集([]interface{}{1}), []byte{1})
		t.AssertEQ(gconv.X取字节集([]interface{}{300}), []byte("[300]"))
		t.AssertEQ(gconv.X取文本切片(value), []string{"123.456"})
		t.AssertEQ(gconv.SliceStr别名(value), []string{"123.456"})
		t.AssertEQ(gconv.SliceIne别名(value), []int{123})
		t.AssertEQ(gconv.SliceUint别名(value), []uint{123})
		t.AssertEQ(gconv.SliceUint32别名(value), []uint32{123})
		t.AssertEQ(gconv.SliceUint64别名(value), []uint64{123})
		t.AssertEQ(gconv.SliceIet32别名(value), []int32{123})
		t.AssertEQ(gconv.SliceInt64别名(value), []int64{123})
		t.AssertEQ(gconv.X取整数切片(value), []int{123})
		t.AssertEQ(gconv.SliceFloat别名(value), []float64{123.456})
		t.AssertEQ(gconv.X取小数切片(value), []float64{123.456})
		t.AssertEQ(gconv.SliceFloat32别名(value), []float32{123.456})
		t.AssertEQ(gconv.SliceFloat64别名(value), []float64{123.456})
		t.AssertEQ(gconv.X取any切片(value), []interface{}{123.456})
		t.AssertEQ(gconv.SliceAny别名(" [26, 27] "), []interface{}{26, 27})
	})
	gtest.C(t, func(t *gtest.T) {
		s := gvar.Vars{
			gvar.X创建(1),
			gvar.X创建(2),
		}
		t.AssertEQ(gconv.SliceInt64别名(s), []int64{1, 2})
	})
}

func Test_Slice_Ints(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取整数切片(nil), nil)
		t.AssertEQ(gconv.X取整数切片("[26, 27]"), []int{26, 27})
		t.AssertEQ(gconv.X取整数切片(" [26, 27] "), []int{26, 27})
		t.AssertEQ(gconv.X取整数切片([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []int{0, 0})
		t.AssertEQ(gconv.X取整数切片([]bool{true, false}), []int{1, 0})
		t.AssertEQ(gconv.X取整数切片([][]byte{{byte(1)}, {byte(2)}}), []int{1, 2})
	})
}

func Test_Slice_Int32s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取整数32位切片(nil), nil)
		t.AssertEQ(gconv.X取整数32位切片(" [26, 27] "), []int32{26, 27})
		t.AssertEQ(gconv.X取整数32位切片([]string{"1", "2"}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]int{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]int8{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]int16{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]int32{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]int64{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]uint{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]uint8{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []int32{0, 0})
		t.AssertEQ(gconv.X取整数32位切片([]uint16{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]uint32{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]uint64{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]bool{true, false}), []int32{1, 0})
		t.AssertEQ(gconv.X取整数32位切片([]float32{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([]float64{1, 2}), []int32{1, 2})
		t.AssertEQ(gconv.X取整数32位切片([][]byte{{byte(1)}, {byte(2)}}), []int32{1, 2})

		s := gvar.Vars{
			gvar.X创建(1),
			gvar.X创建(2),
		}
		t.AssertEQ(gconv.SliceIet32别名(s), []int32{1, 2})
	})
}

func Test_Slice_Int64s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取整数64位切片(nil), nil)
		t.AssertEQ(gconv.X取整数64位切片(" [26, 27] "), []int64{26, 27})
		t.AssertEQ(gconv.X取整数64位切片([]string{"1", "2"}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]int{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]int8{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]int16{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]int32{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]int64{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]uint{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]uint8{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []int64{0, 0})
		t.AssertEQ(gconv.X取整数64位切片([]uint16{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]uint32{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]uint64{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]bool{true, false}), []int64{1, 0})
		t.AssertEQ(gconv.X取整数64位切片([]float32{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([]float64{1, 2}), []int64{1, 2})
		t.AssertEQ(gconv.X取整数64位切片([][]byte{{byte(1)}, {byte(2)}}), []int64{1, 2})

		s := gvar.Vars{
			gvar.X创建(1),
			gvar.X创建(2),
		}
		t.AssertEQ(gconv.X取整数64位切片(s), []int64{1, 2})
	})
}

func Test_Slice_Uints(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取正整数切片(nil), nil)
		t.AssertEQ(gconv.X取正整数切片("1"), []uint{1})
		t.AssertEQ(gconv.X取正整数切片(" [26, 27] "), []uint{26, 27})
		t.AssertEQ(gconv.X取正整数切片([]string{"1", "2"}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]int{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]int8{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]int16{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]int32{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]int64{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]uint{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]uint8{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []uint{0, 0})
		t.AssertEQ(gconv.X取正整数切片([]uint16{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]uint32{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]uint64{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]bool{true, false}), []uint{1, 0})
		t.AssertEQ(gconv.X取正整数切片([]float32{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([]float64{1, 2}), []uint{1, 2})
		t.AssertEQ(gconv.X取正整数切片([][]byte{{byte(1)}, {byte(2)}}), []uint{1, 2})

		s := gvar.Vars{
			gvar.X创建(1),
			gvar.X创建(2),
		}
		t.AssertEQ(gconv.X取正整数切片(s), []uint{1, 2})
	})
}

func Test_Slice_Uint32s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取正整数32位切片(nil), nil)
		t.AssertEQ(gconv.X取正整数32位切片("1"), []uint32{1})
		t.AssertEQ(gconv.X取正整数32位切片(" [26, 27] "), []uint32{26, 27})
		t.AssertEQ(gconv.X取正整数32位切片([]string{"1", "2"}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]int{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]int8{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]int16{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]int32{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]int64{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]uint{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]uint8{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []uint32{0, 0})
		t.AssertEQ(gconv.X取正整数32位切片([]uint16{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]uint32{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]uint64{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]bool{true, false}), []uint32{1, 0})
		t.AssertEQ(gconv.X取正整数32位切片([]float32{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([]float64{1, 2}), []uint32{1, 2})
		t.AssertEQ(gconv.X取正整数32位切片([][]byte{{byte(1)}, {byte(2)}}), []uint32{1, 2})

		s := gvar.Vars{
			gvar.X创建(1),
			gvar.X创建(2),
		}
		t.AssertEQ(gconv.X取正整数32位切片(s), []uint32{1, 2})
	})
}

func Test_Slice_Uint64s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取正整数64位切片(nil), nil)
		t.AssertEQ(gconv.X取正整数64位切片("1"), []uint64{1})
		t.AssertEQ(gconv.X取正整数64位切片(" [26, 27] "), []uint64{26, 27})
		t.AssertEQ(gconv.X取正整数64位切片([]string{"1", "2"}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]int{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]int8{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]int16{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]int32{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]int64{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]uint{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]uint8{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []uint64{0, 0})
		t.AssertEQ(gconv.X取正整数64位切片([]uint16{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]uint64{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]uint64{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]bool{true, false}), []uint64{1, 0})
		t.AssertEQ(gconv.X取正整数64位切片([]float32{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([]float64{1, 2}), []uint64{1, 2})
		t.AssertEQ(gconv.X取正整数64位切片([][]byte{{byte(1)}, {byte(2)}}), []uint64{1, 2})

		s := gvar.Vars{
			gvar.X创建(1),
			gvar.X创建(2),
		}
		t.AssertEQ(gconv.X取正整数64位切片(s), []uint64{1, 2})
	})
}

func Test_Slice_Float32s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取小数32位切片("123.4"), []float32{123.4})
		t.AssertEQ(gconv.X取小数32位切片([]string{"123.4", "123.5"}), []float32{123.4, 123.5})
		t.AssertEQ(gconv.X取小数32位切片([]int{123}), []float32{123})
		t.AssertEQ(gconv.X取小数32位切片([]int8{123}), []float32{123})
		t.AssertEQ(gconv.X取小数32位切片([]int16{123}), []float32{123})
		t.AssertEQ(gconv.X取小数32位切片([]int32{123}), []float32{123})
		t.AssertEQ(gconv.X取小数32位切片([]int64{123}), []float32{123})
		t.AssertEQ(gconv.X取小数32位切片([]uint{123}), []float32{123})
		t.AssertEQ(gconv.X取小数32位切片([]uint8{123}), []float32{123})
		t.AssertEQ(gconv.X取小数32位切片([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []float32{0, 0})
		t.AssertEQ(gconv.X取小数32位切片([]uint16{123}), []float32{123})
		t.AssertEQ(gconv.X取小数32位切片([]uint32{123}), []float32{123})
		t.AssertEQ(gconv.X取小数32位切片([]uint64{123}), []float32{123})
		t.AssertEQ(gconv.X取小数32位切片([]bool{true, false}), []float32{0, 0})
		t.AssertEQ(gconv.X取小数32位切片([]float32{123}), []float32{123})
		t.AssertEQ(gconv.X取小数32位切片([]float64{123}), []float32{123})

		s := gvar.Vars{
			gvar.X创建(1.1),
			gvar.X创建(2.1),
		}
		t.AssertEQ(gconv.SliceFloat32别名(s), []float32{1.1, 2.1})
	})
}

func Test_Slice_Float64s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取小数64位切片("123.4"), []float64{123.4})
		t.AssertEQ(gconv.X取小数64位切片([]string{"123.4", "123.5"}), []float64{123.4, 123.5})
		t.AssertEQ(gconv.X取小数64位切片([]int{123}), []float64{123})
		t.AssertEQ(gconv.X取小数64位切片([]int8{123}), []float64{123})
		t.AssertEQ(gconv.X取小数64位切片([]int16{123}), []float64{123})
		t.AssertEQ(gconv.X取小数64位切片([]int32{123}), []float64{123})
		t.AssertEQ(gconv.X取小数64位切片([]int64{123}), []float64{123})
		t.AssertEQ(gconv.X取小数64位切片([]uint{123}), []float64{123})
		t.AssertEQ(gconv.X取小数64位切片([]uint8{123}), []float64{123})
		t.AssertEQ(gconv.X取小数64位切片([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []float64{0, 0})
		t.AssertEQ(gconv.X取小数64位切片([]uint16{123}), []float64{123})
		t.AssertEQ(gconv.X取小数64位切片([]uint32{123}), []float64{123})
		t.AssertEQ(gconv.X取小数64位切片([]uint64{123}), []float64{123})
		t.AssertEQ(gconv.X取小数64位切片([]bool{true, false}), []float64{0, 0})
		t.AssertEQ(gconv.X取小数64位切片([]float32{123}), []float64{123})
		t.AssertEQ(gconv.X取小数64位切片([]float64{123}), []float64{123})
	})
}

func Test_Slice_Empty(t *testing.T) {
	// Int.
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取整数切片(""), []int{})
		t.Assert(gconv.X取整数切片(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取整数32位切片(""), []int32{})
		t.Assert(gconv.X取整数32位切片(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取整数64位切片(""), []int64{})
		t.Assert(gconv.X取整数64位切片(nil), nil)
	})
	// Uint.
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取正整数切片(""), []uint{})
		t.Assert(gconv.X取正整数切片(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取正整数32位切片(""), []uint32{})
		t.Assert(gconv.X取正整数32位切片(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取正整数64位切片(""), []uint64{})
		t.Assert(gconv.X取正整数64位切片(nil), nil)
	})
	// Float.
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取小数切片(""), []float64{})
		t.Assert(gconv.X取小数切片(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取小数32位切片(""), []float32{})
		t.Assert(gconv.X取小数32位切片(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取小数64位切片(""), []float64{})
		t.Assert(gconv.X取小数64位切片(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取文本切片(""), []string{})
		t.Assert(gconv.X取文本切片(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.SliceAny别名(""), []interface{}{""})
		t.Assert(gconv.SliceAny别名(nil), nil)
	})
}

func Test_Strings(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := []*g.Var{
			g.X泛型类(1),
			g.X泛型类(2),
			g.X泛型类(3),
		}
		t.AssertEQ(gconv.X取文本切片(array), []string{"1", "2", "3"})

		t.AssertEQ(gconv.X取文本切片([]uint8(`["1","2"]`)), []string{"1", "2"})
		t.AssertEQ(gconv.X取文本切片([][]byte{{byte(0)}, {byte(1)}}), []string{"\u0000", "\u0001"})
	})
			//github.com/gogf/gf/issues/1750. md5:b86a24bc52c53801
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取文本切片("123"), []string{"123"})
	})
		// 这段注释引用的是GitHub上的一个 issue，gf（Go Foundation）是一个用Go语言编写的开源框架。3465号issue可能是指该框架中的某个问题或讨论的编号。具体的内容需要查看相关链接才能了解详情。 md5:53810ebfb659d15e
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取文本切片("null"), []string{"null"})
		t.AssertEQ(gconv.X取文本切片([]byte("null")), []string{"110", "117", "108", "108"})
		t.AssertEQ(gconv.X取文本切片("{\"name\":\"wln\"}"), []string{"{\"name\":\"wln\"}"})
	})
}

func Test_Slice_Interfaces(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := gconv.X取any切片([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`))
		t.Assert(len(array), 2)
		t.Assert(array[0].(g.Map)["id"], 1)
		t.Assert(array[0].(g.Map)["name"], "john")
	})
	// map
	gtest.C(t, func(t *gtest.T) {
		array := gconv.X取any切片(g.Map{
			"id":   1,
			"name": "john",
		})
		t.Assert(len(array), 1)
		t.Assert(array[0].(g.Map)["id"], 1)
		t.Assert(array[0].(g.Map)["name"], "john")
	})
	// struct
	gtest.C(t, func(t *gtest.T) {
		type A struct {
			Id   int `json:"id"`
			Name string
		}
		array := gconv.X取any切片(&A{
			Id:   1,
			Name: "john",
		})
		t.Assert(len(array), 1)
		t.Assert(array[0].(*A).Id, 1)
		t.Assert(array[0].(*A).Name, "john")
	})
}

func Test_Slice_PrivateAttribute(t *testing.T) {
	type User struct {
		Id   int    `json:"id"`
		name string `json:"name"`
	}
	gtest.C(t, func(t *gtest.T) {
		user := &User{1, "john"}
		array := gconv.X取any切片(user)
		t.Assert(len(array), 1)
		t.Assert(array[0].(*User).Id, 1)
		t.Assert(array[0].(*User).name, "john")
	})
}

func Test_Slice_Structs(t *testing.T) {
	type Base struct {
		Age int
	}
	type User struct {
		Id   int
		Name string
		Base
	}

	gtest.C(t, func(t *gtest.T) {
		users := make([]User, 0)
		params := []g.Map{
			{"id": 1, "name": "john", "age": 18},
			{"id": 2, "name": "smith", "age": 20},
		}
		err := gconv.Structs(params, &users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, params[0]["id"])
		t.Assert(users[0].Name, params[0]["name"])
		t.Assert(users[0].Age, 18)

		t.Assert(users[1].Id, params[1]["id"])
		t.Assert(users[1].Name, params[1]["name"])
		t.Assert(users[1].Age, 20)
	})

	gtest.C(t, func(t *gtest.T) {
		users := make([]User, 0)
		params := []g.Map{
			{"id": 1, "name": "john", "age": 18},
			{"id": 2, "name": "smith", "age": 20},
		}
		err := gconv.StructsTag(params, &users, "")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, params[0]["id"])
		t.Assert(users[0].Name, params[0]["name"])
		t.Assert(users[0].Age, 18)

		t.Assert(users[1].Id, params[1]["id"])
		t.Assert(users[1].Name, params[1]["name"])
		t.Assert(users[1].Age, 20)
	})
}

func Test_EmptyString_To_CustomType(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Status string
		type Req struct {
			Name     string
			Statuses []Status
			Types    []string
		}
		var (
			req  *Req
			data = g.Map{
				"Name":     "john",
				"Statuses": "",
				"Types":    "",
			}
		)
		err := gconv.Scan(data, &req)
		t.AssertNil(err)
		t.Assert(len(req.Statuses), 0)
		t.Assert(len(req.Types), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		type Status string
		type Req struct {
			Name     string
			Statuses []*Status
			Types    []string
		}
		var (
			req  *Req
			data = g.Map{
				"Name":     "john",
				"Statuses": "",
				"Types":    "",
			}
		)
		err := gconv.Scan(data, &req)
		t.AssertNil(err)
		t.Assert(len(req.Statuses), 0)
		t.Assert(len(req.Types), 0)
	})
}

func Test_SliceMap_WithNilMapValue(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			list1 = []gdb.Record{
				{"name": nil},
			}
			list2 []map[string]any
		)
		list2 = gconv.SliceMap别名(list1)
		t.Assert(len(list2), 1)
		t.Assert(list1[0], list2[0])
		t.Assert(gjson.X变量到json文本PANI(list1), gjson.X变量到json文本PANI(list2))
	})
}
