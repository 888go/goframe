//go:build 屏蔽单元测试

// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package pgsql_test

import (
	_ "github.com/888go/goframe/contrib/drivers/pgsql"

	"context"
	"fmt"

	garray "github.com/888go/goframe/container/garray"
	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

const (
	TableSize   = 10
	TablePrefix = "t_"
	SchemaName  = "test"
	CreateTime  = "2018-10-24 10:00:00"
)

var (
	db         gdb.DB
	configNode gdb.ConfigNode
	ctx        = context.TODO()
)

func init() {
	configNode = gdb.ConfigNode{
		Link: `pgsql:postgres:12345678@tcp(127.0.0.1:5432)`,
	}

	//pgsql只允许连接到指定的数据库。
	//因此，在使用ORM之前，你需要先创建pgsql数据库。
	// md5:9c6007198f234ad1
	gdb.X添加配置组节点(gdb.DefaultGroupName, configNode)
	if r, err := gdb.X创建DB对象(configNode); err != nil {
		gtest.Fatal(err)
	} else {
		db = r
	}

	if configNode.Name == "" {
		schemaTemplate := "SELECT 'CREATE DATABASE %s' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '%s')"
		if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(schemaTemplate, SchemaName, SchemaName)); err != nil {
			gtest.Error(err)
		}

		db = db.X切换数据库(SchemaName)
	} else {
		db = db.X切换数据库(configNode.Name)
	}

}

func createTable(table ...string) string {
	return createTableWithDb(db, table...)
}

func createInitTable(table ...string) string {
	return createInitTableWithDb(db, table...)
}

func createTableWithDb(db gdb.DB, table ...string) (name string) {
	if len(table) > 0 {
		name = table[0]
	} else {
		name = fmt.Sprintf(`%s_%d`, TablePrefix+"test", gtime.X取时间戳纳秒())
	}

	dropTableWithDb(db, name)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
		CREATE TABLE %s (
		   	id bigserial  NOT NULL,
		   	passport varchar(45) NOT NULL,
		   	password varchar(32) NOT NULL,
		   	nickname varchar(45) NOT NULL,
		   	create_time timestamp NOT NULL,
		   	PRIMARY KEY (id)
		) ;`, name,
	)); err != nil {
		gtest.Fatal(err)
	}
	return
}

func dropTable(table string) {
	dropTableWithDb(db, table)
}

func createInitTableWithDb(db gdb.DB, table ...string) (name string) {
	name = createTableWithDb(db, table...)
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

	result, err := db.X插入(ctx, name, array.X取切片())
	gtest.AssertNil(err)

	n, e := result.RowsAffected()
	gtest.Assert(e, nil)
	gtest.Assert(n, TableSize)
	return
}

func dropTableWithDb(db gdb.DB, table string) {
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf("DROP TABLE IF EXISTS %s", table)); err != nil {
		gtest.Error(err)
	}
}
