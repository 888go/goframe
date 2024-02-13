// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gins_test

import (
	"context"
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/gins"
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

var (
	ctx           = context.Background()
	configContent = 文件类.X读文本(
		单元测试类.DataPath("config", "config.toml"),
	)
)

func Test_Config1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNE(configContent, "")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNE(gins.Config(), nil)
	})
}

func Test_Config2(t *testing.T) {
	// relative path
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		dirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		err = 文件类.X创建目录(dirPath)
		t.AssertNil(err)
		defer 文件类.X删除(dirPath)

		name := "config.toml"
		err = 文件类.X写入文本(文件类.X路径生成(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config().X取适配器().(*配置类.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config().X取适配器().(*配置类.AdapterFile).Clear()

		t.Assert(gins.Config().X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config().X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	// 用于gfsnotify回调以刷新配置文件的缓存
	time.Sleep(500 * time.Millisecond)

	// 相对路径，配置文件夹
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		dirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		t.AssertNil(文件类.X创建目录(dirPath))
		defer 文件类.X删除(dirPath)

		name := "config/config.toml"
		err = 文件类.X写入文本(文件类.X路径生成(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config().X取适配器().(*配置类.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config().X取适配器().(*配置类.AdapterFile).Clear()

		t.Assert(gins.Config().X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config().X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)

		// 用于gfsnotify回调以刷新配置文件的缓存
		time.Sleep(500 * time.Millisecond)
	})
}

func Test_Config3(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		dirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		err = 文件类.X创建目录(dirPath)
		t.AssertNil(err)
		defer 文件类.X删除(dirPath)

		name := "test.toml"
		err = 文件类.X写入文本(文件类.X路径生成(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config("test").X取适配器().(*配置类.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config("test").X取适配器().(*配置类.AdapterFile).Clear()
		gins.Config("test").X取适配器().(*配置类.AdapterFile).SetFileName("test.toml")

		t.Assert(gins.Config("test").X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	// 用于gfsnotify回调以刷新配置文件的缓存
	time.Sleep(500 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		dirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		err = 文件类.X创建目录(dirPath)
		t.AssertNil(err)
		defer 文件类.X删除(dirPath)

		name := "config/test.toml"
		err = 文件类.X写入文本(文件类.X路径生成(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config("test").X取适配器().(*配置类.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config("test").X取适配器().(*配置类.AdapterFile).Clear()
		gins.Config("test").X取适配器().(*配置类.AdapterFile).SetFileName("test.toml")

		t.Assert(gins.Config("test").X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	// 用于gfsnotify回调以刷新配置文件的缓存 for next unit testing case.
	time.Sleep(500 * time.Millisecond)
}

func Test_Config4(t *testing.T) {
	// absolute path
	单元测试类.C(t, func(t *单元测试类.T) {
		path := fmt.Sprintf(`%s/%d`, 文件类.X取临时目录(), 时间类.X取时间戳纳秒())
		file := fmt.Sprintf(`%s/%s`, path, "config.toml")
		err := 文件类.X写入文本(file, configContent)
		t.AssertNil(err)
		defer 文件类.X删除(file)
		defer gins.Config().X取适配器().(*配置类.AdapterFile).Clear()

		t.Assert(gins.Config().X取适配器().(*配置类.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config().X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config().X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	time.Sleep(500 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		path := fmt.Sprintf(`%s/%d/config`, 文件类.X取临时目录(), 时间类.X取时间戳纳秒())
		file := fmt.Sprintf(`%s/%s`, path, "config.toml")
		err := 文件类.X写入文本(file, configContent)
		t.AssertNil(err)
		defer 文件类.X删除(file)
		defer gins.Config().X取适配器().(*配置类.AdapterFile).Clear()
		t.Assert(gins.Config().X取适配器().(*配置类.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config().X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config().X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	time.Sleep(500 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		path := fmt.Sprintf(`%s/%d`, 文件类.X取临时目录(), 时间类.X取时间戳纳秒())
		file := fmt.Sprintf(`%s/%s`, path, "test.toml")
		err := 文件类.X写入文本(file, configContent)
		t.AssertNil(err)
		defer 文件类.X删除(file)
		defer gins.Config("test").X取适配器().(*配置类.AdapterFile).Clear()
		gins.Config("test").X取适配器().(*配置类.AdapterFile).SetFileName("test.toml")
		t.Assert(gins.Config("test").X取适配器().(*配置类.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config("test").X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	time.Sleep(500 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		path := fmt.Sprintf(`%s/%d/config`, 文件类.X取临时目录(), 时间类.X取时间戳纳秒())
		file := fmt.Sprintf(`%s/%s`, path, "test.toml")
		err := 文件类.X写入文本(file, configContent)
		t.AssertNil(err)
		defer 文件类.X删除(file)
		defer gins.Config().X取适配器().(*配置类.AdapterFile).Clear()
		gins.Config("test").X取适配器().(*配置类.AdapterFile).SetFileName("test.toml")
		t.Assert(gins.Config("test").X取适配器().(*配置类.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config("test").X取值PANI(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").X取值PANI(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
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

		t.Assert(gins.Config().X取值PANI(ctx, "log-path"), "logs")
	})
}
