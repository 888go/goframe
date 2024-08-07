// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 泛型类_test

import (
	"testing"

	gvar "github.com/888go/goframe/container/gvar"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestVar_Ints(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := gvar.X创建(arr, true)
		t.Assert(objOne.X取整数切片()[0], arr[0])
	})
}

func TestVar_Uints(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := gvar.X创建(arr, true)
		t.Assert(objOne.X取正整数切片()[0], arr[0])
	})
}

func TestVar_Int64s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := gvar.X创建(arr, true)
		t.Assert(objOne.X取整数64位切片()[0], arr[0])
	})
}

func TestVar_Uint64s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := gvar.X创建(arr, true)
		t.Assert(objOne.X取正整数64位切片()[0], arr[0])
	})
}

func TestVar_Floats(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []float64{1, 2, 3, 4, 5}
		objOne := gvar.X创建(arr, true)
		t.Assert(objOne.X取小数切片()[0], arr[0])
	})
}

func TestVar_Float32s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []float32{1, 2, 3, 4, 5}
		objOne := gvar.X创建(arr, true)
		t.AssertEQ(objOne.X取小数32位切片(), arr)
	})
}

func TestVar_Float64s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []float64{1, 2, 3, 4, 5}
		objOne := gvar.X创建(arr, true)
		t.AssertEQ(objOne.X取小数64位切片(), arr)
	})
}

func TestVar_Strings(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []string{"hello", "world"}
		objOne := gvar.X创建(arr, true)
		t.Assert(objOne.X取文本切片()[0], arr[0])
	})
}

func TestVar_Interfaces(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := gvar.X创建(arr, true)
		t.Assert(objOne.X取any切片(), arr)
	})
}

func TestVar_Slice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := gvar.X创建(arr, true)
		t.Assert(objOne.Slice别名(), arr)
	})
}

func TestVar_Array(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := gvar.X创建(arr, false)
		t.Assert(objOne.Array别名(), arr)
	})
}

func TestVar_Vars(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := gvar.X创建(arr, false)
		t.Assert(len(objOne.X取泛型类切片()), 5)
		t.Assert(objOne.X取泛型类切片()[0].X取整数(), 1)
		t.Assert(objOne.X取泛型类切片()[4].X取整数(), 5)

		objEmpty := gvar.X创建([]int{})
		t.Assert(objEmpty.X取泛型类切片(), nil)
	})
}
