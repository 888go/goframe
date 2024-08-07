// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 切片类_test

import (
	"fmt"
	"strings"

	"github.com/888go/goframe/internal/empty"

	garray "github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

func ExampleStrArray_Walk() {
	var array garray.StrArray
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

func ExampleNewStrArray() {
	s := garray.X创建文本()
	s.Append别名("We")
	s.Append别名("are")
	s.Append别名("GF")
	s.Append别名("fans")
	fmt.Println(s.X取切片())

	// Output:
	// [We are GF fans]
}

func ExampleNewStrArraySize() {
	s := garray.X创建文本并按大小(3, 5)
	s.X设置值(0, "We")
	s.X设置值(1, "are")
	s.X设置值(2, "GF")
	s.X设置值(3, "fans")
	fmt.Println(s.X取切片(), s.X取长度(), cap(s.X取切片()))

	// Output:
	// [We are GF] 3 5
}

func ExampleNewStrArrayFrom() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"We", "are", "GF", "fans", "!"})
	fmt.Println(s.X取切片(), s.X取长度(), cap(s.X取切片()))

	// Output:
	// [We are GF fans !] 5 5
}

func ExampleStrArray_At() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"We", "are", "GF", "fans", "!"})
	sAt := s.X取值(2)
	fmt.Println(sAt)

	// Output:
	// GF
}

func ExampleStrArray_Get() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"We", "are", "GF", "fans", "!"})
	sGet, sBool := s.X取值2(3)
	fmt.Println(sGet, sBool)

	// Output:
	// fans true
}

func ExampleStrArray_Set() {
	s := garray.X创建文本并按大小(3, 5)
	s.X设置值(0, "We")
	s.X设置值(1, "are")
	s.X设置值(2, "GF")
	s.X设置值(3, "fans")
	fmt.Println(s.X取切片())

	// Output:
	// [We are GF]
}

func ExampleStrArray_SetArray() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"We", "are", "GF", "fans", "!"})
	fmt.Println(s.X取切片())

	// Output:
	// [We are GF fans !]
}

func ExampleStrArray_Replace() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"We", "are", "GF", "fans", "!"})
	fmt.Println(s.X取切片())
	s.X替换(g.SliceStr别名{"Happy", "coding"})
	fmt.Println(s.X取切片())

	// Output:
	// [We are GF fans !]
	// [Happy coding GF fans !]
}

func ExampleStrArray_Sum() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"3", "5", "10"})
	a := s.X求和()
	fmt.Println(a)

	// Output:
	// 18
}

func ExampleStrArray_Sort() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"b", "d", "a", "c"})
	a := s.X排序递增()
	fmt.Println(a)

	// Output:
	// ["a","b","c","d"]
}

func ExampleStrArray_SortFunc() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"b", "c", "a"})
	fmt.Println(s)
	s.X排序函数(func(v1, v2 string) bool {
		return gstr.X顺序比较(v1, v2) > 0
	})
	fmt.Println(s)
	s.X排序函数(func(v1, v2 string) bool {
		return gstr.X顺序比较(v1, v2) < 0
	})
	fmt.Println(s)

	// Output:
	// ["b","c","a"]
	// ["c","b","a"]
	// ["a","b","c"]
}

func ExampleStrArray_InsertBefore() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d"})
	s.X插入前面(1, "here")
	fmt.Println(s.X取切片())

	// Output:
	// [a here b c d]
}

func ExampleStrArray_InsertAfter() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d"})
	s.X插入后面(1, "here")
	fmt.Println(s.X取切片())

	// Output:
	// [a b here c d]
}

func ExampleStrArray_Remove() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d"})
	s.X删除(1)
	fmt.Println(s.X取切片())

	// Output:
	// [a c d]
}

func ExampleStrArray_RemoveValue() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d"})
	s.X删除值("b")
	fmt.Println(s.X取切片())

	// Output:
	// [a c d]
}

func ExampleStrArray_PushLeft() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d"})
	s.X入栈左("We", "are", "GF", "fans")
	fmt.Println(s.X取切片())

	// Output:
	// [We are GF fans a b c d]
}

