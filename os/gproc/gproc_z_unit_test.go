// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 进程类_test

import (
	"testing"

	gctx "github.com/888go/goframe/os/gctx"
	gproc "github.com/888go/goframe/os/gproc"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_ShellExec(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s, err := gproc.ShellExec(gctx.New(), `echo 123`)
		t.AssertNil(err)
		t.Assert(s, "123\r\n")
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		_, err := gproc.ShellExec(gctx.New(), `NoneExistCommandCall`)
		t.AssertNE(err, nil)
	})
}
