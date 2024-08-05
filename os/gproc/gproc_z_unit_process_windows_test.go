// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

//---build---//go:build windows

package gproc_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_ProcessRun(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		binary := gproc.SearchBinary("go")
		t.AssertNE(binary, "")
		var command = gproc.NewProcess(binary, nil)

		testPath := gtest.DataPath("gobuild")
		filename := filepath.Join(testPath, "main.go")
		output := filepath.Join(testPath, "main.exe")

		command.Args = append(command.Args, "build")
		command.Args = append(command.Args, `-ldflags="-X 'main.TestString=\"test string\"'"`)
		command.Args = append(command.Args, "-o", output)
		command.Args = append(command.Args, filename)

		err := command.Run(gctx.GetInitCtx())
		t.AssertNil(err)

		exists := gfile.Exists(output)
		t.Assert(exists, true)
		defer gfile.Remove(output)

		runCmd := gproc.NewProcess(output, nil)
		var buf strings.Builder
		runCmd.Stdout = &buf
		runCmd.Stderr = &buf
		err = runCmd.Run(gctx.GetInitCtx())
		t.Assert(err, nil)
		t.Assert(buf.String(), `"test string"`)
	})

	gtest.C(t, func(t *gtest.T) {
		binary := gproc.SearchBinary("go")
		t.AssertNE(binary, "")
				// NewProcess(path, args) 参数说明：path：最好不要包含空格. md5:0a5ad3abdc1b7a35
		var command = gproc.NewProcess(binary, nil)

		testPath := gtest.DataPath("gobuild")
		filename := filepath.Join(testPath, "main.go")
		output := filepath.Join(testPath, "main.exe")

		command.Args = append(command.Args, "build")
		command.Args = append(command.Args, `-ldflags="-s -w"`)
		command.Args = append(command.Args, "-o", output)
		command.Args = append(command.Args, filename)

		err := command.Run(gctx.GetInitCtx())
		t.AssertNil(err)

		exists := gfile.Exists(output)
		t.Assert(exists, true)

		defer gfile.Remove(output)
	})
}
