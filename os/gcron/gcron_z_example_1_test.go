// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package 定时cron类_test

import (
	"context"
	"time"

	gcron "github.com/888go/goframe/os/gcron"
	glog "github.com/888go/goframe/os/glog"
)

func ExampleCronAddSingleton() {
	gcron.AddSingleton(ctx, "* * * * * *", func(ctx context.Context) {
		glog.Print(context.TODO(), "doing")
		time.Sleep(2 * time.Second)
	})
	select {}
}
