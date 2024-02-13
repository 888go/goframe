// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 链表类_test

import (
	"container/list"
	"fmt"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/frame/g"
)

func ExampleNew() {
	n := 10
	l := 链表类.New()
	for i := 0; i < n; i++ {
		l.PushBack(i)
	}

	fmt.Println(l.Len())
	fmt.Println(l)
	fmt.Println(l.FrontAll())
	fmt.Println(l.BackAll())

	for i := 0; i < n; i++ {
		fmt.Print(l.PopFront())
	}

	fmt.Println()
	fmt.Println(l.Len())

	// Output:
	// 10
	// [0,1,2,3,4,5,6,7,8,9]
	// [0 1 2 3 4 5 6 7 8 9]
	// [9 8 7 6 5 4 3 2 1 0]
	// 0123456789
	// 0
}

func ExampleNewFrom() {
	n := 10
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 10, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)
	fmt.Println(l.FrontAll())
	fmt.Println(l.BackAll())

	for i := 0; i < n; i++ {
		fmt.Print(l.PopFront())
	}

	fmt.Println()
	fmt.Println(l.Len())

	// Output:
	// 10
	// [1,2,3,4,5,6,7,8,9,10]
	// [1 2 3 4 5 6 7 8 9 10]
	// [10 9 8 7 6 5 4 3 2 1]
	// 12345678910
	// 0
}

func ExampleList_PushFront() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)

	l.PushFront(0)

	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 6
	// [0,1,2,3,4,5]
}

func ExampleList_PushBack() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)

	l.PushBack(6)

	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 6
	// [1,2,3,4,5,6]
}

func ExampleList_PushFronts() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)

	l.PushFronts(g.Slice别名{0, -1, -2, -3, -4})

	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 10
	// [-4,-3,-2,-1,0,1,2,3,4,5]
}

func ExampleList_PushBacks() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)

	l.PushBacks(g.Slice别名{6, 7, 8, 9, 10})

	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 10
	// [1,2,3,4,5,6,7,8,9,10]
}

func ExampleList_PopBack() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)
	fmt.Println(l.PopBack())
	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 5
	// 4
	// [1,2,3,4]
}

func ExampleList_PopFront() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)
	fmt.Println(l.PopFront())
	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 1
	// 4
	// [2,3,4,5]
}

func ExampleList_PopBacks() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)
	fmt.Println(l.PopBacks(2))
	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// [5 4]
	// 3
	// [1,2,3]
}

func ExampleList_PopFronts() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)
	fmt.Println(l.PopFronts(2))
	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// [1 2]
	// 3
	// [3,4,5]
}

func ExampleList_PopBackAll() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)
	fmt.Println(l.PopBackAll())
	fmt.Println(l.Len())

	// Output:
	// 5
	// [1,2,3,4,5]
	// [5 4 3 2 1]
	// 0
}

func ExampleList_PopFrontAll() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)
	fmt.Println(l.PopFrontAll())
	fmt.Println(l.Len())

	// Output:
	// 5
	// [1,2,3,4,5]
	// [1 2 3 4 5]
	// 0
}

func ExampleList_FrontAll() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l)
	fmt.Println(l.FrontAll())

	// Output:
	// [1,2,3,4,5]
	// [1 2 3 4 5]
}

func ExampleList_BackAll() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l)
	fmt.Println(l.BackAll())

	// Output:
	// [1,2,3,4,5]
	// [5 4 3 2 1]
}

func ExampleList_FrontValue() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l)
	fmt.Println(l.FrontValue())

	// Output:
	// [1,2,3,4,5]
	// 1
}

func ExampleList_BackValue() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l)
	fmt.Println(l.BackValue())

	// Output:
	// [1,2,3,4,5]
	// 5
}

func ExampleList_Front() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Front().Value)
	fmt.Println(l)

	e := l.Front()
	l.InsertBefore(e, 0)
	l.InsertAfter(e, "a")

	fmt.Println(l)

	// Output:
	// 1
	// [1,2,3,4,5]
	// [0,1,a,2,3,4,5]
}

func ExampleList_Back() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Back().Value)
	fmt.Println(l)

	e := l.Back()
	l.InsertBefore(e, "a")
	l.InsertAfter(e, 6)

	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// [1,2,3,4,a,5,6]
}

func ExampleList_Len() {
	l := 链表类.NewFrom(g.Slice别名{1, 2, 3, 4, 5})

	fmt.Println(l.Len())

	// Output:
	// 5
}

func ExampleList_Size() {
	l := 链表类.NewFrom(g.Slice别名{1, 2, 3, 4, 5})

	fmt.Println(l.Size())

	// Output:
	// 5
}

func ExampleList_MoveBefore() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Size())
	fmt.Println(l)

	// element of `l`
	e := l.PushBack(6)
	fmt.Println(l.Size())
	fmt.Println(l)

	l.MoveBefore(e, l.Front())

	fmt.Println(l.Size())
	fmt.Println(l)

	// 不是 `l` 的元素
	e = &链表类.Element{Value: 7}
	l.MoveBefore(e, l.Front())

	fmt.Println(l.Size())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 6
	// [1,2,3,4,5,6]
	// 6
	// [6,1,2,3,4,5]
	// 6
	// [6,1,2,3,4,5]
}

