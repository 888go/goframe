// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gqueue_test

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/os/gtimer"
)

func ExampleNew() {
	n := 10
	q := gqueue.New()

	// Producer
	for i := 0; i < n; i++ {
		q.Push(i)
	}

	// 三秒后关闭队列。. md5:02742be1c1ceef32
	gtimer.SetTimeout(context.Background(), time.Second*3, func(ctx context.Context) {
		q.Close()
	})

// 消费者持续读取队列中的数据。
// 如果队列中没有数据，它将阻塞等待。
// 队列的读取是通过队列对象暴露的
// queue.C 属性以及使用 selectIO 多路复用语法来实现的。
// 示例代码如下：
// for {
//     select {
//         case v := <-queue.C:
//             if v != nil {
//                 fmt.Println(v)
//             } else {
//                 return
//             }
//     }
// }
// 在这个例子中，`queue.C` 是一个通道，消费者通过它接收队列中的元素。
// 当有值可读时，`v` 会被赋值并打印；若收到的是 `nil`，则表示某种结束条件，消费者退出循环。
// md5:4bb8650995a22499
	for {
		if v := q.Pop(); v != nil {
			fmt.Print(v)
		} else {
			break
		}
	}

	// Output:
	// 0123456789
}

func ExampleQueue_Push() {
	q := gqueue.New()

	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	// Output:
	// 0
	// 1
	// 2
}

func ExampleQueue_Pop() {
	q := gqueue.New()

	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	fmt.Println(q.Pop())
	q.Close()
	fmt.Println(q.Pop())

	// Output:
	// 0
	// <nil>
}

func ExampleQueue_Close() {
	q := gqueue.New()

	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	time.Sleep(time.Millisecond)
	q.Close()

	fmt.Println(q.Len())
	fmt.Println(q.Pop())

	// May Output:
	// 0
	// <nil>
}

func ExampleQueue_Len() {
	q := gqueue.New()

	q.Push(1)
	q.Push(2)

	fmt.Println(q.Len())

	// May Output:
	// 2
}

func ExampleQueue_Size() {
	q := gqueue.New()

	q.Push(1)
	q.Push(2)

	// Size is alias of Len.
	fmt.Println(q.Size())

	// May Output:
	// 2
}
