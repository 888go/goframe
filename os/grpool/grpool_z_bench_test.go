// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package 协程类_test

import (
	"context"
	"testing"

	grpool "github.com/888go/goframe/os/grpool"
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
		grpool.Add(ctx, increment)
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
		grpool.Add(ctx, increment)
	}
}

func BenchmarkGoroutine2(b *testing.B) {
	b.N = n
	for i := 0; i < b.N; i++ {
		go increment(ctx)
	}
}
