// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package gqueue_test

import (
	"testing"

	"github.com/gogf/gf/v2/container/gqueue"
)

var bn = 20000000

var length = 1000000

var qstatic = gqueue.New(length)

var qdynamic = gqueue.New()

var cany = make(chan interface{}, length)

func Benchmark_Gqueue_StaticPushAndPop(b *testing.B) {
	b.N = bn
	for i := 0; i < b.N; i++ {
		qstatic.Push(i)
		qstatic.Pop()
	}
}

func Benchmark_Gqueue_DynamicPush(b *testing.B) {
	b.N = bn
	for i := 0; i < b.N; i++ {
		qdynamic.Push(i)
	}
}

func Benchmark_Gqueue_DynamicPop(b *testing.B) {
	b.N = bn
	for i := 0; i < b.N; i++ {
		qdynamic.Pop()
	}
}

func Benchmark_Channel_PushAndPop(b *testing.B) {
	b.N = bn
	for i := 0; i < b.N; i++ {
		cany <- i
		<-cany
	}
}