func ExampleStrArray_PushRight() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d"})
	s.X入栈右("We", "are", "GF", "fans")
	fmt.Println(s.X取切片())

	// Output:
	// [a b c d We are GF fans]
}

func ExampleStrArray_PopLeft() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d"})
	s.X出栈左()
	fmt.Println(s.X取切片())

	// Output:
	// [b c d]
}

func ExampleStrArray_PopRight() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d"})
	s.X出栈右()
	fmt.Println(s.X取切片())

	// Output:
	// [a b c]
}

func ExampleStrArray_PopRand() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	r, _ := s.X出栈随机()
	fmt.Println(r)

	// May Output:
	// e
}

func ExampleStrArray_PopRands() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	r := s.X出栈随机多个(2)
	fmt.Println(r)

	// May Output:
	// [e c]
}

func ExampleStrArray_PopLefts() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	r := s.X出栈左多个(2)
	fmt.Println(r)
	fmt.Println(s)

	// Output:
	// [a b]
	// ["c","d","e","f","g","h"]
}

func ExampleStrArray_PopRights() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	r := s.X出栈右多个(2)
	fmt.Println(r)
	fmt.Println(s)

	// Output:
	// [g h]
	// ["a","b","c","d","e","f"]
}

func ExampleStrArray_Range() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	r := s.X取切片并按范围(2, 5)
	fmt.Println(r)

	// Output:
	// [c d e]
}

func ExampleStrArray_SubSlice() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	r := s.X取切片并按数量(3, 4)
	fmt.Println(r)

	// Output:
	// [d e f g]
}

func ExampleStrArray_Append() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"We", "are", "GF", "fans"})
	s.Append别名("a", "b", "c")
	fmt.Println(s)

	// Output:
	// ["We","are","GF","fans","a","b","c"]
}

func ExampleStrArray_Len() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	fmt.Println(s.X取长度())

	// Output:
	// 8
}

func ExampleStrArray_Slice() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	fmt.Println(s.X取切片())

	// Output:
	// [a b c d e f g h]
}

func ExampleStrArray_Interfaces() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	r := s.X取any切片()
	fmt.Println(r)

	// Output:
	// [a b c d e f g h]
}

func ExampleStrArray_Clone() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	r := s.X取副本()
	fmt.Println(r)
	fmt.Println(s)

	// Output:
	// ["a","b","c","d","e","f","g","h"]
	// ["a","b","c","d","e","f","g","h"]
}

func ExampleStrArray_Clear() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	fmt.Println(s)
	fmt.Println(s.X清空())
	fmt.Println(s)

	// Output:
	// ["a","b","c","d","e","f","g","h"]
	// []
	// []
}

func ExampleStrArray_Contains() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	fmt.Println(s.X是否存在("e"))
	fmt.Println(s.X是否存在("z"))

	// Output:
	// true
	// false
}

func ExampleStrArray_ContainsI() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	fmt.Println(s.X是否存在并忽略大小写("E"))
	fmt.Println(s.X是否存在并忽略大小写("z"))

	// Output:
	// true
	// false
}

func ExampleStrArray_Search() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	fmt.Println(s.X查找("e"))
	fmt.Println(s.X查找("z"))

	// Output:
	// 4
	// -1
}

func ExampleStrArray_Unique() {
	s := garray.X创建文本()
	s.X设置切片(g.SliceStr别名{"a", "b", "c", "c", "c", "d", "d"})
	fmt.Println(s.X去重())

	// Output:
	// ["a","b","c","d"]
}

func ExampleStrArray_LockFunc() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c"})
	s.X遍历写锁定(func(array []string) {
		array[len(array)-1] = "GF fans"
	})
	fmt.Println(s)

	// Output:
	// ["a","b","GF fans"]
}

func ExampleStrArray_RLockFunc() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c", "d", "e"})
	s.X遍历读锁定(func(array []string) {
		for i := 0; i < len(array); i++ {
			fmt.Println(array[i])
		}
	})

	// Output:
	// a
	// b
	// c
	// d
	// e
}

