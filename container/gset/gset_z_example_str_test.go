// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。 md5:1d281c30cdc3423b

package gset_test

import (
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
)

// NewStrSet 创建并返回一个新集合，其中包含不重复的元素。
// 参数 `safe` 用于指定是否在并发安全环境下使用集合，默认为 false。 md5:b4b32102d4f1da78
func ExampleNewStrSet() {
	strSet := gset.NewStrSet(true)
	strSet.Add([]string{"str1", "str2", "str3"}...)
	fmt.Println(strSet.Slice())

	// May Output:
	// [str3 str1 str2]
}

// NewStrSetFrom 从`items`创建一个新的集合。 md5:6f9a406a984403d2
func ExampleNewStrSetFrom() {
	strSet := gset.NewStrSetFrom([]string{"str1", "str2", "str3"}, true)
	fmt.Println(strSet.Slice())

	// May Output:
	// [str1 str2 str3]
}

// Add 将一个或多个项目添加到集合中。 md5:316141ff7d4b8e45
func ExampleStrSet_Add() {
	strSet := gset.NewStrSetFrom([]string{"str1", "str2", "str3"}, true)
	strSet.Add("str")
	fmt.Println(strSet.Slice())
	fmt.Println(strSet.AddIfNotExist("str"))

	// May Output:
	// [str str1 str2 str3]
	// false
}

// AddIfNotExist 检查项是否存在于集合中，
// 如果项不存在于集合中，它会将项添加到集合中并返回 true，否则什么都不做并返回 false。 md5:9cff508c42cffd55
func ExampleStrSet_AddIfNotExist() {
	strSet := gset.NewStrSetFrom([]string{"str1", "str2", "str3"}, true)
	strSet.Add("str")
	fmt.Println(strSet.Slice())
	fmt.Println(strSet.AddIfNotExist("str"))

	// May Output:
	// [str str1 str2 str3]
	// false
}

// AddIfNotExistFunc 检查项目是否存在于集合中，
// 如果项目不在集合中且函数 `f` 返回 true，则将其添加到集合并返回 true，
// 否则不做任何操作并返回 false。
// 请注意，函数 `f` 在无写锁的情况下执行。 md5:0a51b9d79022ae82
func ExampleStrSet_AddIfNotExistFunc() {
	strSet := gset.NewStrSetFrom([]string{"str1", "str2", "str3"}, true)
	strSet.Add("str")
	fmt.Println(strSet.Slice())
	fmt.Println(strSet.AddIfNotExistFunc("str5", func() bool {
		return true
	}))

	// May Output:
	// [str1 str2 str3 str]
	// true
}

// AddIfNotExistFunc 检查项目是否存在于集合中，
// 如果项目不在集合中且函数 `f` 返回 true，则将其添加到集合并返回 true，
// 否则不做任何操作并返回 false。
// 请注意，函数 `f` 在无写锁的情况下执行。 md5:0a51b9d79022ae82
func ExampleStrSet_AddIfNotExistFuncLock() {
	strSet := gset.NewStrSetFrom([]string{"str1", "str2", "str3"}, true)
	strSet.Add("str")
	fmt.Println(strSet.Slice())
	fmt.Println(strSet.AddIfNotExistFuncLock("str4", func() bool {
		return true
	}))

	// May Output:
	// [str1 str2 str3 str]
	// true
}

// Clear 删除集合中的所有项。 md5:ce349f0cd3114465
func ExampleStrSet_Clear() {
	strSet := gset.NewStrSetFrom([]string{"str1", "str2", "str3"}, true)
	fmt.Println(strSet.Size())
	strSet.Clear()
	fmt.Println(strSet.Size())

	// Output:
	// 3
	// 0
}

// Complement 返回一个新的集合，该集合是相对于`set`到`full`的补集。
// 这意味着，`newSet`中的所有项都包含在`full`中但不包含在`set`中。
// 如果给定的集合`full`并不是`set`的全集，则它返回`full`与`set`之间的差异。 md5:2116fbb7587db792
func ExampleStrSet_Complement() {
	strSet := gset.NewStrSetFrom([]string{"str1", "str2", "str3", "str4", "str5"}, true)
	s := gset.NewStrSetFrom([]string{"str1", "str2", "str3"}, true)
	fmt.Println(s.Complement(strSet).Slice())

	// May Output:
	// [str4 str5]
}

