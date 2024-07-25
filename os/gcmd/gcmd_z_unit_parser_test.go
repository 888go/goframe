// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package gcmd_test

import (
	"os"
	"testing"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Parse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		os.Args = []string{"gf", "--force", "remove", "-fq", "-p=www", "path", "-n", "root"}
		p, err := gcmd.Parse(map[string]bool{
			"n, name":   true,
			"p, prefix": true,
			"f,force":   false,
			"q,quiet":   false,
		})
		t.AssertNil(err)
		t.Assert(len(p.GetArgAll()), 3)
		t.Assert(p.GetArg(0), "gf")
		t.Assert(p.GetArg(1), "remove")
		t.Assert(p.GetArg(2), "path")
		t.Assert(p.GetArg(2).String(), "path")

		t.Assert(len(p.GetOptAll()), 8)
		t.Assert(p.GetOpt("n"), "root")
		t.Assert(p.GetOpt("name"), "root")
		t.Assert(p.GetOpt("p"), "www")
		t.Assert(p.GetOpt("prefix"), "www")
		t.Assert(p.GetOpt("prefix").String(), "www")

		t.Assert(p.GetOpt("n") != nil, true)
		t.Assert(p.GetOpt("name") != nil, true)
		t.Assert(p.GetOpt("p") != nil, true)
		t.Assert(p.GetOpt("prefix") != nil, true)
		t.Assert(p.GetOpt("f") != nil, true)
		t.Assert(p.GetOpt("force") != nil, true)
		t.Assert(p.GetOpt("q") != nil, true)
		t.Assert(p.GetOpt("quiet") != nil, true)
		t.Assert(p.GetOpt("none") != nil, false)

		_, err = p.MarshalJSON()
		t.AssertNil(err)
	})
}

func Test_ParseArgs(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p, err := gcmd.ParseArgs(
			[]string{"gf", "--force", "remove", "-fq", "-p=www", "path", "-n", "root"},
			map[string]bool{
				"n, name":   true,
				"p, prefix": true,
				"f,force":   false,
				"q,quiet":   false,
			})
		t.AssertNil(err)
		t.Assert(len(p.GetArgAll()), 3)
		t.Assert(p.GetArg(0), "gf")
		t.Assert(p.GetArg(1), "remove")
		t.Assert(p.GetArg(2), "path")
		t.Assert(p.GetArg(2).String(), "path")

		t.Assert(len(p.GetOptAll()), 8)
		t.Assert(p.GetOpt("n"), "root")
		t.Assert(p.GetOpt("name"), "root")
		t.Assert(p.GetOpt("p"), "www")
		t.Assert(p.GetOpt("prefix"), "www")
		t.Assert(p.GetOpt("prefix").String(), "www")

		t.Assert(p.GetOpt("n") != nil, true)
		t.Assert(p.GetOpt("name") != nil, true)
		t.Assert(p.GetOpt("p") != nil, true)
		t.Assert(p.GetOpt("prefix") != nil, true)
		t.Assert(p.GetOpt("f") != nil, true)
		t.Assert(p.GetOpt("force") != nil, true)
		t.Assert(p.GetOpt("q") != nil, true)
		t.Assert(p.GetOpt("quiet") != nil, true)
		t.Assert(p.GetOpt("none") != nil, false)
	})
}
