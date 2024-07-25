// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package mysql_test

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
)

func Test_Master_Slave(t *testing.T) {
	var err error

	gtest.C(t, func(t *gtest.T) {
		_, err = db.Exec(ctx, "CREATE DATABASE IF NOT EXISTS `master` CHARACTER SET UTF8")
		t.AssertNil(err)
		_, err = db.Exec(ctx, "CREATE DATABASE IF NOT EXISTS `slave` CHARACTER SET UTF8")
		t.AssertNil(err)
	})
	defer func() {
		_, _ = db.Exec(ctx, "DROP DATABASE `master`")
		_, _ = db.Exec(ctx, "DROP DATABASE `slave`")
	}()
	var (
		configKey   = guid.S()
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
	gdb.SetConfigGroup(configKey, configGroup)
	masterSlaveDB := g.DB(configKey)
	gtest.C(t, func(t *gtest.T) {
		table := "table_" + guid.S()
		createTableWithDb(masterSlaveDB.Schema("master"), table)
		createTableWithDb(masterSlaveDB.Schema("slave"), table)
		defer dropTableWithDb(masterSlaveDB.Schema("master"), table)
		defer dropTableWithDb(masterSlaveDB.Schema("slave"), table)

		// Data insert to master.
		array := garray.New(true)
		for i := 1; i <= TableSize; i++ {
			array.Append(g.Map{
				"id":          i,
				"passport":    fmt.Sprintf(`user_%d`, i),
				"password":    fmt.Sprintf(`pass_%d`, i),
				"nickname":    fmt.Sprintf(`name_%d`, i),
				"create_time": gtime.NewFromStr(CreateTime).String(),
			})
		}
		_, err = masterSlaveDB.Model(table).Data(array).Insert()
		t.AssertNil(err)

		var count int
		// Auto slave.
		count, err = masterSlaveDB.Model(table).Count()
		t.AssertNil(err)
		t.Assert(count, int64(0))

		// slave.
		count, err = masterSlaveDB.Model(table).Slave().Count()
		t.AssertNil(err)
		t.Assert(count, int64(0))

		// master.
		count, err = masterSlaveDB.Model(table).Master().Count()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
}