// Contains 检查集合是否包含 `item`。 md5:20a3bdc6aeef1d67
func ExampleStrSet_Contains() {
	var set gset.StrSet
	set.Add("a")
	fmt.Println(set.Contains("a"))
	fmt.Println(set.Contains("A"))

	// Output:
	// true
	// false
}

// ContainsI 检查集合中是否存在某个值（忽略大小写）。
// 注意，它内部会遍历整个集合以进行不区分大小写的比较。 md5:851e1bbfa6da1bae
func ExampleStrSet_ContainsI() {
	var set gset.StrSet
	set.Add("a")
	fmt.Println(set.ContainsI("a"))
	fmt.Println(set.ContainsI("A"))

	// Output:
	// true
	// true
}

// Diff 返回一个新的集合，它是 `set` 与 `other` 之间的差集。
// 这意味着，`newSet` 中的所有项目都在 `set` 中，但不在 `other` 中。 md5:6779e6e007651b53
func ExampleStrSet_Diff() {
	s1 := gset.NewStrSetFrom([]string{"a", "b", "c"}, true)
	s2 := gset.NewStrSetFrom([]string{"a", "b", "c", "d"}, true)
	fmt.Println(s2.Diff(s1).Slice())

	// Output:
	// [d]
}

// Equal 检查两个集合是否相等。 md5:105ea4dd39b57fe8
func ExampleStrSet_Equal() {
	s1 := gset.NewStrSetFrom([]string{"a", "b", "c"}, true)
	s2 := gset.NewStrSetFrom([]string{"a", "b", "c", "d"}, true)
	fmt.Println(s2.Equal(s1))

	s3 := gset.NewStrSetFrom([]string{"a", "b", "c"}, true)
	s4 := gset.NewStrSetFrom([]string{"a", "b", "c"}, true)
	fmt.Println(s3.Equal(s4))

	// Output:
	// false
	// true
}

// Intersect 返回一个新的集合，这个集合是 `set` 和 `other` 的交集。
// 这意味着，`newSet` 中的所有元素都既存在于 `set` 中也存在于 `other` 中。 md5:327d3fcc12f06583
func ExampleStrSet_Intersect() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c"}...)
	var s2 gset.StrSet
	s2.Add([]string{"a", "b", "c", "d"}...)
	fmt.Println(s2.Intersect(s1).Slice())

	// May Output:
	// [c a b]
}

// IsSubsetOf 检查当前集合是否是 `other` 的子集. md5:70b7ed1e77ec2f80
func ExampleStrSet_IsSubsetOf() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	var s2 gset.StrSet
	s2.Add([]string{"a", "b", "d"}...)
	fmt.Println(s2.IsSubsetOf(s1))

	// Output:
	// true
}

// Iterator 使用给定的回调函数 `f` 遍历只读集合，如果 `f` 返回 true，则继续遍历；否则停止。 md5:b896360b1cf6fc88
func ExampleStrSet_Iterator() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	s1.Iterator(func(v string) bool {
		fmt.Println("Iterator", v)
		return true
	})

	// May Output:
	// Iterator a
	// Iterator b
	// Iterator c
	// Iterator d
}

// Join 使用字符串 `glue` 连接多个项目。 md5:c8699391999ac788
func ExampleStrSet_Join() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.Join(","))

	// May Output:
	// b,c,d,a
}

// LockFunc 使用回调函数 `f` 为写入操作加锁。 md5:85d746d8a49edab7
func ExampleStrSet_LockFunc() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"1", "2"}...)
	s1.LockFunc(func(m map[string]struct{}) {
		m["3"] = struct{}{}
	})
	fmt.Println(s1.Slice())

	// 可能的输出
	// [2 3 1] md5:294c6ba36e85ea4c

}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func ExampleStrSet_MarshalJSON() {
	type Student struct {
		Id     int
		Name   string
		Scores *gset.StrSet
	}
	s := Student{
		Id:     1,
		Name:   "john",
		Scores: gset.NewStrSetFrom([]string{"100", "99", "98"}, true),
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))

	// May Output:
	// {"Id":1,"Name":"john","Scores":["100","99","98"]}
}

