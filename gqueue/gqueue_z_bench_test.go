// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gqueue_test

import (
	"testing"
	
	"github.com/888go/goframe/gqueue"
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
