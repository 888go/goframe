// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 切片类_test

import (
	"fmt"

	"github.com/888go/goframe/internal/empty"

	garray "github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	gconv "github.com/888go/goframe/util/gconv"
)

func ExampleIntArray_Walk() {
	var array garray.IntArray
	tables := g.SliceInt别名{10, 20}
	prefix := 99
	array.Append别名(tables...)
		// 为给定的表名添加前缀。 md5:dea7405f272e0c9e
	array.X遍历修改(func(value int) int {
		return prefix + value
	})
	fmt.Println(array.X取切片())

	// Output:
	// [109 119]
}

func ExampleNewIntArray() {
	s := garray.X创建整数()
	s.Append别名(10)
	s.Append别名(20)
	s.Append别名(15)
	s.Append别名(30)
	fmt.Println(s.X取切片())

	// Output:
	// [10 20 15 30]
}

func ExampleNewIntArraySize() {
	s := garray.X创建整数并按大小(3, 5)
	s.X设置值(0, 10)
	s.X设置值(1, 20)
	s.X设置值(2, 15)
	s.X设置值(3, 30)
	fmt.Println(s.X取切片(), s.X取长度(), cap(s.X取切片()))

	// Output:
	// [10 20 15] 3 5
}

func ExampleNewIntArrayRange() {
	s := garray.X创建整数并按范围(1, 5, 1)
	fmt.Println(s.X取切片(), s.X取长度(), cap(s.X取切片()))

	// Output:
	// [1 2 3 4 5] 5 8
}

func ExampleNewIntArrayFrom() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30})
	fmt.Println(s.X取切片(), s.X取长度(), cap(s.X取切片()))

	// Output:
	// [10 20 15 30] 4 4
}

func ExampleNewIntArrayFromCopy() {
	s := garray.X创建整数并从切片复制(g.SliceInt别名{10, 20, 15, 30})
	fmt.Println(s.X取切片(), s.X取长度(), cap(s.X取切片()))

	// Output:
	// [10 20 15 30] 4 4
}

func ExampleIntArray_At() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30})
	sAt := s.X取值(2)
	fmt.Println(sAt)

	// Output:
	// 15
}

func ExampleIntArray_Get() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30})
	sGet, sBool := s.X取值2(3)
	fmt.Println(sGet, sBool)
	sGet, sBool = s.X取值2(99)
	fmt.Println(sGet, sBool)

	// Output:
	// 30 true
	// 0 false
}

func ExampleIntArray_Set() {
	s := garray.X创建整数并按大小(3, 5)
	s.X设置值(0, 10)
	s.X设置值(1, 20)
	s.X设置值(2, 15)
	s.X设置值(3, 30)
	fmt.Println(s.X取切片())

	// Output:
	// [10 20 15]
}

func ExampleIntArray_SetArray() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	fmt.Println(s.X取切片())

	// Output:
	// [10 20 15 30]
}

func ExampleIntArray_Replace() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	fmt.Println(s.X取切片())
	s.X替换(g.SliceInt别名{12, 13})
	fmt.Println(s.X取切片())

	// Output:
	// [10 20 15 30]
	// [12 13 15 30]
}

func ExampleIntArray_Sum() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	a := s.X求和()
	fmt.Println(a)

	// Output:
	// 75
}

func ExampleIntArray_Sort() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	a := s.X排序递增()
	fmt.Println(a)

	// Output:
	// [10,15,20,30]
}

func ExampleIntArray_SortFunc() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30})
	fmt.Println(s)
	s.X排序函数(func(v1, v2 int) bool {
		// fmt.Println(v1,v2)
		return v1 > v2
	})
	fmt.Println(s)
	s.X排序函数(func(v1, v2 int) bool {
		return v1 < v2
	})
	fmt.Println(s)

	// Output:
	// [10,20,15,30]
	// [30,20,15,10]
	// [10,15,20,30]
}

func ExampleIntArray_InsertBefore() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	s.X插入前面(1, 99)
	fmt.Println(s.X取切片())

	// Output:
	// [10 99 20 15 30]
}

func ExampleIntArray_InsertAfter() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	s.X插入后面(1, 99)
	fmt.Println(s.X取切片())

	// Output:
	// [10 20 99 15 30]
}

func ExampleIntArray_Remove() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	fmt.Println(s)
	s.X删除(1)
	fmt.Println(s.X取切片())

	// Output:
	// [10,20,15,30]
	// [10 15 30]
}

func ExampleIntArray_RemoveValue() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	fmt.Println(s)
	s.X删除值(20)
	fmt.Println(s.X取切片())

	// Output:
	// [10,20,15,30]
	// [10 15 30]
}

func ExampleIntArray_PushLeft() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	fmt.Println(s)
	s.X入栈左(96, 97, 98, 99)
	fmt.Println(s.X取切片())

	// Output:
	// [10,20,15,30]
	// [96 97 98 99 10 20 15 30]
}

func ExampleIntArray_PushRight() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	fmt.Println(s)
	s.X入栈右(96, 97, 98, 99)
	fmt.Println(s.X取切片())

	// Output:
	// [10,20,15,30]
	// [10 20 15 30 96 97 98 99]
}

