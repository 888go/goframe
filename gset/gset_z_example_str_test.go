// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package 集合类_test

import (
	"encoding/json"
	"fmt"

	"github.com/888go/goframe/gset"
	"github.com/gogf/gf/v2/frame/g"
)

// NewStrSet 创建并返回一个新的不包含重复项的集合。
// 参数`safe`用于指定是否在并发安全的情况下使用集合，默认为false。
// 这里，NewStrSet函数用于创建一个字符串集合，并确保其中的元素互不重复。该函数接受一个可选参数`safe`，它是一个布尔值，表示是否需要保证在并发环境下的安全性。如果不特别指定，那么默认情况下这个集合是不保证并发安全的。
func ExampleNewStrSet() {
	strSet := 集合类.X创建文本(true)
	strSet.X加入([]string{"str1", "str2", "str3"}...)
	fmt.Println(strSet.X取集合切片())

	// May Output:
	// [str3 str1 str2]
}

// NewStrSetFrom 从 `items` 中创建并返回一个新的集合。
func ExampleNewStrSetFrom() {
	strSet := 集合类.X创建文本并按值([]string{"str1", "str2", "str3"}, true)
	fmt.Println(strSet.X取集合切片())

	// May Output:
	// [str1 str2 str3]
}

// Add 向集合中添加一个或多个项目。
func ExampleStrSet_Add() {
	strSet := 集合类.X创建文本并按值([]string{"str1", "str2", "str3"}, true)
	strSet.X加入("str")
	fmt.Println(strSet.X取集合切片())
	fmt.Println(strSet.X加入值并跳过已存在("str"))

	// Mya Output:
	// [str str1 str2 str3]
	// false
}

// AddIfNotExist 检查项是否已存在于集合中，
// 如果该项不在集合中，则将其添加到集合中并返回 true；
// 否则，不进行任何操作并返回 false。
func ExampleStrSet_AddIfNotExist() {
	strSet := 集合类.X创建文本并按值([]string{"str1", "str2", "str3"}, true)
	strSet.X加入("str")
	fmt.Println(strSet.X取集合切片())
	fmt.Println(strSet.X加入值并跳过已存在("str"))

	// Mya Output:
	// [str str1 str2 str3]
	// false
}

// AddIfNotExistFunc 检查项是否存在集合中，
// 如果该项不存在于集合中，并且函数 `f` 返回 true，则将该项添加到集合中并返回 true，
// 否则不做任何操作并返回 false。
// 注意，函数 `f` 在无写入锁的情况下执行。
func ExampleStrSet_AddIfNotExistFunc() {
	strSet := 集合类.X创建文本并按值([]string{"str1", "str2", "str3"}, true)
	strSet.X加入("str")
	fmt.Println(strSet.X取集合切片())
	fmt.Println(strSet.X加入值并跳过已存在_函数("str5", func() bool {
		return true
	}))

	// May Output:
	// [str1 str2 str3 str]
	// true
}

// AddIfNotExistFunc 检查项是否存在集合中，
// 如果该项不存在于集合中，并且函数 `f` 返回 true，则将该项添加到集合中并返回 true，
// 否则不做任何操作并返回 false。
// 注意，函数 `f` 在无写入锁的情况下执行。
func ExampleStrSet_AddIfNotExistFuncLock() {
	strSet := 集合类.X创建文本并按值([]string{"str1", "str2", "str3"}, true)
	strSet.X加入("str")
	fmt.Println(strSet.X取集合切片())
	fmt.Println(strSet.X加入值并跳过已存在_并发安全函数("str4", func() bool {
		return true
	}))

	// May Output:
	// [str1 str2 str3 str]
	// true
}

// 清除删除集合中的所有项。
func ExampleStrSet_Clear() {
	strSet := 集合类.X创建文本并按值([]string{"str1", "str2", "str3"}, true)
	fmt.Println(strSet.X取数量())
	strSet.X清空()
	fmt.Println(strSet.X取数量())

	// Output:
	// 3
	// 0
}

