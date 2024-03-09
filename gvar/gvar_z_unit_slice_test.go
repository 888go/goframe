// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类_test

import (
	"testing"
	
	"github.com/888go/goframe/gvar"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestVar_Ints(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := 泛型类.X创建(arr, true)
		t.Assert(objOne.X取整数数组()[0], arr[0])
	})
}

func TestVar_Uints(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := 泛型类.X创建(arr, true)
		t.Assert(objOne.X取正整数数组()[0], arr[0])
	})
}

func TestVar_Int64s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := 泛型类.X创建(arr, true)
		t.Assert(objOne.X取整数64位数组()[0], arr[0])
	})
}

func TestVar_Uint64s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := 泛型类.X创建(arr, true)
		t.Assert(objOne.X取正整数64位数组()[0], arr[0])
	})
}

func TestVar_Floats(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []float64{1, 2, 3, 4, 5}
		objOne := 泛型类.X创建(arr, true)
		t.Assert(objOne.X取小数数组()[0], arr[0])
	})
}

func TestVar_Float32s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []float32{1, 2, 3, 4, 5}
		objOne := 泛型类.X创建(arr, true)
		t.AssertEQ(objOne.X取小数32位数组(), arr)
	})
}

func TestVar_Float64s(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []float64{1, 2, 3, 4, 5}
		objOne := 泛型类.X创建(arr, true)
		t.AssertEQ(objOne.X取小数64位数组(), arr)
	})
}

func TestVar_Strings(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []string{"hello", "world"}
		objOne := 泛型类.X创建(arr, true)
		t.Assert(objOne.X取文本数组()[0], arr[0])
	})
}

func TestVar_Interfaces(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := 泛型类.X创建(arr, true)
		t.Assert(objOne.X取any数组(), arr)
	})
}

func TestVar_Slice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := 泛型类.X创建(arr, true)
		t.Assert(objOne.Slice别名(), arr)
	})
}

func TestVar_Array(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := 泛型类.X创建(arr, false)
		t.Assert(objOne.Array别名(), arr)
	})
}

func TestVar_Vars(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := 泛型类.X创建(arr, false)
		t.Assert(len(objOne.X取泛型类数组()), 5)
		t.Assert(objOne.X取泛型类数组()[0].X取整数(), 1)
		t.Assert(objOne.X取泛型类数组()[4].X取整数(), 5)

		objEmpty := 泛型类.X创建([]int{})
		t.Assert(objEmpty.X取泛型类数组(), nil)
	})
}
