//go:build 屏蔽单元测试

// 版权所有 2019 gf 作者（https://github.com/gogf/gf）。保留所有权利。
//
// 此源代码形式受麻省理工学院（MIT）许可证的条款约束。
// 如果未随此文件一起分发MIT许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:47e609239e0cb2bc

package dm_test

import (
	"context"
	"fmt"
	"strings"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

var (
	db        gdb.DB
	dblink    gdb.DB
	dbErr     gdb.DB
	ctx       context.Context
	TableSize = 10
)

const (
	TestDBHost  = "127.0.0.1"
	TestDBPort  = "5236"
	TestDBUser  = "SYSDBA"
	TestDBPass  = "SYSDBA001"
	TestDBName  = "SYSDBA"
	TestDBType  = "dm"
	TestCharset = "utf8"
)

type User struct {
	ID          int64     `orm:"id"`
	AccountName string    `orm:"account_name"`
	PwdReset    int64     `orm:"pwd_reset"`
	AttrIndex   int64     `orm:"attr_index"`
	Enabled     int64     `orm:"enabled"`
	Deleted     int64     `orm:"deleted"`
	CreatedBy   string    `orm:"created_by"`
	CreatedTime time.Time `orm:"created_time"`
	UpdatedBy   string    `orm:"updated_by"`
	UpdatedTime time.Time `orm:"updated_time"`
}

func init() {
	node := gdb.ConfigNode{
		Host:             TestDBHost,
		Port:             TestDBPort,
		User:             TestDBUser,
		Pass:             TestDBPass,
		Name:             TestDBName,
		Type:             TestDBType,
		Role:             "master",
		Charset:          TestCharset,
		Weight:           1,
		MaxIdleConnCount: 10,
		MaxOpenConnCount: 10,
		CreatedAt:        "created_time",
		UpdatedAt:        "updated_time",
	}

	// todo
	nodeLink := gdb.ConfigNode{
		Type: TestDBType,
		Name: TestDBName,
		Link: fmt.Sprintf(
			"dm:%s:%s@tcp(%s:%s)/%s?charset=%s",
			TestDBUser, TestDBPass, TestDBHost, TestDBPort, TestDBName, TestCharset,
		),
	}

	nodeErr := gdb.ConfigNode{
		Host:    TestDBHost,
		Port:    TestDBPort,
		User:    TestDBUser,
		Pass:    "1234",
		Name:    TestDBName,
		Type:    TestDBType,
		Role:    "master",
		Charset: TestCharset,
		Weight:  1,
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

func dropTable(table string) {
	count, err := db.X原生SQL查询字段计数(
		ctx,
		"SELECT COUNT(*) FROM all_tables WHERE owner = ? And table_name= ?", TestDBName, strings.ToUpper(table),
	)
	if err != nil {
		gtest.Fatal(err)
	}

	if count == 0 {
		return
	}
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf("DROP TABLE %s", table)); err != nil {
		gtest.Fatal(err)
	}
}

func createTable(table ...string) (name string) {
	if len(table) > 0 {
		name = table[0]
	} else {
		name = fmt.Sprintf("random_%d", gtime.X取时间戳秒())
	}

	dropTable(name)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	CREATE TABLE "%s"
(
"ID" BIGINT NOT NULL,
"ACCOUNT_NAME" VARCHAR(128) DEFAULT '' NOT NULL,
"PWD_RESET" TINYINT DEFAULT 0 NOT NULL,
"ENABLED" INT DEFAULT 1 NOT NULL,
"DELETED" INT DEFAULT 0 NOT NULL,
"ATTR_INDEX" INT DEFAULT 0 ,
"CREATED_BY" VARCHAR(32) DEFAULT '' NOT NULL,
"CREATED_TIME" TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP() NOT NULL,
"UPDATED_BY" VARCHAR(32) DEFAULT '' NOT NULL,
"UPDATED_TIME" TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP() NOT NULL,
NOT CLUSTER PRIMARY KEY("ID")) STORAGE(ON "MAIN", CLUSTERBTR) ;
	`, name)); err != nil {
		gtest.Fatal(err)
	}

	return
}

func createInitTable(table ...string) (name string) {
	name = createTable(table...)
	array := garray.X创建(true)
	for i := 1; i <= TableSize; i++ {
		array.Append别名(g.Map{
			"id":           i,
			"account_name": fmt.Sprintf(`name_%d`, i),
			"pwd_reset":    0,
			"attr_index":   i,
			"create_time":  gtime.X创建并按当前时间().String(),
		})
	}
	result, err := db.X切换数据库(TestDBName).X插入(context.Background(), name, array.X取切片())
	gtest.Assert(err, nil)

	n, e := result.RowsAffected()
	gtest.Assert(e, nil)
	gtest.Assert(n, TableSize)
	return
}

func createTableFalse(table ...string) (name string, err error) {
	if len(table) > 0 {
		name = table[0]
	} else {
		name = fmt.Sprintf("random_%d", gtime.X取时间戳秒())
	}

	dropTable(name)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	CREATE TABLE "%s"
(
"ID" BIGINT NOT NULL,
"ACCOUNT_NAME" VARCHAR(128) DEFAULT '' NOT NULL,
"PWD_RESET" TINYINT DEFAULT 0 NOT NULL,
"ENABLED" INT DEFAULT 1 NOT NULL,
"DELETED" INT DEFAULT 0 NOT NULL,
"INDEX" INT DEFAULT 0 ,
"ATTR_INDEX" INT DEFAULT 0 ,
"CREATED_BY" VARCHAR(32) DEFAULT '' NOT NULL,
"CREATED_TIME" TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP() NOT NULL,
"UPDATED_BY" VARCHAR(32) DEFAULT '' NOT NULL,
"UPDATED_TIME" TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP() NOT NULL,
NOT CLUSTER PRIMARY KEY("ID")) STORAGE(ON "MAIN", CLUSTERBTR) ;
	`, name)); err != nil {
		// gtest.Fatal(err)
		return name, fmt.Errorf("createTableFalse")
	}

	return name, nil
}
