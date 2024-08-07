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

	"github.com/888go/goframe/frame/gins"
	gcfg "github.com/888go/goframe/os/gcfg"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Database(t *testing.T) {
	databaseContent := gfile.X读文本(
		gtest.DataPath("database", "config.toml"),
	)
	gtest.C(t, func(t *gtest.T) {
		var err error
		dirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		err = gfile.X创建目录(dirPath)
		t.AssertNil(err)
		defer gfile.X删除(dirPath)

		name := "config.toml"
		err = gfile.X写入文本(gfile.X路径生成(dirPath, name), databaseContent)
		t.AssertNil(err)

		err = gins.Config().X取适配器().(*gcfg.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config().X取适配器().(*gcfg.AdapterFile).Clear()

						// 用于gfsnotify回调以刷新配置文件的缓存. md5:6c5279392041ab52
		time.Sleep(500 * time.Millisecond)

				// 这段 Go 代码的注释翻译成中文是：打印一条消息，内容是 "gins 测试数据库"，后面跟着从配置（Config()）中获取的键为 "test" 的值。 md5:58e1615a972c88a5
		var (
			db        = gins.Database()
			dbDefault = gins.Database("default")
		)
		t.AssertNE(db, nil)
		t.AssertNE(dbDefault, nil)

		t.Assert(db.X向主节点发送心跳(), nil)
		t.Assert(db.X向从节点发送心跳(), nil)
		t.Assert(dbDefault.X向主节点发送心跳(), nil)
		t.Assert(dbDefault.X向从节点发送心跳(), nil)
	})
}
