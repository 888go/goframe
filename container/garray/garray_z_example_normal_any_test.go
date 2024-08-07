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
)

func ExampleNew() {
	// A normal array.
	a := garray.X创建()

	// Adding items.
	for i := 0; i < 10; i++ {
		a.Append别名(i)
	}

		// 打印数组的长度。 md5:b662aa8b8c6552e9
	fmt.Println(a.X取长度())

	// Print the array items.
	fmt.Println(a.X取切片())

		// 通过索引获取项目。 md5:4c92b4ee2f35a4fa
	fmt.Println(a.X取值2(6))

	// Check item existence.
	fmt.Println(a.X是否存在(6))
	fmt.Println(a.X是否存在(100))

		// 在指定索引前插入项。 md5:608613f475a29c19
	a.X插入后面(9, 11)
		// 在指定索引后插入元素。 md5:c20b7bd92b22a29e
	a.X插入前面(10, 10)

	fmt.Println(a.X取切片())

	// Modify item by index.
	a.X设置值(0, 100)
	fmt.Println(a.X取切片())

	fmt.Println(a.X取值(0))

		// 搜索元素并返回其索引。 md5:3f6b71dc616bddc3
	fmt.Println(a.X查找(5))

	// Remove item by index.
	a.X删除(0)
	fmt.Println(a.X取切片())

		// 清空数组，删除其中的所有元素。 md5:6fbd931483046674
	fmt.Println(a.X取切片())
	a.X清空()
	fmt.Println(a.X取切片())

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
	array := garray.X创建并从切片(g.Slice别名{"a", "b", "c"})
	// Iterator 是 IteratorAsc 的别名，它按照升序遍历数组，并使用给定的回调函数 `f`。如果 `f` 返回 true，则继续遍历；否则停止。
	// md5:d842b3c6584033ab
	array.X遍历(func(k int, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	// IteratorDesc 以降序遍历数组，并使用给定的回调函数`f`进行只读迭代。
	// 如果`f`返回true，则继续遍历；如果返回false，则停止遍历。
	// md5:94f26122239ef7ac
	array.X遍历降序(func(k int, v interface{}) bool {
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
	array := garray.NewFrom别名(g.Slice别名{1, 2, 3, 4, 5, 6, 7, 8, 9})

		// Reverse 函数将数组元素反转顺序。 md5:cc34cd0a2fa08e1c
	fmt.Println(array.X倒排序().X取切片())

	// Output:
	// [9 8 7 6 5 4 3 2 1]
}

func ExampleArray_Shuffle() {
	array := garray.NewFrom别名(g.Slice别名{1, 2, 3, 4, 5, 6, 7, 8, 9})

		// 随机打乱数组。 md5:5897797461d9f11a
	fmt.Println(array.X随机排序().X取切片())
}

func ExampleArray_Rands() {
	array := garray.NewFrom别名(g.Slice别名{1, 2, 3, 4, 5, 6, 7, 8, 9})

	// 随机从数组中获取并返回两个元素。
	// 它不会从数组中删除这些元素。
	// md5:6df7401a881dedbb
	fmt.Println(array.X取值随机多个(2))

	// 随机从数组中选择一个元素并返回。
	// 从数组中删除选中的元素。
	// md5:5c923218de5c63ae
	fmt.Println(array.X出栈随机())
}

func ExampleArray_PopRand() {
	array := garray.NewFrom别名(g.Slice别名{1, 2, 3, 4, 5, 6, 7, 8, 9})

	// 随机从数组中获取并返回两个元素。
	// 它不会从数组中删除这些元素。
	// md5:6df7401a881dedbb
	fmt.Println(array.X取值随机多个(2))

	// 随机从数组中选择一个元素并返回。
	// 从数组中删除选中的元素。
	// md5:5c923218de5c63ae
	fmt.Println(array.X出栈随机())
}

func ExampleArray_Join() {
	array := garray.NewFrom别名(g.Slice别名{"a", "b", "c", "d"})
	fmt.Println(array.X连接(","))

	// Output:
	// a,b,c,d
}

func ExampleArray_Chunk() {
	array := garray.NewFrom别名(g.Slice别名{1, 2, 3, 4, 5, 6, 7, 8, 9})

	// Chunk 将一个数组分割成多个子数组，每个子数组的大小由 `size` 决定。最后一个子数组可能包含少于 `size` 个元素。
	// md5:5de5729fde1de2d3
	fmt.Println(array.X分割(2))

	// Output:
	// [[1 2] [3 4] [5 6] [7 8] [9]]
}

func ExampleArray_PopLeft() {
	array := garray.NewFrom别名([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})

		// 任何以 Pop 开头的函数都会从数组中选取、删除并返回一个元素。 md5:4a1cc3a96b8b68b0

	fmt.Println(array.X出栈左())
	fmt.Println(array.X出栈左多个(2))
	fmt.Println(array.X出栈右())
	fmt.Println(array.X出栈右多个(2))

	// Output:
	// 1 true
	// [2 3]
	// 9 true
	// [7 8]
}

func ExampleArray_PopLefts() {
	array := garray.NewFrom别名([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})

		// 任何以 Pop 开头的函数都会从数组中选取、删除并返回一个元素。 md5:4a1cc3a96b8b68b0

	fmt.Println(array.X出栈左())
	fmt.Println(array.X出栈左多个(2))
	fmt.Println(array.X出栈右())
	fmt.Println(array.X出栈右多个(2))

	// Output:
	// 1 true
	// [2 3]
	// 9 true
	// [7 8]
}

func ExampleArray_PopRight() {
	array := garray.NewFrom别名([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})

		// 任何以 Pop 开头的函数都会从数组中选取、删除并返回一个元素。 md5:4a1cc3a96b8b68b0

	fmt.Println(array.X出栈左())
	fmt.Println(array.X出栈左多个(2))
	fmt.Println(array.X出栈右())
	fmt.Println(array.X出栈右多个(2))

	// Output:
	// 1 true
	// [2 3]
	// 9 true
	// [7 8]
}

func ExampleArray_PopRights() {
	array := garray.NewFrom别名([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})

		// 任何以 Pop 开头的函数都会从数组中选取、删除并返回一个元素。 md5:4a1cc3a96b8b68b0

	fmt.Println(array.X出栈左())
	fmt.Println(array.X出栈左多个(2))
	fmt.Println(array.X出栈右())
	fmt.Println(array.X出栈右多个(2))

	// Output:
	// 1 true
	// [2 3]
	// 9 true
	// [7 8]
}

func ExampleArray_Contains() {
	var array garray.StrArray
	array.Append别名("a")
	fmt.Println(array.X是否存在("a"))
	fmt.Println(array.X是否存在("A"))
	fmt.Println(array.X是否存在并忽略大小写("A"))

	// Output:
	// true
	// false
	// true
}

func ExampleArray_Merge() {
	array1 := garray.NewFrom别名(g.Slice别名{1, 2})
	array2 := garray.NewFrom别名(g.Slice别名{3, 4})
	slice1 := g.Slice别名{5, 6}
	slice2 := []int{7, 8}
	slice3 := []string{"9", "0"}
	fmt.Println(array1.X取切片())
	array1.X合并(array1)
	array1.X合并(array2)
	array1.X合并(slice1)
	array1.X合并(slice2)
	array1.X合并(slice3)
	fmt.Println(array1.X取切片())

	// Output:
	// [1 2]
	// [1 2 1 2 3 4 5 6 7 8 9 0]
}

func ExampleArray_Filter() {
	array1 := garray.NewFrom别名(g.Slice别名{0, 1, 2, nil, "", g.Slice别名{}, "john"})
	array2 := garray.NewFrom别名(g.Slice别名{0, 1, 2, nil, "", g.Slice别名{}, "john"})
	fmt.Printf("%#v\n", array1.X遍历删除(func(index int, value interface{}) bool {
		return empty.IsNil(value)
	}).X取切片())
	fmt.Printf("%#v\n", array2.X遍历删除(func(index int, value interface{}) bool {
		return empty.IsEmpty(value)
	}).X取切片())

	// Output:
	// []interface {}{0, 1, 2, "", []interface {}{}, "john"}
	// []interface {}{1, 2, "john"}
}

func ExampleArray_FilterEmpty() {
	array1 := garray.NewFrom别名(g.Slice别名{0, 1, 2, nil, "", g.Slice别名{}, "john"})
	array2 := garray.NewFrom别名(g.Slice别名{0, 1, 2, nil, "", g.Slice别名{}, "john"})
	fmt.Printf("%#v\n", array1.X删除所有nil().X取切片())
	fmt.Printf("%#v\n", array2.X删除所有空值().X取切片())

	// Output:
	// []interface {}{0, 1, 2, "", []interface {}{}, "john"}
	// []interface {}{1, 2, "john"}
}

func ExampleArray_FilterNil() {
	array1 := garray.NewFrom别名(g.Slice别名{0, 1, 2, nil, "", g.Slice别名{}, "john"})
	array2 := garray.NewFrom别名(g.Slice别名{0, 1, 2, nil, "", g.Slice别名{}, "john"})
	fmt.Printf("%#v\n", array1.X删除所有nil().X取切片())
	fmt.Printf("%#v\n", array2.X删除所有空值().X取切片())

	// Output:
	// []interface {}{0, 1, 2, "", []interface {}{}, "john"}
	// []interface {}{1, 2, "john"}
}
