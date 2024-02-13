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
	databaseContent := 文件类.X读文本(
		单元测试类.DataPath("database", "config.toml"),
	)
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		dirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		err = 文件类.X创建目录(dirPath)
		t.AssertNil(err)
		defer 文件类.X删除(dirPath)

		name := "config.toml"
		err = 文件类.X写入文本(文件类.X路径生成(dirPath, name), databaseContent)
		t.AssertNil(err)

		err = gins.Config().X取适配器().(*配置类.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config().X取适配器().(*配置类.AdapterFile).Clear()

		// 用于gfsnotify回调刷新配置文件的缓存
		time.Sleep(500 * time.Millisecond)

		// 输出 "gins Test_Database" 和 Config() 获取到的 "test" 配置项内容
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
