// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 循环链表类_test

import (
	"fmt"
	
	"github.com/888go/goframe/container/gring"
)

func ExampleNew() {
	// 非并发安全
	循环链表类.New(10)

	// Concurrent safety
	循环链表类.New(10, true)

	// Output:
}

func ExampleRing_Val() {
	r := 循环链表类.New(10)
	r.X设置值(1)
	fmt.Println("Val:", r.X取值())

	r.Next().X设置值("GoFrame")
	fmt.Println("Val:", r.X取值())

	// Output:
	// Val: 1
	// Val: GoFrame
}

func ExampleRing_Len() {
	r1 := 循环链表类.New(10)
	for i := 0; i < 5; i++ {
		r1.X设置值(i).Next()
	}
	fmt.Println("Len:", r1.Len())

	r2 := 循环链表类.New(10, true)
	for i := 0; i < 10; i++ {
		r2.X设置值(i).Next()
	}
	fmt.Println("Len:", r2.Len())

	// Output:
	// Len: 5
	// Len: 10
}

func ExampleRing_Cap() {
	r1 := 循环链表类.New(10)
	for i := 0; i < 5; i++ {
		r1.X设置值(i).Next()
	}
	fmt.Println("Cap:", r1.Cap())

	r2 := 循环链表类.New(10, true)
	for i := 0; i < 10; i++ {
		r2.X设置值(i).Next()
	}
	fmt.Println("Cap:", r2.Cap())

	// Output:
	// Cap: 10
	// Cap: 10
}

func ExampleRing_Set() {
	r := 循环链表类.New(10)
	r.X设置值(1)
	fmt.Println("Val:", r.X取值())

	r.Next().X设置值("GoFrame")
	fmt.Println("Val:", r.X取值())

	// Output:
	// Val: 1
	// Val: GoFrame
}

func ExampleRing_Put() {
	r := 循环链表类.New(10)
	r.Put(1)
	fmt.Println("Val:", r.X取值())
	fmt.Println("Val:", r.Prev().X取值())

	// Output:
	// Val: <nil>
	// Val: 1
}

func ExampleRing_Move() {
	r := 循环链表类.New(10)
	for i := 0; i < 10; i++ {
		r.X设置值(i).Next()
	}
	// ring at Pos 0
	fmt.Println("CurVal:", r.X取值())

	r.Move(5)

	// ring at Pos 5
	fmt.Println("CurVal:", r.X取值())

	// Output:
	// CurVal: 0
	// CurVal: 5
}

func ExampleRing_Prev() {
	r := 循环链表类.New(10)
	for i := 0; i < 5; i++ {
		r.X设置值(i).Next()
	}

	fmt.Println("Prev:", r.Prev().X取值())
	fmt.Println("Prev:", r.Prev().X取值())

	// Output:
	// Prev: 4
	// Prev: 3
}

func ExampleRing_Next() {
	r := 循环链表类.New(10)
	for i := 5; i > 0; i-- {
		r.X设置值(i).Prev()
	}

	fmt.Println("Prev:", r.Next().X取值())
	fmt.Println("Prev:", r.Next().X取值())

	// Output:
	// Prev: 1
	// Prev: 2
}

func ExampleRing_Link_Common() {
	r := 循环链表类.New(10)
	for i := 0; i < 5; i++ {
		r.X设置值(i).Next()
	}

	s := 循环链表类.New(10)
	for i := 0; i < 10; i++ {
		val := i + 5
		s.X设置值(val).Next()
	}

	r.Link(s) // 将环形链表s链接到环形链表r

	fmt.Println("Len:", r.Len())
	fmt.Println("Cap:", r.Cap())
	fmt.Println(r.SlicePrev())
	fmt.Println(r.SliceNext())

	// Output:
	// Len: 15
	// Cap: 20
	// [4 3 2 1 0]
	// [5 6 7 8 9 10 11 12 13 14]
}

func ExampleRing_Link_SameRing() {
	r := 循环链表类.New(10)
	for i := 0; i < 5; i++ {
		r.X设置值(i).Next()
	}

	same_r := r.Link(r.Prev())

	fmt.Println("Len:", same_r.Len())
	fmt.Println("Cap:", same_r.Cap())
	fmt.Println(same_r.SlicePrev())
	fmt.Println(same_r.SliceNext())

	// Output:
	// Len: 1
	// Cap: 1
	// [4]
	// [4]
}

func ExampleRing_Unlink() {
	r := 循环链表类.New(10)
	for i := 0; i < 10; i++ {
		r.X设置值(i).Next()
	}

	fmt.Println("Before Unlink, Len:", r.Len())
	fmt.Println("Before Unlink, Cap:", r.Cap())
	fmt.Println("Before Unlink, ", r.SlicePrev())
	fmt.Println("Before Unlink, ", r.SliceNext())

	r.Unlink(7)

	fmt.Println("After Unlink, Len:", r.Len())
	fmt.Println("After Unlink, Cap:", r.Cap())
	fmt.Println("After Unlink, ", r.SlicePrev())
	fmt.Println("After Unlink, ", r.SliceNext())

	// Output:
	// Before Unlink, Len: 10
	// Before Unlink, Cap: 10
	// Before Unlink,  [0 9 8 7 6 5 4 3 2 1]
	// Before Unlink,  [0 1 2 3 4 5 6 7 8 9]
	// After Unlink, Len: 3
	// After Unlink, Cap: 3
	// After Unlink,  [0 9 8]
	// After Unlink,  [0 8 9]
}

func ExampleRing_RLockIteratorNext() {
	r := 循环链表类.New(10)
	for i := 0; i < 10; i++ {
		r.X设置值(i).Next()
	}

	r.RLockIteratorNext(func(value interface{}) bool {
		if value.(int) < 5 {
			fmt.Println("IteratorNext Success, Value:", value)
			return true
		}

		return false
	})

	// Output:
	// IteratorNext Success, Value: 0
	// IteratorNext Success, Value: 1
	// IteratorNext Success, Value: 2
	// IteratorNext Success, Value: 3
	// IteratorNext Success, Value: 4
}

func ExampleRing_RLockIteratorPrev() {
	r := 循环链表类.New(10)
	for i := 0; i < 10; i++ {
		r.X设置值(i).Next()
	}

	// move r to pos 9
	r.Prev()

	r.RLockIteratorPrev(func(value interface{}) bool {
		if value.(int) >= 5 {
			fmt.Println("IteratorPrev Success, Value:", value)
			return true
		}

		return false
	})

	// Output:
	// IteratorPrev Success, Value: 9
	// IteratorPrev Success, Value: 8
	// IteratorPrev Success, Value: 7
	// IteratorPrev Success, Value: 6
	// IteratorPrev Success, Value: 5
}

func ExampleRing_SliceNext() {
	r := 循环链表类.New(10)
	for i := 0; i < 10; i++ {
		r.X设置值(i).Next()
	}

	fmt.Println(r.SliceNext())

	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}

func ExampleRing_SlicePrev() {
	r := 循环链表类.New(10)
	for i := 0; i < 10; i++ {
		r.X设置值(i).Next()
	}

	fmt.Println(r.SlicePrev())

	// Output:
	// [0 9 8 7 6 5 4 3 2 1]
}
