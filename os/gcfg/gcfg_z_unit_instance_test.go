// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package 配置类

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
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			path = X默认配置文件名称
			err  = 文件类.X写入文本(path, config)
		)
		t.AssertNil(err)
		defer func() {
			t.AssertNil(文件类.X删除(path))
		}()

		c := X取单例对象()
		t.Assert(c.X取值PANI(ctx, "v1"), 1)
		filepath, _ := c.X取适配器().(*AdapterFile).GetFilePath()
		t.AssertEQ(filepath, 文件类.X取当前工作目录()+文件类.Separator+path)
	})
}

func Test_Instance_AutoLocateConfigFile(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(X取单例对象("gf") != nil, true)
	})
	// 自动定位并找到支持的文件扩展名的配置文件。
	单元测试类.C(t, func(t *单元测试类.T) {
		pwd := 文件类.X取当前工作目录()
		t.AssertNil(文件类.X设置当前工作目录(单元测试类.DataPath()))
		defer 文件类.X设置当前工作目录(pwd)
		t.Assert(X取单例对象("c1") != nil, true)
		t.Assert(X取单例对象("c1").X取值PANI(ctx, "my-config"), "1")
		t.Assert(X取单例对象("folder1/c1").X取值PANI(ctx, "my-config"), "2")
	})
	// 自动定位并找到支持的文件扩展名的配置文件。
	单元测试类.C(t, func(t *单元测试类.T) {
		pwd := 文件类.X取当前工作目录()
		t.AssertNil(文件类.X设置当前工作目录(单元测试类.DataPath("folder1")))
		defer 文件类.X设置当前工作目录(pwd)
		t.Assert(X取单例对象("c2").X取值PANI(ctx, "my-config"), 2)
	})
	// 默认配置文件
	单元测试类.C(t, func(t *单元测试类.T) {
		localInstances.X清空()
		pwd := 文件类.X取当前工作目录()
		t.AssertNil(文件类.X设置当前工作目录(单元测试类.DataPath("default")))
		defer 文件类.X设置当前工作目录(pwd)
		t.Assert(X取单例对象().X取值PANI(ctx, "my-config"), 1)

		localInstances.X清空()
		t.AssertNil(环境变量类.X设置值("GF_GCFG_FILE", "config.json"))
		defer 环境变量类.X设置值("GF_GCFG_FILE", "")
		t.Assert(X取单例对象().X取值PANI(ctx, "my-config"), 2)
	})
}

func Test_Instance_EnvPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		环境变量类.X设置值("GF_GCFG_PATH", 单元测试类.DataPath("envpath"))
		defer 环境变量类.X设置值("GF_GCFG_PATH", "")
		t.Assert(X取单例对象("c3") != nil, true)
		t.Assert(X取单例对象("c3").X取值PANI(ctx, "my-config"), "3")
		t.Assert(X取单例对象("c4").X取值PANI(ctx, "my-config"), "4")
		localInstances = map类.X创建StrAny(true)
	})
}

func Test_Instance_EnvFile(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		环境变量类.X设置值("GF_GCFG_PATH", 单元测试类.DataPath("envfile"))
		defer 环境变量类.X设置值("GF_GCFG_PATH", "")
		环境变量类.X设置值("GF_GCFG_FILE", "c6.json")
		defer 环境变量类.X设置值("GF_GCFG_FILE", "")
		t.Assert(X取单例对象().X取值PANI(ctx, "my-config"), "6")
		localInstances = map类.X创建StrAny(true)
	})
}
