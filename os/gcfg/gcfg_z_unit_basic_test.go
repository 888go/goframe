// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。. md5:81db3d7bd1ed4da8

package gcfg_test

import (
	"testing"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
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
			path = gcfg.DefaultConfigFileName
			err  = gfile.PutContents(path, config)
		)
		t.AssertNil(err)
		defer gfile.Remove(path)

		c, err := gcfg.New()
		t.AssertNil(err)
		t.Assert(c.MustGet(ctx, "v1"), 1)
		filepath, _ := c.GetAdapter().(*gcfg.AdapterFile).GetFilePath()
		t.AssertEQ(filepath, gfile.Pwd()+gfile.Separator+path)
	})
}

func Test_Basic2(t *testing.T) {
	config := `log-path = "logs"`
	gtest.C(t, func(t *gtest.T) {
		var (
			path = gcfg.DefaultConfigFileName
			err  = gfile.PutContents(path, config)
		)
		t.AssertNil(err)
		defer func() {
			_ = gfile.Remove(path)
		}()

		c, err := gcfg.New()
		t.AssertNil(err)
		t.Assert(c.MustGet(ctx, "log-path"), "logs")
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
		c, err := gcfg.New()
		t.AssertNil(err)
		c.GetAdapter().(*gcfg.AdapterFile).SetContent(content)
		defer c.GetAdapter().(*gcfg.AdapterFile).ClearContent()
		t.Assert(c.MustGet(ctx, "v1"), 1)
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
		err := gfile.PutContents(path, config)
		t.AssertNil(err)
		defer func() {
			_ = gfile.Remove(path)
		}()

		config, err := gcfg.New()
		t.AssertNil(err)
		c := config.GetAdapter().(*gcfg.AdapterFile)
		c.SetFileName(path)
		t.Assert(c.MustGet(ctx, "v1"), 1)
		t.AssertEQ(c.MustGet(ctx, "v1").Int(), 1)
		t.AssertEQ(c.MustGet(ctx, "v1").Int8(), int8(1))
		t.AssertEQ(c.MustGet(ctx, "v1").Int16(), int16(1))
		t.AssertEQ(c.MustGet(ctx, "v1").Int32(), int32(1))
		t.AssertEQ(c.MustGet(ctx, "v1").Int64(), int64(1))
		t.AssertEQ(c.MustGet(ctx, "v1").Uint(), uint(1))
		t.AssertEQ(c.MustGet(ctx, "v1").Uint8(), uint8(1))
		t.AssertEQ(c.MustGet(ctx, "v1").Uint16(), uint16(1))
		t.AssertEQ(c.MustGet(ctx, "v1").Uint32(), uint32(1))
		t.AssertEQ(c.MustGet(ctx, "v1").Uint64(), uint64(1))

		t.AssertEQ(c.MustGet(ctx, "v1").String(), "1")
		t.AssertEQ(c.MustGet(ctx, "v1").Bool(), true)
		t.AssertEQ(c.MustGet(ctx, "v2").String(), "true")
		t.AssertEQ(c.MustGet(ctx, "v2").Bool(), true)

		t.AssertEQ(c.MustGet(ctx, "v1").String(), "1")
		t.AssertEQ(c.MustGet(ctx, "v4").Float32(), float32(1.234))
		t.AssertEQ(c.MustGet(ctx, "v4").Float64(), float64(1.234))
		t.AssertEQ(c.MustGet(ctx, "v2").String(), "true")
		t.AssertEQ(c.MustGet(ctx, "v2").Bool(), true)
		t.AssertEQ(c.MustGet(ctx, "v3").Bool(), false)

		t.AssertEQ(c.MustGet(ctx, "array").Ints(), []int{1, 2, 3})
		t.AssertEQ(c.MustGet(ctx, "array").Strings(), []string{"1", "2", "3"})
		t.AssertEQ(c.MustGet(ctx, "array").Interfaces(), []interface{}{1, 2, 3})
		t.AssertEQ(c.MustGet(ctx, "redis").Map(), map[string]interface{}{
			"disk":  "127.0.0.1:6379,0",
			"cache": "127.0.0.1:6379,1",
		})
		filepath, _ := c.GetFilePath()
		t.AssertEQ(filepath, gfile.Pwd()+gfile.Separator+path)
	})
}

func TestCfg_Get_WrongConfigFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var err error
		configPath := gfile.Temp(gtime.TimestampNanoStr())
		err = gfile.Mkdir(configPath)
		t.AssertNil(err)
		defer gfile.Remove(configPath)

		defer gfile.Chdir(gfile.Pwd())
		err = gfile.Chdir(configPath)
		t.AssertNil(err)

		err = gfile.PutContents(
			gfile.Join(configPath, "config.yml"),
			"wrong config",
		)
		t.AssertNil(err)
		adapterFile, err := gcfg.NewAdapterFile("config.yml")
		t.AssertNil(err)

		c := gcfg.NewWithAdapter(adapterFile)
		v, err := c.Get(ctx, "name")
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
		c, err := gcfg.New()
		t.AssertNil(err)
		c.GetAdapter().(*gcfg.AdapterFile).SetContent(content)
		defer c.GetAdapter().(*gcfg.AdapterFile).ClearContent()
		t.Assert(c.MustGet(ctx, "v1"), 1)
		t.Assert(c.MustGetWithEnv(ctx, `redis.user`), nil)
		t.Assert(genv.Set("REDIS_USER", `1`), nil)
		defer genv.Remove(`REDIS_USER`)
		t.Assert(c.MustGetWithEnv(ctx, `redis.user`), `1`)
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

		c, err := gcfg.New()
		t.AssertNil(err)
		c.GetAdapter().(*gcfg.AdapterFile).SetContent(content)
		defer c.GetAdapter().(*gcfg.AdapterFile).ClearContent()
		t.Assert(c.MustGet(ctx, "v1"), 1)
		t.Assert(c.MustGetWithCmd(ctx, `redis.user`), nil)

		gcmd.Init([]string{"gf", "--redis.user=2"}...)
		t.Assert(c.MustGetWithCmd(ctx, `redis.user`), `2`)
	})
}