func ExampleStrArray_Merge() {
	s1 := garray.X创建文本()
	s2 := garray.X创建文本()
	s1.X设置切片(g.SliceStr别名{"a", "b", "c"})
	s2.X设置切片(g.SliceStr别名{"d", "e", "f"})
	s1.X合并(s2)
	fmt.Println(s1)

	// Output:
	// ["a","b","c","d","e","f"]
}

func ExampleStrArray_Fill() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	s.X填充(2, 3, "here")
	fmt.Println(s)

	// Output:
	// ["a","b","here","here","here","f","g","h"]
}

func ExampleStrArray_Chunk() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	r := s.X分割(3)
	fmt.Println(r)

	// Output:
	// [[a b c] [d e f] [g h]]
}

func ExampleStrArray_Pad() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c"})
	s.X填满(7, "here")
	fmt.Println(s)
	s.X填满(-10, "there")
	fmt.Println(s)

	// Output:
	// ["a","b","c","here","here","here","here"]
	// ["there","there","there","a","b","c","here","here","here","here"]
}

func ExampleStrArray_Rand() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	fmt.Println(s.X取值随机())

	// May Output:
	// c true
}

func ExampleStrArray_Rands() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	fmt.Println(s.X取值随机多个(3))

	// May Output:
	// [e h e]
}

func ExampleStrArray_Shuffle() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	fmt.Println(s.X随机排序())

	// May Output:
	// ["a","c","e","d","b","g","f","h"]
}

func ExampleStrArray_Reverse() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c", "d", "e", "f", "g", "h"})
	fmt.Println(s.X倒排序())

	// Output:
	// ["h","g","f","e","d","c","b","a"]
}

func ExampleStrArray_Join() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c"})
	fmt.Println(s.X连接(","))

	// Output:
	// a,b,c
}

func ExampleStrArray_CountValues() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c", "c", "c", "d", "d"})
	fmt.Println(s.X统计())

	// Output:
	// map[a:1 b:1 c:3 d:2]
}

func ExampleStrArray_Iterator() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c"})
	s.X遍历(func(k int, v string) bool {
		fmt.Println(k, v)
		return true
	})

	// Output:
	// 0 a
	// 1 b
	// 2 c
}

func ExampleStrArray_IteratorAsc() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c"})
	s.X遍历升序(func(k int, v string) bool {
		fmt.Println(k, v)
		return true
	})

	// Output:
	// 0 a
	// 1 b
	// 2 c
}

func ExampleStrArray_IteratorDesc() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c"})
	s.X遍历降序(func(k int, v string) bool {
		fmt.Println(k, v)
		return true
	})

	// Output:
	// 2 c
	// 1 b
	// 0 a
}

func ExampleStrArray_String() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "c"})
	fmt.Println(s.String())

	// Output:
	// ["a","b","c"]
}

func ExampleStrArray_MarshalJSON() {
	type Student struct {
		Id      int
		Name    string
		Lessons []string
	}
	s := Student{
		Id:      1,
		Name:    "john",
		Lessons: []string{"Math", "English", "Music"},
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))

	// Output:
	// {"Id":1,"Name":"john","Lessons":["Math","English","Music"]}
}

func ExampleStrArray_UnmarshalJSON() {
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

func ExampleStrArray_UnmarshalValue() {
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

func ExampleStrArray_Filter() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"Math", "English", "Sport"})
	s1 := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "", "c", "", "", "d"})
	fmt.Println(s1.X遍历删除(func(index int, value string) bool {
		return empty.IsEmpty(value)
	}))

	fmt.Println(s.X遍历删除(func(index int, value string) bool {
		return strings.Contains(value, "h")
	}))

	// Output:
	// ["a","b","c","d"]
	// ["Sport"]
}

func ExampleStrArray_FilterEmpty() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "", "c", "", "", "d"})
	fmt.Println(s.X删除所有空值())

	// Output:
	// ["a","b","c","d"]
}

func ExampleStrArray_IsEmpty() {
	s := garray.X创建文本并从切片(g.SliceStr别名{"a", "b", "", "c", "", "", "d"})
	fmt.Println(s.X是否为空())
	s1 := garray.X创建文本()
	fmt.Println(s1.X是否为空())

	// Output:
	// false
	// true
}
