// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类_test

import (
	"testing"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Slice(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		value := 123.456
		t.AssertEQ(转换类.X取字节集("123"), []byte("123"))
		t.AssertEQ(转换类.X取字节集([]interface{}{1}), []byte{1})
		t.AssertEQ(转换类.X取字节集([]interface{}{300}), []byte("[300]"))
		t.AssertEQ(转换类.X取文本数组(value), []string{"123.456"})
		t.AssertEQ(转换类.SliceStr别名(value), []string{"123.456"})
		t.AssertEQ(转换类.SliceIne别名(value), []int{123})
		t.AssertEQ(转换类.SliceUint别名(value), []uint{123})
		t.AssertEQ(转换类.SliceUint32别名(value), []uint32{123})
		t.AssertEQ(转换类.SliceUint64别名(value), []uint64{123})
		t.AssertEQ(转换类.SliceIet32别名(value), []int32{123})
		t.AssertEQ(转换类.SliceInt64别名(value), []int64{123})
		t.AssertEQ(转换类.X取整数数组(value), []int{123})
		t.AssertEQ(转换类.SliceFloat别名(value), []float64{123.456})
		t.AssertEQ(转换类.X取小数数组(value), []float64{123.456})
		t.AssertEQ(转换类.SliceFloat32别名(value), []float32{123.456})
		t.AssertEQ(转换类.SliceFloat64别名(value), []float64{123.456})
		t.AssertEQ(转换类.X取any数组(value), []interface{}{123.456})
		t.AssertEQ(转换类.SliceAny别名(" [26, 27] "), []interface{}{26, 27})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		s := 泛型类.Vars{
			泛型类.X创建(1),
			泛型类.X创建(2),
		}
		t.AssertEQ(转换类.SliceInt64别名(s), []int64{1, 2})
	})
}

func Test_Slice_Ints(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取整数数组(nil), nil)
		t.AssertEQ(转换类.X取整数数组("[26, 27]"), []int{26, 27})
		t.AssertEQ(转换类.X取整数数组(" [26, 27] "), []int{26, 27})
		t.AssertEQ(转换类.X取整数数组([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []int{0, 0})
		t.AssertEQ(转换类.X取整数数组([]bool{true, false}), []int{1, 0})
		t.AssertEQ(转换类.X取整数数组([][]byte{{byte(1)}, {byte(2)}}), []int{1, 2})
	})
}

func Test_Slice_Int32s(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取整数32位数组(nil), nil)
		t.AssertEQ(转换类.X取整数32位数组(" [26, 27] "), []int32{26, 27})
		t.AssertEQ(转换类.X取整数32位数组([]string{"1", "2"}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]int{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]int8{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]int16{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]int32{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]int64{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]uint{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]uint8{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []int32{0, 0})
		t.AssertEQ(转换类.X取整数32位数组([]uint16{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]uint32{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]uint64{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]bool{true, false}), []int32{1, 0})
		t.AssertEQ(转换类.X取整数32位数组([]float32{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([]float64{1, 2}), []int32{1, 2})
		t.AssertEQ(转换类.X取整数32位数组([][]byte{{byte(1)}, {byte(2)}}), []int32{1, 2})

		s := 泛型类.Vars{
			泛型类.X创建(1),
			泛型类.X创建(2),
		}
		t.AssertEQ(转换类.SliceIet32别名(s), []int32{1, 2})
	})
}

func Test_Slice_Int64s(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取整数64位数组(nil), nil)
		t.AssertEQ(转换类.X取整数64位数组(" [26, 27] "), []int64{26, 27})
		t.AssertEQ(转换类.X取整数64位数组([]string{"1", "2"}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]int{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]int8{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]int16{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]int32{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]int64{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]uint{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]uint8{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []int64{0, 0})
		t.AssertEQ(转换类.X取整数64位数组([]uint16{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]uint32{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]uint64{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]bool{true, false}), []int64{1, 0})
		t.AssertEQ(转换类.X取整数64位数组([]float32{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([]float64{1, 2}), []int64{1, 2})
		t.AssertEQ(转换类.X取整数64位数组([][]byte{{byte(1)}, {byte(2)}}), []int64{1, 2})

		s := 泛型类.Vars{
			泛型类.X创建(1),
			泛型类.X创建(2),
		}
		t.AssertEQ(转换类.X取整数64位数组(s), []int64{1, 2})
	})
}

