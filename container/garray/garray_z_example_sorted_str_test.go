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

func ExampleSortedStrArray_Walk() {
	var array garray.SortedStrArray
	tables := g.SliceStr别名{"user", "user_detail"}
	prefix := "gf_"
	array.Append别名(tables...)
		// 为给定的表名添加前缀。 md5:dea7405f272e0c9e
	array.X遍历修改(func(value string) string {
		return prefix + value
	})
	fmt.Println(array.X取切片())

	// Output:
	// [gf_user gf_user_detail]
}

func ExampleNewSortedStrArray() {
	s := garray.X创建文本排序()
	s.Append别名("b")
	s.Append别名("d")
	s.Append别名("c")
	s.Append别名("a")
	fmt.Println(s.X取切片())

	// Output:
	// [a b c d]
}

func ExampleNewSortedStrArraySize() {
	s := garray.X创建文本排序并按大小(3)
	s.X设置切片([]string{"b", "d", "a", "c"})
	fmt.Println(s.X取切片(), s.X取长度(), cap(s.X取切片()))

	// Output:
	// [a b c d] 4 4
}

func ExampleNewStrArrayFromCopy() {
	s := garray.X创建文本排序并从切片复制(g.SliceStr别名{"b", "d", "c", "a"})
	fmt.Println(s.X取切片())

	// Output:
	// [a b c d]
}

func ExampleSortedStrArray_At() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "d", "c", "a"})
	sAt := s.X取值(2)
	fmt.Println(s)
	fmt.Println(sAt)

	// Output:
	// ["a","b","c","d"]
	// c

}

func ExampleSortedStrArray_Get() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "d", "c", "a", "e"})
	sGet, sBool := s.X取值2(3)
	fmt.Println(s)
	fmt.Println(sGet, sBool)

	// Output:
	// ["a","b","c","d","e"]
	// d true
}

func ExampleSortedStrArray_SetArray() {
	s := garray.X创建文本排序()
	s.X设置切片([]string{"b", "d", "a", "c"})
	fmt.Println(s.X取切片())

	// Output:
	// [a b c d]
}

func ExampleSortedStrArray_SetUnique() {
	s := garray.X创建文本排序()
	s.X设置切片([]string{"b", "d", "a", "c", "c", "a"})
	fmt.Println(s.X设置去重(true))

	// Output:
	// ["a","b","c","d"]
}

func ExampleSortedStrArray_Sum() {
	s := garray.X创建文本排序()
	s.X设置切片([]string{"5", "3", "2"})
	fmt.Println(s)
	a := s.X求和()
	fmt.Println(a)

	// Output:
	// ["2","3","5"]
	// 10
}

func ExampleSortedStrArray_Sort() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"b", "d", "a", "c"})
	fmt.Println(s)
	a := s.X排序递增()
	fmt.Println(a)

	// Output:
	// ["a","b","c","d"]
	// ["a","b","c","d"]
}

func ExampleSortedStrArray_Remove() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"b", "d", "c", "a"})
	fmt.Println(s.X取切片())
	s.X删除(1)
	fmt.Println(s.X取切片())

	// Output:
	// [a b c d]
	// [a c d]
}

func ExampleSortedStrArray_RemoveValue() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"b", "d", "c", "a"})
	fmt.Println(s.X取切片())
	s.X删除值("b")
	fmt.Println(s.X取切片())

	// Output:
	// [a b c d]
	// [a c d]
}

func ExampleSortedStrArray_PopLeft() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"b", "d", "c", "a"})
	r, _ := s.X出栈左()
	fmt.Println(r)
	fmt.Println(s.X取切片())

	// Output:
	// a
	// [b c d]
}

func ExampleSortedStrArray_PopRight() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"b", "d", "c", "a"})
	fmt.Println(s.X取切片())
	r, _ := s.X出栈右()
	fmt.Println(r)
	fmt.Println(s.X取切片())

	// Output:
	// [a b c d]
	// d
	// [a b c]
}

func ExampleSortedStrArray_PopRights() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	r := s.X出栈右多个(2)
	fmt.Println(r)
	fmt.Println(s)

	// Output:
	// [g h]
	// ["a","b","c","d","e","f"]
}

func ExampleSortedStrArray_Rand() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	r, _ := s.X出栈随机()
	fmt.Println(r)
	fmt.Println(s)

	// May Output:
	// b
	// ["a","c","d","e","f","g","h"]
}

