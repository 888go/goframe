// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"fmt"
	"testing"

	garray "github.com/888go/goframe/container/garray"
	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Master_Slave(t *testing.T) {
	var err error

	gtest.C(t, func(t *gtest.T) {
		_, err = db.X原生SQL执行(ctx, "CREATE DATABASE IF NOT EXISTS `master` CHARACTER SET UTF8")
		t.AssertNil(err)
		_, err = db.X原生SQL执行(ctx, "CREATE DATABASE IF NOT EXISTS `slave` CHARACTER SET UTF8")
		t.AssertNil(err)
	})
	defer func() {
		_, _ = db.X原生SQL执行(ctx, "DROP DATABASE `master`")
		_, _ = db.X原生SQL执行(ctx, "DROP DATABASE `slave`")
	}()
	var (
		configKey   = guid.X生成()
		configGroup = gdb.ConfigGroup{
			gdb.ConfigNode{
				Host:   "127.0.0.1",
				Port:   "3306",
				User:   "root",
				Pass:   "12345678",
				Name:   "master",
				Type:   "mysql",
				Role:   "master",
				Debug:  true,
				Weight: 100,
			},
			gdb.ConfigNode{
				Host:   "127.0.0.1",
				Port:   "3306",
				User:   "root",
				Pass:   "12345678",
				Name:   "slave",
				Type:   "mysql",
				Role:   "slave",
				Debug:  true,
				Weight: 100,
			},
		}
	)
	gdb.X设置组配置(configKey, configGroup)
	masterSlaveDB := g.DB类(configKey)
	gtest.C(t, func(t *gtest.T) {
		table := "table_" + guid.X生成()
		createTableWithDb(masterSlaveDB.X切换数据库("master"), table)
		createTableWithDb(masterSlaveDB.X切换数据库("slave"), table)
		defer dropTableWithDb(masterSlaveDB.X切换数据库("master"), table)
		defer dropTableWithDb(masterSlaveDB.X切换数据库("slave"), table)

		// Data insert to master.
		array := garray.X创建(true)
		for i := 1; i <= TableSize; i++ {
			array.Append别名(g.Map{
				"id":          i,
				"passport":    fmt.Sprintf(`user_%d`, i),
				"password":    fmt.Sprintf(`pass_%d`, i),
				"nickname":    fmt.Sprintf(`name_%d`, i),
				"create_time": gtime.X创建并从文本(CreateTime).String(),
			})
		}
		_, err = masterSlaveDB.X创建Model对象(table).X设置数据(array).X插入()
		t.AssertNil(err)

		var count int
		// Auto slave.
		count, err = masterSlaveDB.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))

		// slave.
		count, err = masterSlaveDB.X创建Model对象(table).X取从节点对象().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))

		// master.
		count, err = masterSlaveDB.X创建Model对象(table).X取主节点对象().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
}
