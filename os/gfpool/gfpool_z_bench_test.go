// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件指针池类

import (
	"os"
	"testing"
)

func Benchmark_OS_Open_Close_ALLFlags(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.OpenFile("/tmp/bench-test", os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}

func Benchmark_GFPool_Open_Close_ALLFlags(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := Open("/tmp/bench-test", os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}

func Benchmark_OS_Open_Close_RDWR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.OpenFile("/tmp/bench-test", os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}

func Benchmark_GFPool_Open_Close_RDWR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := Open("/tmp/bench-test", os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}

func Benchmark_OS_Open_Close_RDONLY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.OpenFile("/tmp/bench-test", os.O_RDONLY, 0666)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}

func Benchmark_GFPool_Open_Close_RDONLY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := Open("/tmp/bench-test", os.O_RDONLY, 0666)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}

func Benchmark_Stat(b *testing.B) {
	f, err := os.Create("/tmp/bench-test-stat")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := 0; i < b.N; i++ {
		f.Stat()
	}
}
