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
		table1 = gtime.TimestampNanoStr() + "_table1"
		table2 = gtime.TimestampNanoStr() + "_table2"
	)
	createTable(table1)
	defer dropTable(table1)
	createTable(table2)
	defer dropTable(table2)

	shardingModel := db.Model(table1).Hook(gdb.HookHandler{
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
		r, err := shardingModel.Insert(g.Map{
			"id":          1,
			"passport":    fmt.Sprintf(`user_%d`, 1),
			"password":    fmt.Sprintf(`pass_%d`, 1),
			"nickname":    fmt.Sprintf(`name_%d`, 1),
			"create_time": gtime.NewFromStr(CreateTime).String(),
		})
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var count int
		count, err = shardingModel.Count()
		t.AssertNil(err)
		t.Assert(count, 1)

		count, err = db.Model(table1).Count()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db.Model(table2).Count()
		t.AssertNil(err)
		t.Assert(count, 1)
	})

	gtest.C(t, func(t *gtest.T) {
		r, err := shardingModel.Where(g.Map{
			"id": 1,
		}).Data(g.Map{
			"passport": fmt.Sprintf(`user_%d`, 2),
			"password": fmt.Sprintf(`pass_%d`, 2),
			"nickname": fmt.Sprintf(`name_%d`, 2),
		}).Update()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var (
			count int
			where = g.Map{"passport": fmt.Sprintf(`user_%d`, 2)}
		)
		count, err = shardingModel.Where(where).Count()
		t.AssertNil(err)
		t.Assert(count, 1)

		count, err = db.Model(table1).Where(where).Count()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db.Model(table2).Where(where).Count()
		t.AssertNil(err)
		t.Assert(count, 1)
	})

	gtest.C(t, func(t *gtest.T) {
		r, err := shardingModel.Where(g.Map{
			"id": 1,
		}).Delete()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var count int
		count, err = shardingModel.Count()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db.Model(table1).Count()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db.Model(table2).Count()
		t.AssertNil(err)
		t.Assert(count, 0)
	})
}

func Test_Model_Sharding_Schema(t *testing.T) {
	var (
		table = gtime.TimestampNanoStr() + "_table"
	)
	createTableWithDb(db, table)
	defer dropTableWithDb(db, table)
	createTableWithDb(db2, table)
	defer dropTableWithDb(db2, table)

	shardingModel := db.Model(table).Hook(gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			in.Table = table
			in.Schema = db2.GetSchema()
			return in.Next(ctx)
		},
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			in.Table = table
			in.Schema = db2.GetSchema()
			return in.Next(ctx)
		},
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			in.Table = table
			in.Schema = db2.GetSchema()
			return in.Next(ctx)
		},
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
			in.Table = table
			in.Schema = db2.GetSchema()
			return in.Next(ctx)
		},
	})
	gtest.C(t, func(t *gtest.T) {
		r, err := shardingModel.Insert(g.Map{
			"id":          1,
			"passport":    fmt.Sprintf(`user_%d`, 1),
			"password":    fmt.Sprintf(`pass_%d`, 1),
			"nickname":    fmt.Sprintf(`name_%d`, 1),
			"create_time": gtime.NewFromStr(CreateTime).String(),
		})
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var count int
		count, err = shardingModel.Count()
		t.AssertNil(err)
		t.Assert(count, 1)

		count, err = db.Model(table).Count()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db2.Model(table).Count()
		t.AssertNil(err)
		t.Assert(count, 1)
	})

	gtest.C(t, func(t *gtest.T) {
		r, err := shardingModel.Where(g.Map{
			"id": 1,
		}).Data(g.Map{
			"passport": fmt.Sprintf(`user_%d`, 2),
			"password": fmt.Sprintf(`pass_%d`, 2),
			"nickname": fmt.Sprintf(`name_%d`, 2),
		}).Update()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var (
			count int
			where = g.Map{"passport": fmt.Sprintf(`user_%d`, 2)}
		)
		count, err = shardingModel.Where(where).Count()
		t.AssertNil(err)
		t.Assert(count, 1)

		count, err = db.Model(table).Where(where).Count()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db2.Model(table).Where(where).Count()
		t.AssertNil(err)
		t.Assert(count, 1)
	})

	gtest.C(t, func(t *gtest.T) {
		r, err := shardingModel.Where(g.Map{
			"id": 1,
		}).Delete()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)

		var count int
		count, err = shardingModel.Count()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db.Model(table).Count()
		t.AssertNil(err)
		t.Assert(count, 0)

		count, err = db2.Model(table).Count()
		t.AssertNil(err)
		t.Assert(count, 0)
	})
}
