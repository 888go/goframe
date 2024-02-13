// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package 配置类_test

import (
	"testing"
	
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/gcmd"
	"github.com/888go/goframe/os/genv"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

func Test_Basic1(t *testing.T) {
	config := `
v1    = 1
v2    = "true"
v3    = "off"
v4    = "1.23"
array = [1,2,3]
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			path = 配置类.X默认配置文件名称
			err  = 文件类.X写入文本(path, config)
		)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		c, err := 配置类.X创建()
		t.AssertNil(err)
		t.Assert(c.X取值PANI(ctx, "v1"), 1)
		filepath, _ := c.X取适配器().(*配置类.AdapterFile).GetFilePath()
		t.AssertEQ(filepath, 文件类.X取当前工作目录()+文件类.Separator+path)
	})
}

func Test_Basic2(t *testing.T) {
	config := `log-path = "logs"`
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			path = 配置类.X默认配置文件名称
			err  = 文件类.X写入文本(path, config)
		)
		t.AssertNil(err)
		defer func() {
			_ = 文件类.X删除(path)
		}()

		c, err := 配置类.X创建()
		t.AssertNil(err)
		t.Assert(c.X取值PANI(ctx, "log-path"), "logs")
	})
}

func Test_Content(t *testing.T) {
	content := `
v1    = 1
v2    = "true"
v3    = "off"
v4    = "1.23"
array = [1,2,3]
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`
	单元测试类.C(t, func(t *单元测试类.T) {
		c, err := 配置类.X创建()
		t.AssertNil(err)
		c.X取适配器().(*配置类.AdapterFile).SetContent(content)
		defer c.X取适配器().(*配置类.AdapterFile).ClearContent()
		t.Assert(c.X取值PANI(ctx, "v1"), 1)
	})
}

func Test_SetFileName(t *testing.T) {
	config := `
{
	"array": [
		1,
		2,
		3
	],
	"redis": {
		"cache": "127.0.0.1:6379,1",
		"disk": "127.0.0.1:6379,0"
	},
	"v1": 1,
	"v2": "true",
	"v3": "off",
	"v4": "1.234"
}
`
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "config.json"
		err := 文件类.X写入文本(path, config)
		t.AssertNil(err)
		defer func() {
			_ = 文件类.X删除(path)
		}()

		config, err := 配置类.X创建()
		t.AssertNil(err)
		c := config.X取适配器().(*配置类.AdapterFile)
		c.SetFileName(path)
		t.Assert(c.MustGet(ctx, "v1"), 1)
		t.AssertEQ(c.MustGet(ctx, "v1").X取整数(), 1)
		t.AssertEQ(c.MustGet(ctx, "v1").X取整数8位(), int8(1))
		t.AssertEQ(c.MustGet(ctx, "v1").X取整数16位(), int16(1))
		t.AssertEQ(c.MustGet(ctx, "v1").X取整数32位(), int32(1))
		t.AssertEQ(c.MustGet(ctx, "v1").X取整数64位(), int64(1))
		t.AssertEQ(c.MustGet(ctx, "v1").X取正整数(), uint(1))
		t.AssertEQ(c.MustGet(ctx, "v1").X取正整数8位(), uint8(1))
		t.AssertEQ(c.MustGet(ctx, "v1").X取正整数16位(), uint16(1))
		t.AssertEQ(c.MustGet(ctx, "v1").X取正整数32位(), uint32(1))
		t.AssertEQ(c.MustGet(ctx, "v1").X取正整数64位(), uint64(1))

		t.AssertEQ(c.MustGet(ctx, "v1").String(), "1")
		t.AssertEQ(c.MustGet(ctx, "v1").X取布尔(), true)
		t.AssertEQ(c.MustGet(ctx, "v2").String(), "true")
		t.AssertEQ(c.MustGet(ctx, "v2").X取布尔(), true)

		t.AssertEQ(c.MustGet(ctx, "v1").String(), "1")
		t.AssertEQ(c.MustGet(ctx, "v4").X取小数32位(), float32(1.234))
		t.AssertEQ(c.MustGet(ctx, "v4").X取小数64位(), float64(1.234))
		t.AssertEQ(c.MustGet(ctx, "v2").String(), "true")
		t.AssertEQ(c.MustGet(ctx, "v2").X取布尔(), true)
		t.AssertEQ(c.MustGet(ctx, "v3").X取布尔(), false)

		t.AssertEQ(c.MustGet(ctx, "array").X取整数数组(), []int{1, 2, 3})
		t.AssertEQ(c.MustGet(ctx, "array").X取文本数组(), []string{"1", "2", "3"})
		t.AssertEQ(c.MustGet(ctx, "array").X取any数组(), []interface{}{1, 2, 3})
		t.AssertEQ(c.MustGet(ctx, "redis").X取Map(), map[string]interface{}{
			"disk":  "127.0.0.1:6379,0",
			"cache": "127.0.0.1:6379,1",
		})
		filepath, _ := c.GetFilePath()
		t.AssertEQ(filepath, 文件类.X取当前工作目录()+文件类.Separator+path)
	})
}

func TestCfg_Get_WrongConfigFile(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		configPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		err = 文件类.X创建目录(configPath)
		t.AssertNil(err)
		defer 文件类.X删除(configPath)

		defer 文件类.X设置当前工作目录(文件类.X取当前工作目录())
		err = 文件类.X设置当前工作目录(configPath)
		t.AssertNil(err)

		err = 文件类.X写入文本(
			文件类.X路径生成(configPath, "config.yml"),
			"wrong config",
		)
		t.AssertNil(err)
		adapterFile, err := 配置类.NewAdapterFile("config.yml")
		t.AssertNil(err)

		c := 配置类.X创建并按适配器(adapterFile)
		v, err := c.X取值(ctx, "name")
		t.AssertNE(err, nil)
		t.Assert(v, nil)
		adapterFile.Clear()
	})
}

func Test_GetWithEnv(t *testing.T) {
	content := `
v1    = 1
v2    = "true"
v3    = "off"
v4    = "1.23"
array = [1,2,3]
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`
	单元测试类.C(t, func(t *单元测试类.T) {
		c, err := 配置类.X创建()
		t.AssertNil(err)
		c.X取适配器().(*配置类.AdapterFile).SetContent(content)
		defer c.X取适配器().(*配置类.AdapterFile).ClearContent()
		t.Assert(c.X取值PANI(ctx, "v1"), 1)
		t.Assert(c.X取值并从环境变量PANI(ctx, `redis.user`), nil)
		t.Assert(环境变量类.X设置值("REDIS_USER", `1`), nil)
		defer 环境变量类.X删除(`REDIS_USER`)
		t.Assert(c.X取值并从环境变量PANI(ctx, `redis.user`), `1`)
	})
}

func Test_GetWithCmd(t *testing.T) {
	content := `
v1    = 1
v2    = "true"
v3    = "off"
v4    = "1.23"
array = [1,2,3]
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`
	单元测试类.C(t, func(t *单元测试类.T) {

		c, err := 配置类.X创建()
		t.AssertNil(err)
		c.X取适配器().(*配置类.AdapterFile).SetContent(content)
		defer c.X取适配器().(*配置类.AdapterFile).ClearContent()
		t.Assert(c.X取值PANI(ctx, "v1"), 1)
		t.Assert(c.X取值并从启动命令PANI_有bug(ctx, `redis.user`), nil)

		cmd类.Init([]string{"gf", "--redis.user=2"}...)
		t.Assert(c.X取值并从启动命令PANI_有bug(ctx, `redis.user`), `2`)
	})
}
