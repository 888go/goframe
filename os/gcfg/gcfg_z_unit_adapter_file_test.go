// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 配置类_test

import (
	"testing"

	gcfg "github.com/888go/goframe/os/gcfg"
	gfile "github.com/888go/goframe/os/gfile"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestAdapterFile_Dump(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		c, err := gcfg.NewAdapterFile("config.yml")
		t.AssertNil(err)

		t.Assert(c.GetFileName(), "config.yml")

		c.Dump()
		c.Data(ctx)
	})

	gtest.C(t, func(t *gtest.T) {
		c, err := gcfg.NewAdapterFile("testdata/default/config.toml")
		t.AssertNil(err)

		c.Dump()
		c.Data(ctx)
		c.GetPaths()
	})

}
func TestAdapterFile_Available(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		c, err := gcfg.NewAdapterFile("testdata/default/config.toml")
		t.AssertNil(err)
		c.Available(ctx)
	})
}

func TestAdapterFile_SetPath(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		c, err := gcfg.NewAdapterFile("config.yml")
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
	gtest.C(t, func(t *gtest.T) {
		c, err := gcfg.NewAdapterFile("config.yml")
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
	gtest.C(t, func(t *gtest.T) {
		c, err := gcfg.NewAdapterFile("config.yml")
		t.AssertNil(err)
		c.SetViolenceCheck(true)
		v, err := c.Get(ctx, "name")
		t.AssertNE(err, nil)
		t.Assert(v, nil)
	})
}

func TestAdapterFile_FilePath(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		c, err := gcfg.NewAdapterFile("config.yml")
		t.AssertNil(err)

		path, _ := c.GetFilePath("tmp")
		t.Assert(path, "")

		path, _ = c.GetFilePath("tmp")
		t.Assert(path, "")
	})
}

func TestAdapterFile_Content(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		c, err := gcfg.NewAdapterFile()
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
	gtest.C(t, func(t *gtest.T) {
		c, err := gcfg.NewAdapterFile("test-cfg-with-utf8-bom")
		t.AssertNil(err)

		t.Assert(c.SetPath("testdata"), nil)
		c.SetFileName("cfg-with-utf8-bom.toml")
		t.Assert(c.MustGet(ctx, "test.testInt"), 1)
		t.Assert(c.MustGet(ctx, "test.testStr"), "test")
	})
}

func TestAdapterFile_Set(t *testing.T) {
	config := `log-path = "logs"`
	gtest.C(t, func(t *gtest.T) {
		var (
			path = gcfg.X默认配置文件名称
			err  = gfile.X写入文本(path, config)
		)
		t.AssertNil(err)
		defer gfile.X删除(path)

		c, err := gcfg.X创建()
		t.Assert(c.X取值PANI(ctx, "log-path").String(), "logs")

		err = c.X取适配器().(*gcfg.AdapterFile).X设置值("log-path", "custom-logs")
		t.AssertNil(err)
		t.Assert(c.X取值PANI(ctx, "log-path").String(), "custom-logs")
	})
}
