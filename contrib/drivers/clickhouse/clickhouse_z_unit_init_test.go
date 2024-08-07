//go:build 屏蔽单元测试

// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package clickhouse_test

import (
	"context"
	"fmt"

	garray "github.com/888go/goframe/container/garray"
	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

const (
	TableSize = 10
	TableName = "user"
)

var (
	db  gdb.DB
	ctx = context.TODO()
)

func init() {
	node := gdb.ConfigNode{
		Host:  "127.0.0.1",
		Port:  "9000",
		User:  "default",
		Name:  "default",
		Type:  "clickhouse",
		Debug: false,
	}
	var err error
	db, err = gdb.X创建DB对象(node)
	gtest.AssertNil(err)
	gtest.AssertNil(db.X向主节点发送心跳())
}

// create table
func createTable(table ...string) string {
	return createTableWithDb(db, table...)
}

// 创建表并插入初始数据. md5:a4bda4ca9da339e1
func createInitTable(table ...string) string {
	return createInitTableWithDb(db, table...)
}

func dropTable(table string) {
	dropTableWithDb(db, table)
}

func createTableWithDb(db gdb.DB, table ...string) (name string) {
	if len(table) > 0 {
		name = table[0]
	} else {
		name = fmt.Sprintf(`%s_%d`, TableName, gtime.X取时间戳纳秒())
	}
	dropTableWithDb(db, name)

	_, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
		CREATE TABLE %s (
		   id bigint unsigned NOT NULL,
		   passport varchar(45),
		   password char(32) NOT NULL,
		   nickname varchar(45) NOT NULL,
		   create_time datetime NOT NULL,
		   PRIMARY KEY (id)
		) ENGINE = MergeTree()
		ORDER BY id ;`,
		name,
	))
	if err != nil {
		gtest.Fatal(err)
	}

	return
}

func createInitTableWithDb(db gdb.DB, table ...string) (name string) {
	name = createTableWithDb(db, table...)
	array := garray.X创建(true)
	for i := 1; i <= TableSize; i++ {
		array.Append别名(g.Map{
			"id":          uint64(i),
			"passport":    fmt.Sprintf(`user_%d`, i),
			"password":    fmt.Sprintf(`pass_%d`, i),
			"nickname":    fmt.Sprintf(`name_%d`, i),
			"create_time": gtime.X创建并按当前时间(),
		})
	}

	result, err := db.X插入(ctx, name, array.X取切片())
	gtest.AssertNil(err)

	if result != nil {
		n, e := result.RowsAffected()
		gtest.Assert(e, nil)
		gtest.Assert(n, TableSize)
	}
	return
}

func dropTableWithDb(db gdb.DB, table string) {
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf("DROP TABLE IF EXISTS `%s`", table)); err != nil {
		gtest.Error(err)
	}
}
