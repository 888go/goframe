//go:build 屏蔽单元测试

// 版权所有 2019 gf 作者（https://github.com/gogf/gf）。保留所有权利。
//
// 此源代码形式受麻省理工学院（MIT）许可证的条款约束。
// 如果未随此文件一起分发MIT许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:47e609239e0cb2bc

package mssql_test

import (
	"context"
	"fmt"

	garray "github.com/888go/goframe/container/garray"
	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

var (
	db     gdb.DB
	dblink gdb.DB
	dbErr  gdb.DB
	ctx    context.Context
)

const (
	TableSize        = 10
	TableName        = "t_user"
	TestSchema1      = "test1"
	TestSchema2      = "test2"
	TableNamePrefix1 = "gf_"
	TestDbUser       = "sa"
	TestDbPass       = "LoremIpsum86"
	CreateTime       = "2018-10-24 10:00:00"
)

func init() {
	node := gdb.ConfigNode{
		Host:             "127.0.0.1",
		Port:             "1433",
		User:             TestDbUser,
		Pass:             TestDbPass,
		Name:             "test",
		Type:             "mssql",
		Role:             "master",
		Charset:          "utf8",
		Weight:           1,
		MaxIdleConnCount: 10,
		MaxOpenConnCount: 10,
	}

	nodeLink := gdb.ConfigNode{
		Type: "mssql",
		Name: "test",
		Link: fmt.Sprintf(
			"mssql:%s:%s@tcp(%s:%s)/%s?encrypt=disable",
			node.User, node.Pass, node.Host, node.Port, node.Name,
		),
	}

	nodeErr := gdb.ConfigNode{
		Type: "mssql",
		Link: fmt.Sprintf("user id=%s;password=%s;server=%s;port=%s;database=%s;encrypt=disable",
			node.User, "node.Pass", node.Host, node.Port, node.Name),
	}

	gdb.X添加配置组节点(gdb.DefaultGroupName, node)
	if r, err := gdb.X创建DB对象(node); err != nil {
		gtest.Fatal(err)
	} else {
		db = r
	}

	gdb.X添加配置组节点("dblink", nodeLink)
	if r, err := gdb.X创建DB对象(nodeLink); err != nil {
		gtest.Fatal(err)
	} else {
		dblink = r
	}

	gdb.X添加配置组节点("dbErr", nodeErr)
	if r, err := gdb.X创建DB对象(nodeErr); err != nil {
		gtest.Fatal(err)
	} else {
		dbErr = r
	}

	ctx = context.Background()
}

func createTable(table ...string) (name string) {
	if len(table) > 0 {
		name = table[0]
	} else {
		name = fmt.Sprintf("user_%d", gtime.X取时间戳秒())
	}

	dropTable(name)

	if _, err := db.X原生SQL执行(context.Background(), fmt.Sprintf(`
		IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='%s' and xtype='U')
		CREATE TABLE %s (
		ID numeric(10,0) NOT NULL,
		PASSPORT VARCHAR(45)  NULL,
		PASSWORD VARCHAR(32)  NULL,
		NICKNAME VARCHAR(45)  NULL,
		CREATE_TIME datetime NULL,
		CREATED_AT datetimeoffset NULL,
		UPDATED_AT datetimeoffset NULL,
		PRIMARY KEY (ID))
	`, name, name)); err != nil {
		gtest.Fatal(err)
	}

	db.X切换数据库("test")
	return
}

func createInitTable(table ...string) (name string) {
	name = createTable(table...)
	array := garray.X创建(true)
	for i := 1; i <= TableSize; i++ {
		array.Append别名(g.Map{
			"id":          i,
			"passport":    fmt.Sprintf(`user_%d`, i),
			"password":    fmt.Sprintf(`pass_%d`, i),
			"nickname":    fmt.Sprintf(`name_%d`, i),
			"create_time": gtime.X创建并按当前时间(),
		})
	}
	result, err := db.X插入(context.Background(), name, array.X取切片())
	gtest.Assert(err, nil)

	n, e := result.RowsAffected()
	gtest.Assert(e, nil)
	gtest.Assert(n, TableSize)
	return
}

func dropTable(table string) {
	if _, err := db.X原生SQL执行(context.Background(), fmt.Sprintf(`
		IF EXISTS (SELECT * FROM sysobjects WHERE name='%s' and xtype='U')
		DROP TABLE %s
	`, table, table)); err != nil {
		gtest.Fatal(err)
	}
}
