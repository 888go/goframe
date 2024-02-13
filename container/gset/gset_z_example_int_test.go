// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package 集合类_test

import (
	"encoding/json"
	"fmt"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/frame/g"
)

// New 创建并返回一个新的集合，其中包含无重复项。
// 参数`safe`用于指定是否在并发安全的情况下使用集合，默认为false。
func ExampleNewIntSet() {
	intSet := 集合类.X创建整数()
	intSet.X加入([]int{1, 2, 3}...)
	fmt.Println(intSet.X取集合数组())

	// May Output:
	// [2 1 3]
}

// NewIntSetFrom 返回一个从`items`创建的新集合。
func ExampleNewFrom() {
	intSet := 集合类.X创建整数并按值([]int{1, 2, 3})
	fmt.Println(intSet.X取集合数组())

	// May Output:
	// [2 1 3]
}

// Add 向集合中添加一个或多个项目。
func ExampleIntSet_Add() {
	intSet := 集合类.X创建整数并按值([]int{1, 2, 3})
	intSet.X加入(1)
	fmt.Println(intSet.X取集合数组())
	fmt.Println(intSet.X加入值并跳过已存在(1))

	// Mya Output:
	// [1 2 3]
	// false
}

// AddIfNotExist 检查项是否已存在于集合中，
// 如果该项不在集合中，则将其添加到集合中并返回 true；
// 否则，不进行任何操作并返回 false。
func ExampleIntSet_AddIfNotExist() {
	intSet := 集合类.X创建整数并按值([]int{1, 2, 3})
	intSet.X加入(1)
	fmt.Println(intSet.X取集合数组())
	fmt.Println(intSet.X加入值并跳过已存在(1))

	// Mya Output:
	// [1 2 3]
	// false
}

// AddIfNotExistFunc 检查项是否存在集合中，
// 如果该项不存在于集合中，并且函数 `f` 返回 true，则将该项添加到集合中并返回 true，
// 否则不做任何操作并返回 false。
// 注意，函数 `f` 在无写入锁的情况下执行。
func ExampleIntSet_AddIfNotExistFunc() {
	intSet := 集合类.X创建整数并按值([]int{1, 2, 3})
	intSet.X加入(1)
	fmt.Println(intSet.X取集合数组())
	fmt.Println(intSet.X加入值并跳过已存在_函数(5, func() bool {
		return true
	}))

	// May Output:
	// [1 2 3]
	// true
}

// AddIfNotExistFunc 检查项是否存在集合中，
// 如果该项不存在于集合中，并且函数 `f` 返回 true，则将该项添加到集合中并返回 true，
// 否则不做任何操作并返回 false。
// 注意，函数 `f` 在无写入锁的情况下执行。
func ExampleIntSet_AddIfNotExistFuncLock() {
	intSet := 集合类.X创建整数并按值([]int{1, 2, 3})
	intSet.X加入(1)
	fmt.Println(intSet.X取集合数组())
	fmt.Println(intSet.X加入值并跳过已存在_并发安全函数(4, func() bool {
		return true
	}))

	// May Output:
	// [1 2 3]
	// true
}

// 清除删除集合中的所有项。
func ExampleIntSet_Clear() {
	intSet := 集合类.X创建整数并按值([]int{1, 2, 3})
	fmt.Println(intSet.X取数量())
	intSet.X清空()
	fmt.Println(intSet.X取数量())

	// Output:
	// 3
	// 0
}

// Complement 返回一个新的集合，它是从 `set` 到 `full` 的补集。
// 这意味着，新集合`newSet`中的所有元素都在`full`中但不在`set`中。
// 如果给定的集合 `full` 不是 `set` 的全集，则返回 `full` 和 `set` 之间的差集。
func ExampleIntSet_Complement() {
	intSet := 集合类.X创建整数并按值([]int{1, 2, 3, 4, 5})
	s := 集合类.X创建整数并按值([]int{1, 2, 3})
	fmt.Println(s.X取补集(intSet).X取集合数组())

	// May Output:
	// [4 5]
}

// Contains 检查集合中是否包含 `item`。
func ExampleIntSet_Contains() {
	var set1 集合类.IntSet
	set1.X加入(1, 4, 5, 6, 7)
	fmt.Println(set1.X是否存在(1))

	var set2 集合类.IntSet
	set2.X加入(1, 4, 5, 6, 7)
	fmt.Println(set2.X是否存在(8))

	// Output:
	// true
	// false
}

// Diff 返回一个新的集合，这个集合是 `set` 与 `other` 的差集。
// 这意味着，新集合 `newSet` 中的所有元素都在 `set` 中，但不在 `other` 中。
func ExampleIntSet_Diff() {
	s1 := 集合类.X创建整数并按值([]int{1, 2, 3})
	s2 := 集合类.X创建整数并按值([]int{1, 2, 3, 4})
	fmt.Println(s2.X取差集(s1).X取集合数组())

	// Output:
	// [4]
}

// Equal 检查两个集合是否相等。
func ExampleIntSet_Equal() {
	s1 := 集合类.X创建整数并按值([]int{1, 2, 3})
	s2 := 集合类.X创建整数并按值([]int{1, 2, 3, 4})
	fmt.Println(s2.X是否相等(s1))

	s3 := 集合类.X创建整数并按值([]int{1, 2, 3})
	s4 := 集合类.X创建整数并按值([]int{1, 2, 3})
	fmt.Println(s3.X是否相等(s4))

	// Output:
	// false
	// true
}

