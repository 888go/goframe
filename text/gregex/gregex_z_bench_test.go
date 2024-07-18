// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package gregex_test//bm:正则类_test

import (
	"regexp"
	"testing"

	"github.com/gogf/gf/v2/text/gregex"
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
