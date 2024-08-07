// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"context"
	"fmt"
	"testing"

	garray "github.com/888go/goframe/container/garray"
	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

const (
	TableSize        = 10
	TableName        = "user"
	TestSchema1      = "test1"
	TestSchema2      = "test2"
	TestPartitionDB  = "test3"
	TableNamePrefix1 = "gf_"
	TestDbUser       = "root"
	TestDbPass       = "12345678"
	CreateTime       = "2018-10-24 10:00:00"
)

var (
	db        gdb.DB
	db2       gdb.DB
	db3       gdb.DB
	dbPrefix  gdb.DB
	dbInvalid gdb.DB
	ctx       = context.TODO()
)

func init() {
	nodeDefault := gdb.ConfigNode{
		Link: fmt.Sprintf("mysql:root:%s@tcp(127.0.0.1:3306)/?loc=Local&parseTime=true", TestDbPass),
	}
	partitionDefault := gdb.ConfigNode{
		Link:  fmt.Sprintf("mysql:root:%s@tcp(127.0.0.1:3307)/?loc=Local&parseTime=true", TestDbPass),
		Debug: true,
	}
	nodePrefix := gdb.ConfigNode{
		Link: fmt.Sprintf("mysql:root:%s@tcp(127.0.0.1:3306)/?loc=Local&parseTime=true", TestDbPass),
	}
	nodePrefix.Prefix = TableNamePrefix1

	nodeInvalid := gdb.ConfigNode{
		Link: fmt.Sprintf("mysql:root:%s@tcp(127.0.0.1:3307)/?loc=Local&parseTime=true", TestDbPass),
	}
	gdb.X添加配置组节点("test", nodeDefault)
	gdb.X添加配置组节点("prefix", nodePrefix)
	gdb.X添加配置组节点("nodeinvalid", nodeInvalid)
	gdb.X添加配置组节点("partition", partitionDefault)
	gdb.X添加配置组节点(gdb.DefaultGroupName, nodeDefault)

	// Default db.
	if r, err := gdb.X创建DB对象并按配置组(); err != nil {
		gtest.Error(err)
	} else {
		db = r
	}
	schemaTemplate := "CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET UTF8"
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(schemaTemplate, TestSchema1)); err != nil {
		gtest.Error(err)
	}
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(schemaTemplate, TestSchema2)); err != nil {
		gtest.Error(err)
	}
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(schemaTemplate, TestPartitionDB)); err != nil {
		gtest.Error(err)
	}
	db = db.X切换数据库(TestSchema1)
	db2 = db.X切换数据库(TestSchema2)
	db3 = db.X切换数据库(TestPartitionDB)
	// Prefix db.
	if r, err := gdb.X创建DB对象并按配置组("prefix"); err != nil {
		gtest.Error(err)
	} else {
		dbPrefix = r
	}
	if _, err := dbPrefix.X原生SQL执行(ctx, fmt.Sprintf(schemaTemplate, TestSchema1)); err != nil {
		gtest.Error(err)
	}
	if _, err := dbPrefix.X原生SQL执行(ctx, fmt.Sprintf(schemaTemplate, TestSchema2)); err != nil {
		gtest.Error(err)
	}
	dbPrefix = dbPrefix.X切换数据库(TestSchema1)

	// Invalid db.
	if r, err := gdb.X创建DB对象并按配置组("nodeinvalid"); err != nil {
		gtest.Error(err)
	} else {
		dbInvalid = r
	}
	dbInvalid = dbInvalid.X切换数据库(TestSchema1)
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
	        id          int(10) unsigned NOT NULL AUTO_INCREMENT,
	        passport    varchar(45) NULL,
	        password    char(32) NULL,
	        nickname    varchar(45) NULL,
	        create_time timestamp(6) NULL,
	        PRIMARY KEY (id)
	    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	    `, name,
	)); err != nil {
		gtest.Fatal(err)
	}
	return name
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

func Test_PartitionTable(t *testing.T) {
	dropShopDBTable()
	createShopDBTable()
	insertShopDBData()

	//defer dropShopDBTable()
	gtest.C(t, func(t *gtest.T) {
		data, err := db3.X设置上下文并取副本(ctx).X创建Model对象("dbx_order").X设置分区名称("p3", "p4").X查询()
		t.AssertNil(err)
		dataLen := len(data)
		t.Assert(dataLen, 5)
		data, err = db3.X设置上下文并取副本(ctx).X创建Model对象("dbx_order").X设置分区名称("p3").X查询()
		t.AssertNil(err)
		dataLen = len(data)
		t.Assert(dataLen, 5)
	})
}
func createShopDBTable() {
	sql := `CREATE TABLE dbx_order (
  id int(11) NOT NULL,
  sales_date date DEFAULT NULL,
  amount decimal(10,2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
PARTITION BY RANGE (YEAR(sales_date))
(PARTITION p1 VALUES LESS THAN (2020) ENGINE = InnoDB,
 PARTITION p2 VALUES LESS THAN (2021) ENGINE = InnoDB,
 PARTITION p3 VALUES LESS THAN (2022) ENGINE = InnoDB,
 PARTITION p4 VALUES LESS THAN MAXVALUE ENGINE = InnoDB);`
	_, err := db3.X原生SQL执行(ctx, sql)
	if err != nil {
		gtest.Fatal(err.Error())
	}
}
func insertShopDBData() {
	data := g.Slice别名{}
	year := 2020
	for i := 1; i <= 5; i++ {
		year++
		data = append(data, g.Map{
			"id":         i,
			"sales_date": fmt.Sprintf("%d-09-21", year),
			"amount":     fmt.Sprintf("1%d.21", i),
		})
	}
	_, err := db3.X创建Model对象("dbx_order").X设置上下文并取副本(ctx).X设置数据(data).X插入()
	if err != nil {
		gtest.Error(err)
	}
}
func dropShopDBTable() {
	if _, err := db3.X原生SQL执行(ctx, "DROP TABLE IF EXISTS `dbx_order`"); err != nil {
		gtest.Error(err)
	}
}