func ExampleSortedStrArray_PopRands() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	r := s.X出栈随机多个(2)
	fmt.Println(r)
	fmt.Println(s)

	// May Output:
	// [d a]
	// ["b","c","e","f","g","h"]
}

func ExampleSortedStrArray_PopLefts() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	r := s.X出栈左多个(2)
	fmt.Println(r)
	fmt.Println(s)

	// Output:
	// [a b]
	// ["c","d","e","f","g","h"]
}

func ExampleSortedStrArray_Range() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	r := s.X取切片并按范围(2, 5)
	fmt.Println(r)

	// Output:
	// [c d e]
}

func ExampleSortedStrArray_SubSlice() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	r := s.X取切片并按数量(3, 4)
	fmt.Println(s.X取切片())
	fmt.Println(r)

	// Output:
	// [a b c d e f g h]
	// [d e f g]
}

func ExampleSortedStrArray_Add() {
	s := garray.X创建文本排序()
	s.X入栈右("b", "d", "c", "a")
	fmt.Println(s)

	// Output:
	// ["a","b","c","d"]
}

func ExampleSortedStrArray_Append() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"b", "d", "c", "a"})
	fmt.Println(s)
	s.Append别名("f", "e", "g")
	fmt.Println(s)

	// Output:
	// ["a","b","c","d"]
	// ["a","b","c","d","e","f","g"]
}

func ExampleSortedStrArray_Len() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	fmt.Println(s)
	fmt.Println(s.X取长度())

	// Output:
	// ["a","b","c","d","e","f","g","h"]
	// 8
}

func ExampleSortedStrArray_Slice() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	fmt.Println(s.X取切片())

	// Output:
	// [a b c d e f g h]
}

func ExampleSortedStrArray_Interfaces() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	r := s.X取any切片()
	fmt.Println(r)

	// Output:
	// [a b c d e f g h]
}

func ExampleSortedStrArray_Clone() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	r := s.X取副本()
	fmt.Println(r)
	fmt.Println(s)

	// Output:
	// ["a","b","c","d","e","f","g","h"]
	// ["a","b","c","d","e","f","g","h"]
}

func ExampleSortedStrArray_Clear() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	fmt.Println(s)
	fmt.Println(s.X清空())
	fmt.Println(s)

	// Output:
	// ["a","b","c","d","e","f","g","h"]
	// []
	// []
}

func ExampleSortedStrArray_Contains() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	fmt.Println(s.X是否存在("e"))
	fmt.Println(s.X是否存在("E"))
	fmt.Println(s.X是否存在("z"))

	// Output:
	// true
	// false
	// false
}

func ExampleSortedStrArray_ContainsI() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	fmt.Println(s)
	fmt.Println(s.X是否存在并忽略大小写("E"))
	fmt.Println(s.X是否存在并忽略大小写("z"))

	// Output:
	// ["a","b","c","d","e","f","g","h"]
	// true
	// false
}

func ExampleSortedStrArray_Search() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	fmt.Println(s)
	fmt.Println(s.X查找("e"))
	fmt.Println(s.X查找("E"))
	fmt.Println(s.X查找("z"))

	// Output:
	// ["a","b","c","d","e","f","g","h"]
	// 4
	// -1
	// -1
}

func ExampleSortedStrArray_Unique() {
	s := garray.X创建文本排序()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "c", "c", "d", "d"})
	fmt.Println(s)
	fmt.Println(s.X去重())

	// Output:
	// ["a","b","c","c","c","d","d"]
	// ["a","b","c","d"]
}

func ExampleSortedStrArray_LockFunc() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "c", "a"})
	s.X遍历写锁定(func(array []string) {
		array[len(array)-1] = "GF fans"
	})
	fmt.Println(s)

	// Output:
	// ["a","b","GF fans"]
}

func ExampleSortedStrArray_RLockFunc() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "c", "a"})
	s.X遍历读锁定(func(array []string) {
		array[len(array)-1] = "GF fans"
		fmt.Println(array[len(array)-1])
	})
	fmt.Println(s)

	// Output:
	// GF fans
	// ["a","b","GF fans"]
}

