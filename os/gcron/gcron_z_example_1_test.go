// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcron_test

import (
	"context"
	"time"
	
	"github.com/888go/goframe/os/gcron"
	"github.com/888go/goframe/os/glog"
)

func ExampleCronAddSingleton() {
	gcron.AddSingleton(ctx, "* * * * * *", func(ctx context.Context) {
		glog.Print(context.TODO(), "doing")
		time.Sleep(2 * time.Second)
	})
	select {}
}
