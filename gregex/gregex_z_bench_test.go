// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package gregex_test

import (
	"regexp"
	"testing"
	
	"github.com/888go/goframe/gregex"
)

var pattern = `(\w+).+\-\-\s*(.+)`

var src = `GF is best! -- John`

func Benchmark_GF_IsMatchString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gregex.IsMatchString(pattern, src)
	}
}

func Benchmark_GF_MatchString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gregex.MatchString(pattern, src)
	}
}

func Benchmark_Compile(b *testing.B) {
	var wcdRegexp = regexp.MustCompile(pattern)
	for i := 0; i < b.N; i++ {
		wcdRegexp.MatchString(src)
	}
}

func Benchmark_Compile_Actual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wcdRegexp := regexp.MustCompile(pattern)
		wcdRegexp.MatchString(src)
	}
}