func Test_Slice_Uints(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取正整数数组(nil), nil)
		t.AssertEQ(转换类.X取正整数数组("1"), []uint{1})
		t.AssertEQ(转换类.X取正整数数组(" [26, 27] "), []uint{26, 27})
		t.AssertEQ(转换类.X取正整数数组([]string{"1", "2"}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]int{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]int8{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]int16{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]int32{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]int64{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]uint{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]uint8{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []uint{0, 0})
		t.AssertEQ(转换类.X取正整数数组([]uint16{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]uint32{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]uint64{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]bool{true, false}), []uint{1, 0})
		t.AssertEQ(转换类.X取正整数数组([]float32{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([]float64{1, 2}), []uint{1, 2})
		t.AssertEQ(转换类.X取正整数数组([][]byte{{byte(1)}, {byte(2)}}), []uint{1, 2})

		s := 泛型类.Vars{
			泛型类.X创建(1),
			泛型类.X创建(2),
		}
		t.AssertEQ(转换类.X取正整数数组(s), []uint{1, 2})
	})
}

func Test_Slice_Uint32s(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取正整数32位数组(nil), nil)
		t.AssertEQ(转换类.X取正整数32位数组("1"), []uint32{1})
		t.AssertEQ(转换类.X取正整数32位数组(" [26, 27] "), []uint32{26, 27})
		t.AssertEQ(转换类.X取正整数32位数组([]string{"1", "2"}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]int{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]int8{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]int16{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]int32{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]int64{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]uint{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]uint8{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []uint32{0, 0})
		t.AssertEQ(转换类.X取正整数32位数组([]uint16{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]uint32{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]uint64{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]bool{true, false}), []uint32{1, 0})
		t.AssertEQ(转换类.X取正整数32位数组([]float32{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([]float64{1, 2}), []uint32{1, 2})
		t.AssertEQ(转换类.X取正整数32位数组([][]byte{{byte(1)}, {byte(2)}}), []uint32{1, 2})

		s := 泛型类.Vars{
			泛型类.X创建(1),
			泛型类.X创建(2),
		}
		t.AssertEQ(转换类.X取正整数32位数组(s), []uint32{1, 2})
	})
}

func Test_Slice_Uint64s(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取正整数64位数组(nil), nil)
		t.AssertEQ(转换类.X取正整数64位数组("1"), []uint64{1})
		t.AssertEQ(转换类.X取正整数64位数组(" [26, 27] "), []uint64{26, 27})
		t.AssertEQ(转换类.X取正整数64位数组([]string{"1", "2"}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]int{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]int8{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]int16{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]int32{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]int64{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]uint{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]uint8{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []uint64{0, 0})
		t.AssertEQ(转换类.X取正整数64位数组([]uint16{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]uint64{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]uint64{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]bool{true, false}), []uint64{1, 0})
		t.AssertEQ(转换类.X取正整数64位数组([]float32{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([]float64{1, 2}), []uint64{1, 2})
		t.AssertEQ(转换类.X取正整数64位数组([][]byte{{byte(1)}, {byte(2)}}), []uint64{1, 2})

		s := 泛型类.Vars{
			泛型类.X创建(1),
			泛型类.X创建(2),
		}
		t.AssertEQ(转换类.X取正整数64位数组(s), []uint64{1, 2})
	})
}

