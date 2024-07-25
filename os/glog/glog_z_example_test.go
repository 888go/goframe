// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package glog_test

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func ExampleContext() {
	ctx := context.WithValue(context.Background(), "Trace-Id", "123456789")
	g.Log().Error(ctx, "runtime error")

	// May Output:
	// 2020-06-08 20:17:03.630 [ERRO] {Trace-Id: 123456789} runtime error
	// Stack:
	// ...
}