func ExampleIntArray_PopLeft() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	fmt.Println(s)
	s.X出栈左()
	fmt.Println(s.X取切片())

	// Output:
	// [10,20,15,30]
	// [20 15 30]
}

func ExampleIntArray_PopRight() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30})
	fmt.Println(s)
	s.X出栈右()
	fmt.Println(s.X取切片())

	// Output:
	// [10,20,15,30]
	// [10 20 15]
}

func ExampleIntArray_PopRand() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60, 70})
	fmt.Println(s)
	r, _ := s.X出栈随机()
	fmt.Println(s)
	fmt.Println(r)

	// May Output:
	// [10,20,15,30,40,50,60,70]
	// [10,20,15,30,40,60,70]
	// 50
}

func ExampleIntArray_PopRands() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	r := s.X出栈随机多个(2)
	fmt.Println(s)
	fmt.Println(r)

	// May Output:
	// [10,20,15,30,40,50,60]
	// [10,20,15,30,40]
	// [50 60]
}

func ExampleIntArray_PopLefts() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	r := s.X出栈左多个(2)
	fmt.Println(s)
	fmt.Println(r)

	// Output:
	// [10,20,15,30,40,50,60]
	// [15,30,40,50,60]
	// [10 20]
}

func ExampleIntArray_PopRights() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	r := s.X出栈右多个(2)
	fmt.Println(s)
	fmt.Println(r)

	// Output:
	// [10,20,15,30,40,50,60]
	// [10,20,15,30,40]
	// [50 60]
}

func ExampleIntArray_Range() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	r := s.X取切片并按范围(2, 5)
	fmt.Println(r)

	// Output:
	// [10,20,15,30,40,50,60]
	// [15 30 40]
}

func ExampleIntArray_SubSlice() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	r := s.X取切片并按数量(3, 4)
	fmt.Println(r)

	// Output:
	// [10,20,15,30,40,50,60]
	// [30 40 50 60]
}

func ExampleIntArray_Append() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	s.Append别名(96, 97, 98)
	fmt.Println(s)

	// Output:
	// [10,20,15,30,40,50,60]
	// [10,20,15,30,40,50,60,96,97,98]
}

func ExampleIntArray_Len() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	fmt.Println(s.X取长度())

	// Output:
	// [10,20,15,30,40,50,60]
	// 7
}

func ExampleIntArray_Slice() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s.X取切片())

	// Output:
	// [10 20 15 30 40 50 60]
}

func ExampleIntArray_Interfaces() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	r := s.X取any切片()
	fmt.Println(r)

	// Output:
	// [10 20 15 30 40 50 60]
}

func ExampleIntArray_Clone() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	r := s.X取副本()
	fmt.Println(r)

	// Output:
	// [10,20,15,30,40,50,60]
	// [10,20,15,30,40,50,60]
}

func ExampleIntArray_Clear() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	fmt.Println(s.X清空())
	fmt.Println(s)

	// Output:
	// [10,20,15,30,40,50,60]
	// []
	// []
}

func ExampleIntArray_Contains() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s.X是否存在(20))
	fmt.Println(s.X是否存在(21))

	// Output:
	// true
	// false
}

func ExampleIntArray_Search() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s.X查找(20))
	fmt.Println(s.X查找(21))

	// Output:
	// 1
	// -1
}

func ExampleIntArray_Unique() {
	s := garray.X创建整数()
	s.X设置切片(g.SliceInt别名{10, 20, 15, 15, 20, 50, 60})
	fmt.Println(s)
	fmt.Println(s.X去重())

	// Output:
	// [10,20,15,15,20,50,60]
	// [10,20,15,50,60]
}

func ExampleIntArray_LockFunc() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	s.X遍历写锁定(func(array []int) {
		for i := 0; i < len(array)-1; i++ {
			fmt.Println(array[i])
		}
	})

	// Output:
	// 10
	// 20
	// 15
	// 30
	// 40
	// 50
}

func ExampleIntArray_RLockFunc() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	s.X遍历读锁定(func(array []int) {
		for i := 0; i < len(array); i++ {
			fmt.Println(array[i])
		}
	})

	// Output:
	// 10
	// 20
	// 15
	// 30
	// 40
	// 50
	// 60
}

func ExampleIntArray_Merge() {
	s1 := garray.X创建整数()
	s2 := garray.X创建整数()
	s1.X设置切片(g.SliceInt别名{10, 20, 15})
	s2.X设置切片(g.SliceInt别名{40, 50, 60})
	fmt.Println(s1)
	fmt.Println(s2)
	s1.X合并(s2)
	fmt.Println(s1)

	// Output:
	// [10,20,15]
	// [40,50,60]
	// [10,20,15,40,50,60]
}

func ExampleIntArray_Fill() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	s.X填充(2, 3, 99)
	fmt.Println(s)

	// Output:
	// [10,20,15,30,40,50,60]
	// [10,20,99,99,99,50,60]
}

