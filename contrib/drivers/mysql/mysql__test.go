// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

import (
	"context"
	"fmt"
	"testing"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
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
	db        db类.DB
	db2       db类.DB
	db3       db类.DB
	dbPrefix  db类.DB
	dbInvalid db类.DB
	ctx       = context.TODO()
)

func init() {
	nodeDefault := db类.X配置项{
		X自定义链接信息: fmt.Sprintf("mysql:root:%s@tcp(127.0.0.1:3306)/?loc=Local&parseTime=true", TestDbPass),
	}
	partitionDefault := db类.X配置项{
		X自定义链接信息:  fmt.Sprintf("mysql:root:%s@tcp(127.0.0.1:3307)/?loc=Local&parseTime=true", TestDbPass),
		X调试模式: true,
	}
	nodePrefix := db类.X配置项{
		X自定义链接信息: fmt.Sprintf("mysql:root:%s@tcp(127.0.0.1:3306)/?loc=Local&parseTime=true", TestDbPass),
	}
	nodePrefix.X表前缀 = TableNamePrefix1

	nodeInvalid := db类.X配置项{
		X自定义链接信息: fmt.Sprintf("mysql:root:%s@tcp(127.0.0.1:3307)/?loc=Local&parseTime=true", TestDbPass),
	}
	db类.X添加配置组节点("test", nodeDefault)
	db类.X添加配置组节点("prefix", nodePrefix)
	db类.X添加配置组节点("nodeinvalid", nodeInvalid)
	db类.X添加配置组节点("partition", partitionDefault)
	db类.X添加配置组节点(db类.DefaultGroupName, nodeDefault)

	// Default db.
	if r, err := db类.X创建DB对象并按配置组(); err != nil {
		单元测试类.Error(err)
	} else {
		db = r
	}
	schemaTemplate := "CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET UTF8"
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(schemaTemplate, TestSchema1)); err != nil {
		单元测试类.Error(err)
	}
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(schemaTemplate, TestSchema2)); err != nil {
		单元测试类.Error(err)
	}
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(schemaTemplate, TestPartitionDB)); err != nil {
		单元测试类.Error(err)
	}
	db = db.X切换数据库(TestSchema1)
	db2 = db.X切换数据库(TestSchema2)
	db3 = db.X切换数据库(TestPartitionDB)
	// Prefix db.
	if r, err := db类.X创建DB对象并按配置组("prefix"); err != nil {
		单元测试类.Error(err)
	} else {
		dbPrefix = r
	}
	if _, err := dbPrefix.X原生SQL执行(ctx, fmt.Sprintf(schemaTemplate, TestSchema1)); err != nil {
		单元测试类.Error(err)
	}
	if _, err := dbPrefix.X原生SQL执行(ctx, fmt.Sprintf(schemaTemplate, TestSchema2)); err != nil {
		单元测试类.Error(err)
	}
	dbPrefix = dbPrefix.X切换数据库(TestSchema1)

	// Invalid db.
	if r, err := db类.X创建DB对象并按配置组("nodeinvalid"); err != nil {
		单元测试类.Error(err)
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

func createTableWithDb(db db类.DB, table ...string) (name string) {
	if len(table) > 0 {
		name = table[0]
	} else {
		name = fmt.Sprintf(`%s_%d`, TableName, 时间类.X取时间戳纳秒())
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
		单元测试类.Fatal(err)
	}
	return name
}

func createInitTableWithDb(db db类.DB, table ...string) (name string) {
	name = createTableWithDb(db, table...)
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

	result, err := db.X插入(ctx, name, array.X取切片())
	单元测试类.AssertNil(err)

	n, e := result.RowsAffected()
	单元测试类.Assert(e, nil)
	单元测试类.Assert(n, TableSize)
	return
}

func dropTableWithDb(db db类.DB, table string) {
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf("DROP TABLE IF EXISTS `%s`", table)); err != nil {
		单元测试类.Error(err)
	}
}

func Test_PartitionTable(t *testing.T) {
	dropShopDBTable()
	createShopDBTable()
	insertShopDBData()

	// 延迟执行 dropShopDBTable() 函数
	单元测试类.C(t, func(t *单元测试类.T) {
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
		单元测试类.Fatal(err.Error())
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
		单元测试类.Error(err)
	}
}
func dropShopDBTable() {
	if _, err := db3.X原生SQL执行(ctx, "DROP TABLE IF EXISTS `dbx_order`"); err != nil {
		单元测试类.Error(err)
	}
}
