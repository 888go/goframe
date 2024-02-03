// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package garray_test

import (
	"fmt"
	
	"github.com/888go/goframe/internal/empty"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/frame/g"
)

func ExampleNew() {
	// A normal array.
	a := garray.New()

	// Adding items.
	for i := 0; i < 10; i++ {
		a.Append(i)
	}

	// 打印数组长度。
	fmt.Println(a.Len())

	// 打印数组元素。
	fmt.Println(a.Slice())

	// 通过索引获取项。
	fmt.Println(a.Get(6))

	// 检查项目是否存在。
	fmt.Println(a.Contains(6))
	fmt.Println(a.Contains(100))

	// 在指定索引之前插入项。
	a.InsertAfter(9, 11)
	// 在指定索引之后插入项目。
	a.InsertBefore(10, 10)

	fmt.Println(a.Slice())

	// 通过索引修改项。
	a.Set(0, 100)
	fmt.Println(a.Slice())

	fmt.Println(a.At(0))

	// 搜索指定项并返回其索引。
	fmt.Println(a.Search(5))

	// 通过索引移除项
	a.Remove(0)
	fmt.Println(a.Slice())

	// 清空数组，移除其所有元素。
	fmt.Println(a.Slice())
	a.Clear()
	fmt.Println(a.Slice())

	// Output:
	// 10
	// [0 1 2 3 4 5 6 7 8 9]
	// 6 true
	// true
	// false
	// [0 1 2 3 4 5 6 7 8 9 10 11]
	// [100 1 2 3 4 5 6 7 8 9 10 11]
	// 100
	// 5
	// [1 2 3 4 5 6 7 8 9 10 11]
	// [1 2 3 4 5 6 7 8 9 10 11]
	// []
}

