// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 配置类

import (
	"context"
	"testing"

	gmap "github.com/888go/goframe/container/gmap"
	genv "github.com/888go/goframe/os/genv"
	gfile "github.com/888go/goframe/os/gfile"
	gtest "github.com/888go/goframe/test/gtest"
)

var (
	ctx = context.TODO()
)

func Test_Instance_Basic(t *testing.T) {
	config := `
array = [1.0, 2.0, 3.0]
v1 = 1.0
v2 = "true"
v3 = "off"
v4 = "1.234"

[redis]
  cache = "127.0.0.1:6379,1"
  disk = "127.0.0.1:6379,0"

`
	gtest.C(t, func(t *gtest.T) {
		var (
			path = X默认配置文件名称
			err  = gfile.X写入文本(path, config)
		)
		t.AssertNil(err)
		defer func() {
			t.AssertNil(gfile.X删除(path))
		}()

		c := X取单例对象()
		t.Assert(c.X取值PANI(ctx, "v1"), 1)
		filepath, _ := c.X取适配器().(*AdapterFile).GetFilePath()
		t.AssertEQ(filepath, gfile.X取当前工作目录()+gfile.Separator+path)
	})
}

func Test_Instance_AutoLocateConfigFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(X取单例对象("gf") != nil, true)
	})
		// 自动定位支持的文件扩展名的配置文件。 md5:941b2ef0c3ebcbf1
	gtest.C(t, func(t *gtest.T) {
		pwd := gfile.X取当前工作目录()
		t.AssertNil(gfile.X设置当前工作目录(gtest.DataPath()))
		defer gfile.X设置当前工作目录(pwd)
		t.Assert(X取单例对象("c1") != nil, true)
		t.Assert(X取单例对象("c1").X取值PANI(ctx, "my-config"), "1")
		t.Assert(X取单例对象("folder1/c1").X取值PANI(ctx, "my-config"), "2")
	})
		// 自动定位支持的文件扩展名的配置文件。 md5:941b2ef0c3ebcbf1
	gtest.C(t, func(t *gtest.T) {
		pwd := gfile.X取当前工作目录()
		t.AssertNil(gfile.X设置当前工作目录(gtest.DataPath("folder1")))
		defer gfile.X设置当前工作目录(pwd)
		t.Assert(X取单例对象("c2").X取值PANI(ctx, "my-config"), 2)
	})
		// 默认配置文件。 md5:bfb03b7e4e99b27b
	gtest.C(t, func(t *gtest.T) {
		localInstances.X清空()
		pwd := gfile.X取当前工作目录()
		t.AssertNil(gfile.X设置当前工作目录(gtest.DataPath("default")))
		defer gfile.X设置当前工作目录(pwd)
		t.Assert(X取单例对象().X取值PANI(ctx, "my-config"), 1)

		localInstances.X清空()
		t.AssertNil(genv.X设置值("GF_GCFG_FILE", "config.json"))
		defer genv.X设置值("GF_GCFG_FILE", "")
		t.Assert(X取单例对象().X取值PANI(ctx, "my-config"), 2)
	})
}

func Test_Instance_EnvPath(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		genv.X设置值("GF_GCFG_PATH", gtest.DataPath("envpath"))
		defer genv.X设置值("GF_GCFG_PATH", "")
		t.Assert(X取单例对象("c3") != nil, true)
		t.Assert(X取单例对象("c3").X取值PANI(ctx, "my-config"), "3")
		t.Assert(X取单例对象("c4").X取值PANI(ctx, "my-config"), "4")
		localInstances = gmap.X创建StrAny(true)
	})
}

func Test_Instance_EnvFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		genv.X设置值("GF_GCFG_PATH", gtest.DataPath("envfile"))
		defer genv.X设置值("GF_GCFG_PATH", "")
		genv.X设置值("GF_GCFG_FILE", "c6.json")
		defer genv.X设置值("GF_GCFG_FILE", "")
		t.Assert(X取单例对象().X取值PANI(ctx, "my-config"), "6")
		localInstances = gmap.X创建StrAny(true)
	})
}
