// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package sqlite_test

import (
	"fmt"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

var (
	db         db类.DB
	dbPrefix   db类.DB
	dbInvalid  db类.DB
	configNode db类.X配置项
	dbDir      = 文件类.X取临时目录("sqlite")
	ctx        = 上下文类.X创建()

	// Error
	ErrorSave = 错误类.X创建错误码(错误码类.CodeNotSupported, `Save operation is not supported by sqlite driver`)
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

	if err := 文件类.X创建目录(dbDir); err != nil {
		单元测试类.Error(err)
	}

	fmt.Println("init sqlite db dir: ", dbDir)

	dbFilePath := 文件类.X路径生成(dbDir, "test.db")
	configNode = db类.X配置项{
		X类型:    "sqlite",
		X自定义链接信息:    fmt.Sprintf(`sqlite::@file(%s)`, dbFilePath),
		X字符集: "utf8",
	}
	nodePrefix := configNode
	nodePrefix.X表前缀 = TableNamePrefix

	nodeInvalid := configNode

	db类.X添加配置组节点(DBGroupTest, configNode)
	db类.X添加配置组节点(DBGroupPrefix, nodePrefix)
	db类.X添加配置组节点(DBGroupInvalid, nodeInvalid)
	db类.X添加配置组节点(db类.DefaultGroupName, configNode)

	// Default db.
	if r, err := db类.X创建DB对象并按配置组(); err != nil {
		单元测试类.Error(err)
	} else {
		db = r
	}

	// Prefix db.
	if r, err := db类.X创建DB对象并按配置组(DBGroupPrefix); err != nil {
		单元测试类.Error(err)
	} else {
		dbPrefix = r
	}

	// Invalid db.
	if r, err := db类.X创建DB对象并按配置组(DBGroupInvalid); err != nil {
		单元测试类.Error(err)
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

func createTableWithDb(db db类.DB, table ...string) (name string) {
	if len(table) > 0 {
		name = table[0]
	} else {
		name = fmt.Sprintf(`%s_%d`, TableName, 时间类.X取时间戳纳秒())
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
		单元测试类.Fatal(err)
	}

	return
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
