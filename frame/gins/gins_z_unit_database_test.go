// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gins_test

import (
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/gins"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Database(t *testing.T) {
	databaseContent := gfile.GetContents(
		gtest.DataPath("database", "config.toml"),
	)
	gtest.C(t, func(t *gtest.T) {
		var err error
		dirPath := gfile.Temp(gtime.TimestampNanoStr())
		err = gfile.Mkdir(dirPath)
		t.AssertNil(err)
		defer gfile.Remove(dirPath)

		name := "config.toml"
		err = gfile.PutContents(gfile.Join(dirPath, name), databaseContent)
		t.AssertNil(err)

		err = gins.Config().GetAdapter().(*gcfg.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config().GetAdapter().(*gcfg.AdapterFile).Clear()

						// 用于gfsnotify回调以刷新配置文件的缓存. md5:6c5279392041ab52
		time.Sleep(500 * time.Millisecond)

				// 这段 Go 代码的注释翻译成中文是：打印一条消息，内容是 "gins 测试数据库"，后面跟着从配置（Config()）中获取的键为 "test" 的值。 md5:58e1615a972c88a5
		var (
			db        = gins.Database()
			dbDefault = gins.Database("default")
		)
		t.AssertNE(db, nil)
		t.AssertNE(dbDefault, nil)

		t.Assert(db.PingMaster(), nil)
		t.Assert(db.PingSlave(), nil)
		t.Assert(dbDefault.PingMaster(), nil)
		t.Assert(dbDefault.PingSlave(), nil)
	})
}
