// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// go test *.go

package 切片类_test

import (
	"strings"
	"testing"

	garray "github.com/888go/goframe/container/garray"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

func Test_Array_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var array garray.Array
		expect := []int{2, 3, 1}
		array.Append别名(2, 3, 1)
		t.Assert(array.X取切片(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array garray.IntArray
		expect := []int{2, 3, 1}
		array.Append别名(2, 3, 1)
		t.Assert(array.X取切片(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array garray.StrArray
		expect := []string{"b", "a"}
		array.Append别名("b", "a")
		t.Assert(array.X取切片(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array garray.SortedArray
		array.X设置排序函数(gutil.X比较整数)
		expect := []int{1, 2, 3}
		array.X入栈右(2, 3, 1)
		t.Assert(array.X取切片(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array garray.SortedIntArray
		expect := []int{1, 2, 3}
		array.X入栈右(2, 3, 1)
		t.Assert(array.X取切片(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array garray.SortedStrArray
		expect := []string{"a", "b", "c"}
		array.X入栈右("c", "a", "b")
		t.Assert(array.X取切片(), expect)
	})
}

func Test_SortedIntArray_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var array garray.SortedIntArray
		expect := []int{1, 2, 3}
		array.X入栈右(2, 3, 1)
		t.Assert(array.X取切片(), expect)
	})
}

func Test_IntArray_Unique(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []int{1, 2, 3, 4, 5, 6}
		array := garray.X创建整数()
		array.Append别名(1, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6)
		array.X去重()
		t.Assert(array.X取切片(), expect)
	})
}

func Test_SortedIntArray1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		array := garray.X创建整数排序()
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
		array := garray.X创建整数排序()
		for i := 0; i <= 10; i++ {
			array.X入栈右(i)
		}
		t.Assert(array.X取切片(), expect)
	})
}

func Test_SortedStrArray1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array1 := garray.X创建文本排序()
		array2 := garray.X创建文本排序(true)
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
		array := garray.X创建文本排序()
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
		array := garray.X创建排序(func(v1, v2 interface{}) int {
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
		array := garray.X创建排序(func1)
		array2 := garray.X创建排序(func1, true)
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
		array1 := garray.NewFromCopy别名(a1)
		t.AssertIN(array1.X出栈随机多个(2), a1)
		t.Assert(len(array1.X出栈随机多个(1)), 1)
		t.Assert(len(array1.X出栈随机多个(9)), 3)
	})
}
