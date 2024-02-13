// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

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