func ExampleArray_Iterator() {
	array := garray.NewArrayFrom(g.Slice{"a", "b", "c"})
// Iterator 是 IteratorAsc 的别名，用于以升序方式对数组进行只读遍历，
// 同时调用给定的回调函数 `f`。
// 若 `f` 返回 true，则继续遍历；若返回 false，则停止遍历。
	array.Iterator(func(k int, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
// IteratorDesc 以降序方式遍历给定回调函数 `f` 的只读数组。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止遍历。
// 这段Go语言代码注释翻译成中文注释如下：
// ```go
// IteratorDesc 函数以降序顺序对给定的只读数组进行迭代，并使用指定的回调函数 `f` 进行处理。
// 若回调函数 `f` 返回值为 true，则会继续进行迭代；若返回值为 false，则停止迭代过程。
	array.IteratorDesc(func(k int, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

	// Output:
	// 0 a
	// 1 b
	// 2 c
	// 2 c
	// 1 b
	// 0 a
}

func ExampleArray_Reverse() {
	array := garray.NewFrom(g.Slice{1, 2, 3, 4, 5, 6, 7, 8, 9})

	// Reverse将数组元素按逆序排列。
	fmt.Println(array.Reverse().Slice())

	// Output:
	// [9 8 7 6 5 4 3 2 1]
}

func ExampleArray_Shuffle() {
	array := garray.NewFrom(g.Slice{1, 2, 3, 4, 5, 6, 7, 8, 9})

	// Shuffle 随机地对数组进行洗牌。
	fmt.Println(array.Shuffle().Slice())
}

func ExampleArray_Rands() {
	array := garray.NewFrom(g.Slice{1, 2, 3, 4, 5, 6, 7, 8, 9})

// 随机从数组中获取并返回 2 个元素。
// 不会从数组中删除这些元素。
	fmt.Println(array.Rands(2))

// 从数组中随机选取并返回一个元素。
// 它会从数组中删除已选取的元素。
	fmt.Println(array.PopRand())
}

func ExampleArray_PopRand() {
	array := garray.NewFrom(g.Slice{1, 2, 3, 4, 5, 6, 7, 8, 9})

// 随机从数组中获取并返回 2 个元素。
// 不会从数组中删除这些元素。
	fmt.Println(array.Rands(2))

// 从数组中随机选取并返回一个元素。
// 它会从数组中删除已选取的元素。
	fmt.Println(array.PopRand())
}

func ExampleArray_Join() {
	array := garray.NewFrom(g.Slice{"a", "b", "c", "d"})
	fmt.Println(array.Join(","))

	// Output:
	// a,b,c,d
}

func ExampleArray_Chunk() {
	array := garray.NewFrom(g.Slice{1, 2, 3, 4, 5, 6, 7, 8, 9})

// Chunk 函数将一个数组分割成多个子数组，
// 每个子数组的大小由参数 `size` 确定。
// 最后一个子数组可能包含少于 size 个元素。
	fmt.Println(array.Chunk(2))

	// Output:
	// [[1 2] [3 4] [5 6] [7 8] [9]]
}

func ExampleArray_PopLeft() {
	array := garray.NewFrom([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})

	// 任何 Pop* 函数都会从数组中挑选、删除并返回一个元素。

	fmt.Println(array.PopLeft())
	fmt.Println(array.PopLefts(2))
	fmt.Println(array.PopRight())
	fmt.Println(array.PopRights(2))

	// Output:
	// 1 true
	// [2 3]
	// 9 true
	// [7 8]
}

func ExampleArray_PopLefts() {
	array := garray.NewFrom([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})

	// 任何 Pop* 函数都会从数组中挑选、删除并返回一个元素。

	fmt.Println(array.PopLeft())
	fmt.Println(array.PopLefts(2))
	fmt.Println(array.PopRight())
	fmt.Println(array.PopRights(2))

	// Output:
	// 1 true
	// [2 3]
	// 9 true
	// [7 8]
}

func ExampleArray_PopRight() {
	array := garray.NewFrom([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})

	// 任何 Pop* 函数都会从数组中挑选、删除并返回一个元素。

	fmt.Println(array.PopLeft())
	fmt.Println(array.PopLefts(2))
	fmt.Println(array.PopRight())
	fmt.Println(array.PopRights(2))

	// Output:
	// 1 true
	// [2 3]
	// 9 true
	// [7 8]
}

func ExampleArray_PopRights() {
	array := garray.NewFrom([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})

	// 任何 Pop* 函数都会从数组中挑选、删除并返回一个元素。

	fmt.Println(array.PopLeft())
	fmt.Println(array.PopLefts(2))
	fmt.Println(array.PopRight())
	fmt.Println(array.PopRights(2))

	// Output:
	// 1 true
	// [2 3]
	// 9 true
	// [7 8]
}

func ExampleArray_Contains() {
	var array garray.StrArray
	array.Append("a")
	fmt.Println(array.Contains("a"))
	fmt.Println(array.Contains("A"))
	fmt.Println(array.ContainsI("A"))

	// Output:
	// true
	// false
	// true
}

func ExampleArray_Merge() {
	array1 := garray.NewFrom(g.Slice{1, 2})
	array2 := garray.NewFrom(g.Slice{3, 4})
	slice1 := g.Slice{5, 6}
	slice2 := []int{7, 8}
	slice3 := []string{"9", "0"}
	fmt.Println(array1.Slice())
	array1.Merge(array1)
	array1.Merge(array2)
	array1.Merge(slice1)
	array1.Merge(slice2)
	array1.Merge(slice3)
	fmt.Println(array1.Slice())

	// Output:
	// [1 2]
	// [1 2 1 2 3 4 5 6 7 8 9 0]
}

func ExampleArray_Filter() {
	array1 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	array2 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	fmt.Printf("%#v\n", array1.Filter(func(index int, value interface{}) bool {
		return empty.IsNil(value)
	}).Slice())
	fmt.Printf("%#v\n", array2.Filter(func(index int, value interface{}) bool {
		return empty.IsEmpty(value)
	}).Slice())

	// Output:
	// []interface {}{0, 1, 2, "", []interface {}{}, "john"}
	// []interface {}{1, 2, "john"}
}

func ExampleArray_FilterEmpty() {
	array1 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	array2 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	fmt.Printf("%#v\n", array1.FilterNil().Slice())
	fmt.Printf("%#v\n", array2.FilterEmpty().Slice())

	// Output:
	// []interface {}{0, 1, 2, "", []interface {}{}, "john"}
	// []interface {}{1, 2, "john"}
}

func ExampleArray_FilterNil() {
	array1 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	array2 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	fmt.Printf("%#v\n", array1.FilterNil().Slice())
	fmt.Printf("%#v\n", array2.FilterEmpty().Slice())

	// Output:
	// []interface {}{0, 1, 2, "", []interface {}{}, "john"}
	// []interface {}{1, 2, "john"}
}
