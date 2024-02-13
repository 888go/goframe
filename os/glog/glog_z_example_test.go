// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类_test

import (
	"context"
	
	"github.com/888go/goframe/frame/g"
)

func ExampleContext() {
	ctx := context.WithValue(context.Background(), "Trace-Id", "123456789")
	g.X日志类().Error(ctx, "runtime error")

	// May Output:
	// 2020-06-08 20:17:03.630 [ERRO] {Trace-Id: 123456789} runtime error
	// Stack:
	// ...
}