func ExampleIntArray_Chunk() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	r := s.X分割(3)
	fmt.Println(r)

	// Output:
	// [10,20,15,30,40,50,60]
	// [[10 20 15] [30 40 50] [60]]
}

func ExampleIntArray_Pad() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	s.X填满(8, 99)
	fmt.Println(s)
	s.X填满(-10, 89)
	fmt.Println(s)

	// Output:
	// [10,20,15,30,40,50,60,99]
	// [89,89,10,20,15,30,40,50,60,99]
}

func ExampleIntArray_Rand() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	fmt.Println(s.X取值随机())

	// May Output:
	// [10,20,15,30,40,50,60]
	// 10 true
}

func ExampleIntArray_Rands() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	fmt.Println(s.X取值随机多个(3))

	// May Output:
	// [10,20,15,30,40,50,60]
	// [20 50 20]
}

func ExampleIntArray_Shuffle() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	fmt.Println(s.X随机排序())

	// May Output:
	// [10,20,15,30,40,50,60]
	// [10,40,15,50,20,60,30]
}

func ExampleIntArray_Reverse() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	fmt.Println(s.X倒排序())

	// Output:
	// [10,20,15,30,40,50,60]
	// [60,50,40,30,15,20,10]
}

func ExampleIntArray_Join() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	fmt.Println(s.X连接(","))

	// Output:
	// [10,20,15,30,40,50,60]
	// 10,20,15,30,40,50,60
}

func ExampleIntArray_CountValues() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 15, 40, 40, 40})
	fmt.Println(s.X统计())

	// Output:
	// map[10:1 15:2 20:1 40:3]
}

func ExampleIntArray_Iterator() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	s.X遍历(func(k int, v int) bool {
		fmt.Println(k, v)
		return true
	})

	// Output:
	// 0 10
	// 1 20
	// 2 15
	// 3 30
	// 4 40
	// 5 50
	// 6 60
}

func ExampleIntArray_IteratorAsc() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	s.X遍历升序(func(k int, v int) bool {
		fmt.Println(k, v)
		return true
	})

	// Output:
	// 0 10
	// 1 20
	// 2 15
	// 3 30
	// 4 40
	// 5 50
	// 6 60
}

func ExampleIntArray_IteratorDesc() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	s.X遍历降序(func(k int, v int) bool {
		fmt.Println(k, v)
		return true
	})

	// Output:
	// 6 60
	// 5 50
	// 4 40
	// 3 30
	// 2 15
	// 1 20
	// 0 10
}

func ExampleIntArray_String() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s)
	fmt.Println(s.String())

	// Output:
	// [10,20,15,30,40,50,60]
	// [10,20,15,30,40,50,60]
}

func ExampleIntArray_MarshalJSON() {
	type Student struct {
		Id     int
		Name   string
		Scores garray.IntArray
	}
	var array garray.IntArray
	array.X设置切片(g.SliceInt别名{98, 97, 96})
	s := Student{
		Id:     1,
		Name:   "john",
		Scores: array,
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))

	// Output:
	// {"Id":1,"Name":"john","Scores":[98,97,96]}
}

func ExampleIntArray_UnmarshalJSON() {
	b := []byte(`{"Id":1,"Name":"john","Scores":[98,96,97]}`)
	type Student struct {
		Id     int
		Name   string
		Scores *garray.IntArray
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)

	// Output:
	// {1 john [98,96,97]}
}

func ExampleIntArray_UnmarshalValue() {
	type Student struct {
		Name   string
		Scores *garray.IntArray
	}

	var s *Student
	gconv.Struct(g.Map{
		"name":   "john",
		"scores": g.SliceInt别名{96, 98, 97},
	}, &s)
	fmt.Println(s)

	// Output:
	// &{john [96,98,97]}
}

func ExampleIntArray_Filter() {
	array1 := garray.X创建整数并从切片(g.SliceInt别名{10, 40, 50, 0, 0, 0, 60})
	array2 := garray.X创建整数并从切片(g.SliceInt别名{10, 4, 51, 5, 45, 50, 56})
	fmt.Println(array1.X遍历删除(func(index int, value int) bool {
		return empty.IsEmpty(value)
	}))
	fmt.Println(array2.X遍历删除(func(index int, value int) bool {
		return value%2 == 0
	}))
	fmt.Println(array2.X遍历删除(func(index int, value int) bool {
		return value%2 == 1
	}))

	// Output:
	// [10,40,50,60]
	// [51,5,45]
	// []
}

func ExampleIntArray_FilterEmpty() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 40, 50, 0, 0, 0, 60})
	fmt.Println(s)
	fmt.Println(s.X删除所有零值())

	// Output:
	// [10,40,50,0,0,0,60]
	// [10,40,50,60]
}

func ExampleIntArray_IsEmpty() {
	s := garray.X创建整数并从切片(g.SliceInt别名{10, 20, 15, 30, 40, 50, 60})
	fmt.Println(s.X是否为空())
	s1 := garray.X创建整数()
	fmt.Println(s1.X是否为空())

	// Output:
	// false
	// true
}