// Intersect 返回一个新的集合，它是从 `set` 到 `other` 的交集。
// 这意味着，`newSet` 中的所有元素都在 `set` 中，并且也在 `other` 中。
func ExampleIntSet_Intersect() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3}...)
	var s2 集合类.IntSet
	s2.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s2.X取交集(s1).X取集合数组())

	// May Output:
	// [1 2 3]
}

// IsSubsetOf 检查当前集合是否为 `other` 的子集
func ExampleIntSet_IsSubsetOf() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	var s2 集合类.IntSet
	s2.X加入([]int{1, 2, 4}...)
	fmt.Println(s2.X是否为子集(s1))

	// Output:
	// true
}

// Iterator 使用给定的回调函数`f`对集合进行只读遍历，
// 如果`f`返回true，则继续遍历；若返回false，则停止遍历。
func ExampleIntSet_Iterator() {
	s1 := 集合类.X创建整数()
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

// Join通过字符串`glue`连接items。
func ExampleIntSet_Join() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s1.X取集合文本(","))

	// May Output:
	// 3,4,1,2
}

// LockFunc 使用回调函数`f`进行写入锁定。
func ExampleIntSet_LockFunc() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2}...)
	s1.X写锁定_函数(func(m map[int]struct{}) {
		m[3] = struct{}{}
	})
	fmt.Println(s1.X取集合数组())

// 可能的输出
// [2 3 1]
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func ExampleIntSet_MarshalJSON() {
	type Student struct {
		Id     int
		Name   string
		Scores *集合类.IntSet
	}
	s := Student{
		Id:     1,
		Name:   "john",
		Scores: 集合类.X创建整数并按值([]int{100, 99, 98}),
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))

	// May Output:
	// {"Id":1,"Name":"john","Scores":[100,99,98]}
}

// Merge 将 `others` 中的元素合并到 `set` 中。
func ExampleIntSet_Merge() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)

	s2 := 集合类.X创建整数()
	fmt.Println(s1.X合并(s2).X取集合数组())

	// May Output:
	// [1 2 3 4]
}

// Pops 随机地从集合中弹出一个元素。
func ExampleIntSet_Pop() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)

	fmt.Println(s1.X出栈())

	// May Output:
	// 1
}

// Pops 随机地从集合中弹出 `size` 个元素。
// 如果 size == -1，则返回所有元素。
func ExampleIntSet_Pops() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	for _, v := range s1.X出栈多个(2) {
		fmt.Println(v)
	}

	// May Output:
	// 1
	// 2
}

// RLockFunc 通过回调函数 `f` 对读取进行加锁。
func ExampleIntSet_RLockFunc() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	s1.X读锁定_函数(func(m map[int]struct{}) {
		fmt.Println(m)
	})

	// Output:
	// map[1:{} 2:{} 3:{} 4:{}]
}

// Remove 从集合中删除`item`。
func ExampleIntSet_Remove() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	s1.X删除(1)
	fmt.Println(s1.X取集合数组())

	// May Output:
	// [3 4 2]
}

// Size 返回集合的大小。
func ExampleIntSet_Size() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s1.X取数量())

	// Output:
	// 4
}

// Slice 返回集合中项目的切片形式。
func ExampleIntSet_Slice() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s1.X取集合数组())

	// May Output:
	// [1, 2, 3, 4]
}

// String 返回 items 作为字符串，其实现方式类似于 json.Marshal。
func ExampleIntSet_String() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s1.String())

	// May Output:
	// [1,2,3,4]
}

// Sum 计算项目总和。注意：项目应转换为 int 类型，
// 否则你可能会得到一个意想不到的结果。
func ExampleIntSet_Sum() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	fmt.Println(s1.X求和())

	// Output:
	// 10
}

// Union 返回一个新的集合，该集合是 `set` 和 `other` 的并集。
// 这意味着，`newSet` 中的所有元素都在 `set` 或者 `other` 中。
func ExampleIntSet_Union() {
	s1 := 集合类.X创建整数()
	s1.X加入([]int{1, 2, 3, 4}...)
	s2 := 集合类.X创建整数()
	s2.X加入([]int{1, 2, 4}...)
	fmt.Println(s1.X取并集(s2).X取集合数组())

	// May Output:
	// [3 4 1 2]
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func ExampleIntSet_UnmarshalJSON() {
	b := []byte(`{"Id":1,"Name":"john","Scores":[100,99,98]}`)
	type Student struct {
		Id     int
		Name   string
		Scores *集合类.IntSet
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)

	// May Output:
	// {1 john [100,99,98]}
}

// UnmarshalValue 是一个接口实现，用于为 set 设置任意类型的值。
func ExampleIntSet_UnmarshalValue() {
	b := []byte(`{"Id":1,"Name":"john","Scores":100,99,98}`)
	type Student struct {
		Id     int
		Name   string
		Scores *集合类.IntSet
	}
	s := Student{}
	json.Unmarshal(b, &s)
	fmt.Println(s)

	// May Output:
	// {1 john [100,99,98]}
}

// Walk 对集合中的每一个元素应用用户提供的函数 `f`。
func ExampleIntSet_Walk() {
	var (
		set   集合类.IntSet
		names = g.SliceInt别名{1, 0}
		delta = 10
	)
	set.X加入(names...)
	// 为给定的表名添加前缀
	set.X遍历修改(func(item int) int {
		return delta + item
	})
	fmt.Println(set.X取集合数组())

	// May Output:
	// [12 60]
}
