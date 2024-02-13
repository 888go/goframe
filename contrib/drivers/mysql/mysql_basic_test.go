// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

import (
	"context"
	"testing"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/test/gtest"
)

func Test_Instance(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := db类.X取单例对象("none")
		t.AssertNE(err, nil)

		db, err := db类.X取单例对象()
		t.AssertNil(err)

		err1 := db.X向主节点发送心跳()
		err2 := db.X向从节点发送心跳()
		t.Assert(err1, nil)
		t.Assert(err2, nil)
	})
}

func Test_Func_FormatSqlWithArgs(t *testing.T) {
	// mysql
	单元测试类.C(t, func(t *单元测试类.T) {
		var s string
		s = db类.X格式化Sql("select * from table where id>=? and sex=?", []interface{}{100, 1})
		t.Assert(s, "select * from table where id>=100 and sex=1")
	})
	// mssql
	单元测试类.C(t, func(t *单元测试类.T) {
		var s string
		s = db类.X格式化Sql("select * from table where id>=@p1 and sex=@p2", []interface{}{100, 1})
		t.Assert(s, "select * from table where id>=100 and sex=1")
	})
	// pgsql
	单元测试类.C(t, func(t *单元测试类.T) {
		var s string
		s = db类.X格式化Sql("select * from table where id>=$1 and sex=$2", []interface{}{100, 1})
		t.Assert(s, "select * from table where id>=100 and sex=1")
	})
	// oracle
	单元测试类.C(t, func(t *单元测试类.T) {
		var s string
		s = db类.X格式化Sql("select * from table where id>=:v1 and sex=:v2", []interface{}{100, 1})
		t.Assert(s, "select * from table where id>=100 and sex=1")
	})
}

func Test_Func_ToSQL(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		sql, err := db类.X捕捉最后一条SQL语句(ctx, func(ctx context.Context) error {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		array, err := db类.X捕捉SQL语句(ctx, func(ctx context.Context) error {
			value, err := db.X设置上下文并取副本(ctx).X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 1).X查询一条值()
			t.Assert(value, "name_1")
			return err
		})
		t.AssertNil(err)
		t.AssertGE(len(array), 1)
	})
}
