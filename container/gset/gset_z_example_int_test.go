// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

package 集合类_test

import (
	"encoding/json"
	"fmt"

	gset "github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/frame/g"
)

// New 创建并返回一个新的集合，其中包含不重复的项目。
// 参数 `safe` 用于指定在并发安全模式下使用集合，其默认为 false。
// md5:db8312fdb3f679d3
func ExampleNewIntSet() {
	intSet := gset.X创建整数()
	intSet.X加入([]int{1, 2, 3}...)
	fmt.Println(intSet.X取集合切片())

	// May Output:
	// [2 1 3]
}

// NewIntSetFrom 根据`items`返回一个新的集合。 md5:7b944f3609c229f9
func ExampleNewFrom() {
	intSet := gset.X创建整数并按值([]int{1, 2, 3})
	fmt.Println(intSet.X取集合切片())

	// May Output:
	// [2 1 3]
}

// Add 将一个或多个项目添加到集合中。 md5:316141ff7d4b8e45
func ExampleIntSet_Add() {
	intSet := gset.X创建整数并按值([]int{1, 2, 3})
	intSet.X加入(1)
	fmt.Println(intSet.X取集合切片())
	fmt.Println(intSet.X加入值并跳过已存在(1))

	// May Output:
	// [1 2 3]
	// false
}

// AddIfNotExist 检查项是否存在于集合中，
// 如果项不存在于集合中，它会将项添加到集合中并返回 true，否则什么都不做并返回 false。
// md5:3a8a0467b52a54c0
func ExampleIntSet_AddIfNotExist() {
	intSet := gset.X创建整数并按值([]int{1, 2, 3})
	intSet.X加入(1)
	fmt.Println(intSet.X取集合切片())
	fmt.Println(intSet.X加入值并跳过已存在(1))

	// May Output:
	// [1 2 3]
	// false
}

// AddIfNotExistFunc 检查项是否存在于集合中，
// 如果项不存在于集合中且函数 `f` 返回 true，它会将项添加到集合中并返回 true，
// 否则，它什么都不做并返回 false。
// 请注意，函数 `f` 在写入锁之外执行。
// md5:a60fff9115523801
func ExampleIntSet_AddIfNotExistFunc() {
	intSet := gset.X创建整数并按值([]int{1, 2, 3})
	intSet.X加入(1)
	fmt.Println(intSet.X取集合切片())
	fmt.Println(intSet.X加入值并跳过已存在_函数(5, func() bool {
		return true
	}))

	// May Output:
	// [1 2 3]
	// true
}

// AddIfNotExistFunc 检查项是否存在于集合中，
// 如果项不存在于集合中且函数 `f` 返回 true，它会将项添加到集合中并返回 true，
// 否则，它什么都不做并返回 false。
// 请注意，函数 `f` 在写入锁之外执行。
// md5:a60fff9115523801
func ExampleIntSet_AddIfNotExistFuncLock() {
	intSet := gset.X创建整数并按值([]int{1, 2, 3})
	intSet.X加入(1)
	fmt.Println(intSet.X取集合切片())
	fmt.Println(intSet.X加入值并跳过已存在_并发安全函数(4, func() bool {
		return true
	}))

	// May Output:
	// [1 2 3]
	// true
}

// Clear 删除集合中的所有项。 md5:ce349f0cd3114465
func ExampleIntSet_Clear() {
	intSet := gset.X创建整数并按值([]int{1, 2, 3})
	fmt.Println(intSet.X取数量())
	intSet.X清空()
	fmt.Println(intSet.X取数量())

	// Output:
	// 3
	// 0
}

// Complement 返回一个新的集合，该集合是相对于`set`到`full`的补集。
// 这意味着，`newSet`中的所有项都包含在`full`中但不包含在`set`中。
// 如果给定的集合`full`并不是`set`的全集，则它返回`full`与`set`之间的差异。
// md5:2116fbb7587db792
func ExampleIntSet_Complement() {
	intSet := gset.X创建整数并按值([]int{1, 2, 3, 4, 5})
	s := gset.X创建整数并按值([]int{1, 2, 3})
	fmt.Println(s.X取补集(intSet).X取集合切片())

	// May Output:
	// [4 5]
}

