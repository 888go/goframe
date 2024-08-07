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
	gcmd "github.com/888go/goframe/os/gcmd"
	genv "github.com/888go/goframe/os/genv"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
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
	gtest.C(t, func(t *gtest.T) {
		var (
			path = gcfg.X默认配置文件名称
			err  = gfile.X写入文本(path, config)
		)
		t.AssertNil(err)
		defer gfile.X删除(path)

		c, err := gcfg.X创建()
		t.AssertNil(err)
		t.Assert(c.X取值PANI(ctx, "v1"), 1)
		filepath, _ := c.X取适配器().(*gcfg.AdapterFile).GetFilePath()
		t.AssertEQ(filepath, gfile.X取当前工作目录()+gfile.Separator+path)
	})
}

func Test_Basic2(t *testing.T) {
	config := `log-path = "logs"`
	gtest.C(t, func(t *gtest.T) {
		var (
			path = gcfg.X默认配置文件名称
			err  = gfile.X写入文本(path, config)
		)
		t.AssertNil(err)
		defer func() {
			_ = gfile.X删除(path)
		}()

		c, err := gcfg.X创建()
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
	gtest.C(t, func(t *gtest.T) {
		c, err := gcfg.X创建()
		t.AssertNil(err)
		c.X取适配器().(*gcfg.AdapterFile).SetContent(content)
		defer c.X取适配器().(*gcfg.AdapterFile).ClearContent()
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
	gtest.C(t, func(t *gtest.T) {
		path := "config.json"
		err := gfile.X写入文本(path, config)
		t.AssertNil(err)
		defer func() {
			_ = gfile.X删除(path)
		}()

		config, err := gcfg.X创建()
		t.AssertNil(err)
		c := config.X取适配器().(*gcfg.AdapterFile)
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

		t.AssertEQ(c.MustGet(ctx, "array").X取整数切片(), []int{1, 2, 3})
		t.AssertEQ(c.MustGet(ctx, "array").X取文本切片(), []string{"1", "2", "3"})
		t.AssertEQ(c.MustGet(ctx, "array").X取any切片(), []interface{}{1, 2, 3})
		t.AssertEQ(c.MustGet(ctx, "redis").X取Map(), map[string]interface{}{
			"disk":  "127.0.0.1:6379,0",
			"cache": "127.0.0.1:6379,1",
		})
		filepath, _ := c.GetFilePath()
		t.AssertEQ(filepath, gfile.X取当前工作目录()+gfile.Separator+path)
	})
}

func TestCfg_Get_WrongConfigFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var err error
		configPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		err = gfile.X创建目录(configPath)
		t.AssertNil(err)
		defer gfile.X删除(configPath)

		defer gfile.X设置当前工作目录(gfile.X取当前工作目录())
		err = gfile.X设置当前工作目录(configPath)
		t.AssertNil(err)

		err = gfile.X写入文本(
			gfile.X路径生成(configPath, "config.yml"),
			"wrong config",
		)
		t.AssertNil(err)
		adapterFile, err := gcfg.NewAdapterFile("config.yml")
		t.AssertNil(err)

		c := gcfg.X创建并按适配器(adapterFile)
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
	gtest.C(t, func(t *gtest.T) {
		c, err := gcfg.X创建()
		t.AssertNil(err)
		c.X取适配器().(*gcfg.AdapterFile).SetContent(content)
		defer c.X取适配器().(*gcfg.AdapterFile).ClearContent()
		t.Assert(c.X取值PANI(ctx, "v1"), 1)
		t.Assert(c.X取值并从环境变量PANI(ctx, `redis.user`), nil)
		t.Assert(genv.X设置值("REDIS_USER", `1`), nil)
		defer genv.X删除(`REDIS_USER`)
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
	gtest.C(t, func(t *gtest.T) {

		c, err := gcfg.X创建()
		t.AssertNil(err)
		c.X取适配器().(*gcfg.AdapterFile).SetContent(content)
		defer c.X取适配器().(*gcfg.AdapterFile).ClearContent()
		t.Assert(c.X取值PANI(ctx, "v1"), 1)
		t.Assert(c.X取值并从启动命令PANI_有bug(ctx, `redis.user`), nil)

		gcmd.Init([]string{"gf", "--redis.user=2"}...)
		t.Assert(c.X取值并从启动命令PANI_有bug(ctx, `redis.user`), `2`)
	})
}
