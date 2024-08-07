// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 工具类

import (
	"context"
	"testing"
)

var (
	m1 = map[string]interface{}{
		"k1": "v1",
	}
	m2 = map[string]interface{}{
		"k2": "v2",
	}
)

func Benchmark_TryCatch(b *testing.B) {
	ctx := context.TODO()
	for i := 0; i < b.N; i++ {
		X异常捕捉并带异常处理(ctx, func(ctx context.Context) {

		}, func(ctx context.Context, err error) {

		})
	}
}

func Benchmark_MapMergeCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MapMergeCopy(m1, m2)
	}
}
