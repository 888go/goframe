// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package utils_test

import (
	"regexp"
	"testing"
	
	"github.com/888go/goframe/internal/utils"
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
