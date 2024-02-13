// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

func Test_Model_Sharding_Table(t *testing.T) {
	var (
		table1 = 时间类.X取文本时间戳纳秒() + "_table1"
		table2 = 时间类.X取文本时间戳纳秒() + "_table2"
	)
	createTable(table1)
	defer dropTable(table1)
	createTable(table2)
	defer dropTable(table2)

	shardingModel := db.X创建Model对象(table1).Hook(db类.HookHandler{
		Select: func(ctx context.Context, in *db类.HookSelectInput) (result db类.Result, err error) {
			in.Table = table2
			return in.Next(ctx)
		},
		Insert: func(ctx context.Context, in *db类.HookInsertInput) (result sql.Result, err error) {
			in.Table = table2
			return in.Next(ctx)
		},
		Update: func(ctx context.Context, in *db类.HookUpdateInput) (result sql.Result, err error) {
			in.Table = table2
			return in.Next(ctx)
		},
		Delete: func(ctx context.Context, in *db类.HookDeleteInput) (result sql.Result, err error) {
			in.Table = table2
			return in.Next(ctx)
		},
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := shardingModel.X插入(g.Map{
			"id":          1,
			"passport":    fmt.Sprintf(`user_%d`, 1),
			"password":    fmt.Sprintf(`pass_%d`, 1),
			"nickname":    fmt.Sprintf(`name_%d`, 1),
			"create_time": 时间类.X创建并从文本(CreateTime).String(),
		})
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var count int
		count, err = shardingModel.X查询行数()
		t.AssertNil(err)
		t.Assert(count, 1)

		count, err = db.X创建Model对象(table1).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db.X创建Model对象(table2).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := shardingModel.X条件(g.Map{
			"id": 1,
		}).X设置数据(g.Map{
			"passport": fmt.Sprintf(`user_%d`, 2),
			"password": fmt.Sprintf(`pass_%d`, 2),
			"nickname": fmt.Sprintf(`name_%d`, 2),
		}).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var (
			count int
			where = g.Map{"passport": fmt.Sprintf(`user_%d`, 2)}
		)
		count, err = shardingModel.X条件(where).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 1)

		count, err = db.X创建Model对象(table1).X条件(where).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db.X创建Model对象(table2).X条件(where).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := shardingModel.X条件(g.Map{
			"id": 1,
		}).X删除()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var count int
		count, err = shardingModel.X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db.X创建Model对象(table1).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db.X创建Model对象(table2).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)
	})
}

func Test_Model_Sharding_Schema(t *testing.T) {
	var (
		table = 时间类.X取文本时间戳纳秒() + "_table"
	)
	createTableWithDb(db, table)
	defer dropTableWithDb(db, table)
	createTableWithDb(db2, table)
	defer dropTableWithDb(db2, table)

	shardingModel := db.X创建Model对象(table).Hook(db类.HookHandler{
		Select: func(ctx context.Context, in *db类.HookSelectInput) (result db类.Result, err error) {
			in.Table = table
			in.Schema = db2.X取默认数据库名称()
			return in.Next(ctx)
		},
		Insert: func(ctx context.Context, in *db类.HookInsertInput) (result sql.Result, err error) {
			in.Table = table
			in.Schema = db2.X取默认数据库名称()
			return in.Next(ctx)
		},
		Update: func(ctx context.Context, in *db类.HookUpdateInput) (result sql.Result, err error) {
			in.Table = table
			in.Schema = db2.X取默认数据库名称()
			return in.Next(ctx)
		},
		Delete: func(ctx context.Context, in *db类.HookDeleteInput) (result sql.Result, err error) {
			in.Table = table
			in.Schema = db2.X取默认数据库名称()
			return in.Next(ctx)
		},
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := shardingModel.X插入(g.Map{
			"id":          1,
			"passport":    fmt.Sprintf(`user_%d`, 1),
			"password":    fmt.Sprintf(`pass_%d`, 1),
			"nickname":    fmt.Sprintf(`name_%d`, 1),
			"create_time": 时间类.X创建并从文本(CreateTime).String(),
		})
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var count int
		count, err = shardingModel.X查询行数()
		t.AssertNil(err)
		t.Assert(count, 1)

		count, err = db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db2.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := shardingModel.X条件(g.Map{
			"id": 1,
		}).X设置数据(g.Map{
			"passport": fmt.Sprintf(`user_%d`, 2),
			"password": fmt.Sprintf(`pass_%d`, 2),
			"nickname": fmt.Sprintf(`name_%d`, 2),
		}).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var (
			count int
			where = g.Map{"passport": fmt.Sprintf(`user_%d`, 2)}
		)
		count, err = shardingModel.X条件(where).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 1)

		count, err = db.X创建Model对象(table).X条件(where).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db2.X创建Model对象(table).X条件(where).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := shardingModel.X条件(g.Map{
			"id": 1,
		}).X删除()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var count int
		count, err = shardingModel.X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db2.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)
	})
}