func ExampleSortedStrArray_Merge() {
	s1 := garray.X创建文本排序()
	s2 := garray.X创建文本排序()
	s1.X设置切片(g.SliceStr别名{"b", "c", "a"})
	s2.X设置切片(g.SliceStr别名{"e", "d", "f"})
	fmt.Println(s1)
	fmt.Println(s2)
	s1.X合并(s2)
	fmt.Println(s1)

	// Output:
	// ["a","b","c"]
	// ["d","e","f"]
	// ["a","b","c","d","e","f"]
}

func ExampleSortedStrArray_Chunk() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	r := s.X分割(3)
	fmt.Println(r)

	// Output:
	// [[a b c] [d e f] [g h]]
}

func ExampleSortedStrArray_Rands() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	fmt.Println(s)
	fmt.Println(s.X取值随机多个(3))

	// May Output:
	// ["a","b","c","d","e","f","g","h"]
	// [h g c]
}

func ExampleSortedStrArray_Join() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"c", "b", "a", "d", "f", "e", "h", "g"})
	fmt.Println(s.X连接(","))

	// Output:
	// a,b,c,d,e,f,g,h
}

func ExampleSortedStrArray_CountValues() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"a", "b", "c", "c", "c", "d", "d"})
	fmt.Println(s.X统计())

	// Output:
	// map[a:1 b:1 c:3 d:2]
}

func ExampleSortedStrArray_Iterator() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "c", "a"})
	s.X遍历(func(k int, v string) bool {
		fmt.Println(k, v)
		return true
	})

	// Output:
	// 0 a
	// 1 b
	// 2 c
}

func ExampleSortedStrArray_IteratorAsc() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "c", "a"})
	s.X遍历升序(func(k int, v string) bool {
		fmt.Println(k, v)
		return true
	})

	// Output:
	// 0 a
	// 1 b
	// 2 c
}

func ExampleSortedStrArray_IteratorDesc() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "c", "a"})
	s.X遍历降序(func(k int, v string) bool {
		fmt.Println(k, v)
		return true
	})

	// Output:
	// 2 c
	// 1 b
	// 0 a
}

func ExampleSortedStrArray_String() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "c", "a"})
	fmt.Println(s.String())

	// Output:
	// ["a","b","c"]
}

func ExampleSortedStrArray_MarshalJSON() {
	type Student struct {
		ID     int
		Name   string
		Levels garray.SortedStrArray
	}
	r := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "c", "a"})
	s := Student{
		ID:     1,
		Name:   "john",
		Levels: *r,
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))

	// Output:
	// {"ID":1,"Name":"john","Levels":["a","b","c"]}
}

func ExampleSortedStrArray_UnmarshalJSON() {
	b := []byte(`{"Id":1,"Name":"john","Lessons":["Math","English","Sport"]}`)
	type Student struct {
		Id      int
		Name    string
		Lessons *garray.StrArray
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)

	// Output:
	// {1 john ["Math","English","Sport"]}
}

func ExampleSortedStrArray_UnmarshalValue() {
	type Student struct {
		Name    string
		Lessons *garray.StrArray
	}
	var s *Student
	gconv.Struct(g.Map{
		"name":    "john",
		"lessons": []byte(`["Math","English","Sport"]`),
	}, &s)
	fmt.Println(s)

	var s1 *Student
	gconv.Struct(g.Map{
		"name":    "john",
		"lessons": g.SliceStr别名{"Math", "English", "Sport"},
	}, &s1)
	fmt.Println(s1)

	// Output:
	// &{john ["Math","English","Sport"]}
	// &{john ["Math","English","Sport"]}
}

func ExampleSortedStrArray_Filter() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "a", "", "c", "", "", "d"})
	fmt.Println(s)
	fmt.Println(s.X遍历删除(func(index int, value string) bool {
		return empty.IsEmpty(value)
	}))

	// Output:
	// ["","","","a","b","c","d"]
	// ["a","b","c","d"]
}

func ExampleSortedStrArray_FilterEmpty() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "a", "", "c", "", "", "d"})
	fmt.Println(s)
	fmt.Println(s.X删除所有空值())

	// Output:
	// ["","","","a","b","c","d"]
	// ["a","b","c","d"]
}

func ExampleSortedStrArray_IsEmpty() {
	s := garray.X创建文本排序并从切片(g.SliceStr别名{"b", "a", "", "c", "", "", "d"})
	fmt.Println(s.X是否为空())
	s1 := garray.X创建文本排序()
	fmt.Println(s1.X是否为空())

	// Output:
	// false
	// true
}
