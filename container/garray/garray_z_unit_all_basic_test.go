// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// go test *.go

package garray_test

import (
	"strings"
	"testing"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

func Test_Array_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var array garray.Array
		expect := []int{2, 3, 1}
		array.Append(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array garray.IntArray
		expect := []int{2, 3, 1}
		array.Append(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array garray.StrArray
		expect := []string{"b", "a"}
		array.Append("b", "a")
		t.Assert(array.Slice(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array garray.SortedArray
		array.SetComparator(gutil.ComparatorInt)
		expect := []int{1, 2, 3}
		array.Add(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array garray.SortedIntArray
		expect := []int{1, 2, 3}
		array.Add(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		var array garray.SortedStrArray
		expect := []string{"a", "b", "c"}
		array.Add("c", "a", "b")
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedIntArray_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var array garray.SortedIntArray
		expect := []int{1, 2, 3}
		array.Add(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
}

func Test_IntArray_Unique(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []int{1, 2, 3, 4, 5, 6}
		array := garray.NewIntArray()
		array.Append(1, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6)
		array.Unique()
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedIntArray1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		array := garray.NewSortedIntArray()
		for i := 10; i > -1; i-- {
			array.Add(i)
		}
		t.Assert(array.Slice(), expect)
		t.Assert(array.Add().Slice(), expect)
	})
}

func Test_SortedIntArray2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		array := garray.NewSortedIntArray()
		for i := 0; i <= 10; i++ {
			array.Add(i)
		}
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedStrArray1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array1 := garray.NewSortedStrArray()
		array2 := garray.NewSortedStrArray(true)
		for i := 10; i > -1; i-- {
			array1.Add(gconv.String(i))
			array2.Add(gconv.String(i))
		}
		t.Assert(array1.Slice(), expect)
		t.Assert(array2.Slice(), expect)
	})

}

func Test_SortedStrArray2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array := garray.NewSortedStrArray()
		for i := 0; i <= 10; i++ {
			array.Add(gconv.String(i))
		}
		t.Assert(array.Slice(), expect)
		array.Add()
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedArray1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array := garray.NewSortedArray(func(v1, v2 interface{}) int {
			return strings.Compare(gconv.String(v1), gconv.String(v2))
		})
		for i := 10; i > -1; i-- {
			array.Add(gconv.String(i))
		}
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedArray2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(gconv.String(v1), gconv.String(v2))
		}
		array := garray.NewSortedArray(func1)
		array2 := garray.NewSortedArray(func1, true)
		for i := 0; i <= 10; i++ {
			array.Add(gconv.String(i))
			array2.Add(gconv.String(i))
		}
		t.Assert(array.Slice(), expect)
		t.Assert(array.Add().Slice(), expect)
		t.Assert(array2.Slice(), expect)
	})
}

func TestNewFromCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{"100", "200", "300", "400", "500", "600"}
		array1 := garray.NewFromCopy(a1)
		t.AssertIN(array1.PopRands(2), a1)
		t.Assert(len(array1.PopRands(1)), 1)
		t.Assert(len(array1.PopRands(9)), 3)
	})
}
