// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package 对象复用类_test

import (
	"sync"
	"testing"
	"time"

	gpool "github.com/888go/goframe/container/gpool"
)

var pool = gpool.X创建(time.Hour, nil)

var syncp = sync.Pool{}

func BenchmarkGPoolPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.X入栈(i)
	}
}

func BenchmarkGPoolGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.X出栈()
	}
}

func BenchmarkSyncPoolPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		syncp.Put(i)
	}
}

func BenchmarkSyncPoolGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		syncp.Get()
	}
}
