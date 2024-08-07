//go:build 屏蔽单元测试

// 版权所有 2019 gf 作者（https://github.com/gogf/gf）。保留所有权利。
//
// 此源代码形式受麻省理工学院（MIT）许可证的条款约束。
// 如果未随此文件一起分发MIT许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:47e609239e0cb2bc

package oracle_test

import (
	"context"
	"fmt"
	"strings"

	_ "github.com/sijms/go-ora/v2"

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
	TestSchema       = "XE"
)

const (
	TestDbIP   = "127.0.0.1"
	TestDbPort = "1521"
	TestDbUser = "system"
	TestDbPass = "oracle"
	TestDbName = "XE"
	TestDbType = "oracle"
)

func init() {
	node := gdb.ConfigNode{
		Host:             TestDbIP,
		Port:             TestDbPort,
		User:             TestDbUser,
		Pass:             TestDbPass,
		Name:             TestDbName,
		Type:             TestDbType,
		Role:             "master",
		Charset:          "utf8",
		Weight:           1,
		MaxIdleConnCount: 10,
		MaxOpenConnCount: 10,
	}

	nodeLink := gdb.ConfigNode{
		Type: TestDbType,
		Name: TestDbName,
		Link: fmt.Sprintf("%s:%s:%s@tcp(%s:%s)/%s",
			TestDbType, TestDbUser, TestDbPass, TestDbIP, TestDbPort, TestDbName,
		),
	}

	nodeErr := gdb.ConfigNode{
		Host:    TestDbIP,
		Port:    TestDbPort,
		User:    TestDbUser,
		Pass:    "1234",
		Name:    TestDbName,
		Type:    TestDbType,
		Role:    "master",
		Charset: "utf8",
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

func createTable(table ...string) (name string) {
	if len(table) > 0 {
		name = table[0]
	} else {
		name = fmt.Sprintf("user_%d", gtime.X取时间戳秒())
	}

	dropTable(name)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
	CREATE TABLE %s (
		ID NUMBER(10) NOT NULL,
		PASSPORT VARCHAR(45) NOT NULL,
		PASSWORD CHAR(32) NOT NULL,
		NICKNAME VARCHAR(45) NOT NULL,
		CREATE_TIME varchar(45),
	    SALARY NUMBER(18,2),
		PRIMARY KEY (ID))
	`, name)); err != nil {
		gtest.Fatal(err)
	}

	// db.Schema("test")
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
			"create_time": gtime.X创建并按当前时间().String(),
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
	count, err := db.X原生SQL查询字段计数(ctx, "SELECT COUNT(*) FROM USER_TABLES WHERE TABLE_NAME = ?", strings.ToUpper(table))
	if err != nil {
		gtest.Fatal(err)
	}

	if count == 0 {
		return
	}
	if _, err = db.X原生SQL执行(ctx, fmt.Sprintf("DROP TABLE %s", table)); err != nil {
		gtest.Fatal(err)
	}
}
