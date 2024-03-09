// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 队列类_test

import (
	"context"
	"fmt"
	"time"
	
	"github.com/888go/goframe/gqueue"
	"github.com/gogf/gf/v2/os/gtimer"
)

func ExampleNew() {
	n := 10
	q := 队列类.X创建()

	// Producer
	for i := 0; i < n; i++ {
		q.X入栈(i)
	}

	// 在三秒钟后关闭队列。
	gtimer.SetTimeout(context.Background(), time.Second*3, func(ctx context.Context) {
		q.X关闭()
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
	q := 队列类.X创建()

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
	q := 队列类.X创建()

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
	q := 队列类.X创建()

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
	q := 队列类.X创建()

	q.X入栈(1)
	q.X入栈(2)

	fmt.Println(q.X取长度())

	// May Output:
	// 2
}

func ExampleQueue_Size() {
	q := 队列类.X创建()

	q.X入栈(1)
	q.X入栈(2)

	// Size 是 Len 的别名。
	fmt.Println(q.Size弃用())

	// May Output:
	// 2
}