func ExampleList_MoveAfter() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Size())
	fmt.Println(l)

	// element of `l`
	e := l.PushFront(0)
	fmt.Println(l.Size())
	fmt.Println(l)

	l.MoveAfter(e, l.Back())

	fmt.Println(l.Size())
	fmt.Println(l)

	// 不是 `l` 的元素
	e = &链表类.Element{Value: -1}
	l.MoveAfter(e, l.Back())

	fmt.Println(l.Size())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 6
	// [0,1,2,3,4,5]
	// 6
	// [1,2,3,4,5,0]
	// 6
	// [1,2,3,4,5,0]
}

func ExampleList_MoveToFront() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Size())
	fmt.Println(l)

	// element of `l`
	l.MoveToFront(l.Back())

	fmt.Println(l.Size())
	fmt.Println(l)

	// 不是 `l` 的元素
	e := &链表类.Element{Value: 6}
	l.MoveToFront(e)

	fmt.Println(l.Size())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 5
	// [5,1,2,3,4]
	// 5
	// [5,1,2,3,4]
}

func ExampleList_MoveToBack() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Size())
	fmt.Println(l)

	// element of `l`
	l.MoveToBack(l.Front())

	fmt.Println(l.Size())
	fmt.Println(l)

	// 不是 `l` 的元素
	e := &链表类.Element{Value: 0}
	l.MoveToBack(e)

	fmt.Println(l.Size())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 5
	// [2,3,4,5,1]
	// 5
	// [2,3,4,5,1]
}

func ExampleList_PushBackList() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Size())
	fmt.Println(l)

	other := 链表类.NewFrom(g.Slice别名{6, 7, 8, 9, 10})

	fmt.Println(other.Size())
	fmt.Println(other)

	l.PushBackList(other)

	fmt.Println(l.Size())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 5
	// [6,7,8,9,10]
	// 10
	// [1,2,3,4,5,6,7,8,9,10]
}

func ExampleList_PushFrontList() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Size())
	fmt.Println(l)

	other := 链表类.NewFrom(g.Slice别名{-4, -3, -2, -1, 0})

	fmt.Println(other.Size())
	fmt.Println(other)

	l.PushFrontList(other)

	fmt.Println(l.Size())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 5
	// [-4,-3,-2,-1,0]
	// 10
	// [-4,-3,-2,-1,0,1,2,3,4,5]
}

func ExampleList_InsertAfter() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)

	l.InsertAfter(l.Front(), "a")
	l.InsertAfter(l.Back(), "b")

	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 7
	// [1,a,2,3,4,5,b]
}

func ExampleList_InsertBefore() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)

	l.InsertBefore(l.Front(), "a")
	l.InsertBefore(l.Back(), "b")

	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 7
	// [a,1,2,3,4,b,5]
}

func ExampleList_Remove() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)

	fmt.Println(l.Remove(l.Front()))
	fmt.Println(l.Remove(l.Back()))

	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 1
	// 5
	// 3
	// [2,3,4]
}

func ExampleList_Removes() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)

	l.Removes([]*链表类.Element{l.Front(), l.Back()})

	fmt.Println(l.Len())
	fmt.Println(l)

	// Output:
	// 5
	// [1,2,3,4,5]
	// 3
	// [2,3,4]
}

func ExampleList_RemoveAll() {
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 5, 1).X取切片())

	fmt.Println(l.Len())
	fmt.Println(l)

	l.RemoveAll()

	fmt.Println(l.Len())

	// Output:
	// 5
	// [1,2,3,4,5]
	// 0
}

func ExampleList_RLockFunc() {
	// 并发安全列表。
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 10, 1).X取切片(), true)
	// 从头部开始迭代读取
	l.RLockFunc(func(list *list.List) {
		length := list.Len()
		if length > 0 {
			for i, e := 0, list.Front(); i < length; i, e = i+1, e.Next() {
				fmt.Print(e.Value)
			}
		}
	})
	fmt.Println()
	// 从尾部开始迭代读取
	l.RLockFunc(func(list *list.List) {
		length := list.Len()
		if length > 0 {
			for i, e := 0, list.Back(); i < length; i, e = i+1, e.Prev() {
				fmt.Print(e.Value)
			}
		}
	})

	fmt.Println()
	// Output:
	// 12345678910
	// 10987654321
}

func ExampleList_IteratorAsc() {
	// 并发安全列表。
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 10, 1).X取切片(), true)
	// 使用IteratorAsc从头部开始进行迭代读取。
	l.IteratorAsc(func(e *链表类.Element) bool {
		fmt.Print(e.Value)
		return true
	})

	// Output:
	// 12345678910
}

func ExampleList_IteratorDesc() {
	// 并发安全列表。
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 10, 1).X取切片(), true)
	// 使用IteratorDesc从尾部开始迭代读取
	l.IteratorDesc(func(e *链表类.Element) bool {
		fmt.Print(e.Value)
		return true
	})
	// Output:
	// 10987654321
}

func ExampleList_LockFunc() {
	// 并发安全列表。
	l := 链表类.NewFrom(数组类.X创建并按范围(1, 10, 1).X取切片(), true)
	// 从头部开始进行迭代写入。
	l.LockFunc(func(list *list.List) {
		length := list.Len()
		if length > 0 {
			for i, e := 0, list.Front(); i < length; i, e = i+1, e.Next() {
				if e.Value == 6 {
					e.Value = "M"
					break
				}
			}
		}
	})
	fmt.Println(l)

	// Output:
	// [1,2,3,4,5,M,7,8,9,10]
}

func ExampleList_Join() {
	var l 链表类.List
	l.PushBacks(g.Slice别名{"a", "b", "c", "d"})

	fmt.Println(l.Join(","))

	// Output:
	// a,b,c,d
}