// Complement 返回一个新的集合，它是从 `set` 到 `full` 的补集。
// 这意味着，新集合`newSet`中的所有元素都在`full`中但不在`set`中。
// 如果给定的集合 `full` 不是 `set` 的全集，则返回 `full` 和 `set` 之间的差集。
func ExampleStrSet_Complement() {
	strSet := 集合类.X创建文本并按值([]string{"str1", "str2", "str3", "str4", "str5"}, true)
	s := 集合类.X创建文本并按值([]string{"str1", "str2", "str3"}, true)
	fmt.Println(s.X取补集(strSet).X取集合切片())

	// May Output:
	// [str4 str5]
}

// Contains 检查集合中是否包含 `item`。
func ExampleStrSet_Contains() {
	var set 集合类.StrSet
	set.X加入("a")
	fmt.Println(set.X是否存在("a"))
	fmt.Println(set.X是否存在("A"))

	// Output:
	// true
	// false
}

// ContainsI 检查某个值是否以不区分大小写的方式存在于集合中。
// 注意：它内部会遍历整个集合，以不区分大小写的方式进行比较。
func ExampleStrSet_ContainsI() {
	var set 集合类.StrSet
	set.X加入("a")
	fmt.Println(set.X是否存在并忽略大小写("a"))
	fmt.Println(set.X是否存在并忽略大小写("A"))

	// Output:
	// true
	// true
}

// Diff 返回一个新的集合，这个集合是 `set` 与 `other` 的差集。
// 这意味着，新集合 `newSet` 中的所有元素都在 `set` 中，但不在 `other` 中。
func ExampleStrSet_Diff() {
	s1 := 集合类.X创建文本并按值([]string{"a", "b", "c"}, true)
	s2 := 集合类.X创建文本并按值([]string{"a", "b", "c", "d"}, true)
	fmt.Println(s2.X取差集(s1).X取集合切片())

	// Output:
	// [d]
}

// Equal 检查两个集合是否相等。
func ExampleStrSet_Equal() {
	s1 := 集合类.X创建文本并按值([]string{"a", "b", "c"}, true)
	s2 := 集合类.X创建文本并按值([]string{"a", "b", "c", "d"}, true)
	fmt.Println(s2.X是否相等(s1))

	s3 := 集合类.X创建文本并按值([]string{"a", "b", "c"}, true)
	s4 := 集合类.X创建文本并按值([]string{"a", "b", "c"}, true)
	fmt.Println(s3.X是否相等(s4))

	// Output:
	// false
	// true
}

// Intersect 返回一个新的集合，它是从 `set` 到 `other` 的交集。
// 这意味着，`newSet` 中的所有元素都在 `set` 中，并且也在 `other` 中。
func ExampleStrSet_Intersect() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c"}...)
	var s2 集合类.StrSet
	s2.X加入([]string{"a", "b", "c", "d"}...)
	fmt.Println(s2.X取交集(s1).X取集合切片())

	// May Output:
	// [c a b]
}

// IsSubsetOf 检查当前集合是否为 `other` 的子集
func ExampleStrSet_IsSubsetOf() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)
	var s2 集合类.StrSet
	s2.X加入([]string{"a", "b", "d"}...)
	fmt.Println(s2.X是否为子集(s1))

	// Output:
	// true
}

// Iterator 使用给定的回调函数`f`对集合进行只读遍历，
// 如果`f`返回true，则继续遍历；若返回false，则停止遍历。
func ExampleStrSet_Iterator() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)
	s1.X遍历(func(v string) bool {
		fmt.Println("Iterator", v)
		return true
	})

	// May Output:
	// Iterator a
	// Iterator b
	// Iterator c
	// Iterator d
}

// Join通过字符串`glue`连接items。
func ExampleStrSet_Join() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.X取集合文本(","))

	// May Output:
	// b,c,d,a
}

