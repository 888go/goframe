// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

//go:build windows

package 进程类_test

import (
	"fmt"
	"path/filepath"
	"testing"

	gctx "github.com/888go/goframe/os/gctx"
	gfile "github.com/888go/goframe/os/gfile"
	gproc "github.com/888go/goframe/os/gproc"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_ShellExec_GoBuild_Windows(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testPath := gtest.DataPath("gobuild")
		filename := filepath.Join(testPath, "main.go")
		output := filepath.Join(testPath, "main.exe")
		cmd := fmt.Sprintf(`go build -ldflags="-s -w" -o %s  %s`, output, filename)

		err := gproc.ShellRun(gctx.X创建(), cmd)
		t.Assert(err, nil)

		exists := gfile.X是否存在(output)
		t.Assert(exists, true)

		defer gfile.X删除(output)
	})

	gtest.C(t, func(t *gtest.T) {
		testPath := gtest.DataPath("gobuild")
		filename := filepath.Join(testPath, "main.go")
		output := filepath.Join(testPath, "main.exe")
		cmd := fmt.Sprintf(`go build -ldflags="-X 'main.TestString=\"test string\"'" -o %s %s`, output, filename)

		err := gproc.ShellRun(gctx.X创建(), cmd)
		t.Assert(err, nil)

		exists := gfile.X是否存在(output)
		t.Assert(exists, true)
		defer gfile.X删除(output)

		result, err := gproc.ShellExec(gctx.X创建(), output)
		t.Assert(err, nil)
		t.Assert(result, `"test string"`)
	})

}

func Test_ShellExec_SpaceDir_Windows(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testPath := gtest.DataPath("shellexec")
		filename := filepath.Join(testPath, "main.go")
						// 使用go build命令，生成名为test.exe的可执行文件，源代码为main.go. md5:3a438d2ac0c99590
		cmd := fmt.Sprintf(`go build -o test.exe %s`, filename)
		r, err := gproc.ShellExec(gctx.X创建(), cmd)
		t.AssertNil(err)
		t.Assert(r, "")

		exists := gfile.X是否存在(filename)
		t.Assert(exists, true)

		outputDir := filepath.Join(testPath, "testdir")
		output := filepath.Join(outputDir, "test.exe")
		err = gfile.X移动("test.exe", output)
		t.AssertNil(err)
		defer gfile.X删除(output)

		expectContent := "123"
		testOutput := filepath.Join(testPath, "space dir", "test.txt")
		cmd = fmt.Sprintf(`%s -c %s -o "%s"`, output, expectContent, testOutput)
		r, err = gproc.ShellExec(gctx.X创建(), cmd)
		t.AssertNil(err)

		exists = gfile.X是否存在(testOutput)
		t.Assert(exists, true)
		defer gfile.X删除(testOutput)

		contents := gfile.X读文本(testOutput)
		t.Assert(contents, expectContent)
	})
	gtest.C(t, func(t *gtest.T) {
		testPath := gtest.DataPath("shellexec")
		filename := filepath.Join(testPath, "main.go")
						// 使用go build命令，生成名为test.exe的可执行文件，源代码为main.go. md5:3a438d2ac0c99590
		cmd := fmt.Sprintf(`go build -o test.exe %s`, filename)
		r, err := gproc.ShellExec(gctx.X创建(), cmd)
		t.AssertNil(err)
		t.Assert(r, "")

		exists := gfile.X是否存在(filename)
		t.Assert(exists, true)

		outputDir := filepath.Join(testPath, "space dir")
		output := filepath.Join(outputDir, "test.exe")
		err = gfile.X移动("test.exe", output)
		t.AssertNil(err)
		defer gfile.X删除(output)

		expectContent := "123"
		testOutput := filepath.Join(testPath, "testdir", "test.txt")
		cmd = fmt.Sprintf(`"%s" -c %s -o %s`, output, expectContent, testOutput)
		r, err = gproc.ShellExec(gctx.X创建(), cmd)
		t.AssertNil(err)

		exists = gfile.X是否存在(testOutput)
		t.Assert(exists, true)
		defer gfile.X删除(testOutput)

		contents := gfile.X读文本(testOutput)
		t.Assert(contents, expectContent)

	})
}
