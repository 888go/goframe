// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 使用 go test 命令测试所有.go文件，启用基准测试（-bench=".*"），并显示内存使用情况统计信息（-benchmem）

package 进程类_test

import (
	"testing"
	
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/gproc"
	"github.com/888go/goframe/test/gtest"
)

func Test_ShellExec(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s, err := 进程类.ShellExec(上下文类.X创建(), `echo 123`)
		t.AssertNil(err)
		t.Assert(s, "123\r\n")//2024-01-14 此处替换成win平台的换行符, 以免在win平台单元测试不过
	})
	// error
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := 进程类.ShellExec(上下文类.X创建(), `NoneExistCommandCall`)
		t.AssertNE(err, nil)
	})
}