// Merge 将 `others` 集合中的项目合并到 `set` 中。 md5:788b02e300c6f440
func ExampleStrSet_Merge() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)

	s2 := gset.NewStrSet(true)
	fmt.Println(s1.Merge(s2).Slice())

	// May Output:
	// [d a b c]
}

// 随机从集合中弹出一个元素。 md5:56ac5a59d1852551
func ExampleStrSet_Pop() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)

	fmt.Println(s1.Pop())

	// May Output:
	// a
}

// Pops 从集合中随机弹出 `size` 个元素。
// 如果 size == -1，它将返回所有元素。 md5:c687f88e0a2df8f2
func ExampleStrSet_Pops() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	for _, v := range s1.Pops(2) {
		fmt.Println(v)
	}

	// May Output:
	// a
	// b
}

// RLockFunc 使用回调函数 `f` 进行读取锁定。 md5:5fe2bf1a85ce319e
func ExampleStrSet_RLockFunc() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	s1.RLockFunc(func(m map[string]struct{}) {
		fmt.Println(m)
	})

	// Output:
	// map[a:{} b:{} c:{} d:{}]
}

// Remove 从集合中删除 `item`。 md5:ab30c696cc44d190
func ExampleStrSet_Remove() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	s1.Remove("a")
	fmt.Println(s1.Slice())

	// May Output:
	// [b c d]
}

// Size 返回集合的大小。 md5:0d55ac576b7779ee
func ExampleStrSet_Size() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.Size())

	// Output:
	// 4
}

// Slice 返回集合中的元素作为切片。 md5:f5bc80ac01ae812b
func ExampleStrSet_Slice() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.Slice())

	// May Output:
	// [a,b,c,d]
}

// String 将 items 转换为字符串，其实现方式类似于 json.Marshal。 md5:cedb10711c2e5dac
func ExampleStrSet_String() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	fmt.Println(s1.String())

	// May Output:
	// "a","b","c","d"
}

// Sum 对项目求和。注意：项目应转换为整型，
// 否则你可能会得到意想不到的结果。 md5:7cca75708fbf4ffc
func ExampleStrSet_Sum() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"1", "2", "3", "4"}...)
	fmt.Println(s1.Sum())

	// Output:
	// 10
}

// Union 返回一个新集合，它是`set`和`other`的并集。
// 意味着，`newSet`中的所有项目都在`set`中或在`other`中。 md5:420e241c3c12e8e6
func ExampleStrSet_Union() {
	s1 := gset.NewStrSet(true)
	s1.Add([]string{"a", "b", "c", "d"}...)
	s2 := gset.NewStrSet(true)
	s2.Add([]string{"a", "b", "d"}...)
	fmt.Println(s1.Union(s2).Slice())

	// May Output:
	// [a b c d]
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func ExampleStrSet_UnmarshalJSON() {
	b := []byte(`{"Id":1,"Name":"john","Scores":["100","99","98"]}`)
	type Student struct {
		Id     int
		Name   string
		Scores *gset.StrSet
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)

	// May Output:
	// {1 john "99","98","100"}
}

// UnmarshalValue 是一个接口实现，用于将任何类型的价值设置为集合。 md5:b119247f684920ad
func ExampleStrSet_UnmarshalValue() {
	b := []byte(`{"Id":1,"Name":"john","Scores":["100","99","98"]}`)
	type Student struct {
		Id     int
		Name   string
		Scores *gset.StrSet
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)

	// May Output:
	// {1 john "99","98","100"}
}

// Walk应用用户提供的函数`f`到集合中的每一项。 md5:d6ceaae555e8a9e6
func ExampleStrSet_Walk() {
	var (
		set    gset.StrSet
		names  = g.SliceStr{"user", "user_detail"}
		prefix = "gf_"
	)
	set.Add(names...)
	// 为给定的表名添加前缀。 md5:dea7405f272e0c9e
	set.Walk(func(item string) string {
		return prefix + item
	})
	fmt.Println(set.Slice())

	// May Output:
	// [gf_user gf_user_detail]
}
