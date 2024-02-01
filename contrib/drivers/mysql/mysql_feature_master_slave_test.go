// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test
import (
	"fmt"
	"testing"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
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

		// 向主库插入数据。
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
