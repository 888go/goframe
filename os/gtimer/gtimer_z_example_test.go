// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package 定时类_test

import (
	"context"
	"fmt"
	"time"

	gtimer "github.com/888go/goframe/os/gtimer"
)

func ExampleAdd() {
	var (
		ctx      = context.Background()
		now      = time.Now()
		interval = 1400 * time.Millisecond
	)
	gtimer.Add(ctx, interval, func(ctx context.Context) {
		fmt.Println(time.Now(), time.Duration(time.Now().UnixNano()-now.UnixNano()))
		now = time.Now()
	})

	select {}
}