// Contains 检查集合是否包含 `item`。 md5:20a3bdc6aeef1d67
func ExampleIntSet_Contains() {
	var set1 gset.IntSet
	set1.X加入(1, 4, 5, 6, 7)
	fmt.Println(set1.X是否存在(1))

	var set2 gset.IntSet
	set2.X加入(1, 4, 5, 6, 7)
	fmt.Println(set2.X是否存在(8))

	// Output:
	// true
	// false
}

// Diff 返回一个新的集合，它是 `set` 与 `other` 之间的差集。
// 这意味着，`newSet` 中的所有项目都在 `set` 中，但不在 `other` 中。
// md5:6779e6e007651b53
func ExampleIntSet_Diff() {
	s1 := gset.X创建整数并按值([]int{1, 2, 3})
	s2 := gset.X创建整数并按值([]int{1, 2, 3, 4})
	fmt.Println(s2.X取差集(s1).X取集合切片())

	// Output:
	// [4]
}

// Equal 检查两个集合是否相等。 md5:105ea4dd39b57fe8
func ExampleIntSet_Equal() {
	s1 := gset.X创建整数并按值([]int{1, 2, 3})
	s2 := gset.X创建整数并按值([]int{1, 2, 3, 4})
	fmt.Println(s2.X是否相等(s1))

	s3 := gset.X创建整数并按值([]int{1, 2, 3})
	s4 := gset.X创建整数并按值([]int{1, 2, 3})
	fmt.Println(s3.X是否相等(s4))

	// Output:
	// false
	// true
}

// Intersect 返回一个新的集合，这个集合是 `set` 和 `other` 的交集。
// 这意味着，`newSet` 中的所有元素都既存在于 `set` 中也存在于 `other` 中。
// md5:327d3fcc12f06583
func ExampleIntSet_Intersect() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3}...)
	var s2 gset.IntSet
	s2.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s2.X取交集(s1).X取集合切片())

	// May Output:
	// [1 2 3]
}

// IsSubsetOf 检查当前集合是否是 `other` 的子集. md5:70b7ed1e77ec2f80
func ExampleIntSet_IsSubsetOf() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	var s2 gset.IntSet
	s2.X加入([]int{1, 2, 4}...)
	fmt.Println(s2.X是否为子集(s1))

	// Output:
	// true
}

// Iterator 使用给定的回调函数 `f` 遍历只读集合，如果 `f` 返回 true，则继续遍历；否则停止。
// md5:b896360b1cf6fc88
func ExampleIntSet_Iterator() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	s1.X遍历(func(v int) bool {
		fmt.Println("Iterator", v)
		return true
	})
	// May Output:
	// Iterator 2
	// Iterator 3
	// Iterator 1
	// Iterator 4
}

// Join 使用字符串 `glue` 连接多个项目。 md5:c8699391999ac788
func ExampleIntSet_Join() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s1.X取集合文本(","))

	// May Output:
	// 3,4,1,2
}

// LockFunc 使用回调函数 `f` 为写入操作加锁。 md5:85d746d8a49edab7
func ExampleIntSet_LockFunc() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2}...)
	s1.X写锁定_函数(func(m map[int]struct{}) {
		m[3] = struct{}{}
	})
	fmt.Println(s1.X取集合切片())

	// 可能的输出
	// [2 3 1]
	// md5:294c6ba36e85ea4c
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func ExampleIntSet_MarshalJSON() {
	type Student struct {
		Id     int
		Name   string
		Scores *gset.IntSet
	}
	s := Student{
		Id:     1,
		Name:   "john",
		Scores: gset.X创建整数并按值([]int{100, 99, 98}),
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))

	// May Output:
	// {"Id":1,"Name":"john","Scores":[100,99,98]}
}

