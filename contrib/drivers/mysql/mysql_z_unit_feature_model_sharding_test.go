// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Model_Sharding_Table(t *testing.T) {
	var (
		table1 = gtime.X取文本时间戳纳秒() + "_table1"
		table2 = gtime.X取文本时间戳纳秒() + "_table2"
	)
	createTable(table1)
	defer dropTable(table1)
	createTable(table2)
	defer dropTable(table2)

	shardingModel := db.X创建Model对象(table1).Hook(gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			in.Table = table2
			return in.Next(ctx)
		},
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			in.Table = table2
			return in.Next(ctx)
		},
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			in.Table = table2
			return in.Next(ctx)
		},
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
			in.Table = table2
			return in.Next(ctx)
		},
	})
	gtest.C(t, func(t *gtest.T) {
		r, err := shardingModel.X插入(g.Map{
			"id":          1,
			"passport":    fmt.Sprintf(`user_%d`, 1),
			"password":    fmt.Sprintf(`pass_%d`, 1),
			"nickname":    fmt.Sprintf(`name_%d`, 1),
			"create_time": gtime.X创建并从文本(CreateTime).String(),
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

	gtest.C(t, func(t *gtest.T) {
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

	gtest.C(t, func(t *gtest.T) {
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
		table = gtime.X取文本时间戳纳秒() + "_table"
	)
	createTableWithDb(db, table)
	defer dropTableWithDb(db, table)
	createTableWithDb(db2, table)
	defer dropTableWithDb(db2, table)

	shardingModel := db.X创建Model对象(table).Hook(gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			in.Table = table
			in.Schema = db2.X取默认数据库名称()
			return in.Next(ctx)
		},
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			in.Table = table
			in.Schema = db2.X取默认数据库名称()
			return in.Next(ctx)
		},
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			in.Table = table
			in.Schema = db2.X取默认数据库名称()
			return in.Next(ctx)
		},
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
			in.Table = table
			in.Schema = db2.X取默认数据库名称()
			return in.Next(ctx)
		},
	})
	gtest.C(t, func(t *gtest.T) {
		r, err := shardingModel.X插入(g.Map{
			"id":          1,
			"passport":    fmt.Sprintf(`user_%d`, 1),
			"password":    fmt.Sprintf(`pass_%d`, 1),
			"nickname":    fmt.Sprintf(`name_%d`, 1),
			"create_time": gtime.X创建并从文本(CreateTime).String(),
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

	gtest.C(t, func(t *gtest.T) {
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

	gtest.C(t, func(t *gtest.T) {
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
