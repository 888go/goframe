	// 版权归GoFrame作者(https:	//goframe.org)所有。保留所有权利。
	//
	// 本源代码形式受MIT许可证条款约束。
	// 如果未随本文件一同分发MIT许可证副本，
	// 您可以在https:	//github.com/gogf/gf处获取。
	// md5:a9832f33b234e3f3

package utils_test

import (
	"regexp"
	"testing"

	"github.com/gogf/gf/v2/internal/utils"
)

var (
	replaceCharReg, _ = regexp.Compile(`[\-\.\_\s]+`)
)

func Benchmark_RemoveSymbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.RemoveSymbols(`-a-b._a c1!@#$%^&*()_+:";'.,'01`)
	}
}

func Benchmark_RegularReplaceChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replaceCharReg.ReplaceAllString(`-a-b._a c1!@#$%^&*()_+:";'.,'01`, "")
	}
}
