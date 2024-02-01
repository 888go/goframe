// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gqueue_test
import (
	"context"
	"fmt"
	"time"
	
	"github.com/888go/goframe/container/gqueue"
	"github.com/888go/goframe/os/gtimer"
	)

func ExampleNew() {
	n := 10
	q := gqueue.New()

	// Producer
	for i := 0; i < n; i++ {
		q.Push(i)
	}

	// 在三秒钟后关闭队列。
	gtimer.SetTimeout(context.Background(), time.Second*3, func(ctx context.Context) {
		q.Close()
	})

// 消费者持续读取队列中的数据。
// 如果队列中没有数据，则会阻塞等待。
// 通过队列对象公开的queue.C属性来读取队列，
// 并使用selectIO多路复用语法进行操作。
// 示例代码如下：
// for {
//     select {
//         case v := <-queue.C: // 从队列channel接收数据
//             if v != nil { // 若接收到的数据不为空
//                 fmt.Println(v) // 输出接收到的数据
//             } else { // 若接收到的数据为空
//                 return // 则退出循环（或程序）
//             }
//     }
// }
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

	// Size 是 Len 的别名。
	fmt.Println(q.Size())

	// May Output:
	// 2
}