// Merge 将 `others` 集合中的项目合并到 `set` 中。 md5:788b02e300c6f440
func ExampleIntSet_Merge() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)

	s2 := gset.X创建整数()
	fmt.Println(s1.X合并(s2).X取集合切片())

	// May Output:
	// [1 2 3 4]
}

// 随机从集合中弹出一个元素。 md5:56ac5a59d1852551
func ExampleIntSet_Pop() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)

	fmt.Println(s1.X出栈())

	// May Output:
	// 1
}

// Pops 从集合中随机弹出 `size` 个元素。
// 如果 size == -1，它将返回所有元素。
// md5:c687f88e0a2df8f2
func ExampleIntSet_Pops() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	for _, v := range s1.X出栈多个(2) {
		fmt.Println(v)
	}

	// May Output:
	// 1
	// 2
}

// RLockFunc 使用回调函数 `f` 进行读取锁定。 md5:5fe2bf1a85ce319e
func ExampleIntSet_RLockFunc() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	s1.X读锁定_函数(func(m map[int]struct{}) {
		fmt.Println(m)
	})

	// Output:
	// map[1:{} 2:{} 3:{} 4:{}]
}

// Remove 从集合中删除 `item`。 md5:ab30c696cc44d190
func ExampleIntSet_Remove() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	s1.X删除(1)
	fmt.Println(s1.X取集合切片())

	// May Output:
	// [3 4 2]
}

// Size 返回集合的大小。 md5:0d55ac576b7779ee
func ExampleIntSet_Size() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s1.X取数量())

	// Output:
	// 4
}

// Slice 返回集合中的元素作为切片。 md5:f5bc80ac01ae812b
func ExampleIntSet_Slice() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s1.X取集合切片())

	// May Output:
	// [1, 2, 3, 4]
}

// String 将 items 转换为字符串，其实现方式类似于 json.Marshal。 md5:cedb10711c2e5dac
func ExampleIntSet_String() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s1.String())

	// May Output:
	// [1,2,3,4]
}

// Sum 对项目求和。注意：项目应转换为整型，
// 否则你可能会得到意想不到的结果。
// md5:7cca75708fbf4ffc
func ExampleIntSet_Sum() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s1.X求和())

	// Output:
	// 10
}

// Union 返回一个新集合，它是`set`和`other`的并集。
// 意味着，`newSet`中的所有项目都在`set`中或在`other`中。
// md5:420e241c3c12e8e6
func ExampleIntSet_Union() {
	s1 := gset.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	s2 := gset.X创建整数()
	s2.X加入([]int{1, 2, 4}...)
	fmt.Println(s1.X取并集(s2).X取集合切片())

	// May Output:
	// [3 4 1 2]
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func ExampleIntSet_UnmarshalJSON() {
	b := []byte(`{"Id":1,"Name":"john","Scores":[100,99,98]}`)
	type Student struct {
		Id     int
		Name   string
		Scores *gset.IntSet
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)

	// May Output:
	// {1 john [100,99,98]}
}

// UnmarshalValue 是一个接口实现，用于将任何类型的价值设置为集合。 md5:b119247f684920ad
func ExampleIntSet_UnmarshalValue() {
	b := []byte(`{"Id":1,"Name":"john","Scores":100,99,98}`)
	type Student struct {
		Id     int
		Name   string
		Scores *gset.IntSet
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)

	// May Output:
	// {1 john [100,99,98]}
}

// Walk应用用户提供的函数`f`到集合中的每一项。 md5:d6ceaae555e8a9e6
func ExampleIntSet_Walk() {
	var (
		set   gset.IntSet
		names = g.SliceInt别名{1, 0}
		delta = 10
	)
	set.X加入(names...)
		// 为给定的表名添加前缀。 md5:dea7405f272e0c9e
	set.X遍历修改(func(item int) int {
		return delta + item
	})
	fmt.Println(set.X取集合切片())

	// May Output:
	// [12 60]
}
