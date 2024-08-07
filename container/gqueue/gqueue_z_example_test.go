// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 队列类_test

import (
	"context"
	"fmt"
	"time"

	gqueue "github.com/888go/goframe/container/gqueue"
	gtimer "github.com/888go/goframe/os/gtimer"
)

func ExampleNew() {
	n := 10
	q := gqueue.X创建()

	// Producer
	for i := 0; i < n; i++ {
		q.X入栈(i)
	}

		// 三秒后关闭队列。 md5:02742be1c1ceef32
	gtimer.SetTimeout别名(context.Background(), time.Second*3, func(ctx context.Context) {
		q.X关闭()
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
		if v := q.X出栈(); v != nil {
			fmt.Print(v)
		} else {
			break
		}
	}

	// Output:
	// 0123456789
}

func ExampleQueue_Push() {
	q := gqueue.X创建()

	for i := 0; i < 10; i++ {
		q.X入栈(i)
	}

	fmt.Println(q.X出栈())
	fmt.Println(q.X出栈())
	fmt.Println(q.X出栈())

	// Output:
	// 0
	// 1
	// 2
}

func ExampleQueue_Pop() {
	q := gqueue.X创建()

	for i := 0; i < 10; i++ {
		q.X入栈(i)
	}

	fmt.Println(q.X出栈())
	q.X关闭()
	fmt.Println(q.X出栈())

	// Output:
	// 0
	// <nil>
}

func ExampleQueue_Close() {
	q := gqueue.X创建()

	for i := 0; i < 10; i++ {
		q.X入栈(i)
	}

	time.Sleep(time.Millisecond)
	q.X关闭()

	fmt.Println(q.X取长度())
	fmt.Println(q.X出栈())

	// May Output:
	// 0
	// <nil>
}

func ExampleQueue_Len() {
	q := gqueue.X创建()

	q.X入栈(1)
	q.X入栈(2)

	fmt.Println(q.X取长度())

	// May Output:
	// 2
}

func ExampleQueue_Size() {
	q := gqueue.X创建()

	q.X入栈(1)
	q.X入栈(2)

	// Size is alias of Len.
	fmt.Println(q.Size弃用())

	// May Output:
	// 2
}
