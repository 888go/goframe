// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 定时类_test

import (
	"context"
	"fmt"
	"time"
	
	"github.com/888go/goframe/os/gtimer"
)

func ExampleAdd() {
	var (
		ctx      = context.Background()
		now      = time.Now()
		interval = 1400 * time.Millisecond
	)
	定时类.X加入循环任务(ctx, interval, func(ctx context.Context) {
		fmt.Println(time.Now(), time.Duration(time.Now().UnixNano()-now.UnixNano()))
		now = time.Now()
	})

	select {}
}
