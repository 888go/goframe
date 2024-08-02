// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package uid类_test

import (
	"testing"

	guid "github.com/888go/goframe/util/guid"
)

func Benchmark_S(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			guid.S()
		}
	})
}

func Benchmark_S_Data_1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			guid.S([]byte("123"))
		}
	})
}

func Benchmark_S_Data_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			guid.S([]byte("123"), []byte("456"))
		}
	})
}
