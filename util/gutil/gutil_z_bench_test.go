// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gutil

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
		TryCatch(ctx, func(ctx context.Context) {

		}, func(ctx context.Context, err error) {

		})
	}
}

func Benchmark_MapMergeCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MapMergeCopy(m1, m2)
	}
}
