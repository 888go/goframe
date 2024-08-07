// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"context"
	"testing"

	gdb "github.com/888go/goframe/database/gdb"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Instance(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_, err := gdb.X取单例对象("none")
		t.AssertNE(err, nil)

		db, err := gdb.X取单例对象()
		t.AssertNil(err)

		err1 := db.X向主节点发送心跳()
		err2 := db.X向从节点发送心跳()
		t.Assert(err1, nil)
		t.Assert(err2, nil)
	})
}

func Test_Func_FormatSqlWithArgs(t *testing.T) {
	// mysql
	gtest.C(t, func(t *gtest.T) {
		var s string
		s = gdb.X格式化Sql("select * from table where id>=? and sex=?", []interface{}{100, 1})
		t.Assert(s, "select * from table where id>=100 and sex=1")
	})
	// mssql
	gtest.C(t, func(t *gtest.T) {
		var s string
		s = gdb.X格式化Sql("select * from table where id>=@p1 and sex=@p2", []interface{}{100, 1})
		t.Assert(s, "select * from table where id>=100 and sex=1")
	})
	// pgsql
	gtest.C(t, func(t *gtest.T) {
		var s string
		s = gdb.X格式化Sql("select * from table where id>=$1 and sex=$2", []interface{}{100, 1})
		t.Assert(s, "select * from table where id>=100 and sex=1")
	})
	// oracle
	gtest.C(t, func(t *gtest.T) {
		var s string
		s = gdb.X格式化Sql("select * from table where id>=:v1 and sex=:v2", []interface{}{100, 1})
		t.Assert(s, "select * from table where id>=100 and sex=1")
	})
}

func Test_Func_ToSQL(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		sql, err := gdb.X捕捉最后一条SQL语句(ctx, func(ctx context.Context) error {
			value, err := db.X设置上下文并取副本(ctx).X创建Model对象(TableName).X字段保留过滤("nickname").X条件("id", 1).X查询一条值()
			t.Assert(value, nil)
			return err
		})
		t.AssertNil(err)
		t.Assert(sql, "SELECT `nickname` FROM `user` WHERE `id`=1 LIMIT 1")
	})
}

func Test_Func_CatchSQL(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		array, err := gdb.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			value, err := db.X设置上下文并取副本(ctx).X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 1).X查询一条值()
			t.Assert(value, "name_1")
			return err
		})
		t.AssertNil(err)
		t.AssertGE(len(array), 1)
	})
}
