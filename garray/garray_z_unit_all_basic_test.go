// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// go test *.go

package 数组类_test

import (
	"strings"
	"testing"
	
	"github.com/888go/goframe/garray"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

func Test_Array_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var array 数组类.Array
		expect := []int{2, 3, 1}
		array.Append别名(2, 3, 1)
		t.Assert(array.X取切片(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array 数组类.IntArray
		expect := []int{2, 3, 1}
		array.Append别名(2, 3, 1)
		t.Assert(array.X取切片(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array 数组类.StrArray
		expect := []string{"b", "a"}
		array.Append别名("b", "a")
		t.Assert(array.X取切片(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array 数组类.SortedArray
		array.X设置排序函数(gutil.ComparatorInt)
		expect := []int{1, 2, 3}
		array.X入栈右(2, 3, 1)
		t.Assert(array.X取切片(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array 数组类.SortedIntArray
		expect := []int{1, 2, 3}
		array.X入栈右(2, 3, 1)
		t.Assert(array.X取切片(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array 数组类.SortedStrArray
		expect := []string{"a", "b", "c"}
		array.X入栈右("c", "a", "b")
		t.Assert(array.X取切片(), expect)
	})
}

func Test_SortedIntArray_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var array 数组类.SortedIntArray
		expect := []int{1, 2, 3}
		array.X入栈右(2, 3, 1)
		t.Assert(array.X取切片(), expect)
	})
}

func Test_IntArray_Unique(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []int{1, 2, 3, 4, 5, 6}
		array := 数组类.X创建整数()
		array.Append别名(1, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6)
		array.X去重()
		t.Assert(array.X取切片(), expect)
	})
}

func Test_SortedIntArray1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		array := 数组类.X创建整数排序()
		for i := 10; i > -1; i-- {
			array.X入栈右(i)
		}
		t.Assert(array.X取切片(), expect)
		t.Assert(array.X入栈右().X取切片(), expect)
	})
}

func Test_SortedIntArray2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		array := 数组类.X创建整数排序()
		for i := 0; i <= 10; i++ {
			array.X入栈右(i)
		}
		t.Assert(array.X取切片(), expect)
	})
}

func Test_SortedStrArray1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array1 := 数组类.X创建文本排序()
		array2 := 数组类.X创建文本排序(true)
		for i := 10; i > -1; i-- {
			array1.X入栈右(gconv.String(i))
			array2.X入栈右(gconv.String(i))
		}
		t.Assert(array1.X取切片(), expect)
		t.Assert(array2.X取切片(), expect)
	})

}

func Test_SortedStrArray2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array := 数组类.X创建文本排序()
		for i := 0; i <= 10; i++ {
			array.X入栈右(gconv.String(i))
		}
		t.Assert(array.X取切片(), expect)
		array.X入栈右()
		t.Assert(array.X取切片(), expect)
	})
}

func Test_SortedArray1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array := 数组类.X创建排序(func(v1, v2 interface{}) int {
			return strings.Compare(gconv.String(v1), gconv.String(v2))
		})
		for i := 10; i > -1; i-- {
			array.X入栈右(gconv.String(i))
		}
		t.Assert(array.X取切片(), expect)
	})
}

func Test_SortedArray2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(gconv.String(v1), gconv.String(v2))
		}
		array := 数组类.X创建排序(func1)
		array2 := 数组类.X创建排序(func1, true)
		for i := 0; i <= 10; i++ {
			array.X入栈右(gconv.String(i))
			array2.X入栈右(gconv.String(i))
		}
		t.Assert(array.X取切片(), expect)
		t.Assert(array.X入栈右().X取切片(), expect)
		t.Assert(array2.X取切片(), expect)
	})
}

func TestNewFromCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{"100", "200", "300", "400", "500", "600"}
		array1 := 数组类.NewFromCopy别名(a1)
		t.AssertIN(array1.X出栈随机多个(2), a1)
		t.Assert(len(array1.X出栈随机多个(1)), 1)
		t.Assert(len(array1.X出栈随机多个(9)), 3)
	})
}
