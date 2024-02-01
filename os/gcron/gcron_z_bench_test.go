// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcron_test
import (
	"context"
	"testing"
	
	"github.com/888go/goframe/os/gcron"
	)

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcron.Add(ctx, "1 1 1 1 1 1", func(ctx context.Context) {

		})
	}
}
