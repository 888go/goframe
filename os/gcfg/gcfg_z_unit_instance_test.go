// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gcfg

import (
	"context"
	"testing"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/os/genv"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
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
			path = DefaultConfigFileName
			err  = gfile.PutContents(path, config)
		)
		t.AssertNil(err)
		defer func() {
			t.AssertNil(gfile.Remove(path))
		}()

		c := Instance()
		t.Assert(c.MustGet(ctx, "v1"), 1)
		filepath, _ := c.GetAdapter().(*AdapterFile).GetFilePath()
		t.AssertEQ(filepath, gfile.Pwd()+gfile.Separator+path)
	})
}

func Test_Instance_AutoLocateConfigFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Instance("gf") != nil, true)
	})
	// 自动定位并找到支持的文件扩展名的配置文件。
	gtest.C(t, func(t *gtest.T) {
		pwd := gfile.Pwd()
		t.AssertNil(gfile.Chdir(gtest.DataPath()))
		defer gfile.Chdir(pwd)
		t.Assert(Instance("c1") != nil, true)
		t.Assert(Instance("c1").MustGet(ctx, "my-config"), "1")
		t.Assert(Instance("folder1/c1").MustGet(ctx, "my-config"), "2")
	})
	// 自动定位并找到支持的文件扩展名的配置文件。
	gtest.C(t, func(t *gtest.T) {
		pwd := gfile.Pwd()
		t.AssertNil(gfile.Chdir(gtest.DataPath("folder1")))
		defer gfile.Chdir(pwd)
		t.Assert(Instance("c2").MustGet(ctx, "my-config"), 2)
	})
	// 默认配置文件
	gtest.C(t, func(t *gtest.T) {
		localInstances.Clear()
		pwd := gfile.Pwd()
		t.AssertNil(gfile.Chdir(gtest.DataPath("default")))
		defer gfile.Chdir(pwd)
		t.Assert(Instance().MustGet(ctx, "my-config"), 1)

		localInstances.Clear()
		t.AssertNil(genv.Set("GF_GCFG_FILE", "config.json"))
		defer genv.Set("GF_GCFG_FILE", "")
		t.Assert(Instance().MustGet(ctx, "my-config"), 2)
	})
}

func Test_Instance_EnvPath(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		genv.Set("GF_GCFG_PATH", gtest.DataPath("envpath"))
		defer genv.Set("GF_GCFG_PATH", "")
		t.Assert(Instance("c3") != nil, true)
		t.Assert(Instance("c3").MustGet(ctx, "my-config"), "3")
		t.Assert(Instance("c4").MustGet(ctx, "my-config"), "4")
		localInstances = gmap.NewStrAnyMap(true)
	})
}

func Test_Instance_EnvFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		genv.Set("GF_GCFG_PATH", gtest.DataPath("envfile"))
		defer genv.Set("GF_GCFG_PATH", "")
		genv.Set("GF_GCFG_FILE", "c6.json")
		defer genv.Set("GF_GCFG_FILE", "")
		t.Assert(Instance().MustGet(ctx, "my-config"), "6")
		localInstances = gmap.NewStrAnyMap(true)
	})
}
