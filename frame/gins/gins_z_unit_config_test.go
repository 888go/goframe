// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gins_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/frame/gins"
	gcfg "github.com/888go/goframe/os/gcfg"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

var (
	ctx           = context.Background()
	configContent = gfile.X读文本(
		gtest.DataPath("config", "config.toml"),
	)
)

func Test_Config1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(configContent, "")
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(gins.Config(), nil)
	})
}

func Test_Config2(t *testing.T) {
	// relative path
	gtest.C(t, func(t *gtest.T) {
		var err error
		dirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		err = gfile.X创建目录(dirPath)
		t.AssertNil(err)
		defer gfile.X删除(dirPath)

		name := "config.toml"
		err = gfile.X写入文本(gfile.X路径生成(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config().X取适配器().(*gcfg.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config().X取适配器().(*gcfg.AdapterFile).Clear()

		t.Assert(gins.Config().X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config().X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
		// 用于gfsnotify回调以刷新配置文件的缓存. md5:6c5279392041ab52
	time.Sleep(500 * time.Millisecond)

			// 相对路径，配置文件夹. md5:cd955a5e034346e9
	gtest.C(t, func(t *gtest.T) {
		var err error
		dirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		t.AssertNil(gfile.X创建目录(dirPath))
		defer gfile.X删除(dirPath)

		name := "config/config.toml"
		err = gfile.X写入文本(gfile.X路径生成(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config().X取适配器().(*gcfg.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config().X取适配器().(*gcfg.AdapterFile).Clear()

		t.Assert(gins.Config().X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config().X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)

			// 用于gfsnotify回调以刷新配置文件的缓存. md5:6c5279392041ab52
		time.Sleep(500 * time.Millisecond)
	})
}

func Test_Config3(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var err error
		dirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		err = gfile.X创建目录(dirPath)
		t.AssertNil(err)
		defer gfile.X删除(dirPath)

		name := "test.toml"
		err = gfile.X写入文本(gfile.X路径生成(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config("test").X取适配器().(*gcfg.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config("test").X取适配器().(*gcfg.AdapterFile).Clear()
		gins.Config("test").X取适配器().(*gcfg.AdapterFile).SetFileName("test.toml")

		t.Assert(gins.Config("test").X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
		// 用于gfsnotify回调以刷新配置文件的缓存. md5:6c5279392041ab52
	time.Sleep(500 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		var err error
		dirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		err = gfile.X创建目录(dirPath)
		t.AssertNil(err)
		defer gfile.X删除(dirPath)

		name := "config/test.toml"
		err = gfile.X写入文本(gfile.X路径生成(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config("test").X取适配器().(*gcfg.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config("test").X取适配器().(*gcfg.AdapterFile).Clear()
		gins.Config("test").X取适配器().(*gcfg.AdapterFile).SetFileName("test.toml")

		t.Assert(gins.Config("test").X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
			// 用于gfsnotify回调以刷新配置文件的缓存. md5:6c5279392041ab52
	time.Sleep(500 * time.Millisecond)
}

func Test_Config4(t *testing.T) {
	// absolute path
	gtest.C(t, func(t *gtest.T) {
		path := fmt.Sprintf(`%s/%d`, gfile.X取临时目录(), gtime.X取时间戳纳秒())
		file := fmt.Sprintf(`%s/%s`, path, "config.toml")
		err := gfile.X写入文本(file, configContent)
		t.AssertNil(err)
		defer gfile.X删除(file)
		defer gins.Config().X取适配器().(*gcfg.AdapterFile).Clear()

		t.Assert(gins.Config().X取适配器().(*gcfg.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config().X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config().X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	time.Sleep(500 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		path := fmt.Sprintf(`%s/%d/config`, gfile.X取临时目录(), gtime.X取时间戳纳秒())
		file := fmt.Sprintf(`%s/%s`, path, "config.toml")
		err := gfile.X写入文本(file, configContent)
		t.AssertNil(err)
		defer gfile.X删除(file)
		defer gins.Config().X取适配器().(*gcfg.AdapterFile).Clear()
		t.Assert(gins.Config().X取适配器().(*gcfg.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config().X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config().X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	time.Sleep(500 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		path := fmt.Sprintf(`%s/%d`, gfile.X取临时目录(), gtime.X取时间戳纳秒())
		file := fmt.Sprintf(`%s/%s`, path, "test.toml")
		err := gfile.X写入文本(file, configContent)
		t.AssertNil(err)
		defer gfile.X删除(file)
		defer gins.Config("test").X取适配器().(*gcfg.AdapterFile).Clear()
		gins.Config("test").X取适配器().(*gcfg.AdapterFile).SetFileName("test.toml")
		t.Assert(gins.Config("test").X取适配器().(*gcfg.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config("test").X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	time.Sleep(500 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		path := fmt.Sprintf(`%s/%d/config`, gfile.X取临时目录(), gtime.X取时间戳纳秒())
		file := fmt.Sprintf(`%s/%s`, path, "test.toml")
		err := gfile.X写入文本(file, configContent)
		t.AssertNil(err)
		defer gfile.X删除(file)
		defer gins.Config().X取适配器().(*gcfg.AdapterFile).Clear()
		gins.Config("test").X取适配器().(*gcfg.AdapterFile).SetFileName("test.toml")
		t.Assert(gins.Config("test").X取适配器().(*gcfg.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config("test").X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
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

		t.Assert(gins.Config().X取值PANI(ctx, "log-path"), "logs")
	})
}
