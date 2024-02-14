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

	单元测试类.C(t, func(t *单元测试类.T) {
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
		configKey   = uid类.X生成()
		configGroup = db类.X配置组{
			db类.X配置项{
				X地址:   "127.0.0.1",
				X端口:   "3306",
				X账号:   "root",
				X密码:   "12345678",
				X名称:   "master",
				X类型:   "mysql",
				X节点角色:   "master",
				X调试模式:  true,
				X负载均衡权重: 100,
			},
			db类.X配置项{
				X地址:   "127.0.0.1",
				X端口:   "3306",
				X账号:   "root",
				X密码:   "12345678",
				X名称:   "slave",
				X类型:   "mysql",
				X节点角色:   "slave",
				X调试模式:  true,
				X负载均衡权重: 100,
			},
		}
	)
	db类.X设置组配置(configKey, configGroup)
	masterSlaveDB := g.DB类(configKey)
	单元测试类.C(t, func(t *单元测试类.T) {
		table := "table_" + uid类.X生成()
		createTableWithDb(masterSlaveDB.X切换数据库("master"), table)
		createTableWithDb(masterSlaveDB.X切换数据库("slave"), table)
		defer dropTableWithDb(masterSlaveDB.X切换数据库("master"), table)
		defer dropTableWithDb(masterSlaveDB.X切换数据库("slave"), table)

		// 向主库插入数据。
		array := 数组类.X创建(true)
		for i := 1; i <= TableSize; i++ {
			array.Append别名(g.Map{
				"id":          i,
				"passport":    fmt.Sprintf(`user_%d`, i),
				"password":    fmt.Sprintf(`pass_%d`, i),
				"nickname":    fmt.Sprintf(`name_%d`, i),
				"create_time": 时间类.X创建并从文本(CreateTime).String(),
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
