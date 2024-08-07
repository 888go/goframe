// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package sqlite_test

import (
	"fmt"

	garray "github.com/888go/goframe/container/garray"
	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gctx "github.com/888go/goframe/os/gctx"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

var (
	db         gdb.DB
	dbPrefix   gdb.DB
	dbInvalid  gdb.DB
	configNode gdb.ConfigNode
	dbDir      = gfile.X取临时目录("sqlite")
	ctx        = gctx.X创建()
)

const (
	TableSize               = 10
	TableName               = "user"
	TableNameWhichIsKeyword = "group"
	TestSchema1             = "test1"
	TestSchema2             = "test2"
	TableNamePrefix         = "gf_"
	CreateTime              = "2018-10-24 10:00:00"
	DBGroupTest             = "test"
	DBGroupPrefix           = "prefix"
	DBGroupInvalid          = "invalid"
)

func init() {
	fmt.Println("init sqlite db start")

	if err := gfile.X创建目录(dbDir); err != nil {
		gtest.Error(err)
	}

	fmt.Println("init sqlite db dir: ", dbDir)

	dbFilePath := gfile.X路径生成(dbDir, "test.db")
	configNode = gdb.ConfigNode{
		Type:    "sqlite",
		Link:    fmt.Sprintf(`sqlite::@file(%s)`, dbFilePath),
		Charset: "utf8",
	}
	nodePrefix := configNode
	nodePrefix.Prefix = TableNamePrefix

	nodeInvalid := configNode

	gdb.X添加配置组节点(DBGroupTest, configNode)
	gdb.X添加配置组节点(DBGroupPrefix, nodePrefix)
	gdb.X添加配置组节点(DBGroupInvalid, nodeInvalid)
	gdb.X添加配置组节点(gdb.DefaultGroupName, configNode)

	// Default db.
	if r, err := gdb.X创建DB对象并按配置组(); err != nil {
		gtest.Error(err)
	} else {
		db = r
	}

	// Prefix db.
	if r, err := gdb.X创建DB对象并按配置组(DBGroupPrefix); err != nil {
		gtest.Error(err)
	} else {
		dbPrefix = r
	}

	// Invalid db.
	if r, err := gdb.X创建DB对象并按配置组(DBGroupInvalid); err != nil {
		gtest.Error(err)
	} else {
		dbInvalid = r
	}

	fmt.Println("init sqlite db finish")
}

func createTable(table ...string) string {
	return createTableWithDb(db, table...)
}

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

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	CREATE TABLE %s (
		id          INTEGER       PRIMARY KEY AUTOINCREMENT
									UNIQUE
									NOT NULL,
		passport    VARCHAR(45)  NOT NULL
									DEFAULT passport,
		password    VARCHAR(128) NOT NULL
									DEFAULT password,
		nickname    VARCHAR(45),
		create_time DATETIME
	);
	`, db.X取Core对象().X底层QuoteWord(name),
	)); err != nil {
		gtest.Fatal(err)
	}

	return
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
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf("DROP TABLE IF EXISTS `%s`", table)); err != nil {
		gtest.Error(err)
	}
}
