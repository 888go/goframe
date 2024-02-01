// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gcmd_test
import (
	"os"
	"testing"
	
	"github.com/888go/goframe/os/gcmd"
	"github.com/888go/goframe/test/gtest"
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