func Test_Slice_Float32s(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取小数32位数组("123.4"), []float32{123.4})
		t.AssertEQ(转换类.X取小数32位数组([]string{"123.4", "123.5"}), []float32{123.4, 123.5})
		t.AssertEQ(转换类.X取小数32位数组([]int{123}), []float32{123})
		t.AssertEQ(转换类.X取小数32位数组([]int8{123}), []float32{123})
		t.AssertEQ(转换类.X取小数32位数组([]int16{123}), []float32{123})
		t.AssertEQ(转换类.X取小数32位数组([]int32{123}), []float32{123})
		t.AssertEQ(转换类.X取小数32位数组([]int64{123}), []float32{123})
		t.AssertEQ(转换类.X取小数32位数组([]uint{123}), []float32{123})
		t.AssertEQ(转换类.X取小数32位数组([]uint8{123}), []float32{123})
		t.AssertEQ(转换类.X取小数32位数组([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []float32{0, 0})
		t.AssertEQ(转换类.X取小数32位数组([]uint16{123}), []float32{123})
		t.AssertEQ(转换类.X取小数32位数组([]uint32{123}), []float32{123})
		t.AssertEQ(转换类.X取小数32位数组([]uint64{123}), []float32{123})
		t.AssertEQ(转换类.X取小数32位数组([]bool{true, false}), []float32{0, 0})
		t.AssertEQ(转换类.X取小数32位数组([]float32{123}), []float32{123})
		t.AssertEQ(转换类.X取小数32位数组([]float64{123}), []float32{123})

		s := 泛型类.Vars{
			泛型类.X创建(1.1),
			泛型类.X创建(2.1),
		}
		t.AssertEQ(转换类.SliceFloat32别名(s), []float32{1.1, 2.1})
	})
}