// LockFunc 使用回调函数`f`进行写入锁定。
func ExampleStrSet_LockFunc() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"1", "2"}...)
	s1.X写锁定_函数(func(m map[string]struct{}) {
		m["3"] = struct{}{}
	})
	fmt.Println(s1.X取集合切片())

	// 可能的输出
	// [2 3 1]

}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func ExampleStrSet_MarshalJSON() {
	type Student struct {
		Id     int
		Name   string
		Scores *集合类.StrSet
	}
	s := Student{
		Id:     1,
		Name:   "john",
		Scores: 集合类.X创建文本并按值([]string{"100", "99", "98"}, true),
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))

	// May Output:
	// {"Id":1,"Name":"john","Scores":["100","99","98"]}
}

// Merge 将 `others` 中的元素合并到 `set` 中。
func ExampleStrSet_Merge() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)

	s2 := 集合类.X创建文本(true)
	fmt.Println(s1.X合并(s2).X取集合切片())

	// May Output:
	// [d a b c]
}

// Pops 随机地从集合中弹出一个元素。
func ExampleStrSet_Pop() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)

	fmt.Println(s1.X出栈())

	// May Output:
	// a
}

// Pops 随机地从集合中弹出 `size` 个元素。
// 如果 size == -1，则返回所有元素。
func ExampleStrSet_Pops() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)
	for _, v := range s1.X出栈多个(2) {
		fmt.Println(v)
	}

	// May Output:
	// a
	// b
}

// RLockFunc 通过回调函数 `f` 对读取进行加锁。
func ExampleStrSet_RLockFunc() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)
	s1.X读锁定_函数(func(m map[string]struct{}) {
		fmt.Println(m)
	})

	// Output:
	// map[a:{} b:{} c:{} d:{}]
}

// Remove 从集合中删除`item`。
func ExampleStrSet_Remove() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)
	s1.X删除("a")
	fmt.Println(s1.X取集合切片())

	// May Output:
	// [b c d]
}

// Size 返回集合的大小。
func ExampleStrSet_Size() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.X取数量())

	// Output:
	// 4
}

// Slice 返回集合中项目的切片形式。
func ExampleStrSet_Slice() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.X取集合切片())

	// May Output:
	// [a,b,c,d]
}

// String 返回 items 作为字符串，其实现方式类似于 json.Marshal。
func ExampleStrSet_String() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.String())

	// May Output:
	// "a","b","c","d"
}

// Sum 计算项目总和。注意：项目应转换为 int 类型，
// 否则你可能会得到一个意想不到的结果。
func ExampleStrSet_Sum() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"1", "2", "3", "4"}...)
	fmt.Println(s1.X求和())

	// Output:
	// 10
}

// Union 返回一个新的集合，该集合是 `set` 和 `other` 的并集。
// 这意味着，`newSet` 中的所有元素都在 `set` 或者 `other` 中。
func ExampleStrSet_Union() {
	s1 := 集合类.X创建文本(true)
	s1.X加入([]string{"a", "b", "c", "d"}...)
	s2 := 集合类.X创建文本(true)
	s2.X加入([]string{"a", "b", "d"}...)
	fmt.Println(s1.X取并集(s2).X取集合切片())

	// May Output:
	// [a b c d]
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func ExampleStrSet_UnmarshalJSON() {
	b := []byte(`{"Id":1,"Name":"john","Scores":["100","99","98"]}`)
	type Student struct {
		Id     int
		Name   string
		Scores *集合类.StrSet
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)

	// May Output:
	// {1 john "99","98","100"}
}

// UnmarshalValue 是一个接口实现，用于为 set 设置任意类型的值。
func ExampleStrSet_UnmarshalValue() {
	b := []byte(`{"Id":1,"Name":"john","Scores":["100","99","98"]}`)
	type Student struct {
		Id     int
		Name   string
		Scores *集合类.StrSet
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)

	// May Output:
	// {1 john "99","98","100"}
}

// Walk 对集合中的每一个元素应用用户提供的函数 `f`。
func ExampleStrSet_Walk() {
	var (
		set    集合类.StrSet
		names  = g.SliceStr{"user", "user_detail"}
		prefix = "gf_"
	)
	set.X加入(names...)
	// 为给定的表名添加前缀
	set.X遍历修改(func(item string) string {
		return prefix + item
	})
	fmt.Println(set.X取集合切片())

	// May Output:
	// [gf_user gf_user_detail]
}
