// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 协程类_test

import (
	"context"
	"testing"
	
	"github.com/888go/goframe/os/grpool"
)

var (
	ctx = context.TODO()
	n   = 500000
)

func increment(ctx context.Context) {
	for i := 0; i < 1000000; i++ {
	}
}

func BenchmarkGrpool_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		协程类.Add(ctx, increment)
	}
}

func BenchmarkGoroutine_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go increment(ctx)
	}
}

func BenchmarkGrpool2(b *testing.B) {
	b.N = n
	for i := 0; i < b.N; i++ {
		协程类.Add(ctx, increment)
	}
}

func BenchmarkGoroutine2(b *testing.B) {
	b.N = n
	for i := 0; i < b.N; i++ {
		go increment(ctx)
	}
}
