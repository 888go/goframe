// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gins_test
import (
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/gins"
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
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

		// 用于gfsnotify回调刷新配置文件的缓存
		time.Sleep(500 * time.Millisecond)

		// 输出 "gins Test_Database" 和 Config() 获取到的 "test" 配置项内容
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
