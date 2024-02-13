// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package 配置类_test

import (
	"testing"
	
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
)

func TestAdapterFile_Dump(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c, err := 配置类.NewAdapterFile("config.yml")
		t.AssertNil(err)

		t.Assert(c.GetFileName(), "config.yml")

		c.Dump()
		c.Data(ctx)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		c, err := 配置类.NewAdapterFile("testdata/default/config.toml")
		t.AssertNil(err)

		c.Dump()
		c.Data(ctx)
		c.GetPaths()
	})

}
func TestAdapterFile_Available(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c, err := 配置类.NewAdapterFile("testdata/default/config.toml")
		t.AssertNil(err)
		c.Available(ctx)
	})
}

func TestAdapterFile_SetPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c, err := 配置类.NewAdapterFile("config.yml")
		t.AssertNil(err)

		err = c.SetPath("/tmp")
		t.AssertNil(err)

		err = c.SetPath("notexist")
		t.AssertNE(err, nil)

		err = c.SetPath("testdata/c1.toml")
		t.AssertNE(err, nil)

		err = c.SetPath("")
		t.AssertNil(err)

		err = c.SetPath("gcfg.go")
		t.AssertNE(err, nil)

		v, err := c.Get(ctx, "name")
		t.AssertNE(err, nil)
		t.Assert(v, nil)
	})
}

func TestAdapterFile_AddPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c, err := 配置类.NewAdapterFile("config.yml")
		t.AssertNil(err)

		err = c.AddPath("/tmp")
		t.AssertNil(err)

		err = c.AddPath("notexist")
		t.AssertNE(err, nil)

		err = c.SetPath("testdata/c1.toml")
		t.AssertNE(err, nil)

		err = c.SetPath("")
		t.AssertNil(err)

		err = c.AddPath("gcfg.go")
		t.AssertNE(err, nil)

		v, err := c.Get(ctx, "name")
		t.AssertNE(err, nil)
		t.Assert(v, nil)
	})
}

func TestAdapterFile_SetViolenceCheck(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c, err := 配置类.NewAdapterFile("config.yml")
		t.AssertNil(err)
		c.SetViolenceCheck(true)
		v, err := c.Get(ctx, "name")
		t.AssertNE(err, nil)
		t.Assert(v, nil)
	})
}

func TestAdapterFile_FilePath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c, err := 配置类.NewAdapterFile("config.yml")
		t.AssertNil(err)

		path, _ := c.GetFilePath("tmp")
		t.Assert(path, "")

		path, _ = c.GetFilePath("tmp")
		t.Assert(path, "")
	})
}

func TestAdapterFile_Content(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c, err := 配置类.NewAdapterFile()
		t.AssertNil(err)

		c.SetContent("gf", "config.yml")
		t.Assert(c.GetContent("config.yml"), "gf")
		c.SetContent("gf1", "config.yml")
		t.Assert(c.GetContent("config.yml"), "gf1")
		c.RemoveContent("config.yml")
		c.ClearContent()
		t.Assert(c.GetContent("name"), "")
	})
}

func TestAdapterFile_With_UTF8_BOM(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c, err := 配置类.NewAdapterFile("test-cfg-with-utf8-bom")
		t.AssertNil(err)

		t.Assert(c.SetPath("testdata"), nil)
		c.SetFileName("cfg-with-utf8-bom.toml")
		t.Assert(c.MustGet(ctx, "test.testInt"), 1)
		t.Assert(c.MustGet(ctx, "test.testStr"), "test")
	})
}

func TestAdapterFile_Set(t *testing.T) {
	config := `log-path = "logs"`
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			path = 配置类.X默认配置文件名称
			err  = 文件类.X写入文本(path, config)
		)
		t.Assert(err, nil)
		defer 文件类.X删除(path)

		c, err := 配置类.X创建()
		t.Assert(c.X取值PANI(ctx, "log-path").String(), "logs")

		err = c.X取适配器().(*配置类.AdapterFile).X设置值("log-path", "custom-logs")
		t.Assert(err, nil)
		t.Assert(c.X取值PANI(ctx, "log-path").String(), "custom-logs")
	})
}
