// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"strings"
	"testing"
)

func Benchmark_TrimRightCharWithStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		path := "//////////"
		strings.TrimRight(path, "/")
	}
}

func Benchmark_TrimRightCharWithSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		path := "//////////"
		for len(path) > 0 && path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}
	}
}

func Benchmark_TrimRightCharWithSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		path := "//////////"
		for {
			if length := len(path); length > 0 && path[length-1] == '/' {
				path = path[:length-1]
			} else {
				break
			}
		}
	}
}