func Test_Slice_Float64s(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取小数64位数组("123.4"), []float64{123.4})
		t.AssertEQ(转换类.X取小数64位数组([]string{"123.4", "123.5"}), []float64{123.4, 123.5})
		t.AssertEQ(转换类.X取小数64位数组([]int{123}), []float64{123})
		t.AssertEQ(转换类.X取小数64位数组([]int8{123}), []float64{123})
		t.AssertEQ(转换类.X取小数64位数组([]int16{123}), []float64{123})
		t.AssertEQ(转换类.X取小数64位数组([]int32{123}), []float64{123})
		t.AssertEQ(转换类.X取小数64位数组([]int64{123}), []float64{123})
		t.AssertEQ(转换类.X取小数64位数组([]uint{123}), []float64{123})
		t.AssertEQ(转换类.X取小数64位数组([]uint8{123}), []float64{123})
		t.AssertEQ(转换类.X取小数64位数组([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`)), []float64{0, 0})
		t.AssertEQ(转换类.X取小数64位数组([]uint16{123}), []float64{123})
		t.AssertEQ(转换类.X取小数64位数组([]uint32{123}), []float64{123})
		t.AssertEQ(转换类.X取小数64位数组([]uint64{123}), []float64{123})
		t.AssertEQ(转换类.X取小数64位数组([]bool{true, false}), []float64{0, 0})
		t.AssertEQ(转换类.X取小数64位数组([]float32{123}), []float64{123})
		t.AssertEQ(转换类.X取小数64位数组([]float64{123}), []float64{123})
	})
}

func Test_Slice_Empty(t *testing.T) {
	// Int.
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取整数数组(""), []int{})
		t.Assert(转换类.X取整数数组(nil), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取整数32位数组(""), []int32{})
		t.Assert(转换类.X取整数32位数组(nil), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取整数64位数组(""), []int64{})
		t.Assert(转换类.X取整数64位数组(nil), nil)
	})
	// Uint.
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取正整数数组(""), []uint{})
		t.Assert(转换类.X取正整数数组(nil), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取正整数32位数组(""), []uint32{})
		t.Assert(转换类.X取正整数32位数组(nil), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取正整数64位数组(""), []uint64{})
		t.Assert(转换类.X取正整数64位数组(nil), nil)
	})
	// Float.
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取小数数组(""), []float64{})
		t.Assert(转换类.X取小数数组(nil), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取小数32位数组(""), []float32{})
		t.Assert(转换类.X取小数32位数组(nil), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取小数64位数组(""), []float64{})
		t.Assert(转换类.X取小数64位数组(nil), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取文本数组(""), []string{})
		t.Assert(转换类.X取文本数组(nil), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.SliceAny别名(""), []interface{}{""})
		t.Assert(转换类.SliceAny别名(nil), nil)
	})
}

func Test_Strings(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := []*g.Var{
			g.X泛型类(1),
			g.X泛型类(2),
			g.X泛型类(3),
		}
		t.AssertEQ(转换类.X取文本数组(array), []string{"1", "2", "3"})

		t.AssertEQ(转换类.X取文本数组([]uint8(`["1","2"]`)), []string{"1", "2"})
		t.AssertEQ(转换类.X取文本数组([][]byte{{byte(0)}, {byte(1)}}), []string{"\u0000", "\u0001"})
	})
	// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题：#1750
// 对应中文注释：
// 参考GitHub上gogf/gf项目编号为1750的问题
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取文本数组("123"), []string{"123"})
	})
}

func Test_Slice_Interfaces(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 转换类.X取any数组([]uint8(`[{"id": 1, "name":"john"},{"id": 2, "name":"huang"}]`))
		t.Assert(len(array), 2)
		t.Assert(array[0].(g.Map)["id"], 1)
		t.Assert(array[0].(g.Map)["name"], "john")
	})
	// map
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 转换类.X取any数组(g.Map{
			"id":   1,
			"name": "john",
		})
		t.Assert(len(array), 1)
		t.Assert(array[0].(g.Map)["id"], 1)
		t.Assert(array[0].(g.Map)["name"], "john")
	})
	// struct
	单元测试类.C(t, func(t *单元测试类.T) {
		type A struct {
			Id   int `json:"id"`
			Name string
		}
		array := 转换类.X取any数组(&A{
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
	单元测试类.C(t, func(t *单元测试类.T) {
		user := &User{1, "john"}
		array := 转换类.X取any数组(user)
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

	单元测试类.C(t, func(t *单元测试类.T) {
		users := make([]User, 0)
		params := []g.Map{
			{"id": 1, "name": "john", "age": 18},
			{"id": 2, "name": "smith", "age": 20},
		}
		err := 转换类.Structs(params, &users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, params[0]["id"])
		t.Assert(users[0].Name, params[0]["name"])
		t.Assert(users[0].Age, 18)

		t.Assert(users[1].Id, params[1]["id"])
		t.Assert(users[1].Name, params[1]["name"])
		t.Assert(users[1].Age, 20)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		users := make([]User, 0)
		params := []g.Map{
			{"id": 1, "name": "john", "age": 18},
			{"id": 2, "name": "smith", "age": 20},
		}
		err := 转换类.StructsTag(params, &users, "")
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
	单元测试类.C(t, func(t *单元测试类.T) {
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
		err := 转换类.Scan(data, &req)
		t.AssertNil(err)
		t.Assert(len(req.Statuses), 0)
		t.Assert(len(req.Types), 0)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
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
		err := 转换类.Scan(data, &req)
		t.AssertNil(err)
		t.Assert(len(req.Statuses), 0)
		t.Assert(len(req.Types), 0)
	})
}

func Test_SliceMap_WithNilMapValue(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			list1 = []db类.Record{
				{"name": nil},
			}
			list2 []map[string]any
		)
		list2 = 转换类.SliceMap别名(list1)
		t.Assert(len(list2), 1)
		t.Assert(list1[0], list2[0])
		t.Assert(json类.X变量到json文本PANI(list1), json类.X变量到json文本PANI(list2))
	})
}
